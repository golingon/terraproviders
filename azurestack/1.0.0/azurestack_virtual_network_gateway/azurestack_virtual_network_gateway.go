// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurestack_virtual_network_gateway

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource azurestack_virtual_network_gateway.
type Resource struct {
	Name      string
	Args      Args
	state     *azurestackVirtualNetworkGatewayState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (avng *Resource) Type() string {
	return "azurestack_virtual_network_gateway"
}

// LocalName returns the local name for [Resource].
func (avng *Resource) LocalName() string {
	return avng.Name
}

// Configuration returns the configuration (args) for [Resource].
func (avng *Resource) Configuration() interface{} {
	return avng.Args
}

// DependOn is used for other resources to depend on [Resource].
func (avng *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(avng)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (avng *Resource) Dependencies() terra.Dependencies {
	return avng.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (avng *Resource) LifecycleManagement() *terra.Lifecycle {
	return avng.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (avng *Resource) Attributes() azurestackVirtualNetworkGatewayAttributes {
	return azurestackVirtualNetworkGatewayAttributes{ref: terra.ReferenceResource(avng)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (avng *Resource) ImportState(state io.Reader) error {
	avng.state = &azurestackVirtualNetworkGatewayState{}
	if err := json.NewDecoder(state).Decode(avng.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", avng.Type(), avng.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (avng *Resource) State() (*azurestackVirtualNetworkGatewayState, bool) {
	return avng.state, avng.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (avng *Resource) StateMust() *azurestackVirtualNetworkGatewayState {
	if avng.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", avng.Type(), avng.LocalName()))
	}
	return avng.state
}

// Args contains the configurations for azurestack_virtual_network_gateway.
type Args struct {
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
	BgpSettings *BgpSettings `hcl:"bgp_settings,block"`
	// IpConfiguration: min=1,max=3
	IpConfiguration []IpConfiguration `hcl:"ip_configuration,block" validate:"min=1,max=3"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
	// VpnClientConfiguration: optional
	VpnClientConfiguration *VpnClientConfiguration `hcl:"vpn_client_configuration,block"`
}

type azurestackVirtualNetworkGatewayAttributes struct {
	ref terra.Reference
}

// ActiveActive returns a reference to field active_active of azurestack_virtual_network_gateway.
func (avng azurestackVirtualNetworkGatewayAttributes) ActiveActive() terra.BoolValue {
	return terra.ReferenceAsBool(avng.ref.Append("active_active"))
}

// DefaultLocalNetworkGatewayId returns a reference to field default_local_network_gateway_id of azurestack_virtual_network_gateway.
func (avng azurestackVirtualNetworkGatewayAttributes) DefaultLocalNetworkGatewayId() terra.StringValue {
	return terra.ReferenceAsString(avng.ref.Append("default_local_network_gateway_id"))
}

// EnableBgp returns a reference to field enable_bgp of azurestack_virtual_network_gateway.
func (avng azurestackVirtualNetworkGatewayAttributes) EnableBgp() terra.BoolValue {
	return terra.ReferenceAsBool(avng.ref.Append("enable_bgp"))
}

// Id returns a reference to field id of azurestack_virtual_network_gateway.
func (avng azurestackVirtualNetworkGatewayAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(avng.ref.Append("id"))
}

// Location returns a reference to field location of azurestack_virtual_network_gateway.
func (avng azurestackVirtualNetworkGatewayAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(avng.ref.Append("location"))
}

// Name returns a reference to field name of azurestack_virtual_network_gateway.
func (avng azurestackVirtualNetworkGatewayAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(avng.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurestack_virtual_network_gateway.
func (avng azurestackVirtualNetworkGatewayAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(avng.ref.Append("resource_group_name"))
}

// Sku returns a reference to field sku of azurestack_virtual_network_gateway.
func (avng azurestackVirtualNetworkGatewayAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(avng.ref.Append("sku"))
}

// Tags returns a reference to field tags of azurestack_virtual_network_gateway.
func (avng azurestackVirtualNetworkGatewayAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](avng.ref.Append("tags"))
}

// Type returns a reference to field type of azurestack_virtual_network_gateway.
func (avng azurestackVirtualNetworkGatewayAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(avng.ref.Append("type"))
}

// VpnType returns a reference to field vpn_type of azurestack_virtual_network_gateway.
func (avng azurestackVirtualNetworkGatewayAttributes) VpnType() terra.StringValue {
	return terra.ReferenceAsString(avng.ref.Append("vpn_type"))
}

func (avng azurestackVirtualNetworkGatewayAttributes) BgpSettings() terra.ListValue[BgpSettingsAttributes] {
	return terra.ReferenceAsList[BgpSettingsAttributes](avng.ref.Append("bgp_settings"))
}

func (avng azurestackVirtualNetworkGatewayAttributes) IpConfiguration() terra.ListValue[IpConfigurationAttributes] {
	return terra.ReferenceAsList[IpConfigurationAttributes](avng.ref.Append("ip_configuration"))
}

func (avng azurestackVirtualNetworkGatewayAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](avng.ref.Append("timeouts"))
}

func (avng azurestackVirtualNetworkGatewayAttributes) VpnClientConfiguration() terra.ListValue[VpnClientConfigurationAttributes] {
	return terra.ReferenceAsList[VpnClientConfigurationAttributes](avng.ref.Append("vpn_client_configuration"))
}

type azurestackVirtualNetworkGatewayState struct {
	ActiveActive                 bool                          `json:"active_active"`
	DefaultLocalNetworkGatewayId string                        `json:"default_local_network_gateway_id"`
	EnableBgp                    bool                          `json:"enable_bgp"`
	Id                           string                        `json:"id"`
	Location                     string                        `json:"location"`
	Name                         string                        `json:"name"`
	ResourceGroupName            string                        `json:"resource_group_name"`
	Sku                          string                        `json:"sku"`
	Tags                         map[string]string             `json:"tags"`
	Type                         string                        `json:"type"`
	VpnType                      string                        `json:"vpn_type"`
	BgpSettings                  []BgpSettingsState            `json:"bgp_settings"`
	IpConfiguration              []IpConfigurationState        `json:"ip_configuration"`
	Timeouts                     *TimeoutsState                `json:"timeouts"`
	VpnClientConfiguration       []VpnClientConfigurationState `json:"vpn_client_configuration"`
}
