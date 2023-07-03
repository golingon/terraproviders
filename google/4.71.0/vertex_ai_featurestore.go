// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	vertexaifeaturestore "github.com/golingon/terraproviders/google/4.71.0/vertexaifeaturestore"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVertexAiFeaturestore creates a new instance of [VertexAiFeaturestore].
func NewVertexAiFeaturestore(name string, args VertexAiFeaturestoreArgs) *VertexAiFeaturestore {
	return &VertexAiFeaturestore{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VertexAiFeaturestore)(nil)

// VertexAiFeaturestore represents the Terraform resource google_vertex_ai_featurestore.
type VertexAiFeaturestore struct {
	Name      string
	Args      VertexAiFeaturestoreArgs
	state     *vertexAiFeaturestoreState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VertexAiFeaturestore].
func (vaf *VertexAiFeaturestore) Type() string {
	return "google_vertex_ai_featurestore"
}

// LocalName returns the local name for [VertexAiFeaturestore].
func (vaf *VertexAiFeaturestore) LocalName() string {
	return vaf.Name
}

// Configuration returns the configuration (args) for [VertexAiFeaturestore].
func (vaf *VertexAiFeaturestore) Configuration() interface{} {
	return vaf.Args
}

// DependOn is used for other resources to depend on [VertexAiFeaturestore].
func (vaf *VertexAiFeaturestore) DependOn() terra.Reference {
	return terra.ReferenceResource(vaf)
}

// Dependencies returns the list of resources [VertexAiFeaturestore] depends_on.
func (vaf *VertexAiFeaturestore) Dependencies() terra.Dependencies {
	return vaf.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VertexAiFeaturestore].
func (vaf *VertexAiFeaturestore) LifecycleManagement() *terra.Lifecycle {
	return vaf.Lifecycle
}

// Attributes returns the attributes for [VertexAiFeaturestore].
func (vaf *VertexAiFeaturestore) Attributes() vertexAiFeaturestoreAttributes {
	return vertexAiFeaturestoreAttributes{ref: terra.ReferenceResource(vaf)}
}

// ImportState imports the given attribute values into [VertexAiFeaturestore]'s state.
func (vaf *VertexAiFeaturestore) ImportState(av io.Reader) error {
	vaf.state = &vertexAiFeaturestoreState{}
	if err := json.NewDecoder(av).Decode(vaf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vaf.Type(), vaf.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VertexAiFeaturestore] has state.
func (vaf *VertexAiFeaturestore) State() (*vertexAiFeaturestoreState, bool) {
	return vaf.state, vaf.state != nil
}

// StateMust returns the state for [VertexAiFeaturestore]. Panics if the state is nil.
func (vaf *VertexAiFeaturestore) StateMust() *vertexAiFeaturestoreState {
	if vaf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vaf.Type(), vaf.LocalName()))
	}
	return vaf.state
}

// VertexAiFeaturestoreArgs contains the configurations for google_vertex_ai_featurestore.
type VertexAiFeaturestoreArgs struct {
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
	EncryptionSpec *vertexaifeaturestore.EncryptionSpec `hcl:"encryption_spec,block"`
	// OnlineServingConfig: optional
	OnlineServingConfig *vertexaifeaturestore.OnlineServingConfig `hcl:"online_serving_config,block"`
	// Timeouts: optional
	Timeouts *vertexaifeaturestore.Timeouts `hcl:"timeouts,block"`
}
type vertexAiFeaturestoreAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_vertex_ai_featurestore.
func (vaf vertexAiFeaturestoreAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(vaf.ref.Append("create_time"))
}

// Etag returns a reference to field etag of google_vertex_ai_featurestore.
func (vaf vertexAiFeaturestoreAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(vaf.ref.Append("etag"))
}

// ForceDestroy returns a reference to field force_destroy of google_vertex_ai_featurestore.
func (vaf vertexAiFeaturestoreAttributes) ForceDestroy() terra.BoolValue {
	return terra.ReferenceAsBool(vaf.ref.Append("force_destroy"))
}

// Id returns a reference to field id of google_vertex_ai_featurestore.
func (vaf vertexAiFeaturestoreAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vaf.ref.Append("id"))
}

// Labels returns a reference to field labels of google_vertex_ai_featurestore.
func (vaf vertexAiFeaturestoreAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vaf.ref.Append("labels"))
}

// Name returns a reference to field name of google_vertex_ai_featurestore.
func (vaf vertexAiFeaturestoreAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vaf.ref.Append("name"))
}

// Project returns a reference to field project of google_vertex_ai_featurestore.
func (vaf vertexAiFeaturestoreAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(vaf.ref.Append("project"))
}

// Region returns a reference to field region of google_vertex_ai_featurestore.
func (vaf vertexAiFeaturestoreAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(vaf.ref.Append("region"))
}

// UpdateTime returns a reference to field update_time of google_vertex_ai_featurestore.
func (vaf vertexAiFeaturestoreAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(vaf.ref.Append("update_time"))
}

func (vaf vertexAiFeaturestoreAttributes) EncryptionSpec() terra.ListValue[vertexaifeaturestore.EncryptionSpecAttributes] {
	return terra.ReferenceAsList[vertexaifeaturestore.EncryptionSpecAttributes](vaf.ref.Append("encryption_spec"))
}

func (vaf vertexAiFeaturestoreAttributes) OnlineServingConfig() terra.ListValue[vertexaifeaturestore.OnlineServingConfigAttributes] {
	return terra.ReferenceAsList[vertexaifeaturestore.OnlineServingConfigAttributes](vaf.ref.Append("online_serving_config"))
}

func (vaf vertexAiFeaturestoreAttributes) Timeouts() vertexaifeaturestore.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vertexaifeaturestore.TimeoutsAttributes](vaf.ref.Append("timeouts"))
}

type vertexAiFeaturestoreState struct {
	CreateTime          string                                          `json:"create_time"`
	Etag                string                                          `json:"etag"`
	ForceDestroy        bool                                            `json:"force_destroy"`
	Id                  string                                          `json:"id"`
	Labels              map[string]string                               `json:"labels"`
	Name                string                                          `json:"name"`
	Project             string                                          `json:"project"`
	Region              string                                          `json:"region"`
	UpdateTime          string                                          `json:"update_time"`
	EncryptionSpec      []vertexaifeaturestore.EncryptionSpecState      `json:"encryption_spec"`
	OnlineServingConfig []vertexaifeaturestore.OnlineServingConfigState `json:"online_serving_config"`
	Timeouts            *vertexaifeaturestore.TimeoutsState             `json:"timeouts"`
}
