// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_cognito_user_group

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_cognito_user_group.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (acug *DataSource) DataSource() string {
	return "aws_cognito_user_group"
}

// LocalName returns the local name for [DataSource].
func (acug *DataSource) LocalName() string {
	return acug.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (acug *DataSource) Configuration() interface{} {
	return acug.Args
}

// Attributes returns the attributes for [DataSource].
func (acug *DataSource) Attributes() dataAwsCognitoUserGroupAttributes {
	return dataAwsCognitoUserGroupAttributes{ref: terra.ReferenceDataSource(acug)}
}

// DataArgs contains the configurations for aws_cognito_user_group.
type DataArgs struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// UserPoolId: string, required
	UserPoolId terra.StringValue `hcl:"user_pool_id,attr" validate:"required"`
}

type dataAwsCognitoUserGroupAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of aws_cognito_user_group.
func (acug dataAwsCognitoUserGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(acug.ref.Append("description"))
}

// Id returns a reference to field id of aws_cognito_user_group.
func (acug dataAwsCognitoUserGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acug.ref.Append("id"))
}

// Name returns a reference to field name of aws_cognito_user_group.
func (acug dataAwsCognitoUserGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(acug.ref.Append("name"))
}

// Precedence returns a reference to field precedence of aws_cognito_user_group.
func (acug dataAwsCognitoUserGroupAttributes) Precedence() terra.NumberValue {
	return terra.ReferenceAsNumber(acug.ref.Append("precedence"))
}

// RoleArn returns a reference to field role_arn of aws_cognito_user_group.
func (acug dataAwsCognitoUserGroupAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(acug.ref.Append("role_arn"))
}

// UserPoolId returns a reference to field user_pool_id of aws_cognito_user_group.
func (acug dataAwsCognitoUserGroupAttributes) UserPoolId() terra.StringValue {
	return terra.ReferenceAsString(acug.ref.Append("user_pool_id"))
}
