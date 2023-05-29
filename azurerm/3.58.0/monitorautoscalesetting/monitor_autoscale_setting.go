// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package monitorautoscalesetting

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Notification struct {
	// Email: optional
	Email *Email `hcl:"email,block"`
	// Webhook: min=0
	Webhook []Webhook `hcl:"webhook,block" validate:"min=0"`
}

type Email struct {
	// CustomEmails: list of string, optional
	CustomEmails terra.ListValue[terra.StringValue] `hcl:"custom_emails,attr"`
	// SendToSubscriptionAdministrator: bool, optional
	SendToSubscriptionAdministrator terra.BoolValue `hcl:"send_to_subscription_administrator,attr"`
	// SendToSubscriptionCoAdministrator: bool, optional
	SendToSubscriptionCoAdministrator terra.BoolValue `hcl:"send_to_subscription_co_administrator,attr"`
}

type Webhook struct {
	// Properties: map of string, optional
	Properties terra.MapValue[terra.StringValue] `hcl:"properties,attr"`
	// ServiceUri: string, required
	ServiceUri terra.StringValue `hcl:"service_uri,attr" validate:"required"`
}

type Profile struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Capacity: required
	Capacity *Capacity `hcl:"capacity,block" validate:"required"`
	// FixedDate: optional
	FixedDate *FixedDate `hcl:"fixed_date,block"`
	// Recurrence: optional
	Recurrence *Recurrence `hcl:"recurrence,block"`
	// Rule: min=0,max=10
	Rule []Rule `hcl:"rule,block" validate:"min=0,max=10"`
}

type Capacity struct {
	// Default: number, required
	Default terra.NumberValue `hcl:"default,attr" validate:"required"`
	// Maximum: number, required
	Maximum terra.NumberValue `hcl:"maximum,attr" validate:"required"`
	// Minimum: number, required
	Minimum terra.NumberValue `hcl:"minimum,attr" validate:"required"`
}

type FixedDate struct {
	// End: string, required
	End terra.StringValue `hcl:"end,attr" validate:"required"`
	// Start: string, required
	Start terra.StringValue `hcl:"start,attr" validate:"required"`
	// Timezone: string, optional
	Timezone terra.StringValue `hcl:"timezone,attr"`
}

type Recurrence struct {
	// Days: list of string, required
	Days terra.ListValue[terra.StringValue] `hcl:"days,attr" validate:"required"`
	// Hours: list of number, required
	Hours terra.ListValue[terra.NumberValue] `hcl:"hours,attr" validate:"required"`
	// Minutes: list of number, required
	Minutes terra.ListValue[terra.NumberValue] `hcl:"minutes,attr" validate:"required"`
	// Timezone: string, optional
	Timezone terra.StringValue `hcl:"timezone,attr"`
}

type Rule struct {
	// MetricTrigger: required
	MetricTrigger *MetricTrigger `hcl:"metric_trigger,block" validate:"required"`
	// ScaleAction: required
	ScaleAction *ScaleAction `hcl:"scale_action,block" validate:"required"`
}

type MetricTrigger struct {
	// DivideByInstanceCount: bool, optional
	DivideByInstanceCount terra.BoolValue `hcl:"divide_by_instance_count,attr"`
	// MetricName: string, required
	MetricName terra.StringValue `hcl:"metric_name,attr" validate:"required"`
	// MetricNamespace: string, optional
	MetricNamespace terra.StringValue `hcl:"metric_namespace,attr"`
	// MetricResourceId: string, required
	MetricResourceId terra.StringValue `hcl:"metric_resource_id,attr" validate:"required"`
	// Operator: string, required
	Operator terra.StringValue `hcl:"operator,attr" validate:"required"`
	// Statistic: string, required
	Statistic terra.StringValue `hcl:"statistic,attr" validate:"required"`
	// Threshold: number, required
	Threshold terra.NumberValue `hcl:"threshold,attr" validate:"required"`
	// TimeAggregation: string, required
	TimeAggregation terra.StringValue `hcl:"time_aggregation,attr" validate:"required"`
	// TimeGrain: string, required
	TimeGrain terra.StringValue `hcl:"time_grain,attr" validate:"required"`
	// TimeWindow: string, required
	TimeWindow terra.StringValue `hcl:"time_window,attr" validate:"required"`
	// Dimensions: min=0
	Dimensions []Dimensions `hcl:"dimensions,block" validate:"min=0"`
}

type Dimensions struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Operator: string, required
	Operator terra.StringValue `hcl:"operator,attr" validate:"required"`
	// Values: list of string, required
	Values terra.ListValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type ScaleAction struct {
	// Cooldown: string, required
	Cooldown terra.StringValue `hcl:"cooldown,attr" validate:"required"`
	// Direction: string, required
	Direction terra.StringValue `hcl:"direction,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Value: number, required
	Value terra.NumberValue `hcl:"value,attr" validate:"required"`
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

type NotificationAttributes struct {
	ref terra.Reference
}

func (n NotificationAttributes) InternalRef() (terra.Reference, error) {
	return n.ref, nil
}

func (n NotificationAttributes) InternalWithRef(ref terra.Reference) NotificationAttributes {
	return NotificationAttributes{ref: ref}
}

func (n NotificationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return n.ref.InternalTokens()
}

func (n NotificationAttributes) Email() terra.ListValue[EmailAttributes] {
	return terra.ReferenceAsList[EmailAttributes](n.ref.Append("email"))
}

func (n NotificationAttributes) Webhook() terra.ListValue[WebhookAttributes] {
	return terra.ReferenceAsList[WebhookAttributes](n.ref.Append("webhook"))
}

type EmailAttributes struct {
	ref terra.Reference
}

func (e EmailAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e EmailAttributes) InternalWithRef(ref terra.Reference) EmailAttributes {
	return EmailAttributes{ref: ref}
}

func (e EmailAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return e.ref.InternalTokens()
}

func (e EmailAttributes) CustomEmails() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](e.ref.Append("custom_emails"))
}

func (e EmailAttributes) SendToSubscriptionAdministrator() terra.BoolValue {
	return terra.ReferenceAsBool(e.ref.Append("send_to_subscription_administrator"))
}

func (e EmailAttributes) SendToSubscriptionCoAdministrator() terra.BoolValue {
	return terra.ReferenceAsBool(e.ref.Append("send_to_subscription_co_administrator"))
}

type WebhookAttributes struct {
	ref terra.Reference
}

func (w WebhookAttributes) InternalRef() (terra.Reference, error) {
	return w.ref, nil
}

func (w WebhookAttributes) InternalWithRef(ref terra.Reference) WebhookAttributes {
	return WebhookAttributes{ref: ref}
}

func (w WebhookAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return w.ref.InternalTokens()
}

func (w WebhookAttributes) Properties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](w.ref.Append("properties"))
}

func (w WebhookAttributes) ServiceUri() terra.StringValue {
	return terra.ReferenceAsString(w.ref.Append("service_uri"))
}

type ProfileAttributes struct {
	ref terra.Reference
}

func (p ProfileAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p ProfileAttributes) InternalWithRef(ref terra.Reference) ProfileAttributes {
	return ProfileAttributes{ref: ref}
}

func (p ProfileAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p ProfileAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("name"))
}

func (p ProfileAttributes) Capacity() terra.ListValue[CapacityAttributes] {
	return terra.ReferenceAsList[CapacityAttributes](p.ref.Append("capacity"))
}

func (p ProfileAttributes) FixedDate() terra.ListValue[FixedDateAttributes] {
	return terra.ReferenceAsList[FixedDateAttributes](p.ref.Append("fixed_date"))
}

func (p ProfileAttributes) Recurrence() terra.ListValue[RecurrenceAttributes] {
	return terra.ReferenceAsList[RecurrenceAttributes](p.ref.Append("recurrence"))
}

func (p ProfileAttributes) Rule() terra.ListValue[RuleAttributes] {
	return terra.ReferenceAsList[RuleAttributes](p.ref.Append("rule"))
}

type CapacityAttributes struct {
	ref terra.Reference
}

func (c CapacityAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c CapacityAttributes) InternalWithRef(ref terra.Reference) CapacityAttributes {
	return CapacityAttributes{ref: ref}
}

func (c CapacityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c CapacityAttributes) Default() terra.NumberValue {
	return terra.ReferenceAsNumber(c.ref.Append("default"))
}

func (c CapacityAttributes) Maximum() terra.NumberValue {
	return terra.ReferenceAsNumber(c.ref.Append("maximum"))
}

func (c CapacityAttributes) Minimum() terra.NumberValue {
	return terra.ReferenceAsNumber(c.ref.Append("minimum"))
}

type FixedDateAttributes struct {
	ref terra.Reference
}

func (fd FixedDateAttributes) InternalRef() (terra.Reference, error) {
	return fd.ref, nil
}

func (fd FixedDateAttributes) InternalWithRef(ref terra.Reference) FixedDateAttributes {
	return FixedDateAttributes{ref: ref}
}

func (fd FixedDateAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fd.ref.InternalTokens()
}

func (fd FixedDateAttributes) End() terra.StringValue {
	return terra.ReferenceAsString(fd.ref.Append("end"))
}

func (fd FixedDateAttributes) Start() terra.StringValue {
	return terra.ReferenceAsString(fd.ref.Append("start"))
}

func (fd FixedDateAttributes) Timezone() terra.StringValue {
	return terra.ReferenceAsString(fd.ref.Append("timezone"))
}

type RecurrenceAttributes struct {
	ref terra.Reference
}

func (r RecurrenceAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RecurrenceAttributes) InternalWithRef(ref terra.Reference) RecurrenceAttributes {
	return RecurrenceAttributes{ref: ref}
}

func (r RecurrenceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RecurrenceAttributes) Days() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](r.ref.Append("days"))
}

func (r RecurrenceAttributes) Hours() terra.ListValue[terra.NumberValue] {
	return terra.ReferenceAsList[terra.NumberValue](r.ref.Append("hours"))
}

func (r RecurrenceAttributes) Minutes() terra.ListValue[terra.NumberValue] {
	return terra.ReferenceAsList[terra.NumberValue](r.ref.Append("minutes"))
}

func (r RecurrenceAttributes) Timezone() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("timezone"))
}

type RuleAttributes struct {
	ref terra.Reference
}

func (r RuleAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RuleAttributes) InternalWithRef(ref terra.Reference) RuleAttributes {
	return RuleAttributes{ref: ref}
}

func (r RuleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RuleAttributes) MetricTrigger() terra.ListValue[MetricTriggerAttributes] {
	return terra.ReferenceAsList[MetricTriggerAttributes](r.ref.Append("metric_trigger"))
}

func (r RuleAttributes) ScaleAction() terra.ListValue[ScaleActionAttributes] {
	return terra.ReferenceAsList[ScaleActionAttributes](r.ref.Append("scale_action"))
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

func (mt MetricTriggerAttributes) DivideByInstanceCount() terra.BoolValue {
	return terra.ReferenceAsBool(mt.ref.Append("divide_by_instance_count"))
}

func (mt MetricTriggerAttributes) MetricName() terra.StringValue {
	return terra.ReferenceAsString(mt.ref.Append("metric_name"))
}

func (mt MetricTriggerAttributes) MetricNamespace() terra.StringValue {
	return terra.ReferenceAsString(mt.ref.Append("metric_namespace"))
}

func (mt MetricTriggerAttributes) MetricResourceId() terra.StringValue {
	return terra.ReferenceAsString(mt.ref.Append("metric_resource_id"))
}

func (mt MetricTriggerAttributes) Operator() terra.StringValue {
	return terra.ReferenceAsString(mt.ref.Append("operator"))
}

func (mt MetricTriggerAttributes) Statistic() terra.StringValue {
	return terra.ReferenceAsString(mt.ref.Append("statistic"))
}

func (mt MetricTriggerAttributes) Threshold() terra.NumberValue {
	return terra.ReferenceAsNumber(mt.ref.Append("threshold"))
}

func (mt MetricTriggerAttributes) TimeAggregation() terra.StringValue {
	return terra.ReferenceAsString(mt.ref.Append("time_aggregation"))
}

func (mt MetricTriggerAttributes) TimeGrain() terra.StringValue {
	return terra.ReferenceAsString(mt.ref.Append("time_grain"))
}

func (mt MetricTriggerAttributes) TimeWindow() terra.StringValue {
	return terra.ReferenceAsString(mt.ref.Append("time_window"))
}

func (mt MetricTriggerAttributes) Dimensions() terra.ListValue[DimensionsAttributes] {
	return terra.ReferenceAsList[DimensionsAttributes](mt.ref.Append("dimensions"))
}

type DimensionsAttributes struct {
	ref terra.Reference
}

func (d DimensionsAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d DimensionsAttributes) InternalWithRef(ref terra.Reference) DimensionsAttributes {
	return DimensionsAttributes{ref: ref}
}

func (d DimensionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d DimensionsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("name"))
}

func (d DimensionsAttributes) Operator() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("operator"))
}

func (d DimensionsAttributes) Values() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](d.ref.Append("values"))
}

type ScaleActionAttributes struct {
	ref terra.Reference
}

func (sa ScaleActionAttributes) InternalRef() (terra.Reference, error) {
	return sa.ref, nil
}

func (sa ScaleActionAttributes) InternalWithRef(ref terra.Reference) ScaleActionAttributes {
	return ScaleActionAttributes{ref: ref}
}

func (sa ScaleActionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sa.ref.InternalTokens()
}

func (sa ScaleActionAttributes) Cooldown() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("cooldown"))
}

func (sa ScaleActionAttributes) Direction() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("direction"))
}

func (sa ScaleActionAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("type"))
}

func (sa ScaleActionAttributes) Value() terra.NumberValue {
	return terra.ReferenceAsNumber(sa.ref.Append("value"))
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

type NotificationState struct {
	Email   []EmailState   `json:"email"`
	Webhook []WebhookState `json:"webhook"`
}

type EmailState struct {
	CustomEmails                      []string `json:"custom_emails"`
	SendToSubscriptionAdministrator   bool     `json:"send_to_subscription_administrator"`
	SendToSubscriptionCoAdministrator bool     `json:"send_to_subscription_co_administrator"`
}

type WebhookState struct {
	Properties map[string]string `json:"properties"`
	ServiceUri string            `json:"service_uri"`
}

type ProfileState struct {
	Name       string            `json:"name"`
	Capacity   []CapacityState   `json:"capacity"`
	FixedDate  []FixedDateState  `json:"fixed_date"`
	Recurrence []RecurrenceState `json:"recurrence"`
	Rule       []RuleState       `json:"rule"`
}

type CapacityState struct {
	Default float64 `json:"default"`
	Maximum float64 `json:"maximum"`
	Minimum float64 `json:"minimum"`
}

type FixedDateState struct {
	End      string `json:"end"`
	Start    string `json:"start"`
	Timezone string `json:"timezone"`
}

type RecurrenceState struct {
	Days     []string  `json:"days"`
	Hours    []float64 `json:"hours"`
	Minutes  []float64 `json:"minutes"`
	Timezone string    `json:"timezone"`
}

type RuleState struct {
	MetricTrigger []MetricTriggerState `json:"metric_trigger"`
	ScaleAction   []ScaleActionState   `json:"scale_action"`
}

type MetricTriggerState struct {
	DivideByInstanceCount bool              `json:"divide_by_instance_count"`
	MetricName            string            `json:"metric_name"`
	MetricNamespace       string            `json:"metric_namespace"`
	MetricResourceId      string            `json:"metric_resource_id"`
	Operator              string            `json:"operator"`
	Statistic             string            `json:"statistic"`
	Threshold             float64           `json:"threshold"`
	TimeAggregation       string            `json:"time_aggregation"`
	TimeGrain             string            `json:"time_grain"`
	TimeWindow            string            `json:"time_window"`
	Dimensions            []DimensionsState `json:"dimensions"`
}

type DimensionsState struct {
	Name     string   `json:"name"`
	Operator string   `json:"operator"`
	Values   []string `json:"values"`
}

type ScaleActionState struct {
	Cooldown  string  `json:"cooldown"`
	Direction string  `json:"direction"`
	Type      string  `json:"type"`
	Value     float64 `json:"value"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}