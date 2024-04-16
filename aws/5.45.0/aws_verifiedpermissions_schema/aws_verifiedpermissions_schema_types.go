// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_verifiedpermissions_schema

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Definition struct {
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type DefinitionAttributes struct {
	ref terra.Reference
}

func (d DefinitionAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d DefinitionAttributes) InternalWithRef(ref terra.Reference) DefinitionAttributes {
	return DefinitionAttributes{ref: ref}
}

func (d DefinitionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d DefinitionAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("value"))
}

type DefinitionState struct {
	Value string `json:"value"`
}
