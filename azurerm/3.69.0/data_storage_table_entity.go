// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datastoragetableentity "github.com/golingon/terraproviders/azurerm/3.69.0/datastoragetableentity"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataStorageTableEntity creates a new instance of [DataStorageTableEntity].
func NewDataStorageTableEntity(name string, args DataStorageTableEntityArgs) *DataStorageTableEntity {
	return &DataStorageTableEntity{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataStorageTableEntity)(nil)

// DataStorageTableEntity represents the Terraform data resource azurerm_storage_table_entity.
type DataStorageTableEntity struct {
	Name string
	Args DataStorageTableEntityArgs
}

// DataSource returns the Terraform object type for [DataStorageTableEntity].
func (ste *DataStorageTableEntity) DataSource() string {
	return "azurerm_storage_table_entity"
}

// LocalName returns the local name for [DataStorageTableEntity].
func (ste *DataStorageTableEntity) LocalName() string {
	return ste.Name
}

// Configuration returns the configuration (args) for [DataStorageTableEntity].
func (ste *DataStorageTableEntity) Configuration() interface{} {
	return ste.Args
}

// Attributes returns the attributes for [DataStorageTableEntity].
func (ste *DataStorageTableEntity) Attributes() dataStorageTableEntityAttributes {
	return dataStorageTableEntityAttributes{ref: terra.ReferenceDataResource(ste)}
}

// DataStorageTableEntityArgs contains the configurations for azurerm_storage_table_entity.
type DataStorageTableEntityArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PartitionKey: string, required
	PartitionKey terra.StringValue `hcl:"partition_key,attr" validate:"required"`
	// RowKey: string, required
	RowKey terra.StringValue `hcl:"row_key,attr" validate:"required"`
	// StorageAccountName: string, required
	StorageAccountName terra.StringValue `hcl:"storage_account_name,attr" validate:"required"`
	// TableName: string, required
	TableName terra.StringValue `hcl:"table_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datastoragetableentity.Timeouts `hcl:"timeouts,block"`
}
type dataStorageTableEntityAttributes struct {
	ref terra.Reference
}

// Entity returns a reference to field entity of azurerm_storage_table_entity.
func (ste dataStorageTableEntityAttributes) Entity() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ste.ref.Append("entity"))
}

// Id returns a reference to field id of azurerm_storage_table_entity.
func (ste dataStorageTableEntityAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ste.ref.Append("id"))
}

// PartitionKey returns a reference to field partition_key of azurerm_storage_table_entity.
func (ste dataStorageTableEntityAttributes) PartitionKey() terra.StringValue {
	return terra.ReferenceAsString(ste.ref.Append("partition_key"))
}

// RowKey returns a reference to field row_key of azurerm_storage_table_entity.
func (ste dataStorageTableEntityAttributes) RowKey() terra.StringValue {
	return terra.ReferenceAsString(ste.ref.Append("row_key"))
}

// StorageAccountName returns a reference to field storage_account_name of azurerm_storage_table_entity.
func (ste dataStorageTableEntityAttributes) StorageAccountName() terra.StringValue {
	return terra.ReferenceAsString(ste.ref.Append("storage_account_name"))
}

// TableName returns a reference to field table_name of azurerm_storage_table_entity.
func (ste dataStorageTableEntityAttributes) TableName() terra.StringValue {
	return terra.ReferenceAsString(ste.ref.Append("table_name"))
}

func (ste dataStorageTableEntityAttributes) Timeouts() datastoragetableentity.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datastoragetableentity.TimeoutsAttributes](ste.ref.Append("timeouts"))
}
