// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package consumptionbudgetresourcegroup

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Filter struct {
	// FilterDimension: min=0
	Dimension []FilterDimension `hcl:"dimension,block" validate:"min=0"`
	// Not: optional
	Not *Not `hcl:"not,block"`
	// FilterTag: min=0
	Tag []FilterTag `hcl:"tag,block" validate:"min=0"`
}

type FilterDimension struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Operator: string, optional
	Operator terra.StringValue `hcl:"operator,attr"`
	// Values: list of string, required
	Values terra.ListValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type Not struct {
	// NotDimension: optional
	Dimension *NotDimension `hcl:"dimension,block"`
	// NotTag: optional
	Tag *NotTag `hcl:"tag,block"`
}

type NotDimension struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Operator: string, optional
	Operator terra.StringValue `hcl:"operator,attr"`
	// Values: list of string, required
	Values terra.ListValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type NotTag struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Operator: string, optional
	Operator terra.StringValue `hcl:"operator,attr"`
	// Values: list of string, required
	Values terra.ListValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type FilterTag struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Operator: string, optional
	Operator terra.StringValue `hcl:"operator,attr"`
	// Values: list of string, required
	Values terra.ListValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type Notification struct {
	// ContactEmails: list of string, optional
	ContactEmails terra.ListValue[terra.StringValue] `hcl:"contact_emails,attr"`
	// ContactGroups: list of string, optional
	ContactGroups terra.ListValue[terra.StringValue] `hcl:"contact_groups,attr"`
	// ContactRoles: list of string, optional
	ContactRoles terra.ListValue[terra.StringValue] `hcl:"contact_roles,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Operator: string, required
	Operator terra.StringValue `hcl:"operator,attr" validate:"required"`
	// Threshold: number, required
	Threshold terra.NumberValue `hcl:"threshold,attr" validate:"required"`
	// ThresholdType: string, optional
	ThresholdType terra.StringValue `hcl:"threshold_type,attr"`
}

type TimePeriod struct {
	// EndDate: string, optional
	EndDate terra.StringValue `hcl:"end_date,attr"`
	// StartDate: string, required
	StartDate terra.StringValue `hcl:"start_date,attr" validate:"required"`
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

type FilterAttributes struct {
	ref terra.Reference
}

func (f FilterAttributes) InternalRef() terra.Reference {
	return f.ref
}

func (f FilterAttributes) InternalWithRef(ref terra.Reference) FilterAttributes {
	return FilterAttributes{ref: ref}
}

func (f FilterAttributes) InternalTokens() hclwrite.Tokens {
	return f.ref.InternalTokens()
}

func (f FilterAttributes) Dimension() terra.SetValue[FilterDimensionAttributes] {
	return terra.ReferenceSet[FilterDimensionAttributes](f.ref.Append("dimension"))
}

func (f FilterAttributes) Not() terra.ListValue[NotAttributes] {
	return terra.ReferenceList[NotAttributes](f.ref.Append("not"))
}

func (f FilterAttributes) Tag() terra.SetValue[FilterTagAttributes] {
	return terra.ReferenceSet[FilterTagAttributes](f.ref.Append("tag"))
}

type FilterDimensionAttributes struct {
	ref terra.Reference
}

func (d FilterDimensionAttributes) InternalRef() terra.Reference {
	return d.ref
}

func (d FilterDimensionAttributes) InternalWithRef(ref terra.Reference) FilterDimensionAttributes {
	return FilterDimensionAttributes{ref: ref}
}

func (d FilterDimensionAttributes) InternalTokens() hclwrite.Tokens {
	return d.ref.InternalTokens()
}

func (d FilterDimensionAttributes) Name() terra.StringValue {
	return terra.ReferenceString(d.ref.Append("name"))
}

func (d FilterDimensionAttributes) Operator() terra.StringValue {
	return terra.ReferenceString(d.ref.Append("operator"))
}

func (d FilterDimensionAttributes) Values() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](d.ref.Append("values"))
}

type NotAttributes struct {
	ref terra.Reference
}

func (n NotAttributes) InternalRef() terra.Reference {
	return n.ref
}

func (n NotAttributes) InternalWithRef(ref terra.Reference) NotAttributes {
	return NotAttributes{ref: ref}
}

func (n NotAttributes) InternalTokens() hclwrite.Tokens {
	return n.ref.InternalTokens()
}

func (n NotAttributes) Dimension() terra.ListValue[NotDimensionAttributes] {
	return terra.ReferenceList[NotDimensionAttributes](n.ref.Append("dimension"))
}

func (n NotAttributes) Tag() terra.ListValue[NotTagAttributes] {
	return terra.ReferenceList[NotTagAttributes](n.ref.Append("tag"))
}

type NotDimensionAttributes struct {
	ref terra.Reference
}

func (d NotDimensionAttributes) InternalRef() terra.Reference {
	return d.ref
}

func (d NotDimensionAttributes) InternalWithRef(ref terra.Reference) NotDimensionAttributes {
	return NotDimensionAttributes{ref: ref}
}

func (d NotDimensionAttributes) InternalTokens() hclwrite.Tokens {
	return d.ref.InternalTokens()
}

func (d NotDimensionAttributes) Name() terra.StringValue {
	return terra.ReferenceString(d.ref.Append("name"))
}

func (d NotDimensionAttributes) Operator() terra.StringValue {
	return terra.ReferenceString(d.ref.Append("operator"))
}

func (d NotDimensionAttributes) Values() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](d.ref.Append("values"))
}

type NotTagAttributes struct {
	ref terra.Reference
}

func (t NotTagAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t NotTagAttributes) InternalWithRef(ref terra.Reference) NotTagAttributes {
	return NotTagAttributes{ref: ref}
}

func (t NotTagAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t NotTagAttributes) Name() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("name"))
}

func (t NotTagAttributes) Operator() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("operator"))
}

func (t NotTagAttributes) Values() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](t.ref.Append("values"))
}

type FilterTagAttributes struct {
	ref terra.Reference
}

func (t FilterTagAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t FilterTagAttributes) InternalWithRef(ref terra.Reference) FilterTagAttributes {
	return FilterTagAttributes{ref: ref}
}

func (t FilterTagAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t FilterTagAttributes) Name() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("name"))
}

func (t FilterTagAttributes) Operator() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("operator"))
}

func (t FilterTagAttributes) Values() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](t.ref.Append("values"))
}

type NotificationAttributes struct {
	ref terra.Reference
}

func (n NotificationAttributes) InternalRef() terra.Reference {
	return n.ref
}

func (n NotificationAttributes) InternalWithRef(ref terra.Reference) NotificationAttributes {
	return NotificationAttributes{ref: ref}
}

func (n NotificationAttributes) InternalTokens() hclwrite.Tokens {
	return n.ref.InternalTokens()
}

func (n NotificationAttributes) ContactEmails() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](n.ref.Append("contact_emails"))
}

func (n NotificationAttributes) ContactGroups() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](n.ref.Append("contact_groups"))
}

func (n NotificationAttributes) ContactRoles() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](n.ref.Append("contact_roles"))
}

func (n NotificationAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceBool(n.ref.Append("enabled"))
}

func (n NotificationAttributes) Operator() terra.StringValue {
	return terra.ReferenceString(n.ref.Append("operator"))
}

func (n NotificationAttributes) Threshold() terra.NumberValue {
	return terra.ReferenceNumber(n.ref.Append("threshold"))
}

func (n NotificationAttributes) ThresholdType() terra.StringValue {
	return terra.ReferenceString(n.ref.Append("threshold_type"))
}

type TimePeriodAttributes struct {
	ref terra.Reference
}

func (tp TimePeriodAttributes) InternalRef() terra.Reference {
	return tp.ref
}

func (tp TimePeriodAttributes) InternalWithRef(ref terra.Reference) TimePeriodAttributes {
	return TimePeriodAttributes{ref: ref}
}

func (tp TimePeriodAttributes) InternalTokens() hclwrite.Tokens {
	return tp.ref.InternalTokens()
}

func (tp TimePeriodAttributes) EndDate() terra.StringValue {
	return terra.ReferenceString(tp.ref.Append("end_date"))
}

func (tp TimePeriodAttributes) StartDate() terra.StringValue {
	return terra.ReferenceString(tp.ref.Append("start_date"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("update"))
}

type FilterState struct {
	Dimension []FilterDimensionState `json:"dimension"`
	Not       []NotState             `json:"not"`
	Tag       []FilterTagState       `json:"tag"`
}

type FilterDimensionState struct {
	Name     string   `json:"name"`
	Operator string   `json:"operator"`
	Values   []string `json:"values"`
}

type NotState struct {
	Dimension []NotDimensionState `json:"dimension"`
	Tag       []NotTagState       `json:"tag"`
}

type NotDimensionState struct {
	Name     string   `json:"name"`
	Operator string   `json:"operator"`
	Values   []string `json:"values"`
}

type NotTagState struct {
	Name     string   `json:"name"`
	Operator string   `json:"operator"`
	Values   []string `json:"values"`
}

type FilterTagState struct {
	Name     string   `json:"name"`
	Operator string   `json:"operator"`
	Values   []string `json:"values"`
}

type NotificationState struct {
	ContactEmails []string `json:"contact_emails"`
	ContactGroups []string `json:"contact_groups"`
	ContactRoles  []string `json:"contact_roles"`
	Enabled       bool     `json:"enabled"`
	Operator      string   `json:"operator"`
	Threshold     float64  `json:"threshold"`
	ThresholdType string   `json:"threshold_type"`
}

type TimePeriodState struct {
	EndDate   string `json:"end_date"`
	StartDate string `json:"start_date"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
