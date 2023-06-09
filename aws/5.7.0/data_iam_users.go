// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataIamUsers creates a new instance of [DataIamUsers].
func NewDataIamUsers(name string, args DataIamUsersArgs) *DataIamUsers {
	return &DataIamUsers{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataIamUsers)(nil)

// DataIamUsers represents the Terraform data resource aws_iam_users.
type DataIamUsers struct {
	Name string
	Args DataIamUsersArgs
}

// DataSource returns the Terraform object type for [DataIamUsers].
func (iu *DataIamUsers) DataSource() string {
	return "aws_iam_users"
}

// LocalName returns the local name for [DataIamUsers].
func (iu *DataIamUsers) LocalName() string {
	return iu.Name
}

// Configuration returns the configuration (args) for [DataIamUsers].
func (iu *DataIamUsers) Configuration() interface{} {
	return iu.Args
}

// Attributes returns the attributes for [DataIamUsers].
func (iu *DataIamUsers) Attributes() dataIamUsersAttributes {
	return dataIamUsersAttributes{ref: terra.ReferenceDataResource(iu)}
}

// DataIamUsersArgs contains the configurations for aws_iam_users.
type DataIamUsersArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NameRegex: string, optional
	NameRegex terra.StringValue `hcl:"name_regex,attr"`
	// PathPrefix: string, optional
	PathPrefix terra.StringValue `hcl:"path_prefix,attr"`
}
type dataIamUsersAttributes struct {
	ref terra.Reference
}

// Arns returns a reference to field arns of aws_iam_users.
func (iu dataIamUsersAttributes) Arns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](iu.ref.Append("arns"))
}

// Id returns a reference to field id of aws_iam_users.
func (iu dataIamUsersAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iu.ref.Append("id"))
}

// NameRegex returns a reference to field name_regex of aws_iam_users.
func (iu dataIamUsersAttributes) NameRegex() terra.StringValue {
	return terra.ReferenceAsString(iu.ref.Append("name_regex"))
}

// Names returns a reference to field names of aws_iam_users.
func (iu dataIamUsersAttributes) Names() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](iu.ref.Append("names"))
}

// PathPrefix returns a reference to field path_prefix of aws_iam_users.
func (iu dataIamUsersAttributes) PathPrefix() terra.StringValue {
	return terra.ReferenceAsString(iu.ref.Append("path_prefix"))
}
