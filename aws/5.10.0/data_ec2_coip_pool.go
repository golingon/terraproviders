// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataec2coippool "github.com/golingon/terraproviders/aws/5.10.0/dataec2coippool"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEc2CoipPool creates a new instance of [DataEc2CoipPool].
func NewDataEc2CoipPool(name string, args DataEc2CoipPoolArgs) *DataEc2CoipPool {
	return &DataEc2CoipPool{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEc2CoipPool)(nil)

// DataEc2CoipPool represents the Terraform data resource aws_ec2_coip_pool.
type DataEc2CoipPool struct {
	Name string
	Args DataEc2CoipPoolArgs
}

// DataSource returns the Terraform object type for [DataEc2CoipPool].
func (ecp *DataEc2CoipPool) DataSource() string {
	return "aws_ec2_coip_pool"
}

// LocalName returns the local name for [DataEc2CoipPool].
func (ecp *DataEc2CoipPool) LocalName() string {
	return ecp.Name
}

// Configuration returns the configuration (args) for [DataEc2CoipPool].
func (ecp *DataEc2CoipPool) Configuration() interface{} {
	return ecp.Args
}

// Attributes returns the attributes for [DataEc2CoipPool].
func (ecp *DataEc2CoipPool) Attributes() dataEc2CoipPoolAttributes {
	return dataEc2CoipPoolAttributes{ref: terra.ReferenceDataResource(ecp)}
}

// DataEc2CoipPoolArgs contains the configurations for aws_ec2_coip_pool.
type DataEc2CoipPoolArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LocalGatewayRouteTableId: string, optional
	LocalGatewayRouteTableId terra.StringValue `hcl:"local_gateway_route_table_id,attr"`
	// PoolId: string, optional
	PoolId terra.StringValue `hcl:"pool_id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Filter: min=0
	Filter []dataec2coippool.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataec2coippool.Timeouts `hcl:"timeouts,block"`
}
type dataEc2CoipPoolAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ec2_coip_pool.
func (ecp dataEc2CoipPoolAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ecp.ref.Append("arn"))
}

// Id returns a reference to field id of aws_ec2_coip_pool.
func (ecp dataEc2CoipPoolAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ecp.ref.Append("id"))
}

// LocalGatewayRouteTableId returns a reference to field local_gateway_route_table_id of aws_ec2_coip_pool.
func (ecp dataEc2CoipPoolAttributes) LocalGatewayRouteTableId() terra.StringValue {
	return terra.ReferenceAsString(ecp.ref.Append("local_gateway_route_table_id"))
}

// PoolCidrs returns a reference to field pool_cidrs of aws_ec2_coip_pool.
func (ecp dataEc2CoipPoolAttributes) PoolCidrs() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ecp.ref.Append("pool_cidrs"))
}

// PoolId returns a reference to field pool_id of aws_ec2_coip_pool.
func (ecp dataEc2CoipPoolAttributes) PoolId() terra.StringValue {
	return terra.ReferenceAsString(ecp.ref.Append("pool_id"))
}

// Tags returns a reference to field tags of aws_ec2_coip_pool.
func (ecp dataEc2CoipPoolAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ecp.ref.Append("tags"))
}

func (ecp dataEc2CoipPoolAttributes) Filter() terra.SetValue[dataec2coippool.FilterAttributes] {
	return terra.ReferenceAsSet[dataec2coippool.FilterAttributes](ecp.ref.Append("filter"))
}

func (ecp dataEc2CoipPoolAttributes) Timeouts() dataec2coippool.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataec2coippool.TimeoutsAttributes](ecp.ref.Append("timeouts"))
}
