// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	datacontainercluster "github.com/golingon/terraproviders/google/4.76.0/datacontainercluster"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataContainerCluster creates a new instance of [DataContainerCluster].
func NewDataContainerCluster(name string, args DataContainerClusterArgs) *DataContainerCluster {
	return &DataContainerCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataContainerCluster)(nil)

// DataContainerCluster represents the Terraform data resource google_container_cluster.
type DataContainerCluster struct {
	Name string
	Args DataContainerClusterArgs
}

// DataSource returns the Terraform object type for [DataContainerCluster].
func (cc *DataContainerCluster) DataSource() string {
	return "google_container_cluster"
}

// LocalName returns the local name for [DataContainerCluster].
func (cc *DataContainerCluster) LocalName() string {
	return cc.Name
}

// Configuration returns the configuration (args) for [DataContainerCluster].
func (cc *DataContainerCluster) Configuration() interface{} {
	return cc.Args
}

// Attributes returns the attributes for [DataContainerCluster].
func (cc *DataContainerCluster) Attributes() dataContainerClusterAttributes {
	return dataContainerClusterAttributes{ref: terra.ReferenceDataResource(cc)}
}

// DataContainerClusterArgs contains the configurations for google_container_cluster.
type DataContainerClusterArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// AddonsConfig: min=0
	AddonsConfig []datacontainercluster.AddonsConfig `hcl:"addons_config,block" validate:"min=0"`
	// AuthenticatorGroupsConfig: min=0
	AuthenticatorGroupsConfig []datacontainercluster.AuthenticatorGroupsConfig `hcl:"authenticator_groups_config,block" validate:"min=0"`
	// BinaryAuthorization: min=0
	BinaryAuthorization []datacontainercluster.BinaryAuthorization `hcl:"binary_authorization,block" validate:"min=0"`
	// ClusterAutoscaling: min=0
	ClusterAutoscaling []datacontainercluster.ClusterAutoscaling `hcl:"cluster_autoscaling,block" validate:"min=0"`
	// ConfidentialNodes: min=0
	ConfidentialNodes []datacontainercluster.ConfidentialNodes `hcl:"confidential_nodes,block" validate:"min=0"`
	// CostManagementConfig: min=0
	CostManagementConfig []datacontainercluster.CostManagementConfig `hcl:"cost_management_config,block" validate:"min=0"`
	// DatabaseEncryption: min=0
	DatabaseEncryption []datacontainercluster.DatabaseEncryption `hcl:"database_encryption,block" validate:"min=0"`
	// DefaultSnatStatus: min=0
	DefaultSnatStatus []datacontainercluster.DefaultSnatStatus `hcl:"default_snat_status,block" validate:"min=0"`
	// DnsConfig: min=0
	DnsConfig []datacontainercluster.DnsConfig `hcl:"dns_config,block" validate:"min=0"`
	// GatewayApiConfig: min=0
	GatewayApiConfig []datacontainercluster.GatewayApiConfig `hcl:"gateway_api_config,block" validate:"min=0"`
	// IpAllocationPolicy: min=0
	IpAllocationPolicy []datacontainercluster.IpAllocationPolicy `hcl:"ip_allocation_policy,block" validate:"min=0"`
	// LoggingConfig: min=0
	LoggingConfig []datacontainercluster.LoggingConfig `hcl:"logging_config,block" validate:"min=0"`
	// MaintenancePolicy: min=0
	MaintenancePolicy []datacontainercluster.MaintenancePolicy `hcl:"maintenance_policy,block" validate:"min=0"`
	// MasterAuth: min=0
	MasterAuth []datacontainercluster.MasterAuth `hcl:"master_auth,block" validate:"min=0"`
	// MasterAuthorizedNetworksConfig: min=0
	MasterAuthorizedNetworksConfig []datacontainercluster.MasterAuthorizedNetworksConfig `hcl:"master_authorized_networks_config,block" validate:"min=0"`
	// MeshCertificates: min=0
	MeshCertificates []datacontainercluster.MeshCertificates `hcl:"mesh_certificates,block" validate:"min=0"`
	// MonitoringConfig: min=0
	MonitoringConfig []datacontainercluster.MonitoringConfig `hcl:"monitoring_config,block" validate:"min=0"`
	// NetworkPolicy: min=0
	NetworkPolicy []datacontainercluster.NetworkPolicy `hcl:"network_policy,block" validate:"min=0"`
	// NodeConfig: min=0
	NodeConfig []datacontainercluster.NodeConfig `hcl:"node_config,block" validate:"min=0"`
	// NodePool: min=0
	NodePool []datacontainercluster.NodePool `hcl:"node_pool,block" validate:"min=0"`
	// NodePoolDefaults: min=0
	NodePoolDefaults []datacontainercluster.NodePoolDefaults `hcl:"node_pool_defaults,block" validate:"min=0"`
	// NotificationConfig: min=0
	NotificationConfig []datacontainercluster.NotificationConfig `hcl:"notification_config,block" validate:"min=0"`
	// PrivateClusterConfig: min=0
	PrivateClusterConfig []datacontainercluster.PrivateClusterConfig `hcl:"private_cluster_config,block" validate:"min=0"`
	// ReleaseChannel: min=0
	ReleaseChannel []datacontainercluster.ReleaseChannel `hcl:"release_channel,block" validate:"min=0"`
	// ResourceUsageExportConfig: min=0
	ResourceUsageExportConfig []datacontainercluster.ResourceUsageExportConfig `hcl:"resource_usage_export_config,block" validate:"min=0"`
	// SecurityPostureConfig: min=0
	SecurityPostureConfig []datacontainercluster.SecurityPostureConfig `hcl:"security_posture_config,block" validate:"min=0"`
	// ServiceExternalIpsConfig: min=0
	ServiceExternalIpsConfig []datacontainercluster.ServiceExternalIpsConfig `hcl:"service_external_ips_config,block" validate:"min=0"`
	// VerticalPodAutoscaling: min=0
	VerticalPodAutoscaling []datacontainercluster.VerticalPodAutoscaling `hcl:"vertical_pod_autoscaling,block" validate:"min=0"`
	// WorkloadIdentityConfig: min=0
	WorkloadIdentityConfig []datacontainercluster.WorkloadIdentityConfig `hcl:"workload_identity_config,block" validate:"min=0"`
}
type dataContainerClusterAttributes struct {
	ref terra.Reference
}

// AllowNetAdmin returns a reference to field allow_net_admin of google_container_cluster.
func (cc dataContainerClusterAttributes) AllowNetAdmin() terra.BoolValue {
	return terra.ReferenceAsBool(cc.ref.Append("allow_net_admin"))
}

// ClusterIpv4Cidr returns a reference to field cluster_ipv4_cidr of google_container_cluster.
func (cc dataContainerClusterAttributes) ClusterIpv4Cidr() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("cluster_ipv4_cidr"))
}

// DatapathProvider returns a reference to field datapath_provider of google_container_cluster.
func (cc dataContainerClusterAttributes) DatapathProvider() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("datapath_provider"))
}

// DefaultMaxPodsPerNode returns a reference to field default_max_pods_per_node of google_container_cluster.
func (cc dataContainerClusterAttributes) DefaultMaxPodsPerNode() terra.NumberValue {
	return terra.ReferenceAsNumber(cc.ref.Append("default_max_pods_per_node"))
}

// Description returns a reference to field description of google_container_cluster.
func (cc dataContainerClusterAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("description"))
}

// EnableAutopilot returns a reference to field enable_autopilot of google_container_cluster.
func (cc dataContainerClusterAttributes) EnableAutopilot() terra.BoolValue {
	return terra.ReferenceAsBool(cc.ref.Append("enable_autopilot"))
}

// EnableBinaryAuthorization returns a reference to field enable_binary_authorization of google_container_cluster.
func (cc dataContainerClusterAttributes) EnableBinaryAuthorization() terra.BoolValue {
	return terra.ReferenceAsBool(cc.ref.Append("enable_binary_authorization"))
}

// EnableIntranodeVisibility returns a reference to field enable_intranode_visibility of google_container_cluster.
func (cc dataContainerClusterAttributes) EnableIntranodeVisibility() terra.BoolValue {
	return terra.ReferenceAsBool(cc.ref.Append("enable_intranode_visibility"))
}

// EnableKubernetesAlpha returns a reference to field enable_kubernetes_alpha of google_container_cluster.
func (cc dataContainerClusterAttributes) EnableKubernetesAlpha() terra.BoolValue {
	return terra.ReferenceAsBool(cc.ref.Append("enable_kubernetes_alpha"))
}

// EnableL4IlbSubsetting returns a reference to field enable_l4_ilb_subsetting of google_container_cluster.
func (cc dataContainerClusterAttributes) EnableL4IlbSubsetting() terra.BoolValue {
	return terra.ReferenceAsBool(cc.ref.Append("enable_l4_ilb_subsetting"))
}

// EnableLegacyAbac returns a reference to field enable_legacy_abac of google_container_cluster.
func (cc dataContainerClusterAttributes) EnableLegacyAbac() terra.BoolValue {
	return terra.ReferenceAsBool(cc.ref.Append("enable_legacy_abac"))
}

// EnableShieldedNodes returns a reference to field enable_shielded_nodes of google_container_cluster.
func (cc dataContainerClusterAttributes) EnableShieldedNodes() terra.BoolValue {
	return terra.ReferenceAsBool(cc.ref.Append("enable_shielded_nodes"))
}

// EnableTpu returns a reference to field enable_tpu of google_container_cluster.
func (cc dataContainerClusterAttributes) EnableTpu() terra.BoolValue {
	return terra.ReferenceAsBool(cc.ref.Append("enable_tpu"))
}

// Endpoint returns a reference to field endpoint of google_container_cluster.
func (cc dataContainerClusterAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("endpoint"))
}

// Id returns a reference to field id of google_container_cluster.
func (cc dataContainerClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("id"))
}

// InitialNodeCount returns a reference to field initial_node_count of google_container_cluster.
func (cc dataContainerClusterAttributes) InitialNodeCount() terra.NumberValue {
	return terra.ReferenceAsNumber(cc.ref.Append("initial_node_count"))
}

// LabelFingerprint returns a reference to field label_fingerprint of google_container_cluster.
func (cc dataContainerClusterAttributes) LabelFingerprint() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("label_fingerprint"))
}

// Location returns a reference to field location of google_container_cluster.
func (cc dataContainerClusterAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("location"))
}

// LoggingService returns a reference to field logging_service of google_container_cluster.
func (cc dataContainerClusterAttributes) LoggingService() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("logging_service"))
}

// MasterVersion returns a reference to field master_version of google_container_cluster.
func (cc dataContainerClusterAttributes) MasterVersion() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("master_version"))
}

// MinMasterVersion returns a reference to field min_master_version of google_container_cluster.
func (cc dataContainerClusterAttributes) MinMasterVersion() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("min_master_version"))
}

// MonitoringService returns a reference to field monitoring_service of google_container_cluster.
func (cc dataContainerClusterAttributes) MonitoringService() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("monitoring_service"))
}

// Name returns a reference to field name of google_container_cluster.
func (cc dataContainerClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("name"))
}

// Network returns a reference to field network of google_container_cluster.
func (cc dataContainerClusterAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("network"))
}

// NetworkingMode returns a reference to field networking_mode of google_container_cluster.
func (cc dataContainerClusterAttributes) NetworkingMode() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("networking_mode"))
}

// NodeLocations returns a reference to field node_locations of google_container_cluster.
func (cc dataContainerClusterAttributes) NodeLocations() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cc.ref.Append("node_locations"))
}

// NodeVersion returns a reference to field node_version of google_container_cluster.
func (cc dataContainerClusterAttributes) NodeVersion() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("node_version"))
}

// Operation returns a reference to field operation of google_container_cluster.
func (cc dataContainerClusterAttributes) Operation() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("operation"))
}

// PrivateIpv6GoogleAccess returns a reference to field private_ipv6_google_access of google_container_cluster.
func (cc dataContainerClusterAttributes) PrivateIpv6GoogleAccess() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("private_ipv6_google_access"))
}

// Project returns a reference to field project of google_container_cluster.
func (cc dataContainerClusterAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("project"))
}

// RemoveDefaultNodePool returns a reference to field remove_default_node_pool of google_container_cluster.
func (cc dataContainerClusterAttributes) RemoveDefaultNodePool() terra.BoolValue {
	return terra.ReferenceAsBool(cc.ref.Append("remove_default_node_pool"))
}

// ResourceLabels returns a reference to field resource_labels of google_container_cluster.
func (cc dataContainerClusterAttributes) ResourceLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cc.ref.Append("resource_labels"))
}

// SelfLink returns a reference to field self_link of google_container_cluster.
func (cc dataContainerClusterAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("self_link"))
}

// ServicesIpv4Cidr returns a reference to field services_ipv4_cidr of google_container_cluster.
func (cc dataContainerClusterAttributes) ServicesIpv4Cidr() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("services_ipv4_cidr"))
}

// Subnetwork returns a reference to field subnetwork of google_container_cluster.
func (cc dataContainerClusterAttributes) Subnetwork() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("subnetwork"))
}

// TpuIpv4CidrBlock returns a reference to field tpu_ipv4_cidr_block of google_container_cluster.
func (cc dataContainerClusterAttributes) TpuIpv4CidrBlock() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("tpu_ipv4_cidr_block"))
}

func (cc dataContainerClusterAttributes) AddonsConfig() terra.ListValue[datacontainercluster.AddonsConfigAttributes] {
	return terra.ReferenceAsList[datacontainercluster.AddonsConfigAttributes](cc.ref.Append("addons_config"))
}

func (cc dataContainerClusterAttributes) AuthenticatorGroupsConfig() terra.ListValue[datacontainercluster.AuthenticatorGroupsConfigAttributes] {
	return terra.ReferenceAsList[datacontainercluster.AuthenticatorGroupsConfigAttributes](cc.ref.Append("authenticator_groups_config"))
}

func (cc dataContainerClusterAttributes) BinaryAuthorization() terra.ListValue[datacontainercluster.BinaryAuthorizationAttributes] {
	return terra.ReferenceAsList[datacontainercluster.BinaryAuthorizationAttributes](cc.ref.Append("binary_authorization"))
}

func (cc dataContainerClusterAttributes) ClusterAutoscaling() terra.ListValue[datacontainercluster.ClusterAutoscalingAttributes] {
	return terra.ReferenceAsList[datacontainercluster.ClusterAutoscalingAttributes](cc.ref.Append("cluster_autoscaling"))
}

func (cc dataContainerClusterAttributes) ConfidentialNodes() terra.ListValue[datacontainercluster.ConfidentialNodesAttributes] {
	return terra.ReferenceAsList[datacontainercluster.ConfidentialNodesAttributes](cc.ref.Append("confidential_nodes"))
}

func (cc dataContainerClusterAttributes) CostManagementConfig() terra.ListValue[datacontainercluster.CostManagementConfigAttributes] {
	return terra.ReferenceAsList[datacontainercluster.CostManagementConfigAttributes](cc.ref.Append("cost_management_config"))
}

func (cc dataContainerClusterAttributes) DatabaseEncryption() terra.ListValue[datacontainercluster.DatabaseEncryptionAttributes] {
	return terra.ReferenceAsList[datacontainercluster.DatabaseEncryptionAttributes](cc.ref.Append("database_encryption"))
}

func (cc dataContainerClusterAttributes) DefaultSnatStatus() terra.ListValue[datacontainercluster.DefaultSnatStatusAttributes] {
	return terra.ReferenceAsList[datacontainercluster.DefaultSnatStatusAttributes](cc.ref.Append("default_snat_status"))
}

func (cc dataContainerClusterAttributes) DnsConfig() terra.ListValue[datacontainercluster.DnsConfigAttributes] {
	return terra.ReferenceAsList[datacontainercluster.DnsConfigAttributes](cc.ref.Append("dns_config"))
}

func (cc dataContainerClusterAttributes) GatewayApiConfig() terra.ListValue[datacontainercluster.GatewayApiConfigAttributes] {
	return terra.ReferenceAsList[datacontainercluster.GatewayApiConfigAttributes](cc.ref.Append("gateway_api_config"))
}

func (cc dataContainerClusterAttributes) IpAllocationPolicy() terra.ListValue[datacontainercluster.IpAllocationPolicyAttributes] {
	return terra.ReferenceAsList[datacontainercluster.IpAllocationPolicyAttributes](cc.ref.Append("ip_allocation_policy"))
}

func (cc dataContainerClusterAttributes) LoggingConfig() terra.ListValue[datacontainercluster.LoggingConfigAttributes] {
	return terra.ReferenceAsList[datacontainercluster.LoggingConfigAttributes](cc.ref.Append("logging_config"))
}

func (cc dataContainerClusterAttributes) MaintenancePolicy() terra.ListValue[datacontainercluster.MaintenancePolicyAttributes] {
	return terra.ReferenceAsList[datacontainercluster.MaintenancePolicyAttributes](cc.ref.Append("maintenance_policy"))
}

func (cc dataContainerClusterAttributes) MasterAuth() terra.ListValue[datacontainercluster.MasterAuthAttributes] {
	return terra.ReferenceAsList[datacontainercluster.MasterAuthAttributes](cc.ref.Append("master_auth"))
}

func (cc dataContainerClusterAttributes) MasterAuthorizedNetworksConfig() terra.ListValue[datacontainercluster.MasterAuthorizedNetworksConfigAttributes] {
	return terra.ReferenceAsList[datacontainercluster.MasterAuthorizedNetworksConfigAttributes](cc.ref.Append("master_authorized_networks_config"))
}

func (cc dataContainerClusterAttributes) MeshCertificates() terra.ListValue[datacontainercluster.MeshCertificatesAttributes] {
	return terra.ReferenceAsList[datacontainercluster.MeshCertificatesAttributes](cc.ref.Append("mesh_certificates"))
}

func (cc dataContainerClusterAttributes) MonitoringConfig() terra.ListValue[datacontainercluster.MonitoringConfigAttributes] {
	return terra.ReferenceAsList[datacontainercluster.MonitoringConfigAttributes](cc.ref.Append("monitoring_config"))
}

func (cc dataContainerClusterAttributes) NetworkPolicy() terra.ListValue[datacontainercluster.NetworkPolicyAttributes] {
	return terra.ReferenceAsList[datacontainercluster.NetworkPolicyAttributes](cc.ref.Append("network_policy"))
}

func (cc dataContainerClusterAttributes) NodeConfig() terra.ListValue[datacontainercluster.NodeConfigAttributes] {
	return terra.ReferenceAsList[datacontainercluster.NodeConfigAttributes](cc.ref.Append("node_config"))
}

func (cc dataContainerClusterAttributes) NodePool() terra.ListValue[datacontainercluster.NodePoolAttributes] {
	return terra.ReferenceAsList[datacontainercluster.NodePoolAttributes](cc.ref.Append("node_pool"))
}

func (cc dataContainerClusterAttributes) NodePoolDefaults() terra.ListValue[datacontainercluster.NodePoolDefaultsAttributes] {
	return terra.ReferenceAsList[datacontainercluster.NodePoolDefaultsAttributes](cc.ref.Append("node_pool_defaults"))
}

func (cc dataContainerClusterAttributes) NotificationConfig() terra.ListValue[datacontainercluster.NotificationConfigAttributes] {
	return terra.ReferenceAsList[datacontainercluster.NotificationConfigAttributes](cc.ref.Append("notification_config"))
}

func (cc dataContainerClusterAttributes) PrivateClusterConfig() terra.ListValue[datacontainercluster.PrivateClusterConfigAttributes] {
	return terra.ReferenceAsList[datacontainercluster.PrivateClusterConfigAttributes](cc.ref.Append("private_cluster_config"))
}

func (cc dataContainerClusterAttributes) ReleaseChannel() terra.ListValue[datacontainercluster.ReleaseChannelAttributes] {
	return terra.ReferenceAsList[datacontainercluster.ReleaseChannelAttributes](cc.ref.Append("release_channel"))
}

func (cc dataContainerClusterAttributes) ResourceUsageExportConfig() terra.ListValue[datacontainercluster.ResourceUsageExportConfigAttributes] {
	return terra.ReferenceAsList[datacontainercluster.ResourceUsageExportConfigAttributes](cc.ref.Append("resource_usage_export_config"))
}

func (cc dataContainerClusterAttributes) SecurityPostureConfig() terra.ListValue[datacontainercluster.SecurityPostureConfigAttributes] {
	return terra.ReferenceAsList[datacontainercluster.SecurityPostureConfigAttributes](cc.ref.Append("security_posture_config"))
}

func (cc dataContainerClusterAttributes) ServiceExternalIpsConfig() terra.ListValue[datacontainercluster.ServiceExternalIpsConfigAttributes] {
	return terra.ReferenceAsList[datacontainercluster.ServiceExternalIpsConfigAttributes](cc.ref.Append("service_external_ips_config"))
}

func (cc dataContainerClusterAttributes) VerticalPodAutoscaling() terra.ListValue[datacontainercluster.VerticalPodAutoscalingAttributes] {
	return terra.ReferenceAsList[datacontainercluster.VerticalPodAutoscalingAttributes](cc.ref.Append("vertical_pod_autoscaling"))
}

func (cc dataContainerClusterAttributes) WorkloadIdentityConfig() terra.ListValue[datacontainercluster.WorkloadIdentityConfigAttributes] {
	return terra.ReferenceAsList[datacontainercluster.WorkloadIdentityConfigAttributes](cc.ref.Append("workload_identity_config"))
}
