// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	virtualnetworkpeering "github.com/golingon/terraproviders/azurerm/3.64.0/virtualnetworkpeering"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVirtualNetworkPeering creates a new instance of [VirtualNetworkPeering].
func NewVirtualNetworkPeering(name string, args VirtualNetworkPeeringArgs) *VirtualNetworkPeering {
	return &VirtualNetworkPeering{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VirtualNetworkPeering)(nil)

// VirtualNetworkPeering represents the Terraform resource azurerm_virtual_network_peering.
type VirtualNetworkPeering struct {
	Name      string
	Args      VirtualNetworkPeeringArgs
	state     *virtualNetworkPeeringState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VirtualNetworkPeering].
func (vnp *VirtualNetworkPeering) Type() string {
	return "azurerm_virtual_network_peering"
}

// LocalName returns the local name for [VirtualNetworkPeering].
func (vnp *VirtualNetworkPeering) LocalName() string {
	return vnp.Name
}

// Configuration returns the configuration (args) for [VirtualNetworkPeering].
func (vnp *VirtualNetworkPeering) Configuration() interface{} {
	return vnp.Args
}

// DependOn is used for other resources to depend on [VirtualNetworkPeering].
func (vnp *VirtualNetworkPeering) DependOn() terra.Reference {
	return terra.ReferenceResource(vnp)
}

// Dependencies returns the list of resources [VirtualNetworkPeering] depends_on.
func (vnp *VirtualNetworkPeering) Dependencies() terra.Dependencies {
	return vnp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VirtualNetworkPeering].
func (vnp *VirtualNetworkPeering) LifecycleManagement() *terra.Lifecycle {
	return vnp.Lifecycle
}

// Attributes returns the attributes for [VirtualNetworkPeering].
func (vnp *VirtualNetworkPeering) Attributes() virtualNetworkPeeringAttributes {
	return virtualNetworkPeeringAttributes{ref: terra.ReferenceResource(vnp)}
}

// ImportState imports the given attribute values into [VirtualNetworkPeering]'s state.
func (vnp *VirtualNetworkPeering) ImportState(av io.Reader) error {
	vnp.state = &virtualNetworkPeeringState{}
	if err := json.NewDecoder(av).Decode(vnp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vnp.Type(), vnp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VirtualNetworkPeering] has state.
func (vnp *VirtualNetworkPeering) State() (*virtualNetworkPeeringState, bool) {
	return vnp.state, vnp.state != nil
}

// StateMust returns the state for [VirtualNetworkPeering]. Panics if the state is nil.
func (vnp *VirtualNetworkPeering) StateMust() *virtualNetworkPeeringState {
	if vnp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vnp.Type(), vnp.LocalName()))
	}
	return vnp.state
}

// VirtualNetworkPeeringArgs contains the configurations for azurerm_virtual_network_peering.
type VirtualNetworkPeeringArgs struct {
	// AllowForwardedTraffic: bool, optional
	AllowForwardedTraffic terra.BoolValue `hcl:"allow_forwarded_traffic,attr"`
	// AllowGatewayTransit: bool, optional
	AllowGatewayTransit terra.BoolValue `hcl:"allow_gateway_transit,attr"`
	// AllowVirtualNetworkAccess: bool, optional
	AllowVirtualNetworkAccess terra.BoolValue `hcl:"allow_virtual_network_access,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RemoteVirtualNetworkId: string, required
	RemoteVirtualNetworkId terra.StringValue `hcl:"remote_virtual_network_id,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Triggers: map of string, optional
	Triggers terra.MapValue[terra.StringValue] `hcl:"triggers,attr"`
	// UseRemoteGateways: bool, optional
	UseRemoteGateways terra.BoolValue `hcl:"use_remote_gateways,attr"`
	// VirtualNetworkName: string, required
	VirtualNetworkName terra.StringValue `hcl:"virtual_network_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *virtualnetworkpeering.Timeouts `hcl:"timeouts,block"`
}
type virtualNetworkPeeringAttributes struct {
	ref terra.Reference
}

// AllowForwardedTraffic returns a reference to field allow_forwarded_traffic of azurerm_virtual_network_peering.
func (vnp virtualNetworkPeeringAttributes) AllowForwardedTraffic() terra.BoolValue {
	return terra.ReferenceAsBool(vnp.ref.Append("allow_forwarded_traffic"))
}

// AllowGatewayTransit returns a reference to field allow_gateway_transit of azurerm_virtual_network_peering.
func (vnp virtualNetworkPeeringAttributes) AllowGatewayTransit() terra.BoolValue {
	return terra.ReferenceAsBool(vnp.ref.Append("allow_gateway_transit"))
}

// AllowVirtualNetworkAccess returns a reference to field allow_virtual_network_access of azurerm_virtual_network_peering.
func (vnp virtualNetworkPeeringAttributes) AllowVirtualNetworkAccess() terra.BoolValue {
	return terra.ReferenceAsBool(vnp.ref.Append("allow_virtual_network_access"))
}

// Id returns a reference to field id of azurerm_virtual_network_peering.
func (vnp virtualNetworkPeeringAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vnp.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_virtual_network_peering.
func (vnp virtualNetworkPeeringAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vnp.ref.Append("name"))
}

// RemoteVirtualNetworkId returns a reference to field remote_virtual_network_id of azurerm_virtual_network_peering.
func (vnp virtualNetworkPeeringAttributes) RemoteVirtualNetworkId() terra.StringValue {
	return terra.ReferenceAsString(vnp.ref.Append("remote_virtual_network_id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_virtual_network_peering.
func (vnp virtualNetworkPeeringAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(vnp.ref.Append("resource_group_name"))
}

// Triggers returns a reference to field triggers of azurerm_virtual_network_peering.
func (vnp virtualNetworkPeeringAttributes) Triggers() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vnp.ref.Append("triggers"))
}

// UseRemoteGateways returns a reference to field use_remote_gateways of azurerm_virtual_network_peering.
func (vnp virtualNetworkPeeringAttributes) UseRemoteGateways() terra.BoolValue {
	return terra.ReferenceAsBool(vnp.ref.Append("use_remote_gateways"))
}

// VirtualNetworkName returns a reference to field virtual_network_name of azurerm_virtual_network_peering.
func (vnp virtualNetworkPeeringAttributes) VirtualNetworkName() terra.StringValue {
	return terra.ReferenceAsString(vnp.ref.Append("virtual_network_name"))
}

func (vnp virtualNetworkPeeringAttributes) Timeouts() virtualnetworkpeering.TimeoutsAttributes {
	return terra.ReferenceAsSingle[virtualnetworkpeering.TimeoutsAttributes](vnp.ref.Append("timeouts"))
}

type virtualNetworkPeeringState struct {
	AllowForwardedTraffic     bool                                 `json:"allow_forwarded_traffic"`
	AllowGatewayTransit       bool                                 `json:"allow_gateway_transit"`
	AllowVirtualNetworkAccess bool                                 `json:"allow_virtual_network_access"`
	Id                        string                               `json:"id"`
	Name                      string                               `json:"name"`
	RemoteVirtualNetworkId    string                               `json:"remote_virtual_network_id"`
	ResourceGroupName         string                               `json:"resource_group_name"`
	Triggers                  map[string]string                    `json:"triggers"`
	UseRemoteGateways         bool                                 `json:"use_remote_gateways"`
	VirtualNetworkName        string                               `json:"virtual_network_name"`
	Timeouts                  *virtualnetworkpeering.TimeoutsState `json:"timeouts"`
}
