// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datavpngateway "github.com/golingon/terraproviders/azurerm/3.64.0/datavpngateway"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataVpnGateway creates a new instance of [DataVpnGateway].
func NewDataVpnGateway(name string, args DataVpnGatewayArgs) *DataVpnGateway {
	return &DataVpnGateway{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataVpnGateway)(nil)

// DataVpnGateway represents the Terraform data resource azurerm_vpn_gateway.
type DataVpnGateway struct {
	Name string
	Args DataVpnGatewayArgs
}

// DataSource returns the Terraform object type for [DataVpnGateway].
func (vg *DataVpnGateway) DataSource() string {
	return "azurerm_vpn_gateway"
}

// LocalName returns the local name for [DataVpnGateway].
func (vg *DataVpnGateway) LocalName() string {
	return vg.Name
}

// Configuration returns the configuration (args) for [DataVpnGateway].
func (vg *DataVpnGateway) Configuration() interface{} {
	return vg.Args
}

// Attributes returns the attributes for [DataVpnGateway].
func (vg *DataVpnGateway) Attributes() dataVpnGatewayAttributes {
	return dataVpnGatewayAttributes{ref: terra.ReferenceDataResource(vg)}
}

// DataVpnGatewayArgs contains the configurations for azurerm_vpn_gateway.
type DataVpnGatewayArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// BgpSettings: min=0
	BgpSettings []datavpngateway.BgpSettings `hcl:"bgp_settings,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datavpngateway.Timeouts `hcl:"timeouts,block"`
}
type dataVpnGatewayAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_vpn_gateway.
func (vg dataVpnGatewayAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vg.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_vpn_gateway.
func (vg dataVpnGatewayAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(vg.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_vpn_gateway.
func (vg dataVpnGatewayAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vg.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_vpn_gateway.
func (vg dataVpnGatewayAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(vg.ref.Append("resource_group_name"))
}

// ScaleUnit returns a reference to field scale_unit of azurerm_vpn_gateway.
func (vg dataVpnGatewayAttributes) ScaleUnit() terra.NumberValue {
	return terra.ReferenceAsNumber(vg.ref.Append("scale_unit"))
}

// Tags returns a reference to field tags of azurerm_vpn_gateway.
func (vg dataVpnGatewayAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vg.ref.Append("tags"))
}

// VirtualHubId returns a reference to field virtual_hub_id of azurerm_vpn_gateway.
func (vg dataVpnGatewayAttributes) VirtualHubId() terra.StringValue {
	return terra.ReferenceAsString(vg.ref.Append("virtual_hub_id"))
}

func (vg dataVpnGatewayAttributes) BgpSettings() terra.ListValue[datavpngateway.BgpSettingsAttributes] {
	return terra.ReferenceAsList[datavpngateway.BgpSettingsAttributes](vg.ref.Append("bgp_settings"))
}

func (vg dataVpnGatewayAttributes) Timeouts() datavpngateway.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datavpngateway.TimeoutsAttributes](vg.ref.Append("timeouts"))
}
