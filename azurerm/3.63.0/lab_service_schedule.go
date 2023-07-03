// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	labserviceschedule "github.com/golingon/terraproviders/azurerm/3.63.0/labserviceschedule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLabServiceSchedule creates a new instance of [LabServiceSchedule].
func NewLabServiceSchedule(name string, args LabServiceScheduleArgs) *LabServiceSchedule {
	return &LabServiceSchedule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LabServiceSchedule)(nil)

// LabServiceSchedule represents the Terraform resource azurerm_lab_service_schedule.
type LabServiceSchedule struct {
	Name      string
	Args      LabServiceScheduleArgs
	state     *labServiceScheduleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LabServiceSchedule].
func (lss *LabServiceSchedule) Type() string {
	return "azurerm_lab_service_schedule"
}

// LocalName returns the local name for [LabServiceSchedule].
func (lss *LabServiceSchedule) LocalName() string {
	return lss.Name
}

// Configuration returns the configuration (args) for [LabServiceSchedule].
func (lss *LabServiceSchedule) Configuration() interface{} {
	return lss.Args
}

// DependOn is used for other resources to depend on [LabServiceSchedule].
func (lss *LabServiceSchedule) DependOn() terra.Reference {
	return terra.ReferenceResource(lss)
}

// Dependencies returns the list of resources [LabServiceSchedule] depends_on.
func (lss *LabServiceSchedule) Dependencies() terra.Dependencies {
	return lss.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LabServiceSchedule].
func (lss *LabServiceSchedule) LifecycleManagement() *terra.Lifecycle {
	return lss.Lifecycle
}

// Attributes returns the attributes for [LabServiceSchedule].
func (lss *LabServiceSchedule) Attributes() labServiceScheduleAttributes {
	return labServiceScheduleAttributes{ref: terra.ReferenceResource(lss)}
}

// ImportState imports the given attribute values into [LabServiceSchedule]'s state.
func (lss *LabServiceSchedule) ImportState(av io.Reader) error {
	lss.state = &labServiceScheduleState{}
	if err := json.NewDecoder(av).Decode(lss.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lss.Type(), lss.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LabServiceSchedule] has state.
func (lss *LabServiceSchedule) State() (*labServiceScheduleState, bool) {
	return lss.state, lss.state != nil
}

// StateMust returns the state for [LabServiceSchedule]. Panics if the state is nil.
func (lss *LabServiceSchedule) StateMust() *labServiceScheduleState {
	if lss.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lss.Type(), lss.LocalName()))
	}
	return lss.state
}

// LabServiceScheduleArgs contains the configurations for azurerm_lab_service_schedule.
type LabServiceScheduleArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LabId: string, required
	LabId terra.StringValue `hcl:"lab_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Notes: string, optional
	Notes terra.StringValue `hcl:"notes,attr"`
	// StartTime: string, optional
	StartTime terra.StringValue `hcl:"start_time,attr"`
	// StopTime: string, required
	StopTime terra.StringValue `hcl:"stop_time,attr" validate:"required"`
	// TimeZone: string, required
	TimeZone terra.StringValue `hcl:"time_zone,attr" validate:"required"`
	// Recurrence: optional
	Recurrence *labserviceschedule.Recurrence `hcl:"recurrence,block"`
	// Timeouts: optional
	Timeouts *labserviceschedule.Timeouts `hcl:"timeouts,block"`
}
type labServiceScheduleAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_lab_service_schedule.
func (lss labServiceScheduleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lss.ref.Append("id"))
}

// LabId returns a reference to field lab_id of azurerm_lab_service_schedule.
func (lss labServiceScheduleAttributes) LabId() terra.StringValue {
	return terra.ReferenceAsString(lss.ref.Append("lab_id"))
}

// Name returns a reference to field name of azurerm_lab_service_schedule.
func (lss labServiceScheduleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lss.ref.Append("name"))
}

// Notes returns a reference to field notes of azurerm_lab_service_schedule.
func (lss labServiceScheduleAttributes) Notes() terra.StringValue {
	return terra.ReferenceAsString(lss.ref.Append("notes"))
}

// StartTime returns a reference to field start_time of azurerm_lab_service_schedule.
func (lss labServiceScheduleAttributes) StartTime() terra.StringValue {
	return terra.ReferenceAsString(lss.ref.Append("start_time"))
}

// StopTime returns a reference to field stop_time of azurerm_lab_service_schedule.
func (lss labServiceScheduleAttributes) StopTime() terra.StringValue {
	return terra.ReferenceAsString(lss.ref.Append("stop_time"))
}

// TimeZone returns a reference to field time_zone of azurerm_lab_service_schedule.
func (lss labServiceScheduleAttributes) TimeZone() terra.StringValue {
	return terra.ReferenceAsString(lss.ref.Append("time_zone"))
}

func (lss labServiceScheduleAttributes) Recurrence() terra.ListValue[labserviceschedule.RecurrenceAttributes] {
	return terra.ReferenceAsList[labserviceschedule.RecurrenceAttributes](lss.ref.Append("recurrence"))
}

func (lss labServiceScheduleAttributes) Timeouts() labserviceschedule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[labserviceschedule.TimeoutsAttributes](lss.ref.Append("timeouts"))
}

type labServiceScheduleState struct {
	Id         string                               `json:"id"`
	LabId      string                               `json:"lab_id"`
	Name       string                               `json:"name"`
	Notes      string                               `json:"notes"`
	StartTime  string                               `json:"start_time"`
	StopTime   string                               `json:"stop_time"`
	TimeZone   string                               `json:"time_zone"`
	Recurrence []labserviceschedule.RecurrenceState `json:"recurrence"`
	Timeouts   *labserviceschedule.TimeoutsState    `json:"timeouts"`
}
