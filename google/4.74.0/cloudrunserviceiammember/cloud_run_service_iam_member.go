// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package cloudrunserviceiammember

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Condition struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Expression: string, required
	Expression terra.StringValue `hcl:"expression,attr" validate:"required"`
	// Title: string, required
	Title terra.StringValue `hcl:"title,attr" validate:"required"`
}

type ConditionAttributes struct {
	ref terra.Reference
}

func (c ConditionAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ConditionAttributes) InternalWithRef(ref terra.Reference) ConditionAttributes {
	return ConditionAttributes{ref: ref}
}

func (c ConditionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ConditionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("description"))
}

func (c ConditionAttributes) Expression() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("expression"))
}

func (c ConditionAttributes) Title() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("title"))
}

type ConditionState struct {
	Description string `json:"description"`
	Expression  string `json:"expression"`
	Title       string `json:"title"`
}
