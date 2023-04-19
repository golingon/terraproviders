// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package pipespipe

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type SourceParameters struct {
	// FilterCriteria: optional
	FilterCriteria *FilterCriteria `hcl:"filter_criteria,block"`
}

type FilterCriteria struct {
	// Filter: min=0,max=5
	Filter []Filter `hcl:"filter,block" validate:"min=0,max=5"`
}

type Filter struct {
	// Pattern: string, required
	Pattern terra.StringValue `hcl:"pattern,attr" validate:"required"`
}

type TargetParameters struct {
	// InputTemplate: string, optional
	InputTemplate terra.StringValue `hcl:"input_template,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type SourceParametersAttributes struct {
	ref terra.Reference
}

func (sp SourceParametersAttributes) InternalRef() (terra.Reference, error) {
	return sp.ref, nil
}

func (sp SourceParametersAttributes) InternalWithRef(ref terra.Reference) SourceParametersAttributes {
	return SourceParametersAttributes{ref: ref}
}

func (sp SourceParametersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sp.ref.InternalTokens()
}

func (sp SourceParametersAttributes) FilterCriteria() terra.ListValue[FilterCriteriaAttributes] {
	return terra.ReferenceAsList[FilterCriteriaAttributes](sp.ref.Append("filter_criteria"))
}

type FilterCriteriaAttributes struct {
	ref terra.Reference
}

func (fc FilterCriteriaAttributes) InternalRef() (terra.Reference, error) {
	return fc.ref, nil
}

func (fc FilterCriteriaAttributes) InternalWithRef(ref terra.Reference) FilterCriteriaAttributes {
	return FilterCriteriaAttributes{ref: ref}
}

func (fc FilterCriteriaAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fc.ref.InternalTokens()
}

func (fc FilterCriteriaAttributes) Filter() terra.ListValue[FilterAttributes] {
	return terra.ReferenceAsList[FilterAttributes](fc.ref.Append("filter"))
}

type FilterAttributes struct {
	ref terra.Reference
}

func (f FilterAttributes) InternalRef() (terra.Reference, error) {
	return f.ref, nil
}

func (f FilterAttributes) InternalWithRef(ref terra.Reference) FilterAttributes {
	return FilterAttributes{ref: ref}
}

func (f FilterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return f.ref.InternalTokens()
}

func (f FilterAttributes) Pattern() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("pattern"))
}

type TargetParametersAttributes struct {
	ref terra.Reference
}

func (tp TargetParametersAttributes) InternalRef() (terra.Reference, error) {
	return tp.ref, nil
}

func (tp TargetParametersAttributes) InternalWithRef(ref terra.Reference) TargetParametersAttributes {
	return TargetParametersAttributes{ref: ref}
}

func (tp TargetParametersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tp.ref.InternalTokens()
}

func (tp TargetParametersAttributes) InputTemplate() terra.StringValue {
	return terra.ReferenceAsString(tp.ref.Append("input_template"))
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

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type SourceParametersState struct {
	FilterCriteria []FilterCriteriaState `json:"filter_criteria"`
}

type FilterCriteriaState struct {
	Filter []FilterState `json:"filter"`
}

type FilterState struct {
	Pattern string `json:"pattern"`
}

type TargetParametersState struct {
	InputTemplate string `json:"input_template"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}