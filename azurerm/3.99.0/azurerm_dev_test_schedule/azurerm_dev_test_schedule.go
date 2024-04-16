// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_dev_test_schedule

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource azurerm_dev_test_schedule.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermDevTestScheduleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (adts *Resource) Type() string {
	return "azurerm_dev_test_schedule"
}

// LocalName returns the local name for [Resource].
func (adts *Resource) LocalName() string {
	return adts.Name
}

// Configuration returns the configuration (args) for [Resource].
func (adts *Resource) Configuration() interface{} {
	return adts.Args
}

// DependOn is used for other resources to depend on [Resource].
func (adts *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(adts)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (adts *Resource) Dependencies() terra.Dependencies {
	return adts.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (adts *Resource) LifecycleManagement() *terra.Lifecycle {
	return adts.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (adts *Resource) Attributes() azurermDevTestScheduleAttributes {
	return azurermDevTestScheduleAttributes{ref: terra.ReferenceResource(adts)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (adts *Resource) ImportState(state io.Reader) error {
	adts.state = &azurermDevTestScheduleState{}
	if err := json.NewDecoder(state).Decode(adts.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", adts.Type(), adts.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (adts *Resource) State() (*azurermDevTestScheduleState, bool) {
	return adts.state, adts.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (adts *Resource) StateMust() *azurermDevTestScheduleState {
	if adts.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", adts.Type(), adts.LocalName()))
	}
	return adts.state
}

// Args contains the configurations for azurerm_dev_test_schedule.
type Args struct {
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
	DailyRecurrence *DailyRecurrence `hcl:"daily_recurrence,block"`
	// HourlyRecurrence: optional
	HourlyRecurrence *HourlyRecurrence `hcl:"hourly_recurrence,block"`
	// NotificationSettings: required
	NotificationSettings *NotificationSettings `hcl:"notification_settings,block" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
	// WeeklyRecurrence: optional
	WeeklyRecurrence *WeeklyRecurrence `hcl:"weekly_recurrence,block"`
}

type azurermDevTestScheduleAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_dev_test_schedule.
func (adts azurermDevTestScheduleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adts.ref.Append("id"))
}

// LabName returns a reference to field lab_name of azurerm_dev_test_schedule.
func (adts azurermDevTestScheduleAttributes) LabName() terra.StringValue {
	return terra.ReferenceAsString(adts.ref.Append("lab_name"))
}

// Location returns a reference to field location of azurerm_dev_test_schedule.
func (adts azurermDevTestScheduleAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(adts.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_dev_test_schedule.
func (adts azurermDevTestScheduleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(adts.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_dev_test_schedule.
func (adts azurermDevTestScheduleAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(adts.ref.Append("resource_group_name"))
}

// Status returns a reference to field status of azurerm_dev_test_schedule.
func (adts azurermDevTestScheduleAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(adts.ref.Append("status"))
}

// Tags returns a reference to field tags of azurerm_dev_test_schedule.
func (adts azurermDevTestScheduleAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](adts.ref.Append("tags"))
}

// TaskType returns a reference to field task_type of azurerm_dev_test_schedule.
func (adts azurermDevTestScheduleAttributes) TaskType() terra.StringValue {
	return terra.ReferenceAsString(adts.ref.Append("task_type"))
}

// TimeZoneId returns a reference to field time_zone_id of azurerm_dev_test_schedule.
func (adts azurermDevTestScheduleAttributes) TimeZoneId() terra.StringValue {
	return terra.ReferenceAsString(adts.ref.Append("time_zone_id"))
}

func (adts azurermDevTestScheduleAttributes) DailyRecurrence() terra.ListValue[DailyRecurrenceAttributes] {
	return terra.ReferenceAsList[DailyRecurrenceAttributes](adts.ref.Append("daily_recurrence"))
}

func (adts azurermDevTestScheduleAttributes) HourlyRecurrence() terra.ListValue[HourlyRecurrenceAttributes] {
	return terra.ReferenceAsList[HourlyRecurrenceAttributes](adts.ref.Append("hourly_recurrence"))
}

func (adts azurermDevTestScheduleAttributes) NotificationSettings() terra.ListValue[NotificationSettingsAttributes] {
	return terra.ReferenceAsList[NotificationSettingsAttributes](adts.ref.Append("notification_settings"))
}

func (adts azurermDevTestScheduleAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](adts.ref.Append("timeouts"))
}

func (adts azurermDevTestScheduleAttributes) WeeklyRecurrence() terra.ListValue[WeeklyRecurrenceAttributes] {
	return terra.ReferenceAsList[WeeklyRecurrenceAttributes](adts.ref.Append("weekly_recurrence"))
}

type azurermDevTestScheduleState struct {
	Id                   string                      `json:"id"`
	LabName              string                      `json:"lab_name"`
	Location             string                      `json:"location"`
	Name                 string                      `json:"name"`
	ResourceGroupName    string                      `json:"resource_group_name"`
	Status               string                      `json:"status"`
	Tags                 map[string]string           `json:"tags"`
	TaskType             string                      `json:"task_type"`
	TimeZoneId           string                      `json:"time_zone_id"`
	DailyRecurrence      []DailyRecurrenceState      `json:"daily_recurrence"`
	HourlyRecurrence     []HourlyRecurrenceState     `json:"hourly_recurrence"`
	NotificationSettings []NotificationSettingsState `json:"notification_settings"`
	Timeouts             *TimeoutsState              `json:"timeouts"`
	WeeklyRecurrence     []WeeklyRecurrenceState     `json:"weekly_recurrence"`
}
