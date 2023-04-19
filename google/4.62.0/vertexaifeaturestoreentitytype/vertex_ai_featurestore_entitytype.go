// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vertexaifeaturestoreentitytype

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type MonitoringConfig struct {
	// CategoricalThresholdConfig: optional
	CategoricalThresholdConfig *CategoricalThresholdConfig `hcl:"categorical_threshold_config,block"`
	// ImportFeaturesAnalysis: optional
	ImportFeaturesAnalysis *ImportFeaturesAnalysis `hcl:"import_features_analysis,block"`
	// NumericalThresholdConfig: optional
	NumericalThresholdConfig *NumericalThresholdConfig `hcl:"numerical_threshold_config,block"`
	// SnapshotAnalysis: optional
	SnapshotAnalysis *SnapshotAnalysis `hcl:"snapshot_analysis,block"`
}

type CategoricalThresholdConfig struct {
	// Value: number, required
	Value terra.NumberValue `hcl:"value,attr" validate:"required"`
}

type ImportFeaturesAnalysis struct {
	// AnomalyDetectionBaseline: string, optional
	AnomalyDetectionBaseline terra.StringValue `hcl:"anomaly_detection_baseline,attr"`
	// State: string, optional
	State terra.StringValue `hcl:"state,attr"`
}

type NumericalThresholdConfig struct {
	// Value: number, required
	Value terra.NumberValue `hcl:"value,attr" validate:"required"`
}

type SnapshotAnalysis struct {
	// Disabled: bool, optional
	Disabled terra.BoolValue `hcl:"disabled,attr"`
	// MonitoringIntervalDays: number, optional
	MonitoringIntervalDays terra.NumberValue `hcl:"monitoring_interval_days,attr"`
	// StalenessDays: number, optional
	StalenessDays terra.NumberValue `hcl:"staleness_days,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type MonitoringConfigAttributes struct {
	ref terra.Reference
}

func (mc MonitoringConfigAttributes) InternalRef() (terra.Reference, error) {
	return mc.ref, nil
}

func (mc MonitoringConfigAttributes) InternalWithRef(ref terra.Reference) MonitoringConfigAttributes {
	return MonitoringConfigAttributes{ref: ref}
}

func (mc MonitoringConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mc.ref.InternalTokens()
}

func (mc MonitoringConfigAttributes) CategoricalThresholdConfig() terra.ListValue[CategoricalThresholdConfigAttributes] {
	return terra.ReferenceAsList[CategoricalThresholdConfigAttributes](mc.ref.Append("categorical_threshold_config"))
}

func (mc MonitoringConfigAttributes) ImportFeaturesAnalysis() terra.ListValue[ImportFeaturesAnalysisAttributes] {
	return terra.ReferenceAsList[ImportFeaturesAnalysisAttributes](mc.ref.Append("import_features_analysis"))
}

func (mc MonitoringConfigAttributes) NumericalThresholdConfig() terra.ListValue[NumericalThresholdConfigAttributes] {
	return terra.ReferenceAsList[NumericalThresholdConfigAttributes](mc.ref.Append("numerical_threshold_config"))
}

func (mc MonitoringConfigAttributes) SnapshotAnalysis() terra.ListValue[SnapshotAnalysisAttributes] {
	return terra.ReferenceAsList[SnapshotAnalysisAttributes](mc.ref.Append("snapshot_analysis"))
}

type CategoricalThresholdConfigAttributes struct {
	ref terra.Reference
}

func (ctc CategoricalThresholdConfigAttributes) InternalRef() (terra.Reference, error) {
	return ctc.ref, nil
}

func (ctc CategoricalThresholdConfigAttributes) InternalWithRef(ref terra.Reference) CategoricalThresholdConfigAttributes {
	return CategoricalThresholdConfigAttributes{ref: ref}
}

func (ctc CategoricalThresholdConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ctc.ref.InternalTokens()
}

func (ctc CategoricalThresholdConfigAttributes) Value() terra.NumberValue {
	return terra.ReferenceAsNumber(ctc.ref.Append("value"))
}

type ImportFeaturesAnalysisAttributes struct {
	ref terra.Reference
}

func (ifa ImportFeaturesAnalysisAttributes) InternalRef() (terra.Reference, error) {
	return ifa.ref, nil
}

func (ifa ImportFeaturesAnalysisAttributes) InternalWithRef(ref terra.Reference) ImportFeaturesAnalysisAttributes {
	return ImportFeaturesAnalysisAttributes{ref: ref}
}

func (ifa ImportFeaturesAnalysisAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ifa.ref.InternalTokens()
}

func (ifa ImportFeaturesAnalysisAttributes) AnomalyDetectionBaseline() terra.StringValue {
	return terra.ReferenceAsString(ifa.ref.Append("anomaly_detection_baseline"))
}

func (ifa ImportFeaturesAnalysisAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(ifa.ref.Append("state"))
}

type NumericalThresholdConfigAttributes struct {
	ref terra.Reference
}

func (ntc NumericalThresholdConfigAttributes) InternalRef() (terra.Reference, error) {
	return ntc.ref, nil
}

func (ntc NumericalThresholdConfigAttributes) InternalWithRef(ref terra.Reference) NumericalThresholdConfigAttributes {
	return NumericalThresholdConfigAttributes{ref: ref}
}

func (ntc NumericalThresholdConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ntc.ref.InternalTokens()
}

func (ntc NumericalThresholdConfigAttributes) Value() terra.NumberValue {
	return terra.ReferenceAsNumber(ntc.ref.Append("value"))
}

type SnapshotAnalysisAttributes struct {
	ref terra.Reference
}

func (sa SnapshotAnalysisAttributes) InternalRef() (terra.Reference, error) {
	return sa.ref, nil
}

func (sa SnapshotAnalysisAttributes) InternalWithRef(ref terra.Reference) SnapshotAnalysisAttributes {
	return SnapshotAnalysisAttributes{ref: ref}
}

func (sa SnapshotAnalysisAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sa.ref.InternalTokens()
}

func (sa SnapshotAnalysisAttributes) Disabled() terra.BoolValue {
	return terra.ReferenceAsBool(sa.ref.Append("disabled"))
}

func (sa SnapshotAnalysisAttributes) MonitoringIntervalDays() terra.NumberValue {
	return terra.ReferenceAsNumber(sa.ref.Append("monitoring_interval_days"))
}

func (sa SnapshotAnalysisAttributes) StalenessDays() terra.NumberValue {
	return terra.ReferenceAsNumber(sa.ref.Append("staleness_days"))
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

type MonitoringConfigState struct {
	CategoricalThresholdConfig []CategoricalThresholdConfigState `json:"categorical_threshold_config"`
	ImportFeaturesAnalysis     []ImportFeaturesAnalysisState     `json:"import_features_analysis"`
	NumericalThresholdConfig   []NumericalThresholdConfigState   `json:"numerical_threshold_config"`
	SnapshotAnalysis           []SnapshotAnalysisState           `json:"snapshot_analysis"`
}

type CategoricalThresholdConfigState struct {
	Value float64 `json:"value"`
}

type ImportFeaturesAnalysisState struct {
	AnomalyDetectionBaseline string `json:"anomaly_detection_baseline"`
	State                    string `json:"state"`
}

type NumericalThresholdConfigState struct {
	Value float64 `json:"value"`
}

type SnapshotAnalysisState struct {
	Disabled               bool    `json:"disabled"`
	MonitoringIntervalDays float64 `json:"monitoring_interval_days"`
	StalenessDays          float64 `json:"staleness_days"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
