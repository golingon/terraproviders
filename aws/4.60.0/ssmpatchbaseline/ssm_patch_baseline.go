// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package ssmpatchbaseline

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ApprovalRule struct {
	// ApproveAfterDays: number, optional
	ApproveAfterDays terra.NumberValue `hcl:"approve_after_days,attr"`
	// ApproveUntilDate: string, optional
	ApproveUntilDate terra.StringValue `hcl:"approve_until_date,attr"`
	// ComplianceLevel: string, optional
	ComplianceLevel terra.StringValue `hcl:"compliance_level,attr"`
	// EnableNonSecurity: bool, optional
	EnableNonSecurity terra.BoolValue `hcl:"enable_non_security,attr"`
	// PatchFilter: min=1,max=10
	PatchFilter []PatchFilter `hcl:"patch_filter,block" validate:"min=1,max=10"`
}

type PatchFilter struct {
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// Values: list of string, required
	Values terra.ListValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type GlobalFilter struct {
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// Values: list of string, required
	Values terra.ListValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type Source struct {
	// Configuration: string, required
	Configuration terra.StringValue `hcl:"configuration,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Products: list of string, required
	Products terra.ListValue[terra.StringValue] `hcl:"products,attr" validate:"required"`
}

type ApprovalRuleAttributes struct {
	ref terra.Reference
}

func (ar ApprovalRuleAttributes) InternalRef() terra.Reference {
	return ar.ref
}

func (ar ApprovalRuleAttributes) InternalWithRef(ref terra.Reference) ApprovalRuleAttributes {
	return ApprovalRuleAttributes{ref: ref}
}

func (ar ApprovalRuleAttributes) InternalTokens() hclwrite.Tokens {
	return ar.ref.InternalTokens()
}

func (ar ApprovalRuleAttributes) ApproveAfterDays() terra.NumberValue {
	return terra.ReferenceAsNumber(ar.ref.Append("approve_after_days"))
}

func (ar ApprovalRuleAttributes) ApproveUntilDate() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("approve_until_date"))
}

func (ar ApprovalRuleAttributes) ComplianceLevel() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("compliance_level"))
}

func (ar ApprovalRuleAttributes) EnableNonSecurity() terra.BoolValue {
	return terra.ReferenceAsBool(ar.ref.Append("enable_non_security"))
}

func (ar ApprovalRuleAttributes) PatchFilter() terra.ListValue[PatchFilterAttributes] {
	return terra.ReferenceAsList[PatchFilterAttributes](ar.ref.Append("patch_filter"))
}

type PatchFilterAttributes struct {
	ref terra.Reference
}

func (pf PatchFilterAttributes) InternalRef() terra.Reference {
	return pf.ref
}

func (pf PatchFilterAttributes) InternalWithRef(ref terra.Reference) PatchFilterAttributes {
	return PatchFilterAttributes{ref: ref}
}

func (pf PatchFilterAttributes) InternalTokens() hclwrite.Tokens {
	return pf.ref.InternalTokens()
}

func (pf PatchFilterAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(pf.ref.Append("key"))
}

func (pf PatchFilterAttributes) Values() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pf.ref.Append("values"))
}

type GlobalFilterAttributes struct {
	ref terra.Reference
}

func (gf GlobalFilterAttributes) InternalRef() terra.Reference {
	return gf.ref
}

func (gf GlobalFilterAttributes) InternalWithRef(ref terra.Reference) GlobalFilterAttributes {
	return GlobalFilterAttributes{ref: ref}
}

func (gf GlobalFilterAttributes) InternalTokens() hclwrite.Tokens {
	return gf.ref.InternalTokens()
}

func (gf GlobalFilterAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(gf.ref.Append("key"))
}

func (gf GlobalFilterAttributes) Values() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](gf.ref.Append("values"))
}

type SourceAttributes struct {
	ref terra.Reference
}

func (s SourceAttributes) InternalRef() terra.Reference {
	return s.ref
}

func (s SourceAttributes) InternalWithRef(ref terra.Reference) SourceAttributes {
	return SourceAttributes{ref: ref}
}

func (s SourceAttributes) InternalTokens() hclwrite.Tokens {
	return s.ref.InternalTokens()
}

func (s SourceAttributes) Configuration() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("configuration"))
}

func (s SourceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("name"))
}

func (s SourceAttributes) Products() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](s.ref.Append("products"))
}

type ApprovalRuleState struct {
	ApproveAfterDays  float64            `json:"approve_after_days"`
	ApproveUntilDate  string             `json:"approve_until_date"`
	ComplianceLevel   string             `json:"compliance_level"`
	EnableNonSecurity bool               `json:"enable_non_security"`
	PatchFilter       []PatchFilterState `json:"patch_filter"`
}

type PatchFilterState struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

type GlobalFilterState struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

type SourceState struct {
	Configuration string   `json:"configuration"`
	Name          string   `json:"name"`
	Products      []string `json:"products"`
}