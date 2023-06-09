// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataprivatednsresolverforwardingrule "github.com/golingon/terraproviders/azurerm/3.63.0/dataprivatednsresolverforwardingrule"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataPrivateDnsResolverForwardingRule creates a new instance of [DataPrivateDnsResolverForwardingRule].
func NewDataPrivateDnsResolverForwardingRule(name string, args DataPrivateDnsResolverForwardingRuleArgs) *DataPrivateDnsResolverForwardingRule {
	return &DataPrivateDnsResolverForwardingRule{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPrivateDnsResolverForwardingRule)(nil)

// DataPrivateDnsResolverForwardingRule represents the Terraform data resource azurerm_private_dns_resolver_forwarding_rule.
type DataPrivateDnsResolverForwardingRule struct {
	Name string
	Args DataPrivateDnsResolverForwardingRuleArgs
}

// DataSource returns the Terraform object type for [DataPrivateDnsResolverForwardingRule].
func (pdrfr *DataPrivateDnsResolverForwardingRule) DataSource() string {
	return "azurerm_private_dns_resolver_forwarding_rule"
}

// LocalName returns the local name for [DataPrivateDnsResolverForwardingRule].
func (pdrfr *DataPrivateDnsResolverForwardingRule) LocalName() string {
	return pdrfr.Name
}

// Configuration returns the configuration (args) for [DataPrivateDnsResolverForwardingRule].
func (pdrfr *DataPrivateDnsResolverForwardingRule) Configuration() interface{} {
	return pdrfr.Args
}

// Attributes returns the attributes for [DataPrivateDnsResolverForwardingRule].
func (pdrfr *DataPrivateDnsResolverForwardingRule) Attributes() dataPrivateDnsResolverForwardingRuleAttributes {
	return dataPrivateDnsResolverForwardingRuleAttributes{ref: terra.ReferenceDataResource(pdrfr)}
}

// DataPrivateDnsResolverForwardingRuleArgs contains the configurations for azurerm_private_dns_resolver_forwarding_rule.
type DataPrivateDnsResolverForwardingRuleArgs struct {
	// DnsForwardingRulesetId: string, required
	DnsForwardingRulesetId terra.StringValue `hcl:"dns_forwarding_ruleset_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// TargetDnsServers: min=0
	TargetDnsServers []dataprivatednsresolverforwardingrule.TargetDnsServers `hcl:"target_dns_servers,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataprivatednsresolverforwardingrule.Timeouts `hcl:"timeouts,block"`
}
type dataPrivateDnsResolverForwardingRuleAttributes struct {
	ref terra.Reference
}

// DnsForwardingRulesetId returns a reference to field dns_forwarding_ruleset_id of azurerm_private_dns_resolver_forwarding_rule.
func (pdrfr dataPrivateDnsResolverForwardingRuleAttributes) DnsForwardingRulesetId() terra.StringValue {
	return terra.ReferenceAsString(pdrfr.ref.Append("dns_forwarding_ruleset_id"))
}

// DomainName returns a reference to field domain_name of azurerm_private_dns_resolver_forwarding_rule.
func (pdrfr dataPrivateDnsResolverForwardingRuleAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(pdrfr.ref.Append("domain_name"))
}

// Enabled returns a reference to field enabled of azurerm_private_dns_resolver_forwarding_rule.
func (pdrfr dataPrivateDnsResolverForwardingRuleAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(pdrfr.ref.Append("enabled"))
}

// Id returns a reference to field id of azurerm_private_dns_resolver_forwarding_rule.
func (pdrfr dataPrivateDnsResolverForwardingRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pdrfr.ref.Append("id"))
}

// Metadata returns a reference to field metadata of azurerm_private_dns_resolver_forwarding_rule.
func (pdrfr dataPrivateDnsResolverForwardingRuleAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pdrfr.ref.Append("metadata"))
}

// Name returns a reference to field name of azurerm_private_dns_resolver_forwarding_rule.
func (pdrfr dataPrivateDnsResolverForwardingRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pdrfr.ref.Append("name"))
}

func (pdrfr dataPrivateDnsResolverForwardingRuleAttributes) TargetDnsServers() terra.ListValue[dataprivatednsresolverforwardingrule.TargetDnsServersAttributes] {
	return terra.ReferenceAsList[dataprivatednsresolverforwardingrule.TargetDnsServersAttributes](pdrfr.ref.Append("target_dns_servers"))
}

func (pdrfr dataPrivateDnsResolverForwardingRuleAttributes) Timeouts() dataprivatednsresolverforwardingrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataprivatednsresolverforwardingrule.TimeoutsAttributes](pdrfr.ref.Append("timeouts"))
}
