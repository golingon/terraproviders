// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVpcIpamPoolCidrAllocation creates a new instance of [VpcIpamPoolCidrAllocation].
func NewVpcIpamPoolCidrAllocation(name string, args VpcIpamPoolCidrAllocationArgs) *VpcIpamPoolCidrAllocation {
	return &VpcIpamPoolCidrAllocation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpcIpamPoolCidrAllocation)(nil)

// VpcIpamPoolCidrAllocation represents the Terraform resource aws_vpc_ipam_pool_cidr_allocation.
type VpcIpamPoolCidrAllocation struct {
	Name      string
	Args      VpcIpamPoolCidrAllocationArgs
	state     *vpcIpamPoolCidrAllocationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VpcIpamPoolCidrAllocation].
func (vipca *VpcIpamPoolCidrAllocation) Type() string {
	return "aws_vpc_ipam_pool_cidr_allocation"
}

// LocalName returns the local name for [VpcIpamPoolCidrAllocation].
func (vipca *VpcIpamPoolCidrAllocation) LocalName() string {
	return vipca.Name
}

// Configuration returns the configuration (args) for [VpcIpamPoolCidrAllocation].
func (vipca *VpcIpamPoolCidrAllocation) Configuration() interface{} {
	return vipca.Args
}

// DependOn is used for other resources to depend on [VpcIpamPoolCidrAllocation].
func (vipca *VpcIpamPoolCidrAllocation) DependOn() terra.Reference {
	return terra.ReferenceResource(vipca)
}

// Dependencies returns the list of resources [VpcIpamPoolCidrAllocation] depends_on.
func (vipca *VpcIpamPoolCidrAllocation) Dependencies() terra.Dependencies {
	return vipca.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VpcIpamPoolCidrAllocation].
func (vipca *VpcIpamPoolCidrAllocation) LifecycleManagement() *terra.Lifecycle {
	return vipca.Lifecycle
}

// Attributes returns the attributes for [VpcIpamPoolCidrAllocation].
func (vipca *VpcIpamPoolCidrAllocation) Attributes() vpcIpamPoolCidrAllocationAttributes {
	return vpcIpamPoolCidrAllocationAttributes{ref: terra.ReferenceResource(vipca)}
}

// ImportState imports the given attribute values into [VpcIpamPoolCidrAllocation]'s state.
func (vipca *VpcIpamPoolCidrAllocation) ImportState(av io.Reader) error {
	vipca.state = &vpcIpamPoolCidrAllocationState{}
	if err := json.NewDecoder(av).Decode(vipca.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vipca.Type(), vipca.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VpcIpamPoolCidrAllocation] has state.
func (vipca *VpcIpamPoolCidrAllocation) State() (*vpcIpamPoolCidrAllocationState, bool) {
	return vipca.state, vipca.state != nil
}

// StateMust returns the state for [VpcIpamPoolCidrAllocation]. Panics if the state is nil.
func (vipca *VpcIpamPoolCidrAllocation) StateMust() *vpcIpamPoolCidrAllocationState {
	if vipca.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vipca.Type(), vipca.LocalName()))
	}
	return vipca.state
}

// VpcIpamPoolCidrAllocationArgs contains the configurations for aws_vpc_ipam_pool_cidr_allocation.
type VpcIpamPoolCidrAllocationArgs struct {
	// Cidr: string, optional
	Cidr terra.StringValue `hcl:"cidr,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisallowedCidrs: set of string, optional
	DisallowedCidrs terra.SetValue[terra.StringValue] `hcl:"disallowed_cidrs,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpamPoolId: string, required
	IpamPoolId terra.StringValue `hcl:"ipam_pool_id,attr" validate:"required"`
	// NetmaskLength: number, optional
	NetmaskLength terra.NumberValue `hcl:"netmask_length,attr"`
}
type vpcIpamPoolCidrAllocationAttributes struct {
	ref terra.Reference
}

// Cidr returns a reference to field cidr of aws_vpc_ipam_pool_cidr_allocation.
func (vipca vpcIpamPoolCidrAllocationAttributes) Cidr() terra.StringValue {
	return terra.ReferenceAsString(vipca.ref.Append("cidr"))
}

// Description returns a reference to field description of aws_vpc_ipam_pool_cidr_allocation.
func (vipca vpcIpamPoolCidrAllocationAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(vipca.ref.Append("description"))
}

// DisallowedCidrs returns a reference to field disallowed_cidrs of aws_vpc_ipam_pool_cidr_allocation.
func (vipca vpcIpamPoolCidrAllocationAttributes) DisallowedCidrs() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vipca.ref.Append("disallowed_cidrs"))
}

// Id returns a reference to field id of aws_vpc_ipam_pool_cidr_allocation.
func (vipca vpcIpamPoolCidrAllocationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vipca.ref.Append("id"))
}

// IpamPoolAllocationId returns a reference to field ipam_pool_allocation_id of aws_vpc_ipam_pool_cidr_allocation.
func (vipca vpcIpamPoolCidrAllocationAttributes) IpamPoolAllocationId() terra.StringValue {
	return terra.ReferenceAsString(vipca.ref.Append("ipam_pool_allocation_id"))
}

// IpamPoolId returns a reference to field ipam_pool_id of aws_vpc_ipam_pool_cidr_allocation.
func (vipca vpcIpamPoolCidrAllocationAttributes) IpamPoolId() terra.StringValue {
	return terra.ReferenceAsString(vipca.ref.Append("ipam_pool_id"))
}

// NetmaskLength returns a reference to field netmask_length of aws_vpc_ipam_pool_cidr_allocation.
func (vipca vpcIpamPoolCidrAllocationAttributes) NetmaskLength() terra.NumberValue {
	return terra.ReferenceAsNumber(vipca.ref.Append("netmask_length"))
}

// ResourceId returns a reference to field resource_id of aws_vpc_ipam_pool_cidr_allocation.
func (vipca vpcIpamPoolCidrAllocationAttributes) ResourceId() terra.StringValue {
	return terra.ReferenceAsString(vipca.ref.Append("resource_id"))
}

// ResourceOwner returns a reference to field resource_owner of aws_vpc_ipam_pool_cidr_allocation.
func (vipca vpcIpamPoolCidrAllocationAttributes) ResourceOwner() terra.StringValue {
	return terra.ReferenceAsString(vipca.ref.Append("resource_owner"))
}

// ResourceType returns a reference to field resource_type of aws_vpc_ipam_pool_cidr_allocation.
func (vipca vpcIpamPoolCidrAllocationAttributes) ResourceType() terra.StringValue {
	return terra.ReferenceAsString(vipca.ref.Append("resource_type"))
}

type vpcIpamPoolCidrAllocationState struct {
	Cidr                 string   `json:"cidr"`
	Description          string   `json:"description"`
	DisallowedCidrs      []string `json:"disallowed_cidrs"`
	Id                   string   `json:"id"`
	IpamPoolAllocationId string   `json:"ipam_pool_allocation_id"`
	IpamPoolId           string   `json:"ipam_pool_id"`
	NetmaskLength        float64  `json:"netmask_length"`
	ResourceId           string   `json:"resource_id"`
	ResourceOwner        string   `json:"resource_owner"`
	ResourceType         string   `json:"resource_type"`
}
