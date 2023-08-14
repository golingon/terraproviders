// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datanetappvolume "github.com/golingon/terraproviders/azurerm/3.69.0/datanetappvolume"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataNetappVolume creates a new instance of [DataNetappVolume].
func NewDataNetappVolume(name string, args DataNetappVolumeArgs) *DataNetappVolume {
	return &DataNetappVolume{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataNetappVolume)(nil)

// DataNetappVolume represents the Terraform data resource azurerm_netapp_volume.
type DataNetappVolume struct {
	Name string
	Args DataNetappVolumeArgs
}

// DataSource returns the Terraform object type for [DataNetappVolume].
func (nv *DataNetappVolume) DataSource() string {
	return "azurerm_netapp_volume"
}

// LocalName returns the local name for [DataNetappVolume].
func (nv *DataNetappVolume) LocalName() string {
	return nv.Name
}

// Configuration returns the configuration (args) for [DataNetappVolume].
func (nv *DataNetappVolume) Configuration() interface{} {
	return nv.Args
}

// Attributes returns the attributes for [DataNetappVolume].
func (nv *DataNetappVolume) Attributes() dataNetappVolumeAttributes {
	return dataNetappVolumeAttributes{ref: terra.ReferenceDataResource(nv)}
}

// DataNetappVolumeArgs contains the configurations for azurerm_netapp_volume.
type DataNetappVolumeArgs struct {
	// AccountName: string, required
	AccountName terra.StringValue `hcl:"account_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PoolName: string, required
	PoolName terra.StringValue `hcl:"pool_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SecurityStyle: string, optional
	SecurityStyle terra.StringValue `hcl:"security_style,attr"`
	// DataProtectionReplication: min=0
	DataProtectionReplication []datanetappvolume.DataProtectionReplication `hcl:"data_protection_replication,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datanetappvolume.Timeouts `hcl:"timeouts,block"`
}
type dataNetappVolumeAttributes struct {
	ref terra.Reference
}

// AccountName returns a reference to field account_name of azurerm_netapp_volume.
func (nv dataNetappVolumeAttributes) AccountName() terra.StringValue {
	return terra.ReferenceAsString(nv.ref.Append("account_name"))
}

// Id returns a reference to field id of azurerm_netapp_volume.
func (nv dataNetappVolumeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nv.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_netapp_volume.
func (nv dataNetappVolumeAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(nv.ref.Append("location"))
}

// MountIpAddresses returns a reference to field mount_ip_addresses of azurerm_netapp_volume.
func (nv dataNetappVolumeAttributes) MountIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nv.ref.Append("mount_ip_addresses"))
}

// Name returns a reference to field name of azurerm_netapp_volume.
func (nv dataNetappVolumeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nv.ref.Append("name"))
}

// NetworkFeatures returns a reference to field network_features of azurerm_netapp_volume.
func (nv dataNetappVolumeAttributes) NetworkFeatures() terra.StringValue {
	return terra.ReferenceAsString(nv.ref.Append("network_features"))
}

// PoolName returns a reference to field pool_name of azurerm_netapp_volume.
func (nv dataNetappVolumeAttributes) PoolName() terra.StringValue {
	return terra.ReferenceAsString(nv.ref.Append("pool_name"))
}

// Protocols returns a reference to field protocols of azurerm_netapp_volume.
func (nv dataNetappVolumeAttributes) Protocols() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nv.ref.Append("protocols"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_netapp_volume.
func (nv dataNetappVolumeAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(nv.ref.Append("resource_group_name"))
}

// SecurityStyle returns a reference to field security_style of azurerm_netapp_volume.
func (nv dataNetappVolumeAttributes) SecurityStyle() terra.StringValue {
	return terra.ReferenceAsString(nv.ref.Append("security_style"))
}

// ServiceLevel returns a reference to field service_level of azurerm_netapp_volume.
func (nv dataNetappVolumeAttributes) ServiceLevel() terra.StringValue {
	return terra.ReferenceAsString(nv.ref.Append("service_level"))
}

// StorageQuotaInGb returns a reference to field storage_quota_in_gb of azurerm_netapp_volume.
func (nv dataNetappVolumeAttributes) StorageQuotaInGb() terra.NumberValue {
	return terra.ReferenceAsNumber(nv.ref.Append("storage_quota_in_gb"))
}

// SubnetId returns a reference to field subnet_id of azurerm_netapp_volume.
func (nv dataNetappVolumeAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(nv.ref.Append("subnet_id"))
}

// VolumePath returns a reference to field volume_path of azurerm_netapp_volume.
func (nv dataNetappVolumeAttributes) VolumePath() terra.StringValue {
	return terra.ReferenceAsString(nv.ref.Append("volume_path"))
}

// Zone returns a reference to field zone of azurerm_netapp_volume.
func (nv dataNetappVolumeAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(nv.ref.Append("zone"))
}

func (nv dataNetappVolumeAttributes) DataProtectionReplication() terra.ListValue[datanetappvolume.DataProtectionReplicationAttributes] {
	return terra.ReferenceAsList[datanetappvolume.DataProtectionReplicationAttributes](nv.ref.Append("data_protection_replication"))
}

func (nv dataNetappVolumeAttributes) Timeouts() datanetappvolume.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datanetappvolume.TimeoutsAttributes](nv.ref.Append("timeouts"))
}