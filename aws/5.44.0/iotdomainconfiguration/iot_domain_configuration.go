// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package iotdomainconfiguration

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type AuthorizerConfig struct {
	// AllowAuthorizerOverride: bool, optional
	AllowAuthorizerOverride terra.BoolValue `hcl:"allow_authorizer_override,attr"`
	// DefaultAuthorizerName: string, optional
	DefaultAuthorizerName terra.StringValue `hcl:"default_authorizer_name,attr"`
}

type TlsConfig struct {
	// SecurityPolicy: string, optional
	SecurityPolicy terra.StringValue `hcl:"security_policy,attr"`
}

type AuthorizerConfigAttributes struct {
	ref terra.Reference
}

func (ac AuthorizerConfigAttributes) InternalRef() (terra.Reference, error) {
	return ac.ref, nil
}

func (ac AuthorizerConfigAttributes) InternalWithRef(ref terra.Reference) AuthorizerConfigAttributes {
	return AuthorizerConfigAttributes{ref: ref}
}

func (ac AuthorizerConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ac.ref.InternalTokens()
}

func (ac AuthorizerConfigAttributes) AllowAuthorizerOverride() terra.BoolValue {
	return terra.ReferenceAsBool(ac.ref.Append("allow_authorizer_override"))
}

func (ac AuthorizerConfigAttributes) DefaultAuthorizerName() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("default_authorizer_name"))
}

type TlsConfigAttributes struct {
	ref terra.Reference
}

func (tc TlsConfigAttributes) InternalRef() (terra.Reference, error) {
	return tc.ref, nil
}

func (tc TlsConfigAttributes) InternalWithRef(ref terra.Reference) TlsConfigAttributes {
	return TlsConfigAttributes{ref: ref}
}

func (tc TlsConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tc.ref.InternalTokens()
}

func (tc TlsConfigAttributes) SecurityPolicy() terra.StringValue {
	return terra.ReferenceAsString(tc.ref.Append("security_policy"))
}

type AuthorizerConfigState struct {
	AllowAuthorizerOverride bool   `json:"allow_authorizer_override"`
	DefaultAuthorizerName   string `json:"default_authorizer_name"`
}

type TlsConfigState struct {
	SecurityPolicy string `json:"security_policy"`
}
