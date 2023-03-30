// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

func NewDataPartition(name string, args DataPartitionArgs) *DataPartition {
	return &DataPartition{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPartition)(nil)

type DataPartition struct {
	Name string
	Args DataPartitionArgs
}

func (p *DataPartition) DataSource() string {
	return "aws_partition"
}

func (p *DataPartition) LocalName() string {
	return p.Name
}

func (p *DataPartition) Configuration() interface{} {
	return p.Args
}

func (p *DataPartition) Attributes() dataPartitionAttributes {
	return dataPartitionAttributes{ref: terra.ReferenceDataResource(p)}
}

type DataPartitionArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type dataPartitionAttributes struct {
	ref terra.Reference
}

func (p dataPartitionAttributes) DnsSuffix() terra.StringValue {
	return terra.ReferenceString(p.ref.Append("dns_suffix"))
}

func (p dataPartitionAttributes) Id() terra.StringValue {
	return terra.ReferenceString(p.ref.Append("id"))
}

func (p dataPartitionAttributes) Partition() terra.StringValue {
	return terra.ReferenceString(p.ref.Append("partition"))
}

func (p dataPartitionAttributes) ReverseDnsPrefix() terra.StringValue {
	return terra.ReferenceString(p.ref.Append("reverse_dns_prefix"))
}
