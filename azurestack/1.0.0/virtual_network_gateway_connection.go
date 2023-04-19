// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurestack

import (
	"encoding/json"
	"fmt"
	virtualnetworkgatewayconnection "github.com/golingon/terraproviders/azurestack/1.0.0/virtualnetworkgatewayconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVirtualNetworkGatewayConnection creates a new instance of [VirtualNetworkGatewayConnection].
func NewVirtualNetworkGatewayConnection(name string, args VirtualNetworkGatewayConnectionArgs) *VirtualNetworkGatewayConnection {
	return &VirtualNetworkGatewayConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VirtualNetworkGatewayConnection)(nil)

// VirtualNetworkGatewayConnection represents the Terraform resource azurestack_virtual_network_gateway_connection.
type VirtualNetworkGatewayConnection struct {
	Name      string
	Args      VirtualNetworkGatewayConnectionArgs
	state     *virtualNetworkGatewayConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VirtualNetworkGatewayConnection].
func (vngc *VirtualNetworkGatewayConnection) Type() string {
	return "azurestack_virtual_network_gateway_connection"
}

// LocalName returns the local name for [VirtualNetworkGatewayConnection].
func (vngc *VirtualNetworkGatewayConnection) LocalName() string {
	return vngc.Name
}

// Configuration returns the configuration (args) for [VirtualNetworkGatewayConnection].
func (vngc *VirtualNetworkGatewayConnection) Configuration() interface{} {
	return vngc.Args
}

// DependOn is used for other resources to depend on [VirtualNetworkGatewayConnection].
func (vngc *VirtualNetworkGatewayConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(vngc)
}

// Dependencies returns the list of resources [VirtualNetworkGatewayConnection] depends_on.
func (vngc *VirtualNetworkGatewayConnection) Dependencies() terra.Dependencies {
	return vngc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VirtualNetworkGatewayConnection].
func (vngc *VirtualNetworkGatewayConnection) LifecycleManagement() *terra.Lifecycle {
	return vngc.Lifecycle
}

// Attributes returns the attributes for [VirtualNetworkGatewayConnection].
func (vngc *VirtualNetworkGatewayConnection) Attributes() virtualNetworkGatewayConnectionAttributes {
	return virtualNetworkGatewayConnectionAttributes{ref: terra.ReferenceResource(vngc)}
}

// ImportState imports the given attribute values into [VirtualNetworkGatewayConnection]'s state.
func (vngc *VirtualNetworkGatewayConnection) ImportState(av io.Reader) error {
	vngc.state = &virtualNetworkGatewayConnectionState{}
	if err := json.NewDecoder(av).Decode(vngc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vngc.Type(), vngc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VirtualNetworkGatewayConnection] has state.
func (vngc *VirtualNetworkGatewayConnection) State() (*virtualNetworkGatewayConnectionState, bool) {
	return vngc.state, vngc.state != nil
}

// StateMust returns the state for [VirtualNetworkGatewayConnection]. Panics if the state is nil.
func (vngc *VirtualNetworkGatewayConnection) StateMust() *virtualNetworkGatewayConnectionState {
	if vngc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vngc.Type(), vngc.LocalName()))
	}
	return vngc.state
}

// VirtualNetworkGatewayConnectionArgs contains the configurations for azurestack_virtual_network_gateway_connection.
type VirtualNetworkGatewayConnectionArgs struct {
	// AuthorizationKey: string, optional
	AuthorizationKey terra.StringValue `hcl:"authorization_key,attr"`
	// EnableBgp: bool, optional
	EnableBgp terra.BoolValue `hcl:"enable_bgp,attr"`
	// ExpressRouteCircuitId: string, optional
	ExpressRouteCircuitId terra.StringValue `hcl:"express_route_circuit_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LocalNetworkGatewayId: string, optional
	LocalNetworkGatewayId terra.StringValue `hcl:"local_network_gateway_id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PeerVirtualNetworkGatewayId: string, optional
	PeerVirtualNetworkGatewayId terra.StringValue `hcl:"peer_virtual_network_gateway_id,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// RoutingWeight: number, optional
	RoutingWeight terra.NumberValue `hcl:"routing_weight,attr"`
	// SharedKey: string, optional
	SharedKey terra.StringValue `hcl:"shared_key,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// UsePolicyBasedTrafficSelectors: bool, optional
	UsePolicyBasedTrafficSelectors terra.BoolValue `hcl:"use_policy_based_traffic_selectors,attr"`
	// VirtualNetworkGatewayId: string, required
	VirtualNetworkGatewayId terra.StringValue `hcl:"virtual_network_gateway_id,attr" validate:"required"`
	// IpsecPolicy: optional
	IpsecPolicy *virtualnetworkgatewayconnection.IpsecPolicy `hcl:"ipsec_policy,block"`
	// Timeouts: optional
	Timeouts *virtualnetworkgatewayconnection.Timeouts `hcl:"timeouts,block"`
}
type virtualNetworkGatewayConnectionAttributes struct {
	ref terra.Reference
}

// AuthorizationKey returns a reference to field authorization_key of azurestack_virtual_network_gateway_connection.
func (vngc virtualNetworkGatewayConnectionAttributes) AuthorizationKey() terra.StringValue {
	return terra.ReferenceAsString(vngc.ref.Append("authorization_key"))
}

// EnableBgp returns a reference to field enable_bgp of azurestack_virtual_network_gateway_connection.
func (vngc virtualNetworkGatewayConnectionAttributes) EnableBgp() terra.BoolValue {
	return terra.ReferenceAsBool(vngc.ref.Append("enable_bgp"))
}

// ExpressRouteCircuitId returns a reference to field express_route_circuit_id of azurestack_virtual_network_gateway_connection.
func (vngc virtualNetworkGatewayConnectionAttributes) ExpressRouteCircuitId() terra.StringValue {
	return terra.ReferenceAsString(vngc.ref.Append("express_route_circuit_id"))
}

// Id returns a reference to field id of azurestack_virtual_network_gateway_connection.
func (vngc virtualNetworkGatewayConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vngc.ref.Append("id"))
}

// LocalNetworkGatewayId returns a reference to field local_network_gateway_id of azurestack_virtual_network_gateway_connection.
func (vngc virtualNetworkGatewayConnectionAttributes) LocalNetworkGatewayId() terra.StringValue {
	return terra.ReferenceAsString(vngc.ref.Append("local_network_gateway_id"))
}

// Location returns a reference to field location of azurestack_virtual_network_gateway_connection.
func (vngc virtualNetworkGatewayConnectionAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(vngc.ref.Append("location"))
}

// Name returns a reference to field name of azurestack_virtual_network_gateway_connection.
func (vngc virtualNetworkGatewayConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vngc.ref.Append("name"))
}

// PeerVirtualNetworkGatewayId returns a reference to field peer_virtual_network_gateway_id of azurestack_virtual_network_gateway_connection.
func (vngc virtualNetworkGatewayConnectionAttributes) PeerVirtualNetworkGatewayId() terra.StringValue {
	return terra.ReferenceAsString(vngc.ref.Append("peer_virtual_network_gateway_id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurestack_virtual_network_gateway_connection.
func (vngc virtualNetworkGatewayConnectionAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(vngc.ref.Append("resource_group_name"))
}

// RoutingWeight returns a reference to field routing_weight of azurestack_virtual_network_gateway_connection.
func (vngc virtualNetworkGatewayConnectionAttributes) RoutingWeight() terra.NumberValue {
	return terra.ReferenceAsNumber(vngc.ref.Append("routing_weight"))
}

// SharedKey returns a reference to field shared_key of azurestack_virtual_network_gateway_connection.
func (vngc virtualNetworkGatewayConnectionAttributes) SharedKey() terra.StringValue {
	return terra.ReferenceAsString(vngc.ref.Append("shared_key"))
}

// Tags returns a reference to field tags of azurestack_virtual_network_gateway_connection.
func (vngc virtualNetworkGatewayConnectionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vngc.ref.Append("tags"))
}

// Type returns a reference to field type of azurestack_virtual_network_gateway_connection.
func (vngc virtualNetworkGatewayConnectionAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(vngc.ref.Append("type"))
}

// UsePolicyBasedTrafficSelectors returns a reference to field use_policy_based_traffic_selectors of azurestack_virtual_network_gateway_connection.
func (vngc virtualNetworkGatewayConnectionAttributes) UsePolicyBasedTrafficSelectors() terra.BoolValue {
	return terra.ReferenceAsBool(vngc.ref.Append("use_policy_based_traffic_selectors"))
}

// VirtualNetworkGatewayId returns a reference to field virtual_network_gateway_id of azurestack_virtual_network_gateway_connection.
func (vngc virtualNetworkGatewayConnectionAttributes) VirtualNetworkGatewayId() terra.StringValue {
	return terra.ReferenceAsString(vngc.ref.Append("virtual_network_gateway_id"))
}

func (vngc virtualNetworkGatewayConnectionAttributes) IpsecPolicy() terra.ListValue[virtualnetworkgatewayconnection.IpsecPolicyAttributes] {
	return terra.ReferenceAsList[virtualnetworkgatewayconnection.IpsecPolicyAttributes](vngc.ref.Append("ipsec_policy"))
}

func (vngc virtualNetworkGatewayConnectionAttributes) Timeouts() virtualnetworkgatewayconnection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[virtualnetworkgatewayconnection.TimeoutsAttributes](vngc.ref.Append("timeouts"))
}

type virtualNetworkGatewayConnectionState struct {
	AuthorizationKey               string                                             `json:"authorization_key"`
	EnableBgp                      bool                                               `json:"enable_bgp"`
	ExpressRouteCircuitId          string                                             `json:"express_route_circuit_id"`
	Id                             string                                             `json:"id"`
	LocalNetworkGatewayId          string                                             `json:"local_network_gateway_id"`
	Location                       string                                             `json:"location"`
	Name                           string                                             `json:"name"`
	PeerVirtualNetworkGatewayId    string                                             `json:"peer_virtual_network_gateway_id"`
	ResourceGroupName              string                                             `json:"resource_group_name"`
	RoutingWeight                  float64                                            `json:"routing_weight"`
	SharedKey                      string                                             `json:"shared_key"`
	Tags                           map[string]string                                  `json:"tags"`
	Type                           string                                             `json:"type"`
	UsePolicyBasedTrafficSelectors bool                                               `json:"use_policy_based_traffic_selectors"`
	VirtualNetworkGatewayId        string                                             `json:"virtual_network_gateway_id"`
	IpsecPolicy                    []virtualnetworkgatewayconnection.IpsecPolicyState `json:"ipsec_policy"`
	Timeouts                       *virtualnetworkgatewayconnection.TimeoutsState     `json:"timeouts"`
}