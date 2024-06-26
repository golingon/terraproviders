// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_msk_serverless_cluster

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type ClientAuthentication struct {
	// ClientAuthenticationSasl: required
	Sasl *ClientAuthenticationSasl `hcl:"sasl,block" validate:"required"`
}

type ClientAuthenticationSasl struct {
	// ClientAuthenticationSaslIam: required
	Iam *ClientAuthenticationSaslIam `hcl:"iam,block" validate:"required"`
}

type ClientAuthenticationSaslIam struct {
	// Enabled: bool, required
	Enabled terra.BoolValue `hcl:"enabled,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
}

type VpcConfig struct {
	// SecurityGroupIds: set of string, optional
	SecurityGroupIds terra.SetValue[terra.StringValue] `hcl:"security_group_ids,attr"`
	// SubnetIds: set of string, required
	SubnetIds terra.SetValue[terra.StringValue] `hcl:"subnet_ids,attr" validate:"required"`
}

type ClientAuthenticationAttributes struct {
	ref terra.Reference
}

func (ca ClientAuthenticationAttributes) InternalRef() (terra.Reference, error) {
	return ca.ref, nil
}

func (ca ClientAuthenticationAttributes) InternalWithRef(ref terra.Reference) ClientAuthenticationAttributes {
	return ClientAuthenticationAttributes{ref: ref}
}

func (ca ClientAuthenticationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ca.ref.InternalTokens()
}

func (ca ClientAuthenticationAttributes) Sasl() terra.ListValue[ClientAuthenticationSaslAttributes] {
	return terra.ReferenceAsList[ClientAuthenticationSaslAttributes](ca.ref.Append("sasl"))
}

type ClientAuthenticationSaslAttributes struct {
	ref terra.Reference
}

func (s ClientAuthenticationSaslAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s ClientAuthenticationSaslAttributes) InternalWithRef(ref terra.Reference) ClientAuthenticationSaslAttributes {
	return ClientAuthenticationSaslAttributes{ref: ref}
}

func (s ClientAuthenticationSaslAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s ClientAuthenticationSaslAttributes) Iam() terra.ListValue[ClientAuthenticationSaslIamAttributes] {
	return terra.ReferenceAsList[ClientAuthenticationSaslIamAttributes](s.ref.Append("iam"))
}

type ClientAuthenticationSaslIamAttributes struct {
	ref terra.Reference
}

func (i ClientAuthenticationSaslIamAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i ClientAuthenticationSaslIamAttributes) InternalWithRef(ref terra.Reference) ClientAuthenticationSaslIamAttributes {
	return ClientAuthenticationSaslIamAttributes{ref: ref}
}

func (i ClientAuthenticationSaslIamAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i ClientAuthenticationSaslIamAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(i.ref.Append("enabled"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

type VpcConfigAttributes struct {
	ref terra.Reference
}

func (vc VpcConfigAttributes) InternalRef() (terra.Reference, error) {
	return vc.ref, nil
}

func (vc VpcConfigAttributes) InternalWithRef(ref terra.Reference) VpcConfigAttributes {
	return VpcConfigAttributes{ref: ref}
}

func (vc VpcConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return vc.ref.InternalTokens()
}

func (vc VpcConfigAttributes) SecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vc.ref.Append("security_group_ids"))
}

func (vc VpcConfigAttributes) SubnetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vc.ref.Append("subnet_ids"))
}

type ClientAuthenticationState struct {
	Sasl []ClientAuthenticationSaslState `json:"sasl"`
}

type ClientAuthenticationSaslState struct {
	Iam []ClientAuthenticationSaslIamState `json:"iam"`
}

type ClientAuthenticationSaslIamState struct {
	Enabled bool `json:"enabled"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}

type VpcConfigState struct {
	SecurityGroupIds []string `json:"security_group_ids"`
	SubnetIds        []string `json:"subnet_ids"`
}
