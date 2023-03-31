// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package s3bucketintelligenttieringconfiguration

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Filter struct {
	// Prefix: string, optional
	Prefix terra.StringValue `hcl:"prefix,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}

type Tiering struct {
	// AccessTier: string, required
	AccessTier terra.StringValue `hcl:"access_tier,attr" validate:"required"`
	// Days: number, required
	Days terra.NumberValue `hcl:"days,attr" validate:"required"`
}

type FilterAttributes struct {
	ref terra.Reference
}

func (f FilterAttributes) InternalRef() terra.Reference {
	return f.ref
}

func (f FilterAttributes) InternalWithRef(ref terra.Reference) FilterAttributes {
	return FilterAttributes{ref: ref}
}

func (f FilterAttributes) InternalTokens() hclwrite.Tokens {
	return f.ref.InternalTokens()
}

func (f FilterAttributes) Prefix() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("prefix"))
}

func (f FilterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](f.ref.Append("tags"))
}

type TieringAttributes struct {
	ref terra.Reference
}

func (t TieringAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TieringAttributes) InternalWithRef(ref terra.Reference) TieringAttributes {
	return TieringAttributes{ref: ref}
}

func (t TieringAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TieringAttributes) AccessTier() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("access_tier"))
}

func (t TieringAttributes) Days() terra.NumberValue {
	return terra.ReferenceAsNumber(t.ref.Append("days"))
}

type FilterState struct {
	Prefix string            `json:"prefix"`
	Tags   map[string]string `json:"tags"`
}

type TieringState struct {
	AccessTier string  `json:"access_tier"`
	Days       float64 `json:"days"`
}
