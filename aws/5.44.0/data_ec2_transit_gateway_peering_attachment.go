// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"github.com/golingon/lingon/pkg/terra"
	dataec2transitgatewaypeeringattachment "github.com/golingon/terraproviders/aws/5.44.0/dataec2transitgatewaypeeringattachment"
)

// NewDataEc2TransitGatewayPeeringAttachment creates a new instance of [DataEc2TransitGatewayPeeringAttachment].
func NewDataEc2TransitGatewayPeeringAttachment(name string, args DataEc2TransitGatewayPeeringAttachmentArgs) *DataEc2TransitGatewayPeeringAttachment {
	return &DataEc2TransitGatewayPeeringAttachment{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEc2TransitGatewayPeeringAttachment)(nil)

// DataEc2TransitGatewayPeeringAttachment represents the Terraform data resource aws_ec2_transit_gateway_peering_attachment.
type DataEc2TransitGatewayPeeringAttachment struct {
	Name string
	Args DataEc2TransitGatewayPeeringAttachmentArgs
}

// DataSource returns the Terraform object type for [DataEc2TransitGatewayPeeringAttachment].
func (etgpa *DataEc2TransitGatewayPeeringAttachment) DataSource() string {
	return "aws_ec2_transit_gateway_peering_attachment"
}

// LocalName returns the local name for [DataEc2TransitGatewayPeeringAttachment].
func (etgpa *DataEc2TransitGatewayPeeringAttachment) LocalName() string {
	return etgpa.Name
}

// Configuration returns the configuration (args) for [DataEc2TransitGatewayPeeringAttachment].
func (etgpa *DataEc2TransitGatewayPeeringAttachment) Configuration() interface{} {
	return etgpa.Args
}

// Attributes returns the attributes for [DataEc2TransitGatewayPeeringAttachment].
func (etgpa *DataEc2TransitGatewayPeeringAttachment) Attributes() dataEc2TransitGatewayPeeringAttachmentAttributes {
	return dataEc2TransitGatewayPeeringAttachmentAttributes{ref: terra.ReferenceDataResource(etgpa)}
}

// DataEc2TransitGatewayPeeringAttachmentArgs contains the configurations for aws_ec2_transit_gateway_peering_attachment.
type DataEc2TransitGatewayPeeringAttachmentArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Filter: min=0
	Filter []dataec2transitgatewaypeeringattachment.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataec2transitgatewaypeeringattachment.Timeouts `hcl:"timeouts,block"`
}
type dataEc2TransitGatewayPeeringAttachmentAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ec2_transit_gateway_peering_attachment.
func (etgpa dataEc2TransitGatewayPeeringAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(etgpa.ref.Append("id"))
}

// PeerAccountId returns a reference to field peer_account_id of aws_ec2_transit_gateway_peering_attachment.
func (etgpa dataEc2TransitGatewayPeeringAttachmentAttributes) PeerAccountId() terra.StringValue {
	return terra.ReferenceAsString(etgpa.ref.Append("peer_account_id"))
}

// PeerRegion returns a reference to field peer_region of aws_ec2_transit_gateway_peering_attachment.
func (etgpa dataEc2TransitGatewayPeeringAttachmentAttributes) PeerRegion() terra.StringValue {
	return terra.ReferenceAsString(etgpa.ref.Append("peer_region"))
}

// PeerTransitGatewayId returns a reference to field peer_transit_gateway_id of aws_ec2_transit_gateway_peering_attachment.
func (etgpa dataEc2TransitGatewayPeeringAttachmentAttributes) PeerTransitGatewayId() terra.StringValue {
	return terra.ReferenceAsString(etgpa.ref.Append("peer_transit_gateway_id"))
}

// State returns a reference to field state of aws_ec2_transit_gateway_peering_attachment.
func (etgpa dataEc2TransitGatewayPeeringAttachmentAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(etgpa.ref.Append("state"))
}

// Tags returns a reference to field tags of aws_ec2_transit_gateway_peering_attachment.
func (etgpa dataEc2TransitGatewayPeeringAttachmentAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](etgpa.ref.Append("tags"))
}

// TransitGatewayId returns a reference to field transit_gateway_id of aws_ec2_transit_gateway_peering_attachment.
func (etgpa dataEc2TransitGatewayPeeringAttachmentAttributes) TransitGatewayId() terra.StringValue {
	return terra.ReferenceAsString(etgpa.ref.Append("transit_gateway_id"))
}

func (etgpa dataEc2TransitGatewayPeeringAttachmentAttributes) Filter() terra.SetValue[dataec2transitgatewaypeeringattachment.FilterAttributes] {
	return terra.ReferenceAsSet[dataec2transitgatewaypeeringattachment.FilterAttributes](etgpa.ref.Append("filter"))
}

func (etgpa dataEc2TransitGatewayPeeringAttachmentAttributes) Timeouts() dataec2transitgatewaypeeringattachment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataec2transitgatewaypeeringattachment.TimeoutsAttributes](etgpa.ref.Append("timeouts"))
}
