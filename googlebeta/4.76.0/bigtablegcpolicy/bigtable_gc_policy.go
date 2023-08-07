// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package bigtablegcpolicy

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type MaxAge struct {
	// Days: number, optional
	Days terra.NumberValue `hcl:"days,attr"`
	// Duration: string, optional
	Duration terra.StringValue `hcl:"duration,attr"`
}

type MaxVersion struct {
	// Number: number, required
	Number terra.NumberValue `hcl:"number,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
}

type MaxAgeAttributes struct {
	ref terra.Reference
}

func (ma MaxAgeAttributes) InternalRef() (terra.Reference, error) {
	return ma.ref, nil
}

func (ma MaxAgeAttributes) InternalWithRef(ref terra.Reference) MaxAgeAttributes {
	return MaxAgeAttributes{ref: ref}
}

func (ma MaxAgeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ma.ref.InternalTokens()
}

func (ma MaxAgeAttributes) Days() terra.NumberValue {
	return terra.ReferenceAsNumber(ma.ref.Append("days"))
}

func (ma MaxAgeAttributes) Duration() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("duration"))
}

type MaxVersionAttributes struct {
	ref terra.Reference
}

func (mv MaxVersionAttributes) InternalRef() (terra.Reference, error) {
	return mv.ref, nil
}

func (mv MaxVersionAttributes) InternalWithRef(ref terra.Reference) MaxVersionAttributes {
	return MaxVersionAttributes{ref: ref}
}

func (mv MaxVersionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mv.ref.InternalTokens()
}

func (mv MaxVersionAttributes) Number() terra.NumberValue {
	return terra.ReferenceAsNumber(mv.ref.Append("number"))
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

type MaxAgeState struct {
	Days     float64 `json:"days"`
	Duration string  `json:"duration"`
}

type MaxVersionState struct {
	Number float64 `json:"number"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}
