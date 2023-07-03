// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataIamRoles creates a new instance of [DataIamRoles].
func NewDataIamRoles(name string, args DataIamRolesArgs) *DataIamRoles {
	return &DataIamRoles{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataIamRoles)(nil)

// DataIamRoles represents the Terraform data resource aws_iam_roles.
type DataIamRoles struct {
	Name string
	Args DataIamRolesArgs
}

// DataSource returns the Terraform object type for [DataIamRoles].
func (ir *DataIamRoles) DataSource() string {
	return "aws_iam_roles"
}

// LocalName returns the local name for [DataIamRoles].
func (ir *DataIamRoles) LocalName() string {
	return ir.Name
}

// Configuration returns the configuration (args) for [DataIamRoles].
func (ir *DataIamRoles) Configuration() interface{} {
	return ir.Args
}

// Attributes returns the attributes for [DataIamRoles].
func (ir *DataIamRoles) Attributes() dataIamRolesAttributes {
	return dataIamRolesAttributes{ref: terra.ReferenceDataResource(ir)}
}

// DataIamRolesArgs contains the configurations for aws_iam_roles.
type DataIamRolesArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NameRegex: string, optional
	NameRegex terra.StringValue `hcl:"name_regex,attr"`
	// PathPrefix: string, optional
	PathPrefix terra.StringValue `hcl:"path_prefix,attr"`
}
type dataIamRolesAttributes struct {
	ref terra.Reference
}

// Arns returns a reference to field arns of aws_iam_roles.
func (ir dataIamRolesAttributes) Arns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ir.ref.Append("arns"))
}

// Id returns a reference to field id of aws_iam_roles.
func (ir dataIamRolesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("id"))
}

// NameRegex returns a reference to field name_regex of aws_iam_roles.
func (ir dataIamRolesAttributes) NameRegex() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("name_regex"))
}

// Names returns a reference to field names of aws_iam_roles.
func (ir dataIamRolesAttributes) Names() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ir.ref.Append("names"))
}

// PathPrefix returns a reference to field path_prefix of aws_iam_roles.
func (ir dataIamRolesAttributes) PathPrefix() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("path_prefix"))
}
