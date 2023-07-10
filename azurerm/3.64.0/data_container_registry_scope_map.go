// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datacontainerregistryscopemap "github.com/golingon/terraproviders/azurerm/3.64.0/datacontainerregistryscopemap"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataContainerRegistryScopeMap creates a new instance of [DataContainerRegistryScopeMap].
func NewDataContainerRegistryScopeMap(name string, args DataContainerRegistryScopeMapArgs) *DataContainerRegistryScopeMap {
	return &DataContainerRegistryScopeMap{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataContainerRegistryScopeMap)(nil)

// DataContainerRegistryScopeMap represents the Terraform data resource azurerm_container_registry_scope_map.
type DataContainerRegistryScopeMap struct {
	Name string
	Args DataContainerRegistryScopeMapArgs
}

// DataSource returns the Terraform object type for [DataContainerRegistryScopeMap].
func (crsm *DataContainerRegistryScopeMap) DataSource() string {
	return "azurerm_container_registry_scope_map"
}

// LocalName returns the local name for [DataContainerRegistryScopeMap].
func (crsm *DataContainerRegistryScopeMap) LocalName() string {
	return crsm.Name
}

// Configuration returns the configuration (args) for [DataContainerRegistryScopeMap].
func (crsm *DataContainerRegistryScopeMap) Configuration() interface{} {
	return crsm.Args
}

// Attributes returns the attributes for [DataContainerRegistryScopeMap].
func (crsm *DataContainerRegistryScopeMap) Attributes() dataContainerRegistryScopeMapAttributes {
	return dataContainerRegistryScopeMapAttributes{ref: terra.ReferenceDataResource(crsm)}
}

// DataContainerRegistryScopeMapArgs contains the configurations for azurerm_container_registry_scope_map.
type DataContainerRegistryScopeMapArgs struct {
	// ContainerRegistryName: string, required
	ContainerRegistryName terra.StringValue `hcl:"container_registry_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datacontainerregistryscopemap.Timeouts `hcl:"timeouts,block"`
}
type dataContainerRegistryScopeMapAttributes struct {
	ref terra.Reference
}

// Actions returns a reference to field actions of azurerm_container_registry_scope_map.
func (crsm dataContainerRegistryScopeMapAttributes) Actions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](crsm.ref.Append("actions"))
}

// ContainerRegistryName returns a reference to field container_registry_name of azurerm_container_registry_scope_map.
func (crsm dataContainerRegistryScopeMapAttributes) ContainerRegistryName() terra.StringValue {
	return terra.ReferenceAsString(crsm.ref.Append("container_registry_name"))
}

// Description returns a reference to field description of azurerm_container_registry_scope_map.
func (crsm dataContainerRegistryScopeMapAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(crsm.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_container_registry_scope_map.
func (crsm dataContainerRegistryScopeMapAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crsm.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_container_registry_scope_map.
func (crsm dataContainerRegistryScopeMapAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crsm.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_container_registry_scope_map.
func (crsm dataContainerRegistryScopeMapAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(crsm.ref.Append("resource_group_name"))
}

func (crsm dataContainerRegistryScopeMapAttributes) Timeouts() datacontainerregistryscopemap.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datacontainerregistryscopemap.TimeoutsAttributes](crsm.ref.Append("timeouts"))
}
