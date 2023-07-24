// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataBinaryAuthorizationAttestorIamPolicy creates a new instance of [DataBinaryAuthorizationAttestorIamPolicy].
func NewDataBinaryAuthorizationAttestorIamPolicy(name string, args DataBinaryAuthorizationAttestorIamPolicyArgs) *DataBinaryAuthorizationAttestorIamPolicy {
	return &DataBinaryAuthorizationAttestorIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataBinaryAuthorizationAttestorIamPolicy)(nil)

// DataBinaryAuthorizationAttestorIamPolicy represents the Terraform data resource google_binary_authorization_attestor_iam_policy.
type DataBinaryAuthorizationAttestorIamPolicy struct {
	Name string
	Args DataBinaryAuthorizationAttestorIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataBinaryAuthorizationAttestorIamPolicy].
func (baaip *DataBinaryAuthorizationAttestorIamPolicy) DataSource() string {
	return "google_binary_authorization_attestor_iam_policy"
}

// LocalName returns the local name for [DataBinaryAuthorizationAttestorIamPolicy].
func (baaip *DataBinaryAuthorizationAttestorIamPolicy) LocalName() string {
	return baaip.Name
}

// Configuration returns the configuration (args) for [DataBinaryAuthorizationAttestorIamPolicy].
func (baaip *DataBinaryAuthorizationAttestorIamPolicy) Configuration() interface{} {
	return baaip.Args
}

// Attributes returns the attributes for [DataBinaryAuthorizationAttestorIamPolicy].
func (baaip *DataBinaryAuthorizationAttestorIamPolicy) Attributes() dataBinaryAuthorizationAttestorIamPolicyAttributes {
	return dataBinaryAuthorizationAttestorIamPolicyAttributes{ref: terra.ReferenceDataResource(baaip)}
}

// DataBinaryAuthorizationAttestorIamPolicyArgs contains the configurations for google_binary_authorization_attestor_iam_policy.
type DataBinaryAuthorizationAttestorIamPolicyArgs struct {
	// Attestor: string, required
	Attestor terra.StringValue `hcl:"attestor,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type dataBinaryAuthorizationAttestorIamPolicyAttributes struct {
	ref terra.Reference
}

// Attestor returns a reference to field attestor of google_binary_authorization_attestor_iam_policy.
func (baaip dataBinaryAuthorizationAttestorIamPolicyAttributes) Attestor() terra.StringValue {
	return terra.ReferenceAsString(baaip.ref.Append("attestor"))
}

// Etag returns a reference to field etag of google_binary_authorization_attestor_iam_policy.
func (baaip dataBinaryAuthorizationAttestorIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(baaip.ref.Append("etag"))
}

// Id returns a reference to field id of google_binary_authorization_attestor_iam_policy.
func (baaip dataBinaryAuthorizationAttestorIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(baaip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_binary_authorization_attestor_iam_policy.
func (baaip dataBinaryAuthorizationAttestorIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(baaip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_binary_authorization_attestor_iam_policy.
func (baaip dataBinaryAuthorizationAttestorIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(baaip.ref.Append("project"))
}
