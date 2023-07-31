// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataroutetable "github.com/golingon/terraproviders/azurerm/3.67.0/dataroutetable"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataRouteTable creates a new instance of [DataRouteTable].
func NewDataRouteTable(name string, args DataRouteTableArgs) *DataRouteTable {
	return &DataRouteTable{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataRouteTable)(nil)

// DataRouteTable represents the Terraform data resource azurerm_route_table.
type DataRouteTable struct {
	Name string
	Args DataRouteTableArgs
}

// DataSource returns the Terraform object type for [DataRouteTable].
func (rt *DataRouteTable) DataSource() string {
	return "azurerm_route_table"
}

// LocalName returns the local name for [DataRouteTable].
func (rt *DataRouteTable) LocalName() string {
	return rt.Name
}

// Configuration returns the configuration (args) for [DataRouteTable].
func (rt *DataRouteTable) Configuration() interface{} {
	return rt.Args
}

// Attributes returns the attributes for [DataRouteTable].
func (rt *DataRouteTable) Attributes() dataRouteTableAttributes {
	return dataRouteTableAttributes{ref: terra.ReferenceDataResource(rt)}
}

// DataRouteTableArgs contains the configurations for azurerm_route_table.
type DataRouteTableArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Route: min=0
	Route []dataroutetable.Route `hcl:"route,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataroutetable.Timeouts `hcl:"timeouts,block"`
}
type dataRouteTableAttributes struct {
	ref terra.Reference
}

// BgpRoutePropagationEnabled returns a reference to field bgp_route_propagation_enabled of azurerm_route_table.
func (rt dataRouteTableAttributes) BgpRoutePropagationEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(rt.ref.Append("bgp_route_propagation_enabled"))
}

// Id returns a reference to field id of azurerm_route_table.
func (rt dataRouteTableAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rt.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_route_table.
func (rt dataRouteTableAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(rt.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_route_table.
func (rt dataRouteTableAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rt.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_route_table.
func (rt dataRouteTableAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(rt.ref.Append("resource_group_name"))
}

// Subnets returns a reference to field subnets of azurerm_route_table.
func (rt dataRouteTableAttributes) Subnets() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rt.ref.Append("subnets"))
}

// Tags returns a reference to field tags of azurerm_route_table.
func (rt dataRouteTableAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rt.ref.Append("tags"))
}

func (rt dataRouteTableAttributes) Route() terra.ListValue[dataroutetable.RouteAttributes] {
	return terra.ReferenceAsList[dataroutetable.RouteAttributes](rt.ref.Append("route"))
}

func (rt dataRouteTableAttributes) Timeouts() dataroutetable.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataroutetable.TimeoutsAttributes](rt.ref.Append("timeouts"))
}
