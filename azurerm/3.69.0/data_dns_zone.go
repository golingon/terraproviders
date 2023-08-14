// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datadnszone "github.com/golingon/terraproviders/azurerm/3.69.0/datadnszone"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataDnsZone creates a new instance of [DataDnsZone].
func NewDataDnsZone(name string, args DataDnsZoneArgs) *DataDnsZone {
	return &DataDnsZone{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDnsZone)(nil)

// DataDnsZone represents the Terraform data resource azurerm_dns_zone.
type DataDnsZone struct {
	Name string
	Args DataDnsZoneArgs
}

// DataSource returns the Terraform object type for [DataDnsZone].
func (dz *DataDnsZone) DataSource() string {
	return "azurerm_dns_zone"
}

// LocalName returns the local name for [DataDnsZone].
func (dz *DataDnsZone) LocalName() string {
	return dz.Name
}

// Configuration returns the configuration (args) for [DataDnsZone].
func (dz *DataDnsZone) Configuration() interface{} {
	return dz.Args
}

// Attributes returns the attributes for [DataDnsZone].
func (dz *DataDnsZone) Attributes() dataDnsZoneAttributes {
	return dataDnsZoneAttributes{ref: terra.ReferenceDataResource(dz)}
}

// DataDnsZoneArgs contains the configurations for azurerm_dns_zone.
type DataDnsZoneArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, optional
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr"`
	// Timeouts: optional
	Timeouts *datadnszone.Timeouts `hcl:"timeouts,block"`
}
type dataDnsZoneAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_dns_zone.
func (dz dataDnsZoneAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dz.ref.Append("id"))
}

// MaxNumberOfRecordSets returns a reference to field max_number_of_record_sets of azurerm_dns_zone.
func (dz dataDnsZoneAttributes) MaxNumberOfRecordSets() terra.NumberValue {
	return terra.ReferenceAsNumber(dz.ref.Append("max_number_of_record_sets"))
}

// Name returns a reference to field name of azurerm_dns_zone.
func (dz dataDnsZoneAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dz.ref.Append("name"))
}

// NameServers returns a reference to field name_servers of azurerm_dns_zone.
func (dz dataDnsZoneAttributes) NameServers() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dz.ref.Append("name_servers"))
}

// NumberOfRecordSets returns a reference to field number_of_record_sets of azurerm_dns_zone.
func (dz dataDnsZoneAttributes) NumberOfRecordSets() terra.NumberValue {
	return terra.ReferenceAsNumber(dz.ref.Append("number_of_record_sets"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_dns_zone.
func (dz dataDnsZoneAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dz.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_dns_zone.
func (dz dataDnsZoneAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dz.ref.Append("tags"))
}

func (dz dataDnsZoneAttributes) Timeouts() datadnszone.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datadnszone.TimeoutsAttributes](dz.ref.Append("timeouts"))
}
