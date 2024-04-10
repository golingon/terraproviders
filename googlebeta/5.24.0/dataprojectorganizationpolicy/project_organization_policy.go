// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package dataprojectorganizationpolicy

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type BooleanPolicy struct{}

type ListPolicy struct {
	// Allow: min=0
	Allow []Allow `hcl:"allow,block" validate:"min=0"`
	// Deny: min=0
	Deny []Deny `hcl:"deny,block" validate:"min=0"`
}

type Allow struct{}

type Deny struct{}

type RestorePolicy struct{}

type BooleanPolicyAttributes struct {
	ref terra.Reference
}

func (bp BooleanPolicyAttributes) InternalRef() (terra.Reference, error) {
	return bp.ref, nil
}

func (bp BooleanPolicyAttributes) InternalWithRef(ref terra.Reference) BooleanPolicyAttributes {
	return BooleanPolicyAttributes{ref: ref}
}

func (bp BooleanPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bp.ref.InternalTokens()
}

func (bp BooleanPolicyAttributes) Enforced() terra.BoolValue {
	return terra.ReferenceAsBool(bp.ref.Append("enforced"))
}

type ListPolicyAttributes struct {
	ref terra.Reference
}

func (lp ListPolicyAttributes) InternalRef() (terra.Reference, error) {
	return lp.ref, nil
}

func (lp ListPolicyAttributes) InternalWithRef(ref terra.Reference) ListPolicyAttributes {
	return ListPolicyAttributes{ref: ref}
}

func (lp ListPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lp.ref.InternalTokens()
}

func (lp ListPolicyAttributes) InheritFromParent() terra.BoolValue {
	return terra.ReferenceAsBool(lp.ref.Append("inherit_from_parent"))
}

func (lp ListPolicyAttributes) SuggestedValue() terra.StringValue {
	return terra.ReferenceAsString(lp.ref.Append("suggested_value"))
}

func (lp ListPolicyAttributes) Allow() terra.ListValue[AllowAttributes] {
	return terra.ReferenceAsList[AllowAttributes](lp.ref.Append("allow"))
}

func (lp ListPolicyAttributes) Deny() terra.ListValue[DenyAttributes] {
	return terra.ReferenceAsList[DenyAttributes](lp.ref.Append("deny"))
}

type AllowAttributes struct {
	ref terra.Reference
}

func (a AllowAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a AllowAttributes) InternalWithRef(ref terra.Reference) AllowAttributes {
	return AllowAttributes{ref: ref}
}

func (a AllowAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a AllowAttributes) All() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("all"))
}

func (a AllowAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](a.ref.Append("values"))
}

type DenyAttributes struct {
	ref terra.Reference
}

func (d DenyAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d DenyAttributes) InternalWithRef(ref terra.Reference) DenyAttributes {
	return DenyAttributes{ref: ref}
}

func (d DenyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d DenyAttributes) All() terra.BoolValue {
	return terra.ReferenceAsBool(d.ref.Append("all"))
}

func (d DenyAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](d.ref.Append("values"))
}

type RestorePolicyAttributes struct {
	ref terra.Reference
}

func (rp RestorePolicyAttributes) InternalRef() (terra.Reference, error) {
	return rp.ref, nil
}

func (rp RestorePolicyAttributes) InternalWithRef(ref terra.Reference) RestorePolicyAttributes {
	return RestorePolicyAttributes{ref: ref}
}

func (rp RestorePolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rp.ref.InternalTokens()
}

func (rp RestorePolicyAttributes) Default() terra.BoolValue {
	return terra.ReferenceAsBool(rp.ref.Append("default"))
}

type BooleanPolicyState struct {
	Enforced bool `json:"enforced"`
}

type ListPolicyState struct {
	InheritFromParent bool         `json:"inherit_from_parent"`
	SuggestedValue    string       `json:"suggested_value"`
	Allow             []AllowState `json:"allow"`
	Deny              []DenyState  `json:"deny"`
}

type AllowState struct {
	All    bool     `json:"all"`
	Values []string `json:"values"`
}

type DenyState struct {
	All    bool     `json:"all"`
	Values []string `json:"values"`
}

type RestorePolicyState struct {
	Default bool `json:"default"`
}
