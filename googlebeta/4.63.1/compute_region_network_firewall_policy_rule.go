// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computeregionnetworkfirewallpolicyrule "github.com/golingon/terraproviders/googlebeta/4.63.1/computeregionnetworkfirewallpolicyrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeRegionNetworkFirewallPolicyRule creates a new instance of [ComputeRegionNetworkFirewallPolicyRule].
func NewComputeRegionNetworkFirewallPolicyRule(name string, args ComputeRegionNetworkFirewallPolicyRuleArgs) *ComputeRegionNetworkFirewallPolicyRule {
	return &ComputeRegionNetworkFirewallPolicyRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeRegionNetworkFirewallPolicyRule)(nil)

// ComputeRegionNetworkFirewallPolicyRule represents the Terraform resource google_compute_region_network_firewall_policy_rule.
type ComputeRegionNetworkFirewallPolicyRule struct {
	Name      string
	Args      ComputeRegionNetworkFirewallPolicyRuleArgs
	state     *computeRegionNetworkFirewallPolicyRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeRegionNetworkFirewallPolicyRule].
func (crnfpr *ComputeRegionNetworkFirewallPolicyRule) Type() string {
	return "google_compute_region_network_firewall_policy_rule"
}

// LocalName returns the local name for [ComputeRegionNetworkFirewallPolicyRule].
func (crnfpr *ComputeRegionNetworkFirewallPolicyRule) LocalName() string {
	return crnfpr.Name
}

// Configuration returns the configuration (args) for [ComputeRegionNetworkFirewallPolicyRule].
func (crnfpr *ComputeRegionNetworkFirewallPolicyRule) Configuration() interface{} {
	return crnfpr.Args
}

// DependOn is used for other resources to depend on [ComputeRegionNetworkFirewallPolicyRule].
func (crnfpr *ComputeRegionNetworkFirewallPolicyRule) DependOn() terra.Reference {
	return terra.ReferenceResource(crnfpr)
}

// Dependencies returns the list of resources [ComputeRegionNetworkFirewallPolicyRule] depends_on.
func (crnfpr *ComputeRegionNetworkFirewallPolicyRule) Dependencies() terra.Dependencies {
	return crnfpr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeRegionNetworkFirewallPolicyRule].
func (crnfpr *ComputeRegionNetworkFirewallPolicyRule) LifecycleManagement() *terra.Lifecycle {
	return crnfpr.Lifecycle
}

// Attributes returns the attributes for [ComputeRegionNetworkFirewallPolicyRule].
func (crnfpr *ComputeRegionNetworkFirewallPolicyRule) Attributes() computeRegionNetworkFirewallPolicyRuleAttributes {
	return computeRegionNetworkFirewallPolicyRuleAttributes{ref: terra.ReferenceResource(crnfpr)}
}

// ImportState imports the given attribute values into [ComputeRegionNetworkFirewallPolicyRule]'s state.
func (crnfpr *ComputeRegionNetworkFirewallPolicyRule) ImportState(av io.Reader) error {
	crnfpr.state = &computeRegionNetworkFirewallPolicyRuleState{}
	if err := json.NewDecoder(av).Decode(crnfpr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crnfpr.Type(), crnfpr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeRegionNetworkFirewallPolicyRule] has state.
func (crnfpr *ComputeRegionNetworkFirewallPolicyRule) State() (*computeRegionNetworkFirewallPolicyRuleState, bool) {
	return crnfpr.state, crnfpr.state != nil
}

// StateMust returns the state for [ComputeRegionNetworkFirewallPolicyRule]. Panics if the state is nil.
func (crnfpr *ComputeRegionNetworkFirewallPolicyRule) StateMust() *computeRegionNetworkFirewallPolicyRuleState {
	if crnfpr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crnfpr.Type(), crnfpr.LocalName()))
	}
	return crnfpr.state
}

// ComputeRegionNetworkFirewallPolicyRuleArgs contains the configurations for google_compute_region_network_firewall_policy_rule.
type ComputeRegionNetworkFirewallPolicyRuleArgs struct {
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
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// RuleName: string, optional
	RuleName terra.StringValue `hcl:"rule_name,attr"`
	// TargetServiceAccounts: list of string, optional
	TargetServiceAccounts terra.ListValue[terra.StringValue] `hcl:"target_service_accounts,attr"`
	// Match: required
	Match *computeregionnetworkfirewallpolicyrule.Match `hcl:"match,block" validate:"required"`
	// TargetSecureTags: min=0
	TargetSecureTags []computeregionnetworkfirewallpolicyrule.TargetSecureTags `hcl:"target_secure_tags,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *computeregionnetworkfirewallpolicyrule.Timeouts `hcl:"timeouts,block"`
}
type computeRegionNetworkFirewallPolicyRuleAttributes struct {
	ref terra.Reference
}

// Action returns a reference to field action of google_compute_region_network_firewall_policy_rule.
func (crnfpr computeRegionNetworkFirewallPolicyRuleAttributes) Action() terra.StringValue {
	return terra.ReferenceAsString(crnfpr.ref.Append("action"))
}

// Description returns a reference to field description of google_compute_region_network_firewall_policy_rule.
func (crnfpr computeRegionNetworkFirewallPolicyRuleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(crnfpr.ref.Append("description"))
}

// Direction returns a reference to field direction of google_compute_region_network_firewall_policy_rule.
func (crnfpr computeRegionNetworkFirewallPolicyRuleAttributes) Direction() terra.StringValue {
	return terra.ReferenceAsString(crnfpr.ref.Append("direction"))
}

// Disabled returns a reference to field disabled of google_compute_region_network_firewall_policy_rule.
func (crnfpr computeRegionNetworkFirewallPolicyRuleAttributes) Disabled() terra.BoolValue {
	return terra.ReferenceAsBool(crnfpr.ref.Append("disabled"))
}

// EnableLogging returns a reference to field enable_logging of google_compute_region_network_firewall_policy_rule.
func (crnfpr computeRegionNetworkFirewallPolicyRuleAttributes) EnableLogging() terra.BoolValue {
	return terra.ReferenceAsBool(crnfpr.ref.Append("enable_logging"))
}

// FirewallPolicy returns a reference to field firewall_policy of google_compute_region_network_firewall_policy_rule.
func (crnfpr computeRegionNetworkFirewallPolicyRuleAttributes) FirewallPolicy() terra.StringValue {
	return terra.ReferenceAsString(crnfpr.ref.Append("firewall_policy"))
}

// Id returns a reference to field id of google_compute_region_network_firewall_policy_rule.
func (crnfpr computeRegionNetworkFirewallPolicyRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crnfpr.ref.Append("id"))
}

// Kind returns a reference to field kind of google_compute_region_network_firewall_policy_rule.
func (crnfpr computeRegionNetworkFirewallPolicyRuleAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(crnfpr.ref.Append("kind"))
}

// Priority returns a reference to field priority of google_compute_region_network_firewall_policy_rule.
func (crnfpr computeRegionNetworkFirewallPolicyRuleAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(crnfpr.ref.Append("priority"))
}

// Project returns a reference to field project of google_compute_region_network_firewall_policy_rule.
func (crnfpr computeRegionNetworkFirewallPolicyRuleAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crnfpr.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_region_network_firewall_policy_rule.
func (crnfpr computeRegionNetworkFirewallPolicyRuleAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(crnfpr.ref.Append("region"))
}

// RuleName returns a reference to field rule_name of google_compute_region_network_firewall_policy_rule.
func (crnfpr computeRegionNetworkFirewallPolicyRuleAttributes) RuleName() terra.StringValue {
	return terra.ReferenceAsString(crnfpr.ref.Append("rule_name"))
}

// RuleTupleCount returns a reference to field rule_tuple_count of google_compute_region_network_firewall_policy_rule.
func (crnfpr computeRegionNetworkFirewallPolicyRuleAttributes) RuleTupleCount() terra.NumberValue {
	return terra.ReferenceAsNumber(crnfpr.ref.Append("rule_tuple_count"))
}

// TargetServiceAccounts returns a reference to field target_service_accounts of google_compute_region_network_firewall_policy_rule.
func (crnfpr computeRegionNetworkFirewallPolicyRuleAttributes) TargetServiceAccounts() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](crnfpr.ref.Append("target_service_accounts"))
}

func (crnfpr computeRegionNetworkFirewallPolicyRuleAttributes) Match() terra.ListValue[computeregionnetworkfirewallpolicyrule.MatchAttributes] {
	return terra.ReferenceAsList[computeregionnetworkfirewallpolicyrule.MatchAttributes](crnfpr.ref.Append("match"))
}

func (crnfpr computeRegionNetworkFirewallPolicyRuleAttributes) TargetSecureTags() terra.ListValue[computeregionnetworkfirewallpolicyrule.TargetSecureTagsAttributes] {
	return terra.ReferenceAsList[computeregionnetworkfirewallpolicyrule.TargetSecureTagsAttributes](crnfpr.ref.Append("target_secure_tags"))
}

func (crnfpr computeRegionNetworkFirewallPolicyRuleAttributes) Timeouts() computeregionnetworkfirewallpolicyrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeregionnetworkfirewallpolicyrule.TimeoutsAttributes](crnfpr.ref.Append("timeouts"))
}

type computeRegionNetworkFirewallPolicyRuleState struct {
	Action                string                                                         `json:"action"`
	Description           string                                                         `json:"description"`
	Direction             string                                                         `json:"direction"`
	Disabled              bool                                                           `json:"disabled"`
	EnableLogging         bool                                                           `json:"enable_logging"`
	FirewallPolicy        string                                                         `json:"firewall_policy"`
	Id                    string                                                         `json:"id"`
	Kind                  string                                                         `json:"kind"`
	Priority              float64                                                        `json:"priority"`
	Project               string                                                         `json:"project"`
	Region                string                                                         `json:"region"`
	RuleName              string                                                         `json:"rule_name"`
	RuleTupleCount        float64                                                        `json:"rule_tuple_count"`
	TargetServiceAccounts []string                                                       `json:"target_service_accounts"`
	Match                 []computeregionnetworkfirewallpolicyrule.MatchState            `json:"match"`
	TargetSecureTags      []computeregionnetworkfirewallpolicyrule.TargetSecureTagsState `json:"target_secure_tags"`
	Timeouts              *computeregionnetworkfirewallpolicyrule.TimeoutsState          `json:"timeouts"`
}
