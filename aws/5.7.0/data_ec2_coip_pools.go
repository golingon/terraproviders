// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataec2coippools "github.com/golingon/terraproviders/aws/5.7.0/dataec2coippools"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEc2CoipPools creates a new instance of [DataEc2CoipPools].
func NewDataEc2CoipPools(name string, args DataEc2CoipPoolsArgs) *DataEc2CoipPools {
	return &DataEc2CoipPools{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEc2CoipPools)(nil)

// DataEc2CoipPools represents the Terraform data resource aws_ec2_coip_pools.
type DataEc2CoipPools struct {
	Name string
	Args DataEc2CoipPoolsArgs
}

// DataSource returns the Terraform object type for [DataEc2CoipPools].
func (ecp *DataEc2CoipPools) DataSource() string {
	return "aws_ec2_coip_pools"
}

// LocalName returns the local name for [DataEc2CoipPools].
func (ecp *DataEc2CoipPools) LocalName() string {
	return ecp.Name
}

// Configuration returns the configuration (args) for [DataEc2CoipPools].
func (ecp *DataEc2CoipPools) Configuration() interface{} {
	return ecp.Args
}

// Attributes returns the attributes for [DataEc2CoipPools].
func (ecp *DataEc2CoipPools) Attributes() dataEc2CoipPoolsAttributes {
	return dataEc2CoipPoolsAttributes{ref: terra.ReferenceDataResource(ecp)}
}

// DataEc2CoipPoolsArgs contains the configurations for aws_ec2_coip_pools.
type DataEc2CoipPoolsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Filter: min=0
	Filter []dataec2coippools.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataec2coippools.Timeouts `hcl:"timeouts,block"`
}
type dataEc2CoipPoolsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ec2_coip_pools.
func (ecp dataEc2CoipPoolsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ecp.ref.Append("id"))
}

// PoolIds returns a reference to field pool_ids of aws_ec2_coip_pools.
func (ecp dataEc2CoipPoolsAttributes) PoolIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ecp.ref.Append("pool_ids"))
}

// Tags returns a reference to field tags of aws_ec2_coip_pools.
func (ecp dataEc2CoipPoolsAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ecp.ref.Append("tags"))
}

func (ecp dataEc2CoipPoolsAttributes) Filter() terra.SetValue[dataec2coippools.FilterAttributes] {
	return terra.ReferenceAsSet[dataec2coippools.FilterAttributes](ecp.ref.Append("filter"))
}

func (ecp dataEc2CoipPoolsAttributes) Timeouts() dataec2coippools.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataec2coippools.TimeoutsAttributes](ecp.ref.Append("timeouts"))
}
