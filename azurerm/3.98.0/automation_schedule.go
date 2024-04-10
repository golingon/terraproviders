// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	automationschedule "github.com/golingon/terraproviders/azurerm/3.98.0/automationschedule"
	"io"
)

// NewAutomationSchedule creates a new instance of [AutomationSchedule].
func NewAutomationSchedule(name string, args AutomationScheduleArgs) *AutomationSchedule {
	return &AutomationSchedule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AutomationSchedule)(nil)

// AutomationSchedule represents the Terraform resource azurerm_automation_schedule.
type AutomationSchedule struct {
	Name      string
	Args      AutomationScheduleArgs
	state     *automationScheduleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AutomationSchedule].
func (as *AutomationSchedule) Type() string {
	return "azurerm_automation_schedule"
}

// LocalName returns the local name for [AutomationSchedule].
func (as *AutomationSchedule) LocalName() string {
	return as.Name
}

// Configuration returns the configuration (args) for [AutomationSchedule].
func (as *AutomationSchedule) Configuration() interface{} {
	return as.Args
}

// DependOn is used for other resources to depend on [AutomationSchedule].
func (as *AutomationSchedule) DependOn() terra.Reference {
	return terra.ReferenceResource(as)
}

// Dependencies returns the list of resources [AutomationSchedule] depends_on.
func (as *AutomationSchedule) Dependencies() terra.Dependencies {
	return as.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AutomationSchedule].
func (as *AutomationSchedule) LifecycleManagement() *terra.Lifecycle {
	return as.Lifecycle
}

// Attributes returns the attributes for [AutomationSchedule].
func (as *AutomationSchedule) Attributes() automationScheduleAttributes {
	return automationScheduleAttributes{ref: terra.ReferenceResource(as)}
}

// ImportState imports the given attribute values into [AutomationSchedule]'s state.
func (as *AutomationSchedule) ImportState(av io.Reader) error {
	as.state = &automationScheduleState{}
	if err := json.NewDecoder(av).Decode(as.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", as.Type(), as.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AutomationSchedule] has state.
func (as *AutomationSchedule) State() (*automationScheduleState, bool) {
	return as.state, as.state != nil
}

// StateMust returns the state for [AutomationSchedule]. Panics if the state is nil.
func (as *AutomationSchedule) StateMust() *automationScheduleState {
	if as.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", as.Type(), as.LocalName()))
	}
	return as.state
}

// AutomationScheduleArgs contains the configurations for azurerm_automation_schedule.
type AutomationScheduleArgs struct {
	// AutomationAccountName: string, required
	AutomationAccountName terra.StringValue `hcl:"automation_account_name,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// ExpiryTime: string, optional
	ExpiryTime terra.StringValue `hcl:"expiry_time,attr"`
	// Frequency: string, required
	Frequency terra.StringValue `hcl:"frequency,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Interval: number, optional
	Interval terra.NumberValue `hcl:"interval,attr"`
	// MonthDays: set of number, optional
	MonthDays terra.SetValue[terra.NumberValue] `hcl:"month_days,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// StartTime: string, optional
	StartTime terra.StringValue `hcl:"start_time,attr"`
	// Timezone: string, optional
	Timezone terra.StringValue `hcl:"timezone,attr"`
	// WeekDays: set of string, optional
	WeekDays terra.SetValue[terra.StringValue] `hcl:"week_days,attr"`
	// MonthlyOccurrence: optional
	MonthlyOccurrence *automationschedule.MonthlyOccurrence `hcl:"monthly_occurrence,block"`
	// Timeouts: optional
	Timeouts *automationschedule.Timeouts `hcl:"timeouts,block"`
}
type automationScheduleAttributes struct {
	ref terra.Reference
}

// AutomationAccountName returns a reference to field automation_account_name of azurerm_automation_schedule.
func (as automationScheduleAttributes) AutomationAccountName() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("automation_account_name"))
}

// Description returns a reference to field description of azurerm_automation_schedule.
func (as automationScheduleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("description"))
}

// ExpiryTime returns a reference to field expiry_time of azurerm_automation_schedule.
func (as automationScheduleAttributes) ExpiryTime() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("expiry_time"))
}

// Frequency returns a reference to field frequency of azurerm_automation_schedule.
func (as automationScheduleAttributes) Frequency() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("frequency"))
}

// Id returns a reference to field id of azurerm_automation_schedule.
func (as automationScheduleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("id"))
}

// Interval returns a reference to field interval of azurerm_automation_schedule.
func (as automationScheduleAttributes) Interval() terra.NumberValue {
	return terra.ReferenceAsNumber(as.ref.Append("interval"))
}

// MonthDays returns a reference to field month_days of azurerm_automation_schedule.
func (as automationScheduleAttributes) MonthDays() terra.SetValue[terra.NumberValue] {
	return terra.ReferenceAsSet[terra.NumberValue](as.ref.Append("month_days"))
}

// Name returns a reference to field name of azurerm_automation_schedule.
func (as automationScheduleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_automation_schedule.
func (as automationScheduleAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("resource_group_name"))
}

// StartTime returns a reference to field start_time of azurerm_automation_schedule.
func (as automationScheduleAttributes) StartTime() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("start_time"))
}

// Timezone returns a reference to field timezone of azurerm_automation_schedule.
func (as automationScheduleAttributes) Timezone() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("timezone"))
}

// WeekDays returns a reference to field week_days of azurerm_automation_schedule.
func (as automationScheduleAttributes) WeekDays() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](as.ref.Append("week_days"))
}

func (as automationScheduleAttributes) MonthlyOccurrence() terra.ListValue[automationschedule.MonthlyOccurrenceAttributes] {
	return terra.ReferenceAsList[automationschedule.MonthlyOccurrenceAttributes](as.ref.Append("monthly_occurrence"))
}

func (as automationScheduleAttributes) Timeouts() automationschedule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[automationschedule.TimeoutsAttributes](as.ref.Append("timeouts"))
}

type automationScheduleState struct {
	AutomationAccountName string                                      `json:"automation_account_name"`
	Description           string                                      `json:"description"`
	ExpiryTime            string                                      `json:"expiry_time"`
	Frequency             string                                      `json:"frequency"`
	Id                    string                                      `json:"id"`
	Interval              float64                                     `json:"interval"`
	MonthDays             []float64                                   `json:"month_days"`
	Name                  string                                      `json:"name"`
	ResourceGroupName     string                                      `json:"resource_group_name"`
	StartTime             string                                      `json:"start_time"`
	Timezone              string                                      `json:"timezone"`
	WeekDays              []string                                    `json:"week_days"`
	MonthlyOccurrence     []automationschedule.MonthlyOccurrenceState `json:"monthly_occurrence"`
	Timeouts              *automationschedule.TimeoutsState           `json:"timeouts"`
}
