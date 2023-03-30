// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewEc2TransitGatewayPeeringAttachment(name string, args Ec2TransitGatewayPeeringAttachmentArgs) *Ec2TransitGatewayPeeringAttachment {
	return &Ec2TransitGatewayPeeringAttachment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Ec2TransitGatewayPeeringAttachment)(nil)

type Ec2TransitGatewayPeeringAttachment struct {
	Name  string
	Args  Ec2TransitGatewayPeeringAttachmentArgs
	state *ec2TransitGatewayPeeringAttachmentState
}

func (etgpa *Ec2TransitGatewayPeeringAttachment) Type() string {
	return "aws_ec2_transit_gateway_peering_attachment"
}

func (etgpa *Ec2TransitGatewayPeeringAttachment) LocalName() string {
	return etgpa.Name
}

func (etgpa *Ec2TransitGatewayPeeringAttachment) Configuration() interface{} {
	return etgpa.Args
}

func (etgpa *Ec2TransitGatewayPeeringAttachment) Attributes() ec2TransitGatewayPeeringAttachmentAttributes {
	return ec2TransitGatewayPeeringAttachmentAttributes{ref: terra.ReferenceResource(etgpa)}
}

func (etgpa *Ec2TransitGatewayPeeringAttachment) ImportState(av io.Reader) error {
	etgpa.state = &ec2TransitGatewayPeeringAttachmentState{}
	if err := json.NewDecoder(av).Decode(etgpa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", etgpa.Type(), etgpa.LocalName(), err)
	}
	return nil
}

func (etgpa *Ec2TransitGatewayPeeringAttachment) State() (*ec2TransitGatewayPeeringAttachmentState, bool) {
	return etgpa.state, etgpa.state != nil
}

func (etgpa *Ec2TransitGatewayPeeringAttachment) StateMust() *ec2TransitGatewayPeeringAttachmentState {
	if etgpa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", etgpa.Type(), etgpa.LocalName()))
	}
	return etgpa.state
}

func (etgpa *Ec2TransitGatewayPeeringAttachment) DependOn() terra.Reference {
	return terra.ReferenceResource(etgpa)
}

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
	// DependsOn contains resources that Ec2TransitGatewayPeeringAttachment depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type ec2TransitGatewayPeeringAttachmentAttributes struct {
	ref terra.Reference
}

func (etgpa ec2TransitGatewayPeeringAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceString(etgpa.ref.Append("id"))
}

func (etgpa ec2TransitGatewayPeeringAttachmentAttributes) PeerAccountId() terra.StringValue {
	return terra.ReferenceString(etgpa.ref.Append("peer_account_id"))
}

func (etgpa ec2TransitGatewayPeeringAttachmentAttributes) PeerRegion() terra.StringValue {
	return terra.ReferenceString(etgpa.ref.Append("peer_region"))
}

func (etgpa ec2TransitGatewayPeeringAttachmentAttributes) PeerTransitGatewayId() terra.StringValue {
	return terra.ReferenceString(etgpa.ref.Append("peer_transit_gateway_id"))
}

func (etgpa ec2TransitGatewayPeeringAttachmentAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](etgpa.ref.Append("tags"))
}

func (etgpa ec2TransitGatewayPeeringAttachmentAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](etgpa.ref.Append("tags_all"))
}

func (etgpa ec2TransitGatewayPeeringAttachmentAttributes) TransitGatewayId() terra.StringValue {
	return terra.ReferenceString(etgpa.ref.Append("transit_gateway_id"))
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
