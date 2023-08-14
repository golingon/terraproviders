// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkInterfaceSgAttachment creates a new instance of [NetworkInterfaceSgAttachment].
func NewNetworkInterfaceSgAttachment(name string, args NetworkInterfaceSgAttachmentArgs) *NetworkInterfaceSgAttachment {
	return &NetworkInterfaceSgAttachment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkInterfaceSgAttachment)(nil)

// NetworkInterfaceSgAttachment represents the Terraform resource aws_network_interface_sg_attachment.
type NetworkInterfaceSgAttachment struct {
	Name      string
	Args      NetworkInterfaceSgAttachmentArgs
	state     *networkInterfaceSgAttachmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkInterfaceSgAttachment].
func (nisa *NetworkInterfaceSgAttachment) Type() string {
	return "aws_network_interface_sg_attachment"
}

// LocalName returns the local name for [NetworkInterfaceSgAttachment].
func (nisa *NetworkInterfaceSgAttachment) LocalName() string {
	return nisa.Name
}

// Configuration returns the configuration (args) for [NetworkInterfaceSgAttachment].
func (nisa *NetworkInterfaceSgAttachment) Configuration() interface{} {
	return nisa.Args
}

// DependOn is used for other resources to depend on [NetworkInterfaceSgAttachment].
func (nisa *NetworkInterfaceSgAttachment) DependOn() terra.Reference {
	return terra.ReferenceResource(nisa)
}

// Dependencies returns the list of resources [NetworkInterfaceSgAttachment] depends_on.
func (nisa *NetworkInterfaceSgAttachment) Dependencies() terra.Dependencies {
	return nisa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkInterfaceSgAttachment].
func (nisa *NetworkInterfaceSgAttachment) LifecycleManagement() *terra.Lifecycle {
	return nisa.Lifecycle
}

// Attributes returns the attributes for [NetworkInterfaceSgAttachment].
func (nisa *NetworkInterfaceSgAttachment) Attributes() networkInterfaceSgAttachmentAttributes {
	return networkInterfaceSgAttachmentAttributes{ref: terra.ReferenceResource(nisa)}
}

// ImportState imports the given attribute values into [NetworkInterfaceSgAttachment]'s state.
func (nisa *NetworkInterfaceSgAttachment) ImportState(av io.Reader) error {
	nisa.state = &networkInterfaceSgAttachmentState{}
	if err := json.NewDecoder(av).Decode(nisa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nisa.Type(), nisa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkInterfaceSgAttachment] has state.
func (nisa *NetworkInterfaceSgAttachment) State() (*networkInterfaceSgAttachmentState, bool) {
	return nisa.state, nisa.state != nil
}

// StateMust returns the state for [NetworkInterfaceSgAttachment]. Panics if the state is nil.
func (nisa *NetworkInterfaceSgAttachment) StateMust() *networkInterfaceSgAttachmentState {
	if nisa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nisa.Type(), nisa.LocalName()))
	}
	return nisa.state
}

// NetworkInterfaceSgAttachmentArgs contains the configurations for aws_network_interface_sg_attachment.
type NetworkInterfaceSgAttachmentArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NetworkInterfaceId: string, required
	NetworkInterfaceId terra.StringValue `hcl:"network_interface_id,attr" validate:"required"`
	// SecurityGroupId: string, required
	SecurityGroupId terra.StringValue `hcl:"security_group_id,attr" validate:"required"`
}
type networkInterfaceSgAttachmentAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_network_interface_sg_attachment.
func (nisa networkInterfaceSgAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nisa.ref.Append("id"))
}

// NetworkInterfaceId returns a reference to field network_interface_id of aws_network_interface_sg_attachment.
func (nisa networkInterfaceSgAttachmentAttributes) NetworkInterfaceId() terra.StringValue {
	return terra.ReferenceAsString(nisa.ref.Append("network_interface_id"))
}

// SecurityGroupId returns a reference to field security_group_id of aws_network_interface_sg_attachment.
func (nisa networkInterfaceSgAttachmentAttributes) SecurityGroupId() terra.StringValue {
	return terra.ReferenceAsString(nisa.ref.Append("security_group_id"))
}

type networkInterfaceSgAttachmentState struct {
	Id                 string `json:"id"`
	NetworkInterfaceId string `json:"network_interface_id"`
	SecurityGroupId    string `json:"security_group_id"`
}