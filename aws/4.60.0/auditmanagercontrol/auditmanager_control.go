// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package auditmanagercontrol

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ControlMappingSources struct {
	// SourceDescription: string, optional
	SourceDescription terra.StringValue `hcl:"source_description,attr"`
	// SourceFrequency: string, optional
	SourceFrequency terra.StringValue `hcl:"source_frequency,attr"`
	// SourceName: string, required
	SourceName terra.StringValue `hcl:"source_name,attr" validate:"required"`
	// SourceSetUpOption: string, required
	SourceSetUpOption terra.StringValue `hcl:"source_set_up_option,attr" validate:"required"`
	// SourceType: string, required
	SourceType terra.StringValue `hcl:"source_type,attr" validate:"required"`
	// TroubleshootingText: string, optional
	TroubleshootingText terra.StringValue `hcl:"troubleshooting_text,attr"`
	// SourceKeyword: min=0
	SourceKeyword []SourceKeyword `hcl:"source_keyword,block" validate:"min=0"`
}

type SourceKeyword struct {
	// KeywordInputType: string, required
	KeywordInputType terra.StringValue `hcl:"keyword_input_type,attr" validate:"required"`
	// KeywordValue: string, required
	KeywordValue terra.StringValue `hcl:"keyword_value,attr" validate:"required"`
}

type ControlMappingSourcesAttributes struct {
	ref terra.Reference
}

func (cms ControlMappingSourcesAttributes) InternalRef() terra.Reference {
	return cms.ref
}

func (cms ControlMappingSourcesAttributes) InternalWithRef(ref terra.Reference) ControlMappingSourcesAttributes {
	return ControlMappingSourcesAttributes{ref: ref}
}

func (cms ControlMappingSourcesAttributes) InternalTokens() hclwrite.Tokens {
	return cms.ref.InternalTokens()
}

func (cms ControlMappingSourcesAttributes) SourceDescription() terra.StringValue {
	return terra.ReferenceString(cms.ref.Append("source_description"))
}

func (cms ControlMappingSourcesAttributes) SourceFrequency() terra.StringValue {
	return terra.ReferenceString(cms.ref.Append("source_frequency"))
}

func (cms ControlMappingSourcesAttributes) SourceId() terra.StringValue {
	return terra.ReferenceString(cms.ref.Append("source_id"))
}

func (cms ControlMappingSourcesAttributes) SourceName() terra.StringValue {
	return terra.ReferenceString(cms.ref.Append("source_name"))
}

func (cms ControlMappingSourcesAttributes) SourceSetUpOption() terra.StringValue {
	return terra.ReferenceString(cms.ref.Append("source_set_up_option"))
}

func (cms ControlMappingSourcesAttributes) SourceType() terra.StringValue {
	return terra.ReferenceString(cms.ref.Append("source_type"))
}

func (cms ControlMappingSourcesAttributes) TroubleshootingText() terra.StringValue {
	return terra.ReferenceString(cms.ref.Append("troubleshooting_text"))
}

func (cms ControlMappingSourcesAttributes) SourceKeyword() terra.ListValue[SourceKeywordAttributes] {
	return terra.ReferenceList[SourceKeywordAttributes](cms.ref.Append("source_keyword"))
}

type SourceKeywordAttributes struct {
	ref terra.Reference
}

func (sk SourceKeywordAttributes) InternalRef() terra.Reference {
	return sk.ref
}

func (sk SourceKeywordAttributes) InternalWithRef(ref terra.Reference) SourceKeywordAttributes {
	return SourceKeywordAttributes{ref: ref}
}

func (sk SourceKeywordAttributes) InternalTokens() hclwrite.Tokens {
	return sk.ref.InternalTokens()
}

func (sk SourceKeywordAttributes) KeywordInputType() terra.StringValue {
	return terra.ReferenceString(sk.ref.Append("keyword_input_type"))
}

func (sk SourceKeywordAttributes) KeywordValue() terra.StringValue {
	return terra.ReferenceString(sk.ref.Append("keyword_value"))
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
