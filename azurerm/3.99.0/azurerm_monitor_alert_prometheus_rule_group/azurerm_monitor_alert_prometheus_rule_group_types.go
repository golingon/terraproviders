// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_monitor_alert_prometheus_rule_group

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Rule struct {
	// Alert: string, optional
	Alert terra.StringValue `hcl:"alert,attr"`
	// Annotations: map of string, optional
	Annotations terra.MapValue[terra.StringValue] `hcl:"annotations,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Expression: string, required
	Expression terra.StringValue `hcl:"expression,attr" validate:"required"`
	// For: string, optional
	For terra.StringValue `hcl:"for,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Record: string, optional
	Record terra.StringValue `hcl:"record,attr"`
	// Severity: number, optional
	Severity terra.NumberValue `hcl:"severity,attr"`
	// RuleAction: min=0,max=5
	Action []RuleAction `hcl:"action,block" validate:"min=0,max=5"`
	// RuleAlertResolution: optional
	AlertResolution *RuleAlertResolution `hcl:"alert_resolution,block"`
}

type RuleAction struct {
	// ActionGroupId: string, required
	ActionGroupId terra.StringValue `hcl:"action_group_id,attr" validate:"required"`
	// ActionProperties: map of string, optional
	ActionProperties terra.MapValue[terra.StringValue] `hcl:"action_properties,attr"`
}

type RuleAlertResolution struct {
	// AutoResolved: bool, optional
	AutoResolved terra.BoolValue `hcl:"auto_resolved,attr"`
	// TimeToResolve: string, optional
	TimeToResolve terra.StringValue `hcl:"time_to_resolve,attr"`
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

func (r RuleAttributes) Alert() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("alert"))
}

func (r RuleAttributes) Annotations() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](r.ref.Append("annotations"))
}

func (r RuleAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(r.ref.Append("enabled"))
}

func (r RuleAttributes) Expression() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("expression"))
}

func (r RuleAttributes) For() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("for"))
}

func (r RuleAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](r.ref.Append("labels"))
}

func (r RuleAttributes) Record() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("record"))
}

func (r RuleAttributes) Severity() terra.NumberValue {
	return terra.ReferenceAsNumber(r.ref.Append("severity"))
}

func (r RuleAttributes) Action() terra.ListValue[RuleActionAttributes] {
	return terra.ReferenceAsList[RuleActionAttributes](r.ref.Append("action"))
}

func (r RuleAttributes) AlertResolution() terra.ListValue[RuleAlertResolutionAttributes] {
	return terra.ReferenceAsList[RuleAlertResolutionAttributes](r.ref.Append("alert_resolution"))
}

type RuleActionAttributes struct {
	ref terra.Reference
}

func (a RuleActionAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a RuleActionAttributes) InternalWithRef(ref terra.Reference) RuleActionAttributes {
	return RuleActionAttributes{ref: ref}
}

func (a RuleActionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a RuleActionAttributes) ActionGroupId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("action_group_id"))
}

func (a RuleActionAttributes) ActionProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](a.ref.Append("action_properties"))
}

type RuleAlertResolutionAttributes struct {
	ref terra.Reference
}

func (ar RuleAlertResolutionAttributes) InternalRef() (terra.Reference, error) {
	return ar.ref, nil
}

func (ar RuleAlertResolutionAttributes) InternalWithRef(ref terra.Reference) RuleAlertResolutionAttributes {
	return RuleAlertResolutionAttributes{ref: ref}
}

func (ar RuleAlertResolutionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ar.ref.InternalTokens()
}

func (ar RuleAlertResolutionAttributes) AutoResolved() terra.BoolValue {
	return terra.ReferenceAsBool(ar.ref.Append("auto_resolved"))
}

func (ar RuleAlertResolutionAttributes) TimeToResolve() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("time_to_resolve"))
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

type RuleState struct {
	Alert           string                     `json:"alert"`
	Annotations     map[string]string          `json:"annotations"`
	Enabled         bool                       `json:"enabled"`
	Expression      string                     `json:"expression"`
	For             string                     `json:"for"`
	Labels          map[string]string          `json:"labels"`
	Record          string                     `json:"record"`
	Severity        float64                    `json:"severity"`
	Action          []RuleActionState          `json:"action"`
	AlertResolution []RuleAlertResolutionState `json:"alert_resolution"`
}

type RuleActionState struct {
	ActionGroupId    string            `json:"action_group_id"`
	ActionProperties map[string]string `json:"action_properties"`
}

type RuleAlertResolutionState struct {
	AutoResolved  bool   `json:"auto_resolved"`
	TimeToResolve string `json:"time_to_resolve"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
