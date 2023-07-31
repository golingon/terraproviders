// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataprivatednscnamerecord "github.com/golingon/terraproviders/azurerm/3.67.0/dataprivatednscnamerecord"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataPrivateDnsCnameRecord creates a new instance of [DataPrivateDnsCnameRecord].
func NewDataPrivateDnsCnameRecord(name string, args DataPrivateDnsCnameRecordArgs) *DataPrivateDnsCnameRecord {
	return &DataPrivateDnsCnameRecord{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPrivateDnsCnameRecord)(nil)

// DataPrivateDnsCnameRecord represents the Terraform data resource azurerm_private_dns_cname_record.
type DataPrivateDnsCnameRecord struct {
	Name string
	Args DataPrivateDnsCnameRecordArgs
}

// DataSource returns the Terraform object type for [DataPrivateDnsCnameRecord].
func (pdcr *DataPrivateDnsCnameRecord) DataSource() string {
	return "azurerm_private_dns_cname_record"
}

// LocalName returns the local name for [DataPrivateDnsCnameRecord].
func (pdcr *DataPrivateDnsCnameRecord) LocalName() string {
	return pdcr.Name
}

// Configuration returns the configuration (args) for [DataPrivateDnsCnameRecord].
func (pdcr *DataPrivateDnsCnameRecord) Configuration() interface{} {
	return pdcr.Args
}

// Attributes returns the attributes for [DataPrivateDnsCnameRecord].
func (pdcr *DataPrivateDnsCnameRecord) Attributes() dataPrivateDnsCnameRecordAttributes {
	return dataPrivateDnsCnameRecordAttributes{ref: terra.ReferenceDataResource(pdcr)}
}

// DataPrivateDnsCnameRecordArgs contains the configurations for azurerm_private_dns_cname_record.
type DataPrivateDnsCnameRecordArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ZoneName: string, required
	ZoneName terra.StringValue `hcl:"zone_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dataprivatednscnamerecord.Timeouts `hcl:"timeouts,block"`
}
type dataPrivateDnsCnameRecordAttributes struct {
	ref terra.Reference
}

// Fqdn returns a reference to field fqdn of azurerm_private_dns_cname_record.
func (pdcr dataPrivateDnsCnameRecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(pdcr.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_private_dns_cname_record.
func (pdcr dataPrivateDnsCnameRecordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pdcr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_private_dns_cname_record.
func (pdcr dataPrivateDnsCnameRecordAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pdcr.ref.Append("name"))
}

// Record returns a reference to field record of azurerm_private_dns_cname_record.
func (pdcr dataPrivateDnsCnameRecordAttributes) Record() terra.StringValue {
	return terra.ReferenceAsString(pdcr.ref.Append("record"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_private_dns_cname_record.
func (pdcr dataPrivateDnsCnameRecordAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(pdcr.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_private_dns_cname_record.
func (pdcr dataPrivateDnsCnameRecordAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pdcr.ref.Append("tags"))
}

// TargetResourceId returns a reference to field target_resource_id of azurerm_private_dns_cname_record.
func (pdcr dataPrivateDnsCnameRecordAttributes) TargetResourceId() terra.StringValue {
	return terra.ReferenceAsString(pdcr.ref.Append("target_resource_id"))
}

// Ttl returns a reference to field ttl of azurerm_private_dns_cname_record.
func (pdcr dataPrivateDnsCnameRecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(pdcr.ref.Append("ttl"))
}

// ZoneName returns a reference to field zone_name of azurerm_private_dns_cname_record.
func (pdcr dataPrivateDnsCnameRecordAttributes) ZoneName() terra.StringValue {
	return terra.ReferenceAsString(pdcr.ref.Append("zone_name"))
}

func (pdcr dataPrivateDnsCnameRecordAttributes) Timeouts() dataprivatednscnamerecord.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataprivatednscnamerecord.TimeoutsAttributes](pdcr.ref.Append("timeouts"))
}
