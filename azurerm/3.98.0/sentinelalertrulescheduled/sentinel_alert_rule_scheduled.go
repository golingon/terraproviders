// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package sentinelalertrulescheduled

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type AlertDetailsOverride struct {
	// DescriptionFormat: string, optional
	DescriptionFormat terra.StringValue `hcl:"description_format,attr"`
	// DisplayNameFormat: string, optional
	DisplayNameFormat terra.StringValue `hcl:"display_name_format,attr"`
	// SeverityColumnName: string, optional
	SeverityColumnName terra.StringValue `hcl:"severity_column_name,attr"`
	// TacticsColumnName: string, optional
	TacticsColumnName terra.StringValue `hcl:"tactics_column_name,attr"`
	// DynamicProperty: min=0
	DynamicProperty []DynamicProperty `hcl:"dynamic_property,block" validate:"min=0"`
}

type DynamicProperty struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type EntityMapping struct {
	// EntityType: string, required
	EntityType terra.StringValue `hcl:"entity_type,attr" validate:"required"`
	// FieldMapping: min=1,max=3
	FieldMapping []FieldMapping `hcl:"field_mapping,block" validate:"min=1,max=3"`
}

type FieldMapping struct {
	// ColumnName: string, required
	ColumnName terra.StringValue `hcl:"column_name,attr" validate:"required"`
	// Identifier: string, required
	Identifier terra.StringValue `hcl:"identifier,attr" validate:"required"`
}

type EventGrouping struct {
	// AggregationMethod: string, required
	AggregationMethod terra.StringValue `hcl:"aggregation_method,attr" validate:"required"`
}

type IncidentConfiguration struct {
	// CreateIncident: bool, required
	CreateIncident terra.BoolValue `hcl:"create_incident,attr" validate:"required"`
	// Grouping: required
	Grouping *Grouping `hcl:"grouping,block" validate:"required"`
}

type Grouping struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// EntityMatchingMethod: string, optional
	EntityMatchingMethod terra.StringValue `hcl:"entity_matching_method,attr"`
	// GroupByAlertDetails: list of string, optional
	GroupByAlertDetails terra.ListValue[terra.StringValue] `hcl:"group_by_alert_details,attr"`
	// GroupByCustomDetails: list of string, optional
	GroupByCustomDetails terra.ListValue[terra.StringValue] `hcl:"group_by_custom_details,attr"`
	// GroupByEntities: list of string, optional
	GroupByEntities terra.ListValue[terra.StringValue] `hcl:"group_by_entities,attr"`
	// LookbackDuration: string, optional
	LookbackDuration terra.StringValue `hcl:"lookback_duration,attr"`
	// ReopenClosedIncidents: bool, optional
	ReopenClosedIncidents terra.BoolValue `hcl:"reopen_closed_incidents,attr"`
}

type SentinelEntityMapping struct {
	// ColumnName: string, required
	ColumnName terra.StringValue `hcl:"column_name,attr" validate:"required"`
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

type AlertDetailsOverrideAttributes struct {
	ref terra.Reference
}

func (ado AlertDetailsOverrideAttributes) InternalRef() (terra.Reference, error) {
	return ado.ref, nil
}

func (ado AlertDetailsOverrideAttributes) InternalWithRef(ref terra.Reference) AlertDetailsOverrideAttributes {
	return AlertDetailsOverrideAttributes{ref: ref}
}

func (ado AlertDetailsOverrideAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ado.ref.InternalTokens()
}

func (ado AlertDetailsOverrideAttributes) DescriptionFormat() terra.StringValue {
	return terra.ReferenceAsString(ado.ref.Append("description_format"))
}

func (ado AlertDetailsOverrideAttributes) DisplayNameFormat() terra.StringValue {
	return terra.ReferenceAsString(ado.ref.Append("display_name_format"))
}

func (ado AlertDetailsOverrideAttributes) SeverityColumnName() terra.StringValue {
	return terra.ReferenceAsString(ado.ref.Append("severity_column_name"))
}

func (ado AlertDetailsOverrideAttributes) TacticsColumnName() terra.StringValue {
	return terra.ReferenceAsString(ado.ref.Append("tactics_column_name"))
}

func (ado AlertDetailsOverrideAttributes) DynamicProperty() terra.ListValue[DynamicPropertyAttributes] {
	return terra.ReferenceAsList[DynamicPropertyAttributes](ado.ref.Append("dynamic_property"))
}

type DynamicPropertyAttributes struct {
	ref terra.Reference
}

func (dp DynamicPropertyAttributes) InternalRef() (terra.Reference, error) {
	return dp.ref, nil
}

func (dp DynamicPropertyAttributes) InternalWithRef(ref terra.Reference) DynamicPropertyAttributes {
	return DynamicPropertyAttributes{ref: ref}
}

func (dp DynamicPropertyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dp.ref.InternalTokens()
}

func (dp DynamicPropertyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("name"))
}

func (dp DynamicPropertyAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("value"))
}

type EntityMappingAttributes struct {
	ref terra.Reference
}

func (em EntityMappingAttributes) InternalRef() (terra.Reference, error) {
	return em.ref, nil
}

func (em EntityMappingAttributes) InternalWithRef(ref terra.Reference) EntityMappingAttributes {
	return EntityMappingAttributes{ref: ref}
}

func (em EntityMappingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return em.ref.InternalTokens()
}

func (em EntityMappingAttributes) EntityType() terra.StringValue {
	return terra.ReferenceAsString(em.ref.Append("entity_type"))
}

func (em EntityMappingAttributes) FieldMapping() terra.ListValue[FieldMappingAttributes] {
	return terra.ReferenceAsList[FieldMappingAttributes](em.ref.Append("field_mapping"))
}

type FieldMappingAttributes struct {
	ref terra.Reference
}

func (fm FieldMappingAttributes) InternalRef() (terra.Reference, error) {
	return fm.ref, nil
}

func (fm FieldMappingAttributes) InternalWithRef(ref terra.Reference) FieldMappingAttributes {
	return FieldMappingAttributes{ref: ref}
}

func (fm FieldMappingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fm.ref.InternalTokens()
}

func (fm FieldMappingAttributes) ColumnName() terra.StringValue {
	return terra.ReferenceAsString(fm.ref.Append("column_name"))
}

func (fm FieldMappingAttributes) Identifier() terra.StringValue {
	return terra.ReferenceAsString(fm.ref.Append("identifier"))
}

type EventGroupingAttributes struct {
	ref terra.Reference
}

func (eg EventGroupingAttributes) InternalRef() (terra.Reference, error) {
	return eg.ref, nil
}

func (eg EventGroupingAttributes) InternalWithRef(ref terra.Reference) EventGroupingAttributes {
	return EventGroupingAttributes{ref: ref}
}

func (eg EventGroupingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return eg.ref.InternalTokens()
}

func (eg EventGroupingAttributes) AggregationMethod() terra.StringValue {
	return terra.ReferenceAsString(eg.ref.Append("aggregation_method"))
}

type IncidentConfigurationAttributes struct {
	ref terra.Reference
}

func (ic IncidentConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return ic.ref, nil
}

func (ic IncidentConfigurationAttributes) InternalWithRef(ref terra.Reference) IncidentConfigurationAttributes {
	return IncidentConfigurationAttributes{ref: ref}
}

func (ic IncidentConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ic.ref.InternalTokens()
}

func (ic IncidentConfigurationAttributes) CreateIncident() terra.BoolValue {
	return terra.ReferenceAsBool(ic.ref.Append("create_incident"))
}

func (ic IncidentConfigurationAttributes) Grouping() terra.ListValue[GroupingAttributes] {
	return terra.ReferenceAsList[GroupingAttributes](ic.ref.Append("grouping"))
}

type GroupingAttributes struct {
	ref terra.Reference
}

func (g GroupingAttributes) InternalRef() (terra.Reference, error) {
	return g.ref, nil
}

func (g GroupingAttributes) InternalWithRef(ref terra.Reference) GroupingAttributes {
	return GroupingAttributes{ref: ref}
}

func (g GroupingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return g.ref.InternalTokens()
}

func (g GroupingAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(g.ref.Append("enabled"))
}

func (g GroupingAttributes) EntityMatchingMethod() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("entity_matching_method"))
}

func (g GroupingAttributes) GroupByAlertDetails() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](g.ref.Append("group_by_alert_details"))
}

func (g GroupingAttributes) GroupByCustomDetails() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](g.ref.Append("group_by_custom_details"))
}

func (g GroupingAttributes) GroupByEntities() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](g.ref.Append("group_by_entities"))
}

func (g GroupingAttributes) LookbackDuration() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("lookback_duration"))
}

func (g GroupingAttributes) ReopenClosedIncidents() terra.BoolValue {
	return terra.ReferenceAsBool(g.ref.Append("reopen_closed_incidents"))
}

type SentinelEntityMappingAttributes struct {
	ref terra.Reference
}

func (sem SentinelEntityMappingAttributes) InternalRef() (terra.Reference, error) {
	return sem.ref, nil
}

func (sem SentinelEntityMappingAttributes) InternalWithRef(ref terra.Reference) SentinelEntityMappingAttributes {
	return SentinelEntityMappingAttributes{ref: ref}
}

func (sem SentinelEntityMappingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sem.ref.InternalTokens()
}

func (sem SentinelEntityMappingAttributes) ColumnName() terra.StringValue {
	return terra.ReferenceAsString(sem.ref.Append("column_name"))
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

type AlertDetailsOverrideState struct {
	DescriptionFormat  string                 `json:"description_format"`
	DisplayNameFormat  string                 `json:"display_name_format"`
	SeverityColumnName string                 `json:"severity_column_name"`
	TacticsColumnName  string                 `json:"tactics_column_name"`
	DynamicProperty    []DynamicPropertyState `json:"dynamic_property"`
}

type DynamicPropertyState struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type EntityMappingState struct {
	EntityType   string              `json:"entity_type"`
	FieldMapping []FieldMappingState `json:"field_mapping"`
}

type FieldMappingState struct {
	ColumnName string `json:"column_name"`
	Identifier string `json:"identifier"`
}

type EventGroupingState struct {
	AggregationMethod string `json:"aggregation_method"`
}

type IncidentConfigurationState struct {
	CreateIncident bool            `json:"create_incident"`
	Grouping       []GroupingState `json:"grouping"`
}

type GroupingState struct {
	Enabled               bool     `json:"enabled"`
	EntityMatchingMethod  string   `json:"entity_matching_method"`
	GroupByAlertDetails   []string `json:"group_by_alert_details"`
	GroupByCustomDetails  []string `json:"group_by_custom_details"`
	GroupByEntities       []string `json:"group_by_entities"`
	LookbackDuration      string   `json:"lookback_duration"`
	ReopenClosedIncidents bool     `json:"reopen_closed_incidents"`
}

type SentinelEntityMappingState struct {
	ColumnName string `json:"column_name"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
