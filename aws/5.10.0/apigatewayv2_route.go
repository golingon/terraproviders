// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	apigatewayv2route "github.com/golingon/terraproviders/aws/5.10.0/apigatewayv2route"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApigatewayv2Route creates a new instance of [Apigatewayv2Route].
func NewApigatewayv2Route(name string, args Apigatewayv2RouteArgs) *Apigatewayv2Route {
	return &Apigatewayv2Route{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Apigatewayv2Route)(nil)

// Apigatewayv2Route represents the Terraform resource aws_apigatewayv2_route.
type Apigatewayv2Route struct {
	Name      string
	Args      Apigatewayv2RouteArgs
	state     *apigatewayv2RouteState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Apigatewayv2Route].
func (ar *Apigatewayv2Route) Type() string {
	return "aws_apigatewayv2_route"
}

// LocalName returns the local name for [Apigatewayv2Route].
func (ar *Apigatewayv2Route) LocalName() string {
	return ar.Name
}

// Configuration returns the configuration (args) for [Apigatewayv2Route].
func (ar *Apigatewayv2Route) Configuration() interface{} {
	return ar.Args
}

// DependOn is used for other resources to depend on [Apigatewayv2Route].
func (ar *Apigatewayv2Route) DependOn() terra.Reference {
	return terra.ReferenceResource(ar)
}

// Dependencies returns the list of resources [Apigatewayv2Route] depends_on.
func (ar *Apigatewayv2Route) Dependencies() terra.Dependencies {
	return ar.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Apigatewayv2Route].
func (ar *Apigatewayv2Route) LifecycleManagement() *terra.Lifecycle {
	return ar.Lifecycle
}

// Attributes returns the attributes for [Apigatewayv2Route].
func (ar *Apigatewayv2Route) Attributes() apigatewayv2RouteAttributes {
	return apigatewayv2RouteAttributes{ref: terra.ReferenceResource(ar)}
}

// ImportState imports the given attribute values into [Apigatewayv2Route]'s state.
func (ar *Apigatewayv2Route) ImportState(av io.Reader) error {
	ar.state = &apigatewayv2RouteState{}
	if err := json.NewDecoder(av).Decode(ar.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ar.Type(), ar.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Apigatewayv2Route] has state.
func (ar *Apigatewayv2Route) State() (*apigatewayv2RouteState, bool) {
	return ar.state, ar.state != nil
}

// StateMust returns the state for [Apigatewayv2Route]. Panics if the state is nil.
func (ar *Apigatewayv2Route) StateMust() *apigatewayv2RouteState {
	if ar.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ar.Type(), ar.LocalName()))
	}
	return ar.state
}

// Apigatewayv2RouteArgs contains the configurations for aws_apigatewayv2_route.
type Apigatewayv2RouteArgs struct {
	// ApiId: string, required
	ApiId terra.StringValue `hcl:"api_id,attr" validate:"required"`
	// ApiKeyRequired: bool, optional
	ApiKeyRequired terra.BoolValue `hcl:"api_key_required,attr"`
	// AuthorizationScopes: set of string, optional
	AuthorizationScopes terra.SetValue[terra.StringValue] `hcl:"authorization_scopes,attr"`
	// AuthorizationType: string, optional
	AuthorizationType terra.StringValue `hcl:"authorization_type,attr"`
	// AuthorizerId: string, optional
	AuthorizerId terra.StringValue `hcl:"authorizer_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ModelSelectionExpression: string, optional
	ModelSelectionExpression terra.StringValue `hcl:"model_selection_expression,attr"`
	// OperationName: string, optional
	OperationName terra.StringValue `hcl:"operation_name,attr"`
	// RequestModels: map of string, optional
	RequestModels terra.MapValue[terra.StringValue] `hcl:"request_models,attr"`
	// RouteKey: string, required
	RouteKey terra.StringValue `hcl:"route_key,attr" validate:"required"`
	// RouteResponseSelectionExpression: string, optional
	RouteResponseSelectionExpression terra.StringValue `hcl:"route_response_selection_expression,attr"`
	// Target: string, optional
	Target terra.StringValue `hcl:"target,attr"`
	// RequestParameter: min=0
	RequestParameter []apigatewayv2route.RequestParameter `hcl:"request_parameter,block" validate:"min=0"`
}
type apigatewayv2RouteAttributes struct {
	ref terra.Reference
}

// ApiId returns a reference to field api_id of aws_apigatewayv2_route.
func (ar apigatewayv2RouteAttributes) ApiId() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("api_id"))
}

// ApiKeyRequired returns a reference to field api_key_required of aws_apigatewayv2_route.
func (ar apigatewayv2RouteAttributes) ApiKeyRequired() terra.BoolValue {
	return terra.ReferenceAsBool(ar.ref.Append("api_key_required"))
}

// AuthorizationScopes returns a reference to field authorization_scopes of aws_apigatewayv2_route.
func (ar apigatewayv2RouteAttributes) AuthorizationScopes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ar.ref.Append("authorization_scopes"))
}

// AuthorizationType returns a reference to field authorization_type of aws_apigatewayv2_route.
func (ar apigatewayv2RouteAttributes) AuthorizationType() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("authorization_type"))
}

// AuthorizerId returns a reference to field authorizer_id of aws_apigatewayv2_route.
func (ar apigatewayv2RouteAttributes) AuthorizerId() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("authorizer_id"))
}

// Id returns a reference to field id of aws_apigatewayv2_route.
func (ar apigatewayv2RouteAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("id"))
}

// ModelSelectionExpression returns a reference to field model_selection_expression of aws_apigatewayv2_route.
func (ar apigatewayv2RouteAttributes) ModelSelectionExpression() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("model_selection_expression"))
}

// OperationName returns a reference to field operation_name of aws_apigatewayv2_route.
func (ar apigatewayv2RouteAttributes) OperationName() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("operation_name"))
}

// RequestModels returns a reference to field request_models of aws_apigatewayv2_route.
func (ar apigatewayv2RouteAttributes) RequestModels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ar.ref.Append("request_models"))
}

// RouteKey returns a reference to field route_key of aws_apigatewayv2_route.
func (ar apigatewayv2RouteAttributes) RouteKey() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("route_key"))
}

// RouteResponseSelectionExpression returns a reference to field route_response_selection_expression of aws_apigatewayv2_route.
func (ar apigatewayv2RouteAttributes) RouteResponseSelectionExpression() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("route_response_selection_expression"))
}

// Target returns a reference to field target of aws_apigatewayv2_route.
func (ar apigatewayv2RouteAttributes) Target() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("target"))
}

func (ar apigatewayv2RouteAttributes) RequestParameter() terra.SetValue[apigatewayv2route.RequestParameterAttributes] {
	return terra.ReferenceAsSet[apigatewayv2route.RequestParameterAttributes](ar.ref.Append("request_parameter"))
}

type apigatewayv2RouteState struct {
	ApiId                            string                                    `json:"api_id"`
	ApiKeyRequired                   bool                                      `json:"api_key_required"`
	AuthorizationScopes              []string                                  `json:"authorization_scopes"`
	AuthorizationType                string                                    `json:"authorization_type"`
	AuthorizerId                     string                                    `json:"authorizer_id"`
	Id                               string                                    `json:"id"`
	ModelSelectionExpression         string                                    `json:"model_selection_expression"`
	OperationName                    string                                    `json:"operation_name"`
	RequestModels                    map[string]string                         `json:"request_models"`
	RouteKey                         string                                    `json:"route_key"`
	RouteResponseSelectionExpression string                                    `json:"route_response_selection_expression"`
	Target                           string                                    `json:"target"`
	RequestParameter                 []apigatewayv2route.RequestParameterState `json:"request_parameter"`
}
