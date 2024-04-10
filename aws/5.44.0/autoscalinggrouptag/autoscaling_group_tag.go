// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package autoscalinggrouptag

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Tag struct {
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// PropagateAtLaunch: bool, required
	PropagateAtLaunch terra.BoolValue `hcl:"propagate_at_launch,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type TagAttributes struct {
	ref terra.Reference
}

func (t TagAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TagAttributes) InternalWithRef(ref terra.Reference) TagAttributes {
	return TagAttributes{ref: ref}
}

func (t TagAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TagAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("key"))
}

func (t TagAttributes) PropagateAtLaunch() terra.BoolValue {
	return terra.ReferenceAsBool(t.ref.Append("propagate_at_launch"))
}

func (t TagAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("value"))
}

type TagState struct {
	Key               string `json:"key"`
	PropagateAtLaunch bool   `json:"propagate_at_launch"`
	Value             string `json:"value"`
}
