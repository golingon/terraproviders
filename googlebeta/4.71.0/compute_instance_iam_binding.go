// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computeinstanceiambinding "github.com/golingon/terraproviders/googlebeta/4.71.0/computeinstanceiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeInstanceIamBinding creates a new instance of [ComputeInstanceIamBinding].
func NewComputeInstanceIamBinding(name string, args ComputeInstanceIamBindingArgs) *ComputeInstanceIamBinding {
	return &ComputeInstanceIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeInstanceIamBinding)(nil)

// ComputeInstanceIamBinding represents the Terraform resource google_compute_instance_iam_binding.
type ComputeInstanceIamBinding struct {
	Name      string
	Args      ComputeInstanceIamBindingArgs
	state     *computeInstanceIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeInstanceIamBinding].
func (ciib *ComputeInstanceIamBinding) Type() string {
	return "google_compute_instance_iam_binding"
}

// LocalName returns the local name for [ComputeInstanceIamBinding].
func (ciib *ComputeInstanceIamBinding) LocalName() string {
	return ciib.Name
}

// Configuration returns the configuration (args) for [ComputeInstanceIamBinding].
func (ciib *ComputeInstanceIamBinding) Configuration() interface{} {
	return ciib.Args
}

// DependOn is used for other resources to depend on [ComputeInstanceIamBinding].
func (ciib *ComputeInstanceIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(ciib)
}

// Dependencies returns the list of resources [ComputeInstanceIamBinding] depends_on.
func (ciib *ComputeInstanceIamBinding) Dependencies() terra.Dependencies {
	return ciib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeInstanceIamBinding].
func (ciib *ComputeInstanceIamBinding) LifecycleManagement() *terra.Lifecycle {
	return ciib.Lifecycle
}

// Attributes returns the attributes for [ComputeInstanceIamBinding].
func (ciib *ComputeInstanceIamBinding) Attributes() computeInstanceIamBindingAttributes {
	return computeInstanceIamBindingAttributes{ref: terra.ReferenceResource(ciib)}
}

// ImportState imports the given attribute values into [ComputeInstanceIamBinding]'s state.
func (ciib *ComputeInstanceIamBinding) ImportState(av io.Reader) error {
	ciib.state = &computeInstanceIamBindingState{}
	if err := json.NewDecoder(av).Decode(ciib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ciib.Type(), ciib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeInstanceIamBinding] has state.
func (ciib *ComputeInstanceIamBinding) State() (*computeInstanceIamBindingState, bool) {
	return ciib.state, ciib.state != nil
}

// StateMust returns the state for [ComputeInstanceIamBinding]. Panics if the state is nil.
func (ciib *ComputeInstanceIamBinding) StateMust() *computeInstanceIamBindingState {
	if ciib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ciib.Type(), ciib.LocalName()))
	}
	return ciib.state
}

// ComputeInstanceIamBindingArgs contains the configurations for google_compute_instance_iam_binding.
type ComputeInstanceIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceName: string, required
	InstanceName terra.StringValue `hcl:"instance_name,attr" validate:"required"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// Condition: optional
	Condition *computeinstanceiambinding.Condition `hcl:"condition,block"`
}
type computeInstanceIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_compute_instance_iam_binding.
func (ciib computeInstanceIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ciib.ref.Append("etag"))
}

// Id returns a reference to field id of google_compute_instance_iam_binding.
func (ciib computeInstanceIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ciib.ref.Append("id"))
}

// InstanceName returns a reference to field instance_name of google_compute_instance_iam_binding.
func (ciib computeInstanceIamBindingAttributes) InstanceName() terra.StringValue {
	return terra.ReferenceAsString(ciib.ref.Append("instance_name"))
}

// Members returns a reference to field members of google_compute_instance_iam_binding.
func (ciib computeInstanceIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ciib.ref.Append("members"))
}

// Project returns a reference to field project of google_compute_instance_iam_binding.
func (ciib computeInstanceIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ciib.ref.Append("project"))
}

// Role returns a reference to field role of google_compute_instance_iam_binding.
func (ciib computeInstanceIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(ciib.ref.Append("role"))
}

// Zone returns a reference to field zone of google_compute_instance_iam_binding.
func (ciib computeInstanceIamBindingAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(ciib.ref.Append("zone"))
}

func (ciib computeInstanceIamBindingAttributes) Condition() terra.ListValue[computeinstanceiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[computeinstanceiambinding.ConditionAttributes](ciib.ref.Append("condition"))
}

type computeInstanceIamBindingState struct {
	Etag         string                                     `json:"etag"`
	Id           string                                     `json:"id"`
	InstanceName string                                     `json:"instance_name"`
	Members      []string                                   `json:"members"`
	Project      string                                     `json:"project"`
	Role         string                                     `json:"role"`
	Zone         string                                     `json:"zone"`
	Condition    []computeinstanceiambinding.ConditionState `json:"condition"`
}
