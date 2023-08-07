// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datadnsaaaarecord "github.com/golingon/terraproviders/azurerm/3.68.0/datadnsaaaarecord"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataDnsAaaaRecord creates a new instance of [DataDnsAaaaRecord].
func NewDataDnsAaaaRecord(name string, args DataDnsAaaaRecordArgs) *DataDnsAaaaRecord {
	return &DataDnsAaaaRecord{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDnsAaaaRecord)(nil)

// DataDnsAaaaRecord represents the Terraform data resource azurerm_dns_aaaa_record.
type DataDnsAaaaRecord struct {
	Name string
	Args DataDnsAaaaRecordArgs
}

// DataSource returns the Terraform object type for [DataDnsAaaaRecord].
func (dar *DataDnsAaaaRecord) DataSource() string {
	return "azurerm_dns_aaaa_record"
}

// LocalName returns the local name for [DataDnsAaaaRecord].
func (dar *DataDnsAaaaRecord) LocalName() string {
	return dar.Name
}

// Configuration returns the configuration (args) for [DataDnsAaaaRecord].
func (dar *DataDnsAaaaRecord) Configuration() interface{} {
	return dar.Args
}

// Attributes returns the attributes for [DataDnsAaaaRecord].
func (dar *DataDnsAaaaRecord) Attributes() dataDnsAaaaRecordAttributes {
	return dataDnsAaaaRecordAttributes{ref: terra.ReferenceDataResource(dar)}
}

// DataDnsAaaaRecordArgs contains the configurations for azurerm_dns_aaaa_record.
type DataDnsAaaaRecordArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ZoneName: string, required
	ZoneName terra.StringValue `hcl:"zone_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datadnsaaaarecord.Timeouts `hcl:"timeouts,block"`
}
type dataDnsAaaaRecordAttributes struct {
	ref terra.Reference
}

// Fqdn returns a reference to field fqdn of azurerm_dns_aaaa_record.
func (dar dataDnsAaaaRecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(dar.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_dns_aaaa_record.
func (dar dataDnsAaaaRecordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dar.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_dns_aaaa_record.
func (dar dataDnsAaaaRecordAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dar.ref.Append("name"))
}

// Records returns a reference to field records of azurerm_dns_aaaa_record.
func (dar dataDnsAaaaRecordAttributes) Records() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dar.ref.Append("records"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_dns_aaaa_record.
func (dar dataDnsAaaaRecordAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dar.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_dns_aaaa_record.
func (dar dataDnsAaaaRecordAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dar.ref.Append("tags"))
}

// TargetResourceId returns a reference to field target_resource_id of azurerm_dns_aaaa_record.
func (dar dataDnsAaaaRecordAttributes) TargetResourceId() terra.StringValue {
	return terra.ReferenceAsString(dar.ref.Append("target_resource_id"))
}

// Ttl returns a reference to field ttl of azurerm_dns_aaaa_record.
func (dar dataDnsAaaaRecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(dar.ref.Append("ttl"))
}

// ZoneName returns a reference to field zone_name of azurerm_dns_aaaa_record.
func (dar dataDnsAaaaRecordAttributes) ZoneName() terra.StringValue {
	return terra.ReferenceAsString(dar.ref.Append("zone_name"))
}

func (dar dataDnsAaaaRecordAttributes) Timeouts() datadnsaaaarecord.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datadnsaaaarecord.TimeoutsAttributes](dar.ref.Append("timeouts"))
}
