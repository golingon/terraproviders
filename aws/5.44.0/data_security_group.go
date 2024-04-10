// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"github.com/golingon/lingon/pkg/terra"
	datasecuritygroup "github.com/golingon/terraproviders/aws/5.44.0/datasecuritygroup"
)

// NewDataSecurityGroup creates a new instance of [DataSecurityGroup].
func NewDataSecurityGroup(name string, args DataSecurityGroupArgs) *DataSecurityGroup {
	return &DataSecurityGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSecurityGroup)(nil)

// DataSecurityGroup represents the Terraform data resource aws_security_group.
type DataSecurityGroup struct {
	Name string
	Args DataSecurityGroupArgs
}

// DataSource returns the Terraform object type for [DataSecurityGroup].
func (sg *DataSecurityGroup) DataSource() string {
	return "aws_security_group"
}

// LocalName returns the local name for [DataSecurityGroup].
func (sg *DataSecurityGroup) LocalName() string {
	return sg.Name
}

// Configuration returns the configuration (args) for [DataSecurityGroup].
func (sg *DataSecurityGroup) Configuration() interface{} {
	return sg.Args
}

// Attributes returns the attributes for [DataSecurityGroup].
func (sg *DataSecurityGroup) Attributes() dataSecurityGroupAttributes {
	return dataSecurityGroupAttributes{ref: terra.ReferenceDataResource(sg)}
}

// DataSecurityGroupArgs contains the configurations for aws_security_group.
type DataSecurityGroupArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// VpcId: string, optional
	VpcId terra.StringValue `hcl:"vpc_id,attr"`
	// Filter: min=0
	Filter []datasecuritygroup.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datasecuritygroup.Timeouts `hcl:"timeouts,block"`
}
type dataSecurityGroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_security_group.
func (sg dataSecurityGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sg.ref.Append("arn"))
}

// Description returns a reference to field description of aws_security_group.
func (sg dataSecurityGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(sg.ref.Append("description"))
}

// Id returns a reference to field id of aws_security_group.
func (sg dataSecurityGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sg.ref.Append("id"))
}

// Name returns a reference to field name of aws_security_group.
func (sg dataSecurityGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sg.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_security_group.
func (sg dataSecurityGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sg.ref.Append("tags"))
}

// VpcId returns a reference to field vpc_id of aws_security_group.
func (sg dataSecurityGroupAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(sg.ref.Append("vpc_id"))
}

func (sg dataSecurityGroupAttributes) Filter() terra.SetValue[datasecuritygroup.FilterAttributes] {
	return terra.ReferenceAsSet[datasecuritygroup.FilterAttributes](sg.ref.Append("filter"))
}

func (sg dataSecurityGroupAttributes) Timeouts() datasecuritygroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datasecuritygroup.TimeoutsAttributes](sg.ref.Append("timeouts"))
}
