// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_ec2_transit_gateway_attachments

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_ec2_transit_gateway_attachments.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (aetga *DataSource) DataSource() string {
	return "aws_ec2_transit_gateway_attachments"
}

// LocalName returns the local name for [DataSource].
func (aetga *DataSource) LocalName() string {
	return aetga.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (aetga *DataSource) Configuration() interface{} {
	return aetga.Args
}

// Attributes returns the attributes for [DataSource].
func (aetga *DataSource) Attributes() dataAwsEc2TransitGatewayAttachmentsAttributes {
	return dataAwsEc2TransitGatewayAttachmentsAttributes{ref: terra.ReferenceDataSource(aetga)}
}

// DataArgs contains the configurations for aws_ec2_transit_gateway_attachments.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Filter: min=0
	Filter []DataFilter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAwsEc2TransitGatewayAttachmentsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ec2_transit_gateway_attachments.
func (aetga dataAwsEc2TransitGatewayAttachmentsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aetga.ref.Append("id"))
}

// Ids returns a reference to field ids of aws_ec2_transit_gateway_attachments.
func (aetga dataAwsEc2TransitGatewayAttachmentsAttributes) Ids() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](aetga.ref.Append("ids"))
}

// Tags returns a reference to field tags of aws_ec2_transit_gateway_attachments.
func (aetga dataAwsEc2TransitGatewayAttachmentsAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aetga.ref.Append("tags"))
}

func (aetga dataAwsEc2TransitGatewayAttachmentsAttributes) Filter() terra.SetValue[DataFilterAttributes] {
	return terra.ReferenceAsSet[DataFilterAttributes](aetga.ref.Append("filter"))
}

func (aetga dataAwsEc2TransitGatewayAttachmentsAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](aetga.ref.Append("timeouts"))
}
