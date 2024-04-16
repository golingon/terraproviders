// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_cloudwatch_event_permission

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Condition struct {
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
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

func (c ConditionAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("key"))
}

func (c ConditionAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("type"))
}

func (c ConditionAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("value"))
}

type ConditionState struct {
	Key   string `json:"key"`
	Type  string `json:"type"`
	Value string `json:"value"`
}
