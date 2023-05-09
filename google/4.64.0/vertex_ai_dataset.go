// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	vertexaidataset "github.com/golingon/terraproviders/google/4.64.0/vertexaidataset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVertexAiDataset creates a new instance of [VertexAiDataset].
func NewVertexAiDataset(name string, args VertexAiDatasetArgs) *VertexAiDataset {
	return &VertexAiDataset{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VertexAiDataset)(nil)

// VertexAiDataset represents the Terraform resource google_vertex_ai_dataset.
type VertexAiDataset struct {
	Name      string
	Args      VertexAiDatasetArgs
	state     *vertexAiDatasetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VertexAiDataset].
func (vad *VertexAiDataset) Type() string {
	return "google_vertex_ai_dataset"
}

// LocalName returns the local name for [VertexAiDataset].
func (vad *VertexAiDataset) LocalName() string {
	return vad.Name
}

// Configuration returns the configuration (args) for [VertexAiDataset].
func (vad *VertexAiDataset) Configuration() interface{} {
	return vad.Args
}

// DependOn is used for other resources to depend on [VertexAiDataset].
func (vad *VertexAiDataset) DependOn() terra.Reference {
	return terra.ReferenceResource(vad)
}

// Dependencies returns the list of resources [VertexAiDataset] depends_on.
func (vad *VertexAiDataset) Dependencies() terra.Dependencies {
	return vad.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VertexAiDataset].
func (vad *VertexAiDataset) LifecycleManagement() *terra.Lifecycle {
	return vad.Lifecycle
}

// Attributes returns the attributes for [VertexAiDataset].
func (vad *VertexAiDataset) Attributes() vertexAiDatasetAttributes {
	return vertexAiDatasetAttributes{ref: terra.ReferenceResource(vad)}
}

// ImportState imports the given attribute values into [VertexAiDataset]'s state.
func (vad *VertexAiDataset) ImportState(av io.Reader) error {
	vad.state = &vertexAiDatasetState{}
	if err := json.NewDecoder(av).Decode(vad.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vad.Type(), vad.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VertexAiDataset] has state.
func (vad *VertexAiDataset) State() (*vertexAiDatasetState, bool) {
	return vad.state, vad.state != nil
}

// StateMust returns the state for [VertexAiDataset]. Panics if the state is nil.
func (vad *VertexAiDataset) StateMust() *vertexAiDatasetState {
	if vad.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vad.Type(), vad.LocalName()))
	}
	return vad.state
}

// VertexAiDatasetArgs contains the configurations for google_vertex_ai_dataset.
type VertexAiDatasetArgs struct {
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// MetadataSchemaUri: string, required
	MetadataSchemaUri terra.StringValue `hcl:"metadata_schema_uri,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// EncryptionSpec: optional
	EncryptionSpec *vertexaidataset.EncryptionSpec `hcl:"encryption_spec,block"`
	// Timeouts: optional
	Timeouts *vertexaidataset.Timeouts `hcl:"timeouts,block"`
}
type vertexAiDatasetAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_vertex_ai_dataset.
func (vad vertexAiDatasetAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(vad.ref.Append("create_time"))
}

// DisplayName returns a reference to field display_name of google_vertex_ai_dataset.
func (vad vertexAiDatasetAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(vad.ref.Append("display_name"))
}

// Id returns a reference to field id of google_vertex_ai_dataset.
func (vad vertexAiDatasetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vad.ref.Append("id"))
}

// Labels returns a reference to field labels of google_vertex_ai_dataset.
func (vad vertexAiDatasetAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vad.ref.Append("labels"))
}

// MetadataSchemaUri returns a reference to field metadata_schema_uri of google_vertex_ai_dataset.
func (vad vertexAiDatasetAttributes) MetadataSchemaUri() terra.StringValue {
	return terra.ReferenceAsString(vad.ref.Append("metadata_schema_uri"))
}

// Name returns a reference to field name of google_vertex_ai_dataset.
func (vad vertexAiDatasetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vad.ref.Append("name"))
}

// Project returns a reference to field project of google_vertex_ai_dataset.
func (vad vertexAiDatasetAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(vad.ref.Append("project"))
}

// Region returns a reference to field region of google_vertex_ai_dataset.
func (vad vertexAiDatasetAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(vad.ref.Append("region"))
}

// UpdateTime returns a reference to field update_time of google_vertex_ai_dataset.
func (vad vertexAiDatasetAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(vad.ref.Append("update_time"))
}

func (vad vertexAiDatasetAttributes) EncryptionSpec() terra.ListValue[vertexaidataset.EncryptionSpecAttributes] {
	return terra.ReferenceAsList[vertexaidataset.EncryptionSpecAttributes](vad.ref.Append("encryption_spec"))
}

func (vad vertexAiDatasetAttributes) Timeouts() vertexaidataset.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vertexaidataset.TimeoutsAttributes](vad.ref.Append("timeouts"))
}

type vertexAiDatasetState struct {
	CreateTime        string                                `json:"create_time"`
	DisplayName       string                                `json:"display_name"`
	Id                string                                `json:"id"`
	Labels            map[string]string                     `json:"labels"`
	MetadataSchemaUri string                                `json:"metadata_schema_uri"`
	Name              string                                `json:"name"`
	Project           string                                `json:"project"`
	Region            string                                `json:"region"`
	UpdateTime        string                                `json:"update_time"`
	EncryptionSpec    []vertexaidataset.EncryptionSpecState `json:"encryption_spec"`
	Timeouts          *vertexaidataset.TimeoutsState        `json:"timeouts"`
}
