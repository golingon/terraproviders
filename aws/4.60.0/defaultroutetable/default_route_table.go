// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package defaultroutetable

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Route struct {
	// CidrBlock: string, optional
	CidrBlock terra.StringValue `hcl:"cidr_block,attr"`
	// CoreNetworkArn: string, optional
	CoreNetworkArn terra.StringValue `hcl:"core_network_arn,attr"`
	// DestinationPrefixListId: string, optional
	DestinationPrefixListId terra.StringValue `hcl:"destination_prefix_list_id,attr"`
	// EgressOnlyGatewayId: string, optional
	EgressOnlyGatewayId terra.StringValue `hcl:"egress_only_gateway_id,attr"`
	// GatewayId: string, optional
	GatewayId terra.StringValue `hcl:"gateway_id,attr"`
	// InstanceId: string, optional
	InstanceId terra.StringValue `hcl:"instance_id,attr"`
	// Ipv6CidrBlock: string, optional
	Ipv6CidrBlock terra.StringValue `hcl:"ipv6_cidr_block,attr"`
	// NatGatewayId: string, optional
	NatGatewayId terra.StringValue `hcl:"nat_gateway_id,attr"`
	// NetworkInterfaceId: string, optional
	NetworkInterfaceId terra.StringValue `hcl:"network_interface_id,attr"`
	// TransitGatewayId: string, optional
	TransitGatewayId terra.StringValue `hcl:"transit_gateway_id,attr"`
	// VpcEndpointId: string, optional
	VpcEndpointId terra.StringValue `hcl:"vpc_endpoint_id,attr"`
	// VpcPeeringConnectionId: string, optional
	VpcPeeringConnectionId terra.StringValue `hcl:"vpc_peering_connection_id,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type RouteAttributes struct {
	ref terra.Reference
}

func (r RouteAttributes) InternalRef() terra.Reference {
	return r.ref
}

func (r RouteAttributes) InternalWithRef(ref terra.Reference) RouteAttributes {
	return RouteAttributes{ref: ref}
}

func (r RouteAttributes) InternalTokens() hclwrite.Tokens {
	return r.ref.InternalTokens()
}

func (r RouteAttributes) CidrBlock() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("cidr_block"))
}

func (r RouteAttributes) CoreNetworkArn() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("core_network_arn"))
}

func (r RouteAttributes) DestinationPrefixListId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("destination_prefix_list_id"))
}

func (r RouteAttributes) EgressOnlyGatewayId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("egress_only_gateway_id"))
}

func (r RouteAttributes) GatewayId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("gateway_id"))
}

func (r RouteAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("instance_id"))
}

func (r RouteAttributes) Ipv6CidrBlock() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("ipv6_cidr_block"))
}

func (r RouteAttributes) NatGatewayId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("nat_gateway_id"))
}

func (r RouteAttributes) NetworkInterfaceId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("network_interface_id"))
}

func (r RouteAttributes) TransitGatewayId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("transit_gateway_id"))
}

func (r RouteAttributes) VpcEndpointId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("vpc_endpoint_id"))
}

func (r RouteAttributes) VpcPeeringConnectionId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("vpc_peering_connection_id"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type RouteState struct {
	CidrBlock               string `json:"cidr_block"`
	CoreNetworkArn          string `json:"core_network_arn"`
	DestinationPrefixListId string `json:"destination_prefix_list_id"`
	EgressOnlyGatewayId     string `json:"egress_only_gateway_id"`
	GatewayId               string `json:"gateway_id"`
	InstanceId              string `json:"instance_id"`
	Ipv6CidrBlock           string `json:"ipv6_cidr_block"`
	NatGatewayId            string `json:"nat_gateway_id"`
	NetworkInterfaceId      string `json:"network_interface_id"`
	TransitGatewayId        string `json:"transit_gateway_id"`
	VpcEndpointId           string `json:"vpc_endpoint_id"`
	VpcPeeringConnectionId  string `json:"vpc_peering_connection_id"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Update string `json:"update"`
}
