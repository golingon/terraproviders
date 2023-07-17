// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiGatewayMethod creates a new instance of [ApiGatewayMethod].
func NewApiGatewayMethod(name string, args ApiGatewayMethodArgs) *ApiGatewayMethod {
	return &ApiGatewayMethod{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiGatewayMethod)(nil)

// ApiGatewayMethod represents the Terraform resource aws_api_gateway_method.
type ApiGatewayMethod struct {
	Name      string
	Args      ApiGatewayMethodArgs
	state     *apiGatewayMethodState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiGatewayMethod].
func (agm *ApiGatewayMethod) Type() string {
	return "aws_api_gateway_method"
}

// LocalName returns the local name for [ApiGatewayMethod].
func (agm *ApiGatewayMethod) LocalName() string {
	return agm.Name
}

// Configuration returns the configuration (args) for [ApiGatewayMethod].
func (agm *ApiGatewayMethod) Configuration() interface{} {
	return agm.Args
}

// DependOn is used for other resources to depend on [ApiGatewayMethod].
func (agm *ApiGatewayMethod) DependOn() terra.Reference {
	return terra.ReferenceResource(agm)
}

// Dependencies returns the list of resources [ApiGatewayMethod] depends_on.
func (agm *ApiGatewayMethod) Dependencies() terra.Dependencies {
	return agm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiGatewayMethod].
func (agm *ApiGatewayMethod) LifecycleManagement() *terra.Lifecycle {
	return agm.Lifecycle
}

// Attributes returns the attributes for [ApiGatewayMethod].
func (agm *ApiGatewayMethod) Attributes() apiGatewayMethodAttributes {
	return apiGatewayMethodAttributes{ref: terra.ReferenceResource(agm)}
}

// ImportState imports the given attribute values into [ApiGatewayMethod]'s state.
func (agm *ApiGatewayMethod) ImportState(av io.Reader) error {
	agm.state = &apiGatewayMethodState{}
	if err := json.NewDecoder(av).Decode(agm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", agm.Type(), agm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiGatewayMethod] has state.
func (agm *ApiGatewayMethod) State() (*apiGatewayMethodState, bool) {
	return agm.state, agm.state != nil
}

// StateMust returns the state for [ApiGatewayMethod]. Panics if the state is nil.
func (agm *ApiGatewayMethod) StateMust() *apiGatewayMethodState {
	if agm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", agm.Type(), agm.LocalName()))
	}
	return agm.state
}

// ApiGatewayMethodArgs contains the configurations for aws_api_gateway_method.
type ApiGatewayMethodArgs struct {
	// ApiKeyRequired: bool, optional
	ApiKeyRequired terra.BoolValue `hcl:"api_key_required,attr"`
	// Authorization: string, required
	Authorization terra.StringValue `hcl:"authorization,attr" validate:"required"`
	// AuthorizationScopes: set of string, optional
	AuthorizationScopes terra.SetValue[terra.StringValue] `hcl:"authorization_scopes,attr"`
	// AuthorizerId: string, optional
	AuthorizerId terra.StringValue `hcl:"authorizer_id,attr"`
	// HttpMethod: string, required
	HttpMethod terra.StringValue `hcl:"http_method,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// OperationName: string, optional
	OperationName terra.StringValue `hcl:"operation_name,attr"`
	// RequestModels: map of string, optional
	RequestModels terra.MapValue[terra.StringValue] `hcl:"request_models,attr"`
	// RequestParameters: map of bool, optional
	RequestParameters terra.MapValue[terra.BoolValue] `hcl:"request_parameters,attr"`
	// RequestValidatorId: string, optional
	RequestValidatorId terra.StringValue `hcl:"request_validator_id,attr"`
	// ResourceId: string, required
	ResourceId terra.StringValue `hcl:"resource_id,attr" validate:"required"`
	// RestApiId: string, required
	RestApiId terra.StringValue `hcl:"rest_api_id,attr" validate:"required"`
}
type apiGatewayMethodAttributes struct {
	ref terra.Reference
}

// ApiKeyRequired returns a reference to field api_key_required of aws_api_gateway_method.
func (agm apiGatewayMethodAttributes) ApiKeyRequired() terra.BoolValue {
	return terra.ReferenceAsBool(agm.ref.Append("api_key_required"))
}

// Authorization returns a reference to field authorization of aws_api_gateway_method.
func (agm apiGatewayMethodAttributes) Authorization() terra.StringValue {
	return terra.ReferenceAsString(agm.ref.Append("authorization"))
}

// AuthorizationScopes returns a reference to field authorization_scopes of aws_api_gateway_method.
func (agm apiGatewayMethodAttributes) AuthorizationScopes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](agm.ref.Append("authorization_scopes"))
}

// AuthorizerId returns a reference to field authorizer_id of aws_api_gateway_method.
func (agm apiGatewayMethodAttributes) AuthorizerId() terra.StringValue {
	return terra.ReferenceAsString(agm.ref.Append("authorizer_id"))
}

// HttpMethod returns a reference to field http_method of aws_api_gateway_method.
func (agm apiGatewayMethodAttributes) HttpMethod() terra.StringValue {
	return terra.ReferenceAsString(agm.ref.Append("http_method"))
}

// Id returns a reference to field id of aws_api_gateway_method.
func (agm apiGatewayMethodAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(agm.ref.Append("id"))
}

// OperationName returns a reference to field operation_name of aws_api_gateway_method.
func (agm apiGatewayMethodAttributes) OperationName() terra.StringValue {
	return terra.ReferenceAsString(agm.ref.Append("operation_name"))
}

// RequestModels returns a reference to field request_models of aws_api_gateway_method.
func (agm apiGatewayMethodAttributes) RequestModels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](agm.ref.Append("request_models"))
}

// RequestParameters returns a reference to field request_parameters of aws_api_gateway_method.
func (agm apiGatewayMethodAttributes) RequestParameters() terra.MapValue[terra.BoolValue] {
	return terra.ReferenceAsMap[terra.BoolValue](agm.ref.Append("request_parameters"))
}

// RequestValidatorId returns a reference to field request_validator_id of aws_api_gateway_method.
func (agm apiGatewayMethodAttributes) RequestValidatorId() terra.StringValue {
	return terra.ReferenceAsString(agm.ref.Append("request_validator_id"))
}

// ResourceId returns a reference to field resource_id of aws_api_gateway_method.
func (agm apiGatewayMethodAttributes) ResourceId() terra.StringValue {
	return terra.ReferenceAsString(agm.ref.Append("resource_id"))
}

// RestApiId returns a reference to field rest_api_id of aws_api_gateway_method.
func (agm apiGatewayMethodAttributes) RestApiId() terra.StringValue {
	return terra.ReferenceAsString(agm.ref.Append("rest_api_id"))
}

type apiGatewayMethodState struct {
	ApiKeyRequired      bool              `json:"api_key_required"`
	Authorization       string            `json:"authorization"`
	AuthorizationScopes []string          `json:"authorization_scopes"`
	AuthorizerId        string            `json:"authorizer_id"`
	HttpMethod          string            `json:"http_method"`
	Id                  string            `json:"id"`
	OperationName       string            `json:"operation_name"`
	RequestModels       map[string]string `json:"request_models"`
	RequestParameters   map[string]bool   `json:"request_parameters"`
	RequestValidatorId  string            `json:"request_validator_id"`
	ResourceId          string            `json:"resource_id"`
	RestApiId           string            `json:"rest_api_id"`
}
