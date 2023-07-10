// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	devtestschedule "github.com/golingon/terraproviders/azurerm/3.64.0/devtestschedule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDevTestSchedule creates a new instance of [DevTestSchedule].
func NewDevTestSchedule(name string, args DevTestScheduleArgs) *DevTestSchedule {
	return &DevTestSchedule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DevTestSchedule)(nil)

// DevTestSchedule represents the Terraform resource azurerm_dev_test_schedule.
type DevTestSchedule struct {
	Name      string
	Args      DevTestScheduleArgs
	state     *devTestScheduleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DevTestSchedule].
func (dts *DevTestSchedule) Type() string {
	return "azurerm_dev_test_schedule"
}

// LocalName returns the local name for [DevTestSchedule].
func (dts *DevTestSchedule) LocalName() string {
	return dts.Name
}

// Configuration returns the configuration (args) for [DevTestSchedule].
func (dts *DevTestSchedule) Configuration() interface{} {
	return dts.Args
}

// DependOn is used for other resources to depend on [DevTestSchedule].
func (dts *DevTestSchedule) DependOn() terra.Reference {
	return terra.ReferenceResource(dts)
}

// Dependencies returns the list of resources [DevTestSchedule] depends_on.
func (dts *DevTestSchedule) Dependencies() terra.Dependencies {
	return dts.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DevTestSchedule].
func (dts *DevTestSchedule) LifecycleManagement() *terra.Lifecycle {
	return dts.Lifecycle
}

// Attributes returns the attributes for [DevTestSchedule].
func (dts *DevTestSchedule) Attributes() devTestScheduleAttributes {
	return devTestScheduleAttributes{ref: terra.ReferenceResource(dts)}
}

// ImportState imports the given attribute values into [DevTestSchedule]'s state.
func (dts *DevTestSchedule) ImportState(av io.Reader) error {
	dts.state = &devTestScheduleState{}
	if err := json.NewDecoder(av).Decode(dts.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dts.Type(), dts.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DevTestSchedule] has state.
func (dts *DevTestSchedule) State() (*devTestScheduleState, bool) {
	return dts.state, dts.state != nil
}

// StateMust returns the state for [DevTestSchedule]. Panics if the state is nil.
func (dts *DevTestSchedule) StateMust() *devTestScheduleState {
	if dts.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dts.Type(), dts.LocalName()))
	}
	return dts.state
}

// DevTestScheduleArgs contains the configurations for azurerm_dev_test_schedule.
type DevTestScheduleArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LabName: string, required
	LabName terra.StringValue `hcl:"lab_name,attr" validate:"required"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Status: string, optional
	Status terra.StringValue `hcl:"status,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TaskType: string, required
	TaskType terra.StringValue `hcl:"task_type,attr" validate:"required"`
	// TimeZoneId: string, required
	TimeZoneId terra.StringValue `hcl:"time_zone_id,attr" validate:"required"`
	// DailyRecurrence: optional
	DailyRecurrence *devtestschedule.DailyRecurrence `hcl:"daily_recurrence,block"`
	// HourlyRecurrence: optional
	HourlyRecurrence *devtestschedule.HourlyRecurrence `hcl:"hourly_recurrence,block"`
	// NotificationSettings: required
	NotificationSettings *devtestschedule.NotificationSettings `hcl:"notification_settings,block" validate:"required"`
	// Timeouts: optional
	Timeouts *devtestschedule.Timeouts `hcl:"timeouts,block"`
	// WeeklyRecurrence: optional
	WeeklyRecurrence *devtestschedule.WeeklyRecurrence `hcl:"weekly_recurrence,block"`
}
type devTestScheduleAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_dev_test_schedule.
func (dts devTestScheduleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dts.ref.Append("id"))
}

// LabName returns a reference to field lab_name of azurerm_dev_test_schedule.
func (dts devTestScheduleAttributes) LabName() terra.StringValue {
	return terra.ReferenceAsString(dts.ref.Append("lab_name"))
}

// Location returns a reference to field location of azurerm_dev_test_schedule.
func (dts devTestScheduleAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dts.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_dev_test_schedule.
func (dts devTestScheduleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dts.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_dev_test_schedule.
func (dts devTestScheduleAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dts.ref.Append("resource_group_name"))
}

// Status returns a reference to field status of azurerm_dev_test_schedule.
func (dts devTestScheduleAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(dts.ref.Append("status"))
}

// Tags returns a reference to field tags of azurerm_dev_test_schedule.
func (dts devTestScheduleAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dts.ref.Append("tags"))
}

// TaskType returns a reference to field task_type of azurerm_dev_test_schedule.
func (dts devTestScheduleAttributes) TaskType() terra.StringValue {
	return terra.ReferenceAsString(dts.ref.Append("task_type"))
}

// TimeZoneId returns a reference to field time_zone_id of azurerm_dev_test_schedule.
func (dts devTestScheduleAttributes) TimeZoneId() terra.StringValue {
	return terra.ReferenceAsString(dts.ref.Append("time_zone_id"))
}

func (dts devTestScheduleAttributes) DailyRecurrence() terra.ListValue[devtestschedule.DailyRecurrenceAttributes] {
	return terra.ReferenceAsList[devtestschedule.DailyRecurrenceAttributes](dts.ref.Append("daily_recurrence"))
}

func (dts devTestScheduleAttributes) HourlyRecurrence() terra.ListValue[devtestschedule.HourlyRecurrenceAttributes] {
	return terra.ReferenceAsList[devtestschedule.HourlyRecurrenceAttributes](dts.ref.Append("hourly_recurrence"))
}

func (dts devTestScheduleAttributes) NotificationSettings() terra.ListValue[devtestschedule.NotificationSettingsAttributes] {
	return terra.ReferenceAsList[devtestschedule.NotificationSettingsAttributes](dts.ref.Append("notification_settings"))
}

func (dts devTestScheduleAttributes) Timeouts() devtestschedule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[devtestschedule.TimeoutsAttributes](dts.ref.Append("timeouts"))
}

func (dts devTestScheduleAttributes) WeeklyRecurrence() terra.ListValue[devtestschedule.WeeklyRecurrenceAttributes] {
	return terra.ReferenceAsList[devtestschedule.WeeklyRecurrenceAttributes](dts.ref.Append("weekly_recurrence"))
}

type devTestScheduleState struct {
	Id                   string                                      `json:"id"`
	LabName              string                                      `json:"lab_name"`
	Location             string                                      `json:"location"`
	Name                 string                                      `json:"name"`
	ResourceGroupName    string                                      `json:"resource_group_name"`
	Status               string                                      `json:"status"`
	Tags                 map[string]string                           `json:"tags"`
	TaskType             string                                      `json:"task_type"`
	TimeZoneId           string                                      `json:"time_zone_id"`
	DailyRecurrence      []devtestschedule.DailyRecurrenceState      `json:"daily_recurrence"`
	HourlyRecurrence     []devtestschedule.HourlyRecurrenceState     `json:"hourly_recurrence"`
	NotificationSettings []devtestschedule.NotificationSettingsState `json:"notification_settings"`
	Timeouts             *devtestschedule.TimeoutsState              `json:"timeouts"`
	WeeklyRecurrence     []devtestschedule.WeeklyRecurrenceState     `json:"weekly_recurrence"`
}
