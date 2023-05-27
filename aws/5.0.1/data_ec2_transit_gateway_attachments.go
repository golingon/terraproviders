// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataec2transitgatewayattachments "github.com/golingon/terraproviders/aws/5.0.1/dataec2transitgatewayattachments"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEc2TransitGatewayAttachments creates a new instance of [DataEc2TransitGatewayAttachments].
func NewDataEc2TransitGatewayAttachments(name string, args DataEc2TransitGatewayAttachmentsArgs) *DataEc2TransitGatewayAttachments {
	return &DataEc2TransitGatewayAttachments{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEc2TransitGatewayAttachments)(nil)

// DataEc2TransitGatewayAttachments represents the Terraform data resource aws_ec2_transit_gateway_attachments.
type DataEc2TransitGatewayAttachments struct {
	Name string
	Args DataEc2TransitGatewayAttachmentsArgs
}

// DataSource returns the Terraform object type for [DataEc2TransitGatewayAttachments].
func (etga *DataEc2TransitGatewayAttachments) DataSource() string {
	return "aws_ec2_transit_gateway_attachments"
}

// LocalName returns the local name for [DataEc2TransitGatewayAttachments].
func (etga *DataEc2TransitGatewayAttachments) LocalName() string {
	return etga.Name
}

// Configuration returns the configuration (args) for [DataEc2TransitGatewayAttachments].
func (etga *DataEc2TransitGatewayAttachments) Configuration() interface{} {
	return etga.Args
}

// Attributes returns the attributes for [DataEc2TransitGatewayAttachments].
func (etga *DataEc2TransitGatewayAttachments) Attributes() dataEc2TransitGatewayAttachmentsAttributes {
	return dataEc2TransitGatewayAttachmentsAttributes{ref: terra.ReferenceDataResource(etga)}
}

// DataEc2TransitGatewayAttachmentsArgs contains the configurations for aws_ec2_transit_gateway_attachments.
type DataEc2TransitGatewayAttachmentsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Filter: min=0
	Filter []dataec2transitgatewayattachments.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataec2transitgatewayattachments.Timeouts `hcl:"timeouts,block"`
}
type dataEc2TransitGatewayAttachmentsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ec2_transit_gateway_attachments.
func (etga dataEc2TransitGatewayAttachmentsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(etga.ref.Append("id"))
}

// Ids returns a reference to field ids of aws_ec2_transit_gateway_attachments.
func (etga dataEc2TransitGatewayAttachmentsAttributes) Ids() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](etga.ref.Append("ids"))
}

// Tags returns a reference to field tags of aws_ec2_transit_gateway_attachments.
func (etga dataEc2TransitGatewayAttachmentsAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](etga.ref.Append("tags"))
}

func (etga dataEc2TransitGatewayAttachmentsAttributes) Filter() terra.SetValue[dataec2transitgatewayattachments.FilterAttributes] {
	return terra.ReferenceAsSet[dataec2transitgatewayattachments.FilterAttributes](etga.ref.Append("filter"))
}

func (etga dataEc2TransitGatewayAttachmentsAttributes) Timeouts() dataec2transitgatewayattachments.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataec2transitgatewayattachments.TimeoutsAttributes](etga.ref.Append("timeouts"))
}
