// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataprivatednsptrrecord "github.com/golingon/terraproviders/azurerm/3.49.0/dataprivatednsptrrecord"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataPrivateDnsPtrRecord creates a new instance of [DataPrivateDnsPtrRecord].
func NewDataPrivateDnsPtrRecord(name string, args DataPrivateDnsPtrRecordArgs) *DataPrivateDnsPtrRecord {
	return &DataPrivateDnsPtrRecord{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPrivateDnsPtrRecord)(nil)

// DataPrivateDnsPtrRecord represents the Terraform data resource azurerm_private_dns_ptr_record.
type DataPrivateDnsPtrRecord struct {
	Name string
	Args DataPrivateDnsPtrRecordArgs
}

// DataSource returns the Terraform object type for [DataPrivateDnsPtrRecord].
func (pdpr *DataPrivateDnsPtrRecord) DataSource() string {
	return "azurerm_private_dns_ptr_record"
}

// LocalName returns the local name for [DataPrivateDnsPtrRecord].
func (pdpr *DataPrivateDnsPtrRecord) LocalName() string {
	return pdpr.Name
}

// Configuration returns the configuration (args) for [DataPrivateDnsPtrRecord].
func (pdpr *DataPrivateDnsPtrRecord) Configuration() interface{} {
	return pdpr.Args
}

// Attributes returns the attributes for [DataPrivateDnsPtrRecord].
func (pdpr *DataPrivateDnsPtrRecord) Attributes() dataPrivateDnsPtrRecordAttributes {
	return dataPrivateDnsPtrRecordAttributes{ref: terra.ReferenceDataResource(pdpr)}
}

// DataPrivateDnsPtrRecordArgs contains the configurations for azurerm_private_dns_ptr_record.
type DataPrivateDnsPtrRecordArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ZoneName: string, required
	ZoneName terra.StringValue `hcl:"zone_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dataprivatednsptrrecord.Timeouts `hcl:"timeouts,block"`
}
type dataPrivateDnsPtrRecordAttributes struct {
	ref terra.Reference
}

// Fqdn returns a reference to field fqdn of azurerm_private_dns_ptr_record.
func (pdpr dataPrivateDnsPtrRecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(pdpr.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_private_dns_ptr_record.
func (pdpr dataPrivateDnsPtrRecordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pdpr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_private_dns_ptr_record.
func (pdpr dataPrivateDnsPtrRecordAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pdpr.ref.Append("name"))
}

// Records returns a reference to field records of azurerm_private_dns_ptr_record.
func (pdpr dataPrivateDnsPtrRecordAttributes) Records() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](pdpr.ref.Append("records"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_private_dns_ptr_record.
func (pdpr dataPrivateDnsPtrRecordAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(pdpr.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_private_dns_ptr_record.
func (pdpr dataPrivateDnsPtrRecordAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pdpr.ref.Append("tags"))
}

// Ttl returns a reference to field ttl of azurerm_private_dns_ptr_record.
func (pdpr dataPrivateDnsPtrRecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(pdpr.ref.Append("ttl"))
}

// ZoneName returns a reference to field zone_name of azurerm_private_dns_ptr_record.
func (pdpr dataPrivateDnsPtrRecordAttributes) ZoneName() terra.StringValue {
	return terra.ReferenceAsString(pdpr.ref.Append("zone_name"))
}

func (pdpr dataPrivateDnsPtrRecordAttributes) Timeouts() dataprivatednsptrrecord.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataprivatednsptrrecord.TimeoutsAttributes](pdpr.ref.Append("timeouts"))
}
