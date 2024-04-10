// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package datacloudquotasquotainfos

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type QuotaInfos struct {
	// DimensionsInfos: min=0
	DimensionsInfos []DimensionsInfos `hcl:"dimensions_infos,block" validate:"min=0"`
	// QuotaIncreaseEligibility: min=0
	QuotaIncreaseEligibility []QuotaIncreaseEligibility `hcl:"quota_increase_eligibility,block" validate:"min=0"`
}

type DimensionsInfos struct {
	// Details: min=0
	Details []Details `hcl:"details,block" validate:"min=0"`
}

type Details struct{}

type QuotaIncreaseEligibility struct{}

type QuotaInfosAttributes struct {
	ref terra.Reference
}

func (qi QuotaInfosAttributes) InternalRef() (terra.Reference, error) {
	return qi.ref, nil
}

func (qi QuotaInfosAttributes) InternalWithRef(ref terra.Reference) QuotaInfosAttributes {
	return QuotaInfosAttributes{ref: ref}
}

func (qi QuotaInfosAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return qi.ref.InternalTokens()
}

func (qi QuotaInfosAttributes) ContainerType() terra.StringValue {
	return terra.ReferenceAsString(qi.ref.Append("container_type"))
}

func (qi QuotaInfosAttributes) Dimensions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](qi.ref.Append("dimensions"))
}

func (qi QuotaInfosAttributes) IsConcurrent() terra.BoolValue {
	return terra.ReferenceAsBool(qi.ref.Append("is_concurrent"))
}

func (qi QuotaInfosAttributes) IsFixed() terra.BoolValue {
	return terra.ReferenceAsBool(qi.ref.Append("is_fixed"))
}

func (qi QuotaInfosAttributes) IsPrecise() terra.BoolValue {
	return terra.ReferenceAsBool(qi.ref.Append("is_precise"))
}

func (qi QuotaInfosAttributes) Metric() terra.StringValue {
	return terra.ReferenceAsString(qi.ref.Append("metric"))
}

func (qi QuotaInfosAttributes) MetricDisplayName() terra.StringValue {
	return terra.ReferenceAsString(qi.ref.Append("metric_display_name"))
}

func (qi QuotaInfosAttributes) MetricUnit() terra.StringValue {
	return terra.ReferenceAsString(qi.ref.Append("metric_unit"))
}

func (qi QuotaInfosAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(qi.ref.Append("name"))
}

func (qi QuotaInfosAttributes) QuotaDisplayName() terra.StringValue {
	return terra.ReferenceAsString(qi.ref.Append("quota_display_name"))
}

func (qi QuotaInfosAttributes) QuotaId() terra.StringValue {
	return terra.ReferenceAsString(qi.ref.Append("quota_id"))
}

func (qi QuotaInfosAttributes) RefreshInterval() terra.StringValue {
	return terra.ReferenceAsString(qi.ref.Append("refresh_interval"))
}

func (qi QuotaInfosAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(qi.ref.Append("service"))
}

func (qi QuotaInfosAttributes) ServiceRequestQuotaUri() terra.StringValue {
	return terra.ReferenceAsString(qi.ref.Append("service_request_quota_uri"))
}

func (qi QuotaInfosAttributes) DimensionsInfos() terra.ListValue[DimensionsInfosAttributes] {
	return terra.ReferenceAsList[DimensionsInfosAttributes](qi.ref.Append("dimensions_infos"))
}

func (qi QuotaInfosAttributes) QuotaIncreaseEligibility() terra.ListValue[QuotaIncreaseEligibilityAttributes] {
	return terra.ReferenceAsList[QuotaIncreaseEligibilityAttributes](qi.ref.Append("quota_increase_eligibility"))
}

type DimensionsInfosAttributes struct {
	ref terra.Reference
}

func (di DimensionsInfosAttributes) InternalRef() (terra.Reference, error) {
	return di.ref, nil
}

func (di DimensionsInfosAttributes) InternalWithRef(ref terra.Reference) DimensionsInfosAttributes {
	return DimensionsInfosAttributes{ref: ref}
}

func (di DimensionsInfosAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return di.ref.InternalTokens()
}

func (di DimensionsInfosAttributes) ApplicableLocations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](di.ref.Append("applicable_locations"))
}

func (di DimensionsInfosAttributes) Dimensions() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](di.ref.Append("dimensions"))
}

func (di DimensionsInfosAttributes) Details() terra.ListValue[DetailsAttributes] {
	return terra.ReferenceAsList[DetailsAttributes](di.ref.Append("details"))
}

type DetailsAttributes struct {
	ref terra.Reference
}

func (d DetailsAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d DetailsAttributes) InternalWithRef(ref terra.Reference) DetailsAttributes {
	return DetailsAttributes{ref: ref}
}

func (d DetailsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d DetailsAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("value"))
}

type QuotaIncreaseEligibilityAttributes struct {
	ref terra.Reference
}

func (qie QuotaIncreaseEligibilityAttributes) InternalRef() (terra.Reference, error) {
	return qie.ref, nil
}

func (qie QuotaIncreaseEligibilityAttributes) InternalWithRef(ref terra.Reference) QuotaIncreaseEligibilityAttributes {
	return QuotaIncreaseEligibilityAttributes{ref: ref}
}

func (qie QuotaIncreaseEligibilityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return qie.ref.InternalTokens()
}

func (qie QuotaIncreaseEligibilityAttributes) IneligibilityReason() terra.StringValue {
	return terra.ReferenceAsString(qie.ref.Append("ineligibility_reason"))
}

func (qie QuotaIncreaseEligibilityAttributes) IsEligible() terra.BoolValue {
	return terra.ReferenceAsBool(qie.ref.Append("is_eligible"))
}

type QuotaInfosState struct {
	ContainerType            string                          `json:"container_type"`
	Dimensions               []string                        `json:"dimensions"`
	IsConcurrent             bool                            `json:"is_concurrent"`
	IsFixed                  bool                            `json:"is_fixed"`
	IsPrecise                bool                            `json:"is_precise"`
	Metric                   string                          `json:"metric"`
	MetricDisplayName        string                          `json:"metric_display_name"`
	MetricUnit               string                          `json:"metric_unit"`
	Name                     string                          `json:"name"`
	QuotaDisplayName         string                          `json:"quota_display_name"`
	QuotaId                  string                          `json:"quota_id"`
	RefreshInterval          string                          `json:"refresh_interval"`
	Service                  string                          `json:"service"`
	ServiceRequestQuotaUri   string                          `json:"service_request_quota_uri"`
	DimensionsInfos          []DimensionsInfosState          `json:"dimensions_infos"`
	QuotaIncreaseEligibility []QuotaIncreaseEligibilityState `json:"quota_increase_eligibility"`
}

type DimensionsInfosState struct {
	ApplicableLocations []string          `json:"applicable_locations"`
	Dimensions          map[string]string `json:"dimensions"`
	Details             []DetailsState    `json:"details"`
}

type DetailsState struct {
	Value string `json:"value"`
}

type QuotaIncreaseEligibilityState struct {
	IneligibilityReason string `json:"ineligibility_reason"`
	IsEligible          bool   `json:"is_eligible"`
}
