// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package ecrregistryscanningconfiguration

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Rule struct {
	// ScanFrequency: string, required
	ScanFrequency terra.StringValue `hcl:"scan_frequency,attr" validate:"required"`
	// RepositoryFilter: min=1
	RepositoryFilter []RepositoryFilter `hcl:"repository_filter,block" validate:"min=1"`
}

type RepositoryFilter struct {
	// Filter: string, required
	Filter terra.StringValue `hcl:"filter,attr" validate:"required"`
	// FilterType: string, required
	FilterType terra.StringValue `hcl:"filter_type,attr" validate:"required"`
}

type RuleAttributes struct {
	ref terra.Reference
}

func (r RuleAttributes) InternalRef() terra.Reference {
	return r.ref
}

func (r RuleAttributes) InternalWithRef(ref terra.Reference) RuleAttributes {
	return RuleAttributes{ref: ref}
}

func (r RuleAttributes) InternalTokens() hclwrite.Tokens {
	return r.ref.InternalTokens()
}

func (r RuleAttributes) ScanFrequency() terra.StringValue {
	return terra.ReferenceString(r.ref.Append("scan_frequency"))
}

func (r RuleAttributes) RepositoryFilter() terra.SetValue[RepositoryFilterAttributes] {
	return terra.ReferenceSet[RepositoryFilterAttributes](r.ref.Append("repository_filter"))
}

type RepositoryFilterAttributes struct {
	ref terra.Reference
}

func (rf RepositoryFilterAttributes) InternalRef() terra.Reference {
	return rf.ref
}

func (rf RepositoryFilterAttributes) InternalWithRef(ref terra.Reference) RepositoryFilterAttributes {
	return RepositoryFilterAttributes{ref: ref}
}

func (rf RepositoryFilterAttributes) InternalTokens() hclwrite.Tokens {
	return rf.ref.InternalTokens()
}

func (rf RepositoryFilterAttributes) Filter() terra.StringValue {
	return terra.ReferenceString(rf.ref.Append("filter"))
}

func (rf RepositoryFilterAttributes) FilterType() terra.StringValue {
	return terra.ReferenceString(rf.ref.Append("filter_type"))
}

type RuleState struct {
	ScanFrequency    string                  `json:"scan_frequency"`
	RepositoryFilter []RepositoryFilterState `json:"repository_filter"`
}

type RepositoryFilterState struct {
	Filter     string `json:"filter"`
	FilterType string `json:"filter_type"`
}
