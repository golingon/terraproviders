// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computesnapshotiambinding "github.com/golingon/terraproviders/googlebeta/4.71.0/computesnapshotiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeSnapshotIamBinding creates a new instance of [ComputeSnapshotIamBinding].
func NewComputeSnapshotIamBinding(name string, args ComputeSnapshotIamBindingArgs) *ComputeSnapshotIamBinding {
	return &ComputeSnapshotIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeSnapshotIamBinding)(nil)

// ComputeSnapshotIamBinding represents the Terraform resource google_compute_snapshot_iam_binding.
type ComputeSnapshotIamBinding struct {
	Name      string
	Args      ComputeSnapshotIamBindingArgs
	state     *computeSnapshotIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeSnapshotIamBinding].
func (csib *ComputeSnapshotIamBinding) Type() string {
	return "google_compute_snapshot_iam_binding"
}

// LocalName returns the local name for [ComputeSnapshotIamBinding].
func (csib *ComputeSnapshotIamBinding) LocalName() string {
	return csib.Name
}

// Configuration returns the configuration (args) for [ComputeSnapshotIamBinding].
func (csib *ComputeSnapshotIamBinding) Configuration() interface{} {
	return csib.Args
}

// DependOn is used for other resources to depend on [ComputeSnapshotIamBinding].
func (csib *ComputeSnapshotIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(csib)
}

// Dependencies returns the list of resources [ComputeSnapshotIamBinding] depends_on.
func (csib *ComputeSnapshotIamBinding) Dependencies() terra.Dependencies {
	return csib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeSnapshotIamBinding].
func (csib *ComputeSnapshotIamBinding) LifecycleManagement() *terra.Lifecycle {
	return csib.Lifecycle
}

// Attributes returns the attributes for [ComputeSnapshotIamBinding].
func (csib *ComputeSnapshotIamBinding) Attributes() computeSnapshotIamBindingAttributes {
	return computeSnapshotIamBindingAttributes{ref: terra.ReferenceResource(csib)}
}

// ImportState imports the given attribute values into [ComputeSnapshotIamBinding]'s state.
func (csib *ComputeSnapshotIamBinding) ImportState(av io.Reader) error {
	csib.state = &computeSnapshotIamBindingState{}
	if err := json.NewDecoder(av).Decode(csib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", csib.Type(), csib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeSnapshotIamBinding] has state.
func (csib *ComputeSnapshotIamBinding) State() (*computeSnapshotIamBindingState, bool) {
	return csib.state, csib.state != nil
}

// StateMust returns the state for [ComputeSnapshotIamBinding]. Panics if the state is nil.
func (csib *ComputeSnapshotIamBinding) StateMust() *computeSnapshotIamBindingState {
	if csib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", csib.Type(), csib.LocalName()))
	}
	return csib.state
}

// ComputeSnapshotIamBindingArgs contains the configurations for google_compute_snapshot_iam_binding.
type ComputeSnapshotIamBindingArgs struct {
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
	Condition *computesnapshotiambinding.Condition `hcl:"condition,block"`
}
type computeSnapshotIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_compute_snapshot_iam_binding.
func (csib computeSnapshotIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(csib.ref.Append("etag"))
}

// Id returns a reference to field id of google_compute_snapshot_iam_binding.
func (csib computeSnapshotIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(csib.ref.Append("id"))
}

// Members returns a reference to field members of google_compute_snapshot_iam_binding.
func (csib computeSnapshotIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](csib.ref.Append("members"))
}

// Name returns a reference to field name of google_compute_snapshot_iam_binding.
func (csib computeSnapshotIamBindingAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(csib.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_snapshot_iam_binding.
func (csib computeSnapshotIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(csib.ref.Append("project"))
}

// Role returns a reference to field role of google_compute_snapshot_iam_binding.
func (csib computeSnapshotIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(csib.ref.Append("role"))
}

func (csib computeSnapshotIamBindingAttributes) Condition() terra.ListValue[computesnapshotiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[computesnapshotiambinding.ConditionAttributes](csib.ref.Append("condition"))
}

type computeSnapshotIamBindingState struct {
	Etag      string                                     `json:"etag"`
	Id        string                                     `json:"id"`
	Members   []string                                   `json:"members"`
	Name      string                                     `json:"name"`
	Project   string                                     `json:"project"`
	Role      string                                     `json:"role"`
	Condition []computesnapshotiambinding.ConditionState `json:"condition"`
}
