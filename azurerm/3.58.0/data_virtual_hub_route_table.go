// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datavirtualhubroutetable "github.com/golingon/terraproviders/azurerm/3.58.0/datavirtualhubroutetable"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataVirtualHubRouteTable creates a new instance of [DataVirtualHubRouteTable].
func NewDataVirtualHubRouteTable(name string, args DataVirtualHubRouteTableArgs) *DataVirtualHubRouteTable {
	return &DataVirtualHubRouteTable{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataVirtualHubRouteTable)(nil)

// DataVirtualHubRouteTable represents the Terraform data resource azurerm_virtual_hub_route_table.
type DataVirtualHubRouteTable struct {
	Name string
	Args DataVirtualHubRouteTableArgs
}

// DataSource returns the Terraform object type for [DataVirtualHubRouteTable].
func (vhrt *DataVirtualHubRouteTable) DataSource() string {
	return "azurerm_virtual_hub_route_table"
}

// LocalName returns the local name for [DataVirtualHubRouteTable].
func (vhrt *DataVirtualHubRouteTable) LocalName() string {
	return vhrt.Name
}

// Configuration returns the configuration (args) for [DataVirtualHubRouteTable].
func (vhrt *DataVirtualHubRouteTable) Configuration() interface{} {
	return vhrt.Args
}

// Attributes returns the attributes for [DataVirtualHubRouteTable].
func (vhrt *DataVirtualHubRouteTable) Attributes() dataVirtualHubRouteTableAttributes {
	return dataVirtualHubRouteTableAttributes{ref: terra.ReferenceDataResource(vhrt)}
}

// DataVirtualHubRouteTableArgs contains the configurations for azurerm_virtual_hub_route_table.
type DataVirtualHubRouteTableArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// VirtualHubName: string, required
	VirtualHubName terra.StringValue `hcl:"virtual_hub_name,attr" validate:"required"`
	// Route: min=0
	Route []datavirtualhubroutetable.Route `hcl:"route,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datavirtualhubroutetable.Timeouts `hcl:"timeouts,block"`
}
type dataVirtualHubRouteTableAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_virtual_hub_route_table.
func (vhrt dataVirtualHubRouteTableAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vhrt.ref.Append("id"))
}

// Labels returns a reference to field labels of azurerm_virtual_hub_route_table.
func (vhrt dataVirtualHubRouteTableAttributes) Labels() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vhrt.ref.Append("labels"))
}

// Name returns a reference to field name of azurerm_virtual_hub_route_table.
func (vhrt dataVirtualHubRouteTableAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vhrt.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_virtual_hub_route_table.
func (vhrt dataVirtualHubRouteTableAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(vhrt.ref.Append("resource_group_name"))
}

// VirtualHubId returns a reference to field virtual_hub_id of azurerm_virtual_hub_route_table.
func (vhrt dataVirtualHubRouteTableAttributes) VirtualHubId() terra.StringValue {
	return terra.ReferenceAsString(vhrt.ref.Append("virtual_hub_id"))
}

// VirtualHubName returns a reference to field virtual_hub_name of azurerm_virtual_hub_route_table.
func (vhrt dataVirtualHubRouteTableAttributes) VirtualHubName() terra.StringValue {
	return terra.ReferenceAsString(vhrt.ref.Append("virtual_hub_name"))
}

func (vhrt dataVirtualHubRouteTableAttributes) Route() terra.ListValue[datavirtualhubroutetable.RouteAttributes] {
	return terra.ReferenceAsList[datavirtualhubroutetable.RouteAttributes](vhrt.ref.Append("route"))
}

func (vhrt dataVirtualHubRouteTableAttributes) Timeouts() datavirtualhubroutetable.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datavirtualhubroutetable.TimeoutsAttributes](vhrt.ref.Append("timeouts"))
}
