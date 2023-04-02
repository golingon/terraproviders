// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package lbsslnegotiationpolicy

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Attribute struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type AttributeAttributes struct {
	ref terra.Reference
}

func (a AttributeAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a AttributeAttributes) InternalWithRef(ref terra.Reference) AttributeAttributes {
	return AttributeAttributes{ref: ref}
}

func (a AttributeAttributes) InternalTokens() hclwrite.Tokens {
	return a.ref.InternalTokens()
}

func (a AttributeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("name"))
}

func (a AttributeAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("value"))
}

type AttributeState struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
