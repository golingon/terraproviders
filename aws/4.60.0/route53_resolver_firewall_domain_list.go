// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewRoute53ResolverFirewallDomainList(name string, args Route53ResolverFirewallDomainListArgs) *Route53ResolverFirewallDomainList {
	return &Route53ResolverFirewallDomainList{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Route53ResolverFirewallDomainList)(nil)

type Route53ResolverFirewallDomainList struct {
	Name  string
	Args  Route53ResolverFirewallDomainListArgs
	state *route53ResolverFirewallDomainListState
}

func (rrfdl *Route53ResolverFirewallDomainList) Type() string {
	return "aws_route53_resolver_firewall_domain_list"
}

func (rrfdl *Route53ResolverFirewallDomainList) LocalName() string {
	return rrfdl.Name
}

func (rrfdl *Route53ResolverFirewallDomainList) Configuration() interface{} {
	return rrfdl.Args
}

func (rrfdl *Route53ResolverFirewallDomainList) Attributes() route53ResolverFirewallDomainListAttributes {
	return route53ResolverFirewallDomainListAttributes{ref: terra.ReferenceResource(rrfdl)}
}

func (rrfdl *Route53ResolverFirewallDomainList) ImportState(av io.Reader) error {
	rrfdl.state = &route53ResolverFirewallDomainListState{}
	if err := json.NewDecoder(av).Decode(rrfdl.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rrfdl.Type(), rrfdl.LocalName(), err)
	}
	return nil
}

func (rrfdl *Route53ResolverFirewallDomainList) State() (*route53ResolverFirewallDomainListState, bool) {
	return rrfdl.state, rrfdl.state != nil
}

func (rrfdl *Route53ResolverFirewallDomainList) StateMust() *route53ResolverFirewallDomainListState {
	if rrfdl.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rrfdl.Type(), rrfdl.LocalName()))
	}
	return rrfdl.state
}

func (rrfdl *Route53ResolverFirewallDomainList) DependOn() terra.Reference {
	return terra.ReferenceResource(rrfdl)
}

type Route53ResolverFirewallDomainListArgs struct {
	// Domains: set of string, optional
	Domains terra.SetValue[terra.StringValue] `hcl:"domains,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// DependsOn contains resources that Route53ResolverFirewallDomainList depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type route53ResolverFirewallDomainListAttributes struct {
	ref terra.Reference
}

func (rrfdl route53ResolverFirewallDomainListAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(rrfdl.ref.Append("arn"))
}

func (rrfdl route53ResolverFirewallDomainListAttributes) Domains() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](rrfdl.ref.Append("domains"))
}

func (rrfdl route53ResolverFirewallDomainListAttributes) Id() terra.StringValue {
	return terra.ReferenceString(rrfdl.ref.Append("id"))
}

func (rrfdl route53ResolverFirewallDomainListAttributes) Name() terra.StringValue {
	return terra.ReferenceString(rrfdl.ref.Append("name"))
}

func (rrfdl route53ResolverFirewallDomainListAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](rrfdl.ref.Append("tags"))
}

func (rrfdl route53ResolverFirewallDomainListAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](rrfdl.ref.Append("tags_all"))
}

type route53ResolverFirewallDomainListState struct {
	Arn     string            `json:"arn"`
	Domains []string          `json:"domains"`
	Id      string            `json:"id"`
	Name    string            `json:"name"`
	Tags    map[string]string `json:"tags"`
	TagsAll map[string]string `json:"tags_all"`
}
