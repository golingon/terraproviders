// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package monitoraaddiagnosticsetting

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Log struct {
	// Category: string, required
	Category terra.StringValue `hcl:"category,attr" validate:"required"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// RetentionPolicy: required
	RetentionPolicy *RetentionPolicy `hcl:"retention_policy,block" validate:"required"`
}

type RetentionPolicy struct {
	// Days: number, optional
	Days terra.NumberValue `hcl:"days,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
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

func (l LogAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(l.ref.Append("enabled"))
}

func (l LogAttributes) RetentionPolicy() terra.ListValue[RetentionPolicyAttributes] {
	return terra.ReferenceAsList[RetentionPolicyAttributes](l.ref.Append("retention_policy"))
}

type RetentionPolicyAttributes struct {
	ref terra.Reference
}

func (rp RetentionPolicyAttributes) InternalRef() (terra.Reference, error) {
	return rp.ref, nil
}

func (rp RetentionPolicyAttributes) InternalWithRef(ref terra.Reference) RetentionPolicyAttributes {
	return RetentionPolicyAttributes{ref: ref}
}

func (rp RetentionPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rp.ref.InternalTokens()
}

func (rp RetentionPolicyAttributes) Days() terra.NumberValue {
	return terra.ReferenceAsNumber(rp.ref.Append("days"))
}

func (rp RetentionPolicyAttributes) Enabled() terra.BoolValue {
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

type LogState struct {
	Category        string                 `json:"category"`
	Enabled         bool                   `json:"enabled"`
	RetentionPolicy []RetentionPolicyState `json:"retention_policy"`
}

type RetentionPolicyState struct {
	Days    float64 `json:"days"`
	Enabled bool    `json:"enabled"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
