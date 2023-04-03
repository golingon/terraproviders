// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	notebooksruntimeiambinding "github.com/golingon/terraproviders/google/4.59.0/notebooksruntimeiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNotebooksRuntimeIamBinding creates a new instance of [NotebooksRuntimeIamBinding].
func NewNotebooksRuntimeIamBinding(name string, args NotebooksRuntimeIamBindingArgs) *NotebooksRuntimeIamBinding {
	return &NotebooksRuntimeIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NotebooksRuntimeIamBinding)(nil)

// NotebooksRuntimeIamBinding represents the Terraform resource google_notebooks_runtime_iam_binding.
type NotebooksRuntimeIamBinding struct {
	Name      string
	Args      NotebooksRuntimeIamBindingArgs
	state     *notebooksRuntimeIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NotebooksRuntimeIamBinding].
func (nrib *NotebooksRuntimeIamBinding) Type() string {
	return "google_notebooks_runtime_iam_binding"
}

// LocalName returns the local name for [NotebooksRuntimeIamBinding].
func (nrib *NotebooksRuntimeIamBinding) LocalName() string {
	return nrib.Name
}

// Configuration returns the configuration (args) for [NotebooksRuntimeIamBinding].
func (nrib *NotebooksRuntimeIamBinding) Configuration() interface{} {
	return nrib.Args
}

// DependOn is used for other resources to depend on [NotebooksRuntimeIamBinding].
func (nrib *NotebooksRuntimeIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(nrib)
}

// Dependencies returns the list of resources [NotebooksRuntimeIamBinding] depends_on.
func (nrib *NotebooksRuntimeIamBinding) Dependencies() terra.Dependencies {
	return nrib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NotebooksRuntimeIamBinding].
func (nrib *NotebooksRuntimeIamBinding) LifecycleManagement() *terra.Lifecycle {
	return nrib.Lifecycle
}

// Attributes returns the attributes for [NotebooksRuntimeIamBinding].
func (nrib *NotebooksRuntimeIamBinding) Attributes() notebooksRuntimeIamBindingAttributes {
	return notebooksRuntimeIamBindingAttributes{ref: terra.ReferenceResource(nrib)}
}

// ImportState imports the given attribute values into [NotebooksRuntimeIamBinding]'s state.
func (nrib *NotebooksRuntimeIamBinding) ImportState(av io.Reader) error {
	nrib.state = &notebooksRuntimeIamBindingState{}
	if err := json.NewDecoder(av).Decode(nrib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nrib.Type(), nrib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NotebooksRuntimeIamBinding] has state.
func (nrib *NotebooksRuntimeIamBinding) State() (*notebooksRuntimeIamBindingState, bool) {
	return nrib.state, nrib.state != nil
}

// StateMust returns the state for [NotebooksRuntimeIamBinding]. Panics if the state is nil.
func (nrib *NotebooksRuntimeIamBinding) StateMust() *notebooksRuntimeIamBindingState {
	if nrib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nrib.Type(), nrib.LocalName()))
	}
	return nrib.state
}

// NotebooksRuntimeIamBindingArgs contains the configurations for google_notebooks_runtime_iam_binding.
type NotebooksRuntimeIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// RuntimeName: string, required
	RuntimeName terra.StringValue `hcl:"runtime_name,attr" validate:"required"`
	// Condition: optional
	Condition *notebooksruntimeiambinding.Condition `hcl:"condition,block"`
}
type notebooksRuntimeIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_notebooks_runtime_iam_binding.
func (nrib notebooksRuntimeIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(nrib.ref.Append("etag"))
}

// Id returns a reference to field id of google_notebooks_runtime_iam_binding.
func (nrib notebooksRuntimeIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nrib.ref.Append("id"))
}

// Location returns a reference to field location of google_notebooks_runtime_iam_binding.
func (nrib notebooksRuntimeIamBindingAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(nrib.ref.Append("location"))
}

// Members returns a reference to field members of google_notebooks_runtime_iam_binding.
func (nrib notebooksRuntimeIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](nrib.ref.Append("members"))
}

// Project returns a reference to field project of google_notebooks_runtime_iam_binding.
func (nrib notebooksRuntimeIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(nrib.ref.Append("project"))
}

// Role returns a reference to field role of google_notebooks_runtime_iam_binding.
func (nrib notebooksRuntimeIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(nrib.ref.Append("role"))
}

// RuntimeName returns a reference to field runtime_name of google_notebooks_runtime_iam_binding.
func (nrib notebooksRuntimeIamBindingAttributes) RuntimeName() terra.StringValue {
	return terra.ReferenceAsString(nrib.ref.Append("runtime_name"))
}

func (nrib notebooksRuntimeIamBindingAttributes) Condition() terra.ListValue[notebooksruntimeiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[notebooksruntimeiambinding.ConditionAttributes](nrib.ref.Append("condition"))
}

type notebooksRuntimeIamBindingState struct {
	Etag        string                                      `json:"etag"`
	Id          string                                      `json:"id"`
	Location    string                                      `json:"location"`
	Members     []string                                    `json:"members"`
	Project     string                                      `json:"project"`
	Role        string                                      `json:"role"`
	RuntimeName string                                      `json:"runtime_name"`
	Condition   []notebooksruntimeiambinding.ConditionState `json:"condition"`
}
