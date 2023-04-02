// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package sagemakerhumantaskui

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type UiTemplate struct {
	// Content: string, optional
	Content terra.StringValue `hcl:"content,attr"`
}

type UiTemplateAttributes struct {
	ref terra.Reference
}

func (ut UiTemplateAttributes) InternalRef() (terra.Reference, error) {
	return ut.ref, nil
}

func (ut UiTemplateAttributes) InternalWithRef(ref terra.Reference) UiTemplateAttributes {
	return UiTemplateAttributes{ref: ref}
}

func (ut UiTemplateAttributes) InternalTokens() hclwrite.Tokens {
	return ut.ref.InternalTokens()
}

func (ut UiTemplateAttributes) Content() terra.StringValue {
	return terra.ReferenceAsString(ut.ref.Append("content"))
}

func (ut UiTemplateAttributes) ContentSha256() terra.StringValue {
	return terra.ReferenceAsString(ut.ref.Append("content_sha256"))
}

func (ut UiTemplateAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(ut.ref.Append("url"))
}

type UiTemplateState struct {
	Content       string `json:"content"`
	ContentSha256 string `json:"content_sha256"`
	Url           string `json:"url"`
}
