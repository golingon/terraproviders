// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	netappvolume "github.com/golingon/terraproviders/azurerm/3.63.0/netappvolume"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetappVolume creates a new instance of [NetappVolume].
func NewNetappVolume(name string, args NetappVolumeArgs) *NetappVolume {
	return &NetappVolume{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetappVolume)(nil)

// NetappVolume represents the Terraform resource azurerm_netapp_volume.
type NetappVolume struct {
	Name      string
	Args      NetappVolumeArgs
	state     *netappVolumeState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetappVolume].
func (nv *NetappVolume) Type() string {
	return "azurerm_netapp_volume"
}

// LocalName returns the local name for [NetappVolume].
func (nv *NetappVolume) LocalName() string {
	return nv.Name
}

// Configuration returns the configuration (args) for [NetappVolume].
func (nv *NetappVolume) Configuration() interface{} {
	return nv.Args
}

// DependOn is used for other resources to depend on [NetappVolume].
func (nv *NetappVolume) DependOn() terra.Reference {
	return terra.ReferenceResource(nv)
}

// Dependencies returns the list of resources [NetappVolume] depends_on.
func (nv *NetappVolume) Dependencies() terra.Dependencies {
	return nv.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetappVolume].
func (nv *NetappVolume) LifecycleManagement() *terra.Lifecycle {
	return nv.Lifecycle
}

// Attributes returns the attributes for [NetappVolume].
func (nv *NetappVolume) Attributes() netappVolumeAttributes {
	return netappVolumeAttributes{ref: terra.ReferenceResource(nv)}
}

// ImportState imports the given attribute values into [NetappVolume]'s state.
func (nv *NetappVolume) ImportState(av io.Reader) error {
	nv.state = &netappVolumeState{}
	if err := json.NewDecoder(av).Decode(nv.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nv.Type(), nv.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetappVolume] has state.
func (nv *NetappVolume) State() (*netappVolumeState, bool) {
	return nv.state, nv.state != nil
}

// StateMust returns the state for [NetappVolume]. Panics if the state is nil.
func (nv *NetappVolume) StateMust() *netappVolumeState {
	if nv.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nv.Type(), nv.LocalName()))
	}
	return nv.state
}

// NetappVolumeArgs contains the configurations for azurerm_netapp_volume.
type NetappVolumeArgs struct {
	// AccountName: string, required
	AccountName terra.StringValue `hcl:"account_name,attr" validate:"required"`
	// AzureVmwareDataStoreEnabled: bool, optional
	AzureVmwareDataStoreEnabled terra.BoolValue `hcl:"azure_vmware_data_store_enabled,attr"`
	// CreateFromSnapshotResourceId: string, optional
	CreateFromSnapshotResourceId terra.StringValue `hcl:"create_from_snapshot_resource_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NetworkFeatures: string, optional
	NetworkFeatures terra.StringValue `hcl:"network_features,attr"`
	// PoolName: string, required
	PoolName terra.StringValue `hcl:"pool_name,attr" validate:"required"`
	// Protocols: set of string, optional
	Protocols terra.SetValue[terra.StringValue] `hcl:"protocols,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SecurityStyle: string, optional
	SecurityStyle terra.StringValue `hcl:"security_style,attr"`
	// ServiceLevel: string, required
	ServiceLevel terra.StringValue `hcl:"service_level,attr" validate:"required"`
	// SnapshotDirectoryVisible: bool, optional
	SnapshotDirectoryVisible terra.BoolValue `hcl:"snapshot_directory_visible,attr"`
	// StorageQuotaInGb: number, required
	StorageQuotaInGb terra.NumberValue `hcl:"storage_quota_in_gb,attr" validate:"required"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// ThroughputInMibps: number, optional
	ThroughputInMibps terra.NumberValue `hcl:"throughput_in_mibps,attr"`
	// VolumePath: string, required
	VolumePath terra.StringValue `hcl:"volume_path,attr" validate:"required"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// DataProtectionReplication: optional
	DataProtectionReplication *netappvolume.DataProtectionReplication `hcl:"data_protection_replication,block"`
	// DataProtectionSnapshotPolicy: optional
	DataProtectionSnapshotPolicy *netappvolume.DataProtectionSnapshotPolicy `hcl:"data_protection_snapshot_policy,block"`
	// ExportPolicyRule: min=0,max=5
	ExportPolicyRule []netappvolume.ExportPolicyRule `hcl:"export_policy_rule,block" validate:"min=0,max=5"`
	// Timeouts: optional
	Timeouts *netappvolume.Timeouts `hcl:"timeouts,block"`
}
type netappVolumeAttributes struct {
	ref terra.Reference
}

// AccountName returns a reference to field account_name of azurerm_netapp_volume.
func (nv netappVolumeAttributes) AccountName() terra.StringValue {
	return terra.ReferenceAsString(nv.ref.Append("account_name"))
}

// AzureVmwareDataStoreEnabled returns a reference to field azure_vmware_data_store_enabled of azurerm_netapp_volume.
func (nv netappVolumeAttributes) AzureVmwareDataStoreEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(nv.ref.Append("azure_vmware_data_store_enabled"))
}

// CreateFromSnapshotResourceId returns a reference to field create_from_snapshot_resource_id of azurerm_netapp_volume.
func (nv netappVolumeAttributes) CreateFromSnapshotResourceId() terra.StringValue {
	return terra.ReferenceAsString(nv.ref.Append("create_from_snapshot_resource_id"))
}

// Id returns a reference to field id of azurerm_netapp_volume.
func (nv netappVolumeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nv.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_netapp_volume.
func (nv netappVolumeAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(nv.ref.Append("location"))
}

// MountIpAddresses returns a reference to field mount_ip_addresses of azurerm_netapp_volume.
func (nv netappVolumeAttributes) MountIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nv.ref.Append("mount_ip_addresses"))
}

// Name returns a reference to field name of azurerm_netapp_volume.
func (nv netappVolumeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nv.ref.Append("name"))
}

// NetworkFeatures returns a reference to field network_features of azurerm_netapp_volume.
func (nv netappVolumeAttributes) NetworkFeatures() terra.StringValue {
	return terra.ReferenceAsString(nv.ref.Append("network_features"))
}

// PoolName returns a reference to field pool_name of azurerm_netapp_volume.
func (nv netappVolumeAttributes) PoolName() terra.StringValue {
	return terra.ReferenceAsString(nv.ref.Append("pool_name"))
}

// Protocols returns a reference to field protocols of azurerm_netapp_volume.
func (nv netappVolumeAttributes) Protocols() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](nv.ref.Append("protocols"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_netapp_volume.
func (nv netappVolumeAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(nv.ref.Append("resource_group_name"))
}

// SecurityStyle returns a reference to field security_style of azurerm_netapp_volume.
func (nv netappVolumeAttributes) SecurityStyle() terra.StringValue {
	return terra.ReferenceAsString(nv.ref.Append("security_style"))
}

// ServiceLevel returns a reference to field service_level of azurerm_netapp_volume.
func (nv netappVolumeAttributes) ServiceLevel() terra.StringValue {
	return terra.ReferenceAsString(nv.ref.Append("service_level"))
}

// SnapshotDirectoryVisible returns a reference to field snapshot_directory_visible of azurerm_netapp_volume.
func (nv netappVolumeAttributes) SnapshotDirectoryVisible() terra.BoolValue {
	return terra.ReferenceAsBool(nv.ref.Append("snapshot_directory_visible"))
}

// StorageQuotaInGb returns a reference to field storage_quota_in_gb of azurerm_netapp_volume.
func (nv netappVolumeAttributes) StorageQuotaInGb() terra.NumberValue {
	return terra.ReferenceAsNumber(nv.ref.Append("storage_quota_in_gb"))
}

// SubnetId returns a reference to field subnet_id of azurerm_netapp_volume.
func (nv netappVolumeAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(nv.ref.Append("subnet_id"))
}

// Tags returns a reference to field tags of azurerm_netapp_volume.
func (nv netappVolumeAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nv.ref.Append("tags"))
}

// ThroughputInMibps returns a reference to field throughput_in_mibps of azurerm_netapp_volume.
func (nv netappVolumeAttributes) ThroughputInMibps() terra.NumberValue {
	return terra.ReferenceAsNumber(nv.ref.Append("throughput_in_mibps"))
}

// VolumePath returns a reference to field volume_path of azurerm_netapp_volume.
func (nv netappVolumeAttributes) VolumePath() terra.StringValue {
	return terra.ReferenceAsString(nv.ref.Append("volume_path"))
}

// Zone returns a reference to field zone of azurerm_netapp_volume.
func (nv netappVolumeAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(nv.ref.Append("zone"))
}

func (nv netappVolumeAttributes) DataProtectionReplication() terra.ListValue[netappvolume.DataProtectionReplicationAttributes] {
	return terra.ReferenceAsList[netappvolume.DataProtectionReplicationAttributes](nv.ref.Append("data_protection_replication"))
}

func (nv netappVolumeAttributes) DataProtectionSnapshotPolicy() terra.ListValue[netappvolume.DataProtectionSnapshotPolicyAttributes] {
	return terra.ReferenceAsList[netappvolume.DataProtectionSnapshotPolicyAttributes](nv.ref.Append("data_protection_snapshot_policy"))
}

func (nv netappVolumeAttributes) ExportPolicyRule() terra.ListValue[netappvolume.ExportPolicyRuleAttributes] {
	return terra.ReferenceAsList[netappvolume.ExportPolicyRuleAttributes](nv.ref.Append("export_policy_rule"))
}

func (nv netappVolumeAttributes) Timeouts() netappvolume.TimeoutsAttributes {
	return terra.ReferenceAsSingle[netappvolume.TimeoutsAttributes](nv.ref.Append("timeouts"))
}

type netappVolumeState struct {
	AccountName                  string                                           `json:"account_name"`
	AzureVmwareDataStoreEnabled  bool                                             `json:"azure_vmware_data_store_enabled"`
	CreateFromSnapshotResourceId string                                           `json:"create_from_snapshot_resource_id"`
	Id                           string                                           `json:"id"`
	Location                     string                                           `json:"location"`
	MountIpAddresses             []string                                         `json:"mount_ip_addresses"`
	Name                         string                                           `json:"name"`
	NetworkFeatures              string                                           `json:"network_features"`
	PoolName                     string                                           `json:"pool_name"`
	Protocols                    []string                                         `json:"protocols"`
	ResourceGroupName            string                                           `json:"resource_group_name"`
	SecurityStyle                string                                           `json:"security_style"`
	ServiceLevel                 string                                           `json:"service_level"`
	SnapshotDirectoryVisible     bool                                             `json:"snapshot_directory_visible"`
	StorageQuotaInGb             float64                                          `json:"storage_quota_in_gb"`
	SubnetId                     string                                           `json:"subnet_id"`
	Tags                         map[string]string                                `json:"tags"`
	ThroughputInMibps            float64                                          `json:"throughput_in_mibps"`
	VolumePath                   string                                           `json:"volume_path"`
	Zone                         string                                           `json:"zone"`
	DataProtectionReplication    []netappvolume.DataProtectionReplicationState    `json:"data_protection_replication"`
	DataProtectionSnapshotPolicy []netappvolume.DataProtectionSnapshotPolicyState `json:"data_protection_snapshot_policy"`
	ExportPolicyRule             []netappvolume.ExportPolicyRuleState             `json:"export_policy_rule"`
	Timeouts                     *netappvolume.TimeoutsState                      `json:"timeouts"`
}
