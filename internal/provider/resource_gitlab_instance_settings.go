package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	gitlab "github.com/xanzy/go-gitlab"
)

var _ = registerResource("gitlab_instance_settings", func() *schema.Resource {
	return &schema.Resource{
		Description: `The ` + "`gitlab_instance_settings`" + ` data source allows to manage the instance application settings.

**Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/settings.html)`,

		CreateContext: resourceGitlabInstanceSettingsCreateOrUpdate, // there's no real "create", just updating. keep it DRY
		ReadContext:   resourceGitlabInstanceSettingsRead,
		UpdateContext: resourceGitlabInstanceSettingsCreateOrUpdate,
		DeleteContext: resourceGitlabInstanceSettingsDelete,

		Schema: map[string]*schema.Schema{
			"admin_mode": {
				Description: "Require administrators to enable Admin Mode by re-authenticating for administrative tasks.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"after_sign_out_path": {
				Description: "Where to redirect users after logout.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			//"restricted_visibility_levels": {
			//	Description: "Selected levels cannot be used by non-Administrator users for groups, projects or snippets. Can take 'private', 'internal' and 'public' as a parameter. Default is 'null' which means there is no restriction.",
			//	Type:        schema.TypeList,
			//	Elem: &schema.Schema{
			//		Type:         schema.TypeString,
			//		ValidateFunc: validation.StringInSlice([]string{"private", "internal", "public"}, false),
			//	},
			//	Computed: true,
			//	Optional: true,
			//},
		},
	}
})

func resourceGitlabInstanceSettingsCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*gitlab.Client)

	// TODO: Only works properly with GetOkExists for some reason. Is this an issue?
	options := &gitlab.UpdateSettingsOptions{}
	if v, ok := d.GetOkExists("admin_mode"); ok {
		options.AdminMode = gitlab.Bool(v.(bool))
	}

	if v, ok := d.GetOk("after_sign_out_path"); ok {
		options.AfterSignOutPath = gitlab.String(v.(string))
	}

	// TODO: How to handle "empty" lists? This should then set the call to "null" apparently? unclear.
	//if v, ok := d.GetOk("restricted_visibility_levels"); ok {
	//	itemsRaw := v.([]interface{})
	//	items := make([]gitlab.VisibilityValue, len(itemsRaw))
	//	for i, raw := range itemsRaw {
	//		items[i] = *stringToVisibilityLevel(raw.(string))
	//	}
	//	options.RestrictedVisibilityLevels = &items
	//}

	log.Printf("[DEBUG] create gitlab instance level settings object")

	_, _, err := client.Settings.UpdateSettings(options, gitlab.WithContext(ctx))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(client.BaseURL().String())

	return resourceGitlabInstanceSettingsRead(ctx, d, meta)
}

func resourceGitlabInstanceSettingsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*gitlab.Client)
	log.Printf("[DEBUG] read gitlab instance level settings")

	// NOTE: As the settings are global per instance, we do not request with any ID.
	//       Also, settings cannot be removed in the remote, so there's no logic to
	//       unset the ID in that case.
	settings, _, err := client.Settings.GetSettings(gitlab.WithContext(ctx))
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("admin_mode", settings.AdminMode)
	d.Set("after_sign_out_path", settings.AfterSignOutPath)
	//d.Set("restricted_visibility_levels", settings.RestrictedVisibilityLevels)

	return nil
}

func resourceGitlabInstanceSettingsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// NOTE: It is not possible to "delete" the settings, as they are global to the gitlab instance.
	//       We could in theory reset everything to the default values as described in the API documentation.
	//       TODO: Decide if we just lose the resource from state, or reset settings to default values.
	//client := meta.(*gitlab.Client)
	log.Printf("[DEBUG] Delete gitlab instance level settings [noop]")
	return nil
}
