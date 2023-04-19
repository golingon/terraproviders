// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package xraygroup

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type InsightsConfiguration struct {
	// InsightsEnabled: bool, required
	InsightsEnabled terra.BoolValue `hcl:"insights_enabled,attr" validate:"required"`
	// NotificationsEnabled: bool, optional
	NotificationsEnabled terra.BoolValue `hcl:"notifications_enabled,attr"`
}

type InsightsConfigurationAttributes struct {
	ref terra.Reference
}

func (ic InsightsConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return ic.ref, nil
}

func (ic InsightsConfigurationAttributes) InternalWithRef(ref terra.Reference) InsightsConfigurationAttributes {
	return InsightsConfigurationAttributes{ref: ref}
}

func (ic InsightsConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ic.ref.InternalTokens()
}

func (ic InsightsConfigurationAttributes) InsightsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ic.ref.Append("insights_enabled"))
}

func (ic InsightsConfigurationAttributes) NotificationsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ic.ref.Append("notifications_enabled"))
}

type InsightsConfigurationState struct {
	InsightsEnabled      bool `json:"insights_enabled"`
	NotificationsEnabled bool `json:"notifications_enabled"`
}
