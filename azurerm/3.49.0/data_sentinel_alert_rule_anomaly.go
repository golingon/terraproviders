// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datasentinelalertruleanomaly "github.com/golingon/terraproviders/azurerm/3.49.0/datasentinelalertruleanomaly"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataSentinelAlertRuleAnomaly(name string, args DataSentinelAlertRuleAnomalyArgs) *DataSentinelAlertRuleAnomaly {
	return &DataSentinelAlertRuleAnomaly{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSentinelAlertRuleAnomaly)(nil)

type DataSentinelAlertRuleAnomaly struct {
	Name string
	Args DataSentinelAlertRuleAnomalyArgs
}

func (sara *DataSentinelAlertRuleAnomaly) DataSource() string {
	return "azurerm_sentinel_alert_rule_anomaly"
}

func (sara *DataSentinelAlertRuleAnomaly) LocalName() string {
	return sara.Name
}

func (sara *DataSentinelAlertRuleAnomaly) Configuration() interface{} {
	return sara.Args
}

func (sara *DataSentinelAlertRuleAnomaly) Attributes() dataSentinelAlertRuleAnomalyAttributes {
	return dataSentinelAlertRuleAnomalyAttributes{ref: terra.ReferenceDataResource(sara)}
}

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

func (sara dataSentinelAlertRuleAnomalyAttributes) AnomalySettingsVersion() terra.NumberValue {
	return terra.ReferenceNumber(sara.ref.Append("anomaly_settings_version"))
}

func (sara dataSentinelAlertRuleAnomalyAttributes) AnomalyVersion() terra.StringValue {
	return terra.ReferenceString(sara.ref.Append("anomaly_version"))
}

func (sara dataSentinelAlertRuleAnomalyAttributes) Description() terra.StringValue {
	return terra.ReferenceString(sara.ref.Append("description"))
}

func (sara dataSentinelAlertRuleAnomalyAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceString(sara.ref.Append("display_name"))
}

func (sara dataSentinelAlertRuleAnomalyAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceBool(sara.ref.Append("enabled"))
}

func (sara dataSentinelAlertRuleAnomalyAttributes) Frequency() terra.StringValue {
	return terra.ReferenceString(sara.ref.Append("frequency"))
}

func (sara dataSentinelAlertRuleAnomalyAttributes) Id() terra.StringValue {
	return terra.ReferenceString(sara.ref.Append("id"))
}

func (sara dataSentinelAlertRuleAnomalyAttributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceString(sara.ref.Append("log_analytics_workspace_id"))
}

func (sara dataSentinelAlertRuleAnomalyAttributes) Mode() terra.StringValue {
	return terra.ReferenceString(sara.ref.Append("mode"))
}

func (sara dataSentinelAlertRuleAnomalyAttributes) Name() terra.StringValue {
	return terra.ReferenceString(sara.ref.Append("name"))
}

func (sara dataSentinelAlertRuleAnomalyAttributes) SettingsDefinitionId() terra.StringValue {
	return terra.ReferenceString(sara.ref.Append("settings_definition_id"))
}

func (sara dataSentinelAlertRuleAnomalyAttributes) Tactics() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](sara.ref.Append("tactics"))
}

func (sara dataSentinelAlertRuleAnomalyAttributes) Techniques() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](sara.ref.Append("techniques"))
}

func (sara dataSentinelAlertRuleAnomalyAttributes) MultiSelectObservation() terra.ListValue[datasentinelalertruleanomaly.MultiSelectObservationAttributes] {
	return terra.ReferenceList[datasentinelalertruleanomaly.MultiSelectObservationAttributes](sara.ref.Append("multi_select_observation"))
}

func (sara dataSentinelAlertRuleAnomalyAttributes) PrioritizedExcludeObservation() terra.ListValue[datasentinelalertruleanomaly.PrioritizedExcludeObservationAttributes] {
	return terra.ReferenceList[datasentinelalertruleanomaly.PrioritizedExcludeObservationAttributes](sara.ref.Append("prioritized_exclude_observation"))
}

func (sara dataSentinelAlertRuleAnomalyAttributes) RequiredDataConnector() terra.ListValue[datasentinelalertruleanomaly.RequiredDataConnectorAttributes] {
	return terra.ReferenceList[datasentinelalertruleanomaly.RequiredDataConnectorAttributes](sara.ref.Append("required_data_connector"))
}

func (sara dataSentinelAlertRuleAnomalyAttributes) SingleSelectObservation() terra.ListValue[datasentinelalertruleanomaly.SingleSelectObservationAttributes] {
	return terra.ReferenceList[datasentinelalertruleanomaly.SingleSelectObservationAttributes](sara.ref.Append("single_select_observation"))
}

func (sara dataSentinelAlertRuleAnomalyAttributes) ThresholdObservation() terra.ListValue[datasentinelalertruleanomaly.ThresholdObservationAttributes] {
	return terra.ReferenceList[datasentinelalertruleanomaly.ThresholdObservationAttributes](sara.ref.Append("threshold_observation"))
}

func (sara dataSentinelAlertRuleAnomalyAttributes) Timeouts() datasentinelalertruleanomaly.TimeoutsAttributes {
	return terra.ReferenceSingle[datasentinelalertruleanomaly.TimeoutsAttributes](sara.ref.Append("timeouts"))
}
