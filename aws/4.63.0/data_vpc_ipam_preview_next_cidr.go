// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datavpcipampreviewnextcidr "github.com/golingon/terraproviders/aws/4.63.0/datavpcipampreviewnextcidr"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataVpcIpamPreviewNextCidr creates a new instance of [DataVpcIpamPreviewNextCidr].
func NewDataVpcIpamPreviewNextCidr(name string, args DataVpcIpamPreviewNextCidrArgs) *DataVpcIpamPreviewNextCidr {
	return &DataVpcIpamPreviewNextCidr{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataVpcIpamPreviewNextCidr)(nil)

// DataVpcIpamPreviewNextCidr represents the Terraform data resource aws_vpc_ipam_preview_next_cidr.
type DataVpcIpamPreviewNextCidr struct {
	Name string
	Args DataVpcIpamPreviewNextCidrArgs
}

// DataSource returns the Terraform object type for [DataVpcIpamPreviewNextCidr].
func (vipnc *DataVpcIpamPreviewNextCidr) DataSource() string {
	return "aws_vpc_ipam_preview_next_cidr"
}

// LocalName returns the local name for [DataVpcIpamPreviewNextCidr].
func (vipnc *DataVpcIpamPreviewNextCidr) LocalName() string {
	return vipnc.Name
}

// Configuration returns the configuration (args) for [DataVpcIpamPreviewNextCidr].
func (vipnc *DataVpcIpamPreviewNextCidr) Configuration() interface{} {
	return vipnc.Args
}

// Attributes returns the attributes for [DataVpcIpamPreviewNextCidr].
func (vipnc *DataVpcIpamPreviewNextCidr) Attributes() dataVpcIpamPreviewNextCidrAttributes {
	return dataVpcIpamPreviewNextCidrAttributes{ref: terra.ReferenceDataResource(vipnc)}
}

// DataVpcIpamPreviewNextCidrArgs contains the configurations for aws_vpc_ipam_preview_next_cidr.
type DataVpcIpamPreviewNextCidrArgs struct {
	// DisallowedCidrs: set of string, optional
	DisallowedCidrs terra.SetValue[terra.StringValue] `hcl:"disallowed_cidrs,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpamPoolId: string, required
	IpamPoolId terra.StringValue `hcl:"ipam_pool_id,attr" validate:"required"`
	// NetmaskLength: number, optional
	NetmaskLength terra.NumberValue `hcl:"netmask_length,attr"`
	// Timeouts: optional
	Timeouts *datavpcipampreviewnextcidr.Timeouts `hcl:"timeouts,block"`
}
type dataVpcIpamPreviewNextCidrAttributes struct {
	ref terra.Reference
}

// Cidr returns a reference to field cidr of aws_vpc_ipam_preview_next_cidr.
func (vipnc dataVpcIpamPreviewNextCidrAttributes) Cidr() terra.StringValue {
	return terra.ReferenceAsString(vipnc.ref.Append("cidr"))
}

// DisallowedCidrs returns a reference to field disallowed_cidrs of aws_vpc_ipam_preview_next_cidr.
func (vipnc dataVpcIpamPreviewNextCidrAttributes) DisallowedCidrs() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vipnc.ref.Append("disallowed_cidrs"))
}

// Id returns a reference to field id of aws_vpc_ipam_preview_next_cidr.
func (vipnc dataVpcIpamPreviewNextCidrAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vipnc.ref.Append("id"))
}

// IpamPoolId returns a reference to field ipam_pool_id of aws_vpc_ipam_preview_next_cidr.
func (vipnc dataVpcIpamPreviewNextCidrAttributes) IpamPoolId() terra.StringValue {
	return terra.ReferenceAsString(vipnc.ref.Append("ipam_pool_id"))
}

// NetmaskLength returns a reference to field netmask_length of aws_vpc_ipam_preview_next_cidr.
func (vipnc dataVpcIpamPreviewNextCidrAttributes) NetmaskLength() terra.NumberValue {
	return terra.ReferenceAsNumber(vipnc.ref.Append("netmask_length"))
}

func (vipnc dataVpcIpamPreviewNextCidrAttributes) Timeouts() datavpcipampreviewnextcidr.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datavpcipampreviewnextcidr.TimeoutsAttributes](vipnc.ref.Append("timeouts"))
}
