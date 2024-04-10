// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"github.com/golingon/lingon/pkg/terra"
	dataconsumptionbudgetresourcegroup "github.com/golingon/terraproviders/azurerm/3.98.0/dataconsumptionbudgetresourcegroup"
)

// NewDataConsumptionBudgetResourceGroup creates a new instance of [DataConsumptionBudgetResourceGroup].
func NewDataConsumptionBudgetResourceGroup(name string, args DataConsumptionBudgetResourceGroupArgs) *DataConsumptionBudgetResourceGroup {
	return &DataConsumptionBudgetResourceGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataConsumptionBudgetResourceGroup)(nil)

// DataConsumptionBudgetResourceGroup represents the Terraform data resource azurerm_consumption_budget_resource_group.
type DataConsumptionBudgetResourceGroup struct {
	Name string
	Args DataConsumptionBudgetResourceGroupArgs
}

// DataSource returns the Terraform object type for [DataConsumptionBudgetResourceGroup].
func (cbrg *DataConsumptionBudgetResourceGroup) DataSource() string {
	return "azurerm_consumption_budget_resource_group"
}

// LocalName returns the local name for [DataConsumptionBudgetResourceGroup].
func (cbrg *DataConsumptionBudgetResourceGroup) LocalName() string {
	return cbrg.Name
}

// Configuration returns the configuration (args) for [DataConsumptionBudgetResourceGroup].
func (cbrg *DataConsumptionBudgetResourceGroup) Configuration() interface{} {
	return cbrg.Args
}

// Attributes returns the attributes for [DataConsumptionBudgetResourceGroup].
func (cbrg *DataConsumptionBudgetResourceGroup) Attributes() dataConsumptionBudgetResourceGroupAttributes {
	return dataConsumptionBudgetResourceGroupAttributes{ref: terra.ReferenceDataResource(cbrg)}
}

// DataConsumptionBudgetResourceGroupArgs contains the configurations for azurerm_consumption_budget_resource_group.
type DataConsumptionBudgetResourceGroupArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupId: string, required
	ResourceGroupId terra.StringValue `hcl:"resource_group_id,attr" validate:"required"`
	// Filter: min=0
	Filter []dataconsumptionbudgetresourcegroup.Filter `hcl:"filter,block" validate:"min=0"`
	// Notification: min=0
	Notification []dataconsumptionbudgetresourcegroup.Notification `hcl:"notification,block" validate:"min=0"`
	// TimePeriod: min=0
	TimePeriod []dataconsumptionbudgetresourcegroup.TimePeriod `hcl:"time_period,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataconsumptionbudgetresourcegroup.Timeouts `hcl:"timeouts,block"`
}
type dataConsumptionBudgetResourceGroupAttributes struct {
	ref terra.Reference
}

// Amount returns a reference to field amount of azurerm_consumption_budget_resource_group.
func (cbrg dataConsumptionBudgetResourceGroupAttributes) Amount() terra.NumberValue {
	return terra.ReferenceAsNumber(cbrg.ref.Append("amount"))
}

// Id returns a reference to field id of azurerm_consumption_budget_resource_group.
func (cbrg dataConsumptionBudgetResourceGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cbrg.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_consumption_budget_resource_group.
func (cbrg dataConsumptionBudgetResourceGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cbrg.ref.Append("name"))
}

// ResourceGroupId returns a reference to field resource_group_id of azurerm_consumption_budget_resource_group.
func (cbrg dataConsumptionBudgetResourceGroupAttributes) ResourceGroupId() terra.StringValue {
	return terra.ReferenceAsString(cbrg.ref.Append("resource_group_id"))
}

// TimeGrain returns a reference to field time_grain of azurerm_consumption_budget_resource_group.
func (cbrg dataConsumptionBudgetResourceGroupAttributes) TimeGrain() terra.StringValue {
	return terra.ReferenceAsString(cbrg.ref.Append("time_grain"))
}

func (cbrg dataConsumptionBudgetResourceGroupAttributes) Filter() terra.ListValue[dataconsumptionbudgetresourcegroup.FilterAttributes] {
	return terra.ReferenceAsList[dataconsumptionbudgetresourcegroup.FilterAttributes](cbrg.ref.Append("filter"))
}

func (cbrg dataConsumptionBudgetResourceGroupAttributes) Notification() terra.ListValue[dataconsumptionbudgetresourcegroup.NotificationAttributes] {
	return terra.ReferenceAsList[dataconsumptionbudgetresourcegroup.NotificationAttributes](cbrg.ref.Append("notification"))
}

func (cbrg dataConsumptionBudgetResourceGroupAttributes) TimePeriod() terra.ListValue[dataconsumptionbudgetresourcegroup.TimePeriodAttributes] {
	return terra.ReferenceAsList[dataconsumptionbudgetresourcegroup.TimePeriodAttributes](cbrg.ref.Append("time_period"))
}

func (cbrg dataConsumptionBudgetResourceGroupAttributes) Timeouts() dataconsumptionbudgetresourcegroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataconsumptionbudgetresourcegroup.TimeoutsAttributes](cbrg.ref.Append("timeouts"))
}
