// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computeregionnetworkfirewallpolicyrule "github.com/golingon/terraproviders/google/4.59.0/computeregionnetworkfirewallpolicyrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewComputeRegionNetworkFirewallPolicyRule(name string, args ComputeRegionNetworkFirewallPolicyRuleArgs) *ComputeRegionNetworkFirewallPolicyRule {
	return &ComputeRegionNetworkFirewallPolicyRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeRegionNetworkFirewallPolicyRule)(nil)

type ComputeRegionNetworkFirewallPolicyRule struct {
	Name  string
	Args  ComputeRegionNetworkFirewallPolicyRuleArgs
	state *computeRegionNetworkFirewallPolicyRuleState
}

func (crnfpr *ComputeRegionNetworkFirewallPolicyRule) Type() string {
	return "google_compute_region_network_firewall_policy_rule"
}

func (crnfpr *ComputeRegionNetworkFirewallPolicyRule) LocalName() string {
	return crnfpr.Name
}

func (crnfpr *ComputeRegionNetworkFirewallPolicyRule) Configuration() interface{} {
	return crnfpr.Args
}

func (crnfpr *ComputeRegionNetworkFirewallPolicyRule) Attributes() computeRegionNetworkFirewallPolicyRuleAttributes {
	return computeRegionNetworkFirewallPolicyRuleAttributes{ref: terra.ReferenceResource(crnfpr)}
}

func (crnfpr *ComputeRegionNetworkFirewallPolicyRule) ImportState(av io.Reader) error {
	crnfpr.state = &computeRegionNetworkFirewallPolicyRuleState{}
	if err := json.NewDecoder(av).Decode(crnfpr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crnfpr.Type(), crnfpr.LocalName(), err)
	}
	return nil
}

func (crnfpr *ComputeRegionNetworkFirewallPolicyRule) State() (*computeRegionNetworkFirewallPolicyRuleState, bool) {
	return crnfpr.state, crnfpr.state != nil
}

func (crnfpr *ComputeRegionNetworkFirewallPolicyRule) StateMust() *computeRegionNetworkFirewallPolicyRuleState {
	if crnfpr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crnfpr.Type(), crnfpr.LocalName()))
	}
	return crnfpr.state
}

func (crnfpr *ComputeRegionNetworkFirewallPolicyRule) DependOn() terra.Reference {
	return terra.ReferenceResource(crnfpr)
}

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
	// DependsOn contains resources that ComputeRegionNetworkFirewallPolicyRule depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type computeRegionNetworkFirewallPolicyRuleAttributes struct {
	ref terra.Reference
}

func (crnfpr computeRegionNetworkFirewallPolicyRuleAttributes) Action() terra.StringValue {
	return terra.ReferenceString(crnfpr.ref.Append("action"))
}

func (crnfpr computeRegionNetworkFirewallPolicyRuleAttributes) Description() terra.StringValue {
	return terra.ReferenceString(crnfpr.ref.Append("description"))
}

func (crnfpr computeRegionNetworkFirewallPolicyRuleAttributes) Direction() terra.StringValue {
	return terra.ReferenceString(crnfpr.ref.Append("direction"))
}

func (crnfpr computeRegionNetworkFirewallPolicyRuleAttributes) Disabled() terra.BoolValue {
	return terra.ReferenceBool(crnfpr.ref.Append("disabled"))
}

func (crnfpr computeRegionNetworkFirewallPolicyRuleAttributes) EnableLogging() terra.BoolValue {
	return terra.ReferenceBool(crnfpr.ref.Append("enable_logging"))
}

func (crnfpr computeRegionNetworkFirewallPolicyRuleAttributes) FirewallPolicy() terra.StringValue {
	return terra.ReferenceString(crnfpr.ref.Append("firewall_policy"))
}

func (crnfpr computeRegionNetworkFirewallPolicyRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceString(crnfpr.ref.Append("id"))
}

func (crnfpr computeRegionNetworkFirewallPolicyRuleAttributes) Kind() terra.StringValue {
	return terra.ReferenceString(crnfpr.ref.Append("kind"))
}

func (crnfpr computeRegionNetworkFirewallPolicyRuleAttributes) Priority() terra.NumberValue {
	return terra.ReferenceNumber(crnfpr.ref.Append("priority"))
}

func (crnfpr computeRegionNetworkFirewallPolicyRuleAttributes) Project() terra.StringValue {
	return terra.ReferenceString(crnfpr.ref.Append("project"))
}

func (crnfpr computeRegionNetworkFirewallPolicyRuleAttributes) Region() terra.StringValue {
	return terra.ReferenceString(crnfpr.ref.Append("region"))
}

func (crnfpr computeRegionNetworkFirewallPolicyRuleAttributes) RuleName() terra.StringValue {
	return terra.ReferenceString(crnfpr.ref.Append("rule_name"))
}

func (crnfpr computeRegionNetworkFirewallPolicyRuleAttributes) RuleTupleCount() terra.NumberValue {
	return terra.ReferenceNumber(crnfpr.ref.Append("rule_tuple_count"))
}

func (crnfpr computeRegionNetworkFirewallPolicyRuleAttributes) TargetServiceAccounts() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](crnfpr.ref.Append("target_service_accounts"))
}

func (crnfpr computeRegionNetworkFirewallPolicyRuleAttributes) Match() terra.ListValue[computeregionnetworkfirewallpolicyrule.MatchAttributes] {
	return terra.ReferenceList[computeregionnetworkfirewallpolicyrule.MatchAttributes](crnfpr.ref.Append("match"))
}

func (crnfpr computeRegionNetworkFirewallPolicyRuleAttributes) TargetSecureTags() terra.ListValue[computeregionnetworkfirewallpolicyrule.TargetSecureTagsAttributes] {
	return terra.ReferenceList[computeregionnetworkfirewallpolicyrule.TargetSecureTagsAttributes](crnfpr.ref.Append("target_secure_tags"))
}

func (crnfpr computeRegionNetworkFirewallPolicyRuleAttributes) Timeouts() computeregionnetworkfirewallpolicyrule.TimeoutsAttributes {
	return terra.ReferenceSingle[computeregionnetworkfirewallpolicyrule.TimeoutsAttributes](crnfpr.ref.Append("timeouts"))
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
