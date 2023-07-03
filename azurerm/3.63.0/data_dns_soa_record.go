// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datadnssoarecord "github.com/golingon/terraproviders/azurerm/3.63.0/datadnssoarecord"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataDnsSoaRecord creates a new instance of [DataDnsSoaRecord].
func NewDataDnsSoaRecord(name string, args DataDnsSoaRecordArgs) *DataDnsSoaRecord {
	return &DataDnsSoaRecord{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDnsSoaRecord)(nil)

// DataDnsSoaRecord represents the Terraform data resource azurerm_dns_soa_record.
type DataDnsSoaRecord struct {
	Name string
	Args DataDnsSoaRecordArgs
}

// DataSource returns the Terraform object type for [DataDnsSoaRecord].
func (dsr *DataDnsSoaRecord) DataSource() string {
	return "azurerm_dns_soa_record"
}

// LocalName returns the local name for [DataDnsSoaRecord].
func (dsr *DataDnsSoaRecord) LocalName() string {
	return dsr.Name
}

// Configuration returns the configuration (args) for [DataDnsSoaRecord].
func (dsr *DataDnsSoaRecord) Configuration() interface{} {
	return dsr.Args
}

// Attributes returns the attributes for [DataDnsSoaRecord].
func (dsr *DataDnsSoaRecord) Attributes() dataDnsSoaRecordAttributes {
	return dataDnsSoaRecordAttributes{ref: terra.ReferenceDataResource(dsr)}
}

// DataDnsSoaRecordArgs contains the configurations for azurerm_dns_soa_record.
type DataDnsSoaRecordArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ZoneName: string, required
	ZoneName terra.StringValue `hcl:"zone_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datadnssoarecord.Timeouts `hcl:"timeouts,block"`
}
type dataDnsSoaRecordAttributes struct {
	ref terra.Reference
}

// Email returns a reference to field email of azurerm_dns_soa_record.
func (dsr dataDnsSoaRecordAttributes) Email() terra.StringValue {
	return terra.ReferenceAsString(dsr.ref.Append("email"))
}

// ExpireTime returns a reference to field expire_time of azurerm_dns_soa_record.
func (dsr dataDnsSoaRecordAttributes) ExpireTime() terra.NumberValue {
	return terra.ReferenceAsNumber(dsr.ref.Append("expire_time"))
}

// Fqdn returns a reference to field fqdn of azurerm_dns_soa_record.
func (dsr dataDnsSoaRecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(dsr.ref.Append("fqdn"))
}

// HostName returns a reference to field host_name of azurerm_dns_soa_record.
func (dsr dataDnsSoaRecordAttributes) HostName() terra.StringValue {
	return terra.ReferenceAsString(dsr.ref.Append("host_name"))
}

// Id returns a reference to field id of azurerm_dns_soa_record.
func (dsr dataDnsSoaRecordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dsr.ref.Append("id"))
}

// MinimumTtl returns a reference to field minimum_ttl of azurerm_dns_soa_record.
func (dsr dataDnsSoaRecordAttributes) MinimumTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(dsr.ref.Append("minimum_ttl"))
}

// Name returns a reference to field name of azurerm_dns_soa_record.
func (dsr dataDnsSoaRecordAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dsr.ref.Append("name"))
}

// RefreshTime returns a reference to field refresh_time of azurerm_dns_soa_record.
func (dsr dataDnsSoaRecordAttributes) RefreshTime() terra.NumberValue {
	return terra.ReferenceAsNumber(dsr.ref.Append("refresh_time"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_dns_soa_record.
func (dsr dataDnsSoaRecordAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dsr.ref.Append("resource_group_name"))
}

// RetryTime returns a reference to field retry_time of azurerm_dns_soa_record.
func (dsr dataDnsSoaRecordAttributes) RetryTime() terra.NumberValue {
	return terra.ReferenceAsNumber(dsr.ref.Append("retry_time"))
}

// SerialNumber returns a reference to field serial_number of azurerm_dns_soa_record.
func (dsr dataDnsSoaRecordAttributes) SerialNumber() terra.NumberValue {
	return terra.ReferenceAsNumber(dsr.ref.Append("serial_number"))
}

// Tags returns a reference to field tags of azurerm_dns_soa_record.
func (dsr dataDnsSoaRecordAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dsr.ref.Append("tags"))
}

// Ttl returns a reference to field ttl of azurerm_dns_soa_record.
func (dsr dataDnsSoaRecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(dsr.ref.Append("ttl"))
}

// ZoneName returns a reference to field zone_name of azurerm_dns_soa_record.
func (dsr dataDnsSoaRecordAttributes) ZoneName() terra.StringValue {
	return terra.ReferenceAsString(dsr.ref.Append("zone_name"))
}

func (dsr dataDnsSoaRecordAttributes) Timeouts() datadnssoarecord.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datadnssoarecord.TimeoutsAttributes](dsr.ref.Append("timeouts"))
}
