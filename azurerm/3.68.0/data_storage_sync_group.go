// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datastoragesyncgroup "github.com/golingon/terraproviders/azurerm/3.68.0/datastoragesyncgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataStorageSyncGroup creates a new instance of [DataStorageSyncGroup].
func NewDataStorageSyncGroup(name string, args DataStorageSyncGroupArgs) *DataStorageSyncGroup {
	return &DataStorageSyncGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataStorageSyncGroup)(nil)

// DataStorageSyncGroup represents the Terraform data resource azurerm_storage_sync_group.
type DataStorageSyncGroup struct {
	Name string
	Args DataStorageSyncGroupArgs
}

// DataSource returns the Terraform object type for [DataStorageSyncGroup].
func (ssg *DataStorageSyncGroup) DataSource() string {
	return "azurerm_storage_sync_group"
}

// LocalName returns the local name for [DataStorageSyncGroup].
func (ssg *DataStorageSyncGroup) LocalName() string {
	return ssg.Name
}

// Configuration returns the configuration (args) for [DataStorageSyncGroup].
func (ssg *DataStorageSyncGroup) Configuration() interface{} {
	return ssg.Args
}

// Attributes returns the attributes for [DataStorageSyncGroup].
func (ssg *DataStorageSyncGroup) Attributes() dataStorageSyncGroupAttributes {
	return dataStorageSyncGroupAttributes{ref: terra.ReferenceDataResource(ssg)}
}

// DataStorageSyncGroupArgs contains the configurations for azurerm_storage_sync_group.
type DataStorageSyncGroupArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// StorageSyncId: string, required
	StorageSyncId terra.StringValue `hcl:"storage_sync_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datastoragesyncgroup.Timeouts `hcl:"timeouts,block"`
}
type dataStorageSyncGroupAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_storage_sync_group.
func (ssg dataStorageSyncGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ssg.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_storage_sync_group.
func (ssg dataStorageSyncGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ssg.ref.Append("name"))
}

// StorageSyncId returns a reference to field storage_sync_id of azurerm_storage_sync_group.
func (ssg dataStorageSyncGroupAttributes) StorageSyncId() terra.StringValue {
	return terra.ReferenceAsString(ssg.ref.Append("storage_sync_id"))
}

func (ssg dataStorageSyncGroupAttributes) Timeouts() datastoragesyncgroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datastoragesyncgroup.TimeoutsAttributes](ssg.ref.Append("timeouts"))
}
