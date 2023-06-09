// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dns

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataNsRecordSet creates a new instance of [DataNsRecordSet].
func NewDataNsRecordSet(name string, args DataNsRecordSetArgs) *DataNsRecordSet {
	return &DataNsRecordSet{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataNsRecordSet)(nil)

// DataNsRecordSet represents the Terraform data resource dns_ns_record_set.
type DataNsRecordSet struct {
	Name string
	Args DataNsRecordSetArgs
}

// DataSource returns the Terraform object type for [DataNsRecordSet].
func (nrs *DataNsRecordSet) DataSource() string {
	return "dns_ns_record_set"
}

// LocalName returns the local name for [DataNsRecordSet].
func (nrs *DataNsRecordSet) LocalName() string {
	return nrs.Name
}

// Configuration returns the configuration (args) for [DataNsRecordSet].
func (nrs *DataNsRecordSet) Configuration() interface{} {
	return nrs.Args
}

// Attributes returns the attributes for [DataNsRecordSet].
func (nrs *DataNsRecordSet) Attributes() dataNsRecordSetAttributes {
	return dataNsRecordSetAttributes{ref: terra.ReferenceDataResource(nrs)}
}

// DataNsRecordSetArgs contains the configurations for dns_ns_record_set.
type DataNsRecordSetArgs struct {
	// Host: string, required
	Host terra.StringValue `hcl:"host,attr" validate:"required"`
}
type dataNsRecordSetAttributes struct {
	ref terra.Reference
}

// Host returns a reference to field host of dns_ns_record_set.
func (nrs dataNsRecordSetAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(nrs.ref.Append("host"))
}

// Id returns a reference to field id of dns_ns_record_set.
func (nrs dataNsRecordSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nrs.ref.Append("id"))
}

// Nameservers returns a reference to field nameservers of dns_ns_record_set.
func (nrs dataNsRecordSetAttributes) Nameservers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nrs.ref.Append("nameservers"))
}
