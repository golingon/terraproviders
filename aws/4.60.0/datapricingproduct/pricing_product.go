// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datapricingproduct

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Filters struct {
	// Field: string, required
	Field terra.StringValue `hcl:"field,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type FiltersAttributes struct {
	ref terra.Reference
}

func (f FiltersAttributes) InternalRef() terra.Reference {
	return f.ref
}

func (f FiltersAttributes) InternalWithRef(ref terra.Reference) FiltersAttributes {
	return FiltersAttributes{ref: ref}
}

func (f FiltersAttributes) InternalTokens() hclwrite.Tokens {
	return f.ref.InternalTokens()
}

func (f FiltersAttributes) Field() terra.StringValue {
	return terra.ReferenceString(f.ref.Append("field"))
}

func (f FiltersAttributes) Value() terra.StringValue {
	return terra.ReferenceString(f.ref.Append("value"))
}

type FiltersState struct {
	Field string `json:"field"`
	Value string `json:"value"`
}
