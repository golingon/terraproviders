// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datamonitorscheduledqueryrulesalert

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Action struct{}

type Trigger struct {
	// MetricTrigger: min=0
	MetricTrigger []MetricTrigger `hcl:"metric_trigger,block" validate:"min=0"`
}

type MetricTrigger struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type ActionAttributes struct {
	ref terra.Reference
}

func (a ActionAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a ActionAttributes) InternalWithRef(ref terra.Reference) ActionAttributes {
	return ActionAttributes{ref: ref}
}

func (a ActionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a ActionAttributes) ActionGroup() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](a.ref.Append("action_group"))
}

func (a ActionAttributes) CustomWebhookPayload() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("custom_webhook_payload"))
}

func (a ActionAttributes) EmailSubject() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("email_subject"))
}

type TriggerAttributes struct {
	ref terra.Reference
}

func (t TriggerAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TriggerAttributes) InternalWithRef(ref terra.Reference) TriggerAttributes {
	return TriggerAttributes{ref: ref}
}

func (t TriggerAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TriggerAttributes) Operator() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("operator"))
}

func (t TriggerAttributes) Threshold() terra.NumberValue {
	return terra.ReferenceAsNumber(t.ref.Append("threshold"))
}

func (t TriggerAttributes) MetricTrigger() terra.SetValue[MetricTriggerAttributes] {
	return terra.ReferenceAsSet[MetricTriggerAttributes](t.ref.Append("metric_trigger"))
}

type MetricTriggerAttributes struct {
	ref terra.Reference
}

func (mt MetricTriggerAttributes) InternalRef() (terra.Reference, error) {
	return mt.ref, nil
}

func (mt MetricTriggerAttributes) InternalWithRef(ref terra.Reference) MetricTriggerAttributes {
	return MetricTriggerAttributes{ref: ref}
}

func (mt MetricTriggerAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mt.ref.InternalTokens()
}

func (mt MetricTriggerAttributes) MetricColumn() terra.StringValue {
	return terra.ReferenceAsString(mt.ref.Append("metric_column"))
}

func (mt MetricTriggerAttributes) MetricTriggerType() terra.StringValue {
	return terra.ReferenceAsString(mt.ref.Append("metric_trigger_type"))
}

func (mt MetricTriggerAttributes) Operator() terra.StringValue {
	return terra.ReferenceAsString(mt.ref.Append("operator"))
}

func (mt MetricTriggerAttributes) Threshold() terra.NumberValue {
	return terra.ReferenceAsNumber(mt.ref.Append("threshold"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type ActionState struct {
	ActionGroup          []string `json:"action_group"`
	CustomWebhookPayload string   `json:"custom_webhook_payload"`
	EmailSubject         string   `json:"email_subject"`
}

type TriggerState struct {
	Operator      string               `json:"operator"`
	Threshold     float64              `json:"threshold"`
	MetricTrigger []MetricTriggerState `json:"metric_trigger"`
}

type MetricTriggerState struct {
	MetricColumn      string  `json:"metric_column"`
	MetricTriggerType string  `json:"metric_trigger_type"`
	Operator          string  `json:"operator"`
	Threshold         float64 `json:"threshold"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}