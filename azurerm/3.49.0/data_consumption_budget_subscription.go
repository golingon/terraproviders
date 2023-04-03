// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataconsumptionbudgetsubscription "github.com/golingon/terraproviders/azurerm/3.49.0/dataconsumptionbudgetsubscription"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataConsumptionBudgetSubscription creates a new instance of [DataConsumptionBudgetSubscription].
func NewDataConsumptionBudgetSubscription(name string, args DataConsumptionBudgetSubscriptionArgs) *DataConsumptionBudgetSubscription {
	return &DataConsumptionBudgetSubscription{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataConsumptionBudgetSubscription)(nil)

// DataConsumptionBudgetSubscription represents the Terraform data resource azurerm_consumption_budget_subscription.
type DataConsumptionBudgetSubscription struct {
	Name string
	Args DataConsumptionBudgetSubscriptionArgs
}

// DataSource returns the Terraform object type for [DataConsumptionBudgetSubscription].
func (cbs *DataConsumptionBudgetSubscription) DataSource() string {
	return "azurerm_consumption_budget_subscription"
}

// LocalName returns the local name for [DataConsumptionBudgetSubscription].
func (cbs *DataConsumptionBudgetSubscription) LocalName() string {
	return cbs.Name
}

// Configuration returns the configuration (args) for [DataConsumptionBudgetSubscription].
func (cbs *DataConsumptionBudgetSubscription) Configuration() interface{} {
	return cbs.Args
}

// Attributes returns the attributes for [DataConsumptionBudgetSubscription].
func (cbs *DataConsumptionBudgetSubscription) Attributes() dataConsumptionBudgetSubscriptionAttributes {
	return dataConsumptionBudgetSubscriptionAttributes{ref: terra.ReferenceDataResource(cbs)}
}

// DataConsumptionBudgetSubscriptionArgs contains the configurations for azurerm_consumption_budget_subscription.
type DataConsumptionBudgetSubscriptionArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SubscriptionId: string, required
	SubscriptionId terra.StringValue `hcl:"subscription_id,attr" validate:"required"`
	// Filter: min=0
	Filter []dataconsumptionbudgetsubscription.Filter `hcl:"filter,block" validate:"min=0"`
	// Notification: min=0
	Notification []dataconsumptionbudgetsubscription.Notification `hcl:"notification,block" validate:"min=0"`
	// TimePeriod: min=0
	TimePeriod []dataconsumptionbudgetsubscription.TimePeriod `hcl:"time_period,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataconsumptionbudgetsubscription.Timeouts `hcl:"timeouts,block"`
}
type dataConsumptionBudgetSubscriptionAttributes struct {
	ref terra.Reference
}

// Amount returns a reference to field amount of azurerm_consumption_budget_subscription.
func (cbs dataConsumptionBudgetSubscriptionAttributes) Amount() terra.NumberValue {
	return terra.ReferenceAsNumber(cbs.ref.Append("amount"))
}

// Id returns a reference to field id of azurerm_consumption_budget_subscription.
func (cbs dataConsumptionBudgetSubscriptionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cbs.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_consumption_budget_subscription.
func (cbs dataConsumptionBudgetSubscriptionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cbs.ref.Append("name"))
}

// SubscriptionId returns a reference to field subscription_id of azurerm_consumption_budget_subscription.
func (cbs dataConsumptionBudgetSubscriptionAttributes) SubscriptionId() terra.StringValue {
	return terra.ReferenceAsString(cbs.ref.Append("subscription_id"))
}

// TimeGrain returns a reference to field time_grain of azurerm_consumption_budget_subscription.
func (cbs dataConsumptionBudgetSubscriptionAttributes) TimeGrain() terra.StringValue {
	return terra.ReferenceAsString(cbs.ref.Append("time_grain"))
}

func (cbs dataConsumptionBudgetSubscriptionAttributes) Filter() terra.ListValue[dataconsumptionbudgetsubscription.FilterAttributes] {
	return terra.ReferenceAsList[dataconsumptionbudgetsubscription.FilterAttributes](cbs.ref.Append("filter"))
}

func (cbs dataConsumptionBudgetSubscriptionAttributes) Notification() terra.ListValue[dataconsumptionbudgetsubscription.NotificationAttributes] {
	return terra.ReferenceAsList[dataconsumptionbudgetsubscription.NotificationAttributes](cbs.ref.Append("notification"))
}

func (cbs dataConsumptionBudgetSubscriptionAttributes) TimePeriod() terra.ListValue[dataconsumptionbudgetsubscription.TimePeriodAttributes] {
	return terra.ReferenceAsList[dataconsumptionbudgetsubscription.TimePeriodAttributes](cbs.ref.Append("time_period"))
}

func (cbs dataConsumptionBudgetSubscriptionAttributes) Timeouts() dataconsumptionbudgetsubscription.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataconsumptionbudgetsubscription.TimeoutsAttributes](cbs.ref.Append("timeouts"))
}
