// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	consumptionbudgetresourcegroup "github.com/golingon/terraproviders/azurerm/3.64.0/consumptionbudgetresourcegroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewConsumptionBudgetResourceGroup creates a new instance of [ConsumptionBudgetResourceGroup].
func NewConsumptionBudgetResourceGroup(name string, args ConsumptionBudgetResourceGroupArgs) *ConsumptionBudgetResourceGroup {
	return &ConsumptionBudgetResourceGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ConsumptionBudgetResourceGroup)(nil)

// ConsumptionBudgetResourceGroup represents the Terraform resource azurerm_consumption_budget_resource_group.
type ConsumptionBudgetResourceGroup struct {
	Name      string
	Args      ConsumptionBudgetResourceGroupArgs
	state     *consumptionBudgetResourceGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ConsumptionBudgetResourceGroup].
func (cbrg *ConsumptionBudgetResourceGroup) Type() string {
	return "azurerm_consumption_budget_resource_group"
}

// LocalName returns the local name for [ConsumptionBudgetResourceGroup].
func (cbrg *ConsumptionBudgetResourceGroup) LocalName() string {
	return cbrg.Name
}

// Configuration returns the configuration (args) for [ConsumptionBudgetResourceGroup].
func (cbrg *ConsumptionBudgetResourceGroup) Configuration() interface{} {
	return cbrg.Args
}

// DependOn is used for other resources to depend on [ConsumptionBudgetResourceGroup].
func (cbrg *ConsumptionBudgetResourceGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(cbrg)
}

// Dependencies returns the list of resources [ConsumptionBudgetResourceGroup] depends_on.
func (cbrg *ConsumptionBudgetResourceGroup) Dependencies() terra.Dependencies {
	return cbrg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ConsumptionBudgetResourceGroup].
func (cbrg *ConsumptionBudgetResourceGroup) LifecycleManagement() *terra.Lifecycle {
	return cbrg.Lifecycle
}

// Attributes returns the attributes for [ConsumptionBudgetResourceGroup].
func (cbrg *ConsumptionBudgetResourceGroup) Attributes() consumptionBudgetResourceGroupAttributes {
	return consumptionBudgetResourceGroupAttributes{ref: terra.ReferenceResource(cbrg)}
}

// ImportState imports the given attribute values into [ConsumptionBudgetResourceGroup]'s state.
func (cbrg *ConsumptionBudgetResourceGroup) ImportState(av io.Reader) error {
	cbrg.state = &consumptionBudgetResourceGroupState{}
	if err := json.NewDecoder(av).Decode(cbrg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cbrg.Type(), cbrg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ConsumptionBudgetResourceGroup] has state.
func (cbrg *ConsumptionBudgetResourceGroup) State() (*consumptionBudgetResourceGroupState, bool) {
	return cbrg.state, cbrg.state != nil
}

// StateMust returns the state for [ConsumptionBudgetResourceGroup]. Panics if the state is nil.
func (cbrg *ConsumptionBudgetResourceGroup) StateMust() *consumptionBudgetResourceGroupState {
	if cbrg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cbrg.Type(), cbrg.LocalName()))
	}
	return cbrg.state
}

// ConsumptionBudgetResourceGroupArgs contains the configurations for azurerm_consumption_budget_resource_group.
type ConsumptionBudgetResourceGroupArgs struct {
	// Amount: number, required
	Amount terra.NumberValue `hcl:"amount,attr" validate:"required"`
	// Etag: string, optional
	Etag terra.StringValue `hcl:"etag,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupId: string, required
	ResourceGroupId terra.StringValue `hcl:"resource_group_id,attr" validate:"required"`
	// TimeGrain: string, optional
	TimeGrain terra.StringValue `hcl:"time_grain,attr"`
	// Filter: optional
	Filter *consumptionbudgetresourcegroup.Filter `hcl:"filter,block"`
	// Notification: min=1,max=5
	Notification []consumptionbudgetresourcegroup.Notification `hcl:"notification,block" validate:"min=1,max=5"`
	// TimePeriod: required
	TimePeriod *consumptionbudgetresourcegroup.TimePeriod `hcl:"time_period,block" validate:"required"`
	// Timeouts: optional
	Timeouts *consumptionbudgetresourcegroup.Timeouts `hcl:"timeouts,block"`
}
type consumptionBudgetResourceGroupAttributes struct {
	ref terra.Reference
}

// Amount returns a reference to field amount of azurerm_consumption_budget_resource_group.
func (cbrg consumptionBudgetResourceGroupAttributes) Amount() terra.NumberValue {
	return terra.ReferenceAsNumber(cbrg.ref.Append("amount"))
}

// Etag returns a reference to field etag of azurerm_consumption_budget_resource_group.
func (cbrg consumptionBudgetResourceGroupAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(cbrg.ref.Append("etag"))
}

// Id returns a reference to field id of azurerm_consumption_budget_resource_group.
func (cbrg consumptionBudgetResourceGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cbrg.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_consumption_budget_resource_group.
func (cbrg consumptionBudgetResourceGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cbrg.ref.Append("name"))
}

// ResourceGroupId returns a reference to field resource_group_id of azurerm_consumption_budget_resource_group.
func (cbrg consumptionBudgetResourceGroupAttributes) ResourceGroupId() terra.StringValue {
	return terra.ReferenceAsString(cbrg.ref.Append("resource_group_id"))
}

// TimeGrain returns a reference to field time_grain of azurerm_consumption_budget_resource_group.
func (cbrg consumptionBudgetResourceGroupAttributes) TimeGrain() terra.StringValue {
	return terra.ReferenceAsString(cbrg.ref.Append("time_grain"))
}

func (cbrg consumptionBudgetResourceGroupAttributes) Filter() terra.ListValue[consumptionbudgetresourcegroup.FilterAttributes] {
	return terra.ReferenceAsList[consumptionbudgetresourcegroup.FilterAttributes](cbrg.ref.Append("filter"))
}

func (cbrg consumptionBudgetResourceGroupAttributes) Notification() terra.SetValue[consumptionbudgetresourcegroup.NotificationAttributes] {
	return terra.ReferenceAsSet[consumptionbudgetresourcegroup.NotificationAttributes](cbrg.ref.Append("notification"))
}

func (cbrg consumptionBudgetResourceGroupAttributes) TimePeriod() terra.ListValue[consumptionbudgetresourcegroup.TimePeriodAttributes] {
	return terra.ReferenceAsList[consumptionbudgetresourcegroup.TimePeriodAttributes](cbrg.ref.Append("time_period"))
}

func (cbrg consumptionBudgetResourceGroupAttributes) Timeouts() consumptionbudgetresourcegroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[consumptionbudgetresourcegroup.TimeoutsAttributes](cbrg.ref.Append("timeouts"))
}

type consumptionBudgetResourceGroupState struct {
	Amount          float64                                            `json:"amount"`
	Etag            string                                             `json:"etag"`
	Id              string                                             `json:"id"`
	Name            string                                             `json:"name"`
	ResourceGroupId string                                             `json:"resource_group_id"`
	TimeGrain       string                                             `json:"time_grain"`
	Filter          []consumptionbudgetresourcegroup.FilterState       `json:"filter"`
	Notification    []consumptionbudgetresourcegroup.NotificationState `json:"notification"`
	TimePeriod      []consumptionbudgetresourcegroup.TimePeriodState   `json:"time_period"`
	Timeouts        *consumptionbudgetresourcegroup.TimeoutsState      `json:"timeouts"`
}
