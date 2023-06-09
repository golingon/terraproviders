// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataComputeRegionBackendServiceIamPolicy creates a new instance of [DataComputeRegionBackendServiceIamPolicy].
func NewDataComputeRegionBackendServiceIamPolicy(name string, args DataComputeRegionBackendServiceIamPolicyArgs) *DataComputeRegionBackendServiceIamPolicy {
	return &DataComputeRegionBackendServiceIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataComputeRegionBackendServiceIamPolicy)(nil)

// DataComputeRegionBackendServiceIamPolicy represents the Terraform data resource google_compute_region_backend_service_iam_policy.
type DataComputeRegionBackendServiceIamPolicy struct {
	Name string
	Args DataComputeRegionBackendServiceIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataComputeRegionBackendServiceIamPolicy].
func (crbsip *DataComputeRegionBackendServiceIamPolicy) DataSource() string {
	return "google_compute_region_backend_service_iam_policy"
}

// LocalName returns the local name for [DataComputeRegionBackendServiceIamPolicy].
func (crbsip *DataComputeRegionBackendServiceIamPolicy) LocalName() string {
	return crbsip.Name
}

// Configuration returns the configuration (args) for [DataComputeRegionBackendServiceIamPolicy].
func (crbsip *DataComputeRegionBackendServiceIamPolicy) Configuration() interface{} {
	return crbsip.Args
}

// Attributes returns the attributes for [DataComputeRegionBackendServiceIamPolicy].
func (crbsip *DataComputeRegionBackendServiceIamPolicy) Attributes() dataComputeRegionBackendServiceIamPolicyAttributes {
	return dataComputeRegionBackendServiceIamPolicyAttributes{ref: terra.ReferenceDataResource(crbsip)}
}

// DataComputeRegionBackendServiceIamPolicyArgs contains the configurations for google_compute_region_backend_service_iam_policy.
type DataComputeRegionBackendServiceIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
}
type dataComputeRegionBackendServiceIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_compute_region_backend_service_iam_policy.
func (crbsip dataComputeRegionBackendServiceIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(crbsip.ref.Append("etag"))
}

// Id returns a reference to field id of google_compute_region_backend_service_iam_policy.
func (crbsip dataComputeRegionBackendServiceIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crbsip.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_region_backend_service_iam_policy.
func (crbsip dataComputeRegionBackendServiceIamPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crbsip.ref.Append("name"))
}

// PolicyData returns a reference to field policy_data of google_compute_region_backend_service_iam_policy.
func (crbsip dataComputeRegionBackendServiceIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(crbsip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_compute_region_backend_service_iam_policy.
func (crbsip dataComputeRegionBackendServiceIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crbsip.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_region_backend_service_iam_policy.
func (crbsip dataComputeRegionBackendServiceIamPolicyAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(crbsip.ref.Append("region"))
}
