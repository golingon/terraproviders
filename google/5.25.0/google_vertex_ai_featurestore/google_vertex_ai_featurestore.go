// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_vertex_ai_featurestore

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

// Resource represents the Terraform resource google_vertex_ai_featurestore.
type Resource struct {
	Name      string
	Args      Args
	state     *googleVertexAiFeaturestoreState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gvaf *Resource) Type() string {
	return "google_vertex_ai_featurestore"
}

// LocalName returns the local name for [Resource].
func (gvaf *Resource) LocalName() string {
	return gvaf.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gvaf *Resource) Configuration() interface{} {
	return gvaf.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gvaf *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gvaf)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gvaf *Resource) Dependencies() terra.Dependencies {
	return gvaf.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gvaf *Resource) LifecycleManagement() *terra.Lifecycle {
	return gvaf.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gvaf *Resource) Attributes() googleVertexAiFeaturestoreAttributes {
	return googleVertexAiFeaturestoreAttributes{ref: terra.ReferenceResource(gvaf)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gvaf *Resource) ImportState(state io.Reader) error {
	gvaf.state = &googleVertexAiFeaturestoreState{}
	if err := json.NewDecoder(state).Decode(gvaf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gvaf.Type(), gvaf.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gvaf *Resource) State() (*googleVertexAiFeaturestoreState, bool) {
	return gvaf.state, gvaf.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gvaf *Resource) StateMust() *googleVertexAiFeaturestoreState {
	if gvaf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gvaf.Type(), gvaf.LocalName()))
	}
	return gvaf.state
}

// Args contains the configurations for google_vertex_ai_featurestore.
type Args struct {
	// ForceDestroy: bool, optional
	ForceDestroy terra.BoolValue `hcl:"force_destroy,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// EncryptionSpec: optional
	EncryptionSpec *EncryptionSpec `hcl:"encryption_spec,block"`
	// OnlineServingConfig: optional
	OnlineServingConfig *OnlineServingConfig `hcl:"online_serving_config,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleVertexAiFeaturestoreAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_vertex_ai_featurestore.
func (gvaf googleVertexAiFeaturestoreAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(gvaf.ref.Append("create_time"))
}

// EffectiveLabels returns a reference to field effective_labels of google_vertex_ai_featurestore.
func (gvaf googleVertexAiFeaturestoreAttributes) EffectiveLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gvaf.ref.Append("effective_labels"))
}

// Etag returns a reference to field etag of google_vertex_ai_featurestore.
func (gvaf googleVertexAiFeaturestoreAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gvaf.ref.Append("etag"))
}

// ForceDestroy returns a reference to field force_destroy of google_vertex_ai_featurestore.
func (gvaf googleVertexAiFeaturestoreAttributes) ForceDestroy() terra.BoolValue {
	return terra.ReferenceAsBool(gvaf.ref.Append("force_destroy"))
}

// Id returns a reference to field id of google_vertex_ai_featurestore.
func (gvaf googleVertexAiFeaturestoreAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gvaf.ref.Append("id"))
}

// Labels returns a reference to field labels of google_vertex_ai_featurestore.
func (gvaf googleVertexAiFeaturestoreAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gvaf.ref.Append("labels"))
}

// Name returns a reference to field name of google_vertex_ai_featurestore.
func (gvaf googleVertexAiFeaturestoreAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gvaf.ref.Append("name"))
}

// Project returns a reference to field project of google_vertex_ai_featurestore.
func (gvaf googleVertexAiFeaturestoreAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gvaf.ref.Append("project"))
}

// Region returns a reference to field region of google_vertex_ai_featurestore.
func (gvaf googleVertexAiFeaturestoreAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(gvaf.ref.Append("region"))
}

// TerraformLabels returns a reference to field terraform_labels of google_vertex_ai_featurestore.
func (gvaf googleVertexAiFeaturestoreAttributes) TerraformLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gvaf.ref.Append("terraform_labels"))
}

// UpdateTime returns a reference to field update_time of google_vertex_ai_featurestore.
func (gvaf googleVertexAiFeaturestoreAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(gvaf.ref.Append("update_time"))
}

func (gvaf googleVertexAiFeaturestoreAttributes) EncryptionSpec() terra.ListValue[EncryptionSpecAttributes] {
	return terra.ReferenceAsList[EncryptionSpecAttributes](gvaf.ref.Append("encryption_spec"))
}

func (gvaf googleVertexAiFeaturestoreAttributes) OnlineServingConfig() terra.ListValue[OnlineServingConfigAttributes] {
	return terra.ReferenceAsList[OnlineServingConfigAttributes](gvaf.ref.Append("online_serving_config"))
}

func (gvaf googleVertexAiFeaturestoreAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gvaf.ref.Append("timeouts"))
}

type googleVertexAiFeaturestoreState struct {
	CreateTime          string                     `json:"create_time"`
	EffectiveLabels     map[string]string          `json:"effective_labels"`
	Etag                string                     `json:"etag"`
	ForceDestroy        bool                       `json:"force_destroy"`
	Id                  string                     `json:"id"`
	Labels              map[string]string          `json:"labels"`
	Name                string                     `json:"name"`
	Project             string                     `json:"project"`
	Region              string                     `json:"region"`
	TerraformLabels     map[string]string          `json:"terraform_labels"`
	UpdateTime          string                     `json:"update_time"`
	EncryptionSpec      []EncryptionSpecState      `json:"encryption_spec"`
	OnlineServingConfig []OnlineServingConfigState `json:"online_serving_config"`
	Timeouts            *TimeoutsState             `json:"timeouts"`
}
