// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datadnsmxrecord "github.com/golingon/terraproviders/azurerm/3.63.0/datadnsmxrecord"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataDnsMxRecord creates a new instance of [DataDnsMxRecord].
func NewDataDnsMxRecord(name string, args DataDnsMxRecordArgs) *DataDnsMxRecord {
	return &DataDnsMxRecord{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDnsMxRecord)(nil)

// DataDnsMxRecord represents the Terraform data resource azurerm_dns_mx_record.
type DataDnsMxRecord struct {
	Name string
	Args DataDnsMxRecordArgs
}

// DataSource returns the Terraform object type for [DataDnsMxRecord].
func (dmr *DataDnsMxRecord) DataSource() string {
	return "azurerm_dns_mx_record"
}

// LocalName returns the local name for [DataDnsMxRecord].
func (dmr *DataDnsMxRecord) LocalName() string {
	return dmr.Name
}

// Configuration returns the configuration (args) for [DataDnsMxRecord].
func (dmr *DataDnsMxRecord) Configuration() interface{} {
	return dmr.Args
}

// Attributes returns the attributes for [DataDnsMxRecord].
func (dmr *DataDnsMxRecord) Attributes() dataDnsMxRecordAttributes {
	return dataDnsMxRecordAttributes{ref: terra.ReferenceDataResource(dmr)}
}

// DataDnsMxRecordArgs contains the configurations for azurerm_dns_mx_record.
type DataDnsMxRecordArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ZoneName: string, required
	ZoneName terra.StringValue `hcl:"zone_name,attr" validate:"required"`
	// Record: min=0
	Record []datadnsmxrecord.Record `hcl:"record,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datadnsmxrecord.Timeouts `hcl:"timeouts,block"`
}
type dataDnsMxRecordAttributes struct {
	ref terra.Reference
}

// Fqdn returns a reference to field fqdn of azurerm_dns_mx_record.
func (dmr dataDnsMxRecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(dmr.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_dns_mx_record.
func (dmr dataDnsMxRecordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dmr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_dns_mx_record.
func (dmr dataDnsMxRecordAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dmr.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_dns_mx_record.
func (dmr dataDnsMxRecordAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dmr.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_dns_mx_record.
func (dmr dataDnsMxRecordAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dmr.ref.Append("tags"))
}

// Ttl returns a reference to field ttl of azurerm_dns_mx_record.
func (dmr dataDnsMxRecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(dmr.ref.Append("ttl"))
}

// ZoneName returns a reference to field zone_name of azurerm_dns_mx_record.
func (dmr dataDnsMxRecordAttributes) ZoneName() terra.StringValue {
	return terra.ReferenceAsString(dmr.ref.Append("zone_name"))
}

func (dmr dataDnsMxRecordAttributes) Record() terra.SetValue[datadnsmxrecord.RecordAttributes] {
	return terra.ReferenceAsSet[datadnsmxrecord.RecordAttributes](dmr.ref.Append("record"))
}

func (dmr dataDnsMxRecordAttributes) Timeouts() datadnsmxrecord.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datadnsmxrecord.TimeoutsAttributes](dmr.ref.Append("timeouts"))
}
