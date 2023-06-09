// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataCloudwatchLogGroups creates a new instance of [DataCloudwatchLogGroups].
func NewDataCloudwatchLogGroups(name string, args DataCloudwatchLogGroupsArgs) *DataCloudwatchLogGroups {
	return &DataCloudwatchLogGroups{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCloudwatchLogGroups)(nil)

// DataCloudwatchLogGroups represents the Terraform data resource aws_cloudwatch_log_groups.
type DataCloudwatchLogGroups struct {
	Name string
	Args DataCloudwatchLogGroupsArgs
}

// DataSource returns the Terraform object type for [DataCloudwatchLogGroups].
func (clg *DataCloudwatchLogGroups) DataSource() string {
	return "aws_cloudwatch_log_groups"
}

// LocalName returns the local name for [DataCloudwatchLogGroups].
func (clg *DataCloudwatchLogGroups) LocalName() string {
	return clg.Name
}

// Configuration returns the configuration (args) for [DataCloudwatchLogGroups].
func (clg *DataCloudwatchLogGroups) Configuration() interface{} {
	return clg.Args
}

// Attributes returns the attributes for [DataCloudwatchLogGroups].
func (clg *DataCloudwatchLogGroups) Attributes() dataCloudwatchLogGroupsAttributes {
	return dataCloudwatchLogGroupsAttributes{ref: terra.ReferenceDataResource(clg)}
}

// DataCloudwatchLogGroupsArgs contains the configurations for aws_cloudwatch_log_groups.
type DataCloudwatchLogGroupsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogGroupNamePrefix: string, optional
	LogGroupNamePrefix terra.StringValue `hcl:"log_group_name_prefix,attr"`
}
type dataCloudwatchLogGroupsAttributes struct {
	ref terra.Reference
}

// Arns returns a reference to field arns of aws_cloudwatch_log_groups.
func (clg dataCloudwatchLogGroupsAttributes) Arns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](clg.ref.Append("arns"))
}

// Id returns a reference to field id of aws_cloudwatch_log_groups.
func (clg dataCloudwatchLogGroupsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(clg.ref.Append("id"))
}

// LogGroupNamePrefix returns a reference to field log_group_name_prefix of aws_cloudwatch_log_groups.
func (clg dataCloudwatchLogGroupsAttributes) LogGroupNamePrefix() terra.StringValue {
	return terra.ReferenceAsString(clg.ref.Append("log_group_name_prefix"))
}

// LogGroupNames returns a reference to field log_group_names of aws_cloudwatch_log_groups.
func (clg dataCloudwatchLogGroupsAttributes) LogGroupNames() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](clg.ref.Append("log_group_names"))
}
