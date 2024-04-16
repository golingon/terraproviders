// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_dax_parameter_group

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Parameters struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type ParametersAttributes struct {
	ref terra.Reference
}

func (p ParametersAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p ParametersAttributes) InternalWithRef(ref terra.Reference) ParametersAttributes {
	return ParametersAttributes{ref: ref}
}

func (p ParametersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p ParametersAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("name"))
}

func (p ParametersAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("value"))
}

type ParametersState struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
