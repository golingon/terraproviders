// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurestack_virtual_network_gateway_connection

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

// Resource represents the Terraform resource azurestack_virtual_network_gateway_connection.
type Resource struct {
	Name      string
	Args      Args
	state     *azurestackVirtualNetworkGatewayConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (avngc *Resource) Type() string {
	return "azurestack_virtual_network_gateway_connection"
}

// LocalName returns the local name for [Resource].
func (avngc *Resource) LocalName() string {
	return avngc.Name
}

// Configuration returns the configuration (args) for [Resource].
func (avngc *Resource) Configuration() interface{} {
	return avngc.Args
}

// DependOn is used for other resources to depend on [Resource].
func (avngc *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(avngc)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (avngc *Resource) Dependencies() terra.Dependencies {
	return avngc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (avngc *Resource) LifecycleManagement() *terra.Lifecycle {
	return avngc.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (avngc *Resource) Attributes() azurestackVirtualNetworkGatewayConnectionAttributes {
	return azurestackVirtualNetworkGatewayConnectionAttributes{ref: terra.ReferenceResource(avngc)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (avngc *Resource) ImportState(state io.Reader) error {
	avngc.state = &azurestackVirtualNetworkGatewayConnectionState{}
	if err := json.NewDecoder(state).Decode(avngc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", avngc.Type(), avngc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (avngc *Resource) State() (*azurestackVirtualNetworkGatewayConnectionState, bool) {
	return avngc.state, avngc.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (avngc *Resource) StateMust() *azurestackVirtualNetworkGatewayConnectionState {
	if avngc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", avngc.Type(), avngc.LocalName()))
	}
	return avngc.state
}

// Args contains the configurations for azurestack_virtual_network_gateway_connection.
type Args struct {
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
	IpsecPolicy *IpsecPolicy `hcl:"ipsec_policy,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurestackVirtualNetworkGatewayConnectionAttributes struct {
	ref terra.Reference
}

// AuthorizationKey returns a reference to field authorization_key of azurestack_virtual_network_gateway_connection.
func (avngc azurestackVirtualNetworkGatewayConnectionAttributes) AuthorizationKey() terra.StringValue {
	return terra.ReferenceAsString(avngc.ref.Append("authorization_key"))
}

// EnableBgp returns a reference to field enable_bgp of azurestack_virtual_network_gateway_connection.
func (avngc azurestackVirtualNetworkGatewayConnectionAttributes) EnableBgp() terra.BoolValue {
	return terra.ReferenceAsBool(avngc.ref.Append("enable_bgp"))
}

// ExpressRouteCircuitId returns a reference to field express_route_circuit_id of azurestack_virtual_network_gateway_connection.
func (avngc azurestackVirtualNetworkGatewayConnectionAttributes) ExpressRouteCircuitId() terra.StringValue {
	return terra.ReferenceAsString(avngc.ref.Append("express_route_circuit_id"))
}

// Id returns a reference to field id of azurestack_virtual_network_gateway_connection.
func (avngc azurestackVirtualNetworkGatewayConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(avngc.ref.Append("id"))
}

// LocalNetworkGatewayId returns a reference to field local_network_gateway_id of azurestack_virtual_network_gateway_connection.
func (avngc azurestackVirtualNetworkGatewayConnectionAttributes) LocalNetworkGatewayId() terra.StringValue {
	return terra.ReferenceAsString(avngc.ref.Append("local_network_gateway_id"))
}

// Location returns a reference to field location of azurestack_virtual_network_gateway_connection.
func (avngc azurestackVirtualNetworkGatewayConnectionAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(avngc.ref.Append("location"))
}

// Name returns a reference to field name of azurestack_virtual_network_gateway_connection.
func (avngc azurestackVirtualNetworkGatewayConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(avngc.ref.Append("name"))
}

// PeerVirtualNetworkGatewayId returns a reference to field peer_virtual_network_gateway_id of azurestack_virtual_network_gateway_connection.
func (avngc azurestackVirtualNetworkGatewayConnectionAttributes) PeerVirtualNetworkGatewayId() terra.StringValue {
	return terra.ReferenceAsString(avngc.ref.Append("peer_virtual_network_gateway_id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurestack_virtual_network_gateway_connection.
func (avngc azurestackVirtualNetworkGatewayConnectionAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(avngc.ref.Append("resource_group_name"))
}

// RoutingWeight returns a reference to field routing_weight of azurestack_virtual_network_gateway_connection.
func (avngc azurestackVirtualNetworkGatewayConnectionAttributes) RoutingWeight() terra.NumberValue {
	return terra.ReferenceAsNumber(avngc.ref.Append("routing_weight"))
}

// SharedKey returns a reference to field shared_key of azurestack_virtual_network_gateway_connection.
func (avngc azurestackVirtualNetworkGatewayConnectionAttributes) SharedKey() terra.StringValue {
	return terra.ReferenceAsString(avngc.ref.Append("shared_key"))
}

// Tags returns a reference to field tags of azurestack_virtual_network_gateway_connection.
func (avngc azurestackVirtualNetworkGatewayConnectionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](avngc.ref.Append("tags"))
}

// Type returns a reference to field type of azurestack_virtual_network_gateway_connection.
func (avngc azurestackVirtualNetworkGatewayConnectionAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(avngc.ref.Append("type"))
}

// UsePolicyBasedTrafficSelectors returns a reference to field use_policy_based_traffic_selectors of azurestack_virtual_network_gateway_connection.
func (avngc azurestackVirtualNetworkGatewayConnectionAttributes) UsePolicyBasedTrafficSelectors() terra.BoolValue {
	return terra.ReferenceAsBool(avngc.ref.Append("use_policy_based_traffic_selectors"))
}

// VirtualNetworkGatewayId returns a reference to field virtual_network_gateway_id of azurestack_virtual_network_gateway_connection.
func (avngc azurestackVirtualNetworkGatewayConnectionAttributes) VirtualNetworkGatewayId() terra.StringValue {
	return terra.ReferenceAsString(avngc.ref.Append("virtual_network_gateway_id"))
}

func (avngc azurestackVirtualNetworkGatewayConnectionAttributes) IpsecPolicy() terra.ListValue[IpsecPolicyAttributes] {
	return terra.ReferenceAsList[IpsecPolicyAttributes](avngc.ref.Append("ipsec_policy"))
}

func (avngc azurestackVirtualNetworkGatewayConnectionAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](avngc.ref.Append("timeouts"))
}

type azurestackVirtualNetworkGatewayConnectionState struct {
	AuthorizationKey               string             `json:"authorization_key"`
	EnableBgp                      bool               `json:"enable_bgp"`
	ExpressRouteCircuitId          string             `json:"express_route_circuit_id"`
	Id                             string             `json:"id"`
	LocalNetworkGatewayId          string             `json:"local_network_gateway_id"`
	Location                       string             `json:"location"`
	Name                           string             `json:"name"`
	PeerVirtualNetworkGatewayId    string             `json:"peer_virtual_network_gateway_id"`
	ResourceGroupName              string             `json:"resource_group_name"`
	RoutingWeight                  float64            `json:"routing_weight"`
	SharedKey                      string             `json:"shared_key"`
	Tags                           map[string]string  `json:"tags"`
	Type                           string             `json:"type"`
	UsePolicyBasedTrafficSelectors bool               `json:"use_policy_based_traffic_selectors"`
	VirtualNetworkGatewayId        string             `json:"virtual_network_gateway_id"`
	IpsecPolicy                    []IpsecPolicyState `json:"ipsec_policy"`
	Timeouts                       *TimeoutsState     `json:"timeouts"`
}
