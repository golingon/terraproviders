// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	kubernetescluster "github.com/golingon/terraproviders/azurerm/3.68.0/kubernetescluster"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKubernetesCluster creates a new instance of [KubernetesCluster].
func NewKubernetesCluster(name string, args KubernetesClusterArgs) *KubernetesCluster {
	return &KubernetesCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KubernetesCluster)(nil)

// KubernetesCluster represents the Terraform resource azurerm_kubernetes_cluster.
type KubernetesCluster struct {
	Name      string
	Args      KubernetesClusterArgs
	state     *kubernetesClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KubernetesCluster].
func (kc *KubernetesCluster) Type() string {
	return "azurerm_kubernetes_cluster"
}

// LocalName returns the local name for [KubernetesCluster].
func (kc *KubernetesCluster) LocalName() string {
	return kc.Name
}

// Configuration returns the configuration (args) for [KubernetesCluster].
func (kc *KubernetesCluster) Configuration() interface{} {
	return kc.Args
}

// DependOn is used for other resources to depend on [KubernetesCluster].
func (kc *KubernetesCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(kc)
}

// Dependencies returns the list of resources [KubernetesCluster] depends_on.
func (kc *KubernetesCluster) Dependencies() terra.Dependencies {
	return kc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KubernetesCluster].
func (kc *KubernetesCluster) LifecycleManagement() *terra.Lifecycle {
	return kc.Lifecycle
}

// Attributes returns the attributes for [KubernetesCluster].
func (kc *KubernetesCluster) Attributes() kubernetesClusterAttributes {
	return kubernetesClusterAttributes{ref: terra.ReferenceResource(kc)}
}

// ImportState imports the given attribute values into [KubernetesCluster]'s state.
func (kc *KubernetesCluster) ImportState(av io.Reader) error {
	kc.state = &kubernetesClusterState{}
	if err := json.NewDecoder(av).Decode(kc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kc.Type(), kc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KubernetesCluster] has state.
func (kc *KubernetesCluster) State() (*kubernetesClusterState, bool) {
	return kc.state, kc.state != nil
}

// StateMust returns the state for [KubernetesCluster]. Panics if the state is nil.
func (kc *KubernetesCluster) StateMust() *kubernetesClusterState {
	if kc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kc.Type(), kc.LocalName()))
	}
	return kc.state
}

// KubernetesClusterArgs contains the configurations for azurerm_kubernetes_cluster.
type KubernetesClusterArgs struct {
	// ApiServerAuthorizedIpRanges: set of string, optional
	ApiServerAuthorizedIpRanges terra.SetValue[terra.StringValue] `hcl:"api_server_authorized_ip_ranges,attr"`
	// AutomaticChannelUpgrade: string, optional
	AutomaticChannelUpgrade terra.StringValue `hcl:"automatic_channel_upgrade,attr"`
	// AzurePolicyEnabled: bool, optional
	AzurePolicyEnabled terra.BoolValue `hcl:"azure_policy_enabled,attr"`
	// CustomCaTrustCertificatesBase64: list of string, optional
	CustomCaTrustCertificatesBase64 terra.ListValue[terra.StringValue] `hcl:"custom_ca_trust_certificates_base64,attr"`
	// DiskEncryptionSetId: string, optional
	DiskEncryptionSetId terra.StringValue `hcl:"disk_encryption_set_id,attr"`
	// DnsPrefix: string, optional
	DnsPrefix terra.StringValue `hcl:"dns_prefix,attr"`
	// DnsPrefixPrivateCluster: string, optional
	DnsPrefixPrivateCluster terra.StringValue `hcl:"dns_prefix_private_cluster,attr"`
	// EdgeZone: string, optional
	EdgeZone terra.StringValue `hcl:"edge_zone,attr"`
	// EnablePodSecurityPolicy: bool, optional
	EnablePodSecurityPolicy terra.BoolValue `hcl:"enable_pod_security_policy,attr"`
	// HttpApplicationRoutingEnabled: bool, optional
	HttpApplicationRoutingEnabled terra.BoolValue `hcl:"http_application_routing_enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ImageCleanerEnabled: bool, optional
	ImageCleanerEnabled terra.BoolValue `hcl:"image_cleaner_enabled,attr"`
	// ImageCleanerIntervalHours: number, optional
	ImageCleanerIntervalHours terra.NumberValue `hcl:"image_cleaner_interval_hours,attr"`
	// KubernetesVersion: string, optional
	KubernetesVersion terra.StringValue `hcl:"kubernetes_version,attr"`
	// LocalAccountDisabled: bool, optional
	LocalAccountDisabled terra.BoolValue `hcl:"local_account_disabled,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NodeOsChannelUpgrade: string, optional
	NodeOsChannelUpgrade terra.StringValue `hcl:"node_os_channel_upgrade,attr"`
	// NodeResourceGroup: string, optional
	NodeResourceGroup terra.StringValue `hcl:"node_resource_group,attr"`
	// OidcIssuerEnabled: bool, optional
	OidcIssuerEnabled terra.BoolValue `hcl:"oidc_issuer_enabled,attr"`
	// OpenServiceMeshEnabled: bool, optional
	OpenServiceMeshEnabled terra.BoolValue `hcl:"open_service_mesh_enabled,attr"`
	// PrivateClusterEnabled: bool, optional
	PrivateClusterEnabled terra.BoolValue `hcl:"private_cluster_enabled,attr"`
	// PrivateClusterPublicFqdnEnabled: bool, optional
	PrivateClusterPublicFqdnEnabled terra.BoolValue `hcl:"private_cluster_public_fqdn_enabled,attr"`
	// PrivateDnsZoneId: string, optional
	PrivateDnsZoneId terra.StringValue `hcl:"private_dns_zone_id,attr"`
	// PublicNetworkAccessEnabled: bool, optional
	PublicNetworkAccessEnabled terra.BoolValue `hcl:"public_network_access_enabled,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// RoleBasedAccessControlEnabled: bool, optional
	RoleBasedAccessControlEnabled terra.BoolValue `hcl:"role_based_access_control_enabled,attr"`
	// RunCommandEnabled: bool, optional
	RunCommandEnabled terra.BoolValue `hcl:"run_command_enabled,attr"`
	// SkuTier: string, optional
	SkuTier terra.StringValue `hcl:"sku_tier,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// WorkloadIdentityEnabled: bool, optional
	WorkloadIdentityEnabled terra.BoolValue `hcl:"workload_identity_enabled,attr"`
	// KubeAdminConfig: min=0
	KubeAdminConfig []kubernetescluster.KubeAdminConfig `hcl:"kube_admin_config,block" validate:"min=0"`
	// KubeConfig: min=0
	KubeConfig []kubernetescluster.KubeConfig `hcl:"kube_config,block" validate:"min=0"`
	// AciConnectorLinux: optional
	AciConnectorLinux *kubernetescluster.AciConnectorLinux `hcl:"aci_connector_linux,block"`
	// ApiServerAccessProfile: optional
	ApiServerAccessProfile *kubernetescluster.ApiServerAccessProfile `hcl:"api_server_access_profile,block"`
	// AutoScalerProfile: optional
	AutoScalerProfile *kubernetescluster.AutoScalerProfile `hcl:"auto_scaler_profile,block"`
	// AzureActiveDirectoryRoleBasedAccessControl: optional
	AzureActiveDirectoryRoleBasedAccessControl *kubernetescluster.AzureActiveDirectoryRoleBasedAccessControl `hcl:"azure_active_directory_role_based_access_control,block"`
	// ConfidentialComputing: optional
	ConfidentialComputing *kubernetescluster.ConfidentialComputing `hcl:"confidential_computing,block"`
	// DefaultNodePool: required
	DefaultNodePool *kubernetescluster.DefaultNodePool `hcl:"default_node_pool,block" validate:"required"`
	// HttpProxyConfig: optional
	HttpProxyConfig *kubernetescluster.HttpProxyConfig `hcl:"http_proxy_config,block"`
	// Identity: optional
	Identity *kubernetescluster.Identity `hcl:"identity,block"`
	// IngressApplicationGateway: optional
	IngressApplicationGateway *kubernetescluster.IngressApplicationGateway `hcl:"ingress_application_gateway,block"`
	// KeyManagementService: optional
	KeyManagementService *kubernetescluster.KeyManagementService `hcl:"key_management_service,block"`
	// KeyVaultSecretsProvider: optional
	KeyVaultSecretsProvider *kubernetescluster.KeyVaultSecretsProvider `hcl:"key_vault_secrets_provider,block"`
	// KubeletIdentity: optional
	KubeletIdentity *kubernetescluster.KubeletIdentity `hcl:"kubelet_identity,block"`
	// LinuxProfile: optional
	LinuxProfile *kubernetescluster.LinuxProfile `hcl:"linux_profile,block"`
	// MaintenanceWindow: optional
	MaintenanceWindow *kubernetescluster.MaintenanceWindow `hcl:"maintenance_window,block"`
	// MaintenanceWindowAutoUpgrade: optional
	MaintenanceWindowAutoUpgrade *kubernetescluster.MaintenanceWindowAutoUpgrade `hcl:"maintenance_window_auto_upgrade,block"`
	// MaintenanceWindowNodeOs: optional
	MaintenanceWindowNodeOs *kubernetescluster.MaintenanceWindowNodeOs `hcl:"maintenance_window_node_os,block"`
	// MicrosoftDefender: optional
	MicrosoftDefender *kubernetescluster.MicrosoftDefender `hcl:"microsoft_defender,block"`
	// MonitorMetrics: optional
	MonitorMetrics *kubernetescluster.MonitorMetrics `hcl:"monitor_metrics,block"`
	// NetworkProfile: optional
	NetworkProfile *kubernetescluster.NetworkProfile `hcl:"network_profile,block"`
	// OmsAgent: optional
	OmsAgent *kubernetescluster.OmsAgent `hcl:"oms_agent,block"`
	// ServiceMeshProfile: optional
	ServiceMeshProfile *kubernetescluster.ServiceMeshProfile `hcl:"service_mesh_profile,block"`
	// ServicePrincipal: optional
	ServicePrincipal *kubernetescluster.ServicePrincipal `hcl:"service_principal,block"`
	// StorageProfile: optional
	StorageProfile *kubernetescluster.StorageProfile `hcl:"storage_profile,block"`
	// Timeouts: optional
	Timeouts *kubernetescluster.Timeouts `hcl:"timeouts,block"`
	// WebAppRouting: optional
	WebAppRouting *kubernetescluster.WebAppRouting `hcl:"web_app_routing,block"`
	// WindowsProfile: optional
	WindowsProfile *kubernetescluster.WindowsProfile `hcl:"windows_profile,block"`
	// WorkloadAutoscalerProfile: optional
	WorkloadAutoscalerProfile *kubernetescluster.WorkloadAutoscalerProfile `hcl:"workload_autoscaler_profile,block"`
}
type kubernetesClusterAttributes struct {
	ref terra.Reference
}

// ApiServerAuthorizedIpRanges returns a reference to field api_server_authorized_ip_ranges of azurerm_kubernetes_cluster.
func (kc kubernetesClusterAttributes) ApiServerAuthorizedIpRanges() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](kc.ref.Append("api_server_authorized_ip_ranges"))
}

// AutomaticChannelUpgrade returns a reference to field automatic_channel_upgrade of azurerm_kubernetes_cluster.
func (kc kubernetesClusterAttributes) AutomaticChannelUpgrade() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("automatic_channel_upgrade"))
}

// AzurePolicyEnabled returns a reference to field azure_policy_enabled of azurerm_kubernetes_cluster.
func (kc kubernetesClusterAttributes) AzurePolicyEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(kc.ref.Append("azure_policy_enabled"))
}

// CustomCaTrustCertificatesBase64 returns a reference to field custom_ca_trust_certificates_base64 of azurerm_kubernetes_cluster.
func (kc kubernetesClusterAttributes) CustomCaTrustCertificatesBase64() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](kc.ref.Append("custom_ca_trust_certificates_base64"))
}

// DiskEncryptionSetId returns a reference to field disk_encryption_set_id of azurerm_kubernetes_cluster.
func (kc kubernetesClusterAttributes) DiskEncryptionSetId() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("disk_encryption_set_id"))
}

// DnsPrefix returns a reference to field dns_prefix of azurerm_kubernetes_cluster.
func (kc kubernetesClusterAttributes) DnsPrefix() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("dns_prefix"))
}

// DnsPrefixPrivateCluster returns a reference to field dns_prefix_private_cluster of azurerm_kubernetes_cluster.
func (kc kubernetesClusterAttributes) DnsPrefixPrivateCluster() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("dns_prefix_private_cluster"))
}

// EdgeZone returns a reference to field edge_zone of azurerm_kubernetes_cluster.
func (kc kubernetesClusterAttributes) EdgeZone() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("edge_zone"))
}

// EnablePodSecurityPolicy returns a reference to field enable_pod_security_policy of azurerm_kubernetes_cluster.
func (kc kubernetesClusterAttributes) EnablePodSecurityPolicy() terra.BoolValue {
	return terra.ReferenceAsBool(kc.ref.Append("enable_pod_security_policy"))
}

// Fqdn returns a reference to field fqdn of azurerm_kubernetes_cluster.
func (kc kubernetesClusterAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("fqdn"))
}

// HttpApplicationRoutingEnabled returns a reference to field http_application_routing_enabled of azurerm_kubernetes_cluster.
func (kc kubernetesClusterAttributes) HttpApplicationRoutingEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(kc.ref.Append("http_application_routing_enabled"))
}

// HttpApplicationRoutingZoneName returns a reference to field http_application_routing_zone_name of azurerm_kubernetes_cluster.
func (kc kubernetesClusterAttributes) HttpApplicationRoutingZoneName() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("http_application_routing_zone_name"))
}

// Id returns a reference to field id of azurerm_kubernetes_cluster.
func (kc kubernetesClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("id"))
}

// ImageCleanerEnabled returns a reference to field image_cleaner_enabled of azurerm_kubernetes_cluster.
func (kc kubernetesClusterAttributes) ImageCleanerEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(kc.ref.Append("image_cleaner_enabled"))
}

// ImageCleanerIntervalHours returns a reference to field image_cleaner_interval_hours of azurerm_kubernetes_cluster.
func (kc kubernetesClusterAttributes) ImageCleanerIntervalHours() terra.NumberValue {
	return terra.ReferenceAsNumber(kc.ref.Append("image_cleaner_interval_hours"))
}

// KubeAdminConfigRaw returns a reference to field kube_admin_config_raw of azurerm_kubernetes_cluster.
func (kc kubernetesClusterAttributes) KubeAdminConfigRaw() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("kube_admin_config_raw"))
}

// KubeConfigRaw returns a reference to field kube_config_raw of azurerm_kubernetes_cluster.
func (kc kubernetesClusterAttributes) KubeConfigRaw() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("kube_config_raw"))
}

// KubernetesVersion returns a reference to field kubernetes_version of azurerm_kubernetes_cluster.
func (kc kubernetesClusterAttributes) KubernetesVersion() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("kubernetes_version"))
}

// LocalAccountDisabled returns a reference to field local_account_disabled of azurerm_kubernetes_cluster.
func (kc kubernetesClusterAttributes) LocalAccountDisabled() terra.BoolValue {
	return terra.ReferenceAsBool(kc.ref.Append("local_account_disabled"))
}

// Location returns a reference to field location of azurerm_kubernetes_cluster.
func (kc kubernetesClusterAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_kubernetes_cluster.
func (kc kubernetesClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("name"))
}

// NodeOsChannelUpgrade returns a reference to field node_os_channel_upgrade of azurerm_kubernetes_cluster.
func (kc kubernetesClusterAttributes) NodeOsChannelUpgrade() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("node_os_channel_upgrade"))
}

// NodeResourceGroup returns a reference to field node_resource_group of azurerm_kubernetes_cluster.
func (kc kubernetesClusterAttributes) NodeResourceGroup() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("node_resource_group"))
}

// NodeResourceGroupId returns a reference to field node_resource_group_id of azurerm_kubernetes_cluster.
func (kc kubernetesClusterAttributes) NodeResourceGroupId() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("node_resource_group_id"))
}

// OidcIssuerEnabled returns a reference to field oidc_issuer_enabled of azurerm_kubernetes_cluster.
func (kc kubernetesClusterAttributes) OidcIssuerEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(kc.ref.Append("oidc_issuer_enabled"))
}

// OidcIssuerUrl returns a reference to field oidc_issuer_url of azurerm_kubernetes_cluster.
func (kc kubernetesClusterAttributes) OidcIssuerUrl() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("oidc_issuer_url"))
}

// OpenServiceMeshEnabled returns a reference to field open_service_mesh_enabled of azurerm_kubernetes_cluster.
func (kc kubernetesClusterAttributes) OpenServiceMeshEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(kc.ref.Append("open_service_mesh_enabled"))
}

// PortalFqdn returns a reference to field portal_fqdn of azurerm_kubernetes_cluster.
func (kc kubernetesClusterAttributes) PortalFqdn() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("portal_fqdn"))
}

// PrivateClusterEnabled returns a reference to field private_cluster_enabled of azurerm_kubernetes_cluster.
func (kc kubernetesClusterAttributes) PrivateClusterEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(kc.ref.Append("private_cluster_enabled"))
}

// PrivateClusterPublicFqdnEnabled returns a reference to field private_cluster_public_fqdn_enabled of azurerm_kubernetes_cluster.
func (kc kubernetesClusterAttributes) PrivateClusterPublicFqdnEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(kc.ref.Append("private_cluster_public_fqdn_enabled"))
}

// PrivateDnsZoneId returns a reference to field private_dns_zone_id of azurerm_kubernetes_cluster.
func (kc kubernetesClusterAttributes) PrivateDnsZoneId() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("private_dns_zone_id"))
}

// PrivateFqdn returns a reference to field private_fqdn of azurerm_kubernetes_cluster.
func (kc kubernetesClusterAttributes) PrivateFqdn() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("private_fqdn"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_kubernetes_cluster.
func (kc kubernetesClusterAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(kc.ref.Append("public_network_access_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_kubernetes_cluster.
func (kc kubernetesClusterAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("resource_group_name"))
}

// RoleBasedAccessControlEnabled returns a reference to field role_based_access_control_enabled of azurerm_kubernetes_cluster.
func (kc kubernetesClusterAttributes) RoleBasedAccessControlEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(kc.ref.Append("role_based_access_control_enabled"))
}

// RunCommandEnabled returns a reference to field run_command_enabled of azurerm_kubernetes_cluster.
func (kc kubernetesClusterAttributes) RunCommandEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(kc.ref.Append("run_command_enabled"))
}

// SkuTier returns a reference to field sku_tier of azurerm_kubernetes_cluster.
func (kc kubernetesClusterAttributes) SkuTier() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("sku_tier"))
}

// Tags returns a reference to field tags of azurerm_kubernetes_cluster.
func (kc kubernetesClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](kc.ref.Append("tags"))
}

// WorkloadIdentityEnabled returns a reference to field workload_identity_enabled of azurerm_kubernetes_cluster.
func (kc kubernetesClusterAttributes) WorkloadIdentityEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(kc.ref.Append("workload_identity_enabled"))
}

func (kc kubernetesClusterAttributes) KubeAdminConfig() terra.ListValue[kubernetescluster.KubeAdminConfigAttributes] {
	return terra.ReferenceAsList[kubernetescluster.KubeAdminConfigAttributes](kc.ref.Append("kube_admin_config"))
}

func (kc kubernetesClusterAttributes) KubeConfig() terra.ListValue[kubernetescluster.KubeConfigAttributes] {
	return terra.ReferenceAsList[kubernetescluster.KubeConfigAttributes](kc.ref.Append("kube_config"))
}

func (kc kubernetesClusterAttributes) AciConnectorLinux() terra.ListValue[kubernetescluster.AciConnectorLinuxAttributes] {
	return terra.ReferenceAsList[kubernetescluster.AciConnectorLinuxAttributes](kc.ref.Append("aci_connector_linux"))
}

func (kc kubernetesClusterAttributes) ApiServerAccessProfile() terra.ListValue[kubernetescluster.ApiServerAccessProfileAttributes] {
	return terra.ReferenceAsList[kubernetescluster.ApiServerAccessProfileAttributes](kc.ref.Append("api_server_access_profile"))
}

func (kc kubernetesClusterAttributes) AutoScalerProfile() terra.ListValue[kubernetescluster.AutoScalerProfileAttributes] {
	return terra.ReferenceAsList[kubernetescluster.AutoScalerProfileAttributes](kc.ref.Append("auto_scaler_profile"))
}

func (kc kubernetesClusterAttributes) AzureActiveDirectoryRoleBasedAccessControl() terra.ListValue[kubernetescluster.AzureActiveDirectoryRoleBasedAccessControlAttributes] {
	return terra.ReferenceAsList[kubernetescluster.AzureActiveDirectoryRoleBasedAccessControlAttributes](kc.ref.Append("azure_active_directory_role_based_access_control"))
}

func (kc kubernetesClusterAttributes) ConfidentialComputing() terra.ListValue[kubernetescluster.ConfidentialComputingAttributes] {
	return terra.ReferenceAsList[kubernetescluster.ConfidentialComputingAttributes](kc.ref.Append("confidential_computing"))
}

func (kc kubernetesClusterAttributes) DefaultNodePool() terra.ListValue[kubernetescluster.DefaultNodePoolAttributes] {
	return terra.ReferenceAsList[kubernetescluster.DefaultNodePoolAttributes](kc.ref.Append("default_node_pool"))
}

func (kc kubernetesClusterAttributes) HttpProxyConfig() terra.ListValue[kubernetescluster.HttpProxyConfigAttributes] {
	return terra.ReferenceAsList[kubernetescluster.HttpProxyConfigAttributes](kc.ref.Append("http_proxy_config"))
}

func (kc kubernetesClusterAttributes) Identity() terra.ListValue[kubernetescluster.IdentityAttributes] {
	return terra.ReferenceAsList[kubernetescluster.IdentityAttributes](kc.ref.Append("identity"))
}

func (kc kubernetesClusterAttributes) IngressApplicationGateway() terra.ListValue[kubernetescluster.IngressApplicationGatewayAttributes] {
	return terra.ReferenceAsList[kubernetescluster.IngressApplicationGatewayAttributes](kc.ref.Append("ingress_application_gateway"))
}

func (kc kubernetesClusterAttributes) KeyManagementService() terra.ListValue[kubernetescluster.KeyManagementServiceAttributes] {
	return terra.ReferenceAsList[kubernetescluster.KeyManagementServiceAttributes](kc.ref.Append("key_management_service"))
}

func (kc kubernetesClusterAttributes) KeyVaultSecretsProvider() terra.ListValue[kubernetescluster.KeyVaultSecretsProviderAttributes] {
	return terra.ReferenceAsList[kubernetescluster.KeyVaultSecretsProviderAttributes](kc.ref.Append("key_vault_secrets_provider"))
}

func (kc kubernetesClusterAttributes) KubeletIdentity() terra.ListValue[kubernetescluster.KubeletIdentityAttributes] {
	return terra.ReferenceAsList[kubernetescluster.KubeletIdentityAttributes](kc.ref.Append("kubelet_identity"))
}

func (kc kubernetesClusterAttributes) LinuxProfile() terra.ListValue[kubernetescluster.LinuxProfileAttributes] {
	return terra.ReferenceAsList[kubernetescluster.LinuxProfileAttributes](kc.ref.Append("linux_profile"))
}

func (kc kubernetesClusterAttributes) MaintenanceWindow() terra.ListValue[kubernetescluster.MaintenanceWindowAttributes] {
	return terra.ReferenceAsList[kubernetescluster.MaintenanceWindowAttributes](kc.ref.Append("maintenance_window"))
}

func (kc kubernetesClusterAttributes) MaintenanceWindowAutoUpgrade() terra.ListValue[kubernetescluster.MaintenanceWindowAutoUpgradeAttributes] {
	return terra.ReferenceAsList[kubernetescluster.MaintenanceWindowAutoUpgradeAttributes](kc.ref.Append("maintenance_window_auto_upgrade"))
}

func (kc kubernetesClusterAttributes) MaintenanceWindowNodeOs() terra.ListValue[kubernetescluster.MaintenanceWindowNodeOsAttributes] {
	return terra.ReferenceAsList[kubernetescluster.MaintenanceWindowNodeOsAttributes](kc.ref.Append("maintenance_window_node_os"))
}

func (kc kubernetesClusterAttributes) MicrosoftDefender() terra.ListValue[kubernetescluster.MicrosoftDefenderAttributes] {
	return terra.ReferenceAsList[kubernetescluster.MicrosoftDefenderAttributes](kc.ref.Append("microsoft_defender"))
}

func (kc kubernetesClusterAttributes) MonitorMetrics() terra.ListValue[kubernetescluster.MonitorMetricsAttributes] {
	return terra.ReferenceAsList[kubernetescluster.MonitorMetricsAttributes](kc.ref.Append("monitor_metrics"))
}

func (kc kubernetesClusterAttributes) NetworkProfile() terra.ListValue[kubernetescluster.NetworkProfileAttributes] {
	return terra.ReferenceAsList[kubernetescluster.NetworkProfileAttributes](kc.ref.Append("network_profile"))
}

func (kc kubernetesClusterAttributes) OmsAgent() terra.ListValue[kubernetescluster.OmsAgentAttributes] {
	return terra.ReferenceAsList[kubernetescluster.OmsAgentAttributes](kc.ref.Append("oms_agent"))
}

func (kc kubernetesClusterAttributes) ServiceMeshProfile() terra.ListValue[kubernetescluster.ServiceMeshProfileAttributes] {
	return terra.ReferenceAsList[kubernetescluster.ServiceMeshProfileAttributes](kc.ref.Append("service_mesh_profile"))
}

func (kc kubernetesClusterAttributes) ServicePrincipal() terra.ListValue[kubernetescluster.ServicePrincipalAttributes] {
	return terra.ReferenceAsList[kubernetescluster.ServicePrincipalAttributes](kc.ref.Append("service_principal"))
}

func (kc kubernetesClusterAttributes) StorageProfile() terra.ListValue[kubernetescluster.StorageProfileAttributes] {
	return terra.ReferenceAsList[kubernetescluster.StorageProfileAttributes](kc.ref.Append("storage_profile"))
}

func (kc kubernetesClusterAttributes) Timeouts() kubernetescluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[kubernetescluster.TimeoutsAttributes](kc.ref.Append("timeouts"))
}

func (kc kubernetesClusterAttributes) WebAppRouting() terra.ListValue[kubernetescluster.WebAppRoutingAttributes] {
	return terra.ReferenceAsList[kubernetescluster.WebAppRoutingAttributes](kc.ref.Append("web_app_routing"))
}

func (kc kubernetesClusterAttributes) WindowsProfile() terra.ListValue[kubernetescluster.WindowsProfileAttributes] {
	return terra.ReferenceAsList[kubernetescluster.WindowsProfileAttributes](kc.ref.Append("windows_profile"))
}

func (kc kubernetesClusterAttributes) WorkloadAutoscalerProfile() terra.ListValue[kubernetescluster.WorkloadAutoscalerProfileAttributes] {
	return terra.ReferenceAsList[kubernetescluster.WorkloadAutoscalerProfileAttributes](kc.ref.Append("workload_autoscaler_profile"))
}

type kubernetesClusterState struct {
	ApiServerAuthorizedIpRanges                []string                                                            `json:"api_server_authorized_ip_ranges"`
	AutomaticChannelUpgrade                    string                                                              `json:"automatic_channel_upgrade"`
	AzurePolicyEnabled                         bool                                                                `json:"azure_policy_enabled"`
	CustomCaTrustCertificatesBase64            []string                                                            `json:"custom_ca_trust_certificates_base64"`
	DiskEncryptionSetId                        string                                                              `json:"disk_encryption_set_id"`
	DnsPrefix                                  string                                                              `json:"dns_prefix"`
	DnsPrefixPrivateCluster                    string                                                              `json:"dns_prefix_private_cluster"`
	EdgeZone                                   string                                                              `json:"edge_zone"`
	EnablePodSecurityPolicy                    bool                                                                `json:"enable_pod_security_policy"`
	Fqdn                                       string                                                              `json:"fqdn"`
	HttpApplicationRoutingEnabled              bool                                                                `json:"http_application_routing_enabled"`
	HttpApplicationRoutingZoneName             string                                                              `json:"http_application_routing_zone_name"`
	Id                                         string                                                              `json:"id"`
	ImageCleanerEnabled                        bool                                                                `json:"image_cleaner_enabled"`
	ImageCleanerIntervalHours                  float64                                                             `json:"image_cleaner_interval_hours"`
	KubeAdminConfigRaw                         string                                                              `json:"kube_admin_config_raw"`
	KubeConfigRaw                              string                                                              `json:"kube_config_raw"`
	KubernetesVersion                          string                                                              `json:"kubernetes_version"`
	LocalAccountDisabled                       bool                                                                `json:"local_account_disabled"`
	Location                                   string                                                              `json:"location"`
	Name                                       string                                                              `json:"name"`
	NodeOsChannelUpgrade                       string                                                              `json:"node_os_channel_upgrade"`
	NodeResourceGroup                          string                                                              `json:"node_resource_group"`
	NodeResourceGroupId                        string                                                              `json:"node_resource_group_id"`
	OidcIssuerEnabled                          bool                                                                `json:"oidc_issuer_enabled"`
	OidcIssuerUrl                              string                                                              `json:"oidc_issuer_url"`
	OpenServiceMeshEnabled                     bool                                                                `json:"open_service_mesh_enabled"`
	PortalFqdn                                 string                                                              `json:"portal_fqdn"`
	PrivateClusterEnabled                      bool                                                                `json:"private_cluster_enabled"`
	PrivateClusterPublicFqdnEnabled            bool                                                                `json:"private_cluster_public_fqdn_enabled"`
	PrivateDnsZoneId                           string                                                              `json:"private_dns_zone_id"`
	PrivateFqdn                                string                                                              `json:"private_fqdn"`
	PublicNetworkAccessEnabled                 bool                                                                `json:"public_network_access_enabled"`
	ResourceGroupName                          string                                                              `json:"resource_group_name"`
	RoleBasedAccessControlEnabled              bool                                                                `json:"role_based_access_control_enabled"`
	RunCommandEnabled                          bool                                                                `json:"run_command_enabled"`
	SkuTier                                    string                                                              `json:"sku_tier"`
	Tags                                       map[string]string                                                   `json:"tags"`
	WorkloadIdentityEnabled                    bool                                                                `json:"workload_identity_enabled"`
	KubeAdminConfig                            []kubernetescluster.KubeAdminConfigState                            `json:"kube_admin_config"`
	KubeConfig                                 []kubernetescluster.KubeConfigState                                 `json:"kube_config"`
	AciConnectorLinux                          []kubernetescluster.AciConnectorLinuxState                          `json:"aci_connector_linux"`
	ApiServerAccessProfile                     []kubernetescluster.ApiServerAccessProfileState                     `json:"api_server_access_profile"`
	AutoScalerProfile                          []kubernetescluster.AutoScalerProfileState                          `json:"auto_scaler_profile"`
	AzureActiveDirectoryRoleBasedAccessControl []kubernetescluster.AzureActiveDirectoryRoleBasedAccessControlState `json:"azure_active_directory_role_based_access_control"`
	ConfidentialComputing                      []kubernetescluster.ConfidentialComputingState                      `json:"confidential_computing"`
	DefaultNodePool                            []kubernetescluster.DefaultNodePoolState                            `json:"default_node_pool"`
	HttpProxyConfig                            []kubernetescluster.HttpProxyConfigState                            `json:"http_proxy_config"`
	Identity                                   []kubernetescluster.IdentityState                                   `json:"identity"`
	IngressApplicationGateway                  []kubernetescluster.IngressApplicationGatewayState                  `json:"ingress_application_gateway"`
	KeyManagementService                       []kubernetescluster.KeyManagementServiceState                       `json:"key_management_service"`
	KeyVaultSecretsProvider                    []kubernetescluster.KeyVaultSecretsProviderState                    `json:"key_vault_secrets_provider"`
	KubeletIdentity                            []kubernetescluster.KubeletIdentityState                            `json:"kubelet_identity"`
	LinuxProfile                               []kubernetescluster.LinuxProfileState                               `json:"linux_profile"`
	MaintenanceWindow                          []kubernetescluster.MaintenanceWindowState                          `json:"maintenance_window"`
	MaintenanceWindowAutoUpgrade               []kubernetescluster.MaintenanceWindowAutoUpgradeState               `json:"maintenance_window_auto_upgrade"`
	MaintenanceWindowNodeOs                    []kubernetescluster.MaintenanceWindowNodeOsState                    `json:"maintenance_window_node_os"`
	MicrosoftDefender                          []kubernetescluster.MicrosoftDefenderState                          `json:"microsoft_defender"`
	MonitorMetrics                             []kubernetescluster.MonitorMetricsState                             `json:"monitor_metrics"`
	NetworkProfile                             []kubernetescluster.NetworkProfileState                             `json:"network_profile"`
	OmsAgent                                   []kubernetescluster.OmsAgentState                                   `json:"oms_agent"`
	ServiceMeshProfile                         []kubernetescluster.ServiceMeshProfileState                         `json:"service_mesh_profile"`
	ServicePrincipal                           []kubernetescluster.ServicePrincipalState                           `json:"service_principal"`
	StorageProfile                             []kubernetescluster.StorageProfileState                             `json:"storage_profile"`
	Timeouts                                   *kubernetescluster.TimeoutsState                                    `json:"timeouts"`
	WebAppRouting                              []kubernetescluster.WebAppRoutingState                              `json:"web_app_routing"`
	WindowsProfile                             []kubernetescluster.WindowsProfileState                             `json:"windows_profile"`
	WorkloadAutoscalerProfile                  []kubernetescluster.WorkloadAutoscalerProfileState                  `json:"workload_autoscaler_profile"`
}
