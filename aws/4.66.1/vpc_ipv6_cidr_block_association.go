// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	vpcipv6cidrblockassociation "github.com/golingon/terraproviders/aws/4.66.1/vpcipv6cidrblockassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVpcIpv6CidrBlockAssociation creates a new instance of [VpcIpv6CidrBlockAssociation].
func NewVpcIpv6CidrBlockAssociation(name string, args VpcIpv6CidrBlockAssociationArgs) *VpcIpv6CidrBlockAssociation {
	return &VpcIpv6CidrBlockAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpcIpv6CidrBlockAssociation)(nil)

// VpcIpv6CidrBlockAssociation represents the Terraform resource aws_vpc_ipv6_cidr_block_association.
type VpcIpv6CidrBlockAssociation struct {
	Name      string
	Args      VpcIpv6CidrBlockAssociationArgs
	state     *vpcIpv6CidrBlockAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VpcIpv6CidrBlockAssociation].
func (vicba *VpcIpv6CidrBlockAssociation) Type() string {
	return "aws_vpc_ipv6_cidr_block_association"
}

// LocalName returns the local name for [VpcIpv6CidrBlockAssociation].
func (vicba *VpcIpv6CidrBlockAssociation) LocalName() string {
	return vicba.Name
}

// Configuration returns the configuration (args) for [VpcIpv6CidrBlockAssociation].
func (vicba *VpcIpv6CidrBlockAssociation) Configuration() interface{} {
	return vicba.Args
}

// DependOn is used for other resources to depend on [VpcIpv6CidrBlockAssociation].
func (vicba *VpcIpv6CidrBlockAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(vicba)
}

// Dependencies returns the list of resources [VpcIpv6CidrBlockAssociation] depends_on.
func (vicba *VpcIpv6CidrBlockAssociation) Dependencies() terra.Dependencies {
	return vicba.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VpcIpv6CidrBlockAssociation].
func (vicba *VpcIpv6CidrBlockAssociation) LifecycleManagement() *terra.Lifecycle {
	return vicba.Lifecycle
}

// Attributes returns the attributes for [VpcIpv6CidrBlockAssociation].
func (vicba *VpcIpv6CidrBlockAssociation) Attributes() vpcIpv6CidrBlockAssociationAttributes {
	return vpcIpv6CidrBlockAssociationAttributes{ref: terra.ReferenceResource(vicba)}
}

// ImportState imports the given attribute values into [VpcIpv6CidrBlockAssociation]'s state.
func (vicba *VpcIpv6CidrBlockAssociation) ImportState(av io.Reader) error {
	vicba.state = &vpcIpv6CidrBlockAssociationState{}
	if err := json.NewDecoder(av).Decode(vicba.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vicba.Type(), vicba.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VpcIpv6CidrBlockAssociation] has state.
func (vicba *VpcIpv6CidrBlockAssociation) State() (*vpcIpv6CidrBlockAssociationState, bool) {
	return vicba.state, vicba.state != nil
}

// StateMust returns the state for [VpcIpv6CidrBlockAssociation]. Panics if the state is nil.
func (vicba *VpcIpv6CidrBlockAssociation) StateMust() *vpcIpv6CidrBlockAssociationState {
	if vicba.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vicba.Type(), vicba.LocalName()))
	}
	return vicba.state
}

// VpcIpv6CidrBlockAssociationArgs contains the configurations for aws_vpc_ipv6_cidr_block_association.
type VpcIpv6CidrBlockAssociationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Ipv6CidrBlock: string, optional
	Ipv6CidrBlock terra.StringValue `hcl:"ipv6_cidr_block,attr"`
	// Ipv6IpamPoolId: string, required
	Ipv6IpamPoolId terra.StringValue `hcl:"ipv6_ipam_pool_id,attr" validate:"required"`
	// Ipv6NetmaskLength: number, optional
	Ipv6NetmaskLength terra.NumberValue `hcl:"ipv6_netmask_length,attr"`
	// VpcId: string, required
	VpcId terra.StringValue `hcl:"vpc_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *vpcipv6cidrblockassociation.Timeouts `hcl:"timeouts,block"`
}
type vpcIpv6CidrBlockAssociationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_vpc_ipv6_cidr_block_association.
func (vicba vpcIpv6CidrBlockAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vicba.ref.Append("id"))
}

// Ipv6CidrBlock returns a reference to field ipv6_cidr_block of aws_vpc_ipv6_cidr_block_association.
func (vicba vpcIpv6CidrBlockAssociationAttributes) Ipv6CidrBlock() terra.StringValue {
	return terra.ReferenceAsString(vicba.ref.Append("ipv6_cidr_block"))
}

// Ipv6IpamPoolId returns a reference to field ipv6_ipam_pool_id of aws_vpc_ipv6_cidr_block_association.
func (vicba vpcIpv6CidrBlockAssociationAttributes) Ipv6IpamPoolId() terra.StringValue {
	return terra.ReferenceAsString(vicba.ref.Append("ipv6_ipam_pool_id"))
}

// Ipv6NetmaskLength returns a reference to field ipv6_netmask_length of aws_vpc_ipv6_cidr_block_association.
func (vicba vpcIpv6CidrBlockAssociationAttributes) Ipv6NetmaskLength() terra.NumberValue {
	return terra.ReferenceAsNumber(vicba.ref.Append("ipv6_netmask_length"))
}

// VpcId returns a reference to field vpc_id of aws_vpc_ipv6_cidr_block_association.
func (vicba vpcIpv6CidrBlockAssociationAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(vicba.ref.Append("vpc_id"))
}

func (vicba vpcIpv6CidrBlockAssociationAttributes) Timeouts() vpcipv6cidrblockassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vpcipv6cidrblockassociation.TimeoutsAttributes](vicba.ref.Append("timeouts"))
}

type vpcIpv6CidrBlockAssociationState struct {
	Id                string                                     `json:"id"`
	Ipv6CidrBlock     string                                     `json:"ipv6_cidr_block"`
	Ipv6IpamPoolId    string                                     `json:"ipv6_ipam_pool_id"`
	Ipv6NetmaskLength float64                                    `json:"ipv6_netmask_length"`
	VpcId             string                                     `json:"vpc_id"`
	Timeouts          *vpcipv6cidrblockassociation.TimeoutsState `json:"timeouts"`
}
