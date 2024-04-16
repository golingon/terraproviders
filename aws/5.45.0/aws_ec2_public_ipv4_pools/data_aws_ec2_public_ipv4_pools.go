// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_ec2_public_ipv4_pools

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_ec2_public_ipv4_pools.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (aepip *DataSource) DataSource() string {
	return "aws_ec2_public_ipv4_pools"
}

// LocalName returns the local name for [DataSource].
func (aepip *DataSource) LocalName() string {
	return aepip.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (aepip *DataSource) Configuration() interface{} {
	return aepip.Args
}

// Attributes returns the attributes for [DataSource].
func (aepip *DataSource) Attributes() dataAwsEc2PublicIpv4PoolsAttributes {
	return dataAwsEc2PublicIpv4PoolsAttributes{ref: terra.ReferenceDataSource(aepip)}
}

// DataArgs contains the configurations for aws_ec2_public_ipv4_pools.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Filter: min=0
	Filter []DataFilter `hcl:"filter,block" validate:"min=0"`
}

type dataAwsEc2PublicIpv4PoolsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ec2_public_ipv4_pools.
func (aepip dataAwsEc2PublicIpv4PoolsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aepip.ref.Append("id"))
}

// PoolIds returns a reference to field pool_ids of aws_ec2_public_ipv4_pools.
func (aepip dataAwsEc2PublicIpv4PoolsAttributes) PoolIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](aepip.ref.Append("pool_ids"))
}

// Tags returns a reference to field tags of aws_ec2_public_ipv4_pools.
func (aepip dataAwsEc2PublicIpv4PoolsAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aepip.ref.Append("tags"))
}

func (aepip dataAwsEc2PublicIpv4PoolsAttributes) Filter() terra.SetValue[DataFilterAttributes] {
	return terra.ReferenceAsSet[DataFilterAttributes](aepip.ref.Append("filter"))
}
