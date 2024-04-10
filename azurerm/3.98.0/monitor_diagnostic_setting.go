// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	monitordiagnosticsetting "github.com/golingon/terraproviders/azurerm/3.98.0/monitordiagnosticsetting"
	"io"
)

// NewMonitorDiagnosticSetting creates a new instance of [MonitorDiagnosticSetting].
func NewMonitorDiagnosticSetting(name string, args MonitorDiagnosticSettingArgs) *MonitorDiagnosticSetting {
	return &MonitorDiagnosticSetting{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MonitorDiagnosticSetting)(nil)

// MonitorDiagnosticSetting represents the Terraform resource azurerm_monitor_diagnostic_setting.
type MonitorDiagnosticSetting struct {
	Name      string
	Args      MonitorDiagnosticSettingArgs
	state     *monitorDiagnosticSettingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MonitorDiagnosticSetting].
func (mds *MonitorDiagnosticSetting) Type() string {
	return "azurerm_monitor_diagnostic_setting"
}

// LocalName returns the local name for [MonitorDiagnosticSetting].
func (mds *MonitorDiagnosticSetting) LocalName() string {
	return mds.Name
}

// Configuration returns the configuration (args) for [MonitorDiagnosticSetting].
func (mds *MonitorDiagnosticSetting) Configuration() interface{} {
	return mds.Args
}

// DependOn is used for other resources to depend on [MonitorDiagnosticSetting].
func (mds *MonitorDiagnosticSetting) DependOn() terra.Reference {
	return terra.ReferenceResource(mds)
}

// Dependencies returns the list of resources [MonitorDiagnosticSetting] depends_on.
func (mds *MonitorDiagnosticSetting) Dependencies() terra.Dependencies {
	return mds.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MonitorDiagnosticSetting].
func (mds *MonitorDiagnosticSetting) LifecycleManagement() *terra.Lifecycle {
	return mds.Lifecycle
}

// Attributes returns the attributes for [MonitorDiagnosticSetting].
func (mds *MonitorDiagnosticSetting) Attributes() monitorDiagnosticSettingAttributes {
	return monitorDiagnosticSettingAttributes{ref: terra.ReferenceResource(mds)}
}

// ImportState imports the given attribute values into [MonitorDiagnosticSetting]'s state.
func (mds *MonitorDiagnosticSetting) ImportState(av io.Reader) error {
	mds.state = &monitorDiagnosticSettingState{}
	if err := json.NewDecoder(av).Decode(mds.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mds.Type(), mds.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MonitorDiagnosticSetting] has state.
func (mds *MonitorDiagnosticSetting) State() (*monitorDiagnosticSettingState, bool) {
	return mds.state, mds.state != nil
}

// StateMust returns the state for [MonitorDiagnosticSetting]. Panics if the state is nil.
func (mds *MonitorDiagnosticSetting) StateMust() *monitorDiagnosticSettingState {
	if mds.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mds.Type(), mds.LocalName()))
	}
	return mds.state
}

// MonitorDiagnosticSettingArgs contains the configurations for azurerm_monitor_diagnostic_setting.
type MonitorDiagnosticSettingArgs struct {
	// EventhubAuthorizationRuleId: string, optional
	EventhubAuthorizationRuleId terra.StringValue `hcl:"eventhub_authorization_rule_id,attr"`
	// EventhubName: string, optional
	EventhubName terra.StringValue `hcl:"eventhub_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogAnalyticsDestinationType: string, optional
	LogAnalyticsDestinationType terra.StringValue `hcl:"log_analytics_destination_type,attr"`
	// LogAnalyticsWorkspaceId: string, optional
	LogAnalyticsWorkspaceId terra.StringValue `hcl:"log_analytics_workspace_id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PartnerSolutionId: string, optional
	PartnerSolutionId terra.StringValue `hcl:"partner_solution_id,attr"`
	// StorageAccountId: string, optional
	StorageAccountId terra.StringValue `hcl:"storage_account_id,attr"`
	// TargetResourceId: string, required
	TargetResourceId terra.StringValue `hcl:"target_resource_id,attr" validate:"required"`
	// EnabledLog: min=0
	EnabledLog []monitordiagnosticsetting.EnabledLog `hcl:"enabled_log,block" validate:"min=0"`
	// Log: min=0
	Log []monitordiagnosticsetting.Log `hcl:"log,block" validate:"min=0"`
	// Metric: min=0
	Metric []monitordiagnosticsetting.Metric `hcl:"metric,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *monitordiagnosticsetting.Timeouts `hcl:"timeouts,block"`
}
type monitorDiagnosticSettingAttributes struct {
	ref terra.Reference
}

// EventhubAuthorizationRuleId returns a reference to field eventhub_authorization_rule_id of azurerm_monitor_diagnostic_setting.
func (mds monitorDiagnosticSettingAttributes) EventhubAuthorizationRuleId() terra.StringValue {
	return terra.ReferenceAsString(mds.ref.Append("eventhub_authorization_rule_id"))
}

// EventhubName returns a reference to field eventhub_name of azurerm_monitor_diagnostic_setting.
func (mds monitorDiagnosticSettingAttributes) EventhubName() terra.StringValue {
	return terra.ReferenceAsString(mds.ref.Append("eventhub_name"))
}

// Id returns a reference to field id of azurerm_monitor_diagnostic_setting.
func (mds monitorDiagnosticSettingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mds.ref.Append("id"))
}

// LogAnalyticsDestinationType returns a reference to field log_analytics_destination_type of azurerm_monitor_diagnostic_setting.
func (mds monitorDiagnosticSettingAttributes) LogAnalyticsDestinationType() terra.StringValue {
	return terra.ReferenceAsString(mds.ref.Append("log_analytics_destination_type"))
}

// LogAnalyticsWorkspaceId returns a reference to field log_analytics_workspace_id of azurerm_monitor_diagnostic_setting.
func (mds monitorDiagnosticSettingAttributes) LogAnalyticsWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(mds.ref.Append("log_analytics_workspace_id"))
}

// Name returns a reference to field name of azurerm_monitor_diagnostic_setting.
func (mds monitorDiagnosticSettingAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mds.ref.Append("name"))
}

// PartnerSolutionId returns a reference to field partner_solution_id of azurerm_monitor_diagnostic_setting.
func (mds monitorDiagnosticSettingAttributes) PartnerSolutionId() terra.StringValue {
	return terra.ReferenceAsString(mds.ref.Append("partner_solution_id"))
}

// StorageAccountId returns a reference to field storage_account_id of azurerm_monitor_diagnostic_setting.
func (mds monitorDiagnosticSettingAttributes) StorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(mds.ref.Append("storage_account_id"))
}

// TargetResourceId returns a reference to field target_resource_id of azurerm_monitor_diagnostic_setting.
func (mds monitorDiagnosticSettingAttributes) TargetResourceId() terra.StringValue {
	return terra.ReferenceAsString(mds.ref.Append("target_resource_id"))
}

func (mds monitorDiagnosticSettingAttributes) EnabledLog() terra.SetValue[monitordiagnosticsetting.EnabledLogAttributes] {
	return terra.ReferenceAsSet[monitordiagnosticsetting.EnabledLogAttributes](mds.ref.Append("enabled_log"))
}

func (mds monitorDiagnosticSettingAttributes) Log() terra.SetValue[monitordiagnosticsetting.LogAttributes] {
	return terra.ReferenceAsSet[monitordiagnosticsetting.LogAttributes](mds.ref.Append("log"))
}

func (mds monitorDiagnosticSettingAttributes) Metric() terra.SetValue[monitordiagnosticsetting.MetricAttributes] {
	return terra.ReferenceAsSet[monitordiagnosticsetting.MetricAttributes](mds.ref.Append("metric"))
}

func (mds monitorDiagnosticSettingAttributes) Timeouts() monitordiagnosticsetting.TimeoutsAttributes {
	return terra.ReferenceAsSingle[monitordiagnosticsetting.TimeoutsAttributes](mds.ref.Append("timeouts"))
}

type monitorDiagnosticSettingState struct {
	EventhubAuthorizationRuleId string                                     `json:"eventhub_authorization_rule_id"`
	EventhubName                string                                     `json:"eventhub_name"`
	Id                          string                                     `json:"id"`
	LogAnalyticsDestinationType string                                     `json:"log_analytics_destination_type"`
	LogAnalyticsWorkspaceId     string                                     `json:"log_analytics_workspace_id"`
	Name                        string                                     `json:"name"`
	PartnerSolutionId           string                                     `json:"partner_solution_id"`
	StorageAccountId            string                                     `json:"storage_account_id"`
	TargetResourceId            string                                     `json:"target_resource_id"`
	EnabledLog                  []monitordiagnosticsetting.EnabledLogState `json:"enabled_log"`
	Log                         []monitordiagnosticsetting.LogState        `json:"log"`
	Metric                      []monitordiagnosticsetting.MetricState     `json:"metric"`
	Timeouts                    *monitordiagnosticsetting.TimeoutsState    `json:"timeouts"`
}
