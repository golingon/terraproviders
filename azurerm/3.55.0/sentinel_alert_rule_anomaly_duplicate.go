// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	sentinelalertruleanomalyduplicate "github.com/golingon/terraproviders/azurerm/3.55.0/sentinelalertruleanomalyduplicate"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSentinelAlertRuleAnomalyDuplicate creates a new instance of [SentinelAlertRuleAnomalyDuplicate].
func NewSentinelAlertRuleAnomalyDuplicate(name string, args SentinelAlertRuleAnomalyDuplicateArgs) *SentinelAlertRuleAnomalyDuplicate {
	return &SentinelAlertRuleAnomalyDuplicate{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SentinelAlertRuleAnomalyDuplicate)(nil)

// SentinelAlertRuleAnomalyDuplicate represents the Terraform resource azurerm_sentinel_alert_rule_anomaly_duplicate.
type SentinelAlertRuleAnomalyDuplicate struct {
	Name      string
	Args      SentinelAlertRuleAnomalyDuplicateArgs
	state     *sentinelAlertRuleAnomalyDuplicateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SentinelAlertRuleAnomalyDuplicate].
func (sarad *SentinelAlertRuleAnomalyDuplicate) Type() string {
	return "azurerm_sentinel_alert_rule_anomaly_duplicate"
}

// LocalName returns the local name for [SentinelAlertRuleAnomalyDuplicate].
func (sarad *SentinelAlertRuleAnomalyDuplicate) LocalName() string {
	return sarad.Name
}

// Configuration returns the configuration (args) for [SentinelAlertRuleAnomalyDuplicate].
func (sarad *SentinelAlertRuleAnomalyDuplicate) Configuration() interface{} {
	return sarad.Args
}

// DependOn is used for other resources to depend on [SentinelAlertRuleAnomalyDuplicate].
func (sarad *SentinelAlertRuleAnomalyDuplicate) DependOn() terra.Reference {
	return terra.ReferenceResource(sarad)
}

// Dependencies returns the list of resources [SentinelAlertRuleAnomalyDuplicate] depends_on.
func (sarad *SentinelAlertRuleAnomalyDuplicate) Dependencies() terra.Dependencies {
	return sarad.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SentinelAlertRuleAnomalyDuplicate].
func (sarad *SentinelAlertRuleAnomalyDuplicate) LifecycleManagement() *terra.Lifecycle {
	return sarad.Lifecycle
}

// Attributes returns the attributes for [SentinelAlertRuleAnomalyDuplicate].
func (sarad *SentinelAlertRuleAnomalyDuplicate) Attributes() sentinelAlertRuleAnomalyDuplicateAttributes {
	return sentinelAlertRuleAnomalyDuplicateAttributes{ref: terra.ReferenceResource(sarad)}
}

// ImportState imports the given attribute values into [SentinelAlertRuleAnomalyDuplicate]'s state.
func (sarad *SentinelAlertRuleAnomalyDuplicate) ImportState(av io.Reader) error {
	sarad.state = &sentinelAlertRuleAnomalyDuplicateState{}
	if err := json.NewDecoder(av).Decode(sarad.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sarad.Type(), sarad.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SentinelAlertRuleAnomalyDuplicate] has state.
func (sarad *SentinelAlertRuleAnomalyDuplicate) State() (*sentinelAlertRuleAnomalyDuplicateState, bool) {
	return sarad.state, sarad.state != nil
}

// StateMust returns the state for [SentinelAlertRuleAnomalyDuplicate]. Panics if the state is nil.
func (sarad *SentinelAlertRuleAnomalyDuplicate) StateMust() *sentinelAlertRuleAnomalyDuplicateState {
	if sarad.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sarad.Type(), sarad.LocalName()))
	}
	return sarad.state
}

// SentinelAlertRuleAnomalyDuplicateArgs contains the configurations for azurerm_sentinel_alert_rule_anomaly_duplicate.
type SentinelAlertRuleAnomalyDuplicateArgs struct {
	// BuiltInRuleId: string, required
	BuiltInRuleId terra.StringValue `hcl:"built_in_rule_id,attr" validate:"required"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Enabled: bool, required
	Enabled terra.BoolValue `hcl:"enabled,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogAnalyticsWorkspaceId: string, required
	LogAnalyticsWorkspaceId terra.StringValue `hcl:"log_analytics_workspace_id,attr" validate:"required"`
	// Mode: string, required
	Mode terra.StringValue `hcl:"mode,attr" validate:"required"`
	// RequiredDataConnector: min=0
	RequiredDataConnector []sentinelalertruleanomalyduplicate.RequiredDataConnector `hcl:"required_data_connector,block" validate:"min=0"`
	// MultiSelectObservation: min=0
	MultiSelectObservation []sentinelalertruleanomalyduplicate.MultiSelectObservation `hcl:"multi_select_observation,block" validate:"min=0"`
	// PrioritizedExcludeObservation: min=0
	PrioritizedExcludeObservation []sentinelalertruleanomalyduplicate.PrioritizedExcludeObservation `hcl:"prioritized_exclude_observation,block" validate:"min=0"`
	// SingleSelectObservation: min=0
	SingleSelectObservation []sentinelalertruleanomalyduplicate.SingleSelectObservation `hcl:"single_select_observation,block" validate:"min=0"`
	// ThresholdObservation: min=0
	ThresholdObservation []sentinelalertruleanomalyduplicate.ThresholdObservation `hcl:"threshold_observation,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *sentinelalertruleanomalyduplicate.Timeouts `hcl:"timeouts,block"`
}
type sentinelAlertRuleAnomalyDuplicateAttributes struct {
	ref terra.Reference
}

// AnomalySettingsVersion returns a reference to field anomaly_settings_version of azurerm_sentinel_alert_rule_anomaly_duplicate.
func (sarad sentinelAlertRuleAnomalyDuplicateAttributes) AnomalySettingsVersion() terra.NumberValue {
	return terra.ReferenceAsNumber(sarad.ref.Append("anomaly_settings_version"))
}

// AnomalyVersion returns a reference to field anomaly_version of azurerm_sentinel_alert_rule_anomaly_duplicate.
func (sarad sentinelAlertRuleAnomalyDuplicateAttributes) AnomalyVersion() terra.StringValue {
	return terra.ReferenceAsString(sarad.ref.Append("anomaly_version"))
}

// BuiltInRuleId returns a reference to field built_in_rule_id of azurerm_sentinel_alert_rule_anomaly_duplicate.
func (sarad sentinelAlertRuleAnomalyDuplicateAttributes) BuiltInRuleId() terra.StringValue {
	return terra.ReferenceAsString(sarad.ref.Append("built_in_rule_id"))
}

// Description returns a reference to field description of azurerm_sentinel_alert_rule_anomaly_duplicate.
func (sarad sentinelAlertRuleAnomalyDuplicateAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(sarad.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_sentinel_alert_rule_anomaly_duplicate.
func (sarad sentinelAlertRuleAnomalyDuplicateAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(sarad.ref.Append("display_name"))
}

// Enabled returns a reference to field enabled of azurerm_sentinel_alert_rule_anomaly_duplicate.
func (sarad sentinelAlertRuleAnomalyDuplicateAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(sarad.ref.Append("enabled"))
}

// Frequency returns a reference to field frequency of azurerm_sentinel_alert_rule_anomaly_duplicate.
func (sarad sentinelAlertRuleAnomalyDuplicateAttributes) Frequency() terra.StringValue {
	return terra.ReferenceAsString(sarad.ref.Append("frequency"))
}

// Id returns a reference to field id of azurerm_sentinel_alert_rule_anomaly_duplicate.
func (sarad sentinelAlertRuleAnomalyDuplicateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sarad.ref.Append("id"))
}

// IsDefaultSettings returns a reference to field is_default_settings of azurerm_sentinel_alert_rule_anomaly_duplicate.
func (sarad sentinelAlertRuleAnomalyDuplicateAttributes) IsDefaultSettings() terra.BoolValue {
	return terra.ReferenceAsBool(sarad.ref.Append("is_default_settings"))
}

// LogAnalyticsWorkspaceId returns a reference to field log_analytics_workspace_id of azurerm_sentinel_alert_rule_anomaly_duplicate.
func (sarad sentinelAlertRuleAnomalyDuplicateAttributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(sarad.ref.Append("log_analytics_workspace_id"))
}

// Mode returns a reference to field mode of azurerm_sentinel_alert_rule_anomaly_duplicate.
func (sarad sentinelAlertRuleAnomalyDuplicateAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(sarad.ref.Append("mode"))
}

// Name returns a reference to field name of azurerm_sentinel_alert_rule_anomaly_duplicate.
func (sarad sentinelAlertRuleAnomalyDuplicateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sarad.ref.Append("name"))
}

// SettingsDefinitionId returns a reference to field settings_definition_id of azurerm_sentinel_alert_rule_anomaly_duplicate.
func (sarad sentinelAlertRuleAnomalyDuplicateAttributes) SettingsDefinitionId() terra.StringValue {
	return terra.ReferenceAsString(sarad.ref.Append("settings_definition_id"))
}

// Tactics returns a reference to field tactics of azurerm_sentinel_alert_rule_anomaly_duplicate.
func (sarad sentinelAlertRuleAnomalyDuplicateAttributes) Tactics() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sarad.ref.Append("tactics"))
}

// Techniques returns a reference to field techniques of azurerm_sentinel_alert_rule_anomaly_duplicate.
func (sarad sentinelAlertRuleAnomalyDuplicateAttributes) Techniques() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sarad.ref.Append("techniques"))
}

func (sarad sentinelAlertRuleAnomalyDuplicateAttributes) RequiredDataConnector() terra.ListValue[sentinelalertruleanomalyduplicate.RequiredDataConnectorAttributes] {
	return terra.ReferenceAsList[sentinelalertruleanomalyduplicate.RequiredDataConnectorAttributes](sarad.ref.Append("required_data_connector"))
}

func (sarad sentinelAlertRuleAnomalyDuplicateAttributes) MultiSelectObservation() terra.ListValue[sentinelalertruleanomalyduplicate.MultiSelectObservationAttributes] {
	return terra.ReferenceAsList[sentinelalertruleanomalyduplicate.MultiSelectObservationAttributes](sarad.ref.Append("multi_select_observation"))
}

func (sarad sentinelAlertRuleAnomalyDuplicateAttributes) PrioritizedExcludeObservation() terra.ListValue[sentinelalertruleanomalyduplicate.PrioritizedExcludeObservationAttributes] {
	return terra.ReferenceAsList[sentinelalertruleanomalyduplicate.PrioritizedExcludeObservationAttributes](sarad.ref.Append("prioritized_exclude_observation"))
}

func (sarad sentinelAlertRuleAnomalyDuplicateAttributes) SingleSelectObservation() terra.ListValue[sentinelalertruleanomalyduplicate.SingleSelectObservationAttributes] {
	return terra.ReferenceAsList[sentinelalertruleanomalyduplicate.SingleSelectObservationAttributes](sarad.ref.Append("single_select_observation"))
}

func (sarad sentinelAlertRuleAnomalyDuplicateAttributes) ThresholdObservation() terra.ListValue[sentinelalertruleanomalyduplicate.ThresholdObservationAttributes] {
	return terra.ReferenceAsList[sentinelalertruleanomalyduplicate.ThresholdObservationAttributes](sarad.ref.Append("threshold_observation"))
}

func (sarad sentinelAlertRuleAnomalyDuplicateAttributes) Timeouts() sentinelalertruleanomalyduplicate.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sentinelalertruleanomalyduplicate.TimeoutsAttributes](sarad.ref.Append("timeouts"))
}

type sentinelAlertRuleAnomalyDuplicateState struct {
	AnomalySettingsVersion        float64                                                                `json:"anomaly_settings_version"`
	AnomalyVersion                string                                                                 `json:"anomaly_version"`
	BuiltInRuleId                 string                                                                 `json:"built_in_rule_id"`
	Description                   string                                                                 `json:"description"`
	DisplayName                   string                                                                 `json:"display_name"`
	Enabled                       bool                                                                   `json:"enabled"`
	Frequency                     string                                                                 `json:"frequency"`
	Id                            string                                                                 `json:"id"`
	IsDefaultSettings             bool                                                                   `json:"is_default_settings"`
	LogAnalyticsWorkspaceId       string                                                                 `json:"log_analytics_workspace_id"`
	Mode                          string                                                                 `json:"mode"`
	Name                          string                                                                 `json:"name"`
	SettingsDefinitionId          string                                                                 `json:"settings_definition_id"`
	Tactics                       []string                                                               `json:"tactics"`
	Techniques                    []string                                                               `json:"techniques"`
	RequiredDataConnector         []sentinelalertruleanomalyduplicate.RequiredDataConnectorState         `json:"required_data_connector"`
	MultiSelectObservation        []sentinelalertruleanomalyduplicate.MultiSelectObservationState        `json:"multi_select_observation"`
	PrioritizedExcludeObservation []sentinelalertruleanomalyduplicate.PrioritizedExcludeObservationState `json:"prioritized_exclude_observation"`
	SingleSelectObservation       []sentinelalertruleanomalyduplicate.SingleSelectObservationState       `json:"single_select_observation"`
	ThresholdObservation          []sentinelalertruleanomalyduplicate.ThresholdObservationState          `json:"threshold_observation"`
	Timeouts                      *sentinelalertruleanomalyduplicate.TimeoutsState                       `json:"timeouts"`
}
