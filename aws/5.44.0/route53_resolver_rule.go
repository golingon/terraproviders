// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	route53resolverrule "github.com/golingon/terraproviders/aws/5.44.0/route53resolverrule"
	"io"
)

// NewRoute53ResolverRule creates a new instance of [Route53ResolverRule].
func NewRoute53ResolverRule(name string, args Route53ResolverRuleArgs) *Route53ResolverRule {
	return &Route53ResolverRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Route53ResolverRule)(nil)

// Route53ResolverRule represents the Terraform resource aws_route53_resolver_rule.
type Route53ResolverRule struct {
	Name      string
	Args      Route53ResolverRuleArgs
	state     *route53ResolverRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Route53ResolverRule].
func (rrr *Route53ResolverRule) Type() string {
	return "aws_route53_resolver_rule"
}

// LocalName returns the local name for [Route53ResolverRule].
func (rrr *Route53ResolverRule) LocalName() string {
	return rrr.Name
}

// Configuration returns the configuration (args) for [Route53ResolverRule].
func (rrr *Route53ResolverRule) Configuration() interface{} {
	return rrr.Args
}

// DependOn is used for other resources to depend on [Route53ResolverRule].
func (rrr *Route53ResolverRule) DependOn() terra.Reference {
	return terra.ReferenceResource(rrr)
}

// Dependencies returns the list of resources [Route53ResolverRule] depends_on.
func (rrr *Route53ResolverRule) Dependencies() terra.Dependencies {
	return rrr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Route53ResolverRule].
func (rrr *Route53ResolverRule) LifecycleManagement() *terra.Lifecycle {
	return rrr.Lifecycle
}

// Attributes returns the attributes for [Route53ResolverRule].
func (rrr *Route53ResolverRule) Attributes() route53ResolverRuleAttributes {
	return route53ResolverRuleAttributes{ref: terra.ReferenceResource(rrr)}
}

// ImportState imports the given attribute values into [Route53ResolverRule]'s state.
func (rrr *Route53ResolverRule) ImportState(av io.Reader) error {
	rrr.state = &route53ResolverRuleState{}
	if err := json.NewDecoder(av).Decode(rrr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rrr.Type(), rrr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Route53ResolverRule] has state.
func (rrr *Route53ResolverRule) State() (*route53ResolverRuleState, bool) {
	return rrr.state, rrr.state != nil
}

// StateMust returns the state for [Route53ResolverRule]. Panics if the state is nil.
func (rrr *Route53ResolverRule) StateMust() *route53ResolverRuleState {
	if rrr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rrr.Type(), rrr.LocalName()))
	}
	return rrr.state
}

// Route53ResolverRuleArgs contains the configurations for aws_route53_resolver_rule.
type Route53ResolverRuleArgs struct {
	// DomainName: string, required
	DomainName terra.StringValue `hcl:"domain_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// ResolverEndpointId: string, optional
	ResolverEndpointId terra.StringValue `hcl:"resolver_endpoint_id,attr"`
	// RuleType: string, required
	RuleType terra.StringValue `hcl:"rule_type,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TargetIp: min=0
	TargetIp []route53resolverrule.TargetIp `hcl:"target_ip,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *route53resolverrule.Timeouts `hcl:"timeouts,block"`
}
type route53ResolverRuleAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_route53_resolver_rule.
func (rrr route53ResolverRuleAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(rrr.ref.Append("arn"))
}

// DomainName returns a reference to field domain_name of aws_route53_resolver_rule.
func (rrr route53ResolverRuleAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(rrr.ref.Append("domain_name"))
}

// Id returns a reference to field id of aws_route53_resolver_rule.
func (rrr route53ResolverRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rrr.ref.Append("id"))
}

// Name returns a reference to field name of aws_route53_resolver_rule.
func (rrr route53ResolverRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rrr.ref.Append("name"))
}

// OwnerId returns a reference to field owner_id of aws_route53_resolver_rule.
func (rrr route53ResolverRuleAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(rrr.ref.Append("owner_id"))
}

// ResolverEndpointId returns a reference to field resolver_endpoint_id of aws_route53_resolver_rule.
func (rrr route53ResolverRuleAttributes) ResolverEndpointId() terra.StringValue {
	return terra.ReferenceAsString(rrr.ref.Append("resolver_endpoint_id"))
}

// RuleType returns a reference to field rule_type of aws_route53_resolver_rule.
func (rrr route53ResolverRuleAttributes) RuleType() terra.StringValue {
	return terra.ReferenceAsString(rrr.ref.Append("rule_type"))
}

// ShareStatus returns a reference to field share_status of aws_route53_resolver_rule.
func (rrr route53ResolverRuleAttributes) ShareStatus() terra.StringValue {
	return terra.ReferenceAsString(rrr.ref.Append("share_status"))
}

// Tags returns a reference to field tags of aws_route53_resolver_rule.
func (rrr route53ResolverRuleAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rrr.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_route53_resolver_rule.
func (rrr route53ResolverRuleAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rrr.ref.Append("tags_all"))
}

func (rrr route53ResolverRuleAttributes) TargetIp() terra.SetValue[route53resolverrule.TargetIpAttributes] {
	return terra.ReferenceAsSet[route53resolverrule.TargetIpAttributes](rrr.ref.Append("target_ip"))
}

func (rrr route53ResolverRuleAttributes) Timeouts() route53resolverrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[route53resolverrule.TimeoutsAttributes](rrr.ref.Append("timeouts"))
}

type route53ResolverRuleState struct {
	Arn                string                              `json:"arn"`
	DomainName         string                              `json:"domain_name"`
	Id                 string                              `json:"id"`
	Name               string                              `json:"name"`
	OwnerId            string                              `json:"owner_id"`
	ResolverEndpointId string                              `json:"resolver_endpoint_id"`
	RuleType           string                              `json:"rule_type"`
	ShareStatus        string                              `json:"share_status"`
	Tags               map[string]string                   `json:"tags"`
	TagsAll            map[string]string                   `json:"tags_all"`
	TargetIp           []route53resolverrule.TargetIpState `json:"target_ip"`
	Timeouts           *route53resolverrule.TimeoutsState  `json:"timeouts"`
}
