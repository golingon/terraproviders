// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	kubernetesclusternodepool "github.com/golingon/terraproviders/azurerm/3.58.0/kubernetesclusternodepool"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKubernetesClusterNodePool creates a new instance of [KubernetesClusterNodePool].
func NewKubernetesClusterNodePool(name string, args KubernetesClusterNodePoolArgs) *KubernetesClusterNodePool {
	return &KubernetesClusterNodePool{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KubernetesClusterNodePool)(nil)

// KubernetesClusterNodePool represents the Terraform resource azurerm_kubernetes_cluster_node_pool.
type KubernetesClusterNodePool struct {
	Name      string
	Args      KubernetesClusterNodePoolArgs
	state     *kubernetesClusterNodePoolState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KubernetesClusterNodePool].
func (kcnp *KubernetesClusterNodePool) Type() string {
	return "azurerm_kubernetes_cluster_node_pool"
}

// LocalName returns the local name for [KubernetesClusterNodePool].
func (kcnp *KubernetesClusterNodePool) LocalName() string {
	return kcnp.Name
}

// Configuration returns the configuration (args) for [KubernetesClusterNodePool].
func (kcnp *KubernetesClusterNodePool) Configuration() interface{} {
	return kcnp.Args
}

// DependOn is used for other resources to depend on [KubernetesClusterNodePool].
func (kcnp *KubernetesClusterNodePool) DependOn() terra.Reference {
	return terra.ReferenceResource(kcnp)
}

// Dependencies returns the list of resources [KubernetesClusterNodePool] depends_on.
func (kcnp *KubernetesClusterNodePool) Dependencies() terra.Dependencies {
	return kcnp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KubernetesClusterNodePool].
func (kcnp *KubernetesClusterNodePool) LifecycleManagement() *terra.Lifecycle {
	return kcnp.Lifecycle
}

// Attributes returns the attributes for [KubernetesClusterNodePool].
func (kcnp *KubernetesClusterNodePool) Attributes() kubernetesClusterNodePoolAttributes {
	return kubernetesClusterNodePoolAttributes{ref: terra.ReferenceResource(kcnp)}
}

// ImportState imports the given attribute values into [KubernetesClusterNodePool]'s state.
func (kcnp *KubernetesClusterNodePool) ImportState(av io.Reader) error {
	kcnp.state = &kubernetesClusterNodePoolState{}
	if err := json.NewDecoder(av).Decode(kcnp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kcnp.Type(), kcnp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KubernetesClusterNodePool] has state.
func (kcnp *KubernetesClusterNodePool) State() (*kubernetesClusterNodePoolState, bool) {
	return kcnp.state, kcnp.state != nil
}

// StateMust returns the state for [KubernetesClusterNodePool]. Panics if the state is nil.
func (kcnp *KubernetesClusterNodePool) StateMust() *kubernetesClusterNodePoolState {
	if kcnp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kcnp.Type(), kcnp.LocalName()))
	}
	return kcnp.state
}

// KubernetesClusterNodePoolArgs contains the configurations for azurerm_kubernetes_cluster_node_pool.
type KubernetesClusterNodePoolArgs struct {
	// CapacityReservationGroupId: string, optional
	CapacityReservationGroupId terra.StringValue `hcl:"capacity_reservation_group_id,attr"`
	// CustomCaTrustEnabled: bool, optional
	CustomCaTrustEnabled terra.BoolValue `hcl:"custom_ca_trust_enabled,attr"`
	// EnableAutoScaling: bool, optional
	EnableAutoScaling terra.BoolValue `hcl:"enable_auto_scaling,attr"`
	// EnableHostEncryption: bool, optional
	EnableHostEncryption terra.BoolValue `hcl:"enable_host_encryption,attr"`
	// EnableNodePublicIp: bool, optional
	EnableNodePublicIp terra.BoolValue `hcl:"enable_node_public_ip,attr"`
	// EvictionPolicy: string, optional
	EvictionPolicy terra.StringValue `hcl:"eviction_policy,attr"`
	// FipsEnabled: bool, optional
	FipsEnabled terra.BoolValue `hcl:"fips_enabled,attr"`
	// HostGroupId: string, optional
	HostGroupId terra.StringValue `hcl:"host_group_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KubeletDiskType: string, optional
	KubeletDiskType terra.StringValue `hcl:"kubelet_disk_type,attr"`
	// KubernetesClusterId: string, required
	KubernetesClusterId terra.StringValue `hcl:"kubernetes_cluster_id,attr" validate:"required"`
	// MaxCount: number, optional
	MaxCount terra.NumberValue `hcl:"max_count,attr"`
	// MaxPods: number, optional
	MaxPods terra.NumberValue `hcl:"max_pods,attr"`
	// MessageOfTheDay: string, optional
	MessageOfTheDay terra.StringValue `hcl:"message_of_the_day,attr"`
	// MinCount: number, optional
	MinCount terra.NumberValue `hcl:"min_count,attr"`
	// Mode: string, optional
	Mode terra.StringValue `hcl:"mode,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NodeCount: number, optional
	NodeCount terra.NumberValue `hcl:"node_count,attr"`
	// NodeLabels: map of string, optional
	NodeLabels terra.MapValue[terra.StringValue] `hcl:"node_labels,attr"`
	// NodePublicIpPrefixId: string, optional
	NodePublicIpPrefixId terra.StringValue `hcl:"node_public_ip_prefix_id,attr"`
	// NodeTaints: list of string, optional
	NodeTaints terra.ListValue[terra.StringValue] `hcl:"node_taints,attr"`
	// OrchestratorVersion: string, optional
	OrchestratorVersion terra.StringValue `hcl:"orchestrator_version,attr"`
	// OsDiskSizeGb: number, optional
	OsDiskSizeGb terra.NumberValue `hcl:"os_disk_size_gb,attr"`
	// OsDiskType: string, optional
	OsDiskType terra.StringValue `hcl:"os_disk_type,attr"`
	// OsSku: string, optional
	OsSku terra.StringValue `hcl:"os_sku,attr"`
	// OsType: string, optional
	OsType terra.StringValue `hcl:"os_type,attr"`
	// PodSubnetId: string, optional
	PodSubnetId terra.StringValue `hcl:"pod_subnet_id,attr"`
	// Priority: string, optional
	Priority terra.StringValue `hcl:"priority,attr"`
	// ProximityPlacementGroupId: string, optional
	ProximityPlacementGroupId terra.StringValue `hcl:"proximity_placement_group_id,attr"`
	// ScaleDownMode: string, optional
	ScaleDownMode terra.StringValue `hcl:"scale_down_mode,attr"`
	// SnapshotId: string, optional
	SnapshotId terra.StringValue `hcl:"snapshot_id,attr"`
	// SpotMaxPrice: number, optional
	SpotMaxPrice terra.NumberValue `hcl:"spot_max_price,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// UltraSsdEnabled: bool, optional
	UltraSsdEnabled terra.BoolValue `hcl:"ultra_ssd_enabled,attr"`
	// VmSize: string, required
	VmSize terra.StringValue `hcl:"vm_size,attr" validate:"required"`
	// VnetSubnetId: string, optional
	VnetSubnetId terra.StringValue `hcl:"vnet_subnet_id,attr"`
	// WorkloadRuntime: string, optional
	WorkloadRuntime terra.StringValue `hcl:"workload_runtime,attr"`
	// Zones: set of string, optional
	Zones terra.SetValue[terra.StringValue] `hcl:"zones,attr"`
	// KubeletConfig: optional
	KubeletConfig *kubernetesclusternodepool.KubeletConfig `hcl:"kubelet_config,block"`
	// LinuxOsConfig: optional
	LinuxOsConfig *kubernetesclusternodepool.LinuxOsConfig `hcl:"linux_os_config,block"`
	// NodeNetworkProfile: optional
	NodeNetworkProfile *kubernetesclusternodepool.NodeNetworkProfile `hcl:"node_network_profile,block"`
	// Timeouts: optional
	Timeouts *kubernetesclusternodepool.Timeouts `hcl:"timeouts,block"`
	// UpgradeSettings: optional
	UpgradeSettings *kubernetesclusternodepool.UpgradeSettings `hcl:"upgrade_settings,block"`
	// WindowsProfile: optional
	WindowsProfile *kubernetesclusternodepool.WindowsProfile `hcl:"windows_profile,block"`
}
type kubernetesClusterNodePoolAttributes struct {
	ref terra.Reference
}

// CapacityReservationGroupId returns a reference to field capacity_reservation_group_id of azurerm_kubernetes_cluster_node_pool.
func (kcnp kubernetesClusterNodePoolAttributes) CapacityReservationGroupId() terra.StringValue {
	return terra.ReferenceAsString(kcnp.ref.Append("capacity_reservation_group_id"))
}

// CustomCaTrustEnabled returns a reference to field custom_ca_trust_enabled of azurerm_kubernetes_cluster_node_pool.
func (kcnp kubernetesClusterNodePoolAttributes) CustomCaTrustEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(kcnp.ref.Append("custom_ca_trust_enabled"))
}

// EnableAutoScaling returns a reference to field enable_auto_scaling of azurerm_kubernetes_cluster_node_pool.
func (kcnp kubernetesClusterNodePoolAttributes) EnableAutoScaling() terra.BoolValue {
	return terra.ReferenceAsBool(kcnp.ref.Append("enable_auto_scaling"))
}

// EnableHostEncryption returns a reference to field enable_host_encryption of azurerm_kubernetes_cluster_node_pool.
func (kcnp kubernetesClusterNodePoolAttributes) EnableHostEncryption() terra.BoolValue {
	return terra.ReferenceAsBool(kcnp.ref.Append("enable_host_encryption"))
}

// EnableNodePublicIp returns a reference to field enable_node_public_ip of azurerm_kubernetes_cluster_node_pool.
func (kcnp kubernetesClusterNodePoolAttributes) EnableNodePublicIp() terra.BoolValue {
	return terra.ReferenceAsBool(kcnp.ref.Append("enable_node_public_ip"))
}

// EvictionPolicy returns a reference to field eviction_policy of azurerm_kubernetes_cluster_node_pool.
func (kcnp kubernetesClusterNodePoolAttributes) EvictionPolicy() terra.StringValue {
	return terra.ReferenceAsString(kcnp.ref.Append("eviction_policy"))
}

// FipsEnabled returns a reference to field fips_enabled of azurerm_kubernetes_cluster_node_pool.
func (kcnp kubernetesClusterNodePoolAttributes) FipsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(kcnp.ref.Append("fips_enabled"))
}

// HostGroupId returns a reference to field host_group_id of azurerm_kubernetes_cluster_node_pool.
func (kcnp kubernetesClusterNodePoolAttributes) HostGroupId() terra.StringValue {
	return terra.ReferenceAsString(kcnp.ref.Append("host_group_id"))
}

// Id returns a reference to field id of azurerm_kubernetes_cluster_node_pool.
func (kcnp kubernetesClusterNodePoolAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kcnp.ref.Append("id"))
}

// KubeletDiskType returns a reference to field kubelet_disk_type of azurerm_kubernetes_cluster_node_pool.
func (kcnp kubernetesClusterNodePoolAttributes) KubeletDiskType() terra.StringValue {
	return terra.ReferenceAsString(kcnp.ref.Append("kubelet_disk_type"))
}

// KubernetesClusterId returns a reference to field kubernetes_cluster_id of azurerm_kubernetes_cluster_node_pool.
func (kcnp kubernetesClusterNodePoolAttributes) KubernetesClusterId() terra.StringValue {
	return terra.ReferenceAsString(kcnp.ref.Append("kubernetes_cluster_id"))
}

// MaxCount returns a reference to field max_count of azurerm_kubernetes_cluster_node_pool.
func (kcnp kubernetesClusterNodePoolAttributes) MaxCount() terra.NumberValue {
	return terra.ReferenceAsNumber(kcnp.ref.Append("max_count"))
}

// MaxPods returns a reference to field max_pods of azurerm_kubernetes_cluster_node_pool.
func (kcnp kubernetesClusterNodePoolAttributes) MaxPods() terra.NumberValue {
	return terra.ReferenceAsNumber(kcnp.ref.Append("max_pods"))
}

// MessageOfTheDay returns a reference to field message_of_the_day of azurerm_kubernetes_cluster_node_pool.
func (kcnp kubernetesClusterNodePoolAttributes) MessageOfTheDay() terra.StringValue {
	return terra.ReferenceAsString(kcnp.ref.Append("message_of_the_day"))
}

// MinCount returns a reference to field min_count of azurerm_kubernetes_cluster_node_pool.
func (kcnp kubernetesClusterNodePoolAttributes) MinCount() terra.NumberValue {
	return terra.ReferenceAsNumber(kcnp.ref.Append("min_count"))
}

// Mode returns a reference to field mode of azurerm_kubernetes_cluster_node_pool.
func (kcnp kubernetesClusterNodePoolAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(kcnp.ref.Append("mode"))
}

// Name returns a reference to field name of azurerm_kubernetes_cluster_node_pool.
func (kcnp kubernetesClusterNodePoolAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kcnp.ref.Append("name"))
}

// NodeCount returns a reference to field node_count of azurerm_kubernetes_cluster_node_pool.
func (kcnp kubernetesClusterNodePoolAttributes) NodeCount() terra.NumberValue {
	return terra.ReferenceAsNumber(kcnp.ref.Append("node_count"))
}

// NodeLabels returns a reference to field node_labels of azurerm_kubernetes_cluster_node_pool.
func (kcnp kubernetesClusterNodePoolAttributes) NodeLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](kcnp.ref.Append("node_labels"))
}

// NodePublicIpPrefixId returns a reference to field node_public_ip_prefix_id of azurerm_kubernetes_cluster_node_pool.
func (kcnp kubernetesClusterNodePoolAttributes) NodePublicIpPrefixId() terra.StringValue {
	return terra.ReferenceAsString(kcnp.ref.Append("node_public_ip_prefix_id"))
}

// NodeTaints returns a reference to field node_taints of azurerm_kubernetes_cluster_node_pool.
func (kcnp kubernetesClusterNodePoolAttributes) NodeTaints() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](kcnp.ref.Append("node_taints"))
}

// OrchestratorVersion returns a reference to field orchestrator_version of azurerm_kubernetes_cluster_node_pool.
func (kcnp kubernetesClusterNodePoolAttributes) OrchestratorVersion() terra.StringValue {
	return terra.ReferenceAsString(kcnp.ref.Append("orchestrator_version"))
}

// OsDiskSizeGb returns a reference to field os_disk_size_gb of azurerm_kubernetes_cluster_node_pool.
func (kcnp kubernetesClusterNodePoolAttributes) OsDiskSizeGb() terra.NumberValue {
	return terra.ReferenceAsNumber(kcnp.ref.Append("os_disk_size_gb"))
}

// OsDiskType returns a reference to field os_disk_type of azurerm_kubernetes_cluster_node_pool.
func (kcnp kubernetesClusterNodePoolAttributes) OsDiskType() terra.StringValue {
	return terra.ReferenceAsString(kcnp.ref.Append("os_disk_type"))
}

// OsSku returns a reference to field os_sku of azurerm_kubernetes_cluster_node_pool.
func (kcnp kubernetesClusterNodePoolAttributes) OsSku() terra.StringValue {
	return terra.ReferenceAsString(kcnp.ref.Append("os_sku"))
}

// OsType returns a reference to field os_type of azurerm_kubernetes_cluster_node_pool.
func (kcnp kubernetesClusterNodePoolAttributes) OsType() terra.StringValue {
	return terra.ReferenceAsString(kcnp.ref.Append("os_type"))
}

// PodSubnetId returns a reference to field pod_subnet_id of azurerm_kubernetes_cluster_node_pool.
func (kcnp kubernetesClusterNodePoolAttributes) PodSubnetId() terra.StringValue {
	return terra.ReferenceAsString(kcnp.ref.Append("pod_subnet_id"))
}

// Priority returns a reference to field priority of azurerm_kubernetes_cluster_node_pool.
func (kcnp kubernetesClusterNodePoolAttributes) Priority() terra.StringValue {
	return terra.ReferenceAsString(kcnp.ref.Append("priority"))
}

// ProximityPlacementGroupId returns a reference to field proximity_placement_group_id of azurerm_kubernetes_cluster_node_pool.
func (kcnp kubernetesClusterNodePoolAttributes) ProximityPlacementGroupId() terra.StringValue {
	return terra.ReferenceAsString(kcnp.ref.Append("proximity_placement_group_id"))
}

// ScaleDownMode returns a reference to field scale_down_mode of azurerm_kubernetes_cluster_node_pool.
func (kcnp kubernetesClusterNodePoolAttributes) ScaleDownMode() terra.StringValue {
	return terra.ReferenceAsString(kcnp.ref.Append("scale_down_mode"))
}

// SnapshotId returns a reference to field snapshot_id of azurerm_kubernetes_cluster_node_pool.
func (kcnp kubernetesClusterNodePoolAttributes) SnapshotId() terra.StringValue {
	return terra.ReferenceAsString(kcnp.ref.Append("snapshot_id"))
}

// SpotMaxPrice returns a reference to field spot_max_price of azurerm_kubernetes_cluster_node_pool.
func (kcnp kubernetesClusterNodePoolAttributes) SpotMaxPrice() terra.NumberValue {
	return terra.ReferenceAsNumber(kcnp.ref.Append("spot_max_price"))
}

// Tags returns a reference to field tags of azurerm_kubernetes_cluster_node_pool.
func (kcnp kubernetesClusterNodePoolAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](kcnp.ref.Append("tags"))
}

// UltraSsdEnabled returns a reference to field ultra_ssd_enabled of azurerm_kubernetes_cluster_node_pool.
func (kcnp kubernetesClusterNodePoolAttributes) UltraSsdEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(kcnp.ref.Append("ultra_ssd_enabled"))
}

// VmSize returns a reference to field vm_size of azurerm_kubernetes_cluster_node_pool.
func (kcnp kubernetesClusterNodePoolAttributes) VmSize() terra.StringValue {
	return terra.ReferenceAsString(kcnp.ref.Append("vm_size"))
}

// VnetSubnetId returns a reference to field vnet_subnet_id of azurerm_kubernetes_cluster_node_pool.
func (kcnp kubernetesClusterNodePoolAttributes) VnetSubnetId() terra.StringValue {
	return terra.ReferenceAsString(kcnp.ref.Append("vnet_subnet_id"))
}

// WorkloadRuntime returns a reference to field workload_runtime of azurerm_kubernetes_cluster_node_pool.
func (kcnp kubernetesClusterNodePoolAttributes) WorkloadRuntime() terra.StringValue {
	return terra.ReferenceAsString(kcnp.ref.Append("workload_runtime"))
}

// Zones returns a reference to field zones of azurerm_kubernetes_cluster_node_pool.
func (kcnp kubernetesClusterNodePoolAttributes) Zones() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](kcnp.ref.Append("zones"))
}

func (kcnp kubernetesClusterNodePoolAttributes) KubeletConfig() terra.ListValue[kubernetesclusternodepool.KubeletConfigAttributes] {
	return terra.ReferenceAsList[kubernetesclusternodepool.KubeletConfigAttributes](kcnp.ref.Append("kubelet_config"))
}

func (kcnp kubernetesClusterNodePoolAttributes) LinuxOsConfig() terra.ListValue[kubernetesclusternodepool.LinuxOsConfigAttributes] {
	return terra.ReferenceAsList[kubernetesclusternodepool.LinuxOsConfigAttributes](kcnp.ref.Append("linux_os_config"))
}

func (kcnp kubernetesClusterNodePoolAttributes) NodeNetworkProfile() terra.ListValue[kubernetesclusternodepool.NodeNetworkProfileAttributes] {
	return terra.ReferenceAsList[kubernetesclusternodepool.NodeNetworkProfileAttributes](kcnp.ref.Append("node_network_profile"))
}

func (kcnp kubernetesClusterNodePoolAttributes) Timeouts() kubernetesclusternodepool.TimeoutsAttributes {
	return terra.ReferenceAsSingle[kubernetesclusternodepool.TimeoutsAttributes](kcnp.ref.Append("timeouts"))
}

func (kcnp kubernetesClusterNodePoolAttributes) UpgradeSettings() terra.ListValue[kubernetesclusternodepool.UpgradeSettingsAttributes] {
	return terra.ReferenceAsList[kubernetesclusternodepool.UpgradeSettingsAttributes](kcnp.ref.Append("upgrade_settings"))
}

func (kcnp kubernetesClusterNodePoolAttributes) WindowsProfile() terra.ListValue[kubernetesclusternodepool.WindowsProfileAttributes] {
	return terra.ReferenceAsList[kubernetesclusternodepool.WindowsProfileAttributes](kcnp.ref.Append("windows_profile"))
}

type kubernetesClusterNodePoolState struct {
	CapacityReservationGroupId string                                              `json:"capacity_reservation_group_id"`
	CustomCaTrustEnabled       bool                                                `json:"custom_ca_trust_enabled"`
	EnableAutoScaling          bool                                                `json:"enable_auto_scaling"`
	EnableHostEncryption       bool                                                `json:"enable_host_encryption"`
	EnableNodePublicIp         bool                                                `json:"enable_node_public_ip"`
	EvictionPolicy             string                                              `json:"eviction_policy"`
	FipsEnabled                bool                                                `json:"fips_enabled"`
	HostGroupId                string                                              `json:"host_group_id"`
	Id                         string                                              `json:"id"`
	KubeletDiskType            string                                              `json:"kubelet_disk_type"`
	KubernetesClusterId        string                                              `json:"kubernetes_cluster_id"`
	MaxCount                   float64                                             `json:"max_count"`
	MaxPods                    float64                                             `json:"max_pods"`
	MessageOfTheDay            string                                              `json:"message_of_the_day"`
	MinCount                   float64                                             `json:"min_count"`
	Mode                       string                                              `json:"mode"`
	Name                       string                                              `json:"name"`
	NodeCount                  float64                                             `json:"node_count"`
	NodeLabels                 map[string]string                                   `json:"node_labels"`
	NodePublicIpPrefixId       string                                              `json:"node_public_ip_prefix_id"`
	NodeTaints                 []string                                            `json:"node_taints"`
	OrchestratorVersion        string                                              `json:"orchestrator_version"`
	OsDiskSizeGb               float64                                             `json:"os_disk_size_gb"`
	OsDiskType                 string                                              `json:"os_disk_type"`
	OsSku                      string                                              `json:"os_sku"`
	OsType                     string                                              `json:"os_type"`
	PodSubnetId                string                                              `json:"pod_subnet_id"`
	Priority                   string                                              `json:"priority"`
	ProximityPlacementGroupId  string                                              `json:"proximity_placement_group_id"`
	ScaleDownMode              string                                              `json:"scale_down_mode"`
	SnapshotId                 string                                              `json:"snapshot_id"`
	SpotMaxPrice               float64                                             `json:"spot_max_price"`
	Tags                       map[string]string                                   `json:"tags"`
	UltraSsdEnabled            bool                                                `json:"ultra_ssd_enabled"`
	VmSize                     string                                              `json:"vm_size"`
	VnetSubnetId               string                                              `json:"vnet_subnet_id"`
	WorkloadRuntime            string                                              `json:"workload_runtime"`
	Zones                      []string                                            `json:"zones"`
	KubeletConfig              []kubernetesclusternodepool.KubeletConfigState      `json:"kubelet_config"`
	LinuxOsConfig              []kubernetesclusternodepool.LinuxOsConfigState      `json:"linux_os_config"`
	NodeNetworkProfile         []kubernetesclusternodepool.NodeNetworkProfileState `json:"node_network_profile"`
	Timeouts                   *kubernetesclusternodepool.TimeoutsState            `json:"timeouts"`
	UpgradeSettings            []kubernetesclusternodepool.UpgradeSettingsState    `json:"upgrade_settings"`
	WindowsProfile             []kubernetesclusternodepool.WindowsProfileState     `json:"windows_profile"`
}
