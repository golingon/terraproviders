// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_private_dns_resolver_forwarding_rule

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurerm_private_dns_resolver_forwarding_rule.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (apdrfr *DataSource) DataSource() string {
	return "azurerm_private_dns_resolver_forwarding_rule"
}

// LocalName returns the local name for [DataSource].
func (apdrfr *DataSource) LocalName() string {
	return apdrfr.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (apdrfr *DataSource) Configuration() interface{} {
	return apdrfr.Args
}

// Attributes returns the attributes for [DataSource].
func (apdrfr *DataSource) Attributes() dataAzurermPrivateDnsResolverForwardingRuleAttributes {
	return dataAzurermPrivateDnsResolverForwardingRuleAttributes{ref: terra.ReferenceDataSource(apdrfr)}
}

// DataArgs contains the configurations for azurerm_private_dns_resolver_forwarding_rule.
type DataArgs struct {
	// DnsForwardingRulesetId: string, required
	DnsForwardingRulesetId terra.StringValue `hcl:"dns_forwarding_ruleset_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAzurermPrivateDnsResolverForwardingRuleAttributes struct {
	ref terra.Reference
}

// DnsForwardingRulesetId returns a reference to field dns_forwarding_ruleset_id of azurerm_private_dns_resolver_forwarding_rule.
func (apdrfr dataAzurermPrivateDnsResolverForwardingRuleAttributes) DnsForwardingRulesetId() terra.StringValue {
	return terra.ReferenceAsString(apdrfr.ref.Append("dns_forwarding_ruleset_id"))
}

// DomainName returns a reference to field domain_name of azurerm_private_dns_resolver_forwarding_rule.
func (apdrfr dataAzurermPrivateDnsResolverForwardingRuleAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(apdrfr.ref.Append("domain_name"))
}

// Enabled returns a reference to field enabled of azurerm_private_dns_resolver_forwarding_rule.
func (apdrfr dataAzurermPrivateDnsResolverForwardingRuleAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(apdrfr.ref.Append("enabled"))
}

// Id returns a reference to field id of azurerm_private_dns_resolver_forwarding_rule.
func (apdrfr dataAzurermPrivateDnsResolverForwardingRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(apdrfr.ref.Append("id"))
}

// Metadata returns a reference to field metadata of azurerm_private_dns_resolver_forwarding_rule.
func (apdrfr dataAzurermPrivateDnsResolverForwardingRuleAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](apdrfr.ref.Append("metadata"))
}

// Name returns a reference to field name of azurerm_private_dns_resolver_forwarding_rule.
func (apdrfr dataAzurermPrivateDnsResolverForwardingRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(apdrfr.ref.Append("name"))
}

func (apdrfr dataAzurermPrivateDnsResolverForwardingRuleAttributes) TargetDnsServers() terra.ListValue[DataTargetDnsServersAttributes] {
	return terra.ReferenceAsList[DataTargetDnsServersAttributes](apdrfr.ref.Append("target_dns_servers"))
}

func (apdrfr dataAzurermPrivateDnsResolverForwardingRuleAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](apdrfr.ref.Append("timeouts"))
}
