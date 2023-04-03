// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datavirtualnetwork "github.com/golingon/terraproviders/azurerm/3.49.0/datavirtualnetwork"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataVirtualNetwork creates a new instance of [DataVirtualNetwork].
func NewDataVirtualNetwork(name string, args DataVirtualNetworkArgs) *DataVirtualNetwork {
	return &DataVirtualNetwork{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataVirtualNetwork)(nil)

// DataVirtualNetwork represents the Terraform data resource azurerm_virtual_network.
type DataVirtualNetwork struct {
	Name string
	Args DataVirtualNetworkArgs
}

// DataSource returns the Terraform object type for [DataVirtualNetwork].
func (vn *DataVirtualNetwork) DataSource() string {
	return "azurerm_virtual_network"
}

// LocalName returns the local name for [DataVirtualNetwork].
func (vn *DataVirtualNetwork) LocalName() string {
	return vn.Name
}

// Configuration returns the configuration (args) for [DataVirtualNetwork].
func (vn *DataVirtualNetwork) Configuration() interface{} {
	return vn.Args
}

// Attributes returns the attributes for [DataVirtualNetwork].
func (vn *DataVirtualNetwork) Attributes() dataVirtualNetworkAttributes {
	return dataVirtualNetworkAttributes{ref: terra.ReferenceDataResource(vn)}
}

// DataVirtualNetworkArgs contains the configurations for azurerm_virtual_network.
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

// AddressSpace returns a reference to field address_space of azurerm_virtual_network.
func (vn dataVirtualNetworkAttributes) AddressSpace() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vn.ref.Append("address_space"))
}

// DnsServers returns a reference to field dns_servers of azurerm_virtual_network.
func (vn dataVirtualNetworkAttributes) DnsServers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vn.ref.Append("dns_servers"))
}

// Guid returns a reference to field guid of azurerm_virtual_network.
func (vn dataVirtualNetworkAttributes) Guid() terra.StringValue {
	return terra.ReferenceAsString(vn.ref.Append("guid"))
}

// Id returns a reference to field id of azurerm_virtual_network.
func (vn dataVirtualNetworkAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vn.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_virtual_network.
func (vn dataVirtualNetworkAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(vn.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_virtual_network.
func (vn dataVirtualNetworkAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vn.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_virtual_network.
func (vn dataVirtualNetworkAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(vn.ref.Append("resource_group_name"))
}

// Subnets returns a reference to field subnets of azurerm_virtual_network.
func (vn dataVirtualNetworkAttributes) Subnets() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vn.ref.Append("subnets"))
}

// Tags returns a reference to field tags of azurerm_virtual_network.
func (vn dataVirtualNetworkAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vn.ref.Append("tags"))
}

// VnetPeerings returns a reference to field vnet_peerings of azurerm_virtual_network.
func (vn dataVirtualNetworkAttributes) VnetPeerings() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vn.ref.Append("vnet_peerings"))
}

// VnetPeeringsAddresses returns a reference to field vnet_peerings_addresses of azurerm_virtual_network.
func (vn dataVirtualNetworkAttributes) VnetPeeringsAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vn.ref.Append("vnet_peerings_addresses"))
}

func (vn dataVirtualNetworkAttributes) Timeouts() datavirtualnetwork.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datavirtualnetwork.TimeoutsAttributes](vn.ref.Append("timeouts"))
}
