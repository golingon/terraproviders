// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package ssoadmininstanceaccesscontrolattributes

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Attribute struct {
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// Value: min=1
	Value []Value `hcl:"value,block" validate:"min=1"`
}

type Value struct {
	// Source: set of string, required
	Source terra.SetValue[terra.StringValue] `hcl:"source,attr" validate:"required"`
}

type AttributeAttributes struct {
	ref terra.Reference
}

func (a AttributeAttributes) InternalRef() terra.Reference {
	return a.ref
}

func (a AttributeAttributes) InternalWithRef(ref terra.Reference) AttributeAttributes {
	return AttributeAttributes{ref: ref}
}

func (a AttributeAttributes) InternalTokens() hclwrite.Tokens {
	return a.ref.InternalTokens()
}

func (a AttributeAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("key"))
}

func (a AttributeAttributes) Value() terra.SetValue[ValueAttributes] {
	return terra.ReferenceAsSet[ValueAttributes](a.ref.Append("value"))
}

type ValueAttributes struct {
	ref terra.Reference
}

func (v ValueAttributes) InternalRef() terra.Reference {
	return v.ref
}

func (v ValueAttributes) InternalWithRef(ref terra.Reference) ValueAttributes {
	return ValueAttributes{ref: ref}
}

func (v ValueAttributes) InternalTokens() hclwrite.Tokens {
	return v.ref.InternalTokens()
}

func (v ValueAttributes) Source() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](v.ref.Append("source"))
}

type AttributeState struct {
	Key   string       `json:"key"`
	Value []ValueState `json:"value"`
}

type ValueState struct {
	Source []string `json:"source"`
}
