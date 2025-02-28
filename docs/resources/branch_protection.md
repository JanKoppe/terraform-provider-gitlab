---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "gitlab_branch_protection Resource - terraform-provider-gitlab"
subcategory: ""
description: |-
  The gitlab_branch_protection resource allows to manage the lifecycle of a protected branch of a repository.
  ~> The allowed_to_push, allowed_to_merge, allowed_to_unprotect, unprotect_access_level and code_owner_approval_required attributes require a GitLab Enterprise instance.
  Upstream API: GitLab REST API docs https://docs.gitlab.com/ee/api/protected_branches.html
---

# gitlab_branch_protection (Resource)

The `gitlab_branch_protection` resource allows to manage the lifecycle of a protected branch of a repository.

~> The `allowed_to_push`, `allowed_to_merge`, `allowed_to_unprotect`, `unprotect_access_level` and `code_owner_approval_required` attributes require a GitLab Enterprise instance.

**Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/protected_branches.html)

## Example Usage

```terraform
resource "gitlab_branch_protection" "BranchProtect" {
  project                      = "12345"
  branch                       = "BranchProtected"
  push_access_level            = "developer"
  merge_access_level           = "developer"
  unprotect_access_level       = "developer"
  allow_force_push             = true
  code_owner_approval_required = true
  allowed_to_push {
    user_id = 5
  }
  allowed_to_push {
    user_id = 521
  }
  allowed_to_merge {
    user_id = 15
  }
  allowed_to_merge {
    user_id = 37
  }
  allowed_to_unprotect {
    user_id = 15
  }
  allowed_to_unprotect {
    group_id = 42
  }
}

# Example using dynamic block
resource "gitlab_branch_protection" "main" {
  project                = "12345"
  branch                 = "main"
  push_access_level      = "maintainer"
  merge_access_level     = "maintainer"
  unprotect_access_level = "maintainer"

  dynamic "allowed_to_push" {
    for_each = [50, 55, 60]
    content {
      user_id = allowed_to_push.value
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **branch** (String) Name of the branch.
- **project** (String) The id of the project.

### Optional

- **allow_force_push** (Boolean) Can be set to true to allow users with push access to force push.
- **allowed_to_merge** (Block Set) Defines permissions for action. (see [below for nested schema](#nestedblock--allowed_to_merge))
- **allowed_to_push** (Block Set) Defines permissions for action. (see [below for nested schema](#nestedblock--allowed_to_push))
- **allowed_to_unprotect** (Block Set) Defines permissions for action. (see [below for nested schema](#nestedblock--allowed_to_unprotect))
- **code_owner_approval_required** (Boolean) Can be set to true to require code owner approval before merging.
- **id** (String) The ID of this resource.
- **merge_access_level** (String) Access levels allowed to merge. Valid values are: `no one`, `developer`, `maintainer`.
- **push_access_level** (String) Access levels allowed to push. Valid values are: `no one`, `developer`, `maintainer`.
- **unprotect_access_level** (String) Access levels allowed to unprotect. Valid values are: `developer`, `maintainer`.

### Read-Only

- **branch_protection_id** (Number) The ID of the branch protection (not the branch name).

<a id="nestedblock--allowed_to_merge"></a>
### Nested Schema for `allowed_to_merge`

Optional:

- **group_id** (Number) The ID of a GitLab group allowed to perform the relevant action. Mutually exclusive with `user_id`.
- **user_id** (Number) The ID of a GitLab user allowed to perform the relevant action. Mutually exclusive with `group_id`.

Read-Only:

- **access_level** (String) Level of access.
- **access_level_description** (String) Readable description of level of access.


<a id="nestedblock--allowed_to_push"></a>
### Nested Schema for `allowed_to_push`

Optional:

- **group_id** (Number) The ID of a GitLab group allowed to perform the relevant action. Mutually exclusive with `user_id`.
- **user_id** (Number) The ID of a GitLab user allowed to perform the relevant action. Mutually exclusive with `group_id`.

Read-Only:

- **access_level** (String) Level of access.
- **access_level_description** (String) Readable description of level of access.


<a id="nestedblock--allowed_to_unprotect"></a>
### Nested Schema for `allowed_to_unprotect`

Optional:

- **group_id** (Number) The ID of a GitLab group allowed to perform the relevant action. Mutually exclusive with `user_id`.
- **user_id** (Number) The ID of a GitLab user allowed to perform the relevant action. Mutually exclusive with `group_id`.

Read-Only:

- **access_level** (String) Level of access.
- **access_level_description** (String) Readable description of level of access.

## Import

Import is supported using the following syntax:

```shell
# Gitlab protected branches can be imported with a key composed of `<project_id>:<branch>`, e.g.
terraform import gitlab_branch_protection.BranchProtect "12345:main"
```
