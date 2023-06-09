// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	automationjobschedule "github.com/golingon/terraproviders/azurerm/3.49.0/automationjobschedule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAutomationJobSchedule creates a new instance of [AutomationJobSchedule].
func NewAutomationJobSchedule(name string, args AutomationJobScheduleArgs) *AutomationJobSchedule {
	return &AutomationJobSchedule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AutomationJobSchedule)(nil)

// AutomationJobSchedule represents the Terraform resource azurerm_automation_job_schedule.
type AutomationJobSchedule struct {
	Name      string
	Args      AutomationJobScheduleArgs
	state     *automationJobScheduleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AutomationJobSchedule].
func (ajs *AutomationJobSchedule) Type() string {
	return "azurerm_automation_job_schedule"
}

// LocalName returns the local name for [AutomationJobSchedule].
func (ajs *AutomationJobSchedule) LocalName() string {
	return ajs.Name
}

// Configuration returns the configuration (args) for [AutomationJobSchedule].
func (ajs *AutomationJobSchedule) Configuration() interface{} {
	return ajs.Args
}

// DependOn is used for other resources to depend on [AutomationJobSchedule].
func (ajs *AutomationJobSchedule) DependOn() terra.Reference {
	return terra.ReferenceResource(ajs)
}

// Dependencies returns the list of resources [AutomationJobSchedule] depends_on.
func (ajs *AutomationJobSchedule) Dependencies() terra.Dependencies {
	return ajs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AutomationJobSchedule].
func (ajs *AutomationJobSchedule) LifecycleManagement() *terra.Lifecycle {
	return ajs.Lifecycle
}

// Attributes returns the attributes for [AutomationJobSchedule].
func (ajs *AutomationJobSchedule) Attributes() automationJobScheduleAttributes {
	return automationJobScheduleAttributes{ref: terra.ReferenceResource(ajs)}
}

// ImportState imports the given attribute values into [AutomationJobSchedule]'s state.
func (ajs *AutomationJobSchedule) ImportState(av io.Reader) error {
	ajs.state = &automationJobScheduleState{}
	if err := json.NewDecoder(av).Decode(ajs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ajs.Type(), ajs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AutomationJobSchedule] has state.
func (ajs *AutomationJobSchedule) State() (*automationJobScheduleState, bool) {
	return ajs.state, ajs.state != nil
}

// StateMust returns the state for [AutomationJobSchedule]. Panics if the state is nil.
func (ajs *AutomationJobSchedule) StateMust() *automationJobScheduleState {
	if ajs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ajs.Type(), ajs.LocalName()))
	}
	return ajs.state
}

// AutomationJobScheduleArgs contains the configurations for azurerm_automation_job_schedule.
type AutomationJobScheduleArgs struct {
	// AutomationAccountName: string, required
	AutomationAccountName terra.StringValue `hcl:"automation_account_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// JobScheduleId: string, optional
	JobScheduleId terra.StringValue `hcl:"job_schedule_id,attr"`
	// Parameters: map of string, optional
	Parameters terra.MapValue[terra.StringValue] `hcl:"parameters,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// RunOn: string, optional
	RunOn terra.StringValue `hcl:"run_on,attr"`
	// RunbookName: string, required
	RunbookName terra.StringValue `hcl:"runbook_name,attr" validate:"required"`
	// ScheduleName: string, required
	ScheduleName terra.StringValue `hcl:"schedule_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *automationjobschedule.Timeouts `hcl:"timeouts,block"`
}
type automationJobScheduleAttributes struct {
	ref terra.Reference
}

// AutomationAccountName returns a reference to field automation_account_name of azurerm_automation_job_schedule.
func (ajs automationJobScheduleAttributes) AutomationAccountName() terra.StringValue {
	return terra.ReferenceAsString(ajs.ref.Append("automation_account_name"))
}

// Id returns a reference to field id of azurerm_automation_job_schedule.
func (ajs automationJobScheduleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ajs.ref.Append("id"))
}

// JobScheduleId returns a reference to field job_schedule_id of azurerm_automation_job_schedule.
func (ajs automationJobScheduleAttributes) JobScheduleId() terra.StringValue {
	return terra.ReferenceAsString(ajs.ref.Append("job_schedule_id"))
}

// Parameters returns a reference to field parameters of azurerm_automation_job_schedule.
func (ajs automationJobScheduleAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ajs.ref.Append("parameters"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_automation_job_schedule.
func (ajs automationJobScheduleAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ajs.ref.Append("resource_group_name"))
}

// RunOn returns a reference to field run_on of azurerm_automation_job_schedule.
func (ajs automationJobScheduleAttributes) RunOn() terra.StringValue {
	return terra.ReferenceAsString(ajs.ref.Append("run_on"))
}

// RunbookName returns a reference to field runbook_name of azurerm_automation_job_schedule.
func (ajs automationJobScheduleAttributes) RunbookName() terra.StringValue {
	return terra.ReferenceAsString(ajs.ref.Append("runbook_name"))
}

// ScheduleName returns a reference to field schedule_name of azurerm_automation_job_schedule.
func (ajs automationJobScheduleAttributes) ScheduleName() terra.StringValue {
	return terra.ReferenceAsString(ajs.ref.Append("schedule_name"))
}

func (ajs automationJobScheduleAttributes) Timeouts() automationjobschedule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[automationjobschedule.TimeoutsAttributes](ajs.ref.Append("timeouts"))
}

type automationJobScheduleState struct {
	AutomationAccountName string                               `json:"automation_account_name"`
	Id                    string                               `json:"id"`
	JobScheduleId         string                               `json:"job_schedule_id"`
	Parameters            map[string]string                    `json:"parameters"`
	ResourceGroupName     string                               `json:"resource_group_name"`
	RunOn                 string                               `json:"run_on"`
	RunbookName           string                               `json:"runbook_name"`
	ScheduleName          string                               `json:"schedule_name"`
	Timeouts              *automationjobschedule.TimeoutsState `json:"timeouts"`
}
