// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	loganalyticsdataexportrule "github.com/golingon/terraproviders/azurerm/3.64.0/loganalyticsdataexportrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLogAnalyticsDataExportRule creates a new instance of [LogAnalyticsDataExportRule].
func NewLogAnalyticsDataExportRule(name string, args LogAnalyticsDataExportRuleArgs) *LogAnalyticsDataExportRule {
	return &LogAnalyticsDataExportRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LogAnalyticsDataExportRule)(nil)

// LogAnalyticsDataExportRule represents the Terraform resource azurerm_log_analytics_data_export_rule.
type LogAnalyticsDataExportRule struct {
	Name      string
	Args      LogAnalyticsDataExportRuleArgs
	state     *logAnalyticsDataExportRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LogAnalyticsDataExportRule].
func (lader *LogAnalyticsDataExportRule) Type() string {
	return "azurerm_log_analytics_data_export_rule"
}

// LocalName returns the local name for [LogAnalyticsDataExportRule].
func (lader *LogAnalyticsDataExportRule) LocalName() string {
	return lader.Name
}

// Configuration returns the configuration (args) for [LogAnalyticsDataExportRule].
func (lader *LogAnalyticsDataExportRule) Configuration() interface{} {
	return lader.Args
}

// DependOn is used for other resources to depend on [LogAnalyticsDataExportRule].
func (lader *LogAnalyticsDataExportRule) DependOn() terra.Reference {
	return terra.ReferenceResource(lader)
}

// Dependencies returns the list of resources [LogAnalyticsDataExportRule] depends_on.
func (lader *LogAnalyticsDataExportRule) Dependencies() terra.Dependencies {
	return lader.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LogAnalyticsDataExportRule].
func (lader *LogAnalyticsDataExportRule) LifecycleManagement() *terra.Lifecycle {
	return lader.Lifecycle
}

// Attributes returns the attributes for [LogAnalyticsDataExportRule].
func (lader *LogAnalyticsDataExportRule) Attributes() logAnalyticsDataExportRuleAttributes {
	return logAnalyticsDataExportRuleAttributes{ref: terra.ReferenceResource(lader)}
}

// ImportState imports the given attribute values into [LogAnalyticsDataExportRule]'s state.
func (lader *LogAnalyticsDataExportRule) ImportState(av io.Reader) error {
	lader.state = &logAnalyticsDataExportRuleState{}
	if err := json.NewDecoder(av).Decode(lader.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lader.Type(), lader.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LogAnalyticsDataExportRule] has state.
func (lader *LogAnalyticsDataExportRule) State() (*logAnalyticsDataExportRuleState, bool) {
	return lader.state, lader.state != nil
}

// StateMust returns the state for [LogAnalyticsDataExportRule]. Panics if the state is nil.
func (lader *LogAnalyticsDataExportRule) StateMust() *logAnalyticsDataExportRuleState {
	if lader.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lader.Type(), lader.LocalName()))
	}
	return lader.state
}

// LogAnalyticsDataExportRuleArgs contains the configurations for azurerm_log_analytics_data_export_rule.
type LogAnalyticsDataExportRuleArgs struct {
	// DestinationResourceId: string, required
	DestinationResourceId terra.StringValue `hcl:"destination_resource_id,attr" validate:"required"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// TableNames: set of string, required
	TableNames terra.SetValue[terra.StringValue] `hcl:"table_names,attr" validate:"required"`
	// WorkspaceResourceId: string, required
	WorkspaceResourceId terra.StringValue `hcl:"workspace_resource_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *loganalyticsdataexportrule.Timeouts `hcl:"timeouts,block"`
}
type logAnalyticsDataExportRuleAttributes struct {
	ref terra.Reference
}

// DestinationResourceId returns a reference to field destination_resource_id of azurerm_log_analytics_data_export_rule.
func (lader logAnalyticsDataExportRuleAttributes) DestinationResourceId() terra.StringValue {
	return terra.ReferenceAsString(lader.ref.Append("destination_resource_id"))
}

// Enabled returns a reference to field enabled of azurerm_log_analytics_data_export_rule.
func (lader logAnalyticsDataExportRuleAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(lader.ref.Append("enabled"))
}

// ExportRuleId returns a reference to field export_rule_id of azurerm_log_analytics_data_export_rule.
func (lader logAnalyticsDataExportRuleAttributes) ExportRuleId() terra.StringValue {
	return terra.ReferenceAsString(lader.ref.Append("export_rule_id"))
}

// Id returns a reference to field id of azurerm_log_analytics_data_export_rule.
func (lader logAnalyticsDataExportRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lader.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_log_analytics_data_export_rule.
func (lader logAnalyticsDataExportRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lader.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_log_analytics_data_export_rule.
func (lader logAnalyticsDataExportRuleAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(lader.ref.Append("resource_group_name"))
}

// TableNames returns a reference to field table_names of azurerm_log_analytics_data_export_rule.
func (lader logAnalyticsDataExportRuleAttributes) TableNames() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](lader.ref.Append("table_names"))
}

// WorkspaceResourceId returns a reference to field workspace_resource_id of azurerm_log_analytics_data_export_rule.
func (lader logAnalyticsDataExportRuleAttributes) WorkspaceResourceId() terra.StringValue {
	return terra.ReferenceAsString(lader.ref.Append("workspace_resource_id"))
}

func (lader logAnalyticsDataExportRuleAttributes) Timeouts() loganalyticsdataexportrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[loganalyticsdataexportrule.TimeoutsAttributes](lader.ref.Append("timeouts"))
}

type logAnalyticsDataExportRuleState struct {
	DestinationResourceId string                                    `json:"destination_resource_id"`
	Enabled               bool                                      `json:"enabled"`
	ExportRuleId          string                                    `json:"export_rule_id"`
	Id                    string                                    `json:"id"`
	Name                  string                                    `json:"name"`
	ResourceGroupName     string                                    `json:"resource_group_name"`
	TableNames            []string                                  `json:"table_names"`
	WorkspaceResourceId   string                                    `json:"workspace_resource_id"`
	Timeouts              *loganalyticsdataexportrule.TimeoutsState `json:"timeouts"`
}
