// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package cognitoriskconfiguration

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AccountTakeoverRiskConfiguration struct {
	// AccountTakeoverRiskConfigurationActions: required
	Actions *AccountTakeoverRiskConfigurationActions `hcl:"actions,block" validate:"required"`
	// NotifyConfiguration: required
	NotifyConfiguration *NotifyConfiguration `hcl:"notify_configuration,block" validate:"required"`
}

type AccountTakeoverRiskConfigurationActions struct {
	// HighAction: optional
	HighAction *HighAction `hcl:"high_action,block"`
	// LowAction: optional
	LowAction *LowAction `hcl:"low_action,block"`
	// MediumAction: optional
	MediumAction *MediumAction `hcl:"medium_action,block"`
}

type HighAction struct {
	// EventAction: string, required
	EventAction terra.StringValue `hcl:"event_action,attr" validate:"required"`
	// Notify: bool, required
	Notify terra.BoolValue `hcl:"notify,attr" validate:"required"`
}

type LowAction struct {
	// EventAction: string, required
	EventAction terra.StringValue `hcl:"event_action,attr" validate:"required"`
	// Notify: bool, required
	Notify terra.BoolValue `hcl:"notify,attr" validate:"required"`
}

type MediumAction struct {
	// EventAction: string, required
	EventAction terra.StringValue `hcl:"event_action,attr" validate:"required"`
	// Notify: bool, required
	Notify terra.BoolValue `hcl:"notify,attr" validate:"required"`
}

type NotifyConfiguration struct {
	// From: string, optional
	From terra.StringValue `hcl:"from,attr"`
	// ReplyTo: string, optional
	ReplyTo terra.StringValue `hcl:"reply_to,attr"`
	// SourceArn: string, required
	SourceArn terra.StringValue `hcl:"source_arn,attr" validate:"required"`
	// BlockEmail: optional
	BlockEmail *BlockEmail `hcl:"block_email,block"`
	// MfaEmail: optional
	MfaEmail *MfaEmail `hcl:"mfa_email,block"`
	// NoActionEmail: optional
	NoActionEmail *NoActionEmail `hcl:"no_action_email,block"`
}

type BlockEmail struct {
	// HtmlBody: string, required
	HtmlBody terra.StringValue `hcl:"html_body,attr" validate:"required"`
	// Subject: string, required
	Subject terra.StringValue `hcl:"subject,attr" validate:"required"`
	// TextBody: string, required
	TextBody terra.StringValue `hcl:"text_body,attr" validate:"required"`
}

type MfaEmail struct {
	// HtmlBody: string, required
	HtmlBody terra.StringValue `hcl:"html_body,attr" validate:"required"`
	// Subject: string, required
	Subject terra.StringValue `hcl:"subject,attr" validate:"required"`
	// TextBody: string, required
	TextBody terra.StringValue `hcl:"text_body,attr" validate:"required"`
}

type NoActionEmail struct {
	// HtmlBody: string, required
	HtmlBody terra.StringValue `hcl:"html_body,attr" validate:"required"`
	// Subject: string, required
	Subject terra.StringValue `hcl:"subject,attr" validate:"required"`
	// TextBody: string, required
	TextBody terra.StringValue `hcl:"text_body,attr" validate:"required"`
}

type CompromisedCredentialsRiskConfiguration struct {
	// EventFilter: set of string, optional
	EventFilter terra.SetValue[terra.StringValue] `hcl:"event_filter,attr"`
	// CompromisedCredentialsRiskConfigurationActions: required
	Actions *CompromisedCredentialsRiskConfigurationActions `hcl:"actions,block" validate:"required"`
}

type CompromisedCredentialsRiskConfigurationActions struct {
	// EventAction: string, required
	EventAction terra.StringValue `hcl:"event_action,attr" validate:"required"`
}

type RiskExceptionConfiguration struct {
	// BlockedIpRangeList: set of string, optional
	BlockedIpRangeList terra.SetValue[terra.StringValue] `hcl:"blocked_ip_range_list,attr"`
	// SkippedIpRangeList: set of string, optional
	SkippedIpRangeList terra.SetValue[terra.StringValue] `hcl:"skipped_ip_range_list,attr"`
}

type AccountTakeoverRiskConfigurationAttributes struct {
	ref terra.Reference
}

func (atrc AccountTakeoverRiskConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return atrc.ref, nil
}

func (atrc AccountTakeoverRiskConfigurationAttributes) InternalWithRef(ref terra.Reference) AccountTakeoverRiskConfigurationAttributes {
	return AccountTakeoverRiskConfigurationAttributes{ref: ref}
}

func (atrc AccountTakeoverRiskConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return atrc.ref.InternalTokens()
}

func (atrc AccountTakeoverRiskConfigurationAttributes) Actions() terra.ListValue[AccountTakeoverRiskConfigurationActionsAttributes] {
	return terra.ReferenceAsList[AccountTakeoverRiskConfigurationActionsAttributes](atrc.ref.Append("actions"))
}

func (atrc AccountTakeoverRiskConfigurationAttributes) NotifyConfiguration() terra.ListValue[NotifyConfigurationAttributes] {
	return terra.ReferenceAsList[NotifyConfigurationAttributes](atrc.ref.Append("notify_configuration"))
}

type AccountTakeoverRiskConfigurationActionsAttributes struct {
	ref terra.Reference
}

func (a AccountTakeoverRiskConfigurationActionsAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a AccountTakeoverRiskConfigurationActionsAttributes) InternalWithRef(ref terra.Reference) AccountTakeoverRiskConfigurationActionsAttributes {
	return AccountTakeoverRiskConfigurationActionsAttributes{ref: ref}
}

func (a AccountTakeoverRiskConfigurationActionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a AccountTakeoverRiskConfigurationActionsAttributes) HighAction() terra.ListValue[HighActionAttributes] {
	return terra.ReferenceAsList[HighActionAttributes](a.ref.Append("high_action"))
}

func (a AccountTakeoverRiskConfigurationActionsAttributes) LowAction() terra.ListValue[LowActionAttributes] {
	return terra.ReferenceAsList[LowActionAttributes](a.ref.Append("low_action"))
}

func (a AccountTakeoverRiskConfigurationActionsAttributes) MediumAction() terra.ListValue[MediumActionAttributes] {
	return terra.ReferenceAsList[MediumActionAttributes](a.ref.Append("medium_action"))
}

type HighActionAttributes struct {
	ref terra.Reference
}

func (ha HighActionAttributes) InternalRef() (terra.Reference, error) {
	return ha.ref, nil
}

func (ha HighActionAttributes) InternalWithRef(ref terra.Reference) HighActionAttributes {
	return HighActionAttributes{ref: ref}
}

func (ha HighActionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ha.ref.InternalTokens()
}

func (ha HighActionAttributes) EventAction() terra.StringValue {
	return terra.ReferenceAsString(ha.ref.Append("event_action"))
}

func (ha HighActionAttributes) Notify() terra.BoolValue {
	return terra.ReferenceAsBool(ha.ref.Append("notify"))
}

type LowActionAttributes struct {
	ref terra.Reference
}

func (la LowActionAttributes) InternalRef() (terra.Reference, error) {
	return la.ref, nil
}

func (la LowActionAttributes) InternalWithRef(ref terra.Reference) LowActionAttributes {
	return LowActionAttributes{ref: ref}
}

func (la LowActionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return la.ref.InternalTokens()
}

func (la LowActionAttributes) EventAction() terra.StringValue {
	return terra.ReferenceAsString(la.ref.Append("event_action"))
}

func (la LowActionAttributes) Notify() terra.BoolValue {
	return terra.ReferenceAsBool(la.ref.Append("notify"))
}

type MediumActionAttributes struct {
	ref terra.Reference
}

func (ma MediumActionAttributes) InternalRef() (terra.Reference, error) {
	return ma.ref, nil
}

func (ma MediumActionAttributes) InternalWithRef(ref terra.Reference) MediumActionAttributes {
	return MediumActionAttributes{ref: ref}
}

func (ma MediumActionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ma.ref.InternalTokens()
}

func (ma MediumActionAttributes) EventAction() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("event_action"))
}

func (ma MediumActionAttributes) Notify() terra.BoolValue {
	return terra.ReferenceAsBool(ma.ref.Append("notify"))
}

type NotifyConfigurationAttributes struct {
	ref terra.Reference
}

func (nc NotifyConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return nc.ref, nil
}

func (nc NotifyConfigurationAttributes) InternalWithRef(ref terra.Reference) NotifyConfigurationAttributes {
	return NotifyConfigurationAttributes{ref: ref}
}

func (nc NotifyConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return nc.ref.InternalTokens()
}

func (nc NotifyConfigurationAttributes) From() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("from"))
}

func (nc NotifyConfigurationAttributes) ReplyTo() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("reply_to"))
}

func (nc NotifyConfigurationAttributes) SourceArn() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("source_arn"))
}

func (nc NotifyConfigurationAttributes) BlockEmail() terra.ListValue[BlockEmailAttributes] {
	return terra.ReferenceAsList[BlockEmailAttributes](nc.ref.Append("block_email"))
}

func (nc NotifyConfigurationAttributes) MfaEmail() terra.ListValue[MfaEmailAttributes] {
	return terra.ReferenceAsList[MfaEmailAttributes](nc.ref.Append("mfa_email"))
}

func (nc NotifyConfigurationAttributes) NoActionEmail() terra.ListValue[NoActionEmailAttributes] {
	return terra.ReferenceAsList[NoActionEmailAttributes](nc.ref.Append("no_action_email"))
}

type BlockEmailAttributes struct {
	ref terra.Reference
}

func (be BlockEmailAttributes) InternalRef() (terra.Reference, error) {
	return be.ref, nil
}

func (be BlockEmailAttributes) InternalWithRef(ref terra.Reference) BlockEmailAttributes {
	return BlockEmailAttributes{ref: ref}
}

func (be BlockEmailAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return be.ref.InternalTokens()
}

func (be BlockEmailAttributes) HtmlBody() terra.StringValue {
	return terra.ReferenceAsString(be.ref.Append("html_body"))
}

func (be BlockEmailAttributes) Subject() terra.StringValue {
	return terra.ReferenceAsString(be.ref.Append("subject"))
}

func (be BlockEmailAttributes) TextBody() terra.StringValue {
	return terra.ReferenceAsString(be.ref.Append("text_body"))
}

type MfaEmailAttributes struct {
	ref terra.Reference
}

func (me MfaEmailAttributes) InternalRef() (terra.Reference, error) {
	return me.ref, nil
}

func (me MfaEmailAttributes) InternalWithRef(ref terra.Reference) MfaEmailAttributes {
	return MfaEmailAttributes{ref: ref}
}

func (me MfaEmailAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return me.ref.InternalTokens()
}

func (me MfaEmailAttributes) HtmlBody() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("html_body"))
}

func (me MfaEmailAttributes) Subject() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("subject"))
}

func (me MfaEmailAttributes) TextBody() terra.StringValue {
	return terra.ReferenceAsString(me.ref.Append("text_body"))
}

type NoActionEmailAttributes struct {
	ref terra.Reference
}

func (nae NoActionEmailAttributes) InternalRef() (terra.Reference, error) {
	return nae.ref, nil
}

func (nae NoActionEmailAttributes) InternalWithRef(ref terra.Reference) NoActionEmailAttributes {
	return NoActionEmailAttributes{ref: ref}
}

func (nae NoActionEmailAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return nae.ref.InternalTokens()
}

func (nae NoActionEmailAttributes) HtmlBody() terra.StringValue {
	return terra.ReferenceAsString(nae.ref.Append("html_body"))
}

func (nae NoActionEmailAttributes) Subject() terra.StringValue {
	return terra.ReferenceAsString(nae.ref.Append("subject"))
}

func (nae NoActionEmailAttributes) TextBody() terra.StringValue {
	return terra.ReferenceAsString(nae.ref.Append("text_body"))
}

type CompromisedCredentialsRiskConfigurationAttributes struct {
	ref terra.Reference
}

func (ccrc CompromisedCredentialsRiskConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return ccrc.ref, nil
}

func (ccrc CompromisedCredentialsRiskConfigurationAttributes) InternalWithRef(ref terra.Reference) CompromisedCredentialsRiskConfigurationAttributes {
	return CompromisedCredentialsRiskConfigurationAttributes{ref: ref}
}

func (ccrc CompromisedCredentialsRiskConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ccrc.ref.InternalTokens()
}

func (ccrc CompromisedCredentialsRiskConfigurationAttributes) EventFilter() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ccrc.ref.Append("event_filter"))
}

func (ccrc CompromisedCredentialsRiskConfigurationAttributes) Actions() terra.ListValue[CompromisedCredentialsRiskConfigurationActionsAttributes] {
	return terra.ReferenceAsList[CompromisedCredentialsRiskConfigurationActionsAttributes](ccrc.ref.Append("actions"))
}

type CompromisedCredentialsRiskConfigurationActionsAttributes struct {
	ref terra.Reference
}

func (a CompromisedCredentialsRiskConfigurationActionsAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a CompromisedCredentialsRiskConfigurationActionsAttributes) InternalWithRef(ref terra.Reference) CompromisedCredentialsRiskConfigurationActionsAttributes {
	return CompromisedCredentialsRiskConfigurationActionsAttributes{ref: ref}
}

func (a CompromisedCredentialsRiskConfigurationActionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a CompromisedCredentialsRiskConfigurationActionsAttributes) EventAction() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("event_action"))
}

type RiskExceptionConfigurationAttributes struct {
	ref terra.Reference
}

func (rec RiskExceptionConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return rec.ref, nil
}

func (rec RiskExceptionConfigurationAttributes) InternalWithRef(ref terra.Reference) RiskExceptionConfigurationAttributes {
	return RiskExceptionConfigurationAttributes{ref: ref}
}

func (rec RiskExceptionConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rec.ref.InternalTokens()
}

func (rec RiskExceptionConfigurationAttributes) BlockedIpRangeList() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rec.ref.Append("blocked_ip_range_list"))
}

func (rec RiskExceptionConfigurationAttributes) SkippedIpRangeList() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rec.ref.Append("skipped_ip_range_list"))
}

type AccountTakeoverRiskConfigurationState struct {
	Actions             []AccountTakeoverRiskConfigurationActionsState `json:"actions"`
	NotifyConfiguration []NotifyConfigurationState                     `json:"notify_configuration"`
}

type AccountTakeoverRiskConfigurationActionsState struct {
	HighAction   []HighActionState   `json:"high_action"`
	LowAction    []LowActionState    `json:"low_action"`
	MediumAction []MediumActionState `json:"medium_action"`
}

type HighActionState struct {
	EventAction string `json:"event_action"`
	Notify      bool   `json:"notify"`
}

type LowActionState struct {
	EventAction string `json:"event_action"`
	Notify      bool   `json:"notify"`
}

type MediumActionState struct {
	EventAction string `json:"event_action"`
	Notify      bool   `json:"notify"`
}

type NotifyConfigurationState struct {
	From          string               `json:"from"`
	ReplyTo       string               `json:"reply_to"`
	SourceArn     string               `json:"source_arn"`
	BlockEmail    []BlockEmailState    `json:"block_email"`
	MfaEmail      []MfaEmailState      `json:"mfa_email"`
	NoActionEmail []NoActionEmailState `json:"no_action_email"`
}

type BlockEmailState struct {
	HtmlBody string `json:"html_body"`
	Subject  string `json:"subject"`
	TextBody string `json:"text_body"`
}

type MfaEmailState struct {
	HtmlBody string `json:"html_body"`
	Subject  string `json:"subject"`
	TextBody string `json:"text_body"`
}

type NoActionEmailState struct {
	HtmlBody string `json:"html_body"`
	Subject  string `json:"subject"`
	TextBody string `json:"text_body"`
}

type CompromisedCredentialsRiskConfigurationState struct {
	EventFilter []string                                              `json:"event_filter"`
	Actions     []CompromisedCredentialsRiskConfigurationActionsState `json:"actions"`
}

type CompromisedCredentialsRiskConfigurationActionsState struct {
	EventAction string `json:"event_action"`
}

type RiskExceptionConfigurationState struct {
	BlockedIpRangeList []string `json:"blocked_ip_range_list"`
	SkippedIpRangeList []string `json:"skipped_ip_range_list"`
}
