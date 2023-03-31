// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataprivatednsarecord "github.com/golingon/terraproviders/azurerm/3.49.0/dataprivatednsarecord"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataPrivateDnsARecord(name string, args DataPrivateDnsARecordArgs) *DataPrivateDnsARecord {
	return &DataPrivateDnsARecord{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPrivateDnsARecord)(nil)

type DataPrivateDnsARecord struct {
	Name string
	Args DataPrivateDnsARecordArgs
}

func (pdar *DataPrivateDnsARecord) DataSource() string {
	return "azurerm_private_dns_a_record"
}

func (pdar *DataPrivateDnsARecord) LocalName() string {
	return pdar.Name
}

func (pdar *DataPrivateDnsARecord) Configuration() interface{} {
	return pdar.Args
}

func (pdar *DataPrivateDnsARecord) Attributes() dataPrivateDnsARecordAttributes {
	return dataPrivateDnsARecordAttributes{ref: terra.ReferenceDataResource(pdar)}
}

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

func (pdar dataPrivateDnsARecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceString(pdar.ref.Append("fqdn"))
}

func (pdar dataPrivateDnsARecordAttributes) Id() terra.StringValue {
	return terra.ReferenceString(pdar.ref.Append("id"))
}

func (pdar dataPrivateDnsARecordAttributes) Name() terra.StringValue {
	return terra.ReferenceString(pdar.ref.Append("name"))
}

func (pdar dataPrivateDnsARecordAttributes) Records() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](pdar.ref.Append("records"))
}

func (pdar dataPrivateDnsARecordAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(pdar.ref.Append("resource_group_name"))
}

func (pdar dataPrivateDnsARecordAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](pdar.ref.Append("tags"))
}

func (pdar dataPrivateDnsARecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceNumber(pdar.ref.Append("ttl"))
}

func (pdar dataPrivateDnsARecordAttributes) ZoneName() terra.StringValue {
	return terra.ReferenceString(pdar.ref.Append("zone_name"))
}

func (pdar dataPrivateDnsARecordAttributes) Timeouts() dataprivatednsarecord.TimeoutsAttributes {
	return terra.ReferenceSingle[dataprivatednsarecord.TimeoutsAttributes](pdar.ref.Append("timeouts"))
}
