// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataec2transitgateway "github.com/golingon/terraproviders/aws/5.9.0/dataec2transitgateway"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEc2TransitGateway creates a new instance of [DataEc2TransitGateway].
func NewDataEc2TransitGateway(name string, args DataEc2TransitGatewayArgs) *DataEc2TransitGateway {
	return &DataEc2TransitGateway{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEc2TransitGateway)(nil)

// DataEc2TransitGateway represents the Terraform data resource aws_ec2_transit_gateway.
type DataEc2TransitGateway struct {
	Name string
	Args DataEc2TransitGatewayArgs
}

// DataSource returns the Terraform object type for [DataEc2TransitGateway].
func (etg *DataEc2TransitGateway) DataSource() string {
	return "aws_ec2_transit_gateway"
}

// LocalName returns the local name for [DataEc2TransitGateway].
func (etg *DataEc2TransitGateway) LocalName() string {
	return etg.Name
}

// Configuration returns the configuration (args) for [DataEc2TransitGateway].
func (etg *DataEc2TransitGateway) Configuration() interface{} {
	return etg.Args
}

// Attributes returns the attributes for [DataEc2TransitGateway].
func (etg *DataEc2TransitGateway) Attributes() dataEc2TransitGatewayAttributes {
	return dataEc2TransitGatewayAttributes{ref: terra.ReferenceDataResource(etg)}
}

// DataEc2TransitGatewayArgs contains the configurations for aws_ec2_transit_gateway.
type DataEc2TransitGatewayArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Filter: min=0
	Filter []dataec2transitgateway.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataec2transitgateway.Timeouts `hcl:"timeouts,block"`
}
type dataEc2TransitGatewayAttributes struct {
	ref terra.Reference
}

// AmazonSideAsn returns a reference to field amazon_side_asn of aws_ec2_transit_gateway.
func (etg dataEc2TransitGatewayAttributes) AmazonSideAsn() terra.NumberValue {
	return terra.ReferenceAsNumber(etg.ref.Append("amazon_side_asn"))
}

// Arn returns a reference to field arn of aws_ec2_transit_gateway.
func (etg dataEc2TransitGatewayAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(etg.ref.Append("arn"))
}

// AssociationDefaultRouteTableId returns a reference to field association_default_route_table_id of aws_ec2_transit_gateway.
func (etg dataEc2TransitGatewayAttributes) AssociationDefaultRouteTableId() terra.StringValue {
	return terra.ReferenceAsString(etg.ref.Append("association_default_route_table_id"))
}

// AutoAcceptSharedAttachments returns a reference to field auto_accept_shared_attachments of aws_ec2_transit_gateway.
func (etg dataEc2TransitGatewayAttributes) AutoAcceptSharedAttachments() terra.StringValue {
	return terra.ReferenceAsString(etg.ref.Append("auto_accept_shared_attachments"))
}

// DefaultRouteTableAssociation returns a reference to field default_route_table_association of aws_ec2_transit_gateway.
func (etg dataEc2TransitGatewayAttributes) DefaultRouteTableAssociation() terra.StringValue {
	return terra.ReferenceAsString(etg.ref.Append("default_route_table_association"))
}

// DefaultRouteTablePropagation returns a reference to field default_route_table_propagation of aws_ec2_transit_gateway.
func (etg dataEc2TransitGatewayAttributes) DefaultRouteTablePropagation() terra.StringValue {
	return terra.ReferenceAsString(etg.ref.Append("default_route_table_propagation"))
}

// Description returns a reference to field description of aws_ec2_transit_gateway.
func (etg dataEc2TransitGatewayAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(etg.ref.Append("description"))
}

// DnsSupport returns a reference to field dns_support of aws_ec2_transit_gateway.
func (etg dataEc2TransitGatewayAttributes) DnsSupport() terra.StringValue {
	return terra.ReferenceAsString(etg.ref.Append("dns_support"))
}

// Id returns a reference to field id of aws_ec2_transit_gateway.
func (etg dataEc2TransitGatewayAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(etg.ref.Append("id"))
}

// MulticastSupport returns a reference to field multicast_support of aws_ec2_transit_gateway.
func (etg dataEc2TransitGatewayAttributes) MulticastSupport() terra.StringValue {
	return terra.ReferenceAsString(etg.ref.Append("multicast_support"))
}

// OwnerId returns a reference to field owner_id of aws_ec2_transit_gateway.
func (etg dataEc2TransitGatewayAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(etg.ref.Append("owner_id"))
}

// PropagationDefaultRouteTableId returns a reference to field propagation_default_route_table_id of aws_ec2_transit_gateway.
func (etg dataEc2TransitGatewayAttributes) PropagationDefaultRouteTableId() terra.StringValue {
	return terra.ReferenceAsString(etg.ref.Append("propagation_default_route_table_id"))
}

// Tags returns a reference to field tags of aws_ec2_transit_gateway.
func (etg dataEc2TransitGatewayAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](etg.ref.Append("tags"))
}

// TransitGatewayCidrBlocks returns a reference to field transit_gateway_cidr_blocks of aws_ec2_transit_gateway.
func (etg dataEc2TransitGatewayAttributes) TransitGatewayCidrBlocks() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](etg.ref.Append("transit_gateway_cidr_blocks"))
}

// VpnEcmpSupport returns a reference to field vpn_ecmp_support of aws_ec2_transit_gateway.
func (etg dataEc2TransitGatewayAttributes) VpnEcmpSupport() terra.StringValue {
	return terra.ReferenceAsString(etg.ref.Append("vpn_ecmp_support"))
}

func (etg dataEc2TransitGatewayAttributes) Filter() terra.SetValue[dataec2transitgateway.FilterAttributes] {
	return terra.ReferenceAsSet[dataec2transitgateway.FilterAttributes](etg.ref.Append("filter"))
}

func (etg dataEc2TransitGatewayAttributes) Timeouts() dataec2transitgateway.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataec2transitgateway.TimeoutsAttributes](etg.ref.Append("timeouts"))
}
