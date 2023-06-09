// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEc2TransitGatewayMulticastGroupMember creates a new instance of [Ec2TransitGatewayMulticastGroupMember].
func NewEc2TransitGatewayMulticastGroupMember(name string, args Ec2TransitGatewayMulticastGroupMemberArgs) *Ec2TransitGatewayMulticastGroupMember {
	return &Ec2TransitGatewayMulticastGroupMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Ec2TransitGatewayMulticastGroupMember)(nil)

// Ec2TransitGatewayMulticastGroupMember represents the Terraform resource aws_ec2_transit_gateway_multicast_group_member.
type Ec2TransitGatewayMulticastGroupMember struct {
	Name      string
	Args      Ec2TransitGatewayMulticastGroupMemberArgs
	state     *ec2TransitGatewayMulticastGroupMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Ec2TransitGatewayMulticastGroupMember].
func (etgmgm *Ec2TransitGatewayMulticastGroupMember) Type() string {
	return "aws_ec2_transit_gateway_multicast_group_member"
}

// LocalName returns the local name for [Ec2TransitGatewayMulticastGroupMember].
func (etgmgm *Ec2TransitGatewayMulticastGroupMember) LocalName() string {
	return etgmgm.Name
}

// Configuration returns the configuration (args) for [Ec2TransitGatewayMulticastGroupMember].
func (etgmgm *Ec2TransitGatewayMulticastGroupMember) Configuration() interface{} {
	return etgmgm.Args
}

// DependOn is used for other resources to depend on [Ec2TransitGatewayMulticastGroupMember].
func (etgmgm *Ec2TransitGatewayMulticastGroupMember) DependOn() terra.Reference {
	return terra.ReferenceResource(etgmgm)
}

// Dependencies returns the list of resources [Ec2TransitGatewayMulticastGroupMember] depends_on.
func (etgmgm *Ec2TransitGatewayMulticastGroupMember) Dependencies() terra.Dependencies {
	return etgmgm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Ec2TransitGatewayMulticastGroupMember].
func (etgmgm *Ec2TransitGatewayMulticastGroupMember) LifecycleManagement() *terra.Lifecycle {
	return etgmgm.Lifecycle
}

// Attributes returns the attributes for [Ec2TransitGatewayMulticastGroupMember].
func (etgmgm *Ec2TransitGatewayMulticastGroupMember) Attributes() ec2TransitGatewayMulticastGroupMemberAttributes {
	return ec2TransitGatewayMulticastGroupMemberAttributes{ref: terra.ReferenceResource(etgmgm)}
}

// ImportState imports the given attribute values into [Ec2TransitGatewayMulticastGroupMember]'s state.
func (etgmgm *Ec2TransitGatewayMulticastGroupMember) ImportState(av io.Reader) error {
	etgmgm.state = &ec2TransitGatewayMulticastGroupMemberState{}
	if err := json.NewDecoder(av).Decode(etgmgm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", etgmgm.Type(), etgmgm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Ec2TransitGatewayMulticastGroupMember] has state.
func (etgmgm *Ec2TransitGatewayMulticastGroupMember) State() (*ec2TransitGatewayMulticastGroupMemberState, bool) {
	return etgmgm.state, etgmgm.state != nil
}

// StateMust returns the state for [Ec2TransitGatewayMulticastGroupMember]. Panics if the state is nil.
func (etgmgm *Ec2TransitGatewayMulticastGroupMember) StateMust() *ec2TransitGatewayMulticastGroupMemberState {
	if etgmgm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", etgmgm.Type(), etgmgm.LocalName()))
	}
	return etgmgm.state
}

// Ec2TransitGatewayMulticastGroupMemberArgs contains the configurations for aws_ec2_transit_gateway_multicast_group_member.
type Ec2TransitGatewayMulticastGroupMemberArgs struct {
	// GroupIpAddress: string, required
	GroupIpAddress terra.StringValue `hcl:"group_ip_address,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NetworkInterfaceId: string, required
	NetworkInterfaceId terra.StringValue `hcl:"network_interface_id,attr" validate:"required"`
	// TransitGatewayMulticastDomainId: string, required
	TransitGatewayMulticastDomainId terra.StringValue `hcl:"transit_gateway_multicast_domain_id,attr" validate:"required"`
}
type ec2TransitGatewayMulticastGroupMemberAttributes struct {
	ref terra.Reference
}

// GroupIpAddress returns a reference to field group_ip_address of aws_ec2_transit_gateway_multicast_group_member.
func (etgmgm ec2TransitGatewayMulticastGroupMemberAttributes) GroupIpAddress() terra.StringValue {
	return terra.ReferenceAsString(etgmgm.ref.Append("group_ip_address"))
}

// Id returns a reference to field id of aws_ec2_transit_gateway_multicast_group_member.
func (etgmgm ec2TransitGatewayMulticastGroupMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(etgmgm.ref.Append("id"))
}

// NetworkInterfaceId returns a reference to field network_interface_id of aws_ec2_transit_gateway_multicast_group_member.
func (etgmgm ec2TransitGatewayMulticastGroupMemberAttributes) NetworkInterfaceId() terra.StringValue {
	return terra.ReferenceAsString(etgmgm.ref.Append("network_interface_id"))
}

// TransitGatewayMulticastDomainId returns a reference to field transit_gateway_multicast_domain_id of aws_ec2_transit_gateway_multicast_group_member.
func (etgmgm ec2TransitGatewayMulticastGroupMemberAttributes) TransitGatewayMulticastDomainId() terra.StringValue {
	return terra.ReferenceAsString(etgmgm.ref.Append("transit_gateway_multicast_domain_id"))
}

type ec2TransitGatewayMulticastGroupMemberState struct {
	GroupIpAddress                  string `json:"group_ip_address"`
	Id                              string `json:"id"`
	NetworkInterfaceId              string `json:"network_interface_id"`
	TransitGatewayMulticastDomainId string `json:"transit_gateway_multicast_domain_id"`
}
