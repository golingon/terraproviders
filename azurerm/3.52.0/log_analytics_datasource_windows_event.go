// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	loganalyticsdatasourcewindowsevent "github.com/golingon/terraproviders/azurerm/3.52.0/loganalyticsdatasourcewindowsevent"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLogAnalyticsDatasourceWindowsEvent creates a new instance of [LogAnalyticsDatasourceWindowsEvent].
func NewLogAnalyticsDatasourceWindowsEvent(name string, args LogAnalyticsDatasourceWindowsEventArgs) *LogAnalyticsDatasourceWindowsEvent {
	return &LogAnalyticsDatasourceWindowsEvent{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LogAnalyticsDatasourceWindowsEvent)(nil)

// LogAnalyticsDatasourceWindowsEvent represents the Terraform resource azurerm_log_analytics_datasource_windows_event.
type LogAnalyticsDatasourceWindowsEvent struct {
	Name      string
	Args      LogAnalyticsDatasourceWindowsEventArgs
	state     *logAnalyticsDatasourceWindowsEventState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LogAnalyticsDatasourceWindowsEvent].
func (ladwe *LogAnalyticsDatasourceWindowsEvent) Type() string {
	return "azurerm_log_analytics_datasource_windows_event"
}

// LocalName returns the local name for [LogAnalyticsDatasourceWindowsEvent].
func (ladwe *LogAnalyticsDatasourceWindowsEvent) LocalName() string {
	return ladwe.Name
}

// Configuration returns the configuration (args) for [LogAnalyticsDatasourceWindowsEvent].
func (ladwe *LogAnalyticsDatasourceWindowsEvent) Configuration() interface{} {
	return ladwe.Args
}

// DependOn is used for other resources to depend on [LogAnalyticsDatasourceWindowsEvent].
func (ladwe *LogAnalyticsDatasourceWindowsEvent) DependOn() terra.Reference {
	return terra.ReferenceResource(ladwe)
}

// Dependencies returns the list of resources [LogAnalyticsDatasourceWindowsEvent] depends_on.
func (ladwe *LogAnalyticsDatasourceWindowsEvent) Dependencies() terra.Dependencies {
	return ladwe.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LogAnalyticsDatasourceWindowsEvent].
func (ladwe *LogAnalyticsDatasourceWindowsEvent) LifecycleManagement() *terra.Lifecycle {
	return ladwe.Lifecycle
}

// Attributes returns the attributes for [LogAnalyticsDatasourceWindowsEvent].
func (ladwe *LogAnalyticsDatasourceWindowsEvent) Attributes() logAnalyticsDatasourceWindowsEventAttributes {
	return logAnalyticsDatasourceWindowsEventAttributes{ref: terra.ReferenceResource(ladwe)}
}

// ImportState imports the given attribute values into [LogAnalyticsDatasourceWindowsEvent]'s state.
func (ladwe *LogAnalyticsDatasourceWindowsEvent) ImportState(av io.Reader) error {
	ladwe.state = &logAnalyticsDatasourceWindowsEventState{}
	if err := json.NewDecoder(av).Decode(ladwe.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ladwe.Type(), ladwe.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LogAnalyticsDatasourceWindowsEvent] has state.
func (ladwe *LogAnalyticsDatasourceWindowsEvent) State() (*logAnalyticsDatasourceWindowsEventState, bool) {
	return ladwe.state, ladwe.state != nil
}

// StateMust returns the state for [LogAnalyticsDatasourceWindowsEvent]. Panics if the state is nil.
func (ladwe *LogAnalyticsDatasourceWindowsEvent) StateMust() *logAnalyticsDatasourceWindowsEventState {
	if ladwe.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ladwe.Type(), ladwe.LocalName()))
	}
	return ladwe.state
}

// LogAnalyticsDatasourceWindowsEventArgs contains the configurations for azurerm_log_analytics_datasource_windows_event.
type LogAnalyticsDatasourceWindowsEventArgs struct {
	// EventLogName: string, required
	EventLogName terra.StringValue `hcl:"event_log_name,attr" validate:"required"`
	// EventTypes: set of string, required
	EventTypes terra.SetValue[terra.StringValue] `hcl:"event_types,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// WorkspaceName: string, required
	WorkspaceName terra.StringValue `hcl:"workspace_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *loganalyticsdatasourcewindowsevent.Timeouts `hcl:"timeouts,block"`
}
type logAnalyticsDatasourceWindowsEventAttributes struct {
	ref terra.Reference
}

// EventLogName returns a reference to field event_log_name of azurerm_log_analytics_datasource_windows_event.
func (ladwe logAnalyticsDatasourceWindowsEventAttributes) EventLogName() terra.StringValue {
	return terra.ReferenceAsString(ladwe.ref.Append("event_log_name"))
}

// EventTypes returns a reference to field event_types of azurerm_log_analytics_datasource_windows_event.
func (ladwe logAnalyticsDatasourceWindowsEventAttributes) EventTypes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ladwe.ref.Append("event_types"))
}

// Id returns a reference to field id of azurerm_log_analytics_datasource_windows_event.
func (ladwe logAnalyticsDatasourceWindowsEventAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ladwe.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_log_analytics_datasource_windows_event.
func (ladwe logAnalyticsDatasourceWindowsEventAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ladwe.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_log_analytics_datasource_windows_event.
func (ladwe logAnalyticsDatasourceWindowsEventAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ladwe.ref.Append("resource_group_name"))
}

// WorkspaceName returns a reference to field workspace_name of azurerm_log_analytics_datasource_windows_event.
func (ladwe logAnalyticsDatasourceWindowsEventAttributes) WorkspaceName() terra.StringValue {
	return terra.ReferenceAsString(ladwe.ref.Append("workspace_name"))
}

func (ladwe logAnalyticsDatasourceWindowsEventAttributes) Timeouts() loganalyticsdatasourcewindowsevent.TimeoutsAttributes {
	return terra.ReferenceAsSingle[loganalyticsdatasourcewindowsevent.TimeoutsAttributes](ladwe.ref.Append("timeouts"))
}

type logAnalyticsDatasourceWindowsEventState struct {
	EventLogName      string                                            `json:"event_log_name"`
	EventTypes        []string                                          `json:"event_types"`
	Id                string                                            `json:"id"`
	Name              string                                            `json:"name"`
	ResourceGroupName string                                            `json:"resource_group_name"`
	WorkspaceName     string                                            `json:"workspace_name"`
	Timeouts          *loganalyticsdatasourcewindowsevent.TimeoutsState `json:"timeouts"`
}
