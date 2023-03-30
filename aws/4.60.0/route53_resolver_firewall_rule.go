// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewRoute53ResolverFirewallRule(name string, args Route53ResolverFirewallRuleArgs) *Route53ResolverFirewallRule {
	return &Route53ResolverFirewallRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Route53ResolverFirewallRule)(nil)

type Route53ResolverFirewallRule struct {
	Name  string
	Args  Route53ResolverFirewallRuleArgs
	state *route53ResolverFirewallRuleState
}

func (rrfr *Route53ResolverFirewallRule) Type() string {
	return "aws_route53_resolver_firewall_rule"
}

func (rrfr *Route53ResolverFirewallRule) LocalName() string {
	return rrfr.Name
}

func (rrfr *Route53ResolverFirewallRule) Configuration() interface{} {
	return rrfr.Args
}

func (rrfr *Route53ResolverFirewallRule) Attributes() route53ResolverFirewallRuleAttributes {
	return route53ResolverFirewallRuleAttributes{ref: terra.ReferenceResource(rrfr)}
}

func (rrfr *Route53ResolverFirewallRule) ImportState(av io.Reader) error {
	rrfr.state = &route53ResolverFirewallRuleState{}
	if err := json.NewDecoder(av).Decode(rrfr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rrfr.Type(), rrfr.LocalName(), err)
	}
	return nil
}

func (rrfr *Route53ResolverFirewallRule) State() (*route53ResolverFirewallRuleState, bool) {
	return rrfr.state, rrfr.state != nil
}

func (rrfr *Route53ResolverFirewallRule) StateMust() *route53ResolverFirewallRuleState {
	if rrfr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rrfr.Type(), rrfr.LocalName()))
	}
	return rrfr.state
}

func (rrfr *Route53ResolverFirewallRule) DependOn() terra.Reference {
	return terra.ReferenceResource(rrfr)
}

type Route53ResolverFirewallRuleArgs struct {
	// Action: string, required
	Action terra.StringValue `hcl:"action,attr" validate:"required"`
	// BlockOverrideDnsType: string, optional
	BlockOverrideDnsType terra.StringValue `hcl:"block_override_dns_type,attr"`
	// BlockOverrideDomain: string, optional
	BlockOverrideDomain terra.StringValue `hcl:"block_override_domain,attr"`
	// BlockOverrideTtl: number, optional
	BlockOverrideTtl terra.NumberValue `hcl:"block_override_ttl,attr"`
	// BlockResponse: string, optional
	BlockResponse terra.StringValue `hcl:"block_response,attr"`
	// FirewallDomainListId: string, required
	FirewallDomainListId terra.StringValue `hcl:"firewall_domain_list_id,attr" validate:"required"`
	// FirewallRuleGroupId: string, required
	FirewallRuleGroupId terra.StringValue `hcl:"firewall_rule_group_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Priority: number, required
	Priority terra.NumberValue `hcl:"priority,attr" validate:"required"`
	// DependsOn contains resources that Route53ResolverFirewallRule depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type route53ResolverFirewallRuleAttributes struct {
	ref terra.Reference
}

func (rrfr route53ResolverFirewallRuleAttributes) Action() terra.StringValue {
	return terra.ReferenceString(rrfr.ref.Append("action"))
}

func (rrfr route53ResolverFirewallRuleAttributes) BlockOverrideDnsType() terra.StringValue {
	return terra.ReferenceString(rrfr.ref.Append("block_override_dns_type"))
}

func (rrfr route53ResolverFirewallRuleAttributes) BlockOverrideDomain() terra.StringValue {
	return terra.ReferenceString(rrfr.ref.Append("block_override_domain"))
}

func (rrfr route53ResolverFirewallRuleAttributes) BlockOverrideTtl() terra.NumberValue {
	return terra.ReferenceNumber(rrfr.ref.Append("block_override_ttl"))
}

func (rrfr route53ResolverFirewallRuleAttributes) BlockResponse() terra.StringValue {
	return terra.ReferenceString(rrfr.ref.Append("block_response"))
}

func (rrfr route53ResolverFirewallRuleAttributes) FirewallDomainListId() terra.StringValue {
	return terra.ReferenceString(rrfr.ref.Append("firewall_domain_list_id"))
}

func (rrfr route53ResolverFirewallRuleAttributes) FirewallRuleGroupId() terra.StringValue {
	return terra.ReferenceString(rrfr.ref.Append("firewall_rule_group_id"))
}

func (rrfr route53ResolverFirewallRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceString(rrfr.ref.Append("id"))
}

func (rrfr route53ResolverFirewallRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceString(rrfr.ref.Append("name"))
}

func (rrfr route53ResolverFirewallRuleAttributes) Priority() terra.NumberValue {
	return terra.ReferenceNumber(rrfr.ref.Append("priority"))
}

type route53ResolverFirewallRuleState struct {
	Action               string  `json:"action"`
	BlockOverrideDnsType string  `json:"block_override_dns_type"`
	BlockOverrideDomain  string  `json:"block_override_domain"`
	BlockOverrideTtl     float64 `json:"block_override_ttl"`
	BlockResponse        string  `json:"block_response"`
	FirewallDomainListId string  `json:"firewall_domain_list_id"`
	FirewallRuleGroupId  string  `json:"firewall_rule_group_id"`
	Id                   string  `json:"id"`
	Name                 string  `json:"name"`
	Priority             float64 `json:"priority"`
}
