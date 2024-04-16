// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_cloudfront_continuous_deployment_policy

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type StagingDistributionDnsNames struct {
	// Items: set of string, optional
	Items terra.SetValue[terra.StringValue] `hcl:"items,attr"`
	// Quantity: number, required
	Quantity terra.NumberValue `hcl:"quantity,attr" validate:"required"`
}

type TrafficConfig struct {
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// TrafficConfigSingleHeaderConfig: min=0
	SingleHeaderConfig []TrafficConfigSingleHeaderConfig `hcl:"single_header_config,block" validate:"min=0"`
	// TrafficConfigSingleWeightConfig: min=0
	SingleWeightConfig []TrafficConfigSingleWeightConfig `hcl:"single_weight_config,block" validate:"min=0"`
}

type TrafficConfigSingleHeaderConfig struct {
	// Header: string, required
	Header terra.StringValue `hcl:"header,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type TrafficConfigSingleWeightConfig struct {
	// Weight: number, required
	Weight terra.NumberValue `hcl:"weight,attr" validate:"required"`
	// TrafficConfigSingleWeightConfigSessionStickinessConfig: min=0
	SessionStickinessConfig []TrafficConfigSingleWeightConfigSessionStickinessConfig `hcl:"session_stickiness_config,block" validate:"min=0"`
}

type TrafficConfigSingleWeightConfigSessionStickinessConfig struct {
	// IdleTtl: number, required
	IdleTtl terra.NumberValue `hcl:"idle_ttl,attr" validate:"required"`
	// MaximumTtl: number, required
	MaximumTtl terra.NumberValue `hcl:"maximum_ttl,attr" validate:"required"`
}

type StagingDistributionDnsNamesAttributes struct {
	ref terra.Reference
}

func (sddn StagingDistributionDnsNamesAttributes) InternalRef() (terra.Reference, error) {
	return sddn.ref, nil
}

func (sddn StagingDistributionDnsNamesAttributes) InternalWithRef(ref terra.Reference) StagingDistributionDnsNamesAttributes {
	return StagingDistributionDnsNamesAttributes{ref: ref}
}

func (sddn StagingDistributionDnsNamesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sddn.ref.InternalTokens()
}

func (sddn StagingDistributionDnsNamesAttributes) Items() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sddn.ref.Append("items"))
}

func (sddn StagingDistributionDnsNamesAttributes) Quantity() terra.NumberValue {
	return terra.ReferenceAsNumber(sddn.ref.Append("quantity"))
}

type TrafficConfigAttributes struct {
	ref terra.Reference
}

func (tc TrafficConfigAttributes) InternalRef() (terra.Reference, error) {
	return tc.ref, nil
}

func (tc TrafficConfigAttributes) InternalWithRef(ref terra.Reference) TrafficConfigAttributes {
	return TrafficConfigAttributes{ref: ref}
}

func (tc TrafficConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tc.ref.InternalTokens()
}

func (tc TrafficConfigAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(tc.ref.Append("type"))
}

func (tc TrafficConfigAttributes) SingleHeaderConfig() terra.ListValue[TrafficConfigSingleHeaderConfigAttributes] {
	return terra.ReferenceAsList[TrafficConfigSingleHeaderConfigAttributes](tc.ref.Append("single_header_config"))
}

func (tc TrafficConfigAttributes) SingleWeightConfig() terra.ListValue[TrafficConfigSingleWeightConfigAttributes] {
	return terra.ReferenceAsList[TrafficConfigSingleWeightConfigAttributes](tc.ref.Append("single_weight_config"))
}

type TrafficConfigSingleHeaderConfigAttributes struct {
	ref terra.Reference
}

func (shc TrafficConfigSingleHeaderConfigAttributes) InternalRef() (terra.Reference, error) {
	return shc.ref, nil
}

func (shc TrafficConfigSingleHeaderConfigAttributes) InternalWithRef(ref terra.Reference) TrafficConfigSingleHeaderConfigAttributes {
	return TrafficConfigSingleHeaderConfigAttributes{ref: ref}
}

func (shc TrafficConfigSingleHeaderConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return shc.ref.InternalTokens()
}

func (shc TrafficConfigSingleHeaderConfigAttributes) Header() terra.StringValue {
	return terra.ReferenceAsString(shc.ref.Append("header"))
}

func (shc TrafficConfigSingleHeaderConfigAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(shc.ref.Append("value"))
}

type TrafficConfigSingleWeightConfigAttributes struct {
	ref terra.Reference
}

func (swc TrafficConfigSingleWeightConfigAttributes) InternalRef() (terra.Reference, error) {
	return swc.ref, nil
}

func (swc TrafficConfigSingleWeightConfigAttributes) InternalWithRef(ref terra.Reference) TrafficConfigSingleWeightConfigAttributes {
	return TrafficConfigSingleWeightConfigAttributes{ref: ref}
}

func (swc TrafficConfigSingleWeightConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return swc.ref.InternalTokens()
}

func (swc TrafficConfigSingleWeightConfigAttributes) Weight() terra.NumberValue {
	return terra.ReferenceAsNumber(swc.ref.Append("weight"))
}

func (swc TrafficConfigSingleWeightConfigAttributes) SessionStickinessConfig() terra.ListValue[TrafficConfigSingleWeightConfigSessionStickinessConfigAttributes] {
	return terra.ReferenceAsList[TrafficConfigSingleWeightConfigSessionStickinessConfigAttributes](swc.ref.Append("session_stickiness_config"))
}

type TrafficConfigSingleWeightConfigSessionStickinessConfigAttributes struct {
	ref terra.Reference
}

func (ssc TrafficConfigSingleWeightConfigSessionStickinessConfigAttributes) InternalRef() (terra.Reference, error) {
	return ssc.ref, nil
}

func (ssc TrafficConfigSingleWeightConfigSessionStickinessConfigAttributes) InternalWithRef(ref terra.Reference) TrafficConfigSingleWeightConfigSessionStickinessConfigAttributes {
	return TrafficConfigSingleWeightConfigSessionStickinessConfigAttributes{ref: ref}
}

func (ssc TrafficConfigSingleWeightConfigSessionStickinessConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ssc.ref.InternalTokens()
}

func (ssc TrafficConfigSingleWeightConfigSessionStickinessConfigAttributes) IdleTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(ssc.ref.Append("idle_ttl"))
}

func (ssc TrafficConfigSingleWeightConfigSessionStickinessConfigAttributes) MaximumTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(ssc.ref.Append("maximum_ttl"))
}

type StagingDistributionDnsNamesState struct {
	Items    []string `json:"items"`
	Quantity float64  `json:"quantity"`
}

type TrafficConfigState struct {
	Type               string                                 `json:"type"`
	SingleHeaderConfig []TrafficConfigSingleHeaderConfigState `json:"single_header_config"`
	SingleWeightConfig []TrafficConfigSingleWeightConfigState `json:"single_weight_config"`
}

type TrafficConfigSingleHeaderConfigState struct {
	Header string `json:"header"`
	Value  string `json:"value"`
}

type TrafficConfigSingleWeightConfigState struct {
	Weight                  float64                                                       `json:"weight"`
	SessionStickinessConfig []TrafficConfigSingleWeightConfigSessionStickinessConfigState `json:"session_stickiness_config"`
}

type TrafficConfigSingleWeightConfigSessionStickinessConfigState struct {
	IdleTtl    float64 `json:"idle_ttl"`
	MaximumTtl float64 `json:"maximum_ttl"`
}
