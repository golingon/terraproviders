// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataec2publicipv4pools "github.com/golingon/terraproviders/aws/5.12.0/dataec2publicipv4pools"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEc2PublicIpv4Pools creates a new instance of [DataEc2PublicIpv4Pools].
func NewDataEc2PublicIpv4Pools(name string, args DataEc2PublicIpv4PoolsArgs) *DataEc2PublicIpv4Pools {
	return &DataEc2PublicIpv4Pools{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEc2PublicIpv4Pools)(nil)

// DataEc2PublicIpv4Pools represents the Terraform data resource aws_ec2_public_ipv4_pools.
type DataEc2PublicIpv4Pools struct {
	Name string
	Args DataEc2PublicIpv4PoolsArgs
}

// DataSource returns the Terraform object type for [DataEc2PublicIpv4Pools].
func (epip *DataEc2PublicIpv4Pools) DataSource() string {
	return "aws_ec2_public_ipv4_pools"
}

// LocalName returns the local name for [DataEc2PublicIpv4Pools].
func (epip *DataEc2PublicIpv4Pools) LocalName() string {
	return epip.Name
}

// Configuration returns the configuration (args) for [DataEc2PublicIpv4Pools].
func (epip *DataEc2PublicIpv4Pools) Configuration() interface{} {
	return epip.Args
}

// Attributes returns the attributes for [DataEc2PublicIpv4Pools].
func (epip *DataEc2PublicIpv4Pools) Attributes() dataEc2PublicIpv4PoolsAttributes {
	return dataEc2PublicIpv4PoolsAttributes{ref: terra.ReferenceDataResource(epip)}
}

// DataEc2PublicIpv4PoolsArgs contains the configurations for aws_ec2_public_ipv4_pools.
type DataEc2PublicIpv4PoolsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Filter: min=0
	Filter []dataec2publicipv4pools.Filter `hcl:"filter,block" validate:"min=0"`
}
type dataEc2PublicIpv4PoolsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ec2_public_ipv4_pools.
func (epip dataEc2PublicIpv4PoolsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(epip.ref.Append("id"))
}

// PoolIds returns a reference to field pool_ids of aws_ec2_public_ipv4_pools.
func (epip dataEc2PublicIpv4PoolsAttributes) PoolIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](epip.ref.Append("pool_ids"))
}

// Tags returns a reference to field tags of aws_ec2_public_ipv4_pools.
func (epip dataEc2PublicIpv4PoolsAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](epip.ref.Append("tags"))
}

func (epip dataEc2PublicIpv4PoolsAttributes) Filter() terra.SetValue[dataec2publicipv4pools.FilterAttributes] {
	return terra.ReferenceAsSet[dataec2publicipv4pools.FilterAttributes](epip.ref.Append("filter"))
}