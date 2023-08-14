// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datacontainerregistrytoken "github.com/golingon/terraproviders/azurerm/3.69.0/datacontainerregistrytoken"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataContainerRegistryToken creates a new instance of [DataContainerRegistryToken].
func NewDataContainerRegistryToken(name string, args DataContainerRegistryTokenArgs) *DataContainerRegistryToken {
	return &DataContainerRegistryToken{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataContainerRegistryToken)(nil)

// DataContainerRegistryToken represents the Terraform data resource azurerm_container_registry_token.
type DataContainerRegistryToken struct {
	Name string
	Args DataContainerRegistryTokenArgs
}

// DataSource returns the Terraform object type for [DataContainerRegistryToken].
func (crt *DataContainerRegistryToken) DataSource() string {
	return "azurerm_container_registry_token"
}

// LocalName returns the local name for [DataContainerRegistryToken].
func (crt *DataContainerRegistryToken) LocalName() string {
	return crt.Name
}

// Configuration returns the configuration (args) for [DataContainerRegistryToken].
func (crt *DataContainerRegistryToken) Configuration() interface{} {
	return crt.Args
}

// Attributes returns the attributes for [DataContainerRegistryToken].
func (crt *DataContainerRegistryToken) Attributes() dataContainerRegistryTokenAttributes {
	return dataContainerRegistryTokenAttributes{ref: terra.ReferenceDataResource(crt)}
}

// DataContainerRegistryTokenArgs contains the configurations for azurerm_container_registry_token.
type DataContainerRegistryTokenArgs struct {
	// ContainerRegistryName: string, required
	ContainerRegistryName terra.StringValue `hcl:"container_registry_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datacontainerregistrytoken.Timeouts `hcl:"timeouts,block"`
}
type dataContainerRegistryTokenAttributes struct {
	ref terra.Reference
}

// ContainerRegistryName returns a reference to field container_registry_name of azurerm_container_registry_token.
func (crt dataContainerRegistryTokenAttributes) ContainerRegistryName() terra.StringValue {
	return terra.ReferenceAsString(crt.ref.Append("container_registry_name"))
}

// Enabled returns a reference to field enabled of azurerm_container_registry_token.
func (crt dataContainerRegistryTokenAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(crt.ref.Append("enabled"))
}

// Id returns a reference to field id of azurerm_container_registry_token.
func (crt dataContainerRegistryTokenAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crt.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_container_registry_token.
func (crt dataContainerRegistryTokenAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crt.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_container_registry_token.
func (crt dataContainerRegistryTokenAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(crt.ref.Append("resource_group_name"))
}

// ScopeMapId returns a reference to field scope_map_id of azurerm_container_registry_token.
func (crt dataContainerRegistryTokenAttributes) ScopeMapId() terra.StringValue {
	return terra.ReferenceAsString(crt.ref.Append("scope_map_id"))
}

func (crt dataContainerRegistryTokenAttributes) Timeouts() datacontainerregistrytoken.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datacontainerregistrytoken.TimeoutsAttributes](crt.ref.Append("timeouts"))
}
