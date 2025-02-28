package provider

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/xanzy/go-gitlab"
)

var _ = registerResource("gitlab_deploy_token", func() *schema.Resource {
	// lintignore: XR002 // TODO: Resolve this tfproviderlint issue
	return &schema.Resource{
		Description: `The ` + "`gitlab_deploy_token`" + ` resource allows to manage the lifecycle of group and project deploy tokens.

**Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/deploy_tokens.html)`,

		CreateContext: resourceGitlabDeployTokenCreate,
		ReadContext:   resourceGitlabDeployTokenRead,
		DeleteContext: resourceGitlabDeployTokenDelete,

		Schema: map[string]*schema.Schema{
			"project": {
				Description:  "The name or id of the project to add the deploy token to.",
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"project", "group"},
				ForceNew:     true,
			},
			"group": {
				Description:  "The name or id of the group to add the deploy token to.",
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"project", "group"},
				ForceNew:     true,
			},
			"name": {
				Description: "A name to describe the deploy token with.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"username": {
				Description: "A username for the deploy token. Default is `gitlab+deploy-token-{n}`.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"expires_at": {
				Description:      "Time the token will expire it, RFC3339 format. Will not expire per default.",
				Type:             schema.TypeString,
				Optional:         true,
				ValidateFunc:     validation.IsRFC3339Time,
				DiffSuppressFunc: expiresAtSuppressFunc,
				ForceNew:         true,
			},
			"scopes": {
				Description: "Valid values: `read_repository`, `read_registry`, `read_package_registry`, `write_registry`, `write_package_registry`.",
				Type:        schema.TypeSet,
				Required:    true,
				ForceNew:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validation.StringInSlice(
						[]string{
							"read_registry",
							"read_repository",
							"read_package_registry",
							"write_registry",
							"write_package_registry",
						}, false),
				},
			},

			"token": {
				Description: "The secret token. This is only populated when creating a new deploy token.",
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
			},
		},
	}
})

func expiresAtSuppressFunc(k, old, new string, d *schema.ResourceData) bool {
	oldDate, oldDateErr := time.Parse(time.RFC3339, old)
	newDate, newDateErr := time.Parse(time.RFC3339, new)
	if oldDateErr != nil || newDateErr != nil {
		return false
	}
	return oldDate == newDate
}

func resourceGitlabDeployTokenCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*gitlab.Client)
	project, isProject := d.GetOk("project")
	group, isGroup := d.GetOk("group")

	var expiresAt *time.Time
	var err error

	if exp, ok := d.GetOk("expires_at"); ok {
		parsedExpiresAt, err := time.Parse(time.RFC3339, exp.(string))
		expiresAt = &parsedExpiresAt
		if err != nil {
			return diag.Errorf("Invalid expires_at date: %v", err)
		}
	}

	scopes := stringSetToStringSlice(d.Get("scopes").(*schema.Set))

	var deployToken *gitlab.DeployToken

	if isProject {
		options := &gitlab.CreateProjectDeployTokenOptions{
			Name:      gitlab.String(d.Get("name").(string)),
			Username:  gitlab.String(d.Get("username").(string)),
			ExpiresAt: expiresAt,
			Scopes:    scopes,
		}

		log.Printf("[DEBUG] Create GitLab deploy token %s in project %s", *options.Name, project.(string))

		deployToken, _, err = client.DeployTokens.CreateProjectDeployToken(project, options, gitlab.WithContext(ctx))

	} else if isGroup {
		options := &gitlab.CreateGroupDeployTokenOptions{
			Name:      gitlab.String(d.Get("name").(string)),
			Username:  gitlab.String(d.Get("username").(string)),
			ExpiresAt: expiresAt,
			Scopes:    scopes,
		}

		log.Printf("[DEBUG] Create GitLab deploy token %s in group %s", *options.Name, group.(string))

		deployToken, _, err = client.DeployTokens.CreateGroupDeployToken(group, options, gitlab.WithContext(ctx))
	}

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%d", deployToken.ID))

	// Token is only available on creation
	d.Set("token", deployToken.Token)
	d.Set("username", deployToken.Username)

	return nil
}

func resourceGitlabDeployTokenRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*gitlab.Client)
	project, isProject := d.GetOk("project")
	group, isGroup := d.GetOk("group")
	deployTokenID, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	var deployToken *gitlab.DeployToken

	extractDeployToken := func(paginatedDeployTokens []*gitlab.DeployToken) *gitlab.DeployToken {
		for _, token := range paginatedDeployTokens {
			if token.ID == deployTokenID {
				return token
			}
		}
		return nil
	}

	if isProject {
		log.Printf("[DEBUG] Read GitLab deploy token %d in project %s", deployTokenID, project.(string))
		options := gitlab.ListProjectDeployTokensOptions{
			Page:    1,
			PerPage: 20,
		}
		for options.Page != 0 && deployToken == nil {
			paginatedDeployTokens, resp, err := client.DeployTokens.ListProjectDeployTokens(project, &options, gitlab.WithContext(ctx))
			if err != nil {
				return diag.FromErr(err)
			}
			deployToken = extractDeployToken(paginatedDeployTokens)
			options.Page = resp.NextPage
		}
	} else if isGroup {
		log.Printf("[DEBUG] Read GitLab deploy token %d in group %s", deployTokenID, group.(string))
		options := gitlab.ListGroupDeployTokensOptions{
			Page:    1,
			PerPage: 20,
		}
		for options.Page != 0 && deployToken == nil {
			paginatedDeployTokens, resp, err := client.DeployTokens.ListGroupDeployTokens(group, &options, gitlab.WithContext(ctx))
			if err != nil {
				return diag.FromErr(err)
			}
			deployToken = extractDeployToken(paginatedDeployTokens)
			options.Page = resp.NextPage
		}
	}

	if deployToken == nil {
		log.Printf("[DEBUG] GitLab deploy token %d in was not found, removing from state", deployTokenID)
		d.SetId("")
		return nil
	}

	d.Set("name", deployToken.Name)
	d.Set("username", deployToken.Username)

	if deployToken.ExpiresAt != nil {
		d.Set("expires_at", deployToken.ExpiresAt.Format(time.RFC3339))
	}

	if err := d.Set("scopes", deployToken.Scopes); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGitlabDeployTokenDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*gitlab.Client)
	project, isProject := d.GetOk("project")
	group, isGroup := d.GetOk("group")
	deployTokenID, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	var response *gitlab.Response

	if isProject {
		log.Printf("[DEBUG] Delete GitLab deploy token %d in project %s", deployTokenID, project.(string))
		response, err = client.DeployTokens.DeleteProjectDeployToken(project, deployTokenID, gitlab.WithContext(ctx))

	} else if isGroup {
		log.Printf("[DEBUG] Delete GitLab deploy token %d in group %s", deployTokenID, group.(string))
		response, err = client.DeployTokens.DeleteGroupDeployToken(group, deployTokenID, gitlab.WithContext(ctx))
	}
	if err != nil {
		return diag.FromErr(err)
	}

	// StatusNoContent = 204
	// Success with no body
	if response.StatusCode != http.StatusNoContent {
		return diag.Errorf("Invalid status code returned: %s", response.Status)
	}

	return nil
}
