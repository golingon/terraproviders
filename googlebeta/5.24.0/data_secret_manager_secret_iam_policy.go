// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import "github.com/golingon/lingon/pkg/terra"

// NewDataSecretManagerSecretIamPolicy creates a new instance of [DataSecretManagerSecretIamPolicy].
func NewDataSecretManagerSecretIamPolicy(name string, args DataSecretManagerSecretIamPolicyArgs) *DataSecretManagerSecretIamPolicy {
	return &DataSecretManagerSecretIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSecretManagerSecretIamPolicy)(nil)

// DataSecretManagerSecretIamPolicy represents the Terraform data resource google_secret_manager_secret_iam_policy.
type DataSecretManagerSecretIamPolicy struct {
	Name string
	Args DataSecretManagerSecretIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataSecretManagerSecretIamPolicy].
func (smsip *DataSecretManagerSecretIamPolicy) DataSource() string {
	return "google_secret_manager_secret_iam_policy"
}

// LocalName returns the local name for [DataSecretManagerSecretIamPolicy].
func (smsip *DataSecretManagerSecretIamPolicy) LocalName() string {
	return smsip.Name
}

// Configuration returns the configuration (args) for [DataSecretManagerSecretIamPolicy].
func (smsip *DataSecretManagerSecretIamPolicy) Configuration() interface{} {
	return smsip.Args
}

// Attributes returns the attributes for [DataSecretManagerSecretIamPolicy].
func (smsip *DataSecretManagerSecretIamPolicy) Attributes() dataSecretManagerSecretIamPolicyAttributes {
	return dataSecretManagerSecretIamPolicyAttributes{ref: terra.ReferenceDataResource(smsip)}
}

// DataSecretManagerSecretIamPolicyArgs contains the configurations for google_secret_manager_secret_iam_policy.
type DataSecretManagerSecretIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// SecretId: string, required
	SecretId terra.StringValue `hcl:"secret_id,attr" validate:"required"`
}
type dataSecretManagerSecretIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_secret_manager_secret_iam_policy.
func (smsip dataSecretManagerSecretIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(smsip.ref.Append("etag"))
}

// Id returns a reference to field id of google_secret_manager_secret_iam_policy.
func (smsip dataSecretManagerSecretIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(smsip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_secret_manager_secret_iam_policy.
func (smsip dataSecretManagerSecretIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(smsip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_secret_manager_secret_iam_policy.
func (smsip dataSecretManagerSecretIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(smsip.ref.Append("project"))
}

// SecretId returns a reference to field secret_id of google_secret_manager_secret_iam_policy.
func (smsip dataSecretManagerSecretIamPolicyAttributes) SecretId() terra.StringValue {
	return terra.ReferenceAsString(smsip.ref.Append("secret_id"))
}
