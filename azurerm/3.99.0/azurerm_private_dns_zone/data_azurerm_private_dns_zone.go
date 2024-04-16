// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_private_dns_zone

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurerm_private_dns_zone.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (apdz *DataSource) DataSource() string {
	return "azurerm_private_dns_zone"
}

// LocalName returns the local name for [DataSource].
func (apdz *DataSource) LocalName() string {
	return apdz.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (apdz *DataSource) Configuration() interface{} {
	return apdz.Args
}

// Attributes returns the attributes for [DataSource].
func (apdz *DataSource) Attributes() dataAzurermPrivateDnsZoneAttributes {
	return dataAzurermPrivateDnsZoneAttributes{ref: terra.ReferenceDataSource(apdz)}
}

// DataArgs contains the configurations for azurerm_private_dns_zone.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, optional
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAzurermPrivateDnsZoneAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_private_dns_zone.
func (apdz dataAzurermPrivateDnsZoneAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(apdz.ref.Append("id"))
}

// MaxNumberOfRecordSets returns a reference to field max_number_of_record_sets of azurerm_private_dns_zone.
func (apdz dataAzurermPrivateDnsZoneAttributes) MaxNumberOfRecordSets() terra.NumberValue {
	return terra.ReferenceAsNumber(apdz.ref.Append("max_number_of_record_sets"))
}

// MaxNumberOfVirtualNetworkLinks returns a reference to field max_number_of_virtual_network_links of azurerm_private_dns_zone.
func (apdz dataAzurermPrivateDnsZoneAttributes) MaxNumberOfVirtualNetworkLinks() terra.NumberValue {
	return terra.ReferenceAsNumber(apdz.ref.Append("max_number_of_virtual_network_links"))
}

// MaxNumberOfVirtualNetworkLinksWithRegistration returns a reference to field max_number_of_virtual_network_links_with_registration of azurerm_private_dns_zone.
func (apdz dataAzurermPrivateDnsZoneAttributes) MaxNumberOfVirtualNetworkLinksWithRegistration() terra.NumberValue {
	return terra.ReferenceAsNumber(apdz.ref.Append("max_number_of_virtual_network_links_with_registration"))
}

// Name returns a reference to field name of azurerm_private_dns_zone.
func (apdz dataAzurermPrivateDnsZoneAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(apdz.ref.Append("name"))
}

// NumberOfRecordSets returns a reference to field number_of_record_sets of azurerm_private_dns_zone.
func (apdz dataAzurermPrivateDnsZoneAttributes) NumberOfRecordSets() terra.NumberValue {
	return terra.ReferenceAsNumber(apdz.ref.Append("number_of_record_sets"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_private_dns_zone.
func (apdz dataAzurermPrivateDnsZoneAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(apdz.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_private_dns_zone.
func (apdz dataAzurermPrivateDnsZoneAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](apdz.ref.Append("tags"))
}

func (apdz dataAzurermPrivateDnsZoneAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](apdz.ref.Append("timeouts"))
}
