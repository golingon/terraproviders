// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	cognitomanageduserpoolclient "github.com/golingon/terraproviders/aws/5.12.0/cognitomanageduserpoolclient"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCognitoManagedUserPoolClient creates a new instance of [CognitoManagedUserPoolClient].
func NewCognitoManagedUserPoolClient(name string, args CognitoManagedUserPoolClientArgs) *CognitoManagedUserPoolClient {
	return &CognitoManagedUserPoolClient{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CognitoManagedUserPoolClient)(nil)

// CognitoManagedUserPoolClient represents the Terraform resource aws_cognito_managed_user_pool_client.
type CognitoManagedUserPoolClient struct {
	Name      string
	Args      CognitoManagedUserPoolClientArgs
	state     *cognitoManagedUserPoolClientState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CognitoManagedUserPoolClient].
func (cmupc *CognitoManagedUserPoolClient) Type() string {
	return "aws_cognito_managed_user_pool_client"
}

// LocalName returns the local name for [CognitoManagedUserPoolClient].
func (cmupc *CognitoManagedUserPoolClient) LocalName() string {
	return cmupc.Name
}

// Configuration returns the configuration (args) for [CognitoManagedUserPoolClient].
func (cmupc *CognitoManagedUserPoolClient) Configuration() interface{} {
	return cmupc.Args
}

// DependOn is used for other resources to depend on [CognitoManagedUserPoolClient].
func (cmupc *CognitoManagedUserPoolClient) DependOn() terra.Reference {
	return terra.ReferenceResource(cmupc)
}

// Dependencies returns the list of resources [CognitoManagedUserPoolClient] depends_on.
func (cmupc *CognitoManagedUserPoolClient) Dependencies() terra.Dependencies {
	return cmupc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CognitoManagedUserPoolClient].
func (cmupc *CognitoManagedUserPoolClient) LifecycleManagement() *terra.Lifecycle {
	return cmupc.Lifecycle
}

// Attributes returns the attributes for [CognitoManagedUserPoolClient].
func (cmupc *CognitoManagedUserPoolClient) Attributes() cognitoManagedUserPoolClientAttributes {
	return cognitoManagedUserPoolClientAttributes{ref: terra.ReferenceResource(cmupc)}
}

// ImportState imports the given attribute values into [CognitoManagedUserPoolClient]'s state.
func (cmupc *CognitoManagedUserPoolClient) ImportState(av io.Reader) error {
	cmupc.state = &cognitoManagedUserPoolClientState{}
	if err := json.NewDecoder(av).Decode(cmupc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cmupc.Type(), cmupc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CognitoManagedUserPoolClient] has state.
func (cmupc *CognitoManagedUserPoolClient) State() (*cognitoManagedUserPoolClientState, bool) {
	return cmupc.state, cmupc.state != nil
}

// StateMust returns the state for [CognitoManagedUserPoolClient]. Panics if the state is nil.
func (cmupc *CognitoManagedUserPoolClient) StateMust() *cognitoManagedUserPoolClientState {
	if cmupc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cmupc.Type(), cmupc.LocalName()))
	}
	return cmupc.state
}

// CognitoManagedUserPoolClientArgs contains the configurations for aws_cognito_managed_user_pool_client.
type CognitoManagedUserPoolClientArgs struct {
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
	// IdTokenValidity: number, optional
	IdTokenValidity terra.NumberValue `hcl:"id_token_validity,attr"`
	// LogoutUrls: set of string, optional
	LogoutUrls terra.SetValue[terra.StringValue] `hcl:"logout_urls,attr"`
	// NamePattern: string, optional
	NamePattern terra.StringValue `hcl:"name_pattern,attr"`
	// NamePrefix: string, optional
	NamePrefix terra.StringValue `hcl:"name_prefix,attr"`
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
	AnalyticsConfiguration []cognitomanageduserpoolclient.AnalyticsConfiguration `hcl:"analytics_configuration,block" validate:"min=0"`
	// TokenValidityUnits: min=0
	TokenValidityUnits []cognitomanageduserpoolclient.TokenValidityUnits `hcl:"token_validity_units,block" validate:"min=0"`
}
type cognitoManagedUserPoolClientAttributes struct {
	ref terra.Reference
}

// AccessTokenValidity returns a reference to field access_token_validity of aws_cognito_managed_user_pool_client.
func (cmupc cognitoManagedUserPoolClientAttributes) AccessTokenValidity() terra.NumberValue {
	return terra.ReferenceAsNumber(cmupc.ref.Append("access_token_validity"))
}

// AllowedOauthFlows returns a reference to field allowed_oauth_flows of aws_cognito_managed_user_pool_client.
func (cmupc cognitoManagedUserPoolClientAttributes) AllowedOauthFlows() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cmupc.ref.Append("allowed_oauth_flows"))
}

// AllowedOauthFlowsUserPoolClient returns a reference to field allowed_oauth_flows_user_pool_client of aws_cognito_managed_user_pool_client.
func (cmupc cognitoManagedUserPoolClientAttributes) AllowedOauthFlowsUserPoolClient() terra.BoolValue {
	return terra.ReferenceAsBool(cmupc.ref.Append("allowed_oauth_flows_user_pool_client"))
}

// AllowedOauthScopes returns a reference to field allowed_oauth_scopes of aws_cognito_managed_user_pool_client.
func (cmupc cognitoManagedUserPoolClientAttributes) AllowedOauthScopes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cmupc.ref.Append("allowed_oauth_scopes"))
}

// AuthSessionValidity returns a reference to field auth_session_validity of aws_cognito_managed_user_pool_client.
func (cmupc cognitoManagedUserPoolClientAttributes) AuthSessionValidity() terra.NumberValue {
	return terra.ReferenceAsNumber(cmupc.ref.Append("auth_session_validity"))
}

// CallbackUrls returns a reference to field callback_urls of aws_cognito_managed_user_pool_client.
func (cmupc cognitoManagedUserPoolClientAttributes) CallbackUrls() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cmupc.ref.Append("callback_urls"))
}

// ClientSecret returns a reference to field client_secret of aws_cognito_managed_user_pool_client.
func (cmupc cognitoManagedUserPoolClientAttributes) ClientSecret() terra.StringValue {
	return terra.ReferenceAsString(cmupc.ref.Append("client_secret"))
}

// DefaultRedirectUri returns a reference to field default_redirect_uri of aws_cognito_managed_user_pool_client.
func (cmupc cognitoManagedUserPoolClientAttributes) DefaultRedirectUri() terra.StringValue {
	return terra.ReferenceAsString(cmupc.ref.Append("default_redirect_uri"))
}

// EnablePropagateAdditionalUserContextData returns a reference to field enable_propagate_additional_user_context_data of aws_cognito_managed_user_pool_client.
func (cmupc cognitoManagedUserPoolClientAttributes) EnablePropagateAdditionalUserContextData() terra.BoolValue {
	return terra.ReferenceAsBool(cmupc.ref.Append("enable_propagate_additional_user_context_data"))
}

// EnableTokenRevocation returns a reference to field enable_token_revocation of aws_cognito_managed_user_pool_client.
func (cmupc cognitoManagedUserPoolClientAttributes) EnableTokenRevocation() terra.BoolValue {
	return terra.ReferenceAsBool(cmupc.ref.Append("enable_token_revocation"))
}

// ExplicitAuthFlows returns a reference to field explicit_auth_flows of aws_cognito_managed_user_pool_client.
func (cmupc cognitoManagedUserPoolClientAttributes) ExplicitAuthFlows() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cmupc.ref.Append("explicit_auth_flows"))
}

// Id returns a reference to field id of aws_cognito_managed_user_pool_client.
func (cmupc cognitoManagedUserPoolClientAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cmupc.ref.Append("id"))
}

// IdTokenValidity returns a reference to field id_token_validity of aws_cognito_managed_user_pool_client.
func (cmupc cognitoManagedUserPoolClientAttributes) IdTokenValidity() terra.NumberValue {
	return terra.ReferenceAsNumber(cmupc.ref.Append("id_token_validity"))
}

// LogoutUrls returns a reference to field logout_urls of aws_cognito_managed_user_pool_client.
func (cmupc cognitoManagedUserPoolClientAttributes) LogoutUrls() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cmupc.ref.Append("logout_urls"))
}

// Name returns a reference to field name of aws_cognito_managed_user_pool_client.
func (cmupc cognitoManagedUserPoolClientAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cmupc.ref.Append("name"))
}

// NamePattern returns a reference to field name_pattern of aws_cognito_managed_user_pool_client.
func (cmupc cognitoManagedUserPoolClientAttributes) NamePattern() terra.StringValue {
	return terra.ReferenceAsString(cmupc.ref.Append("name_pattern"))
}

// NamePrefix returns a reference to field name_prefix of aws_cognito_managed_user_pool_client.
func (cmupc cognitoManagedUserPoolClientAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(cmupc.ref.Append("name_prefix"))
}

// PreventUserExistenceErrors returns a reference to field prevent_user_existence_errors of aws_cognito_managed_user_pool_client.
func (cmupc cognitoManagedUserPoolClientAttributes) PreventUserExistenceErrors() terra.StringValue {
	return terra.ReferenceAsString(cmupc.ref.Append("prevent_user_existence_errors"))
}

// ReadAttributes returns a reference to field read_attributes of aws_cognito_managed_user_pool_client.
func (cmupc cognitoManagedUserPoolClientAttributes) ReadAttributes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cmupc.ref.Append("read_attributes"))
}

// RefreshTokenValidity returns a reference to field refresh_token_validity of aws_cognito_managed_user_pool_client.
func (cmupc cognitoManagedUserPoolClientAttributes) RefreshTokenValidity() terra.NumberValue {
	return terra.ReferenceAsNumber(cmupc.ref.Append("refresh_token_validity"))
}

// SupportedIdentityProviders returns a reference to field supported_identity_providers of aws_cognito_managed_user_pool_client.
func (cmupc cognitoManagedUserPoolClientAttributes) SupportedIdentityProviders() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cmupc.ref.Append("supported_identity_providers"))
}

// UserPoolId returns a reference to field user_pool_id of aws_cognito_managed_user_pool_client.
func (cmupc cognitoManagedUserPoolClientAttributes) UserPoolId() terra.StringValue {
	return terra.ReferenceAsString(cmupc.ref.Append("user_pool_id"))
}

// WriteAttributes returns a reference to field write_attributes of aws_cognito_managed_user_pool_client.
func (cmupc cognitoManagedUserPoolClientAttributes) WriteAttributes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cmupc.ref.Append("write_attributes"))
}

func (cmupc cognitoManagedUserPoolClientAttributes) AnalyticsConfiguration() terra.ListValue[cognitomanageduserpoolclient.AnalyticsConfigurationAttributes] {
	return terra.ReferenceAsList[cognitomanageduserpoolclient.AnalyticsConfigurationAttributes](cmupc.ref.Append("analytics_configuration"))
}

func (cmupc cognitoManagedUserPoolClientAttributes) TokenValidityUnits() terra.ListValue[cognitomanageduserpoolclient.TokenValidityUnitsAttributes] {
	return terra.ReferenceAsList[cognitomanageduserpoolclient.TokenValidityUnitsAttributes](cmupc.ref.Append("token_validity_units"))
}

type cognitoManagedUserPoolClientState struct {
	AccessTokenValidity                      float64                                                    `json:"access_token_validity"`
	AllowedOauthFlows                        []string                                                   `json:"allowed_oauth_flows"`
	AllowedOauthFlowsUserPoolClient          bool                                                       `json:"allowed_oauth_flows_user_pool_client"`
	AllowedOauthScopes                       []string                                                   `json:"allowed_oauth_scopes"`
	AuthSessionValidity                      float64                                                    `json:"auth_session_validity"`
	CallbackUrls                             []string                                                   `json:"callback_urls"`
	ClientSecret                             string                                                     `json:"client_secret"`
	DefaultRedirectUri                       string                                                     `json:"default_redirect_uri"`
	EnablePropagateAdditionalUserContextData bool                                                       `json:"enable_propagate_additional_user_context_data"`
	EnableTokenRevocation                    bool                                                       `json:"enable_token_revocation"`
	ExplicitAuthFlows                        []string                                                   `json:"explicit_auth_flows"`
	Id                                       string                                                     `json:"id"`
	IdTokenValidity                          float64                                                    `json:"id_token_validity"`
	LogoutUrls                               []string                                                   `json:"logout_urls"`
	Name                                     string                                                     `json:"name"`
	NamePattern                              string                                                     `json:"name_pattern"`
	NamePrefix                               string                                                     `json:"name_prefix"`
	PreventUserExistenceErrors               string                                                     `json:"prevent_user_existence_errors"`
	ReadAttributes                           []string                                                   `json:"read_attributes"`
	RefreshTokenValidity                     float64                                                    `json:"refresh_token_validity"`
	SupportedIdentityProviders               []string                                                   `json:"supported_identity_providers"`
	UserPoolId                               string                                                     `json:"user_pool_id"`
	WriteAttributes                          []string                                                   `json:"write_attributes"`
	AnalyticsConfiguration                   []cognitomanageduserpoolclient.AnalyticsConfigurationState `json:"analytics_configuration"`
	TokenValidityUnits                       []cognitomanageduserpoolclient.TokenValidityUnitsState     `json:"token_validity_units"`
}
