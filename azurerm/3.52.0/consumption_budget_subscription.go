// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	consumptionbudgetsubscription "github.com/golingon/terraproviders/azurerm/3.52.0/consumptionbudgetsubscription"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewConsumptionBudgetSubscription creates a new instance of [ConsumptionBudgetSubscription].
func NewConsumptionBudgetSubscription(name string, args ConsumptionBudgetSubscriptionArgs) *ConsumptionBudgetSubscription {
	return &ConsumptionBudgetSubscription{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ConsumptionBudgetSubscription)(nil)

// ConsumptionBudgetSubscription represents the Terraform resource azurerm_consumption_budget_subscription.
type ConsumptionBudgetSubscription struct {
	Name      string
	Args      ConsumptionBudgetSubscriptionArgs
	state     *consumptionBudgetSubscriptionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ConsumptionBudgetSubscription].
func (cbs *ConsumptionBudgetSubscription) Type() string {
	return "azurerm_consumption_budget_subscription"
}

// LocalName returns the local name for [ConsumptionBudgetSubscription].
func (cbs *ConsumptionBudgetSubscription) LocalName() string {
	return cbs.Name
}

// Configuration returns the configuration (args) for [ConsumptionBudgetSubscription].
func (cbs *ConsumptionBudgetSubscription) Configuration() interface{} {
	return cbs.Args
}

// DependOn is used for other resources to depend on [ConsumptionBudgetSubscription].
func (cbs *ConsumptionBudgetSubscription) DependOn() terra.Reference {
	return terra.ReferenceResource(cbs)
}

// Dependencies returns the list of resources [ConsumptionBudgetSubscription] depends_on.
func (cbs *ConsumptionBudgetSubscription) Dependencies() terra.Dependencies {
	return cbs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ConsumptionBudgetSubscription].
func (cbs *ConsumptionBudgetSubscription) LifecycleManagement() *terra.Lifecycle {
	return cbs.Lifecycle
}

// Attributes returns the attributes for [ConsumptionBudgetSubscription].
func (cbs *ConsumptionBudgetSubscription) Attributes() consumptionBudgetSubscriptionAttributes {
	return consumptionBudgetSubscriptionAttributes{ref: terra.ReferenceResource(cbs)}
}

// ImportState imports the given attribute values into [ConsumptionBudgetSubscription]'s state.
func (cbs *ConsumptionBudgetSubscription) ImportState(av io.Reader) error {
	cbs.state = &consumptionBudgetSubscriptionState{}
	if err := json.NewDecoder(av).Decode(cbs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cbs.Type(), cbs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ConsumptionBudgetSubscription] has state.
func (cbs *ConsumptionBudgetSubscription) State() (*consumptionBudgetSubscriptionState, bool) {
	return cbs.state, cbs.state != nil
}

// StateMust returns the state for [ConsumptionBudgetSubscription]. Panics if the state is nil.
func (cbs *ConsumptionBudgetSubscription) StateMust() *consumptionBudgetSubscriptionState {
	if cbs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cbs.Type(), cbs.LocalName()))
	}
	return cbs.state
}

// ConsumptionBudgetSubscriptionArgs contains the configurations for azurerm_consumption_budget_subscription.
type ConsumptionBudgetSubscriptionArgs struct {
	// Amount: number, required
	Amount terra.NumberValue `hcl:"amount,attr" validate:"required"`
	// Etag: string, optional
	Etag terra.StringValue `hcl:"etag,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SubscriptionId: string, required
	SubscriptionId terra.StringValue `hcl:"subscription_id,attr" validate:"required"`
	// TimeGrain: string, optional
	TimeGrain terra.StringValue `hcl:"time_grain,attr"`
	// Filter: optional
	Filter *consumptionbudgetsubscription.Filter `hcl:"filter,block"`
	// Notification: min=1,max=5
	Notification []consumptionbudgetsubscription.Notification `hcl:"notification,block" validate:"min=1,max=5"`
	// TimePeriod: required
	TimePeriod *consumptionbudgetsubscription.TimePeriod `hcl:"time_period,block" validate:"required"`
	// Timeouts: optional
	Timeouts *consumptionbudgetsubscription.Timeouts `hcl:"timeouts,block"`
}
type consumptionBudgetSubscriptionAttributes struct {
	ref terra.Reference
}

// Amount returns a reference to field amount of azurerm_consumption_budget_subscription.
func (cbs consumptionBudgetSubscriptionAttributes) Amount() terra.NumberValue {
	return terra.ReferenceAsNumber(cbs.ref.Append("amount"))
}

// Etag returns a reference to field etag of azurerm_consumption_budget_subscription.
func (cbs consumptionBudgetSubscriptionAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(cbs.ref.Append("etag"))
}

// Id returns a reference to field id of azurerm_consumption_budget_subscription.
func (cbs consumptionBudgetSubscriptionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cbs.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_consumption_budget_subscription.
func (cbs consumptionBudgetSubscriptionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cbs.ref.Append("name"))
}

// SubscriptionId returns a reference to field subscription_id of azurerm_consumption_budget_subscription.
func (cbs consumptionBudgetSubscriptionAttributes) SubscriptionId() terra.StringValue {
	return terra.ReferenceAsString(cbs.ref.Append("subscription_id"))
}

// TimeGrain returns a reference to field time_grain of azurerm_consumption_budget_subscription.
func (cbs consumptionBudgetSubscriptionAttributes) TimeGrain() terra.StringValue {
	return terra.ReferenceAsString(cbs.ref.Append("time_grain"))
}

func (cbs consumptionBudgetSubscriptionAttributes) Filter() terra.ListValue[consumptionbudgetsubscription.FilterAttributes] {
	return terra.ReferenceAsList[consumptionbudgetsubscription.FilterAttributes](cbs.ref.Append("filter"))
}

func (cbs consumptionBudgetSubscriptionAttributes) Notification() terra.SetValue[consumptionbudgetsubscription.NotificationAttributes] {
	return terra.ReferenceAsSet[consumptionbudgetsubscription.NotificationAttributes](cbs.ref.Append("notification"))
}

func (cbs consumptionBudgetSubscriptionAttributes) TimePeriod() terra.ListValue[consumptionbudgetsubscription.TimePeriodAttributes] {
	return terra.ReferenceAsList[consumptionbudgetsubscription.TimePeriodAttributes](cbs.ref.Append("time_period"))
}

func (cbs consumptionBudgetSubscriptionAttributes) Timeouts() consumptionbudgetsubscription.TimeoutsAttributes {
	return terra.ReferenceAsSingle[consumptionbudgetsubscription.TimeoutsAttributes](cbs.ref.Append("timeouts"))
}

type consumptionBudgetSubscriptionState struct {
	Amount         float64                                           `json:"amount"`
	Etag           string                                            `json:"etag"`
	Id             string                                            `json:"id"`
	Name           string                                            `json:"name"`
	SubscriptionId string                                            `json:"subscription_id"`
	TimeGrain      string                                            `json:"time_grain"`
	Filter         []consumptionbudgetsubscription.FilterState       `json:"filter"`
	Notification   []consumptionbudgetsubscription.NotificationState `json:"notification"`
	TimePeriod     []consumptionbudgetsubscription.TimePeriodState   `json:"time_period"`
	Timeouts       *consumptionbudgetsubscription.TimeoutsState      `json:"timeouts"`
}
