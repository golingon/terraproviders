// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataPartition creates a new instance of [DataPartition].
func NewDataPartition(name string, args DataPartitionArgs) *DataPartition {
	return &DataPartition{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPartition)(nil)

// DataPartition represents the Terraform data resource aws_partition.
type DataPartition struct {
	Name string
	Args DataPartitionArgs
}

// DataSource returns the Terraform object type for [DataPartition].
func (p *DataPartition) DataSource() string {
	return "aws_partition"
}

// LocalName returns the local name for [DataPartition].
func (p *DataPartition) LocalName() string {
	return p.Name
}

// Configuration returns the configuration (args) for [DataPartition].
func (p *DataPartition) Configuration() interface{} {
	return p.Args
}

// Attributes returns the attributes for [DataPartition].
func (p *DataPartition) Attributes() dataPartitionAttributes {
	return dataPartitionAttributes{ref: terra.ReferenceDataResource(p)}
}

// DataPartitionArgs contains the configurations for aws_partition.
type DataPartitionArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type dataPartitionAttributes struct {
	ref terra.Reference
}

// DnsSuffix returns a reference to field dns_suffix of aws_partition.
func (p dataPartitionAttributes) DnsSuffix() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("dns_suffix"))
}

// Id returns a reference to field id of aws_partition.
func (p dataPartitionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("id"))
}

// Partition returns a reference to field partition of aws_partition.
func (p dataPartitionAttributes) Partition() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("partition"))
}

// ReverseDnsPrefix returns a reference to field reverse_dns_prefix of aws_partition.
func (p dataPartitionAttributes) ReverseDnsPrefix() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("reverse_dns_prefix"))
}
