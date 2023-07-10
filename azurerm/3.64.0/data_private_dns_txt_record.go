// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataprivatednstxtrecord "github.com/golingon/terraproviders/azurerm/3.64.0/dataprivatednstxtrecord"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataPrivateDnsTxtRecord creates a new instance of [DataPrivateDnsTxtRecord].
func NewDataPrivateDnsTxtRecord(name string, args DataPrivateDnsTxtRecordArgs) *DataPrivateDnsTxtRecord {
	return &DataPrivateDnsTxtRecord{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPrivateDnsTxtRecord)(nil)

// DataPrivateDnsTxtRecord represents the Terraform data resource azurerm_private_dns_txt_record.
type DataPrivateDnsTxtRecord struct {
	Name string
	Args DataPrivateDnsTxtRecordArgs
}

// DataSource returns the Terraform object type for [DataPrivateDnsTxtRecord].
func (pdtr *DataPrivateDnsTxtRecord) DataSource() string {
	return "azurerm_private_dns_txt_record"
}

// LocalName returns the local name for [DataPrivateDnsTxtRecord].
func (pdtr *DataPrivateDnsTxtRecord) LocalName() string {
	return pdtr.Name
}

// Configuration returns the configuration (args) for [DataPrivateDnsTxtRecord].
func (pdtr *DataPrivateDnsTxtRecord) Configuration() interface{} {
	return pdtr.Args
}

// Attributes returns the attributes for [DataPrivateDnsTxtRecord].
func (pdtr *DataPrivateDnsTxtRecord) Attributes() dataPrivateDnsTxtRecordAttributes {
	return dataPrivateDnsTxtRecordAttributes{ref: terra.ReferenceDataResource(pdtr)}
}

// DataPrivateDnsTxtRecordArgs contains the configurations for azurerm_private_dns_txt_record.
type DataPrivateDnsTxtRecordArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ZoneName: string, required
	ZoneName terra.StringValue `hcl:"zone_name,attr" validate:"required"`
	// Record: min=0
	Record []dataprivatednstxtrecord.Record `hcl:"record,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataprivatednstxtrecord.Timeouts `hcl:"timeouts,block"`
}
type dataPrivateDnsTxtRecordAttributes struct {
	ref terra.Reference
}

// Fqdn returns a reference to field fqdn of azurerm_private_dns_txt_record.
func (pdtr dataPrivateDnsTxtRecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(pdtr.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_private_dns_txt_record.
func (pdtr dataPrivateDnsTxtRecordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pdtr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_private_dns_txt_record.
func (pdtr dataPrivateDnsTxtRecordAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pdtr.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_private_dns_txt_record.
func (pdtr dataPrivateDnsTxtRecordAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(pdtr.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_private_dns_txt_record.
func (pdtr dataPrivateDnsTxtRecordAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pdtr.ref.Append("tags"))
}

// Ttl returns a reference to field ttl of azurerm_private_dns_txt_record.
func (pdtr dataPrivateDnsTxtRecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(pdtr.ref.Append("ttl"))
}

// ZoneName returns a reference to field zone_name of azurerm_private_dns_txt_record.
func (pdtr dataPrivateDnsTxtRecordAttributes) ZoneName() terra.StringValue {
	return terra.ReferenceAsString(pdtr.ref.Append("zone_name"))
}

func (pdtr dataPrivateDnsTxtRecordAttributes) Record() terra.SetValue[dataprivatednstxtrecord.RecordAttributes] {
	return terra.ReferenceAsSet[dataprivatednstxtrecord.RecordAttributes](pdtr.ref.Append("record"))
}

func (pdtr dataPrivateDnsTxtRecordAttributes) Timeouts() dataprivatednstxtrecord.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataprivatednstxtrecord.TimeoutsAttributes](pdtr.ref.Append("timeouts"))
}
