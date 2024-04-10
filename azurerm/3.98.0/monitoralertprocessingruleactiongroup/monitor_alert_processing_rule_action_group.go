// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package monitoralertprocessingruleactiongroup

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Condition struct {
	// AlertContext: optional
	AlertContext *AlertContext `hcl:"alert_context,block"`
	// AlertRuleId: optional
	AlertRuleId *AlertRuleId `hcl:"alert_rule_id,block"`
	// AlertRuleName: optional
	AlertRuleName *AlertRuleName `hcl:"alert_rule_name,block"`
	// Description: optional
	Description *Description `hcl:"description,block"`
	// MonitorCondition: optional
	MonitorCondition *MonitorCondition `hcl:"monitor_condition,block"`
	// MonitorService: optional
	MonitorService *MonitorService `hcl:"monitor_service,block"`
	// Severity: optional
	Severity *Severity `hcl:"severity,block"`
	// SignalType: optional
	SignalType *SignalType `hcl:"signal_type,block"`
	// TargetResource: optional
	TargetResource *TargetResource `hcl:"target_resource,block"`
	// TargetResourceGroup: optional
	TargetResourceGroup *TargetResourceGroup `hcl:"target_resource_group,block"`
	// TargetResourceType: optional
	TargetResourceType *TargetResourceType `hcl:"target_resource_type,block"`
}

type AlertContext struct {
	// Operator: string, required
	Operator terra.StringValue `hcl:"operator,attr" validate:"required"`
	// Values: list of string, required
	Values terra.ListValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type AlertRuleId struct {
	// Operator: string, required
	Operator terra.StringValue `hcl:"operator,attr" validate:"required"`
	// Values: list of string, required
	Values terra.ListValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type AlertRuleName struct {
	// Operator: string, required
	Operator terra.StringValue `hcl:"operator,attr" validate:"required"`
	// Values: list of string, required
	Values terra.ListValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type Description struct {
	// Operator: string, required
	Operator terra.StringValue `hcl:"operator,attr" validate:"required"`
	// Values: list of string, required
	Values terra.ListValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type MonitorCondition struct {
	// Operator: string, required
	Operator terra.StringValue `hcl:"operator,attr" validate:"required"`
	// Values: list of string, required
	Values terra.ListValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type MonitorService struct {
	// Operator: string, required
	Operator terra.StringValue `hcl:"operator,attr" validate:"required"`
	// Values: list of string, required
	Values terra.ListValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type Severity struct {
	// Operator: string, required
	Operator terra.StringValue `hcl:"operator,attr" validate:"required"`
	// Values: list of string, required
	Values terra.ListValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type SignalType struct {
	// Operator: string, required
	Operator terra.StringValue `hcl:"operator,attr" validate:"required"`
	// Values: list of string, required
	Values terra.ListValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type TargetResource struct {
	// Operator: string, required
	Operator terra.StringValue `hcl:"operator,attr" validate:"required"`
	// Values: list of string, required
	Values terra.ListValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type TargetResourceGroup struct {
	// Operator: string, required
	Operator terra.StringValue `hcl:"operator,attr" validate:"required"`
	// Values: list of string, required
	Values terra.ListValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type TargetResourceType struct {
	// Operator: string, required
	Operator terra.StringValue `hcl:"operator,attr" validate:"required"`
	// Values: list of string, required
	Values terra.ListValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type Schedule struct {
	// EffectiveFrom: string, optional
	EffectiveFrom terra.StringValue `hcl:"effective_from,attr"`
	// EffectiveUntil: string, optional
	EffectiveUntil terra.StringValue `hcl:"effective_until,attr"`
	// TimeZone: string, optional
	TimeZone terra.StringValue `hcl:"time_zone,attr"`
	// Recurrence: optional
	Recurrence *Recurrence `hcl:"recurrence,block"`
}

type Recurrence struct {
	// Daily: min=0
	Daily []Daily `hcl:"daily,block" validate:"min=0"`
	// Monthly: min=0
	Monthly []Monthly `hcl:"monthly,block" validate:"min=0"`
	// Weekly: min=0
	Weekly []Weekly `hcl:"weekly,block" validate:"min=0"`
}

type Daily struct {
	// EndTime: string, required
	EndTime terra.StringValue `hcl:"end_time,attr" validate:"required"`
	// StartTime: string, required
	StartTime terra.StringValue `hcl:"start_time,attr" validate:"required"`
}

type Monthly struct {
	// DaysOfMonth: list of number, required
	DaysOfMonth terra.ListValue[terra.NumberValue] `hcl:"days_of_month,attr" validate:"required"`
	// EndTime: string, optional
	EndTime terra.StringValue `hcl:"end_time,attr"`
	// StartTime: string, optional
	StartTime terra.StringValue `hcl:"start_time,attr"`
}

type Weekly struct {
	// DaysOfWeek: list of string, required
	DaysOfWeek terra.ListValue[terra.StringValue] `hcl:"days_of_week,attr" validate:"required"`
	// EndTime: string, optional
	EndTime terra.StringValue `hcl:"end_time,attr"`
	// StartTime: string, optional
	StartTime terra.StringValue `hcl:"start_time,attr"`
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

type ConditionAttributes struct {
	ref terra.Reference
}

func (c ConditionAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ConditionAttributes) InternalWithRef(ref terra.Reference) ConditionAttributes {
	return ConditionAttributes{ref: ref}
}

func (c ConditionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ConditionAttributes) AlertContext() terra.ListValue[AlertContextAttributes] {
	return terra.ReferenceAsList[AlertContextAttributes](c.ref.Append("alert_context"))
}

func (c ConditionAttributes) AlertRuleId() terra.ListValue[AlertRuleIdAttributes] {
	return terra.ReferenceAsList[AlertRuleIdAttributes](c.ref.Append("alert_rule_id"))
}

func (c ConditionAttributes) AlertRuleName() terra.ListValue[AlertRuleNameAttributes] {
	return terra.ReferenceAsList[AlertRuleNameAttributes](c.ref.Append("alert_rule_name"))
}

func (c ConditionAttributes) Description() terra.ListValue[DescriptionAttributes] {
	return terra.ReferenceAsList[DescriptionAttributes](c.ref.Append("description"))
}

func (c ConditionAttributes) MonitorCondition() terra.ListValue[MonitorConditionAttributes] {
	return terra.ReferenceAsList[MonitorConditionAttributes](c.ref.Append("monitor_condition"))
}

func (c ConditionAttributes) MonitorService() terra.ListValue[MonitorServiceAttributes] {
	return terra.ReferenceAsList[MonitorServiceAttributes](c.ref.Append("monitor_service"))
}

func (c ConditionAttributes) Severity() terra.ListValue[SeverityAttributes] {
	return terra.ReferenceAsList[SeverityAttributes](c.ref.Append("severity"))
}

func (c ConditionAttributes) SignalType() terra.ListValue[SignalTypeAttributes] {
	return terra.ReferenceAsList[SignalTypeAttributes](c.ref.Append("signal_type"))
}

func (c ConditionAttributes) TargetResource() terra.ListValue[TargetResourceAttributes] {
	return terra.ReferenceAsList[TargetResourceAttributes](c.ref.Append("target_resource"))
}

func (c ConditionAttributes) TargetResourceGroup() terra.ListValue[TargetResourceGroupAttributes] {
	return terra.ReferenceAsList[TargetResourceGroupAttributes](c.ref.Append("target_resource_group"))
}

func (c ConditionAttributes) TargetResourceType() terra.ListValue[TargetResourceTypeAttributes] {
	return terra.ReferenceAsList[TargetResourceTypeAttributes](c.ref.Append("target_resource_type"))
}

type AlertContextAttributes struct {
	ref terra.Reference
}

func (ac AlertContextAttributes) InternalRef() (terra.Reference, error) {
	return ac.ref, nil
}

func (ac AlertContextAttributes) InternalWithRef(ref terra.Reference) AlertContextAttributes {
	return AlertContextAttributes{ref: ref}
}

func (ac AlertContextAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ac.ref.InternalTokens()
}

func (ac AlertContextAttributes) Operator() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("operator"))
}

func (ac AlertContextAttributes) Values() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ac.ref.Append("values"))
}

type AlertRuleIdAttributes struct {
	ref terra.Reference
}

func (ari AlertRuleIdAttributes) InternalRef() (terra.Reference, error) {
	return ari.ref, nil
}

func (ari AlertRuleIdAttributes) InternalWithRef(ref terra.Reference) AlertRuleIdAttributes {
	return AlertRuleIdAttributes{ref: ref}
}

func (ari AlertRuleIdAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ari.ref.InternalTokens()
}

func (ari AlertRuleIdAttributes) Operator() terra.StringValue {
	return terra.ReferenceAsString(ari.ref.Append("operator"))
}

func (ari AlertRuleIdAttributes) Values() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ari.ref.Append("values"))
}

type AlertRuleNameAttributes struct {
	ref terra.Reference
}

func (arn AlertRuleNameAttributes) InternalRef() (terra.Reference, error) {
	return arn.ref, nil
}

func (arn AlertRuleNameAttributes) InternalWithRef(ref terra.Reference) AlertRuleNameAttributes {
	return AlertRuleNameAttributes{ref: ref}
}

func (arn AlertRuleNameAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return arn.ref.InternalTokens()
}

func (arn AlertRuleNameAttributes) Operator() terra.StringValue {
	return terra.ReferenceAsString(arn.ref.Append("operator"))
}

func (arn AlertRuleNameAttributes) Values() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](arn.ref.Append("values"))
}

type DescriptionAttributes struct {
	ref terra.Reference
}

func (d DescriptionAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d DescriptionAttributes) InternalWithRef(ref terra.Reference) DescriptionAttributes {
	return DescriptionAttributes{ref: ref}
}

func (d DescriptionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d DescriptionAttributes) Operator() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("operator"))
}

func (d DescriptionAttributes) Values() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](d.ref.Append("values"))
}

type MonitorConditionAttributes struct {
	ref terra.Reference
}

func (mc MonitorConditionAttributes) InternalRef() (terra.Reference, error) {
	return mc.ref, nil
}

func (mc MonitorConditionAttributes) InternalWithRef(ref terra.Reference) MonitorConditionAttributes {
	return MonitorConditionAttributes{ref: ref}
}

func (mc MonitorConditionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mc.ref.InternalTokens()
}

func (mc MonitorConditionAttributes) Operator() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("operator"))
}

func (mc MonitorConditionAttributes) Values() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](mc.ref.Append("values"))
}

type MonitorServiceAttributes struct {
	ref terra.Reference
}

func (ms MonitorServiceAttributes) InternalRef() (terra.Reference, error) {
	return ms.ref, nil
}

func (ms MonitorServiceAttributes) InternalWithRef(ref terra.Reference) MonitorServiceAttributes {
	return MonitorServiceAttributes{ref: ref}
}

func (ms MonitorServiceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ms.ref.InternalTokens()
}

func (ms MonitorServiceAttributes) Operator() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("operator"))
}

func (ms MonitorServiceAttributes) Values() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ms.ref.Append("values"))
}

type SeverityAttributes struct {
	ref terra.Reference
}

func (s SeverityAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SeverityAttributes) InternalWithRef(ref terra.Reference) SeverityAttributes {
	return SeverityAttributes{ref: ref}
}

func (s SeverityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SeverityAttributes) Operator() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("operator"))
}

func (s SeverityAttributes) Values() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](s.ref.Append("values"))
}

type SignalTypeAttributes struct {
	ref terra.Reference
}

func (st SignalTypeAttributes) InternalRef() (terra.Reference, error) {
	return st.ref, nil
}

func (st SignalTypeAttributes) InternalWithRef(ref terra.Reference) SignalTypeAttributes {
	return SignalTypeAttributes{ref: ref}
}

func (st SignalTypeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return st.ref.InternalTokens()
}

func (st SignalTypeAttributes) Operator() terra.StringValue {
	return terra.ReferenceAsString(st.ref.Append("operator"))
}

func (st SignalTypeAttributes) Values() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](st.ref.Append("values"))
}

type TargetResourceAttributes struct {
	ref terra.Reference
}

func (tr TargetResourceAttributes) InternalRef() (terra.Reference, error) {
	return tr.ref, nil
}

func (tr TargetResourceAttributes) InternalWithRef(ref terra.Reference) TargetResourceAttributes {
	return TargetResourceAttributes{ref: ref}
}

func (tr TargetResourceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tr.ref.InternalTokens()
}

func (tr TargetResourceAttributes) Operator() terra.StringValue {
	return terra.ReferenceAsString(tr.ref.Append("operator"))
}

func (tr TargetResourceAttributes) Values() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](tr.ref.Append("values"))
}

type TargetResourceGroupAttributes struct {
	ref terra.Reference
}

func (trg TargetResourceGroupAttributes) InternalRef() (terra.Reference, error) {
	return trg.ref, nil
}

func (trg TargetResourceGroupAttributes) InternalWithRef(ref terra.Reference) TargetResourceGroupAttributes {
	return TargetResourceGroupAttributes{ref: ref}
}

func (trg TargetResourceGroupAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return trg.ref.InternalTokens()
}

func (trg TargetResourceGroupAttributes) Operator() terra.StringValue {
	return terra.ReferenceAsString(trg.ref.Append("operator"))
}

func (trg TargetResourceGroupAttributes) Values() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](trg.ref.Append("values"))
}

type TargetResourceTypeAttributes struct {
	ref terra.Reference
}

func (trt TargetResourceTypeAttributes) InternalRef() (terra.Reference, error) {
	return trt.ref, nil
}

func (trt TargetResourceTypeAttributes) InternalWithRef(ref terra.Reference) TargetResourceTypeAttributes {
	return TargetResourceTypeAttributes{ref: ref}
}

func (trt TargetResourceTypeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return trt.ref.InternalTokens()
}

func (trt TargetResourceTypeAttributes) Operator() terra.StringValue {
	return terra.ReferenceAsString(trt.ref.Append("operator"))
}

func (trt TargetResourceTypeAttributes) Values() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](trt.ref.Append("values"))
}

type ScheduleAttributes struct {
	ref terra.Reference
}

func (s ScheduleAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s ScheduleAttributes) InternalWithRef(ref terra.Reference) ScheduleAttributes {
	return ScheduleAttributes{ref: ref}
}

func (s ScheduleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s ScheduleAttributes) EffectiveFrom() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("effective_from"))
}

func (s ScheduleAttributes) EffectiveUntil() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("effective_until"))
}

func (s ScheduleAttributes) TimeZone() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("time_zone"))
}

func (s ScheduleAttributes) Recurrence() terra.ListValue[RecurrenceAttributes] {
	return terra.ReferenceAsList[RecurrenceAttributes](s.ref.Append("recurrence"))
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

func (r RecurrenceAttributes) Daily() terra.ListValue[DailyAttributes] {
	return terra.ReferenceAsList[DailyAttributes](r.ref.Append("daily"))
}

func (r RecurrenceAttributes) Monthly() terra.ListValue[MonthlyAttributes] {
	return terra.ReferenceAsList[MonthlyAttributes](r.ref.Append("monthly"))
}

func (r RecurrenceAttributes) Weekly() terra.ListValue[WeeklyAttributes] {
	return terra.ReferenceAsList[WeeklyAttributes](r.ref.Append("weekly"))
}

type DailyAttributes struct {
	ref terra.Reference
}

func (d DailyAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d DailyAttributes) InternalWithRef(ref terra.Reference) DailyAttributes {
	return DailyAttributes{ref: ref}
}

func (d DailyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d DailyAttributes) EndTime() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("end_time"))
}

func (d DailyAttributes) StartTime() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("start_time"))
}

type MonthlyAttributes struct {
	ref terra.Reference
}

func (m MonthlyAttributes) InternalRef() (terra.Reference, error) {
	return m.ref, nil
}

func (m MonthlyAttributes) InternalWithRef(ref terra.Reference) MonthlyAttributes {
	return MonthlyAttributes{ref: ref}
}

func (m MonthlyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return m.ref.InternalTokens()
}

func (m MonthlyAttributes) DaysOfMonth() terra.ListValue[terra.NumberValue] {
	return terra.ReferenceAsList[terra.NumberValue](m.ref.Append("days_of_month"))
}

func (m MonthlyAttributes) EndTime() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("end_time"))
}

func (m MonthlyAttributes) StartTime() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("start_time"))
}

type WeeklyAttributes struct {
	ref terra.Reference
}

func (w WeeklyAttributes) InternalRef() (terra.Reference, error) {
	return w.ref, nil
}

func (w WeeklyAttributes) InternalWithRef(ref terra.Reference) WeeklyAttributes {
	return WeeklyAttributes{ref: ref}
}

func (w WeeklyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return w.ref.InternalTokens()
}

func (w WeeklyAttributes) DaysOfWeek() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](w.ref.Append("days_of_week"))
}

func (w WeeklyAttributes) EndTime() terra.StringValue {
	return terra.ReferenceAsString(w.ref.Append("end_time"))
}

func (w WeeklyAttributes) StartTime() terra.StringValue {
	return terra.ReferenceAsString(w.ref.Append("start_time"))
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

type ConditionState struct {
	AlertContext        []AlertContextState        `json:"alert_context"`
	AlertRuleId         []AlertRuleIdState         `json:"alert_rule_id"`
	AlertRuleName       []AlertRuleNameState       `json:"alert_rule_name"`
	Description         []DescriptionState         `json:"description"`
	MonitorCondition    []MonitorConditionState    `json:"monitor_condition"`
	MonitorService      []MonitorServiceState      `json:"monitor_service"`
	Severity            []SeverityState            `json:"severity"`
	SignalType          []SignalTypeState          `json:"signal_type"`
	TargetResource      []TargetResourceState      `json:"target_resource"`
	TargetResourceGroup []TargetResourceGroupState `json:"target_resource_group"`
	TargetResourceType  []TargetResourceTypeState  `json:"target_resource_type"`
}

type AlertContextState struct {
	Operator string   `json:"operator"`
	Values   []string `json:"values"`
}

type AlertRuleIdState struct {
	Operator string   `json:"operator"`
	Values   []string `json:"values"`
}

type AlertRuleNameState struct {
	Operator string   `json:"operator"`
	Values   []string `json:"values"`
}

type DescriptionState struct {
	Operator string   `json:"operator"`
	Values   []string `json:"values"`
}

type MonitorConditionState struct {
	Operator string   `json:"operator"`
	Values   []string `json:"values"`
}

type MonitorServiceState struct {
	Operator string   `json:"operator"`
	Values   []string `json:"values"`
}

type SeverityState struct {
	Operator string   `json:"operator"`
	Values   []string `json:"values"`
}

type SignalTypeState struct {
	Operator string   `json:"operator"`
	Values   []string `json:"values"`
}

type TargetResourceState struct {
	Operator string   `json:"operator"`
	Values   []string `json:"values"`
}

type TargetResourceGroupState struct {
	Operator string   `json:"operator"`
	Values   []string `json:"values"`
}

type TargetResourceTypeState struct {
	Operator string   `json:"operator"`
	Values   []string `json:"values"`
}

type ScheduleState struct {
	EffectiveFrom  string            `json:"effective_from"`
	EffectiveUntil string            `json:"effective_until"`
	TimeZone       string            `json:"time_zone"`
	Recurrence     []RecurrenceState `json:"recurrence"`
}

type RecurrenceState struct {
	Daily   []DailyState   `json:"daily"`
	Monthly []MonthlyState `json:"monthly"`
	Weekly  []WeeklyState  `json:"weekly"`
}

type DailyState struct {
	EndTime   string `json:"end_time"`
	StartTime string `json:"start_time"`
}

type MonthlyState struct {
	DaysOfMonth []float64 `json:"days_of_month"`
	EndTime     string    `json:"end_time"`
	StartTime   string    `json:"start_time"`
}

type WeeklyState struct {
	DaysOfWeek []string `json:"days_of_week"`
	EndTime    string   `json:"end_time"`
	StartTime  string   `json:"start_time"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
