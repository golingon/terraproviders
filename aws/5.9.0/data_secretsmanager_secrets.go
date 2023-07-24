// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datasecretsmanagersecrets "github.com/golingon/terraproviders/aws/5.9.0/datasecretsmanagersecrets"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSecretsmanagerSecrets creates a new instance of [DataSecretsmanagerSecrets].
func NewDataSecretsmanagerSecrets(name string, args DataSecretsmanagerSecretsArgs) *DataSecretsmanagerSecrets {
	return &DataSecretsmanagerSecrets{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSecretsmanagerSecrets)(nil)

// DataSecretsmanagerSecrets represents the Terraform data resource aws_secretsmanager_secrets.
type DataSecretsmanagerSecrets struct {
	Name string
	Args DataSecretsmanagerSecretsArgs
}

// DataSource returns the Terraform object type for [DataSecretsmanagerSecrets].
func (ss *DataSecretsmanagerSecrets) DataSource() string {
	return "aws_secretsmanager_secrets"
}

// LocalName returns the local name for [DataSecretsmanagerSecrets].
func (ss *DataSecretsmanagerSecrets) LocalName() string {
	return ss.Name
}

// Configuration returns the configuration (args) for [DataSecretsmanagerSecrets].
func (ss *DataSecretsmanagerSecrets) Configuration() interface{} {
	return ss.Args
}

// Attributes returns the attributes for [DataSecretsmanagerSecrets].
func (ss *DataSecretsmanagerSecrets) Attributes() dataSecretsmanagerSecretsAttributes {
	return dataSecretsmanagerSecretsAttributes{ref: terra.ReferenceDataResource(ss)}
}

// DataSecretsmanagerSecretsArgs contains the configurations for aws_secretsmanager_secrets.
type DataSecretsmanagerSecretsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Filter: min=0
	Filter []datasecretsmanagersecrets.Filter `hcl:"filter,block" validate:"min=0"`
}
type dataSecretsmanagerSecretsAttributes struct {
	ref terra.Reference
}

// Arns returns a reference to field arns of aws_secretsmanager_secrets.
func (ss dataSecretsmanagerSecretsAttributes) Arns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ss.ref.Append("arns"))
}

// Id returns a reference to field id of aws_secretsmanager_secrets.
func (ss dataSecretsmanagerSecretsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("id"))
}

// Names returns a reference to field names of aws_secretsmanager_secrets.
func (ss dataSecretsmanagerSecretsAttributes) Names() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ss.ref.Append("names"))
}

func (ss dataSecretsmanagerSecretsAttributes) Filter() terra.SetValue[datasecretsmanagersecrets.FilterAttributes] {
	return terra.ReferenceAsSet[datasecretsmanagersecrets.FilterAttributes](ss.ref.Append("filter"))
}
