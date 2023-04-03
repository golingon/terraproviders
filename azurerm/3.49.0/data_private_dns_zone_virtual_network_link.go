// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataprivatednszonevirtualnetworklink "github.com/golingon/terraproviders/azurerm/3.49.0/dataprivatednszonevirtualnetworklink"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataPrivateDnsZoneVirtualNetworkLink creates a new instance of [DataPrivateDnsZoneVirtualNetworkLink].
func NewDataPrivateDnsZoneVirtualNetworkLink(name string, args DataPrivateDnsZoneVirtualNetworkLinkArgs) *DataPrivateDnsZoneVirtualNetworkLink {
	return &DataPrivateDnsZoneVirtualNetworkLink{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPrivateDnsZoneVirtualNetworkLink)(nil)

// DataPrivateDnsZoneVirtualNetworkLink represents the Terraform data resource azurerm_private_dns_zone_virtual_network_link.
type DataPrivateDnsZoneVirtualNetworkLink struct {
	Name string
	Args DataPrivateDnsZoneVirtualNetworkLinkArgs
}

// DataSource returns the Terraform object type for [DataPrivateDnsZoneVirtualNetworkLink].
func (pdzvnl *DataPrivateDnsZoneVirtualNetworkLink) DataSource() string {
	return "azurerm_private_dns_zone_virtual_network_link"
}

// LocalName returns the local name for [DataPrivateDnsZoneVirtualNetworkLink].
func (pdzvnl *DataPrivateDnsZoneVirtualNetworkLink) LocalName() string {
	return pdzvnl.Name
}

// Configuration returns the configuration (args) for [DataPrivateDnsZoneVirtualNetworkLink].
func (pdzvnl *DataPrivateDnsZoneVirtualNetworkLink) Configuration() interface{} {
	return pdzvnl.Args
}

// Attributes returns the attributes for [DataPrivateDnsZoneVirtualNetworkLink].
func (pdzvnl *DataPrivateDnsZoneVirtualNetworkLink) Attributes() dataPrivateDnsZoneVirtualNetworkLinkAttributes {
	return dataPrivateDnsZoneVirtualNetworkLinkAttributes{ref: terra.ReferenceDataResource(pdzvnl)}
}

// DataPrivateDnsZoneVirtualNetworkLinkArgs contains the configurations for azurerm_private_dns_zone_virtual_network_link.
type DataPrivateDnsZoneVirtualNetworkLinkArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PrivateDnsZoneName: string, required
	PrivateDnsZoneName terra.StringValue `hcl:"private_dns_zone_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dataprivatednszonevirtualnetworklink.Timeouts `hcl:"timeouts,block"`
}
type dataPrivateDnsZoneVirtualNetworkLinkAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_private_dns_zone_virtual_network_link.
func (pdzvnl dataPrivateDnsZoneVirtualNetworkLinkAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pdzvnl.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_private_dns_zone_virtual_network_link.
func (pdzvnl dataPrivateDnsZoneVirtualNetworkLinkAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pdzvnl.ref.Append("name"))
}

// PrivateDnsZoneName returns a reference to field private_dns_zone_name of azurerm_private_dns_zone_virtual_network_link.
func (pdzvnl dataPrivateDnsZoneVirtualNetworkLinkAttributes) PrivateDnsZoneName() terra.StringValue {
	return terra.ReferenceAsString(pdzvnl.ref.Append("private_dns_zone_name"))
}

// RegistrationEnabled returns a reference to field registration_enabled of azurerm_private_dns_zone_virtual_network_link.
func (pdzvnl dataPrivateDnsZoneVirtualNetworkLinkAttributes) RegistrationEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(pdzvnl.ref.Append("registration_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_private_dns_zone_virtual_network_link.
func (pdzvnl dataPrivateDnsZoneVirtualNetworkLinkAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(pdzvnl.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_private_dns_zone_virtual_network_link.
func (pdzvnl dataPrivateDnsZoneVirtualNetworkLinkAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pdzvnl.ref.Append("tags"))
}

// VirtualNetworkId returns a reference to field virtual_network_id of azurerm_private_dns_zone_virtual_network_link.
func (pdzvnl dataPrivateDnsZoneVirtualNetworkLinkAttributes) VirtualNetworkId() terra.StringValue {
	return terra.ReferenceAsString(pdzvnl.ref.Append("virtual_network_id"))
}

func (pdzvnl dataPrivateDnsZoneVirtualNetworkLinkAttributes) Timeouts() dataprivatednszonevirtualnetworklink.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataprivatednszonevirtualnetworklink.TimeoutsAttributes](pdzvnl.ref.Append("timeouts"))
}
