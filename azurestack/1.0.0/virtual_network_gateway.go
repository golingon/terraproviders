// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurestack

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	virtualnetworkgateway "github.com/golingon/terraproviders/azurestack/1.0.0/virtualnetworkgateway"
	"io"
)

// NewVirtualNetworkGateway creates a new instance of [VirtualNetworkGateway].
func NewVirtualNetworkGateway(name string, args VirtualNetworkGatewayArgs) *VirtualNetworkGateway {
	return &VirtualNetworkGateway{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VirtualNetworkGateway)(nil)

// VirtualNetworkGateway represents the Terraform resource azurestack_virtual_network_gateway.
type VirtualNetworkGateway struct {
	Name      string
	Args      VirtualNetworkGatewayArgs
	state     *virtualNetworkGatewayState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VirtualNetworkGateway].
func (vng *VirtualNetworkGateway) Type() string {
	return "azurestack_virtual_network_gateway"
}

// LocalName returns the local name for [VirtualNetworkGateway].
func (vng *VirtualNetworkGateway) LocalName() string {
	return vng.Name
}

// Configuration returns the configuration (args) for [VirtualNetworkGateway].
func (vng *VirtualNetworkGateway) Configuration() interface{} {
	return vng.Args
}

// DependOn is used for other resources to depend on [VirtualNetworkGateway].
func (vng *VirtualNetworkGateway) DependOn() terra.Reference {
	return terra.ReferenceResource(vng)
}

// Dependencies returns the list of resources [VirtualNetworkGateway] depends_on.
func (vng *VirtualNetworkGateway) Dependencies() terra.Dependencies {
	return vng.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VirtualNetworkGateway].
func (vng *VirtualNetworkGateway) LifecycleManagement() *terra.Lifecycle {
	return vng.Lifecycle
}

// Attributes returns the attributes for [VirtualNetworkGateway].
func (vng *VirtualNetworkGateway) Attributes() virtualNetworkGatewayAttributes {
	return virtualNetworkGatewayAttributes{ref: terra.ReferenceResource(vng)}
}

// ImportState imports the given attribute values into [VirtualNetworkGateway]'s state.
func (vng *VirtualNetworkGateway) ImportState(av io.Reader) error {
	vng.state = &virtualNetworkGatewayState{}
	if err := json.NewDecoder(av).Decode(vng.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vng.Type(), vng.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VirtualNetworkGateway] has state.
func (vng *VirtualNetworkGateway) State() (*virtualNetworkGatewayState, bool) {
	return vng.state, vng.state != nil
}

// StateMust returns the state for [VirtualNetworkGateway]. Panics if the state is nil.
func (vng *VirtualNetworkGateway) StateMust() *virtualNetworkGatewayState {
	if vng.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vng.Type(), vng.LocalName()))
	}
	return vng.state
}

// VirtualNetworkGatewayArgs contains the configurations for azurestack_virtual_network_gateway.
type VirtualNetworkGatewayArgs struct {
	// ActiveActive: bool, optional
	ActiveActive terra.BoolValue `hcl:"active_active,attr"`
	// DefaultLocalNetworkGatewayId: string, optional
	DefaultLocalNetworkGatewayId terra.StringValue `hcl:"default_local_network_gateway_id,attr"`
	// EnableBgp: bool, optional
	EnableBgp terra.BoolValue `hcl:"enable_bgp,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Sku: string, required
	Sku terra.StringValue `hcl:"sku,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// VpnType: string, optional
	VpnType terra.StringValue `hcl:"vpn_type,attr"`
	// BgpSettings: optional
	BgpSettings *virtualnetworkgateway.BgpSettings `hcl:"bgp_settings,block"`
	// IpConfiguration: min=1,max=3
	IpConfiguration []virtualnetworkgateway.IpConfiguration `hcl:"ip_configuration,block" validate:"min=1,max=3"`
	// Timeouts: optional
	Timeouts *virtualnetworkgateway.Timeouts `hcl:"timeouts,block"`
	// VpnClientConfiguration: optional
	VpnClientConfiguration *virtualnetworkgateway.VpnClientConfiguration `hcl:"vpn_client_configuration,block"`
}
type virtualNetworkGatewayAttributes struct {
	ref terra.Reference
}

// ActiveActive returns a reference to field active_active of azurestack_virtual_network_gateway.
func (vng virtualNetworkGatewayAttributes) ActiveActive() terra.BoolValue {
	return terra.ReferenceAsBool(vng.ref.Append("active_active"))
}

// DefaultLocalNetworkGatewayId returns a reference to field default_local_network_gateway_id of azurestack_virtual_network_gateway.
func (vng virtualNetworkGatewayAttributes) DefaultLocalNetworkGatewayId() terra.StringValue {
	return terra.ReferenceAsString(vng.ref.Append("default_local_network_gateway_id"))
}

// EnableBgp returns a reference to field enable_bgp of azurestack_virtual_network_gateway.
func (vng virtualNetworkGatewayAttributes) EnableBgp() terra.BoolValue {
	return terra.ReferenceAsBool(vng.ref.Append("enable_bgp"))
}

// Id returns a reference to field id of azurestack_virtual_network_gateway.
func (vng virtualNetworkGatewayAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vng.ref.Append("id"))
}

// Location returns a reference to field location of azurestack_virtual_network_gateway.
func (vng virtualNetworkGatewayAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(vng.ref.Append("location"))
}

// Name returns a reference to field name of azurestack_virtual_network_gateway.
func (vng virtualNetworkGatewayAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vng.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurestack_virtual_network_gateway.
func (vng virtualNetworkGatewayAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(vng.ref.Append("resource_group_name"))
}

// Sku returns a reference to field sku of azurestack_virtual_network_gateway.
func (vng virtualNetworkGatewayAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(vng.ref.Append("sku"))
}

// Tags returns a reference to field tags of azurestack_virtual_network_gateway.
func (vng virtualNetworkGatewayAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vng.ref.Append("tags"))
}

// Type returns a reference to field type of azurestack_virtual_network_gateway.
func (vng virtualNetworkGatewayAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(vng.ref.Append("type"))
}

// VpnType returns a reference to field vpn_type of azurestack_virtual_network_gateway.
func (vng virtualNetworkGatewayAttributes) VpnType() terra.StringValue {
	return terra.ReferenceAsString(vng.ref.Append("vpn_type"))
}

func (vng virtualNetworkGatewayAttributes) BgpSettings() terra.ListValue[virtualnetworkgateway.BgpSettingsAttributes] {
	return terra.ReferenceAsList[virtualnetworkgateway.BgpSettingsAttributes](vng.ref.Append("bgp_settings"))
}

func (vng virtualNetworkGatewayAttributes) IpConfiguration() terra.ListValue[virtualnetworkgateway.IpConfigurationAttributes] {
	return terra.ReferenceAsList[virtualnetworkgateway.IpConfigurationAttributes](vng.ref.Append("ip_configuration"))
}

func (vng virtualNetworkGatewayAttributes) Timeouts() virtualnetworkgateway.TimeoutsAttributes {
	return terra.ReferenceAsSingle[virtualnetworkgateway.TimeoutsAttributes](vng.ref.Append("timeouts"))
}

func (vng virtualNetworkGatewayAttributes) VpnClientConfiguration() terra.ListValue[virtualnetworkgateway.VpnClientConfigurationAttributes] {
	return terra.ReferenceAsList[virtualnetworkgateway.VpnClientConfigurationAttributes](vng.ref.Append("vpn_client_configuration"))
}

type virtualNetworkGatewayState struct {
	ActiveActive                 bool                                                `json:"active_active"`
	DefaultLocalNetworkGatewayId string                                              `json:"default_local_network_gateway_id"`
	EnableBgp                    bool                                                `json:"enable_bgp"`
	Id                           string                                              `json:"id"`
	Location                     string                                              `json:"location"`
	Name                         string                                              `json:"name"`
	ResourceGroupName            string                                              `json:"resource_group_name"`
	Sku                          string                                              `json:"sku"`
	Tags                         map[string]string                                   `json:"tags"`
	Type                         string                                              `json:"type"`
	VpnType                      string                                              `json:"vpn_type"`
	BgpSettings                  []virtualnetworkgateway.BgpSettingsState            `json:"bgp_settings"`
	IpConfiguration              []virtualnetworkgateway.IpConfigurationState        `json:"ip_configuration"`
	Timeouts                     *virtualnetworkgateway.TimeoutsState                `json:"timeouts"`
	VpnClientConfiguration       []virtualnetworkgateway.VpnClientConfigurationState `json:"vpn_client_configuration"`
}
