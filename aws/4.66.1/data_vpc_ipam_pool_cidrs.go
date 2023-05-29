// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datavpcipampoolcidrs "github.com/golingon/terraproviders/aws/4.66.1/datavpcipampoolcidrs"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataVpcIpamPoolCidrs creates a new instance of [DataVpcIpamPoolCidrs].
func NewDataVpcIpamPoolCidrs(name string, args DataVpcIpamPoolCidrsArgs) *DataVpcIpamPoolCidrs {
	return &DataVpcIpamPoolCidrs{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataVpcIpamPoolCidrs)(nil)

// DataVpcIpamPoolCidrs represents the Terraform data resource aws_vpc_ipam_pool_cidrs.
type DataVpcIpamPoolCidrs struct {
	Name string
	Args DataVpcIpamPoolCidrsArgs
}

// DataSource returns the Terraform object type for [DataVpcIpamPoolCidrs].
func (vipc *DataVpcIpamPoolCidrs) DataSource() string {
	return "aws_vpc_ipam_pool_cidrs"
}

// LocalName returns the local name for [DataVpcIpamPoolCidrs].
func (vipc *DataVpcIpamPoolCidrs) LocalName() string {
	return vipc.Name
}

// Configuration returns the configuration (args) for [DataVpcIpamPoolCidrs].
func (vipc *DataVpcIpamPoolCidrs) Configuration() interface{} {
	return vipc.Args
}

// Attributes returns the attributes for [DataVpcIpamPoolCidrs].
func (vipc *DataVpcIpamPoolCidrs) Attributes() dataVpcIpamPoolCidrsAttributes {
	return dataVpcIpamPoolCidrsAttributes{ref: terra.ReferenceDataResource(vipc)}
}

// DataVpcIpamPoolCidrsArgs contains the configurations for aws_vpc_ipam_pool_cidrs.
type DataVpcIpamPoolCidrsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpamPoolId: string, required
	IpamPoolId terra.StringValue `hcl:"ipam_pool_id,attr" validate:"required"`
	// IpamPoolCidrs: min=0
	IpamPoolCidrs []datavpcipampoolcidrs.IpamPoolCidrs `hcl:"ipam_pool_cidrs,block" validate:"min=0"`
	// Filter: min=0
	Filter []datavpcipampoolcidrs.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datavpcipampoolcidrs.Timeouts `hcl:"timeouts,block"`
}
type dataVpcIpamPoolCidrsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_vpc_ipam_pool_cidrs.
func (vipc dataVpcIpamPoolCidrsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vipc.ref.Append("id"))
}

// IpamPoolId returns a reference to field ipam_pool_id of aws_vpc_ipam_pool_cidrs.
func (vipc dataVpcIpamPoolCidrsAttributes) IpamPoolId() terra.StringValue {
	return terra.ReferenceAsString(vipc.ref.Append("ipam_pool_id"))
}

func (vipc dataVpcIpamPoolCidrsAttributes) IpamPoolCidrs() terra.SetValue[datavpcipampoolcidrs.IpamPoolCidrsAttributes] {
	return terra.ReferenceAsSet[datavpcipampoolcidrs.IpamPoolCidrsAttributes](vipc.ref.Append("ipam_pool_cidrs"))
}

func (vipc dataVpcIpamPoolCidrsAttributes) Filter() terra.SetValue[datavpcipampoolcidrs.FilterAttributes] {
	return terra.ReferenceAsSet[datavpcipampoolcidrs.FilterAttributes](vipc.ref.Append("filter"))
}

func (vipc dataVpcIpamPoolCidrsAttributes) Timeouts() datavpcipampoolcidrs.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datavpcipampoolcidrs.TimeoutsAttributes](vipc.ref.Append("timeouts"))
}