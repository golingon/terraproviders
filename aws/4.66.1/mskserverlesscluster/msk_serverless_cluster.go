// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package mskserverlesscluster

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ClientAuthentication struct {
	// Sasl: required
	Sasl *Sasl `hcl:"sasl,block" validate:"required"`
}

type Sasl struct {
	// Iam: required
	Iam *Iam `hcl:"iam,block" validate:"required"`
}

type Iam struct {
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

func (ca ClientAuthenticationAttributes) Sasl() terra.ListValue[SaslAttributes] {
	return terra.ReferenceAsList[SaslAttributes](ca.ref.Append("sasl"))
}

type SaslAttributes struct {
	ref terra.Reference
}

func (s SaslAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SaslAttributes) InternalWithRef(ref terra.Reference) SaslAttributes {
	return SaslAttributes{ref: ref}
}

func (s SaslAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SaslAttributes) Iam() terra.ListValue[IamAttributes] {
	return terra.ReferenceAsList[IamAttributes](s.ref.Append("iam"))
}

type IamAttributes struct {
	ref terra.Reference
}

func (i IamAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i IamAttributes) InternalWithRef(ref terra.Reference) IamAttributes {
	return IamAttributes{ref: ref}
}

func (i IamAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i IamAttributes) Enabled() terra.BoolValue {
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
	Sasl []SaslState `json:"sasl"`
}

type SaslState struct {
	Iam []IamState `json:"iam"`
}

type IamState struct {
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
