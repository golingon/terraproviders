// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataKmsSecret creates a new instance of [DataKmsSecret].
func NewDataKmsSecret(name string, args DataKmsSecretArgs) *DataKmsSecret {
	return &DataKmsSecret{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataKmsSecret)(nil)

// DataKmsSecret represents the Terraform data resource google_kms_secret.
type DataKmsSecret struct {
	Name string
	Args DataKmsSecretArgs
}

// DataSource returns the Terraform object type for [DataKmsSecret].
func (ks *DataKmsSecret) DataSource() string {
	return "google_kms_secret"
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

// DataKmsSecretArgs contains the configurations for google_kms_secret.
type DataKmsSecretArgs struct {
	// AdditionalAuthenticatedData: string, optional
	AdditionalAuthenticatedData terra.StringValue `hcl:"additional_authenticated_data,attr"`
	// Ciphertext: string, required
	Ciphertext terra.StringValue `hcl:"ciphertext,attr" validate:"required"`
	// CryptoKey: string, required
	CryptoKey terra.StringValue `hcl:"crypto_key,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type dataKmsSecretAttributes struct {
	ref terra.Reference
}

// AdditionalAuthenticatedData returns a reference to field additional_authenticated_data of google_kms_secret.
func (ks dataKmsSecretAttributes) AdditionalAuthenticatedData() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("additional_authenticated_data"))
}

// Ciphertext returns a reference to field ciphertext of google_kms_secret.
func (ks dataKmsSecretAttributes) Ciphertext() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("ciphertext"))
}

// CryptoKey returns a reference to field crypto_key of google_kms_secret.
func (ks dataKmsSecretAttributes) CryptoKey() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("crypto_key"))
}

// Id returns a reference to field id of google_kms_secret.
func (ks dataKmsSecretAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("id"))
}

// Plaintext returns a reference to field plaintext of google_kms_secret.
func (ks dataKmsSecretAttributes) Plaintext() terra.StringValue {
	return terra.ReferenceAsString(ks.ref.Append("plaintext"))
}