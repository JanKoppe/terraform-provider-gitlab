package provider

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	gitlab "github.com/xanzy/go-gitlab"
)

var _ = registerResource("gitlab_project_variable", func() *schema.Resource {
	return &schema.Resource{
		Description: `The ` + "`" + `gitlab_project_variable` + "`" + ` resource allows to manage the lifecycle of a CI/CD variable for a project.

~> **Important:** If your GitLab version is older than 13.4, you may see nondeterministic behavior when updating or deleting gitlab_project_variable resources with non-unique keys, for example if there is another variable with the same key and different environment scope. See [this GitLab issue](https://gitlab.com/gitlab-org/gitlab/-/issues/9912).

**Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/project_level_variables.html)`,

		CreateContext: resourceGitlabProjectVariableCreate,
		ReadContext:   resourceGitlabProjectVariableRead,
		UpdateContext: resourceGitlabProjectVariableUpdate,
		DeleteContext: resourceGitlabProjectVariableDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"project": {
				Description: "The name or id of the project.",
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
			},
			"key": {
				Description:  "The name of the variable.",
				Type:         schema.TypeString,
				ForceNew:     true,
				Required:     true,
				ValidateFunc: StringIsGitlabVariableName,
			},
			"value": {
				Description: "The value of the variable.",
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
			},
			"variable_type": {
				Description:  "The type of a variable. Available types are: env_var (default) and file.",
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "env_var",
				ValidateFunc: StringIsGitlabVariableType,
			},
			"protected": {
				Description: "If set to `true`, the variable will be passed only to pipelines running on protected branches and tags. Defaults to `false`.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"masked": {
				Description: "If set to `true`, the variable will be masked if it would have been written to the logs. Defaults to `false`.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"environment_scope": {
				Description: "The environment_scope of the variable. Defaults to `*`.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "*",
				// Versions of GitLab prior to 13.4 cannot update environment_scope.
				ForceNew: true,
			},
		},
	}
})

func resourceGitlabProjectVariableCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*gitlab.Client)

	project := d.Get("project").(string)
	key := d.Get("key").(string)
	value := d.Get("value").(string)
	variableType := stringToVariableType(d.Get("variable_type").(string))
	protected := d.Get("protected").(bool)
	masked := d.Get("masked").(bool)
	environmentScope := d.Get("environment_scope").(string)

	options := gitlab.CreateProjectVariableOptions{
		Key:              &key,
		Value:            &value,
		VariableType:     variableType,
		Protected:        &protected,
		Masked:           &masked,
		EnvironmentScope: &environmentScope,
	}

	id := strings.Join([]string{project, key, environmentScope}, ":")

	log.Printf("[DEBUG] create gitlab project variable %q", id)

	_, _, err := client.ProjectVariables.CreateVariable(project, &options, gitlab.WithContext(ctx))
	if err != nil {
		return augmentVariableClientError(d, err)
	}

	d.SetId(id)

	return resourceGitlabProjectVariableRead(ctx, d, meta)
}

func resourceGitlabProjectVariableRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*gitlab.Client)

	var (
		project          string
		key              string
		environmentScope string
	)

	// An older version of this resource used the ID format "project:key".
	// For backwards compatibility we still support the old format.
	parts := strings.SplitN(d.Id(), ":", 4)
	switch len(parts) {
	case 2:
		project = parts[0]
		key = parts[1]
		environmentScope = d.Get("environment_scope").(string)
	case 3:
		project = parts[0]
		key = parts[1]
		environmentScope = parts[2]
	default:
		return diag.Errorf(`Failed to parse project variable ID %q: expected format project:key or project:key:environment_scope`, d.Id())
	}

	log.Printf("[DEBUG] read gitlab project variable %q", d.Id())

	v, err := getProjectVariable(ctx, client, project, key, environmentScope)
	if err != nil {
		if errors.Is(err, errProjectVariableNotExist) {
			log.Printf("[DEBUG] read gitlab project variable %q was not found", d.Id())
			d.SetId("")
			return nil
		}
		return augmentVariableClientError(d, err)
	}

	d.Set("key", v.Key)
	d.Set("value", v.Value)
	d.Set("variable_type", v.VariableType)
	d.Set("project", project)
	d.Set("protected", v.Protected)
	d.Set("masked", v.Masked)
	d.Set("environment_scope", v.EnvironmentScope)
	return nil
}

func resourceGitlabProjectVariableUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*gitlab.Client)

	project := d.Get("project").(string)
	key := d.Get("key").(string)
	value := d.Get("value").(string)
	variableType := stringToVariableType(d.Get("variable_type").(string))
	protected := d.Get("protected").(bool)
	masked := d.Get("masked").(bool)
	environmentScope := d.Get("environment_scope").(string)

	options := &gitlab.UpdateProjectVariableOptions{
		Value:            &value,
		VariableType:     variableType,
		Protected:        &protected,
		Masked:           &masked,
		EnvironmentScope: &environmentScope,
	}
	log.Printf("[DEBUG] update gitlab project variable %q", d.Id())

	_, _, err := client.ProjectVariables.UpdateVariable(project, key, options, withEnvironmentScopeFilter(ctx, environmentScope))
	if err != nil {
		return augmentVariableClientError(d, err)
	}

	return resourceGitlabProjectVariableRead(ctx, d, meta)
}

func resourceGitlabProjectVariableDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*gitlab.Client)
	project := d.Get("project").(string)
	key := d.Get("key").(string)
	environmentScope := d.Get("environment_scope").(string)
	log.Printf("[DEBUG] Delete gitlab project variable %q", d.Id())

	// Note that the environment_scope filter is added here to support GitLab versions >= 13.4,
	// but it will be ignored in prior versions, causing nondeterministic destroy behavior when
	// destroying or updating scoped variables.
	// ref: https://gitlab.com/gitlab-org/gitlab/-/merge_requests/39209
	_, err := client.ProjectVariables.RemoveVariable(project, key, nil, withEnvironmentScopeFilter(ctx, environmentScope))
	return augmentVariableClientError(d, err)
}

var errProjectVariableNotExist = errors.New("project variable does not exist")

func getProjectVariable(ctx context.Context, client *gitlab.Client, project interface{}, key, environmentScope string) (*gitlab.ProjectVariable, error) {
	// List and filter variables manually to support GitLab versions < v13.4 (2020-08-22)
	// ref: https://gitlab.com/gitlab-org/gitlab/-/merge_requests/39209

	page := 1

	for {
		projectVariables, resp, err := client.ProjectVariables.ListVariables(project, &gitlab.ListProjectVariablesOptions{Page: page}, gitlab.WithContext(ctx))
		if err != nil {
			return nil, err
		}

		for _, v := range projectVariables {
			if v.Key == key && v.EnvironmentScope == environmentScope {
				return v, nil
			}
		}

		if resp.NextPage == 0 {
			return nil, errProjectVariableNotExist
		}

		page++
	}
}
