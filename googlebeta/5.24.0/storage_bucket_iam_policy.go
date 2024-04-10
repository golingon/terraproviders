// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewStorageBucketIamPolicy creates a new instance of [StorageBucketIamPolicy].
func NewStorageBucketIamPolicy(name string, args StorageBucketIamPolicyArgs) *StorageBucketIamPolicy {
	return &StorageBucketIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StorageBucketIamPolicy)(nil)

// StorageBucketIamPolicy represents the Terraform resource google_storage_bucket_iam_policy.
type StorageBucketIamPolicy struct {
	Name      string
	Args      StorageBucketIamPolicyArgs
	state     *storageBucketIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StorageBucketIamPolicy].
func (sbip *StorageBucketIamPolicy) Type() string {
	return "google_storage_bucket_iam_policy"
}

// LocalName returns the local name for [StorageBucketIamPolicy].
func (sbip *StorageBucketIamPolicy) LocalName() string {
	return sbip.Name
}

// Configuration returns the configuration (args) for [StorageBucketIamPolicy].
func (sbip *StorageBucketIamPolicy) Configuration() interface{} {
	return sbip.Args
}

// DependOn is used for other resources to depend on [StorageBucketIamPolicy].
func (sbip *StorageBucketIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(sbip)
}

// Dependencies returns the list of resources [StorageBucketIamPolicy] depends_on.
func (sbip *StorageBucketIamPolicy) Dependencies() terra.Dependencies {
	return sbip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StorageBucketIamPolicy].
func (sbip *StorageBucketIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return sbip.Lifecycle
}

// Attributes returns the attributes for [StorageBucketIamPolicy].
func (sbip *StorageBucketIamPolicy) Attributes() storageBucketIamPolicyAttributes {
	return storageBucketIamPolicyAttributes{ref: terra.ReferenceResource(sbip)}
}

// ImportState imports the given attribute values into [StorageBucketIamPolicy]'s state.
func (sbip *StorageBucketIamPolicy) ImportState(av io.Reader) error {
	sbip.state = &storageBucketIamPolicyState{}
	if err := json.NewDecoder(av).Decode(sbip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sbip.Type(), sbip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StorageBucketIamPolicy] has state.
func (sbip *StorageBucketIamPolicy) State() (*storageBucketIamPolicyState, bool) {
	return sbip.state, sbip.state != nil
}

// StateMust returns the state for [StorageBucketIamPolicy]. Panics if the state is nil.
func (sbip *StorageBucketIamPolicy) StateMust() *storageBucketIamPolicyState {
	if sbip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sbip.Type(), sbip.LocalName()))
	}
	return sbip.state
}

// StorageBucketIamPolicyArgs contains the configurations for google_storage_bucket_iam_policy.
type StorageBucketIamPolicyArgs struct {
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
}
type storageBucketIamPolicyAttributes struct {
	ref terra.Reference
}

// Bucket returns a reference to field bucket of google_storage_bucket_iam_policy.
func (sbip storageBucketIamPolicyAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(sbip.ref.Append("bucket"))
}

// Etag returns a reference to field etag of google_storage_bucket_iam_policy.
func (sbip storageBucketIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(sbip.ref.Append("etag"))
}

// Id returns a reference to field id of google_storage_bucket_iam_policy.
func (sbip storageBucketIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sbip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_storage_bucket_iam_policy.
func (sbip storageBucketIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(sbip.ref.Append("policy_data"))
}

type storageBucketIamPolicyState struct {
	Bucket     string `json:"bucket"`
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	PolicyData string `json:"policy_data"`
}
