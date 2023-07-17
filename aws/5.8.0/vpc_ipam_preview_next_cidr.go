// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVpcIpamPreviewNextCidr creates a new instance of [VpcIpamPreviewNextCidr].
func NewVpcIpamPreviewNextCidr(name string, args VpcIpamPreviewNextCidrArgs) *VpcIpamPreviewNextCidr {
	return &VpcIpamPreviewNextCidr{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpcIpamPreviewNextCidr)(nil)

// VpcIpamPreviewNextCidr represents the Terraform resource aws_vpc_ipam_preview_next_cidr.
type VpcIpamPreviewNextCidr struct {
	Name      string
	Args      VpcIpamPreviewNextCidrArgs
	state     *vpcIpamPreviewNextCidrState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VpcIpamPreviewNextCidr].
func (vipnc *VpcIpamPreviewNextCidr) Type() string {
	return "aws_vpc_ipam_preview_next_cidr"
}

// LocalName returns the local name for [VpcIpamPreviewNextCidr].
func (vipnc *VpcIpamPreviewNextCidr) LocalName() string {
	return vipnc.Name
}

// Configuration returns the configuration (args) for [VpcIpamPreviewNextCidr].
func (vipnc *VpcIpamPreviewNextCidr) Configuration() interface{} {
	return vipnc.Args
}

// DependOn is used for other resources to depend on [VpcIpamPreviewNextCidr].
func (vipnc *VpcIpamPreviewNextCidr) DependOn() terra.Reference {
	return terra.ReferenceResource(vipnc)
}

// Dependencies returns the list of resources [VpcIpamPreviewNextCidr] depends_on.
func (vipnc *VpcIpamPreviewNextCidr) Dependencies() terra.Dependencies {
	return vipnc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VpcIpamPreviewNextCidr].
func (vipnc *VpcIpamPreviewNextCidr) LifecycleManagement() *terra.Lifecycle {
	return vipnc.Lifecycle
}

// Attributes returns the attributes for [VpcIpamPreviewNextCidr].
func (vipnc *VpcIpamPreviewNextCidr) Attributes() vpcIpamPreviewNextCidrAttributes {
	return vpcIpamPreviewNextCidrAttributes{ref: terra.ReferenceResource(vipnc)}
}

// ImportState imports the given attribute values into [VpcIpamPreviewNextCidr]'s state.
func (vipnc *VpcIpamPreviewNextCidr) ImportState(av io.Reader) error {
	vipnc.state = &vpcIpamPreviewNextCidrState{}
	if err := json.NewDecoder(av).Decode(vipnc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vipnc.Type(), vipnc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VpcIpamPreviewNextCidr] has state.
func (vipnc *VpcIpamPreviewNextCidr) State() (*vpcIpamPreviewNextCidrState, bool) {
	return vipnc.state, vipnc.state != nil
}

// StateMust returns the state for [VpcIpamPreviewNextCidr]. Panics if the state is nil.
func (vipnc *VpcIpamPreviewNextCidr) StateMust() *vpcIpamPreviewNextCidrState {
	if vipnc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vipnc.Type(), vipnc.LocalName()))
	}
	return vipnc.state
}

// VpcIpamPreviewNextCidrArgs contains the configurations for aws_vpc_ipam_preview_next_cidr.
type VpcIpamPreviewNextCidrArgs struct {
	// DisallowedCidrs: set of string, optional
	DisallowedCidrs terra.SetValue[terra.StringValue] `hcl:"disallowed_cidrs,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpamPoolId: string, required
	IpamPoolId terra.StringValue `hcl:"ipam_pool_id,attr" validate:"required"`
	// NetmaskLength: number, optional
	NetmaskLength terra.NumberValue `hcl:"netmask_length,attr"`
}
type vpcIpamPreviewNextCidrAttributes struct {
	ref terra.Reference
}

// Cidr returns a reference to field cidr of aws_vpc_ipam_preview_next_cidr.
func (vipnc vpcIpamPreviewNextCidrAttributes) Cidr() terra.StringValue {
	return terra.ReferenceAsString(vipnc.ref.Append("cidr"))
}

// DisallowedCidrs returns a reference to field disallowed_cidrs of aws_vpc_ipam_preview_next_cidr.
func (vipnc vpcIpamPreviewNextCidrAttributes) DisallowedCidrs() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vipnc.ref.Append("disallowed_cidrs"))
}

// Id returns a reference to field id of aws_vpc_ipam_preview_next_cidr.
func (vipnc vpcIpamPreviewNextCidrAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vipnc.ref.Append("id"))
}

// IpamPoolId returns a reference to field ipam_pool_id of aws_vpc_ipam_preview_next_cidr.
func (vipnc vpcIpamPreviewNextCidrAttributes) IpamPoolId() terra.StringValue {
	return terra.ReferenceAsString(vipnc.ref.Append("ipam_pool_id"))
}

// NetmaskLength returns a reference to field netmask_length of aws_vpc_ipam_preview_next_cidr.
func (vipnc vpcIpamPreviewNextCidrAttributes) NetmaskLength() terra.NumberValue {
	return terra.ReferenceAsNumber(vipnc.ref.Append("netmask_length"))
}

type vpcIpamPreviewNextCidrState struct {
	Cidr            string   `json:"cidr"`
	DisallowedCidrs []string `json:"disallowed_cidrs"`
	Id              string   `json:"id"`
	IpamPoolId      string   `json:"ipam_pool_id"`
	NetmaskLength   float64  `json:"netmask_length"`
}
