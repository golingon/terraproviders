// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataComputeSubnetworkIamPolicy creates a new instance of [DataComputeSubnetworkIamPolicy].
func NewDataComputeSubnetworkIamPolicy(name string, args DataComputeSubnetworkIamPolicyArgs) *DataComputeSubnetworkIamPolicy {
	return &DataComputeSubnetworkIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataComputeSubnetworkIamPolicy)(nil)

// DataComputeSubnetworkIamPolicy represents the Terraform data resource google_compute_subnetwork_iam_policy.
type DataComputeSubnetworkIamPolicy struct {
	Name string
	Args DataComputeSubnetworkIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataComputeSubnetworkIamPolicy].
func (csip *DataComputeSubnetworkIamPolicy) DataSource() string {
	return "google_compute_subnetwork_iam_policy"
}

// LocalName returns the local name for [DataComputeSubnetworkIamPolicy].
func (csip *DataComputeSubnetworkIamPolicy) LocalName() string {
	return csip.Name
}

// Configuration returns the configuration (args) for [DataComputeSubnetworkIamPolicy].
func (csip *DataComputeSubnetworkIamPolicy) Configuration() interface{} {
	return csip.Args
}

// Attributes returns the attributes for [DataComputeSubnetworkIamPolicy].
func (csip *DataComputeSubnetworkIamPolicy) Attributes() dataComputeSubnetworkIamPolicyAttributes {
	return dataComputeSubnetworkIamPolicyAttributes{ref: terra.ReferenceDataResource(csip)}
}

// DataComputeSubnetworkIamPolicyArgs contains the configurations for google_compute_subnetwork_iam_policy.
type DataComputeSubnetworkIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Subnetwork: string, required
	Subnetwork terra.StringValue `hcl:"subnetwork,attr" validate:"required"`
}
type dataComputeSubnetworkIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_compute_subnetwork_iam_policy.
func (csip dataComputeSubnetworkIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(csip.ref.Append("etag"))
}

// Id returns a reference to field id of google_compute_subnetwork_iam_policy.
func (csip dataComputeSubnetworkIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(csip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_compute_subnetwork_iam_policy.
func (csip dataComputeSubnetworkIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(csip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_compute_subnetwork_iam_policy.
func (csip dataComputeSubnetworkIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(csip.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_subnetwork_iam_policy.
func (csip dataComputeSubnetworkIamPolicyAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(csip.ref.Append("region"))
}

// Subnetwork returns a reference to field subnetwork of google_compute_subnetwork_iam_policy.
func (csip dataComputeSubnetworkIamPolicyAttributes) Subnetwork() terra.StringValue {
	return terra.ReferenceAsString(csip.ref.Append("subnetwork"))
}