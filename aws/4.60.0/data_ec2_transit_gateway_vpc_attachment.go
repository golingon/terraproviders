// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataec2transitgatewayvpcattachment "github.com/golingon/terraproviders/aws/4.60.0/dataec2transitgatewayvpcattachment"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEc2TransitGatewayVpcAttachment creates a new instance of [DataEc2TransitGatewayVpcAttachment].
func NewDataEc2TransitGatewayVpcAttachment(name string, args DataEc2TransitGatewayVpcAttachmentArgs) *DataEc2TransitGatewayVpcAttachment {
	return &DataEc2TransitGatewayVpcAttachment{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEc2TransitGatewayVpcAttachment)(nil)

// DataEc2TransitGatewayVpcAttachment represents the Terraform data resource aws_ec2_transit_gateway_vpc_attachment.
type DataEc2TransitGatewayVpcAttachment struct {
	Name string
	Args DataEc2TransitGatewayVpcAttachmentArgs
}

// DataSource returns the Terraform object type for [DataEc2TransitGatewayVpcAttachment].
func (etgva *DataEc2TransitGatewayVpcAttachment) DataSource() string {
	return "aws_ec2_transit_gateway_vpc_attachment"
}

// LocalName returns the local name for [DataEc2TransitGatewayVpcAttachment].
func (etgva *DataEc2TransitGatewayVpcAttachment) LocalName() string {
	return etgva.Name
}

// Configuration returns the configuration (args) for [DataEc2TransitGatewayVpcAttachment].
func (etgva *DataEc2TransitGatewayVpcAttachment) Configuration() interface{} {
	return etgva.Args
}

// Attributes returns the attributes for [DataEc2TransitGatewayVpcAttachment].
func (etgva *DataEc2TransitGatewayVpcAttachment) Attributes() dataEc2TransitGatewayVpcAttachmentAttributes {
	return dataEc2TransitGatewayVpcAttachmentAttributes{ref: terra.ReferenceDataResource(etgva)}
}

// DataEc2TransitGatewayVpcAttachmentArgs contains the configurations for aws_ec2_transit_gateway_vpc_attachment.
type DataEc2TransitGatewayVpcAttachmentArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Filter: min=0
	Filter []dataec2transitgatewayvpcattachment.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataec2transitgatewayvpcattachment.Timeouts `hcl:"timeouts,block"`
}
type dataEc2TransitGatewayVpcAttachmentAttributes struct {
	ref terra.Reference
}

// ApplianceModeSupport returns a reference to field appliance_mode_support of aws_ec2_transit_gateway_vpc_attachment.
func (etgva dataEc2TransitGatewayVpcAttachmentAttributes) ApplianceModeSupport() terra.StringValue {
	return terra.ReferenceAsString(etgva.ref.Append("appliance_mode_support"))
}

// DnsSupport returns a reference to field dns_support of aws_ec2_transit_gateway_vpc_attachment.
func (etgva dataEc2TransitGatewayVpcAttachmentAttributes) DnsSupport() terra.StringValue {
	return terra.ReferenceAsString(etgva.ref.Append("dns_support"))
}

// Id returns a reference to field id of aws_ec2_transit_gateway_vpc_attachment.
func (etgva dataEc2TransitGatewayVpcAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(etgva.ref.Append("id"))
}

// Ipv6Support returns a reference to field ipv6_support of aws_ec2_transit_gateway_vpc_attachment.
func (etgva dataEc2TransitGatewayVpcAttachmentAttributes) Ipv6Support() terra.StringValue {
	return terra.ReferenceAsString(etgva.ref.Append("ipv6_support"))
}

// SubnetIds returns a reference to field subnet_ids of aws_ec2_transit_gateway_vpc_attachment.
func (etgva dataEc2TransitGatewayVpcAttachmentAttributes) SubnetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](etgva.ref.Append("subnet_ids"))
}

// Tags returns a reference to field tags of aws_ec2_transit_gateway_vpc_attachment.
func (etgva dataEc2TransitGatewayVpcAttachmentAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](etgva.ref.Append("tags"))
}

// TransitGatewayId returns a reference to field transit_gateway_id of aws_ec2_transit_gateway_vpc_attachment.
func (etgva dataEc2TransitGatewayVpcAttachmentAttributes) TransitGatewayId() terra.StringValue {
	return terra.ReferenceAsString(etgva.ref.Append("transit_gateway_id"))
}

// VpcId returns a reference to field vpc_id of aws_ec2_transit_gateway_vpc_attachment.
func (etgva dataEc2TransitGatewayVpcAttachmentAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(etgva.ref.Append("vpc_id"))
}

// VpcOwnerId returns a reference to field vpc_owner_id of aws_ec2_transit_gateway_vpc_attachment.
func (etgva dataEc2TransitGatewayVpcAttachmentAttributes) VpcOwnerId() terra.StringValue {
	return terra.ReferenceAsString(etgva.ref.Append("vpc_owner_id"))
}

func (etgva dataEc2TransitGatewayVpcAttachmentAttributes) Filter() terra.SetValue[dataec2transitgatewayvpcattachment.FilterAttributes] {
	return terra.ReferenceAsSet[dataec2transitgatewayvpcattachment.FilterAttributes](etgva.ref.Append("filter"))
}

func (etgva dataEc2TransitGatewayVpcAttachmentAttributes) Timeouts() dataec2transitgatewayvpcattachment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataec2transitgatewayvpcattachment.TimeoutsAttributes](etgva.ref.Append("timeouts"))
}
