// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datacomputeglobalforwardingrule

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type MetadataFilters struct {
	// FilterLabels: min=0
	FilterLabels []FilterLabels `hcl:"filter_labels,block" validate:"min=0"`
}

type FilterLabels struct{}

type MetadataFiltersAttributes struct {
	ref terra.Reference
}

func (mf MetadataFiltersAttributes) InternalRef() (terra.Reference, error) {
	return mf.ref, nil
}

func (mf MetadataFiltersAttributes) InternalWithRef(ref terra.Reference) MetadataFiltersAttributes {
	return MetadataFiltersAttributes{ref: ref}
}

func (mf MetadataFiltersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mf.ref.InternalTokens()
}

func (mf MetadataFiltersAttributes) FilterMatchCriteria() terra.StringValue {
	return terra.ReferenceAsString(mf.ref.Append("filter_match_criteria"))
}

func (mf MetadataFiltersAttributes) FilterLabels() terra.ListValue[FilterLabelsAttributes] {
	return terra.ReferenceAsList[FilterLabelsAttributes](mf.ref.Append("filter_labels"))
}

type FilterLabelsAttributes struct {
	ref terra.Reference
}

func (fl FilterLabelsAttributes) InternalRef() (terra.Reference, error) {
	return fl.ref, nil
}

func (fl FilterLabelsAttributes) InternalWithRef(ref terra.Reference) FilterLabelsAttributes {
	return FilterLabelsAttributes{ref: ref}
}

func (fl FilterLabelsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fl.ref.InternalTokens()
}

func (fl FilterLabelsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(fl.ref.Append("name"))
}

func (fl FilterLabelsAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(fl.ref.Append("value"))
}

type MetadataFiltersState struct {
	FilterMatchCriteria string              `json:"filter_match_criteria"`
	FilterLabels        []FilterLabelsState `json:"filter_labels"`
}

type FilterLabelsState struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}