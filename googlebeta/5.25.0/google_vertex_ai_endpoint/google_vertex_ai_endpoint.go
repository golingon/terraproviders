// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_vertex_ai_endpoint

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

// Resource represents the Terraform resource google_vertex_ai_endpoint.
type Resource struct {
	Name      string
	Args      Args
	state     *googleVertexAiEndpointState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gvae *Resource) Type() string {
	return "google_vertex_ai_endpoint"
}

// LocalName returns the local name for [Resource].
func (gvae *Resource) LocalName() string {
	return gvae.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gvae *Resource) Configuration() interface{} {
	return gvae.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gvae *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gvae)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gvae *Resource) Dependencies() terra.Dependencies {
	return gvae.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gvae *Resource) LifecycleManagement() *terra.Lifecycle {
	return gvae.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gvae *Resource) Attributes() googleVertexAiEndpointAttributes {
	return googleVertexAiEndpointAttributes{ref: terra.ReferenceResource(gvae)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gvae *Resource) ImportState(state io.Reader) error {
	gvae.state = &googleVertexAiEndpointState{}
	if err := json.NewDecoder(state).Decode(gvae.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gvae.Type(), gvae.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gvae *Resource) State() (*googleVertexAiEndpointState, bool) {
	return gvae.state, gvae.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gvae *Resource) StateMust() *googleVertexAiEndpointState {
	if gvae.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gvae.Type(), gvae.LocalName()))
	}
	return gvae.state
}

// Args contains the configurations for google_vertex_ai_endpoint.
type Args struct {
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
	// EncryptionSpec: optional
	EncryptionSpec *EncryptionSpec `hcl:"encryption_spec,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleVertexAiEndpointAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_vertex_ai_endpoint.
func (gvae googleVertexAiEndpointAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(gvae.ref.Append("create_time"))
}

// Description returns a reference to field description of google_vertex_ai_endpoint.
func (gvae googleVertexAiEndpointAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gvae.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of google_vertex_ai_endpoint.
func (gvae googleVertexAiEndpointAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(gvae.ref.Append("display_name"))
}

// EffectiveLabels returns a reference to field effective_labels of google_vertex_ai_endpoint.
func (gvae googleVertexAiEndpointAttributes) EffectiveLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gvae.ref.Append("effective_labels"))
}

// Etag returns a reference to field etag of google_vertex_ai_endpoint.
func (gvae googleVertexAiEndpointAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gvae.ref.Append("etag"))
}

// Id returns a reference to field id of google_vertex_ai_endpoint.
func (gvae googleVertexAiEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gvae.ref.Append("id"))
}

// Labels returns a reference to field labels of google_vertex_ai_endpoint.
func (gvae googleVertexAiEndpointAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gvae.ref.Append("labels"))
}

// Location returns a reference to field location of google_vertex_ai_endpoint.
func (gvae googleVertexAiEndpointAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gvae.ref.Append("location"))
}

// ModelDeploymentMonitoringJob returns a reference to field model_deployment_monitoring_job of google_vertex_ai_endpoint.
func (gvae googleVertexAiEndpointAttributes) ModelDeploymentMonitoringJob() terra.StringValue {
	return terra.ReferenceAsString(gvae.ref.Append("model_deployment_monitoring_job"))
}

// Name returns a reference to field name of google_vertex_ai_endpoint.
func (gvae googleVertexAiEndpointAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gvae.ref.Append("name"))
}

// Network returns a reference to field network of google_vertex_ai_endpoint.
func (gvae googleVertexAiEndpointAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(gvae.ref.Append("network"))
}

// Project returns a reference to field project of google_vertex_ai_endpoint.
func (gvae googleVertexAiEndpointAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gvae.ref.Append("project"))
}

// Region returns a reference to field region of google_vertex_ai_endpoint.
func (gvae googleVertexAiEndpointAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(gvae.ref.Append("region"))
}

// TerraformLabels returns a reference to field terraform_labels of google_vertex_ai_endpoint.
func (gvae googleVertexAiEndpointAttributes) TerraformLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gvae.ref.Append("terraform_labels"))
}

// UpdateTime returns a reference to field update_time of google_vertex_ai_endpoint.
func (gvae googleVertexAiEndpointAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(gvae.ref.Append("update_time"))
}

func (gvae googleVertexAiEndpointAttributes) DeployedModels() terra.ListValue[DeployedModelsAttributes] {
	return terra.ReferenceAsList[DeployedModelsAttributes](gvae.ref.Append("deployed_models"))
}

func (gvae googleVertexAiEndpointAttributes) EncryptionSpec() terra.ListValue[EncryptionSpecAttributes] {
	return terra.ReferenceAsList[EncryptionSpecAttributes](gvae.ref.Append("encryption_spec"))
}

func (gvae googleVertexAiEndpointAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gvae.ref.Append("timeouts"))
}

type googleVertexAiEndpointState struct {
	CreateTime                   string                `json:"create_time"`
	Description                  string                `json:"description"`
	DisplayName                  string                `json:"display_name"`
	EffectiveLabels              map[string]string     `json:"effective_labels"`
	Etag                         string                `json:"etag"`
	Id                           string                `json:"id"`
	Labels                       map[string]string     `json:"labels"`
	Location                     string                `json:"location"`
	ModelDeploymentMonitoringJob string                `json:"model_deployment_monitoring_job"`
	Name                         string                `json:"name"`
	Network                      string                `json:"network"`
	Project                      string                `json:"project"`
	Region                       string                `json:"region"`
	TerraformLabels              map[string]string     `json:"terraform_labels"`
	UpdateTime                   string                `json:"update_time"`
	DeployedModels               []DeployedModelsState `json:"deployed_models"`
	EncryptionSpec               []EncryptionSpecState `json:"encryption_spec"`
	Timeouts                     *TimeoutsState        `json:"timeouts"`
}
