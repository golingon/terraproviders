// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package glaciervault

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Notification struct {
	// Events: set of string, required
	Events terra.SetValue[terra.StringValue] `hcl:"events,attr" validate:"required"`
	// SnsTopic: string, required
	SnsTopic terra.StringValue `hcl:"sns_topic,attr" validate:"required"`
}

type NotificationAttributes struct {
	ref terra.Reference
}

func (n NotificationAttributes) InternalRef() (terra.Reference, error) {
	return n.ref, nil
}

func (n NotificationAttributes) InternalWithRef(ref terra.Reference) NotificationAttributes {
	return NotificationAttributes{ref: ref}
}

func (n NotificationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return n.ref.InternalTokens()
}

func (n NotificationAttributes) Events() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](n.ref.Append("events"))
}

func (n NotificationAttributes) SnsTopic() terra.StringValue {
	return terra.ReferenceAsString(n.ref.Append("sns_topic"))
}

type NotificationState struct {
	Events   []string `json:"events"`
	SnsTopic string   `json:"sns_topic"`
}