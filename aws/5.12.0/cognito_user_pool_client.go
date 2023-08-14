// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	cognitouserpoolclient "github.com/golingon/terraproviders/aws/5.12.0/cognitouserpoolclient"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCognitoUserPoolClient creates a new instance of [CognitoUserPoolClient].
func NewCognitoUserPoolClient(name string, args CognitoUserPoolClientArgs) *CognitoUserPoolClient {
	return &CognitoUserPoolClient{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CognitoUserPoolClient)(nil)

// CognitoUserPoolClient represents the Terraform resource aws_cognito_user_pool_client.
type CognitoUserPoolClient struct {
	Name      string
	Args      CognitoUserPoolClientArgs
	state     *cognitoUserPoolClientState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CognitoUserPoolClient].
func (cupc *CognitoUserPoolClient) Type() string {
	return "aws_cognito_user_pool_client"
}

// LocalName returns the local name for [CognitoUserPoolClient].
func (cupc *CognitoUserPoolClient) LocalName() string {
	return cupc.Name
}

// Configuration returns the configuration (args) for [CognitoUserPoolClient].
func (cupc *CognitoUserPoolClient) Configuration() interface{} {
	return cupc.Args
}

// DependOn is used for other resources to depend on [CognitoUserPoolClient].
func (cupc *CognitoUserPoolClient) DependOn() terra.Reference {
	return terra.ReferenceResource(cupc)
}

// Dependencies returns the list of resources [CognitoUserPoolClient] depends_on.
func (cupc *CognitoUserPoolClient) Dependencies() terra.Dependencies {
	return cupc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CognitoUserPoolClient].
func (cupc *CognitoUserPoolClient) LifecycleManagement() *terra.Lifecycle {
	return cupc.Lifecycle
}

// Attributes returns the attributes for [CognitoUserPoolClient].
func (cupc *CognitoUserPoolClient) Attributes() cognitoUserPoolClientAttributes {
	return cognitoUserPoolClientAttributes{ref: terra.ReferenceResource(cupc)}
}

// ImportState imports the given attribute values into [CognitoUserPoolClient]'s state.
func (cupc *CognitoUserPoolClient) ImportState(av io.Reader) error {
	cupc.state = &cognitoUserPoolClientState{}
	if err := json.NewDecoder(av).Decode(cupc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cupc.Type(), cupc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CognitoUserPoolClient] has state.
func (cupc *CognitoUserPoolClient) State() (*cognitoUserPoolClientState, bool) {
	return cupc.state, cupc.state != nil
}

// StateMust returns the state for [CognitoUserPoolClient]. Panics if the state is nil.
func (cupc *CognitoUserPoolClient) StateMust() *cognitoUserPoolClientState {
	if cupc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cupc.Type(), cupc.LocalName()))
	}
	return cupc.state
}

// CognitoUserPoolClientArgs contains the configurations for aws_cognito_user_pool_client.
type CognitoUserPoolClientArgs struct {
	// AccessTokenValidity: number, optional
	AccessTokenValidity terra.NumberValue `hcl:"access_token_validity,attr"`
	// AllowedOauthFlows: set of string, optional
	AllowedOauthFlows terra.SetValue[terra.StringValue] `hcl:"allowed_oauth_flows,attr"`
	// AllowedOauthFlowsUserPoolClient: bool, optional
	AllowedOauthFlowsUserPoolClient terra.BoolValue `hcl:"allowed_oauth_flows_user_pool_client,attr"`
	// AllowedOauthScopes: set of string, optional
	AllowedOauthScopes terra.SetValue[terra.StringValue] `hcl:"allowed_oauth_scopes,attr"`
	// AuthSessionValidity: number, optional
	AuthSessionValidity terra.NumberValue `hcl:"auth_session_validity,attr"`
	// CallbackUrls: set of string, optional
	CallbackUrls terra.SetValue[terra.StringValue] `hcl:"callback_urls,attr"`
	// DefaultRedirectUri: string, optional
	DefaultRedirectUri terra.StringValue `hcl:"default_redirect_uri,attr"`
	// EnablePropagateAdditionalUserContextData: bool, optional
	EnablePropagateAdditionalUserContextData terra.BoolValue `hcl:"enable_propagate_additional_user_context_data,attr"`
	// EnableTokenRevocation: bool, optional
	EnableTokenRevocation terra.BoolValue `hcl:"enable_token_revocation,attr"`
	// ExplicitAuthFlows: set of string, optional
	ExplicitAuthFlows terra.SetValue[terra.StringValue] `hcl:"explicit_auth_flows,attr"`
	// GenerateSecret: bool, optional
	GenerateSecret terra.BoolValue `hcl:"generate_secret,attr"`
	// IdTokenValidity: number, optional
	IdTokenValidity terra.NumberValue `hcl:"id_token_validity,attr"`
	// LogoutUrls: set of string, optional
	LogoutUrls terra.SetValue[terra.StringValue] `hcl:"logout_urls,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PreventUserExistenceErrors: string, optional
	PreventUserExistenceErrors terra.StringValue `hcl:"prevent_user_existence_errors,attr"`
	// ReadAttributes: set of string, optional
	ReadAttributes terra.SetValue[terra.StringValue] `hcl:"read_attributes,attr"`
	// RefreshTokenValidity: number, optional
	RefreshTokenValidity terra.NumberValue `hcl:"refresh_token_validity,attr"`
	// SupportedIdentityProviders: set of string, optional
	SupportedIdentityProviders terra.SetValue[terra.StringValue] `hcl:"supported_identity_providers,attr"`
	// UserPoolId: string, required
	UserPoolId terra.StringValue `hcl:"user_pool_id,attr" validate:"required"`
	// WriteAttributes: set of string, optional
	WriteAttributes terra.SetValue[terra.StringValue] `hcl:"write_attributes,attr"`
	// AnalyticsConfiguration: min=0
	AnalyticsConfiguration []cognitouserpoolclient.AnalyticsConfiguration `hcl:"analytics_configuration,block" validate:"min=0"`
	// TokenValidityUnits: min=0
	TokenValidityUnits []cognitouserpoolclient.TokenValidityUnits `hcl:"token_validity_units,block" validate:"min=0"`
}
type cognitoUserPoolClientAttributes struct {
	ref terra.Reference
}

// AccessTokenValidity returns a reference to field access_token_validity of aws_cognito_user_pool_client.
func (cupc cognitoUserPoolClientAttributes) AccessTokenValidity() terra.NumberValue {
	return terra.ReferenceAsNumber(cupc.ref.Append("access_token_validity"))
}

// AllowedOauthFlows returns a reference to field allowed_oauth_flows of aws_cognito_user_pool_client.
func (cupc cognitoUserPoolClientAttributes) AllowedOauthFlows() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cupc.ref.Append("allowed_oauth_flows"))
}

// AllowedOauthFlowsUserPoolClient returns a reference to field allowed_oauth_flows_user_pool_client of aws_cognito_user_pool_client.
func (cupc cognitoUserPoolClientAttributes) AllowedOauthFlowsUserPoolClient() terra.BoolValue {
	return terra.ReferenceAsBool(cupc.ref.Append("allowed_oauth_flows_user_pool_client"))
}

// AllowedOauthScopes returns a reference to field allowed_oauth_scopes of aws_cognito_user_pool_client.
func (cupc cognitoUserPoolClientAttributes) AllowedOauthScopes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cupc.ref.Append("allowed_oauth_scopes"))
}

// AuthSessionValidity returns a reference to field auth_session_validity of aws_cognito_user_pool_client.
func (cupc cognitoUserPoolClientAttributes) AuthSessionValidity() terra.NumberValue {
	return terra.ReferenceAsNumber(cupc.ref.Append("auth_session_validity"))
}

// CallbackUrls returns a reference to field callback_urls of aws_cognito_user_pool_client.
func (cupc cognitoUserPoolClientAttributes) CallbackUrls() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cupc.ref.Append("callback_urls"))
}

// ClientSecret returns a reference to field client_secret of aws_cognito_user_pool_client.
func (cupc cognitoUserPoolClientAttributes) ClientSecret() terra.StringValue {
	return terra.ReferenceAsString(cupc.ref.Append("client_secret"))
}

// DefaultRedirectUri returns a reference to field default_redirect_uri of aws_cognito_user_pool_client.
func (cupc cognitoUserPoolClientAttributes) DefaultRedirectUri() terra.StringValue {
	return terra.ReferenceAsString(cupc.ref.Append("default_redirect_uri"))
}

// EnablePropagateAdditionalUserContextData returns a reference to field enable_propagate_additional_user_context_data of aws_cognito_user_pool_client.
func (cupc cognitoUserPoolClientAttributes) EnablePropagateAdditionalUserContextData() terra.BoolValue {
	return terra.ReferenceAsBool(cupc.ref.Append("enable_propagate_additional_user_context_data"))
}

// EnableTokenRevocation returns a reference to field enable_token_revocation of aws_cognito_user_pool_client.
func (cupc cognitoUserPoolClientAttributes) EnableTokenRevocation() terra.BoolValue {
	return terra.ReferenceAsBool(cupc.ref.Append("enable_token_revocation"))
}

// ExplicitAuthFlows returns a reference to field explicit_auth_flows of aws_cognito_user_pool_client.
func (cupc cognitoUserPoolClientAttributes) ExplicitAuthFlows() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cupc.ref.Append("explicit_auth_flows"))
}

// GenerateSecret returns a reference to field generate_secret of aws_cognito_user_pool_client.
func (cupc cognitoUserPoolClientAttributes) GenerateSecret() terra.BoolValue {
	return terra.ReferenceAsBool(cupc.ref.Append("generate_secret"))
}

// Id returns a reference to field id of aws_cognito_user_pool_client.
func (cupc cognitoUserPoolClientAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cupc.ref.Append("id"))
}

// IdTokenValidity returns a reference to field id_token_validity of aws_cognito_user_pool_client.
func (cupc cognitoUserPoolClientAttributes) IdTokenValidity() terra.NumberValue {
	return terra.ReferenceAsNumber(cupc.ref.Append("id_token_validity"))
}

// LogoutUrls returns a reference to field logout_urls of aws_cognito_user_pool_client.
func (cupc cognitoUserPoolClientAttributes) LogoutUrls() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cupc.ref.Append("logout_urls"))
}

// Name returns a reference to field name of aws_cognito_user_pool_client.
func (cupc cognitoUserPoolClientAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cupc.ref.Append("name"))
}

// PreventUserExistenceErrors returns a reference to field prevent_user_existence_errors of aws_cognito_user_pool_client.
func (cupc cognitoUserPoolClientAttributes) PreventUserExistenceErrors() terra.StringValue {
	return terra.ReferenceAsString(cupc.ref.Append("prevent_user_existence_errors"))
}

// ReadAttributes returns a reference to field read_attributes of aws_cognito_user_pool_client.
func (cupc cognitoUserPoolClientAttributes) ReadAttributes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cupc.ref.Append("read_attributes"))
}

// RefreshTokenValidity returns a reference to field refresh_token_validity of aws_cognito_user_pool_client.
func (cupc cognitoUserPoolClientAttributes) RefreshTokenValidity() terra.NumberValue {
	return terra.ReferenceAsNumber(cupc.ref.Append("refresh_token_validity"))
}

// SupportedIdentityProviders returns a reference to field supported_identity_providers of aws_cognito_user_pool_client.
func (cupc cognitoUserPoolClientAttributes) SupportedIdentityProviders() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cupc.ref.Append("supported_identity_providers"))
}

// UserPoolId returns a reference to field user_pool_id of aws_cognito_user_pool_client.
func (cupc cognitoUserPoolClientAttributes) UserPoolId() terra.StringValue {
	return terra.ReferenceAsString(cupc.ref.Append("user_pool_id"))
}

// WriteAttributes returns a reference to field write_attributes of aws_cognito_user_pool_client.
func (cupc cognitoUserPoolClientAttributes) WriteAttributes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cupc.ref.Append("write_attributes"))
}

func (cupc cognitoUserPoolClientAttributes) AnalyticsConfiguration() terra.ListValue[cognitouserpoolclient.AnalyticsConfigurationAttributes] {
	return terra.ReferenceAsList[cognitouserpoolclient.AnalyticsConfigurationAttributes](cupc.ref.Append("analytics_configuration"))
}

func (cupc cognitoUserPoolClientAttributes) TokenValidityUnits() terra.ListValue[cognitouserpoolclient.TokenValidityUnitsAttributes] {
	return terra.ReferenceAsList[cognitouserpoolclient.TokenValidityUnitsAttributes](cupc.ref.Append("token_validity_units"))
}

type cognitoUserPoolClientState struct {
	AccessTokenValidity                      float64                                             `json:"access_token_validity"`
	AllowedOauthFlows                        []string                                            `json:"allowed_oauth_flows"`
	AllowedOauthFlowsUserPoolClient          bool                                                `json:"allowed_oauth_flows_user_pool_client"`
	AllowedOauthScopes                       []string                                            `json:"allowed_oauth_scopes"`
	AuthSessionValidity                      float64                                             `json:"auth_session_validity"`
	CallbackUrls                             []string                                            `json:"callback_urls"`
	ClientSecret                             string                                              `json:"client_secret"`
	DefaultRedirectUri                       string                                              `json:"default_redirect_uri"`
	EnablePropagateAdditionalUserContextData bool                                                `json:"enable_propagate_additional_user_context_data"`
	EnableTokenRevocation                    bool                                                `json:"enable_token_revocation"`
	ExplicitAuthFlows                        []string                                            `json:"explicit_auth_flows"`
	GenerateSecret                           bool                                                `json:"generate_secret"`
	Id                                       string                                              `json:"id"`
	IdTokenValidity                          float64                                             `json:"id_token_validity"`
	LogoutUrls                               []string                                            `json:"logout_urls"`
	Name                                     string                                              `json:"name"`
	PreventUserExistenceErrors               string                                              `json:"prevent_user_existence_errors"`
	ReadAttributes                           []string                                            `json:"read_attributes"`
	RefreshTokenValidity                     float64                                             `json:"refresh_token_validity"`
	SupportedIdentityProviders               []string                                            `json:"supported_identity_providers"`
	UserPoolId                               string                                              `json:"user_pool_id"`
	WriteAttributes                          []string                                            `json:"write_attributes"`
	AnalyticsConfiguration                   []cognitouserpoolclient.AnalyticsConfigurationState `json:"analytics_configuration"`
	TokenValidityUnits                       []cognitouserpoolclient.TokenValidityUnitsState     `json:"token_validity_units"`
}
