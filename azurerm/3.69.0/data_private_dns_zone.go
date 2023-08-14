// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataprivatednszone "github.com/golingon/terraproviders/azurerm/3.69.0/dataprivatednszone"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataPrivateDnsZone creates a new instance of [DataPrivateDnsZone].
func NewDataPrivateDnsZone(name string, args DataPrivateDnsZoneArgs) *DataPrivateDnsZone {
	return &DataPrivateDnsZone{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPrivateDnsZone)(nil)

// DataPrivateDnsZone represents the Terraform data resource azurerm_private_dns_zone.
type DataPrivateDnsZone struct {
	Name string
	Args DataPrivateDnsZoneArgs
}

// DataSource returns the Terraform object type for [DataPrivateDnsZone].
func (pdz *DataPrivateDnsZone) DataSource() string {
	return "azurerm_private_dns_zone"
}

// LocalName returns the local name for [DataPrivateDnsZone].
func (pdz *DataPrivateDnsZone) LocalName() string {
	return pdz.Name
}

// Configuration returns the configuration (args) for [DataPrivateDnsZone].
func (pdz *DataPrivateDnsZone) Configuration() interface{} {
	return pdz.Args
}

// Attributes returns the attributes for [DataPrivateDnsZone].
func (pdz *DataPrivateDnsZone) Attributes() dataPrivateDnsZoneAttributes {
	return dataPrivateDnsZoneAttributes{ref: terra.ReferenceDataResource(pdz)}
}

// DataPrivateDnsZoneArgs contains the configurations for azurerm_private_dns_zone.
type DataPrivateDnsZoneArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, optional
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *dataprivatednszone.Timeouts `hcl:"timeouts,block"`
}
type dataPrivateDnsZoneAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_private_dns_zone.
func (pdz dataPrivateDnsZoneAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pdz.ref.Append("id"))
}

// MaxNumberOfRecordSets returns a reference to field max_number_of_record_sets of azurerm_private_dns_zone.
func (pdz dataPrivateDnsZoneAttributes) MaxNumberOfRecordSets() terra.NumberValue {
	return terra.ReferenceAsNumber(pdz.ref.Append("max_number_of_record_sets"))
}

// MaxNumberOfVirtualNetworkLinks returns a reference to field max_number_of_virtual_network_links of azurerm_private_dns_zone.
func (pdz dataPrivateDnsZoneAttributes) MaxNumberOfVirtualNetworkLinks() terra.NumberValue {
	return terra.ReferenceAsNumber(pdz.ref.Append("max_number_of_virtual_network_links"))
}

// MaxNumberOfVirtualNetworkLinksWithRegistration returns a reference to field max_number_of_virtual_network_links_with_registration of azurerm_private_dns_zone.
func (pdz dataPrivateDnsZoneAttributes) MaxNumberOfVirtualNetworkLinksWithRegistration() terra.NumberValue {
	return terra.ReferenceAsNumber(pdz.ref.Append("max_number_of_virtual_network_links_with_registration"))
}

// Name returns a reference to field name of azurerm_private_dns_zone.
func (pdz dataPrivateDnsZoneAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pdz.ref.Append("name"))
}

// NumberOfRecordSets returns a reference to field number_of_record_sets of azurerm_private_dns_zone.
func (pdz dataPrivateDnsZoneAttributes) NumberOfRecordSets() terra.NumberValue {
	return terra.ReferenceAsNumber(pdz.ref.Append("number_of_record_sets"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_private_dns_zone.
func (pdz dataPrivateDnsZoneAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(pdz.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_private_dns_zone.
func (pdz dataPrivateDnsZoneAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pdz.ref.Append("tags"))
}

func (pdz dataPrivateDnsZoneAttributes) Timeouts() dataprivatednszone.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataprivatednszone.TimeoutsAttributes](pdz.ref.Append("timeouts"))
}