// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	vertexaiindexendpoint "github.com/golingon/terraproviders/googlebeta/5.24.0/vertexaiindexendpoint"
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
	// PublicEndpointEnabled: bool, optional
	PublicEndpointEnabled terra.BoolValue `hcl:"public_endpoint_enabled,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// PrivateServiceConnectConfig: optional
	PrivateServiceConnectConfig *vertexaiindexendpoint.PrivateServiceConnectConfig `hcl:"private_service_connect_config,block"`
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

// EffectiveLabels returns a reference to field effective_labels of google_vertex_ai_index_endpoint.
func (vaie vertexAiIndexEndpointAttributes) EffectiveLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vaie.ref.Append("effective_labels"))
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

// PublicEndpointDomainName returns a reference to field public_endpoint_domain_name of google_vertex_ai_index_endpoint.
func (vaie vertexAiIndexEndpointAttributes) PublicEndpointDomainName() terra.StringValue {
	return terra.ReferenceAsString(vaie.ref.Append("public_endpoint_domain_name"))
}

// PublicEndpointEnabled returns a reference to field public_endpoint_enabled of google_vertex_ai_index_endpoint.
func (vaie vertexAiIndexEndpointAttributes) PublicEndpointEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(vaie.ref.Append("public_endpoint_enabled"))
}

// Region returns a reference to field region of google_vertex_ai_index_endpoint.
func (vaie vertexAiIndexEndpointAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(vaie.ref.Append("region"))
}

// TerraformLabels returns a reference to field terraform_labels of google_vertex_ai_index_endpoint.
func (vaie vertexAiIndexEndpointAttributes) TerraformLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vaie.ref.Append("terraform_labels"))
}

// UpdateTime returns a reference to field update_time of google_vertex_ai_index_endpoint.
func (vaie vertexAiIndexEndpointAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(vaie.ref.Append("update_time"))
}

func (vaie vertexAiIndexEndpointAttributes) PrivateServiceConnectConfig() terra.ListValue[vertexaiindexendpoint.PrivateServiceConnectConfigAttributes] {
	return terra.ReferenceAsList[vertexaiindexendpoint.PrivateServiceConnectConfigAttributes](vaie.ref.Append("private_service_connect_config"))
}

func (vaie vertexAiIndexEndpointAttributes) Timeouts() vertexaiindexendpoint.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vertexaiindexendpoint.TimeoutsAttributes](vaie.ref.Append("timeouts"))
}

type vertexAiIndexEndpointState struct {
	CreateTime                  string                                                   `json:"create_time"`
	Description                 string                                                   `json:"description"`
	DisplayName                 string                                                   `json:"display_name"`
	EffectiveLabels             map[string]string                                        `json:"effective_labels"`
	Etag                        string                                                   `json:"etag"`
	Id                          string                                                   `json:"id"`
	Labels                      map[string]string                                        `json:"labels"`
	Name                        string                                                   `json:"name"`
	Network                     string                                                   `json:"network"`
	Project                     string                                                   `json:"project"`
	PublicEndpointDomainName    string                                                   `json:"public_endpoint_domain_name"`
	PublicEndpointEnabled       bool                                                     `json:"public_endpoint_enabled"`
	Region                      string                                                   `json:"region"`
	TerraformLabels             map[string]string                                        `json:"terraform_labels"`
	UpdateTime                  string                                                   `json:"update_time"`
	PrivateServiceConnectConfig []vertexaiindexendpoint.PrivateServiceConnectConfigState `json:"private_service_connect_config"`
	Timeouts                    *vertexaiindexendpoint.TimeoutsState                     `json:"timeouts"`
}
