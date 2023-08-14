// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datadnsptrrecord "github.com/golingon/terraproviders/azurerm/3.69.0/datadnsptrrecord"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataDnsPtrRecord creates a new instance of [DataDnsPtrRecord].
func NewDataDnsPtrRecord(name string, args DataDnsPtrRecordArgs) *DataDnsPtrRecord {
	return &DataDnsPtrRecord{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDnsPtrRecord)(nil)

// DataDnsPtrRecord represents the Terraform data resource azurerm_dns_ptr_record.
type DataDnsPtrRecord struct {
	Name string
	Args DataDnsPtrRecordArgs
}

// DataSource returns the Terraform object type for [DataDnsPtrRecord].
func (dpr *DataDnsPtrRecord) DataSource() string {
	return "azurerm_dns_ptr_record"
}

// LocalName returns the local name for [DataDnsPtrRecord].
func (dpr *DataDnsPtrRecord) LocalName() string {
	return dpr.Name
}

// Configuration returns the configuration (args) for [DataDnsPtrRecord].
func (dpr *DataDnsPtrRecord) Configuration() interface{} {
	return dpr.Args
}

// Attributes returns the attributes for [DataDnsPtrRecord].
func (dpr *DataDnsPtrRecord) Attributes() dataDnsPtrRecordAttributes {
	return dataDnsPtrRecordAttributes{ref: terra.ReferenceDataResource(dpr)}
}

// DataDnsPtrRecordArgs contains the configurations for azurerm_dns_ptr_record.
type DataDnsPtrRecordArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ZoneName: string, required
	ZoneName terra.StringValue `hcl:"zone_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datadnsptrrecord.Timeouts `hcl:"timeouts,block"`
}
type dataDnsPtrRecordAttributes struct {
	ref terra.Reference
}

// Fqdn returns a reference to field fqdn of azurerm_dns_ptr_record.
func (dpr dataDnsPtrRecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(dpr.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_dns_ptr_record.
func (dpr dataDnsPtrRecordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dpr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_dns_ptr_record.
func (dpr dataDnsPtrRecordAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dpr.ref.Append("name"))
}

// Records returns a reference to field records of azurerm_dns_ptr_record.
func (dpr dataDnsPtrRecordAttributes) Records() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dpr.ref.Append("records"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_dns_ptr_record.
func (dpr dataDnsPtrRecordAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dpr.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_dns_ptr_record.
func (dpr dataDnsPtrRecordAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dpr.ref.Append("tags"))
}

// Ttl returns a reference to field ttl of azurerm_dns_ptr_record.
func (dpr dataDnsPtrRecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(dpr.ref.Append("ttl"))
}

// ZoneName returns a reference to field zone_name of azurerm_dns_ptr_record.
func (dpr dataDnsPtrRecordAttributes) ZoneName() terra.StringValue {
	return terra.ReferenceAsString(dpr.ref.Append("zone_name"))
}

func (dpr dataDnsPtrRecordAttributes) Timeouts() datadnsptrrecord.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datadnsptrrecord.TimeoutsAttributes](dpr.ref.Append("timeouts"))
}
