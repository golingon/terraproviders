// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datadnstxtrecord "github.com/golingon/terraproviders/azurerm/3.66.0/datadnstxtrecord"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataDnsTxtRecord creates a new instance of [DataDnsTxtRecord].
func NewDataDnsTxtRecord(name string, args DataDnsTxtRecordArgs) *DataDnsTxtRecord {
	return &DataDnsTxtRecord{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDnsTxtRecord)(nil)

// DataDnsTxtRecord represents the Terraform data resource azurerm_dns_txt_record.
type DataDnsTxtRecord struct {
	Name string
	Args DataDnsTxtRecordArgs
}

// DataSource returns the Terraform object type for [DataDnsTxtRecord].
func (dtr *DataDnsTxtRecord) DataSource() string {
	return "azurerm_dns_txt_record"
}

// LocalName returns the local name for [DataDnsTxtRecord].
func (dtr *DataDnsTxtRecord) LocalName() string {
	return dtr.Name
}

// Configuration returns the configuration (args) for [DataDnsTxtRecord].
func (dtr *DataDnsTxtRecord) Configuration() interface{} {
	return dtr.Args
}

// Attributes returns the attributes for [DataDnsTxtRecord].
func (dtr *DataDnsTxtRecord) Attributes() dataDnsTxtRecordAttributes {
	return dataDnsTxtRecordAttributes{ref: terra.ReferenceDataResource(dtr)}
}

// DataDnsTxtRecordArgs contains the configurations for azurerm_dns_txt_record.
type DataDnsTxtRecordArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ZoneName: string, required
	ZoneName terra.StringValue `hcl:"zone_name,attr" validate:"required"`
	// Record: min=0
	Record []datadnstxtrecord.Record `hcl:"record,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datadnstxtrecord.Timeouts `hcl:"timeouts,block"`
}
type dataDnsTxtRecordAttributes struct {
	ref terra.Reference
}

// Fqdn returns a reference to field fqdn of azurerm_dns_txt_record.
func (dtr dataDnsTxtRecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(dtr.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_dns_txt_record.
func (dtr dataDnsTxtRecordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dtr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_dns_txt_record.
func (dtr dataDnsTxtRecordAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dtr.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_dns_txt_record.
func (dtr dataDnsTxtRecordAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dtr.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_dns_txt_record.
func (dtr dataDnsTxtRecordAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dtr.ref.Append("tags"))
}

// Ttl returns a reference to field ttl of azurerm_dns_txt_record.
func (dtr dataDnsTxtRecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(dtr.ref.Append("ttl"))
}

// ZoneName returns a reference to field zone_name of azurerm_dns_txt_record.
func (dtr dataDnsTxtRecordAttributes) ZoneName() terra.StringValue {
	return terra.ReferenceAsString(dtr.ref.Append("zone_name"))
}

func (dtr dataDnsTxtRecordAttributes) Record() terra.SetValue[datadnstxtrecord.RecordAttributes] {
	return terra.ReferenceAsSet[datadnstxtrecord.RecordAttributes](dtr.ref.Append("record"))
}

func (dtr dataDnsTxtRecordAttributes) Timeouts() datadnstxtrecord.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datadnstxtrecord.TimeoutsAttributes](dtr.ref.Append("timeouts"))
}
