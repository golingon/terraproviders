// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVpnGatewayAttachment creates a new instance of [VpnGatewayAttachment].
func NewVpnGatewayAttachment(name string, args VpnGatewayAttachmentArgs) *VpnGatewayAttachment {
	return &VpnGatewayAttachment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpnGatewayAttachment)(nil)

// VpnGatewayAttachment represents the Terraform resource aws_vpn_gateway_attachment.
type VpnGatewayAttachment struct {
	Name      string
	Args      VpnGatewayAttachmentArgs
	state     *vpnGatewayAttachmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VpnGatewayAttachment].
func (vga *VpnGatewayAttachment) Type() string {
	return "aws_vpn_gateway_attachment"
}

// LocalName returns the local name for [VpnGatewayAttachment].
func (vga *VpnGatewayAttachment) LocalName() string {
	return vga.Name
}

// Configuration returns the configuration (args) for [VpnGatewayAttachment].
func (vga *VpnGatewayAttachment) Configuration() interface{} {
	return vga.Args
}

// DependOn is used for other resources to depend on [VpnGatewayAttachment].
func (vga *VpnGatewayAttachment) DependOn() terra.Reference {
	return terra.ReferenceResource(vga)
}

// Dependencies returns the list of resources [VpnGatewayAttachment] depends_on.
func (vga *VpnGatewayAttachment) Dependencies() terra.Dependencies {
	return vga.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VpnGatewayAttachment].
func (vga *VpnGatewayAttachment) LifecycleManagement() *terra.Lifecycle {
	return vga.Lifecycle
}

// Attributes returns the attributes for [VpnGatewayAttachment].
func (vga *VpnGatewayAttachment) Attributes() vpnGatewayAttachmentAttributes {
	return vpnGatewayAttachmentAttributes{ref: terra.ReferenceResource(vga)}
}

// ImportState imports the given attribute values into [VpnGatewayAttachment]'s state.
func (vga *VpnGatewayAttachment) ImportState(av io.Reader) error {
	vga.state = &vpnGatewayAttachmentState{}
	if err := json.NewDecoder(av).Decode(vga.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vga.Type(), vga.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VpnGatewayAttachment] has state.
func (vga *VpnGatewayAttachment) State() (*vpnGatewayAttachmentState, bool) {
	return vga.state, vga.state != nil
}

// StateMust returns the state for [VpnGatewayAttachment]. Panics if the state is nil.
func (vga *VpnGatewayAttachment) StateMust() *vpnGatewayAttachmentState {
	if vga.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vga.Type(), vga.LocalName()))
	}
	return vga.state
}

// VpnGatewayAttachmentArgs contains the configurations for aws_vpn_gateway_attachment.
type VpnGatewayAttachmentArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// VpcId: string, required
	VpcId terra.StringValue `hcl:"vpc_id,attr" validate:"required"`
	// VpnGatewayId: string, required
	VpnGatewayId terra.StringValue `hcl:"vpn_gateway_id,attr" validate:"required"`
}
type vpnGatewayAttachmentAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_vpn_gateway_attachment.
func (vga vpnGatewayAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vga.ref.Append("id"))
}

// VpcId returns a reference to field vpc_id of aws_vpn_gateway_attachment.
func (vga vpnGatewayAttachmentAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(vga.ref.Append("vpc_id"))
}

// VpnGatewayId returns a reference to field vpn_gateway_id of aws_vpn_gateway_attachment.
func (vga vpnGatewayAttachmentAttributes) VpnGatewayId() terra.StringValue {
	return terra.ReferenceAsString(vga.ref.Append("vpn_gateway_id"))
}

type vpnGatewayAttachmentState struct {
	Id           string `json:"id"`
	VpcId        string `json:"vpc_id"`
	VpnGatewayId string `json:"vpn_gateway_id"`
}
