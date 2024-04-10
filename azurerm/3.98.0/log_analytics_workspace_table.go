// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	loganalyticsworkspacetable "github.com/golingon/terraproviders/azurerm/3.98.0/loganalyticsworkspacetable"
	"io"
)

// NewLogAnalyticsWorkspaceTable creates a new instance of [LogAnalyticsWorkspaceTable].
func NewLogAnalyticsWorkspaceTable(name string, args LogAnalyticsWorkspaceTableArgs) *LogAnalyticsWorkspaceTable {
	return &LogAnalyticsWorkspaceTable{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LogAnalyticsWorkspaceTable)(nil)

// LogAnalyticsWorkspaceTable represents the Terraform resource azurerm_log_analytics_workspace_table.
type LogAnalyticsWorkspaceTable struct {
	Name      string
	Args      LogAnalyticsWorkspaceTableArgs
	state     *logAnalyticsWorkspaceTableState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LogAnalyticsWorkspaceTable].
func (lawt *LogAnalyticsWorkspaceTable) Type() string {
	return "azurerm_log_analytics_workspace_table"
}

// LocalName returns the local name for [LogAnalyticsWorkspaceTable].
func (lawt *LogAnalyticsWorkspaceTable) LocalName() string {
	return lawt.Name
}

// Configuration returns the configuration (args) for [LogAnalyticsWorkspaceTable].
func (lawt *LogAnalyticsWorkspaceTable) Configuration() interface{} {
	return lawt.Args
}

// DependOn is used for other resources to depend on [LogAnalyticsWorkspaceTable].
func (lawt *LogAnalyticsWorkspaceTable) DependOn() terra.Reference {
	return terra.ReferenceResource(lawt)
}

// Dependencies returns the list of resources [LogAnalyticsWorkspaceTable] depends_on.
func (lawt *LogAnalyticsWorkspaceTable) Dependencies() terra.Dependencies {
	return lawt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LogAnalyticsWorkspaceTable].
func (lawt *LogAnalyticsWorkspaceTable) LifecycleManagement() *terra.Lifecycle {
	return lawt.Lifecycle
}

// Attributes returns the attributes for [LogAnalyticsWorkspaceTable].
func (lawt *LogAnalyticsWorkspaceTable) Attributes() logAnalyticsWorkspaceTableAttributes {
	return logAnalyticsWorkspaceTableAttributes{ref: terra.ReferenceResource(lawt)}
}

// ImportState imports the given attribute values into [LogAnalyticsWorkspaceTable]'s state.
func (lawt *LogAnalyticsWorkspaceTable) ImportState(av io.Reader) error {
	lawt.state = &logAnalyticsWorkspaceTableState{}
	if err := json.NewDecoder(av).Decode(lawt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lawt.Type(), lawt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LogAnalyticsWorkspaceTable] has state.
func (lawt *LogAnalyticsWorkspaceTable) State() (*logAnalyticsWorkspaceTableState, bool) {
	return lawt.state, lawt.state != nil
}

// StateMust returns the state for [LogAnalyticsWorkspaceTable]. Panics if the state is nil.
func (lawt *LogAnalyticsWorkspaceTable) StateMust() *logAnalyticsWorkspaceTableState {
	if lawt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lawt.Type(), lawt.LocalName()))
	}
	return lawt.state
}

// LogAnalyticsWorkspaceTableArgs contains the configurations for azurerm_log_analytics_workspace_table.
type LogAnalyticsWorkspaceTableArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Plan: string, optional
	Plan terra.StringValue `hcl:"plan,attr"`
	// RetentionInDays: number, optional
	RetentionInDays terra.NumberValue `hcl:"retention_in_days,attr"`
	// TotalRetentionInDays: number, optional
	TotalRetentionInDays terra.NumberValue `hcl:"total_retention_in_days,attr"`
	// WorkspaceId: string, required
	WorkspaceId terra.StringValue `hcl:"workspace_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *loganalyticsworkspacetable.Timeouts `hcl:"timeouts,block"`
}
type logAnalyticsWorkspaceTableAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_log_analytics_workspace_table.
func (lawt logAnalyticsWorkspaceTableAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lawt.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_log_analytics_workspace_table.
func (lawt logAnalyticsWorkspaceTableAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lawt.ref.Append("name"))
}

// Plan returns a reference to field plan of azurerm_log_analytics_workspace_table.
func (lawt logAnalyticsWorkspaceTableAttributes) Plan() terra.StringValue {
	return terra.ReferenceAsString(lawt.ref.Append("plan"))
}

// RetentionInDays returns a reference to field retention_in_days of azurerm_log_analytics_workspace_table.
func (lawt logAnalyticsWorkspaceTableAttributes) RetentionInDays() terra.NumberValue {
	return terra.ReferenceAsNumber(lawt.ref.Append("retention_in_days"))
}

// TotalRetentionInDays returns a reference to field total_retention_in_days of azurerm_log_analytics_workspace_table.
func (lawt logAnalyticsWorkspaceTableAttributes) TotalRetentionInDays() terra.NumberValue {
	return terra.ReferenceAsNumber(lawt.ref.Append("total_retention_in_days"))
}

// WorkspaceId returns a reference to field workspace_id of azurerm_log_analytics_workspace_table.
func (lawt logAnalyticsWorkspaceTableAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(lawt.ref.Append("workspace_id"))
}

func (lawt logAnalyticsWorkspaceTableAttributes) Timeouts() loganalyticsworkspacetable.TimeoutsAttributes {
	return terra.ReferenceAsSingle[loganalyticsworkspacetable.TimeoutsAttributes](lawt.ref.Append("timeouts"))
}

type logAnalyticsWorkspaceTableState struct {
	Id                   string                                    `json:"id"`
	Name                 string                                    `json:"name"`
	Plan                 string                                    `json:"plan"`
	RetentionInDays      float64                                   `json:"retention_in_days"`
	TotalRetentionInDays float64                                   `json:"total_retention_in_days"`
	WorkspaceId          string                                    `json:"workspace_id"`
	Timeouts             *loganalyticsworkspacetable.TimeoutsState `json:"timeouts"`
}
