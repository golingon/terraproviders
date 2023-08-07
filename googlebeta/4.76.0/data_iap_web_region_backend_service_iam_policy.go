// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataIapWebRegionBackendServiceIamPolicy creates a new instance of [DataIapWebRegionBackendServiceIamPolicy].
func NewDataIapWebRegionBackendServiceIamPolicy(name string, args DataIapWebRegionBackendServiceIamPolicyArgs) *DataIapWebRegionBackendServiceIamPolicy {
	return &DataIapWebRegionBackendServiceIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataIapWebRegionBackendServiceIamPolicy)(nil)

// DataIapWebRegionBackendServiceIamPolicy represents the Terraform data resource google_iap_web_region_backend_service_iam_policy.
type DataIapWebRegionBackendServiceIamPolicy struct {
	Name string
	Args DataIapWebRegionBackendServiceIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataIapWebRegionBackendServiceIamPolicy].
func (iwrbsip *DataIapWebRegionBackendServiceIamPolicy) DataSource() string {
	return "google_iap_web_region_backend_service_iam_policy"
}

// LocalName returns the local name for [DataIapWebRegionBackendServiceIamPolicy].
func (iwrbsip *DataIapWebRegionBackendServiceIamPolicy) LocalName() string {
	return iwrbsip.Name
}

// Configuration returns the configuration (args) for [DataIapWebRegionBackendServiceIamPolicy].
func (iwrbsip *DataIapWebRegionBackendServiceIamPolicy) Configuration() interface{} {
	return iwrbsip.Args
}

// Attributes returns the attributes for [DataIapWebRegionBackendServiceIamPolicy].
func (iwrbsip *DataIapWebRegionBackendServiceIamPolicy) Attributes() dataIapWebRegionBackendServiceIamPolicyAttributes {
	return dataIapWebRegionBackendServiceIamPolicyAttributes{ref: terra.ReferenceDataResource(iwrbsip)}
}

// DataIapWebRegionBackendServiceIamPolicyArgs contains the configurations for google_iap_web_region_backend_service_iam_policy.
type DataIapWebRegionBackendServiceIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// WebRegionBackendService: string, required
	WebRegionBackendService terra.StringValue `hcl:"web_region_backend_service,attr" validate:"required"`
}
type dataIapWebRegionBackendServiceIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_iap_web_region_backend_service_iam_policy.
func (iwrbsip dataIapWebRegionBackendServiceIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(iwrbsip.ref.Append("etag"))
}

// Id returns a reference to field id of google_iap_web_region_backend_service_iam_policy.
func (iwrbsip dataIapWebRegionBackendServiceIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iwrbsip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_iap_web_region_backend_service_iam_policy.
func (iwrbsip dataIapWebRegionBackendServiceIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(iwrbsip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_iap_web_region_backend_service_iam_policy.
func (iwrbsip dataIapWebRegionBackendServiceIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(iwrbsip.ref.Append("project"))
}

// Region returns a reference to field region of google_iap_web_region_backend_service_iam_policy.
func (iwrbsip dataIapWebRegionBackendServiceIamPolicyAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(iwrbsip.ref.Append("region"))
}

// WebRegionBackendService returns a reference to field web_region_backend_service of google_iap_web_region_backend_service_iam_policy.
func (iwrbsip dataIapWebRegionBackendServiceIamPolicyAttributes) WebRegionBackendService() terra.StringValue {
	return terra.ReferenceAsString(iwrbsip.ref.Append("web_region_backend_service"))
}
