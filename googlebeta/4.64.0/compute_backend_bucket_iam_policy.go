// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeBackendBucketIamPolicy creates a new instance of [ComputeBackendBucketIamPolicy].
func NewComputeBackendBucketIamPolicy(name string, args ComputeBackendBucketIamPolicyArgs) *ComputeBackendBucketIamPolicy {
	return &ComputeBackendBucketIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeBackendBucketIamPolicy)(nil)

// ComputeBackendBucketIamPolicy represents the Terraform resource google_compute_backend_bucket_iam_policy.
type ComputeBackendBucketIamPolicy struct {
	Name      string
	Args      ComputeBackendBucketIamPolicyArgs
	state     *computeBackendBucketIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeBackendBucketIamPolicy].
func (cbbip *ComputeBackendBucketIamPolicy) Type() string {
	return "google_compute_backend_bucket_iam_policy"
}

// LocalName returns the local name for [ComputeBackendBucketIamPolicy].
func (cbbip *ComputeBackendBucketIamPolicy) LocalName() string {
	return cbbip.Name
}

// Configuration returns the configuration (args) for [ComputeBackendBucketIamPolicy].
func (cbbip *ComputeBackendBucketIamPolicy) Configuration() interface{} {
	return cbbip.Args
}

// DependOn is used for other resources to depend on [ComputeBackendBucketIamPolicy].
func (cbbip *ComputeBackendBucketIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(cbbip)
}

// Dependencies returns the list of resources [ComputeBackendBucketIamPolicy] depends_on.
func (cbbip *ComputeBackendBucketIamPolicy) Dependencies() terra.Dependencies {
	return cbbip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeBackendBucketIamPolicy].
func (cbbip *ComputeBackendBucketIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return cbbip.Lifecycle
}

// Attributes returns the attributes for [ComputeBackendBucketIamPolicy].
func (cbbip *ComputeBackendBucketIamPolicy) Attributes() computeBackendBucketIamPolicyAttributes {
	return computeBackendBucketIamPolicyAttributes{ref: terra.ReferenceResource(cbbip)}
}

// ImportState imports the given attribute values into [ComputeBackendBucketIamPolicy]'s state.
func (cbbip *ComputeBackendBucketIamPolicy) ImportState(av io.Reader) error {
	cbbip.state = &computeBackendBucketIamPolicyState{}
	if err := json.NewDecoder(av).Decode(cbbip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cbbip.Type(), cbbip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeBackendBucketIamPolicy] has state.
func (cbbip *ComputeBackendBucketIamPolicy) State() (*computeBackendBucketIamPolicyState, bool) {
	return cbbip.state, cbbip.state != nil
}

// StateMust returns the state for [ComputeBackendBucketIamPolicy]. Panics if the state is nil.
func (cbbip *ComputeBackendBucketIamPolicy) StateMust() *computeBackendBucketIamPolicyState {
	if cbbip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cbbip.Type(), cbbip.LocalName()))
	}
	return cbbip.state
}

// ComputeBackendBucketIamPolicyArgs contains the configurations for google_compute_backend_bucket_iam_policy.
type ComputeBackendBucketIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type computeBackendBucketIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_compute_backend_bucket_iam_policy.
func (cbbip computeBackendBucketIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(cbbip.ref.Append("etag"))
}

// Id returns a reference to field id of google_compute_backend_bucket_iam_policy.
func (cbbip computeBackendBucketIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cbbip.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_backend_bucket_iam_policy.
func (cbbip computeBackendBucketIamPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cbbip.ref.Append("name"))
}

// PolicyData returns a reference to field policy_data of google_compute_backend_bucket_iam_policy.
func (cbbip computeBackendBucketIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(cbbip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_compute_backend_bucket_iam_policy.
func (cbbip computeBackendBucketIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cbbip.ref.Append("project"))
}

type computeBackendBucketIamPolicyState struct {
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	Name       string `json:"name"`
	PolicyData string `json:"policy_data"`
	Project    string `json:"project"`
}
