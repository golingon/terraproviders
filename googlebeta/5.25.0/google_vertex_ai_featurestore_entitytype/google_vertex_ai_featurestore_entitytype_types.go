// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_vertex_ai_featurestore_entitytype

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type MonitoringConfig struct {
	// MonitoringConfigCategoricalThresholdConfig: optional
	CategoricalThresholdConfig *MonitoringConfigCategoricalThresholdConfig `hcl:"categorical_threshold_config,block"`
	// MonitoringConfigImportFeaturesAnalysis: optional
	ImportFeaturesAnalysis *MonitoringConfigImportFeaturesAnalysis `hcl:"import_features_analysis,block"`
	// MonitoringConfigNumericalThresholdConfig: optional
	NumericalThresholdConfig *MonitoringConfigNumericalThresholdConfig `hcl:"numerical_threshold_config,block"`
	// MonitoringConfigSnapshotAnalysis: optional
	SnapshotAnalysis *MonitoringConfigSnapshotAnalysis `hcl:"snapshot_analysis,block"`
}

type MonitoringConfigCategoricalThresholdConfig struct {
	// Value: number, required
	Value terra.NumberValue `hcl:"value,attr" validate:"required"`
}

type MonitoringConfigImportFeaturesAnalysis struct {
	// AnomalyDetectionBaseline: string, optional
	AnomalyDetectionBaseline terra.StringValue `hcl:"anomaly_detection_baseline,attr"`
	// State: string, optional
	State terra.StringValue `hcl:"state,attr"`
}

type MonitoringConfigNumericalThresholdConfig struct {
	// Value: number, required
	Value terra.NumberValue `hcl:"value,attr" validate:"required"`
}

type MonitoringConfigSnapshotAnalysis struct {
	// Disabled: bool, optional
	Disabled terra.BoolValue `hcl:"disabled,attr"`
	// MonitoringInterval: string, optional
	MonitoringInterval terra.StringValue `hcl:"monitoring_interval,attr"`
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

func (mc MonitoringConfigAttributes) CategoricalThresholdConfig() terra.ListValue[MonitoringConfigCategoricalThresholdConfigAttributes] {
	return terra.ReferenceAsList[MonitoringConfigCategoricalThresholdConfigAttributes](mc.ref.Append("categorical_threshold_config"))
}

func (mc MonitoringConfigAttributes) ImportFeaturesAnalysis() terra.ListValue[MonitoringConfigImportFeaturesAnalysisAttributes] {
	return terra.ReferenceAsList[MonitoringConfigImportFeaturesAnalysisAttributes](mc.ref.Append("import_features_analysis"))
}

func (mc MonitoringConfigAttributes) NumericalThresholdConfig() terra.ListValue[MonitoringConfigNumericalThresholdConfigAttributes] {
	return terra.ReferenceAsList[MonitoringConfigNumericalThresholdConfigAttributes](mc.ref.Append("numerical_threshold_config"))
}

func (mc MonitoringConfigAttributes) SnapshotAnalysis() terra.ListValue[MonitoringConfigSnapshotAnalysisAttributes] {
	return terra.ReferenceAsList[MonitoringConfigSnapshotAnalysisAttributes](mc.ref.Append("snapshot_analysis"))
}

type MonitoringConfigCategoricalThresholdConfigAttributes struct {
	ref terra.Reference
}

func (ctc MonitoringConfigCategoricalThresholdConfigAttributes) InternalRef() (terra.Reference, error) {
	return ctc.ref, nil
}

func (ctc MonitoringConfigCategoricalThresholdConfigAttributes) InternalWithRef(ref terra.Reference) MonitoringConfigCategoricalThresholdConfigAttributes {
	return MonitoringConfigCategoricalThresholdConfigAttributes{ref: ref}
}

func (ctc MonitoringConfigCategoricalThresholdConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ctc.ref.InternalTokens()
}

func (ctc MonitoringConfigCategoricalThresholdConfigAttributes) Value() terra.NumberValue {
	return terra.ReferenceAsNumber(ctc.ref.Append("value"))
}

type MonitoringConfigImportFeaturesAnalysisAttributes struct {
	ref terra.Reference
}

func (ifa MonitoringConfigImportFeaturesAnalysisAttributes) InternalRef() (terra.Reference, error) {
	return ifa.ref, nil
}

func (ifa MonitoringConfigImportFeaturesAnalysisAttributes) InternalWithRef(ref terra.Reference) MonitoringConfigImportFeaturesAnalysisAttributes {
	return MonitoringConfigImportFeaturesAnalysisAttributes{ref: ref}
}

func (ifa MonitoringConfigImportFeaturesAnalysisAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ifa.ref.InternalTokens()
}

func (ifa MonitoringConfigImportFeaturesAnalysisAttributes) AnomalyDetectionBaseline() terra.StringValue {
	return terra.ReferenceAsString(ifa.ref.Append("anomaly_detection_baseline"))
}

func (ifa MonitoringConfigImportFeaturesAnalysisAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(ifa.ref.Append("state"))
}

type MonitoringConfigNumericalThresholdConfigAttributes struct {
	ref terra.Reference
}

func (ntc MonitoringConfigNumericalThresholdConfigAttributes) InternalRef() (terra.Reference, error) {
	return ntc.ref, nil
}

func (ntc MonitoringConfigNumericalThresholdConfigAttributes) InternalWithRef(ref terra.Reference) MonitoringConfigNumericalThresholdConfigAttributes {
	return MonitoringConfigNumericalThresholdConfigAttributes{ref: ref}
}

func (ntc MonitoringConfigNumericalThresholdConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ntc.ref.InternalTokens()
}

func (ntc MonitoringConfigNumericalThresholdConfigAttributes) Value() terra.NumberValue {
	return terra.ReferenceAsNumber(ntc.ref.Append("value"))
}

type MonitoringConfigSnapshotAnalysisAttributes struct {
	ref terra.Reference
}

func (sa MonitoringConfigSnapshotAnalysisAttributes) InternalRef() (terra.Reference, error) {
	return sa.ref, nil
}

func (sa MonitoringConfigSnapshotAnalysisAttributes) InternalWithRef(ref terra.Reference) MonitoringConfigSnapshotAnalysisAttributes {
	return MonitoringConfigSnapshotAnalysisAttributes{ref: ref}
}

func (sa MonitoringConfigSnapshotAnalysisAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sa.ref.InternalTokens()
}

func (sa MonitoringConfigSnapshotAnalysisAttributes) Disabled() terra.BoolValue {
	return terra.ReferenceAsBool(sa.ref.Append("disabled"))
}

func (sa MonitoringConfigSnapshotAnalysisAttributes) MonitoringInterval() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("monitoring_interval"))
}

func (sa MonitoringConfigSnapshotAnalysisAttributes) MonitoringIntervalDays() terra.NumberValue {
	return terra.ReferenceAsNumber(sa.ref.Append("monitoring_interval_days"))
}

func (sa MonitoringConfigSnapshotAnalysisAttributes) StalenessDays() terra.NumberValue {
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
	CategoricalThresholdConfig []MonitoringConfigCategoricalThresholdConfigState `json:"categorical_threshold_config"`
	ImportFeaturesAnalysis     []MonitoringConfigImportFeaturesAnalysisState     `json:"import_features_analysis"`
	NumericalThresholdConfig   []MonitoringConfigNumericalThresholdConfigState   `json:"numerical_threshold_config"`
	SnapshotAnalysis           []MonitoringConfigSnapshotAnalysisState           `json:"snapshot_analysis"`
}

type MonitoringConfigCategoricalThresholdConfigState struct {
	Value float64 `json:"value"`
}

type MonitoringConfigImportFeaturesAnalysisState struct {
	AnomalyDetectionBaseline string `json:"anomaly_detection_baseline"`
	State                    string `json:"state"`
}

type MonitoringConfigNumericalThresholdConfigState struct {
	Value float64 `json:"value"`
}

type MonitoringConfigSnapshotAnalysisState struct {
	Disabled               bool    `json:"disabled"`
	MonitoringInterval     string  `json:"monitoring_interval"`
	MonitoringIntervalDays float64 `json:"monitoring_interval_days"`
	StalenessDays          float64 `json:"staleness_days"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
