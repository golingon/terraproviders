// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datalocalnetworkgateway "github.com/golingon/terraproviders/azurerm/3.67.0/datalocalnetworkgateway"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataLocalNetworkGateway creates a new instance of [DataLocalNetworkGateway].
func NewDataLocalNetworkGateway(name string, args DataLocalNetworkGatewayArgs) *DataLocalNetworkGateway {
	return &DataLocalNetworkGateway{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataLocalNetworkGateway)(nil)

// DataLocalNetworkGateway represents the Terraform data resource azurerm_local_network_gateway.
type DataLocalNetworkGateway struct {
	Name string
	Args DataLocalNetworkGatewayArgs
}

// DataSource returns the Terraform object type for [DataLocalNetworkGateway].
func (lng *DataLocalNetworkGateway) DataSource() string {
	return "azurerm_local_network_gateway"
}

// LocalName returns the local name for [DataLocalNetworkGateway].
func (lng *DataLocalNetworkGateway) LocalName() string {
	return lng.Name
}

// Configuration returns the configuration (args) for [DataLocalNetworkGateway].
func (lng *DataLocalNetworkGateway) Configuration() interface{} {
	return lng.Args
}

// Attributes returns the attributes for [DataLocalNetworkGateway].
func (lng *DataLocalNetworkGateway) Attributes() dataLocalNetworkGatewayAttributes {
	return dataLocalNetworkGatewayAttributes{ref: terra.ReferenceDataResource(lng)}
}

// DataLocalNetworkGatewayArgs contains the configurations for azurerm_local_network_gateway.
type DataLocalNetworkGatewayArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// BgpSettings: min=0
	BgpSettings []datalocalnetworkgateway.BgpSettings `hcl:"bgp_settings,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datalocalnetworkgateway.Timeouts `hcl:"timeouts,block"`
}
type dataLocalNetworkGatewayAttributes struct {
	ref terra.Reference
}

// AddressSpace returns a reference to field address_space of azurerm_local_network_gateway.
func (lng dataLocalNetworkGatewayAttributes) AddressSpace() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](lng.ref.Append("address_space"))
}

// GatewayAddress returns a reference to field gateway_address of azurerm_local_network_gateway.
func (lng dataLocalNetworkGatewayAttributes) GatewayAddress() terra.StringValue {
	return terra.ReferenceAsString(lng.ref.Append("gateway_address"))
}

// GatewayFqdn returns a reference to field gateway_fqdn of azurerm_local_network_gateway.
func (lng dataLocalNetworkGatewayAttributes) GatewayFqdn() terra.StringValue {
	return terra.ReferenceAsString(lng.ref.Append("gateway_fqdn"))
}

// Id returns a reference to field id of azurerm_local_network_gateway.
func (lng dataLocalNetworkGatewayAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lng.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_local_network_gateway.
func (lng dataLocalNetworkGatewayAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(lng.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_local_network_gateway.
func (lng dataLocalNetworkGatewayAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lng.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_local_network_gateway.
func (lng dataLocalNetworkGatewayAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(lng.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_local_network_gateway.
func (lng dataLocalNetworkGatewayAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lng.ref.Append("tags"))
}

func (lng dataLocalNetworkGatewayAttributes) BgpSettings() terra.ListValue[datalocalnetworkgateway.BgpSettingsAttributes] {
	return terra.ReferenceAsList[datalocalnetworkgateway.BgpSettingsAttributes](lng.ref.Append("bgp_settings"))
}

func (lng dataLocalNetworkGatewayAttributes) Timeouts() datalocalnetworkgateway.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datalocalnetworkgateway.TimeoutsAttributes](lng.ref.Append("timeouts"))
}
