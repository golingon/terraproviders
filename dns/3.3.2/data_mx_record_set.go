// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dns

import (
	datamxrecordset "github.com/golingon/terraproviders/dns/3.3.2/datamxrecordset"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataMxRecordSet creates a new instance of [DataMxRecordSet].
func NewDataMxRecordSet(name string, args DataMxRecordSetArgs) *DataMxRecordSet {
	return &DataMxRecordSet{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataMxRecordSet)(nil)

// DataMxRecordSet represents the Terraform data resource dns_mx_record_set.
type DataMxRecordSet struct {
	Name string
	Args DataMxRecordSetArgs
}

// DataSource returns the Terraform object type for [DataMxRecordSet].
func (mrs *DataMxRecordSet) DataSource() string {
	return "dns_mx_record_set"
}

// LocalName returns the local name for [DataMxRecordSet].
func (mrs *DataMxRecordSet) LocalName() string {
	return mrs.Name
}

// Configuration returns the configuration (args) for [DataMxRecordSet].
func (mrs *DataMxRecordSet) Configuration() interface{} {
	return mrs.Args
}

// Attributes returns the attributes for [DataMxRecordSet].
func (mrs *DataMxRecordSet) Attributes() dataMxRecordSetAttributes {
	return dataMxRecordSetAttributes{ref: terra.ReferenceDataResource(mrs)}
}

// DataMxRecordSetArgs contains the configurations for dns_mx_record_set.
type DataMxRecordSetArgs struct {
	// Domain: string, required
	Domain terra.StringValue `hcl:"domain,attr" validate:"required"`
	// Mx: min=0
	Mx []datamxrecordset.Mx `hcl:"mx,block" validate:"min=0"`
}
type dataMxRecordSetAttributes struct {
	ref terra.Reference
}

// Domain returns a reference to field domain of dns_mx_record_set.
func (mrs dataMxRecordSetAttributes) Domain() terra.StringValue {
	return terra.ReferenceAsString(mrs.ref.Append("domain"))
}

// Id returns a reference to field id of dns_mx_record_set.
func (mrs dataMxRecordSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mrs.ref.Append("id"))
}

func (mrs dataMxRecordSetAttributes) Mx() terra.ListValue[datamxrecordset.MxAttributes] {
	return terra.ReferenceAsList[datamxrecordset.MxAttributes](mrs.ref.Append("mx"))
}
