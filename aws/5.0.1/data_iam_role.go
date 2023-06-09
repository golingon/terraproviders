// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataiamrole "github.com/golingon/terraproviders/aws/5.0.1/dataiamrole"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataIamRole creates a new instance of [DataIamRole].
func NewDataIamRole(name string, args DataIamRoleArgs) *DataIamRole {
	return &DataIamRole{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataIamRole)(nil)

// DataIamRole represents the Terraform data resource aws_iam_role.
type DataIamRole struct {
	Name string
	Args DataIamRoleArgs
}

// DataSource returns the Terraform object type for [DataIamRole].
func (ir *DataIamRole) DataSource() string {
	return "aws_iam_role"
}

// LocalName returns the local name for [DataIamRole].
func (ir *DataIamRole) LocalName() string {
	return ir.Name
}

// Configuration returns the configuration (args) for [DataIamRole].
func (ir *DataIamRole) Configuration() interface{} {
	return ir.Args
}

// Attributes returns the attributes for [DataIamRole].
func (ir *DataIamRole) Attributes() dataIamRoleAttributes {
	return dataIamRoleAttributes{ref: terra.ReferenceDataResource(ir)}
}

// DataIamRoleArgs contains the configurations for aws_iam_role.
type DataIamRoleArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// RoleLastUsed: min=0
	RoleLastUsed []dataiamrole.RoleLastUsed `hcl:"role_last_used,block" validate:"min=0"`
}
type dataIamRoleAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_iam_role.
func (ir dataIamRoleAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("arn"))
}

// AssumeRolePolicy returns a reference to field assume_role_policy of aws_iam_role.
func (ir dataIamRoleAttributes) AssumeRolePolicy() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("assume_role_policy"))
}

// CreateDate returns a reference to field create_date of aws_iam_role.
func (ir dataIamRoleAttributes) CreateDate() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("create_date"))
}

// Description returns a reference to field description of aws_iam_role.
func (ir dataIamRoleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("description"))
}

// Id returns a reference to field id of aws_iam_role.
func (ir dataIamRoleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("id"))
}

// MaxSessionDuration returns a reference to field max_session_duration of aws_iam_role.
func (ir dataIamRoleAttributes) MaxSessionDuration() terra.NumberValue {
	return terra.ReferenceAsNumber(ir.ref.Append("max_session_duration"))
}

// Name returns a reference to field name of aws_iam_role.
func (ir dataIamRoleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("name"))
}

// Path returns a reference to field path of aws_iam_role.
func (ir dataIamRoleAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("path"))
}

// PermissionsBoundary returns a reference to field permissions_boundary of aws_iam_role.
func (ir dataIamRoleAttributes) PermissionsBoundary() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("permissions_boundary"))
}

// Tags returns a reference to field tags of aws_iam_role.
func (ir dataIamRoleAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ir.ref.Append("tags"))
}

// UniqueId returns a reference to field unique_id of aws_iam_role.
func (ir dataIamRoleAttributes) UniqueId() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("unique_id"))
}

func (ir dataIamRoleAttributes) RoleLastUsed() terra.ListValue[dataiamrole.RoleLastUsedAttributes] {
	return terra.ReferenceAsList[dataiamrole.RoleLastUsedAttributes](ir.ref.Append("role_last_used"))
}
