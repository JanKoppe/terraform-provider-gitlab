package provider

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/xanzy/go-gitlab"
)

func TestAccGitlabDeployToken_basic(t *testing.T) {
	var deployToken gitlab.DeployToken
	rInt := acctest.RandInt()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckGitlabDeployTokenDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGitlabDeployTokenConfig(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGitlabDeployTokenExists("gitlab_deploy_token.foo", &deployToken),
					testAccCheckGitlabDeployTokenAttributes(&deployToken, &testAccCheckGitlabDeployTokenExpectedAttributes{
						Name:     fmt.Sprintf("deployToken-%d", rInt),
						Username: "my-username",
					}),
				),
			},
		},
	})
}
func TestAccGitlabDeployToken_pagination(t *testing.T) {
	testAccCheck(t)

	testGroup := testAccCreateGroups(t, 1)[0]
	testProject := testAccCreateProject(t)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckGitlabDeployTokenDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGitlabDeployTokenPaginationConfig(25, testGroup.ID, testProject.ID),
			},
			// In case pagination wouldn't properly work, we would get that the plan isn't empty,
			// because some of the deploy tokens wouldn't be in the first page and therefore
			// considered non-existing, ...
			{
				Config:   testAccGitlabDeployTokenPaginationConfig(25, testGroup.ID, testProject.ID),
				PlanOnly: true,
			},
		},
	})
}

func testAccCheckGitlabDeployTokenExists(n string, deployToken *gitlab.DeployToken) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not Found: %s", n)
		}

		deployTokenID, err := strconv.Atoi(rs.Primary.ID)
		if err != nil {
			return err
		}

		projectName := rs.Primary.Attributes["project"]
		groupName := rs.Primary.Attributes["group"]

		var gotDeployTokens []*gitlab.DeployToken

		if projectName != "" {
			gotDeployTokens, _, err = testGitlabClient.DeployTokens.ListProjectDeployTokens(projectName, nil)
		} else if groupName != "" {
			gotDeployTokens, _, err = testGitlabClient.DeployTokens.ListGroupDeployTokens(groupName, nil)
		} else {
			return fmt.Errorf("No project or group ID is set")
		}

		if err != nil {
			return err
		}

		for _, token := range gotDeployTokens {
			if token.ID == deployTokenID {
				*deployToken = *token
				return nil
			}
		}

		return fmt.Errorf("Deploy Token doesn't exist")
	}
}

type testAccCheckGitlabDeployTokenExpectedAttributes struct {
	Name     string
	Username string
}

func testAccCheckGitlabDeployTokenAttributes(deployToken *gitlab.DeployToken, want *testAccCheckGitlabDeployTokenExpectedAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if deployToken.Name != want.Name {
			return fmt.Errorf("got name %q; want %q", deployToken.Name, want.Name)
		}

		if deployToken.Username != want.Username {
			return fmt.Errorf("got username %q; want %q", deployToken.Username, want.Username)
		}

		return nil
	}
}

func testAccCheckGitlabDeployTokenDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "gitlab_deploy_token" {
			continue
		}

		deployTokenID, err := strconv.Atoi(rs.Primary.ID)
		if err != nil {
			return err
		}

		project := rs.Primary.Attributes["project"]
		group := rs.Primary.Attributes["group"]

		var gotDeployTokens []*gitlab.DeployToken

		if project != "" {
			gotDeployTokens, _, err = testGitlabClient.DeployTokens.ListProjectDeployTokens(project, nil)
		} else if group != "" {
			gotDeployTokens, _, err = testGitlabClient.DeployTokens.ListGroupDeployTokens(group, nil)
		} else {
			return fmt.Errorf("somehow neither project nor group were set")
		}

		if err == nil {
			for _, token := range gotDeployTokens {
				if token.ID == deployTokenID {
					return fmt.Errorf("Deploy token still exists")
				}
			}
		}

		if !is404(err) {
			return err
		}
	}

	return nil
}

func testAccGitlabDeployTokenConfig(rInt int) string {
	return fmt.Sprintf(`
resource "gitlab_project" "foo" {
  name        = "foo-%d"
  description = "Terraform acceptance test"

  # So that acceptance tests can be run in a gitlab organization
  # with no billing
  visibility_level = "public"
}

resource "gitlab_deploy_token" "foo" {
  project  = "${gitlab_project.foo.id}"
  name     = "deployToken-%d"
  username = "my-username"

  expires_at = "2021-03-14T07:20:50.000Z"

  scopes = [
	"read_registry",
	"read_repository",
	"read_package_registry",
	"write_registry",
	"write_package_registry",
  ]
}
  `, rInt, rInt)
}

func testAccGitlabDeployTokenPaginationConfig(numberOfTokens int, groupID int, projectID int) string {
	return fmt.Sprintf(`
resource "gitlab_deploy_token" "example_group" {
  group  = %d
  name   = "deploy-token-${count.index}"
  scopes = ["read_registry"]

  count = %d
}

resource "gitlab_deploy_token" "example_project" {
  project  = %d
  name   = "deploy-token-${count.index}"
  scopes = ["read_registry"]

  count = %d
}
  `, groupID, numberOfTokens, projectID, numberOfTokens)
}

type expiresAtSuppressFuncTest struct {
	description string
	old         string
	new         string
	expected    bool
}

func TestExpiresAtSuppressFunc(t *testing.T) {
	testcases := []expiresAtSuppressFuncTest{
		{
			description: "same dates without millis",
			old:         "2025-03-14T00:00:00Z",
			new:         "2025-03-14T00:00:00Z",
			expected:    true,
		}, {
			description: "different date without millis",
			old:         "2025-03-14T00:00:00Z",
			new:         "2025-03-14T11:11:11Z",
			expected:    false,
		}, {
			description: "same date with and without millis",
			old:         "2025-03-14T00:00:00Z",
			new:         "2025-03-14T00:00:00.000Z",
			expected:    true,
		}, {
			description: "cannot parse new date",
			old:         "2025-03-14T00:00:00Z",
			new:         "invalid-date",
			expected:    false,
		},
	}

	for _, test := range testcases {
		t.Run(test.description, func(t *testing.T) {
			actual := expiresAtSuppressFunc("", test.old, test.new, nil)
			if actual != test.expected {
				t.Fatalf("FAIL\n\told: %s, new: %s\n\texpected: %t\n\tactual: %t",
					test.old, test.new, test.expected, actual)
			}
		})
	}
}
