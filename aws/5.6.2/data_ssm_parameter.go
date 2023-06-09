// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataSsmParameter creates a new instance of [DataSsmParameter].
func NewDataSsmParameter(name string, args DataSsmParameterArgs) *DataSsmParameter {
	return &DataSsmParameter{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSsmParameter)(nil)

// DataSsmParameter represents the Terraform data resource aws_ssm_parameter.
type DataSsmParameter struct {
	Name string
	Args DataSsmParameterArgs
}

// DataSource returns the Terraform object type for [DataSsmParameter].
func (sp *DataSsmParameter) DataSource() string {
	return "aws_ssm_parameter"
}

// LocalName returns the local name for [DataSsmParameter].
func (sp *DataSsmParameter) LocalName() string {
	return sp.Name
}

// Configuration returns the configuration (args) for [DataSsmParameter].
func (sp *DataSsmParameter) Configuration() interface{} {
	return sp.Args
}

// Attributes returns the attributes for [DataSsmParameter].
func (sp *DataSsmParameter) Attributes() dataSsmParameterAttributes {
	return dataSsmParameterAttributes{ref: terra.ReferenceDataResource(sp)}
}

// DataSsmParameterArgs contains the configurations for aws_ssm_parameter.
type DataSsmParameterArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// WithDecryption: bool, optional
	WithDecryption terra.BoolValue `hcl:"with_decryption,attr"`
}
type dataSsmParameterAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ssm_parameter.
func (sp dataSsmParameterAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("arn"))
}

// Id returns a reference to field id of aws_ssm_parameter.
func (sp dataSsmParameterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("id"))
}

// Name returns a reference to field name of aws_ssm_parameter.
func (sp dataSsmParameterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("name"))
}

// Type returns a reference to field type of aws_ssm_parameter.
func (sp dataSsmParameterAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("type"))
}

// Value returns a reference to field value of aws_ssm_parameter.
func (sp dataSsmParameterAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("value"))
}

// Version returns a reference to field version of aws_ssm_parameter.
func (sp dataSsmParameterAttributes) Version() terra.NumberValue {
	return terra.ReferenceAsNumber(sp.ref.Append("version"))
}

// WithDecryption returns a reference to field with_decryption of aws_ssm_parameter.
func (sp dataSsmParameterAttributes) WithDecryption() terra.BoolValue {
	return terra.ReferenceAsBool(sp.ref.Append("with_decryption"))
}
