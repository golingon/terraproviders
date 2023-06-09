// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package monitoringslo

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type BasicSli struct {
	// Location: set of string, optional
	Location terra.SetValue[terra.StringValue] `hcl:"location,attr"`
	// Method: set of string, optional
	Method terra.SetValue[terra.StringValue] `hcl:"method,attr"`
	// Version: set of string, optional
	Version terra.SetValue[terra.StringValue] `hcl:"version,attr"`
	// BasicSliAvailability: optional
	Availability *BasicSliAvailability `hcl:"availability,block"`
	// BasicSliLatency: optional
	Latency *BasicSliLatency `hcl:"latency,block"`
}

type BasicSliAvailability struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
}

type BasicSliLatency struct {
	// Threshold: string, required
	Threshold terra.StringValue `hcl:"threshold,attr" validate:"required"`
}

type RequestBasedSli struct {
	// RequestBasedSliDistributionCut: optional
	DistributionCut *RequestBasedSliDistributionCut `hcl:"distribution_cut,block"`
	// RequestBasedSliGoodTotalRatio: optional
	GoodTotalRatio *RequestBasedSliGoodTotalRatio `hcl:"good_total_ratio,block"`
}

type RequestBasedSliDistributionCut struct {
	// DistributionFilter: string, required
	DistributionFilter terra.StringValue `hcl:"distribution_filter,attr" validate:"required"`
	// RequestBasedSliDistributionCutRange: required
	Range *RequestBasedSliDistributionCutRange `hcl:"range,block" validate:"required"`
}

type RequestBasedSliDistributionCutRange struct {
	// Max: number, optional
	Max terra.NumberValue `hcl:"max,attr"`
	// Min: number, optional
	Min terra.NumberValue `hcl:"min,attr"`
}

type RequestBasedSliGoodTotalRatio struct {
	// BadServiceFilter: string, optional
	BadServiceFilter terra.StringValue `hcl:"bad_service_filter,attr"`
	// GoodServiceFilter: string, optional
	GoodServiceFilter terra.StringValue `hcl:"good_service_filter,attr"`
	// TotalServiceFilter: string, optional
	TotalServiceFilter terra.StringValue `hcl:"total_service_filter,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type WindowsBasedSli struct {
	// GoodBadMetricFilter: string, optional
	GoodBadMetricFilter terra.StringValue `hcl:"good_bad_metric_filter,attr"`
	// WindowPeriod: string, optional
	WindowPeriod terra.StringValue `hcl:"window_period,attr"`
	// GoodTotalRatioThreshold: optional
	GoodTotalRatioThreshold *GoodTotalRatioThreshold `hcl:"good_total_ratio_threshold,block"`
	// MetricMeanInRange: optional
	MetricMeanInRange *MetricMeanInRange `hcl:"metric_mean_in_range,block"`
	// MetricSumInRange: optional
	MetricSumInRange *MetricSumInRange `hcl:"metric_sum_in_range,block"`
}

type GoodTotalRatioThreshold struct {
	// Threshold: number, optional
	Threshold terra.NumberValue `hcl:"threshold,attr"`
	// BasicSliPerformance: optional
	BasicSliPerformance *BasicSliPerformance `hcl:"basic_sli_performance,block"`
	// Performance: optional
	Performance *Performance `hcl:"performance,block"`
}

type BasicSliPerformance struct {
	// Location: set of string, optional
	Location terra.SetValue[terra.StringValue] `hcl:"location,attr"`
	// Method: set of string, optional
	Method terra.SetValue[terra.StringValue] `hcl:"method,attr"`
	// Version: set of string, optional
	Version terra.SetValue[terra.StringValue] `hcl:"version,attr"`
	// BasicSliPerformanceAvailability: optional
	Availability *BasicSliPerformanceAvailability `hcl:"availability,block"`
	// BasicSliPerformanceLatency: optional
	Latency *BasicSliPerformanceLatency `hcl:"latency,block"`
}

type BasicSliPerformanceAvailability struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
}

type BasicSliPerformanceLatency struct {
	// Threshold: string, required
	Threshold terra.StringValue `hcl:"threshold,attr" validate:"required"`
}

type Performance struct {
	// PerformanceDistributionCut: optional
	DistributionCut *PerformanceDistributionCut `hcl:"distribution_cut,block"`
	// PerformanceGoodTotalRatio: optional
	GoodTotalRatio *PerformanceGoodTotalRatio `hcl:"good_total_ratio,block"`
}

type PerformanceDistributionCut struct {
	// DistributionFilter: string, required
	DistributionFilter terra.StringValue `hcl:"distribution_filter,attr" validate:"required"`
	// PerformanceDistributionCutRange: required
	Range *PerformanceDistributionCutRange `hcl:"range,block" validate:"required"`
}

type PerformanceDistributionCutRange struct {
	// Max: number, optional
	Max terra.NumberValue `hcl:"max,attr"`
	// Min: number, optional
	Min terra.NumberValue `hcl:"min,attr"`
}

type PerformanceGoodTotalRatio struct {
	// BadServiceFilter: string, optional
	BadServiceFilter terra.StringValue `hcl:"bad_service_filter,attr"`
	// GoodServiceFilter: string, optional
	GoodServiceFilter terra.StringValue `hcl:"good_service_filter,attr"`
	// TotalServiceFilter: string, optional
	TotalServiceFilter terra.StringValue `hcl:"total_service_filter,attr"`
}

type MetricMeanInRange struct {
	// TimeSeries: string, required
	TimeSeries terra.StringValue `hcl:"time_series,attr" validate:"required"`
	// MetricMeanInRangeRange: required
	Range *MetricMeanInRangeRange `hcl:"range,block" validate:"required"`
}

type MetricMeanInRangeRange struct {
	// Max: number, optional
	Max terra.NumberValue `hcl:"max,attr"`
	// Min: number, optional
	Min terra.NumberValue `hcl:"min,attr"`
}

type MetricSumInRange struct {
	// TimeSeries: string, required
	TimeSeries terra.StringValue `hcl:"time_series,attr" validate:"required"`
	// MetricSumInRangeRange: required
	Range *MetricSumInRangeRange `hcl:"range,block" validate:"required"`
}

type MetricSumInRangeRange struct {
	// Max: number, optional
	Max terra.NumberValue `hcl:"max,attr"`
	// Min: number, optional
	Min terra.NumberValue `hcl:"min,attr"`
}

type BasicSliAttributes struct {
	ref terra.Reference
}

func (bs BasicSliAttributes) InternalRef() (terra.Reference, error) {
	return bs.ref, nil
}

func (bs BasicSliAttributes) InternalWithRef(ref terra.Reference) BasicSliAttributes {
	return BasicSliAttributes{ref: ref}
}

func (bs BasicSliAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bs.ref.InternalTokens()
}

func (bs BasicSliAttributes) Location() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](bs.ref.Append("location"))
}

func (bs BasicSliAttributes) Method() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](bs.ref.Append("method"))
}

func (bs BasicSliAttributes) Version() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](bs.ref.Append("version"))
}

func (bs BasicSliAttributes) Availability() terra.ListValue[BasicSliAvailabilityAttributes] {
	return terra.ReferenceAsList[BasicSliAvailabilityAttributes](bs.ref.Append("availability"))
}

func (bs BasicSliAttributes) Latency() terra.ListValue[BasicSliLatencyAttributes] {
	return terra.ReferenceAsList[BasicSliLatencyAttributes](bs.ref.Append("latency"))
}

type BasicSliAvailabilityAttributes struct {
	ref terra.Reference
}

func (a BasicSliAvailabilityAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a BasicSliAvailabilityAttributes) InternalWithRef(ref terra.Reference) BasicSliAvailabilityAttributes {
	return BasicSliAvailabilityAttributes{ref: ref}
}

func (a BasicSliAvailabilityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a BasicSliAvailabilityAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("enabled"))
}

type BasicSliLatencyAttributes struct {
	ref terra.Reference
}

func (l BasicSliLatencyAttributes) InternalRef() (terra.Reference, error) {
	return l.ref, nil
}

func (l BasicSliLatencyAttributes) InternalWithRef(ref terra.Reference) BasicSliLatencyAttributes {
	return BasicSliLatencyAttributes{ref: ref}
}

func (l BasicSliLatencyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return l.ref.InternalTokens()
}

func (l BasicSliLatencyAttributes) Threshold() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("threshold"))
}

type RequestBasedSliAttributes struct {
	ref terra.Reference
}

func (rbs RequestBasedSliAttributes) InternalRef() (terra.Reference, error) {
	return rbs.ref, nil
}

func (rbs RequestBasedSliAttributes) InternalWithRef(ref terra.Reference) RequestBasedSliAttributes {
	return RequestBasedSliAttributes{ref: ref}
}

func (rbs RequestBasedSliAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rbs.ref.InternalTokens()
}

func (rbs RequestBasedSliAttributes) DistributionCut() terra.ListValue[RequestBasedSliDistributionCutAttributes] {
	return terra.ReferenceAsList[RequestBasedSliDistributionCutAttributes](rbs.ref.Append("distribution_cut"))
}

func (rbs RequestBasedSliAttributes) GoodTotalRatio() terra.ListValue[RequestBasedSliGoodTotalRatioAttributes] {
	return terra.ReferenceAsList[RequestBasedSliGoodTotalRatioAttributes](rbs.ref.Append("good_total_ratio"))
}

type RequestBasedSliDistributionCutAttributes struct {
	ref terra.Reference
}

func (dc RequestBasedSliDistributionCutAttributes) InternalRef() (terra.Reference, error) {
	return dc.ref, nil
}

func (dc RequestBasedSliDistributionCutAttributes) InternalWithRef(ref terra.Reference) RequestBasedSliDistributionCutAttributes {
	return RequestBasedSliDistributionCutAttributes{ref: ref}
}

func (dc RequestBasedSliDistributionCutAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dc.ref.InternalTokens()
}

func (dc RequestBasedSliDistributionCutAttributes) DistributionFilter() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("distribution_filter"))
}

func (dc RequestBasedSliDistributionCutAttributes) Range() terra.ListValue[RequestBasedSliDistributionCutRangeAttributes] {
	return terra.ReferenceAsList[RequestBasedSliDistributionCutRangeAttributes](dc.ref.Append("range"))
}

type RequestBasedSliDistributionCutRangeAttributes struct {
	ref terra.Reference
}

func (r RequestBasedSliDistributionCutRangeAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RequestBasedSliDistributionCutRangeAttributes) InternalWithRef(ref terra.Reference) RequestBasedSliDistributionCutRangeAttributes {
	return RequestBasedSliDistributionCutRangeAttributes{ref: ref}
}

func (r RequestBasedSliDistributionCutRangeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RequestBasedSliDistributionCutRangeAttributes) Max() terra.NumberValue {
	return terra.ReferenceAsNumber(r.ref.Append("max"))
}

func (r RequestBasedSliDistributionCutRangeAttributes) Min() terra.NumberValue {
	return terra.ReferenceAsNumber(r.ref.Append("min"))
}

type RequestBasedSliGoodTotalRatioAttributes struct {
	ref terra.Reference
}

func (gtr RequestBasedSliGoodTotalRatioAttributes) InternalRef() (terra.Reference, error) {
	return gtr.ref, nil
}

func (gtr RequestBasedSliGoodTotalRatioAttributes) InternalWithRef(ref terra.Reference) RequestBasedSliGoodTotalRatioAttributes {
	return RequestBasedSliGoodTotalRatioAttributes{ref: ref}
}

func (gtr RequestBasedSliGoodTotalRatioAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return gtr.ref.InternalTokens()
}

func (gtr RequestBasedSliGoodTotalRatioAttributes) BadServiceFilter() terra.StringValue {
	return terra.ReferenceAsString(gtr.ref.Append("bad_service_filter"))
}

func (gtr RequestBasedSliGoodTotalRatioAttributes) GoodServiceFilter() terra.StringValue {
	return terra.ReferenceAsString(gtr.ref.Append("good_service_filter"))
}

func (gtr RequestBasedSliGoodTotalRatioAttributes) TotalServiceFilter() terra.StringValue {
	return terra.ReferenceAsString(gtr.ref.Append("total_service_filter"))
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

type WindowsBasedSliAttributes struct {
	ref terra.Reference
}

func (wbs WindowsBasedSliAttributes) InternalRef() (terra.Reference, error) {
	return wbs.ref, nil
}

func (wbs WindowsBasedSliAttributes) InternalWithRef(ref terra.Reference) WindowsBasedSliAttributes {
	return WindowsBasedSliAttributes{ref: ref}
}

func (wbs WindowsBasedSliAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return wbs.ref.InternalTokens()
}

func (wbs WindowsBasedSliAttributes) GoodBadMetricFilter() terra.StringValue {
	return terra.ReferenceAsString(wbs.ref.Append("good_bad_metric_filter"))
}

func (wbs WindowsBasedSliAttributes) WindowPeriod() terra.StringValue {
	return terra.ReferenceAsString(wbs.ref.Append("window_period"))
}

func (wbs WindowsBasedSliAttributes) GoodTotalRatioThreshold() terra.ListValue[GoodTotalRatioThresholdAttributes] {
	return terra.ReferenceAsList[GoodTotalRatioThresholdAttributes](wbs.ref.Append("good_total_ratio_threshold"))
}

func (wbs WindowsBasedSliAttributes) MetricMeanInRange() terra.ListValue[MetricMeanInRangeAttributes] {
	return terra.ReferenceAsList[MetricMeanInRangeAttributes](wbs.ref.Append("metric_mean_in_range"))
}

func (wbs WindowsBasedSliAttributes) MetricSumInRange() terra.ListValue[MetricSumInRangeAttributes] {
	return terra.ReferenceAsList[MetricSumInRangeAttributes](wbs.ref.Append("metric_sum_in_range"))
}

type GoodTotalRatioThresholdAttributes struct {
	ref terra.Reference
}

func (gtrt GoodTotalRatioThresholdAttributes) InternalRef() (terra.Reference, error) {
	return gtrt.ref, nil
}

func (gtrt GoodTotalRatioThresholdAttributes) InternalWithRef(ref terra.Reference) GoodTotalRatioThresholdAttributes {
	return GoodTotalRatioThresholdAttributes{ref: ref}
}

func (gtrt GoodTotalRatioThresholdAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return gtrt.ref.InternalTokens()
}

func (gtrt GoodTotalRatioThresholdAttributes) Threshold() terra.NumberValue {
	return terra.ReferenceAsNumber(gtrt.ref.Append("threshold"))
}

func (gtrt GoodTotalRatioThresholdAttributes) BasicSliPerformance() terra.ListValue[BasicSliPerformanceAttributes] {
	return terra.ReferenceAsList[BasicSliPerformanceAttributes](gtrt.ref.Append("basic_sli_performance"))
}

func (gtrt GoodTotalRatioThresholdAttributes) Performance() terra.ListValue[PerformanceAttributes] {
	return terra.ReferenceAsList[PerformanceAttributes](gtrt.ref.Append("performance"))
}

type BasicSliPerformanceAttributes struct {
	ref terra.Reference
}

func (bsp BasicSliPerformanceAttributes) InternalRef() (terra.Reference, error) {
	return bsp.ref, nil
}

func (bsp BasicSliPerformanceAttributes) InternalWithRef(ref terra.Reference) BasicSliPerformanceAttributes {
	return BasicSliPerformanceAttributes{ref: ref}
}

func (bsp BasicSliPerformanceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bsp.ref.InternalTokens()
}

func (bsp BasicSliPerformanceAttributes) Location() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](bsp.ref.Append("location"))
}

func (bsp BasicSliPerformanceAttributes) Method() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](bsp.ref.Append("method"))
}

func (bsp BasicSliPerformanceAttributes) Version() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](bsp.ref.Append("version"))
}

func (bsp BasicSliPerformanceAttributes) Availability() terra.ListValue[BasicSliPerformanceAvailabilityAttributes] {
	return terra.ReferenceAsList[BasicSliPerformanceAvailabilityAttributes](bsp.ref.Append("availability"))
}

func (bsp BasicSliPerformanceAttributes) Latency() terra.ListValue[BasicSliPerformanceLatencyAttributes] {
	return terra.ReferenceAsList[BasicSliPerformanceLatencyAttributes](bsp.ref.Append("latency"))
}

type BasicSliPerformanceAvailabilityAttributes struct {
	ref terra.Reference
}

func (a BasicSliPerformanceAvailabilityAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a BasicSliPerformanceAvailabilityAttributes) InternalWithRef(ref terra.Reference) BasicSliPerformanceAvailabilityAttributes {
	return BasicSliPerformanceAvailabilityAttributes{ref: ref}
}

func (a BasicSliPerformanceAvailabilityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a BasicSliPerformanceAvailabilityAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("enabled"))
}

type BasicSliPerformanceLatencyAttributes struct {
	ref terra.Reference
}

func (l BasicSliPerformanceLatencyAttributes) InternalRef() (terra.Reference, error) {
	return l.ref, nil
}

func (l BasicSliPerformanceLatencyAttributes) InternalWithRef(ref terra.Reference) BasicSliPerformanceLatencyAttributes {
	return BasicSliPerformanceLatencyAttributes{ref: ref}
}

func (l BasicSliPerformanceLatencyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return l.ref.InternalTokens()
}

func (l BasicSliPerformanceLatencyAttributes) Threshold() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("threshold"))
}

type PerformanceAttributes struct {
	ref terra.Reference
}

func (p PerformanceAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p PerformanceAttributes) InternalWithRef(ref terra.Reference) PerformanceAttributes {
	return PerformanceAttributes{ref: ref}
}

func (p PerformanceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p PerformanceAttributes) DistributionCut() terra.ListValue[PerformanceDistributionCutAttributes] {
	return terra.ReferenceAsList[PerformanceDistributionCutAttributes](p.ref.Append("distribution_cut"))
}

func (p PerformanceAttributes) GoodTotalRatio() terra.ListValue[PerformanceGoodTotalRatioAttributes] {
	return terra.ReferenceAsList[PerformanceGoodTotalRatioAttributes](p.ref.Append("good_total_ratio"))
}

type PerformanceDistributionCutAttributes struct {
	ref terra.Reference
}

func (dc PerformanceDistributionCutAttributes) InternalRef() (terra.Reference, error) {
	return dc.ref, nil
}

func (dc PerformanceDistributionCutAttributes) InternalWithRef(ref terra.Reference) PerformanceDistributionCutAttributes {
	return PerformanceDistributionCutAttributes{ref: ref}
}

func (dc PerformanceDistributionCutAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dc.ref.InternalTokens()
}

func (dc PerformanceDistributionCutAttributes) DistributionFilter() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("distribution_filter"))
}

func (dc PerformanceDistributionCutAttributes) Range() terra.ListValue[PerformanceDistributionCutRangeAttributes] {
	return terra.ReferenceAsList[PerformanceDistributionCutRangeAttributes](dc.ref.Append("range"))
}

type PerformanceDistributionCutRangeAttributes struct {
	ref terra.Reference
}

func (r PerformanceDistributionCutRangeAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r PerformanceDistributionCutRangeAttributes) InternalWithRef(ref terra.Reference) PerformanceDistributionCutRangeAttributes {
	return PerformanceDistributionCutRangeAttributes{ref: ref}
}

func (r PerformanceDistributionCutRangeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r PerformanceDistributionCutRangeAttributes) Max() terra.NumberValue {
	return terra.ReferenceAsNumber(r.ref.Append("max"))
}

func (r PerformanceDistributionCutRangeAttributes) Min() terra.NumberValue {
	return terra.ReferenceAsNumber(r.ref.Append("min"))
}

type PerformanceGoodTotalRatioAttributes struct {
	ref terra.Reference
}

func (gtr PerformanceGoodTotalRatioAttributes) InternalRef() (terra.Reference, error) {
	return gtr.ref, nil
}

func (gtr PerformanceGoodTotalRatioAttributes) InternalWithRef(ref terra.Reference) PerformanceGoodTotalRatioAttributes {
	return PerformanceGoodTotalRatioAttributes{ref: ref}
}

func (gtr PerformanceGoodTotalRatioAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return gtr.ref.InternalTokens()
}

func (gtr PerformanceGoodTotalRatioAttributes) BadServiceFilter() terra.StringValue {
	return terra.ReferenceAsString(gtr.ref.Append("bad_service_filter"))
}

func (gtr PerformanceGoodTotalRatioAttributes) GoodServiceFilter() terra.StringValue {
	return terra.ReferenceAsString(gtr.ref.Append("good_service_filter"))
}

func (gtr PerformanceGoodTotalRatioAttributes) TotalServiceFilter() terra.StringValue {
	return terra.ReferenceAsString(gtr.ref.Append("total_service_filter"))
}

type MetricMeanInRangeAttributes struct {
	ref terra.Reference
}

func (mmir MetricMeanInRangeAttributes) InternalRef() (terra.Reference, error) {
	return mmir.ref, nil
}

func (mmir MetricMeanInRangeAttributes) InternalWithRef(ref terra.Reference) MetricMeanInRangeAttributes {
	return MetricMeanInRangeAttributes{ref: ref}
}

func (mmir MetricMeanInRangeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mmir.ref.InternalTokens()
}

func (mmir MetricMeanInRangeAttributes) TimeSeries() terra.StringValue {
	return terra.ReferenceAsString(mmir.ref.Append("time_series"))
}

func (mmir MetricMeanInRangeAttributes) Range() terra.ListValue[MetricMeanInRangeRangeAttributes] {
	return terra.ReferenceAsList[MetricMeanInRangeRangeAttributes](mmir.ref.Append("range"))
}

type MetricMeanInRangeRangeAttributes struct {
	ref terra.Reference
}

func (r MetricMeanInRangeRangeAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r MetricMeanInRangeRangeAttributes) InternalWithRef(ref terra.Reference) MetricMeanInRangeRangeAttributes {
	return MetricMeanInRangeRangeAttributes{ref: ref}
}

func (r MetricMeanInRangeRangeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r MetricMeanInRangeRangeAttributes) Max() terra.NumberValue {
	return terra.ReferenceAsNumber(r.ref.Append("max"))
}

func (r MetricMeanInRangeRangeAttributes) Min() terra.NumberValue {
	return terra.ReferenceAsNumber(r.ref.Append("min"))
}

type MetricSumInRangeAttributes struct {
	ref terra.Reference
}

func (msir MetricSumInRangeAttributes) InternalRef() (terra.Reference, error) {
	return msir.ref, nil
}

func (msir MetricSumInRangeAttributes) InternalWithRef(ref terra.Reference) MetricSumInRangeAttributes {
	return MetricSumInRangeAttributes{ref: ref}
}

func (msir MetricSumInRangeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return msir.ref.InternalTokens()
}

func (msir MetricSumInRangeAttributes) TimeSeries() terra.StringValue {
	return terra.ReferenceAsString(msir.ref.Append("time_series"))
}

func (msir MetricSumInRangeAttributes) Range() terra.ListValue[MetricSumInRangeRangeAttributes] {
	return terra.ReferenceAsList[MetricSumInRangeRangeAttributes](msir.ref.Append("range"))
}

type MetricSumInRangeRangeAttributes struct {
	ref terra.Reference
}

func (r MetricSumInRangeRangeAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r MetricSumInRangeRangeAttributes) InternalWithRef(ref terra.Reference) MetricSumInRangeRangeAttributes {
	return MetricSumInRangeRangeAttributes{ref: ref}
}

func (r MetricSumInRangeRangeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r MetricSumInRangeRangeAttributes) Max() terra.NumberValue {
	return terra.ReferenceAsNumber(r.ref.Append("max"))
}

func (r MetricSumInRangeRangeAttributes) Min() terra.NumberValue {
	return terra.ReferenceAsNumber(r.ref.Append("min"))
}

type BasicSliState struct {
	Location     []string                    `json:"location"`
	Method       []string                    `json:"method"`
	Version      []string                    `json:"version"`
	Availability []BasicSliAvailabilityState `json:"availability"`
	Latency      []BasicSliLatencyState      `json:"latency"`
}

type BasicSliAvailabilityState struct {
	Enabled bool `json:"enabled"`
}

type BasicSliLatencyState struct {
	Threshold string `json:"threshold"`
}

type RequestBasedSliState struct {
	DistributionCut []RequestBasedSliDistributionCutState `json:"distribution_cut"`
	GoodTotalRatio  []RequestBasedSliGoodTotalRatioState  `json:"good_total_ratio"`
}

type RequestBasedSliDistributionCutState struct {
	DistributionFilter string                                     `json:"distribution_filter"`
	Range              []RequestBasedSliDistributionCutRangeState `json:"range"`
}

type RequestBasedSliDistributionCutRangeState struct {
	Max float64 `json:"max"`
	Min float64 `json:"min"`
}

type RequestBasedSliGoodTotalRatioState struct {
	BadServiceFilter   string `json:"bad_service_filter"`
	GoodServiceFilter  string `json:"good_service_filter"`
	TotalServiceFilter string `json:"total_service_filter"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

type WindowsBasedSliState struct {
	GoodBadMetricFilter     string                         `json:"good_bad_metric_filter"`
	WindowPeriod            string                         `json:"window_period"`
	GoodTotalRatioThreshold []GoodTotalRatioThresholdState `json:"good_total_ratio_threshold"`
	MetricMeanInRange       []MetricMeanInRangeState       `json:"metric_mean_in_range"`
	MetricSumInRange        []MetricSumInRangeState        `json:"metric_sum_in_range"`
}

type GoodTotalRatioThresholdState struct {
	Threshold           float64                    `json:"threshold"`
	BasicSliPerformance []BasicSliPerformanceState `json:"basic_sli_performance"`
	Performance         []PerformanceState         `json:"performance"`
}

type BasicSliPerformanceState struct {
	Location     []string                               `json:"location"`
	Method       []string                               `json:"method"`
	Version      []string                               `json:"version"`
	Availability []BasicSliPerformanceAvailabilityState `json:"availability"`
	Latency      []BasicSliPerformanceLatencyState      `json:"latency"`
}

type BasicSliPerformanceAvailabilityState struct {
	Enabled bool `json:"enabled"`
}

type BasicSliPerformanceLatencyState struct {
	Threshold string `json:"threshold"`
}

type PerformanceState struct {
	DistributionCut []PerformanceDistributionCutState `json:"distribution_cut"`
	GoodTotalRatio  []PerformanceGoodTotalRatioState  `json:"good_total_ratio"`
}

type PerformanceDistributionCutState struct {
	DistributionFilter string                                 `json:"distribution_filter"`
	Range              []PerformanceDistributionCutRangeState `json:"range"`
}

type PerformanceDistributionCutRangeState struct {
	Max float64 `json:"max"`
	Min float64 `json:"min"`
}

type PerformanceGoodTotalRatioState struct {
	BadServiceFilter   string `json:"bad_service_filter"`
	GoodServiceFilter  string `json:"good_service_filter"`
	TotalServiceFilter string `json:"total_service_filter"`
}

type MetricMeanInRangeState struct {
	TimeSeries string                        `json:"time_series"`
	Range      []MetricMeanInRangeRangeState `json:"range"`
}

type MetricMeanInRangeRangeState struct {
	Max float64 `json:"max"`
	Min float64 `json:"min"`
}

type MetricSumInRangeState struct {
	TimeSeries string                       `json:"time_series"`
	Range      []MetricSumInRangeRangeState `json:"range"`
}

type MetricSumInRangeRangeState struct {
	Max float64 `json:"max"`
	Min float64 `json:"min"`
}
