// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datakubernetescluster "github.com/golingon/terraproviders/azurerm/3.64.0/datakubernetescluster"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataKubernetesCluster creates a new instance of [DataKubernetesCluster].
func NewDataKubernetesCluster(name string, args DataKubernetesClusterArgs) *DataKubernetesCluster {
	return &DataKubernetesCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataKubernetesCluster)(nil)

// DataKubernetesCluster represents the Terraform data resource azurerm_kubernetes_cluster.
type DataKubernetesCluster struct {
	Name string
	Args DataKubernetesClusterArgs
}

// DataSource returns the Terraform object type for [DataKubernetesCluster].
func (kc *DataKubernetesCluster) DataSource() string {
	return "azurerm_kubernetes_cluster"
}

// LocalName returns the local name for [DataKubernetesCluster].
func (kc *DataKubernetesCluster) LocalName() string {
	return kc.Name
}

// Configuration returns the configuration (args) for [DataKubernetesCluster].
func (kc *DataKubernetesCluster) Configuration() interface{} {
	return kc.Args
}

// Attributes returns the attributes for [DataKubernetesCluster].
func (kc *DataKubernetesCluster) Attributes() dataKubernetesClusterAttributes {
	return dataKubernetesClusterAttributes{ref: terra.ReferenceDataResource(kc)}
}

// DataKubernetesClusterArgs contains the configurations for azurerm_kubernetes_cluster.
type DataKubernetesClusterArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// AciConnectorLinux: min=0
	AciConnectorLinux []datakubernetescluster.AciConnectorLinux `hcl:"aci_connector_linux,block" validate:"min=0"`
	// AgentPoolProfile: min=0
	AgentPoolProfile []datakubernetescluster.AgentPoolProfile `hcl:"agent_pool_profile,block" validate:"min=0"`
	// AzureActiveDirectoryRoleBasedAccessControl: min=0
	AzureActiveDirectoryRoleBasedAccessControl []datakubernetescluster.AzureActiveDirectoryRoleBasedAccessControl `hcl:"azure_active_directory_role_based_access_control,block" validate:"min=0"`
	// Identity: min=0
	Identity []datakubernetescluster.Identity `hcl:"identity,block" validate:"min=0"`
	// IngressApplicationGateway: min=0
	IngressApplicationGateway []datakubernetescluster.IngressApplicationGateway `hcl:"ingress_application_gateway,block" validate:"min=0"`
	// KeyManagementService: min=0
	KeyManagementService []datakubernetescluster.KeyManagementService `hcl:"key_management_service,block" validate:"min=0"`
	// KeyVaultSecretsProvider: min=0
	KeyVaultSecretsProvider []datakubernetescluster.KeyVaultSecretsProvider `hcl:"key_vault_secrets_provider,block" validate:"min=0"`
	// KubeAdminConfig: min=0
	KubeAdminConfig []datakubernetescluster.KubeAdminConfig `hcl:"kube_admin_config,block" validate:"min=0"`
	// KubeConfig: min=0
	KubeConfig []datakubernetescluster.KubeConfig `hcl:"kube_config,block" validate:"min=0"`
	// KubeletIdentity: min=0
	KubeletIdentity []datakubernetescluster.KubeletIdentity `hcl:"kubelet_identity,block" validate:"min=0"`
	// LinuxProfile: min=0
	LinuxProfile []datakubernetescluster.LinuxProfile `hcl:"linux_profile,block" validate:"min=0"`
	// MicrosoftDefender: min=0
	MicrosoftDefender []datakubernetescluster.MicrosoftDefender `hcl:"microsoft_defender,block" validate:"min=0"`
	// NetworkProfile: min=0
	NetworkProfile []datakubernetescluster.NetworkProfile `hcl:"network_profile,block" validate:"min=0"`
	// OmsAgent: min=0
	OmsAgent []datakubernetescluster.OmsAgent `hcl:"oms_agent,block" validate:"min=0"`
	// ServicePrincipal: min=0
	ServicePrincipal []datakubernetescluster.ServicePrincipal `hcl:"service_principal,block" validate:"min=0"`
	// StorageProfile: min=0
	StorageProfile []datakubernetescluster.StorageProfile `hcl:"storage_profile,block" validate:"min=0"`
	// WindowsProfile: min=0
	WindowsProfile []datakubernetescluster.WindowsProfile `hcl:"windows_profile,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datakubernetescluster.Timeouts `hcl:"timeouts,block"`
}
type dataKubernetesClusterAttributes struct {
	ref terra.Reference
}

// ApiServerAuthorizedIpRanges returns a reference to field api_server_authorized_ip_ranges of azurerm_kubernetes_cluster.
func (kc dataKubernetesClusterAttributes) ApiServerAuthorizedIpRanges() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](kc.ref.Append("api_server_authorized_ip_ranges"))
}

// AzurePolicyEnabled returns a reference to field azure_policy_enabled of azurerm_kubernetes_cluster.
func (kc dataKubernetesClusterAttributes) AzurePolicyEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(kc.ref.Append("azure_policy_enabled"))
}

// CustomCaTrustCertificatesBase64 returns a reference to field custom_ca_trust_certificates_base64 of azurerm_kubernetes_cluster.
func (kc dataKubernetesClusterAttributes) CustomCaTrustCertificatesBase64() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](kc.ref.Append("custom_ca_trust_certificates_base64"))
}

// DiskEncryptionSetId returns a reference to field disk_encryption_set_id of azurerm_kubernetes_cluster.
func (kc dataKubernetesClusterAttributes) DiskEncryptionSetId() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("disk_encryption_set_id"))
}

// DnsPrefix returns a reference to field dns_prefix of azurerm_kubernetes_cluster.
func (kc dataKubernetesClusterAttributes) DnsPrefix() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("dns_prefix"))
}

// Fqdn returns a reference to field fqdn of azurerm_kubernetes_cluster.
func (kc dataKubernetesClusterAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("fqdn"))
}

// HttpApplicationRoutingEnabled returns a reference to field http_application_routing_enabled of azurerm_kubernetes_cluster.
func (kc dataKubernetesClusterAttributes) HttpApplicationRoutingEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(kc.ref.Append("http_application_routing_enabled"))
}

// HttpApplicationRoutingZoneName returns a reference to field http_application_routing_zone_name of azurerm_kubernetes_cluster.
func (kc dataKubernetesClusterAttributes) HttpApplicationRoutingZoneName() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("http_application_routing_zone_name"))
}

// Id returns a reference to field id of azurerm_kubernetes_cluster.
func (kc dataKubernetesClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("id"))
}

// KubeAdminConfigRaw returns a reference to field kube_admin_config_raw of azurerm_kubernetes_cluster.
func (kc dataKubernetesClusterAttributes) KubeAdminConfigRaw() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("kube_admin_config_raw"))
}

// KubeConfigRaw returns a reference to field kube_config_raw of azurerm_kubernetes_cluster.
func (kc dataKubernetesClusterAttributes) KubeConfigRaw() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("kube_config_raw"))
}

// KubernetesVersion returns a reference to field kubernetes_version of azurerm_kubernetes_cluster.
func (kc dataKubernetesClusterAttributes) KubernetesVersion() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("kubernetes_version"))
}

// Location returns a reference to field location of azurerm_kubernetes_cluster.
func (kc dataKubernetesClusterAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_kubernetes_cluster.
func (kc dataKubernetesClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("name"))
}

// NodeResourceGroup returns a reference to field node_resource_group of azurerm_kubernetes_cluster.
func (kc dataKubernetesClusterAttributes) NodeResourceGroup() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("node_resource_group"))
}

// NodeResourceGroupId returns a reference to field node_resource_group_id of azurerm_kubernetes_cluster.
func (kc dataKubernetesClusterAttributes) NodeResourceGroupId() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("node_resource_group_id"))
}

// OidcIssuerEnabled returns a reference to field oidc_issuer_enabled of azurerm_kubernetes_cluster.
func (kc dataKubernetesClusterAttributes) OidcIssuerEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(kc.ref.Append("oidc_issuer_enabled"))
}

// OidcIssuerUrl returns a reference to field oidc_issuer_url of azurerm_kubernetes_cluster.
func (kc dataKubernetesClusterAttributes) OidcIssuerUrl() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("oidc_issuer_url"))
}

// OpenServiceMeshEnabled returns a reference to field open_service_mesh_enabled of azurerm_kubernetes_cluster.
func (kc dataKubernetesClusterAttributes) OpenServiceMeshEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(kc.ref.Append("open_service_mesh_enabled"))
}

// PrivateClusterEnabled returns a reference to field private_cluster_enabled of azurerm_kubernetes_cluster.
func (kc dataKubernetesClusterAttributes) PrivateClusterEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(kc.ref.Append("private_cluster_enabled"))
}

// PrivateFqdn returns a reference to field private_fqdn of azurerm_kubernetes_cluster.
func (kc dataKubernetesClusterAttributes) PrivateFqdn() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("private_fqdn"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_kubernetes_cluster.
func (kc dataKubernetesClusterAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("resource_group_name"))
}

// RoleBasedAccessControlEnabled returns a reference to field role_based_access_control_enabled of azurerm_kubernetes_cluster.
func (kc dataKubernetesClusterAttributes) RoleBasedAccessControlEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(kc.ref.Append("role_based_access_control_enabled"))
}

// Tags returns a reference to field tags of azurerm_kubernetes_cluster.
func (kc dataKubernetesClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](kc.ref.Append("tags"))
}

func (kc dataKubernetesClusterAttributes) AciConnectorLinux() terra.ListValue[datakubernetescluster.AciConnectorLinuxAttributes] {
	return terra.ReferenceAsList[datakubernetescluster.AciConnectorLinuxAttributes](kc.ref.Append("aci_connector_linux"))
}

func (kc dataKubernetesClusterAttributes) AgentPoolProfile() terra.ListValue[datakubernetescluster.AgentPoolProfileAttributes] {
	return terra.ReferenceAsList[datakubernetescluster.AgentPoolProfileAttributes](kc.ref.Append("agent_pool_profile"))
}

func (kc dataKubernetesClusterAttributes) AzureActiveDirectoryRoleBasedAccessControl() terra.ListValue[datakubernetescluster.AzureActiveDirectoryRoleBasedAccessControlAttributes] {
	return terra.ReferenceAsList[datakubernetescluster.AzureActiveDirectoryRoleBasedAccessControlAttributes](kc.ref.Append("azure_active_directory_role_based_access_control"))
}

func (kc dataKubernetesClusterAttributes) Identity() terra.ListValue[datakubernetescluster.IdentityAttributes] {
	return terra.ReferenceAsList[datakubernetescluster.IdentityAttributes](kc.ref.Append("identity"))
}

func (kc dataKubernetesClusterAttributes) IngressApplicationGateway() terra.ListValue[datakubernetescluster.IngressApplicationGatewayAttributes] {
	return terra.ReferenceAsList[datakubernetescluster.IngressApplicationGatewayAttributes](kc.ref.Append("ingress_application_gateway"))
}

func (kc dataKubernetesClusterAttributes) KeyManagementService() terra.ListValue[datakubernetescluster.KeyManagementServiceAttributes] {
	return terra.ReferenceAsList[datakubernetescluster.KeyManagementServiceAttributes](kc.ref.Append("key_management_service"))
}

func (kc dataKubernetesClusterAttributes) KeyVaultSecretsProvider() terra.ListValue[datakubernetescluster.KeyVaultSecretsProviderAttributes] {
	return terra.ReferenceAsList[datakubernetescluster.KeyVaultSecretsProviderAttributes](kc.ref.Append("key_vault_secrets_provider"))
}

func (kc dataKubernetesClusterAttributes) KubeAdminConfig() terra.ListValue[datakubernetescluster.KubeAdminConfigAttributes] {
	return terra.ReferenceAsList[datakubernetescluster.KubeAdminConfigAttributes](kc.ref.Append("kube_admin_config"))
}

func (kc dataKubernetesClusterAttributes) KubeConfig() terra.ListValue[datakubernetescluster.KubeConfigAttributes] {
	return terra.ReferenceAsList[datakubernetescluster.KubeConfigAttributes](kc.ref.Append("kube_config"))
}

func (kc dataKubernetesClusterAttributes) KubeletIdentity() terra.ListValue[datakubernetescluster.KubeletIdentityAttributes] {
	return terra.ReferenceAsList[datakubernetescluster.KubeletIdentityAttributes](kc.ref.Append("kubelet_identity"))
}

func (kc dataKubernetesClusterAttributes) LinuxProfile() terra.ListValue[datakubernetescluster.LinuxProfileAttributes] {
	return terra.ReferenceAsList[datakubernetescluster.LinuxProfileAttributes](kc.ref.Append("linux_profile"))
}

func (kc dataKubernetesClusterAttributes) MicrosoftDefender() terra.ListValue[datakubernetescluster.MicrosoftDefenderAttributes] {
	return terra.ReferenceAsList[datakubernetescluster.MicrosoftDefenderAttributes](kc.ref.Append("microsoft_defender"))
}

func (kc dataKubernetesClusterAttributes) NetworkProfile() terra.ListValue[datakubernetescluster.NetworkProfileAttributes] {
	return terra.ReferenceAsList[datakubernetescluster.NetworkProfileAttributes](kc.ref.Append("network_profile"))
}

func (kc dataKubernetesClusterAttributes) OmsAgent() terra.ListValue[datakubernetescluster.OmsAgentAttributes] {
	return terra.ReferenceAsList[datakubernetescluster.OmsAgentAttributes](kc.ref.Append("oms_agent"))
}

func (kc dataKubernetesClusterAttributes) ServicePrincipal() terra.ListValue[datakubernetescluster.ServicePrincipalAttributes] {
	return terra.ReferenceAsList[datakubernetescluster.ServicePrincipalAttributes](kc.ref.Append("service_principal"))
}

func (kc dataKubernetesClusterAttributes) StorageProfile() terra.ListValue[datakubernetescluster.StorageProfileAttributes] {
	return terra.ReferenceAsList[datakubernetescluster.StorageProfileAttributes](kc.ref.Append("storage_profile"))
}

func (kc dataKubernetesClusterAttributes) WindowsProfile() terra.ListValue[datakubernetescluster.WindowsProfileAttributes] {
	return terra.ReferenceAsList[datakubernetescluster.WindowsProfileAttributes](kc.ref.Append("windows_profile"))
}

func (kc dataKubernetesClusterAttributes) Timeouts() datakubernetescluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datakubernetescluster.TimeoutsAttributes](kc.ref.Append("timeouts"))
}