// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import "github.com/golingon/lingon/pkg/terra"

// NewDataIapWebTypeComputeIamPolicy creates a new instance of [DataIapWebTypeComputeIamPolicy].
func NewDataIapWebTypeComputeIamPolicy(name string, args DataIapWebTypeComputeIamPolicyArgs) *DataIapWebTypeComputeIamPolicy {
	return &DataIapWebTypeComputeIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataIapWebTypeComputeIamPolicy)(nil)

// DataIapWebTypeComputeIamPolicy represents the Terraform data resource google_iap_web_type_compute_iam_policy.
type DataIapWebTypeComputeIamPolicy struct {
	Name string
	Args DataIapWebTypeComputeIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataIapWebTypeComputeIamPolicy].
func (iwtcip *DataIapWebTypeComputeIamPolicy) DataSource() string {
	return "google_iap_web_type_compute_iam_policy"
}

// LocalName returns the local name for [DataIapWebTypeComputeIamPolicy].
func (iwtcip *DataIapWebTypeComputeIamPolicy) LocalName() string {
	return iwtcip.Name
}

// Configuration returns the configuration (args) for [DataIapWebTypeComputeIamPolicy].
func (iwtcip *DataIapWebTypeComputeIamPolicy) Configuration() interface{} {
	return iwtcip.Args
}

// Attributes returns the attributes for [DataIapWebTypeComputeIamPolicy].
func (iwtcip *DataIapWebTypeComputeIamPolicy) Attributes() dataIapWebTypeComputeIamPolicyAttributes {
	return dataIapWebTypeComputeIamPolicyAttributes{ref: terra.ReferenceDataResource(iwtcip)}
}

// DataIapWebTypeComputeIamPolicyArgs contains the configurations for google_iap_web_type_compute_iam_policy.
type DataIapWebTypeComputeIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type dataIapWebTypeComputeIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_iap_web_type_compute_iam_policy.
func (iwtcip dataIapWebTypeComputeIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(iwtcip.ref.Append("etag"))
}

// Id returns a reference to field id of google_iap_web_type_compute_iam_policy.
func (iwtcip dataIapWebTypeComputeIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iwtcip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_iap_web_type_compute_iam_policy.
func (iwtcip dataIapWebTypeComputeIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(iwtcip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_iap_web_type_compute_iam_policy.
func (iwtcip dataIapWebTypeComputeIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(iwtcip.ref.Append("project"))
}
