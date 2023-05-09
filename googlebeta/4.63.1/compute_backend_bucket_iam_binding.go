// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computebackendbucketiambinding "github.com/golingon/terraproviders/googlebeta/4.63.1/computebackendbucketiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeBackendBucketIamBinding creates a new instance of [ComputeBackendBucketIamBinding].
func NewComputeBackendBucketIamBinding(name string, args ComputeBackendBucketIamBindingArgs) *ComputeBackendBucketIamBinding {
	return &ComputeBackendBucketIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeBackendBucketIamBinding)(nil)

// ComputeBackendBucketIamBinding represents the Terraform resource google_compute_backend_bucket_iam_binding.
type ComputeBackendBucketIamBinding struct {
	Name      string
	Args      ComputeBackendBucketIamBindingArgs
	state     *computeBackendBucketIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeBackendBucketIamBinding].
func (cbbib *ComputeBackendBucketIamBinding) Type() string {
	return "google_compute_backend_bucket_iam_binding"
}

// LocalName returns the local name for [ComputeBackendBucketIamBinding].
func (cbbib *ComputeBackendBucketIamBinding) LocalName() string {
	return cbbib.Name
}

// Configuration returns the configuration (args) for [ComputeBackendBucketIamBinding].
func (cbbib *ComputeBackendBucketIamBinding) Configuration() interface{} {
	return cbbib.Args
}

// DependOn is used for other resources to depend on [ComputeBackendBucketIamBinding].
func (cbbib *ComputeBackendBucketIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(cbbib)
}

// Dependencies returns the list of resources [ComputeBackendBucketIamBinding] depends_on.
func (cbbib *ComputeBackendBucketIamBinding) Dependencies() terra.Dependencies {
	return cbbib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeBackendBucketIamBinding].
func (cbbib *ComputeBackendBucketIamBinding) LifecycleManagement() *terra.Lifecycle {
	return cbbib.Lifecycle
}

// Attributes returns the attributes for [ComputeBackendBucketIamBinding].
func (cbbib *ComputeBackendBucketIamBinding) Attributes() computeBackendBucketIamBindingAttributes {
	return computeBackendBucketIamBindingAttributes{ref: terra.ReferenceResource(cbbib)}
}

// ImportState imports the given attribute values into [ComputeBackendBucketIamBinding]'s state.
func (cbbib *ComputeBackendBucketIamBinding) ImportState(av io.Reader) error {
	cbbib.state = &computeBackendBucketIamBindingState{}
	if err := json.NewDecoder(av).Decode(cbbib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cbbib.Type(), cbbib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeBackendBucketIamBinding] has state.
func (cbbib *ComputeBackendBucketIamBinding) State() (*computeBackendBucketIamBindingState, bool) {
	return cbbib.state, cbbib.state != nil
}

// StateMust returns the state for [ComputeBackendBucketIamBinding]. Panics if the state is nil.
func (cbbib *ComputeBackendBucketIamBinding) StateMust() *computeBackendBucketIamBindingState {
	if cbbib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cbbib.Type(), cbbib.LocalName()))
	}
	return cbbib.state
}

// ComputeBackendBucketIamBindingArgs contains the configurations for google_compute_backend_bucket_iam_binding.
type ComputeBackendBucketIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *computebackendbucketiambinding.Condition `hcl:"condition,block"`
}
type computeBackendBucketIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_compute_backend_bucket_iam_binding.
func (cbbib computeBackendBucketIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(cbbib.ref.Append("etag"))
}

// Id returns a reference to field id of google_compute_backend_bucket_iam_binding.
func (cbbib computeBackendBucketIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cbbib.ref.Append("id"))
}

// Members returns a reference to field members of google_compute_backend_bucket_iam_binding.
func (cbbib computeBackendBucketIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cbbib.ref.Append("members"))
}

// Name returns a reference to field name of google_compute_backend_bucket_iam_binding.
func (cbbib computeBackendBucketIamBindingAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cbbib.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_backend_bucket_iam_binding.
func (cbbib computeBackendBucketIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cbbib.ref.Append("project"))
}

// Role returns a reference to field role of google_compute_backend_bucket_iam_binding.
func (cbbib computeBackendBucketIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(cbbib.ref.Append("role"))
}

func (cbbib computeBackendBucketIamBindingAttributes) Condition() terra.ListValue[computebackendbucketiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[computebackendbucketiambinding.ConditionAttributes](cbbib.ref.Append("condition"))
}

type computeBackendBucketIamBindingState struct {
	Etag      string                                          `json:"etag"`
	Id        string                                          `json:"id"`
	Members   []string                                        `json:"members"`
	Name      string                                          `json:"name"`
	Project   string                                          `json:"project"`
	Role      string                                          `json:"role"`
	Condition []computebackendbucketiambinding.ConditionState `json:"condition"`
}
