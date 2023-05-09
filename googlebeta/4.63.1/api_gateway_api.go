// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	apigatewayapi "github.com/golingon/terraproviders/googlebeta/4.63.1/apigatewayapi"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiGatewayApi creates a new instance of [ApiGatewayApi].
func NewApiGatewayApi(name string, args ApiGatewayApiArgs) *ApiGatewayApi {
	return &ApiGatewayApi{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiGatewayApi)(nil)

// ApiGatewayApi represents the Terraform resource google_api_gateway_api.
type ApiGatewayApi struct {
	Name      string
	Args      ApiGatewayApiArgs
	state     *apiGatewayApiState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiGatewayApi].
func (aga *ApiGatewayApi) Type() string {
	return "google_api_gateway_api"
}

// LocalName returns the local name for [ApiGatewayApi].
func (aga *ApiGatewayApi) LocalName() string {
	return aga.Name
}

// Configuration returns the configuration (args) for [ApiGatewayApi].
func (aga *ApiGatewayApi) Configuration() interface{} {
	return aga.Args
}

// DependOn is used for other resources to depend on [ApiGatewayApi].
func (aga *ApiGatewayApi) DependOn() terra.Reference {
	return terra.ReferenceResource(aga)
}

// Dependencies returns the list of resources [ApiGatewayApi] depends_on.
func (aga *ApiGatewayApi) Dependencies() terra.Dependencies {
	return aga.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiGatewayApi].
func (aga *ApiGatewayApi) LifecycleManagement() *terra.Lifecycle {
	return aga.Lifecycle
}

// Attributes returns the attributes for [ApiGatewayApi].
func (aga *ApiGatewayApi) Attributes() apiGatewayApiAttributes {
	return apiGatewayApiAttributes{ref: terra.ReferenceResource(aga)}
}

// ImportState imports the given attribute values into [ApiGatewayApi]'s state.
func (aga *ApiGatewayApi) ImportState(av io.Reader) error {
	aga.state = &apiGatewayApiState{}
	if err := json.NewDecoder(av).Decode(aga.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aga.Type(), aga.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiGatewayApi] has state.
func (aga *ApiGatewayApi) State() (*apiGatewayApiState, bool) {
	return aga.state, aga.state != nil
}

// StateMust returns the state for [ApiGatewayApi]. Panics if the state is nil.
func (aga *ApiGatewayApi) StateMust() *apiGatewayApiState {
	if aga.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aga.Type(), aga.LocalName()))
	}
	return aga.state
}

// ApiGatewayApiArgs contains the configurations for google_api_gateway_api.
type ApiGatewayApiArgs struct {
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
	Timeouts *apigatewayapi.Timeouts `hcl:"timeouts,block"`
}
type apiGatewayApiAttributes struct {
	ref terra.Reference
}

// ApiId returns a reference to field api_id of google_api_gateway_api.
func (aga apiGatewayApiAttributes) ApiId() terra.StringValue {
	return terra.ReferenceAsString(aga.ref.Append("api_id"))
}

// CreateTime returns a reference to field create_time of google_api_gateway_api.
func (aga apiGatewayApiAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(aga.ref.Append("create_time"))
}

// DisplayName returns a reference to field display_name of google_api_gateway_api.
func (aga apiGatewayApiAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(aga.ref.Append("display_name"))
}

// Id returns a reference to field id of google_api_gateway_api.
func (aga apiGatewayApiAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aga.ref.Append("id"))
}

// Labels returns a reference to field labels of google_api_gateway_api.
func (aga apiGatewayApiAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aga.ref.Append("labels"))
}

// ManagedService returns a reference to field managed_service of google_api_gateway_api.
func (aga apiGatewayApiAttributes) ManagedService() terra.StringValue {
	return terra.ReferenceAsString(aga.ref.Append("managed_service"))
}

// Name returns a reference to field name of google_api_gateway_api.
func (aga apiGatewayApiAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aga.ref.Append("name"))
}

// Project returns a reference to field project of google_api_gateway_api.
func (aga apiGatewayApiAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(aga.ref.Append("project"))
}

func (aga apiGatewayApiAttributes) Timeouts() apigatewayapi.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apigatewayapi.TimeoutsAttributes](aga.ref.Append("timeouts"))
}

type apiGatewayApiState struct {
	ApiId          string                       `json:"api_id"`
	CreateTime     string                       `json:"create_time"`
	DisplayName    string                       `json:"display_name"`
	Id             string                       `json:"id"`
	Labels         map[string]string            `json:"labels"`
	ManagedService string                       `json:"managed_service"`
	Name           string                       `json:"name"`
	Project        string                       `json:"project"`
	Timeouts       *apigatewayapi.TimeoutsState `json:"timeouts"`
}
