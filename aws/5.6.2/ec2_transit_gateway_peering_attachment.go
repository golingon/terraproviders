// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEc2TransitGatewayPeeringAttachment creates a new instance of [Ec2TransitGatewayPeeringAttachment].
func NewEc2TransitGatewayPeeringAttachment(name string, args Ec2TransitGatewayPeeringAttachmentArgs) *Ec2TransitGatewayPeeringAttachment {
	return &Ec2TransitGatewayPeeringAttachment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Ec2TransitGatewayPeeringAttachment)(nil)

// Ec2TransitGatewayPeeringAttachment represents the Terraform resource aws_ec2_transit_gateway_peering_attachment.
type Ec2TransitGatewayPeeringAttachment struct {
	Name      string
	Args      Ec2TransitGatewayPeeringAttachmentArgs
	state     *ec2TransitGatewayPeeringAttachmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Ec2TransitGatewayPeeringAttachment].
func (etgpa *Ec2TransitGatewayPeeringAttachment) Type() string {
	return "aws_ec2_transit_gateway_peering_attachment"
}

// LocalName returns the local name for [Ec2TransitGatewayPeeringAttachment].
func (etgpa *Ec2TransitGatewayPeeringAttachment) LocalName() string {
	return etgpa.Name
}

// Configuration returns the configuration (args) for [Ec2TransitGatewayPeeringAttachment].
func (etgpa *Ec2TransitGatewayPeeringAttachment) Configuration() interface{} {
	return etgpa.Args
}

// DependOn is used for other resources to depend on [Ec2TransitGatewayPeeringAttachment].
func (etgpa *Ec2TransitGatewayPeeringAttachment) DependOn() terra.Reference {
	return terra.ReferenceResource(etgpa)
}

// Dependencies returns the list of resources [Ec2TransitGatewayPeeringAttachment] depends_on.
func (etgpa *Ec2TransitGatewayPeeringAttachment) Dependencies() terra.Dependencies {
	return etgpa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Ec2TransitGatewayPeeringAttachment].
func (etgpa *Ec2TransitGatewayPeeringAttachment) LifecycleManagement() *terra.Lifecycle {
	return etgpa.Lifecycle
}

// Attributes returns the attributes for [Ec2TransitGatewayPeeringAttachment].
func (etgpa *Ec2TransitGatewayPeeringAttachment) Attributes() ec2TransitGatewayPeeringAttachmentAttributes {
	return ec2TransitGatewayPeeringAttachmentAttributes{ref: terra.ReferenceResource(etgpa)}
}

// ImportState imports the given attribute values into [Ec2TransitGatewayPeeringAttachment]'s state.
func (etgpa *Ec2TransitGatewayPeeringAttachment) ImportState(av io.Reader) error {
	etgpa.state = &ec2TransitGatewayPeeringAttachmentState{}
	if err := json.NewDecoder(av).Decode(etgpa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", etgpa.Type(), etgpa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Ec2TransitGatewayPeeringAttachment] has state.
func (etgpa *Ec2TransitGatewayPeeringAttachment) State() (*ec2TransitGatewayPeeringAttachmentState, bool) {
	return etgpa.state, etgpa.state != nil
}

// StateMust returns the state for [Ec2TransitGatewayPeeringAttachment]. Panics if the state is nil.
func (etgpa *Ec2TransitGatewayPeeringAttachment) StateMust() *ec2TransitGatewayPeeringAttachmentState {
	if etgpa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", etgpa.Type(), etgpa.LocalName()))
	}
	return etgpa.state
}

// Ec2TransitGatewayPeeringAttachmentArgs contains the configurations for aws_ec2_transit_gateway_peering_attachment.
type Ec2TransitGatewayPeeringAttachmentArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PeerAccountId: string, optional
	PeerAccountId terra.StringValue `hcl:"peer_account_id,attr"`
	// PeerRegion: string, required
	PeerRegion terra.StringValue `hcl:"peer_region,attr" validate:"required"`
	// PeerTransitGatewayId: string, required
	PeerTransitGatewayId terra.StringValue `hcl:"peer_transit_gateway_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TransitGatewayId: string, required
	TransitGatewayId terra.StringValue `hcl:"transit_gateway_id,attr" validate:"required"`
}
type ec2TransitGatewayPeeringAttachmentAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ec2_transit_gateway_peering_attachment.
func (etgpa ec2TransitGatewayPeeringAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(etgpa.ref.Append("id"))
}

// PeerAccountId returns a reference to field peer_account_id of aws_ec2_transit_gateway_peering_attachment.
func (etgpa ec2TransitGatewayPeeringAttachmentAttributes) PeerAccountId() terra.StringValue {
	return terra.ReferenceAsString(etgpa.ref.Append("peer_account_id"))
}

// PeerRegion returns a reference to field peer_region of aws_ec2_transit_gateway_peering_attachment.
func (etgpa ec2TransitGatewayPeeringAttachmentAttributes) PeerRegion() terra.StringValue {
	return terra.ReferenceAsString(etgpa.ref.Append("peer_region"))
}

// PeerTransitGatewayId returns a reference to field peer_transit_gateway_id of aws_ec2_transit_gateway_peering_attachment.
func (etgpa ec2TransitGatewayPeeringAttachmentAttributes) PeerTransitGatewayId() terra.StringValue {
	return terra.ReferenceAsString(etgpa.ref.Append("peer_transit_gateway_id"))
}

// Tags returns a reference to field tags of aws_ec2_transit_gateway_peering_attachment.
func (etgpa ec2TransitGatewayPeeringAttachmentAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](etgpa.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_ec2_transit_gateway_peering_attachment.
func (etgpa ec2TransitGatewayPeeringAttachmentAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](etgpa.ref.Append("tags_all"))
}

// TransitGatewayId returns a reference to field transit_gateway_id of aws_ec2_transit_gateway_peering_attachment.
func (etgpa ec2TransitGatewayPeeringAttachmentAttributes) TransitGatewayId() terra.StringValue {
	return terra.ReferenceAsString(etgpa.ref.Append("transit_gateway_id"))
}

type ec2TransitGatewayPeeringAttachmentState struct {
	Id                   string            `json:"id"`
	PeerAccountId        string            `json:"peer_account_id"`
	PeerRegion           string            `json:"peer_region"`
	PeerTransitGatewayId string            `json:"peer_transit_gateway_id"`
	Tags                 map[string]string `json:"tags"`
	TagsAll              map[string]string `json:"tags_all"`
	TransitGatewayId     string            `json:"transit_gateway_id"`
}