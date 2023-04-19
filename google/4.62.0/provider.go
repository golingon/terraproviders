// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	provider "github.com/golingon/terraproviders/google/4.62.0/provider"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewProvider(args ProviderArgs) *Provider {
	return &Provider{Args: args}
}

var _ terra.Provider = (*Provider)(nil)

type Provider struct {
	Args ProviderArgs
}

// LocalName returns the provider local name for [Provider].
func (p *Provider) LocalName() string {
	return "google"
}

// Source returns the provider source for [Provider].
func (p *Provider) Source() string {
	return "hashicorp/google"
}

// Version returns the provider version for [Provider].
func (p *Provider) Version() string {
	return "4.62.0"
}

// Configuration returns the configuration (args) for [Provider].
func (p *Provider) Configuration() interface{} {
	return p.Args
}

// ProviderArgs contains the configurations for provider.
type ProviderArgs struct {
	// AccessApprovalCustomEndpoint: string, optional
	AccessApprovalCustomEndpoint terra.StringValue `hcl:"access_approval_custom_endpoint,attr"`
	// AccessContextManagerCustomEndpoint: string, optional
	AccessContextManagerCustomEndpoint terra.StringValue `hcl:"access_context_manager_custom_endpoint,attr"`
	// AccessToken: string, optional
	AccessToken terra.StringValue `hcl:"access_token,attr"`
	// ActiveDirectoryCustomEndpoint: string, optional
	ActiveDirectoryCustomEndpoint terra.StringValue `hcl:"active_directory_custom_endpoint,attr"`
	// AlloydbCustomEndpoint: string, optional
	AlloydbCustomEndpoint terra.StringValue `hcl:"alloydb_custom_endpoint,attr"`
	// ApigeeCustomEndpoint: string, optional
	ApigeeCustomEndpoint terra.StringValue `hcl:"apigee_custom_endpoint,attr"`
	// ApikeysCustomEndpoint: string, optional
	ApikeysCustomEndpoint terra.StringValue `hcl:"apikeys_custom_endpoint,attr"`
	// AppEngineCustomEndpoint: string, optional
	AppEngineCustomEndpoint terra.StringValue `hcl:"app_engine_custom_endpoint,attr"`
	// ArtifactRegistryCustomEndpoint: string, optional
	ArtifactRegistryCustomEndpoint terra.StringValue `hcl:"artifact_registry_custom_endpoint,attr"`
	// AssuredWorkloadsCustomEndpoint: string, optional
	AssuredWorkloadsCustomEndpoint terra.StringValue `hcl:"assured_workloads_custom_endpoint,attr"`
	// BeyondcorpCustomEndpoint: string, optional
	BeyondcorpCustomEndpoint terra.StringValue `hcl:"beyondcorp_custom_endpoint,attr"`
	// BigQueryCustomEndpoint: string, optional
	BigQueryCustomEndpoint terra.StringValue `hcl:"big_query_custom_endpoint,attr"`
	// BigqueryAnalyticsHubCustomEndpoint: string, optional
	BigqueryAnalyticsHubCustomEndpoint terra.StringValue `hcl:"bigquery_analytics_hub_custom_endpoint,attr"`
	// BigqueryConnectionCustomEndpoint: string, optional
	BigqueryConnectionCustomEndpoint terra.StringValue `hcl:"bigquery_connection_custom_endpoint,attr"`
	// BigqueryDataTransferCustomEndpoint: string, optional
	BigqueryDataTransferCustomEndpoint terra.StringValue `hcl:"bigquery_data_transfer_custom_endpoint,attr"`
	// BigqueryDatapolicyCustomEndpoint: string, optional
	BigqueryDatapolicyCustomEndpoint terra.StringValue `hcl:"bigquery_datapolicy_custom_endpoint,attr"`
	// BigqueryReservationCustomEndpoint: string, optional
	BigqueryReservationCustomEndpoint terra.StringValue `hcl:"bigquery_reservation_custom_endpoint,attr"`
	// BigtableCustomEndpoint: string, optional
	BigtableCustomEndpoint terra.StringValue `hcl:"bigtable_custom_endpoint,attr"`
	// BillingCustomEndpoint: string, optional
	BillingCustomEndpoint terra.StringValue `hcl:"billing_custom_endpoint,attr"`
	// BillingProject: string, optional
	BillingProject terra.StringValue `hcl:"billing_project,attr"`
	// BinaryAuthorizationCustomEndpoint: string, optional
	BinaryAuthorizationCustomEndpoint terra.StringValue `hcl:"binary_authorization_custom_endpoint,attr"`
	// CertificateManagerCustomEndpoint: string, optional
	CertificateManagerCustomEndpoint terra.StringValue `hcl:"certificate_manager_custom_endpoint,attr"`
	// CloudAssetCustomEndpoint: string, optional
	CloudAssetCustomEndpoint terra.StringValue `hcl:"cloud_asset_custom_endpoint,attr"`
	// CloudBillingCustomEndpoint: string, optional
	CloudBillingCustomEndpoint terra.StringValue `hcl:"cloud_billing_custom_endpoint,attr"`
	// CloudBuildCustomEndpoint: string, optional
	CloudBuildCustomEndpoint terra.StringValue `hcl:"cloud_build_custom_endpoint,attr"`
	// CloudBuildWorkerPoolCustomEndpoint: string, optional
	CloudBuildWorkerPoolCustomEndpoint terra.StringValue `hcl:"cloud_build_worker_pool_custom_endpoint,attr"`
	// CloudFunctionsCustomEndpoint: string, optional
	CloudFunctionsCustomEndpoint terra.StringValue `hcl:"cloud_functions_custom_endpoint,attr"`
	// CloudIdentityCustomEndpoint: string, optional
	CloudIdentityCustomEndpoint terra.StringValue `hcl:"cloud_identity_custom_endpoint,attr"`
	// CloudIdsCustomEndpoint: string, optional
	CloudIdsCustomEndpoint terra.StringValue `hcl:"cloud_ids_custom_endpoint,attr"`
	// CloudIotCustomEndpoint: string, optional
	CloudIotCustomEndpoint terra.StringValue `hcl:"cloud_iot_custom_endpoint,attr"`
	// CloudResourceManagerCustomEndpoint: string, optional
	CloudResourceManagerCustomEndpoint terra.StringValue `hcl:"cloud_resource_manager_custom_endpoint,attr"`
	// CloudRunCustomEndpoint: string, optional
	CloudRunCustomEndpoint terra.StringValue `hcl:"cloud_run_custom_endpoint,attr"`
	// CloudRunV2CustomEndpoint: string, optional
	CloudRunV2CustomEndpoint terra.StringValue `hcl:"cloud_run_v2_custom_endpoint,attr"`
	// CloudSchedulerCustomEndpoint: string, optional
	CloudSchedulerCustomEndpoint terra.StringValue `hcl:"cloud_scheduler_custom_endpoint,attr"`
	// CloudTasksCustomEndpoint: string, optional
	CloudTasksCustomEndpoint terra.StringValue `hcl:"cloud_tasks_custom_endpoint,attr"`
	// ClouddeployCustomEndpoint: string, optional
	ClouddeployCustomEndpoint terra.StringValue `hcl:"clouddeploy_custom_endpoint,attr"`
	// Cloudfunctions2CustomEndpoint: string, optional
	Cloudfunctions2CustomEndpoint terra.StringValue `hcl:"cloudfunctions2_custom_endpoint,attr"`
	// ComposerCustomEndpoint: string, optional
	ComposerCustomEndpoint terra.StringValue `hcl:"composer_custom_endpoint,attr"`
	// ComputeCustomEndpoint: string, optional
	ComputeCustomEndpoint terra.StringValue `hcl:"compute_custom_endpoint,attr"`
	// ContainerAnalysisCustomEndpoint: string, optional
	ContainerAnalysisCustomEndpoint terra.StringValue `hcl:"container_analysis_custom_endpoint,attr"`
	// ContainerAttachedCustomEndpoint: string, optional
	ContainerAttachedCustomEndpoint terra.StringValue `hcl:"container_attached_custom_endpoint,attr"`
	// ContainerAwsCustomEndpoint: string, optional
	ContainerAwsCustomEndpoint terra.StringValue `hcl:"container_aws_custom_endpoint,attr"`
	// ContainerAzureCustomEndpoint: string, optional
	ContainerAzureCustomEndpoint terra.StringValue `hcl:"container_azure_custom_endpoint,attr"`
	// ContainerCustomEndpoint: string, optional
	ContainerCustomEndpoint terra.StringValue `hcl:"container_custom_endpoint,attr"`
	// Credentials: string, optional
	Credentials terra.StringValue `hcl:"credentials,attr"`
	// DataCatalogCustomEndpoint: string, optional
	DataCatalogCustomEndpoint terra.StringValue `hcl:"data_catalog_custom_endpoint,attr"`
	// DataFusionCustomEndpoint: string, optional
	DataFusionCustomEndpoint terra.StringValue `hcl:"data_fusion_custom_endpoint,attr"`
	// DataLossPreventionCustomEndpoint: string, optional
	DataLossPreventionCustomEndpoint terra.StringValue `hcl:"data_loss_prevention_custom_endpoint,attr"`
	// DataflowCustomEndpoint: string, optional
	DataflowCustomEndpoint terra.StringValue `hcl:"dataflow_custom_endpoint,attr"`
	// DataplexCustomEndpoint: string, optional
	DataplexCustomEndpoint terra.StringValue `hcl:"dataplex_custom_endpoint,attr"`
	// DataprocCustomEndpoint: string, optional
	DataprocCustomEndpoint terra.StringValue `hcl:"dataproc_custom_endpoint,attr"`
	// DataprocMetastoreCustomEndpoint: string, optional
	DataprocMetastoreCustomEndpoint terra.StringValue `hcl:"dataproc_metastore_custom_endpoint,attr"`
	// DatastoreCustomEndpoint: string, optional
	DatastoreCustomEndpoint terra.StringValue `hcl:"datastore_custom_endpoint,attr"`
	// DatastreamCustomEndpoint: string, optional
	DatastreamCustomEndpoint terra.StringValue `hcl:"datastream_custom_endpoint,attr"`
	// DeploymentManagerCustomEndpoint: string, optional
	DeploymentManagerCustomEndpoint terra.StringValue `hcl:"deployment_manager_custom_endpoint,attr"`
	// DialogflowCustomEndpoint: string, optional
	DialogflowCustomEndpoint terra.StringValue `hcl:"dialogflow_custom_endpoint,attr"`
	// DialogflowCxCustomEndpoint: string, optional
	DialogflowCxCustomEndpoint terra.StringValue `hcl:"dialogflow_cx_custom_endpoint,attr"`
	// DnsCustomEndpoint: string, optional
	DnsCustomEndpoint terra.StringValue `hcl:"dns_custom_endpoint,attr"`
	// DocumentAiCustomEndpoint: string, optional
	DocumentAiCustomEndpoint terra.StringValue `hcl:"document_ai_custom_endpoint,attr"`
	// EssentialContactsCustomEndpoint: string, optional
	EssentialContactsCustomEndpoint terra.StringValue `hcl:"essential_contacts_custom_endpoint,attr"`
	// EventarcCustomEndpoint: string, optional
	EventarcCustomEndpoint terra.StringValue `hcl:"eventarc_custom_endpoint,attr"`
	// FilestoreCustomEndpoint: string, optional
	FilestoreCustomEndpoint terra.StringValue `hcl:"filestore_custom_endpoint,attr"`
	// FirebaserulesCustomEndpoint: string, optional
	FirebaserulesCustomEndpoint terra.StringValue `hcl:"firebaserules_custom_endpoint,attr"`
	// FirestoreCustomEndpoint: string, optional
	FirestoreCustomEndpoint terra.StringValue `hcl:"firestore_custom_endpoint,attr"`
	// GameServicesCustomEndpoint: string, optional
	GameServicesCustomEndpoint terra.StringValue `hcl:"game_services_custom_endpoint,attr"`
	// GkeBackupCustomEndpoint: string, optional
	GkeBackupCustomEndpoint terra.StringValue `hcl:"gke_backup_custom_endpoint,attr"`
	// GkeHubCustomEndpoint: string, optional
	GkeHubCustomEndpoint terra.StringValue `hcl:"gke_hub_custom_endpoint,attr"`
	// HealthcareCustomEndpoint: string, optional
	HealthcareCustomEndpoint terra.StringValue `hcl:"healthcare_custom_endpoint,attr"`
	// Iam2CustomEndpoint: string, optional
	Iam2CustomEndpoint terra.StringValue `hcl:"iam2_custom_endpoint,attr"`
	// IamBetaCustomEndpoint: string, optional
	IamBetaCustomEndpoint terra.StringValue `hcl:"iam_beta_custom_endpoint,attr"`
	// IamCredentialsCustomEndpoint: string, optional
	IamCredentialsCustomEndpoint terra.StringValue `hcl:"iam_credentials_custom_endpoint,attr"`
	// IamCustomEndpoint: string, optional
	IamCustomEndpoint terra.StringValue `hcl:"iam_custom_endpoint,attr"`
	// IamWorkforcePoolCustomEndpoint: string, optional
	IamWorkforcePoolCustomEndpoint terra.StringValue `hcl:"iam_workforce_pool_custom_endpoint,attr"`
	// IapCustomEndpoint: string, optional
	IapCustomEndpoint terra.StringValue `hcl:"iap_custom_endpoint,attr"`
	// IdentityPlatformCustomEndpoint: string, optional
	IdentityPlatformCustomEndpoint terra.StringValue `hcl:"identity_platform_custom_endpoint,attr"`
	// ImpersonateServiceAccount: string, optional
	ImpersonateServiceAccount terra.StringValue `hcl:"impersonate_service_account,attr"`
	// ImpersonateServiceAccountDelegates: list of string, optional
	ImpersonateServiceAccountDelegates terra.ListValue[terra.StringValue] `hcl:"impersonate_service_account_delegates,attr"`
	// KmsCustomEndpoint: string, optional
	KmsCustomEndpoint terra.StringValue `hcl:"kms_custom_endpoint,attr"`
	// LoggingCustomEndpoint: string, optional
	LoggingCustomEndpoint terra.StringValue `hcl:"logging_custom_endpoint,attr"`
	// MemcacheCustomEndpoint: string, optional
	MemcacheCustomEndpoint terra.StringValue `hcl:"memcache_custom_endpoint,attr"`
	// MlEngineCustomEndpoint: string, optional
	MlEngineCustomEndpoint terra.StringValue `hcl:"ml_engine_custom_endpoint,attr"`
	// MonitoringCustomEndpoint: string, optional
	MonitoringCustomEndpoint terra.StringValue `hcl:"monitoring_custom_endpoint,attr"`
	// NetworkConnectivityCustomEndpoint: string, optional
	NetworkConnectivityCustomEndpoint terra.StringValue `hcl:"network_connectivity_custom_endpoint,attr"`
	// NetworkManagementCustomEndpoint: string, optional
	NetworkManagementCustomEndpoint terra.StringValue `hcl:"network_management_custom_endpoint,attr"`
	// NetworkServicesCustomEndpoint: string, optional
	NetworkServicesCustomEndpoint terra.StringValue `hcl:"network_services_custom_endpoint,attr"`
	// NotebooksCustomEndpoint: string, optional
	NotebooksCustomEndpoint terra.StringValue `hcl:"notebooks_custom_endpoint,attr"`
	// OrgPolicyCustomEndpoint: string, optional
	OrgPolicyCustomEndpoint terra.StringValue `hcl:"org_policy_custom_endpoint,attr"`
	// OsConfigCustomEndpoint: string, optional
	OsConfigCustomEndpoint terra.StringValue `hcl:"os_config_custom_endpoint,attr"`
	// OsLoginCustomEndpoint: string, optional
	OsLoginCustomEndpoint terra.StringValue `hcl:"os_login_custom_endpoint,attr"`
	// PrivatecaCustomEndpoint: string, optional
	PrivatecaCustomEndpoint terra.StringValue `hcl:"privateca_custom_endpoint,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// PubsubCustomEndpoint: string, optional
	PubsubCustomEndpoint terra.StringValue `hcl:"pubsub_custom_endpoint,attr"`
	// PubsubLiteCustomEndpoint: string, optional
	PubsubLiteCustomEndpoint terra.StringValue `hcl:"pubsub_lite_custom_endpoint,attr"`
	// RecaptchaEnterpriseCustomEndpoint: string, optional
	RecaptchaEnterpriseCustomEndpoint terra.StringValue `hcl:"recaptcha_enterprise_custom_endpoint,attr"`
	// RedisCustomEndpoint: string, optional
	RedisCustomEndpoint terra.StringValue `hcl:"redis_custom_endpoint,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// RequestReason: string, optional
	RequestReason terra.StringValue `hcl:"request_reason,attr"`
	// RequestTimeout: string, optional
	RequestTimeout terra.StringValue `hcl:"request_timeout,attr"`
	// ResourceManagerCustomEndpoint: string, optional
	ResourceManagerCustomEndpoint terra.StringValue `hcl:"resource_manager_custom_endpoint,attr"`
	// ResourceManagerV3CustomEndpoint: string, optional
	ResourceManagerV3CustomEndpoint terra.StringValue `hcl:"resource_manager_v3_custom_endpoint,attr"`
	// Scopes: list of string, optional
	Scopes terra.ListValue[terra.StringValue] `hcl:"scopes,attr"`
	// SecretManagerCustomEndpoint: string, optional
	SecretManagerCustomEndpoint terra.StringValue `hcl:"secret_manager_custom_endpoint,attr"`
	// SecurityCenterCustomEndpoint: string, optional
	SecurityCenterCustomEndpoint terra.StringValue `hcl:"security_center_custom_endpoint,attr"`
	// ServiceManagementCustomEndpoint: string, optional
	ServiceManagementCustomEndpoint terra.StringValue `hcl:"service_management_custom_endpoint,attr"`
	// ServiceNetworkingCustomEndpoint: string, optional
	ServiceNetworkingCustomEndpoint terra.StringValue `hcl:"service_networking_custom_endpoint,attr"`
	// ServiceUsageCustomEndpoint: string, optional
	ServiceUsageCustomEndpoint terra.StringValue `hcl:"service_usage_custom_endpoint,attr"`
	// SourceRepoCustomEndpoint: string, optional
	SourceRepoCustomEndpoint terra.StringValue `hcl:"source_repo_custom_endpoint,attr"`
	// SpannerCustomEndpoint: string, optional
	SpannerCustomEndpoint terra.StringValue `hcl:"spanner_custom_endpoint,attr"`
	// SqlCustomEndpoint: string, optional
	SqlCustomEndpoint terra.StringValue `hcl:"sql_custom_endpoint,attr"`
	// StorageCustomEndpoint: string, optional
	StorageCustomEndpoint terra.StringValue `hcl:"storage_custom_endpoint,attr"`
	// StorageTransferCustomEndpoint: string, optional
	StorageTransferCustomEndpoint terra.StringValue `hcl:"storage_transfer_custom_endpoint,attr"`
	// TagsCustomEndpoint: string, optional
	TagsCustomEndpoint terra.StringValue `hcl:"tags_custom_endpoint,attr"`
	// TagsLocationCustomEndpoint: string, optional
	TagsLocationCustomEndpoint terra.StringValue `hcl:"tags_location_custom_endpoint,attr"`
	// TpuCustomEndpoint: string, optional
	TpuCustomEndpoint terra.StringValue `hcl:"tpu_custom_endpoint,attr"`
	// UserProjectOverride: bool, optional
	UserProjectOverride terra.BoolValue `hcl:"user_project_override,attr"`
	// VertexAiCustomEndpoint: string, optional
	VertexAiCustomEndpoint terra.StringValue `hcl:"vertex_ai_custom_endpoint,attr"`
	// VpcAccessCustomEndpoint: string, optional
	VpcAccessCustomEndpoint terra.StringValue `hcl:"vpc_access_custom_endpoint,attr"`
	// WorkflowsCustomEndpoint: string, optional
	WorkflowsCustomEndpoint terra.StringValue `hcl:"workflows_custom_endpoint,attr"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// Batching: optional
	Batching *provider.Batching `hcl:"batching,block"`
}
