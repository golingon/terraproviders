// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	vpngateway "github.com/golingon/terraproviders/azurerm/3.69.0/vpngateway"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVpnGateway creates a new instance of [VpnGateway].
func NewVpnGateway(name string, args VpnGatewayArgs) *VpnGateway {
	return &VpnGateway{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpnGateway)(nil)

// VpnGateway represents the Terraform resource azurerm_vpn_gateway.
type VpnGateway struct {
	Name      string
	Args      VpnGatewayArgs
	state     *vpnGatewayState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VpnGateway].
func (vg *VpnGateway) Type() string {
	return "azurerm_vpn_gateway"
}

// LocalName returns the local name for [VpnGateway].
func (vg *VpnGateway) LocalName() string {
	return vg.Name
}

// Configuration returns the configuration (args) for [VpnGateway].
func (vg *VpnGateway) Configuration() interface{} {
	return vg.Args
}

// DependOn is used for other resources to depend on [VpnGateway].
func (vg *VpnGateway) DependOn() terra.Reference {
	return terra.ReferenceResource(vg)
}

// Dependencies returns the list of resources [VpnGateway] depends_on.
func (vg *VpnGateway) Dependencies() terra.Dependencies {
	return vg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VpnGateway].
func (vg *VpnGateway) LifecycleManagement() *terra.Lifecycle {
	return vg.Lifecycle
}

// Attributes returns the attributes for [VpnGateway].
func (vg *VpnGateway) Attributes() vpnGatewayAttributes {
	return vpnGatewayAttributes{ref: terra.ReferenceResource(vg)}
}

// ImportState imports the given attribute values into [VpnGateway]'s state.
func (vg *VpnGateway) ImportState(av io.Reader) error {
	vg.state = &vpnGatewayState{}
	if err := json.NewDecoder(av).Decode(vg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vg.Type(), vg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VpnGateway] has state.
func (vg *VpnGateway) State() (*vpnGatewayState, bool) {
	return vg.state, vg.state != nil
}

// StateMust returns the state for [VpnGateway]. Panics if the state is nil.
func (vg *VpnGateway) StateMust() *vpnGatewayState {
	if vg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vg.Type(), vg.LocalName()))
	}
	return vg.state
}

// VpnGatewayArgs contains the configurations for azurerm_vpn_gateway.
type VpnGatewayArgs struct {
	// BgpRouteTranslationForNatEnabled: bool, optional
	BgpRouteTranslationForNatEnabled terra.BoolValue `hcl:"bgp_route_translation_for_nat_enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// RoutingPreference: string, optional
	RoutingPreference terra.StringValue `hcl:"routing_preference,attr"`
	// ScaleUnit: number, optional
	ScaleUnit terra.NumberValue `hcl:"scale_unit,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// VirtualHubId: string, required
	VirtualHubId terra.StringValue `hcl:"virtual_hub_id,attr" validate:"required"`
	// BgpSettings: optional
	BgpSettings *vpngateway.BgpSettings `hcl:"bgp_settings,block"`
	// Timeouts: optional
	Timeouts *vpngateway.Timeouts `hcl:"timeouts,block"`
}
type vpnGatewayAttributes struct {
	ref terra.Reference
}

// BgpRouteTranslationForNatEnabled returns a reference to field bgp_route_translation_for_nat_enabled of azurerm_vpn_gateway.
func (vg vpnGatewayAttributes) BgpRouteTranslationForNatEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(vg.ref.Append("bgp_route_translation_for_nat_enabled"))
}

// Id returns a reference to field id of azurerm_vpn_gateway.
func (vg vpnGatewayAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vg.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_vpn_gateway.
func (vg vpnGatewayAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(vg.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_vpn_gateway.
func (vg vpnGatewayAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vg.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_vpn_gateway.
func (vg vpnGatewayAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(vg.ref.Append("resource_group_name"))
}

// RoutingPreference returns a reference to field routing_preference of azurerm_vpn_gateway.
func (vg vpnGatewayAttributes) RoutingPreference() terra.StringValue {
	return terra.ReferenceAsString(vg.ref.Append("routing_preference"))
}

// ScaleUnit returns a reference to field scale_unit of azurerm_vpn_gateway.
func (vg vpnGatewayAttributes) ScaleUnit() terra.NumberValue {
	return terra.ReferenceAsNumber(vg.ref.Append("scale_unit"))
}

// Tags returns a reference to field tags of azurerm_vpn_gateway.
func (vg vpnGatewayAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vg.ref.Append("tags"))
}

// VirtualHubId returns a reference to field virtual_hub_id of azurerm_vpn_gateway.
func (vg vpnGatewayAttributes) VirtualHubId() terra.StringValue {
	return terra.ReferenceAsString(vg.ref.Append("virtual_hub_id"))
}

func (vg vpnGatewayAttributes) BgpSettings() terra.ListValue[vpngateway.BgpSettingsAttributes] {
	return terra.ReferenceAsList[vpngateway.BgpSettingsAttributes](vg.ref.Append("bgp_settings"))
}

func (vg vpnGatewayAttributes) Timeouts() vpngateway.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vpngateway.TimeoutsAttributes](vg.ref.Append("timeouts"))
}

type vpnGatewayState struct {
	BgpRouteTranslationForNatEnabled bool                          `json:"bgp_route_translation_for_nat_enabled"`
	Id                               string                        `json:"id"`
	Location                         string                        `json:"location"`
	Name                             string                        `json:"name"`
	ResourceGroupName                string                        `json:"resource_group_name"`
	RoutingPreference                string                        `json:"routing_preference"`
	ScaleUnit                        float64                       `json:"scale_unit"`
	Tags                             map[string]string             `json:"tags"`
	VirtualHubId                     string                        `json:"virtual_hub_id"`
	BgpSettings                      []vpngateway.BgpSettingsState `json:"bgp_settings"`
	Timeouts                         *vpngateway.TimeoutsState     `json:"timeouts"`
}
