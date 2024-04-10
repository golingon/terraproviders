// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"github.com/golingon/lingon/pkg/terra"
	datavirtualhub "github.com/golingon/terraproviders/azurerm/3.98.0/datavirtualhub"
)

// NewDataVirtualHub creates a new instance of [DataVirtualHub].
func NewDataVirtualHub(name string, args DataVirtualHubArgs) *DataVirtualHub {
	return &DataVirtualHub{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataVirtualHub)(nil)

// DataVirtualHub represents the Terraform data resource azurerm_virtual_hub.
type DataVirtualHub struct {
	Name string
	Args DataVirtualHubArgs
}

// DataSource returns the Terraform object type for [DataVirtualHub].
func (vh *DataVirtualHub) DataSource() string {
	return "azurerm_virtual_hub"
}

// LocalName returns the local name for [DataVirtualHub].
func (vh *DataVirtualHub) LocalName() string {
	return vh.Name
}

// Configuration returns the configuration (args) for [DataVirtualHub].
func (vh *DataVirtualHub) Configuration() interface{} {
	return vh.Args
}

// Attributes returns the attributes for [DataVirtualHub].
func (vh *DataVirtualHub) Attributes() dataVirtualHubAttributes {
	return dataVirtualHubAttributes{ref: terra.ReferenceDataResource(vh)}
}

// DataVirtualHubArgs contains the configurations for azurerm_virtual_hub.
type DataVirtualHubArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datavirtualhub.Timeouts `hcl:"timeouts,block"`
}
type dataVirtualHubAttributes struct {
	ref terra.Reference
}

// AddressPrefix returns a reference to field address_prefix of azurerm_virtual_hub.
func (vh dataVirtualHubAttributes) AddressPrefix() terra.StringValue {
	return terra.ReferenceAsString(vh.ref.Append("address_prefix"))
}

// DefaultRouteTableId returns a reference to field default_route_table_id of azurerm_virtual_hub.
func (vh dataVirtualHubAttributes) DefaultRouteTableId() terra.StringValue {
	return terra.ReferenceAsString(vh.ref.Append("default_route_table_id"))
}

// Id returns a reference to field id of azurerm_virtual_hub.
func (vh dataVirtualHubAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vh.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_virtual_hub.
func (vh dataVirtualHubAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(vh.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_virtual_hub.
func (vh dataVirtualHubAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vh.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_virtual_hub.
func (vh dataVirtualHubAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(vh.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_virtual_hub.
func (vh dataVirtualHubAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vh.ref.Append("tags"))
}

// VirtualRouterAsn returns a reference to field virtual_router_asn of azurerm_virtual_hub.
func (vh dataVirtualHubAttributes) VirtualRouterAsn() terra.NumberValue {
	return terra.ReferenceAsNumber(vh.ref.Append("virtual_router_asn"))
}

// VirtualRouterIps returns a reference to field virtual_router_ips of azurerm_virtual_hub.
func (vh dataVirtualHubAttributes) VirtualRouterIps() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vh.ref.Append("virtual_router_ips"))
}

// VirtualWanId returns a reference to field virtual_wan_id of azurerm_virtual_hub.
func (vh dataVirtualHubAttributes) VirtualWanId() terra.StringValue {
	return terra.ReferenceAsString(vh.ref.Append("virtual_wan_id"))
}

func (vh dataVirtualHubAttributes) Timeouts() datavirtualhub.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datavirtualhub.TimeoutsAttributes](vh.ref.Append("timeouts"))
}
