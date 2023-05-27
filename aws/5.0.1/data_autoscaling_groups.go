// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataautoscalinggroups "github.com/golingon/terraproviders/aws/5.0.1/dataautoscalinggroups"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataAutoscalingGroups creates a new instance of [DataAutoscalingGroups].
func NewDataAutoscalingGroups(name string, args DataAutoscalingGroupsArgs) *DataAutoscalingGroups {
	return &DataAutoscalingGroups{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAutoscalingGroups)(nil)

// DataAutoscalingGroups represents the Terraform data resource aws_autoscaling_groups.
type DataAutoscalingGroups struct {
	Name string
	Args DataAutoscalingGroupsArgs
}

// DataSource returns the Terraform object type for [DataAutoscalingGroups].
func (ag *DataAutoscalingGroups) DataSource() string {
	return "aws_autoscaling_groups"
}

// LocalName returns the local name for [DataAutoscalingGroups].
func (ag *DataAutoscalingGroups) LocalName() string {
	return ag.Name
}

// Configuration returns the configuration (args) for [DataAutoscalingGroups].
func (ag *DataAutoscalingGroups) Configuration() interface{} {
	return ag.Args
}

// Attributes returns the attributes for [DataAutoscalingGroups].
func (ag *DataAutoscalingGroups) Attributes() dataAutoscalingGroupsAttributes {
	return dataAutoscalingGroupsAttributes{ref: terra.ReferenceDataResource(ag)}
}

// DataAutoscalingGroupsArgs contains the configurations for aws_autoscaling_groups.
type DataAutoscalingGroupsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Names: list of string, optional
	Names terra.ListValue[terra.StringValue] `hcl:"names,attr"`
	// Filter: min=0
	Filter []dataautoscalinggroups.Filter `hcl:"filter,block" validate:"min=0"`
}
type dataAutoscalingGroupsAttributes struct {
	ref terra.Reference
}

// Arns returns a reference to field arns of aws_autoscaling_groups.
func (ag dataAutoscalingGroupsAttributes) Arns() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ag.ref.Append("arns"))
}

// Id returns a reference to field id of aws_autoscaling_groups.
func (ag dataAutoscalingGroupsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ag.ref.Append("id"))
}

// Names returns a reference to field names of aws_autoscaling_groups.
func (ag dataAutoscalingGroupsAttributes) Names() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ag.ref.Append("names"))
}

func (ag dataAutoscalingGroupsAttributes) Filter() terra.SetValue[dataautoscalinggroups.FilterAttributes] {
	return terra.ReferenceAsSet[dataautoscalinggroups.FilterAttributes](ag.ref.Append("filter"))
}
