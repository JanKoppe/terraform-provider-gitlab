package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGitlabInstanceSettings_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "gitlab_instance_settings" "this" {
						admin_mode = true
						after_sign_out_path = "https://registry.terraform.io/providers/gitlabhq/gitlab"
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("gitlab_instance_settings.this", "id"),
					resource.TestCheckResourceAttr("gitlab_instance_settings.this", "admin_mode", "true"),
					resource.TestCheckResourceAttr("gitlab_instance_settings.this", "after_sign_out_path", "https://registry.terraform.io/providers/gitlabhq/gitlab"),
				),
			},
			{
				Config: `
					resource "gitlab_instance_settings" "this" {
						admin_mode = false
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("gitlab_instance_settings.this", "id"),
					resource.TestCheckResourceAttr("gitlab_instance_settings.this", "admin_mode", "false"),
				),
			},
		},
	})
}
