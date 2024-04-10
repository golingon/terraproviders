// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import "github.com/golingon/lingon/pkg/terra"

// NewDataComputeInstanceIamPolicy creates a new instance of [DataComputeInstanceIamPolicy].
func NewDataComputeInstanceIamPolicy(name string, args DataComputeInstanceIamPolicyArgs) *DataComputeInstanceIamPolicy {
	return &DataComputeInstanceIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataComputeInstanceIamPolicy)(nil)

// DataComputeInstanceIamPolicy represents the Terraform data resource google_compute_instance_iam_policy.
type DataComputeInstanceIamPolicy struct {
	Name string
	Args DataComputeInstanceIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataComputeInstanceIamPolicy].
func (ciip *DataComputeInstanceIamPolicy) DataSource() string {
	return "google_compute_instance_iam_policy"
}

// LocalName returns the local name for [DataComputeInstanceIamPolicy].
func (ciip *DataComputeInstanceIamPolicy) LocalName() string {
	return ciip.Name
}

// Configuration returns the configuration (args) for [DataComputeInstanceIamPolicy].
func (ciip *DataComputeInstanceIamPolicy) Configuration() interface{} {
	return ciip.Args
}

// Attributes returns the attributes for [DataComputeInstanceIamPolicy].
func (ciip *DataComputeInstanceIamPolicy) Attributes() dataComputeInstanceIamPolicyAttributes {
	return dataComputeInstanceIamPolicyAttributes{ref: terra.ReferenceDataResource(ciip)}
}

// DataComputeInstanceIamPolicyArgs contains the configurations for google_compute_instance_iam_policy.
type DataComputeInstanceIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceName: string, required
	InstanceName terra.StringValue `hcl:"instance_name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
}
type dataComputeInstanceIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_compute_instance_iam_policy.
func (ciip dataComputeInstanceIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ciip.ref.Append("etag"))
}

// Id returns a reference to field id of google_compute_instance_iam_policy.
func (ciip dataComputeInstanceIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ciip.ref.Append("id"))
}

// InstanceName returns a reference to field instance_name of google_compute_instance_iam_policy.
func (ciip dataComputeInstanceIamPolicyAttributes) InstanceName() terra.StringValue {
	return terra.ReferenceAsString(ciip.ref.Append("instance_name"))
}

// PolicyData returns a reference to field policy_data of google_compute_instance_iam_policy.
func (ciip dataComputeInstanceIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(ciip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_compute_instance_iam_policy.
func (ciip dataComputeInstanceIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ciip.ref.Append("project"))
}

// Zone returns a reference to field zone of google_compute_instance_iam_policy.
func (ciip dataComputeInstanceIamPolicyAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(ciip.ref.Append("zone"))
}
