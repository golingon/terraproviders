// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataStorageBucketIamPolicy creates a new instance of [DataStorageBucketIamPolicy].
func NewDataStorageBucketIamPolicy(name string, args DataStorageBucketIamPolicyArgs) *DataStorageBucketIamPolicy {
	return &DataStorageBucketIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataStorageBucketIamPolicy)(nil)

// DataStorageBucketIamPolicy represents the Terraform data resource google_storage_bucket_iam_policy.
type DataStorageBucketIamPolicy struct {
	Name string
	Args DataStorageBucketIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataStorageBucketIamPolicy].
func (sbip *DataStorageBucketIamPolicy) DataSource() string {
	return "google_storage_bucket_iam_policy"
}

// LocalName returns the local name for [DataStorageBucketIamPolicy].
func (sbip *DataStorageBucketIamPolicy) LocalName() string {
	return sbip.Name
}

// Configuration returns the configuration (args) for [DataStorageBucketIamPolicy].
func (sbip *DataStorageBucketIamPolicy) Configuration() interface{} {
	return sbip.Args
}

// Attributes returns the attributes for [DataStorageBucketIamPolicy].
func (sbip *DataStorageBucketIamPolicy) Attributes() dataStorageBucketIamPolicyAttributes {
	return dataStorageBucketIamPolicyAttributes{ref: terra.ReferenceDataResource(sbip)}
}

// DataStorageBucketIamPolicyArgs contains the configurations for google_storage_bucket_iam_policy.
type DataStorageBucketIamPolicyArgs struct {
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type dataStorageBucketIamPolicyAttributes struct {
	ref terra.Reference
}

// Bucket returns a reference to field bucket of google_storage_bucket_iam_policy.
func (sbip dataStorageBucketIamPolicyAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(sbip.ref.Append("bucket"))
}

// Etag returns a reference to field etag of google_storage_bucket_iam_policy.
func (sbip dataStorageBucketIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(sbip.ref.Append("etag"))
}

// Id returns a reference to field id of google_storage_bucket_iam_policy.
func (sbip dataStorageBucketIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sbip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_storage_bucket_iam_policy.
func (sbip dataStorageBucketIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(sbip.ref.Append("policy_data"))
}
