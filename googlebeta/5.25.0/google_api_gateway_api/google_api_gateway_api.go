// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_api_gateway_api

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

// Resource represents the Terraform resource google_api_gateway_api.
type Resource struct {
	Name      string
	Args      Args
	state     *googleApiGatewayApiState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gaga *Resource) Type() string {
	return "google_api_gateway_api"
}

// LocalName returns the local name for [Resource].
func (gaga *Resource) LocalName() string {
	return gaga.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gaga *Resource) Configuration() interface{} {
	return gaga.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gaga *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gaga)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gaga *Resource) Dependencies() terra.Dependencies {
	return gaga.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gaga *Resource) LifecycleManagement() *terra.Lifecycle {
	return gaga.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gaga *Resource) Attributes() googleApiGatewayApiAttributes {
	return googleApiGatewayApiAttributes{ref: terra.ReferenceResource(gaga)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gaga *Resource) ImportState(state io.Reader) error {
	gaga.state = &googleApiGatewayApiState{}
	if err := json.NewDecoder(state).Decode(gaga.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gaga.Type(), gaga.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gaga *Resource) State() (*googleApiGatewayApiState, bool) {
	return gaga.state, gaga.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gaga *Resource) StateMust() *googleApiGatewayApiState {
	if gaga.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gaga.Type(), gaga.LocalName()))
	}
	return gaga.state
}

// Args contains the configurations for google_api_gateway_api.
type Args struct {
	// ApiId: string, required
	ApiId terra.StringValue `hcl:"api_id,attr" validate:"required"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// ManagedService: string, optional
	ManagedService terra.StringValue `hcl:"managed_service,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleApiGatewayApiAttributes struct {
	ref terra.Reference
}

// ApiId returns a reference to field api_id of google_api_gateway_api.
func (gaga googleApiGatewayApiAttributes) ApiId() terra.StringValue {
	return terra.ReferenceAsString(gaga.ref.Append("api_id"))
}

// CreateTime returns a reference to field create_time of google_api_gateway_api.
func (gaga googleApiGatewayApiAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(gaga.ref.Append("create_time"))
}

// DisplayName returns a reference to field display_name of google_api_gateway_api.
func (gaga googleApiGatewayApiAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(gaga.ref.Append("display_name"))
}

// EffectiveLabels returns a reference to field effective_labels of google_api_gateway_api.
func (gaga googleApiGatewayApiAttributes) EffectiveLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gaga.ref.Append("effective_labels"))
}

// Id returns a reference to field id of google_api_gateway_api.
func (gaga googleApiGatewayApiAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gaga.ref.Append("id"))
}

// Labels returns a reference to field labels of google_api_gateway_api.
func (gaga googleApiGatewayApiAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gaga.ref.Append("labels"))
}

// ManagedService returns a reference to field managed_service of google_api_gateway_api.
func (gaga googleApiGatewayApiAttributes) ManagedService() terra.StringValue {
	return terra.ReferenceAsString(gaga.ref.Append("managed_service"))
}

// Name returns a reference to field name of google_api_gateway_api.
func (gaga googleApiGatewayApiAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gaga.ref.Append("name"))
}

// Project returns a reference to field project of google_api_gateway_api.
func (gaga googleApiGatewayApiAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gaga.ref.Append("project"))
}

// TerraformLabels returns a reference to field terraform_labels of google_api_gateway_api.
func (gaga googleApiGatewayApiAttributes) TerraformLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gaga.ref.Append("terraform_labels"))
}

func (gaga googleApiGatewayApiAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gaga.ref.Append("timeouts"))
}

type googleApiGatewayApiState struct {
	ApiId           string            `json:"api_id"`
	CreateTime      string            `json:"create_time"`
	DisplayName     string            `json:"display_name"`
	EffectiveLabels map[string]string `json:"effective_labels"`
	Id              string            `json:"id"`
	Labels          map[string]string `json:"labels"`
	ManagedService  string            `json:"managed_service"`
	Name            string            `json:"name"`
	Project         string            `json:"project"`
	TerraformLabels map[string]string `json:"terraform_labels"`
	Timeouts        *TimeoutsState    `json:"timeouts"`
}
