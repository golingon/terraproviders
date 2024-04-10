// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package servicequotasservicequota

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type UsageMetric struct {
	// MetricDimensions: min=0
	MetricDimensions []MetricDimensions `hcl:"metric_dimensions,block" validate:"min=0"`
}

type MetricDimensions struct{}

type UsageMetricAttributes struct {
	ref terra.Reference
}

func (um UsageMetricAttributes) InternalRef() (terra.Reference, error) {
	return um.ref, nil
}

func (um UsageMetricAttributes) InternalWithRef(ref terra.Reference) UsageMetricAttributes {
	return UsageMetricAttributes{ref: ref}
}

func (um UsageMetricAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return um.ref.InternalTokens()
}

func (um UsageMetricAttributes) MetricName() terra.StringValue {
	return terra.ReferenceAsString(um.ref.Append("metric_name"))
}

func (um UsageMetricAttributes) MetricNamespace() terra.StringValue {
	return terra.ReferenceAsString(um.ref.Append("metric_namespace"))
}

func (um UsageMetricAttributes) MetricStatisticRecommendation() terra.StringValue {
	return terra.ReferenceAsString(um.ref.Append("metric_statistic_recommendation"))
}

func (um UsageMetricAttributes) MetricDimensions() terra.ListValue[MetricDimensionsAttributes] {
	return terra.ReferenceAsList[MetricDimensionsAttributes](um.ref.Append("metric_dimensions"))
}

type MetricDimensionsAttributes struct {
	ref terra.Reference
}

func (md MetricDimensionsAttributes) InternalRef() (terra.Reference, error) {
	return md.ref, nil
}

func (md MetricDimensionsAttributes) InternalWithRef(ref terra.Reference) MetricDimensionsAttributes {
	return MetricDimensionsAttributes{ref: ref}
}

func (md MetricDimensionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return md.ref.InternalTokens()
}

func (md MetricDimensionsAttributes) Class() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("class"))
}

func (md MetricDimensionsAttributes) Resource() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("resource"))
}

func (md MetricDimensionsAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("service"))
}

func (md MetricDimensionsAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("type"))
}

type UsageMetricState struct {
	MetricName                    string                  `json:"metric_name"`
	MetricNamespace               string                  `json:"metric_namespace"`
	MetricStatisticRecommendation string                  `json:"metric_statistic_recommendation"`
	MetricDimensions              []MetricDimensionsState `json:"metric_dimensions"`
}

type MetricDimensionsState struct {
	Class    string `json:"class"`
	Resource string `json:"resource"`
	Service  string `json:"service"`
	Type     string `json:"type"`
}
