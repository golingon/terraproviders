// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datassmmaintenancewindows

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Filter struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Values: list of string, required
	Values terra.ListValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
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

func (f FilterAttributes) Name() terra.StringValue {
	return terra.ReferenceString(f.ref.Append("name"))
}

func (f FilterAttributes) Values() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](f.ref.Append("values"))
}

type FilterState struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}
