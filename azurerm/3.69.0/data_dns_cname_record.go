// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datadnscnamerecord "github.com/golingon/terraproviders/azurerm/3.69.0/datadnscnamerecord"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataDnsCnameRecord creates a new instance of [DataDnsCnameRecord].
func NewDataDnsCnameRecord(name string, args DataDnsCnameRecordArgs) *DataDnsCnameRecord {
	return &DataDnsCnameRecord{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDnsCnameRecord)(nil)

// DataDnsCnameRecord represents the Terraform data resource azurerm_dns_cname_record.
type DataDnsCnameRecord struct {
	Name string
	Args DataDnsCnameRecordArgs
}

// DataSource returns the Terraform object type for [DataDnsCnameRecord].
func (dcr *DataDnsCnameRecord) DataSource() string {
	return "azurerm_dns_cname_record"
}

// LocalName returns the local name for [DataDnsCnameRecord].
func (dcr *DataDnsCnameRecord) LocalName() string {
	return dcr.Name
}

// Configuration returns the configuration (args) for [DataDnsCnameRecord].
func (dcr *DataDnsCnameRecord) Configuration() interface{} {
	return dcr.Args
}

// Attributes returns the attributes for [DataDnsCnameRecord].
func (dcr *DataDnsCnameRecord) Attributes() dataDnsCnameRecordAttributes {
	return dataDnsCnameRecordAttributes{ref: terra.ReferenceDataResource(dcr)}
}

// DataDnsCnameRecordArgs contains the configurations for azurerm_dns_cname_record.
type DataDnsCnameRecordArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ZoneName: string, required
	ZoneName terra.StringValue `hcl:"zone_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datadnscnamerecord.Timeouts `hcl:"timeouts,block"`
}
type dataDnsCnameRecordAttributes struct {
	ref terra.Reference
}

// Fqdn returns a reference to field fqdn of azurerm_dns_cname_record.
func (dcr dataDnsCnameRecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(dcr.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_dns_cname_record.
func (dcr dataDnsCnameRecordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dcr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_dns_cname_record.
func (dcr dataDnsCnameRecordAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dcr.ref.Append("name"))
}

// Record returns a reference to field record of azurerm_dns_cname_record.
func (dcr dataDnsCnameRecordAttributes) Record() terra.StringValue {
	return terra.ReferenceAsString(dcr.ref.Append("record"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_dns_cname_record.
func (dcr dataDnsCnameRecordAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dcr.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_dns_cname_record.
func (dcr dataDnsCnameRecordAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dcr.ref.Append("tags"))
}

// TargetResourceId returns a reference to field target_resource_id of azurerm_dns_cname_record.
func (dcr dataDnsCnameRecordAttributes) TargetResourceId() terra.StringValue {
	return terra.ReferenceAsString(dcr.ref.Append("target_resource_id"))
}

// Ttl returns a reference to field ttl of azurerm_dns_cname_record.
func (dcr dataDnsCnameRecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(dcr.ref.Append("ttl"))
}

// ZoneName returns a reference to field zone_name of azurerm_dns_cname_record.
func (dcr dataDnsCnameRecordAttributes) ZoneName() terra.StringValue {
	return terra.ReferenceAsString(dcr.ref.Append("zone_name"))
}

func (dcr dataDnsCnameRecordAttributes) Timeouts() datadnscnamerecord.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datadnscnamerecord.TimeoutsAttributes](dcr.ref.Append("timeouts"))
}
