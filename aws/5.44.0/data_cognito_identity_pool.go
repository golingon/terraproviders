// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"github.com/golingon/lingon/pkg/terra"
	datacognitoidentitypool "github.com/golingon/terraproviders/aws/5.44.0/datacognitoidentitypool"
)

// NewDataCognitoIdentityPool creates a new instance of [DataCognitoIdentityPool].
func NewDataCognitoIdentityPool(name string, args DataCognitoIdentityPoolArgs) *DataCognitoIdentityPool {
	return &DataCognitoIdentityPool{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCognitoIdentityPool)(nil)

// DataCognitoIdentityPool represents the Terraform data resource aws_cognito_identity_pool.
type DataCognitoIdentityPool struct {
	Name string
	Args DataCognitoIdentityPoolArgs
}

// DataSource returns the Terraform object type for [DataCognitoIdentityPool].
func (cip *DataCognitoIdentityPool) DataSource() string {
	return "aws_cognito_identity_pool"
}

// LocalName returns the local name for [DataCognitoIdentityPool].
func (cip *DataCognitoIdentityPool) LocalName() string {
	return cip.Name
}

// Configuration returns the configuration (args) for [DataCognitoIdentityPool].
func (cip *DataCognitoIdentityPool) Configuration() interface{} {
	return cip.Args
}

// Attributes returns the attributes for [DataCognitoIdentityPool].
func (cip *DataCognitoIdentityPool) Attributes() dataCognitoIdentityPoolAttributes {
	return dataCognitoIdentityPoolAttributes{ref: terra.ReferenceDataResource(cip)}
}

// DataCognitoIdentityPoolArgs contains the configurations for aws_cognito_identity_pool.
type DataCognitoIdentityPoolArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IdentityPoolName: string, required
	IdentityPoolName terra.StringValue `hcl:"identity_pool_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// CognitoIdentityProviders: min=0
	CognitoIdentityProviders []datacognitoidentitypool.CognitoIdentityProviders `hcl:"cognito_identity_providers,block" validate:"min=0"`
}
type dataCognitoIdentityPoolAttributes struct {
	ref terra.Reference
}

// AllowClassicFlow returns a reference to field allow_classic_flow of aws_cognito_identity_pool.
func (cip dataCognitoIdentityPoolAttributes) AllowClassicFlow() terra.BoolValue {
	return terra.ReferenceAsBool(cip.ref.Append("allow_classic_flow"))
}

// AllowUnauthenticatedIdentities returns a reference to field allow_unauthenticated_identities of aws_cognito_identity_pool.
func (cip dataCognitoIdentityPoolAttributes) AllowUnauthenticatedIdentities() terra.BoolValue {
	return terra.ReferenceAsBool(cip.ref.Append("allow_unauthenticated_identities"))
}

// Arn returns a reference to field arn of aws_cognito_identity_pool.
func (cip dataCognitoIdentityPoolAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cip.ref.Append("arn"))
}

// DeveloperProviderName returns a reference to field developer_provider_name of aws_cognito_identity_pool.
func (cip dataCognitoIdentityPoolAttributes) DeveloperProviderName() terra.StringValue {
	return terra.ReferenceAsString(cip.ref.Append("developer_provider_name"))
}

// Id returns a reference to field id of aws_cognito_identity_pool.
func (cip dataCognitoIdentityPoolAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cip.ref.Append("id"))
}

// IdentityPoolName returns a reference to field identity_pool_name of aws_cognito_identity_pool.
func (cip dataCognitoIdentityPoolAttributes) IdentityPoolName() terra.StringValue {
	return terra.ReferenceAsString(cip.ref.Append("identity_pool_name"))
}

// OpenidConnectProviderArns returns a reference to field openid_connect_provider_arns of aws_cognito_identity_pool.
func (cip dataCognitoIdentityPoolAttributes) OpenidConnectProviderArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cip.ref.Append("openid_connect_provider_arns"))
}

// SamlProviderArns returns a reference to field saml_provider_arns of aws_cognito_identity_pool.
func (cip dataCognitoIdentityPoolAttributes) SamlProviderArns() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cip.ref.Append("saml_provider_arns"))
}

// SupportedLoginProviders returns a reference to field supported_login_providers of aws_cognito_identity_pool.
func (cip dataCognitoIdentityPoolAttributes) SupportedLoginProviders() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cip.ref.Append("supported_login_providers"))
}

// Tags returns a reference to field tags of aws_cognito_identity_pool.
func (cip dataCognitoIdentityPoolAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cip.ref.Append("tags"))
}

func (cip dataCognitoIdentityPoolAttributes) CognitoIdentityProviders() terra.SetValue[datacognitoidentitypool.CognitoIdentityProvidersAttributes] {
	return terra.ReferenceAsSet[datacognitoidentitypool.CognitoIdentityProvidersAttributes](cip.ref.Append("cognito_identity_providers"))
}
