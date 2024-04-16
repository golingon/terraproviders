// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_gke_hub_namespace

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

// Resource represents the Terraform resource google_gke_hub_namespace.
type Resource struct {
	Name      string
	Args      Args
	state     *googleGkeHubNamespaceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gghn *Resource) Type() string {
	return "google_gke_hub_namespace"
}

// LocalName returns the local name for [Resource].
func (gghn *Resource) LocalName() string {
	return gghn.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gghn *Resource) Configuration() interface{} {
	return gghn.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gghn *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gghn)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gghn *Resource) Dependencies() terra.Dependencies {
	return gghn.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gghn *Resource) LifecycleManagement() *terra.Lifecycle {
	return gghn.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gghn *Resource) Attributes() googleGkeHubNamespaceAttributes {
	return googleGkeHubNamespaceAttributes{ref: terra.ReferenceResource(gghn)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gghn *Resource) ImportState(state io.Reader) error {
	gghn.state = &googleGkeHubNamespaceState{}
	if err := json.NewDecoder(state).Decode(gghn.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gghn.Type(), gghn.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gghn *Resource) State() (*googleGkeHubNamespaceState, bool) {
	return gghn.state, gghn.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gghn *Resource) StateMust() *googleGkeHubNamespaceState {
	if gghn.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gghn.Type(), gghn.LocalName()))
	}
	return gghn.state
}

// Args contains the configurations for google_gke_hub_namespace.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// NamespaceLabels: map of string, optional
	NamespaceLabels terra.MapValue[terra.StringValue] `hcl:"namespace_labels,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Scope: string, required
	Scope terra.StringValue `hcl:"scope,attr" validate:"required"`
	// ScopeId: string, required
	ScopeId terra.StringValue `hcl:"scope_id,attr" validate:"required"`
	// ScopeNamespaceId: string, required
	ScopeNamespaceId terra.StringValue `hcl:"scope_namespace_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleGkeHubNamespaceAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_gke_hub_namespace.
func (gghn googleGkeHubNamespaceAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(gghn.ref.Append("create_time"))
}

// DeleteTime returns a reference to field delete_time of google_gke_hub_namespace.
func (gghn googleGkeHubNamespaceAttributes) DeleteTime() terra.StringValue {
	return terra.ReferenceAsString(gghn.ref.Append("delete_time"))
}

// EffectiveLabels returns a reference to field effective_labels of google_gke_hub_namespace.
func (gghn googleGkeHubNamespaceAttributes) EffectiveLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gghn.ref.Append("effective_labels"))
}

// Id returns a reference to field id of google_gke_hub_namespace.
func (gghn googleGkeHubNamespaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gghn.ref.Append("id"))
}

// Labels returns a reference to field labels of google_gke_hub_namespace.
func (gghn googleGkeHubNamespaceAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gghn.ref.Append("labels"))
}

// Name returns a reference to field name of google_gke_hub_namespace.
func (gghn googleGkeHubNamespaceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gghn.ref.Append("name"))
}

// NamespaceLabels returns a reference to field namespace_labels of google_gke_hub_namespace.
func (gghn googleGkeHubNamespaceAttributes) NamespaceLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gghn.ref.Append("namespace_labels"))
}

// Project returns a reference to field project of google_gke_hub_namespace.
func (gghn googleGkeHubNamespaceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gghn.ref.Append("project"))
}

// Scope returns a reference to field scope of google_gke_hub_namespace.
func (gghn googleGkeHubNamespaceAttributes) Scope() terra.StringValue {
	return terra.ReferenceAsString(gghn.ref.Append("scope"))
}

// ScopeId returns a reference to field scope_id of google_gke_hub_namespace.
func (gghn googleGkeHubNamespaceAttributes) ScopeId() terra.StringValue {
	return terra.ReferenceAsString(gghn.ref.Append("scope_id"))
}

// ScopeNamespaceId returns a reference to field scope_namespace_id of google_gke_hub_namespace.
func (gghn googleGkeHubNamespaceAttributes) ScopeNamespaceId() terra.StringValue {
	return terra.ReferenceAsString(gghn.ref.Append("scope_namespace_id"))
}

// TerraformLabels returns a reference to field terraform_labels of google_gke_hub_namespace.
func (gghn googleGkeHubNamespaceAttributes) TerraformLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gghn.ref.Append("terraform_labels"))
}

// Uid returns a reference to field uid of google_gke_hub_namespace.
func (gghn googleGkeHubNamespaceAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(gghn.ref.Append("uid"))
}

// UpdateTime returns a reference to field update_time of google_gke_hub_namespace.
func (gghn googleGkeHubNamespaceAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(gghn.ref.Append("update_time"))
}

func (gghn googleGkeHubNamespaceAttributes) State() terra.ListValue[StateAttributes] {
	return terra.ReferenceAsList[StateAttributes](gghn.ref.Append("state"))
}

func (gghn googleGkeHubNamespaceAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gghn.ref.Append("timeouts"))
}

type googleGkeHubNamespaceState struct {
	CreateTime       string            `json:"create_time"`
	DeleteTime       string            `json:"delete_time"`
	EffectiveLabels  map[string]string `json:"effective_labels"`
	Id               string            `json:"id"`
	Labels           map[string]string `json:"labels"`
	Name             string            `json:"name"`
	NamespaceLabels  map[string]string `json:"namespace_labels"`
	Project          string            `json:"project"`
	Scope            string            `json:"scope"`
	ScopeId          string            `json:"scope_id"`
	ScopeNamespaceId string            `json:"scope_namespace_id"`
	TerraformLabels  map[string]string `json:"terraform_labels"`
	Uid              string            `json:"uid"`
	UpdateTime       string            `json:"update_time"`
	State            []StateState      `json:"state"`
	Timeouts         *TimeoutsState    `json:"timeouts"`
}
