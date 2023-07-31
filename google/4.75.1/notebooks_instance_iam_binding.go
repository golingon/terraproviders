// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	notebooksinstanceiambinding "github.com/golingon/terraproviders/google/4.75.1/notebooksinstanceiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNotebooksInstanceIamBinding creates a new instance of [NotebooksInstanceIamBinding].
func NewNotebooksInstanceIamBinding(name string, args NotebooksInstanceIamBindingArgs) *NotebooksInstanceIamBinding {
	return &NotebooksInstanceIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NotebooksInstanceIamBinding)(nil)

// NotebooksInstanceIamBinding represents the Terraform resource google_notebooks_instance_iam_binding.
type NotebooksInstanceIamBinding struct {
	Name      string
	Args      NotebooksInstanceIamBindingArgs
	state     *notebooksInstanceIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NotebooksInstanceIamBinding].
func (niib *NotebooksInstanceIamBinding) Type() string {
	return "google_notebooks_instance_iam_binding"
}

// LocalName returns the local name for [NotebooksInstanceIamBinding].
func (niib *NotebooksInstanceIamBinding) LocalName() string {
	return niib.Name
}

// Configuration returns the configuration (args) for [NotebooksInstanceIamBinding].
func (niib *NotebooksInstanceIamBinding) Configuration() interface{} {
	return niib.Args
}

// DependOn is used for other resources to depend on [NotebooksInstanceIamBinding].
func (niib *NotebooksInstanceIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(niib)
}

// Dependencies returns the list of resources [NotebooksInstanceIamBinding] depends_on.
func (niib *NotebooksInstanceIamBinding) Dependencies() terra.Dependencies {
	return niib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NotebooksInstanceIamBinding].
func (niib *NotebooksInstanceIamBinding) LifecycleManagement() *terra.Lifecycle {
	return niib.Lifecycle
}

// Attributes returns the attributes for [NotebooksInstanceIamBinding].
func (niib *NotebooksInstanceIamBinding) Attributes() notebooksInstanceIamBindingAttributes {
	return notebooksInstanceIamBindingAttributes{ref: terra.ReferenceResource(niib)}
}

// ImportState imports the given attribute values into [NotebooksInstanceIamBinding]'s state.
func (niib *NotebooksInstanceIamBinding) ImportState(av io.Reader) error {
	niib.state = &notebooksInstanceIamBindingState{}
	if err := json.NewDecoder(av).Decode(niib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", niib.Type(), niib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NotebooksInstanceIamBinding] has state.
func (niib *NotebooksInstanceIamBinding) State() (*notebooksInstanceIamBindingState, bool) {
	return niib.state, niib.state != nil
}

// StateMust returns the state for [NotebooksInstanceIamBinding]. Panics if the state is nil.
func (niib *NotebooksInstanceIamBinding) StateMust() *notebooksInstanceIamBindingState {
	if niib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", niib.Type(), niib.LocalName()))
	}
	return niib.state
}

// NotebooksInstanceIamBindingArgs contains the configurations for google_notebooks_instance_iam_binding.
type NotebooksInstanceIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceName: string, required
	InstanceName terra.StringValue `hcl:"instance_name,attr" validate:"required"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *notebooksinstanceiambinding.Condition `hcl:"condition,block"`
}
type notebooksInstanceIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_notebooks_instance_iam_binding.
func (niib notebooksInstanceIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(niib.ref.Append("etag"))
}

// Id returns a reference to field id of google_notebooks_instance_iam_binding.
func (niib notebooksInstanceIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(niib.ref.Append("id"))
}

// InstanceName returns a reference to field instance_name of google_notebooks_instance_iam_binding.
func (niib notebooksInstanceIamBindingAttributes) InstanceName() terra.StringValue {
	return terra.ReferenceAsString(niib.ref.Append("instance_name"))
}

// Location returns a reference to field location of google_notebooks_instance_iam_binding.
func (niib notebooksInstanceIamBindingAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(niib.ref.Append("location"))
}

// Members returns a reference to field members of google_notebooks_instance_iam_binding.
func (niib notebooksInstanceIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](niib.ref.Append("members"))
}

// Project returns a reference to field project of google_notebooks_instance_iam_binding.
func (niib notebooksInstanceIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(niib.ref.Append("project"))
}

// Role returns a reference to field role of google_notebooks_instance_iam_binding.
func (niib notebooksInstanceIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(niib.ref.Append("role"))
}

func (niib notebooksInstanceIamBindingAttributes) Condition() terra.ListValue[notebooksinstanceiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[notebooksinstanceiambinding.ConditionAttributes](niib.ref.Append("condition"))
}

type notebooksInstanceIamBindingState struct {
	Etag         string                                       `json:"etag"`
	Id           string                                       `json:"id"`
	InstanceName string                                       `json:"instance_name"`
	Location     string                                       `json:"location"`
	Members      []string                                     `json:"members"`
	Project      string                                       `json:"project"`
	Role         string                                       `json:"role"`
	Condition    []notebooksinstanceiambinding.ConditionState `json:"condition"`
}
