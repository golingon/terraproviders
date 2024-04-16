// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_vpc_ipv6_cidr_block_association

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_vpc_ipv6_cidr_block_association.
type Resource struct {
	Name      string
	Args      Args
	state     *awsVpcIpv6CidrBlockAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (avicba *Resource) Type() string {
	return "aws_vpc_ipv6_cidr_block_association"
}

// LocalName returns the local name for [Resource].
func (avicba *Resource) LocalName() string {
	return avicba.Name
}

// Configuration returns the configuration (args) for [Resource].
func (avicba *Resource) Configuration() interface{} {
	return avicba.Args
}

// DependOn is used for other resources to depend on [Resource].
func (avicba *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(avicba)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (avicba *Resource) Dependencies() terra.Dependencies {
	return avicba.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (avicba *Resource) LifecycleManagement() *terra.Lifecycle {
	return avicba.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (avicba *Resource) Attributes() awsVpcIpv6CidrBlockAssociationAttributes {
	return awsVpcIpv6CidrBlockAssociationAttributes{ref: terra.ReferenceResource(avicba)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (avicba *Resource) ImportState(state io.Reader) error {
	avicba.state = &awsVpcIpv6CidrBlockAssociationState{}
	if err := json.NewDecoder(state).Decode(avicba.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", avicba.Type(), avicba.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (avicba *Resource) State() (*awsVpcIpv6CidrBlockAssociationState, bool) {
	return avicba.state, avicba.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (avicba *Resource) StateMust() *awsVpcIpv6CidrBlockAssociationState {
	if avicba.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", avicba.Type(), avicba.LocalName()))
	}
	return avicba.state
}

// Args contains the configurations for aws_vpc_ipv6_cidr_block_association.
type Args struct {
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
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsVpcIpv6CidrBlockAssociationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_vpc_ipv6_cidr_block_association.
func (avicba awsVpcIpv6CidrBlockAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(avicba.ref.Append("id"))
}

// Ipv6CidrBlock returns a reference to field ipv6_cidr_block of aws_vpc_ipv6_cidr_block_association.
func (avicba awsVpcIpv6CidrBlockAssociationAttributes) Ipv6CidrBlock() terra.StringValue {
	return terra.ReferenceAsString(avicba.ref.Append("ipv6_cidr_block"))
}

// Ipv6IpamPoolId returns a reference to field ipv6_ipam_pool_id of aws_vpc_ipv6_cidr_block_association.
func (avicba awsVpcIpv6CidrBlockAssociationAttributes) Ipv6IpamPoolId() terra.StringValue {
	return terra.ReferenceAsString(avicba.ref.Append("ipv6_ipam_pool_id"))
}

// Ipv6NetmaskLength returns a reference to field ipv6_netmask_length of aws_vpc_ipv6_cidr_block_association.
func (avicba awsVpcIpv6CidrBlockAssociationAttributes) Ipv6NetmaskLength() terra.NumberValue {
	return terra.ReferenceAsNumber(avicba.ref.Append("ipv6_netmask_length"))
}

// VpcId returns a reference to field vpc_id of aws_vpc_ipv6_cidr_block_association.
func (avicba awsVpcIpv6CidrBlockAssociationAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(avicba.ref.Append("vpc_id"))
}

func (avicba awsVpcIpv6CidrBlockAssociationAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](avicba.ref.Append("timeouts"))
}

type awsVpcIpv6CidrBlockAssociationState struct {
	Id                string         `json:"id"`
	Ipv6CidrBlock     string         `json:"ipv6_cidr_block"`
	Ipv6IpamPoolId    string         `json:"ipv6_ipam_pool_id"`
	Ipv6NetmaskLength float64        `json:"ipv6_netmask_length"`
	VpcId             string         `json:"vpc_id"`
	Timeouts          *TimeoutsState `json:"timeouts"`
}
