// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computenetworkfirewallpolicyrule "github.com/golingon/terraproviders/google/4.75.1/computenetworkfirewallpolicyrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeNetworkFirewallPolicyRule creates a new instance of [ComputeNetworkFirewallPolicyRule].
func NewComputeNetworkFirewallPolicyRule(name string, args ComputeNetworkFirewallPolicyRuleArgs) *ComputeNetworkFirewallPolicyRule {
	return &ComputeNetworkFirewallPolicyRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeNetworkFirewallPolicyRule)(nil)

// ComputeNetworkFirewallPolicyRule represents the Terraform resource google_compute_network_firewall_policy_rule.
type ComputeNetworkFirewallPolicyRule struct {
	Name      string
	Args      ComputeNetworkFirewallPolicyRuleArgs
	state     *computeNetworkFirewallPolicyRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeNetworkFirewallPolicyRule].
func (cnfpr *ComputeNetworkFirewallPolicyRule) Type() string {
	return "google_compute_network_firewall_policy_rule"
}

// LocalName returns the local name for [ComputeNetworkFirewallPolicyRule].
func (cnfpr *ComputeNetworkFirewallPolicyRule) LocalName() string {
	return cnfpr.Name
}

// Configuration returns the configuration (args) for [ComputeNetworkFirewallPolicyRule].
func (cnfpr *ComputeNetworkFirewallPolicyRule) Configuration() interface{} {
	return cnfpr.Args
}

// DependOn is used for other resources to depend on [ComputeNetworkFirewallPolicyRule].
func (cnfpr *ComputeNetworkFirewallPolicyRule) DependOn() terra.Reference {
	return terra.ReferenceResource(cnfpr)
}

// Dependencies returns the list of resources [ComputeNetworkFirewallPolicyRule] depends_on.
func (cnfpr *ComputeNetworkFirewallPolicyRule) Dependencies() terra.Dependencies {
	return cnfpr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeNetworkFirewallPolicyRule].
func (cnfpr *ComputeNetworkFirewallPolicyRule) LifecycleManagement() *terra.Lifecycle {
	return cnfpr.Lifecycle
}

// Attributes returns the attributes for [ComputeNetworkFirewallPolicyRule].
func (cnfpr *ComputeNetworkFirewallPolicyRule) Attributes() computeNetworkFirewallPolicyRuleAttributes {
	return computeNetworkFirewallPolicyRuleAttributes{ref: terra.ReferenceResource(cnfpr)}
}

// ImportState imports the given attribute values into [ComputeNetworkFirewallPolicyRule]'s state.
func (cnfpr *ComputeNetworkFirewallPolicyRule) ImportState(av io.Reader) error {
	cnfpr.state = &computeNetworkFirewallPolicyRuleState{}
	if err := json.NewDecoder(av).Decode(cnfpr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cnfpr.Type(), cnfpr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeNetworkFirewallPolicyRule] has state.
func (cnfpr *ComputeNetworkFirewallPolicyRule) State() (*computeNetworkFirewallPolicyRuleState, bool) {
	return cnfpr.state, cnfpr.state != nil
}

// StateMust returns the state for [ComputeNetworkFirewallPolicyRule]. Panics if the state is nil.
func (cnfpr *ComputeNetworkFirewallPolicyRule) StateMust() *computeNetworkFirewallPolicyRuleState {
	if cnfpr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cnfpr.Type(), cnfpr.LocalName()))
	}
	return cnfpr.state
}

// ComputeNetworkFirewallPolicyRuleArgs contains the configurations for google_compute_network_firewall_policy_rule.
type ComputeNetworkFirewallPolicyRuleArgs struct {
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
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// RuleName: string, optional
	RuleName terra.StringValue `hcl:"rule_name,attr"`
	// TargetServiceAccounts: list of string, optional
	TargetServiceAccounts terra.ListValue[terra.StringValue] `hcl:"target_service_accounts,attr"`
	// Match: required
	Match *computenetworkfirewallpolicyrule.Match `hcl:"match,block" validate:"required"`
	// TargetSecureTags: min=0
	TargetSecureTags []computenetworkfirewallpolicyrule.TargetSecureTags `hcl:"target_secure_tags,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *computenetworkfirewallpolicyrule.Timeouts `hcl:"timeouts,block"`
}
type computeNetworkFirewallPolicyRuleAttributes struct {
	ref terra.Reference
}

// Action returns a reference to field action of google_compute_network_firewall_policy_rule.
func (cnfpr computeNetworkFirewallPolicyRuleAttributes) Action() terra.StringValue {
	return terra.ReferenceAsString(cnfpr.ref.Append("action"))
}

// Description returns a reference to field description of google_compute_network_firewall_policy_rule.
func (cnfpr computeNetworkFirewallPolicyRuleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cnfpr.ref.Append("description"))
}

// Direction returns a reference to field direction of google_compute_network_firewall_policy_rule.
func (cnfpr computeNetworkFirewallPolicyRuleAttributes) Direction() terra.StringValue {
	return terra.ReferenceAsString(cnfpr.ref.Append("direction"))
}

// Disabled returns a reference to field disabled of google_compute_network_firewall_policy_rule.
func (cnfpr computeNetworkFirewallPolicyRuleAttributes) Disabled() terra.BoolValue {
	return terra.ReferenceAsBool(cnfpr.ref.Append("disabled"))
}

// EnableLogging returns a reference to field enable_logging of google_compute_network_firewall_policy_rule.
func (cnfpr computeNetworkFirewallPolicyRuleAttributes) EnableLogging() terra.BoolValue {
	return terra.ReferenceAsBool(cnfpr.ref.Append("enable_logging"))
}

// FirewallPolicy returns a reference to field firewall_policy of google_compute_network_firewall_policy_rule.
func (cnfpr computeNetworkFirewallPolicyRuleAttributes) FirewallPolicy() terra.StringValue {
	return terra.ReferenceAsString(cnfpr.ref.Append("firewall_policy"))
}

// Id returns a reference to field id of google_compute_network_firewall_policy_rule.
func (cnfpr computeNetworkFirewallPolicyRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cnfpr.ref.Append("id"))
}

// Kind returns a reference to field kind of google_compute_network_firewall_policy_rule.
func (cnfpr computeNetworkFirewallPolicyRuleAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(cnfpr.ref.Append("kind"))
}

// Priority returns a reference to field priority of google_compute_network_firewall_policy_rule.
func (cnfpr computeNetworkFirewallPolicyRuleAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(cnfpr.ref.Append("priority"))
}

// Project returns a reference to field project of google_compute_network_firewall_policy_rule.
func (cnfpr computeNetworkFirewallPolicyRuleAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cnfpr.ref.Append("project"))
}

// RuleName returns a reference to field rule_name of google_compute_network_firewall_policy_rule.
func (cnfpr computeNetworkFirewallPolicyRuleAttributes) RuleName() terra.StringValue {
	return terra.ReferenceAsString(cnfpr.ref.Append("rule_name"))
}

// RuleTupleCount returns a reference to field rule_tuple_count of google_compute_network_firewall_policy_rule.
func (cnfpr computeNetworkFirewallPolicyRuleAttributes) RuleTupleCount() terra.NumberValue {
	return terra.ReferenceAsNumber(cnfpr.ref.Append("rule_tuple_count"))
}

// TargetServiceAccounts returns a reference to field target_service_accounts of google_compute_network_firewall_policy_rule.
func (cnfpr computeNetworkFirewallPolicyRuleAttributes) TargetServiceAccounts() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cnfpr.ref.Append("target_service_accounts"))
}

func (cnfpr computeNetworkFirewallPolicyRuleAttributes) Match() terra.ListValue[computenetworkfirewallpolicyrule.MatchAttributes] {
	return terra.ReferenceAsList[computenetworkfirewallpolicyrule.MatchAttributes](cnfpr.ref.Append("match"))
}

func (cnfpr computeNetworkFirewallPolicyRuleAttributes) TargetSecureTags() terra.ListValue[computenetworkfirewallpolicyrule.TargetSecureTagsAttributes] {
	return terra.ReferenceAsList[computenetworkfirewallpolicyrule.TargetSecureTagsAttributes](cnfpr.ref.Append("target_secure_tags"))
}

func (cnfpr computeNetworkFirewallPolicyRuleAttributes) Timeouts() computenetworkfirewallpolicyrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computenetworkfirewallpolicyrule.TimeoutsAttributes](cnfpr.ref.Append("timeouts"))
}

type computeNetworkFirewallPolicyRuleState struct {
	Action                string                                                   `json:"action"`
	Description           string                                                   `json:"description"`
	Direction             string                                                   `json:"direction"`
	Disabled              bool                                                     `json:"disabled"`
	EnableLogging         bool                                                     `json:"enable_logging"`
	FirewallPolicy        string                                                   `json:"firewall_policy"`
	Id                    string                                                   `json:"id"`
	Kind                  string                                                   `json:"kind"`
	Priority              float64                                                  `json:"priority"`
	Project               string                                                   `json:"project"`
	RuleName              string                                                   `json:"rule_name"`
	RuleTupleCount        float64                                                  `json:"rule_tuple_count"`
	TargetServiceAccounts []string                                                 `json:"target_service_accounts"`
	Match                 []computenetworkfirewallpolicyrule.MatchState            `json:"match"`
	TargetSecureTags      []computenetworkfirewallpolicyrule.TargetSecureTagsState `json:"target_secure_tags"`
	Timeouts              *computenetworkfirewallpolicyrule.TimeoutsState          `json:"timeouts"`
}
