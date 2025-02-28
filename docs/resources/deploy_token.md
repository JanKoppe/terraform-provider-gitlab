---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "gitlab_deploy_token Resource - terraform-provider-gitlab"
subcategory: ""
description: |-
  The gitlab_deploy_token resource allows to manage the lifecycle of group and project deploy tokens.
  Upstream API: GitLab REST API docs https://docs.gitlab.com/ee/api/deploy_tokens.html
---

# gitlab_deploy_token (Resource)

The `gitlab_deploy_token` resource allows to manage the lifecycle of group and project deploy tokens.

**Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/deploy_tokens.html)

## Example Usage

```terraform
# Example Usage - Project
resource "gitlab_deploy_token" "example" {
  project    = "example/deploying"
  name       = "Example deploy token"
  username   = "example-username"
  expires_at = "2020-03-14T00:00:00.000Z"

  scopes = ["read_repository", "read_registry"]
}

resource "gitlab_deploy_token" "example-two" {
  project    = "12345678"
  name       = "Example deploy token expires in 24h"
  expires_at = timeadd(timestamp(), "24h")
}

# Example Usage - Group
resource "gitlab_deploy_token" "example" {
  group = "example/deploying"
  name  = "Example group deploy token"

  scopes = ["read_repository"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **name** (String) A name to describe the deploy token with.
- **scopes** (Set of String) Valid values: `read_repository`, `read_registry`, `read_package_registry`, `write_registry`, `write_package_registry`.

### Optional

- **expires_at** (String) Time the token will expire it, RFC3339 format. Will not expire per default.
- **group** (String) The name or id of the group to add the deploy token to.
- **id** (String) The ID of this resource.
- **project** (String) The name or id of the project to add the deploy token to.
- **username** (String) A username for the deploy token. Default is `gitlab+deploy-token-{n}`.

### Read-Only

- **token** (String, Sensitive) The secret token. This is only populated when creating a new deploy token.


