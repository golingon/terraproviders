// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	vertexaiendpoint "github.com/golingon/terraproviders/google/4.64.0/vertexaiendpoint"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVertexAiEndpoint creates a new instance of [VertexAiEndpoint].
func NewVertexAiEndpoint(name string, args VertexAiEndpointArgs) *VertexAiEndpoint {
	return &VertexAiEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VertexAiEndpoint)(nil)

// VertexAiEndpoint represents the Terraform resource google_vertex_ai_endpoint.
type VertexAiEndpoint struct {
	Name      string
	Args      VertexAiEndpointArgs
	state     *vertexAiEndpointState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VertexAiEndpoint].
func (vae *VertexAiEndpoint) Type() string {
	return "google_vertex_ai_endpoint"
}

// LocalName returns the local name for [VertexAiEndpoint].
func (vae *VertexAiEndpoint) LocalName() string {
	return vae.Name
}

// Configuration returns the configuration (args) for [VertexAiEndpoint].
func (vae *VertexAiEndpoint) Configuration() interface{} {
	return vae.Args
}

// DependOn is used for other resources to depend on [VertexAiEndpoint].
func (vae *VertexAiEndpoint) DependOn() terra.Reference {
	return terra.ReferenceResource(vae)
}

// Dependencies returns the list of resources [VertexAiEndpoint] depends_on.
func (vae *VertexAiEndpoint) Dependencies() terra.Dependencies {
	return vae.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VertexAiEndpoint].
func (vae *VertexAiEndpoint) LifecycleManagement() *terra.Lifecycle {
	return vae.Lifecycle
}

// Attributes returns the attributes for [VertexAiEndpoint].
func (vae *VertexAiEndpoint) Attributes() vertexAiEndpointAttributes {
	return vertexAiEndpointAttributes{ref: terra.ReferenceResource(vae)}
}

// ImportState imports the given attribute values into [VertexAiEndpoint]'s state.
func (vae *VertexAiEndpoint) ImportState(av io.Reader) error {
	vae.state = &vertexAiEndpointState{}
	if err := json.NewDecoder(av).Decode(vae.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vae.Type(), vae.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VertexAiEndpoint] has state.
func (vae *VertexAiEndpoint) State() (*vertexAiEndpointState, bool) {
	return vae.state, vae.state != nil
}

// StateMust returns the state for [VertexAiEndpoint]. Panics if the state is nil.
func (vae *VertexAiEndpoint) StateMust() *vertexAiEndpointState {
	if vae.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vae.Type(), vae.LocalName()))
	}
	return vae.state
}

// VertexAiEndpointArgs contains the configurations for google_vertex_ai_endpoint.
type VertexAiEndpointArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Network: string, optional
	Network terra.StringValue `hcl:"network,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// DeployedModels: min=0
	DeployedModels []vertexaiendpoint.DeployedModels `hcl:"deployed_models,block" validate:"min=0"`
	// EncryptionSpec: optional
	EncryptionSpec *vertexaiendpoint.EncryptionSpec `hcl:"encryption_spec,block"`
	// Timeouts: optional
	Timeouts *vertexaiendpoint.Timeouts `hcl:"timeouts,block"`
}
type vertexAiEndpointAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_vertex_ai_endpoint.
func (vae vertexAiEndpointAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(vae.ref.Append("create_time"))
}

// Description returns a reference to field description of google_vertex_ai_endpoint.
func (vae vertexAiEndpointAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(vae.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of google_vertex_ai_endpoint.
func (vae vertexAiEndpointAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(vae.ref.Append("display_name"))
}

// Etag returns a reference to field etag of google_vertex_ai_endpoint.
func (vae vertexAiEndpointAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(vae.ref.Append("etag"))
}

// Id returns a reference to field id of google_vertex_ai_endpoint.
func (vae vertexAiEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vae.ref.Append("id"))
}

// Labels returns a reference to field labels of google_vertex_ai_endpoint.
func (vae vertexAiEndpointAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vae.ref.Append("labels"))
}

// Location returns a reference to field location of google_vertex_ai_endpoint.
func (vae vertexAiEndpointAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(vae.ref.Append("location"))
}

// ModelDeploymentMonitoringJob returns a reference to field model_deployment_monitoring_job of google_vertex_ai_endpoint.
func (vae vertexAiEndpointAttributes) ModelDeploymentMonitoringJob() terra.StringValue {
	return terra.ReferenceAsString(vae.ref.Append("model_deployment_monitoring_job"))
}

// Name returns a reference to field name of google_vertex_ai_endpoint.
func (vae vertexAiEndpointAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vae.ref.Append("name"))
}

// Network returns a reference to field network of google_vertex_ai_endpoint.
func (vae vertexAiEndpointAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(vae.ref.Append("network"))
}

// Project returns a reference to field project of google_vertex_ai_endpoint.
func (vae vertexAiEndpointAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(vae.ref.Append("project"))
}

// Region returns a reference to field region of google_vertex_ai_endpoint.
func (vae vertexAiEndpointAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(vae.ref.Append("region"))
}

// UpdateTime returns a reference to field update_time of google_vertex_ai_endpoint.
func (vae vertexAiEndpointAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(vae.ref.Append("update_time"))
}

func (vae vertexAiEndpointAttributes) DeployedModels() terra.ListValue[vertexaiendpoint.DeployedModelsAttributes] {
	return terra.ReferenceAsList[vertexaiendpoint.DeployedModelsAttributes](vae.ref.Append("deployed_models"))
}

func (vae vertexAiEndpointAttributes) EncryptionSpec() terra.ListValue[vertexaiendpoint.EncryptionSpecAttributes] {
	return terra.ReferenceAsList[vertexaiendpoint.EncryptionSpecAttributes](vae.ref.Append("encryption_spec"))
}

func (vae vertexAiEndpointAttributes) Timeouts() vertexaiendpoint.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vertexaiendpoint.TimeoutsAttributes](vae.ref.Append("timeouts"))
}

type vertexAiEndpointState struct {
	CreateTime                   string                                 `json:"create_time"`
	Description                  string                                 `json:"description"`
	DisplayName                  string                                 `json:"display_name"`
	Etag                         string                                 `json:"etag"`
	Id                           string                                 `json:"id"`
	Labels                       map[string]string                      `json:"labels"`
	Location                     string                                 `json:"location"`
	ModelDeploymentMonitoringJob string                                 `json:"model_deployment_monitoring_job"`
	Name                         string                                 `json:"name"`
	Network                      string                                 `json:"network"`
	Project                      string                                 `json:"project"`
	Region                       string                                 `json:"region"`
	UpdateTime                   string                                 `json:"update_time"`
	DeployedModels               []vertexaiendpoint.DeployedModelsState `json:"deployed_models"`
	EncryptionSpec               []vertexaiendpoint.EncryptionSpecState `json:"encryption_spec"`
	Timeouts                     *vertexaiendpoint.TimeoutsState        `json:"timeouts"`
}
