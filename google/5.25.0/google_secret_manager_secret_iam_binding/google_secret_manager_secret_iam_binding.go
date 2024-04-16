// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_secret_manager_secret_iam_binding

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

// Resource represents the Terraform resource google_secret_manager_secret_iam_binding.
type Resource struct {
	Name      string
	Args      Args
	state     *googleSecretManagerSecretIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gsmsib *Resource) Type() string {
	return "google_secret_manager_secret_iam_binding"
}

// LocalName returns the local name for [Resource].
func (gsmsib *Resource) LocalName() string {
	return gsmsib.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gsmsib *Resource) Configuration() interface{} {
	return gsmsib.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gsmsib *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gsmsib)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gsmsib *Resource) Dependencies() terra.Dependencies {
	return gsmsib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gsmsib *Resource) LifecycleManagement() *terra.Lifecycle {
	return gsmsib.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gsmsib *Resource) Attributes() googleSecretManagerSecretIamBindingAttributes {
	return googleSecretManagerSecretIamBindingAttributes{ref: terra.ReferenceResource(gsmsib)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gsmsib *Resource) ImportState(state io.Reader) error {
	gsmsib.state = &googleSecretManagerSecretIamBindingState{}
	if err := json.NewDecoder(state).Decode(gsmsib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gsmsib.Type(), gsmsib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gsmsib *Resource) State() (*googleSecretManagerSecretIamBindingState, bool) {
	return gsmsib.state, gsmsib.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gsmsib *Resource) StateMust() *googleSecretManagerSecretIamBindingState {
	if gsmsib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gsmsib.Type(), gsmsib.LocalName()))
	}
	return gsmsib.state
}

// Args contains the configurations for google_secret_manager_secret_iam_binding.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// SecretId: string, required
	SecretId terra.StringValue `hcl:"secret_id,attr" validate:"required"`
	// Condition: optional
	Condition *Condition `hcl:"condition,block"`
}

type googleSecretManagerSecretIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_secret_manager_secret_iam_binding.
func (gsmsib googleSecretManagerSecretIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gsmsib.ref.Append("etag"))
}

// Id returns a reference to field id of google_secret_manager_secret_iam_binding.
func (gsmsib googleSecretManagerSecretIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gsmsib.ref.Append("id"))
}

// Members returns a reference to field members of google_secret_manager_secret_iam_binding.
func (gsmsib googleSecretManagerSecretIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](gsmsib.ref.Append("members"))
}

// Project returns a reference to field project of google_secret_manager_secret_iam_binding.
func (gsmsib googleSecretManagerSecretIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gsmsib.ref.Append("project"))
}

// Role returns a reference to field role of google_secret_manager_secret_iam_binding.
func (gsmsib googleSecretManagerSecretIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(gsmsib.ref.Append("role"))
}

// SecretId returns a reference to field secret_id of google_secret_manager_secret_iam_binding.
func (gsmsib googleSecretManagerSecretIamBindingAttributes) SecretId() terra.StringValue {
	return terra.ReferenceAsString(gsmsib.ref.Append("secret_id"))
}

func (gsmsib googleSecretManagerSecretIamBindingAttributes) Condition() terra.ListValue[ConditionAttributes] {
	return terra.ReferenceAsList[ConditionAttributes](gsmsib.ref.Append("condition"))
}

type googleSecretManagerSecretIamBindingState struct {
	Etag      string           `json:"etag"`
	Id        string           `json:"id"`
	Members   []string         `json:"members"`
	Project   string           `json:"project"`
	Role      string           `json:"role"`
	SecretId  string           `json:"secret_id"`
	Condition []ConditionState `json:"condition"`
}
