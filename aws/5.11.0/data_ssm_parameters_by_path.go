// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataSsmParametersByPath creates a new instance of [DataSsmParametersByPath].
func NewDataSsmParametersByPath(name string, args DataSsmParametersByPathArgs) *DataSsmParametersByPath {
	return &DataSsmParametersByPath{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSsmParametersByPath)(nil)

// DataSsmParametersByPath represents the Terraform data resource aws_ssm_parameters_by_path.
type DataSsmParametersByPath struct {
	Name string
	Args DataSsmParametersByPathArgs
}

// DataSource returns the Terraform object type for [DataSsmParametersByPath].
func (spbp *DataSsmParametersByPath) DataSource() string {
	return "aws_ssm_parameters_by_path"
}

// LocalName returns the local name for [DataSsmParametersByPath].
func (spbp *DataSsmParametersByPath) LocalName() string {
	return spbp.Name
}

// Configuration returns the configuration (args) for [DataSsmParametersByPath].
func (spbp *DataSsmParametersByPath) Configuration() interface{} {
	return spbp.Args
}

// Attributes returns the attributes for [DataSsmParametersByPath].
func (spbp *DataSsmParametersByPath) Attributes() dataSsmParametersByPathAttributes {
	return dataSsmParametersByPathAttributes{ref: terra.ReferenceDataResource(spbp)}
}

// DataSsmParametersByPathArgs contains the configurations for aws_ssm_parameters_by_path.
type DataSsmParametersByPathArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Path: string, required
	Path terra.StringValue `hcl:"path,attr" validate:"required"`
	// Recursive: bool, optional
	Recursive terra.BoolValue `hcl:"recursive,attr"`
	// WithDecryption: bool, optional
	WithDecryption terra.BoolValue `hcl:"with_decryption,attr"`
}
type dataSsmParametersByPathAttributes struct {
	ref terra.Reference
}

// Arns returns a reference to field arns of aws_ssm_parameters_by_path.
func (spbp dataSsmParametersByPathAttributes) Arns() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](spbp.ref.Append("arns"))
}

// Id returns a reference to field id of aws_ssm_parameters_by_path.
func (spbp dataSsmParametersByPathAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(spbp.ref.Append("id"))
}

// Names returns a reference to field names of aws_ssm_parameters_by_path.
func (spbp dataSsmParametersByPathAttributes) Names() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](spbp.ref.Append("names"))
}

// Path returns a reference to field path of aws_ssm_parameters_by_path.
func (spbp dataSsmParametersByPathAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(spbp.ref.Append("path"))
}

// Recursive returns a reference to field recursive of aws_ssm_parameters_by_path.
func (spbp dataSsmParametersByPathAttributes) Recursive() terra.BoolValue {
	return terra.ReferenceAsBool(spbp.ref.Append("recursive"))
}

// Types returns a reference to field types of aws_ssm_parameters_by_path.
func (spbp dataSsmParametersByPathAttributes) Types() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](spbp.ref.Append("types"))
}

// Values returns a reference to field values of aws_ssm_parameters_by_path.
func (spbp dataSsmParametersByPathAttributes) Values() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](spbp.ref.Append("values"))
}

// WithDecryption returns a reference to field with_decryption of aws_ssm_parameters_by_path.
func (spbp dataSsmParametersByPathAttributes) WithDecryption() terra.BoolValue {
	return terra.ReferenceAsBool(spbp.ref.Append("with_decryption"))
}
