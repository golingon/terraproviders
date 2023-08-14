// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataprivatednssoarecord "github.com/golingon/terraproviders/azurerm/3.68.0/dataprivatednssoarecord"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataPrivateDnsSoaRecord creates a new instance of [DataPrivateDnsSoaRecord].
func NewDataPrivateDnsSoaRecord(name string, args DataPrivateDnsSoaRecordArgs) *DataPrivateDnsSoaRecord {
	return &DataPrivateDnsSoaRecord{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPrivateDnsSoaRecord)(nil)

// DataPrivateDnsSoaRecord represents the Terraform data resource azurerm_private_dns_soa_record.
type DataPrivateDnsSoaRecord struct {
	Name string
	Args DataPrivateDnsSoaRecordArgs
}

// DataSource returns the Terraform object type for [DataPrivateDnsSoaRecord].
func (pdsr *DataPrivateDnsSoaRecord) DataSource() string {
	return "azurerm_private_dns_soa_record"
}

// LocalName returns the local name for [DataPrivateDnsSoaRecord].
func (pdsr *DataPrivateDnsSoaRecord) LocalName() string {
	return pdsr.Name
}

// Configuration returns the configuration (args) for [DataPrivateDnsSoaRecord].
func (pdsr *DataPrivateDnsSoaRecord) Configuration() interface{} {
	return pdsr.Args
}

// Attributes returns the attributes for [DataPrivateDnsSoaRecord].
func (pdsr *DataPrivateDnsSoaRecord) Attributes() dataPrivateDnsSoaRecordAttributes {
	return dataPrivateDnsSoaRecordAttributes{ref: terra.ReferenceDataResource(pdsr)}
}

// DataPrivateDnsSoaRecordArgs contains the configurations for azurerm_private_dns_soa_record.
type DataPrivateDnsSoaRecordArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ZoneName: string, required
	ZoneName terra.StringValue `hcl:"zone_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dataprivatednssoarecord.Timeouts `hcl:"timeouts,block"`
}
type dataPrivateDnsSoaRecordAttributes struct {
	ref terra.Reference
}

// Email returns a reference to field email of azurerm_private_dns_soa_record.
func (pdsr dataPrivateDnsSoaRecordAttributes) Email() terra.StringValue {
	return terra.ReferenceAsString(pdsr.ref.Append("email"))
}

// ExpireTime returns a reference to field expire_time of azurerm_private_dns_soa_record.
func (pdsr dataPrivateDnsSoaRecordAttributes) ExpireTime() terra.NumberValue {
	return terra.ReferenceAsNumber(pdsr.ref.Append("expire_time"))
}

// Fqdn returns a reference to field fqdn of azurerm_private_dns_soa_record.
func (pdsr dataPrivateDnsSoaRecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(pdsr.ref.Append("fqdn"))
}

// HostName returns a reference to field host_name of azurerm_private_dns_soa_record.
func (pdsr dataPrivateDnsSoaRecordAttributes) HostName() terra.StringValue {
	return terra.ReferenceAsString(pdsr.ref.Append("host_name"))
}

// Id returns a reference to field id of azurerm_private_dns_soa_record.
func (pdsr dataPrivateDnsSoaRecordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pdsr.ref.Append("id"))
}

// MinimumTtl returns a reference to field minimum_ttl of azurerm_private_dns_soa_record.
func (pdsr dataPrivateDnsSoaRecordAttributes) MinimumTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(pdsr.ref.Append("minimum_ttl"))
}

// Name returns a reference to field name of azurerm_private_dns_soa_record.
func (pdsr dataPrivateDnsSoaRecordAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pdsr.ref.Append("name"))
}

// RefreshTime returns a reference to field refresh_time of azurerm_private_dns_soa_record.
func (pdsr dataPrivateDnsSoaRecordAttributes) RefreshTime() terra.NumberValue {
	return terra.ReferenceAsNumber(pdsr.ref.Append("refresh_time"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_private_dns_soa_record.
func (pdsr dataPrivateDnsSoaRecordAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(pdsr.ref.Append("resource_group_name"))
}

// RetryTime returns a reference to field retry_time of azurerm_private_dns_soa_record.
func (pdsr dataPrivateDnsSoaRecordAttributes) RetryTime() terra.NumberValue {
	return terra.ReferenceAsNumber(pdsr.ref.Append("retry_time"))
}

// SerialNumber returns a reference to field serial_number of azurerm_private_dns_soa_record.
func (pdsr dataPrivateDnsSoaRecordAttributes) SerialNumber() terra.NumberValue {
	return terra.ReferenceAsNumber(pdsr.ref.Append("serial_number"))
}

// Tags returns a reference to field tags of azurerm_private_dns_soa_record.
func (pdsr dataPrivateDnsSoaRecordAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pdsr.ref.Append("tags"))
}

// Ttl returns a reference to field ttl of azurerm_private_dns_soa_record.
func (pdsr dataPrivateDnsSoaRecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(pdsr.ref.Append("ttl"))
}

// ZoneName returns a reference to field zone_name of azurerm_private_dns_soa_record.
func (pdsr dataPrivateDnsSoaRecordAttributes) ZoneName() terra.StringValue {
	return terra.ReferenceAsString(pdsr.ref.Append("zone_name"))
}

func (pdsr dataPrivateDnsSoaRecordAttributes) Timeouts() dataprivatednssoarecord.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataprivatednssoarecord.TimeoutsAttributes](pdsr.ref.Append("timeouts"))
}