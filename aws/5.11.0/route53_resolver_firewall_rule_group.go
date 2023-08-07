// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRoute53ResolverFirewallRuleGroup creates a new instance of [Route53ResolverFirewallRuleGroup].
func NewRoute53ResolverFirewallRuleGroup(name string, args Route53ResolverFirewallRuleGroupArgs) *Route53ResolverFirewallRuleGroup {
	return &Route53ResolverFirewallRuleGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Route53ResolverFirewallRuleGroup)(nil)

// Route53ResolverFirewallRuleGroup represents the Terraform resource aws_route53_resolver_firewall_rule_group.
type Route53ResolverFirewallRuleGroup struct {
	Name      string
	Args      Route53ResolverFirewallRuleGroupArgs
	state     *route53ResolverFirewallRuleGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Route53ResolverFirewallRuleGroup].
func (rrfrg *Route53ResolverFirewallRuleGroup) Type() string {
	return "aws_route53_resolver_firewall_rule_group"
}

// LocalName returns the local name for [Route53ResolverFirewallRuleGroup].
func (rrfrg *Route53ResolverFirewallRuleGroup) LocalName() string {
	return rrfrg.Name
}

// Configuration returns the configuration (args) for [Route53ResolverFirewallRuleGroup].
func (rrfrg *Route53ResolverFirewallRuleGroup) Configuration() interface{} {
	return rrfrg.Args
}

// DependOn is used for other resources to depend on [Route53ResolverFirewallRuleGroup].
func (rrfrg *Route53ResolverFirewallRuleGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(rrfrg)
}

// Dependencies returns the list of resources [Route53ResolverFirewallRuleGroup] depends_on.
func (rrfrg *Route53ResolverFirewallRuleGroup) Dependencies() terra.Dependencies {
	return rrfrg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Route53ResolverFirewallRuleGroup].
func (rrfrg *Route53ResolverFirewallRuleGroup) LifecycleManagement() *terra.Lifecycle {
	return rrfrg.Lifecycle
}

// Attributes returns the attributes for [Route53ResolverFirewallRuleGroup].
func (rrfrg *Route53ResolverFirewallRuleGroup) Attributes() route53ResolverFirewallRuleGroupAttributes {
	return route53ResolverFirewallRuleGroupAttributes{ref: terra.ReferenceResource(rrfrg)}
}

// ImportState imports the given attribute values into [Route53ResolverFirewallRuleGroup]'s state.
func (rrfrg *Route53ResolverFirewallRuleGroup) ImportState(av io.Reader) error {
	rrfrg.state = &route53ResolverFirewallRuleGroupState{}
	if err := json.NewDecoder(av).Decode(rrfrg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rrfrg.Type(), rrfrg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Route53ResolverFirewallRuleGroup] has state.
func (rrfrg *Route53ResolverFirewallRuleGroup) State() (*route53ResolverFirewallRuleGroupState, bool) {
	return rrfrg.state, rrfrg.state != nil
}

// StateMust returns the state for [Route53ResolverFirewallRuleGroup]. Panics if the state is nil.
func (rrfrg *Route53ResolverFirewallRuleGroup) StateMust() *route53ResolverFirewallRuleGroupState {
	if rrfrg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rrfrg.Type(), rrfrg.LocalName()))
	}
	return rrfrg.state
}

// Route53ResolverFirewallRuleGroupArgs contains the configurations for aws_route53_resolver_firewall_rule_group.
type Route53ResolverFirewallRuleGroupArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}
type route53ResolverFirewallRuleGroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_route53_resolver_firewall_rule_group.
func (rrfrg route53ResolverFirewallRuleGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(rrfrg.ref.Append("arn"))
}

// Id returns a reference to field id of aws_route53_resolver_firewall_rule_group.
func (rrfrg route53ResolverFirewallRuleGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rrfrg.ref.Append("id"))
}

// Name returns a reference to field name of aws_route53_resolver_firewall_rule_group.
func (rrfrg route53ResolverFirewallRuleGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rrfrg.ref.Append("name"))
}

// OwnerId returns a reference to field owner_id of aws_route53_resolver_firewall_rule_group.
func (rrfrg route53ResolverFirewallRuleGroupAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(rrfrg.ref.Append("owner_id"))
}

// ShareStatus returns a reference to field share_status of aws_route53_resolver_firewall_rule_group.
func (rrfrg route53ResolverFirewallRuleGroupAttributes) ShareStatus() terra.StringValue {
	return terra.ReferenceAsString(rrfrg.ref.Append("share_status"))
}

// Tags returns a reference to field tags of aws_route53_resolver_firewall_rule_group.
func (rrfrg route53ResolverFirewallRuleGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rrfrg.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_route53_resolver_firewall_rule_group.
func (rrfrg route53ResolverFirewallRuleGroupAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rrfrg.ref.Append("tags_all"))
}

type route53ResolverFirewallRuleGroupState struct {
	Arn         string            `json:"arn"`
	Id          string            `json:"id"`
	Name        string            `json:"name"`
	OwnerId     string            `json:"owner_id"`
	ShareStatus string            `json:"share_status"`
	Tags        map[string]string `json:"tags"`
	TagsAll     map[string]string `json:"tags_all"`
}
