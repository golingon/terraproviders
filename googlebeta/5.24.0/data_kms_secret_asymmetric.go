// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import "github.com/golingon/lingon/pkg/terra"

// NewDataKmsSecretAsymmetric creates a new instance of [DataKmsSecretAsymmetric].
func NewDataKmsSecretAsymmetric(name string, args DataKmsSecretAsymmetricArgs) *DataKmsSecretAsymmetric {
	return &DataKmsSecretAsymmetric{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataKmsSecretAsymmetric)(nil)

// DataKmsSecretAsymmetric represents the Terraform data resource google_kms_secret_asymmetric.
type DataKmsSecretAsymmetric struct {
	Name string
	Args DataKmsSecretAsymmetricArgs
}

// DataSource returns the Terraform object type for [DataKmsSecretAsymmetric].
func (ksa *DataKmsSecretAsymmetric) DataSource() string {
	return "google_kms_secret_asymmetric"
}

// LocalName returns the local name for [DataKmsSecretAsymmetric].
func (ksa *DataKmsSecretAsymmetric) LocalName() string {
	return ksa.Name
}

// Configuration returns the configuration (args) for [DataKmsSecretAsymmetric].
func (ksa *DataKmsSecretAsymmetric) Configuration() interface{} {
	return ksa.Args
}

// Attributes returns the attributes for [DataKmsSecretAsymmetric].
func (ksa *DataKmsSecretAsymmetric) Attributes() dataKmsSecretAsymmetricAttributes {
	return dataKmsSecretAsymmetricAttributes{ref: terra.ReferenceDataResource(ksa)}
}

// DataKmsSecretAsymmetricArgs contains the configurations for google_kms_secret_asymmetric.
type DataKmsSecretAsymmetricArgs struct {
	// Ciphertext: string, required
	Ciphertext terra.StringValue `hcl:"ciphertext,attr" validate:"required"`
	// Crc32: string, optional
	Crc32 terra.StringValue `hcl:"crc32,attr"`
	// CryptoKeyVersion: string, required
	CryptoKeyVersion terra.StringValue `hcl:"crypto_key_version,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type dataKmsSecretAsymmetricAttributes struct {
	ref terra.Reference
}

// Ciphertext returns a reference to field ciphertext of google_kms_secret_asymmetric.
func (ksa dataKmsSecretAsymmetricAttributes) Ciphertext() terra.StringValue {
	return terra.ReferenceAsString(ksa.ref.Append("ciphertext"))
}

// Crc32 returns a reference to field crc32 of google_kms_secret_asymmetric.
func (ksa dataKmsSecretAsymmetricAttributes) Crc32() terra.StringValue {
	return terra.ReferenceAsString(ksa.ref.Append("crc32"))
}

// CryptoKeyVersion returns a reference to field crypto_key_version of google_kms_secret_asymmetric.
func (ksa dataKmsSecretAsymmetricAttributes) CryptoKeyVersion() terra.StringValue {
	return terra.ReferenceAsString(ksa.ref.Append("crypto_key_version"))
}

// Id returns a reference to field id of google_kms_secret_asymmetric.
func (ksa dataKmsSecretAsymmetricAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ksa.ref.Append("id"))
}

// Plaintext returns a reference to field plaintext of google_kms_secret_asymmetric.
func (ksa dataKmsSecretAsymmetricAttributes) Plaintext() terra.StringValue {
	return terra.ReferenceAsString(ksa.ref.Append("plaintext"))
}
