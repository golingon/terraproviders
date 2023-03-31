// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dialogflowcxpage

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type EntryFulfillment struct {
	// ReturnPartialResponses: bool, optional
	ReturnPartialResponses terra.BoolValue `hcl:"return_partial_responses,attr"`
	// Tag: string, optional
	Tag terra.StringValue `hcl:"tag,attr"`
	// Webhook: string, optional
	Webhook terra.StringValue `hcl:"webhook,attr"`
	// EntryFulfillmentMessages: min=0
	Messages []EntryFulfillmentMessages `hcl:"messages,block" validate:"min=0"`
}

type EntryFulfillmentMessages struct {
	// EntryFulfillmentMessagesText: optional
	Text *EntryFulfillmentMessagesText `hcl:"text,block"`
}

type EntryFulfillmentMessagesText struct {
	// Text: list of string, optional
	Text terra.ListValue[terra.StringValue] `hcl:"text,attr"`
}

type EventHandlers struct {
	// Event: string, optional
	Event terra.StringValue `hcl:"event,attr"`
	// TargetFlow: string, optional
	TargetFlow terra.StringValue `hcl:"target_flow,attr"`
	// TargetPage: string, optional
	TargetPage terra.StringValue `hcl:"target_page,attr"`
	// EventHandlersTriggerFulfillment: optional
	TriggerFulfillment *EventHandlersTriggerFulfillment `hcl:"trigger_fulfillment,block"`
}

type EventHandlersTriggerFulfillment struct {
	// ReturnPartialResponses: bool, optional
	ReturnPartialResponses terra.BoolValue `hcl:"return_partial_responses,attr"`
	// Tag: string, optional
	Tag terra.StringValue `hcl:"tag,attr"`
	// Webhook: string, optional
	Webhook terra.StringValue `hcl:"webhook,attr"`
	// EventHandlersTriggerFulfillmentMessages: min=0
	Messages []EventHandlersTriggerFulfillmentMessages `hcl:"messages,block" validate:"min=0"`
}

type EventHandlersTriggerFulfillmentMessages struct {
	// EventHandlersTriggerFulfillmentMessagesText: optional
	Text *EventHandlersTriggerFulfillmentMessagesText `hcl:"text,block"`
}

type EventHandlersTriggerFulfillmentMessagesText struct {
	// Text: list of string, optional
	Text terra.ListValue[terra.StringValue] `hcl:"text,attr"`
}

type Form struct {
	// Parameters: min=0
	Parameters []Parameters `hcl:"parameters,block" validate:"min=0"`
}

type Parameters struct {
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// EntityType: string, optional
	EntityType terra.StringValue `hcl:"entity_type,attr"`
	// IsList: bool, optional
	IsList terra.BoolValue `hcl:"is_list,attr"`
	// Redact: bool, optional
	Redact terra.BoolValue `hcl:"redact,attr"`
	// Required: bool, optional
	Required terra.BoolValue `hcl:"required,attr"`
	// FillBehavior: optional
	FillBehavior *FillBehavior `hcl:"fill_behavior,block"`
}

type FillBehavior struct {
	// InitialPromptFulfillment: optional
	InitialPromptFulfillment *InitialPromptFulfillment `hcl:"initial_prompt_fulfillment,block"`
}

type InitialPromptFulfillment struct {
	// ReturnPartialResponses: bool, optional
	ReturnPartialResponses terra.BoolValue `hcl:"return_partial_responses,attr"`
	// Tag: string, optional
	Tag terra.StringValue `hcl:"tag,attr"`
	// Webhook: string, optional
	Webhook terra.StringValue `hcl:"webhook,attr"`
	// InitialPromptFulfillmentMessages: min=0
	Messages []InitialPromptFulfillmentMessages `hcl:"messages,block" validate:"min=0"`
}

type InitialPromptFulfillmentMessages struct {
	// InitialPromptFulfillmentMessagesText: optional
	Text *InitialPromptFulfillmentMessagesText `hcl:"text,block"`
}

type InitialPromptFulfillmentMessagesText struct {
	// Text: list of string, optional
	Text terra.ListValue[terra.StringValue] `hcl:"text,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type TransitionRoutes struct {
	// Condition: string, optional
	Condition terra.StringValue `hcl:"condition,attr"`
	// Intent: string, optional
	Intent terra.StringValue `hcl:"intent,attr"`
	// TargetFlow: string, optional
	TargetFlow terra.StringValue `hcl:"target_flow,attr"`
	// TargetPage: string, optional
	TargetPage terra.StringValue `hcl:"target_page,attr"`
	// TransitionRoutesTriggerFulfillment: optional
	TriggerFulfillment *TransitionRoutesTriggerFulfillment `hcl:"trigger_fulfillment,block"`
}

type TransitionRoutesTriggerFulfillment struct {
	// ReturnPartialResponses: bool, optional
	ReturnPartialResponses terra.BoolValue `hcl:"return_partial_responses,attr"`
	// Tag: string, optional
	Tag terra.StringValue `hcl:"tag,attr"`
	// Webhook: string, optional
	Webhook terra.StringValue `hcl:"webhook,attr"`
	// TransitionRoutesTriggerFulfillmentMessages: min=0
	Messages []TransitionRoutesTriggerFulfillmentMessages `hcl:"messages,block" validate:"min=0"`
}

type TransitionRoutesTriggerFulfillmentMessages struct {
	// TransitionRoutesTriggerFulfillmentMessagesText: optional
	Text *TransitionRoutesTriggerFulfillmentMessagesText `hcl:"text,block"`
}

type TransitionRoutesTriggerFulfillmentMessagesText struct {
	// Text: list of string, optional
	Text terra.ListValue[terra.StringValue] `hcl:"text,attr"`
}

type EntryFulfillmentAttributes struct {
	ref terra.Reference
}

func (ef EntryFulfillmentAttributes) InternalRef() terra.Reference {
	return ef.ref
}

func (ef EntryFulfillmentAttributes) InternalWithRef(ref terra.Reference) EntryFulfillmentAttributes {
	return EntryFulfillmentAttributes{ref: ref}
}

func (ef EntryFulfillmentAttributes) InternalTokens() hclwrite.Tokens {
	return ef.ref.InternalTokens()
}

func (ef EntryFulfillmentAttributes) ReturnPartialResponses() terra.BoolValue {
	return terra.ReferenceBool(ef.ref.Append("return_partial_responses"))
}

func (ef EntryFulfillmentAttributes) Tag() terra.StringValue {
	return terra.ReferenceString(ef.ref.Append("tag"))
}

func (ef EntryFulfillmentAttributes) Webhook() terra.StringValue {
	return terra.ReferenceString(ef.ref.Append("webhook"))
}

func (ef EntryFulfillmentAttributes) Messages() terra.ListValue[EntryFulfillmentMessagesAttributes] {
	return terra.ReferenceList[EntryFulfillmentMessagesAttributes](ef.ref.Append("messages"))
}

type EntryFulfillmentMessagesAttributes struct {
	ref terra.Reference
}

func (m EntryFulfillmentMessagesAttributes) InternalRef() terra.Reference {
	return m.ref
}

func (m EntryFulfillmentMessagesAttributes) InternalWithRef(ref terra.Reference) EntryFulfillmentMessagesAttributes {
	return EntryFulfillmentMessagesAttributes{ref: ref}
}

func (m EntryFulfillmentMessagesAttributes) InternalTokens() hclwrite.Tokens {
	return m.ref.InternalTokens()
}

func (m EntryFulfillmentMessagesAttributes) Text() terra.ListValue[EntryFulfillmentMessagesTextAttributes] {
	return terra.ReferenceList[EntryFulfillmentMessagesTextAttributes](m.ref.Append("text"))
}

type EntryFulfillmentMessagesTextAttributes struct {
	ref terra.Reference
}

func (t EntryFulfillmentMessagesTextAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t EntryFulfillmentMessagesTextAttributes) InternalWithRef(ref terra.Reference) EntryFulfillmentMessagesTextAttributes {
	return EntryFulfillmentMessagesTextAttributes{ref: ref}
}

func (t EntryFulfillmentMessagesTextAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t EntryFulfillmentMessagesTextAttributes) AllowPlaybackInterruption() terra.BoolValue {
	return terra.ReferenceBool(t.ref.Append("allow_playback_interruption"))
}

func (t EntryFulfillmentMessagesTextAttributes) Text() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](t.ref.Append("text"))
}

type EventHandlersAttributes struct {
	ref terra.Reference
}

func (eh EventHandlersAttributes) InternalRef() terra.Reference {
	return eh.ref
}

func (eh EventHandlersAttributes) InternalWithRef(ref terra.Reference) EventHandlersAttributes {
	return EventHandlersAttributes{ref: ref}
}

func (eh EventHandlersAttributes) InternalTokens() hclwrite.Tokens {
	return eh.ref.InternalTokens()
}

func (eh EventHandlersAttributes) Event() terra.StringValue {
	return terra.ReferenceString(eh.ref.Append("event"))
}

func (eh EventHandlersAttributes) Name() terra.StringValue {
	return terra.ReferenceString(eh.ref.Append("name"))
}

func (eh EventHandlersAttributes) TargetFlow() terra.StringValue {
	return terra.ReferenceString(eh.ref.Append("target_flow"))
}

func (eh EventHandlersAttributes) TargetPage() terra.StringValue {
	return terra.ReferenceString(eh.ref.Append("target_page"))
}

func (eh EventHandlersAttributes) TriggerFulfillment() terra.ListValue[EventHandlersTriggerFulfillmentAttributes] {
	return terra.ReferenceList[EventHandlersTriggerFulfillmentAttributes](eh.ref.Append("trigger_fulfillment"))
}

type EventHandlersTriggerFulfillmentAttributes struct {
	ref terra.Reference
}

func (tf EventHandlersTriggerFulfillmentAttributes) InternalRef() terra.Reference {
	return tf.ref
}

func (tf EventHandlersTriggerFulfillmentAttributes) InternalWithRef(ref terra.Reference) EventHandlersTriggerFulfillmentAttributes {
	return EventHandlersTriggerFulfillmentAttributes{ref: ref}
}

func (tf EventHandlersTriggerFulfillmentAttributes) InternalTokens() hclwrite.Tokens {
	return tf.ref.InternalTokens()
}

func (tf EventHandlersTriggerFulfillmentAttributes) ReturnPartialResponses() terra.BoolValue {
	return terra.ReferenceBool(tf.ref.Append("return_partial_responses"))
}

func (tf EventHandlersTriggerFulfillmentAttributes) Tag() terra.StringValue {
	return terra.ReferenceString(tf.ref.Append("tag"))
}

func (tf EventHandlersTriggerFulfillmentAttributes) Webhook() terra.StringValue {
	return terra.ReferenceString(tf.ref.Append("webhook"))
}

func (tf EventHandlersTriggerFulfillmentAttributes) Messages() terra.ListValue[EventHandlersTriggerFulfillmentMessagesAttributes] {
	return terra.ReferenceList[EventHandlersTriggerFulfillmentMessagesAttributes](tf.ref.Append("messages"))
}

type EventHandlersTriggerFulfillmentMessagesAttributes struct {
	ref terra.Reference
}

func (m EventHandlersTriggerFulfillmentMessagesAttributes) InternalRef() terra.Reference {
	return m.ref
}

func (m EventHandlersTriggerFulfillmentMessagesAttributes) InternalWithRef(ref terra.Reference) EventHandlersTriggerFulfillmentMessagesAttributes {
	return EventHandlersTriggerFulfillmentMessagesAttributes{ref: ref}
}

func (m EventHandlersTriggerFulfillmentMessagesAttributes) InternalTokens() hclwrite.Tokens {
	return m.ref.InternalTokens()
}

func (m EventHandlersTriggerFulfillmentMessagesAttributes) Text() terra.ListValue[EventHandlersTriggerFulfillmentMessagesTextAttributes] {
	return terra.ReferenceList[EventHandlersTriggerFulfillmentMessagesTextAttributes](m.ref.Append("text"))
}

type EventHandlersTriggerFulfillmentMessagesTextAttributes struct {
	ref terra.Reference
}

func (t EventHandlersTriggerFulfillmentMessagesTextAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t EventHandlersTriggerFulfillmentMessagesTextAttributes) InternalWithRef(ref terra.Reference) EventHandlersTriggerFulfillmentMessagesTextAttributes {
	return EventHandlersTriggerFulfillmentMessagesTextAttributes{ref: ref}
}

func (t EventHandlersTriggerFulfillmentMessagesTextAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t EventHandlersTriggerFulfillmentMessagesTextAttributes) AllowPlaybackInterruption() terra.BoolValue {
	return terra.ReferenceBool(t.ref.Append("allow_playback_interruption"))
}

func (t EventHandlersTriggerFulfillmentMessagesTextAttributes) Text() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](t.ref.Append("text"))
}

type FormAttributes struct {
	ref terra.Reference
}

func (f FormAttributes) InternalRef() terra.Reference {
	return f.ref
}

func (f FormAttributes) InternalWithRef(ref terra.Reference) FormAttributes {
	return FormAttributes{ref: ref}
}

func (f FormAttributes) InternalTokens() hclwrite.Tokens {
	return f.ref.InternalTokens()
}

func (f FormAttributes) Parameters() terra.ListValue[ParametersAttributes] {
	return terra.ReferenceList[ParametersAttributes](f.ref.Append("parameters"))
}

type ParametersAttributes struct {
	ref terra.Reference
}

func (p ParametersAttributes) InternalRef() terra.Reference {
	return p.ref
}

func (p ParametersAttributes) InternalWithRef(ref terra.Reference) ParametersAttributes {
	return ParametersAttributes{ref: ref}
}

func (p ParametersAttributes) InternalTokens() hclwrite.Tokens {
	return p.ref.InternalTokens()
}

func (p ParametersAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceString(p.ref.Append("display_name"))
}

func (p ParametersAttributes) EntityType() terra.StringValue {
	return terra.ReferenceString(p.ref.Append("entity_type"))
}

func (p ParametersAttributes) IsList() terra.BoolValue {
	return terra.ReferenceBool(p.ref.Append("is_list"))
}

func (p ParametersAttributes) Redact() terra.BoolValue {
	return terra.ReferenceBool(p.ref.Append("redact"))
}

func (p ParametersAttributes) Required() terra.BoolValue {
	return terra.ReferenceBool(p.ref.Append("required"))
}

func (p ParametersAttributes) FillBehavior() terra.ListValue[FillBehaviorAttributes] {
	return terra.ReferenceList[FillBehaviorAttributes](p.ref.Append("fill_behavior"))
}

type FillBehaviorAttributes struct {
	ref terra.Reference
}

func (fb FillBehaviorAttributes) InternalRef() terra.Reference {
	return fb.ref
}

func (fb FillBehaviorAttributes) InternalWithRef(ref terra.Reference) FillBehaviorAttributes {
	return FillBehaviorAttributes{ref: ref}
}

func (fb FillBehaviorAttributes) InternalTokens() hclwrite.Tokens {
	return fb.ref.InternalTokens()
}

func (fb FillBehaviorAttributes) InitialPromptFulfillment() terra.ListValue[InitialPromptFulfillmentAttributes] {
	return terra.ReferenceList[InitialPromptFulfillmentAttributes](fb.ref.Append("initial_prompt_fulfillment"))
}

type InitialPromptFulfillmentAttributes struct {
	ref terra.Reference
}

func (ipf InitialPromptFulfillmentAttributes) InternalRef() terra.Reference {
	return ipf.ref
}

func (ipf InitialPromptFulfillmentAttributes) InternalWithRef(ref terra.Reference) InitialPromptFulfillmentAttributes {
	return InitialPromptFulfillmentAttributes{ref: ref}
}

func (ipf InitialPromptFulfillmentAttributes) InternalTokens() hclwrite.Tokens {
	return ipf.ref.InternalTokens()
}

func (ipf InitialPromptFulfillmentAttributes) ReturnPartialResponses() terra.BoolValue {
	return terra.ReferenceBool(ipf.ref.Append("return_partial_responses"))
}

func (ipf InitialPromptFulfillmentAttributes) Tag() terra.StringValue {
	return terra.ReferenceString(ipf.ref.Append("tag"))
}

func (ipf InitialPromptFulfillmentAttributes) Webhook() terra.StringValue {
	return terra.ReferenceString(ipf.ref.Append("webhook"))
}

func (ipf InitialPromptFulfillmentAttributes) Messages() terra.ListValue[InitialPromptFulfillmentMessagesAttributes] {
	return terra.ReferenceList[InitialPromptFulfillmentMessagesAttributes](ipf.ref.Append("messages"))
}

type InitialPromptFulfillmentMessagesAttributes struct {
	ref terra.Reference
}

func (m InitialPromptFulfillmentMessagesAttributes) InternalRef() terra.Reference {
	return m.ref
}

func (m InitialPromptFulfillmentMessagesAttributes) InternalWithRef(ref terra.Reference) InitialPromptFulfillmentMessagesAttributes {
	return InitialPromptFulfillmentMessagesAttributes{ref: ref}
}

func (m InitialPromptFulfillmentMessagesAttributes) InternalTokens() hclwrite.Tokens {
	return m.ref.InternalTokens()
}

func (m InitialPromptFulfillmentMessagesAttributes) Text() terra.ListValue[InitialPromptFulfillmentMessagesTextAttributes] {
	return terra.ReferenceList[InitialPromptFulfillmentMessagesTextAttributes](m.ref.Append("text"))
}

type InitialPromptFulfillmentMessagesTextAttributes struct {
	ref terra.Reference
}

func (t InitialPromptFulfillmentMessagesTextAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t InitialPromptFulfillmentMessagesTextAttributes) InternalWithRef(ref terra.Reference) InitialPromptFulfillmentMessagesTextAttributes {
	return InitialPromptFulfillmentMessagesTextAttributes{ref: ref}
}

func (t InitialPromptFulfillmentMessagesTextAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t InitialPromptFulfillmentMessagesTextAttributes) AllowPlaybackInterruption() terra.BoolValue {
	return terra.ReferenceBool(t.ref.Append("allow_playback_interruption"))
}

func (t InitialPromptFulfillmentMessagesTextAttributes) Text() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](t.ref.Append("text"))
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
	return terra.ReferenceString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("update"))
}

type TransitionRoutesAttributes struct {
	ref terra.Reference
}

func (tr TransitionRoutesAttributes) InternalRef() terra.Reference {
	return tr.ref
}

func (tr TransitionRoutesAttributes) InternalWithRef(ref terra.Reference) TransitionRoutesAttributes {
	return TransitionRoutesAttributes{ref: ref}
}

func (tr TransitionRoutesAttributes) InternalTokens() hclwrite.Tokens {
	return tr.ref.InternalTokens()
}

func (tr TransitionRoutesAttributes) Condition() terra.StringValue {
	return terra.ReferenceString(tr.ref.Append("condition"))
}

func (tr TransitionRoutesAttributes) Intent() terra.StringValue {
	return terra.ReferenceString(tr.ref.Append("intent"))
}

func (tr TransitionRoutesAttributes) Name() terra.StringValue {
	return terra.ReferenceString(tr.ref.Append("name"))
}

func (tr TransitionRoutesAttributes) TargetFlow() terra.StringValue {
	return terra.ReferenceString(tr.ref.Append("target_flow"))
}

func (tr TransitionRoutesAttributes) TargetPage() terra.StringValue {
	return terra.ReferenceString(tr.ref.Append("target_page"))
}

func (tr TransitionRoutesAttributes) TriggerFulfillment() terra.ListValue[TransitionRoutesTriggerFulfillmentAttributes] {
	return terra.ReferenceList[TransitionRoutesTriggerFulfillmentAttributes](tr.ref.Append("trigger_fulfillment"))
}

type TransitionRoutesTriggerFulfillmentAttributes struct {
	ref terra.Reference
}

func (tf TransitionRoutesTriggerFulfillmentAttributes) InternalRef() terra.Reference {
	return tf.ref
}

func (tf TransitionRoutesTriggerFulfillmentAttributes) InternalWithRef(ref terra.Reference) TransitionRoutesTriggerFulfillmentAttributes {
	return TransitionRoutesTriggerFulfillmentAttributes{ref: ref}
}

func (tf TransitionRoutesTriggerFulfillmentAttributes) InternalTokens() hclwrite.Tokens {
	return tf.ref.InternalTokens()
}

func (tf TransitionRoutesTriggerFulfillmentAttributes) ReturnPartialResponses() terra.BoolValue {
	return terra.ReferenceBool(tf.ref.Append("return_partial_responses"))
}

func (tf TransitionRoutesTriggerFulfillmentAttributes) Tag() terra.StringValue {
	return terra.ReferenceString(tf.ref.Append("tag"))
}

func (tf TransitionRoutesTriggerFulfillmentAttributes) Webhook() terra.StringValue {
	return terra.ReferenceString(tf.ref.Append("webhook"))
}

func (tf TransitionRoutesTriggerFulfillmentAttributes) Messages() terra.ListValue[TransitionRoutesTriggerFulfillmentMessagesAttributes] {
	return terra.ReferenceList[TransitionRoutesTriggerFulfillmentMessagesAttributes](tf.ref.Append("messages"))
}

type TransitionRoutesTriggerFulfillmentMessagesAttributes struct {
	ref terra.Reference
}

func (m TransitionRoutesTriggerFulfillmentMessagesAttributes) InternalRef() terra.Reference {
	return m.ref
}

func (m TransitionRoutesTriggerFulfillmentMessagesAttributes) InternalWithRef(ref terra.Reference) TransitionRoutesTriggerFulfillmentMessagesAttributes {
	return TransitionRoutesTriggerFulfillmentMessagesAttributes{ref: ref}
}

func (m TransitionRoutesTriggerFulfillmentMessagesAttributes) InternalTokens() hclwrite.Tokens {
	return m.ref.InternalTokens()
}

func (m TransitionRoutesTriggerFulfillmentMessagesAttributes) Text() terra.ListValue[TransitionRoutesTriggerFulfillmentMessagesTextAttributes] {
	return terra.ReferenceList[TransitionRoutesTriggerFulfillmentMessagesTextAttributes](m.ref.Append("text"))
}

type TransitionRoutesTriggerFulfillmentMessagesTextAttributes struct {
	ref terra.Reference
}

func (t TransitionRoutesTriggerFulfillmentMessagesTextAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TransitionRoutesTriggerFulfillmentMessagesTextAttributes) InternalWithRef(ref terra.Reference) TransitionRoutesTriggerFulfillmentMessagesTextAttributes {
	return TransitionRoutesTriggerFulfillmentMessagesTextAttributes{ref: ref}
}

func (t TransitionRoutesTriggerFulfillmentMessagesTextAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TransitionRoutesTriggerFulfillmentMessagesTextAttributes) AllowPlaybackInterruption() terra.BoolValue {
	return terra.ReferenceBool(t.ref.Append("allow_playback_interruption"))
}

func (t TransitionRoutesTriggerFulfillmentMessagesTextAttributes) Text() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](t.ref.Append("text"))
}

type EntryFulfillmentState struct {
	ReturnPartialResponses bool                            `json:"return_partial_responses"`
	Tag                    string                          `json:"tag"`
	Webhook                string                          `json:"webhook"`
	Messages               []EntryFulfillmentMessagesState `json:"messages"`
}

type EntryFulfillmentMessagesState struct {
	Text []EntryFulfillmentMessagesTextState `json:"text"`
}

type EntryFulfillmentMessagesTextState struct {
	AllowPlaybackInterruption bool     `json:"allow_playback_interruption"`
	Text                      []string `json:"text"`
}

type EventHandlersState struct {
	Event              string                                 `json:"event"`
	Name               string                                 `json:"name"`
	TargetFlow         string                                 `json:"target_flow"`
	TargetPage         string                                 `json:"target_page"`
	TriggerFulfillment []EventHandlersTriggerFulfillmentState `json:"trigger_fulfillment"`
}

type EventHandlersTriggerFulfillmentState struct {
	ReturnPartialResponses bool                                           `json:"return_partial_responses"`
	Tag                    string                                         `json:"tag"`
	Webhook                string                                         `json:"webhook"`
	Messages               []EventHandlersTriggerFulfillmentMessagesState `json:"messages"`
}

type EventHandlersTriggerFulfillmentMessagesState struct {
	Text []EventHandlersTriggerFulfillmentMessagesTextState `json:"text"`
}

type EventHandlersTriggerFulfillmentMessagesTextState struct {
	AllowPlaybackInterruption bool     `json:"allow_playback_interruption"`
	Text                      []string `json:"text"`
}

type FormState struct {
	Parameters []ParametersState `json:"parameters"`
}

type ParametersState struct {
	DisplayName  string              `json:"display_name"`
	EntityType   string              `json:"entity_type"`
	IsList       bool                `json:"is_list"`
	Redact       bool                `json:"redact"`
	Required     bool                `json:"required"`
	FillBehavior []FillBehaviorState `json:"fill_behavior"`
}

type FillBehaviorState struct {
	InitialPromptFulfillment []InitialPromptFulfillmentState `json:"initial_prompt_fulfillment"`
}

type InitialPromptFulfillmentState struct {
	ReturnPartialResponses bool                                    `json:"return_partial_responses"`
	Tag                    string                                  `json:"tag"`
	Webhook                string                                  `json:"webhook"`
	Messages               []InitialPromptFulfillmentMessagesState `json:"messages"`
}

type InitialPromptFulfillmentMessagesState struct {
	Text []InitialPromptFulfillmentMessagesTextState `json:"text"`
}

type InitialPromptFulfillmentMessagesTextState struct {
	AllowPlaybackInterruption bool     `json:"allow_playback_interruption"`
	Text                      []string `json:"text"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

type TransitionRoutesState struct {
	Condition          string                                    `json:"condition"`
	Intent             string                                    `json:"intent"`
	Name               string                                    `json:"name"`
	TargetFlow         string                                    `json:"target_flow"`
	TargetPage         string                                    `json:"target_page"`
	TriggerFulfillment []TransitionRoutesTriggerFulfillmentState `json:"trigger_fulfillment"`
}

type TransitionRoutesTriggerFulfillmentState struct {
	ReturnPartialResponses bool                                              `json:"return_partial_responses"`
	Tag                    string                                            `json:"tag"`
	Webhook                string                                            `json:"webhook"`
	Messages               []TransitionRoutesTriggerFulfillmentMessagesState `json:"messages"`
}

type TransitionRoutesTriggerFulfillmentMessagesState struct {
	Text []TransitionRoutesTriggerFulfillmentMessagesTextState `json:"text"`
}

type TransitionRoutesTriggerFulfillmentMessagesTextState struct {
	AllowPlaybackInterruption bool     `json:"allow_playback_interruption"`
	Text                      []string `json:"text"`
}
