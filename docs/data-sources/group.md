---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "gitlab_group Data Source - terraform-provider-gitlab"
subcategory: ""
description: |-
  The gitlab_group data source allows details of a group to be retrieved by its id or full path.
  Upstream API: GitLab REST API docs https://docs.gitlab.com/ee/api/groups.html#details-of-a-group
---

# gitlab_group (Data Source)

The `gitlab_group` data source allows details of a group to be retrieved by its id or full path.

**Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/groups.html#details-of-a-group)

## Example Usage

```terraform
# By group's ID
data "gitlab_group" "foo" {
  group_id = 123
}

# By group's full path
data "gitlab_group" "foo" {
  full_path = "foo/bar"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **full_path** (String) The full path of the group.
- **group_id** (Number) The ID of the group.
- **id** (String) The ID of this resource.

### Read-Only

- **default_branch_protection** (Number) Whether developers and maintainers can push to the applicable default branch.
- **description** (String) The description of the group.
- **full_name** (String) The full name of the group.
- **lfs_enabled** (Boolean) Boolean, is LFS enabled for projects in this group.
- **name** (String) The name of this group.
- **parent_id** (Number) Integer, ID of the parent group.
- **path** (String) The path of the group.
- **prevent_forking_outside_group** (Boolean) When enabled, users can not fork projects from this group to external namespaces.
- **request_access_enabled** (Boolean) Boolean, is request for access enabled to the group.
- **runners_token** (String, Sensitive) The group level registration token to use during runner setup.
- **visibility_level** (String) Visibility level of the group. Possible values are `private`, `internal`, `public`.
- **web_url** (String) Web URL of the group.


