// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_dns_srv_record

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurerm_dns_srv_record.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (adsr *DataSource) DataSource() string {
	return "azurerm_dns_srv_record"
}

// LocalName returns the local name for [DataSource].
func (adsr *DataSource) LocalName() string {
	return adsr.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (adsr *DataSource) Configuration() interface{} {
	return adsr.Args
}

// Attributes returns the attributes for [DataSource].
func (adsr *DataSource) Attributes() dataAzurermDnsSrvRecordAttributes {
	return dataAzurermDnsSrvRecordAttributes{ref: terra.ReferenceDataSource(adsr)}
}

// DataArgs contains the configurations for azurerm_dns_srv_record.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ZoneName: string, required
	ZoneName terra.StringValue `hcl:"zone_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAzurermDnsSrvRecordAttributes struct {
	ref terra.Reference
}

// Fqdn returns a reference to field fqdn of azurerm_dns_srv_record.
func (adsr dataAzurermDnsSrvRecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(adsr.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_dns_srv_record.
func (adsr dataAzurermDnsSrvRecordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adsr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_dns_srv_record.
func (adsr dataAzurermDnsSrvRecordAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(adsr.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_dns_srv_record.
func (adsr dataAzurermDnsSrvRecordAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(adsr.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_dns_srv_record.
func (adsr dataAzurermDnsSrvRecordAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](adsr.ref.Append("tags"))
}

// Ttl returns a reference to field ttl of azurerm_dns_srv_record.
func (adsr dataAzurermDnsSrvRecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(adsr.ref.Append("ttl"))
}

// ZoneName returns a reference to field zone_name of azurerm_dns_srv_record.
func (adsr dataAzurermDnsSrvRecordAttributes) ZoneName() terra.StringValue {
	return terra.ReferenceAsString(adsr.ref.Append("zone_name"))
}

func (adsr dataAzurermDnsSrvRecordAttributes) Record() terra.SetValue[DataRecordAttributes] {
	return terra.ReferenceAsSet[DataRecordAttributes](adsr.ref.Append("record"))
}

func (adsr dataAzurermDnsSrvRecordAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](adsr.ref.Append("timeouts"))
}
