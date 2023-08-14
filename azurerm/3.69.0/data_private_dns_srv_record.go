// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataprivatednssrvrecord "github.com/golingon/terraproviders/azurerm/3.69.0/dataprivatednssrvrecord"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataPrivateDnsSrvRecord creates a new instance of [DataPrivateDnsSrvRecord].
func NewDataPrivateDnsSrvRecord(name string, args DataPrivateDnsSrvRecordArgs) *DataPrivateDnsSrvRecord {
	return &DataPrivateDnsSrvRecord{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPrivateDnsSrvRecord)(nil)

// DataPrivateDnsSrvRecord represents the Terraform data resource azurerm_private_dns_srv_record.
type DataPrivateDnsSrvRecord struct {
	Name string
	Args DataPrivateDnsSrvRecordArgs
}

// DataSource returns the Terraform object type for [DataPrivateDnsSrvRecord].
func (pdsr *DataPrivateDnsSrvRecord) DataSource() string {
	return "azurerm_private_dns_srv_record"
}

// LocalName returns the local name for [DataPrivateDnsSrvRecord].
func (pdsr *DataPrivateDnsSrvRecord) LocalName() string {
	return pdsr.Name
}

// Configuration returns the configuration (args) for [DataPrivateDnsSrvRecord].
func (pdsr *DataPrivateDnsSrvRecord) Configuration() interface{} {
	return pdsr.Args
}

// Attributes returns the attributes for [DataPrivateDnsSrvRecord].
func (pdsr *DataPrivateDnsSrvRecord) Attributes() dataPrivateDnsSrvRecordAttributes {
	return dataPrivateDnsSrvRecordAttributes{ref: terra.ReferenceDataResource(pdsr)}
}

// DataPrivateDnsSrvRecordArgs contains the configurations for azurerm_private_dns_srv_record.
type DataPrivateDnsSrvRecordArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ZoneName: string, required
	ZoneName terra.StringValue `hcl:"zone_name,attr" validate:"required"`
	// Record: min=0
	Record []dataprivatednssrvrecord.Record `hcl:"record,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataprivatednssrvrecord.Timeouts `hcl:"timeouts,block"`
}
type dataPrivateDnsSrvRecordAttributes struct {
	ref terra.Reference
}

// Fqdn returns a reference to field fqdn of azurerm_private_dns_srv_record.
func (pdsr dataPrivateDnsSrvRecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(pdsr.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_private_dns_srv_record.
func (pdsr dataPrivateDnsSrvRecordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pdsr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_private_dns_srv_record.
func (pdsr dataPrivateDnsSrvRecordAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pdsr.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_private_dns_srv_record.
func (pdsr dataPrivateDnsSrvRecordAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(pdsr.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_private_dns_srv_record.
func (pdsr dataPrivateDnsSrvRecordAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pdsr.ref.Append("tags"))
}

// Ttl returns a reference to field ttl of azurerm_private_dns_srv_record.
func (pdsr dataPrivateDnsSrvRecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(pdsr.ref.Append("ttl"))
}

// ZoneName returns a reference to field zone_name of azurerm_private_dns_srv_record.
func (pdsr dataPrivateDnsSrvRecordAttributes) ZoneName() terra.StringValue {
	return terra.ReferenceAsString(pdsr.ref.Append("zone_name"))
}

func (pdsr dataPrivateDnsSrvRecordAttributes) Record() terra.SetValue[dataprivatednssrvrecord.RecordAttributes] {
	return terra.ReferenceAsSet[dataprivatednssrvrecord.RecordAttributes](pdsr.ref.Append("record"))
}

func (pdsr dataPrivateDnsSrvRecordAttributes) Timeouts() dataprivatednssrvrecord.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataprivatednssrvrecord.TimeoutsAttributes](pdsr.ref.Append("timeouts"))
}
