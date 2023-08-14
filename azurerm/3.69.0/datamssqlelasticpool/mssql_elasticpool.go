// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datamssqlelasticpool

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Sku struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type SkuAttributes struct {
	ref terra.Reference
}

func (s SkuAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SkuAttributes) InternalWithRef(ref terra.Reference) SkuAttributes {
	return SkuAttributes{ref: ref}
}

func (s SkuAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SkuAttributes) Capacity() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("capacity"))
}

func (s SkuAttributes) Family() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("family"))
}

func (s SkuAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("name"))
}

func (s SkuAttributes) Tier() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("tier"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type SkuState struct {
	Capacity float64 `json:"capacity"`
	Family   string  `json:"family"`
	Name     string  `json:"name"`
	Tier     string  `json:"tier"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
