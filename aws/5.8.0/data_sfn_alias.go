// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datasfnalias "github.com/golingon/terraproviders/aws/5.8.0/datasfnalias"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSfnAlias creates a new instance of [DataSfnAlias].
func NewDataSfnAlias(name string, args DataSfnAliasArgs) *DataSfnAlias {
	return &DataSfnAlias{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSfnAlias)(nil)

// DataSfnAlias represents the Terraform data resource aws_sfn_alias.
type DataSfnAlias struct {
	Name string
	Args DataSfnAliasArgs
}

// DataSource returns the Terraform object type for [DataSfnAlias].
func (sa *DataSfnAlias) DataSource() string {
	return "aws_sfn_alias"
}

// LocalName returns the local name for [DataSfnAlias].
func (sa *DataSfnAlias) LocalName() string {
	return sa.Name
}

// Configuration returns the configuration (args) for [DataSfnAlias].
func (sa *DataSfnAlias) Configuration() interface{} {
	return sa.Args
}

// Attributes returns the attributes for [DataSfnAlias].
func (sa *DataSfnAlias) Attributes() dataSfnAliasAttributes {
	return dataSfnAliasAttributes{ref: terra.ReferenceDataResource(sa)}
}

// DataSfnAliasArgs contains the configurations for aws_sfn_alias.
type DataSfnAliasArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// StatemachineArn: string, required
	StatemachineArn terra.StringValue `hcl:"statemachine_arn,attr" validate:"required"`
	// RoutingConfiguration: min=0
	RoutingConfiguration []datasfnalias.RoutingConfiguration `hcl:"routing_configuration,block" validate:"min=0"`
}
type dataSfnAliasAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_sfn_alias.
func (sa dataSfnAliasAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("arn"))
}

// CreationDate returns a reference to field creation_date of aws_sfn_alias.
func (sa dataSfnAliasAttributes) CreationDate() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("creation_date"))
}

// Description returns a reference to field description of aws_sfn_alias.
func (sa dataSfnAliasAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("description"))
}

// Id returns a reference to field id of aws_sfn_alias.
func (sa dataSfnAliasAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("id"))
}

// Name returns a reference to field name of aws_sfn_alias.
func (sa dataSfnAliasAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("name"))
}

// StatemachineArn returns a reference to field statemachine_arn of aws_sfn_alias.
func (sa dataSfnAliasAttributes) StatemachineArn() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("statemachine_arn"))
}

func (sa dataSfnAliasAttributes) RoutingConfiguration() terra.ListValue[datasfnalias.RoutingConfigurationAttributes] {
	return terra.ReferenceAsList[datasfnalias.RoutingConfigurationAttributes](sa.ref.Append("routing_configuration"))
}
