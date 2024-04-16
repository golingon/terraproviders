// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_cloud_quotas_quota_preference

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type QuotaConfig struct {
	// Annotations: map of string, optional
	Annotations terra.MapValue[terra.StringValue] `hcl:"annotations,attr"`
	// PreferredValue: string, required
	PreferredValue terra.StringValue `hcl:"preferred_value,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type QuotaConfigAttributes struct {
	ref terra.Reference
}

func (qc QuotaConfigAttributes) InternalRef() (terra.Reference, error) {
	return qc.ref, nil
}

func (qc QuotaConfigAttributes) InternalWithRef(ref terra.Reference) QuotaConfigAttributes {
	return QuotaConfigAttributes{ref: ref}
}

func (qc QuotaConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return qc.ref.InternalTokens()
}

func (qc QuotaConfigAttributes) Annotations() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](qc.ref.Append("annotations"))
}

func (qc QuotaConfigAttributes) GrantedValue() terra.StringValue {
	return terra.ReferenceAsString(qc.ref.Append("granted_value"))
}

func (qc QuotaConfigAttributes) PreferredValue() terra.StringValue {
	return terra.ReferenceAsString(qc.ref.Append("preferred_value"))
}

func (qc QuotaConfigAttributes) RequestOrigin() terra.StringValue {
	return terra.ReferenceAsString(qc.ref.Append("request_origin"))
}

func (qc QuotaConfigAttributes) StateDetail() terra.StringValue {
	return terra.ReferenceAsString(qc.ref.Append("state_detail"))
}

func (qc QuotaConfigAttributes) TraceId() terra.StringValue {
	return terra.ReferenceAsString(qc.ref.Append("trace_id"))
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

type QuotaConfigState struct {
	Annotations    map[string]string `json:"annotations"`
	GrantedValue   string            `json:"granted_value"`
	PreferredValue string            `json:"preferred_value"`
	RequestOrigin  string            `json:"request_origin"`
	StateDetail    string            `json:"state_detail"`
	TraceId        string            `json:"trace_id"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
