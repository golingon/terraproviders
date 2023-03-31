// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datadnsnsrecord "github.com/golingon/terraproviders/azurerm/3.49.0/datadnsnsrecord"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataDnsNsRecord(name string, args DataDnsNsRecordArgs) *DataDnsNsRecord {
	return &DataDnsNsRecord{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDnsNsRecord)(nil)

type DataDnsNsRecord struct {
	Name string
	Args DataDnsNsRecordArgs
}

func (dnr *DataDnsNsRecord) DataSource() string {
	return "azurerm_dns_ns_record"
}

func (dnr *DataDnsNsRecord) LocalName() string {
	return dnr.Name
}

func (dnr *DataDnsNsRecord) Configuration() interface{} {
	return dnr.Args
}

func (dnr *DataDnsNsRecord) Attributes() dataDnsNsRecordAttributes {
	return dataDnsNsRecordAttributes{ref: terra.ReferenceDataResource(dnr)}
}

type DataDnsNsRecordArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ZoneName: string, required
	ZoneName terra.StringValue `hcl:"zone_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datadnsnsrecord.Timeouts `hcl:"timeouts,block"`
}
type dataDnsNsRecordAttributes struct {
	ref terra.Reference
}

func (dnr dataDnsNsRecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceString(dnr.ref.Append("fqdn"))
}

func (dnr dataDnsNsRecordAttributes) Id() terra.StringValue {
	return terra.ReferenceString(dnr.ref.Append("id"))
}

func (dnr dataDnsNsRecordAttributes) Name() terra.StringValue {
	return terra.ReferenceString(dnr.ref.Append("name"))
}

func (dnr dataDnsNsRecordAttributes) Records() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](dnr.ref.Append("records"))
}

func (dnr dataDnsNsRecordAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(dnr.ref.Append("resource_group_name"))
}

func (dnr dataDnsNsRecordAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](dnr.ref.Append("tags"))
}

func (dnr dataDnsNsRecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceNumber(dnr.ref.Append("ttl"))
}

func (dnr dataDnsNsRecordAttributes) ZoneName() terra.StringValue {
	return terra.ReferenceString(dnr.ref.Append("zone_name"))
}

func (dnr dataDnsNsRecordAttributes) Timeouts() datadnsnsrecord.TimeoutsAttributes {
	return terra.ReferenceSingle[datadnsnsrecord.TimeoutsAttributes](dnr.ref.Append("timeouts"))
}
