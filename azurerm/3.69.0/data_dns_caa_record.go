// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datadnscaarecord "github.com/golingon/terraproviders/azurerm/3.69.0/datadnscaarecord"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataDnsCaaRecord creates a new instance of [DataDnsCaaRecord].
func NewDataDnsCaaRecord(name string, args DataDnsCaaRecordArgs) *DataDnsCaaRecord {
	return &DataDnsCaaRecord{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDnsCaaRecord)(nil)

// DataDnsCaaRecord represents the Terraform data resource azurerm_dns_caa_record.
type DataDnsCaaRecord struct {
	Name string
	Args DataDnsCaaRecordArgs
}

// DataSource returns the Terraform object type for [DataDnsCaaRecord].
func (dcr *DataDnsCaaRecord) DataSource() string {
	return "azurerm_dns_caa_record"
}

// LocalName returns the local name for [DataDnsCaaRecord].
func (dcr *DataDnsCaaRecord) LocalName() string {
	return dcr.Name
}

// Configuration returns the configuration (args) for [DataDnsCaaRecord].
func (dcr *DataDnsCaaRecord) Configuration() interface{} {
	return dcr.Args
}

// Attributes returns the attributes for [DataDnsCaaRecord].
func (dcr *DataDnsCaaRecord) Attributes() dataDnsCaaRecordAttributes {
	return dataDnsCaaRecordAttributes{ref: terra.ReferenceDataResource(dcr)}
}

// DataDnsCaaRecordArgs contains the configurations for azurerm_dns_caa_record.
type DataDnsCaaRecordArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ZoneName: string, required
	ZoneName terra.StringValue `hcl:"zone_name,attr" validate:"required"`
	// Record: min=0
	Record []datadnscaarecord.Record `hcl:"record,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datadnscaarecord.Timeouts `hcl:"timeouts,block"`
}
type dataDnsCaaRecordAttributes struct {
	ref terra.Reference
}

// Fqdn returns a reference to field fqdn of azurerm_dns_caa_record.
func (dcr dataDnsCaaRecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(dcr.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_dns_caa_record.
func (dcr dataDnsCaaRecordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dcr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_dns_caa_record.
func (dcr dataDnsCaaRecordAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dcr.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_dns_caa_record.
func (dcr dataDnsCaaRecordAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dcr.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_dns_caa_record.
func (dcr dataDnsCaaRecordAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dcr.ref.Append("tags"))
}

// Ttl returns a reference to field ttl of azurerm_dns_caa_record.
func (dcr dataDnsCaaRecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(dcr.ref.Append("ttl"))
}

// ZoneName returns a reference to field zone_name of azurerm_dns_caa_record.
func (dcr dataDnsCaaRecordAttributes) ZoneName() terra.StringValue {
	return terra.ReferenceAsString(dcr.ref.Append("zone_name"))
}

func (dcr dataDnsCaaRecordAttributes) Record() terra.SetValue[datadnscaarecord.RecordAttributes] {
	return terra.ReferenceAsSet[datadnscaarecord.RecordAttributes](dcr.ref.Append("record"))
}

func (dcr dataDnsCaaRecordAttributes) Timeouts() datadnscaarecord.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datadnscaarecord.TimeoutsAttributes](dcr.ref.Append("timeouts"))
}
