// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_route53_resolver_endpoint

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataFilter struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Values: list of string, required
	Values terra.ListValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type DataFilterAttributes struct {
	ref terra.Reference
}

func (f DataFilterAttributes) InternalRef() (terra.Reference, error) {
	return f.ref, nil
}

func (f DataFilterAttributes) InternalWithRef(ref terra.Reference) DataFilterAttributes {
	return DataFilterAttributes{ref: ref}
}

func (f DataFilterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return f.ref.InternalTokens()
}

func (f DataFilterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("name"))
}

func (f DataFilterAttributes) Values() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](f.ref.Append("values"))
}

type DataFilterState struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}
