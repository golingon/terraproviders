// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package containeranalysisnote

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AttestationAuthority struct {
	// Hint: required
	Hint *Hint `hcl:"hint,block" validate:"required"`
}

type Hint struct {
	// HumanReadableName: string, required
	HumanReadableName terra.StringValue `hcl:"human_readable_name,attr" validate:"required"`
}

type RelatedUrl struct {
	// Label: string, optional
	Label terra.StringValue `hcl:"label,attr"`
	// Url: string, required
	Url terra.StringValue `hcl:"url,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type AttestationAuthorityAttributes struct {
	ref terra.Reference
}

func (aa AttestationAuthorityAttributes) InternalRef() terra.Reference {
	return aa.ref
}

func (aa AttestationAuthorityAttributes) InternalWithRef(ref terra.Reference) AttestationAuthorityAttributes {
	return AttestationAuthorityAttributes{ref: ref}
}

func (aa AttestationAuthorityAttributes) InternalTokens() hclwrite.Tokens {
	return aa.ref.InternalTokens()
}

func (aa AttestationAuthorityAttributes) Hint() terra.ListValue[HintAttributes] {
	return terra.ReferenceList[HintAttributes](aa.ref.Append("hint"))
}

type HintAttributes struct {
	ref terra.Reference
}

func (h HintAttributes) InternalRef() terra.Reference {
	return h.ref
}

func (h HintAttributes) InternalWithRef(ref terra.Reference) HintAttributes {
	return HintAttributes{ref: ref}
}

func (h HintAttributes) InternalTokens() hclwrite.Tokens {
	return h.ref.InternalTokens()
}

func (h HintAttributes) HumanReadableName() terra.StringValue {
	return terra.ReferenceString(h.ref.Append("human_readable_name"))
}

type RelatedUrlAttributes struct {
	ref terra.Reference
}

func (ru RelatedUrlAttributes) InternalRef() terra.Reference {
	return ru.ref
}

func (ru RelatedUrlAttributes) InternalWithRef(ref terra.Reference) RelatedUrlAttributes {
	return RelatedUrlAttributes{ref: ref}
}

func (ru RelatedUrlAttributes) InternalTokens() hclwrite.Tokens {
	return ru.ref.InternalTokens()
}

func (ru RelatedUrlAttributes) Label() terra.StringValue {
	return terra.ReferenceString(ru.ref.Append("label"))
}

func (ru RelatedUrlAttributes) Url() terra.StringValue {
	return terra.ReferenceString(ru.ref.Append("url"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("update"))
}

type AttestationAuthorityState struct {
	Hint []HintState `json:"hint"`
}

type HintState struct {
	HumanReadableName string `json:"human_readable_name"`
}

type RelatedUrlState struct {
	Label string `json:"label"`
	Url   string `json:"url"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
