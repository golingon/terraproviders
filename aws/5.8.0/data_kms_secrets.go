// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datakmssecrets "github.com/golingon/terraproviders/aws/5.8.0/datakmssecrets"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataKmsSecrets creates a new instance of [DataKmsSecrets].
func NewDataKmsSecrets(name string, args DataKmsSecretsArgs) *DataKmsSecrets {
	return &DataKmsSecrets{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataKmsSecrets)(nil)

// DataKmsSecrets represents the Terraform data resource aws_kms_secrets.
type DataKmsSecrets struct {
	Name string
	Args DataKmsSecretsArgs
}

// DataSource returns the Terraform object type for [DataKmsSecrets].
func (ks *DataKmsSecrets) DataSource() string {
	return "aws_kms_secrets"
}

// LocalName returns the local name for [DataKmsSecrets].
func (ks *DataKmsSecrets) LocalName() string {
	return ks.Name
}

// Configuration returns the configuration (args) for [DataKmsSecrets].
func (ks *DataKmsSecrets) Configuration() interface{} {
	return ks.Args
}

// Attributes returns the attributes for [DataKmsSecrets].
func (ks *DataKmsSecrets) Attributes() dataKmsSecretsAttributes {
	return dataKmsSecretsAttributes{ref: terra.ReferenceDataResource(ks)}
}

// DataKmsSecretsArgs contains the configurations for aws_kms_secrets.
type DataKmsSecretsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Secret: min=1
	Secret []datakmssecrets.Secret `hcl:"secret,block" validate:"min=1"`
}
type dataKmsSecretsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_kms_secrets.
func (ks dataKmsSecretsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("id"))
}

// Plaintext returns a reference to field plaintext of aws_kms_secrets.
func (ks dataKmsSecretsAttributes) Plaintext() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ks.ref.Append("plaintext"))
}

func (ks dataKmsSecretsAttributes) Secret() terra.SetValue[datakmssecrets.SecretAttributes] {
	return terra.ReferenceAsSet[datakmssecrets.SecretAttributes](ks.ref.Append("secret"))
}
