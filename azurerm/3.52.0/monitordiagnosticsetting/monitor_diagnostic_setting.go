// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package monitordiagnosticsetting

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type EnabledLog struct {
	// Category: string, optional
	Category terra.StringValue `hcl:"category,attr"`
	// CategoryGroup: string, optional
	CategoryGroup terra.StringValue `hcl:"category_group,attr"`
	// EnabledLogRetentionPolicy: optional
	RetentionPolicy *EnabledLogRetentionPolicy `hcl:"retention_policy,block"`
}

type EnabledLogRetentionPolicy struct {
	// Days: number, optional
	Days terra.NumberValue `hcl:"days,attr"`
	// Enabled: bool, required
	Enabled terra.BoolValue `hcl:"enabled,attr" validate:"required"`
}

type Log struct {
	// Category: string, optional
	Category terra.StringValue `hcl:"category,attr"`
	// CategoryGroup: string, optional
	CategoryGroup terra.StringValue `hcl:"category_group,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// LogRetentionPolicy: optional
	RetentionPolicy *LogRetentionPolicy `hcl:"retention_policy,block"`
}

type LogRetentionPolicy struct {
	// Days: number, optional
	Days terra.NumberValue `hcl:"days,attr"`
	// Enabled: bool, required
	Enabled terra.BoolValue `hcl:"enabled,attr" validate:"required"`
}

type Metric struct {
	// Category: string, required
	Category terra.StringValue `hcl:"category,attr" validate:"required"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// MetricRetentionPolicy: optional
	RetentionPolicy *MetricRetentionPolicy `hcl:"retention_policy,block"`
}

type MetricRetentionPolicy struct {
	// Days: number, optional
	Days terra.NumberValue `hcl:"days,attr"`
	// Enabled: bool, required
	Enabled terra.BoolValue `hcl:"enabled,attr" validate:"required"`
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

type EnabledLogAttributes struct {
	ref terra.Reference
}

func (el EnabledLogAttributes) InternalRef() (terra.Reference, error) {
	return el.ref, nil
}

func (el EnabledLogAttributes) InternalWithRef(ref terra.Reference) EnabledLogAttributes {
	return EnabledLogAttributes{ref: ref}
}

func (el EnabledLogAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return el.ref.InternalTokens()
}

func (el EnabledLogAttributes) Category() terra.StringValue {
	return terra.ReferenceAsString(el.ref.Append("category"))
}

func (el EnabledLogAttributes) CategoryGroup() terra.StringValue {
	return terra.ReferenceAsString(el.ref.Append("category_group"))
}

func (el EnabledLogAttributes) RetentionPolicy() terra.ListValue[EnabledLogRetentionPolicyAttributes] {
	return terra.ReferenceAsList[EnabledLogRetentionPolicyAttributes](el.ref.Append("retention_policy"))
}

type EnabledLogRetentionPolicyAttributes struct {
	ref terra.Reference
}

func (rp EnabledLogRetentionPolicyAttributes) InternalRef() (terra.Reference, error) {
	return rp.ref, nil
}

func (rp EnabledLogRetentionPolicyAttributes) InternalWithRef(ref terra.Reference) EnabledLogRetentionPolicyAttributes {
	return EnabledLogRetentionPolicyAttributes{ref: ref}
}

func (rp EnabledLogRetentionPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rp.ref.InternalTokens()
}

func (rp EnabledLogRetentionPolicyAttributes) Days() terra.NumberValue {
	return terra.ReferenceAsNumber(rp.ref.Append("days"))
}

func (rp EnabledLogRetentionPolicyAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(rp.ref.Append("enabled"))
}

type LogAttributes struct {
	ref terra.Reference
}

func (l LogAttributes) InternalRef() (terra.Reference, error) {
	return l.ref, nil
}

func (l LogAttributes) InternalWithRef(ref terra.Reference) LogAttributes {
	return LogAttributes{ref: ref}
}

func (l LogAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return l.ref.InternalTokens()
}

func (l LogAttributes) Category() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("category"))
}

func (l LogAttributes) CategoryGroup() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("category_group"))
}

func (l LogAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(l.ref.Append("enabled"))
}

func (l LogAttributes) RetentionPolicy() terra.ListValue[LogRetentionPolicyAttributes] {
	return terra.ReferenceAsList[LogRetentionPolicyAttributes](l.ref.Append("retention_policy"))
}

type LogRetentionPolicyAttributes struct {
	ref terra.Reference
}

func (rp LogRetentionPolicyAttributes) InternalRef() (terra.Reference, error) {
	return rp.ref, nil
}

func (rp LogRetentionPolicyAttributes) InternalWithRef(ref terra.Reference) LogRetentionPolicyAttributes {
	return LogRetentionPolicyAttributes{ref: ref}
}

func (rp LogRetentionPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rp.ref.InternalTokens()
}

func (rp LogRetentionPolicyAttributes) Days() terra.NumberValue {
	return terra.ReferenceAsNumber(rp.ref.Append("days"))
}

func (rp LogRetentionPolicyAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(rp.ref.Append("enabled"))
}

type MetricAttributes struct {
	ref terra.Reference
}

func (m MetricAttributes) InternalRef() (terra.Reference, error) {
	return m.ref, nil
}

func (m MetricAttributes) InternalWithRef(ref terra.Reference) MetricAttributes {
	return MetricAttributes{ref: ref}
}

func (m MetricAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return m.ref.InternalTokens()
}

func (m MetricAttributes) Category() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("category"))
}

func (m MetricAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(m.ref.Append("enabled"))
}

func (m MetricAttributes) RetentionPolicy() terra.ListValue[MetricRetentionPolicyAttributes] {
	return terra.ReferenceAsList[MetricRetentionPolicyAttributes](m.ref.Append("retention_policy"))
}

type MetricRetentionPolicyAttributes struct {
	ref terra.Reference
}

func (rp MetricRetentionPolicyAttributes) InternalRef() (terra.Reference, error) {
	return rp.ref, nil
}

func (rp MetricRetentionPolicyAttributes) InternalWithRef(ref terra.Reference) MetricRetentionPolicyAttributes {
	return MetricRetentionPolicyAttributes{ref: ref}
}

func (rp MetricRetentionPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rp.ref.InternalTokens()
}

func (rp MetricRetentionPolicyAttributes) Days() terra.NumberValue {
	return terra.ReferenceAsNumber(rp.ref.Append("days"))
}

func (rp MetricRetentionPolicyAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(rp.ref.Append("enabled"))
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

type EnabledLogState struct {
	Category        string                           `json:"category"`
	CategoryGroup   string                           `json:"category_group"`
	RetentionPolicy []EnabledLogRetentionPolicyState `json:"retention_policy"`
}

type EnabledLogRetentionPolicyState struct {
	Days    float64 `json:"days"`
	Enabled bool    `json:"enabled"`
}

type LogState struct {
	Category        string                    `json:"category"`
	CategoryGroup   string                    `json:"category_group"`
	Enabled         bool                      `json:"enabled"`
	RetentionPolicy []LogRetentionPolicyState `json:"retention_policy"`
}

type LogRetentionPolicyState struct {
	Days    float64 `json:"days"`
	Enabled bool    `json:"enabled"`
}

type MetricState struct {
	Category        string                       `json:"category"`
	Enabled         bool                         `json:"enabled"`
	RetentionPolicy []MetricRetentionPolicyState `json:"retention_policy"`
}

type MetricRetentionPolicyState struct {
	Days    float64 `json:"days"`
	Enabled bool    `json:"enabled"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}