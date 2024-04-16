// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_iap_web_iam_binding

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

// Resource represents the Terraform resource google_iap_web_iam_binding.
type Resource struct {
	Name      string
	Args      Args
	state     *googleIapWebIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (giwib *Resource) Type() string {
	return "google_iap_web_iam_binding"
}

// LocalName returns the local name for [Resource].
func (giwib *Resource) LocalName() string {
	return giwib.Name
}

// Configuration returns the configuration (args) for [Resource].
func (giwib *Resource) Configuration() interface{} {
	return giwib.Args
}

// DependOn is used for other resources to depend on [Resource].
func (giwib *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(giwib)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (giwib *Resource) Dependencies() terra.Dependencies {
	return giwib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (giwib *Resource) LifecycleManagement() *terra.Lifecycle {
	return giwib.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (giwib *Resource) Attributes() googleIapWebIamBindingAttributes {
	return googleIapWebIamBindingAttributes{ref: terra.ReferenceResource(giwib)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (giwib *Resource) ImportState(state io.Reader) error {
	giwib.state = &googleIapWebIamBindingState{}
	if err := json.NewDecoder(state).Decode(giwib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", giwib.Type(), giwib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (giwib *Resource) State() (*googleIapWebIamBindingState, bool) {
	return giwib.state, giwib.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (giwib *Resource) StateMust() *googleIapWebIamBindingState {
	if giwib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", giwib.Type(), giwib.LocalName()))
	}
	return giwib.state
}

// Args contains the configurations for google_iap_web_iam_binding.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *Condition `hcl:"condition,block"`
}

type googleIapWebIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_iap_web_iam_binding.
func (giwib googleIapWebIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(giwib.ref.Append("etag"))
}

// Id returns a reference to field id of google_iap_web_iam_binding.
func (giwib googleIapWebIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(giwib.ref.Append("id"))
}

// Members returns a reference to field members of google_iap_web_iam_binding.
func (giwib googleIapWebIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](giwib.ref.Append("members"))
}

// Project returns a reference to field project of google_iap_web_iam_binding.
func (giwib googleIapWebIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(giwib.ref.Append("project"))
}

// Role returns a reference to field role of google_iap_web_iam_binding.
func (giwib googleIapWebIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(giwib.ref.Append("role"))
}

func (giwib googleIapWebIamBindingAttributes) Condition() terra.ListValue[ConditionAttributes] {
	return terra.ReferenceAsList[ConditionAttributes](giwib.ref.Append("condition"))
}

type googleIapWebIamBindingState struct {
	Etag      string           `json:"etag"`
	Id        string           `json:"id"`
	Members   []string         `json:"members"`
	Project   string           `json:"project"`
	Role      string           `json:"role"`
	Condition []ConditionState `json:"condition"`
}
