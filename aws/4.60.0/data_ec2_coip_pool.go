// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataec2coippool "github.com/golingon/terraproviders/aws/4.60.0/dataec2coippool"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataEc2CoipPool(name string, args DataEc2CoipPoolArgs) *DataEc2CoipPool {
	return &DataEc2CoipPool{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEc2CoipPool)(nil)

type DataEc2CoipPool struct {
	Name string
	Args DataEc2CoipPoolArgs
}

func (ecp *DataEc2CoipPool) DataSource() string {
	return "aws_ec2_coip_pool"
}

func (ecp *DataEc2CoipPool) LocalName() string {
	return ecp.Name
}

func (ecp *DataEc2CoipPool) Configuration() interface{} {
	return ecp.Args
}

func (ecp *DataEc2CoipPool) Attributes() dataEc2CoipPoolAttributes {
	return dataEc2CoipPoolAttributes{ref: terra.ReferenceDataResource(ecp)}
}

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

func (ecp dataEc2CoipPoolAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(ecp.ref.Append("arn"))
}

func (ecp dataEc2CoipPoolAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ecp.ref.Append("id"))
}

func (ecp dataEc2CoipPoolAttributes) LocalGatewayRouteTableId() terra.StringValue {
	return terra.ReferenceString(ecp.ref.Append("local_gateway_route_table_id"))
}

func (ecp dataEc2CoipPoolAttributes) PoolCidrs() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](ecp.ref.Append("pool_cidrs"))
}

func (ecp dataEc2CoipPoolAttributes) PoolId() terra.StringValue {
	return terra.ReferenceString(ecp.ref.Append("pool_id"))
}

func (ecp dataEc2CoipPoolAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](ecp.ref.Append("tags"))
}

func (ecp dataEc2CoipPoolAttributes) Filter() terra.SetValue[dataec2coippool.FilterAttributes] {
	return terra.ReferenceSet[dataec2coippool.FilterAttributes](ecp.ref.Append("filter"))
}

func (ecp dataEc2CoipPoolAttributes) Timeouts() dataec2coippool.TimeoutsAttributes {
	return terra.ReferenceSingle[dataec2coippool.TimeoutsAttributes](ecp.ref.Append("timeouts"))
}
