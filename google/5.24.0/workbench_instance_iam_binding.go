// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	workbenchinstanceiambinding "github.com/golingon/terraproviders/google/5.24.0/workbenchinstanceiambinding"
	"io"
)

// NewWorkbenchInstanceIamBinding creates a new instance of [WorkbenchInstanceIamBinding].
func NewWorkbenchInstanceIamBinding(name string, args WorkbenchInstanceIamBindingArgs) *WorkbenchInstanceIamBinding {
	return &WorkbenchInstanceIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WorkbenchInstanceIamBinding)(nil)

// WorkbenchInstanceIamBinding represents the Terraform resource google_workbench_instance_iam_binding.
type WorkbenchInstanceIamBinding struct {
	Name      string
	Args      WorkbenchInstanceIamBindingArgs
	state     *workbenchInstanceIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WorkbenchInstanceIamBinding].
func (wiib *WorkbenchInstanceIamBinding) Type() string {
	return "google_workbench_instance_iam_binding"
}

// LocalName returns the local name for [WorkbenchInstanceIamBinding].
func (wiib *WorkbenchInstanceIamBinding) LocalName() string {
	return wiib.Name
}

// Configuration returns the configuration (args) for [WorkbenchInstanceIamBinding].
func (wiib *WorkbenchInstanceIamBinding) Configuration() interface{} {
	return wiib.Args
}

// DependOn is used for other resources to depend on [WorkbenchInstanceIamBinding].
func (wiib *WorkbenchInstanceIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(wiib)
}

// Dependencies returns the list of resources [WorkbenchInstanceIamBinding] depends_on.
func (wiib *WorkbenchInstanceIamBinding) Dependencies() terra.Dependencies {
	return wiib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WorkbenchInstanceIamBinding].
func (wiib *WorkbenchInstanceIamBinding) LifecycleManagement() *terra.Lifecycle {
	return wiib.Lifecycle
}

// Attributes returns the attributes for [WorkbenchInstanceIamBinding].
func (wiib *WorkbenchInstanceIamBinding) Attributes() workbenchInstanceIamBindingAttributes {
	return workbenchInstanceIamBindingAttributes{ref: terra.ReferenceResource(wiib)}
}

// ImportState imports the given attribute values into [WorkbenchInstanceIamBinding]'s state.
func (wiib *WorkbenchInstanceIamBinding) ImportState(av io.Reader) error {
	wiib.state = &workbenchInstanceIamBindingState{}
	if err := json.NewDecoder(av).Decode(wiib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wiib.Type(), wiib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WorkbenchInstanceIamBinding] has state.
func (wiib *WorkbenchInstanceIamBinding) State() (*workbenchInstanceIamBindingState, bool) {
	return wiib.state, wiib.state != nil
}

// StateMust returns the state for [WorkbenchInstanceIamBinding]. Panics if the state is nil.
func (wiib *WorkbenchInstanceIamBinding) StateMust() *workbenchInstanceIamBindingState {
	if wiib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wiib.Type(), wiib.LocalName()))
	}
	return wiib.state
}

// WorkbenchInstanceIamBindingArgs contains the configurations for google_workbench_instance_iam_binding.
type WorkbenchInstanceIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *workbenchinstanceiambinding.Condition `hcl:"condition,block"`
}
type workbenchInstanceIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_workbench_instance_iam_binding.
func (wiib workbenchInstanceIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(wiib.ref.Append("etag"))
}

// Id returns a reference to field id of google_workbench_instance_iam_binding.
func (wiib workbenchInstanceIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wiib.ref.Append("id"))
}

// Location returns a reference to field location of google_workbench_instance_iam_binding.
func (wiib workbenchInstanceIamBindingAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(wiib.ref.Append("location"))
}

// Members returns a reference to field members of google_workbench_instance_iam_binding.
func (wiib workbenchInstanceIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](wiib.ref.Append("members"))
}

// Name returns a reference to field name of google_workbench_instance_iam_binding.
func (wiib workbenchInstanceIamBindingAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wiib.ref.Append("name"))
}

// Project returns a reference to field project of google_workbench_instance_iam_binding.
func (wiib workbenchInstanceIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(wiib.ref.Append("project"))
}

// Role returns a reference to field role of google_workbench_instance_iam_binding.
func (wiib workbenchInstanceIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(wiib.ref.Append("role"))
}

func (wiib workbenchInstanceIamBindingAttributes) Condition() terra.ListValue[workbenchinstanceiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[workbenchinstanceiambinding.ConditionAttributes](wiib.ref.Append("condition"))
}

type workbenchInstanceIamBindingState struct {
	Etag      string                                       `json:"etag"`
	Id        string                                       `json:"id"`
	Location  string                                       `json:"location"`
	Members   []string                                     `json:"members"`
	Name      string                                       `json:"name"`
	Project   string                                       `json:"project"`
	Role      string                                       `json:"role"`
	Condition []workbenchinstanceiambinding.ConditionState `json:"condition"`
}
