// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	consumptionbudgetmanagementgroup "github.com/golingon/terraproviders/azurerm/3.98.0/consumptionbudgetmanagementgroup"
	"io"
)

// NewConsumptionBudgetManagementGroup creates a new instance of [ConsumptionBudgetManagementGroup].
func NewConsumptionBudgetManagementGroup(name string, args ConsumptionBudgetManagementGroupArgs) *ConsumptionBudgetManagementGroup {
	return &ConsumptionBudgetManagementGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ConsumptionBudgetManagementGroup)(nil)

// ConsumptionBudgetManagementGroup represents the Terraform resource azurerm_consumption_budget_management_group.
type ConsumptionBudgetManagementGroup struct {
	Name      string
	Args      ConsumptionBudgetManagementGroupArgs
	state     *consumptionBudgetManagementGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ConsumptionBudgetManagementGroup].
func (cbmg *ConsumptionBudgetManagementGroup) Type() string {
	return "azurerm_consumption_budget_management_group"
}

// LocalName returns the local name for [ConsumptionBudgetManagementGroup].
func (cbmg *ConsumptionBudgetManagementGroup) LocalName() string {
	return cbmg.Name
}

// Configuration returns the configuration (args) for [ConsumptionBudgetManagementGroup].
func (cbmg *ConsumptionBudgetManagementGroup) Configuration() interface{} {
	return cbmg.Args
}

// DependOn is used for other resources to depend on [ConsumptionBudgetManagementGroup].
func (cbmg *ConsumptionBudgetManagementGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(cbmg)
}

// Dependencies returns the list of resources [ConsumptionBudgetManagementGroup] depends_on.
func (cbmg *ConsumptionBudgetManagementGroup) Dependencies() terra.Dependencies {
	return cbmg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ConsumptionBudgetManagementGroup].
func (cbmg *ConsumptionBudgetManagementGroup) LifecycleManagement() *terra.Lifecycle {
	return cbmg.Lifecycle
}

// Attributes returns the attributes for [ConsumptionBudgetManagementGroup].
func (cbmg *ConsumptionBudgetManagementGroup) Attributes() consumptionBudgetManagementGroupAttributes {
	return consumptionBudgetManagementGroupAttributes{ref: terra.ReferenceResource(cbmg)}
}

// ImportState imports the given attribute values into [ConsumptionBudgetManagementGroup]'s state.
func (cbmg *ConsumptionBudgetManagementGroup) ImportState(av io.Reader) error {
	cbmg.state = &consumptionBudgetManagementGroupState{}
	if err := json.NewDecoder(av).Decode(cbmg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cbmg.Type(), cbmg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ConsumptionBudgetManagementGroup] has state.
func (cbmg *ConsumptionBudgetManagementGroup) State() (*consumptionBudgetManagementGroupState, bool) {
	return cbmg.state, cbmg.state != nil
}

// StateMust returns the state for [ConsumptionBudgetManagementGroup]. Panics if the state is nil.
func (cbmg *ConsumptionBudgetManagementGroup) StateMust() *consumptionBudgetManagementGroupState {
	if cbmg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cbmg.Type(), cbmg.LocalName()))
	}
	return cbmg.state
}

// ConsumptionBudgetManagementGroupArgs contains the configurations for azurerm_consumption_budget_management_group.
type ConsumptionBudgetManagementGroupArgs struct {
	// Amount: number, required
	Amount terra.NumberValue `hcl:"amount,attr" validate:"required"`
	// Etag: string, optional
	Etag terra.StringValue `hcl:"etag,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ManagementGroupId: string, required
	ManagementGroupId terra.StringValue `hcl:"management_group_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// TimeGrain: string, optional
	TimeGrain terra.StringValue `hcl:"time_grain,attr"`
	// Filter: optional
	Filter *consumptionbudgetmanagementgroup.Filter `hcl:"filter,block"`
	// Notification: min=1,max=5
	Notification []consumptionbudgetmanagementgroup.Notification `hcl:"notification,block" validate:"min=1,max=5"`
	// TimePeriod: required
	TimePeriod *consumptionbudgetmanagementgroup.TimePeriod `hcl:"time_period,block" validate:"required"`
	// Timeouts: optional
	Timeouts *consumptionbudgetmanagementgroup.Timeouts `hcl:"timeouts,block"`
}
type consumptionBudgetManagementGroupAttributes struct {
	ref terra.Reference
}

// Amount returns a reference to field amount of azurerm_consumption_budget_management_group.
func (cbmg consumptionBudgetManagementGroupAttributes) Amount() terra.NumberValue {
	return terra.ReferenceAsNumber(cbmg.ref.Append("amount"))
}

// Etag returns a reference to field etag of azurerm_consumption_budget_management_group.
func (cbmg consumptionBudgetManagementGroupAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(cbmg.ref.Append("etag"))
}

// Id returns a reference to field id of azurerm_consumption_budget_management_group.
func (cbmg consumptionBudgetManagementGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cbmg.ref.Append("id"))
}

// ManagementGroupId returns a reference to field management_group_id of azurerm_consumption_budget_management_group.
func (cbmg consumptionBudgetManagementGroupAttributes) ManagementGroupId() terra.StringValue {
	return terra.ReferenceAsString(cbmg.ref.Append("management_group_id"))
}

// Name returns a reference to field name of azurerm_consumption_budget_management_group.
func (cbmg consumptionBudgetManagementGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cbmg.ref.Append("name"))
}

// TimeGrain returns a reference to field time_grain of azurerm_consumption_budget_management_group.
func (cbmg consumptionBudgetManagementGroupAttributes) TimeGrain() terra.StringValue {
	return terra.ReferenceAsString(cbmg.ref.Append("time_grain"))
}

func (cbmg consumptionBudgetManagementGroupAttributes) Filter() terra.ListValue[consumptionbudgetmanagementgroup.FilterAttributes] {
	return terra.ReferenceAsList[consumptionbudgetmanagementgroup.FilterAttributes](cbmg.ref.Append("filter"))
}

func (cbmg consumptionBudgetManagementGroupAttributes) Notification() terra.SetValue[consumptionbudgetmanagementgroup.NotificationAttributes] {
	return terra.ReferenceAsSet[consumptionbudgetmanagementgroup.NotificationAttributes](cbmg.ref.Append("notification"))
}

func (cbmg consumptionBudgetManagementGroupAttributes) TimePeriod() terra.ListValue[consumptionbudgetmanagementgroup.TimePeriodAttributes] {
	return terra.ReferenceAsList[consumptionbudgetmanagementgroup.TimePeriodAttributes](cbmg.ref.Append("time_period"))
}

func (cbmg consumptionBudgetManagementGroupAttributes) Timeouts() consumptionbudgetmanagementgroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[consumptionbudgetmanagementgroup.TimeoutsAttributes](cbmg.ref.Append("timeouts"))
}

type consumptionBudgetManagementGroupState struct {
	Amount            float64                                              `json:"amount"`
	Etag              string                                               `json:"etag"`
	Id                string                                               `json:"id"`
	ManagementGroupId string                                               `json:"management_group_id"`
	Name              string                                               `json:"name"`
	TimeGrain         string                                               `json:"time_grain"`
	Filter            []consumptionbudgetmanagementgroup.FilterState       `json:"filter"`
	Notification      []consumptionbudgetmanagementgroup.NotificationState `json:"notification"`
	TimePeriod        []consumptionbudgetmanagementgroup.TimePeriodState   `json:"time_period"`
	Timeouts          *consumptionbudgetmanagementgroup.TimeoutsState      `json:"timeouts"`
}
