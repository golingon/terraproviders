// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataIapWebBackendServiceIamPolicy creates a new instance of [DataIapWebBackendServiceIamPolicy].
func NewDataIapWebBackendServiceIamPolicy(name string, args DataIapWebBackendServiceIamPolicyArgs) *DataIapWebBackendServiceIamPolicy {
	return &DataIapWebBackendServiceIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataIapWebBackendServiceIamPolicy)(nil)

// DataIapWebBackendServiceIamPolicy represents the Terraform data resource google_iap_web_backend_service_iam_policy.
type DataIapWebBackendServiceIamPolicy struct {
	Name string
	Args DataIapWebBackendServiceIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataIapWebBackendServiceIamPolicy].
func (iwbsip *DataIapWebBackendServiceIamPolicy) DataSource() string {
	return "google_iap_web_backend_service_iam_policy"
}

// LocalName returns the local name for [DataIapWebBackendServiceIamPolicy].
func (iwbsip *DataIapWebBackendServiceIamPolicy) LocalName() string {
	return iwbsip.Name
}

// Configuration returns the configuration (args) for [DataIapWebBackendServiceIamPolicy].
func (iwbsip *DataIapWebBackendServiceIamPolicy) Configuration() interface{} {
	return iwbsip.Args
}

// Attributes returns the attributes for [DataIapWebBackendServiceIamPolicy].
func (iwbsip *DataIapWebBackendServiceIamPolicy) Attributes() dataIapWebBackendServiceIamPolicyAttributes {
	return dataIapWebBackendServiceIamPolicyAttributes{ref: terra.ReferenceDataResource(iwbsip)}
}

// DataIapWebBackendServiceIamPolicyArgs contains the configurations for google_iap_web_backend_service_iam_policy.
type DataIapWebBackendServiceIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// WebBackendService: string, required
	WebBackendService terra.StringValue `hcl:"web_backend_service,attr" validate:"required"`
}
type dataIapWebBackendServiceIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_iap_web_backend_service_iam_policy.
func (iwbsip dataIapWebBackendServiceIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(iwbsip.ref.Append("etag"))
}

// Id returns a reference to field id of google_iap_web_backend_service_iam_policy.
func (iwbsip dataIapWebBackendServiceIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iwbsip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_iap_web_backend_service_iam_policy.
func (iwbsip dataIapWebBackendServiceIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(iwbsip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_iap_web_backend_service_iam_policy.
func (iwbsip dataIapWebBackendServiceIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(iwbsip.ref.Append("project"))
}

// WebBackendService returns a reference to field web_backend_service of google_iap_web_backend_service_iam_policy.
func (iwbsip dataIapWebBackendServiceIamPolicyAttributes) WebBackendService() terra.StringValue {
	return terra.ReferenceAsString(iwbsip.ref.Append("web_backend_service"))
}
