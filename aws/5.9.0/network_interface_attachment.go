// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkInterfaceAttachment creates a new instance of [NetworkInterfaceAttachment].
func NewNetworkInterfaceAttachment(name string, args NetworkInterfaceAttachmentArgs) *NetworkInterfaceAttachment {
	return &NetworkInterfaceAttachment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkInterfaceAttachment)(nil)

// NetworkInterfaceAttachment represents the Terraform resource aws_network_interface_attachment.
type NetworkInterfaceAttachment struct {
	Name      string
	Args      NetworkInterfaceAttachmentArgs
	state     *networkInterfaceAttachmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkInterfaceAttachment].
func (nia *NetworkInterfaceAttachment) Type() string {
	return "aws_network_interface_attachment"
}

// LocalName returns the local name for [NetworkInterfaceAttachment].
func (nia *NetworkInterfaceAttachment) LocalName() string {
	return nia.Name
}

// Configuration returns the configuration (args) for [NetworkInterfaceAttachment].
func (nia *NetworkInterfaceAttachment) Configuration() interface{} {
	return nia.Args
}

// DependOn is used for other resources to depend on [NetworkInterfaceAttachment].
func (nia *NetworkInterfaceAttachment) DependOn() terra.Reference {
	return terra.ReferenceResource(nia)
}

// Dependencies returns the list of resources [NetworkInterfaceAttachment] depends_on.
func (nia *NetworkInterfaceAttachment) Dependencies() terra.Dependencies {
	return nia.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkInterfaceAttachment].
func (nia *NetworkInterfaceAttachment) LifecycleManagement() *terra.Lifecycle {
	return nia.Lifecycle
}

// Attributes returns the attributes for [NetworkInterfaceAttachment].
func (nia *NetworkInterfaceAttachment) Attributes() networkInterfaceAttachmentAttributes {
	return networkInterfaceAttachmentAttributes{ref: terra.ReferenceResource(nia)}
}

// ImportState imports the given attribute values into [NetworkInterfaceAttachment]'s state.
func (nia *NetworkInterfaceAttachment) ImportState(av io.Reader) error {
	nia.state = &networkInterfaceAttachmentState{}
	if err := json.NewDecoder(av).Decode(nia.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nia.Type(), nia.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkInterfaceAttachment] has state.
func (nia *NetworkInterfaceAttachment) State() (*networkInterfaceAttachmentState, bool) {
	return nia.state, nia.state != nil
}

// StateMust returns the state for [NetworkInterfaceAttachment]. Panics if the state is nil.
func (nia *NetworkInterfaceAttachment) StateMust() *networkInterfaceAttachmentState {
	if nia.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nia.Type(), nia.LocalName()))
	}
	return nia.state
}

// NetworkInterfaceAttachmentArgs contains the configurations for aws_network_interface_attachment.
type NetworkInterfaceAttachmentArgs struct {
	// DeviceIndex: number, required
	DeviceIndex terra.NumberValue `hcl:"device_index,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceId: string, required
	InstanceId terra.StringValue `hcl:"instance_id,attr" validate:"required"`
	// NetworkInterfaceId: string, required
	NetworkInterfaceId terra.StringValue `hcl:"network_interface_id,attr" validate:"required"`
}
type networkInterfaceAttachmentAttributes struct {
	ref terra.Reference
}

// AttachmentId returns a reference to field attachment_id of aws_network_interface_attachment.
func (nia networkInterfaceAttachmentAttributes) AttachmentId() terra.StringValue {
	return terra.ReferenceAsString(nia.ref.Append("attachment_id"))
}

// DeviceIndex returns a reference to field device_index of aws_network_interface_attachment.
func (nia networkInterfaceAttachmentAttributes) DeviceIndex() terra.NumberValue {
	return terra.ReferenceAsNumber(nia.ref.Append("device_index"))
}

// Id returns a reference to field id of aws_network_interface_attachment.
func (nia networkInterfaceAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nia.ref.Append("id"))
}

// InstanceId returns a reference to field instance_id of aws_network_interface_attachment.
func (nia networkInterfaceAttachmentAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(nia.ref.Append("instance_id"))
}

// NetworkInterfaceId returns a reference to field network_interface_id of aws_network_interface_attachment.
func (nia networkInterfaceAttachmentAttributes) NetworkInterfaceId() terra.StringValue {
	return terra.ReferenceAsString(nia.ref.Append("network_interface_id"))
}

// Status returns a reference to field status of aws_network_interface_attachment.
func (nia networkInterfaceAttachmentAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(nia.ref.Append("status"))
}

type networkInterfaceAttachmentState struct {
	AttachmentId       string  `json:"attachment_id"`
	DeviceIndex        float64 `json:"device_index"`
	Id                 string  `json:"id"`
	InstanceId         string  `json:"instance_id"`
	NetworkInterfaceId string  `json:"network_interface_id"`
	Status             string  `json:"status"`
}