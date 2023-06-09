// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	vertexaimetadatastore "github.com/golingon/terraproviders/googlebeta/4.72.1/vertexaimetadatastore"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVertexAiMetadataStore creates a new instance of [VertexAiMetadataStore].
func NewVertexAiMetadataStore(name string, args VertexAiMetadataStoreArgs) *VertexAiMetadataStore {
	return &VertexAiMetadataStore{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VertexAiMetadataStore)(nil)

// VertexAiMetadataStore represents the Terraform resource google_vertex_ai_metadata_store.
type VertexAiMetadataStore struct {
	Name      string
	Args      VertexAiMetadataStoreArgs
	state     *vertexAiMetadataStoreState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VertexAiMetadataStore].
func (vams *VertexAiMetadataStore) Type() string {
	return "google_vertex_ai_metadata_store"
}

// LocalName returns the local name for [VertexAiMetadataStore].
func (vams *VertexAiMetadataStore) LocalName() string {
	return vams.Name
}

// Configuration returns the configuration (args) for [VertexAiMetadataStore].
func (vams *VertexAiMetadataStore) Configuration() interface{} {
	return vams.Args
}

// DependOn is used for other resources to depend on [VertexAiMetadataStore].
func (vams *VertexAiMetadataStore) DependOn() terra.Reference {
	return terra.ReferenceResource(vams)
}

// Dependencies returns the list of resources [VertexAiMetadataStore] depends_on.
func (vams *VertexAiMetadataStore) Dependencies() terra.Dependencies {
	return vams.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VertexAiMetadataStore].
func (vams *VertexAiMetadataStore) LifecycleManagement() *terra.Lifecycle {
	return vams.Lifecycle
}

// Attributes returns the attributes for [VertexAiMetadataStore].
func (vams *VertexAiMetadataStore) Attributes() vertexAiMetadataStoreAttributes {
	return vertexAiMetadataStoreAttributes{ref: terra.ReferenceResource(vams)}
}

// ImportState imports the given attribute values into [VertexAiMetadataStore]'s state.
func (vams *VertexAiMetadataStore) ImportState(av io.Reader) error {
	vams.state = &vertexAiMetadataStoreState{}
	if err := json.NewDecoder(av).Decode(vams.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vams.Type(), vams.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VertexAiMetadataStore] has state.
func (vams *VertexAiMetadataStore) State() (*vertexAiMetadataStoreState, bool) {
	return vams.state, vams.state != nil
}

// StateMust returns the state for [VertexAiMetadataStore]. Panics if the state is nil.
func (vams *VertexAiMetadataStore) StateMust() *vertexAiMetadataStoreState {
	if vams.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vams.Type(), vams.LocalName()))
	}
	return vams.state
}

// VertexAiMetadataStoreArgs contains the configurations for google_vertex_ai_metadata_store.
type VertexAiMetadataStoreArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// State: min=0
	State []vertexaimetadatastore.State `hcl:"state,block" validate:"min=0"`
	// EncryptionSpec: optional
	EncryptionSpec *vertexaimetadatastore.EncryptionSpec `hcl:"encryption_spec,block"`
	// Timeouts: optional
	Timeouts *vertexaimetadatastore.Timeouts `hcl:"timeouts,block"`
}
type vertexAiMetadataStoreAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_vertex_ai_metadata_store.
func (vams vertexAiMetadataStoreAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(vams.ref.Append("create_time"))
}

// Description returns a reference to field description of google_vertex_ai_metadata_store.
func (vams vertexAiMetadataStoreAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(vams.ref.Append("description"))
}

// Id returns a reference to field id of google_vertex_ai_metadata_store.
func (vams vertexAiMetadataStoreAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vams.ref.Append("id"))
}

// Name returns a reference to field name of google_vertex_ai_metadata_store.
func (vams vertexAiMetadataStoreAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vams.ref.Append("name"))
}

// Project returns a reference to field project of google_vertex_ai_metadata_store.
func (vams vertexAiMetadataStoreAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(vams.ref.Append("project"))
}

// Region returns a reference to field region of google_vertex_ai_metadata_store.
func (vams vertexAiMetadataStoreAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(vams.ref.Append("region"))
}

// UpdateTime returns a reference to field update_time of google_vertex_ai_metadata_store.
func (vams vertexAiMetadataStoreAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(vams.ref.Append("update_time"))
}

func (vams vertexAiMetadataStoreAttributes) State() terra.ListValue[vertexaimetadatastore.StateAttributes] {
	return terra.ReferenceAsList[vertexaimetadatastore.StateAttributes](vams.ref.Append("state"))
}

func (vams vertexAiMetadataStoreAttributes) EncryptionSpec() terra.ListValue[vertexaimetadatastore.EncryptionSpecAttributes] {
	return terra.ReferenceAsList[vertexaimetadatastore.EncryptionSpecAttributes](vams.ref.Append("encryption_spec"))
}

func (vams vertexAiMetadataStoreAttributes) Timeouts() vertexaimetadatastore.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vertexaimetadatastore.TimeoutsAttributes](vams.ref.Append("timeouts"))
}

type vertexAiMetadataStoreState struct {
	CreateTime     string                                      `json:"create_time"`
	Description    string                                      `json:"description"`
	Id             string                                      `json:"id"`
	Name           string                                      `json:"name"`
	Project        string                                      `json:"project"`
	Region         string                                      `json:"region"`
	UpdateTime     string                                      `json:"update_time"`
	State          []vertexaimetadatastore.StateState          `json:"state"`
	EncryptionSpec []vertexaimetadatastore.EncryptionSpecState `json:"encryption_spec"`
	Timeouts       *vertexaimetadatastore.TimeoutsState        `json:"timeouts"`
}
