// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	monitorscheduledqueryrulesalertv2 "github.com/golingon/terraproviders/azurerm/3.58.0/monitorscheduledqueryrulesalertv2"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMonitorScheduledQueryRulesAlertV2 creates a new instance of [MonitorScheduledQueryRulesAlertV2].
func NewMonitorScheduledQueryRulesAlertV2(name string, args MonitorScheduledQueryRulesAlertV2Args) *MonitorScheduledQueryRulesAlertV2 {
	return &MonitorScheduledQueryRulesAlertV2{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MonitorScheduledQueryRulesAlertV2)(nil)

// MonitorScheduledQueryRulesAlertV2 represents the Terraform resource azurerm_monitor_scheduled_query_rules_alert_v2.
type MonitorScheduledQueryRulesAlertV2 struct {
	Name      string
	Args      MonitorScheduledQueryRulesAlertV2Args
	state     *monitorScheduledQueryRulesAlertV2State
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MonitorScheduledQueryRulesAlertV2].
func (msqrav *MonitorScheduledQueryRulesAlertV2) Type() string {
	return "azurerm_monitor_scheduled_query_rules_alert_v2"
}

// LocalName returns the local name for [MonitorScheduledQueryRulesAlertV2].
func (msqrav *MonitorScheduledQueryRulesAlertV2) LocalName() string {
	return msqrav.Name
}

// Configuration returns the configuration (args) for [MonitorScheduledQueryRulesAlertV2].
func (msqrav *MonitorScheduledQueryRulesAlertV2) Configuration() interface{} {
	return msqrav.Args
}

// DependOn is used for other resources to depend on [MonitorScheduledQueryRulesAlertV2].
func (msqrav *MonitorScheduledQueryRulesAlertV2) DependOn() terra.Reference {
	return terra.ReferenceResource(msqrav)
}

// Dependencies returns the list of resources [MonitorScheduledQueryRulesAlertV2] depends_on.
func (msqrav *MonitorScheduledQueryRulesAlertV2) Dependencies() terra.Dependencies {
	return msqrav.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MonitorScheduledQueryRulesAlertV2].
func (msqrav *MonitorScheduledQueryRulesAlertV2) LifecycleManagement() *terra.Lifecycle {
	return msqrav.Lifecycle
}

// Attributes returns the attributes for [MonitorScheduledQueryRulesAlertV2].
func (msqrav *MonitorScheduledQueryRulesAlertV2) Attributes() monitorScheduledQueryRulesAlertV2Attributes {
	return monitorScheduledQueryRulesAlertV2Attributes{ref: terra.ReferenceResource(msqrav)}
}

// ImportState imports the given attribute values into [MonitorScheduledQueryRulesAlertV2]'s state.
func (msqrav *MonitorScheduledQueryRulesAlertV2) ImportState(av io.Reader) error {
	msqrav.state = &monitorScheduledQueryRulesAlertV2State{}
	if err := json.NewDecoder(av).Decode(msqrav.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", msqrav.Type(), msqrav.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MonitorScheduledQueryRulesAlertV2] has state.
func (msqrav *MonitorScheduledQueryRulesAlertV2) State() (*monitorScheduledQueryRulesAlertV2State, bool) {
	return msqrav.state, msqrav.state != nil
}

// StateMust returns the state for [MonitorScheduledQueryRulesAlertV2]. Panics if the state is nil.
func (msqrav *MonitorScheduledQueryRulesAlertV2) StateMust() *monitorScheduledQueryRulesAlertV2State {
	if msqrav.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", msqrav.Type(), msqrav.LocalName()))
	}
	return msqrav.state
}

// MonitorScheduledQueryRulesAlertV2Args contains the configurations for azurerm_monitor_scheduled_query_rules_alert_v2.
type MonitorScheduledQueryRulesAlertV2Args struct {
	// AutoMitigationEnabled: bool, optional
	AutoMitigationEnabled terra.BoolValue `hcl:"auto_mitigation_enabled,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// EvaluationFrequency: string, optional
	EvaluationFrequency terra.StringValue `hcl:"evaluation_frequency,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// MuteActionsAfterAlertDuration: string, optional
	MuteActionsAfterAlertDuration terra.StringValue `hcl:"mute_actions_after_alert_duration,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// QueryTimeRangeOverride: string, optional
	QueryTimeRangeOverride terra.StringValue `hcl:"query_time_range_override,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Scopes: list of string, required
	Scopes terra.ListValue[terra.StringValue] `hcl:"scopes,attr" validate:"required"`
	// Severity: number, required
	Severity terra.NumberValue `hcl:"severity,attr" validate:"required"`
	// SkipQueryValidation: bool, optional
	SkipQueryValidation terra.BoolValue `hcl:"skip_query_validation,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TargetResourceTypes: list of string, optional
	TargetResourceTypes terra.ListValue[terra.StringValue] `hcl:"target_resource_types,attr"`
	// WindowDuration: string, required
	WindowDuration terra.StringValue `hcl:"window_duration,attr" validate:"required"`
	// WorkspaceAlertsStorageEnabled: bool, optional
	WorkspaceAlertsStorageEnabled terra.BoolValue `hcl:"workspace_alerts_storage_enabled,attr"`
	// Action: optional
	Action *monitorscheduledqueryrulesalertv2.Action `hcl:"action,block"`
	// Criteria: min=1
	Criteria []monitorscheduledqueryrulesalertv2.Criteria `hcl:"criteria,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *monitorscheduledqueryrulesalertv2.Timeouts `hcl:"timeouts,block"`
}
type monitorScheduledQueryRulesAlertV2Attributes struct {
	ref terra.Reference
}

// AutoMitigationEnabled returns a reference to field auto_mitigation_enabled of azurerm_monitor_scheduled_query_rules_alert_v2.
func (msqrav monitorScheduledQueryRulesAlertV2Attributes) AutoMitigationEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(msqrav.ref.Append("auto_mitigation_enabled"))
}

// CreatedWithApiVersion returns a reference to field created_with_api_version of azurerm_monitor_scheduled_query_rules_alert_v2.
func (msqrav monitorScheduledQueryRulesAlertV2Attributes) CreatedWithApiVersion() terra.StringValue {
	return terra.ReferenceAsString(msqrav.ref.Append("created_with_api_version"))
}

// Description returns a reference to field description of azurerm_monitor_scheduled_query_rules_alert_v2.
func (msqrav monitorScheduledQueryRulesAlertV2Attributes) Description() terra.StringValue {
	return terra.ReferenceAsString(msqrav.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_monitor_scheduled_query_rules_alert_v2.
func (msqrav monitorScheduledQueryRulesAlertV2Attributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(msqrav.ref.Append("display_name"))
}

// Enabled returns a reference to field enabled of azurerm_monitor_scheduled_query_rules_alert_v2.
func (msqrav monitorScheduledQueryRulesAlertV2Attributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(msqrav.ref.Append("enabled"))
}

// EvaluationFrequency returns a reference to field evaluation_frequency of azurerm_monitor_scheduled_query_rules_alert_v2.
func (msqrav monitorScheduledQueryRulesAlertV2Attributes) EvaluationFrequency() terra.StringValue {
	return terra.ReferenceAsString(msqrav.ref.Append("evaluation_frequency"))
}

// Id returns a reference to field id of azurerm_monitor_scheduled_query_rules_alert_v2.
func (msqrav monitorScheduledQueryRulesAlertV2Attributes) Id() terra.StringValue {
	return terra.ReferenceAsString(msqrav.ref.Append("id"))
}

// IsALegacyLogAnalyticsRule returns a reference to field is_a_legacy_log_analytics_rule of azurerm_monitor_scheduled_query_rules_alert_v2.
func (msqrav monitorScheduledQueryRulesAlertV2Attributes) IsALegacyLogAnalyticsRule() terra.BoolValue {
	return terra.ReferenceAsBool(msqrav.ref.Append("is_a_legacy_log_analytics_rule"))
}

// IsWorkspaceAlertsStorageConfigured returns a reference to field is_workspace_alerts_storage_configured of azurerm_monitor_scheduled_query_rules_alert_v2.
func (msqrav monitorScheduledQueryRulesAlertV2Attributes) IsWorkspaceAlertsStorageConfigured() terra.BoolValue {
	return terra.ReferenceAsBool(msqrav.ref.Append("is_workspace_alerts_storage_configured"))
}

// Location returns a reference to field location of azurerm_monitor_scheduled_query_rules_alert_v2.
func (msqrav monitorScheduledQueryRulesAlertV2Attributes) Location() terra.StringValue {
	return terra.ReferenceAsString(msqrav.ref.Append("location"))
}

// MuteActionsAfterAlertDuration returns a reference to field mute_actions_after_alert_duration of azurerm_monitor_scheduled_query_rules_alert_v2.
func (msqrav monitorScheduledQueryRulesAlertV2Attributes) MuteActionsAfterAlertDuration() terra.StringValue {
	return terra.ReferenceAsString(msqrav.ref.Append("mute_actions_after_alert_duration"))
}

// Name returns a reference to field name of azurerm_monitor_scheduled_query_rules_alert_v2.
func (msqrav monitorScheduledQueryRulesAlertV2Attributes) Name() terra.StringValue {
	return terra.ReferenceAsString(msqrav.ref.Append("name"))
}

// QueryTimeRangeOverride returns a reference to field query_time_range_override of azurerm_monitor_scheduled_query_rules_alert_v2.
func (msqrav monitorScheduledQueryRulesAlertV2Attributes) QueryTimeRangeOverride() terra.StringValue {
	return terra.ReferenceAsString(msqrav.ref.Append("query_time_range_override"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_monitor_scheduled_query_rules_alert_v2.
func (msqrav monitorScheduledQueryRulesAlertV2Attributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(msqrav.ref.Append("resource_group_name"))
}

// Scopes returns a reference to field scopes of azurerm_monitor_scheduled_query_rules_alert_v2.
func (msqrav monitorScheduledQueryRulesAlertV2Attributes) Scopes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](msqrav.ref.Append("scopes"))
}

// Severity returns a reference to field severity of azurerm_monitor_scheduled_query_rules_alert_v2.
func (msqrav monitorScheduledQueryRulesAlertV2Attributes) Severity() terra.NumberValue {
	return terra.ReferenceAsNumber(msqrav.ref.Append("severity"))
}

// SkipQueryValidation returns a reference to field skip_query_validation of azurerm_monitor_scheduled_query_rules_alert_v2.
func (msqrav monitorScheduledQueryRulesAlertV2Attributes) SkipQueryValidation() terra.BoolValue {
	return terra.ReferenceAsBool(msqrav.ref.Append("skip_query_validation"))
}

// Tags returns a reference to field tags of azurerm_monitor_scheduled_query_rules_alert_v2.
func (msqrav monitorScheduledQueryRulesAlertV2Attributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](msqrav.ref.Append("tags"))
}

// TargetResourceTypes returns a reference to field target_resource_types of azurerm_monitor_scheduled_query_rules_alert_v2.
func (msqrav monitorScheduledQueryRulesAlertV2Attributes) TargetResourceTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](msqrav.ref.Append("target_resource_types"))
}

// WindowDuration returns a reference to field window_duration of azurerm_monitor_scheduled_query_rules_alert_v2.
func (msqrav monitorScheduledQueryRulesAlertV2Attributes) WindowDuration() terra.StringValue {
	return terra.ReferenceAsString(msqrav.ref.Append("window_duration"))
}

// WorkspaceAlertsStorageEnabled returns a reference to field workspace_alerts_storage_enabled of azurerm_monitor_scheduled_query_rules_alert_v2.
func (msqrav monitorScheduledQueryRulesAlertV2Attributes) WorkspaceAlertsStorageEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(msqrav.ref.Append("workspace_alerts_storage_enabled"))
}

func (msqrav monitorScheduledQueryRulesAlertV2Attributes) Action() terra.ListValue[monitorscheduledqueryrulesalertv2.ActionAttributes] {
	return terra.ReferenceAsList[monitorscheduledqueryrulesalertv2.ActionAttributes](msqrav.ref.Append("action"))
}

func (msqrav monitorScheduledQueryRulesAlertV2Attributes) Criteria() terra.ListValue[monitorscheduledqueryrulesalertv2.CriteriaAttributes] {
	return terra.ReferenceAsList[monitorscheduledqueryrulesalertv2.CriteriaAttributes](msqrav.ref.Append("criteria"))
}

func (msqrav monitorScheduledQueryRulesAlertV2Attributes) Timeouts() monitorscheduledqueryrulesalertv2.TimeoutsAttributes {
	return terra.ReferenceAsSingle[monitorscheduledqueryrulesalertv2.TimeoutsAttributes](msqrav.ref.Append("timeouts"))
}

type monitorScheduledQueryRulesAlertV2State struct {
	AutoMitigationEnabled              bool                                              `json:"auto_mitigation_enabled"`
	CreatedWithApiVersion              string                                            `json:"created_with_api_version"`
	Description                        string                                            `json:"description"`
	DisplayName                        string                                            `json:"display_name"`
	Enabled                            bool                                              `json:"enabled"`
	EvaluationFrequency                string                                            `json:"evaluation_frequency"`
	Id                                 string                                            `json:"id"`
	IsALegacyLogAnalyticsRule          bool                                              `json:"is_a_legacy_log_analytics_rule"`
	IsWorkspaceAlertsStorageConfigured bool                                              `json:"is_workspace_alerts_storage_configured"`
	Location                           string                                            `json:"location"`
	MuteActionsAfterAlertDuration      string                                            `json:"mute_actions_after_alert_duration"`
	Name                               string                                            `json:"name"`
	QueryTimeRangeOverride             string                                            `json:"query_time_range_override"`
	ResourceGroupName                  string                                            `json:"resource_group_name"`
	Scopes                             []string                                          `json:"scopes"`
	Severity                           float64                                           `json:"severity"`
	SkipQueryValidation                bool                                              `json:"skip_query_validation"`
	Tags                               map[string]string                                 `json:"tags"`
	TargetResourceTypes                []string                                          `json:"target_resource_types"`
	WindowDuration                     string                                            `json:"window_duration"`
	WorkspaceAlertsStorageEnabled      bool                                              `json:"workspace_alerts_storage_enabled"`
	Action                             []monitorscheduledqueryrulesalertv2.ActionState   `json:"action"`
	Criteria                           []monitorscheduledqueryrulesalertv2.CriteriaState `json:"criteria"`
	Timeouts                           *monitorscheduledqueryrulesalertv2.TimeoutsState  `json:"timeouts"`
}
