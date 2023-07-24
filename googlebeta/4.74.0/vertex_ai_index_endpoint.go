// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	vertexaiindexendpoint "github.com/golingon/terraproviders/googlebeta/4.74.0/vertexaiindexendpoint"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVertexAiIndexEndpoint creates a new instance of [VertexAiIndexEndpoint].
func NewVertexAiIndexEndpoint(name string, args VertexAiIndexEndpointArgs) *VertexAiIndexEndpoint {
	return &VertexAiIndexEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VertexAiIndexEndpoint)(nil)

// VertexAiIndexEndpoint represents the Terraform resource google_vertex_ai_index_endpoint.
type VertexAiIndexEndpoint struct {
	Name      string
	Args      VertexAiIndexEndpointArgs
	state     *vertexAiIndexEndpointState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VertexAiIndexEndpoint].
func (vaie *VertexAiIndexEndpoint) Type() string {
	return "google_vertex_ai_index_endpoint"
}

// LocalName returns the local name for [VertexAiIndexEndpoint].
func (vaie *VertexAiIndexEndpoint) LocalName() string {
	return vaie.Name
}

// Configuration returns the configuration (args) for [VertexAiIndexEndpoint].
func (vaie *VertexAiIndexEndpoint) Configuration() interface{} {
	return vaie.Args
}

// DependOn is used for other resources to depend on [VertexAiIndexEndpoint].
func (vaie *VertexAiIndexEndpoint) DependOn() terra.Reference {
	return terra.ReferenceResource(vaie)
}

// Dependencies returns the list of resources [VertexAiIndexEndpoint] depends_on.
func (vaie *VertexAiIndexEndpoint) Dependencies() terra.Dependencies {
	return vaie.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VertexAiIndexEndpoint].
func (vaie *VertexAiIndexEndpoint) LifecycleManagement() *terra.Lifecycle {
	return vaie.Lifecycle
}

// Attributes returns the attributes for [VertexAiIndexEndpoint].
func (vaie *VertexAiIndexEndpoint) Attributes() vertexAiIndexEndpointAttributes {
	return vertexAiIndexEndpointAttributes{ref: terra.ReferenceResource(vaie)}
}

// ImportState imports the given attribute values into [VertexAiIndexEndpoint]'s state.
func (vaie *VertexAiIndexEndpoint) ImportState(av io.Reader) error {
	vaie.state = &vertexAiIndexEndpointState{}
	if err := json.NewDecoder(av).Decode(vaie.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vaie.Type(), vaie.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VertexAiIndexEndpoint] has state.
func (vaie *VertexAiIndexEndpoint) State() (*vertexAiIndexEndpointState, bool) {
	return vaie.state, vaie.state != nil
}

// StateMust returns the state for [VertexAiIndexEndpoint]. Panics if the state is nil.
func (vaie *VertexAiIndexEndpoint) StateMust() *vertexAiIndexEndpointState {
	if vaie.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vaie.Type(), vaie.LocalName()))
	}
	return vaie.state
}

// VertexAiIndexEndpointArgs contains the configurations for google_vertex_ai_index_endpoint.
type VertexAiIndexEndpointArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Network: string, optional
	Network terra.StringValue `hcl:"network,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Timeouts: optional
	Timeouts *vertexaiindexendpoint.Timeouts `hcl:"timeouts,block"`
}
type vertexAiIndexEndpointAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_vertex_ai_index_endpoint.
func (vaie vertexAiIndexEndpointAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(vaie.ref.Append("create_time"))
}

// Description returns a reference to field description of google_vertex_ai_index_endpoint.
func (vaie vertexAiIndexEndpointAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(vaie.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of google_vertex_ai_index_endpoint.
func (vaie vertexAiIndexEndpointAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(vaie.ref.Append("display_name"))
}

// Etag returns a reference to field etag of google_vertex_ai_index_endpoint.
func (vaie vertexAiIndexEndpointAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(vaie.ref.Append("etag"))
}

// Id returns a reference to field id of google_vertex_ai_index_endpoint.
func (vaie vertexAiIndexEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vaie.ref.Append("id"))
}

// Labels returns a reference to field labels of google_vertex_ai_index_endpoint.
func (vaie vertexAiIndexEndpointAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vaie.ref.Append("labels"))
}

// Name returns a reference to field name of google_vertex_ai_index_endpoint.
func (vaie vertexAiIndexEndpointAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vaie.ref.Append("name"))
}

// Network returns a reference to field network of google_vertex_ai_index_endpoint.
func (vaie vertexAiIndexEndpointAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(vaie.ref.Append("network"))
}

// Project returns a reference to field project of google_vertex_ai_index_endpoint.
func (vaie vertexAiIndexEndpointAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(vaie.ref.Append("project"))
}

// Region returns a reference to field region of google_vertex_ai_index_endpoint.
func (vaie vertexAiIndexEndpointAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(vaie.ref.Append("region"))
}

// UpdateTime returns a reference to field update_time of google_vertex_ai_index_endpoint.
func (vaie vertexAiIndexEndpointAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(vaie.ref.Append("update_time"))
}

func (vaie vertexAiIndexEndpointAttributes) Timeouts() vertexaiindexendpoint.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vertexaiindexendpoint.TimeoutsAttributes](vaie.ref.Append("timeouts"))
}

type vertexAiIndexEndpointState struct {
	CreateTime  string                               `json:"create_time"`
	Description string                               `json:"description"`
	DisplayName string                               `json:"display_name"`
	Etag        string                               `json:"etag"`
	Id          string                               `json:"id"`
	Labels      map[string]string                    `json:"labels"`
	Name        string                               `json:"name"`
	Network     string                               `json:"network"`
	Project     string                               `json:"project"`
	Region      string                               `json:"region"`
	UpdateTime  string                               `json:"update_time"`
	Timeouts    *vertexaiindexendpoint.TimeoutsState `json:"timeouts"`
}
