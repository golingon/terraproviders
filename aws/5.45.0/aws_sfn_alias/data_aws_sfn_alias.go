// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_sfn_alias

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_sfn_alias.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (asa *DataSource) DataSource() string {
	return "aws_sfn_alias"
}

// LocalName returns the local name for [DataSource].
func (asa *DataSource) LocalName() string {
	return asa.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (asa *DataSource) Configuration() interface{} {
	return asa.Args
}

// Attributes returns the attributes for [DataSource].
func (asa *DataSource) Attributes() dataAwsSfnAliasAttributes {
	return dataAwsSfnAliasAttributes{ref: terra.ReferenceDataSource(asa)}
}

// DataArgs contains the configurations for aws_sfn_alias.
type DataArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// StatemachineArn: string, required
	StatemachineArn terra.StringValue `hcl:"statemachine_arn,attr" validate:"required"`
}

type dataAwsSfnAliasAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_sfn_alias.
func (asa dataAwsSfnAliasAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(asa.ref.Append("arn"))
}

// CreationDate returns a reference to field creation_date of aws_sfn_alias.
func (asa dataAwsSfnAliasAttributes) CreationDate() terra.StringValue {
	return terra.ReferenceAsString(asa.ref.Append("creation_date"))
}

// Description returns a reference to field description of aws_sfn_alias.
func (asa dataAwsSfnAliasAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(asa.ref.Append("description"))
}

// Id returns a reference to field id of aws_sfn_alias.
func (asa dataAwsSfnAliasAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asa.ref.Append("id"))
}

// Name returns a reference to field name of aws_sfn_alias.
func (asa dataAwsSfnAliasAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(asa.ref.Append("name"))
}

// StatemachineArn returns a reference to field statemachine_arn of aws_sfn_alias.
func (asa dataAwsSfnAliasAttributes) StatemachineArn() terra.StringValue {
	return terra.ReferenceAsString(asa.ref.Append("statemachine_arn"))
}

func (asa dataAwsSfnAliasAttributes) RoutingConfiguration() terra.ListValue[DataRoutingConfigurationAttributes] {
	return terra.ReferenceAsList[DataRoutingConfigurationAttributes](asa.ref.Append("routing_configuration"))
}
