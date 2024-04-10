// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package cloudtrail

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type AdvancedEventSelector struct {
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// FieldSelector: min=1
	FieldSelector []FieldSelector `hcl:"field_selector,block" validate:"min=1"`
}

type FieldSelector struct {
	// EndsWith: list of string, optional
	EndsWith terra.ListValue[terra.StringValue] `hcl:"ends_with,attr"`
	// Equals: list of string, optional
	Equals terra.ListValue[terra.StringValue] `hcl:"equals,attr"`
	// Field: string, required
	Field terra.StringValue `hcl:"field,attr" validate:"required"`
	// NotEndsWith: list of string, optional
	NotEndsWith terra.ListValue[terra.StringValue] `hcl:"not_ends_with,attr"`
	// NotEquals: list of string, optional
	NotEquals terra.ListValue[terra.StringValue] `hcl:"not_equals,attr"`
	// NotStartsWith: list of string, optional
	NotStartsWith terra.ListValue[terra.StringValue] `hcl:"not_starts_with,attr"`
	// StartsWith: list of string, optional
	StartsWith terra.ListValue[terra.StringValue] `hcl:"starts_with,attr"`
}

type EventSelector struct {
	// ExcludeManagementEventSources: set of string, optional
	ExcludeManagementEventSources terra.SetValue[terra.StringValue] `hcl:"exclude_management_event_sources,attr"`
	// IncludeManagementEvents: bool, optional
	IncludeManagementEvents terra.BoolValue `hcl:"include_management_events,attr"`
	// ReadWriteType: string, optional
	ReadWriteType terra.StringValue `hcl:"read_write_type,attr"`
	// DataResource: min=0
	DataResource []DataResource `hcl:"data_resource,block" validate:"min=0"`
}

type DataResource struct {
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Values: list of string, required
	Values terra.ListValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type InsightSelector struct {
	// InsightType: string, required
	InsightType terra.StringValue `hcl:"insight_type,attr" validate:"required"`
}

type AdvancedEventSelectorAttributes struct {
	ref terra.Reference
}

func (aes AdvancedEventSelectorAttributes) InternalRef() (terra.Reference, error) {
	return aes.ref, nil
}

func (aes AdvancedEventSelectorAttributes) InternalWithRef(ref terra.Reference) AdvancedEventSelectorAttributes {
	return AdvancedEventSelectorAttributes{ref: ref}
}

func (aes AdvancedEventSelectorAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return aes.ref.InternalTokens()
}

func (aes AdvancedEventSelectorAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aes.ref.Append("name"))
}

func (aes AdvancedEventSelectorAttributes) FieldSelector() terra.SetValue[FieldSelectorAttributes] {
	return terra.ReferenceAsSet[FieldSelectorAttributes](aes.ref.Append("field_selector"))
}

type FieldSelectorAttributes struct {
	ref terra.Reference
}

func (fs FieldSelectorAttributes) InternalRef() (terra.Reference, error) {
	return fs.ref, nil
}

func (fs FieldSelectorAttributes) InternalWithRef(ref terra.Reference) FieldSelectorAttributes {
	return FieldSelectorAttributes{ref: ref}
}

func (fs FieldSelectorAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fs.ref.InternalTokens()
}

func (fs FieldSelectorAttributes) EndsWith() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](fs.ref.Append("ends_with"))
}

func (fs FieldSelectorAttributes) Equals() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](fs.ref.Append("equals"))
}

func (fs FieldSelectorAttributes) Field() terra.StringValue {
	return terra.ReferenceAsString(fs.ref.Append("field"))
}

func (fs FieldSelectorAttributes) NotEndsWith() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](fs.ref.Append("not_ends_with"))
}

func (fs FieldSelectorAttributes) NotEquals() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](fs.ref.Append("not_equals"))
}

func (fs FieldSelectorAttributes) NotStartsWith() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](fs.ref.Append("not_starts_with"))
}

func (fs FieldSelectorAttributes) StartsWith() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](fs.ref.Append("starts_with"))
}

type EventSelectorAttributes struct {
	ref terra.Reference
}

func (es EventSelectorAttributes) InternalRef() (terra.Reference, error) {
	return es.ref, nil
}

func (es EventSelectorAttributes) InternalWithRef(ref terra.Reference) EventSelectorAttributes {
	return EventSelectorAttributes{ref: ref}
}

func (es EventSelectorAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return es.ref.InternalTokens()
}

func (es EventSelectorAttributes) ExcludeManagementEventSources() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](es.ref.Append("exclude_management_event_sources"))
}

func (es EventSelectorAttributes) IncludeManagementEvents() terra.BoolValue {
	return terra.ReferenceAsBool(es.ref.Append("include_management_events"))
}

func (es EventSelectorAttributes) ReadWriteType() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("read_write_type"))
}

func (es EventSelectorAttributes) DataResource() terra.ListValue[DataResourceAttributes] {
	return terra.ReferenceAsList[DataResourceAttributes](es.ref.Append("data_resource"))
}

type DataResourceAttributes struct {
	ref terra.Reference
}

func (dr DataResourceAttributes) InternalRef() (terra.Reference, error) {
	return dr.ref, nil
}

func (dr DataResourceAttributes) InternalWithRef(ref terra.Reference) DataResourceAttributes {
	return DataResourceAttributes{ref: ref}
}

func (dr DataResourceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dr.ref.InternalTokens()
}

func (dr DataResourceAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(dr.ref.Append("type"))
}

func (dr DataResourceAttributes) Values() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dr.ref.Append("values"))
}

type InsightSelectorAttributes struct {
	ref terra.Reference
}

func (is InsightSelectorAttributes) InternalRef() (terra.Reference, error) {
	return is.ref, nil
}

func (is InsightSelectorAttributes) InternalWithRef(ref terra.Reference) InsightSelectorAttributes {
	return InsightSelectorAttributes{ref: ref}
}

func (is InsightSelectorAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return is.ref.InternalTokens()
}

func (is InsightSelectorAttributes) InsightType() terra.StringValue {
	return terra.ReferenceAsString(is.ref.Append("insight_type"))
}

type AdvancedEventSelectorState struct {
	Name          string               `json:"name"`
	FieldSelector []FieldSelectorState `json:"field_selector"`
}

type FieldSelectorState struct {
	EndsWith      []string `json:"ends_with"`
	Equals        []string `json:"equals"`
	Field         string   `json:"field"`
	NotEndsWith   []string `json:"not_ends_with"`
	NotEquals     []string `json:"not_equals"`
	NotStartsWith []string `json:"not_starts_with"`
	StartsWith    []string `json:"starts_with"`
}

type EventSelectorState struct {
	ExcludeManagementEventSources []string            `json:"exclude_management_event_sources"`
	IncludeManagementEvents       bool                `json:"include_management_events"`
	ReadWriteType                 string              `json:"read_write_type"`
	DataResource                  []DataResourceState `json:"data_resource"`
}

type DataResourceState struct {
	Type   string   `json:"type"`
	Values []string `json:"values"`
}

type InsightSelectorState struct {
	InsightType string `json:"insight_type"`
}
