// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datavirtualnetwork "github.com/golingon/terraproviders/azurerm/3.49.0/datavirtualnetwork"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataVirtualNetwork(name string, args DataVirtualNetworkArgs) *DataVirtualNetwork {
	return &DataVirtualNetwork{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataVirtualNetwork)(nil)

type DataVirtualNetwork struct {
	Name string
	Args DataVirtualNetworkArgs
}

func (vn *DataVirtualNetwork) DataSource() string {
	return "azurerm_virtual_network"
}

func (vn *DataVirtualNetwork) LocalName() string {
	return vn.Name
}

func (vn *DataVirtualNetwork) Configuration() interface{} {
	return vn.Args
}

func (vn *DataVirtualNetwork) Attributes() dataVirtualNetworkAttributes {
	return dataVirtualNetworkAttributes{ref: terra.ReferenceDataResource(vn)}
}

type DataVirtualNetworkArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datavirtualnetwork.Timeouts `hcl:"timeouts,block"`
}
type dataVirtualNetworkAttributes struct {
	ref terra.Reference
}

func (vn dataVirtualNetworkAttributes) AddressSpace() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](vn.ref.Append("address_space"))
}

func (vn dataVirtualNetworkAttributes) DnsServers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](vn.ref.Append("dns_servers"))
}

func (vn dataVirtualNetworkAttributes) Guid() terra.StringValue {
	return terra.ReferenceString(vn.ref.Append("guid"))
}

func (vn dataVirtualNetworkAttributes) Id() terra.StringValue {
	return terra.ReferenceString(vn.ref.Append("id"))
}

func (vn dataVirtualNetworkAttributes) Location() terra.StringValue {
	return terra.ReferenceString(vn.ref.Append("location"))
}

func (vn dataVirtualNetworkAttributes) Name() terra.StringValue {
	return terra.ReferenceString(vn.ref.Append("name"))
}

func (vn dataVirtualNetworkAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(vn.ref.Append("resource_group_name"))
}

func (vn dataVirtualNetworkAttributes) Subnets() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](vn.ref.Append("subnets"))
}

func (vn dataVirtualNetworkAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](vn.ref.Append("tags"))
}

func (vn dataVirtualNetworkAttributes) VnetPeerings() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](vn.ref.Append("vnet_peerings"))
}

func (vn dataVirtualNetworkAttributes) VnetPeeringsAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](vn.ref.Append("vnet_peerings_addresses"))
}

func (vn dataVirtualNetworkAttributes) Timeouts() datavirtualnetwork.TimeoutsAttributes {
	return terra.ReferenceSingle[datavirtualnetwork.TimeoutsAttributes](vn.ref.Append("timeouts"))
}
