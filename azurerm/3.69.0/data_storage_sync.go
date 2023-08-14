// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datastoragesync "github.com/golingon/terraproviders/azurerm/3.69.0/datastoragesync"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataStorageSync creates a new instance of [DataStorageSync].
func NewDataStorageSync(name string, args DataStorageSyncArgs) *DataStorageSync {
	return &DataStorageSync{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataStorageSync)(nil)

// DataStorageSync represents the Terraform data resource azurerm_storage_sync.
type DataStorageSync struct {
	Name string
	Args DataStorageSyncArgs
}

// DataSource returns the Terraform object type for [DataStorageSync].
func (ss *DataStorageSync) DataSource() string {
	return "azurerm_storage_sync"
}

// LocalName returns the local name for [DataStorageSync].
func (ss *DataStorageSync) LocalName() string {
	return ss.Name
}

// Configuration returns the configuration (args) for [DataStorageSync].
func (ss *DataStorageSync) Configuration() interface{} {
	return ss.Args
}

// Attributes returns the attributes for [DataStorageSync].
func (ss *DataStorageSync) Attributes() dataStorageSyncAttributes {
	return dataStorageSyncAttributes{ref: terra.ReferenceDataResource(ss)}
}

// DataStorageSyncArgs contains the configurations for azurerm_storage_sync.
type DataStorageSyncArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datastoragesync.Timeouts `hcl:"timeouts,block"`
}
type dataStorageSyncAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_storage_sync.
func (ss dataStorageSyncAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("id"))
}

// IncomingTrafficPolicy returns a reference to field incoming_traffic_policy of azurerm_storage_sync.
func (ss dataStorageSyncAttributes) IncomingTrafficPolicy() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("incoming_traffic_policy"))
}

// Location returns a reference to field location of azurerm_storage_sync.
func (ss dataStorageSyncAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_storage_sync.
func (ss dataStorageSyncAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_storage_sync.
func (ss dataStorageSyncAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_storage_sync.
func (ss dataStorageSyncAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ss.ref.Append("tags"))
}

func (ss dataStorageSyncAttributes) Timeouts() datastoragesync.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datastoragesync.TimeoutsAttributes](ss.ref.Append("timeouts"))
}
