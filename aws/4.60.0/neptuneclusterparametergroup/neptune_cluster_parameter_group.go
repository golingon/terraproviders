// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package neptuneclusterparametergroup

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Parameter struct {
	// ApplyMethod: string, optional
	ApplyMethod terra.StringValue `hcl:"apply_method,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type ParameterAttributes struct {
	ref terra.Reference
}

func (p ParameterAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p ParameterAttributes) InternalWithRef(ref terra.Reference) ParameterAttributes {
	return ParameterAttributes{ref: ref}
}

func (p ParameterAttributes) InternalTokens() hclwrite.Tokens {
	return p.ref.InternalTokens()
}

func (p ParameterAttributes) ApplyMethod() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("apply_method"))
}

func (p ParameterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("name"))
}

func (p ParameterAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("value"))
}

type ParameterState struct {
	ApplyMethod string `json:"apply_method"`
	Name        string `json:"name"`
	Value       string `json:"value"`
}
