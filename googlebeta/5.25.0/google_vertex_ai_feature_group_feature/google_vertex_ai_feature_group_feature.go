// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_vertex_ai_feature_group_feature

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

// Resource represents the Terraform resource google_vertex_ai_feature_group_feature.
type Resource struct {
	Name      string
	Args      Args
	state     *googleVertexAiFeatureGroupFeatureState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gvafgf *Resource) Type() string {
	return "google_vertex_ai_feature_group_feature"
}

// LocalName returns the local name for [Resource].
func (gvafgf *Resource) LocalName() string {
	return gvafgf.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gvafgf *Resource) Configuration() interface{} {
	return gvafgf.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gvafgf *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gvafgf)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gvafgf *Resource) Dependencies() terra.Dependencies {
	return gvafgf.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gvafgf *Resource) LifecycleManagement() *terra.Lifecycle {
	return gvafgf.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gvafgf *Resource) Attributes() googleVertexAiFeatureGroupFeatureAttributes {
	return googleVertexAiFeatureGroupFeatureAttributes{ref: terra.ReferenceResource(gvafgf)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gvafgf *Resource) ImportState(state io.Reader) error {
	gvafgf.state = &googleVertexAiFeatureGroupFeatureState{}
	if err := json.NewDecoder(state).Decode(gvafgf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gvafgf.Type(), gvafgf.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gvafgf *Resource) State() (*googleVertexAiFeatureGroupFeatureState, bool) {
	return gvafgf.state, gvafgf.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gvafgf *Resource) StateMust() *googleVertexAiFeatureGroupFeatureState {
	if gvafgf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gvafgf.Type(), gvafgf.LocalName()))
	}
	return gvafgf.state
}

// Args contains the configurations for google_vertex_ai_feature_group_feature.
type Args struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// FeatureGroup: string, required
	FeatureGroup terra.StringValue `hcl:"feature_group,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, required
	Region terra.StringValue `hcl:"region,attr" validate:"required"`
	// VersionColumnName: string, optional
	VersionColumnName terra.StringValue `hcl:"version_column_name,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleVertexAiFeatureGroupFeatureAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_vertex_ai_feature_group_feature.
func (gvafgf googleVertexAiFeatureGroupFeatureAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(gvafgf.ref.Append("create_time"))
}

// Description returns a reference to field description of google_vertex_ai_feature_group_feature.
func (gvafgf googleVertexAiFeatureGroupFeatureAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gvafgf.ref.Append("description"))
}

// EffectiveLabels returns a reference to field effective_labels of google_vertex_ai_feature_group_feature.
func (gvafgf googleVertexAiFeatureGroupFeatureAttributes) EffectiveLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gvafgf.ref.Append("effective_labels"))
}

// FeatureGroup returns a reference to field feature_group of google_vertex_ai_feature_group_feature.
func (gvafgf googleVertexAiFeatureGroupFeatureAttributes) FeatureGroup() terra.StringValue {
	return terra.ReferenceAsString(gvafgf.ref.Append("feature_group"))
}

// Id returns a reference to field id of google_vertex_ai_feature_group_feature.
func (gvafgf googleVertexAiFeatureGroupFeatureAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gvafgf.ref.Append("id"))
}

// Labels returns a reference to field labels of google_vertex_ai_feature_group_feature.
func (gvafgf googleVertexAiFeatureGroupFeatureAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gvafgf.ref.Append("labels"))
}

// Name returns a reference to field name of google_vertex_ai_feature_group_feature.
func (gvafgf googleVertexAiFeatureGroupFeatureAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gvafgf.ref.Append("name"))
}

// Project returns a reference to field project of google_vertex_ai_feature_group_feature.
func (gvafgf googleVertexAiFeatureGroupFeatureAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gvafgf.ref.Append("project"))
}

// Region returns a reference to field region of google_vertex_ai_feature_group_feature.
func (gvafgf googleVertexAiFeatureGroupFeatureAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(gvafgf.ref.Append("region"))
}

// TerraformLabels returns a reference to field terraform_labels of google_vertex_ai_feature_group_feature.
func (gvafgf googleVertexAiFeatureGroupFeatureAttributes) TerraformLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gvafgf.ref.Append("terraform_labels"))
}

// UpdateTime returns a reference to field update_time of google_vertex_ai_feature_group_feature.
func (gvafgf googleVertexAiFeatureGroupFeatureAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(gvafgf.ref.Append("update_time"))
}

// VersionColumnName returns a reference to field version_column_name of google_vertex_ai_feature_group_feature.
func (gvafgf googleVertexAiFeatureGroupFeatureAttributes) VersionColumnName() terra.StringValue {
	return terra.ReferenceAsString(gvafgf.ref.Append("version_column_name"))
}

func (gvafgf googleVertexAiFeatureGroupFeatureAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gvafgf.ref.Append("timeouts"))
}

type googleVertexAiFeatureGroupFeatureState struct {
	CreateTime        string            `json:"create_time"`
	Description       string            `json:"description"`
	EffectiveLabels   map[string]string `json:"effective_labels"`
	FeatureGroup      string            `json:"feature_group"`
	Id                string            `json:"id"`
	Labels            map[string]string `json:"labels"`
	Name              string            `json:"name"`
	Project           string            `json:"project"`
	Region            string            `json:"region"`
	TerraformLabels   map[string]string `json:"terraform_labels"`
	UpdateTime        string            `json:"update_time"`
	VersionColumnName string            `json:"version_column_name"`
	Timeouts          *TimeoutsState    `json:"timeouts"`
}
