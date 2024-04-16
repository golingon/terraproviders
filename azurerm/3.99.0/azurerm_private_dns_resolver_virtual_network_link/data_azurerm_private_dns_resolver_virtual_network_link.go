// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_private_dns_resolver_virtual_network_link

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurerm_private_dns_resolver_virtual_network_link.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (apdrvnl *DataSource) DataSource() string {
	return "azurerm_private_dns_resolver_virtual_network_link"
}

// LocalName returns the local name for [DataSource].
func (apdrvnl *DataSource) LocalName() string {
	return apdrvnl.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (apdrvnl *DataSource) Configuration() interface{} {
	return apdrvnl.Args
}

// Attributes returns the attributes for [DataSource].
func (apdrvnl *DataSource) Attributes() dataAzurermPrivateDnsResolverVirtualNetworkLinkAttributes {
	return dataAzurermPrivateDnsResolverVirtualNetworkLinkAttributes{ref: terra.ReferenceDataSource(apdrvnl)}
}

// DataArgs contains the configurations for azurerm_private_dns_resolver_virtual_network_link.
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

type dataAzurermPrivateDnsResolverVirtualNetworkLinkAttributes struct {
	ref terra.Reference
}

// DnsForwardingRulesetId returns a reference to field dns_forwarding_ruleset_id of azurerm_private_dns_resolver_virtual_network_link.
func (apdrvnl dataAzurermPrivateDnsResolverVirtualNetworkLinkAttributes) DnsForwardingRulesetId() terra.StringValue {
	return terra.ReferenceAsString(apdrvnl.ref.Append("dns_forwarding_ruleset_id"))
}

// Id returns a reference to field id of azurerm_private_dns_resolver_virtual_network_link.
func (apdrvnl dataAzurermPrivateDnsResolverVirtualNetworkLinkAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(apdrvnl.ref.Append("id"))
}

// Metadata returns a reference to field metadata of azurerm_private_dns_resolver_virtual_network_link.
func (apdrvnl dataAzurermPrivateDnsResolverVirtualNetworkLinkAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](apdrvnl.ref.Append("metadata"))
}

// Name returns a reference to field name of azurerm_private_dns_resolver_virtual_network_link.
func (apdrvnl dataAzurermPrivateDnsResolverVirtualNetworkLinkAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(apdrvnl.ref.Append("name"))
}

// VirtualNetworkId returns a reference to field virtual_network_id of azurerm_private_dns_resolver_virtual_network_link.
func (apdrvnl dataAzurermPrivateDnsResolverVirtualNetworkLinkAttributes) VirtualNetworkId() terra.StringValue {
	return terra.ReferenceAsString(apdrvnl.ref.Append("virtual_network_id"))
}

func (apdrvnl dataAzurermPrivateDnsResolverVirtualNetworkLinkAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](apdrvnl.ref.Append("timeouts"))
}
