// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_web_pubsub_hub

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type EventHandler struct {
	// SystemEvents: set of string, optional
	SystemEvents terra.SetValue[terra.StringValue] `hcl:"system_events,attr"`
	// UrlTemplate: string, required
	UrlTemplate terra.StringValue `hcl:"url_template,attr" validate:"required"`
	// UserEventPattern: string, optional
	UserEventPattern terra.StringValue `hcl:"user_event_pattern,attr"`
	// EventHandlerAuth: optional
	Auth *EventHandlerAuth `hcl:"auth,block"`
}

type EventHandlerAuth struct {
	// ManagedIdentityId: string, required
	ManagedIdentityId terra.StringValue `hcl:"managed_identity_id,attr" validate:"required"`
}

type EventListener struct {
	// EventhubName: string, required
	EventhubName terra.StringValue `hcl:"eventhub_name,attr" validate:"required"`
	// EventhubNamespaceName: string, required
	EventhubNamespaceName terra.StringValue `hcl:"eventhub_namespace_name,attr" validate:"required"`
	// SystemEventNameFilter: list of string, optional
	SystemEventNameFilter terra.ListValue[terra.StringValue] `hcl:"system_event_name_filter,attr"`
	// UserEventNameFilter: list of string, optional
	UserEventNameFilter terra.ListValue[terra.StringValue] `hcl:"user_event_name_filter,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type EventHandlerAttributes struct {
	ref terra.Reference
}

func (eh EventHandlerAttributes) InternalRef() (terra.Reference, error) {
	return eh.ref, nil
}

func (eh EventHandlerAttributes) InternalWithRef(ref terra.Reference) EventHandlerAttributes {
	return EventHandlerAttributes{ref: ref}
}

func (eh EventHandlerAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return eh.ref.InternalTokens()
}

func (eh EventHandlerAttributes) SystemEvents() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](eh.ref.Append("system_events"))
}

func (eh EventHandlerAttributes) UrlTemplate() terra.StringValue {
	return terra.ReferenceAsString(eh.ref.Append("url_template"))
}

func (eh EventHandlerAttributes) UserEventPattern() terra.StringValue {
	return terra.ReferenceAsString(eh.ref.Append("user_event_pattern"))
}

func (eh EventHandlerAttributes) Auth() terra.ListValue[EventHandlerAuthAttributes] {
	return terra.ReferenceAsList[EventHandlerAuthAttributes](eh.ref.Append("auth"))
}

type EventHandlerAuthAttributes struct {
	ref terra.Reference
}

func (a EventHandlerAuthAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a EventHandlerAuthAttributes) InternalWithRef(ref terra.Reference) EventHandlerAuthAttributes {
	return EventHandlerAuthAttributes{ref: ref}
}

func (a EventHandlerAuthAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a EventHandlerAuthAttributes) ManagedIdentityId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("managed_identity_id"))
}

type EventListenerAttributes struct {
	ref terra.Reference
}

func (el EventListenerAttributes) InternalRef() (terra.Reference, error) {
	return el.ref, nil
}

func (el EventListenerAttributes) InternalWithRef(ref terra.Reference) EventListenerAttributes {
	return EventListenerAttributes{ref: ref}
}

func (el EventListenerAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return el.ref.InternalTokens()
}

func (el EventListenerAttributes) EventhubName() terra.StringValue {
	return terra.ReferenceAsString(el.ref.Append("eventhub_name"))
}

func (el EventListenerAttributes) EventhubNamespaceName() terra.StringValue {
	return terra.ReferenceAsString(el.ref.Append("eventhub_namespace_name"))
}

func (el EventListenerAttributes) SystemEventNameFilter() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](el.ref.Append("system_event_name_filter"))
}

func (el EventListenerAttributes) UserEventNameFilter() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](el.ref.Append("user_event_name_filter"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type EventHandlerState struct {
	SystemEvents     []string                `json:"system_events"`
	UrlTemplate      string                  `json:"url_template"`
	UserEventPattern string                  `json:"user_event_pattern"`
	Auth             []EventHandlerAuthState `json:"auth"`
}

type EventHandlerAuthState struct {
	ManagedIdentityId string `json:"managed_identity_id"`
}

type EventListenerState struct {
	EventhubName          string   `json:"eventhub_name"`
	EventhubNamespaceName string   `json:"eventhub_namespace_name"`
	SystemEventNameFilter []string `json:"system_event_name_filter"`
	UserEventNameFilter   []string `json:"user_event_name_filter"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
