// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package wafgeomatchset

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type GeoMatchConstraint struct {
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type GeoMatchConstraintAttributes struct {
	ref terra.Reference
}

func (gmc GeoMatchConstraintAttributes) InternalRef() terra.Reference {
	return gmc.ref
}

func (gmc GeoMatchConstraintAttributes) InternalWithRef(ref terra.Reference) GeoMatchConstraintAttributes {
	return GeoMatchConstraintAttributes{ref: ref}
}

func (gmc GeoMatchConstraintAttributes) InternalTokens() hclwrite.Tokens {
	return gmc.ref.InternalTokens()
}

func (gmc GeoMatchConstraintAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(gmc.ref.Append("type"))
}

func (gmc GeoMatchConstraintAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(gmc.ref.Append("value"))
}

type GeoMatchConstraintState struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
