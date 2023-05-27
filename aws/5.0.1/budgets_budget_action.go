// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	budgetsbudgetaction "github.com/golingon/terraproviders/aws/5.0.1/budgetsbudgetaction"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBudgetsBudgetAction creates a new instance of [BudgetsBudgetAction].
func NewBudgetsBudgetAction(name string, args BudgetsBudgetActionArgs) *BudgetsBudgetAction {
	return &BudgetsBudgetAction{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BudgetsBudgetAction)(nil)

// BudgetsBudgetAction represents the Terraform resource aws_budgets_budget_action.
type BudgetsBudgetAction struct {
	Name      string
	Args      BudgetsBudgetActionArgs
	state     *budgetsBudgetActionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BudgetsBudgetAction].
func (bba *BudgetsBudgetAction) Type() string {
	return "aws_budgets_budget_action"
}

// LocalName returns the local name for [BudgetsBudgetAction].
func (bba *BudgetsBudgetAction) LocalName() string {
	return bba.Name
}

// Configuration returns the configuration (args) for [BudgetsBudgetAction].
func (bba *BudgetsBudgetAction) Configuration() interface{} {
	return bba.Args
}

// DependOn is used for other resources to depend on [BudgetsBudgetAction].
func (bba *BudgetsBudgetAction) DependOn() terra.Reference {
	return terra.ReferenceResource(bba)
}

// Dependencies returns the list of resources [BudgetsBudgetAction] depends_on.
func (bba *BudgetsBudgetAction) Dependencies() terra.Dependencies {
	return bba.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BudgetsBudgetAction].
func (bba *BudgetsBudgetAction) LifecycleManagement() *terra.Lifecycle {
	return bba.Lifecycle
}

// Attributes returns the attributes for [BudgetsBudgetAction].
func (bba *BudgetsBudgetAction) Attributes() budgetsBudgetActionAttributes {
	return budgetsBudgetActionAttributes{ref: terra.ReferenceResource(bba)}
}

// ImportState imports the given attribute values into [BudgetsBudgetAction]'s state.
func (bba *BudgetsBudgetAction) ImportState(av io.Reader) error {
	bba.state = &budgetsBudgetActionState{}
	if err := json.NewDecoder(av).Decode(bba.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bba.Type(), bba.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BudgetsBudgetAction] has state.
func (bba *BudgetsBudgetAction) State() (*budgetsBudgetActionState, bool) {
	return bba.state, bba.state != nil
}

// StateMust returns the state for [BudgetsBudgetAction]. Panics if the state is nil.
func (bba *BudgetsBudgetAction) StateMust() *budgetsBudgetActionState {
	if bba.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bba.Type(), bba.LocalName()))
	}
	return bba.state
}

// BudgetsBudgetActionArgs contains the configurations for aws_budgets_budget_action.
type BudgetsBudgetActionArgs struct {
	// AccountId: string, optional
	AccountId terra.StringValue `hcl:"account_id,attr"`
	// ActionType: string, required
	ActionType terra.StringValue `hcl:"action_type,attr" validate:"required"`
	// ApprovalModel: string, required
	ApprovalModel terra.StringValue `hcl:"approval_model,attr" validate:"required"`
	// BudgetName: string, required
	BudgetName terra.StringValue `hcl:"budget_name,attr" validate:"required"`
	// ExecutionRoleArn: string, required
	ExecutionRoleArn terra.StringValue `hcl:"execution_role_arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NotificationType: string, required
	NotificationType terra.StringValue `hcl:"notification_type,attr" validate:"required"`
	// ActionThreshold: required
	ActionThreshold *budgetsbudgetaction.ActionThreshold `hcl:"action_threshold,block" validate:"required"`
	// Definition: required
	Definition *budgetsbudgetaction.Definition `hcl:"definition,block" validate:"required"`
	// Subscriber: min=1,max=11
	Subscriber []budgetsbudgetaction.Subscriber `hcl:"subscriber,block" validate:"min=1,max=11"`
	// Timeouts: optional
	Timeouts *budgetsbudgetaction.Timeouts `hcl:"timeouts,block"`
}
type budgetsBudgetActionAttributes struct {
	ref terra.Reference
}

// AccountId returns a reference to field account_id of aws_budgets_budget_action.
func (bba budgetsBudgetActionAttributes) AccountId() terra.StringValue {
	return terra.ReferenceAsString(bba.ref.Append("account_id"))
}

// ActionId returns a reference to field action_id of aws_budgets_budget_action.
func (bba budgetsBudgetActionAttributes) ActionId() terra.StringValue {
	return terra.ReferenceAsString(bba.ref.Append("action_id"))
}

// ActionType returns a reference to field action_type of aws_budgets_budget_action.
func (bba budgetsBudgetActionAttributes) ActionType() terra.StringValue {
	return terra.ReferenceAsString(bba.ref.Append("action_type"))
}

// ApprovalModel returns a reference to field approval_model of aws_budgets_budget_action.
func (bba budgetsBudgetActionAttributes) ApprovalModel() terra.StringValue {
	return terra.ReferenceAsString(bba.ref.Append("approval_model"))
}

// Arn returns a reference to field arn of aws_budgets_budget_action.
func (bba budgetsBudgetActionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(bba.ref.Append("arn"))
}

// BudgetName returns a reference to field budget_name of aws_budgets_budget_action.
func (bba budgetsBudgetActionAttributes) BudgetName() terra.StringValue {
	return terra.ReferenceAsString(bba.ref.Append("budget_name"))
}

// ExecutionRoleArn returns a reference to field execution_role_arn of aws_budgets_budget_action.
func (bba budgetsBudgetActionAttributes) ExecutionRoleArn() terra.StringValue {
	return terra.ReferenceAsString(bba.ref.Append("execution_role_arn"))
}

// Id returns a reference to field id of aws_budgets_budget_action.
func (bba budgetsBudgetActionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bba.ref.Append("id"))
}

// NotificationType returns a reference to field notification_type of aws_budgets_budget_action.
func (bba budgetsBudgetActionAttributes) NotificationType() terra.StringValue {
	return terra.ReferenceAsString(bba.ref.Append("notification_type"))
}

// Status returns a reference to field status of aws_budgets_budget_action.
func (bba budgetsBudgetActionAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(bba.ref.Append("status"))
}

func (bba budgetsBudgetActionAttributes) ActionThreshold() terra.ListValue[budgetsbudgetaction.ActionThresholdAttributes] {
	return terra.ReferenceAsList[budgetsbudgetaction.ActionThresholdAttributes](bba.ref.Append("action_threshold"))
}

func (bba budgetsBudgetActionAttributes) Definition() terra.ListValue[budgetsbudgetaction.DefinitionAttributes] {
	return terra.ReferenceAsList[budgetsbudgetaction.DefinitionAttributes](bba.ref.Append("definition"))
}

func (bba budgetsBudgetActionAttributes) Subscriber() terra.SetValue[budgetsbudgetaction.SubscriberAttributes] {
	return terra.ReferenceAsSet[budgetsbudgetaction.SubscriberAttributes](bba.ref.Append("subscriber"))
}

func (bba budgetsBudgetActionAttributes) Timeouts() budgetsbudgetaction.TimeoutsAttributes {
	return terra.ReferenceAsSingle[budgetsbudgetaction.TimeoutsAttributes](bba.ref.Append("timeouts"))
}

type budgetsBudgetActionState struct {
	AccountId        string                                     `json:"account_id"`
	ActionId         string                                     `json:"action_id"`
	ActionType       string                                     `json:"action_type"`
	ApprovalModel    string                                     `json:"approval_model"`
	Arn              string                                     `json:"arn"`
	BudgetName       string                                     `json:"budget_name"`
	ExecutionRoleArn string                                     `json:"execution_role_arn"`
	Id               string                                     `json:"id"`
	NotificationType string                                     `json:"notification_type"`
	Status           string                                     `json:"status"`
	ActionThreshold  []budgetsbudgetaction.ActionThresholdState `json:"action_threshold"`
	Definition       []budgetsbudgetaction.DefinitionState      `json:"definition"`
	Subscriber       []budgetsbudgetaction.SubscriberState      `json:"subscriber"`
	Timeouts         *budgetsbudgetaction.TimeoutsState         `json:"timeouts"`
}
