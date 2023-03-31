// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package wafwebacl

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type DefaultAction struct {
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type LoggingConfiguration struct {
	// LogDestination: string, required
	LogDestination terra.StringValue `hcl:"log_destination,attr" validate:"required"`
	// RedactedFields: optional
	RedactedFields *RedactedFields `hcl:"redacted_fields,block"`
}

type RedactedFields struct {
	// FieldToMatch: min=1
	FieldToMatch []FieldToMatch `hcl:"field_to_match,block" validate:"min=1"`
}

type FieldToMatch struct {
	// Data: string, optional
	Data terra.StringValue `hcl:"data,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type Rules struct {
	// Priority: number, required
	Priority terra.NumberValue `hcl:"priority,attr" validate:"required"`
	// RuleId: string, required
	RuleId terra.StringValue `hcl:"rule_id,attr" validate:"required"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// Action: optional
	Action *Action `hcl:"action,block"`
	// OverrideAction: optional
	OverrideAction *OverrideAction `hcl:"override_action,block"`
}

type Action struct {
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type OverrideAction struct {
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type DefaultActionAttributes struct {
	ref terra.Reference
}

func (da DefaultActionAttributes) InternalRef() terra.Reference {
	return da.ref
}

func (da DefaultActionAttributes) InternalWithRef(ref terra.Reference) DefaultActionAttributes {
	return DefaultActionAttributes{ref: ref}
}

func (da DefaultActionAttributes) InternalTokens() hclwrite.Tokens {
	return da.ref.InternalTokens()
}

func (da DefaultActionAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("type"))
}

type LoggingConfigurationAttributes struct {
	ref terra.Reference
}

func (lc LoggingConfigurationAttributes) InternalRef() terra.Reference {
	return lc.ref
}

func (lc LoggingConfigurationAttributes) InternalWithRef(ref terra.Reference) LoggingConfigurationAttributes {
	return LoggingConfigurationAttributes{ref: ref}
}

func (lc LoggingConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return lc.ref.InternalTokens()
}

func (lc LoggingConfigurationAttributes) LogDestination() terra.StringValue {
	return terra.ReferenceAsString(lc.ref.Append("log_destination"))
}

func (lc LoggingConfigurationAttributes) RedactedFields() terra.ListValue[RedactedFieldsAttributes] {
	return terra.ReferenceAsList[RedactedFieldsAttributes](lc.ref.Append("redacted_fields"))
}

type RedactedFieldsAttributes struct {
	ref terra.Reference
}

func (rf RedactedFieldsAttributes) InternalRef() terra.Reference {
	return rf.ref
}

func (rf RedactedFieldsAttributes) InternalWithRef(ref terra.Reference) RedactedFieldsAttributes {
	return RedactedFieldsAttributes{ref: ref}
}

func (rf RedactedFieldsAttributes) InternalTokens() hclwrite.Tokens {
	return rf.ref.InternalTokens()
}

func (rf RedactedFieldsAttributes) FieldToMatch() terra.SetValue[FieldToMatchAttributes] {
	return terra.ReferenceAsSet[FieldToMatchAttributes](rf.ref.Append("field_to_match"))
}

type FieldToMatchAttributes struct {
	ref terra.Reference
}

func (ftm FieldToMatchAttributes) InternalRef() terra.Reference {
	return ftm.ref
}

func (ftm FieldToMatchAttributes) InternalWithRef(ref terra.Reference) FieldToMatchAttributes {
	return FieldToMatchAttributes{ref: ref}
}

func (ftm FieldToMatchAttributes) InternalTokens() hclwrite.Tokens {
	return ftm.ref.InternalTokens()
}

func (ftm FieldToMatchAttributes) Data() terra.StringValue {
	return terra.ReferenceAsString(ftm.ref.Append("data"))
}

func (ftm FieldToMatchAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ftm.ref.Append("type"))
}

type RulesAttributes struct {
	ref terra.Reference
}

func (r RulesAttributes) InternalRef() terra.Reference {
	return r.ref
}

func (r RulesAttributes) InternalWithRef(ref terra.Reference) RulesAttributes {
	return RulesAttributes{ref: ref}
}

func (r RulesAttributes) InternalTokens() hclwrite.Tokens {
	return r.ref.InternalTokens()
}

func (r RulesAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(r.ref.Append("priority"))
}

func (r RulesAttributes) RuleId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("rule_id"))
}

func (r RulesAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("type"))
}

func (r RulesAttributes) Action() terra.ListValue[ActionAttributes] {
	return terra.ReferenceAsList[ActionAttributes](r.ref.Append("action"))
}

func (r RulesAttributes) OverrideAction() terra.ListValue[OverrideActionAttributes] {
	return terra.ReferenceAsList[OverrideActionAttributes](r.ref.Append("override_action"))
}

type ActionAttributes struct {
	ref terra.Reference
}

func (a ActionAttributes) InternalRef() terra.Reference {
	return a.ref
}

func (a ActionAttributes) InternalWithRef(ref terra.Reference) ActionAttributes {
	return ActionAttributes{ref: ref}
}

func (a ActionAttributes) InternalTokens() hclwrite.Tokens {
	return a.ref.InternalTokens()
}

func (a ActionAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("type"))
}

type OverrideActionAttributes struct {
	ref terra.Reference
}

func (oa OverrideActionAttributes) InternalRef() terra.Reference {
	return oa.ref
}

func (oa OverrideActionAttributes) InternalWithRef(ref terra.Reference) OverrideActionAttributes {
	return OverrideActionAttributes{ref: ref}
}

func (oa OverrideActionAttributes) InternalTokens() hclwrite.Tokens {
	return oa.ref.InternalTokens()
}

func (oa OverrideActionAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(oa.ref.Append("type"))
}

type DefaultActionState struct {
	Type string `json:"type"`
}

type LoggingConfigurationState struct {
	LogDestination string                `json:"log_destination"`
	RedactedFields []RedactedFieldsState `json:"redacted_fields"`
}

type RedactedFieldsState struct {
	FieldToMatch []FieldToMatchState `json:"field_to_match"`
}

type FieldToMatchState struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

type RulesState struct {
	Priority       float64               `json:"priority"`
	RuleId         string                `json:"rule_id"`
	Type           string                `json:"type"`
	Action         []ActionState         `json:"action"`
	OverrideAction []OverrideActionState `json:"override_action"`
}

type ActionState struct {
	Type string `json:"type"`
}

type OverrideActionState struct {
	Type string `json:"type"`
}