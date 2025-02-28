---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "gitlab_project_protected_branch Data Source - terraform-provider-gitlab"
subcategory: ""
description: |-
  The gitlab_protected_branch data source allows details of a protected branch to be retrieved by its name and the project it belongs to.
  Upstream API: GitLab REST API docs https://docs.gitlab.com/ee/api/protected_branches.html#get-a-single-protected-branch-or-wildcard-protected-branch
---

# gitlab_project_protected_branch (Data Source)

The `gitlab_protected_branch` data source allows details of a protected branch to be retrieved by its name and the project it belongs to.

**Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/protected_branches.html#get-a-single-protected-branch-or-wildcard-protected-branch)

## Example Usage

```terraform
data "gitlab_project_protected_branch" "example" {
  project_id = 30
  name       = "main"
}

data "gitlab_project_protected_branch" "example" {
  project_id = "foo/bar/baz"
  name       = "main"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **name** (String) The name of the protected branch.
- **project_id** (String) The integer or path with namespace that uniquely identifies the project.

### Read-Only

- **allow_force_push** (Boolean) Whether force push is allowed.
- **code_owner_approval_required** (Boolean) Reject code pushes that change files listed in the CODEOWNERS file.
- **id** (Number) The ID of this resource.
- **merge_access_levels** (List of Object) Describes which access levels, users, or groups are allowed to perform the action. (see [below for nested schema](#nestedatt--merge_access_levels))
- **push_access_levels** (List of Object) Describes which access levels, users, or groups are allowed to perform the action. (see [below for nested schema](#nestedatt--push_access_levels))

<a id="nestedatt--merge_access_levels"></a>
### Nested Schema for `merge_access_levels`

Read-Only:

- **access_level** (String)
- **access_level_description** (String)
- **group_id** (Number)
- **user_id** (Number)


<a id="nestedatt--push_access_levels"></a>
### Nested Schema for `push_access_levels`

Read-Only:

- **access_level** (String)
- **access_level_description** (String)
- **group_id** (Number)
- **user_id** (Number)


