// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	costmanagementscheduledaction "github.com/golingon/terraproviders/azurerm/3.98.0/costmanagementscheduledaction"
	"io"
)

// NewCostManagementScheduledAction creates a new instance of [CostManagementScheduledAction].
func NewCostManagementScheduledAction(name string, args CostManagementScheduledActionArgs) *CostManagementScheduledAction {
	return &CostManagementScheduledAction{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CostManagementScheduledAction)(nil)

// CostManagementScheduledAction represents the Terraform resource azurerm_cost_management_scheduled_action.
type CostManagementScheduledAction struct {
	Name      string
	Args      CostManagementScheduledActionArgs
	state     *costManagementScheduledActionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CostManagementScheduledAction].
func (cmsa *CostManagementScheduledAction) Type() string {
	return "azurerm_cost_management_scheduled_action"
}

// LocalName returns the local name for [CostManagementScheduledAction].
func (cmsa *CostManagementScheduledAction) LocalName() string {
	return cmsa.Name
}

// Configuration returns the configuration (args) for [CostManagementScheduledAction].
func (cmsa *CostManagementScheduledAction) Configuration() interface{} {
	return cmsa.Args
}

// DependOn is used for other resources to depend on [CostManagementScheduledAction].
func (cmsa *CostManagementScheduledAction) DependOn() terra.Reference {
	return terra.ReferenceResource(cmsa)
}

// Dependencies returns the list of resources [CostManagementScheduledAction] depends_on.
func (cmsa *CostManagementScheduledAction) Dependencies() terra.Dependencies {
	return cmsa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CostManagementScheduledAction].
func (cmsa *CostManagementScheduledAction) LifecycleManagement() *terra.Lifecycle {
	return cmsa.Lifecycle
}

// Attributes returns the attributes for [CostManagementScheduledAction].
func (cmsa *CostManagementScheduledAction) Attributes() costManagementScheduledActionAttributes {
	return costManagementScheduledActionAttributes{ref: terra.ReferenceResource(cmsa)}
}

// ImportState imports the given attribute values into [CostManagementScheduledAction]'s state.
func (cmsa *CostManagementScheduledAction) ImportState(av io.Reader) error {
	cmsa.state = &costManagementScheduledActionState{}
	if err := json.NewDecoder(av).Decode(cmsa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cmsa.Type(), cmsa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CostManagementScheduledAction] has state.
func (cmsa *CostManagementScheduledAction) State() (*costManagementScheduledActionState, bool) {
	return cmsa.state, cmsa.state != nil
}

// StateMust returns the state for [CostManagementScheduledAction]. Panics if the state is nil.
func (cmsa *CostManagementScheduledAction) StateMust() *costManagementScheduledActionState {
	if cmsa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cmsa.Type(), cmsa.LocalName()))
	}
	return cmsa.state
}

// CostManagementScheduledActionArgs contains the configurations for azurerm_cost_management_scheduled_action.
type CostManagementScheduledActionArgs struct {
	// DayOfMonth: number, optional
	DayOfMonth terra.NumberValue `hcl:"day_of_month,attr"`
	// DaysOfWeek: list of string, optional
	DaysOfWeek terra.ListValue[terra.StringValue] `hcl:"days_of_week,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// EmailAddressSender: string, required
	EmailAddressSender terra.StringValue `hcl:"email_address_sender,attr" validate:"required"`
	// EmailAddresses: list of string, required
	EmailAddresses terra.ListValue[terra.StringValue] `hcl:"email_addresses,attr" validate:"required"`
	// EmailSubject: string, required
	EmailSubject terra.StringValue `hcl:"email_subject,attr" validate:"required"`
	// EndDate: string, required
	EndDate terra.StringValue `hcl:"end_date,attr" validate:"required"`
	// Frequency: string, required
	Frequency terra.StringValue `hcl:"frequency,attr" validate:"required"`
	// HourOfDay: number, optional
	HourOfDay terra.NumberValue `hcl:"hour_of_day,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Message: string, optional
	Message terra.StringValue `hcl:"message,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// StartDate: string, required
	StartDate terra.StringValue `hcl:"start_date,attr" validate:"required"`
	// ViewId: string, required
	ViewId terra.StringValue `hcl:"view_id,attr" validate:"required"`
	// WeeksOfMonth: list of string, optional
	WeeksOfMonth terra.ListValue[terra.StringValue] `hcl:"weeks_of_month,attr"`
	// Timeouts: optional
	Timeouts *costmanagementscheduledaction.Timeouts `hcl:"timeouts,block"`
}
type costManagementScheduledActionAttributes struct {
	ref terra.Reference
}

// DayOfMonth returns a reference to field day_of_month of azurerm_cost_management_scheduled_action.
func (cmsa costManagementScheduledActionAttributes) DayOfMonth() terra.NumberValue {
	return terra.ReferenceAsNumber(cmsa.ref.Append("day_of_month"))
}

// DaysOfWeek returns a reference to field days_of_week of azurerm_cost_management_scheduled_action.
func (cmsa costManagementScheduledActionAttributes) DaysOfWeek() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cmsa.ref.Append("days_of_week"))
}

// DisplayName returns a reference to field display_name of azurerm_cost_management_scheduled_action.
func (cmsa costManagementScheduledActionAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(cmsa.ref.Append("display_name"))
}

// EmailAddressSender returns a reference to field email_address_sender of azurerm_cost_management_scheduled_action.
func (cmsa costManagementScheduledActionAttributes) EmailAddressSender() terra.StringValue {
	return terra.ReferenceAsString(cmsa.ref.Append("email_address_sender"))
}

// EmailAddresses returns a reference to field email_addresses of azurerm_cost_management_scheduled_action.
func (cmsa costManagementScheduledActionAttributes) EmailAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cmsa.ref.Append("email_addresses"))
}

// EmailSubject returns a reference to field email_subject of azurerm_cost_management_scheduled_action.
func (cmsa costManagementScheduledActionAttributes) EmailSubject() terra.StringValue {
	return terra.ReferenceAsString(cmsa.ref.Append("email_subject"))
}

// EndDate returns a reference to field end_date of azurerm_cost_management_scheduled_action.
func (cmsa costManagementScheduledActionAttributes) EndDate() terra.StringValue {
	return terra.ReferenceAsString(cmsa.ref.Append("end_date"))
}

// Frequency returns a reference to field frequency of azurerm_cost_management_scheduled_action.
func (cmsa costManagementScheduledActionAttributes) Frequency() terra.StringValue {
	return terra.ReferenceAsString(cmsa.ref.Append("frequency"))
}

// HourOfDay returns a reference to field hour_of_day of azurerm_cost_management_scheduled_action.
func (cmsa costManagementScheduledActionAttributes) HourOfDay() terra.NumberValue {
	return terra.ReferenceAsNumber(cmsa.ref.Append("hour_of_day"))
}

// Id returns a reference to field id of azurerm_cost_management_scheduled_action.
func (cmsa costManagementScheduledActionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cmsa.ref.Append("id"))
}

// Message returns a reference to field message of azurerm_cost_management_scheduled_action.
func (cmsa costManagementScheduledActionAttributes) Message() terra.StringValue {
	return terra.ReferenceAsString(cmsa.ref.Append("message"))
}

// Name returns a reference to field name of azurerm_cost_management_scheduled_action.
func (cmsa costManagementScheduledActionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cmsa.ref.Append("name"))
}

// StartDate returns a reference to field start_date of azurerm_cost_management_scheduled_action.
func (cmsa costManagementScheduledActionAttributes) StartDate() terra.StringValue {
	return terra.ReferenceAsString(cmsa.ref.Append("start_date"))
}

// ViewId returns a reference to field view_id of azurerm_cost_management_scheduled_action.
func (cmsa costManagementScheduledActionAttributes) ViewId() terra.StringValue {
	return terra.ReferenceAsString(cmsa.ref.Append("view_id"))
}

// WeeksOfMonth returns a reference to field weeks_of_month of azurerm_cost_management_scheduled_action.
func (cmsa costManagementScheduledActionAttributes) WeeksOfMonth() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cmsa.ref.Append("weeks_of_month"))
}

func (cmsa costManagementScheduledActionAttributes) Timeouts() costmanagementscheduledaction.TimeoutsAttributes {
	return terra.ReferenceAsSingle[costmanagementscheduledaction.TimeoutsAttributes](cmsa.ref.Append("timeouts"))
}

type costManagementScheduledActionState struct {
	DayOfMonth         float64                                      `json:"day_of_month"`
	DaysOfWeek         []string                                     `json:"days_of_week"`
	DisplayName        string                                       `json:"display_name"`
	EmailAddressSender string                                       `json:"email_address_sender"`
	EmailAddresses     []string                                     `json:"email_addresses"`
	EmailSubject       string                                       `json:"email_subject"`
	EndDate            string                                       `json:"end_date"`
	Frequency          string                                       `json:"frequency"`
	HourOfDay          float64                                      `json:"hour_of_day"`
	Id                 string                                       `json:"id"`
	Message            string                                       `json:"message"`
	Name               string                                       `json:"name"`
	StartDate          string                                       `json:"start_date"`
	ViewId             string                                       `json:"view_id"`
	WeeksOfMonth       []string                                     `json:"weeks_of_month"`
	Timeouts           *costmanagementscheduledaction.TimeoutsState `json:"timeouts"`
}
