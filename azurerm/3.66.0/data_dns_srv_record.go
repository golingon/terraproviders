// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datadnssrvrecord "github.com/golingon/terraproviders/azurerm/3.66.0/datadnssrvrecord"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataDnsSrvRecord creates a new instance of [DataDnsSrvRecord].
func NewDataDnsSrvRecord(name string, args DataDnsSrvRecordArgs) *DataDnsSrvRecord {
	return &DataDnsSrvRecord{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDnsSrvRecord)(nil)

// DataDnsSrvRecord represents the Terraform data resource azurerm_dns_srv_record.
type DataDnsSrvRecord struct {
	Name string
	Args DataDnsSrvRecordArgs
}

// DataSource returns the Terraform object type for [DataDnsSrvRecord].
func (dsr *DataDnsSrvRecord) DataSource() string {
	return "azurerm_dns_srv_record"
}

// LocalName returns the local name for [DataDnsSrvRecord].
func (dsr *DataDnsSrvRecord) LocalName() string {
	return dsr.Name
}

// Configuration returns the configuration (args) for [DataDnsSrvRecord].
func (dsr *DataDnsSrvRecord) Configuration() interface{} {
	return dsr.Args
}

// Attributes returns the attributes for [DataDnsSrvRecord].
func (dsr *DataDnsSrvRecord) Attributes() dataDnsSrvRecordAttributes {
	return dataDnsSrvRecordAttributes{ref: terra.ReferenceDataResource(dsr)}
}

// DataDnsSrvRecordArgs contains the configurations for azurerm_dns_srv_record.
type DataDnsSrvRecordArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ZoneName: string, required
	ZoneName terra.StringValue `hcl:"zone_name,attr" validate:"required"`
	// Record: min=0
	Record []datadnssrvrecord.Record `hcl:"record,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datadnssrvrecord.Timeouts `hcl:"timeouts,block"`
}
type dataDnsSrvRecordAttributes struct {
	ref terra.Reference
}

// Fqdn returns a reference to field fqdn of azurerm_dns_srv_record.
func (dsr dataDnsSrvRecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(dsr.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_dns_srv_record.
func (dsr dataDnsSrvRecordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dsr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_dns_srv_record.
func (dsr dataDnsSrvRecordAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dsr.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_dns_srv_record.
func (dsr dataDnsSrvRecordAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dsr.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_dns_srv_record.
func (dsr dataDnsSrvRecordAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dsr.ref.Append("tags"))
}

// Ttl returns a reference to field ttl of azurerm_dns_srv_record.
func (dsr dataDnsSrvRecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(dsr.ref.Append("ttl"))
}

// ZoneName returns a reference to field zone_name of azurerm_dns_srv_record.
func (dsr dataDnsSrvRecordAttributes) ZoneName() terra.StringValue {
	return terra.ReferenceAsString(dsr.ref.Append("zone_name"))
}

func (dsr dataDnsSrvRecordAttributes) Record() terra.SetValue[datadnssrvrecord.RecordAttributes] {
	return terra.ReferenceAsSet[datadnssrvrecord.RecordAttributes](dsr.ref.Append("record"))
}

func (dsr dataDnsSrvRecordAttributes) Timeouts() datadnssrvrecord.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datadnssrvrecord.TimeoutsAttributes](dsr.ref.Append("timeouts"))
}
