// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_vertex_ai_metadata_store

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

// Resource represents the Terraform resource google_vertex_ai_metadata_store.
type Resource struct {
	Name      string
	Args      Args
	state     *googleVertexAiMetadataStoreState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gvams *Resource) Type() string {
	return "google_vertex_ai_metadata_store"
}

// LocalName returns the local name for [Resource].
func (gvams *Resource) LocalName() string {
	return gvams.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gvams *Resource) Configuration() interface{} {
	return gvams.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gvams *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gvams)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gvams *Resource) Dependencies() terra.Dependencies {
	return gvams.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gvams *Resource) LifecycleManagement() *terra.Lifecycle {
	return gvams.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gvams *Resource) Attributes() googleVertexAiMetadataStoreAttributes {
	return googleVertexAiMetadataStoreAttributes{ref: terra.ReferenceResource(gvams)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gvams *Resource) ImportState(state io.Reader) error {
	gvams.state = &googleVertexAiMetadataStoreState{}
	if err := json.NewDecoder(state).Decode(gvams.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gvams.Type(), gvams.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gvams *Resource) State() (*googleVertexAiMetadataStoreState, bool) {
	return gvams.state, gvams.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gvams *Resource) StateMust() *googleVertexAiMetadataStoreState {
	if gvams.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gvams.Type(), gvams.LocalName()))
	}
	return gvams.state
}

// Args contains the configurations for google_vertex_ai_metadata_store.
type Args struct {
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
	// EncryptionSpec: optional
	EncryptionSpec *EncryptionSpec `hcl:"encryption_spec,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleVertexAiMetadataStoreAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_vertex_ai_metadata_store.
func (gvams googleVertexAiMetadataStoreAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(gvams.ref.Append("create_time"))
}

// Description returns a reference to field description of google_vertex_ai_metadata_store.
func (gvams googleVertexAiMetadataStoreAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gvams.ref.Append("description"))
}

// Id returns a reference to field id of google_vertex_ai_metadata_store.
func (gvams googleVertexAiMetadataStoreAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gvams.ref.Append("id"))
}

// Name returns a reference to field name of google_vertex_ai_metadata_store.
func (gvams googleVertexAiMetadataStoreAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gvams.ref.Append("name"))
}

// Project returns a reference to field project of google_vertex_ai_metadata_store.
func (gvams googleVertexAiMetadataStoreAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gvams.ref.Append("project"))
}

// Region returns a reference to field region of google_vertex_ai_metadata_store.
func (gvams googleVertexAiMetadataStoreAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(gvams.ref.Append("region"))
}

// UpdateTime returns a reference to field update_time of google_vertex_ai_metadata_store.
func (gvams googleVertexAiMetadataStoreAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(gvams.ref.Append("update_time"))
}

func (gvams googleVertexAiMetadataStoreAttributes) State() terra.ListValue[StateAttributes] {
	return terra.ReferenceAsList[StateAttributes](gvams.ref.Append("state"))
}

func (gvams googleVertexAiMetadataStoreAttributes) EncryptionSpec() terra.ListValue[EncryptionSpecAttributes] {
	return terra.ReferenceAsList[EncryptionSpecAttributes](gvams.ref.Append("encryption_spec"))
}

func (gvams googleVertexAiMetadataStoreAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gvams.ref.Append("timeouts"))
}

type googleVertexAiMetadataStoreState struct {
	CreateTime     string                `json:"create_time"`
	Description    string                `json:"description"`
	Id             string                `json:"id"`
	Name           string                `json:"name"`
	Project        string                `json:"project"`
	Region         string                `json:"region"`
	UpdateTime     string                `json:"update_time"`
	State          []StateState          `json:"state"`
	EncryptionSpec []EncryptionSpecState `json:"encryption_spec"`
	Timeouts       *TimeoutsState        `json:"timeouts"`
}
