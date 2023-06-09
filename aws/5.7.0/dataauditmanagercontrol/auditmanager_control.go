// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataauditmanagercontrol

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ControlMappingSources struct {
	// SourceKeyword: min=0
	SourceKeyword []SourceKeyword `hcl:"source_keyword,block" validate:"min=0"`
}

type SourceKeyword struct{}

type ControlMappingSourcesAttributes struct {
	ref terra.Reference
}

func (cms ControlMappingSourcesAttributes) InternalRef() (terra.Reference, error) {
	return cms.ref, nil
}

func (cms ControlMappingSourcesAttributes) InternalWithRef(ref terra.Reference) ControlMappingSourcesAttributes {
	return ControlMappingSourcesAttributes{ref: ref}
}

func (cms ControlMappingSourcesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cms.ref.InternalTokens()
}

func (cms ControlMappingSourcesAttributes) SourceDescription() terra.StringValue {
	return terra.ReferenceAsString(cms.ref.Append("source_description"))
}

func (cms ControlMappingSourcesAttributes) SourceFrequency() terra.StringValue {
	return terra.ReferenceAsString(cms.ref.Append("source_frequency"))
}

func (cms ControlMappingSourcesAttributes) SourceId() terra.StringValue {
	return terra.ReferenceAsString(cms.ref.Append("source_id"))
}

func (cms ControlMappingSourcesAttributes) SourceName() terra.StringValue {
	return terra.ReferenceAsString(cms.ref.Append("source_name"))
}

func (cms ControlMappingSourcesAttributes) SourceSetUpOption() terra.StringValue {
	return terra.ReferenceAsString(cms.ref.Append("source_set_up_option"))
}

func (cms ControlMappingSourcesAttributes) SourceType() terra.StringValue {
	return terra.ReferenceAsString(cms.ref.Append("source_type"))
}

func (cms ControlMappingSourcesAttributes) TroubleshootingText() terra.StringValue {
	return terra.ReferenceAsString(cms.ref.Append("troubleshooting_text"))
}

func (cms ControlMappingSourcesAttributes) SourceKeyword() terra.ListValue[SourceKeywordAttributes] {
	return terra.ReferenceAsList[SourceKeywordAttributes](cms.ref.Append("source_keyword"))
}

type SourceKeywordAttributes struct {
	ref terra.Reference
}

func (sk SourceKeywordAttributes) InternalRef() (terra.Reference, error) {
	return sk.ref, nil
}

func (sk SourceKeywordAttributes) InternalWithRef(ref terra.Reference) SourceKeywordAttributes {
	return SourceKeywordAttributes{ref: ref}
}

func (sk SourceKeywordAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sk.ref.InternalTokens()
}

func (sk SourceKeywordAttributes) KeywordInputType() terra.StringValue {
	return terra.ReferenceAsString(sk.ref.Append("keyword_input_type"))
}

func (sk SourceKeywordAttributes) KeywordValue() terra.StringValue {
	return terra.ReferenceAsString(sk.ref.Append("keyword_value"))
}

type ControlMappingSourcesState struct {
	SourceDescription   string               `json:"source_description"`
	SourceFrequency     string               `json:"source_frequency"`
	SourceId            string               `json:"source_id"`
	SourceName          string               `json:"source_name"`
	SourceSetUpOption   string               `json:"source_set_up_option"`
	SourceType          string               `json:"source_type"`
	TroubleshootingText string               `json:"troubleshooting_text"`
	SourceKeyword       []SourceKeywordState `json:"source_keyword"`
}

type SourceKeywordState struct {
	KeywordInputType string `json:"keyword_input_type"`
	KeywordValue     string `json:"keyword_value"`
}
