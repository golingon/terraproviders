// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataec2networkinsightspath "github.com/golingon/terraproviders/aws/5.9.0/dataec2networkinsightspath"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEc2NetworkInsightsPath creates a new instance of [DataEc2NetworkInsightsPath].
func NewDataEc2NetworkInsightsPath(name string, args DataEc2NetworkInsightsPathArgs) *DataEc2NetworkInsightsPath {
	return &DataEc2NetworkInsightsPath{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEc2NetworkInsightsPath)(nil)

// DataEc2NetworkInsightsPath represents the Terraform data resource aws_ec2_network_insights_path.
type DataEc2NetworkInsightsPath struct {
	Name string
	Args DataEc2NetworkInsightsPathArgs
}

// DataSource returns the Terraform object type for [DataEc2NetworkInsightsPath].
func (enip *DataEc2NetworkInsightsPath) DataSource() string {
	return "aws_ec2_network_insights_path"
}

// LocalName returns the local name for [DataEc2NetworkInsightsPath].
func (enip *DataEc2NetworkInsightsPath) LocalName() string {
	return enip.Name
}

// Configuration returns the configuration (args) for [DataEc2NetworkInsightsPath].
func (enip *DataEc2NetworkInsightsPath) Configuration() interface{} {
	return enip.Args
}

// Attributes returns the attributes for [DataEc2NetworkInsightsPath].
func (enip *DataEc2NetworkInsightsPath) Attributes() dataEc2NetworkInsightsPathAttributes {
	return dataEc2NetworkInsightsPathAttributes{ref: terra.ReferenceDataResource(enip)}
}

// DataEc2NetworkInsightsPathArgs contains the configurations for aws_ec2_network_insights_path.
type DataEc2NetworkInsightsPathArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NetworkInsightsPathId: string, optional
	NetworkInsightsPathId terra.StringValue `hcl:"network_insights_path_id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Filter: min=0
	Filter []dataec2networkinsightspath.Filter `hcl:"filter,block" validate:"min=0"`
}
type dataEc2NetworkInsightsPathAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ec2_network_insights_path.
func (enip dataEc2NetworkInsightsPathAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(enip.ref.Append("arn"))
}

// Destination returns a reference to field destination of aws_ec2_network_insights_path.
func (enip dataEc2NetworkInsightsPathAttributes) Destination() terra.StringValue {
	return terra.ReferenceAsString(enip.ref.Append("destination"))
}

// DestinationIp returns a reference to field destination_ip of aws_ec2_network_insights_path.
func (enip dataEc2NetworkInsightsPathAttributes) DestinationIp() terra.StringValue {
	return terra.ReferenceAsString(enip.ref.Append("destination_ip"))
}

// DestinationPort returns a reference to field destination_port of aws_ec2_network_insights_path.
func (enip dataEc2NetworkInsightsPathAttributes) DestinationPort() terra.NumberValue {
	return terra.ReferenceAsNumber(enip.ref.Append("destination_port"))
}

// Id returns a reference to field id of aws_ec2_network_insights_path.
func (enip dataEc2NetworkInsightsPathAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(enip.ref.Append("id"))
}

// NetworkInsightsPathId returns a reference to field network_insights_path_id of aws_ec2_network_insights_path.
func (enip dataEc2NetworkInsightsPathAttributes) NetworkInsightsPathId() terra.StringValue {
	return terra.ReferenceAsString(enip.ref.Append("network_insights_path_id"))
}

// Protocol returns a reference to field protocol of aws_ec2_network_insights_path.
func (enip dataEc2NetworkInsightsPathAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(enip.ref.Append("protocol"))
}

// Source returns a reference to field source of aws_ec2_network_insights_path.
func (enip dataEc2NetworkInsightsPathAttributes) Source() terra.StringValue {
	return terra.ReferenceAsString(enip.ref.Append("source"))
}

// SourceIp returns a reference to field source_ip of aws_ec2_network_insights_path.
func (enip dataEc2NetworkInsightsPathAttributes) SourceIp() terra.StringValue {
	return terra.ReferenceAsString(enip.ref.Append("source_ip"))
}

// Tags returns a reference to field tags of aws_ec2_network_insights_path.
func (enip dataEc2NetworkInsightsPathAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](enip.ref.Append("tags"))
}

func (enip dataEc2NetworkInsightsPathAttributes) Filter() terra.SetValue[dataec2networkinsightspath.FilterAttributes] {
	return terra.ReferenceAsSet[dataec2networkinsightspath.FilterAttributes](enip.ref.Append("filter"))
}
