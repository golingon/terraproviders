// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datacognitouserpoolclient "github.com/golingon/terraproviders/aws/5.8.0/datacognitouserpoolclient"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataCognitoUserPoolClient creates a new instance of [DataCognitoUserPoolClient].
func NewDataCognitoUserPoolClient(name string, args DataCognitoUserPoolClientArgs) *DataCognitoUserPoolClient {
	return &DataCognitoUserPoolClient{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCognitoUserPoolClient)(nil)

// DataCognitoUserPoolClient represents the Terraform data resource aws_cognito_user_pool_client.
type DataCognitoUserPoolClient struct {
	Name string
	Args DataCognitoUserPoolClientArgs
}

// DataSource returns the Terraform object type for [DataCognitoUserPoolClient].
func (cupc *DataCognitoUserPoolClient) DataSource() string {
	return "aws_cognito_user_pool_client"
}

// LocalName returns the local name for [DataCognitoUserPoolClient].
func (cupc *DataCognitoUserPoolClient) LocalName() string {
	return cupc.Name
}

// Configuration returns the configuration (args) for [DataCognitoUserPoolClient].
func (cupc *DataCognitoUserPoolClient) Configuration() interface{} {
	return cupc.Args
}

// Attributes returns the attributes for [DataCognitoUserPoolClient].
func (cupc *DataCognitoUserPoolClient) Attributes() dataCognitoUserPoolClientAttributes {
	return dataCognitoUserPoolClientAttributes{ref: terra.ReferenceDataResource(cupc)}
}

// DataCognitoUserPoolClientArgs contains the configurations for aws_cognito_user_pool_client.
type DataCognitoUserPoolClientArgs struct {
	// ClientId: string, required
	ClientId terra.StringValue `hcl:"client_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// UserPoolId: string, required
	UserPoolId terra.StringValue `hcl:"user_pool_id,attr" validate:"required"`
	// AnalyticsConfiguration: min=0
	AnalyticsConfiguration []datacognitouserpoolclient.AnalyticsConfiguration `hcl:"analytics_configuration,block" validate:"min=0"`
	// TokenValidityUnits: min=0
	TokenValidityUnits []datacognitouserpoolclient.TokenValidityUnits `hcl:"token_validity_units,block" validate:"min=0"`
}
type dataCognitoUserPoolClientAttributes struct {
	ref terra.Reference
}

// AccessTokenValidity returns a reference to field access_token_validity of aws_cognito_user_pool_client.
func (cupc dataCognitoUserPoolClientAttributes) AccessTokenValidity() terra.NumberValue {
	return terra.ReferenceAsNumber(cupc.ref.Append("access_token_validity"))
}

// AllowedOauthFlows returns a reference to field allowed_oauth_flows of aws_cognito_user_pool_client.
func (cupc dataCognitoUserPoolClientAttributes) AllowedOauthFlows() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cupc.ref.Append("allowed_oauth_flows"))
}

// AllowedOauthFlowsUserPoolClient returns a reference to field allowed_oauth_flows_user_pool_client of aws_cognito_user_pool_client.
func (cupc dataCognitoUserPoolClientAttributes) AllowedOauthFlowsUserPoolClient() terra.BoolValue {
	return terra.ReferenceAsBool(cupc.ref.Append("allowed_oauth_flows_user_pool_client"))
}

// AllowedOauthScopes returns a reference to field allowed_oauth_scopes of aws_cognito_user_pool_client.
func (cupc dataCognitoUserPoolClientAttributes) AllowedOauthScopes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cupc.ref.Append("allowed_oauth_scopes"))
}

// CallbackUrls returns a reference to field callback_urls of aws_cognito_user_pool_client.
func (cupc dataCognitoUserPoolClientAttributes) CallbackUrls() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cupc.ref.Append("callback_urls"))
}

// ClientId returns a reference to field client_id of aws_cognito_user_pool_client.
func (cupc dataCognitoUserPoolClientAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(cupc.ref.Append("client_id"))
}

// ClientSecret returns a reference to field client_secret of aws_cognito_user_pool_client.
func (cupc dataCognitoUserPoolClientAttributes) ClientSecret() terra.StringValue {
	return terra.ReferenceAsString(cupc.ref.Append("client_secret"))
}

// DefaultRedirectUri returns a reference to field default_redirect_uri of aws_cognito_user_pool_client.
func (cupc dataCognitoUserPoolClientAttributes) DefaultRedirectUri() terra.StringValue {
	return terra.ReferenceAsString(cupc.ref.Append("default_redirect_uri"))
}

// EnablePropagateAdditionalUserContextData returns a reference to field enable_propagate_additional_user_context_data of aws_cognito_user_pool_client.
func (cupc dataCognitoUserPoolClientAttributes) EnablePropagateAdditionalUserContextData() terra.BoolValue {
	return terra.ReferenceAsBool(cupc.ref.Append("enable_propagate_additional_user_context_data"))
}

// EnableTokenRevocation returns a reference to field enable_token_revocation of aws_cognito_user_pool_client.
func (cupc dataCognitoUserPoolClientAttributes) EnableTokenRevocation() terra.BoolValue {
	return terra.ReferenceAsBool(cupc.ref.Append("enable_token_revocation"))
}

// ExplicitAuthFlows returns a reference to field explicit_auth_flows of aws_cognito_user_pool_client.
func (cupc dataCognitoUserPoolClientAttributes) ExplicitAuthFlows() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cupc.ref.Append("explicit_auth_flows"))
}

// GenerateSecret returns a reference to field generate_secret of aws_cognito_user_pool_client.
func (cupc dataCognitoUserPoolClientAttributes) GenerateSecret() terra.BoolValue {
	return terra.ReferenceAsBool(cupc.ref.Append("generate_secret"))
}

// Id returns a reference to field id of aws_cognito_user_pool_client.
func (cupc dataCognitoUserPoolClientAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cupc.ref.Append("id"))
}

// IdTokenValidity returns a reference to field id_token_validity of aws_cognito_user_pool_client.
func (cupc dataCognitoUserPoolClientAttributes) IdTokenValidity() terra.NumberValue {
	return terra.ReferenceAsNumber(cupc.ref.Append("id_token_validity"))
}

// LogoutUrls returns a reference to field logout_urls of aws_cognito_user_pool_client.
func (cupc dataCognitoUserPoolClientAttributes) LogoutUrls() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cupc.ref.Append("logout_urls"))
}

// Name returns a reference to field name of aws_cognito_user_pool_client.
func (cupc dataCognitoUserPoolClientAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cupc.ref.Append("name"))
}

// PreventUserExistenceErrors returns a reference to field prevent_user_existence_errors of aws_cognito_user_pool_client.
func (cupc dataCognitoUserPoolClientAttributes) PreventUserExistenceErrors() terra.StringValue {
	return terra.ReferenceAsString(cupc.ref.Append("prevent_user_existence_errors"))
}

// ReadAttributes returns a reference to field read_attributes of aws_cognito_user_pool_client.
func (cupc dataCognitoUserPoolClientAttributes) ReadAttributes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cupc.ref.Append("read_attributes"))
}

// RefreshTokenValidity returns a reference to field refresh_token_validity of aws_cognito_user_pool_client.
func (cupc dataCognitoUserPoolClientAttributes) RefreshTokenValidity() terra.NumberValue {
	return terra.ReferenceAsNumber(cupc.ref.Append("refresh_token_validity"))
}

// SupportedIdentityProviders returns a reference to field supported_identity_providers of aws_cognito_user_pool_client.
func (cupc dataCognitoUserPoolClientAttributes) SupportedIdentityProviders() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cupc.ref.Append("supported_identity_providers"))
}

// UserPoolId returns a reference to field user_pool_id of aws_cognito_user_pool_client.
func (cupc dataCognitoUserPoolClientAttributes) UserPoolId() terra.StringValue {
	return terra.ReferenceAsString(cupc.ref.Append("user_pool_id"))
}

// WriteAttributes returns a reference to field write_attributes of aws_cognito_user_pool_client.
func (cupc dataCognitoUserPoolClientAttributes) WriteAttributes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cupc.ref.Append("write_attributes"))
}

func (cupc dataCognitoUserPoolClientAttributes) AnalyticsConfiguration() terra.ListValue[datacognitouserpoolclient.AnalyticsConfigurationAttributes] {
	return terra.ReferenceAsList[datacognitouserpoolclient.AnalyticsConfigurationAttributes](cupc.ref.Append("analytics_configuration"))
}

func (cupc dataCognitoUserPoolClientAttributes) TokenValidityUnits() terra.ListValue[datacognitouserpoolclient.TokenValidityUnitsAttributes] {
	return terra.ReferenceAsList[datacognitouserpoolclient.TokenValidityUnitsAttributes](cupc.ref.Append("token_validity_units"))
}
