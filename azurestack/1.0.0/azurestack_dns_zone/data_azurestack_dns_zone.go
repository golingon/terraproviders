// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurestack_dns_zone

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurestack_dns_zone.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (adz *DataSource) DataSource() string {
	return "azurestack_dns_zone"
}

// LocalName returns the local name for [DataSource].
func (adz *DataSource) LocalName() string {
	return adz.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (adz *DataSource) Configuration() interface{} {
	return adz.Args
}

// Attributes returns the attributes for [DataSource].
func (adz *DataSource) Attributes() dataAzurestackDnsZoneAttributes {
	return dataAzurestackDnsZoneAttributes{ref: terra.ReferenceDataSource(adz)}
}

// DataArgs contains the configurations for azurestack_dns_zone.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, optional
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr"`
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAzurestackDnsZoneAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurestack_dns_zone.
func (adz dataAzurestackDnsZoneAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adz.ref.Append("id"))
}

// MaxNumberOfRecordSets returns a reference to field max_number_of_record_sets of azurestack_dns_zone.
func (adz dataAzurestackDnsZoneAttributes) MaxNumberOfRecordSets() terra.NumberValue {
	return terra.ReferenceAsNumber(adz.ref.Append("max_number_of_record_sets"))
}

// Name returns a reference to field name of azurestack_dns_zone.
func (adz dataAzurestackDnsZoneAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(adz.ref.Append("name"))
}

// NameServers returns a reference to field name_servers of azurestack_dns_zone.
func (adz dataAzurestackDnsZoneAttributes) NameServers() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](adz.ref.Append("name_servers"))
}

// NumberOfRecordSets returns a reference to field number_of_record_sets of azurestack_dns_zone.
func (adz dataAzurestackDnsZoneAttributes) NumberOfRecordSets() terra.NumberValue {
	return terra.ReferenceAsNumber(adz.ref.Append("number_of_record_sets"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurestack_dns_zone.
func (adz dataAzurestackDnsZoneAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(adz.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurestack_dns_zone.
func (adz dataAzurestackDnsZoneAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](adz.ref.Append("tags"))
}

func (adz dataAzurestackDnsZoneAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](adz.ref.Append("timeouts"))
}
