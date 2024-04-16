// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_apigatewayv2_integration

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type ResponseParameters struct {
	// Mappings: map of string, required
	Mappings terra.MapValue[terra.StringValue] `hcl:"mappings,attr" validate:"required"`
	// StatusCode: string, required
	StatusCode terra.StringValue `hcl:"status_code,attr" validate:"required"`
}

type TlsConfig struct {
	// ServerNameToVerify: string, optional
	ServerNameToVerify terra.StringValue `hcl:"server_name_to_verify,attr"`
}

type ResponseParametersAttributes struct {
	ref terra.Reference
}

func (rp ResponseParametersAttributes) InternalRef() (terra.Reference, error) {
	return rp.ref, nil
}

func (rp ResponseParametersAttributes) InternalWithRef(ref terra.Reference) ResponseParametersAttributes {
	return ResponseParametersAttributes{ref: ref}
}

func (rp ResponseParametersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rp.ref.InternalTokens()
}

func (rp ResponseParametersAttributes) Mappings() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rp.ref.Append("mappings"))
}

func (rp ResponseParametersAttributes) StatusCode() terra.StringValue {
	return terra.ReferenceAsString(rp.ref.Append("status_code"))
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

func (tc TlsConfigAttributes) ServerNameToVerify() terra.StringValue {
	return terra.ReferenceAsString(tc.ref.Append("server_name_to_verify"))
}

type ResponseParametersState struct {
	Mappings   map[string]string `json:"mappings"`
	StatusCode string            `json:"status_code"`
}

type TlsConfigState struct {
	ServerNameToVerify string `json:"server_name_to_verify"`
}
