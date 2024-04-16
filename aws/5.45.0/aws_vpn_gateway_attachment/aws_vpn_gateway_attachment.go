// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_vpn_gateway_attachment

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

// Resource represents the Terraform resource aws_vpn_gateway_attachment.
type Resource struct {
	Name      string
	Args      Args
	state     *awsVpnGatewayAttachmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (avga *Resource) Type() string {
	return "aws_vpn_gateway_attachment"
}

// LocalName returns the local name for [Resource].
func (avga *Resource) LocalName() string {
	return avga.Name
}

// Configuration returns the configuration (args) for [Resource].
func (avga *Resource) Configuration() interface{} {
	return avga.Args
}

// DependOn is used for other resources to depend on [Resource].
func (avga *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(avga)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (avga *Resource) Dependencies() terra.Dependencies {
	return avga.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (avga *Resource) LifecycleManagement() *terra.Lifecycle {
	return avga.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (avga *Resource) Attributes() awsVpnGatewayAttachmentAttributes {
	return awsVpnGatewayAttachmentAttributes{ref: terra.ReferenceResource(avga)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (avga *Resource) ImportState(state io.Reader) error {
	avga.state = &awsVpnGatewayAttachmentState{}
	if err := json.NewDecoder(state).Decode(avga.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", avga.Type(), avga.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (avga *Resource) State() (*awsVpnGatewayAttachmentState, bool) {
	return avga.state, avga.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (avga *Resource) StateMust() *awsVpnGatewayAttachmentState {
	if avga.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", avga.Type(), avga.LocalName()))
	}
	return avga.state
}

// Args contains the configurations for aws_vpn_gateway_attachment.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// VpcId: string, required
	VpcId terra.StringValue `hcl:"vpc_id,attr" validate:"required"`
	// VpnGatewayId: string, required
	VpnGatewayId terra.StringValue `hcl:"vpn_gateway_id,attr" validate:"required"`
}

type awsVpnGatewayAttachmentAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_vpn_gateway_attachment.
func (avga awsVpnGatewayAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(avga.ref.Append("id"))
}

// VpcId returns a reference to field vpc_id of aws_vpn_gateway_attachment.
func (avga awsVpnGatewayAttachmentAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(avga.ref.Append("vpc_id"))
}

// VpnGatewayId returns a reference to field vpn_gateway_id of aws_vpn_gateway_attachment.
func (avga awsVpnGatewayAttachmentAttributes) VpnGatewayId() terra.StringValue {
	return terra.ReferenceAsString(avga.ref.Append("vpn_gateway_id"))
}

type awsVpnGatewayAttachmentState struct {
	Id           string `json:"id"`
	VpcId        string `json:"vpc_id"`
	VpnGatewayId string `json:"vpn_gateway_id"`
}
