// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApigatewayv2RouteResponse creates a new instance of [Apigatewayv2RouteResponse].
func NewApigatewayv2RouteResponse(name string, args Apigatewayv2RouteResponseArgs) *Apigatewayv2RouteResponse {
	return &Apigatewayv2RouteResponse{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Apigatewayv2RouteResponse)(nil)

// Apigatewayv2RouteResponse represents the Terraform resource aws_apigatewayv2_route_response.
type Apigatewayv2RouteResponse struct {
	Name      string
	Args      Apigatewayv2RouteResponseArgs
	state     *apigatewayv2RouteResponseState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Apigatewayv2RouteResponse].
func (arr *Apigatewayv2RouteResponse) Type() string {
	return "aws_apigatewayv2_route_response"
}

// LocalName returns the local name for [Apigatewayv2RouteResponse].
func (arr *Apigatewayv2RouteResponse) LocalName() string {
	return arr.Name
}

// Configuration returns the configuration (args) for [Apigatewayv2RouteResponse].
func (arr *Apigatewayv2RouteResponse) Configuration() interface{} {
	return arr.Args
}

// DependOn is used for other resources to depend on [Apigatewayv2RouteResponse].
func (arr *Apigatewayv2RouteResponse) DependOn() terra.Reference {
	return terra.ReferenceResource(arr)
}

// Dependencies returns the list of resources [Apigatewayv2RouteResponse] depends_on.
func (arr *Apigatewayv2RouteResponse) Dependencies() terra.Dependencies {
	return arr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Apigatewayv2RouteResponse].
func (arr *Apigatewayv2RouteResponse) LifecycleManagement() *terra.Lifecycle {
	return arr.Lifecycle
}

// Attributes returns the attributes for [Apigatewayv2RouteResponse].
func (arr *Apigatewayv2RouteResponse) Attributes() apigatewayv2RouteResponseAttributes {
	return apigatewayv2RouteResponseAttributes{ref: terra.ReferenceResource(arr)}
}

// ImportState imports the given attribute values into [Apigatewayv2RouteResponse]'s state.
func (arr *Apigatewayv2RouteResponse) ImportState(av io.Reader) error {
	arr.state = &apigatewayv2RouteResponseState{}
	if err := json.NewDecoder(av).Decode(arr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", arr.Type(), arr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Apigatewayv2RouteResponse] has state.
func (arr *Apigatewayv2RouteResponse) State() (*apigatewayv2RouteResponseState, bool) {
	return arr.state, arr.state != nil
}

// StateMust returns the state for [Apigatewayv2RouteResponse]. Panics if the state is nil.
func (arr *Apigatewayv2RouteResponse) StateMust() *apigatewayv2RouteResponseState {
	if arr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", arr.Type(), arr.LocalName()))
	}
	return arr.state
}

// Apigatewayv2RouteResponseArgs contains the configurations for aws_apigatewayv2_route_response.
type Apigatewayv2RouteResponseArgs struct {
	// ApiId: string, required
	ApiId terra.StringValue `hcl:"api_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ModelSelectionExpression: string, optional
	ModelSelectionExpression terra.StringValue `hcl:"model_selection_expression,attr"`
	// ResponseModels: map of string, optional
	ResponseModels terra.MapValue[terra.StringValue] `hcl:"response_models,attr"`
	// RouteId: string, required
	RouteId terra.StringValue `hcl:"route_id,attr" validate:"required"`
	// RouteResponseKey: string, required
	RouteResponseKey terra.StringValue `hcl:"route_response_key,attr" validate:"required"`
}
type apigatewayv2RouteResponseAttributes struct {
	ref terra.Reference
}

// ApiId returns a reference to field api_id of aws_apigatewayv2_route_response.
func (arr apigatewayv2RouteResponseAttributes) ApiId() terra.StringValue {
	return terra.ReferenceAsString(arr.ref.Append("api_id"))
}

// Id returns a reference to field id of aws_apigatewayv2_route_response.
func (arr apigatewayv2RouteResponseAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(arr.ref.Append("id"))
}

// ModelSelectionExpression returns a reference to field model_selection_expression of aws_apigatewayv2_route_response.
func (arr apigatewayv2RouteResponseAttributes) ModelSelectionExpression() terra.StringValue {
	return terra.ReferenceAsString(arr.ref.Append("model_selection_expression"))
}

// ResponseModels returns a reference to field response_models of aws_apigatewayv2_route_response.
func (arr apigatewayv2RouteResponseAttributes) ResponseModels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](arr.ref.Append("response_models"))
}

// RouteId returns a reference to field route_id of aws_apigatewayv2_route_response.
func (arr apigatewayv2RouteResponseAttributes) RouteId() terra.StringValue {
	return terra.ReferenceAsString(arr.ref.Append("route_id"))
}

// RouteResponseKey returns a reference to field route_response_key of aws_apigatewayv2_route_response.
func (arr apigatewayv2RouteResponseAttributes) RouteResponseKey() terra.StringValue {
	return terra.ReferenceAsString(arr.ref.Append("route_response_key"))
}

type apigatewayv2RouteResponseState struct {
	ApiId                    string            `json:"api_id"`
	Id                       string            `json:"id"`
	ModelSelectionExpression string            `json:"model_selection_expression"`
	ResponseModels           map[string]string `json:"response_models"`
	RouteId                  string            `json:"route_id"`
	RouteResponseKey         string            `json:"route_response_key"`
}
