// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package lexintent

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ConclusionStatement struct {
	// ResponseCard: string, optional
	ResponseCard terra.StringValue `hcl:"response_card,attr"`
	// ConclusionStatementMessage: min=1,max=15
	Message []ConclusionStatementMessage `hcl:"message,block" validate:"min=1,max=15"`
}

type ConclusionStatementMessage struct {
	// Content: string, required
	Content terra.StringValue `hcl:"content,attr" validate:"required"`
	// ContentType: string, required
	ContentType terra.StringValue `hcl:"content_type,attr" validate:"required"`
	// GroupNumber: number, optional
	GroupNumber terra.NumberValue `hcl:"group_number,attr"`
}

type ConfirmationPrompt struct {
	// MaxAttempts: number, required
	MaxAttempts terra.NumberValue `hcl:"max_attempts,attr" validate:"required"`
	// ResponseCard: string, optional
	ResponseCard terra.StringValue `hcl:"response_card,attr"`
	// ConfirmationPromptMessage: min=1,max=15
	Message []ConfirmationPromptMessage `hcl:"message,block" validate:"min=1,max=15"`
}

type ConfirmationPromptMessage struct {
	// Content: string, required
	Content terra.StringValue `hcl:"content,attr" validate:"required"`
	// ContentType: string, required
	ContentType terra.StringValue `hcl:"content_type,attr" validate:"required"`
	// GroupNumber: number, optional
	GroupNumber terra.NumberValue `hcl:"group_number,attr"`
}

type DialogCodeHook struct {
	// MessageVersion: string, required
	MessageVersion terra.StringValue `hcl:"message_version,attr" validate:"required"`
	// Uri: string, required
	Uri terra.StringValue `hcl:"uri,attr" validate:"required"`
}

type FollowUpPrompt struct {
	// Prompt: required
	Prompt *Prompt `hcl:"prompt,block" validate:"required"`
	// FollowUpPromptRejectionStatement: required
	RejectionStatement *FollowUpPromptRejectionStatement `hcl:"rejection_statement,block" validate:"required"`
}

type Prompt struct {
	// MaxAttempts: number, required
	MaxAttempts terra.NumberValue `hcl:"max_attempts,attr" validate:"required"`
	// ResponseCard: string, optional
	ResponseCard terra.StringValue `hcl:"response_card,attr"`
	// PromptMessage: min=1,max=15
	Message []PromptMessage `hcl:"message,block" validate:"min=1,max=15"`
}

type PromptMessage struct {
	// Content: string, required
	Content terra.StringValue `hcl:"content,attr" validate:"required"`
	// ContentType: string, required
	ContentType terra.StringValue `hcl:"content_type,attr" validate:"required"`
	// GroupNumber: number, optional
	GroupNumber terra.NumberValue `hcl:"group_number,attr"`
}

type FollowUpPromptRejectionStatement struct {
	// ResponseCard: string, optional
	ResponseCard terra.StringValue `hcl:"response_card,attr"`
	// FollowUpPromptRejectionStatementMessage: min=1,max=15
	Message []FollowUpPromptRejectionStatementMessage `hcl:"message,block" validate:"min=1,max=15"`
}

type FollowUpPromptRejectionStatementMessage struct {
	// Content: string, required
	Content terra.StringValue `hcl:"content,attr" validate:"required"`
	// ContentType: string, required
	ContentType terra.StringValue `hcl:"content_type,attr" validate:"required"`
	// GroupNumber: number, optional
	GroupNumber terra.NumberValue `hcl:"group_number,attr"`
}

type FulfillmentActivity struct {
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// CodeHook: optional
	CodeHook *CodeHook `hcl:"code_hook,block"`
}

type CodeHook struct {
	// MessageVersion: string, required
	MessageVersion terra.StringValue `hcl:"message_version,attr" validate:"required"`
	// Uri: string, required
	Uri terra.StringValue `hcl:"uri,attr" validate:"required"`
}

type RejectionStatement struct {
	// ResponseCard: string, optional
	ResponseCard terra.StringValue `hcl:"response_card,attr"`
	// RejectionStatementMessage: min=1,max=15
	Message []RejectionStatementMessage `hcl:"message,block" validate:"min=1,max=15"`
}

type RejectionStatementMessage struct {
	// Content: string, required
	Content terra.StringValue `hcl:"content,attr" validate:"required"`
	// ContentType: string, required
	ContentType terra.StringValue `hcl:"content_type,attr" validate:"required"`
	// GroupNumber: number, optional
	GroupNumber terra.NumberValue `hcl:"group_number,attr"`
}

type Slot struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Priority: number, optional
	Priority terra.NumberValue `hcl:"priority,attr"`
	// ResponseCard: string, optional
	ResponseCard terra.StringValue `hcl:"response_card,attr"`
	// SampleUtterances: list of string, optional
	SampleUtterances terra.ListValue[terra.StringValue] `hcl:"sample_utterances,attr"`
	// SlotConstraint: string, required
	SlotConstraint terra.StringValue `hcl:"slot_constraint,attr" validate:"required"`
	// SlotType: string, required
	SlotType terra.StringValue `hcl:"slot_type,attr" validate:"required"`
	// SlotTypeVersion: string, optional
	SlotTypeVersion terra.StringValue `hcl:"slot_type_version,attr"`
	// ValueElicitationPrompt: optional
	ValueElicitationPrompt *ValueElicitationPrompt `hcl:"value_elicitation_prompt,block"`
}

type ValueElicitationPrompt struct {
	// MaxAttempts: number, required
	MaxAttempts terra.NumberValue `hcl:"max_attempts,attr" validate:"required"`
	// ResponseCard: string, optional
	ResponseCard terra.StringValue `hcl:"response_card,attr"`
	// ValueElicitationPromptMessage: min=1,max=15
	Message []ValueElicitationPromptMessage `hcl:"message,block" validate:"min=1,max=15"`
}

type ValueElicitationPromptMessage struct {
	// Content: string, required
	Content terra.StringValue `hcl:"content,attr" validate:"required"`
	// ContentType: string, required
	ContentType terra.StringValue `hcl:"content_type,attr" validate:"required"`
	// GroupNumber: number, optional
	GroupNumber terra.NumberValue `hcl:"group_number,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type ConclusionStatementAttributes struct {
	ref terra.Reference
}

func (cs ConclusionStatementAttributes) InternalRef() terra.Reference {
	return cs.ref
}

func (cs ConclusionStatementAttributes) InternalWithRef(ref terra.Reference) ConclusionStatementAttributes {
	return ConclusionStatementAttributes{ref: ref}
}

func (cs ConclusionStatementAttributes) InternalTokens() hclwrite.Tokens {
	return cs.ref.InternalTokens()
}

func (cs ConclusionStatementAttributes) ResponseCard() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("response_card"))
}

func (cs ConclusionStatementAttributes) Message() terra.SetValue[ConclusionStatementMessageAttributes] {
	return terra.ReferenceAsSet[ConclusionStatementMessageAttributes](cs.ref.Append("message"))
}

type ConclusionStatementMessageAttributes struct {
	ref terra.Reference
}

func (m ConclusionStatementMessageAttributes) InternalRef() terra.Reference {
	return m.ref
}

func (m ConclusionStatementMessageAttributes) InternalWithRef(ref terra.Reference) ConclusionStatementMessageAttributes {
	return ConclusionStatementMessageAttributes{ref: ref}
}

func (m ConclusionStatementMessageAttributes) InternalTokens() hclwrite.Tokens {
	return m.ref.InternalTokens()
}

func (m ConclusionStatementMessageAttributes) Content() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("content"))
}

func (m ConclusionStatementMessageAttributes) ContentType() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("content_type"))
}

func (m ConclusionStatementMessageAttributes) GroupNumber() terra.NumberValue {
	return terra.ReferenceAsNumber(m.ref.Append("group_number"))
}

type ConfirmationPromptAttributes struct {
	ref terra.Reference
}

func (cp ConfirmationPromptAttributes) InternalRef() terra.Reference {
	return cp.ref
}

func (cp ConfirmationPromptAttributes) InternalWithRef(ref terra.Reference) ConfirmationPromptAttributes {
	return ConfirmationPromptAttributes{ref: ref}
}

func (cp ConfirmationPromptAttributes) InternalTokens() hclwrite.Tokens {
	return cp.ref.InternalTokens()
}

func (cp ConfirmationPromptAttributes) MaxAttempts() terra.NumberValue {
	return terra.ReferenceAsNumber(cp.ref.Append("max_attempts"))
}

func (cp ConfirmationPromptAttributes) ResponseCard() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("response_card"))
}

func (cp ConfirmationPromptAttributes) Message() terra.SetValue[ConfirmationPromptMessageAttributes] {
	return terra.ReferenceAsSet[ConfirmationPromptMessageAttributes](cp.ref.Append("message"))
}

type ConfirmationPromptMessageAttributes struct {
	ref terra.Reference
}

func (m ConfirmationPromptMessageAttributes) InternalRef() terra.Reference {
	return m.ref
}

func (m ConfirmationPromptMessageAttributes) InternalWithRef(ref terra.Reference) ConfirmationPromptMessageAttributes {
	return ConfirmationPromptMessageAttributes{ref: ref}
}

func (m ConfirmationPromptMessageAttributes) InternalTokens() hclwrite.Tokens {
	return m.ref.InternalTokens()
}

func (m ConfirmationPromptMessageAttributes) Content() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("content"))
}

func (m ConfirmationPromptMessageAttributes) ContentType() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("content_type"))
}

func (m ConfirmationPromptMessageAttributes) GroupNumber() terra.NumberValue {
	return terra.ReferenceAsNumber(m.ref.Append("group_number"))
}

type DialogCodeHookAttributes struct {
	ref terra.Reference
}

func (dch DialogCodeHookAttributes) InternalRef() terra.Reference {
	return dch.ref
}

func (dch DialogCodeHookAttributes) InternalWithRef(ref terra.Reference) DialogCodeHookAttributes {
	return DialogCodeHookAttributes{ref: ref}
}

func (dch DialogCodeHookAttributes) InternalTokens() hclwrite.Tokens {
	return dch.ref.InternalTokens()
}

func (dch DialogCodeHookAttributes) MessageVersion() terra.StringValue {
	return terra.ReferenceAsString(dch.ref.Append("message_version"))
}

func (dch DialogCodeHookAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(dch.ref.Append("uri"))
}

type FollowUpPromptAttributes struct {
	ref terra.Reference
}

func (fup FollowUpPromptAttributes) InternalRef() terra.Reference {
	return fup.ref
}

func (fup FollowUpPromptAttributes) InternalWithRef(ref terra.Reference) FollowUpPromptAttributes {
	return FollowUpPromptAttributes{ref: ref}
}

func (fup FollowUpPromptAttributes) InternalTokens() hclwrite.Tokens {
	return fup.ref.InternalTokens()
}

func (fup FollowUpPromptAttributes) Prompt() terra.ListValue[PromptAttributes] {
	return terra.ReferenceAsList[PromptAttributes](fup.ref.Append("prompt"))
}

func (fup FollowUpPromptAttributes) RejectionStatement() terra.ListValue[FollowUpPromptRejectionStatementAttributes] {
	return terra.ReferenceAsList[FollowUpPromptRejectionStatementAttributes](fup.ref.Append("rejection_statement"))
}

type PromptAttributes struct {
	ref terra.Reference
}

func (p PromptAttributes) InternalRef() terra.Reference {
	return p.ref
}

func (p PromptAttributes) InternalWithRef(ref terra.Reference) PromptAttributes {
	return PromptAttributes{ref: ref}
}

func (p PromptAttributes) InternalTokens() hclwrite.Tokens {
	return p.ref.InternalTokens()
}

func (p PromptAttributes) MaxAttempts() terra.NumberValue {
	return terra.ReferenceAsNumber(p.ref.Append("max_attempts"))
}

func (p PromptAttributes) ResponseCard() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("response_card"))
}

func (p PromptAttributes) Message() terra.SetValue[PromptMessageAttributes] {
	return terra.ReferenceAsSet[PromptMessageAttributes](p.ref.Append("message"))
}

type PromptMessageAttributes struct {
	ref terra.Reference
}

func (m PromptMessageAttributes) InternalRef() terra.Reference {
	return m.ref
}

func (m PromptMessageAttributes) InternalWithRef(ref terra.Reference) PromptMessageAttributes {
	return PromptMessageAttributes{ref: ref}
}

func (m PromptMessageAttributes) InternalTokens() hclwrite.Tokens {
	return m.ref.InternalTokens()
}

func (m PromptMessageAttributes) Content() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("content"))
}

func (m PromptMessageAttributes) ContentType() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("content_type"))
}

func (m PromptMessageAttributes) GroupNumber() terra.NumberValue {
	return terra.ReferenceAsNumber(m.ref.Append("group_number"))
}

type FollowUpPromptRejectionStatementAttributes struct {
	ref terra.Reference
}

func (rs FollowUpPromptRejectionStatementAttributes) InternalRef() terra.Reference {
	return rs.ref
}

func (rs FollowUpPromptRejectionStatementAttributes) InternalWithRef(ref terra.Reference) FollowUpPromptRejectionStatementAttributes {
	return FollowUpPromptRejectionStatementAttributes{ref: ref}
}

func (rs FollowUpPromptRejectionStatementAttributes) InternalTokens() hclwrite.Tokens {
	return rs.ref.InternalTokens()
}

func (rs FollowUpPromptRejectionStatementAttributes) ResponseCard() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("response_card"))
}

func (rs FollowUpPromptRejectionStatementAttributes) Message() terra.SetValue[FollowUpPromptRejectionStatementMessageAttributes] {
	return terra.ReferenceAsSet[FollowUpPromptRejectionStatementMessageAttributes](rs.ref.Append("message"))
}

type FollowUpPromptRejectionStatementMessageAttributes struct {
	ref terra.Reference
}

func (m FollowUpPromptRejectionStatementMessageAttributes) InternalRef() terra.Reference {
	return m.ref
}

func (m FollowUpPromptRejectionStatementMessageAttributes) InternalWithRef(ref terra.Reference) FollowUpPromptRejectionStatementMessageAttributes {
	return FollowUpPromptRejectionStatementMessageAttributes{ref: ref}
}

func (m FollowUpPromptRejectionStatementMessageAttributes) InternalTokens() hclwrite.Tokens {
	return m.ref.InternalTokens()
}

func (m FollowUpPromptRejectionStatementMessageAttributes) Content() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("content"))
}

func (m FollowUpPromptRejectionStatementMessageAttributes) ContentType() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("content_type"))
}

func (m FollowUpPromptRejectionStatementMessageAttributes) GroupNumber() terra.NumberValue {
	return terra.ReferenceAsNumber(m.ref.Append("group_number"))
}

type FulfillmentActivityAttributes struct {
	ref terra.Reference
}

func (fa FulfillmentActivityAttributes) InternalRef() terra.Reference {
	return fa.ref
}

func (fa FulfillmentActivityAttributes) InternalWithRef(ref terra.Reference) FulfillmentActivityAttributes {
	return FulfillmentActivityAttributes{ref: ref}
}

func (fa FulfillmentActivityAttributes) InternalTokens() hclwrite.Tokens {
	return fa.ref.InternalTokens()
}

func (fa FulfillmentActivityAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(fa.ref.Append("type"))
}

func (fa FulfillmentActivityAttributes) CodeHook() terra.ListValue[CodeHookAttributes] {
	return terra.ReferenceAsList[CodeHookAttributes](fa.ref.Append("code_hook"))
}

type CodeHookAttributes struct {
	ref terra.Reference
}

func (ch CodeHookAttributes) InternalRef() terra.Reference {
	return ch.ref
}

func (ch CodeHookAttributes) InternalWithRef(ref terra.Reference) CodeHookAttributes {
	return CodeHookAttributes{ref: ref}
}

func (ch CodeHookAttributes) InternalTokens() hclwrite.Tokens {
	return ch.ref.InternalTokens()
}

func (ch CodeHookAttributes) MessageVersion() terra.StringValue {
	return terra.ReferenceAsString(ch.ref.Append("message_version"))
}

func (ch CodeHookAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(ch.ref.Append("uri"))
}

type RejectionStatementAttributes struct {
	ref terra.Reference
}

func (rs RejectionStatementAttributes) InternalRef() terra.Reference {
	return rs.ref
}

func (rs RejectionStatementAttributes) InternalWithRef(ref terra.Reference) RejectionStatementAttributes {
	return RejectionStatementAttributes{ref: ref}
}

func (rs RejectionStatementAttributes) InternalTokens() hclwrite.Tokens {
	return rs.ref.InternalTokens()
}

func (rs RejectionStatementAttributes) ResponseCard() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("response_card"))
}

func (rs RejectionStatementAttributes) Message() terra.SetValue[RejectionStatementMessageAttributes] {
	return terra.ReferenceAsSet[RejectionStatementMessageAttributes](rs.ref.Append("message"))
}

type RejectionStatementMessageAttributes struct {
	ref terra.Reference
}

func (m RejectionStatementMessageAttributes) InternalRef() terra.Reference {
	return m.ref
}

func (m RejectionStatementMessageAttributes) InternalWithRef(ref terra.Reference) RejectionStatementMessageAttributes {
	return RejectionStatementMessageAttributes{ref: ref}
}

func (m RejectionStatementMessageAttributes) InternalTokens() hclwrite.Tokens {
	return m.ref.InternalTokens()
}

func (m RejectionStatementMessageAttributes) Content() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("content"))
}

func (m RejectionStatementMessageAttributes) ContentType() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("content_type"))
}

func (m RejectionStatementMessageAttributes) GroupNumber() terra.NumberValue {
	return terra.ReferenceAsNumber(m.ref.Append("group_number"))
}

type SlotAttributes struct {
	ref terra.Reference
}

func (s SlotAttributes) InternalRef() terra.Reference {
	return s.ref
}

func (s SlotAttributes) InternalWithRef(ref terra.Reference) SlotAttributes {
	return SlotAttributes{ref: ref}
}

func (s SlotAttributes) InternalTokens() hclwrite.Tokens {
	return s.ref.InternalTokens()
}

func (s SlotAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("description"))
}

func (s SlotAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("name"))
}

func (s SlotAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("priority"))
}

func (s SlotAttributes) ResponseCard() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("response_card"))
}

func (s SlotAttributes) SampleUtterances() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](s.ref.Append("sample_utterances"))
}

func (s SlotAttributes) SlotConstraint() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("slot_constraint"))
}

func (s SlotAttributes) SlotType() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("slot_type"))
}

func (s SlotAttributes) SlotTypeVersion() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("slot_type_version"))
}

func (s SlotAttributes) ValueElicitationPrompt() terra.ListValue[ValueElicitationPromptAttributes] {
	return terra.ReferenceAsList[ValueElicitationPromptAttributes](s.ref.Append("value_elicitation_prompt"))
}

type ValueElicitationPromptAttributes struct {
	ref terra.Reference
}

func (vep ValueElicitationPromptAttributes) InternalRef() terra.Reference {
	return vep.ref
}

func (vep ValueElicitationPromptAttributes) InternalWithRef(ref terra.Reference) ValueElicitationPromptAttributes {
	return ValueElicitationPromptAttributes{ref: ref}
}

func (vep ValueElicitationPromptAttributes) InternalTokens() hclwrite.Tokens {
	return vep.ref.InternalTokens()
}

func (vep ValueElicitationPromptAttributes) MaxAttempts() terra.NumberValue {
	return terra.ReferenceAsNumber(vep.ref.Append("max_attempts"))
}

func (vep ValueElicitationPromptAttributes) ResponseCard() terra.StringValue {
	return terra.ReferenceAsString(vep.ref.Append("response_card"))
}

func (vep ValueElicitationPromptAttributes) Message() terra.SetValue[ValueElicitationPromptMessageAttributes] {
	return terra.ReferenceAsSet[ValueElicitationPromptMessageAttributes](vep.ref.Append("message"))
}

type ValueElicitationPromptMessageAttributes struct {
	ref terra.Reference
}

func (m ValueElicitationPromptMessageAttributes) InternalRef() terra.Reference {
	return m.ref
}

func (m ValueElicitationPromptMessageAttributes) InternalWithRef(ref terra.Reference) ValueElicitationPromptMessageAttributes {
	return ValueElicitationPromptMessageAttributes{ref: ref}
}

func (m ValueElicitationPromptMessageAttributes) InternalTokens() hclwrite.Tokens {
	return m.ref.InternalTokens()
}

func (m ValueElicitationPromptMessageAttributes) Content() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("content"))
}

func (m ValueElicitationPromptMessageAttributes) ContentType() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("content_type"))
}

func (m ValueElicitationPromptMessageAttributes) GroupNumber() terra.NumberValue {
	return terra.ReferenceAsNumber(m.ref.Append("group_number"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type ConclusionStatementState struct {
	ResponseCard string                            `json:"response_card"`
	Message      []ConclusionStatementMessageState `json:"message"`
}

type ConclusionStatementMessageState struct {
	Content     string  `json:"content"`
	ContentType string  `json:"content_type"`
	GroupNumber float64 `json:"group_number"`
}

type ConfirmationPromptState struct {
	MaxAttempts  float64                          `json:"max_attempts"`
	ResponseCard string                           `json:"response_card"`
	Message      []ConfirmationPromptMessageState `json:"message"`
}

type ConfirmationPromptMessageState struct {
	Content     string  `json:"content"`
	ContentType string  `json:"content_type"`
	GroupNumber float64 `json:"group_number"`
}

type DialogCodeHookState struct {
	MessageVersion string `json:"message_version"`
	Uri            string `json:"uri"`
}

type FollowUpPromptState struct {
	Prompt             []PromptState                           `json:"prompt"`
	RejectionStatement []FollowUpPromptRejectionStatementState `json:"rejection_statement"`
}

type PromptState struct {
	MaxAttempts  float64              `json:"max_attempts"`
	ResponseCard string               `json:"response_card"`
	Message      []PromptMessageState `json:"message"`
}

type PromptMessageState struct {
	Content     string  `json:"content"`
	ContentType string  `json:"content_type"`
	GroupNumber float64 `json:"group_number"`
}

type FollowUpPromptRejectionStatementState struct {
	ResponseCard string                                         `json:"response_card"`
	Message      []FollowUpPromptRejectionStatementMessageState `json:"message"`
}

type FollowUpPromptRejectionStatementMessageState struct {
	Content     string  `json:"content"`
	ContentType string  `json:"content_type"`
	GroupNumber float64 `json:"group_number"`
}

type FulfillmentActivityState struct {
	Type     string          `json:"type"`
	CodeHook []CodeHookState `json:"code_hook"`
}

type CodeHookState struct {
	MessageVersion string `json:"message_version"`
	Uri            string `json:"uri"`
}

type RejectionStatementState struct {
	ResponseCard string                           `json:"response_card"`
	Message      []RejectionStatementMessageState `json:"message"`
}

type RejectionStatementMessageState struct {
	Content     string  `json:"content"`
	ContentType string  `json:"content_type"`
	GroupNumber float64 `json:"group_number"`
}

type SlotState struct {
	Description            string                        `json:"description"`
	Name                   string                        `json:"name"`
	Priority               float64                       `json:"priority"`
	ResponseCard           string                        `json:"response_card"`
	SampleUtterances       []string                      `json:"sample_utterances"`
	SlotConstraint         string                        `json:"slot_constraint"`
	SlotType               string                        `json:"slot_type"`
	SlotTypeVersion        string                        `json:"slot_type_version"`
	ValueElicitationPrompt []ValueElicitationPromptState `json:"value_elicitation_prompt"`
}

type ValueElicitationPromptState struct {
	MaxAttempts  float64                              `json:"max_attempts"`
	ResponseCard string                               `json:"response_card"`
	Message      []ValueElicitationPromptMessageState `json:"message"`
}

type ValueElicitationPromptMessageState struct {
	Content     string  `json:"content"`
	ContentType string  `json:"content_type"`
	GroupNumber float64 `json:"group_number"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
