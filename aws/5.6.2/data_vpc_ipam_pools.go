// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datavpcipampools "github.com/golingon/terraproviders/aws/5.6.2/datavpcipampools"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataVpcIpamPools creates a new instance of [DataVpcIpamPools].
func NewDataVpcIpamPools(name string, args DataVpcIpamPoolsArgs) *DataVpcIpamPools {
	return &DataVpcIpamPools{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataVpcIpamPools)(nil)

// DataVpcIpamPools represents the Terraform data resource aws_vpc_ipam_pools.
type DataVpcIpamPools struct {
	Name string
	Args DataVpcIpamPoolsArgs
}

// DataSource returns the Terraform object type for [DataVpcIpamPools].
func (vip *DataVpcIpamPools) DataSource() string {
	return "aws_vpc_ipam_pools"
}

// LocalName returns the local name for [DataVpcIpamPools].
func (vip *DataVpcIpamPools) LocalName() string {
	return vip.Name
}

// Configuration returns the configuration (args) for [DataVpcIpamPools].
func (vip *DataVpcIpamPools) Configuration() interface{} {
	return vip.Args
}

// Attributes returns the attributes for [DataVpcIpamPools].
func (vip *DataVpcIpamPools) Attributes() dataVpcIpamPoolsAttributes {
	return dataVpcIpamPoolsAttributes{ref: terra.ReferenceDataResource(vip)}
}

// DataVpcIpamPoolsArgs contains the configurations for aws_vpc_ipam_pools.
type DataVpcIpamPoolsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpamPools: min=0
	IpamPools []datavpcipampools.IpamPools `hcl:"ipam_pools,block" validate:"min=0"`
	// Filter: min=0
	Filter []datavpcipampools.Filter `hcl:"filter,block" validate:"min=0"`
}
type dataVpcIpamPoolsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_vpc_ipam_pools.
func (vip dataVpcIpamPoolsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vip.ref.Append("id"))
}

func (vip dataVpcIpamPoolsAttributes) IpamPools() terra.SetValue[datavpcipampools.IpamPoolsAttributes] {
	return terra.ReferenceAsSet[datavpcipampools.IpamPoolsAttributes](vip.ref.Append("ipam_pools"))
}

func (vip dataVpcIpamPoolsAttributes) Filter() terra.SetValue[datavpcipampools.FilterAttributes] {
	return terra.ReferenceAsSet[datavpcipampools.FilterAttributes](vip.ref.Append("filter"))
}
