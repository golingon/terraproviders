// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataIamUser creates a new instance of [DataIamUser].
func NewDataIamUser(name string, args DataIamUserArgs) *DataIamUser {
	return &DataIamUser{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataIamUser)(nil)

// DataIamUser represents the Terraform data resource aws_iam_user.
type DataIamUser struct {
	Name string
	Args DataIamUserArgs
}

// DataSource returns the Terraform object type for [DataIamUser].
func (iu *DataIamUser) DataSource() string {
	return "aws_iam_user"
}

// LocalName returns the local name for [DataIamUser].
func (iu *DataIamUser) LocalName() string {
	return iu.Name
}

// Configuration returns the configuration (args) for [DataIamUser].
func (iu *DataIamUser) Configuration() interface{} {
	return iu.Args
}

// Attributes returns the attributes for [DataIamUser].
func (iu *DataIamUser) Attributes() dataIamUserAttributes {
	return dataIamUserAttributes{ref: terra.ReferenceDataResource(iu)}
}

// DataIamUserArgs contains the configurations for aws_iam_user.
type DataIamUserArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// UserName: string, required
	UserName terra.StringValue `hcl:"user_name,attr" validate:"required"`
}
type dataIamUserAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_iam_user.
func (iu dataIamUserAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(iu.ref.Append("arn"))
}

// Id returns a reference to field id of aws_iam_user.
func (iu dataIamUserAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iu.ref.Append("id"))
}

// Path returns a reference to field path of aws_iam_user.
func (iu dataIamUserAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(iu.ref.Append("path"))
}

// PermissionsBoundary returns a reference to field permissions_boundary of aws_iam_user.
func (iu dataIamUserAttributes) PermissionsBoundary() terra.StringValue {
	return terra.ReferenceAsString(iu.ref.Append("permissions_boundary"))
}

// Tags returns a reference to field tags of aws_iam_user.
func (iu dataIamUserAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](iu.ref.Append("tags"))
}

// UserId returns a reference to field user_id of aws_iam_user.
func (iu dataIamUserAttributes) UserId() terra.StringValue {
	return terra.ReferenceAsString(iu.ref.Append("user_id"))
}

// UserName returns a reference to field user_name of aws_iam_user.
func (iu dataIamUserAttributes) UserName() terra.StringValue {
	return terra.ReferenceAsString(iu.ref.Append("user_name"))
}