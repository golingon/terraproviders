// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import "github.com/golingon/lingon/pkg/terra"

// NewDataKmsCryptoKeyIamPolicy creates a new instance of [DataKmsCryptoKeyIamPolicy].
func NewDataKmsCryptoKeyIamPolicy(name string, args DataKmsCryptoKeyIamPolicyArgs) *DataKmsCryptoKeyIamPolicy {
	return &DataKmsCryptoKeyIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataKmsCryptoKeyIamPolicy)(nil)

// DataKmsCryptoKeyIamPolicy represents the Terraform data resource google_kms_crypto_key_iam_policy.
type DataKmsCryptoKeyIamPolicy struct {
	Name string
	Args DataKmsCryptoKeyIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataKmsCryptoKeyIamPolicy].
func (kckip *DataKmsCryptoKeyIamPolicy) DataSource() string {
	return "google_kms_crypto_key_iam_policy"
}

// LocalName returns the local name for [DataKmsCryptoKeyIamPolicy].
func (kckip *DataKmsCryptoKeyIamPolicy) LocalName() string {
	return kckip.Name
}

// Configuration returns the configuration (args) for [DataKmsCryptoKeyIamPolicy].
func (kckip *DataKmsCryptoKeyIamPolicy) Configuration() interface{} {
	return kckip.Args
}

// Attributes returns the attributes for [DataKmsCryptoKeyIamPolicy].
func (kckip *DataKmsCryptoKeyIamPolicy) Attributes() dataKmsCryptoKeyIamPolicyAttributes {
	return dataKmsCryptoKeyIamPolicyAttributes{ref: terra.ReferenceDataResource(kckip)}
}

// DataKmsCryptoKeyIamPolicyArgs contains the configurations for google_kms_crypto_key_iam_policy.
type DataKmsCryptoKeyIamPolicyArgs struct {
	// CryptoKeyId: string, required
	CryptoKeyId terra.StringValue `hcl:"crypto_key_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type dataKmsCryptoKeyIamPolicyAttributes struct {
	ref terra.Reference
}

// CryptoKeyId returns a reference to field crypto_key_id of google_kms_crypto_key_iam_policy.
func (kckip dataKmsCryptoKeyIamPolicyAttributes) CryptoKeyId() terra.StringValue {
	return terra.ReferenceAsString(kckip.ref.Append("crypto_key_id"))
}

// Etag returns a reference to field etag of google_kms_crypto_key_iam_policy.
func (kckip dataKmsCryptoKeyIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(kckip.ref.Append("etag"))
}

// Id returns a reference to field id of google_kms_crypto_key_iam_policy.
func (kckip dataKmsCryptoKeyIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kckip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_kms_crypto_key_iam_policy.
func (kckip dataKmsCryptoKeyIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(kckip.ref.Append("policy_data"))
}
