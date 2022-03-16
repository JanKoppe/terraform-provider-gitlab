package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceGitlabInstanceSettings_basic(t *testing.T) {
	testAccCheck(t)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					data "gitlab_instance_settings" "this" {}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.gitlab_instance_settings.this", "id"),
					resource.TestCheckResourceAttr("data.gitlab_instance_settings.this", "admin_mode", "false"),
				),
			},
		},
	})
}
