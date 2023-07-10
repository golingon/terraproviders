// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computebackendserviceiambinding "github.com/golingon/terraproviders/googlebeta/4.72.1/computebackendserviceiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeBackendServiceIamBinding creates a new instance of [ComputeBackendServiceIamBinding].
func NewComputeBackendServiceIamBinding(name string, args ComputeBackendServiceIamBindingArgs) *ComputeBackendServiceIamBinding {
	return &ComputeBackendServiceIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeBackendServiceIamBinding)(nil)

// ComputeBackendServiceIamBinding represents the Terraform resource google_compute_backend_service_iam_binding.
type ComputeBackendServiceIamBinding struct {
	Name      string
	Args      ComputeBackendServiceIamBindingArgs
	state     *computeBackendServiceIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeBackendServiceIamBinding].
func (cbsib *ComputeBackendServiceIamBinding) Type() string {
	return "google_compute_backend_service_iam_binding"
}

// LocalName returns the local name for [ComputeBackendServiceIamBinding].
func (cbsib *ComputeBackendServiceIamBinding) LocalName() string {
	return cbsib.Name
}

// Configuration returns the configuration (args) for [ComputeBackendServiceIamBinding].
func (cbsib *ComputeBackendServiceIamBinding) Configuration() interface{} {
	return cbsib.Args
}

// DependOn is used for other resources to depend on [ComputeBackendServiceIamBinding].
func (cbsib *ComputeBackendServiceIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(cbsib)
}

// Dependencies returns the list of resources [ComputeBackendServiceIamBinding] depends_on.
func (cbsib *ComputeBackendServiceIamBinding) Dependencies() terra.Dependencies {
	return cbsib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeBackendServiceIamBinding].
func (cbsib *ComputeBackendServiceIamBinding) LifecycleManagement() *terra.Lifecycle {
	return cbsib.Lifecycle
}

// Attributes returns the attributes for [ComputeBackendServiceIamBinding].
func (cbsib *ComputeBackendServiceIamBinding) Attributes() computeBackendServiceIamBindingAttributes {
	return computeBackendServiceIamBindingAttributes{ref: terra.ReferenceResource(cbsib)}
}

// ImportState imports the given attribute values into [ComputeBackendServiceIamBinding]'s state.
func (cbsib *ComputeBackendServiceIamBinding) ImportState(av io.Reader) error {
	cbsib.state = &computeBackendServiceIamBindingState{}
	if err := json.NewDecoder(av).Decode(cbsib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cbsib.Type(), cbsib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeBackendServiceIamBinding] has state.
func (cbsib *ComputeBackendServiceIamBinding) State() (*computeBackendServiceIamBindingState, bool) {
	return cbsib.state, cbsib.state != nil
}

// StateMust returns the state for [ComputeBackendServiceIamBinding]. Panics if the state is nil.
func (cbsib *ComputeBackendServiceIamBinding) StateMust() *computeBackendServiceIamBindingState {
	if cbsib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cbsib.Type(), cbsib.LocalName()))
	}
	return cbsib.state
}

// ComputeBackendServiceIamBindingArgs contains the configurations for google_compute_backend_service_iam_binding.
type ComputeBackendServiceIamBindingArgs struct {
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
	Condition *computebackendserviceiambinding.Condition `hcl:"condition,block"`
}
type computeBackendServiceIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_compute_backend_service_iam_binding.
func (cbsib computeBackendServiceIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(cbsib.ref.Append("etag"))
}

// Id returns a reference to field id of google_compute_backend_service_iam_binding.
func (cbsib computeBackendServiceIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cbsib.ref.Append("id"))
}

// Members returns a reference to field members of google_compute_backend_service_iam_binding.
func (cbsib computeBackendServiceIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cbsib.ref.Append("members"))
}

// Name returns a reference to field name of google_compute_backend_service_iam_binding.
func (cbsib computeBackendServiceIamBindingAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cbsib.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_backend_service_iam_binding.
func (cbsib computeBackendServiceIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cbsib.ref.Append("project"))
}

// Role returns a reference to field role of google_compute_backend_service_iam_binding.
func (cbsib computeBackendServiceIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(cbsib.ref.Append("role"))
}

func (cbsib computeBackendServiceIamBindingAttributes) Condition() terra.ListValue[computebackendserviceiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[computebackendserviceiambinding.ConditionAttributes](cbsib.ref.Append("condition"))
}

type computeBackendServiceIamBindingState struct {
	Etag      string                                           `json:"etag"`
	Id        string                                           `json:"id"`
	Members   []string                                         `json:"members"`
	Name      string                                           `json:"name"`
	Project   string                                           `json:"project"`
	Role      string                                           `json:"role"`
	Condition []computebackendserviceiambinding.ConditionState `json:"condition"`
}
