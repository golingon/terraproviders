// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_vertex_ai_featurestore_entitytype_feature

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

// Resource represents the Terraform resource google_vertex_ai_featurestore_entitytype_feature.
type Resource struct {
	Name      string
	Args      Args
	state     *googleVertexAiFeaturestoreEntitytypeFeatureState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gvafef *Resource) Type() string {
	return "google_vertex_ai_featurestore_entitytype_feature"
}

// LocalName returns the local name for [Resource].
func (gvafef *Resource) LocalName() string {
	return gvafef.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gvafef *Resource) Configuration() interface{} {
	return gvafef.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gvafef *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gvafef)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gvafef *Resource) Dependencies() terra.Dependencies {
	return gvafef.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gvafef *Resource) LifecycleManagement() *terra.Lifecycle {
	return gvafef.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gvafef *Resource) Attributes() googleVertexAiFeaturestoreEntitytypeFeatureAttributes {
	return googleVertexAiFeaturestoreEntitytypeFeatureAttributes{ref: terra.ReferenceResource(gvafef)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gvafef *Resource) ImportState(state io.Reader) error {
	gvafef.state = &googleVertexAiFeaturestoreEntitytypeFeatureState{}
	if err := json.NewDecoder(state).Decode(gvafef.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gvafef.Type(), gvafef.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gvafef *Resource) State() (*googleVertexAiFeaturestoreEntitytypeFeatureState, bool) {
	return gvafef.state, gvafef.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gvafef *Resource) StateMust() *googleVertexAiFeaturestoreEntitytypeFeatureState {
	if gvafef.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gvafef.Type(), gvafef.LocalName()))
	}
	return gvafef.state
}

// Args contains the configurations for google_vertex_ai_featurestore_entitytype_feature.
type Args struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Entitytype: string, required
	Entitytype terra.StringValue `hcl:"entitytype,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// ValueType: string, required
	ValueType terra.StringValue `hcl:"value_type,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleVertexAiFeaturestoreEntitytypeFeatureAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_vertex_ai_featurestore_entitytype_feature.
func (gvafef googleVertexAiFeaturestoreEntitytypeFeatureAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(gvafef.ref.Append("create_time"))
}

// Description returns a reference to field description of google_vertex_ai_featurestore_entitytype_feature.
func (gvafef googleVertexAiFeaturestoreEntitytypeFeatureAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gvafef.ref.Append("description"))
}

// EffectiveLabels returns a reference to field effective_labels of google_vertex_ai_featurestore_entitytype_feature.
func (gvafef googleVertexAiFeaturestoreEntitytypeFeatureAttributes) EffectiveLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gvafef.ref.Append("effective_labels"))
}

// Entitytype returns a reference to field entitytype of google_vertex_ai_featurestore_entitytype_feature.
func (gvafef googleVertexAiFeaturestoreEntitytypeFeatureAttributes) Entitytype() terra.StringValue {
	return terra.ReferenceAsString(gvafef.ref.Append("entitytype"))
}

// Etag returns a reference to field etag of google_vertex_ai_featurestore_entitytype_feature.
func (gvafef googleVertexAiFeaturestoreEntitytypeFeatureAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gvafef.ref.Append("etag"))
}

// Id returns a reference to field id of google_vertex_ai_featurestore_entitytype_feature.
func (gvafef googleVertexAiFeaturestoreEntitytypeFeatureAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gvafef.ref.Append("id"))
}

// Labels returns a reference to field labels of google_vertex_ai_featurestore_entitytype_feature.
func (gvafef googleVertexAiFeaturestoreEntitytypeFeatureAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gvafef.ref.Append("labels"))
}

// Name returns a reference to field name of google_vertex_ai_featurestore_entitytype_feature.
func (gvafef googleVertexAiFeaturestoreEntitytypeFeatureAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gvafef.ref.Append("name"))
}

// Region returns a reference to field region of google_vertex_ai_featurestore_entitytype_feature.
func (gvafef googleVertexAiFeaturestoreEntitytypeFeatureAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(gvafef.ref.Append("region"))
}

// TerraformLabels returns a reference to field terraform_labels of google_vertex_ai_featurestore_entitytype_feature.
func (gvafef googleVertexAiFeaturestoreEntitytypeFeatureAttributes) TerraformLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gvafef.ref.Append("terraform_labels"))
}

// UpdateTime returns a reference to field update_time of google_vertex_ai_featurestore_entitytype_feature.
func (gvafef googleVertexAiFeaturestoreEntitytypeFeatureAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(gvafef.ref.Append("update_time"))
}

// ValueType returns a reference to field value_type of google_vertex_ai_featurestore_entitytype_feature.
func (gvafef googleVertexAiFeaturestoreEntitytypeFeatureAttributes) ValueType() terra.StringValue {
	return terra.ReferenceAsString(gvafef.ref.Append("value_type"))
}

func (gvafef googleVertexAiFeaturestoreEntitytypeFeatureAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gvafef.ref.Append("timeouts"))
}

type googleVertexAiFeaturestoreEntitytypeFeatureState struct {
	CreateTime      string            `json:"create_time"`
	Description     string            `json:"description"`
	EffectiveLabels map[string]string `json:"effective_labels"`
	Entitytype      string            `json:"entitytype"`
	Etag            string            `json:"etag"`
	Id              string            `json:"id"`
	Labels          map[string]string `json:"labels"`
	Name            string            `json:"name"`
	Region          string            `json:"region"`
	TerraformLabels map[string]string `json:"terraform_labels"`
	UpdateTime      string            `json:"update_time"`
	ValueType       string            `json:"value_type"`
	Timeouts        *TimeoutsState    `json:"timeouts"`
}
