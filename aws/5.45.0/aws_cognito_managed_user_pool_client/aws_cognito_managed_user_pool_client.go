// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_cognito_managed_user_pool_client

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

// Resource represents the Terraform resource aws_cognito_managed_user_pool_client.
type Resource struct {
	Name      string
	Args      Args
	state     *awsCognitoManagedUserPoolClientState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (acmupc *Resource) Type() string {
	return "aws_cognito_managed_user_pool_client"
}

// LocalName returns the local name for [Resource].
func (acmupc *Resource) LocalName() string {
	return acmupc.Name
}

// Configuration returns the configuration (args) for [Resource].
func (acmupc *Resource) Configuration() interface{} {
	return acmupc.Args
}

// DependOn is used for other resources to depend on [Resource].
func (acmupc *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(acmupc)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (acmupc *Resource) Dependencies() terra.Dependencies {
	return acmupc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (acmupc *Resource) LifecycleManagement() *terra.Lifecycle {
	return acmupc.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (acmupc *Resource) Attributes() awsCognitoManagedUserPoolClientAttributes {
	return awsCognitoManagedUserPoolClientAttributes{ref: terra.ReferenceResource(acmupc)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (acmupc *Resource) ImportState(state io.Reader) error {
	acmupc.state = &awsCognitoManagedUserPoolClientState{}
	if err := json.NewDecoder(state).Decode(acmupc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acmupc.Type(), acmupc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (acmupc *Resource) State() (*awsCognitoManagedUserPoolClientState, bool) {
	return acmupc.state, acmupc.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (acmupc *Resource) StateMust() *awsCognitoManagedUserPoolClientState {
	if acmupc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acmupc.Type(), acmupc.LocalName()))
	}
	return acmupc.state
}

// Args contains the configurations for aws_cognito_managed_user_pool_client.
type Args struct {
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
	AnalyticsConfiguration []AnalyticsConfiguration `hcl:"analytics_configuration,block" validate:"min=0"`
	// TokenValidityUnits: min=0
	TokenValidityUnits []TokenValidityUnits `hcl:"token_validity_units,block" validate:"min=0"`
}

type awsCognitoManagedUserPoolClientAttributes struct {
	ref terra.Reference
}

// AccessTokenValidity returns a reference to field access_token_validity of aws_cognito_managed_user_pool_client.
func (acmupc awsCognitoManagedUserPoolClientAttributes) AccessTokenValidity() terra.NumberValue {
	return terra.ReferenceAsNumber(acmupc.ref.Append("access_token_validity"))
}

// AllowedOauthFlows returns a reference to field allowed_oauth_flows of aws_cognito_managed_user_pool_client.
func (acmupc awsCognitoManagedUserPoolClientAttributes) AllowedOauthFlows() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](acmupc.ref.Append("allowed_oauth_flows"))
}

// AllowedOauthFlowsUserPoolClient returns a reference to field allowed_oauth_flows_user_pool_client of aws_cognito_managed_user_pool_client.
func (acmupc awsCognitoManagedUserPoolClientAttributes) AllowedOauthFlowsUserPoolClient() terra.BoolValue {
	return terra.ReferenceAsBool(acmupc.ref.Append("allowed_oauth_flows_user_pool_client"))
}

// AllowedOauthScopes returns a reference to field allowed_oauth_scopes of aws_cognito_managed_user_pool_client.
func (acmupc awsCognitoManagedUserPoolClientAttributes) AllowedOauthScopes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](acmupc.ref.Append("allowed_oauth_scopes"))
}

// AuthSessionValidity returns a reference to field auth_session_validity of aws_cognito_managed_user_pool_client.
func (acmupc awsCognitoManagedUserPoolClientAttributes) AuthSessionValidity() terra.NumberValue {
	return terra.ReferenceAsNumber(acmupc.ref.Append("auth_session_validity"))
}

// CallbackUrls returns a reference to field callback_urls of aws_cognito_managed_user_pool_client.
func (acmupc awsCognitoManagedUserPoolClientAttributes) CallbackUrls() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](acmupc.ref.Append("callback_urls"))
}

// ClientSecret returns a reference to field client_secret of aws_cognito_managed_user_pool_client.
func (acmupc awsCognitoManagedUserPoolClientAttributes) ClientSecret() terra.StringValue {
	return terra.ReferenceAsString(acmupc.ref.Append("client_secret"))
}

// DefaultRedirectUri returns a reference to field default_redirect_uri of aws_cognito_managed_user_pool_client.
func (acmupc awsCognitoManagedUserPoolClientAttributes) DefaultRedirectUri() terra.StringValue {
	return terra.ReferenceAsString(acmupc.ref.Append("default_redirect_uri"))
}

// EnablePropagateAdditionalUserContextData returns a reference to field enable_propagate_additional_user_context_data of aws_cognito_managed_user_pool_client.
func (acmupc awsCognitoManagedUserPoolClientAttributes) EnablePropagateAdditionalUserContextData() terra.BoolValue {
	return terra.ReferenceAsBool(acmupc.ref.Append("enable_propagate_additional_user_context_data"))
}

// EnableTokenRevocation returns a reference to field enable_token_revocation of aws_cognito_managed_user_pool_client.
func (acmupc awsCognitoManagedUserPoolClientAttributes) EnableTokenRevocation() terra.BoolValue {
	return terra.ReferenceAsBool(acmupc.ref.Append("enable_token_revocation"))
}

// ExplicitAuthFlows returns a reference to field explicit_auth_flows of aws_cognito_managed_user_pool_client.
func (acmupc awsCognitoManagedUserPoolClientAttributes) ExplicitAuthFlows() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](acmupc.ref.Append("explicit_auth_flows"))
}

// Id returns a reference to field id of aws_cognito_managed_user_pool_client.
func (acmupc awsCognitoManagedUserPoolClientAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acmupc.ref.Append("id"))
}

// IdTokenValidity returns a reference to field id_token_validity of aws_cognito_managed_user_pool_client.
func (acmupc awsCognitoManagedUserPoolClientAttributes) IdTokenValidity() terra.NumberValue {
	return terra.ReferenceAsNumber(acmupc.ref.Append("id_token_validity"))
}

// LogoutUrls returns a reference to field logout_urls of aws_cognito_managed_user_pool_client.
func (acmupc awsCognitoManagedUserPoolClientAttributes) LogoutUrls() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](acmupc.ref.Append("logout_urls"))
}

// Name returns a reference to field name of aws_cognito_managed_user_pool_client.
func (acmupc awsCognitoManagedUserPoolClientAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(acmupc.ref.Append("name"))
}

// NamePattern returns a reference to field name_pattern of aws_cognito_managed_user_pool_client.
func (acmupc awsCognitoManagedUserPoolClientAttributes) NamePattern() terra.StringValue {
	return terra.ReferenceAsString(acmupc.ref.Append("name_pattern"))
}

// NamePrefix returns a reference to field name_prefix of aws_cognito_managed_user_pool_client.
func (acmupc awsCognitoManagedUserPoolClientAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(acmupc.ref.Append("name_prefix"))
}

// PreventUserExistenceErrors returns a reference to field prevent_user_existence_errors of aws_cognito_managed_user_pool_client.
func (acmupc awsCognitoManagedUserPoolClientAttributes) PreventUserExistenceErrors() terra.StringValue {
	return terra.ReferenceAsString(acmupc.ref.Append("prevent_user_existence_errors"))
}

// ReadAttributes returns a reference to field read_attributes of aws_cognito_managed_user_pool_client.
func (acmupc awsCognitoManagedUserPoolClientAttributes) ReadAttributes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](acmupc.ref.Append("read_attributes"))
}

// RefreshTokenValidity returns a reference to field refresh_token_validity of aws_cognito_managed_user_pool_client.
func (acmupc awsCognitoManagedUserPoolClientAttributes) RefreshTokenValidity() terra.NumberValue {
	return terra.ReferenceAsNumber(acmupc.ref.Append("refresh_token_validity"))
}

// SupportedIdentityProviders returns a reference to field supported_identity_providers of aws_cognito_managed_user_pool_client.
func (acmupc awsCognitoManagedUserPoolClientAttributes) SupportedIdentityProviders() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](acmupc.ref.Append("supported_identity_providers"))
}

// UserPoolId returns a reference to field user_pool_id of aws_cognito_managed_user_pool_client.
func (acmupc awsCognitoManagedUserPoolClientAttributes) UserPoolId() terra.StringValue {
	return terra.ReferenceAsString(acmupc.ref.Append("user_pool_id"))
}

// WriteAttributes returns a reference to field write_attributes of aws_cognito_managed_user_pool_client.
func (acmupc awsCognitoManagedUserPoolClientAttributes) WriteAttributes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](acmupc.ref.Append("write_attributes"))
}

func (acmupc awsCognitoManagedUserPoolClientAttributes) AnalyticsConfiguration() terra.ListValue[AnalyticsConfigurationAttributes] {
	return terra.ReferenceAsList[AnalyticsConfigurationAttributes](acmupc.ref.Append("analytics_configuration"))
}

func (acmupc awsCognitoManagedUserPoolClientAttributes) TokenValidityUnits() terra.ListValue[TokenValidityUnitsAttributes] {
	return terra.ReferenceAsList[TokenValidityUnitsAttributes](acmupc.ref.Append("token_validity_units"))
}

type awsCognitoManagedUserPoolClientState struct {
	AccessTokenValidity                      float64                       `json:"access_token_validity"`
	AllowedOauthFlows                        []string                      `json:"allowed_oauth_flows"`
	AllowedOauthFlowsUserPoolClient          bool                          `json:"allowed_oauth_flows_user_pool_client"`
	AllowedOauthScopes                       []string                      `json:"allowed_oauth_scopes"`
	AuthSessionValidity                      float64                       `json:"auth_session_validity"`
	CallbackUrls                             []string                      `json:"callback_urls"`
	ClientSecret                             string                        `json:"client_secret"`
	DefaultRedirectUri                       string                        `json:"default_redirect_uri"`
	EnablePropagateAdditionalUserContextData bool                          `json:"enable_propagate_additional_user_context_data"`
	EnableTokenRevocation                    bool                          `json:"enable_token_revocation"`
	ExplicitAuthFlows                        []string                      `json:"explicit_auth_flows"`
	Id                                       string                        `json:"id"`
	IdTokenValidity                          float64                       `json:"id_token_validity"`
	LogoutUrls                               []string                      `json:"logout_urls"`
	Name                                     string                        `json:"name"`
	NamePattern                              string                        `json:"name_pattern"`
	NamePrefix                               string                        `json:"name_prefix"`
	PreventUserExistenceErrors               string                        `json:"prevent_user_existence_errors"`
	ReadAttributes                           []string                      `json:"read_attributes"`
	RefreshTokenValidity                     float64                       `json:"refresh_token_validity"`
	SupportedIdentityProviders               []string                      `json:"supported_identity_providers"`
	UserPoolId                               string                        `json:"user_pool_id"`
	WriteAttributes                          []string                      `json:"write_attributes"`
	AnalyticsConfiguration                   []AnalyticsConfigurationState `json:"analytics_configuration"`
	TokenValidityUnits                       []TokenValidityUnitsState     `json:"token_validity_units"`
}
