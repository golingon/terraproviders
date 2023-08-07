// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	sentinelalertruleanomalybuiltin "github.com/golingon/terraproviders/azurerm/3.68.0/sentinelalertruleanomalybuiltin"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSentinelAlertRuleAnomalyBuiltIn creates a new instance of [SentinelAlertRuleAnomalyBuiltIn].
func NewSentinelAlertRuleAnomalyBuiltIn(name string, args SentinelAlertRuleAnomalyBuiltInArgs) *SentinelAlertRuleAnomalyBuiltIn {
	return &SentinelAlertRuleAnomalyBuiltIn{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SentinelAlertRuleAnomalyBuiltIn)(nil)

// SentinelAlertRuleAnomalyBuiltIn represents the Terraform resource azurerm_sentinel_alert_rule_anomaly_built_in.
type SentinelAlertRuleAnomalyBuiltIn struct {
	Name      string
	Args      SentinelAlertRuleAnomalyBuiltInArgs
	state     *sentinelAlertRuleAnomalyBuiltInState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SentinelAlertRuleAnomalyBuiltIn].
func (sarabi *SentinelAlertRuleAnomalyBuiltIn) Type() string {
	return "azurerm_sentinel_alert_rule_anomaly_built_in"
}

// LocalName returns the local name for [SentinelAlertRuleAnomalyBuiltIn].
func (sarabi *SentinelAlertRuleAnomalyBuiltIn) LocalName() string {
	return sarabi.Name
}

// Configuration returns the configuration (args) for [SentinelAlertRuleAnomalyBuiltIn].
func (sarabi *SentinelAlertRuleAnomalyBuiltIn) Configuration() interface{} {
	return sarabi.Args
}

// DependOn is used for other resources to depend on [SentinelAlertRuleAnomalyBuiltIn].
func (sarabi *SentinelAlertRuleAnomalyBuiltIn) DependOn() terra.Reference {
	return terra.ReferenceResource(sarabi)
}

// Dependencies returns the list of resources [SentinelAlertRuleAnomalyBuiltIn] depends_on.
func (sarabi *SentinelAlertRuleAnomalyBuiltIn) Dependencies() terra.Dependencies {
	return sarabi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SentinelAlertRuleAnomalyBuiltIn].
func (sarabi *SentinelAlertRuleAnomalyBuiltIn) LifecycleManagement() *terra.Lifecycle {
	return sarabi.Lifecycle
}

// Attributes returns the attributes for [SentinelAlertRuleAnomalyBuiltIn].
func (sarabi *SentinelAlertRuleAnomalyBuiltIn) Attributes() sentinelAlertRuleAnomalyBuiltInAttributes {
	return sentinelAlertRuleAnomalyBuiltInAttributes{ref: terra.ReferenceResource(sarabi)}
}

// ImportState imports the given attribute values into [SentinelAlertRuleAnomalyBuiltIn]'s state.
func (sarabi *SentinelAlertRuleAnomalyBuiltIn) ImportState(av io.Reader) error {
	sarabi.state = &sentinelAlertRuleAnomalyBuiltInState{}
	if err := json.NewDecoder(av).Decode(sarabi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sarabi.Type(), sarabi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SentinelAlertRuleAnomalyBuiltIn] has state.
func (sarabi *SentinelAlertRuleAnomalyBuiltIn) State() (*sentinelAlertRuleAnomalyBuiltInState, bool) {
	return sarabi.state, sarabi.state != nil
}

// StateMust returns the state for [SentinelAlertRuleAnomalyBuiltIn]. Panics if the state is nil.
func (sarabi *SentinelAlertRuleAnomalyBuiltIn) StateMust() *sentinelAlertRuleAnomalyBuiltInState {
	if sarabi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sarabi.Type(), sarabi.LocalName()))
	}
	return sarabi.state
}

// SentinelAlertRuleAnomalyBuiltInArgs contains the configurations for azurerm_sentinel_alert_rule_anomaly_built_in.
type SentinelAlertRuleAnomalyBuiltInArgs struct {
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Enabled: bool, required
	Enabled terra.BoolValue `hcl:"enabled,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogAnalyticsWorkspaceId: string, required
	LogAnalyticsWorkspaceId terra.StringValue `hcl:"log_analytics_workspace_id,attr" validate:"required"`
	// Mode: string, required
	Mode terra.StringValue `hcl:"mode,attr" validate:"required"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// MultiSelectObservation: min=0
	MultiSelectObservation []sentinelalertruleanomalybuiltin.MultiSelectObservation `hcl:"multi_select_observation,block" validate:"min=0"`
	// PrioritizedExcludeObservation: min=0
	PrioritizedExcludeObservation []sentinelalertruleanomalybuiltin.PrioritizedExcludeObservation `hcl:"prioritized_exclude_observation,block" validate:"min=0"`
	// RequiredDataConnector: min=0
	RequiredDataConnector []sentinelalertruleanomalybuiltin.RequiredDataConnector `hcl:"required_data_connector,block" validate:"min=0"`
	// SingleSelectObservation: min=0
	SingleSelectObservation []sentinelalertruleanomalybuiltin.SingleSelectObservation `hcl:"single_select_observation,block" validate:"min=0"`
	// ThresholdObservation: min=0
	ThresholdObservation []sentinelalertruleanomalybuiltin.ThresholdObservation `hcl:"threshold_observation,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *sentinelalertruleanomalybuiltin.Timeouts `hcl:"timeouts,block"`
}
type sentinelAlertRuleAnomalyBuiltInAttributes struct {
	ref terra.Reference
}

// AnomalySettingsVersion returns a reference to field anomaly_settings_version of azurerm_sentinel_alert_rule_anomaly_built_in.
func (sarabi sentinelAlertRuleAnomalyBuiltInAttributes) AnomalySettingsVersion() terra.NumberValue {
	return terra.ReferenceAsNumber(sarabi.ref.Append("anomaly_settings_version"))
}

// AnomalyVersion returns a reference to field anomaly_version of azurerm_sentinel_alert_rule_anomaly_built_in.
func (sarabi sentinelAlertRuleAnomalyBuiltInAttributes) AnomalyVersion() terra.StringValue {
	return terra.ReferenceAsString(sarabi.ref.Append("anomaly_version"))
}

// Description returns a reference to field description of azurerm_sentinel_alert_rule_anomaly_built_in.
func (sarabi sentinelAlertRuleAnomalyBuiltInAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(sarabi.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_sentinel_alert_rule_anomaly_built_in.
func (sarabi sentinelAlertRuleAnomalyBuiltInAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(sarabi.ref.Append("display_name"))
}

// Enabled returns a reference to field enabled of azurerm_sentinel_alert_rule_anomaly_built_in.
func (sarabi sentinelAlertRuleAnomalyBuiltInAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(sarabi.ref.Append("enabled"))
}

// Frequency returns a reference to field frequency of azurerm_sentinel_alert_rule_anomaly_built_in.
func (sarabi sentinelAlertRuleAnomalyBuiltInAttributes) Frequency() terra.StringValue {
	return terra.ReferenceAsString(sarabi.ref.Append("frequency"))
}

// Id returns a reference to field id of azurerm_sentinel_alert_rule_anomaly_built_in.
func (sarabi sentinelAlertRuleAnomalyBuiltInAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sarabi.ref.Append("id"))
}

// LogAnalyticsWorkspaceId returns a reference to field log_analytics_workspace_id of azurerm_sentinel_alert_rule_anomaly_built_in.
func (sarabi sentinelAlertRuleAnomalyBuiltInAttributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(sarabi.ref.Append("log_analytics_workspace_id"))
}

// Mode returns a reference to field mode of azurerm_sentinel_alert_rule_anomaly_built_in.
func (sarabi sentinelAlertRuleAnomalyBuiltInAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(sarabi.ref.Append("mode"))
}

// Name returns a reference to field name of azurerm_sentinel_alert_rule_anomaly_built_in.
func (sarabi sentinelAlertRuleAnomalyBuiltInAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sarabi.ref.Append("name"))
}

// SettingsDefinitionId returns a reference to field settings_definition_id of azurerm_sentinel_alert_rule_anomaly_built_in.
func (sarabi sentinelAlertRuleAnomalyBuiltInAttributes) SettingsDefinitionId() terra.StringValue {
	return terra.ReferenceAsString(sarabi.ref.Append("settings_definition_id"))
}

// Tactics returns a reference to field tactics of azurerm_sentinel_alert_rule_anomaly_built_in.
func (sarabi sentinelAlertRuleAnomalyBuiltInAttributes) Tactics() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sarabi.ref.Append("tactics"))
}

// Techniques returns a reference to field techniques of azurerm_sentinel_alert_rule_anomaly_built_in.
func (sarabi sentinelAlertRuleAnomalyBuiltInAttributes) Techniques() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sarabi.ref.Append("techniques"))
}

func (sarabi sentinelAlertRuleAnomalyBuiltInAttributes) MultiSelectObservation() terra.ListValue[sentinelalertruleanomalybuiltin.MultiSelectObservationAttributes] {
	return terra.ReferenceAsList[sentinelalertruleanomalybuiltin.MultiSelectObservationAttributes](sarabi.ref.Append("multi_select_observation"))
}

func (sarabi sentinelAlertRuleAnomalyBuiltInAttributes) PrioritizedExcludeObservation() terra.ListValue[sentinelalertruleanomalybuiltin.PrioritizedExcludeObservationAttributes] {
	return terra.ReferenceAsList[sentinelalertruleanomalybuiltin.PrioritizedExcludeObservationAttributes](sarabi.ref.Append("prioritized_exclude_observation"))
}

func (sarabi sentinelAlertRuleAnomalyBuiltInAttributes) RequiredDataConnector() terra.ListValue[sentinelalertruleanomalybuiltin.RequiredDataConnectorAttributes] {
	return terra.ReferenceAsList[sentinelalertruleanomalybuiltin.RequiredDataConnectorAttributes](sarabi.ref.Append("required_data_connector"))
}

func (sarabi sentinelAlertRuleAnomalyBuiltInAttributes) SingleSelectObservation() terra.ListValue[sentinelalertruleanomalybuiltin.SingleSelectObservationAttributes] {
	return terra.ReferenceAsList[sentinelalertruleanomalybuiltin.SingleSelectObservationAttributes](sarabi.ref.Append("single_select_observation"))
}

func (sarabi sentinelAlertRuleAnomalyBuiltInAttributes) ThresholdObservation() terra.ListValue[sentinelalertruleanomalybuiltin.ThresholdObservationAttributes] {
	return terra.ReferenceAsList[sentinelalertruleanomalybuiltin.ThresholdObservationAttributes](sarabi.ref.Append("threshold_observation"))
}

func (sarabi sentinelAlertRuleAnomalyBuiltInAttributes) Timeouts() sentinelalertruleanomalybuiltin.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sentinelalertruleanomalybuiltin.TimeoutsAttributes](sarabi.ref.Append("timeouts"))
}

type sentinelAlertRuleAnomalyBuiltInState struct {
	AnomalySettingsVersion        float64                                                              `json:"anomaly_settings_version"`
	AnomalyVersion                string                                                               `json:"anomaly_version"`
	Description                   string                                                               `json:"description"`
	DisplayName                   string                                                               `json:"display_name"`
	Enabled                       bool                                                                 `json:"enabled"`
	Frequency                     string                                                               `json:"frequency"`
	Id                            string                                                               `json:"id"`
	LogAnalyticsWorkspaceId       string                                                               `json:"log_analytics_workspace_id"`
	Mode                          string                                                               `json:"mode"`
	Name                          string                                                               `json:"name"`
	SettingsDefinitionId          string                                                               `json:"settings_definition_id"`
	Tactics                       []string                                                             `json:"tactics"`
	Techniques                    []string                                                             `json:"techniques"`
	MultiSelectObservation        []sentinelalertruleanomalybuiltin.MultiSelectObservationState        `json:"multi_select_observation"`
	PrioritizedExcludeObservation []sentinelalertruleanomalybuiltin.PrioritizedExcludeObservationState `json:"prioritized_exclude_observation"`
	RequiredDataConnector         []sentinelalertruleanomalybuiltin.RequiredDataConnectorState         `json:"required_data_connector"`
	SingleSelectObservation       []sentinelalertruleanomalybuiltin.SingleSelectObservationState       `json:"single_select_observation"`
	ThresholdObservation          []sentinelalertruleanomalybuiltin.ThresholdObservationState          `json:"threshold_observation"`
	Timeouts                      *sentinelalertruleanomalybuiltin.TimeoutsState                       `json:"timeouts"`
}
