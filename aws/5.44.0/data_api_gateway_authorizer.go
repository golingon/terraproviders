// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import "github.com/golingon/lingon/pkg/terra"

// NewDataApiGatewayAuthorizer creates a new instance of [DataApiGatewayAuthorizer].
func NewDataApiGatewayAuthorizer(name string, args DataApiGatewayAuthorizerArgs) *DataApiGatewayAuthorizer {
	return &DataApiGatewayAuthorizer{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataApiGatewayAuthorizer)(nil)

// DataApiGatewayAuthorizer represents the Terraform data resource aws_api_gateway_authorizer.
type DataApiGatewayAuthorizer struct {
	Name string
	Args DataApiGatewayAuthorizerArgs
}

// DataSource returns the Terraform object type for [DataApiGatewayAuthorizer].
func (aga *DataApiGatewayAuthorizer) DataSource() string {
	return "aws_api_gateway_authorizer"
}

// LocalName returns the local name for [DataApiGatewayAuthorizer].
func (aga *DataApiGatewayAuthorizer) LocalName() string {
	return aga.Name
}

// Configuration returns the configuration (args) for [DataApiGatewayAuthorizer].
func (aga *DataApiGatewayAuthorizer) Configuration() interface{} {
	return aga.Args
}

// Attributes returns the attributes for [DataApiGatewayAuthorizer].
func (aga *DataApiGatewayAuthorizer) Attributes() dataApiGatewayAuthorizerAttributes {
	return dataApiGatewayAuthorizerAttributes{ref: terra.ReferenceDataResource(aga)}
}

// DataApiGatewayAuthorizerArgs contains the configurations for aws_api_gateway_authorizer.
type DataApiGatewayAuthorizerArgs struct {
	// AuthorizerId: string, required
	AuthorizerId terra.StringValue `hcl:"authorizer_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RestApiId: string, required
	RestApiId terra.StringValue `hcl:"rest_api_id,attr" validate:"required"`
}
type dataApiGatewayAuthorizerAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_api_gateway_authorizer.
func (aga dataApiGatewayAuthorizerAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(aga.ref.Append("arn"))
}

// AuthorizerCredentials returns a reference to field authorizer_credentials of aws_api_gateway_authorizer.
func (aga dataApiGatewayAuthorizerAttributes) AuthorizerCredentials() terra.StringValue {
	return terra.ReferenceAsString(aga.ref.Append("authorizer_credentials"))
}

// AuthorizerId returns a reference to field authorizer_id of aws_api_gateway_authorizer.
func (aga dataApiGatewayAuthorizerAttributes) AuthorizerId() terra.StringValue {
	return terra.ReferenceAsString(aga.ref.Append("authorizer_id"))
}

// AuthorizerResultTtlInSeconds returns a reference to field authorizer_result_ttl_in_seconds of aws_api_gateway_authorizer.
func (aga dataApiGatewayAuthorizerAttributes) AuthorizerResultTtlInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(aga.ref.Append("authorizer_result_ttl_in_seconds"))
}

// AuthorizerUri returns a reference to field authorizer_uri of aws_api_gateway_authorizer.
func (aga dataApiGatewayAuthorizerAttributes) AuthorizerUri() terra.StringValue {
	return terra.ReferenceAsString(aga.ref.Append("authorizer_uri"))
}

// Id returns a reference to field id of aws_api_gateway_authorizer.
func (aga dataApiGatewayAuthorizerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aga.ref.Append("id"))
}

// IdentitySource returns a reference to field identity_source of aws_api_gateway_authorizer.
func (aga dataApiGatewayAuthorizerAttributes) IdentitySource() terra.StringValue {
	return terra.ReferenceAsString(aga.ref.Append("identity_source"))
}

// IdentityValidationExpression returns a reference to field identity_validation_expression of aws_api_gateway_authorizer.
func (aga dataApiGatewayAuthorizerAttributes) IdentityValidationExpression() terra.StringValue {
	return terra.ReferenceAsString(aga.ref.Append("identity_validation_expression"))
}

// Name returns a reference to field name of aws_api_gateway_authorizer.
func (aga dataApiGatewayAuthorizerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aga.ref.Append("name"))
}

// ProviderArns returns a reference to field provider_arns of aws_api_gateway_authorizer.
func (aga dataApiGatewayAuthorizerAttributes) ProviderArns() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](aga.ref.Append("provider_arns"))
}

// RestApiId returns a reference to field rest_api_id of aws_api_gateway_authorizer.
func (aga dataApiGatewayAuthorizerAttributes) RestApiId() terra.StringValue {
	return terra.ReferenceAsString(aga.ref.Append("rest_api_id"))
}

// Type returns a reference to field type of aws_api_gateway_authorizer.
func (aga dataApiGatewayAuthorizerAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(aga.ref.Append("type"))
}
