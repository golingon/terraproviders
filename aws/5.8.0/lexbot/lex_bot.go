// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package lexbot

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AbortStatement struct {
	// ResponseCard: string, optional
	ResponseCard terra.StringValue `hcl:"response_card,attr"`
	// AbortStatementMessage: min=1,max=15
	Message []AbortStatementMessage `hcl:"message,block" validate:"min=1,max=15"`
}

type AbortStatementMessage struct {
	// Content: string, required
	Content terra.StringValue `hcl:"content,attr" validate:"required"`
	// ContentType: string, required
	ContentType terra.StringValue `hcl:"content_type,attr" validate:"required"`
	// GroupNumber: number, optional
	GroupNumber terra.NumberValue `hcl:"group_number,attr"`
}

type ClarificationPrompt struct {
	// MaxAttempts: number, required
	MaxAttempts terra.NumberValue `hcl:"max_attempts,attr" validate:"required"`
	// ResponseCard: string, optional
	ResponseCard terra.StringValue `hcl:"response_card,attr"`
	// ClarificationPromptMessage: min=1,max=15
	Message []ClarificationPromptMessage `hcl:"message,block" validate:"min=1,max=15"`
}

type ClarificationPromptMessage struct {
	// Content: string, required
	Content terra.StringValue `hcl:"content,attr" validate:"required"`
	// ContentType: string, required
	ContentType terra.StringValue `hcl:"content_type,attr" validate:"required"`
	// GroupNumber: number, optional
	GroupNumber terra.NumberValue `hcl:"group_number,attr"`
}

type Intent struct {
	// IntentName: string, required
	IntentName terra.StringValue `hcl:"intent_name,attr" validate:"required"`
	// IntentVersion: string, required
	IntentVersion terra.StringValue `hcl:"intent_version,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type AbortStatementAttributes struct {
	ref terra.Reference
}

func (as AbortStatementAttributes) InternalRef() (terra.Reference, error) {
	return as.ref, nil
}

func (as AbortStatementAttributes) InternalWithRef(ref terra.Reference) AbortStatementAttributes {
	return AbortStatementAttributes{ref: ref}
}

func (as AbortStatementAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return as.ref.InternalTokens()
}

func (as AbortStatementAttributes) ResponseCard() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("response_card"))
}

func (as AbortStatementAttributes) Message() terra.SetValue[AbortStatementMessageAttributes] {
	return terra.ReferenceAsSet[AbortStatementMessageAttributes](as.ref.Append("message"))
}

type AbortStatementMessageAttributes struct {
	ref terra.Reference
}

func (m AbortStatementMessageAttributes) InternalRef() (terra.Reference, error) {
	return m.ref, nil
}

func (m AbortStatementMessageAttributes) InternalWithRef(ref terra.Reference) AbortStatementMessageAttributes {
	return AbortStatementMessageAttributes{ref: ref}
}

func (m AbortStatementMessageAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return m.ref.InternalTokens()
}

func (m AbortStatementMessageAttributes) Content() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("content"))
}

func (m AbortStatementMessageAttributes) ContentType() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("content_type"))
}

func (m AbortStatementMessageAttributes) GroupNumber() terra.NumberValue {
	return terra.ReferenceAsNumber(m.ref.Append("group_number"))
}

type ClarificationPromptAttributes struct {
	ref terra.Reference
}

func (cp ClarificationPromptAttributes) InternalRef() (terra.Reference, error) {
	return cp.ref, nil
}

func (cp ClarificationPromptAttributes) InternalWithRef(ref terra.Reference) ClarificationPromptAttributes {
	return ClarificationPromptAttributes{ref: ref}
}

func (cp ClarificationPromptAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cp.ref.InternalTokens()
}

func (cp ClarificationPromptAttributes) MaxAttempts() terra.NumberValue {
	return terra.ReferenceAsNumber(cp.ref.Append("max_attempts"))
}

func (cp ClarificationPromptAttributes) ResponseCard() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("response_card"))
}

func (cp ClarificationPromptAttributes) Message() terra.SetValue[ClarificationPromptMessageAttributes] {
	return terra.ReferenceAsSet[ClarificationPromptMessageAttributes](cp.ref.Append("message"))
}

type ClarificationPromptMessageAttributes struct {
	ref terra.Reference
}

func (m ClarificationPromptMessageAttributes) InternalRef() (terra.Reference, error) {
	return m.ref, nil
}

func (m ClarificationPromptMessageAttributes) InternalWithRef(ref terra.Reference) ClarificationPromptMessageAttributes {
	return ClarificationPromptMessageAttributes{ref: ref}
}

func (m ClarificationPromptMessageAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return m.ref.InternalTokens()
}

func (m ClarificationPromptMessageAttributes) Content() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("content"))
}

func (m ClarificationPromptMessageAttributes) ContentType() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("content_type"))
}

func (m ClarificationPromptMessageAttributes) GroupNumber() terra.NumberValue {
	return terra.ReferenceAsNumber(m.ref.Append("group_number"))
}

type IntentAttributes struct {
	ref terra.Reference
}

func (i IntentAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i IntentAttributes) InternalWithRef(ref terra.Reference) IntentAttributes {
	return IntentAttributes{ref: ref}
}

func (i IntentAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i IntentAttributes) IntentName() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("intent_name"))
}

func (i IntentAttributes) IntentVersion() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("intent_version"))
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

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type AbortStatementState struct {
	ResponseCard string                       `json:"response_card"`
	Message      []AbortStatementMessageState `json:"message"`
}

type AbortStatementMessageState struct {
	Content     string  `json:"content"`
	ContentType string  `json:"content_type"`
	GroupNumber float64 `json:"group_number"`
}

type ClarificationPromptState struct {
	MaxAttempts  float64                           `json:"max_attempts"`
	ResponseCard string                            `json:"response_card"`
	Message      []ClarificationPromptMessageState `json:"message"`
}

type ClarificationPromptMessageState struct {
	Content     string  `json:"content"`
	ContentType string  `json:"content_type"`
	GroupNumber float64 `json:"group_number"`
}

type IntentState struct {
	IntentName    string `json:"intent_name"`
	IntentVersion string `json:"intent_version"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
