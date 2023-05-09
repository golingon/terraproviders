// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	vpcipv4cidrblockassociation "github.com/golingon/terraproviders/aws/4.66.1/vpcipv4cidrblockassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVpcIpv4CidrBlockAssociation creates a new instance of [VpcIpv4CidrBlockAssociation].
func NewVpcIpv4CidrBlockAssociation(name string, args VpcIpv4CidrBlockAssociationArgs) *VpcIpv4CidrBlockAssociation {
	return &VpcIpv4CidrBlockAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpcIpv4CidrBlockAssociation)(nil)

// VpcIpv4CidrBlockAssociation represents the Terraform resource aws_vpc_ipv4_cidr_block_association.
type VpcIpv4CidrBlockAssociation struct {
	Name      string
	Args      VpcIpv4CidrBlockAssociationArgs
	state     *vpcIpv4CidrBlockAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VpcIpv4CidrBlockAssociation].
func (vicba *VpcIpv4CidrBlockAssociation) Type() string {
	return "aws_vpc_ipv4_cidr_block_association"
}

// LocalName returns the local name for [VpcIpv4CidrBlockAssociation].
func (vicba *VpcIpv4CidrBlockAssociation) LocalName() string {
	return vicba.Name
}

// Configuration returns the configuration (args) for [VpcIpv4CidrBlockAssociation].
func (vicba *VpcIpv4CidrBlockAssociation) Configuration() interface{} {
	return vicba.Args
}

// DependOn is used for other resources to depend on [VpcIpv4CidrBlockAssociation].
func (vicba *VpcIpv4CidrBlockAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(vicba)
}

// Dependencies returns the list of resources [VpcIpv4CidrBlockAssociation] depends_on.
func (vicba *VpcIpv4CidrBlockAssociation) Dependencies() terra.Dependencies {
	return vicba.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VpcIpv4CidrBlockAssociation].
func (vicba *VpcIpv4CidrBlockAssociation) LifecycleManagement() *terra.Lifecycle {
	return vicba.Lifecycle
}

// Attributes returns the attributes for [VpcIpv4CidrBlockAssociation].
func (vicba *VpcIpv4CidrBlockAssociation) Attributes() vpcIpv4CidrBlockAssociationAttributes {
	return vpcIpv4CidrBlockAssociationAttributes{ref: terra.ReferenceResource(vicba)}
}

// ImportState imports the given attribute values into [VpcIpv4CidrBlockAssociation]'s state.
func (vicba *VpcIpv4CidrBlockAssociation) ImportState(av io.Reader) error {
	vicba.state = &vpcIpv4CidrBlockAssociationState{}
	if err := json.NewDecoder(av).Decode(vicba.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vicba.Type(), vicba.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VpcIpv4CidrBlockAssociation] has state.
func (vicba *VpcIpv4CidrBlockAssociation) State() (*vpcIpv4CidrBlockAssociationState, bool) {
	return vicba.state, vicba.state != nil
}

// StateMust returns the state for [VpcIpv4CidrBlockAssociation]. Panics if the state is nil.
func (vicba *VpcIpv4CidrBlockAssociation) StateMust() *vpcIpv4CidrBlockAssociationState {
	if vicba.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vicba.Type(), vicba.LocalName()))
	}
	return vicba.state
}

// VpcIpv4CidrBlockAssociationArgs contains the configurations for aws_vpc_ipv4_cidr_block_association.
type VpcIpv4CidrBlockAssociationArgs struct {
	// CidrBlock: string, optional
	CidrBlock terra.StringValue `hcl:"cidr_block,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Ipv4IpamPoolId: string, optional
	Ipv4IpamPoolId terra.StringValue `hcl:"ipv4_ipam_pool_id,attr"`
	// Ipv4NetmaskLength: number, optional
	Ipv4NetmaskLength terra.NumberValue `hcl:"ipv4_netmask_length,attr"`
	// VpcId: string, required
	VpcId terra.StringValue `hcl:"vpc_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *vpcipv4cidrblockassociation.Timeouts `hcl:"timeouts,block"`
}
type vpcIpv4CidrBlockAssociationAttributes struct {
	ref terra.Reference
}

// CidrBlock returns a reference to field cidr_block of aws_vpc_ipv4_cidr_block_association.
func (vicba vpcIpv4CidrBlockAssociationAttributes) CidrBlock() terra.StringValue {
	return terra.ReferenceAsString(vicba.ref.Append("cidr_block"))
}

// Id returns a reference to field id of aws_vpc_ipv4_cidr_block_association.
func (vicba vpcIpv4CidrBlockAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vicba.ref.Append("id"))
}

// Ipv4IpamPoolId returns a reference to field ipv4_ipam_pool_id of aws_vpc_ipv4_cidr_block_association.
func (vicba vpcIpv4CidrBlockAssociationAttributes) Ipv4IpamPoolId() terra.StringValue {
	return terra.ReferenceAsString(vicba.ref.Append("ipv4_ipam_pool_id"))
}

// Ipv4NetmaskLength returns a reference to field ipv4_netmask_length of aws_vpc_ipv4_cidr_block_association.
func (vicba vpcIpv4CidrBlockAssociationAttributes) Ipv4NetmaskLength() terra.NumberValue {
	return terra.ReferenceAsNumber(vicba.ref.Append("ipv4_netmask_length"))
}

// VpcId returns a reference to field vpc_id of aws_vpc_ipv4_cidr_block_association.
func (vicba vpcIpv4CidrBlockAssociationAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(vicba.ref.Append("vpc_id"))
}

func (vicba vpcIpv4CidrBlockAssociationAttributes) Timeouts() vpcipv4cidrblockassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vpcipv4cidrblockassociation.TimeoutsAttributes](vicba.ref.Append("timeouts"))
}

type vpcIpv4CidrBlockAssociationState struct {
	CidrBlock         string                                     `json:"cidr_block"`
	Id                string                                     `json:"id"`
	Ipv4IpamPoolId    string                                     `json:"ipv4_ipam_pool_id"`
	Ipv4NetmaskLength float64                                    `json:"ipv4_netmask_length"`
	VpcId             string                                     `json:"vpc_id"`
	Timeouts          *vpcipv4cidrblockassociation.TimeoutsState `json:"timeouts"`
}
