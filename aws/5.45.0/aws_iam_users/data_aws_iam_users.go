// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_iam_users

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_iam_users.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (aiu *DataSource) DataSource() string {
	return "aws_iam_users"
}

// LocalName returns the local name for [DataSource].
func (aiu *DataSource) LocalName() string {
	return aiu.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (aiu *DataSource) Configuration() interface{} {
	return aiu.Args
}

// Attributes returns the attributes for [DataSource].
func (aiu *DataSource) Attributes() dataAwsIamUsersAttributes {
	return dataAwsIamUsersAttributes{ref: terra.ReferenceDataSource(aiu)}
}

// DataArgs contains the configurations for aws_iam_users.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NameRegex: string, optional
	NameRegex terra.StringValue `hcl:"name_regex,attr"`
	// PathPrefix: string, optional
	PathPrefix terra.StringValue `hcl:"path_prefix,attr"`
}

type dataAwsIamUsersAttributes struct {
	ref terra.Reference
}

// Arns returns a reference to field arns of aws_iam_users.
func (aiu dataAwsIamUsersAttributes) Arns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](aiu.ref.Append("arns"))
}

// Id returns a reference to field id of aws_iam_users.
func (aiu dataAwsIamUsersAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aiu.ref.Append("id"))
}

// NameRegex returns a reference to field name_regex of aws_iam_users.
func (aiu dataAwsIamUsersAttributes) NameRegex() terra.StringValue {
	return terra.ReferenceAsString(aiu.ref.Append("name_regex"))
}

// Names returns a reference to field names of aws_iam_users.
func (aiu dataAwsIamUsersAttributes) Names() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](aiu.ref.Append("names"))
}

// PathPrefix returns a reference to field path_prefix of aws_iam_users.
func (aiu dataAwsIamUsersAttributes) PathPrefix() terra.StringValue {
	return terra.ReferenceAsString(aiu.ref.Append("path_prefix"))
}
