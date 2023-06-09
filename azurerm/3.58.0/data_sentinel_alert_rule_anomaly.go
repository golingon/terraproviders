// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datasentinelalertruleanomaly "github.com/golingon/terraproviders/azurerm/3.58.0/datasentinelalertruleanomaly"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSentinelAlertRuleAnomaly creates a new instance of [DataSentinelAlertRuleAnomaly].
func NewDataSentinelAlertRuleAnomaly(name string, args DataSentinelAlertRuleAnomalyArgs) *DataSentinelAlertRuleAnomaly {
	return &DataSentinelAlertRuleAnomaly{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSentinelAlertRuleAnomaly)(nil)

// DataSentinelAlertRuleAnomaly represents the Terraform data resource azurerm_sentinel_alert_rule_anomaly.
type DataSentinelAlertRuleAnomaly struct {
	Name string
	Args DataSentinelAlertRuleAnomalyArgs
}

// DataSource returns the Terraform object type for [DataSentinelAlertRuleAnomaly].
func (sara *DataSentinelAlertRuleAnomaly) DataSource() string {
	return "azurerm_sentinel_alert_rule_anomaly"
}

// LocalName returns the local name for [DataSentinelAlertRuleAnomaly].
func (sara *DataSentinelAlertRuleAnomaly) LocalName() string {
	return sara.Name
}

// Configuration returns the configuration (args) for [DataSentinelAlertRuleAnomaly].
func (sara *DataSentinelAlertRuleAnomaly) Configuration() interface{} {
	return sara.Args
}

// Attributes returns the attributes for [DataSentinelAlertRuleAnomaly].
func (sara *DataSentinelAlertRuleAnomaly) Attributes() dataSentinelAlertRuleAnomalyAttributes {
	return dataSentinelAlertRuleAnomalyAttributes{ref: terra.ReferenceDataResource(sara)}
}

// DataSentinelAlertRuleAnomalyArgs contains the configurations for azurerm_sentinel_alert_rule_anomaly.
type DataSentinelAlertRuleAnomalyArgs struct {
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogAnalyticsWorkspaceId: string, required
	LogAnalyticsWorkspaceId terra.StringValue `hcl:"log_analytics_workspace_id,attr" validate:"required"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// MultiSelectObservation: min=0
	MultiSelectObservation []datasentinelalertruleanomaly.MultiSelectObservation `hcl:"multi_select_observation,block" validate:"min=0"`
	// PrioritizedExcludeObservation: min=0
	PrioritizedExcludeObservation []datasentinelalertruleanomaly.PrioritizedExcludeObservation `hcl:"prioritized_exclude_observation,block" validate:"min=0"`
	// RequiredDataConnector: min=0
	RequiredDataConnector []datasentinelalertruleanomaly.RequiredDataConnector `hcl:"required_data_connector,block" validate:"min=0"`
	// SingleSelectObservation: min=0
	SingleSelectObservation []datasentinelalertruleanomaly.SingleSelectObservation `hcl:"single_select_observation,block" validate:"min=0"`
	// ThresholdObservation: min=0
	ThresholdObservation []datasentinelalertruleanomaly.ThresholdObservation `hcl:"threshold_observation,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datasentinelalertruleanomaly.Timeouts `hcl:"timeouts,block"`
}
type dataSentinelAlertRuleAnomalyAttributes struct {
	ref terra.Reference
}

// AnomalySettingsVersion returns a reference to field anomaly_settings_version of azurerm_sentinel_alert_rule_anomaly.
func (sara dataSentinelAlertRuleAnomalyAttributes) AnomalySettingsVersion() terra.NumberValue {
	return terra.ReferenceAsNumber(sara.ref.Append("anomaly_settings_version"))
}

// AnomalyVersion returns a reference to field anomaly_version of azurerm_sentinel_alert_rule_anomaly.
func (sara dataSentinelAlertRuleAnomalyAttributes) AnomalyVersion() terra.StringValue {
	return terra.ReferenceAsString(sara.ref.Append("anomaly_version"))
}

// Description returns a reference to field description of azurerm_sentinel_alert_rule_anomaly.
func (sara dataSentinelAlertRuleAnomalyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(sara.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_sentinel_alert_rule_anomaly.
func (sara dataSentinelAlertRuleAnomalyAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(sara.ref.Append("display_name"))
}

// Enabled returns a reference to field enabled of azurerm_sentinel_alert_rule_anomaly.
func (sara dataSentinelAlertRuleAnomalyAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(sara.ref.Append("enabled"))
}

// Frequency returns a reference to field frequency of azurerm_sentinel_alert_rule_anomaly.
func (sara dataSentinelAlertRuleAnomalyAttributes) Frequency() terra.StringValue {
	return terra.ReferenceAsString(sara.ref.Append("frequency"))
}

// Id returns a reference to field id of azurerm_sentinel_alert_rule_anomaly.
func (sara dataSentinelAlertRuleAnomalyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sara.ref.Append("id"))
}

// LogAnalyticsWorkspaceId returns a reference to field log_analytics_workspace_id of azurerm_sentinel_alert_rule_anomaly.
func (sara dataSentinelAlertRuleAnomalyAttributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(sara.ref.Append("log_analytics_workspace_id"))
}

// Mode returns a reference to field mode of azurerm_sentinel_alert_rule_anomaly.
func (sara dataSentinelAlertRuleAnomalyAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(sara.ref.Append("mode"))
}

// Name returns a reference to field name of azurerm_sentinel_alert_rule_anomaly.
func (sara dataSentinelAlertRuleAnomalyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sara.ref.Append("name"))
}

// SettingsDefinitionId returns a reference to field settings_definition_id of azurerm_sentinel_alert_rule_anomaly.
func (sara dataSentinelAlertRuleAnomalyAttributes) SettingsDefinitionId() terra.StringValue {
	return terra.ReferenceAsString(sara.ref.Append("settings_definition_id"))
}

// Tactics returns a reference to field tactics of azurerm_sentinel_alert_rule_anomaly.
func (sara dataSentinelAlertRuleAnomalyAttributes) Tactics() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sara.ref.Append("tactics"))
}

// Techniques returns a reference to field techniques of azurerm_sentinel_alert_rule_anomaly.
func (sara dataSentinelAlertRuleAnomalyAttributes) Techniques() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sara.ref.Append("techniques"))
}

func (sara dataSentinelAlertRuleAnomalyAttributes) MultiSelectObservation() terra.ListValue[datasentinelalertruleanomaly.MultiSelectObservationAttributes] {
	return terra.ReferenceAsList[datasentinelalertruleanomaly.MultiSelectObservationAttributes](sara.ref.Append("multi_select_observation"))
}

func (sara dataSentinelAlertRuleAnomalyAttributes) PrioritizedExcludeObservation() terra.ListValue[datasentinelalertruleanomaly.PrioritizedExcludeObservationAttributes] {
	return terra.ReferenceAsList[datasentinelalertruleanomaly.PrioritizedExcludeObservationAttributes](sara.ref.Append("prioritized_exclude_observation"))
}

func (sara dataSentinelAlertRuleAnomalyAttributes) RequiredDataConnector() terra.ListValue[datasentinelalertruleanomaly.RequiredDataConnectorAttributes] {
	return terra.ReferenceAsList[datasentinelalertruleanomaly.RequiredDataConnectorAttributes](sara.ref.Append("required_data_connector"))
}

func (sara dataSentinelAlertRuleAnomalyAttributes) SingleSelectObservation() terra.ListValue[datasentinelalertruleanomaly.SingleSelectObservationAttributes] {
	return terra.ReferenceAsList[datasentinelalertruleanomaly.SingleSelectObservationAttributes](sara.ref.Append("single_select_observation"))
}

func (sara dataSentinelAlertRuleAnomalyAttributes) ThresholdObservation() terra.ListValue[datasentinelalertruleanomaly.ThresholdObservationAttributes] {
	return terra.ReferenceAsList[datasentinelalertruleanomaly.ThresholdObservationAttributes](sara.ref.Append("threshold_observation"))
}

func (sara dataSentinelAlertRuleAnomalyAttributes) Timeouts() datasentinelalertruleanomaly.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datasentinelalertruleanomaly.TimeoutsAttributes](sara.ref.Append("timeouts"))
}
