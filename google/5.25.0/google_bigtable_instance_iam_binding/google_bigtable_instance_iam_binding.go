// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_bigtable_instance_iam_binding

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

// Resource represents the Terraform resource google_bigtable_instance_iam_binding.
type Resource struct {
	Name      string
	Args      Args
	state     *googleBigtableInstanceIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gbiib *Resource) Type() string {
	return "google_bigtable_instance_iam_binding"
}

// LocalName returns the local name for [Resource].
func (gbiib *Resource) LocalName() string {
	return gbiib.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gbiib *Resource) Configuration() interface{} {
	return gbiib.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gbiib *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gbiib)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gbiib *Resource) Dependencies() terra.Dependencies {
	return gbiib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gbiib *Resource) LifecycleManagement() *terra.Lifecycle {
	return gbiib.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gbiib *Resource) Attributes() googleBigtableInstanceIamBindingAttributes {
	return googleBigtableInstanceIamBindingAttributes{ref: terra.ReferenceResource(gbiib)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gbiib *Resource) ImportState(state io.Reader) error {
	gbiib.state = &googleBigtableInstanceIamBindingState{}
	if err := json.NewDecoder(state).Decode(gbiib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gbiib.Type(), gbiib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gbiib *Resource) State() (*googleBigtableInstanceIamBindingState, bool) {
	return gbiib.state, gbiib.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gbiib *Resource) StateMust() *googleBigtableInstanceIamBindingState {
	if gbiib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gbiib.Type(), gbiib.LocalName()))
	}
	return gbiib.state
}

// Args contains the configurations for google_bigtable_instance_iam_binding.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Instance: string, required
	Instance terra.StringValue `hcl:"instance,attr" validate:"required"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *Condition `hcl:"condition,block"`
}

type googleBigtableInstanceIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_bigtable_instance_iam_binding.
func (gbiib googleBigtableInstanceIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gbiib.ref.Append("etag"))
}

// Id returns a reference to field id of google_bigtable_instance_iam_binding.
func (gbiib googleBigtableInstanceIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gbiib.ref.Append("id"))
}

// Instance returns a reference to field instance of google_bigtable_instance_iam_binding.
func (gbiib googleBigtableInstanceIamBindingAttributes) Instance() terra.StringValue {
	return terra.ReferenceAsString(gbiib.ref.Append("instance"))
}

// Members returns a reference to field members of google_bigtable_instance_iam_binding.
func (gbiib googleBigtableInstanceIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](gbiib.ref.Append("members"))
}

// Project returns a reference to field project of google_bigtable_instance_iam_binding.
func (gbiib googleBigtableInstanceIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gbiib.ref.Append("project"))
}

// Role returns a reference to field role of google_bigtable_instance_iam_binding.
func (gbiib googleBigtableInstanceIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(gbiib.ref.Append("role"))
}

func (gbiib googleBigtableInstanceIamBindingAttributes) Condition() terra.ListValue[ConditionAttributes] {
	return terra.ReferenceAsList[ConditionAttributes](gbiib.ref.Append("condition"))
}

type googleBigtableInstanceIamBindingState struct {
	Etag      string           `json:"etag"`
	Id        string           `json:"id"`
	Instance  string           `json:"instance"`
	Members   []string         `json:"members"`
	Project   string           `json:"project"`
	Role      string           `json:"role"`
	Condition []ConditionState `json:"condition"`
}
