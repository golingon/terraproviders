// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_compute_network_firewall_policy_rule

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource google_compute_network_firewall_policy_rule.
type Resource struct {
	Name      string
	Args      Args
	state     *googleComputeNetworkFirewallPolicyRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gcnfpr *Resource) Type() string {
	return "google_compute_network_firewall_policy_rule"
}

// LocalName returns the local name for [Resource].
func (gcnfpr *Resource) LocalName() string {
	return gcnfpr.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gcnfpr *Resource) Configuration() interface{} {
	return gcnfpr.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gcnfpr *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gcnfpr)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gcnfpr *Resource) Dependencies() terra.Dependencies {
	return gcnfpr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gcnfpr *Resource) LifecycleManagement() *terra.Lifecycle {
	return gcnfpr.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gcnfpr *Resource) Attributes() googleComputeNetworkFirewallPolicyRuleAttributes {
	return googleComputeNetworkFirewallPolicyRuleAttributes{ref: terra.ReferenceResource(gcnfpr)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gcnfpr *Resource) ImportState(state io.Reader) error {
	gcnfpr.state = &googleComputeNetworkFirewallPolicyRuleState{}
	if err := json.NewDecoder(state).Decode(gcnfpr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gcnfpr.Type(), gcnfpr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gcnfpr *Resource) State() (*googleComputeNetworkFirewallPolicyRuleState, bool) {
	return gcnfpr.state, gcnfpr.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gcnfpr *Resource) StateMust() *googleComputeNetworkFirewallPolicyRuleState {
	if gcnfpr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gcnfpr.Type(), gcnfpr.LocalName()))
	}
	return gcnfpr.state
}

// Args contains the configurations for google_compute_network_firewall_policy_rule.
type Args struct {
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
	Match *Match `hcl:"match,block" validate:"required"`
	// TargetSecureTags: min=0
	TargetSecureTags []TargetSecureTags `hcl:"target_secure_tags,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleComputeNetworkFirewallPolicyRuleAttributes struct {
	ref terra.Reference
}

// Action returns a reference to field action of google_compute_network_firewall_policy_rule.
func (gcnfpr googleComputeNetworkFirewallPolicyRuleAttributes) Action() terra.StringValue {
	return terra.ReferenceAsString(gcnfpr.ref.Append("action"))
}

// Description returns a reference to field description of google_compute_network_firewall_policy_rule.
func (gcnfpr googleComputeNetworkFirewallPolicyRuleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gcnfpr.ref.Append("description"))
}

// Direction returns a reference to field direction of google_compute_network_firewall_policy_rule.
func (gcnfpr googleComputeNetworkFirewallPolicyRuleAttributes) Direction() terra.StringValue {
	return terra.ReferenceAsString(gcnfpr.ref.Append("direction"))
}

// Disabled returns a reference to field disabled of google_compute_network_firewall_policy_rule.
func (gcnfpr googleComputeNetworkFirewallPolicyRuleAttributes) Disabled() terra.BoolValue {
	return terra.ReferenceAsBool(gcnfpr.ref.Append("disabled"))
}

// EnableLogging returns a reference to field enable_logging of google_compute_network_firewall_policy_rule.
func (gcnfpr googleComputeNetworkFirewallPolicyRuleAttributes) EnableLogging() terra.BoolValue {
	return terra.ReferenceAsBool(gcnfpr.ref.Append("enable_logging"))
}

// FirewallPolicy returns a reference to field firewall_policy of google_compute_network_firewall_policy_rule.
func (gcnfpr googleComputeNetworkFirewallPolicyRuleAttributes) FirewallPolicy() terra.StringValue {
	return terra.ReferenceAsString(gcnfpr.ref.Append("firewall_policy"))
}

// Id returns a reference to field id of google_compute_network_firewall_policy_rule.
func (gcnfpr googleComputeNetworkFirewallPolicyRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gcnfpr.ref.Append("id"))
}

// Kind returns a reference to field kind of google_compute_network_firewall_policy_rule.
func (gcnfpr googleComputeNetworkFirewallPolicyRuleAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(gcnfpr.ref.Append("kind"))
}

// Priority returns a reference to field priority of google_compute_network_firewall_policy_rule.
func (gcnfpr googleComputeNetworkFirewallPolicyRuleAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(gcnfpr.ref.Append("priority"))
}

// Project returns a reference to field project of google_compute_network_firewall_policy_rule.
func (gcnfpr googleComputeNetworkFirewallPolicyRuleAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gcnfpr.ref.Append("project"))
}

// RuleName returns a reference to field rule_name of google_compute_network_firewall_policy_rule.
func (gcnfpr googleComputeNetworkFirewallPolicyRuleAttributes) RuleName() terra.StringValue {
	return terra.ReferenceAsString(gcnfpr.ref.Append("rule_name"))
}

// RuleTupleCount returns a reference to field rule_tuple_count of google_compute_network_firewall_policy_rule.
func (gcnfpr googleComputeNetworkFirewallPolicyRuleAttributes) RuleTupleCount() terra.NumberValue {
	return terra.ReferenceAsNumber(gcnfpr.ref.Append("rule_tuple_count"))
}

// TargetServiceAccounts returns a reference to field target_service_accounts of google_compute_network_firewall_policy_rule.
func (gcnfpr googleComputeNetworkFirewallPolicyRuleAttributes) TargetServiceAccounts() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](gcnfpr.ref.Append("target_service_accounts"))
}

func (gcnfpr googleComputeNetworkFirewallPolicyRuleAttributes) Match() terra.ListValue[MatchAttributes] {
	return terra.ReferenceAsList[MatchAttributes](gcnfpr.ref.Append("match"))
}

func (gcnfpr googleComputeNetworkFirewallPolicyRuleAttributes) TargetSecureTags() terra.ListValue[TargetSecureTagsAttributes] {
	return terra.ReferenceAsList[TargetSecureTagsAttributes](gcnfpr.ref.Append("target_secure_tags"))
}

func (gcnfpr googleComputeNetworkFirewallPolicyRuleAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gcnfpr.ref.Append("timeouts"))
}

type googleComputeNetworkFirewallPolicyRuleState struct {
	Action                string                  `json:"action"`
	Description           string                  `json:"description"`
	Direction             string                  `json:"direction"`
	Disabled              bool                    `json:"disabled"`
	EnableLogging         bool                    `json:"enable_logging"`
	FirewallPolicy        string                  `json:"firewall_policy"`
	Id                    string                  `json:"id"`
	Kind                  string                  `json:"kind"`
	Priority              float64                 `json:"priority"`
	Project               string                  `json:"project"`
	RuleName              string                  `json:"rule_name"`
	RuleTupleCount        float64                 `json:"rule_tuple_count"`
	TargetServiceAccounts []string                `json:"target_service_accounts"`
	Match                 []MatchState            `json:"match"`
	TargetSecureTags      []TargetSecureTagsState `json:"target_secure_tags"`
	Timeouts              *TimeoutsState          `json:"timeouts"`
}
