// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurestack

import (
	datavirtualnetworkgatewayconnection "github.com/golingon/terraproviders/azurestack/1.0.0/datavirtualnetworkgatewayconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataVirtualNetworkGatewayConnection creates a new instance of [DataVirtualNetworkGatewayConnection].
func NewDataVirtualNetworkGatewayConnection(name string, args DataVirtualNetworkGatewayConnectionArgs) *DataVirtualNetworkGatewayConnection {
	return &DataVirtualNetworkGatewayConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataVirtualNetworkGatewayConnection)(nil)

// DataVirtualNetworkGatewayConnection represents the Terraform data resource azurestack_virtual_network_gateway_connection.
type DataVirtualNetworkGatewayConnection struct {
	Name string
	Args DataVirtualNetworkGatewayConnectionArgs
}

// DataSource returns the Terraform object type for [DataVirtualNetworkGatewayConnection].
func (vngc *DataVirtualNetworkGatewayConnection) DataSource() string {
	return "azurestack_virtual_network_gateway_connection"
}

// LocalName returns the local name for [DataVirtualNetworkGatewayConnection].
func (vngc *DataVirtualNetworkGatewayConnection) LocalName() string {
	return vngc.Name
}

// Configuration returns the configuration (args) for [DataVirtualNetworkGatewayConnection].
func (vngc *DataVirtualNetworkGatewayConnection) Configuration() interface{} {
	return vngc.Args
}

// Attributes returns the attributes for [DataVirtualNetworkGatewayConnection].
func (vngc *DataVirtualNetworkGatewayConnection) Attributes() dataVirtualNetworkGatewayConnectionAttributes {
	return dataVirtualNetworkGatewayConnectionAttributes{ref: terra.ReferenceDataResource(vngc)}
}

// DataVirtualNetworkGatewayConnectionArgs contains the configurations for azurestack_virtual_network_gateway_connection.
type DataVirtualNetworkGatewayConnectionArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// IpsecPolicy: min=0
	IpsecPolicy []datavirtualnetworkgatewayconnection.IpsecPolicy `hcl:"ipsec_policy,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datavirtualnetworkgatewayconnection.Timeouts `hcl:"timeouts,block"`
}
type dataVirtualNetworkGatewayConnectionAttributes struct {
	ref terra.Reference
}

// AuthorizationKey returns a reference to field authorization_key of azurestack_virtual_network_gateway_connection.
func (vngc dataVirtualNetworkGatewayConnectionAttributes) AuthorizationKey() terra.StringValue {
	return terra.ReferenceAsString(vngc.ref.Append("authorization_key"))
}

// EgressBytesTransferred returns a reference to field egress_bytes_transferred of azurestack_virtual_network_gateway_connection.
func (vngc dataVirtualNetworkGatewayConnectionAttributes) EgressBytesTransferred() terra.NumberValue {
	return terra.ReferenceAsNumber(vngc.ref.Append("egress_bytes_transferred"))
}

// EnableBgp returns a reference to field enable_bgp of azurestack_virtual_network_gateway_connection.
func (vngc dataVirtualNetworkGatewayConnectionAttributes) EnableBgp() terra.BoolValue {
	return terra.ReferenceAsBool(vngc.ref.Append("enable_bgp"))
}

// ExpressRouteCircuitId returns a reference to field express_route_circuit_id of azurestack_virtual_network_gateway_connection.
func (vngc dataVirtualNetworkGatewayConnectionAttributes) ExpressRouteCircuitId() terra.StringValue {
	return terra.ReferenceAsString(vngc.ref.Append("express_route_circuit_id"))
}

// Id returns a reference to field id of azurestack_virtual_network_gateway_connection.
func (vngc dataVirtualNetworkGatewayConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vngc.ref.Append("id"))
}

// IngressBytesTransferred returns a reference to field ingress_bytes_transferred of azurestack_virtual_network_gateway_connection.
func (vngc dataVirtualNetworkGatewayConnectionAttributes) IngressBytesTransferred() terra.NumberValue {
	return terra.ReferenceAsNumber(vngc.ref.Append("ingress_bytes_transferred"))
}

// LocalNetworkGatewayId returns a reference to field local_network_gateway_id of azurestack_virtual_network_gateway_connection.
func (vngc dataVirtualNetworkGatewayConnectionAttributes) LocalNetworkGatewayId() terra.StringValue {
	return terra.ReferenceAsString(vngc.ref.Append("local_network_gateway_id"))
}

// Location returns a reference to field location of azurestack_virtual_network_gateway_connection.
func (vngc dataVirtualNetworkGatewayConnectionAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(vngc.ref.Append("location"))
}

// Name returns a reference to field name of azurestack_virtual_network_gateway_connection.
func (vngc dataVirtualNetworkGatewayConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vngc.ref.Append("name"))
}

// PeerVirtualNetworkGatewayId returns a reference to field peer_virtual_network_gateway_id of azurestack_virtual_network_gateway_connection.
func (vngc dataVirtualNetworkGatewayConnectionAttributes) PeerVirtualNetworkGatewayId() terra.StringValue {
	return terra.ReferenceAsString(vngc.ref.Append("peer_virtual_network_gateway_id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurestack_virtual_network_gateway_connection.
func (vngc dataVirtualNetworkGatewayConnectionAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(vngc.ref.Append("resource_group_name"))
}

// ResourceGuid returns a reference to field resource_guid of azurestack_virtual_network_gateway_connection.
func (vngc dataVirtualNetworkGatewayConnectionAttributes) ResourceGuid() terra.StringValue {
	return terra.ReferenceAsString(vngc.ref.Append("resource_guid"))
}

// RoutingWeight returns a reference to field routing_weight of azurestack_virtual_network_gateway_connection.
func (vngc dataVirtualNetworkGatewayConnectionAttributes) RoutingWeight() terra.NumberValue {
	return terra.ReferenceAsNumber(vngc.ref.Append("routing_weight"))
}

// SharedKey returns a reference to field shared_key of azurestack_virtual_network_gateway_connection.
func (vngc dataVirtualNetworkGatewayConnectionAttributes) SharedKey() terra.StringValue {
	return terra.ReferenceAsString(vngc.ref.Append("shared_key"))
}

// Tags returns a reference to field tags of azurestack_virtual_network_gateway_connection.
func (vngc dataVirtualNetworkGatewayConnectionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vngc.ref.Append("tags"))
}

// Type returns a reference to field type of azurestack_virtual_network_gateway_connection.
func (vngc dataVirtualNetworkGatewayConnectionAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(vngc.ref.Append("type"))
}

// UsePolicyBasedTrafficSelectors returns a reference to field use_policy_based_traffic_selectors of azurestack_virtual_network_gateway_connection.
func (vngc dataVirtualNetworkGatewayConnectionAttributes) UsePolicyBasedTrafficSelectors() terra.BoolValue {
	return terra.ReferenceAsBool(vngc.ref.Append("use_policy_based_traffic_selectors"))
}

// VirtualNetworkGatewayId returns a reference to field virtual_network_gateway_id of azurestack_virtual_network_gateway_connection.
func (vngc dataVirtualNetworkGatewayConnectionAttributes) VirtualNetworkGatewayId() terra.StringValue {
	return terra.ReferenceAsString(vngc.ref.Append("virtual_network_gateway_id"))
}

func (vngc dataVirtualNetworkGatewayConnectionAttributes) IpsecPolicy() terra.ListValue[datavirtualnetworkgatewayconnection.IpsecPolicyAttributes] {
	return terra.ReferenceAsList[datavirtualnetworkgatewayconnection.IpsecPolicyAttributes](vngc.ref.Append("ipsec_policy"))
}

func (vngc dataVirtualNetworkGatewayConnectionAttributes) Timeouts() datavirtualnetworkgatewayconnection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datavirtualnetworkgatewayconnection.TimeoutsAttributes](vngc.ref.Append("timeouts"))
}
