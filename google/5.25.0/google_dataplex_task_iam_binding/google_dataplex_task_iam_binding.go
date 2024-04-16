// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_dataplex_task_iam_binding

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

// Resource represents the Terraform resource google_dataplex_task_iam_binding.
type Resource struct {
	Name      string
	Args      Args
	state     *googleDataplexTaskIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gdtib *Resource) Type() string {
	return "google_dataplex_task_iam_binding"
}

// LocalName returns the local name for [Resource].
func (gdtib *Resource) LocalName() string {
	return gdtib.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gdtib *Resource) Configuration() interface{} {
	return gdtib.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gdtib *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gdtib)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gdtib *Resource) Dependencies() terra.Dependencies {
	return gdtib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gdtib *Resource) LifecycleManagement() *terra.Lifecycle {
	return gdtib.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gdtib *Resource) Attributes() googleDataplexTaskIamBindingAttributes {
	return googleDataplexTaskIamBindingAttributes{ref: terra.ReferenceResource(gdtib)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gdtib *Resource) ImportState(state io.Reader) error {
	gdtib.state = &googleDataplexTaskIamBindingState{}
	if err := json.NewDecoder(state).Decode(gdtib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gdtib.Type(), gdtib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gdtib *Resource) State() (*googleDataplexTaskIamBindingState, bool) {
	return gdtib.state, gdtib.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gdtib *Resource) StateMust() *googleDataplexTaskIamBindingState {
	if gdtib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gdtib.Type(), gdtib.LocalName()))
	}
	return gdtib.state
}

// Args contains the configurations for google_dataplex_task_iam_binding.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Lake: string, required
	Lake terra.StringValue `hcl:"lake,attr" validate:"required"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// TaskId: string, required
	TaskId terra.StringValue `hcl:"task_id,attr" validate:"required"`
	// Condition: optional
	Condition *Condition `hcl:"condition,block"`
}

type googleDataplexTaskIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_dataplex_task_iam_binding.
func (gdtib googleDataplexTaskIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gdtib.ref.Append("etag"))
}

// Id returns a reference to field id of google_dataplex_task_iam_binding.
func (gdtib googleDataplexTaskIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gdtib.ref.Append("id"))
}

// Lake returns a reference to field lake of google_dataplex_task_iam_binding.
func (gdtib googleDataplexTaskIamBindingAttributes) Lake() terra.StringValue {
	return terra.ReferenceAsString(gdtib.ref.Append("lake"))
}

// Location returns a reference to field location of google_dataplex_task_iam_binding.
func (gdtib googleDataplexTaskIamBindingAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gdtib.ref.Append("location"))
}

// Members returns a reference to field members of google_dataplex_task_iam_binding.
func (gdtib googleDataplexTaskIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](gdtib.ref.Append("members"))
}

// Project returns a reference to field project of google_dataplex_task_iam_binding.
func (gdtib googleDataplexTaskIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gdtib.ref.Append("project"))
}

// Role returns a reference to field role of google_dataplex_task_iam_binding.
func (gdtib googleDataplexTaskIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(gdtib.ref.Append("role"))
}

// TaskId returns a reference to field task_id of google_dataplex_task_iam_binding.
func (gdtib googleDataplexTaskIamBindingAttributes) TaskId() terra.StringValue {
	return terra.ReferenceAsString(gdtib.ref.Append("task_id"))
}

func (gdtib googleDataplexTaskIamBindingAttributes) Condition() terra.ListValue[ConditionAttributes] {
	return terra.ReferenceAsList[ConditionAttributes](gdtib.ref.Append("condition"))
}

type googleDataplexTaskIamBindingState struct {
	Etag      string           `json:"etag"`
	Id        string           `json:"id"`
	Lake      string           `json:"lake"`
	Location  string           `json:"location"`
	Members   []string         `json:"members"`
	Project   string           `json:"project"`
	Role      string           `json:"role"`
	TaskId    string           `json:"task_id"`
	Condition []ConditionState `json:"condition"`
}
