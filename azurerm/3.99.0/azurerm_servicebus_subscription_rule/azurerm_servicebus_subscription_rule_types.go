// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_servicebus_subscription_rule

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type CorrelationFilter struct {
	// ContentType: string, optional
	ContentType terra.StringValue `hcl:"content_type,attr"`
	// CorrelationId: string, optional
	CorrelationId terra.StringValue `hcl:"correlation_id,attr"`
	// Label: string, optional
	Label terra.StringValue `hcl:"label,attr"`
	// MessageId: string, optional
	MessageId terra.StringValue `hcl:"message_id,attr"`
	// Properties: map of string, optional
	Properties terra.MapValue[terra.StringValue] `hcl:"properties,attr"`
	// ReplyTo: string, optional
	ReplyTo terra.StringValue `hcl:"reply_to,attr"`
	// ReplyToSessionId: string, optional
	ReplyToSessionId terra.StringValue `hcl:"reply_to_session_id,attr"`
	// SessionId: string, optional
	SessionId terra.StringValue `hcl:"session_id,attr"`
	// To: string, optional
	To terra.StringValue `hcl:"to,attr"`
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

type CorrelationFilterAttributes struct {
	ref terra.Reference
}

func (cf CorrelationFilterAttributes) InternalRef() (terra.Reference, error) {
	return cf.ref, nil
}

func (cf CorrelationFilterAttributes) InternalWithRef(ref terra.Reference) CorrelationFilterAttributes {
	return CorrelationFilterAttributes{ref: ref}
}

func (cf CorrelationFilterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cf.ref.InternalTokens()
}

func (cf CorrelationFilterAttributes) ContentType() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("content_type"))
}

func (cf CorrelationFilterAttributes) CorrelationId() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("correlation_id"))
}

func (cf CorrelationFilterAttributes) Label() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("label"))
}

func (cf CorrelationFilterAttributes) MessageId() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("message_id"))
}

func (cf CorrelationFilterAttributes) Properties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cf.ref.Append("properties"))
}

func (cf CorrelationFilterAttributes) ReplyTo() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("reply_to"))
}

func (cf CorrelationFilterAttributes) ReplyToSessionId() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("reply_to_session_id"))
}

func (cf CorrelationFilterAttributes) SessionId() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("session_id"))
}

func (cf CorrelationFilterAttributes) To() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("to"))
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

type CorrelationFilterState struct {
	ContentType      string            `json:"content_type"`
	CorrelationId    string            `json:"correlation_id"`
	Label            string            `json:"label"`
	MessageId        string            `json:"message_id"`
	Properties       map[string]string `json:"properties"`
	ReplyTo          string            `json:"reply_to"`
	ReplyToSessionId string            `json:"reply_to_session_id"`
	SessionId        string            `json:"session_id"`
	To               string            `json:"to"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
