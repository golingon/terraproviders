// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_connect_prompt

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_connect_prompt.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (acp *DataSource) DataSource() string {
	return "aws_connect_prompt"
}

// LocalName returns the local name for [DataSource].
func (acp *DataSource) LocalName() string {
	return acp.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (acp *DataSource) Configuration() interface{} {
	return acp.Args
}

// Attributes returns the attributes for [DataSource].
func (acp *DataSource) Attributes() dataAwsConnectPromptAttributes {
	return dataAwsConnectPromptAttributes{ref: terra.ReferenceDataSource(acp)}
}

// DataArgs contains the configurations for aws_connect_prompt.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceId: string, required
	InstanceId terra.StringValue `hcl:"instance_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}

type dataAwsConnectPromptAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_connect_prompt.
func (acp dataAwsConnectPromptAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(acp.ref.Append("arn"))
}

// Id returns a reference to field id of aws_connect_prompt.
func (acp dataAwsConnectPromptAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acp.ref.Append("id"))
}

// InstanceId returns a reference to field instance_id of aws_connect_prompt.
func (acp dataAwsConnectPromptAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(acp.ref.Append("instance_id"))
}

// Name returns a reference to field name of aws_connect_prompt.
func (acp dataAwsConnectPromptAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(acp.ref.Append("name"))
}

// PromptId returns a reference to field prompt_id of aws_connect_prompt.
func (acp dataAwsConnectPromptAttributes) PromptId() terra.StringValue {
	return terra.ReferenceAsString(acp.ref.Append("prompt_id"))
}
