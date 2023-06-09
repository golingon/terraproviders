// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiGatewayAuthorizer creates a new instance of [ApiGatewayAuthorizer].
func NewApiGatewayAuthorizer(name string, args ApiGatewayAuthorizerArgs) *ApiGatewayAuthorizer {
	return &ApiGatewayAuthorizer{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiGatewayAuthorizer)(nil)

// ApiGatewayAuthorizer represents the Terraform resource aws_api_gateway_authorizer.
type ApiGatewayAuthorizer struct {
	Name      string
	Args      ApiGatewayAuthorizerArgs
	state     *apiGatewayAuthorizerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiGatewayAuthorizer].
func (aga *ApiGatewayAuthorizer) Type() string {
	return "aws_api_gateway_authorizer"
}

// LocalName returns the local name for [ApiGatewayAuthorizer].
func (aga *ApiGatewayAuthorizer) LocalName() string {
	return aga.Name
}

// Configuration returns the configuration (args) for [ApiGatewayAuthorizer].
func (aga *ApiGatewayAuthorizer) Configuration() interface{} {
	return aga.Args
}

// DependOn is used for other resources to depend on [ApiGatewayAuthorizer].
func (aga *ApiGatewayAuthorizer) DependOn() terra.Reference {
	return terra.ReferenceResource(aga)
}

// Dependencies returns the list of resources [ApiGatewayAuthorizer] depends_on.
func (aga *ApiGatewayAuthorizer) Dependencies() terra.Dependencies {
	return aga.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiGatewayAuthorizer].
func (aga *ApiGatewayAuthorizer) LifecycleManagement() *terra.Lifecycle {
	return aga.Lifecycle
}

// Attributes returns the attributes for [ApiGatewayAuthorizer].
func (aga *ApiGatewayAuthorizer) Attributes() apiGatewayAuthorizerAttributes {
	return apiGatewayAuthorizerAttributes{ref: terra.ReferenceResource(aga)}
}

// ImportState imports the given attribute values into [ApiGatewayAuthorizer]'s state.
func (aga *ApiGatewayAuthorizer) ImportState(av io.Reader) error {
	aga.state = &apiGatewayAuthorizerState{}
	if err := json.NewDecoder(av).Decode(aga.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aga.Type(), aga.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiGatewayAuthorizer] has state.
func (aga *ApiGatewayAuthorizer) State() (*apiGatewayAuthorizerState, bool) {
	return aga.state, aga.state != nil
}

// StateMust returns the state for [ApiGatewayAuthorizer]. Panics if the state is nil.
func (aga *ApiGatewayAuthorizer) StateMust() *apiGatewayAuthorizerState {
	if aga.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aga.Type(), aga.LocalName()))
	}
	return aga.state
}

// ApiGatewayAuthorizerArgs contains the configurations for aws_api_gateway_authorizer.
type ApiGatewayAuthorizerArgs struct {
	// AuthorizerCredentials: string, optional
	AuthorizerCredentials terra.StringValue `hcl:"authorizer_credentials,attr"`
	// AuthorizerResultTtlInSeconds: number, optional
	AuthorizerResultTtlInSeconds terra.NumberValue `hcl:"authorizer_result_ttl_in_seconds,attr"`
	// AuthorizerUri: string, optional
	AuthorizerUri terra.StringValue `hcl:"authorizer_uri,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IdentitySource: string, optional
	IdentitySource terra.StringValue `hcl:"identity_source,attr"`
	// IdentityValidationExpression: string, optional
	IdentityValidationExpression terra.StringValue `hcl:"identity_validation_expression,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ProviderArns: set of string, optional
	ProviderArns terra.SetValue[terra.StringValue] `hcl:"provider_arns,attr"`
	// RestApiId: string, required
	RestApiId terra.StringValue `hcl:"rest_api_id,attr" validate:"required"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
}
type apiGatewayAuthorizerAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_api_gateway_authorizer.
func (aga apiGatewayAuthorizerAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(aga.ref.Append("arn"))
}

// AuthorizerCredentials returns a reference to field authorizer_credentials of aws_api_gateway_authorizer.
func (aga apiGatewayAuthorizerAttributes) AuthorizerCredentials() terra.StringValue {
	return terra.ReferenceAsString(aga.ref.Append("authorizer_credentials"))
}

// AuthorizerResultTtlInSeconds returns a reference to field authorizer_result_ttl_in_seconds of aws_api_gateway_authorizer.
func (aga apiGatewayAuthorizerAttributes) AuthorizerResultTtlInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(aga.ref.Append("authorizer_result_ttl_in_seconds"))
}

// AuthorizerUri returns a reference to field authorizer_uri of aws_api_gateway_authorizer.
func (aga apiGatewayAuthorizerAttributes) AuthorizerUri() terra.StringValue {
	return terra.ReferenceAsString(aga.ref.Append("authorizer_uri"))
}

// Id returns a reference to field id of aws_api_gateway_authorizer.
func (aga apiGatewayAuthorizerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aga.ref.Append("id"))
}

// IdentitySource returns a reference to field identity_source of aws_api_gateway_authorizer.
func (aga apiGatewayAuthorizerAttributes) IdentitySource() terra.StringValue {
	return terra.ReferenceAsString(aga.ref.Append("identity_source"))
}

// IdentityValidationExpression returns a reference to field identity_validation_expression of aws_api_gateway_authorizer.
func (aga apiGatewayAuthorizerAttributes) IdentityValidationExpression() terra.StringValue {
	return terra.ReferenceAsString(aga.ref.Append("identity_validation_expression"))
}

// Name returns a reference to field name of aws_api_gateway_authorizer.
func (aga apiGatewayAuthorizerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aga.ref.Append("name"))
}

// ProviderArns returns a reference to field provider_arns of aws_api_gateway_authorizer.
func (aga apiGatewayAuthorizerAttributes) ProviderArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](aga.ref.Append("provider_arns"))
}

// RestApiId returns a reference to field rest_api_id of aws_api_gateway_authorizer.
func (aga apiGatewayAuthorizerAttributes) RestApiId() terra.StringValue {
	return terra.ReferenceAsString(aga.ref.Append("rest_api_id"))
}

// Type returns a reference to field type of aws_api_gateway_authorizer.
func (aga apiGatewayAuthorizerAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(aga.ref.Append("type"))
}

type apiGatewayAuthorizerState struct {
	Arn                          string   `json:"arn"`
	AuthorizerCredentials        string   `json:"authorizer_credentials"`
	AuthorizerResultTtlInSeconds float64  `json:"authorizer_result_ttl_in_seconds"`
	AuthorizerUri                string   `json:"authorizer_uri"`
	Id                           string   `json:"id"`
	IdentitySource               string   `json:"identity_source"`
	IdentityValidationExpression string   `json:"identity_validation_expression"`
	Name                         string   `json:"name"`
	ProviderArns                 []string `json:"provider_arns"`
	RestApiId                    string   `json:"rest_api_id"`
	Type                         string   `json:"type"`
}
