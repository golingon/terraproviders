// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	datavertexaiindex "github.com/golingon/terraproviders/googlebeta/4.76.0/datavertexaiindex"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataVertexAiIndex creates a new instance of [DataVertexAiIndex].
func NewDataVertexAiIndex(name string, args DataVertexAiIndexArgs) *DataVertexAiIndex {
	return &DataVertexAiIndex{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataVertexAiIndex)(nil)

// DataVertexAiIndex represents the Terraform data resource google_vertex_ai_index.
type DataVertexAiIndex struct {
	Name string
	Args DataVertexAiIndexArgs
}

// DataSource returns the Terraform object type for [DataVertexAiIndex].
func (vai *DataVertexAiIndex) DataSource() string {
	return "google_vertex_ai_index"
}

// LocalName returns the local name for [DataVertexAiIndex].
func (vai *DataVertexAiIndex) LocalName() string {
	return vai.Name
}

// Configuration returns the configuration (args) for [DataVertexAiIndex].
func (vai *DataVertexAiIndex) Configuration() interface{} {
	return vai.Args
}

// Attributes returns the attributes for [DataVertexAiIndex].
func (vai *DataVertexAiIndex) Attributes() dataVertexAiIndexAttributes {
	return dataVertexAiIndexAttributes{ref: terra.ReferenceDataResource(vai)}
}

// DataVertexAiIndexArgs contains the configurations for google_vertex_ai_index.
type DataVertexAiIndexArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, required
	Region terra.StringValue `hcl:"region,attr" validate:"required"`
	// DeployedIndexes: min=0
	DeployedIndexes []datavertexaiindex.DeployedIndexes `hcl:"deployed_indexes,block" validate:"min=0"`
	// IndexStats: min=0
	IndexStats []datavertexaiindex.IndexStats `hcl:"index_stats,block" validate:"min=0"`
	// Metadata: min=0
	Metadata []datavertexaiindex.Metadata `hcl:"metadata,block" validate:"min=0"`
}
type dataVertexAiIndexAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_vertex_ai_index.
func (vai dataVertexAiIndexAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(vai.ref.Append("create_time"))
}

// Description returns a reference to field description of google_vertex_ai_index.
func (vai dataVertexAiIndexAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(vai.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of google_vertex_ai_index.
func (vai dataVertexAiIndexAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(vai.ref.Append("display_name"))
}

// Etag returns a reference to field etag of google_vertex_ai_index.
func (vai dataVertexAiIndexAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(vai.ref.Append("etag"))
}

// Id returns a reference to field id of google_vertex_ai_index.
func (vai dataVertexAiIndexAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vai.ref.Append("id"))
}

// IndexUpdateMethod returns a reference to field index_update_method of google_vertex_ai_index.
func (vai dataVertexAiIndexAttributes) IndexUpdateMethod() terra.StringValue {
	return terra.ReferenceAsString(vai.ref.Append("index_update_method"))
}

// Labels returns a reference to field labels of google_vertex_ai_index.
func (vai dataVertexAiIndexAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vai.ref.Append("labels"))
}

// MetadataSchemaUri returns a reference to field metadata_schema_uri of google_vertex_ai_index.
func (vai dataVertexAiIndexAttributes) MetadataSchemaUri() terra.StringValue {
	return terra.ReferenceAsString(vai.ref.Append("metadata_schema_uri"))
}

// Name returns a reference to field name of google_vertex_ai_index.
func (vai dataVertexAiIndexAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vai.ref.Append("name"))
}

// Project returns a reference to field project of google_vertex_ai_index.
func (vai dataVertexAiIndexAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(vai.ref.Append("project"))
}

// Region returns a reference to field region of google_vertex_ai_index.
func (vai dataVertexAiIndexAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(vai.ref.Append("region"))
}

// UpdateTime returns a reference to field update_time of google_vertex_ai_index.
func (vai dataVertexAiIndexAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(vai.ref.Append("update_time"))
}

func (vai dataVertexAiIndexAttributes) DeployedIndexes() terra.ListValue[datavertexaiindex.DeployedIndexesAttributes] {
	return terra.ReferenceAsList[datavertexaiindex.DeployedIndexesAttributes](vai.ref.Append("deployed_indexes"))
}

func (vai dataVertexAiIndexAttributes) IndexStats() terra.ListValue[datavertexaiindex.IndexStatsAttributes] {
	return terra.ReferenceAsList[datavertexaiindex.IndexStatsAttributes](vai.ref.Append("index_stats"))
}

func (vai dataVertexAiIndexAttributes) Metadata() terra.ListValue[datavertexaiindex.MetadataAttributes] {
	return terra.ReferenceAsList[datavertexaiindex.MetadataAttributes](vai.ref.Append("metadata"))
}
