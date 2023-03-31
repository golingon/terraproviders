// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import "github.com/volvo-cars/lingon/pkg/terra"

func NewDataDnsRecordSet(name string, args DataDnsRecordSetArgs) *DataDnsRecordSet {
	return &DataDnsRecordSet{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDnsRecordSet)(nil)

type DataDnsRecordSet struct {
	Name string
	Args DataDnsRecordSetArgs
}

func (drs *DataDnsRecordSet) DataSource() string {
	return "google_dns_record_set"
}

func (drs *DataDnsRecordSet) LocalName() string {
	return drs.Name
}

func (drs *DataDnsRecordSet) Configuration() interface{} {
	return drs.Args
}

func (drs *DataDnsRecordSet) Attributes() dataDnsRecordSetAttributes {
	return dataDnsRecordSetAttributes{ref: terra.ReferenceDataResource(drs)}
}

type DataDnsRecordSetArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ManagedZone: string, required
	ManagedZone terra.StringValue `hcl:"managed_zone,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}
type dataDnsRecordSetAttributes struct {
	ref terra.Reference
}

func (drs dataDnsRecordSetAttributes) Id() terra.StringValue {
	return terra.ReferenceString(drs.ref.Append("id"))
}

func (drs dataDnsRecordSetAttributes) ManagedZone() terra.StringValue {
	return terra.ReferenceString(drs.ref.Append("managed_zone"))
}

func (drs dataDnsRecordSetAttributes) Name() terra.StringValue {
	return terra.ReferenceString(drs.ref.Append("name"))
}

func (drs dataDnsRecordSetAttributes) Project() terra.StringValue {
	return terra.ReferenceString(drs.ref.Append("project"))
}

func (drs dataDnsRecordSetAttributes) Rrdatas() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](drs.ref.Append("rrdatas"))
}

func (drs dataDnsRecordSetAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceNumber(drs.ref.Append("ttl"))
}

func (drs dataDnsRecordSetAttributes) Type() terra.StringValue {
	return terra.ReferenceString(drs.ref.Append("type"))
}
