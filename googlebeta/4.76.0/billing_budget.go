// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	billingbudget "github.com/golingon/terraproviders/googlebeta/4.76.0/billingbudget"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBillingBudget creates a new instance of [BillingBudget].
func NewBillingBudget(name string, args BillingBudgetArgs) *BillingBudget {
	return &BillingBudget{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BillingBudget)(nil)

// BillingBudget represents the Terraform resource google_billing_budget.
type BillingBudget struct {
	Name      string
	Args      BillingBudgetArgs
	state     *billingBudgetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BillingBudget].
func (bb *BillingBudget) Type() string {
	return "google_billing_budget"
}

// LocalName returns the local name for [BillingBudget].
func (bb *BillingBudget) LocalName() string {
	return bb.Name
}

// Configuration returns the configuration (args) for [BillingBudget].
func (bb *BillingBudget) Configuration() interface{} {
	return bb.Args
}

// DependOn is used for other resources to depend on [BillingBudget].
func (bb *BillingBudget) DependOn() terra.Reference {
	return terra.ReferenceResource(bb)
}

// Dependencies returns the list of resources [BillingBudget] depends_on.
func (bb *BillingBudget) Dependencies() terra.Dependencies {
	return bb.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BillingBudget].
func (bb *BillingBudget) LifecycleManagement() *terra.Lifecycle {
	return bb.Lifecycle
}

// Attributes returns the attributes for [BillingBudget].
func (bb *BillingBudget) Attributes() billingBudgetAttributes {
	return billingBudgetAttributes{ref: terra.ReferenceResource(bb)}
}

// ImportState imports the given attribute values into [BillingBudget]'s state.
func (bb *BillingBudget) ImportState(av io.Reader) error {
	bb.state = &billingBudgetState{}
	if err := json.NewDecoder(av).Decode(bb.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bb.Type(), bb.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BillingBudget] has state.
func (bb *BillingBudget) State() (*billingBudgetState, bool) {
	return bb.state, bb.state != nil
}

// StateMust returns the state for [BillingBudget]. Panics if the state is nil.
func (bb *BillingBudget) StateMust() *billingBudgetState {
	if bb.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bb.Type(), bb.LocalName()))
	}
	return bb.state
}

// BillingBudgetArgs contains the configurations for google_billing_budget.
type BillingBudgetArgs struct {
	// BillingAccount: string, required
	BillingAccount terra.StringValue `hcl:"billing_account,attr" validate:"required"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// AllUpdatesRule: optional
	AllUpdatesRule *billingbudget.AllUpdatesRule `hcl:"all_updates_rule,block"`
	// Amount: required
	Amount *billingbudget.Amount `hcl:"amount,block" validate:"required"`
	// BudgetFilter: optional
	BudgetFilter *billingbudget.BudgetFilter `hcl:"budget_filter,block"`
	// ThresholdRules: min=0
	ThresholdRules []billingbudget.ThresholdRules `hcl:"threshold_rules,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *billingbudget.Timeouts `hcl:"timeouts,block"`
}
type billingBudgetAttributes struct {
	ref terra.Reference
}

// BillingAccount returns a reference to field billing_account of google_billing_budget.
func (bb billingBudgetAttributes) BillingAccount() terra.StringValue {
	return terra.ReferenceAsString(bb.ref.Append("billing_account"))
}

// DisplayName returns a reference to field display_name of google_billing_budget.
func (bb billingBudgetAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(bb.ref.Append("display_name"))
}

// Id returns a reference to field id of google_billing_budget.
func (bb billingBudgetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bb.ref.Append("id"))
}

// Name returns a reference to field name of google_billing_budget.
func (bb billingBudgetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bb.ref.Append("name"))
}

func (bb billingBudgetAttributes) AllUpdatesRule() terra.ListValue[billingbudget.AllUpdatesRuleAttributes] {
	return terra.ReferenceAsList[billingbudget.AllUpdatesRuleAttributes](bb.ref.Append("all_updates_rule"))
}

func (bb billingBudgetAttributes) Amount() terra.ListValue[billingbudget.AmountAttributes] {
	return terra.ReferenceAsList[billingbudget.AmountAttributes](bb.ref.Append("amount"))
}

func (bb billingBudgetAttributes) BudgetFilter() terra.ListValue[billingbudget.BudgetFilterAttributes] {
	return terra.ReferenceAsList[billingbudget.BudgetFilterAttributes](bb.ref.Append("budget_filter"))
}

func (bb billingBudgetAttributes) ThresholdRules() terra.ListValue[billingbudget.ThresholdRulesAttributes] {
	return terra.ReferenceAsList[billingbudget.ThresholdRulesAttributes](bb.ref.Append("threshold_rules"))
}

func (bb billingBudgetAttributes) Timeouts() billingbudget.TimeoutsAttributes {
	return terra.ReferenceAsSingle[billingbudget.TimeoutsAttributes](bb.ref.Append("timeouts"))
}

type billingBudgetState struct {
	BillingAccount string                              `json:"billing_account"`
	DisplayName    string                              `json:"display_name"`
	Id             string                              `json:"id"`
	Name           string                              `json:"name"`
	AllUpdatesRule []billingbudget.AllUpdatesRuleState `json:"all_updates_rule"`
	Amount         []billingbudget.AmountState         `json:"amount"`
	BudgetFilter   []billingbudget.BudgetFilterState   `json:"budget_filter"`
	ThresholdRules []billingbudget.ThresholdRulesState `json:"threshold_rules"`
	Timeouts       *billingbudget.TimeoutsState        `json:"timeouts"`
}
