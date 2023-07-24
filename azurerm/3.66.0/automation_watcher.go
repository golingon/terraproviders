// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	automationwatcher "github.com/golingon/terraproviders/azurerm/3.66.0/automationwatcher"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAutomationWatcher creates a new instance of [AutomationWatcher].
func NewAutomationWatcher(name string, args AutomationWatcherArgs) *AutomationWatcher {
	return &AutomationWatcher{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AutomationWatcher)(nil)

// AutomationWatcher represents the Terraform resource azurerm_automation_watcher.
type AutomationWatcher struct {
	Name      string
	Args      AutomationWatcherArgs
	state     *automationWatcherState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AutomationWatcher].
func (aw *AutomationWatcher) Type() string {
	return "azurerm_automation_watcher"
}

// LocalName returns the local name for [AutomationWatcher].
func (aw *AutomationWatcher) LocalName() string {
	return aw.Name
}

// Configuration returns the configuration (args) for [AutomationWatcher].
func (aw *AutomationWatcher) Configuration() interface{} {
	return aw.Args
}

// DependOn is used for other resources to depend on [AutomationWatcher].
func (aw *AutomationWatcher) DependOn() terra.Reference {
	return terra.ReferenceResource(aw)
}

// Dependencies returns the list of resources [AutomationWatcher] depends_on.
func (aw *AutomationWatcher) Dependencies() terra.Dependencies {
	return aw.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AutomationWatcher].
func (aw *AutomationWatcher) LifecycleManagement() *terra.Lifecycle {
	return aw.Lifecycle
}

// Attributes returns the attributes for [AutomationWatcher].
func (aw *AutomationWatcher) Attributes() automationWatcherAttributes {
	return automationWatcherAttributes{ref: terra.ReferenceResource(aw)}
}

// ImportState imports the given attribute values into [AutomationWatcher]'s state.
func (aw *AutomationWatcher) ImportState(av io.Reader) error {
	aw.state = &automationWatcherState{}
	if err := json.NewDecoder(av).Decode(aw.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aw.Type(), aw.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AutomationWatcher] has state.
func (aw *AutomationWatcher) State() (*automationWatcherState, bool) {
	return aw.state, aw.state != nil
}

// StateMust returns the state for [AutomationWatcher]. Panics if the state is nil.
func (aw *AutomationWatcher) StateMust() *automationWatcherState {
	if aw.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aw.Type(), aw.LocalName()))
	}
	return aw.state
}

// AutomationWatcherArgs contains the configurations for azurerm_automation_watcher.
type AutomationWatcherArgs struct {
	// AutomationAccountId: string, required
	AutomationAccountId terra.StringValue `hcl:"automation_account_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Etag: string, optional
	Etag terra.StringValue `hcl:"etag,attr"`
	// ExecutionFrequencyInSeconds: number, required
	ExecutionFrequencyInSeconds terra.NumberValue `hcl:"execution_frequency_in_seconds,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ScriptName: string, required
	ScriptName terra.StringValue `hcl:"script_name,attr" validate:"required"`
	// ScriptParameters: map of string, optional
	ScriptParameters terra.MapValue[terra.StringValue] `hcl:"script_parameters,attr"`
	// ScriptRunOn: string, required
	ScriptRunOn terra.StringValue `hcl:"script_run_on,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *automationwatcher.Timeouts `hcl:"timeouts,block"`
}
type automationWatcherAttributes struct {
	ref terra.Reference
}

// AutomationAccountId returns a reference to field automation_account_id of azurerm_automation_watcher.
func (aw automationWatcherAttributes) AutomationAccountId() terra.StringValue {
	return terra.ReferenceAsString(aw.ref.Append("automation_account_id"))
}

// Description returns a reference to field description of azurerm_automation_watcher.
func (aw automationWatcherAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(aw.ref.Append("description"))
}

// Etag returns a reference to field etag of azurerm_automation_watcher.
func (aw automationWatcherAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(aw.ref.Append("etag"))
}

// ExecutionFrequencyInSeconds returns a reference to field execution_frequency_in_seconds of azurerm_automation_watcher.
func (aw automationWatcherAttributes) ExecutionFrequencyInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(aw.ref.Append("execution_frequency_in_seconds"))
}

// Id returns a reference to field id of azurerm_automation_watcher.
func (aw automationWatcherAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aw.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_automation_watcher.
func (aw automationWatcherAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(aw.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_automation_watcher.
func (aw automationWatcherAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aw.ref.Append("name"))
}

// ScriptName returns a reference to field script_name of azurerm_automation_watcher.
func (aw automationWatcherAttributes) ScriptName() terra.StringValue {
	return terra.ReferenceAsString(aw.ref.Append("script_name"))
}

// ScriptParameters returns a reference to field script_parameters of azurerm_automation_watcher.
func (aw automationWatcherAttributes) ScriptParameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aw.ref.Append("script_parameters"))
}

// ScriptRunOn returns a reference to field script_run_on of azurerm_automation_watcher.
func (aw automationWatcherAttributes) ScriptRunOn() terra.StringValue {
	return terra.ReferenceAsString(aw.ref.Append("script_run_on"))
}

// Status returns a reference to field status of azurerm_automation_watcher.
func (aw automationWatcherAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(aw.ref.Append("status"))
}

// Tags returns a reference to field tags of azurerm_automation_watcher.
func (aw automationWatcherAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aw.ref.Append("tags"))
}

func (aw automationWatcherAttributes) Timeouts() automationwatcher.TimeoutsAttributes {
	return terra.ReferenceAsSingle[automationwatcher.TimeoutsAttributes](aw.ref.Append("timeouts"))
}

type automationWatcherState struct {
	AutomationAccountId         string                           `json:"automation_account_id"`
	Description                 string                           `json:"description"`
	Etag                        string                           `json:"etag"`
	ExecutionFrequencyInSeconds float64                          `json:"execution_frequency_in_seconds"`
	Id                          string                           `json:"id"`
	Location                    string                           `json:"location"`
	Name                        string                           `json:"name"`
	ScriptName                  string                           `json:"script_name"`
	ScriptParameters            map[string]string                `json:"script_parameters"`
	ScriptRunOn                 string                           `json:"script_run_on"`
	Status                      string                           `json:"status"`
	Tags                        map[string]string                `json:"tags"`
	Timeouts                    *automationwatcher.TimeoutsState `json:"timeouts"`
}
