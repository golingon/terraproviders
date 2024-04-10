// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewEc2TransitGatewayPeeringAttachmentAccepter creates a new instance of [Ec2TransitGatewayPeeringAttachmentAccepter].
func NewEc2TransitGatewayPeeringAttachmentAccepter(name string, args Ec2TransitGatewayPeeringAttachmentAccepterArgs) *Ec2TransitGatewayPeeringAttachmentAccepter {
	return &Ec2TransitGatewayPeeringAttachmentAccepter{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Ec2TransitGatewayPeeringAttachmentAccepter)(nil)

// Ec2TransitGatewayPeeringAttachmentAccepter represents the Terraform resource aws_ec2_transit_gateway_peering_attachment_accepter.
type Ec2TransitGatewayPeeringAttachmentAccepter struct {
	Name      string
	Args      Ec2TransitGatewayPeeringAttachmentAccepterArgs
	state     *ec2TransitGatewayPeeringAttachmentAccepterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Ec2TransitGatewayPeeringAttachmentAccepter].
func (etgpaa *Ec2TransitGatewayPeeringAttachmentAccepter) Type() string {
	return "aws_ec2_transit_gateway_peering_attachment_accepter"
}

// LocalName returns the local name for [Ec2TransitGatewayPeeringAttachmentAccepter].
func (etgpaa *Ec2TransitGatewayPeeringAttachmentAccepter) LocalName() string {
	return etgpaa.Name
}

// Configuration returns the configuration (args) for [Ec2TransitGatewayPeeringAttachmentAccepter].
func (etgpaa *Ec2TransitGatewayPeeringAttachmentAccepter) Configuration() interface{} {
	return etgpaa.Args
}

// DependOn is used for other resources to depend on [Ec2TransitGatewayPeeringAttachmentAccepter].
func (etgpaa *Ec2TransitGatewayPeeringAttachmentAccepter) DependOn() terra.Reference {
	return terra.ReferenceResource(etgpaa)
}

// Dependencies returns the list of resources [Ec2TransitGatewayPeeringAttachmentAccepter] depends_on.
func (etgpaa *Ec2TransitGatewayPeeringAttachmentAccepter) Dependencies() terra.Dependencies {
	return etgpaa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Ec2TransitGatewayPeeringAttachmentAccepter].
func (etgpaa *Ec2TransitGatewayPeeringAttachmentAccepter) LifecycleManagement() *terra.Lifecycle {
	return etgpaa.Lifecycle
}

// Attributes returns the attributes for [Ec2TransitGatewayPeeringAttachmentAccepter].
func (etgpaa *Ec2TransitGatewayPeeringAttachmentAccepter) Attributes() ec2TransitGatewayPeeringAttachmentAccepterAttributes {
	return ec2TransitGatewayPeeringAttachmentAccepterAttributes{ref: terra.ReferenceResource(etgpaa)}
}

// ImportState imports the given attribute values into [Ec2TransitGatewayPeeringAttachmentAccepter]'s state.
func (etgpaa *Ec2TransitGatewayPeeringAttachmentAccepter) ImportState(av io.Reader) error {
	etgpaa.state = &ec2TransitGatewayPeeringAttachmentAccepterState{}
	if err := json.NewDecoder(av).Decode(etgpaa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", etgpaa.Type(), etgpaa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Ec2TransitGatewayPeeringAttachmentAccepter] has state.
func (etgpaa *Ec2TransitGatewayPeeringAttachmentAccepter) State() (*ec2TransitGatewayPeeringAttachmentAccepterState, bool) {
	return etgpaa.state, etgpaa.state != nil
}

// StateMust returns the state for [Ec2TransitGatewayPeeringAttachmentAccepter]. Panics if the state is nil.
func (etgpaa *Ec2TransitGatewayPeeringAttachmentAccepter) StateMust() *ec2TransitGatewayPeeringAttachmentAccepterState {
	if etgpaa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", etgpaa.Type(), etgpaa.LocalName()))
	}
	return etgpaa.state
}

// Ec2TransitGatewayPeeringAttachmentAccepterArgs contains the configurations for aws_ec2_transit_gateway_peering_attachment_accepter.
type Ec2TransitGatewayPeeringAttachmentAccepterArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TransitGatewayAttachmentId: string, required
	TransitGatewayAttachmentId terra.StringValue `hcl:"transit_gateway_attachment_id,attr" validate:"required"`
}
type ec2TransitGatewayPeeringAttachmentAccepterAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ec2_transit_gateway_peering_attachment_accepter.
func (etgpaa ec2TransitGatewayPeeringAttachmentAccepterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(etgpaa.ref.Append("id"))
}

// PeerAccountId returns a reference to field peer_account_id of aws_ec2_transit_gateway_peering_attachment_accepter.
func (etgpaa ec2TransitGatewayPeeringAttachmentAccepterAttributes) PeerAccountId() terra.StringValue {
	return terra.ReferenceAsString(etgpaa.ref.Append("peer_account_id"))
}

// PeerRegion returns a reference to field peer_region of aws_ec2_transit_gateway_peering_attachment_accepter.
func (etgpaa ec2TransitGatewayPeeringAttachmentAccepterAttributes) PeerRegion() terra.StringValue {
	return terra.ReferenceAsString(etgpaa.ref.Append("peer_region"))
}

// PeerTransitGatewayId returns a reference to field peer_transit_gateway_id of aws_ec2_transit_gateway_peering_attachment_accepter.
func (etgpaa ec2TransitGatewayPeeringAttachmentAccepterAttributes) PeerTransitGatewayId() terra.StringValue {
	return terra.ReferenceAsString(etgpaa.ref.Append("peer_transit_gateway_id"))
}

// Tags returns a reference to field tags of aws_ec2_transit_gateway_peering_attachment_accepter.
func (etgpaa ec2TransitGatewayPeeringAttachmentAccepterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](etgpaa.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_ec2_transit_gateway_peering_attachment_accepter.
func (etgpaa ec2TransitGatewayPeeringAttachmentAccepterAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](etgpaa.ref.Append("tags_all"))
}

// TransitGatewayAttachmentId returns a reference to field transit_gateway_attachment_id of aws_ec2_transit_gateway_peering_attachment_accepter.
func (etgpaa ec2TransitGatewayPeeringAttachmentAccepterAttributes) TransitGatewayAttachmentId() terra.StringValue {
	return terra.ReferenceAsString(etgpaa.ref.Append("transit_gateway_attachment_id"))
}

// TransitGatewayId returns a reference to field transit_gateway_id of aws_ec2_transit_gateway_peering_attachment_accepter.
func (etgpaa ec2TransitGatewayPeeringAttachmentAccepterAttributes) TransitGatewayId() terra.StringValue {
	return terra.ReferenceAsString(etgpaa.ref.Append("transit_gateway_id"))
}

type ec2TransitGatewayPeeringAttachmentAccepterState struct {
	Id                         string            `json:"id"`
	PeerAccountId              string            `json:"peer_account_id"`
	PeerRegion                 string            `json:"peer_region"`
	PeerTransitGatewayId       string            `json:"peer_transit_gateway_id"`
	Tags                       map[string]string `json:"tags"`
	TagsAll                    map[string]string `json:"tags_all"`
	TransitGatewayAttachmentId string            `json:"transit_gateway_attachment_id"`
	TransitGatewayId           string            `json:"transit_gateway_id"`
}
