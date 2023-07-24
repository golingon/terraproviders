// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computemachineimageiambinding "github.com/golingon/terraproviders/googlebeta/4.74.0/computemachineimageiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeMachineImageIamBinding creates a new instance of [ComputeMachineImageIamBinding].
func NewComputeMachineImageIamBinding(name string, args ComputeMachineImageIamBindingArgs) *ComputeMachineImageIamBinding {
	return &ComputeMachineImageIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeMachineImageIamBinding)(nil)

// ComputeMachineImageIamBinding represents the Terraform resource google_compute_machine_image_iam_binding.
type ComputeMachineImageIamBinding struct {
	Name      string
	Args      ComputeMachineImageIamBindingArgs
	state     *computeMachineImageIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeMachineImageIamBinding].
func (cmiib *ComputeMachineImageIamBinding) Type() string {
	return "google_compute_machine_image_iam_binding"
}

// LocalName returns the local name for [ComputeMachineImageIamBinding].
func (cmiib *ComputeMachineImageIamBinding) LocalName() string {
	return cmiib.Name
}

// Configuration returns the configuration (args) for [ComputeMachineImageIamBinding].
func (cmiib *ComputeMachineImageIamBinding) Configuration() interface{} {
	return cmiib.Args
}

// DependOn is used for other resources to depend on [ComputeMachineImageIamBinding].
func (cmiib *ComputeMachineImageIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(cmiib)
}

// Dependencies returns the list of resources [ComputeMachineImageIamBinding] depends_on.
func (cmiib *ComputeMachineImageIamBinding) Dependencies() terra.Dependencies {
	return cmiib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeMachineImageIamBinding].
func (cmiib *ComputeMachineImageIamBinding) LifecycleManagement() *terra.Lifecycle {
	return cmiib.Lifecycle
}

// Attributes returns the attributes for [ComputeMachineImageIamBinding].
func (cmiib *ComputeMachineImageIamBinding) Attributes() computeMachineImageIamBindingAttributes {
	return computeMachineImageIamBindingAttributes{ref: terra.ReferenceResource(cmiib)}
}

// ImportState imports the given attribute values into [ComputeMachineImageIamBinding]'s state.
func (cmiib *ComputeMachineImageIamBinding) ImportState(av io.Reader) error {
	cmiib.state = &computeMachineImageIamBindingState{}
	if err := json.NewDecoder(av).Decode(cmiib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cmiib.Type(), cmiib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeMachineImageIamBinding] has state.
func (cmiib *ComputeMachineImageIamBinding) State() (*computeMachineImageIamBindingState, bool) {
	return cmiib.state, cmiib.state != nil
}

// StateMust returns the state for [ComputeMachineImageIamBinding]. Panics if the state is nil.
func (cmiib *ComputeMachineImageIamBinding) StateMust() *computeMachineImageIamBindingState {
	if cmiib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cmiib.Type(), cmiib.LocalName()))
	}
	return cmiib.state
}

// ComputeMachineImageIamBindingArgs contains the configurations for google_compute_machine_image_iam_binding.
type ComputeMachineImageIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MachineImage: string, required
	MachineImage terra.StringValue `hcl:"machine_image,attr" validate:"required"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *computemachineimageiambinding.Condition `hcl:"condition,block"`
}
type computeMachineImageIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_compute_machine_image_iam_binding.
func (cmiib computeMachineImageIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(cmiib.ref.Append("etag"))
}

// Id returns a reference to field id of google_compute_machine_image_iam_binding.
func (cmiib computeMachineImageIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cmiib.ref.Append("id"))
}

// MachineImage returns a reference to field machine_image of google_compute_machine_image_iam_binding.
func (cmiib computeMachineImageIamBindingAttributes) MachineImage() terra.StringValue {
	return terra.ReferenceAsString(cmiib.ref.Append("machine_image"))
}

// Members returns a reference to field members of google_compute_machine_image_iam_binding.
func (cmiib computeMachineImageIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cmiib.ref.Append("members"))
}

// Project returns a reference to field project of google_compute_machine_image_iam_binding.
func (cmiib computeMachineImageIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cmiib.ref.Append("project"))
}

// Role returns a reference to field role of google_compute_machine_image_iam_binding.
func (cmiib computeMachineImageIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(cmiib.ref.Append("role"))
}

func (cmiib computeMachineImageIamBindingAttributes) Condition() terra.ListValue[computemachineimageiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[computemachineimageiambinding.ConditionAttributes](cmiib.ref.Append("condition"))
}

type computeMachineImageIamBindingState struct {
	Etag         string                                         `json:"etag"`
	Id           string                                         `json:"id"`
	MachineImage string                                         `json:"machine_image"`
	Members      []string                                       `json:"members"`
	Project      string                                         `json:"project"`
	Role         string                                         `json:"role"`
	Condition    []computemachineimageiambinding.ConditionState `json:"condition"`
}
