// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package invitation

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Message struct {
	// AdditionalRecipients: list of string, optional
	AdditionalRecipients terra.ListValue[terra.StringValue] `hcl:"additional_recipients,attr"`
	// Body: string, optional
	Body terra.StringValue `hcl:"body,attr"`
	// Language: string, optional
	Language terra.StringValue `hcl:"language,attr"`
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

type MessageAttributes struct {
	ref terra.Reference
}

func (m MessageAttributes) InternalRef() (terra.Reference, error) {
	return m.ref, nil
}

func (m MessageAttributes) InternalWithRef(ref terra.Reference) MessageAttributes {
	return MessageAttributes{ref: ref}
}

func (m MessageAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return m.ref.InternalTokens()
}

func (m MessageAttributes) AdditionalRecipients() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](m.ref.Append("additional_recipients"))
}

func (m MessageAttributes) Body() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("body"))
}

func (m MessageAttributes) Language() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("language"))
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

type MessageState struct {
	AdditionalRecipients []string `json:"additional_recipients"`
	Body                 string   `json:"body"`
	Language             string   `json:"language"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}