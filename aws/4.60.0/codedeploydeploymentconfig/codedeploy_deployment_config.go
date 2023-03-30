// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package codedeploydeploymentconfig

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type MinimumHealthyHosts struct {
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// Value: number, optional
	Value terra.NumberValue `hcl:"value,attr"`
}

type TrafficRoutingConfig struct {
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// TimeBasedCanary: optional
	TimeBasedCanary *TimeBasedCanary `hcl:"time_based_canary,block"`
	// TimeBasedLinear: optional
	TimeBasedLinear *TimeBasedLinear `hcl:"time_based_linear,block"`
}

type TimeBasedCanary struct {
	// Interval: number, optional
	Interval terra.NumberValue `hcl:"interval,attr"`
	// Percentage: number, optional
	Percentage terra.NumberValue `hcl:"percentage,attr"`
}

type TimeBasedLinear struct {
	// Interval: number, optional
	Interval terra.NumberValue `hcl:"interval,attr"`
	// Percentage: number, optional
	Percentage terra.NumberValue `hcl:"percentage,attr"`
}

type MinimumHealthyHostsAttributes struct {
	ref terra.Reference
}

func (mhh MinimumHealthyHostsAttributes) InternalRef() terra.Reference {
	return mhh.ref
}

func (mhh MinimumHealthyHostsAttributes) InternalWithRef(ref terra.Reference) MinimumHealthyHostsAttributes {
	return MinimumHealthyHostsAttributes{ref: ref}
}

func (mhh MinimumHealthyHostsAttributes) InternalTokens() hclwrite.Tokens {
	return mhh.ref.InternalTokens()
}

func (mhh MinimumHealthyHostsAttributes) Type() terra.StringValue {
	return terra.ReferenceString(mhh.ref.Append("type"))
}

func (mhh MinimumHealthyHostsAttributes) Value() terra.NumberValue {
	return terra.ReferenceNumber(mhh.ref.Append("value"))
}

type TrafficRoutingConfigAttributes struct {
	ref terra.Reference
}

func (trc TrafficRoutingConfigAttributes) InternalRef() terra.Reference {
	return trc.ref
}

func (trc TrafficRoutingConfigAttributes) InternalWithRef(ref terra.Reference) TrafficRoutingConfigAttributes {
	return TrafficRoutingConfigAttributes{ref: ref}
}

func (trc TrafficRoutingConfigAttributes) InternalTokens() hclwrite.Tokens {
	return trc.ref.InternalTokens()
}

func (trc TrafficRoutingConfigAttributes) Type() terra.StringValue {
	return terra.ReferenceString(trc.ref.Append("type"))
}

func (trc TrafficRoutingConfigAttributes) TimeBasedCanary() terra.ListValue[TimeBasedCanaryAttributes] {
	return terra.ReferenceList[TimeBasedCanaryAttributes](trc.ref.Append("time_based_canary"))
}

func (trc TrafficRoutingConfigAttributes) TimeBasedLinear() terra.ListValue[TimeBasedLinearAttributes] {
	return terra.ReferenceList[TimeBasedLinearAttributes](trc.ref.Append("time_based_linear"))
}

type TimeBasedCanaryAttributes struct {
	ref terra.Reference
}

func (tbc TimeBasedCanaryAttributes) InternalRef() terra.Reference {
	return tbc.ref
}

func (tbc TimeBasedCanaryAttributes) InternalWithRef(ref terra.Reference) TimeBasedCanaryAttributes {
	return TimeBasedCanaryAttributes{ref: ref}
}

func (tbc TimeBasedCanaryAttributes) InternalTokens() hclwrite.Tokens {
	return tbc.ref.InternalTokens()
}

func (tbc TimeBasedCanaryAttributes) Interval() terra.NumberValue {
	return terra.ReferenceNumber(tbc.ref.Append("interval"))
}

func (tbc TimeBasedCanaryAttributes) Percentage() terra.NumberValue {
	return terra.ReferenceNumber(tbc.ref.Append("percentage"))
}

type TimeBasedLinearAttributes struct {
	ref terra.Reference
}

func (tbl TimeBasedLinearAttributes) InternalRef() terra.Reference {
	return tbl.ref
}

func (tbl TimeBasedLinearAttributes) InternalWithRef(ref terra.Reference) TimeBasedLinearAttributes {
	return TimeBasedLinearAttributes{ref: ref}
}

func (tbl TimeBasedLinearAttributes) InternalTokens() hclwrite.Tokens {
	return tbl.ref.InternalTokens()
}

func (tbl TimeBasedLinearAttributes) Interval() terra.NumberValue {
	return terra.ReferenceNumber(tbl.ref.Append("interval"))
}

func (tbl TimeBasedLinearAttributes) Percentage() terra.NumberValue {
	return terra.ReferenceNumber(tbl.ref.Append("percentage"))
}

type MinimumHealthyHostsState struct {
	Type  string  `json:"type"`
	Value float64 `json:"value"`
}

type TrafficRoutingConfigState struct {
	Type            string                 `json:"type"`
	TimeBasedCanary []TimeBasedCanaryState `json:"time_based_canary"`
	TimeBasedLinear []TimeBasedLinearState `json:"time_based_linear"`
}

type TimeBasedCanaryState struct {
	Interval   float64 `json:"interval"`
	Percentage float64 `json:"percentage"`
}

type TimeBasedLinearState struct {
	Interval   float64 `json:"interval"`
	Percentage float64 `json:"percentage"`
}
