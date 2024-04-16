// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_storage_sync_group

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurerm_storage_sync_group.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (assg *DataSource) DataSource() string {
	return "azurerm_storage_sync_group"
}

// LocalName returns the local name for [DataSource].
func (assg *DataSource) LocalName() string {
	return assg.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (assg *DataSource) Configuration() interface{} {
	return assg.Args
}

// Attributes returns the attributes for [DataSource].
func (assg *DataSource) Attributes() dataAzurermStorageSyncGroupAttributes {
	return dataAzurermStorageSyncGroupAttributes{ref: terra.ReferenceDataSource(assg)}
}

// DataArgs contains the configurations for azurerm_storage_sync_group.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// StorageSyncId: string, required
	StorageSyncId terra.StringValue `hcl:"storage_sync_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAzurermStorageSyncGroupAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_storage_sync_group.
func (assg dataAzurermStorageSyncGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(assg.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_storage_sync_group.
func (assg dataAzurermStorageSyncGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(assg.ref.Append("name"))
}

// StorageSyncId returns a reference to field storage_sync_id of azurerm_storage_sync_group.
func (assg dataAzurermStorageSyncGroupAttributes) StorageSyncId() terra.StringValue {
	return terra.ReferenceAsString(assg.ref.Append("storage_sync_id"))
}

func (assg dataAzurermStorageSyncGroupAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](assg.ref.Append("timeouts"))
}
