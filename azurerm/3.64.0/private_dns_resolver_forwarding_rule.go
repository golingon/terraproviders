// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	privatednsresolverforwardingrule "github.com/golingon/terraproviders/azurerm/3.64.0/privatednsresolverforwardingrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPrivateDnsResolverForwardingRule creates a new instance of [PrivateDnsResolverForwardingRule].
func NewPrivateDnsResolverForwardingRule(name string, args PrivateDnsResolverForwardingRuleArgs) *PrivateDnsResolverForwardingRule {
	return &PrivateDnsResolverForwardingRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PrivateDnsResolverForwardingRule)(nil)

// PrivateDnsResolverForwardingRule represents the Terraform resource azurerm_private_dns_resolver_forwarding_rule.
type PrivateDnsResolverForwardingRule struct {
	Name      string
	Args      PrivateDnsResolverForwardingRuleArgs
	state     *privateDnsResolverForwardingRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PrivateDnsResolverForwardingRule].
func (pdrfr *PrivateDnsResolverForwardingRule) Type() string {
	return "azurerm_private_dns_resolver_forwarding_rule"
}

// LocalName returns the local name for [PrivateDnsResolverForwardingRule].
func (pdrfr *PrivateDnsResolverForwardingRule) LocalName() string {
	return pdrfr.Name
}

// Configuration returns the configuration (args) for [PrivateDnsResolverForwardingRule].
func (pdrfr *PrivateDnsResolverForwardingRule) Configuration() interface{} {
	return pdrfr.Args
}

// DependOn is used for other resources to depend on [PrivateDnsResolverForwardingRule].
func (pdrfr *PrivateDnsResolverForwardingRule) DependOn() terra.Reference {
	return terra.ReferenceResource(pdrfr)
}

// Dependencies returns the list of resources [PrivateDnsResolverForwardingRule] depends_on.
func (pdrfr *PrivateDnsResolverForwardingRule) Dependencies() terra.Dependencies {
	return pdrfr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PrivateDnsResolverForwardingRule].
func (pdrfr *PrivateDnsResolverForwardingRule) LifecycleManagement() *terra.Lifecycle {
	return pdrfr.Lifecycle
}

// Attributes returns the attributes for [PrivateDnsResolverForwardingRule].
func (pdrfr *PrivateDnsResolverForwardingRule) Attributes() privateDnsResolverForwardingRuleAttributes {
	return privateDnsResolverForwardingRuleAttributes{ref: terra.ReferenceResource(pdrfr)}
}

// ImportState imports the given attribute values into [PrivateDnsResolverForwardingRule]'s state.
func (pdrfr *PrivateDnsResolverForwardingRule) ImportState(av io.Reader) error {
	pdrfr.state = &privateDnsResolverForwardingRuleState{}
	if err := json.NewDecoder(av).Decode(pdrfr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pdrfr.Type(), pdrfr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PrivateDnsResolverForwardingRule] has state.
func (pdrfr *PrivateDnsResolverForwardingRule) State() (*privateDnsResolverForwardingRuleState, bool) {
	return pdrfr.state, pdrfr.state != nil
}

// StateMust returns the state for [PrivateDnsResolverForwardingRule]. Panics if the state is nil.
func (pdrfr *PrivateDnsResolverForwardingRule) StateMust() *privateDnsResolverForwardingRuleState {
	if pdrfr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pdrfr.Type(), pdrfr.LocalName()))
	}
	return pdrfr.state
}

// PrivateDnsResolverForwardingRuleArgs contains the configurations for azurerm_private_dns_resolver_forwarding_rule.
type PrivateDnsResolverForwardingRuleArgs struct {
	// DnsForwardingRulesetId: string, required
	DnsForwardingRulesetId terra.StringValue `hcl:"dns_forwarding_ruleset_id,attr" validate:"required"`
	// DomainName: string, required
	DomainName terra.StringValue `hcl:"domain_name,attr" validate:"required"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Metadata: map of string, optional
	Metadata terra.MapValue[terra.StringValue] `hcl:"metadata,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// TargetDnsServers: min=1
	TargetDnsServers []privatednsresolverforwardingrule.TargetDnsServers `hcl:"target_dns_servers,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *privatednsresolverforwardingrule.Timeouts `hcl:"timeouts,block"`
}
type privateDnsResolverForwardingRuleAttributes struct {
	ref terra.Reference
}

// DnsForwardingRulesetId returns a reference to field dns_forwarding_ruleset_id of azurerm_private_dns_resolver_forwarding_rule.
func (pdrfr privateDnsResolverForwardingRuleAttributes) DnsForwardingRulesetId() terra.StringValue {
	return terra.ReferenceAsString(pdrfr.ref.Append("dns_forwarding_ruleset_id"))
}

// DomainName returns a reference to field domain_name of azurerm_private_dns_resolver_forwarding_rule.
func (pdrfr privateDnsResolverForwardingRuleAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(pdrfr.ref.Append("domain_name"))
}

// Enabled returns a reference to field enabled of azurerm_private_dns_resolver_forwarding_rule.
func (pdrfr privateDnsResolverForwardingRuleAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(pdrfr.ref.Append("enabled"))
}

// Id returns a reference to field id of azurerm_private_dns_resolver_forwarding_rule.
func (pdrfr privateDnsResolverForwardingRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pdrfr.ref.Append("id"))
}

// Metadata returns a reference to field metadata of azurerm_private_dns_resolver_forwarding_rule.
func (pdrfr privateDnsResolverForwardingRuleAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pdrfr.ref.Append("metadata"))
}

// Name returns a reference to field name of azurerm_private_dns_resolver_forwarding_rule.
func (pdrfr privateDnsResolverForwardingRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pdrfr.ref.Append("name"))
}

func (pdrfr privateDnsResolverForwardingRuleAttributes) TargetDnsServers() terra.ListValue[privatednsresolverforwardingrule.TargetDnsServersAttributes] {
	return terra.ReferenceAsList[privatednsresolverforwardingrule.TargetDnsServersAttributes](pdrfr.ref.Append("target_dns_servers"))
}

func (pdrfr privateDnsResolverForwardingRuleAttributes) Timeouts() privatednsresolverforwardingrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[privatednsresolverforwardingrule.TimeoutsAttributes](pdrfr.ref.Append("timeouts"))
}

type privateDnsResolverForwardingRuleState struct {
	DnsForwardingRulesetId string                                                   `json:"dns_forwarding_ruleset_id"`
	DomainName             string                                                   `json:"domain_name"`
	Enabled                bool                                                     `json:"enabled"`
	Id                     string                                                   `json:"id"`
	Metadata               map[string]string                                        `json:"metadata"`
	Name                   string                                                   `json:"name"`
	TargetDnsServers       []privatednsresolverforwardingrule.TargetDnsServersState `json:"target_dns_servers"`
	Timeouts               *privatednsresolverforwardingrule.TimeoutsState          `json:"timeouts"`
}
