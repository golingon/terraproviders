// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataprivatednsmxrecord "github.com/golingon/terraproviders/azurerm/3.64.0/dataprivatednsmxrecord"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataPrivateDnsMxRecord creates a new instance of [DataPrivateDnsMxRecord].
func NewDataPrivateDnsMxRecord(name string, args DataPrivateDnsMxRecordArgs) *DataPrivateDnsMxRecord {
	return &DataPrivateDnsMxRecord{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPrivateDnsMxRecord)(nil)

// DataPrivateDnsMxRecord represents the Terraform data resource azurerm_private_dns_mx_record.
type DataPrivateDnsMxRecord struct {
	Name string
	Args DataPrivateDnsMxRecordArgs
}

// DataSource returns the Terraform object type for [DataPrivateDnsMxRecord].
func (pdmr *DataPrivateDnsMxRecord) DataSource() string {
	return "azurerm_private_dns_mx_record"
}

// LocalName returns the local name for [DataPrivateDnsMxRecord].
func (pdmr *DataPrivateDnsMxRecord) LocalName() string {
	return pdmr.Name
}

// Configuration returns the configuration (args) for [DataPrivateDnsMxRecord].
func (pdmr *DataPrivateDnsMxRecord) Configuration() interface{} {
	return pdmr.Args
}

// Attributes returns the attributes for [DataPrivateDnsMxRecord].
func (pdmr *DataPrivateDnsMxRecord) Attributes() dataPrivateDnsMxRecordAttributes {
	return dataPrivateDnsMxRecordAttributes{ref: terra.ReferenceDataResource(pdmr)}
}

// DataPrivateDnsMxRecordArgs contains the configurations for azurerm_private_dns_mx_record.
type DataPrivateDnsMxRecordArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ZoneName: string, required
	ZoneName terra.StringValue `hcl:"zone_name,attr" validate:"required"`
	// Record: min=0
	Record []dataprivatednsmxrecord.Record `hcl:"record,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataprivatednsmxrecord.Timeouts `hcl:"timeouts,block"`
}
type dataPrivateDnsMxRecordAttributes struct {
	ref terra.Reference
}

// Fqdn returns a reference to field fqdn of azurerm_private_dns_mx_record.
func (pdmr dataPrivateDnsMxRecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(pdmr.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_private_dns_mx_record.
func (pdmr dataPrivateDnsMxRecordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pdmr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_private_dns_mx_record.
func (pdmr dataPrivateDnsMxRecordAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pdmr.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_private_dns_mx_record.
func (pdmr dataPrivateDnsMxRecordAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(pdmr.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_private_dns_mx_record.
func (pdmr dataPrivateDnsMxRecordAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pdmr.ref.Append("tags"))
}

// Ttl returns a reference to field ttl of azurerm_private_dns_mx_record.
func (pdmr dataPrivateDnsMxRecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(pdmr.ref.Append("ttl"))
}

// ZoneName returns a reference to field zone_name of azurerm_private_dns_mx_record.
func (pdmr dataPrivateDnsMxRecordAttributes) ZoneName() terra.StringValue {
	return terra.ReferenceAsString(pdmr.ref.Append("zone_name"))
}

func (pdmr dataPrivateDnsMxRecordAttributes) Record() terra.SetValue[dataprivatednsmxrecord.RecordAttributes] {
	return terra.ReferenceAsSet[dataprivatednsmxrecord.RecordAttributes](pdmr.ref.Append("record"))
}

func (pdmr dataPrivateDnsMxRecordAttributes) Timeouts() dataprivatednsmxrecord.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataprivatednsmxrecord.TimeoutsAttributes](pdmr.ref.Append("timeouts"))
}
