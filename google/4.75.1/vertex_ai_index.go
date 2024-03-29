// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	vertexaiindex "github.com/golingon/terraproviders/google/4.75.1/vertexaiindex"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVertexAiIndex creates a new instance of [VertexAiIndex].
func NewVertexAiIndex(name string, args VertexAiIndexArgs) *VertexAiIndex {
	return &VertexAiIndex{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VertexAiIndex)(nil)

// VertexAiIndex represents the Terraform resource google_vertex_ai_index.
type VertexAiIndex struct {
	Name      string
	Args      VertexAiIndexArgs
	state     *vertexAiIndexState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VertexAiIndex].
func (vai *VertexAiIndex) Type() string {
	return "google_vertex_ai_index"
}

// LocalName returns the local name for [VertexAiIndex].
func (vai *VertexAiIndex) LocalName() string {
	return vai.Name
}

// Configuration returns the configuration (args) for [VertexAiIndex].
func (vai *VertexAiIndex) Configuration() interface{} {
	return vai.Args
}

// DependOn is used for other resources to depend on [VertexAiIndex].
func (vai *VertexAiIndex) DependOn() terra.Reference {
	return terra.ReferenceResource(vai)
}

// Dependencies returns the list of resources [VertexAiIndex] depends_on.
func (vai *VertexAiIndex) Dependencies() terra.Dependencies {
	return vai.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VertexAiIndex].
func (vai *VertexAiIndex) LifecycleManagement() *terra.Lifecycle {
	return vai.Lifecycle
}

// Attributes returns the attributes for [VertexAiIndex].
func (vai *VertexAiIndex) Attributes() vertexAiIndexAttributes {
	return vertexAiIndexAttributes{ref: terra.ReferenceResource(vai)}
}

// ImportState imports the given attribute values into [VertexAiIndex]'s state.
func (vai *VertexAiIndex) ImportState(av io.Reader) error {
	vai.state = &vertexAiIndexState{}
	if err := json.NewDecoder(av).Decode(vai.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vai.Type(), vai.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VertexAiIndex] has state.
func (vai *VertexAiIndex) State() (*vertexAiIndexState, bool) {
	return vai.state, vai.state != nil
}

// StateMust returns the state for [VertexAiIndex]. Panics if the state is nil.
func (vai *VertexAiIndex) StateMust() *vertexAiIndexState {
	if vai.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vai.Type(), vai.LocalName()))
	}
	return vai.state
}

// VertexAiIndexArgs contains the configurations for google_vertex_ai_index.
type VertexAiIndexArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IndexUpdateMethod: string, optional
	IndexUpdateMethod terra.StringValue `hcl:"index_update_method,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// DeployedIndexes: min=0
	DeployedIndexes []vertexaiindex.DeployedIndexes `hcl:"deployed_indexes,block" validate:"min=0"`
	// IndexStats: min=0
	IndexStats []vertexaiindex.IndexStats `hcl:"index_stats,block" validate:"min=0"`
	// Metadata: optional
	Metadata *vertexaiindex.Metadata `hcl:"metadata,block"`
	// Timeouts: optional
	Timeouts *vertexaiindex.Timeouts `hcl:"timeouts,block"`
}
type vertexAiIndexAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_vertex_ai_index.
func (vai vertexAiIndexAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(vai.ref.Append("create_time"))
}

// Description returns a reference to field description of google_vertex_ai_index.
func (vai vertexAiIndexAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(vai.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of google_vertex_ai_index.
func (vai vertexAiIndexAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(vai.ref.Append("display_name"))
}

// Etag returns a reference to field etag of google_vertex_ai_index.
func (vai vertexAiIndexAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(vai.ref.Append("etag"))
}

// Id returns a reference to field id of google_vertex_ai_index.
func (vai vertexAiIndexAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vai.ref.Append("id"))
}

// IndexUpdateMethod returns a reference to field index_update_method of google_vertex_ai_index.
func (vai vertexAiIndexAttributes) IndexUpdateMethod() terra.StringValue {
	return terra.ReferenceAsString(vai.ref.Append("index_update_method"))
}

// Labels returns a reference to field labels of google_vertex_ai_index.
func (vai vertexAiIndexAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vai.ref.Append("labels"))
}

// MetadataSchemaUri returns a reference to field metadata_schema_uri of google_vertex_ai_index.
func (vai vertexAiIndexAttributes) MetadataSchemaUri() terra.StringValue {
	return terra.ReferenceAsString(vai.ref.Append("metadata_schema_uri"))
}

// Name returns a reference to field name of google_vertex_ai_index.
func (vai vertexAiIndexAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vai.ref.Append("name"))
}

// Project returns a reference to field project of google_vertex_ai_index.
func (vai vertexAiIndexAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(vai.ref.Append("project"))
}

// Region returns a reference to field region of google_vertex_ai_index.
func (vai vertexAiIndexAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(vai.ref.Append("region"))
}

// UpdateTime returns a reference to field update_time of google_vertex_ai_index.
func (vai vertexAiIndexAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(vai.ref.Append("update_time"))
}

func (vai vertexAiIndexAttributes) DeployedIndexes() terra.ListValue[vertexaiindex.DeployedIndexesAttributes] {
	return terra.ReferenceAsList[vertexaiindex.DeployedIndexesAttributes](vai.ref.Append("deployed_indexes"))
}

func (vai vertexAiIndexAttributes) IndexStats() terra.ListValue[vertexaiindex.IndexStatsAttributes] {
	return terra.ReferenceAsList[vertexaiindex.IndexStatsAttributes](vai.ref.Append("index_stats"))
}

func (vai vertexAiIndexAttributes) Metadata() terra.ListValue[vertexaiindex.MetadataAttributes] {
	return terra.ReferenceAsList[vertexaiindex.MetadataAttributes](vai.ref.Append("metadata"))
}

func (vai vertexAiIndexAttributes) Timeouts() vertexaiindex.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vertexaiindex.TimeoutsAttributes](vai.ref.Append("timeouts"))
}

type vertexAiIndexState struct {
	CreateTime        string                               `json:"create_time"`
	Description       string                               `json:"description"`
	DisplayName       string                               `json:"display_name"`
	Etag              string                               `json:"etag"`
	Id                string                               `json:"id"`
	IndexUpdateMethod string                               `json:"index_update_method"`
	Labels            map[string]string                    `json:"labels"`
	MetadataSchemaUri string                               `json:"metadata_schema_uri"`
	Name              string                               `json:"name"`
	Project           string                               `json:"project"`
	Region            string                               `json:"region"`
	UpdateTime        string                               `json:"update_time"`
	DeployedIndexes   []vertexaiindex.DeployedIndexesState `json:"deployed_indexes"`
	IndexStats        []vertexaiindex.IndexStatsState      `json:"index_stats"`
	Metadata          []vertexaiindex.MetadataState        `json:"metadata"`
	Timeouts          *vertexaiindex.TimeoutsState         `json:"timeouts"`
}
