// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	containercluster "github.com/golingon/terraproviders/google/4.75.1/containercluster"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewContainerCluster creates a new instance of [ContainerCluster].
func NewContainerCluster(name string, args ContainerClusterArgs) *ContainerCluster {
	return &ContainerCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ContainerCluster)(nil)

// ContainerCluster represents the Terraform resource google_container_cluster.
type ContainerCluster struct {
	Name      string
	Args      ContainerClusterArgs
	state     *containerClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ContainerCluster].
func (cc *ContainerCluster) Type() string {
	return "google_container_cluster"
}

// LocalName returns the local name for [ContainerCluster].
func (cc *ContainerCluster) LocalName() string {
	return cc.Name
}

// Configuration returns the configuration (args) for [ContainerCluster].
func (cc *ContainerCluster) Configuration() interface{} {
	return cc.Args
}

// DependOn is used for other resources to depend on [ContainerCluster].
func (cc *ContainerCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(cc)
}

// Dependencies returns the list of resources [ContainerCluster] depends_on.
func (cc *ContainerCluster) Dependencies() terra.Dependencies {
	return cc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ContainerCluster].
func (cc *ContainerCluster) LifecycleManagement() *terra.Lifecycle {
	return cc.Lifecycle
}

// Attributes returns the attributes for [ContainerCluster].
func (cc *ContainerCluster) Attributes() containerClusterAttributes {
	return containerClusterAttributes{ref: terra.ReferenceResource(cc)}
}

// ImportState imports the given attribute values into [ContainerCluster]'s state.
func (cc *ContainerCluster) ImportState(av io.Reader) error {
	cc.state = &containerClusterState{}
	if err := json.NewDecoder(av).Decode(cc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cc.Type(), cc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ContainerCluster] has state.
func (cc *ContainerCluster) State() (*containerClusterState, bool) {
	return cc.state, cc.state != nil
}

// StateMust returns the state for [ContainerCluster]. Panics if the state is nil.
func (cc *ContainerCluster) StateMust() *containerClusterState {
	if cc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cc.Type(), cc.LocalName()))
	}
	return cc.state
}

// ContainerClusterArgs contains the configurations for google_container_cluster.
type ContainerClusterArgs struct {
	// ClusterIpv4Cidr: string, optional
	ClusterIpv4Cidr terra.StringValue `hcl:"cluster_ipv4_cidr,attr"`
	// DatapathProvider: string, optional
	DatapathProvider terra.StringValue `hcl:"datapath_provider,attr"`
	// DefaultMaxPodsPerNode: number, optional
	DefaultMaxPodsPerNode terra.NumberValue `hcl:"default_max_pods_per_node,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// EnableAutopilot: bool, optional
	EnableAutopilot terra.BoolValue `hcl:"enable_autopilot,attr"`
	// EnableBinaryAuthorization: bool, optional
	EnableBinaryAuthorization terra.BoolValue `hcl:"enable_binary_authorization,attr"`
	// EnableIntranodeVisibility: bool, optional
	EnableIntranodeVisibility terra.BoolValue `hcl:"enable_intranode_visibility,attr"`
	// EnableKubernetesAlpha: bool, optional
	EnableKubernetesAlpha terra.BoolValue `hcl:"enable_kubernetes_alpha,attr"`
	// EnableL4IlbSubsetting: bool, optional
	EnableL4IlbSubsetting terra.BoolValue `hcl:"enable_l4_ilb_subsetting,attr"`
	// EnableLegacyAbac: bool, optional
	EnableLegacyAbac terra.BoolValue `hcl:"enable_legacy_abac,attr"`
	// EnableShieldedNodes: bool, optional
	EnableShieldedNodes terra.BoolValue `hcl:"enable_shielded_nodes,attr"`
	// EnableTpu: bool, optional
	EnableTpu terra.BoolValue `hcl:"enable_tpu,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InitialNodeCount: number, optional
	InitialNodeCount terra.NumberValue `hcl:"initial_node_count,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// LoggingService: string, optional
	LoggingService terra.StringValue `hcl:"logging_service,attr"`
	// MinMasterVersion: string, optional
	MinMasterVersion terra.StringValue `hcl:"min_master_version,attr"`
	// MonitoringService: string, optional
	MonitoringService terra.StringValue `hcl:"monitoring_service,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Network: string, optional
	Network terra.StringValue `hcl:"network,attr"`
	// NetworkingMode: string, optional
	NetworkingMode terra.StringValue `hcl:"networking_mode,attr"`
	// NodeLocations: set of string, optional
	NodeLocations terra.SetValue[terra.StringValue] `hcl:"node_locations,attr"`
	// NodeVersion: string, optional
	NodeVersion terra.StringValue `hcl:"node_version,attr"`
	// PrivateIpv6GoogleAccess: string, optional
	PrivateIpv6GoogleAccess terra.StringValue `hcl:"private_ipv6_google_access,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// RemoveDefaultNodePool: bool, optional
	RemoveDefaultNodePool terra.BoolValue `hcl:"remove_default_node_pool,attr"`
	// ResourceLabels: map of string, optional
	ResourceLabels terra.MapValue[terra.StringValue] `hcl:"resource_labels,attr"`
	// Subnetwork: string, optional
	Subnetwork terra.StringValue `hcl:"subnetwork,attr"`
	// AddonsConfig: optional
	AddonsConfig *containercluster.AddonsConfig `hcl:"addons_config,block"`
	// AuthenticatorGroupsConfig: optional
	AuthenticatorGroupsConfig *containercluster.AuthenticatorGroupsConfig `hcl:"authenticator_groups_config,block"`
	// BinaryAuthorization: optional
	BinaryAuthorization *containercluster.BinaryAuthorization `hcl:"binary_authorization,block"`
	// ClusterAutoscaling: optional
	ClusterAutoscaling *containercluster.ClusterAutoscaling `hcl:"cluster_autoscaling,block"`
	// ConfidentialNodes: optional
	ConfidentialNodes *containercluster.ConfidentialNodes `hcl:"confidential_nodes,block"`
	// CostManagementConfig: optional
	CostManagementConfig *containercluster.CostManagementConfig `hcl:"cost_management_config,block"`
	// DatabaseEncryption: optional
	DatabaseEncryption *containercluster.DatabaseEncryption `hcl:"database_encryption,block"`
	// DefaultSnatStatus: optional
	DefaultSnatStatus *containercluster.DefaultSnatStatus `hcl:"default_snat_status,block"`
	// DnsConfig: optional
	DnsConfig *containercluster.DnsConfig `hcl:"dns_config,block"`
	// GatewayApiConfig: optional
	GatewayApiConfig *containercluster.GatewayApiConfig `hcl:"gateway_api_config,block"`
	// IpAllocationPolicy: optional
	IpAllocationPolicy *containercluster.IpAllocationPolicy `hcl:"ip_allocation_policy,block"`
	// LoggingConfig: optional
	LoggingConfig *containercluster.LoggingConfig `hcl:"logging_config,block"`
	// MaintenancePolicy: optional
	MaintenancePolicy *containercluster.MaintenancePolicy `hcl:"maintenance_policy,block"`
	// MasterAuth: optional
	MasterAuth *containercluster.MasterAuth `hcl:"master_auth,block"`
	// MasterAuthorizedNetworksConfig: optional
	MasterAuthorizedNetworksConfig *containercluster.MasterAuthorizedNetworksConfig `hcl:"master_authorized_networks_config,block"`
	// MeshCertificates: optional
	MeshCertificates *containercluster.MeshCertificates `hcl:"mesh_certificates,block"`
	// MonitoringConfig: optional
	MonitoringConfig *containercluster.MonitoringConfig `hcl:"monitoring_config,block"`
	// NetworkPolicy: optional
	NetworkPolicy *containercluster.NetworkPolicy `hcl:"network_policy,block"`
	// NodeConfig: optional
	NodeConfig *containercluster.NodeConfig `hcl:"node_config,block"`
	// NodePool: min=0
	NodePool []containercluster.NodePool `hcl:"node_pool,block" validate:"min=0"`
	// NodePoolDefaults: optional
	NodePoolDefaults *containercluster.NodePoolDefaults `hcl:"node_pool_defaults,block"`
	// NotificationConfig: optional
	NotificationConfig *containercluster.NotificationConfig `hcl:"notification_config,block"`
	// PrivateClusterConfig: optional
	PrivateClusterConfig *containercluster.PrivateClusterConfig `hcl:"private_cluster_config,block"`
	// ReleaseChannel: optional
	ReleaseChannel *containercluster.ReleaseChannel `hcl:"release_channel,block"`
	// ResourceUsageExportConfig: optional
	ResourceUsageExportConfig *containercluster.ResourceUsageExportConfig `hcl:"resource_usage_export_config,block"`
	// SecurityPostureConfig: optional
	SecurityPostureConfig *containercluster.SecurityPostureConfig `hcl:"security_posture_config,block"`
	// ServiceExternalIpsConfig: optional
	ServiceExternalIpsConfig *containercluster.ServiceExternalIpsConfig `hcl:"service_external_ips_config,block"`
	// Timeouts: optional
	Timeouts *containercluster.Timeouts `hcl:"timeouts,block"`
	// VerticalPodAutoscaling: optional
	VerticalPodAutoscaling *containercluster.VerticalPodAutoscaling `hcl:"vertical_pod_autoscaling,block"`
	// WorkloadIdentityConfig: optional
	WorkloadIdentityConfig *containercluster.WorkloadIdentityConfig `hcl:"workload_identity_config,block"`
}
type containerClusterAttributes struct {
	ref terra.Reference
}

// ClusterIpv4Cidr returns a reference to field cluster_ipv4_cidr of google_container_cluster.
func (cc containerClusterAttributes) ClusterIpv4Cidr() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("cluster_ipv4_cidr"))
}

// DatapathProvider returns a reference to field datapath_provider of google_container_cluster.
func (cc containerClusterAttributes) DatapathProvider() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("datapath_provider"))
}

// DefaultMaxPodsPerNode returns a reference to field default_max_pods_per_node of google_container_cluster.
func (cc containerClusterAttributes) DefaultMaxPodsPerNode() terra.NumberValue {
	return terra.ReferenceAsNumber(cc.ref.Append("default_max_pods_per_node"))
}

// Description returns a reference to field description of google_container_cluster.
func (cc containerClusterAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("description"))
}

// EnableAutopilot returns a reference to field enable_autopilot of google_container_cluster.
func (cc containerClusterAttributes) EnableAutopilot() terra.BoolValue {
	return terra.ReferenceAsBool(cc.ref.Append("enable_autopilot"))
}

// EnableBinaryAuthorization returns a reference to field enable_binary_authorization of google_container_cluster.
func (cc containerClusterAttributes) EnableBinaryAuthorization() terra.BoolValue {
	return terra.ReferenceAsBool(cc.ref.Append("enable_binary_authorization"))
}

// EnableIntranodeVisibility returns a reference to field enable_intranode_visibility of google_container_cluster.
func (cc containerClusterAttributes) EnableIntranodeVisibility() terra.BoolValue {
	return terra.ReferenceAsBool(cc.ref.Append("enable_intranode_visibility"))
}

// EnableKubernetesAlpha returns a reference to field enable_kubernetes_alpha of google_container_cluster.
func (cc containerClusterAttributes) EnableKubernetesAlpha() terra.BoolValue {
	return terra.ReferenceAsBool(cc.ref.Append("enable_kubernetes_alpha"))
}

// EnableL4IlbSubsetting returns a reference to field enable_l4_ilb_subsetting of google_container_cluster.
func (cc containerClusterAttributes) EnableL4IlbSubsetting() terra.BoolValue {
	return terra.ReferenceAsBool(cc.ref.Append("enable_l4_ilb_subsetting"))
}

// EnableLegacyAbac returns a reference to field enable_legacy_abac of google_container_cluster.
func (cc containerClusterAttributes) EnableLegacyAbac() terra.BoolValue {
	return terra.ReferenceAsBool(cc.ref.Append("enable_legacy_abac"))
}

// EnableShieldedNodes returns a reference to field enable_shielded_nodes of google_container_cluster.
func (cc containerClusterAttributes) EnableShieldedNodes() terra.BoolValue {
	return terra.ReferenceAsBool(cc.ref.Append("enable_shielded_nodes"))
}

// EnableTpu returns a reference to field enable_tpu of google_container_cluster.
func (cc containerClusterAttributes) EnableTpu() terra.BoolValue {
	return terra.ReferenceAsBool(cc.ref.Append("enable_tpu"))
}

// Endpoint returns a reference to field endpoint of google_container_cluster.
func (cc containerClusterAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("endpoint"))
}

// Id returns a reference to field id of google_container_cluster.
func (cc containerClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("id"))
}

// InitialNodeCount returns a reference to field initial_node_count of google_container_cluster.
func (cc containerClusterAttributes) InitialNodeCount() terra.NumberValue {
	return terra.ReferenceAsNumber(cc.ref.Append("initial_node_count"))
}

// LabelFingerprint returns a reference to field label_fingerprint of google_container_cluster.
func (cc containerClusterAttributes) LabelFingerprint() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("label_fingerprint"))
}

// Location returns a reference to field location of google_container_cluster.
func (cc containerClusterAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("location"))
}

// LoggingService returns a reference to field logging_service of google_container_cluster.
func (cc containerClusterAttributes) LoggingService() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("logging_service"))
}

// MasterVersion returns a reference to field master_version of google_container_cluster.
func (cc containerClusterAttributes) MasterVersion() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("master_version"))
}

// MinMasterVersion returns a reference to field min_master_version of google_container_cluster.
func (cc containerClusterAttributes) MinMasterVersion() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("min_master_version"))
}

// MonitoringService returns a reference to field monitoring_service of google_container_cluster.
func (cc containerClusterAttributes) MonitoringService() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("monitoring_service"))
}

// Name returns a reference to field name of google_container_cluster.
func (cc containerClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("name"))
}

// Network returns a reference to field network of google_container_cluster.
func (cc containerClusterAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("network"))
}

// NetworkingMode returns a reference to field networking_mode of google_container_cluster.
func (cc containerClusterAttributes) NetworkingMode() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("networking_mode"))
}

// NodeLocations returns a reference to field node_locations of google_container_cluster.
func (cc containerClusterAttributes) NodeLocations() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cc.ref.Append("node_locations"))
}

// NodeVersion returns a reference to field node_version of google_container_cluster.
func (cc containerClusterAttributes) NodeVersion() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("node_version"))
}

// Operation returns a reference to field operation of google_container_cluster.
func (cc containerClusterAttributes) Operation() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("operation"))
}

// PrivateIpv6GoogleAccess returns a reference to field private_ipv6_google_access of google_container_cluster.
func (cc containerClusterAttributes) PrivateIpv6GoogleAccess() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("private_ipv6_google_access"))
}

// Project returns a reference to field project of google_container_cluster.
func (cc containerClusterAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("project"))
}

// RemoveDefaultNodePool returns a reference to field remove_default_node_pool of google_container_cluster.
func (cc containerClusterAttributes) RemoveDefaultNodePool() terra.BoolValue {
	return terra.ReferenceAsBool(cc.ref.Append("remove_default_node_pool"))
}

// ResourceLabels returns a reference to field resource_labels of google_container_cluster.
func (cc containerClusterAttributes) ResourceLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cc.ref.Append("resource_labels"))
}

// SelfLink returns a reference to field self_link of google_container_cluster.
func (cc containerClusterAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("self_link"))
}

// ServicesIpv4Cidr returns a reference to field services_ipv4_cidr of google_container_cluster.
func (cc containerClusterAttributes) ServicesIpv4Cidr() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("services_ipv4_cidr"))
}

// Subnetwork returns a reference to field subnetwork of google_container_cluster.
func (cc containerClusterAttributes) Subnetwork() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("subnetwork"))
}

// TpuIpv4CidrBlock returns a reference to field tpu_ipv4_cidr_block of google_container_cluster.
func (cc containerClusterAttributes) TpuIpv4CidrBlock() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("tpu_ipv4_cidr_block"))
}

func (cc containerClusterAttributes) AddonsConfig() terra.ListValue[containercluster.AddonsConfigAttributes] {
	return terra.ReferenceAsList[containercluster.AddonsConfigAttributes](cc.ref.Append("addons_config"))
}

func (cc containerClusterAttributes) AuthenticatorGroupsConfig() terra.ListValue[containercluster.AuthenticatorGroupsConfigAttributes] {
	return terra.ReferenceAsList[containercluster.AuthenticatorGroupsConfigAttributes](cc.ref.Append("authenticator_groups_config"))
}

func (cc containerClusterAttributes) BinaryAuthorization() terra.ListValue[containercluster.BinaryAuthorizationAttributes] {
	return terra.ReferenceAsList[containercluster.BinaryAuthorizationAttributes](cc.ref.Append("binary_authorization"))
}

func (cc containerClusterAttributes) ClusterAutoscaling() terra.ListValue[containercluster.ClusterAutoscalingAttributes] {
	return terra.ReferenceAsList[containercluster.ClusterAutoscalingAttributes](cc.ref.Append("cluster_autoscaling"))
}

func (cc containerClusterAttributes) ConfidentialNodes() terra.ListValue[containercluster.ConfidentialNodesAttributes] {
	return terra.ReferenceAsList[containercluster.ConfidentialNodesAttributes](cc.ref.Append("confidential_nodes"))
}

func (cc containerClusterAttributes) CostManagementConfig() terra.ListValue[containercluster.CostManagementConfigAttributes] {
	return terra.ReferenceAsList[containercluster.CostManagementConfigAttributes](cc.ref.Append("cost_management_config"))
}

func (cc containerClusterAttributes) DatabaseEncryption() terra.ListValue[containercluster.DatabaseEncryptionAttributes] {
	return terra.ReferenceAsList[containercluster.DatabaseEncryptionAttributes](cc.ref.Append("database_encryption"))
}

func (cc containerClusterAttributes) DefaultSnatStatus() terra.ListValue[containercluster.DefaultSnatStatusAttributes] {
	return terra.ReferenceAsList[containercluster.DefaultSnatStatusAttributes](cc.ref.Append("default_snat_status"))
}

func (cc containerClusterAttributes) DnsConfig() terra.ListValue[containercluster.DnsConfigAttributes] {
	return terra.ReferenceAsList[containercluster.DnsConfigAttributes](cc.ref.Append("dns_config"))
}

func (cc containerClusterAttributes) GatewayApiConfig() terra.ListValue[containercluster.GatewayApiConfigAttributes] {
	return terra.ReferenceAsList[containercluster.GatewayApiConfigAttributes](cc.ref.Append("gateway_api_config"))
}

func (cc containerClusterAttributes) IpAllocationPolicy() terra.ListValue[containercluster.IpAllocationPolicyAttributes] {
	return terra.ReferenceAsList[containercluster.IpAllocationPolicyAttributes](cc.ref.Append("ip_allocation_policy"))
}

func (cc containerClusterAttributes) LoggingConfig() terra.ListValue[containercluster.LoggingConfigAttributes] {
	return terra.ReferenceAsList[containercluster.LoggingConfigAttributes](cc.ref.Append("logging_config"))
}

func (cc containerClusterAttributes) MaintenancePolicy() terra.ListValue[containercluster.MaintenancePolicyAttributes] {
	return terra.ReferenceAsList[containercluster.MaintenancePolicyAttributes](cc.ref.Append("maintenance_policy"))
}

func (cc containerClusterAttributes) MasterAuth() terra.ListValue[containercluster.MasterAuthAttributes] {
	return terra.ReferenceAsList[containercluster.MasterAuthAttributes](cc.ref.Append("master_auth"))
}

func (cc containerClusterAttributes) MasterAuthorizedNetworksConfig() terra.ListValue[containercluster.MasterAuthorizedNetworksConfigAttributes] {
	return terra.ReferenceAsList[containercluster.MasterAuthorizedNetworksConfigAttributes](cc.ref.Append("master_authorized_networks_config"))
}

func (cc containerClusterAttributes) MeshCertificates() terra.ListValue[containercluster.MeshCertificatesAttributes] {
	return terra.ReferenceAsList[containercluster.MeshCertificatesAttributes](cc.ref.Append("mesh_certificates"))
}

func (cc containerClusterAttributes) MonitoringConfig() terra.ListValue[containercluster.MonitoringConfigAttributes] {
	return terra.ReferenceAsList[containercluster.MonitoringConfigAttributes](cc.ref.Append("monitoring_config"))
}

func (cc containerClusterAttributes) NetworkPolicy() terra.ListValue[containercluster.NetworkPolicyAttributes] {
	return terra.ReferenceAsList[containercluster.NetworkPolicyAttributes](cc.ref.Append("network_policy"))
}

func (cc containerClusterAttributes) NodeConfig() terra.ListValue[containercluster.NodeConfigAttributes] {
	return terra.ReferenceAsList[containercluster.NodeConfigAttributes](cc.ref.Append("node_config"))
}

func (cc containerClusterAttributes) NodePool() terra.ListValue[containercluster.NodePoolAttributes] {
	return terra.ReferenceAsList[containercluster.NodePoolAttributes](cc.ref.Append("node_pool"))
}

func (cc containerClusterAttributes) NodePoolDefaults() terra.ListValue[containercluster.NodePoolDefaultsAttributes] {
	return terra.ReferenceAsList[containercluster.NodePoolDefaultsAttributes](cc.ref.Append("node_pool_defaults"))
}

func (cc containerClusterAttributes) NotificationConfig() terra.ListValue[containercluster.NotificationConfigAttributes] {
	return terra.ReferenceAsList[containercluster.NotificationConfigAttributes](cc.ref.Append("notification_config"))
}

func (cc containerClusterAttributes) PrivateClusterConfig() terra.ListValue[containercluster.PrivateClusterConfigAttributes] {
	return terra.ReferenceAsList[containercluster.PrivateClusterConfigAttributes](cc.ref.Append("private_cluster_config"))
}

func (cc containerClusterAttributes) ReleaseChannel() terra.ListValue[containercluster.ReleaseChannelAttributes] {
	return terra.ReferenceAsList[containercluster.ReleaseChannelAttributes](cc.ref.Append("release_channel"))
}

func (cc containerClusterAttributes) ResourceUsageExportConfig() terra.ListValue[containercluster.ResourceUsageExportConfigAttributes] {
	return terra.ReferenceAsList[containercluster.ResourceUsageExportConfigAttributes](cc.ref.Append("resource_usage_export_config"))
}

func (cc containerClusterAttributes) SecurityPostureConfig() terra.ListValue[containercluster.SecurityPostureConfigAttributes] {
	return terra.ReferenceAsList[containercluster.SecurityPostureConfigAttributes](cc.ref.Append("security_posture_config"))
}

func (cc containerClusterAttributes) ServiceExternalIpsConfig() terra.ListValue[containercluster.ServiceExternalIpsConfigAttributes] {
	return terra.ReferenceAsList[containercluster.ServiceExternalIpsConfigAttributes](cc.ref.Append("service_external_ips_config"))
}

func (cc containerClusterAttributes) Timeouts() containercluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[containercluster.TimeoutsAttributes](cc.ref.Append("timeouts"))
}

func (cc containerClusterAttributes) VerticalPodAutoscaling() terra.ListValue[containercluster.VerticalPodAutoscalingAttributes] {
	return terra.ReferenceAsList[containercluster.VerticalPodAutoscalingAttributes](cc.ref.Append("vertical_pod_autoscaling"))
}

func (cc containerClusterAttributes) WorkloadIdentityConfig() terra.ListValue[containercluster.WorkloadIdentityConfigAttributes] {
	return terra.ReferenceAsList[containercluster.WorkloadIdentityConfigAttributes](cc.ref.Append("workload_identity_config"))
}

type containerClusterState struct {
	ClusterIpv4Cidr                string                                                 `json:"cluster_ipv4_cidr"`
	DatapathProvider               string                                                 `json:"datapath_provider"`
	DefaultMaxPodsPerNode          float64                                                `json:"default_max_pods_per_node"`
	Description                    string                                                 `json:"description"`
	EnableAutopilot                bool                                                   `json:"enable_autopilot"`
	EnableBinaryAuthorization      bool                                                   `json:"enable_binary_authorization"`
	EnableIntranodeVisibility      bool                                                   `json:"enable_intranode_visibility"`
	EnableKubernetesAlpha          bool                                                   `json:"enable_kubernetes_alpha"`
	EnableL4IlbSubsetting          bool                                                   `json:"enable_l4_ilb_subsetting"`
	EnableLegacyAbac               bool                                                   `json:"enable_legacy_abac"`
	EnableShieldedNodes            bool                                                   `json:"enable_shielded_nodes"`
	EnableTpu                      bool                                                   `json:"enable_tpu"`
	Endpoint                       string                                                 `json:"endpoint"`
	Id                             string                                                 `json:"id"`
	InitialNodeCount               float64                                                `json:"initial_node_count"`
	LabelFingerprint               string                                                 `json:"label_fingerprint"`
	Location                       string                                                 `json:"location"`
	LoggingService                 string                                                 `json:"logging_service"`
	MasterVersion                  string                                                 `json:"master_version"`
	MinMasterVersion               string                                                 `json:"min_master_version"`
	MonitoringService              string                                                 `json:"monitoring_service"`
	Name                           string                                                 `json:"name"`
	Network                        string                                                 `json:"network"`
	NetworkingMode                 string                                                 `json:"networking_mode"`
	NodeLocations                  []string                                               `json:"node_locations"`
	NodeVersion                    string                                                 `json:"node_version"`
	Operation                      string                                                 `json:"operation"`
	PrivateIpv6GoogleAccess        string                                                 `json:"private_ipv6_google_access"`
	Project                        string                                                 `json:"project"`
	RemoveDefaultNodePool          bool                                                   `json:"remove_default_node_pool"`
	ResourceLabels                 map[string]string                                      `json:"resource_labels"`
	SelfLink                       string                                                 `json:"self_link"`
	ServicesIpv4Cidr               string                                                 `json:"services_ipv4_cidr"`
	Subnetwork                     string                                                 `json:"subnetwork"`
	TpuIpv4CidrBlock               string                                                 `json:"tpu_ipv4_cidr_block"`
	AddonsConfig                   []containercluster.AddonsConfigState                   `json:"addons_config"`
	AuthenticatorGroupsConfig      []containercluster.AuthenticatorGroupsConfigState      `json:"authenticator_groups_config"`
	BinaryAuthorization            []containercluster.BinaryAuthorizationState            `json:"binary_authorization"`
	ClusterAutoscaling             []containercluster.ClusterAutoscalingState             `json:"cluster_autoscaling"`
	ConfidentialNodes              []containercluster.ConfidentialNodesState              `json:"confidential_nodes"`
	CostManagementConfig           []containercluster.CostManagementConfigState           `json:"cost_management_config"`
	DatabaseEncryption             []containercluster.DatabaseEncryptionState             `json:"database_encryption"`
	DefaultSnatStatus              []containercluster.DefaultSnatStatusState              `json:"default_snat_status"`
	DnsConfig                      []containercluster.DnsConfigState                      `json:"dns_config"`
	GatewayApiConfig               []containercluster.GatewayApiConfigState               `json:"gateway_api_config"`
	IpAllocationPolicy             []containercluster.IpAllocationPolicyState             `json:"ip_allocation_policy"`
	LoggingConfig                  []containercluster.LoggingConfigState                  `json:"logging_config"`
	MaintenancePolicy              []containercluster.MaintenancePolicyState              `json:"maintenance_policy"`
	MasterAuth                     []containercluster.MasterAuthState                     `json:"master_auth"`
	MasterAuthorizedNetworksConfig []containercluster.MasterAuthorizedNetworksConfigState `json:"master_authorized_networks_config"`
	MeshCertificates               []containercluster.MeshCertificatesState               `json:"mesh_certificates"`
	MonitoringConfig               []containercluster.MonitoringConfigState               `json:"monitoring_config"`
	NetworkPolicy                  []containercluster.NetworkPolicyState                  `json:"network_policy"`
	NodeConfig                     []containercluster.NodeConfigState                     `json:"node_config"`
	NodePool                       []containercluster.NodePoolState                       `json:"node_pool"`
	NodePoolDefaults               []containercluster.NodePoolDefaultsState               `json:"node_pool_defaults"`
	NotificationConfig             []containercluster.NotificationConfigState             `json:"notification_config"`
	PrivateClusterConfig           []containercluster.PrivateClusterConfigState           `json:"private_cluster_config"`
	ReleaseChannel                 []containercluster.ReleaseChannelState                 `json:"release_channel"`
	ResourceUsageExportConfig      []containercluster.ResourceUsageExportConfigState      `json:"resource_usage_export_config"`
	SecurityPostureConfig          []containercluster.SecurityPostureConfigState          `json:"security_posture_config"`
	ServiceExternalIpsConfig       []containercluster.ServiceExternalIpsConfigState       `json:"service_external_ips_config"`
	Timeouts                       *containercluster.TimeoutsState                        `json:"timeouts"`
	VerticalPodAutoscaling         []containercluster.VerticalPodAutoscalingState         `json:"vertical_pod_autoscaling"`
	WorkloadIdentityConfig         []containercluster.WorkloadIdentityConfigState         `json:"workload_identity_config"`
}
