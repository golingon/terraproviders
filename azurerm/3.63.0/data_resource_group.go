// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataresourcegroup "github.com/golingon/terraproviders/azurerm/3.63.0/dataresourcegroup"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataResourceGroup creates a new instance of [DataResourceGroup].
func NewDataResourceGroup(name string, args DataResourceGroupArgs) *DataResourceGroup {
	return &DataResourceGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataResourceGroup)(nil)

// DataResourceGroup represents the Terraform data resource azurerm_resource_group.
type DataResourceGroup struct {
	Name string
	Args DataResourceGroupArgs
}

// DataSource returns the Terraform object type for [DataResourceGroup].
func (rg *DataResourceGroup) DataSource() string {
	return "azurerm_resource_group"
}

// LocalName returns the local name for [DataResourceGroup].
func (rg *DataResourceGroup) LocalName() string {
	return rg.Name
}

// Configuration returns the configuration (args) for [DataResourceGroup].
func (rg *DataResourceGroup) Configuration() interface{} {
	return rg.Args
}

// Attributes returns the attributes for [DataResourceGroup].
func (rg *DataResourceGroup) Attributes() dataResourceGroupAttributes {
	return dataResourceGroupAttributes{ref: terra.ReferenceDataResource(rg)}
}

// DataResourceGroupArgs contains the configurations for azurerm_resource_group.
type DataResourceGroupArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dataresourcegroup.Timeouts `hcl:"timeouts,block"`
}
type dataResourceGroupAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_resource_group.
func (rg dataResourceGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rg.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_resource_group.
func (rg dataResourceGroupAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(rg.ref.Append("location"))
}

// ManagedBy returns a reference to field managed_by of azurerm_resource_group.
func (rg dataResourceGroupAttributes) ManagedBy() terra.StringValue {
	return terra.ReferenceAsString(rg.ref.Append("managed_by"))
}

// Name returns a reference to field name of azurerm_resource_group.
func (rg dataResourceGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rg.ref.Append("name"))
}

// Tags returns a reference to field tags of azurerm_resource_group.
func (rg dataResourceGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rg.ref.Append("tags"))
}

func (rg dataResourceGroupAttributes) Timeouts() dataresourcegroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataresourcegroup.TimeoutsAttributes](rg.ref.Append("timeouts"))
}
