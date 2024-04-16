// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_dns_caa_record

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurerm_dns_caa_record.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (adcr *DataSource) DataSource() string {
	return "azurerm_dns_caa_record"
}

// LocalName returns the local name for [DataSource].
func (adcr *DataSource) LocalName() string {
	return adcr.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (adcr *DataSource) Configuration() interface{} {
	return adcr.Args
}

// Attributes returns the attributes for [DataSource].
func (adcr *DataSource) Attributes() dataAzurermDnsCaaRecordAttributes {
	return dataAzurermDnsCaaRecordAttributes{ref: terra.ReferenceDataSource(adcr)}
}

// DataArgs contains the configurations for azurerm_dns_caa_record.
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

type dataAzurermDnsCaaRecordAttributes struct {
	ref terra.Reference
}

// Fqdn returns a reference to field fqdn of azurerm_dns_caa_record.
func (adcr dataAzurermDnsCaaRecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(adcr.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_dns_caa_record.
func (adcr dataAzurermDnsCaaRecordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adcr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_dns_caa_record.
func (adcr dataAzurermDnsCaaRecordAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(adcr.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_dns_caa_record.
func (adcr dataAzurermDnsCaaRecordAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(adcr.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_dns_caa_record.
func (adcr dataAzurermDnsCaaRecordAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](adcr.ref.Append("tags"))
}

// Ttl returns a reference to field ttl of azurerm_dns_caa_record.
func (adcr dataAzurermDnsCaaRecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(adcr.ref.Append("ttl"))
}

// ZoneName returns a reference to field zone_name of azurerm_dns_caa_record.
func (adcr dataAzurermDnsCaaRecordAttributes) ZoneName() terra.StringValue {
	return terra.ReferenceAsString(adcr.ref.Append("zone_name"))
}

func (adcr dataAzurermDnsCaaRecordAttributes) Record() terra.SetValue[DataRecordAttributes] {
	return terra.ReferenceAsSet[DataRecordAttributes](adcr.ref.Append("record"))
}

func (adcr dataAzurermDnsCaaRecordAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](adcr.ref.Append("timeouts"))
}
