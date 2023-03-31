// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package sagemakerworkforce

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type CognitoConfig struct {
	// ClientId: string, required
	ClientId terra.StringValue `hcl:"client_id,attr" validate:"required"`
	// UserPool: string, required
	UserPool terra.StringValue `hcl:"user_pool,attr" validate:"required"`
}

type OidcConfig struct {
	// AuthorizationEndpoint: string, required
	AuthorizationEndpoint terra.StringValue `hcl:"authorization_endpoint,attr" validate:"required"`
	// ClientId: string, required
	ClientId terra.StringValue `hcl:"client_id,attr" validate:"required"`
	// ClientSecret: string, required
	ClientSecret terra.StringValue `hcl:"client_secret,attr" validate:"required"`
	// Issuer: string, required
	Issuer terra.StringValue `hcl:"issuer,attr" validate:"required"`
	// JwksUri: string, required
	JwksUri terra.StringValue `hcl:"jwks_uri,attr" validate:"required"`
	// LogoutEndpoint: string, required
	LogoutEndpoint terra.StringValue `hcl:"logout_endpoint,attr" validate:"required"`
	// TokenEndpoint: string, required
	TokenEndpoint terra.StringValue `hcl:"token_endpoint,attr" validate:"required"`
	// UserInfoEndpoint: string, required
	UserInfoEndpoint terra.StringValue `hcl:"user_info_endpoint,attr" validate:"required"`
}

type SourceIpConfig struct {
	// Cidrs: set of string, required
	Cidrs terra.SetValue[terra.StringValue] `hcl:"cidrs,attr" validate:"required"`
}

type WorkforceVpcConfig struct {
	// SecurityGroupIds: set of string, optional
	SecurityGroupIds terra.SetValue[terra.StringValue] `hcl:"security_group_ids,attr"`
	// Subnets: set of string, optional
	Subnets terra.SetValue[terra.StringValue] `hcl:"subnets,attr"`
	// VpcId: string, optional
	VpcId terra.StringValue `hcl:"vpc_id,attr"`
}

type CognitoConfigAttributes struct {
	ref terra.Reference
}

func (cc CognitoConfigAttributes) InternalRef() terra.Reference {
	return cc.ref
}

func (cc CognitoConfigAttributes) InternalWithRef(ref terra.Reference) CognitoConfigAttributes {
	return CognitoConfigAttributes{ref: ref}
}

func (cc CognitoConfigAttributes) InternalTokens() hclwrite.Tokens {
	return cc.ref.InternalTokens()
}

func (cc CognitoConfigAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("client_id"))
}

func (cc CognitoConfigAttributes) UserPool() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("user_pool"))
}

type OidcConfigAttributes struct {
	ref terra.Reference
}

func (oc OidcConfigAttributes) InternalRef() terra.Reference {
	return oc.ref
}

func (oc OidcConfigAttributes) InternalWithRef(ref terra.Reference) OidcConfigAttributes {
	return OidcConfigAttributes{ref: ref}
}

func (oc OidcConfigAttributes) InternalTokens() hclwrite.Tokens {
	return oc.ref.InternalTokens()
}

func (oc OidcConfigAttributes) AuthorizationEndpoint() terra.StringValue {
	return terra.ReferenceAsString(oc.ref.Append("authorization_endpoint"))
}

func (oc OidcConfigAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(oc.ref.Append("client_id"))
}

func (oc OidcConfigAttributes) ClientSecret() terra.StringValue {
	return terra.ReferenceAsString(oc.ref.Append("client_secret"))
}

func (oc OidcConfigAttributes) Issuer() terra.StringValue {
	return terra.ReferenceAsString(oc.ref.Append("issuer"))
}

func (oc OidcConfigAttributes) JwksUri() terra.StringValue {
	return terra.ReferenceAsString(oc.ref.Append("jwks_uri"))
}

func (oc OidcConfigAttributes) LogoutEndpoint() terra.StringValue {
	return terra.ReferenceAsString(oc.ref.Append("logout_endpoint"))
}

func (oc OidcConfigAttributes) TokenEndpoint() terra.StringValue {
	return terra.ReferenceAsString(oc.ref.Append("token_endpoint"))
}

func (oc OidcConfigAttributes) UserInfoEndpoint() terra.StringValue {
	return terra.ReferenceAsString(oc.ref.Append("user_info_endpoint"))
}

type SourceIpConfigAttributes struct {
	ref terra.Reference
}

func (sic SourceIpConfigAttributes) InternalRef() terra.Reference {
	return sic.ref
}

func (sic SourceIpConfigAttributes) InternalWithRef(ref terra.Reference) SourceIpConfigAttributes {
	return SourceIpConfigAttributes{ref: ref}
}

func (sic SourceIpConfigAttributes) InternalTokens() hclwrite.Tokens {
	return sic.ref.InternalTokens()
}

func (sic SourceIpConfigAttributes) Cidrs() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sic.ref.Append("cidrs"))
}

type WorkforceVpcConfigAttributes struct {
	ref terra.Reference
}

func (wvc WorkforceVpcConfigAttributes) InternalRef() terra.Reference {
	return wvc.ref
}

func (wvc WorkforceVpcConfigAttributes) InternalWithRef(ref terra.Reference) WorkforceVpcConfigAttributes {
	return WorkforceVpcConfigAttributes{ref: ref}
}

func (wvc WorkforceVpcConfigAttributes) InternalTokens() hclwrite.Tokens {
	return wvc.ref.InternalTokens()
}

func (wvc WorkforceVpcConfigAttributes) SecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](wvc.ref.Append("security_group_ids"))
}

func (wvc WorkforceVpcConfigAttributes) Subnets() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](wvc.ref.Append("subnets"))
}

func (wvc WorkforceVpcConfigAttributes) VpcEndpointId() terra.StringValue {
	return terra.ReferenceAsString(wvc.ref.Append("vpc_endpoint_id"))
}

func (wvc WorkforceVpcConfigAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(wvc.ref.Append("vpc_id"))
}

type CognitoConfigState struct {
	ClientId string `json:"client_id"`
	UserPool string `json:"user_pool"`
}

type OidcConfigState struct {
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	ClientId              string `json:"client_id"`
	ClientSecret          string `json:"client_secret"`
	Issuer                string `json:"issuer"`
	JwksUri               string `json:"jwks_uri"`
	LogoutEndpoint        string `json:"logout_endpoint"`
	TokenEndpoint         string `json:"token_endpoint"`
	UserInfoEndpoint      string `json:"user_info_endpoint"`
}

type SourceIpConfigState struct {
	Cidrs []string `json:"cidrs"`
}

type WorkforceVpcConfigState struct {
	SecurityGroupIds []string `json:"security_group_ids"`
	Subnets          []string `json:"subnets"`
	VpcEndpointId    string   `json:"vpc_endpoint_id"`
	VpcId            string   `json:"vpc_id"`
}