// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataroutetable

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Associations struct{}

type Routes struct{}

type Filter struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Values: set of string, required
	Values terra.SetValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type AssociationsAttributes struct {
	ref terra.Reference
}

func (a AssociationsAttributes) InternalRef() terra.Reference {
	return a.ref
}

func (a AssociationsAttributes) InternalWithRef(ref terra.Reference) AssociationsAttributes {
	return AssociationsAttributes{ref: ref}
}

func (a AssociationsAttributes) InternalTokens() hclwrite.Tokens {
	return a.ref.InternalTokens()
}

func (a AssociationsAttributes) GatewayId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("gateway_id"))
}

func (a AssociationsAttributes) Main() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("main"))
}

func (a AssociationsAttributes) RouteTableAssociationId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("route_table_association_id"))
}

func (a AssociationsAttributes) RouteTableId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("route_table_id"))
}

func (a AssociationsAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("subnet_id"))
}

type RoutesAttributes struct {
	ref terra.Reference
}

func (r RoutesAttributes) InternalRef() terra.Reference {
	return r.ref
}

func (r RoutesAttributes) InternalWithRef(ref terra.Reference) RoutesAttributes {
	return RoutesAttributes{ref: ref}
}

func (r RoutesAttributes) InternalTokens() hclwrite.Tokens {
	return r.ref.InternalTokens()
}

func (r RoutesAttributes) CarrierGatewayId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("carrier_gateway_id"))
}

func (r RoutesAttributes) CidrBlock() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("cidr_block"))
}

func (r RoutesAttributes) CoreNetworkArn() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("core_network_arn"))
}

func (r RoutesAttributes) DestinationPrefixListId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("destination_prefix_list_id"))
}

func (r RoutesAttributes) EgressOnlyGatewayId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("egress_only_gateway_id"))
}

func (r RoutesAttributes) GatewayId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("gateway_id"))
}

func (r RoutesAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("instance_id"))
}

func (r RoutesAttributes) Ipv6CidrBlock() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("ipv6_cidr_block"))
}

func (r RoutesAttributes) LocalGatewayId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("local_gateway_id"))
}

func (r RoutesAttributes) NatGatewayId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("nat_gateway_id"))
}

func (r RoutesAttributes) NetworkInterfaceId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("network_interface_id"))
}

func (r RoutesAttributes) TransitGatewayId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("transit_gateway_id"))
}

func (r RoutesAttributes) VpcEndpointId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("vpc_endpoint_id"))
}

func (r RoutesAttributes) VpcPeeringConnectionId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("vpc_peering_connection_id"))
}

type FilterAttributes struct {
	ref terra.Reference
}

func (f FilterAttributes) InternalRef() terra.Reference {
	return f.ref
}

func (f FilterAttributes) InternalWithRef(ref terra.Reference) FilterAttributes {
	return FilterAttributes{ref: ref}
}

func (f FilterAttributes) InternalTokens() hclwrite.Tokens {
	return f.ref.InternalTokens()
}

func (f FilterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("name"))
}

func (f FilterAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](f.ref.Append("values"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type AssociationsState struct {
	GatewayId               string `json:"gateway_id"`
	Main                    bool   `json:"main"`
	RouteTableAssociationId string `json:"route_table_association_id"`
	RouteTableId            string `json:"route_table_id"`
	SubnetId                string `json:"subnet_id"`
}

type RoutesState struct {
	CarrierGatewayId        string `json:"carrier_gateway_id"`
	CidrBlock               string `json:"cidr_block"`
	CoreNetworkArn          string `json:"core_network_arn"`
	DestinationPrefixListId string `json:"destination_prefix_list_id"`
	EgressOnlyGatewayId     string `json:"egress_only_gateway_id"`
	GatewayId               string `json:"gateway_id"`
	InstanceId              string `json:"instance_id"`
	Ipv6CidrBlock           string `json:"ipv6_cidr_block"`
	LocalGatewayId          string `json:"local_gateway_id"`
	NatGatewayId            string `json:"nat_gateway_id"`
	NetworkInterfaceId      string `json:"network_interface_id"`
	TransitGatewayId        string `json:"transit_gateway_id"`
	VpcEndpointId           string `json:"vpc_endpoint_id"`
	VpcPeeringConnectionId  string `json:"vpc_peering_connection_id"`
}

type FilterState struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}