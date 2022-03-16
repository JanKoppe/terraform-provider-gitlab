package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	gitlab "github.com/xanzy/go-gitlab"
)

var _ = registerDataSource("gitlab_instance_settings", func() *schema.Resource {
	return &schema.Resource{
		Description: `The ` + "`gitlab_instance_settings`" + ` data source allows retrieving the instances application settings.

**Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/settings.html)`,

		ReadContext: dataSourceGitlabInstanceSettingsRead,

		Schema: map[string]*schema.Schema{
			"admin_mode": {
				Description: "Require administrators to enable Admin Mode by re-authenticating for administrative tasks.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			//"abuse_notification_email": {
			//	Description: "If set, [abuse reports](../user/admin_area/review_abuse_reports.md) are sent to this address. Abuse reports are always available in the Admin Area.",
			//	Type:        schema.TypeString,
			//	Computed:    true,
			//	Optional:    true,
			//},
			"after_sign_out_path": {
				Description: "Where to redirect users after logout.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"after_sign_up_text": {
				Description: "Text shown to the user after signing up.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"akismet_api_key": {
				Description: "API key for Akismet spam protection.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"akismet_enabled": {
				Description: "(**If enabled, requires:** 'akismet_api_key') Enable or disable Akismet spam protection.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"allow_group_owners_to_manage_ldap": {
				Description: "Set to 'true' to allow group owners to manage LDAP.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"allow_local_requests_from_system_hooks": {
				Description: "Allow requests to the local network from system hooks.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"allow_local_requests_from_web_hooks_and_services": {
				Description: "Allow requests to the local network from web hooks and services.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"archive_builds_in_human_readable": {
				Description: "Set the duration for which the jobs are considered as old and expired. After that time passes, the jobs are archived and no longer able to be retried. Make it empty to never expire jobs. It has to be no less than 1 day, for example: <code>15 days</code>, <code>1 month</code>, <code>2 years</code>.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"asset_proxy_enabled": {
				Description: "(**If enabled, requires:** 'asset_proxy_url') Enable proxying of assets. GitLab restart is required to apply changes.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"asset_proxy_secret_key": {
				Description: "Shared secret with the asset proxy server. GitLab restart is required to apply changes.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"asset_proxy_url": {
				Description: "URL of the asset proxy server. GitLab restart is required to apply changes.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"authorized_keys_enabled": {
				Description: "By default, we write to the 'authorized_keys' file to support Git over SSH without additional configuration. GitLab can be optimized to authenticate SSH keys via the database file. Only disable this if you have configured your OpenSSH server to use the AuthorizedKeysCommand.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"auto_devops_domain": {
				Description: "Specify a domain to use by default for every project's Auto Review Apps and Auto Deploy stages.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"auto_devops_enabled": {
				Description: "Enable Auto DevOps for projects by default. It automatically builds, tests, and deploys applications based on a predefined CI/CD configuration.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			//"automatic_purchased_storage_allocation": {
			//	Description: "Enabling this permits automatic allocation of purchased storage in a namespace.",
			//	Type:        schema.TypeBool,
			//	Computed:    true,
			//	Optional:    true,
			//},
			"check_namespace_plan": {
				Description: "Enabling this makes only licensed EE features available to projects if the project namespace's plan includes the feature or if the project is public.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"commit_email_hostname": {
				Description: "Custom hostname (for private commit emails).",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			//"container_expiration_policies_enable_historic_entries": {
			//	Description: "Enable [cleanup policies](../user/packages/container_registry/reduce_container_registry_storage.md#enable-the-cleanup-policy) for all projects.",
			//	Type:        schema.TypeBool,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"container_registry_cleanup_tags_service_max_list_size": {
			//	Description: "The maximum number of tags that can be deleted in a single execution of [cleanup policies](../user/packages/container_registry/reduce_container_registry_storage.md#set-cleanup-limits-to-conserve-resources).",
			//	Type:        schema.TypeInt,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"container_registry_delete_tags_service_timeout": {
			//	Description: "The maximum time, in seconds, that the cleanup process can take to delete a batch of tags for [cleanup policies](../user/packages/container_registry/reduce_container_registry_storage.md#set-cleanup-limits-to-conserve-resources).",
			//	Type:        schema.TypeInt,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"container_registry_expiration_policies_caching": {
			//	Description: "Caching during the execution of [cleanup policies](../user/packages/container_registry/reduce_container_registry_storage.md#set-cleanup-limits-to-conserve-resources).",
			//	Type:        schema.TypeBool,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"container_registry_expiration_policies_worker_capacity": {
			//	Description: "Number of workers for [cleanup policies](../user/packages/container_registry/reduce_container_registry_storage.md#set-cleanup-limits-to-conserve-resources).",
			//	Type:        schema.TypeInt,
			//	Computed:    true,
			//	Optional:    true,
			//},
			"container_registry_token_expire_delay": {
				Description: "Container Registry token duration in minutes.",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			//"deactivate_dormant_users": {
			//	Description: "Enable [automatic deactivation of dormant users](../user/admin_area/moderate_users.md#automatically-deactivate-dormant-users).",
			//	Type:        schema.TypeBool,
			//	Computed:    true,
			//	Optional:    true,
			//},
			"default_artifacts_expire_in": {
				Description: "Set the default expiration time for each job's artifacts.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			//"default_branch_name": {
			//	Description: "[Instance-level custom initial branch name](../user/project/repository/branches/default.md#instance-level-custom-initial-branch-name) ([introduced](https://gitlab.com/gitlab-org/gitlab/-/issues/225258) in GitLab 13.2).",
			//	Type:        schema.TypeString,
			//	Computed:    true,
			//	Optional:    true,
			//},
			"default_branch_protection": {
				Description: "Determine if developers can push to the default branch. Can take: '0' _(not protected, both users with the Developer role or Maintainer role can push new commits and force push)_, '1' _(partially protected, users with the Developer role or Maintainer role can push new commits, but cannot force push)_ or '2' _(fully protected, users with the Developer or Maintainer role cannot push new commits, but users with the Developer or Maintainer role can; no one can force push)_ as a parameter. Default is '2'.",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			//"default_ci_config_path": {
			//	Description: "Default CI/CD configuration file and path for new projects ('.gitlab-ci.yml' if not set).",
			//	Type:        schema.TypeString,
			//	Computed:    true,
			//	Optional:    true,
			//},
			"default_group_visibility": {
				Description: "What visibility level new groups receive. Can take 'private', 'internal' and 'public' as a parameter. Default is 'private'.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"default_project_creation": {
				Description: "Default project creation protection. Can take: '0' _(No one)_, '1' _(Maintainers)_ or '2' _(Developers + Maintainers)_",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			"default_project_visibility": {
				Description: "What visibility level new projects receive. Can take 'private', 'internal' and 'public' as a parameter. Default is 'private'.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"default_projects_limit": {
				Description: "Project limit per user. Default is '100000'.",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			"default_snippet_visibility": {
				Description: "What visibility level new snippets receive. Can take 'private', 'internal' and 'public' as a parameter. Default is 'private'.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			//"delayed_project_deletion": {
			//	Description: "Enable delayed project deletion by default in new groups. Default is 'false'.",
			//	Type:        schema.TypeBool,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"deletion_adjourned_period": {
			//	Description: "The number of days to wait before deleting a project or group that is marked for deletion. Value must be between 0 and 90.",
			//	Type:        schema.TypeInt,
			//	Computed:    true,
			//	Optional:    true,
			//},
			"diff_max_patch_bytes": {
				Description: "Maximum [diff patch size](../user/admin_area/diff_limits.md), in bytes.",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			//"diff_max_files": {
			//	Description: "Maximum [files in a diff](../user/admin_area/diff_limits.md).",
			//	Type:        schema.TypeInt,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"diff_max_lines": {
			//	Description: "Maximum [lines in a diff](../user/admin_area/diff_limits.md).",
			//	Type:        schema.TypeInt,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"disable_feed_token": {
			//	Description: "Disable display of RSS/Atom and calendar feed tokens ([introduced](https://gitlab.com/gitlab-org/gitlab/-/issues/231493) in GitLab 13.7)",
			//	Type:        schema.TypeBool,
			//	Computed:    true,
			//	Optional:    true,
			//},
			"dns_rebinding_protection_enabled": {
				Description: "Enforce DNS rebinding attack protection.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			//"domain_denylist_enabled": {
			//	Description: "(**If enabled, requires:** 'domain_denylist') Allows blocking sign-ups from emails from specific domains.",
			//	Type:        schema.TypeBool,
			//	Computed:    true,
			//	Optional:    true,
			//},
			"dsa_key_restriction": {
				Description: "The minimum allowed bit length of an uploaded DSA key. Default is '0' (no restriction). '-1' disables DSA keys.",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			"ecdsa_key_restriction": {
				Description: "The minimum allowed curve size (in bits) of an uploaded ECDSA key. Default is '0' (no restriction). '-1' disables ECDSA keys.",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			//"ecdsa_sk_key_restriction": {
			//	Description: "The minimum allowed curve size (in bits) of an uploaded ECDSA_SK key. Default is '0' (no restriction). '-1' disables ECDSA_SK keys.",
			//	Type:        schema.TypeInt,
			//	Computed:    true,
			//	Optional:    true,
			//},
			"ed25519_key_restriction": {
				Description: "The minimum allowed curve size (in bits) of an uploaded ED25519 key. Default is '0' (no restriction). '-1' disables ED25519 keys.",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			//"ed25519_sk_key_restriction": {
			//	Description: "The minimum allowed curve size (in bits) of an uploaded ED25519_SK key. Default is '0' (no restriction). '-1' disables ED25519_SK keys.",
			//	Type:        schema.TypeInt,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"eks_access_key_id": {
			//	Description: "AWS IAM access key ID.",
			//	Type:        schema.TypeString,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"eks_account_id": {
			//	Description: "Amazon account ID.",
			//	Type:        schema.TypeString,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"eks_integration_enabled": {
			//	Description: "Enable integration with Amazon EKS.",
			//	Type:        schema.TypeBool,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"eks_secret_access_key": {
			//	Description: "AWS IAM secret access key.",
			//	Type:        schema.TypeString,
			//	Computed:    true,
			//	Optional:    true,
			//},
			"elasticsearch_aws_access_key": {
				Description: "AWS IAM access key.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"elasticsearch_aws_region": {
				Description: "The AWS region the Elasticsearch domain is configured.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"elasticsearch_aws_secret_access_key": {
				Description: "AWS IAM secret access key.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"elasticsearch_aws": {
				Description: "Enable the use of AWS hosted Elasticsearch.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			//"elasticsearch_indexed_field_length_limit": {
			//	Description: "Maximum size of text fields to index by Elasticsearch. 0 value means no limit. This does not apply to repository and wiki indexing.",
			//	Type:        schema.TypeInt,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"elasticsearch_indexed_file_size_limit_kb": {
			//	Description: "Maximum size of repository and wiki files that are indexed by Elasticsearch.",
			//	Type:        schema.TypeInt,
			//	Computed:    true,
			//	Optional:    true,
			//},
			"elasticsearch_indexing": {
				Description: "Enable Elasticsearch indexing.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"elasticsearch_limit_indexing": {
				Description: "Limit Elasticsearch to index certain namespaces and projects.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			//"elasticsearch_max_bulk_concurrency": {
			//	Description: "Maximum concurrency of Elasticsearch bulk requests per indexing operation. This only applies to repository indexing operations.",
			//	Type:        schema.TypeInt,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"elasticsearch_max_bulk_size_mb": {
			//	Description: "Maximum size of Elasticsearch bulk indexing requests in MB. This only applies to repository indexing operations.",
			//	Type:        schema.TypeInt,
			//	Computed:    true,
			//	Optional:    true,
			//},
			"elasticsearch_search": {
				Description: "Enable Elasticsearch search.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"elasticsearch_url": {
				Description: "The URL to use for connecting to Elasticsearch. Use a comma-separated list to support cluster (for example, 'http://localhost:9200, http://localhost:9201').",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed: true,
				Optional: true,
			},
			//"elasticsearch_username": {
			//	Description: "The 'username' of your Elasticsearch instance.",
			//	Type:        schema.TypeString,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"elasticsearch_password": {
			//	Description: "The password of your Elasticsearch instance.",
			//	Type:        schema.TypeString,
			//	Computed:    true,
			//	Optional:    true,
			//},
			"email_additional_text": {
				Description: "Additional text added to the bottom of every email for legal/auditing/compliance reasons.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"email_author_in_body": {
				Description: "Some email servers do not support overriding the email sender name. Enable this option to include the name of the author of the issue, merge request or comment in the email body instead.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"enabled_git_access_protocol": {
				Description: "Enabled protocols for Git access. Allowed values are: 'ssh', 'http', and 'nil' to allow both protocols.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			//"enforce_namespace_storage_limit": {
			//	Description: "Enabling this permits enforcement of namespace storage limits.",
			//	Type:        schema.TypeBool,
			//	Computed:    true,
			//	Optional:    true,
			//},
			"enforce_terms": {
				Description: "(**If enabled, requires:** 'terms') Enforce application ToS to all users.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"external_auth_client_cert": {
				Description: "(**If enabled, requires:** 'external_auth_client_key') The certificate to use to authenticate with the external authorization service.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"external_auth_client_key_pass": {
				Description: "Passphrase to use for the private key when authenticating with the external service this is encrypted when stored.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"external_auth_client_key": {
				Description: "Private key for the certificate when authentication is required for the external authorization service, this is encrypted when stored.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"external_authorization_service_default_label": {
				Description: "The default classification label to use when requesting authorization and no classification label has been specified on the project.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"external_authorization_service_enabled": {
				Description: "(**If enabled, requires:** 'external_authorization_service_default_label', 'external_authorization_service_timeout' and 'external_authorization_service_url') Enable using an external authorization service for accessing projects.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"external_authorization_service_timeout": {
				Description: "The timeout after which an authorization request is aborted, in seconds. When a request times out, access is denied to the user. (min: 0.001, max: 10, step: 0.001).",
				Type:        schema.TypeFloat,
				Computed:    true,
				Optional:    true,
			},
			"external_authorization_service_url": {
				Description: "URL to which authorization requests are directed.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			//"external_pipeline_validation_service_url": {
			//	Description: "URL to use for pipeline validation requests.",
			//	Type:        schema.TypeString,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"external_pipeline_validation_service_token": {
			//	Description: "Optional. Token to include as the 'X-Gitlab-Token' header in requests to the URL in 'external_pipeline_validation_service_url'.",
			//	Type:        schema.TypeString,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"external_pipeline_validation_service_timeout": {
			//	Description: "How long to wait for a response from the pipeline validation service. Assumes 'OK' if it times out.",
			//	Type:        schema.TypeInt,
			//	Computed:    true,
			//	Optional:    true,
			//},
			"file_template_project_id": {
				Description: "The ID of a project to load custom file templates from.",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			"first_day_of_week": {
				Description: "Start day of the week for calendar views and date pickers. Valid values are '0' (default) for Sunday, '1' for Monday, and '6' for Saturday.",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			"geo_node_allowed_ips": {
				Description: "Comma-separated list of IPs and CIDRs of allowed secondary nodes. For example, '1.1.1.1, 2.2.2.0/24'.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"geo_status_timeout": {
				Description: "The amount of seconds after which a request to get a secondary node status times out.",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			//"git_two_factor_session_expiry": {
			//	Description: "Maximum duration (in minutes) of a session for Git operations when 2FA is enabled.",
			//	Type:        schema.TypeInt,
			//	Computed:    true,
			//	Optional:    true,
			//},
			"gitaly_timeout_default": {
				Description: "Default Gitaly timeout, in seconds. This timeout is not enforced for Git fetch/push operations or Sidekiq jobs. Set to '0' to disable timeouts.",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			"gitaly_timeout_fast": {
				Description: "Gitaly fast operation timeout, in seconds. Some Gitaly operations are expected to be fast. If they exceed this threshold, there may be a problem with a storage shard and 'failing fast' can help maintain the stability of the GitLab instance. Set to '0' to disable timeouts.",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			"gitaly_timeout_medium": {
				Description: "Medium Gitaly timeout, in seconds. This should be a value between the Fast and the Default timeout. Set to '0' to disable timeouts.",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			"grafana_enabled": {
				Description: "Enable Grafana.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"grafana_url": {
				Description: "Grafana URL.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"gravatar_enabled": {
				Description: "Enable Gravatar.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"hashed_storage_enabled": {
				Description: "Create new projects using hashed storage paths: Enable immutable, hash-based paths and repository names to store repositories on disk. This prevents repositories from having to be moved or renamed when the Project URL changes and may improve disk I/O performance. (Always enabled in GitLab versions 13.0 and later, configuration is scheduled for removal in 14.0)",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"help_page_hide_commercial_content": {
				Description: "Hide marketing-related entries from help.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"help_page_support_url": {
				Description: "Alternate support URL for help page and help dropdown.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"help_page_text": {
				Description: "Custom text displayed on the help page.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"help_text": {
				Description: "GitLab server administrator information.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"hide_third_party_offers": {
				Description: "Do not display offers from third parties in GitLab.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"home_page_url": {
				Description: "Redirect to this URL when not logged in.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"housekeeping_bitmaps_enabled": {
				Description: "Enable Git pack file bitmap creation.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"housekeeping_enabled": {
				Description: "(**If enabled, requires:** 'housekeeping_bitmaps_enabled', 'housekeeping_full_repack_period', 'housekeeping_gc_period', and 'housekeeping_incremental_repack_period') Enable or disable Git housekeeping.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"housekeeping_full_repack_period": {
				Description: "Number of Git pushes after which an incremental 'git repack' is run.",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			"housekeeping_gc_period": {
				Description: "Number of Git pushes after which 'git gc' is run.",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			"housekeeping_incremental_repack_period": {
				Description: "Number of Git pushes after which an incremental 'git repack' is run.",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			"html_emails_enabled": {
				Description: "Enable HTML emails.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			//"in_product_marketing_emails_enabled": {
			//	Description: "Enable [in-product marketing emails](../user/profile/notifications.md#global-notification-settings). Enabled by default.",
			//	Type:        schema.TypeBool,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"invisible_captcha_enabled": {
			//	Description: "Enable Invisible CAPTCHA spam detection during sign-up. Disabled by default.",
			//	Type:        schema.TypeBool,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"issues_create_limit": {
			//	Description: "Max number of issue creation requests per minute per user. Disabled by default.",
			//	Type:        schema.TypeInt,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"keep_latest_artifact": {
			//	Description: "Prevent the deletion of the artifacts from the most recent successful jobs, regardless of the expiry time. Enabled by default.",
			//	Type:        schema.TypeBool,
			//	Computed:    true,
			//	Optional:    true,
			//},
			"local_markdown_version": {
				Description: "Increase this value when any cached Markdown should be invalidated.",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			//"mailgun_signing_key": {
			//	Description: "The Mailgun HTTP webhook signing key for receiving events from webhook.",
			//	Type:        schema.TypeString,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"mailgun_events_enabled": {
			//	Description: "Enable Mailgun event receiver.",
			//	Type:        schema.TypeBool,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"maintenance_mode_message": {
			//	Description: "Message displayed when instance is in maintenance mode.",
			//	Type:        schema.TypeString,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"maintenance_mode": {
			//	Description: "When instance is in maintenance mode, non-administrative users can sign in with read-only access and make read-only API requests.",
			//	Type:        schema.TypeBool,
			//	Computed:    true,
			//	Optional:    true,
			//},
			"max_artifacts_size": {
				Description: "Maximum artifacts size in MB.",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			"max_attachment_size": {
				Description: "Limit attachment size in MB.",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			//"max_import_size": {
			//	Description: "Maximum import size in MB. 0 for unlimited. Default = 0 (unlimited) [Modified](https://gitlab.com/gitlab-org/gitlab/-/issues/251106) from 50MB to 0 in GitLab 13.8.",
			//	Type:        schema.TypeInt,
			//	Computed:    true,
			//	Optional:    true,
			//},
			"max_pages_size": {
				Description: "Maximum size of pages repositories in MB.",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			//"max_personal_access_token_lifetime": {
			//	Description: "Maximum allowable lifetime for personal access tokens in days.",
			//	Type:        schema.TypeInt,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"max_ssh_key_lifetime": {
			//	Description: "Maximum allowable lifetime for SSH keys in days. [Introduced](https://gitlab.com/gitlab-org/gitlab/-/issues/1007) in GitLab 14.6.",
			//	Type:        schema.TypeInt,
			//	Computed:    true,
			//	Optional:    true,
			//},
			"metrics_method_call_threshold": {
				Description: "A method call is only tracked when it takes longer than the given amount of milliseconds.",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			"mirror_available": {
				Description: "Allow repository mirroring to configured by project Maintainers. If disabled, only Administrators can configure repository mirroring.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"mirror_capacity_threshold": {
				Description: "Minimum capacity to be available before scheduling more mirrors preemptively.",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			"mirror_max_capacity": {
				Description: "Maximum number of mirrors that can be synchronizing at the same time.",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			"mirror_max_delay": {
				Description: "Maximum time (in minutes) between updates that a mirror can have when scheduled to synchronize.",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			//"npm_package_requests_forwarding": {
			//	Description: "Use npmjs.org as a default remote repository when the package is not found in the GitLab Package Registry for npm.",
			//	Type:        schema.TypeBool,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"pypi_package_requests_forwarding": {
			//	Description: "Use pypi.org as a default remote repository when the package is not found in the GitLab Package Registry for PyPI.",
			//	Type:        schema.TypeBool,
			//	Computed:    true,
			//	Optional:    true,
			//},
			"pages_domain_verification_enabled": {
				Description: "Require users to prove ownership of custom domains. Domain verification is an essential security measure for public GitLab sites. Users are required to demonstrate they control a domain before it is enabled.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"password_authentication_enabled_for_git": {
				Description: "Enable authentication for Git over HTTP(S) via a GitLab account password. Default is 'true'.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"password_authentication_enabled_for_web": {
				Description: "Enable authentication for the web interface via a GitLab account password. Default is 'true'.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"performance_bar_allowed_group_path": {
				Description: "Path of the group that is allowed to toggle the performance bar.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			//"personal_access_token_prefix": {
			//	Description: "Prefix for all generated personal access tokens.",
			//	Type:        schema.TypeString,
			//	Computed:    true,
			//	Optional:    true,
			//},
			"plantuml_enabled": {
				Description: "(**If enabled, requires:** 'plantuml_url') Enable PlantUML integration. Default is 'false'.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"plantuml_url": {
				Description: "The PlantUML instance URL for integration.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"polling_interval_multiplier": {
				Description: "Interval multiplier used by endpoints that perform polling. Set to '0' to disable polling.",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			"project_export_enabled": {
				Description: "Enable project export.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"prometheus_metrics_enabled": {
				Description: "Enable Prometheus metrics.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"protected_ci_variables": {
				Description: "CI/CD variables are protected by default.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"pseudonymizer_enabled": {
				Description: "When enabled, GitLab runs a background job that produces pseudonymized CSVs of the GitLab database to upload to your configured object storage directory.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"push_event_activities_limit": {
				Description: "Number of changes (branches or tags) in a single push to determine whether individual push events or bulk push events are created. [Bulk push events are created](../user/admin_area/settings/push_event_activities_limit.md) if it surpasses that value.",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			"push_event_hooks_limit": {
				Description: "Number of changes (branches or tags) in a single push to determine whether webhooks and services fire or not. Webhooks and services aren't submitted if it surpasses that value.",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			//"rate_limiting_response_text": {
			//	Description: "When rate limiting is enabled via the 'throttle_*' settings, send this plain text response when a rate limit is exceeded. 'Retry later' is sent if this is blank.",
			//	Type:        schema.TypeString,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"raw_blob_request_limit": {
			//	Description: "Max number of requests per minute for each raw path. Default: 300. To disable throttling set to 0.",
			//	Type:        schema.TypeInt,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"search_rate_limit": {
			//	Description: "Max number of requests per minute for performing a search while authenticated. Default: 30. To disable throttling set to 0.",
			//	Type:        schema.TypeInt,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"search_rate_limit_unauthenticated": {
			//	Description: "Max number of requests per minute for performing a search while unauthenticated. Default: 10. To disable throttling set to 0.",
			//	Type:        schema.TypeInt,
			//	Computed:    true,
			//	Optional:    true,
			//},
			"recaptcha_enabled": {
				Description: "(**If enabled, requires:** 'recaptcha_private_key' and 'recaptcha_site_key') Enable reCAPTCHA.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"recaptcha_private_key": {
				Description: "Private key for reCAPTCHA.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"recaptcha_site_key": {
				Description: "Site key for reCAPTCHA.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"receive_max_input_size": {
				Description: "Maximum push size (MB).",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			"repository_checks_enabled": {
				Description: "GitLab periodically runs 'git fsck' in all project and wiki repositories to look for silent disk corruption issues.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"repository_size_limit": {
				Description: "Size limit per repository (MB)",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			//"require_admin_approval_after_user_signup": {
			//	Description: "When enabled, any user that signs up for an account using the registration form is placed under a **Pending approval** state and has to be explicitly [approved](../user/admin_area/moderate_users.md) by an administrator.",
			//	Type:        schema.TypeBool,
			//	Computed:    true,
			//	Optional:    true,
			//},
			"require_two_factor_authentication": {
				Description: "(**If enabled, requires:** 'two_factor_grace_period') Require all users to set up Two-factor authentication.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"rsa_key_restriction": {
				Description: "The minimum allowed bit length of an uploaded RSA key. Default is '0' (no restriction). '-1' disables RSA keys.",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			"send_user_confirmation_email": {
				Description: "Send confirmation email on sign-up.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"session_expire_delay": {
				Description: "Session duration in minutes. GitLab restart is required to apply changes.",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			"shared_runners_enabled": {
				Description: "(**If enabled, requires:** 'shared_runners_text' and 'shared_runners_minutes') Enable shared runners for new projects.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"shared_runners_minutes": {
				Description: "Set the maximum number of CI/CD minutes that a group can use on shared runners per month.",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			"shared_runners_text": {
				Description: "Shared runners text.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			//"sidekiq_job_limiter_mode": {
			//	Description: "'track' or 'compress'. Sets the behavior for [Sidekiq job size limits](../user/admin_area/settings/sidekiq_job_limits.md). Default: 'compress'.",
			//	Type:        schema.TypeString,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"sidekiq_job_limiter_compression_threshold_bytes": {
			//	Description: "The threshold in bytes at which Sidekiq jobs are compressed before being stored in Redis. Default: 100 000 bytes (100KB).",
			//	Type:        schema.TypeInt,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"sidekiq_job_limiter_limit_bytes": {
			//	Description: "The threshold in bytes at which Sidekiq jobs are rejected. Default: 0 bytes (doesn't reject any job).",
			//	Type:        schema.TypeInt,
			//	Computed:    true,
			//	Optional:    true,
			//},
			"sign_in_text": {
				Description: "Text on the login page.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"signup_enabled": {
				Description: "Enable registration. Default is 'true'.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"slack_app_enabled": {
				Description: "(**If enabled, requires:** 'slack_app_id', 'slack_app_secret' and 'slack_app_secret') Enable Slack app.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"slack_app_id": {
				Description: "The app ID of the Slack-app.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"slack_app_secret": {
				Description: "The app secret of the Slack-app.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"slack_app_verification_token": {
				Description: "The verification token of the Slack-app.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			//"snippet_size_limit": {
			//	Description: "Max snippet content size in **bytes**. Default: 52428800 Bytes (50MB).",
			//	Type:        schema.TypeInt,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"snowplow_app_id": {
			//	Description: "The Snowplow site name / application ID. (for example, 'gitlab')",
			//	Type:        schema.TypeString,
			//	Computed:    true,
			//	Optional:    true,
			//},
			"snowplow_collector_hostname": {
				Description: "The Snowplow collector hostname. (for example, 'snowplow.trx.gitlab.net')",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"snowplow_cookie_domain": {
				Description: "The Snowplow cookie domain. (for example, '.gitlab.com')",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"snowplow_enabled": {
				Description: "Enable snowplow tracking.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			//"sourcegraph_enabled": {
			//	Description: "Enables Sourcegraph integration. Default is 'false'. **If enabled, requires** 'sourcegraph_url'.",
			//	Type:        schema.TypeBool,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"sourcegraph_public_only": {
			//	Description: "Blocks Sourcegraph from being loaded on private and internal projects. Default is 'true'.",
			//	Type:        schema.TypeBool,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"sourcegraph_url": {
			//	Description: "The Sourcegraph instance URL for integration.",
			//	Type:        schema.TypeString,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"spam_check_endpoint_enabled": {
			//	Description: "Enables spam checking using external Spam Check API endpoint. Default is 'false'.",
			//	Type:        schema.TypeBool,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"spam_check_endpoint_url": {
			//	Description: "URL of the external Spamcheck service endpoint. Valid URI schemes are 'grpc' or 'tls'. Specifying 'tls' forces communication to be encrypted.",
			//	Type:        schema.TypeString,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"spam_check_api_key": {
			//	Description: "API key used by GitLab for accessing the Spam Check service endpoint.",
			//	Type:        schema.TypeString,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"suggest_pipeline_enabled": {
			//	Description: "Enable pipeline suggestion banner.",
			//	Type:        schema.TypeBool,
			//	Computed:    true,
			//	Optional:    true,
			//},
			"terminal_max_session_time": {
				Description: "Maximum time for web terminal websocket connection (in seconds). Set to '0' for unlimited time.",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			"terms": {
				Description: "(**Required by:** 'enforce_terms') Markdown content for the ToS.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"throttle_authenticated_api_enabled": {
				Description: "(**If enabled, requires:** 'throttle_authenticated_api_period_in_seconds' and 'throttle_authenticated_api_requests_per_period') Enable authenticated API request rate limit. Helps reduce request volume (for example, from crawlers or abusive bots).",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"throttle_authenticated_api_period_in_seconds": {
				Description: "Rate limit period (in seconds).",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			"throttle_authenticated_api_requests_per_period": {
				Description: "Maximum requests per period per user.",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			"throttle_authenticated_web_enabled": {
				Description: "(**If enabled, requires:** 'throttle_authenticated_web_period_in_seconds' and 'throttle_authenticated_web_requests_per_period') Enable authenticated web request rate limit. Helps reduce request volume (for example, from crawlers or abusive bots).",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"throttle_authenticated_web_period_in_seconds": {
				Description: "Rate limit period (in seconds).",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			"throttle_authenticated_web_requests_per_period": {
				Description: "Maximum requests per period per user.",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			//"throttle_unauthenticated_api_enabled": {
			//	Description: "(**If enabled, requires:** 'throttle_unauthenticated_api_period_in_seconds' and 'throttle_unauthenticated_api_requests_per_period') Enable unauthenticated API request rate limit. Helps reduce request volume (for example, from crawlers or abusive bots).",
			//	Type:        schema.TypeBool,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"throttle_unauthenticated_api_period_in_seconds": {
			//	Description: "Rate limit period in seconds.",
			//	Type:        schema.TypeInt,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"throttle_unauthenticated_api_requests_per_period": {
			//	Description: "Max requests per period per IP.",
			//	Type:        schema.TypeInt,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"throttle_unauthenticated_web_enabled": {
			//	Description: "(**If enabled, requires:** 'throttle_unauthenticated_web_period_in_seconds' and 'throttle_unauthenticated_web_requests_per_period') Enable unauthenticated web request rate limit. Helps reduce request volume (for example, from crawlers or abusive bots).",
			//	Type:        schema.TypeBool,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"throttle_unauthenticated_web_period_in_seconds": {
			//	Description: "Rate limit period in seconds.",
			//	Type:        schema.TypeInt,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"throttle_unauthenticated_web_requests_per_period": {
			//	Description: "Max requests per period per IP.",
			//	Type:        schema.TypeInt,
			//	Computed:    true,
			//	Optional:    true,
			//},
			"time_tracking_limit_to_hours": {
				Description: "Limit display of time tracking units to hours. Default is 'false'.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"two_factor_grace_period": {
				Description: "Amount of time (in hours) that users are allowed to skip forced configuration of two-factor authentication.",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			"unique_ips_limit_enabled": {
				Description: "(**If enabled, requires:** 'unique_ips_limit_per_user' and 'unique_ips_limit_time_window') Limit sign in from multiple IPs.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"unique_ips_limit_per_user": {
				Description: "Maximum number of IPs per user.",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			"unique_ips_limit_time_window": {
				Description: "How many seconds an IP is counted towards the limit.",
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
			},
			"usage_ping_enabled": {
				Description: "Every week GitLab reports license usage back to GitLab, Inc.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			//"user_deactivation_emails_enabled": {
			//	Description: "Send an email to users upon account deactivation.",
			//	Type:        schema.TypeBool,
			//	Computed:    true,
			//	Optional:    true,
			//},
			"user_default_external": {
				Description: "Newly registered users are external by default.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"user_default_internal_regex": {
				Description: "Specify an email address regex pattern to identify default internal users.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
			"user_oauth_applications": {
				Description: "Allow users to register any application to use GitLab as an OAuth provider.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"user_show_add_ssh_key_message": {
				Description: "When set to 'false' disable the 'You won't be able to pull or push project code via SSH' warning shown to users with no uploaded SSH key.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			"version_check_enabled": {
				Description: "Let GitLab inform you when an update is available.",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			//"whats_new_variant": {
			//	Description: "What's new variant, possible values: 'all_tiers', 'current_tier', and 'disabled'.",
			//	Type:        schema.TypeString,
			//	Computed:    true,
			//	Optional:    true,
			//},
			"web_ide_clientside_preview_enabled": {
				Description: "Live Preview (allow live previews of JavaScript projects in the Web IDE using CodeSandbox Live Preview).",
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
			},
			//"wiki_page_max_content_bytes": {
			//	Description: "Maximum wiki page content size in **bytes**. Default: 52428800 Bytes (50 MB). The minimum value is 1024 bytes.",
			//	Type:        schema.TypeInt,
			//	Computed:    true,
			//	Optional:    true,
			//},
			//"asset_proxy_allowlist": {
			//	Description: "Assets that match these domain(s) are **not** proxied. Wildcards allowed. Your GitLab installation URL is automatically allowlisted. GitLab restart is required to apply changes.",
			//	Type:        schema.TypeList,
			//	Elem: &schema.Schema{
			//		Type: schema.TypeString,
			//	},
			//	Computed: true,
			//	Optional: true,
			//},
			"disabled_oauth_sign_in_sources": {
				Description: "Disabled OAuth sign-in sources.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed: true,
				Optional: true,
			},
			//"domain_denylist": {
			//	Description: "Users with email addresses that match these domain(s) **cannot** sign up. Wildcards allowed. Use separate lines for multiple entries. Ex: 'domain.com', '*.domain.com'.",
			//	Type:        schema.TypeList,
			//	Elem: &schema.Schema{
			//		Type: schema.TypeString,
			//	},
			//	Computed: true,
			//	Optional: true,
			//},
			//"domain_allowlist": {
			//	Description: "Force people to use only corporate emails for sign-up. Default is 'null', meaning there is no restriction.",
			//	Type:        schema.TypeList,
			//	Elem: &schema.Schema{
			//		Type: schema.TypeString,
			//	},
			//	Computed: true,
			//	Optional: true,
			//},
			"elasticsearch_namespace_ids": {
				Description: "The namespaces to index via Elasticsearch if 'elasticsearch_limit_indexing' is enabled.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Computed: true,
				Optional: true,
			},
			"elasticsearch_project_ids": {
				Description: "The projects to index via Elasticsearch if 'elasticsearch_limit_indexing' is enabled.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Computed: true,
				Optional: true,
			},
			"import_sources": {
				Description: "Sources to allow project import from, possible values: 'github', 'bitbucket', 'bitbucket_server', 'gitlab', 'fogbugz', 'git', 'gitlab_project', 'gitea', 'manifest', and 'phabricator'.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				// TODO: validate possible values
				Computed: true,
				Optional: true,
			},
			"outbound_local_requests_whitelist": {
				Description: "Define a list of trusted domains or IP addresses to which local requests are allowed when local requests for hooks and services are disabled.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed: true,
				Optional: true,
			},
			//"repository_storages_weighted": {
			//	Description: "(GitLab 13.1 and later) Hash of names of taken from 'gitlab.yml' to [weights](../administration/repository_storage_paths.md#configure-where-new-repositories-are-stored). New projects are created in one of these stores, chosen by a weighted random selection.",
			//	Type:        schema.TypeMap,
			//	Elem: &schema.Schema{
			//		Type: schema.TypeInt,
			//	},
			//	Computed: true,
			//	Optional: true,
			//},
			"repository_storages": {
				Description: "(GitLab 13.0 and earlier) List of names of enabled storage paths, taken from 'gitlab.yml'. New projects are created in one of these stores, chosen at random.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed: true,
				Optional: true,
			},
			"restricted_visibility_levels": {
				Description: "Selected levels cannot be used by non-Administrator users for groups, projects or snippets. Can take 'private', 'internal' and 'public' as a parameter. Default is 'null' which means there is no restriction.",
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				// TODO: validate possible values
				Computed: true,
				Optional: true,
			},
		},
	}
})

func dataSourceGitlabInstanceSettingsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*gitlab.Client)

	settings, _, err := client.Settings.GetSettings(gitlab.WithContext(ctx))
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("admin_mode", settings.AdminMode)
	//d.Set("abuse_notification_email", settings.AbuseNotificationEmail)
	d.Set("after_sign_out_path", settings.AfterSignOutPath)
	d.Set("after_sign_up_text", settings.AfterSignUpText)
	d.Set("akismet_api_key", settings.AkismetAPIKey)
	d.Set("akismet_enabled", settings.AkismetEnabled)
	d.Set("allow_group_owners_to_manage_ldap", settings.AllowGroupOwnersToManageLDAP)
	d.Set("allow_local_requests_from_system_hooks", settings.AllowLocalRequestsFromSystemHooks)
	d.Set("allow_local_requests_from_web_hooks_and_services", settings.AllowLocalRequestsFromWebHooksAndServices)
	d.Set("archive_builds_in_human_readable", settings.ArchiveBuildsInHumanReadable)
	d.Set("asset_proxy_enabled", settings.AssetProxyEnabled)
	d.Set("asset_proxy_secret_key", settings.AssetProxySecretKey)
	d.Set("asset_proxy_url", settings.AssetProxyURL)
	d.Set("authorized_keys_enabled", settings.AuthorizedKeysEnabled)
	d.Set("auto_devops_domain", settings.AutoDevOpsDomain)
	d.Set("auto_devops_enabled", settings.AutoDevOpsEnabled)
	//d.Set("automatic_purchased_storage_allocation", settings.AutomaticPurchasedStorageAllocation)
	d.Set("check_namespace_plan", settings.CheckNamespacePlan)
	d.Set("commit_email_hostname", settings.CommitEmailHostname)
	//d.Set("container_expiration_policies_enable_historic_entries", settings.ContainerExpirationPoliciesEnableHistoricEntries)
	//d.Set("container_registry_cleanup_tags_service_max_list_size", settings.ContainerRegistryCleanupTagsServiceMaxListSize)
	//d.Set("container_registry_delete_tags_service_timeout", settings.ContainerRegistryDeleteTagsServiceTimeout)
	//d.Set("container_registry_expiration_policies_caching", settings.ContainerRegistryExpirationPoliciesCaching)
	//d.Set("container_registry_expiration_policies_worker_capacity", settings.ContainerRegistryExpirationPoliciesWorkerCapacity)
	d.Set("container_registry_token_expire_delay", settings.ContainerRegistryTokenExpireDelay)
	//d.Set("deactivate_dormant_users", settings.DeactivateDormantUsers)
	d.Set("default_artifacts_expire_in", settings.DefaultArtifactsExpireIn)
	//d.Set("default_branch_name", settings.DefaultBranchName)
	d.Set("default_branch_protection", settings.DefaultBranchProtection)
	//d.Set("default_ci_config_path", settings.DefaultCIConfigPath)
	d.Set("default_group_visibility", settings.DefaultGroupVisibility)
	d.Set("default_project_creation", settings.DefaultProjectCreation)
	d.Set("default_project_visibility", settings.DefaultProjectVisibility)
	d.Set("default_projects_limit", settings.DefaultProjectsLimit)
	d.Set("default_snippet_visibility", settings.DefaultSnippetVisibility)
	//d.Set("delayed_project_deletion", settings.DelayedProjectDeletion)
	//d.Set("deletion_adjourned_period", settings.DeletionAdjournedPeriod)
	d.Set("diff_max_patch_bytes", settings.DiffMaxPatchBytes)
	//d.Set("diff_max_files", settings.DiffMaxFiles)
	//d.Set("diff_max_lines", settings.DiffMaxLines)
	//d.Set("disable_feed_token", settings.DisableFeedToken)
	d.Set("dns_rebinding_protection_enabled", settings.DNSRebindingProtectionEnabled)
	//d.Set("domain_denylist_enabled", settings.DomainDenylistEnabled) [sic: blacklist in go-gitlab]
	d.Set("dsa_key_restriction", settings.DSAKeyRestriction)
	d.Set("ecdsa_key_restriction", settings.ECDSAKeyRestriction)
	//d.Set("ecdsa_sk_key_restriction", settings.ECDSASKKeyRestriction)
	d.Set("ed25519_key_restriction", settings.Ed25519KeyRestriction)
	//d.Set("ed25519_sk_key_restriction", settings.Ed25519SkKeyRestriction)
	//d.Set("eks_access_key_id", settings.EksAccessKeyId)
	//d.Set("eks_account_id", settings.EksAccountId)
	//d.Set("eks_integration_enabled", settings.EksIntegrationEnabled)
	//d.Set("eks_secret_access_key", settings.EksSecretAccessKey)
	d.Set("elasticsearch_aws_access_key", settings.ElasticsearchAWSAccessKey)
	d.Set("elasticsearch_aws_region", settings.ElasticsearchAWSRegion)
	d.Set("elasticsearch_aws_secret_access_key", settings.ElasticsearchAWSSecretAccessKey)
	d.Set("elasticsearch_aws", settings.ElasticsearchAWS)
	//d.Set("elasticsearch_indexed_field_length_limit", settings.ElasticsearchIndexedFieldLengthLimit)
	//d.Set("elasticsearch_indexed_file_size_limit_kb", settings.ElasticsearchIndexedFileSizeLimitKb)
	d.Set("elasticsearch_indexing", settings.ElasticsearchIndexing)
	d.Set("elasticsearch_limit_indexing", settings.ElasticsearchLimitIndexing)
	//d.Set("elasticsearch_max_bulk_concurrency", settings.ElasticsearchMaxBulkConcurrency)
	//d.Set("elasticsearch_max_bulk_size_mb", settings.ElasticsearchMaxBulkSizeMb)
	d.Set("elasticsearch_search", settings.ElasticsearchSearch)
	d.Set("elasticsearch_url", settings.ElasticsearchURL)
	//d.Set("elasticsearch_username", settings.ElasticsearchUsername)
	//d.Set("elasticsearch_password", settings.ElasticsearchPassword)
	d.Set("email_additional_text", settings.EmailAdditionalText)
	d.Set("email_author_in_body", settings.EmailAuthorInBody)
	d.Set("enabled_git_access_protocol", settings.EnabledGitAccessProtocol)
	//d.Set("enforce_namespace_storage_limit", settings.EnforceNamespaceStorageLimit)
	d.Set("enforce_terms", settings.EnforceTerms)
	d.Set("external_auth_client_cert", settings.ExternalAuthClientCert)
	d.Set("external_auth_client_key_pass", settings.ExternalAuthClientKeyPass)
	d.Set("external_auth_client_key", settings.ExternalAuthClientKey)
	d.Set("external_authorization_service_default_label", settings.ExternalAuthorizationServiceDefaultLabel)
	d.Set("external_authorization_service_enabled", settings.ExternalAuthorizationServiceEnabled)
	d.Set("external_authorization_service_timeout", settings.ExternalAuthorizationServiceTimeout)
	d.Set("external_authorization_service_url", settings.ExternalAuthorizationServiceURL)
	//d.Set("external_pipeline_validation_service_url", settings.ExternalPipelineValidationServiceURL)
	//d.Set("external_pipeline_validation_service_token", settings.ExternalPipelineValidationServiceToken)
	//d.Set("external_pipeline_validation_service_timeout", settings.ExternalPipelineValidationServiceTimeout)
	d.Set("file_template_project_id", settings.FileTemplateProjectID)
	d.Set("first_day_of_week", settings.FirstDayOfWeek)
	d.Set("geo_node_allowed_ips", settings.GeoNodeAllowedIPs)
	d.Set("geo_status_timeout", settings.GeoStatusTimeout)
	//d.Set("git_two_factor_session_expiry", settings.GitTwoFactorSessionExpiry)
	d.Set("gitaly_timeout_default", settings.GitalyTimeoutDefault)
	d.Set("gitaly_timeout_fast", settings.GitalyTimeoutFast)
	d.Set("gitaly_timeout_medium", settings.GitalyTimeoutMedium)
	d.Set("grafana_enabled", settings.GrafanaEnabled)
	d.Set("grafana_url", settings.GrafanaURL)
	d.Set("gravatar_enabled", settings.GravatarEnabled)
	d.Set("hashed_storage_enabled", settings.HashedStorageEnabled)
	d.Set("help_page_hide_commercial_content", settings.HelpPageHideCommercialContent)
	d.Set("help_page_support_url", settings.HelpPageSupportURL)
	d.Set("help_page_text", settings.HelpPageText)
	d.Set("help_text", settings.HelpText)
	d.Set("hide_third_party_offers", settings.HideThirdPartyOffers)
	d.Set("home_page_url", settings.HomePageURL)
	d.Set("housekeeping_bitmaps_enabled", settings.HousekeepingBitmapsEnabled)
	d.Set("housekeeping_enabled", settings.HousekeepingEnabled)
	d.Set("housekeeping_full_repack_period", settings.HousekeepingFullRepackPeriod)
	d.Set("housekeeping_gc_period", settings.HousekeepingGcPeriod)
	d.Set("housekeeping_incremental_repack_period", settings.HousekeepingIncrementalRepackPeriod)
	d.Set("html_emails_enabled", settings.HTMLEmailsEnabled)
	//d.Set("in_product_marketing_emails_enabled", settings.InProductMarketingEmailsEnabled)
	//d.Set("invisible_captcha_enabled", settings.InvisibleCaptchaEnabled)
	//d.Set("issues_create_limit", settings.IssuesCreateLimit)
	//d.Set("keep_latest_artifact", settings.KeepLatestArtifact)
	d.Set("local_markdown_version", settings.LocalMarkdownVersion)
	//d.Set("mailgun_signing_key", settings.MailgunSigningKey)
	//d.Set("mailgun_events_enabled", settings.MailgunEventsEnabled)
	//d.Set("maintenance_mode_message", settings.MaintenanceModeMessage)
	//d.Set("maintenance_mode", settings.MaintenanceMode)
	d.Set("max_artifacts_size", settings.MaxArtifactsSize)
	d.Set("max_attachment_size", settings.MaxAttachmentSize)
	//d.Set("max_import_size", settings.MaxImportSize)
	d.Set("max_pages_size", settings.MaxPagesSize)
	//d.Set("max_personal_access_token_lifetime", settings.MaxPersonalAccessTokenLifetime)
	//d.Set("max_ssh_key_lifetime", settings.MaxSSHKeyLifetime)
	d.Set("metrics_method_call_threshold", settings.MetricsMethodCallThreshold)
	d.Set("mirror_available", settings.MirrorAvailable)
	d.Set("mirror_capacity_threshold", settings.MirrorCapacityThreshold)
	d.Set("mirror_max_capacity", settings.MirrorMaxCapacity)
	d.Set("mirror_max_delay", settings.MirrorMaxDelay)
	//d.Set("npm_package_requests_forwarding", settings.NpmPackageRequestsForwarding)
	//d.Set("pypi_package_requests_forwarding", settings.PypiPackageRequestsForwarding)
	d.Set("pages_domain_verification_enabled", settings.PagesDomainVerificationEnabled)
	d.Set("password_authentication_enabled_for_git", settings.PasswordAuthenticationEnabledForGit)
	d.Set("password_authentication_enabled_for_web", settings.PasswordAuthenticationEnabledForWeb)
	d.Set("performance_bar_allowed_group_path", settings.PerformanceBarAllowedGroupPath)
	//d.Set("personal_access_token_prefix", settings.PersonalAccessTokenPrefix)
	d.Set("plantuml_enabled", settings.PlantumlEnabled)
	d.Set("plantuml_url", settings.PlantumlURL)
	d.Set("polling_interval_multiplier", settings.PollingIntervalMultiplier)
	d.Set("project_export_enabled", settings.ProjectExportEnabled)
	d.Set("prometheus_metrics_enabled", settings.PrometheusMetricsEnabled)
	d.Set("protected_ci_variables", settings.ProtectedCIVariables)
	d.Set("pseudonymizer_enabled", settings.PseudonymizerEnabled)
	d.Set("push_event_activities_limit", settings.PushEventActivitiesLimit)
	d.Set("push_event_hooks_limit", settings.PushEventHooksLimit)
	//d.Set("rate_limiting_response_text", settings.RateLimitingResponseText)
	//d.Set("raw_blob_request_limit", settings.RawBlobRequestLimit)
	//d.Set("search_rate_limit", settings.SearchRateLimit)
	//d.Set("search_rate_limit_unauthenticated", settings.SearchRateLimitUnauthenticated)
	d.Set("recaptcha_enabled", settings.RecaptchaEnabled)
	d.Set("recaptcha_private_key", settings.RecaptchaPrivateKey)
	d.Set("recaptcha_site_key", settings.RecaptchaSiteKey)
	d.Set("receive_max_input_size", settings.ReceiveMaxInputSize)
	d.Set("repository_checks_enabled", settings.RepositoryChecksEnabled)
	d.Set("repository_size_limit", settings.RepositorySizeLimit)
	//d.Set("require_admin_approval_after_user_signup", settings.RequireAdminApprovalAfterUserSignup)
	d.Set("require_two_factor_authentication", settings.RequireTwoFactorAuthentication)
	d.Set("rsa_key_restriction", settings.RsaKeyRestriction)
	d.Set("send_user_confirmation_email", settings.SendUserConfirmationEmail)
	d.Set("session_expire_delay", settings.SessionExpireDelay)
	d.Set("shared_runners_enabled", settings.SharedRunnersEnabled)
	d.Set("shared_runners_minutes", settings.SharedRunnersMinutes)
	d.Set("shared_runners_text", settings.SharedRunnersText)
	//d.Set("sidekiq_job_limiter_mode", settings.SidekiqJobLimiterMode)
	//d.Set("sidekiq_job_limiter_compression_threshold_bytes", settings.SidekiqJobLimiterCompressionThresholdBytes)
	//d.Set("sidekiq_job_limiter_limit_bytes", settings.SidekiqJobLimiterLimitBytes)
	d.Set("sign_in_text", settings.SignInText)
	d.Set("signup_enabled", settings.SignupEnabled)
	d.Set("slack_app_enabled", settings.SlackAppEnabled)
	d.Set("slack_app_id", settings.SlackAppID)
	d.Set("slack_app_secret", settings.SlackAppSecret)
	d.Set("slack_app_verification_token", settings.SlackAppVerificationToken)
	//d.Set("snippet_size_limit", settings.SnippetSizeLimit)
	//d.Set("snowplow_app_id", settings.SnowplowAppID) [sic: SnowplowSiteID in go-gitlab]
	d.Set("snowplow_collector_hostname", settings.SnowplowCollectorHostname)
	d.Set("snowplow_cookie_domain", settings.SnowplowCookieDomain)
	d.Set("snowplow_enabled", settings.SnowplowEnabled)
	//d.Set("sourcegraph_enabled", settings.SourcegraphEnabled)
	//d.Set("sourcegraph_public_only", settings.SourcegraphPublicOnly)
	//d.Set("sourcegraph_url", settings.SourcegraphUrl)
	//d.Set("spam_check_endpoint_enabled", settings.SpamCheckEndpointEnabled)
	//d.Set("spam_check_endpoint_url", settings.SpamCheckEndpointUrl)
	//d.Set("spam_check_api_key", settings.SpamCheckApiKey)
	//d.Set("suggest_pipeline_enabled", settings.SuggestPipelineEnabled)
	d.Set("terminal_max_session_time", settings.TerminalMaxSessionTime)
	d.Set("terms", settings.Terms)
	d.Set("throttle_authenticated_api_enabled", settings.ThrottleAuthenticatedAPIEnabled)
	d.Set("throttle_authenticated_api_period_in_seconds", settings.ThrottleAuthenticatedAPIPeriodInSeconds)
	d.Set("throttle_authenticated_api_requests_per_period", settings.ThrottleAuthenticatedAPIRequestsPerPeriod)
	d.Set("throttle_authenticated_web_enabled", settings.ThrottleAuthenticatedWebEnabled)
	d.Set("throttle_authenticated_web_period_in_seconds", settings.ThrottleAuthenticatedWebPeriodInSeconds)
	d.Set("throttle_authenticated_web_requests_per_period", settings.ThrottleAuthenticatedWebRequestsPerPeriod)
	//d.Set("throttle_unauthenticated_api_enabled", settings.ThrottleUnauthenticatedAPIEnabled)
	//d.Set("throttle_unauthenticated_api_period_in_seconds", settings.ThrottleUnauthenticatedAPIPeriodInSeconds)
	//d.Set("throttle_unauthenticated_api_requests_per_period", settings.ThrottleUnauthenticatedAPIRequestsPerPeriod)
	//d.Set("throttle_unauthenticated_web_enabled", settings.ThrottleUnauthenticatedWebEnabled)
	//d.Set("throttle_unauthenticated_web_period_in_seconds", settings.ThrottleUnauthenticatedWebPeriodInSeconds)
	//d.Set("throttle_unauthenticated_web_requests_per_period", settings.ThrottleUnauthenticatedWebRequestsPerPeriod)
	d.Set("time_tracking_limit_to_hours", settings.TimeTrackingLimitToHours)
	d.Set("two_factor_grace_period", settings.TwoFactorGracePeriod)
	d.Set("unique_ips_limit_enabled", settings.UniqueIPsLimitEnabled)
	d.Set("unique_ips_limit_per_user", settings.UniqueIPsLimitPerUser)
	d.Set("unique_ips_limit_time_window", settings.UniqueIPsLimitTimeWindow)
	d.Set("usage_ping_enabled", settings.UsagePingEnabled)
	//d.Set("user_deactivation_emails_enabled", settings.UserDeactivationEmailsEnabled)
	d.Set("user_default_external", settings.UserDefaultExternal)
	d.Set("user_default_internal_regex", settings.UserDefaultInternalRegex)
	d.Set("user_oauth_applications", settings.UserOauthApplications)
	d.Set("user_show_add_ssh_key_message", settings.UserShowAddSSHKeyMessage)
	d.Set("version_check_enabled", settings.VersionCheckEnabled)
	//d.Set("whats_new_variant", settings.WhatsNewVariant)
	d.Set("web_ide_clientside_preview_enabled", settings.WebIDEClientsidePreviewEnabled)
	//d.Set("wiki_page_max_content_bytes", settings.WikiPageMaxContentBytes)
	//d.Set("asset_proxy_allowlist", settings.AssetProxyAllowlist) [sic: whitelist in go-gitlab]
	d.Set("disabled_oauth_sign_in_sources", settings.DisabledOauthSignInSources)
	//d.Set("domain_denylist", settings.DomainDenylist) [sic: blacklist in go-gitlab]
	//d.Set("domain_allowlist", settings.DomainAllowlist) [sic: whitelist in go-gitlab]
	d.Set("elasticsearch_namespace_ids", settings.ElasticsearchNamespaceIDs)
	d.Set("elasticsearch_project_ids", settings.ElasticsearchProjectIDs)
	d.Set("import_sources", settings.ImportSources)
	d.Set("outbound_local_requests_whitelist", settings.OutboundLocalRequestsWhitelist)
	//d.Set("repository_storages_weighted", settings.RepositoryStoragesWeighted)
	d.Set("repository_storages", settings.RepositoryStorages)
	d.Set("restricted_visibility_levels", settings.RestrictedVisibilityLevels)

	// NOTE: the settings are global and not "created". Therefore there is no id
	//       available or required. We just set the url as a dummy id instead.
	d.SetId(client.BaseURL().String())
	return nil
}
