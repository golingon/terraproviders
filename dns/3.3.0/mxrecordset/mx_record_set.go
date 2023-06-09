// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package mxrecordset

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Mx struct {
	// Exchange: string, required
	Exchange terra.StringValue `hcl:"exchange,attr" validate:"required"`
	// Preference: number, required
	Preference terra.NumberValue `hcl:"preference,attr" validate:"required"`
}

type MxAttributes struct {
	ref terra.Reference
}

func (m MxAttributes) InternalRef() (terra.Reference, error) {
	return m.ref, nil
}

func (m MxAttributes) InternalWithRef(ref terra.Reference) MxAttributes {
	return MxAttributes{ref: ref}
}

func (m MxAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return m.ref.InternalTokens()
}

func (m MxAttributes) Exchange() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("exchange"))
}

func (m MxAttributes) Preference() terra.NumberValue {
	return terra.ReferenceAsNumber(m.ref.Append("preference"))
}

type MxState struct {
	Exchange   string  `json:"exchange"`
	Preference float64 `json:"preference"`
}
