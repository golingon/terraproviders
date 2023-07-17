// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	storagetableentity "github.com/golingon/terraproviders/azurerm/3.65.0/storagetableentity"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStorageTableEntity creates a new instance of [StorageTableEntity].
func NewStorageTableEntity(name string, args StorageTableEntityArgs) *StorageTableEntity {
	return &StorageTableEntity{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StorageTableEntity)(nil)

// StorageTableEntity represents the Terraform resource azurerm_storage_table_entity.
type StorageTableEntity struct {
	Name      string
	Args      StorageTableEntityArgs
	state     *storageTableEntityState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StorageTableEntity].
func (ste *StorageTableEntity) Type() string {
	return "azurerm_storage_table_entity"
}

// LocalName returns the local name for [StorageTableEntity].
func (ste *StorageTableEntity) LocalName() string {
	return ste.Name
}

// Configuration returns the configuration (args) for [StorageTableEntity].
func (ste *StorageTableEntity) Configuration() interface{} {
	return ste.Args
}

// DependOn is used for other resources to depend on [StorageTableEntity].
func (ste *StorageTableEntity) DependOn() terra.Reference {
	return terra.ReferenceResource(ste)
}

// Dependencies returns the list of resources [StorageTableEntity] depends_on.
func (ste *StorageTableEntity) Dependencies() terra.Dependencies {
	return ste.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StorageTableEntity].
func (ste *StorageTableEntity) LifecycleManagement() *terra.Lifecycle {
	return ste.Lifecycle
}

// Attributes returns the attributes for [StorageTableEntity].
func (ste *StorageTableEntity) Attributes() storageTableEntityAttributes {
	return storageTableEntityAttributes{ref: terra.ReferenceResource(ste)}
}

// ImportState imports the given attribute values into [StorageTableEntity]'s state.
func (ste *StorageTableEntity) ImportState(av io.Reader) error {
	ste.state = &storageTableEntityState{}
	if err := json.NewDecoder(av).Decode(ste.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ste.Type(), ste.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StorageTableEntity] has state.
func (ste *StorageTableEntity) State() (*storageTableEntityState, bool) {
	return ste.state, ste.state != nil
}

// StateMust returns the state for [StorageTableEntity]. Panics if the state is nil.
func (ste *StorageTableEntity) StateMust() *storageTableEntityState {
	if ste.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ste.Type(), ste.LocalName()))
	}
	return ste.state
}

// StorageTableEntityArgs contains the configurations for azurerm_storage_table_entity.
type StorageTableEntityArgs struct {
	// Entity: map of string, required
	Entity terra.MapValue[terra.StringValue] `hcl:"entity,attr" validate:"required"`
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
	Timeouts *storagetableentity.Timeouts `hcl:"timeouts,block"`
}
type storageTableEntityAttributes struct {
	ref terra.Reference
}

// Entity returns a reference to field entity of azurerm_storage_table_entity.
func (ste storageTableEntityAttributes) Entity() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ste.ref.Append("entity"))
}

// Id returns a reference to field id of azurerm_storage_table_entity.
func (ste storageTableEntityAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ste.ref.Append("id"))
}

// PartitionKey returns a reference to field partition_key of azurerm_storage_table_entity.
func (ste storageTableEntityAttributes) PartitionKey() terra.StringValue {
	return terra.ReferenceAsString(ste.ref.Append("partition_key"))
}

// RowKey returns a reference to field row_key of azurerm_storage_table_entity.
func (ste storageTableEntityAttributes) RowKey() terra.StringValue {
	return terra.ReferenceAsString(ste.ref.Append("row_key"))
}

// StorageAccountName returns a reference to field storage_account_name of azurerm_storage_table_entity.
func (ste storageTableEntityAttributes) StorageAccountName() terra.StringValue {
	return terra.ReferenceAsString(ste.ref.Append("storage_account_name"))
}

// TableName returns a reference to field table_name of azurerm_storage_table_entity.
func (ste storageTableEntityAttributes) TableName() terra.StringValue {
	return terra.ReferenceAsString(ste.ref.Append("table_name"))
}

func (ste storageTableEntityAttributes) Timeouts() storagetableentity.TimeoutsAttributes {
	return terra.ReferenceAsSingle[storagetableentity.TimeoutsAttributes](ste.ref.Append("timeouts"))
}

type storageTableEntityState struct {
	Entity             map[string]string                 `json:"entity"`
	Id                 string                            `json:"id"`
	PartitionKey       string                            `json:"partition_key"`
	RowKey             string                            `json:"row_key"`
	StorageAccountName string                            `json:"storage_account_name"`
	TableName          string                            `json:"table_name"`
	Timeouts           *storagetableentity.TimeoutsState `json:"timeouts"`
}