---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "gitlab_projects Data Source - terraform-provider-gitlab"
subcategory: ""
description: |-
  The gitlab_projects data source allows details of multiple projects to be retrieved. Optionally filtered by the set attributes.
  -> This data source supports all available filters exposed by the xanzy/go-gitlab package, which might not expose all available filters exposed by the Gitlab APIs.
  -> The owner sub-attributes are only populated if the Gitlab token used has an administrator scope.
  Upstream API: GitLab REST API docs https://docs.gitlab.com/ee/api/projects.html#list-all-projects
---

# gitlab_projects (Data Source)

The `gitlab_projects` data source allows details of multiple projects to be retrieved. Optionally filtered by the set attributes.

-> This data source supports all available filters exposed by the xanzy/go-gitlab package, which might not expose all available filters exposed by the Gitlab APIs.

-> The [owner sub-attributes](#nestedobjatt--projects--owner) are only populated if the Gitlab token used has an administrator scope.

**Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/projects.html#list-all-projects)

## Example Usage

```terraform
# List projects within a group tree
data "gitlab_group" "mygroup" {
  full_path = "mygroup"
}

data "gitlab_projects" "group_projects" {
  group_id          = data.gitlab_group.mygroup.id
  order_by          = "name"
  include_subgroups = true
  with_shared       = false
}

# List projects using the search syntax
data "gitlab_projects" "projects" {
  search     = "postgresql"
  visibility = "private"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **archived** (Boolean) Limit by archived status.
- **group_id** (Number) The ID of the group owned by the authenticated user to look projects for within. Cannot be used with `min_access_level`, `with_programming_language` or `statistics`.
- **id** (String) The ID of this resource.
- **include_subgroups** (Boolean) Include projects in subgroups of this group. Default is `false`. Needs `group_id`.
- **max_queryable_pages** (Number) The maximum number of project results pages that may be queried. Prevents overloading your Gitlab instance in case of a misconfiguration.
- **membership** (Boolean) Limit by projects that the current user is a member of.
- **min_access_level** (Number) Limit to projects where current user has at least this access level, refer to the [official documentation](https://docs.gitlab.com/ee/api/members.html) for values. Cannot be used with `group_id`.
- **order_by** (String) Return projects ordered by `id`, `name`, `path`, `created_at`, `updated_at`, or `last_activity_at` fields. Default is `created_at`.
- **owned** (Boolean) Limit by projects owned by the current user.
- **page** (Number) The first page to begin the query on.
- **per_page** (Number) The number of results to return per page.
- **search** (String) Return list of authorized projects matching the search criteria.
- **simple** (Boolean) Return only the ID, URL, name, and path of each project.
- **sort** (String) Return projects sorted in `asc` or `desc` order. Default is `desc`.
- **starred** (Boolean) Limit by projects starred by the current user.
- **statistics** (Boolean) Include project statistics. Cannot be used with `group_id`.
- **visibility** (String) Limit by visibility `public`, `internal`, or `private`.
- **with_custom_attributes** (Boolean) Include custom attributes in response _(admins only)_.
- **with_issues_enabled** (Boolean) Limit by projects with issues feature enabled. Default is `false`.
- **with_merge_requests_enabled** (Boolean) Limit by projects with merge requests feature enabled. Default is `false`.
- **with_programming_language** (String) Limit by projects which use the given programming language. Cannot be used with `group_id`.
- **with_shared** (Boolean) Include projects shared to this group. Default is `true`. Needs `group_id`.

### Read-Only

- **projects** (List of Object) A list containing the projects matching the supplied arguments (see [below for nested schema](#nestedatt--projects))

<a id="nestedatt--projects"></a>
### Nested Schema for `projects`

Read-Only:

- **_links** (Map of String)
- **allow_merge_on_skipped_pipeline** (Boolean)
- **approvals_before_merge** (Number)
- **archived** (Boolean)
- **avatar_url** (String)
- **build_coverage_regex** (String)
- **ci_config_path** (String)
- **ci_forward_deployment_enabled** (Boolean)
- **container_registry_enabled** (Boolean)
- **created_at** (String)
- **creator_id** (Number)
- **custom_attributes** (List of Map of String)
- **default_branch** (String)
- **description** (String)
- **forked_from_project** (List of Object) (see [below for nested schema](#nestedobjatt--projects--forked_from_project))
- **forks_count** (Number)
- **http_url_to_repo** (String)
- **id** (Number)
- **import_error** (String)
- **import_status** (String)
- **issues_enabled** (Boolean)
- **jobs_enabled** (Boolean)
- **last_activity_at** (String)
- **lfs_enabled** (Boolean)
- **merge_method** (String)
- **merge_pipelines_enabled** (Boolean)
- **merge_requests_enabled** (Boolean)
- **merge_trains_enabled** (Boolean)
- **mirror** (Boolean)
- **mirror_overwrites_diverged_branches** (Boolean)
- **mirror_trigger_builds** (Boolean)
- **mirror_user_id** (Number)
- **name** (String)
- **name_with_namespace** (String)
- **namespace** (List of Object) (see [below for nested schema](#nestedobjatt--projects--namespace))
- **only_allow_merge_if_all_discussions_are_resolved** (Boolean)
- **only_allow_merge_if_pipeline_succeeds** (Boolean)
- **only_mirror_protected_branches** (Boolean)
- **open_issues_count** (Number)
- **owner** (List of Object) (see [below for nested schema](#nestedobjatt--projects--owner))
- **packages_enabled** (Boolean)
- **path** (String)
- **path_with_namespace** (String)
- **permissions** (List of Object) (see [below for nested schema](#nestedobjatt--projects--permissions))
- **public** (Boolean)
- **public_builds** (Boolean)
- **readme_url** (String)
- **request_access_enabled** (Boolean)
- **resolve_outdated_diff_discussions** (Boolean)
- **runners_token** (String)
- **shared_runners_enabled** (Boolean)
- **shared_with_groups** (List of Object) (see [below for nested schema](#nestedobjatt--projects--shared_with_groups))
- **snippets_enabled** (Boolean)
- **ssh_url_to_repo** (String)
- **star_count** (Number)
- **statistics** (Map of Number)
- **tag_list** (Set of String)
- **visibility** (String)
- **web_url** (String)
- **wiki_enabled** (Boolean)

<a id="nestedobjatt--projects--forked_from_project"></a>
### Nested Schema for `projects.forked_from_project`

Read-Only:

- **http_url_to_repo** (String)
- **id** (Number)
- **name** (String)
- **name_with_namespace** (String)
- **path** (String)
- **path_with_namespace** (String)
- **web_url** (String)


<a id="nestedobjatt--projects--namespace"></a>
### Nested Schema for `projects.namespace`

Read-Only:

- **full_path** (String)
- **id** (Number)
- **kind** (String)
- **name** (String)
- **path** (String)


<a id="nestedobjatt--projects--owner"></a>
### Nested Schema for `projects.owner`

Read-Only:

- **avatar_url** (String)
- **id** (Number)
- **name** (String)
- **state** (String)
- **username** (String)
- **website_url** (String)


<a id="nestedobjatt--projects--permissions"></a>
### Nested Schema for `projects.permissions`

Read-Only:

- **group_access** (Map of Number)
- **project_access** (Map of Number)


<a id="nestedobjatt--projects--shared_with_groups"></a>
### Nested Schema for `projects.shared_with_groups`

Read-Only:

- **group_access_level** (String)
- **group_id** (Number)
- **group_name** (String)


