// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	logicapptriggerrecurrence "github.com/golingon/terraproviders/azurerm/3.49.0/logicapptriggerrecurrence"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewLogicAppTriggerRecurrence(name string, args LogicAppTriggerRecurrenceArgs) *LogicAppTriggerRecurrence {
	return &LogicAppTriggerRecurrence{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LogicAppTriggerRecurrence)(nil)

type LogicAppTriggerRecurrence struct {
	Name  string
	Args  LogicAppTriggerRecurrenceArgs
	state *logicAppTriggerRecurrenceState
}

func (latr *LogicAppTriggerRecurrence) Type() string {
	return "azurerm_logic_app_trigger_recurrence"
}

func (latr *LogicAppTriggerRecurrence) LocalName() string {
	return latr.Name
}

func (latr *LogicAppTriggerRecurrence) Configuration() interface{} {
	return latr.Args
}

func (latr *LogicAppTriggerRecurrence) Attributes() logicAppTriggerRecurrenceAttributes {
	return logicAppTriggerRecurrenceAttributes{ref: terra.ReferenceResource(latr)}
}

func (latr *LogicAppTriggerRecurrence) ImportState(av io.Reader) error {
	latr.state = &logicAppTriggerRecurrenceState{}
	if err := json.NewDecoder(av).Decode(latr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", latr.Type(), latr.LocalName(), err)
	}
	return nil
}

func (latr *LogicAppTriggerRecurrence) State() (*logicAppTriggerRecurrenceState, bool) {
	return latr.state, latr.state != nil
}

func (latr *LogicAppTriggerRecurrence) StateMust() *logicAppTriggerRecurrenceState {
	if latr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", latr.Type(), latr.LocalName()))
	}
	return latr.state
}

func (latr *LogicAppTriggerRecurrence) DependOn() terra.Reference {
	return terra.ReferenceResource(latr)
}

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
	// DependsOn contains resources that LogicAppTriggerRecurrence depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type logicAppTriggerRecurrenceAttributes struct {
	ref terra.Reference
}

func (latr logicAppTriggerRecurrenceAttributes) Frequency() terra.StringValue {
	return terra.ReferenceString(latr.ref.Append("frequency"))
}

func (latr logicAppTriggerRecurrenceAttributes) Id() terra.StringValue {
	return terra.ReferenceString(latr.ref.Append("id"))
}

func (latr logicAppTriggerRecurrenceAttributes) Interval() terra.NumberValue {
	return terra.ReferenceNumber(latr.ref.Append("interval"))
}

func (latr logicAppTriggerRecurrenceAttributes) LogicAppId() terra.StringValue {
	return terra.ReferenceString(latr.ref.Append("logic_app_id"))
}

func (latr logicAppTriggerRecurrenceAttributes) Name() terra.StringValue {
	return terra.ReferenceString(latr.ref.Append("name"))
}

func (latr logicAppTriggerRecurrenceAttributes) StartTime() terra.StringValue {
	return terra.ReferenceString(latr.ref.Append("start_time"))
}

func (latr logicAppTriggerRecurrenceAttributes) TimeZone() terra.StringValue {
	return terra.ReferenceString(latr.ref.Append("time_zone"))
}

func (latr logicAppTriggerRecurrenceAttributes) Schedule() terra.ListValue[logicapptriggerrecurrence.ScheduleAttributes] {
	return terra.ReferenceList[logicapptriggerrecurrence.ScheduleAttributes](latr.ref.Append("schedule"))
}

func (latr logicAppTriggerRecurrenceAttributes) Timeouts() logicapptriggerrecurrence.TimeoutsAttributes {
	return terra.ReferenceSingle[logicapptriggerrecurrence.TimeoutsAttributes](latr.ref.Append("timeouts"))
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
