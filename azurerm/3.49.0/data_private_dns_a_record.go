// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataprivatednsarecord "github.com/golingon/terraproviders/azurerm/3.49.0/dataprivatednsarecord"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataPrivateDnsARecord creates a new instance of [DataPrivateDnsARecord].
func NewDataPrivateDnsARecord(name string, args DataPrivateDnsARecordArgs) *DataPrivateDnsARecord {
	return &DataPrivateDnsARecord{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPrivateDnsARecord)(nil)

// DataPrivateDnsARecord represents the Terraform data resource azurerm_private_dns_a_record.
type DataPrivateDnsARecord struct {
	Name string
	Args DataPrivateDnsARecordArgs
}

// DataSource returns the Terraform object type for [DataPrivateDnsARecord].
func (pdar *DataPrivateDnsARecord) DataSource() string {
	return "azurerm_private_dns_a_record"
}

// LocalName returns the local name for [DataPrivateDnsARecord].
func (pdar *DataPrivateDnsARecord) LocalName() string {
	return pdar.Name
}

// Configuration returns the configuration (args) for [DataPrivateDnsARecord].
func (pdar *DataPrivateDnsARecord) Configuration() interface{} {
	return pdar.Args
}

// Attributes returns the attributes for [DataPrivateDnsARecord].
func (pdar *DataPrivateDnsARecord) Attributes() dataPrivateDnsARecordAttributes {
	return dataPrivateDnsARecordAttributes{ref: terra.ReferenceDataResource(pdar)}
}

// DataPrivateDnsARecordArgs contains the configurations for azurerm_private_dns_a_record.
type DataPrivateDnsARecordArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ZoneName: string, required
	ZoneName terra.StringValue `hcl:"zone_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dataprivatednsarecord.Timeouts `hcl:"timeouts,block"`
}
type dataPrivateDnsARecordAttributes struct {
	ref terra.Reference
}

// Fqdn returns a reference to field fqdn of azurerm_private_dns_a_record.
func (pdar dataPrivateDnsARecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(pdar.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_private_dns_a_record.
func (pdar dataPrivateDnsARecordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pdar.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_private_dns_a_record.
func (pdar dataPrivateDnsARecordAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pdar.ref.Append("name"))
}

// Records returns a reference to field records of azurerm_private_dns_a_record.
func (pdar dataPrivateDnsARecordAttributes) Records() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](pdar.ref.Append("records"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_private_dns_a_record.
func (pdar dataPrivateDnsARecordAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(pdar.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_private_dns_a_record.
func (pdar dataPrivateDnsARecordAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pdar.ref.Append("tags"))
}

// Ttl returns a reference to field ttl of azurerm_private_dns_a_record.
func (pdar dataPrivateDnsARecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(pdar.ref.Append("ttl"))
}

// ZoneName returns a reference to field zone_name of azurerm_private_dns_a_record.
func (pdar dataPrivateDnsARecordAttributes) ZoneName() terra.StringValue {
	return terra.ReferenceAsString(pdar.ref.Append("zone_name"))
}

func (pdar dataPrivateDnsARecordAttributes) Timeouts() dataprivatednsarecord.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataprivatednsarecord.TimeoutsAttributes](pdar.ref.Append("timeouts"))
}
