// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	logicapptriggerrecurrence "github.com/golingon/terraproviders/azurerm/3.58.0/logicapptriggerrecurrence"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLogicAppTriggerRecurrence creates a new instance of [LogicAppTriggerRecurrence].
func NewLogicAppTriggerRecurrence(name string, args LogicAppTriggerRecurrenceArgs) *LogicAppTriggerRecurrence {
	return &LogicAppTriggerRecurrence{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LogicAppTriggerRecurrence)(nil)

// LogicAppTriggerRecurrence represents the Terraform resource azurerm_logic_app_trigger_recurrence.
type LogicAppTriggerRecurrence struct {
	Name      string
	Args      LogicAppTriggerRecurrenceArgs
	state     *logicAppTriggerRecurrenceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LogicAppTriggerRecurrence].
func (latr *LogicAppTriggerRecurrence) Type() string {
	return "azurerm_logic_app_trigger_recurrence"
}

// LocalName returns the local name for [LogicAppTriggerRecurrence].
func (latr *LogicAppTriggerRecurrence) LocalName() string {
	return latr.Name
}

// Configuration returns the configuration (args) for [LogicAppTriggerRecurrence].
func (latr *LogicAppTriggerRecurrence) Configuration() interface{} {
	return latr.Args
}

// DependOn is used for other resources to depend on [LogicAppTriggerRecurrence].
func (latr *LogicAppTriggerRecurrence) DependOn() terra.Reference {
	return terra.ReferenceResource(latr)
}

// Dependencies returns the list of resources [LogicAppTriggerRecurrence] depends_on.
func (latr *LogicAppTriggerRecurrence) Dependencies() terra.Dependencies {
	return latr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LogicAppTriggerRecurrence].
func (latr *LogicAppTriggerRecurrence) LifecycleManagement() *terra.Lifecycle {
	return latr.Lifecycle
}

// Attributes returns the attributes for [LogicAppTriggerRecurrence].
func (latr *LogicAppTriggerRecurrence) Attributes() logicAppTriggerRecurrenceAttributes {
	return logicAppTriggerRecurrenceAttributes{ref: terra.ReferenceResource(latr)}
}

// ImportState imports the given attribute values into [LogicAppTriggerRecurrence]'s state.
func (latr *LogicAppTriggerRecurrence) ImportState(av io.Reader) error {
	latr.state = &logicAppTriggerRecurrenceState{}
	if err := json.NewDecoder(av).Decode(latr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", latr.Type(), latr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LogicAppTriggerRecurrence] has state.
func (latr *LogicAppTriggerRecurrence) State() (*logicAppTriggerRecurrenceState, bool) {
	return latr.state, latr.state != nil
}

// StateMust returns the state for [LogicAppTriggerRecurrence]. Panics if the state is nil.
func (latr *LogicAppTriggerRecurrence) StateMust() *logicAppTriggerRecurrenceState {
	if latr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", latr.Type(), latr.LocalName()))
	}
	return latr.state
}

// LogicAppTriggerRecurrenceArgs contains the configurations for azurerm_logic_app_trigger_recurrence.
type LogicAppTriggerRecurrenceArgs struct {
	// Frequency: string, required
	Frequency terra.StringValue `hcl:"frequency,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Interval: number, required
	Interval terra.NumberValue `hcl:"interval,attr" validate:"required"`
	// LogicAppId: string, required
	LogicAppId terra.StringValue `hcl:"logic_app_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// StartTime: string, optional
	StartTime terra.StringValue `hcl:"start_time,attr"`
	// TimeZone: string, optional
	TimeZone terra.StringValue `hcl:"time_zone,attr"`
	// Schedule: optional
	Schedule *logicapptriggerrecurrence.Schedule `hcl:"schedule,block"`
	// Timeouts: optional
	Timeouts *logicapptriggerrecurrence.Timeouts `hcl:"timeouts,block"`
}
type logicAppTriggerRecurrenceAttributes struct {
	ref terra.Reference
}

// Frequency returns a reference to field frequency of azurerm_logic_app_trigger_recurrence.
func (latr logicAppTriggerRecurrenceAttributes) Frequency() terra.StringValue {
	return terra.ReferenceAsString(latr.ref.Append("frequency"))
}

// Id returns a reference to field id of azurerm_logic_app_trigger_recurrence.
func (latr logicAppTriggerRecurrenceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(latr.ref.Append("id"))
}

// Interval returns a reference to field interval of azurerm_logic_app_trigger_recurrence.
func (latr logicAppTriggerRecurrenceAttributes) Interval() terra.NumberValue {
	return terra.ReferenceAsNumber(latr.ref.Append("interval"))
}

// LogicAppId returns a reference to field logic_app_id of azurerm_logic_app_trigger_recurrence.
func (latr logicAppTriggerRecurrenceAttributes) LogicAppId() terra.StringValue {
	return terra.ReferenceAsString(latr.ref.Append("logic_app_id"))
}

// Name returns a reference to field name of azurerm_logic_app_trigger_recurrence.
func (latr logicAppTriggerRecurrenceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(latr.ref.Append("name"))
}

// StartTime returns a reference to field start_time of azurerm_logic_app_trigger_recurrence.
func (latr logicAppTriggerRecurrenceAttributes) StartTime() terra.StringValue {
	return terra.ReferenceAsString(latr.ref.Append("start_time"))
}

// TimeZone returns a reference to field time_zone of azurerm_logic_app_trigger_recurrence.
func (latr logicAppTriggerRecurrenceAttributes) TimeZone() terra.StringValue {
	return terra.ReferenceAsString(latr.ref.Append("time_zone"))
}

func (latr logicAppTriggerRecurrenceAttributes) Schedule() terra.ListValue[logicapptriggerrecurrence.ScheduleAttributes] {
	return terra.ReferenceAsList[logicapptriggerrecurrence.ScheduleAttributes](latr.ref.Append("schedule"))
}

func (latr logicAppTriggerRecurrenceAttributes) Timeouts() logicapptriggerrecurrence.TimeoutsAttributes {
	return terra.ReferenceAsSingle[logicapptriggerrecurrence.TimeoutsAttributes](latr.ref.Append("timeouts"))
}

type logicAppTriggerRecurrenceState struct {
	Frequency  string                                    `json:"frequency"`
	Id         string                                    `json:"id"`
	Interval   float64                                   `json:"interval"`
	LogicAppId string                                    `json:"logic_app_id"`
	Name       string                                    `json:"name"`
	StartTime  string                                    `json:"start_time"`
	TimeZone   string                                    `json:"time_zone"`
	Schedule   []logicapptriggerrecurrence.ScheduleState `json:"schedule"`
	Timeouts   *logicapptriggerrecurrence.TimeoutsState  `json:"timeouts"`
}
