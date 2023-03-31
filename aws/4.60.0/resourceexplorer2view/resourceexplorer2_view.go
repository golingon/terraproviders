// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package resourceexplorer2view

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Filters struct {
	// FilterString: string, required
	FilterString terra.StringValue `hcl:"filter_string,attr" validate:"required"`
}

type IncludedProperty struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
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

func (f FiltersAttributes) FilterString() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("filter_string"))
}

type IncludedPropertyAttributes struct {
	ref terra.Reference
}

func (ip IncludedPropertyAttributes) InternalRef() terra.Reference {
	return ip.ref
}

func (ip IncludedPropertyAttributes) InternalWithRef(ref terra.Reference) IncludedPropertyAttributes {
	return IncludedPropertyAttributes{ref: ref}
}

func (ip IncludedPropertyAttributes) InternalTokens() hclwrite.Tokens {
	return ip.ref.InternalTokens()
}

func (ip IncludedPropertyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("name"))
}

type FiltersState struct {
	FilterString string `json:"filter_string"`
}

type IncludedPropertyState struct {
	Name string `json:"name"`
}
