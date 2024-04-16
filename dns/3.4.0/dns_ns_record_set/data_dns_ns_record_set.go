// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package dns_ns_record_set

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource dns_ns_record_set.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (dnrs *DataSource) DataSource() string {
	return "dns_ns_record_set"
}

// LocalName returns the local name for [DataSource].
func (dnrs *DataSource) LocalName() string {
	return dnrs.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (dnrs *DataSource) Configuration() interface{} {
	return dnrs.Args
}

// Attributes returns the attributes for [DataSource].
func (dnrs *DataSource) Attributes() dataDnsNsRecordSetAttributes {
	return dataDnsNsRecordSetAttributes{ref: terra.ReferenceDataSource(dnrs)}
}

// DataArgs contains the configurations for dns_ns_record_set.
type DataArgs struct {
	// Host: string, required
	Host terra.StringValue `hcl:"host,attr" validate:"required"`
}

type dataDnsNsRecordSetAttributes struct {
	ref terra.Reference
}

// Host returns a reference to field host of dns_ns_record_set.
func (dnrs dataDnsNsRecordSetAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(dnrs.ref.Append("host"))
}

// Id returns a reference to field id of dns_ns_record_set.
func (dnrs dataDnsNsRecordSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dnrs.ref.Append("id"))
}

// Nameservers returns a reference to field nameservers of dns_ns_record_set.
func (dnrs dataDnsNsRecordSetAttributes) Nameservers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dnrs.ref.Append("nameservers"))
}
