// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computefirewallpolicyrule "github.com/golingon/terraproviders/google/4.75.1/computefirewallpolicyrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeFirewallPolicyRule creates a new instance of [ComputeFirewallPolicyRule].
func NewComputeFirewallPolicyRule(name string, args ComputeFirewallPolicyRuleArgs) *ComputeFirewallPolicyRule {
	return &ComputeFirewallPolicyRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeFirewallPolicyRule)(nil)

// ComputeFirewallPolicyRule represents the Terraform resource google_compute_firewall_policy_rule.
type ComputeFirewallPolicyRule struct {
	Name      string
	Args      ComputeFirewallPolicyRuleArgs
	state     *computeFirewallPolicyRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeFirewallPolicyRule].
func (cfpr *ComputeFirewallPolicyRule) Type() string {
	return "google_compute_firewall_policy_rule"
}

// LocalName returns the local name for [ComputeFirewallPolicyRule].
func (cfpr *ComputeFirewallPolicyRule) LocalName() string {
	return cfpr.Name
}

// Configuration returns the configuration (args) for [ComputeFirewallPolicyRule].
func (cfpr *ComputeFirewallPolicyRule) Configuration() interface{} {
	return cfpr.Args
}

// DependOn is used for other resources to depend on [ComputeFirewallPolicyRule].
func (cfpr *ComputeFirewallPolicyRule) DependOn() terra.Reference {
	return terra.ReferenceResource(cfpr)
}

// Dependencies returns the list of resources [ComputeFirewallPolicyRule] depends_on.
func (cfpr *ComputeFirewallPolicyRule) Dependencies() terra.Dependencies {
	return cfpr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeFirewallPolicyRule].
func (cfpr *ComputeFirewallPolicyRule) LifecycleManagement() *terra.Lifecycle {
	return cfpr.Lifecycle
}

// Attributes returns the attributes for [ComputeFirewallPolicyRule].
func (cfpr *ComputeFirewallPolicyRule) Attributes() computeFirewallPolicyRuleAttributes {
	return computeFirewallPolicyRuleAttributes{ref: terra.ReferenceResource(cfpr)}
}

// ImportState imports the given attribute values into [ComputeFirewallPolicyRule]'s state.
func (cfpr *ComputeFirewallPolicyRule) ImportState(av io.Reader) error {
	cfpr.state = &computeFirewallPolicyRuleState{}
	if err := json.NewDecoder(av).Decode(cfpr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cfpr.Type(), cfpr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeFirewallPolicyRule] has state.
func (cfpr *ComputeFirewallPolicyRule) State() (*computeFirewallPolicyRuleState, bool) {
	return cfpr.state, cfpr.state != nil
}

// StateMust returns the state for [ComputeFirewallPolicyRule]. Panics if the state is nil.
func (cfpr *ComputeFirewallPolicyRule) StateMust() *computeFirewallPolicyRuleState {
	if cfpr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cfpr.Type(), cfpr.LocalName()))
	}
	return cfpr.state
}

// ComputeFirewallPolicyRuleArgs contains the configurations for google_compute_firewall_policy_rule.
type ComputeFirewallPolicyRuleArgs struct {
	// Action: string, required
	Action terra.StringValue `hcl:"action,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Direction: string, required
	Direction terra.StringValue `hcl:"direction,attr" validate:"required"`
	// Disabled: bool, optional
	Disabled terra.BoolValue `hcl:"disabled,attr"`
	// EnableLogging: bool, optional
	EnableLogging terra.BoolValue `hcl:"enable_logging,attr"`
	// FirewallPolicy: string, required
	FirewallPolicy terra.StringValue `hcl:"firewall_policy,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Priority: number, required
	Priority terra.NumberValue `hcl:"priority,attr" validate:"required"`
	// TargetResources: list of string, optional
	TargetResources terra.ListValue[terra.StringValue] `hcl:"target_resources,attr"`
	// TargetServiceAccounts: list of string, optional
	TargetServiceAccounts terra.ListValue[terra.StringValue] `hcl:"target_service_accounts,attr"`
	// Match: required
	Match *computefirewallpolicyrule.Match `hcl:"match,block" validate:"required"`
	// Timeouts: optional
	Timeouts *computefirewallpolicyrule.Timeouts `hcl:"timeouts,block"`
}
type computeFirewallPolicyRuleAttributes struct {
	ref terra.Reference
}

// Action returns a reference to field action of google_compute_firewall_policy_rule.
func (cfpr computeFirewallPolicyRuleAttributes) Action() terra.StringValue {
	return terra.ReferenceAsString(cfpr.ref.Append("action"))
}

// Description returns a reference to field description of google_compute_firewall_policy_rule.
func (cfpr computeFirewallPolicyRuleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cfpr.ref.Append("description"))
}

// Direction returns a reference to field direction of google_compute_firewall_policy_rule.
func (cfpr computeFirewallPolicyRuleAttributes) Direction() terra.StringValue {
	return terra.ReferenceAsString(cfpr.ref.Append("direction"))
}

// Disabled returns a reference to field disabled of google_compute_firewall_policy_rule.
func (cfpr computeFirewallPolicyRuleAttributes) Disabled() terra.BoolValue {
	return terra.ReferenceAsBool(cfpr.ref.Append("disabled"))
}

// EnableLogging returns a reference to field enable_logging of google_compute_firewall_policy_rule.
func (cfpr computeFirewallPolicyRuleAttributes) EnableLogging() terra.BoolValue {
	return terra.ReferenceAsBool(cfpr.ref.Append("enable_logging"))
}

// FirewallPolicy returns a reference to field firewall_policy of google_compute_firewall_policy_rule.
func (cfpr computeFirewallPolicyRuleAttributes) FirewallPolicy() terra.StringValue {
	return terra.ReferenceAsString(cfpr.ref.Append("firewall_policy"))
}

// Id returns a reference to field id of google_compute_firewall_policy_rule.
func (cfpr computeFirewallPolicyRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cfpr.ref.Append("id"))
}

// Kind returns a reference to field kind of google_compute_firewall_policy_rule.
func (cfpr computeFirewallPolicyRuleAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(cfpr.ref.Append("kind"))
}

// Priority returns a reference to field priority of google_compute_firewall_policy_rule.
func (cfpr computeFirewallPolicyRuleAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(cfpr.ref.Append("priority"))
}

// RuleTupleCount returns a reference to field rule_tuple_count of google_compute_firewall_policy_rule.
func (cfpr computeFirewallPolicyRuleAttributes) RuleTupleCount() terra.NumberValue {
	return terra.ReferenceAsNumber(cfpr.ref.Append("rule_tuple_count"))
}

// TargetResources returns a reference to field target_resources of google_compute_firewall_policy_rule.
func (cfpr computeFirewallPolicyRuleAttributes) TargetResources() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cfpr.ref.Append("target_resources"))
}

// TargetServiceAccounts returns a reference to field target_service_accounts of google_compute_firewall_policy_rule.
func (cfpr computeFirewallPolicyRuleAttributes) TargetServiceAccounts() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cfpr.ref.Append("target_service_accounts"))
}

func (cfpr computeFirewallPolicyRuleAttributes) Match() terra.ListValue[computefirewallpolicyrule.MatchAttributes] {
	return terra.ReferenceAsList[computefirewallpolicyrule.MatchAttributes](cfpr.ref.Append("match"))
}

func (cfpr computeFirewallPolicyRuleAttributes) Timeouts() computefirewallpolicyrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computefirewallpolicyrule.TimeoutsAttributes](cfpr.ref.Append("timeouts"))
}

type computeFirewallPolicyRuleState struct {
	Action                string                                   `json:"action"`
	Description           string                                   `json:"description"`
	Direction             string                                   `json:"direction"`
	Disabled              bool                                     `json:"disabled"`
	EnableLogging         bool                                     `json:"enable_logging"`
	FirewallPolicy        string                                   `json:"firewall_policy"`
	Id                    string                                   `json:"id"`
	Kind                  string                                   `json:"kind"`
	Priority              float64                                  `json:"priority"`
	RuleTupleCount        float64                                  `json:"rule_tuple_count"`
	TargetResources       []string                                 `json:"target_resources"`
	TargetServiceAccounts []string                                 `json:"target_service_accounts"`
	Match                 []computefirewallpolicyrule.MatchState   `json:"match"`
	Timeouts              *computefirewallpolicyrule.TimeoutsState `json:"timeouts"`
}
