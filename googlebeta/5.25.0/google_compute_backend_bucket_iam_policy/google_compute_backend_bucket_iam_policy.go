// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_compute_backend_bucket_iam_policy

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource google_compute_backend_bucket_iam_policy.
type Resource struct {
	Name      string
	Args      Args
	state     *googleComputeBackendBucketIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gcbbip *Resource) Type() string {
	return "google_compute_backend_bucket_iam_policy"
}

// LocalName returns the local name for [Resource].
func (gcbbip *Resource) LocalName() string {
	return gcbbip.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gcbbip *Resource) Configuration() interface{} {
	return gcbbip.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gcbbip *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gcbbip)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gcbbip *Resource) Dependencies() terra.Dependencies {
	return gcbbip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gcbbip *Resource) LifecycleManagement() *terra.Lifecycle {
	return gcbbip.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gcbbip *Resource) Attributes() googleComputeBackendBucketIamPolicyAttributes {
	return googleComputeBackendBucketIamPolicyAttributes{ref: terra.ReferenceResource(gcbbip)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gcbbip *Resource) ImportState(state io.Reader) error {
	gcbbip.state = &googleComputeBackendBucketIamPolicyState{}
	if err := json.NewDecoder(state).Decode(gcbbip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gcbbip.Type(), gcbbip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gcbbip *Resource) State() (*googleComputeBackendBucketIamPolicyState, bool) {
	return gcbbip.state, gcbbip.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gcbbip *Resource) StateMust() *googleComputeBackendBucketIamPolicyState {
	if gcbbip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gcbbip.Type(), gcbbip.LocalName()))
	}
	return gcbbip.state
}

// Args contains the configurations for google_compute_backend_bucket_iam_policy.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}

type googleComputeBackendBucketIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_compute_backend_bucket_iam_policy.
func (gcbbip googleComputeBackendBucketIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gcbbip.ref.Append("etag"))
}

// Id returns a reference to field id of google_compute_backend_bucket_iam_policy.
func (gcbbip googleComputeBackendBucketIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gcbbip.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_backend_bucket_iam_policy.
func (gcbbip googleComputeBackendBucketIamPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gcbbip.ref.Append("name"))
}

// PolicyData returns a reference to field policy_data of google_compute_backend_bucket_iam_policy.
func (gcbbip googleComputeBackendBucketIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(gcbbip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_compute_backend_bucket_iam_policy.
func (gcbbip googleComputeBackendBucketIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gcbbip.ref.Append("project"))
}

type googleComputeBackendBucketIamPolicyState struct {
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	Name       string `json:"name"`
	PolicyData string `json:"policy_data"`
	Project    string `json:"project"`
}
