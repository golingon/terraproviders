// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package cloudwatchmetricstream

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ExcludeFilter struct {
	// Namespace: string, required
	Namespace terra.StringValue `hcl:"namespace,attr" validate:"required"`
}

type IncludeFilter struct {
	// Namespace: string, required
	Namespace terra.StringValue `hcl:"namespace,attr" validate:"required"`
}

type StatisticsConfiguration struct {
	// AdditionalStatistics: set of string, required
	AdditionalStatistics terra.SetValue[terra.StringValue] `hcl:"additional_statistics,attr" validate:"required"`
	// IncludeMetric: min=1
	IncludeMetric []IncludeMetric `hcl:"include_metric,block" validate:"min=1"`
}

type IncludeMetric struct {
	// MetricName: string, required
	MetricName terra.StringValue `hcl:"metric_name,attr" validate:"required"`
	// Namespace: string, required
	Namespace terra.StringValue `hcl:"namespace,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type ExcludeFilterAttributes struct {
	ref terra.Reference
}

func (ef ExcludeFilterAttributes) InternalRef() terra.Reference {
	return ef.ref
}

func (ef ExcludeFilterAttributes) InternalWithRef(ref terra.Reference) ExcludeFilterAttributes {
	return ExcludeFilterAttributes{ref: ref}
}

func (ef ExcludeFilterAttributes) InternalTokens() hclwrite.Tokens {
	return ef.ref.InternalTokens()
}

func (ef ExcludeFilterAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(ef.ref.Append("namespace"))
}

type IncludeFilterAttributes struct {
	ref terra.Reference
}

func (_if IncludeFilterAttributes) InternalRef() terra.Reference {
	return _if.ref
}

func (_if IncludeFilterAttributes) InternalWithRef(ref terra.Reference) IncludeFilterAttributes {
	return IncludeFilterAttributes{ref: ref}
}

func (_if IncludeFilterAttributes) InternalTokens() hclwrite.Tokens {
	return _if.ref.InternalTokens()
}

func (_if IncludeFilterAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(_if.ref.Append("namespace"))
}

type StatisticsConfigurationAttributes struct {
	ref terra.Reference
}

func (sc StatisticsConfigurationAttributes) InternalRef() terra.Reference {
	return sc.ref
}

func (sc StatisticsConfigurationAttributes) InternalWithRef(ref terra.Reference) StatisticsConfigurationAttributes {
	return StatisticsConfigurationAttributes{ref: ref}
}

func (sc StatisticsConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return sc.ref.InternalTokens()
}

func (sc StatisticsConfigurationAttributes) AdditionalStatistics() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sc.ref.Append("additional_statistics"))
}

func (sc StatisticsConfigurationAttributes) IncludeMetric() terra.SetValue[IncludeMetricAttributes] {
	return terra.ReferenceAsSet[IncludeMetricAttributes](sc.ref.Append("include_metric"))
}

type IncludeMetricAttributes struct {
	ref terra.Reference
}

func (im IncludeMetricAttributes) InternalRef() terra.Reference {
	return im.ref
}

func (im IncludeMetricAttributes) InternalWithRef(ref terra.Reference) IncludeMetricAttributes {
	return IncludeMetricAttributes{ref: ref}
}

func (im IncludeMetricAttributes) InternalTokens() hclwrite.Tokens {
	return im.ref.InternalTokens()
}

func (im IncludeMetricAttributes) MetricName() terra.StringValue {
	return terra.ReferenceAsString(im.ref.Append("metric_name"))
}

func (im IncludeMetricAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(im.ref.Append("namespace"))
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
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type ExcludeFilterState struct {
	Namespace string `json:"namespace"`
}

type IncludeFilterState struct {
	Namespace string `json:"namespace"`
}

type StatisticsConfigurationState struct {
	AdditionalStatistics []string             `json:"additional_statistics"`
	IncludeMetric        []IncludeMetricState `json:"include_metric"`
}

type IncludeMetricState struct {
	MetricName string `json:"metric_name"`
	Namespace  string `json:"namespace"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
