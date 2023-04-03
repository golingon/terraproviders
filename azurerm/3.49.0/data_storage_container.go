// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datastoragecontainer "github.com/golingon/terraproviders/azurerm/3.49.0/datastoragecontainer"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataStorageContainer creates a new instance of [DataStorageContainer].
func NewDataStorageContainer(name string, args DataStorageContainerArgs) *DataStorageContainer {
	return &DataStorageContainer{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataStorageContainer)(nil)

// DataStorageContainer represents the Terraform data resource azurerm_storage_container.
type DataStorageContainer struct {
	Name string
	Args DataStorageContainerArgs
}

// DataSource returns the Terraform object type for [DataStorageContainer].
func (sc *DataStorageContainer) DataSource() string {
	return "azurerm_storage_container"
}

// LocalName returns the local name for [DataStorageContainer].
func (sc *DataStorageContainer) LocalName() string {
	return sc.Name
}

// Configuration returns the configuration (args) for [DataStorageContainer].
func (sc *DataStorageContainer) Configuration() interface{} {
	return sc.Args
}

// Attributes returns the attributes for [DataStorageContainer].
func (sc *DataStorageContainer) Attributes() dataStorageContainerAttributes {
	return dataStorageContainerAttributes{ref: terra.ReferenceDataResource(sc)}
}

// DataStorageContainerArgs contains the configurations for azurerm_storage_container.
type DataStorageContainerArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Metadata: map of string, optional
	Metadata terra.MapValue[terra.StringValue] `hcl:"metadata,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// StorageAccountName: string, required
	StorageAccountName terra.StringValue `hcl:"storage_account_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datastoragecontainer.Timeouts `hcl:"timeouts,block"`
}
type dataStorageContainerAttributes struct {
	ref terra.Reference
}

// ContainerAccessType returns a reference to field container_access_type of azurerm_storage_container.
func (sc dataStorageContainerAttributes) ContainerAccessType() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("container_access_type"))
}

// HasImmutabilityPolicy returns a reference to field has_immutability_policy of azurerm_storage_container.
func (sc dataStorageContainerAttributes) HasImmutabilityPolicy() terra.BoolValue {
	return terra.ReferenceAsBool(sc.ref.Append("has_immutability_policy"))
}

// HasLegalHold returns a reference to field has_legal_hold of azurerm_storage_container.
func (sc dataStorageContainerAttributes) HasLegalHold() terra.BoolValue {
	return terra.ReferenceAsBool(sc.ref.Append("has_legal_hold"))
}

// Id returns a reference to field id of azurerm_storage_container.
func (sc dataStorageContainerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("id"))
}

// Metadata returns a reference to field metadata of azurerm_storage_container.
func (sc dataStorageContainerAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sc.ref.Append("metadata"))
}

// Name returns a reference to field name of azurerm_storage_container.
func (sc dataStorageContainerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("name"))
}

// ResourceManagerId returns a reference to field resource_manager_id of azurerm_storage_container.
func (sc dataStorageContainerAttributes) ResourceManagerId() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("resource_manager_id"))
}

// StorageAccountName returns a reference to field storage_account_name of azurerm_storage_container.
func (sc dataStorageContainerAttributes) StorageAccountName() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("storage_account_name"))
}

func (sc dataStorageContainerAttributes) Timeouts() datastoragecontainer.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datastoragecontainer.TimeoutsAttributes](sc.ref.Append("timeouts"))
}
