// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	vpcipampoolcidr "github.com/golingon/terraproviders/aws/5.0.1/vpcipampoolcidr"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVpcIpamPoolCidr creates a new instance of [VpcIpamPoolCidr].
func NewVpcIpamPoolCidr(name string, args VpcIpamPoolCidrArgs) *VpcIpamPoolCidr {
	return &VpcIpamPoolCidr{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpcIpamPoolCidr)(nil)

// VpcIpamPoolCidr represents the Terraform resource aws_vpc_ipam_pool_cidr.
type VpcIpamPoolCidr struct {
	Name      string
	Args      VpcIpamPoolCidrArgs
	state     *vpcIpamPoolCidrState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VpcIpamPoolCidr].
func (vipc *VpcIpamPoolCidr) Type() string {
	return "aws_vpc_ipam_pool_cidr"
}

// LocalName returns the local name for [VpcIpamPoolCidr].
func (vipc *VpcIpamPoolCidr) LocalName() string {
	return vipc.Name
}

// Configuration returns the configuration (args) for [VpcIpamPoolCidr].
func (vipc *VpcIpamPoolCidr) Configuration() interface{} {
	return vipc.Args
}

// DependOn is used for other resources to depend on [VpcIpamPoolCidr].
func (vipc *VpcIpamPoolCidr) DependOn() terra.Reference {
	return terra.ReferenceResource(vipc)
}

// Dependencies returns the list of resources [VpcIpamPoolCidr] depends_on.
func (vipc *VpcIpamPoolCidr) Dependencies() terra.Dependencies {
	return vipc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VpcIpamPoolCidr].
func (vipc *VpcIpamPoolCidr) LifecycleManagement() *terra.Lifecycle {
	return vipc.Lifecycle
}

// Attributes returns the attributes for [VpcIpamPoolCidr].
func (vipc *VpcIpamPoolCidr) Attributes() vpcIpamPoolCidrAttributes {
	return vpcIpamPoolCidrAttributes{ref: terra.ReferenceResource(vipc)}
}

// ImportState imports the given attribute values into [VpcIpamPoolCidr]'s state.
func (vipc *VpcIpamPoolCidr) ImportState(av io.Reader) error {
	vipc.state = &vpcIpamPoolCidrState{}
	if err := json.NewDecoder(av).Decode(vipc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vipc.Type(), vipc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VpcIpamPoolCidr] has state.
func (vipc *VpcIpamPoolCidr) State() (*vpcIpamPoolCidrState, bool) {
	return vipc.state, vipc.state != nil
}

// StateMust returns the state for [VpcIpamPoolCidr]. Panics if the state is nil.
func (vipc *VpcIpamPoolCidr) StateMust() *vpcIpamPoolCidrState {
	if vipc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vipc.Type(), vipc.LocalName()))
	}
	return vipc.state
}

// VpcIpamPoolCidrArgs contains the configurations for aws_vpc_ipam_pool_cidr.
type VpcIpamPoolCidrArgs struct {
	// Cidr: string, optional
	Cidr terra.StringValue `hcl:"cidr,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpamPoolId: string, required
	IpamPoolId terra.StringValue `hcl:"ipam_pool_id,attr" validate:"required"`
	// NetmaskLength: number, optional
	NetmaskLength terra.NumberValue `hcl:"netmask_length,attr"`
	// CidrAuthorizationContext: optional
	CidrAuthorizationContext *vpcipampoolcidr.CidrAuthorizationContext `hcl:"cidr_authorization_context,block"`
	// Timeouts: optional
	Timeouts *vpcipampoolcidr.Timeouts `hcl:"timeouts,block"`
}
type vpcIpamPoolCidrAttributes struct {
	ref terra.Reference
}

// Cidr returns a reference to field cidr of aws_vpc_ipam_pool_cidr.
func (vipc vpcIpamPoolCidrAttributes) Cidr() terra.StringValue {
	return terra.ReferenceAsString(vipc.ref.Append("cidr"))
}

// Id returns a reference to field id of aws_vpc_ipam_pool_cidr.
func (vipc vpcIpamPoolCidrAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vipc.ref.Append("id"))
}

// IpamPoolCidrId returns a reference to field ipam_pool_cidr_id of aws_vpc_ipam_pool_cidr.
func (vipc vpcIpamPoolCidrAttributes) IpamPoolCidrId() terra.StringValue {
	return terra.ReferenceAsString(vipc.ref.Append("ipam_pool_cidr_id"))
}

// IpamPoolId returns a reference to field ipam_pool_id of aws_vpc_ipam_pool_cidr.
func (vipc vpcIpamPoolCidrAttributes) IpamPoolId() terra.StringValue {
	return terra.ReferenceAsString(vipc.ref.Append("ipam_pool_id"))
}

// NetmaskLength returns a reference to field netmask_length of aws_vpc_ipam_pool_cidr.
func (vipc vpcIpamPoolCidrAttributes) NetmaskLength() terra.NumberValue {
	return terra.ReferenceAsNumber(vipc.ref.Append("netmask_length"))
}

func (vipc vpcIpamPoolCidrAttributes) CidrAuthorizationContext() terra.ListValue[vpcipampoolcidr.CidrAuthorizationContextAttributes] {
	return terra.ReferenceAsList[vpcipampoolcidr.CidrAuthorizationContextAttributes](vipc.ref.Append("cidr_authorization_context"))
}

func (vipc vpcIpamPoolCidrAttributes) Timeouts() vpcipampoolcidr.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vpcipampoolcidr.TimeoutsAttributes](vipc.ref.Append("timeouts"))
}

type vpcIpamPoolCidrState struct {
	Cidr                     string                                          `json:"cidr"`
	Id                       string                                          `json:"id"`
	IpamPoolCidrId           string                                          `json:"ipam_pool_cidr_id"`
	IpamPoolId               string                                          `json:"ipam_pool_id"`
	NetmaskLength            float64                                         `json:"netmask_length"`
	CidrAuthorizationContext []vpcipampoolcidr.CidrAuthorizationContextState `json:"cidr_authorization_context"`
	Timeouts                 *vpcipampoolcidr.TimeoutsState                  `json:"timeouts"`
}
