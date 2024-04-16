// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_security_group

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_security_group.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (asg *DataSource) DataSource() string {
	return "aws_security_group"
}

// LocalName returns the local name for [DataSource].
func (asg *DataSource) LocalName() string {
	return asg.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (asg *DataSource) Configuration() interface{} {
	return asg.Args
}

// Attributes returns the attributes for [DataSource].
func (asg *DataSource) Attributes() dataAwsSecurityGroupAttributes {
	return dataAwsSecurityGroupAttributes{ref: terra.ReferenceDataSource(asg)}
}

// DataArgs contains the configurations for aws_security_group.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// VpcId: string, optional
	VpcId terra.StringValue `hcl:"vpc_id,attr"`
	// Filter: min=0
	Filter []DataFilter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAwsSecurityGroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_security_group.
func (asg dataAwsSecurityGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(asg.ref.Append("arn"))
}

// Description returns a reference to field description of aws_security_group.
func (asg dataAwsSecurityGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(asg.ref.Append("description"))
}

// Id returns a reference to field id of aws_security_group.
func (asg dataAwsSecurityGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asg.ref.Append("id"))
}

// Name returns a reference to field name of aws_security_group.
func (asg dataAwsSecurityGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(asg.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_security_group.
func (asg dataAwsSecurityGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](asg.ref.Append("tags"))
}

// VpcId returns a reference to field vpc_id of aws_security_group.
func (asg dataAwsSecurityGroupAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(asg.ref.Append("vpc_id"))
}

func (asg dataAwsSecurityGroupAttributes) Filter() terra.SetValue[DataFilterAttributes] {
	return terra.ReferenceAsSet[DataFilterAttributes](asg.ref.Append("filter"))
}

func (asg dataAwsSecurityGroupAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](asg.ref.Append("timeouts"))
}
