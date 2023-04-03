// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataipgroup "github.com/golingon/terraproviders/azurerm/3.49.0/dataipgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataIpGroup creates a new instance of [DataIpGroup].
func NewDataIpGroup(name string, args DataIpGroupArgs) *DataIpGroup {
	return &DataIpGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataIpGroup)(nil)

// DataIpGroup represents the Terraform data resource azurerm_ip_group.
type DataIpGroup struct {
	Name string
	Args DataIpGroupArgs
}

// DataSource returns the Terraform object type for [DataIpGroup].
func (ig *DataIpGroup) DataSource() string {
	return "azurerm_ip_group"
}

// LocalName returns the local name for [DataIpGroup].
func (ig *DataIpGroup) LocalName() string {
	return ig.Name
}

// Configuration returns the configuration (args) for [DataIpGroup].
func (ig *DataIpGroup) Configuration() interface{} {
	return ig.Args
}

// Attributes returns the attributes for [DataIpGroup].
func (ig *DataIpGroup) Attributes() dataIpGroupAttributes {
	return dataIpGroupAttributes{ref: terra.ReferenceDataResource(ig)}
}

// DataIpGroupArgs contains the configurations for azurerm_ip_group.
type DataIpGroupArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dataipgroup.Timeouts `hcl:"timeouts,block"`
}
type dataIpGroupAttributes struct {
	ref terra.Reference
}

// Cidrs returns a reference to field cidrs of azurerm_ip_group.
func (ig dataIpGroupAttributes) Cidrs() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ig.ref.Append("cidrs"))
}

// Id returns a reference to field id of azurerm_ip_group.
func (ig dataIpGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_ip_group.
func (ig dataIpGroupAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_ip_group.
func (ig dataIpGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_ip_group.
func (ig dataIpGroupAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_ip_group.
func (ig dataIpGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ig.ref.Append("tags"))
}

func (ig dataIpGroupAttributes) Timeouts() dataipgroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataipgroup.TimeoutsAttributes](ig.ref.Append("timeouts"))
}