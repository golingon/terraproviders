// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataCloudformationStack creates a new instance of [DataCloudformationStack].
func NewDataCloudformationStack(name string, args DataCloudformationStackArgs) *DataCloudformationStack {
	return &DataCloudformationStack{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCloudformationStack)(nil)

// DataCloudformationStack represents the Terraform data resource aws_cloudformation_stack.
type DataCloudformationStack struct {
	Name string
	Args DataCloudformationStackArgs
}

// DataSource returns the Terraform object type for [DataCloudformationStack].
func (cs *DataCloudformationStack) DataSource() string {
	return "aws_cloudformation_stack"
}

// LocalName returns the local name for [DataCloudformationStack].
func (cs *DataCloudformationStack) LocalName() string {
	return cs.Name
}

// Configuration returns the configuration (args) for [DataCloudformationStack].
func (cs *DataCloudformationStack) Configuration() interface{} {
	return cs.Args
}

// Attributes returns the attributes for [DataCloudformationStack].
func (cs *DataCloudformationStack) Attributes() dataCloudformationStackAttributes {
	return dataCloudformationStackAttributes{ref: terra.ReferenceDataResource(cs)}
}

// DataCloudformationStackArgs contains the configurations for aws_cloudformation_stack.
type DataCloudformationStackArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}
type dataCloudformationStackAttributes struct {
	ref terra.Reference
}

// Capabilities returns a reference to field capabilities of aws_cloudformation_stack.
func (cs dataCloudformationStackAttributes) Capabilities() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cs.ref.Append("capabilities"))
}

// Description returns a reference to field description of aws_cloudformation_stack.
func (cs dataCloudformationStackAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("description"))
}

// DisableRollback returns a reference to field disable_rollback of aws_cloudformation_stack.
func (cs dataCloudformationStackAttributes) DisableRollback() terra.BoolValue {
	return terra.ReferenceAsBool(cs.ref.Append("disable_rollback"))
}

// IamRoleArn returns a reference to field iam_role_arn of aws_cloudformation_stack.
func (cs dataCloudformationStackAttributes) IamRoleArn() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("iam_role_arn"))
}

// Id returns a reference to field id of aws_cloudformation_stack.
func (cs dataCloudformationStackAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("id"))
}

// Name returns a reference to field name of aws_cloudformation_stack.
func (cs dataCloudformationStackAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("name"))
}

// NotificationArns returns a reference to field notification_arns of aws_cloudformation_stack.
func (cs dataCloudformationStackAttributes) NotificationArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cs.ref.Append("notification_arns"))
}

// Outputs returns a reference to field outputs of aws_cloudformation_stack.
func (cs dataCloudformationStackAttributes) Outputs() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cs.ref.Append("outputs"))
}

// Parameters returns a reference to field parameters of aws_cloudformation_stack.
func (cs dataCloudformationStackAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cs.ref.Append("parameters"))
}

// Tags returns a reference to field tags of aws_cloudformation_stack.
func (cs dataCloudformationStackAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cs.ref.Append("tags"))
}

// TemplateBody returns a reference to field template_body of aws_cloudformation_stack.
func (cs dataCloudformationStackAttributes) TemplateBody() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("template_body"))
}

// TimeoutInMinutes returns a reference to field timeout_in_minutes of aws_cloudformation_stack.
func (cs dataCloudformationStackAttributes) TimeoutInMinutes() terra.NumberValue {
	return terra.ReferenceAsNumber(cs.ref.Append("timeout_in_minutes"))
}