// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataroute "github.com/golingon/terraproviders/aws/5.11.0/dataroute"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataRoute creates a new instance of [DataRoute].
func NewDataRoute(name string, args DataRouteArgs) *DataRoute {
	return &DataRoute{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataRoute)(nil)

// DataRoute represents the Terraform data resource aws_route.
type DataRoute struct {
	Name string
	Args DataRouteArgs
}

// DataSource returns the Terraform object type for [DataRoute].
func (r *DataRoute) DataSource() string {
	return "aws_route"
}

// LocalName returns the local name for [DataRoute].
func (r *DataRoute) LocalName() string {
	return r.Name
}

// Configuration returns the configuration (args) for [DataRoute].
func (r *DataRoute) Configuration() interface{} {
	return r.Args
}

// Attributes returns the attributes for [DataRoute].
func (r *DataRoute) Attributes() dataRouteAttributes {
	return dataRouteAttributes{ref: terra.ReferenceDataResource(r)}
}

// DataRouteArgs contains the configurations for aws_route.
type DataRouteArgs struct {
	// CarrierGatewayId: string, optional
	CarrierGatewayId terra.StringValue `hcl:"carrier_gateway_id,attr"`
	// CoreNetworkArn: string, optional
	CoreNetworkArn terra.StringValue `hcl:"core_network_arn,attr"`
	// DestinationCidrBlock: string, optional
	DestinationCidrBlock terra.StringValue `hcl:"destination_cidr_block,attr"`
	// DestinationIpv6CidrBlock: string, optional
	DestinationIpv6CidrBlock terra.StringValue `hcl:"destination_ipv6_cidr_block,attr"`
	// DestinationPrefixListId: string, optional
	DestinationPrefixListId terra.StringValue `hcl:"destination_prefix_list_id,attr"`
	// EgressOnlyGatewayId: string, optional
	EgressOnlyGatewayId terra.StringValue `hcl:"egress_only_gateway_id,attr"`
	// GatewayId: string, optional
	GatewayId terra.StringValue `hcl:"gateway_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceId: string, optional
	InstanceId terra.StringValue `hcl:"instance_id,attr"`
	// LocalGatewayId: string, optional
	LocalGatewayId terra.StringValue `hcl:"local_gateway_id,attr"`
	// NatGatewayId: string, optional
	NatGatewayId terra.StringValue `hcl:"nat_gateway_id,attr"`
	// NetworkInterfaceId: string, optional
	NetworkInterfaceId terra.StringValue `hcl:"network_interface_id,attr"`
	// RouteTableId: string, required
	RouteTableId terra.StringValue `hcl:"route_table_id,attr" validate:"required"`
	// TransitGatewayId: string, optional
	TransitGatewayId terra.StringValue `hcl:"transit_gateway_id,attr"`
	// VpcPeeringConnectionId: string, optional
	VpcPeeringConnectionId terra.StringValue `hcl:"vpc_peering_connection_id,attr"`
	// Timeouts: optional
	Timeouts *dataroute.Timeouts `hcl:"timeouts,block"`
}
type dataRouteAttributes struct {
	ref terra.Reference
}

// CarrierGatewayId returns a reference to field carrier_gateway_id of aws_route.
func (r dataRouteAttributes) CarrierGatewayId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("carrier_gateway_id"))
}

// CoreNetworkArn returns a reference to field core_network_arn of aws_route.
func (r dataRouteAttributes) CoreNetworkArn() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("core_network_arn"))
}

// DestinationCidrBlock returns a reference to field destination_cidr_block of aws_route.
func (r dataRouteAttributes) DestinationCidrBlock() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("destination_cidr_block"))
}

// DestinationIpv6CidrBlock returns a reference to field destination_ipv6_cidr_block of aws_route.
func (r dataRouteAttributes) DestinationIpv6CidrBlock() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("destination_ipv6_cidr_block"))
}

// DestinationPrefixListId returns a reference to field destination_prefix_list_id of aws_route.
func (r dataRouteAttributes) DestinationPrefixListId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("destination_prefix_list_id"))
}

// EgressOnlyGatewayId returns a reference to field egress_only_gateway_id of aws_route.
func (r dataRouteAttributes) EgressOnlyGatewayId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("egress_only_gateway_id"))
}

// GatewayId returns a reference to field gateway_id of aws_route.
func (r dataRouteAttributes) GatewayId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("gateway_id"))
}

// Id returns a reference to field id of aws_route.
func (r dataRouteAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("id"))
}

// InstanceId returns a reference to field instance_id of aws_route.
func (r dataRouteAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("instance_id"))
}

// LocalGatewayId returns a reference to field local_gateway_id of aws_route.
func (r dataRouteAttributes) LocalGatewayId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("local_gateway_id"))
}

// NatGatewayId returns a reference to field nat_gateway_id of aws_route.
func (r dataRouteAttributes) NatGatewayId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("nat_gateway_id"))
}

// NetworkInterfaceId returns a reference to field network_interface_id of aws_route.
func (r dataRouteAttributes) NetworkInterfaceId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("network_interface_id"))
}

// RouteTableId returns a reference to field route_table_id of aws_route.
func (r dataRouteAttributes) RouteTableId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("route_table_id"))
}

// TransitGatewayId returns a reference to field transit_gateway_id of aws_route.
func (r dataRouteAttributes) TransitGatewayId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("transit_gateway_id"))
}

// VpcPeeringConnectionId returns a reference to field vpc_peering_connection_id of aws_route.
func (r dataRouteAttributes) VpcPeeringConnectionId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("vpc_peering_connection_id"))
}

func (r dataRouteAttributes) Timeouts() dataroute.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataroute.TimeoutsAttributes](r.ref.Append("timeouts"))
}
