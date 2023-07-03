// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataprivatednsresolvervirtualnetworklink "github.com/golingon/terraproviders/azurerm/3.63.0/dataprivatednsresolvervirtualnetworklink"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataPrivateDnsResolverVirtualNetworkLink creates a new instance of [DataPrivateDnsResolverVirtualNetworkLink].
func NewDataPrivateDnsResolverVirtualNetworkLink(name string, args DataPrivateDnsResolverVirtualNetworkLinkArgs) *DataPrivateDnsResolverVirtualNetworkLink {
	return &DataPrivateDnsResolverVirtualNetworkLink{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPrivateDnsResolverVirtualNetworkLink)(nil)

// DataPrivateDnsResolverVirtualNetworkLink represents the Terraform data resource azurerm_private_dns_resolver_virtual_network_link.
type DataPrivateDnsResolverVirtualNetworkLink struct {
	Name string
	Args DataPrivateDnsResolverVirtualNetworkLinkArgs
}

// DataSource returns the Terraform object type for [DataPrivateDnsResolverVirtualNetworkLink].
func (pdrvnl *DataPrivateDnsResolverVirtualNetworkLink) DataSource() string {
	return "azurerm_private_dns_resolver_virtual_network_link"
}

// LocalName returns the local name for [DataPrivateDnsResolverVirtualNetworkLink].
func (pdrvnl *DataPrivateDnsResolverVirtualNetworkLink) LocalName() string {
	return pdrvnl.Name
}

// Configuration returns the configuration (args) for [DataPrivateDnsResolverVirtualNetworkLink].
func (pdrvnl *DataPrivateDnsResolverVirtualNetworkLink) Configuration() interface{} {
	return pdrvnl.Args
}

// Attributes returns the attributes for [DataPrivateDnsResolverVirtualNetworkLink].
func (pdrvnl *DataPrivateDnsResolverVirtualNetworkLink) Attributes() dataPrivateDnsResolverVirtualNetworkLinkAttributes {
	return dataPrivateDnsResolverVirtualNetworkLinkAttributes{ref: terra.ReferenceDataResource(pdrvnl)}
}

// DataPrivateDnsResolverVirtualNetworkLinkArgs contains the configurations for azurerm_private_dns_resolver_virtual_network_link.
type DataPrivateDnsResolverVirtualNetworkLinkArgs struct {
	// DnsForwardingRulesetId: string, required
	DnsForwardingRulesetId terra.StringValue `hcl:"dns_forwarding_ruleset_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dataprivatednsresolvervirtualnetworklink.Timeouts `hcl:"timeouts,block"`
}
type dataPrivateDnsResolverVirtualNetworkLinkAttributes struct {
	ref terra.Reference
}

// DnsForwardingRulesetId returns a reference to field dns_forwarding_ruleset_id of azurerm_private_dns_resolver_virtual_network_link.
func (pdrvnl dataPrivateDnsResolverVirtualNetworkLinkAttributes) DnsForwardingRulesetId() terra.StringValue {
	return terra.ReferenceAsString(pdrvnl.ref.Append("dns_forwarding_ruleset_id"))
}

// Id returns a reference to field id of azurerm_private_dns_resolver_virtual_network_link.
func (pdrvnl dataPrivateDnsResolverVirtualNetworkLinkAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pdrvnl.ref.Append("id"))
}

// Metadata returns a reference to field metadata of azurerm_private_dns_resolver_virtual_network_link.
func (pdrvnl dataPrivateDnsResolverVirtualNetworkLinkAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pdrvnl.ref.Append("metadata"))
}

// Name returns a reference to field name of azurerm_private_dns_resolver_virtual_network_link.
func (pdrvnl dataPrivateDnsResolverVirtualNetworkLinkAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pdrvnl.ref.Append("name"))
}

// VirtualNetworkId returns a reference to field virtual_network_id of azurerm_private_dns_resolver_virtual_network_link.
func (pdrvnl dataPrivateDnsResolverVirtualNetworkLinkAttributes) VirtualNetworkId() terra.StringValue {
	return terra.ReferenceAsString(pdrvnl.ref.Append("virtual_network_id"))
}

func (pdrvnl dataPrivateDnsResolverVirtualNetworkLinkAttributes) Timeouts() dataprivatednsresolvervirtualnetworklink.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataprivatednsresolvervirtualnetworklink.TimeoutsAttributes](pdrvnl.ref.Append("timeouts"))
}
