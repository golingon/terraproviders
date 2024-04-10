// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import "github.com/golingon/lingon/pkg/terra"

// NewDataKmsKeyRingIamPolicy creates a new instance of [DataKmsKeyRingIamPolicy].
func NewDataKmsKeyRingIamPolicy(name string, args DataKmsKeyRingIamPolicyArgs) *DataKmsKeyRingIamPolicy {
	return &DataKmsKeyRingIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataKmsKeyRingIamPolicy)(nil)

// DataKmsKeyRingIamPolicy represents the Terraform data resource google_kms_key_ring_iam_policy.
type DataKmsKeyRingIamPolicy struct {
	Name string
	Args DataKmsKeyRingIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataKmsKeyRingIamPolicy].
func (kkrip *DataKmsKeyRingIamPolicy) DataSource() string {
	return "google_kms_key_ring_iam_policy"
}

// LocalName returns the local name for [DataKmsKeyRingIamPolicy].
func (kkrip *DataKmsKeyRingIamPolicy) LocalName() string {
	return kkrip.Name
}

// Configuration returns the configuration (args) for [DataKmsKeyRingIamPolicy].
func (kkrip *DataKmsKeyRingIamPolicy) Configuration() interface{} {
	return kkrip.Args
}

// Attributes returns the attributes for [DataKmsKeyRingIamPolicy].
func (kkrip *DataKmsKeyRingIamPolicy) Attributes() dataKmsKeyRingIamPolicyAttributes {
	return dataKmsKeyRingIamPolicyAttributes{ref: terra.ReferenceDataResource(kkrip)}
}

// DataKmsKeyRingIamPolicyArgs contains the configurations for google_kms_key_ring_iam_policy.
type DataKmsKeyRingIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyRingId: string, required
	KeyRingId terra.StringValue `hcl:"key_ring_id,attr" validate:"required"`
}
type dataKmsKeyRingIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_kms_key_ring_iam_policy.
func (kkrip dataKmsKeyRingIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(kkrip.ref.Append("etag"))
}

// Id returns a reference to field id of google_kms_key_ring_iam_policy.
func (kkrip dataKmsKeyRingIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kkrip.ref.Append("id"))
}

// KeyRingId returns a reference to field key_ring_id of google_kms_key_ring_iam_policy.
func (kkrip dataKmsKeyRingIamPolicyAttributes) KeyRingId() terra.StringValue {
	return terra.ReferenceAsString(kkrip.ref.Append("key_ring_id"))
}

// PolicyData returns a reference to field policy_data of google_kms_key_ring_iam_policy.
func (kkrip dataKmsKeyRingIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(kkrip.ref.Append("policy_data"))
}
