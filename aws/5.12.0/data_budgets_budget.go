// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	databudgetsbudget "github.com/golingon/terraproviders/aws/5.12.0/databudgetsbudget"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataBudgetsBudget creates a new instance of [DataBudgetsBudget].
func NewDataBudgetsBudget(name string, args DataBudgetsBudgetArgs) *DataBudgetsBudget {
	return &DataBudgetsBudget{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataBudgetsBudget)(nil)

// DataBudgetsBudget represents the Terraform data resource aws_budgets_budget.
type DataBudgetsBudget struct {
	Name string
	Args DataBudgetsBudgetArgs
}

// DataSource returns the Terraform object type for [DataBudgetsBudget].
func (bb *DataBudgetsBudget) DataSource() string {
	return "aws_budgets_budget"
}

// LocalName returns the local name for [DataBudgetsBudget].
func (bb *DataBudgetsBudget) LocalName() string {
	return bb.Name
}

// Configuration returns the configuration (args) for [DataBudgetsBudget].
func (bb *DataBudgetsBudget) Configuration() interface{} {
	return bb.Args
}

// Attributes returns the attributes for [DataBudgetsBudget].
func (bb *DataBudgetsBudget) Attributes() dataBudgetsBudgetAttributes {
	return dataBudgetsBudgetAttributes{ref: terra.ReferenceDataResource(bb)}
}

// DataBudgetsBudgetArgs contains the configurations for aws_budgets_budget.
type DataBudgetsBudgetArgs struct {
	// AccountId: string, optional
	AccountId terra.StringValue `hcl:"account_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NamePrefix: string, optional
	NamePrefix terra.StringValue `hcl:"name_prefix,attr"`
	// AutoAdjustData: min=0
	AutoAdjustData []databudgetsbudget.AutoAdjustData `hcl:"auto_adjust_data,block" validate:"min=0"`
	// BudgetLimit: min=0
	BudgetLimit []databudgetsbudget.BudgetLimit `hcl:"budget_limit,block" validate:"min=0"`
	// CalculatedSpend: min=0
	CalculatedSpend []databudgetsbudget.CalculatedSpend `hcl:"calculated_spend,block" validate:"min=0"`
	// CostFilter: min=0
	CostFilter []databudgetsbudget.CostFilter `hcl:"cost_filter,block" validate:"min=0"`
	// CostTypes: min=0
	CostTypes []databudgetsbudget.CostTypes `hcl:"cost_types,block" validate:"min=0"`
	// Notification: min=0
	Notification []databudgetsbudget.Notification `hcl:"notification,block" validate:"min=0"`
	// PlannedLimit: min=0
	PlannedLimit []databudgetsbudget.PlannedLimit `hcl:"planned_limit,block" validate:"min=0"`
}
type dataBudgetsBudgetAttributes struct {
	ref terra.Reference
}

// AccountId returns a reference to field account_id of aws_budgets_budget.
func (bb dataBudgetsBudgetAttributes) AccountId() terra.StringValue {
	return terra.ReferenceAsString(bb.ref.Append("account_id"))
}

// Arn returns a reference to field arn of aws_budgets_budget.
func (bb dataBudgetsBudgetAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(bb.ref.Append("arn"))
}

// BudgetExceeded returns a reference to field budget_exceeded of aws_budgets_budget.
func (bb dataBudgetsBudgetAttributes) BudgetExceeded() terra.BoolValue {
	return terra.ReferenceAsBool(bb.ref.Append("budget_exceeded"))
}

// BudgetType returns a reference to field budget_type of aws_budgets_budget.
func (bb dataBudgetsBudgetAttributes) BudgetType() terra.StringValue {
	return terra.ReferenceAsString(bb.ref.Append("budget_type"))
}

// Id returns a reference to field id of aws_budgets_budget.
func (bb dataBudgetsBudgetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bb.ref.Append("id"))
}

// Name returns a reference to field name of aws_budgets_budget.
func (bb dataBudgetsBudgetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bb.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_budgets_budget.
func (bb dataBudgetsBudgetAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(bb.ref.Append("name_prefix"))
}

// TimePeriodEnd returns a reference to field time_period_end of aws_budgets_budget.
func (bb dataBudgetsBudgetAttributes) TimePeriodEnd() terra.StringValue {
	return terra.ReferenceAsString(bb.ref.Append("time_period_end"))
}

// TimePeriodStart returns a reference to field time_period_start of aws_budgets_budget.
func (bb dataBudgetsBudgetAttributes) TimePeriodStart() terra.StringValue {
	return terra.ReferenceAsString(bb.ref.Append("time_period_start"))
}

// TimeUnit returns a reference to field time_unit of aws_budgets_budget.
func (bb dataBudgetsBudgetAttributes) TimeUnit() terra.StringValue {
	return terra.ReferenceAsString(bb.ref.Append("time_unit"))
}

func (bb dataBudgetsBudgetAttributes) AutoAdjustData() terra.ListValue[databudgetsbudget.AutoAdjustDataAttributes] {
	return terra.ReferenceAsList[databudgetsbudget.AutoAdjustDataAttributes](bb.ref.Append("auto_adjust_data"))
}

func (bb dataBudgetsBudgetAttributes) BudgetLimit() terra.ListValue[databudgetsbudget.BudgetLimitAttributes] {
	return terra.ReferenceAsList[databudgetsbudget.BudgetLimitAttributes](bb.ref.Append("budget_limit"))
}

func (bb dataBudgetsBudgetAttributes) CalculatedSpend() terra.ListValue[databudgetsbudget.CalculatedSpendAttributes] {
	return terra.ReferenceAsList[databudgetsbudget.CalculatedSpendAttributes](bb.ref.Append("calculated_spend"))
}

func (bb dataBudgetsBudgetAttributes) CostFilter() terra.SetValue[databudgetsbudget.CostFilterAttributes] {
	return terra.ReferenceAsSet[databudgetsbudget.CostFilterAttributes](bb.ref.Append("cost_filter"))
}

func (bb dataBudgetsBudgetAttributes) CostTypes() terra.ListValue[databudgetsbudget.CostTypesAttributes] {
	return terra.ReferenceAsList[databudgetsbudget.CostTypesAttributes](bb.ref.Append("cost_types"))
}

func (bb dataBudgetsBudgetAttributes) Notification() terra.SetValue[databudgetsbudget.NotificationAttributes] {
	return terra.ReferenceAsSet[databudgetsbudget.NotificationAttributes](bb.ref.Append("notification"))
}

func (bb dataBudgetsBudgetAttributes) PlannedLimit() terra.SetValue[databudgetsbudget.PlannedLimitAttributes] {
	return terra.ReferenceAsSet[databudgetsbudget.PlannedLimitAttributes](bb.ref.Append("planned_limit"))
}
