// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datasecuritygroups "github.com/golingon/terraproviders/aws/4.60.0/datasecuritygroups"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSecurityGroups creates a new instance of [DataSecurityGroups].
func NewDataSecurityGroups(name string, args DataSecurityGroupsArgs) *DataSecurityGroups {
	return &DataSecurityGroups{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSecurityGroups)(nil)

// DataSecurityGroups represents the Terraform data resource aws_security_groups.
type DataSecurityGroups struct {
	Name string
	Args DataSecurityGroupsArgs
}

// DataSource returns the Terraform object type for [DataSecurityGroups].
func (sg *DataSecurityGroups) DataSource() string {
	return "aws_security_groups"
}

// LocalName returns the local name for [DataSecurityGroups].
func (sg *DataSecurityGroups) LocalName() string {
	return sg.Name
}

// Configuration returns the configuration (args) for [DataSecurityGroups].
func (sg *DataSecurityGroups) Configuration() interface{} {
	return sg.Args
}

// Attributes returns the attributes for [DataSecurityGroups].
func (sg *DataSecurityGroups) Attributes() dataSecurityGroupsAttributes {
	return dataSecurityGroupsAttributes{ref: terra.ReferenceDataResource(sg)}
}

// DataSecurityGroupsArgs contains the configurations for aws_security_groups.
type DataSecurityGroupsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Filter: min=0
	Filter []datasecuritygroups.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datasecuritygroups.Timeouts `hcl:"timeouts,block"`
}
type dataSecurityGroupsAttributes struct {
	ref terra.Reference
}

// Arns returns a reference to field arns of aws_security_groups.
func (sg dataSecurityGroupsAttributes) Arns() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sg.ref.Append("arns"))
}

// Id returns a reference to field id of aws_security_groups.
func (sg dataSecurityGroupsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sg.ref.Append("id"))
}

// Ids returns a reference to field ids of aws_security_groups.
func (sg dataSecurityGroupsAttributes) Ids() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sg.ref.Append("ids"))
}

// Tags returns a reference to field tags of aws_security_groups.
func (sg dataSecurityGroupsAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sg.ref.Append("tags"))
}

// VpcIds returns a reference to field vpc_ids of aws_security_groups.
func (sg dataSecurityGroupsAttributes) VpcIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sg.ref.Append("vpc_ids"))
}

func (sg dataSecurityGroupsAttributes) Filter() terra.SetValue[datasecuritygroups.FilterAttributes] {
	return terra.ReferenceAsSet[datasecuritygroups.FilterAttributes](sg.ref.Append("filter"))
}

func (sg dataSecurityGroupsAttributes) Timeouts() datasecuritygroups.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datasecuritygroups.TimeoutsAttributes](sg.ref.Append("timeouts"))
}
