// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_gke_hub_feature_iam_binding

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

// Resource represents the Terraform resource google_gke_hub_feature_iam_binding.
type Resource struct {
	Name      string
	Args      Args
	state     *googleGkeHubFeatureIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gghfib *Resource) Type() string {
	return "google_gke_hub_feature_iam_binding"
}

// LocalName returns the local name for [Resource].
func (gghfib *Resource) LocalName() string {
	return gghfib.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gghfib *Resource) Configuration() interface{} {
	return gghfib.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gghfib *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gghfib)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gghfib *Resource) Dependencies() terra.Dependencies {
	return gghfib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gghfib *Resource) LifecycleManagement() *terra.Lifecycle {
	return gghfib.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gghfib *Resource) Attributes() googleGkeHubFeatureIamBindingAttributes {
	return googleGkeHubFeatureIamBindingAttributes{ref: terra.ReferenceResource(gghfib)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gghfib *Resource) ImportState(state io.Reader) error {
	gghfib.state = &googleGkeHubFeatureIamBindingState{}
	if err := json.NewDecoder(state).Decode(gghfib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gghfib.Type(), gghfib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gghfib *Resource) State() (*googleGkeHubFeatureIamBindingState, bool) {
	return gghfib.state, gghfib.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gghfib *Resource) StateMust() *googleGkeHubFeatureIamBindingState {
	if gghfib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gghfib.Type(), gghfib.LocalName()))
	}
	return gghfib.state
}

// Args contains the configurations for google_gke_hub_feature_iam_binding.
type Args struct {
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
	Condition *Condition `hcl:"condition,block"`
}

type googleGkeHubFeatureIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_gke_hub_feature_iam_binding.
func (gghfib googleGkeHubFeatureIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gghfib.ref.Append("etag"))
}

// Id returns a reference to field id of google_gke_hub_feature_iam_binding.
func (gghfib googleGkeHubFeatureIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gghfib.ref.Append("id"))
}

// Location returns a reference to field location of google_gke_hub_feature_iam_binding.
func (gghfib googleGkeHubFeatureIamBindingAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gghfib.ref.Append("location"))
}

// Members returns a reference to field members of google_gke_hub_feature_iam_binding.
func (gghfib googleGkeHubFeatureIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](gghfib.ref.Append("members"))
}

// Name returns a reference to field name of google_gke_hub_feature_iam_binding.
func (gghfib googleGkeHubFeatureIamBindingAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gghfib.ref.Append("name"))
}

// Project returns a reference to field project of google_gke_hub_feature_iam_binding.
func (gghfib googleGkeHubFeatureIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gghfib.ref.Append("project"))
}

// Role returns a reference to field role of google_gke_hub_feature_iam_binding.
func (gghfib googleGkeHubFeatureIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(gghfib.ref.Append("role"))
}

func (gghfib googleGkeHubFeatureIamBindingAttributes) Condition() terra.ListValue[ConditionAttributes] {
	return terra.ReferenceAsList[ConditionAttributes](gghfib.ref.Append("condition"))
}

type googleGkeHubFeatureIamBindingState struct {
	Etag      string           `json:"etag"`
	Id        string           `json:"id"`
	Location  string           `json:"location"`
	Members   []string         `json:"members"`
	Name      string           `json:"name"`
	Project   string           `json:"project"`
	Role      string           `json:"role"`
	Condition []ConditionState `json:"condition"`
}
