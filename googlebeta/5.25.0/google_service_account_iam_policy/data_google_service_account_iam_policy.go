// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_service_account_iam_policy

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource google_service_account_iam_policy.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (gsaip *DataSource) DataSource() string {
	return "google_service_account_iam_policy"
}

// LocalName returns the local name for [DataSource].
func (gsaip *DataSource) LocalName() string {
	return gsaip.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (gsaip *DataSource) Configuration() interface{} {
	return gsaip.Args
}

// Attributes returns the attributes for [DataSource].
func (gsaip *DataSource) Attributes() dataGoogleServiceAccountIamPolicyAttributes {
	return dataGoogleServiceAccountIamPolicyAttributes{ref: terra.ReferenceDataSource(gsaip)}
}

// DataArgs contains the configurations for google_service_account_iam_policy.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ServiceAccountId: string, required
	ServiceAccountId terra.StringValue `hcl:"service_account_id,attr" validate:"required"`
}

type dataGoogleServiceAccountIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_service_account_iam_policy.
func (gsaip dataGoogleServiceAccountIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gsaip.ref.Append("etag"))
}

// Id returns a reference to field id of google_service_account_iam_policy.
func (gsaip dataGoogleServiceAccountIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gsaip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_service_account_iam_policy.
func (gsaip dataGoogleServiceAccountIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(gsaip.ref.Append("policy_data"))
}

// ServiceAccountId returns a reference to field service_account_id of google_service_account_iam_policy.
func (gsaip dataGoogleServiceAccountIamPolicyAttributes) ServiceAccountId() terra.StringValue {
	return terra.ReferenceAsString(gsaip.ref.Append("service_account_id"))
}
