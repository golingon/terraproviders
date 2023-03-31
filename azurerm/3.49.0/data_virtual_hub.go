// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datavirtualhub "github.com/golingon/terraproviders/azurerm/3.49.0/datavirtualhub"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataVirtualHub(name string, args DataVirtualHubArgs) *DataVirtualHub {
	return &DataVirtualHub{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataVirtualHub)(nil)

type DataVirtualHub struct {
	Name string
	Args DataVirtualHubArgs
}

func (vh *DataVirtualHub) DataSource() string {
	return "azurerm_virtual_hub"
}

func (vh *DataVirtualHub) LocalName() string {
	return vh.Name
}

func (vh *DataVirtualHub) Configuration() interface{} {
	return vh.Args
}

func (vh *DataVirtualHub) Attributes() dataVirtualHubAttributes {
	return dataVirtualHubAttributes{ref: terra.ReferenceDataResource(vh)}
}

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

func (vh dataVirtualHubAttributes) AddressPrefix() terra.StringValue {
	return terra.ReferenceString(vh.ref.Append("address_prefix"))
}

func (vh dataVirtualHubAttributes) DefaultRouteTableId() terra.StringValue {
	return terra.ReferenceString(vh.ref.Append("default_route_table_id"))
}

func (vh dataVirtualHubAttributes) Id() terra.StringValue {
	return terra.ReferenceString(vh.ref.Append("id"))
}

func (vh dataVirtualHubAttributes) Location() terra.StringValue {
	return terra.ReferenceString(vh.ref.Append("location"))
}

func (vh dataVirtualHubAttributes) Name() terra.StringValue {
	return terra.ReferenceString(vh.ref.Append("name"))
}

func (vh dataVirtualHubAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(vh.ref.Append("resource_group_name"))
}

func (vh dataVirtualHubAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](vh.ref.Append("tags"))
}

func (vh dataVirtualHubAttributes) VirtualRouterAsn() terra.NumberValue {
	return terra.ReferenceNumber(vh.ref.Append("virtual_router_asn"))
}

func (vh dataVirtualHubAttributes) VirtualRouterIps() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](vh.ref.Append("virtual_router_ips"))
}

func (vh dataVirtualHubAttributes) VirtualWanId() terra.StringValue {
	return terra.ReferenceString(vh.ref.Append("virtual_wan_id"))
}

func (vh dataVirtualHubAttributes) Timeouts() datavirtualhub.TimeoutsAttributes {
	return terra.ReferenceSingle[datavirtualhub.TimeoutsAttributes](vh.ref.Append("timeouts"))
}
