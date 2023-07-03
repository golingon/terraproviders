// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataComputeBackendBucketIamPolicy creates a new instance of [DataComputeBackendBucketIamPolicy].
func NewDataComputeBackendBucketIamPolicy(name string, args DataComputeBackendBucketIamPolicyArgs) *DataComputeBackendBucketIamPolicy {
	return &DataComputeBackendBucketIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataComputeBackendBucketIamPolicy)(nil)

// DataComputeBackendBucketIamPolicy represents the Terraform data resource google_compute_backend_bucket_iam_policy.
type DataComputeBackendBucketIamPolicy struct {
	Name string
	Args DataComputeBackendBucketIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataComputeBackendBucketIamPolicy].
func (cbbip *DataComputeBackendBucketIamPolicy) DataSource() string {
	return "google_compute_backend_bucket_iam_policy"
}

// LocalName returns the local name for [DataComputeBackendBucketIamPolicy].
func (cbbip *DataComputeBackendBucketIamPolicy) LocalName() string {
	return cbbip.Name
}

// Configuration returns the configuration (args) for [DataComputeBackendBucketIamPolicy].
func (cbbip *DataComputeBackendBucketIamPolicy) Configuration() interface{} {
	return cbbip.Args
}

// Attributes returns the attributes for [DataComputeBackendBucketIamPolicy].
func (cbbip *DataComputeBackendBucketIamPolicy) Attributes() dataComputeBackendBucketIamPolicyAttributes {
	return dataComputeBackendBucketIamPolicyAttributes{ref: terra.ReferenceDataResource(cbbip)}
}

// DataComputeBackendBucketIamPolicyArgs contains the configurations for google_compute_backend_bucket_iam_policy.
type DataComputeBackendBucketIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type dataComputeBackendBucketIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_compute_backend_bucket_iam_policy.
func (cbbip dataComputeBackendBucketIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(cbbip.ref.Append("etag"))
}

// Id returns a reference to field id of google_compute_backend_bucket_iam_policy.
func (cbbip dataComputeBackendBucketIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cbbip.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_backend_bucket_iam_policy.
func (cbbip dataComputeBackendBucketIamPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cbbip.ref.Append("name"))
}

// PolicyData returns a reference to field policy_data of google_compute_backend_bucket_iam_policy.
func (cbbip dataComputeBackendBucketIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(cbbip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_compute_backend_bucket_iam_policy.
func (cbbip dataComputeBackendBucketIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cbbip.ref.Append("project"))
}