// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datakmssecret "github.com/golingon/terraproviders/aws/5.10.0/datakmssecret"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataKmsSecret creates a new instance of [DataKmsSecret].
func NewDataKmsSecret(name string, args DataKmsSecretArgs) *DataKmsSecret {
	return &DataKmsSecret{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataKmsSecret)(nil)

// DataKmsSecret represents the Terraform data resource aws_kms_secret.
type DataKmsSecret struct {
	Name string
	Args DataKmsSecretArgs
}

// DataSource returns the Terraform object type for [DataKmsSecret].
func (ks *DataKmsSecret) DataSource() string {
	return "aws_kms_secret"
}

// LocalName returns the local name for [DataKmsSecret].
func (ks *DataKmsSecret) LocalName() string {
	return ks.Name
}

// Configuration returns the configuration (args) for [DataKmsSecret].
func (ks *DataKmsSecret) Configuration() interface{} {
	return ks.Args
}

// Attributes returns the attributes for [DataKmsSecret].
func (ks *DataKmsSecret) Attributes() dataKmsSecretAttributes {
	return dataKmsSecretAttributes{ref: terra.ReferenceDataResource(ks)}
}

// DataKmsSecretArgs contains the configurations for aws_kms_secret.
type DataKmsSecretArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Secret: min=1
	Secret []datakmssecret.Secret `hcl:"secret,block" validate:"min=1"`
}
type dataKmsSecretAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_kms_secret.
func (ks dataKmsSecretAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("id"))
}

func (ks dataKmsSecretAttributes) Secret() terra.SetValue[datakmssecret.SecretAttributes] {
	return terra.ReferenceAsSet[datakmssecret.SecretAttributes](ks.ref.Append("secret"))
}
