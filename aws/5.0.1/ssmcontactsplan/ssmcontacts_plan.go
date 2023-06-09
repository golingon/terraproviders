// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package ssmcontactsplan

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Stage struct {
	// DurationInMinutes: number, required
	DurationInMinutes terra.NumberValue `hcl:"duration_in_minutes,attr" validate:"required"`
	// Target: min=0
	Target []Target `hcl:"target,block" validate:"min=0"`
}

type Target struct {
	// ChannelTargetInfo: optional
	ChannelTargetInfo *ChannelTargetInfo `hcl:"channel_target_info,block"`
	// ContactTargetInfo: optional
	ContactTargetInfo *ContactTargetInfo `hcl:"contact_target_info,block"`
}

type ChannelTargetInfo struct {
	// ContactChannelId: string, required
	ContactChannelId terra.StringValue `hcl:"contact_channel_id,attr" validate:"required"`
	// RetryIntervalInMinutes: number, optional
	RetryIntervalInMinutes terra.NumberValue `hcl:"retry_interval_in_minutes,attr"`
}

type ContactTargetInfo struct {
	// ContactId: string, optional
	ContactId terra.StringValue `hcl:"contact_id,attr"`
	// IsEssential: bool, required
	IsEssential terra.BoolValue `hcl:"is_essential,attr" validate:"required"`
}

type StageAttributes struct {
	ref terra.Reference
}

func (s StageAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s StageAttributes) InternalWithRef(ref terra.Reference) StageAttributes {
	return StageAttributes{ref: ref}
}

func (s StageAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s StageAttributes) DurationInMinutes() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("duration_in_minutes"))
}

func (s StageAttributes) Target() terra.ListValue[TargetAttributes] {
	return terra.ReferenceAsList[TargetAttributes](s.ref.Append("target"))
}

type TargetAttributes struct {
	ref terra.Reference
}

func (t TargetAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TargetAttributes) InternalWithRef(ref terra.Reference) TargetAttributes {
	return TargetAttributes{ref: ref}
}

func (t TargetAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TargetAttributes) ChannelTargetInfo() terra.ListValue[ChannelTargetInfoAttributes] {
	return terra.ReferenceAsList[ChannelTargetInfoAttributes](t.ref.Append("channel_target_info"))
}

func (t TargetAttributes) ContactTargetInfo() terra.ListValue[ContactTargetInfoAttributes] {
	return terra.ReferenceAsList[ContactTargetInfoAttributes](t.ref.Append("contact_target_info"))
}

type ChannelTargetInfoAttributes struct {
	ref terra.Reference
}

func (cti ChannelTargetInfoAttributes) InternalRef() (terra.Reference, error) {
	return cti.ref, nil
}

func (cti ChannelTargetInfoAttributes) InternalWithRef(ref terra.Reference) ChannelTargetInfoAttributes {
	return ChannelTargetInfoAttributes{ref: ref}
}

func (cti ChannelTargetInfoAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cti.ref.InternalTokens()
}

func (cti ChannelTargetInfoAttributes) ContactChannelId() terra.StringValue {
	return terra.ReferenceAsString(cti.ref.Append("contact_channel_id"))
}

func (cti ChannelTargetInfoAttributes) RetryIntervalInMinutes() terra.NumberValue {
	return terra.ReferenceAsNumber(cti.ref.Append("retry_interval_in_minutes"))
}

type ContactTargetInfoAttributes struct {
	ref terra.Reference
}

func (cti ContactTargetInfoAttributes) InternalRef() (terra.Reference, error) {
	return cti.ref, nil
}

func (cti ContactTargetInfoAttributes) InternalWithRef(ref terra.Reference) ContactTargetInfoAttributes {
	return ContactTargetInfoAttributes{ref: ref}
}

func (cti ContactTargetInfoAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cti.ref.InternalTokens()
}

func (cti ContactTargetInfoAttributes) ContactId() terra.StringValue {
	return terra.ReferenceAsString(cti.ref.Append("contact_id"))
}

func (cti ContactTargetInfoAttributes) IsEssential() terra.BoolValue {
	return terra.ReferenceAsBool(cti.ref.Append("is_essential"))
}

type StageState struct {
	DurationInMinutes float64       `json:"duration_in_minutes"`
	Target            []TargetState `json:"target"`
}

type TargetState struct {
	ChannelTargetInfo []ChannelTargetInfoState `json:"channel_target_info"`
	ContactTargetInfo []ContactTargetInfoState `json:"contact_target_info"`
}

type ChannelTargetInfoState struct {
	ContactChannelId       string  `json:"contact_channel_id"`
	RetryIntervalInMinutes float64 `json:"retry_interval_in_minutes"`
}

type ContactTargetInfoState struct {
	ContactId   string `json:"contact_id"`
	IsEssential bool   `json:"is_essential"`
}
