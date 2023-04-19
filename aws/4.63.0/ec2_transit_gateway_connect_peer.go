// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ec2transitgatewayconnectpeer "github.com/golingon/terraproviders/aws/4.63.0/ec2transitgatewayconnectpeer"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEc2TransitGatewayConnectPeer creates a new instance of [Ec2TransitGatewayConnectPeer].
func NewEc2TransitGatewayConnectPeer(name string, args Ec2TransitGatewayConnectPeerArgs) *Ec2TransitGatewayConnectPeer {
	return &Ec2TransitGatewayConnectPeer{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Ec2TransitGatewayConnectPeer)(nil)

// Ec2TransitGatewayConnectPeer represents the Terraform resource aws_ec2_transit_gateway_connect_peer.
type Ec2TransitGatewayConnectPeer struct {
	Name      string
	Args      Ec2TransitGatewayConnectPeerArgs
	state     *ec2TransitGatewayConnectPeerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Ec2TransitGatewayConnectPeer].
func (etgcp *Ec2TransitGatewayConnectPeer) Type() string {
	return "aws_ec2_transit_gateway_connect_peer"
}

// LocalName returns the local name for [Ec2TransitGatewayConnectPeer].
func (etgcp *Ec2TransitGatewayConnectPeer) LocalName() string {
	return etgcp.Name
}

// Configuration returns the configuration (args) for [Ec2TransitGatewayConnectPeer].
func (etgcp *Ec2TransitGatewayConnectPeer) Configuration() interface{} {
	return etgcp.Args
}

// DependOn is used for other resources to depend on [Ec2TransitGatewayConnectPeer].
func (etgcp *Ec2TransitGatewayConnectPeer) DependOn() terra.Reference {
	return terra.ReferenceResource(etgcp)
}

// Dependencies returns the list of resources [Ec2TransitGatewayConnectPeer] depends_on.
func (etgcp *Ec2TransitGatewayConnectPeer) Dependencies() terra.Dependencies {
	return etgcp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Ec2TransitGatewayConnectPeer].
func (etgcp *Ec2TransitGatewayConnectPeer) LifecycleManagement() *terra.Lifecycle {
	return etgcp.Lifecycle
}

// Attributes returns the attributes for [Ec2TransitGatewayConnectPeer].
func (etgcp *Ec2TransitGatewayConnectPeer) Attributes() ec2TransitGatewayConnectPeerAttributes {
	return ec2TransitGatewayConnectPeerAttributes{ref: terra.ReferenceResource(etgcp)}
}

// ImportState imports the given attribute values into [Ec2TransitGatewayConnectPeer]'s state.
func (etgcp *Ec2TransitGatewayConnectPeer) ImportState(av io.Reader) error {
	etgcp.state = &ec2TransitGatewayConnectPeerState{}
	if err := json.NewDecoder(av).Decode(etgcp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", etgcp.Type(), etgcp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Ec2TransitGatewayConnectPeer] has state.
func (etgcp *Ec2TransitGatewayConnectPeer) State() (*ec2TransitGatewayConnectPeerState, bool) {
	return etgcp.state, etgcp.state != nil
}

// StateMust returns the state for [Ec2TransitGatewayConnectPeer]. Panics if the state is nil.
func (etgcp *Ec2TransitGatewayConnectPeer) StateMust() *ec2TransitGatewayConnectPeerState {
	if etgcp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", etgcp.Type(), etgcp.LocalName()))
	}
	return etgcp.state
}

// Ec2TransitGatewayConnectPeerArgs contains the configurations for aws_ec2_transit_gateway_connect_peer.
type Ec2TransitGatewayConnectPeerArgs struct {
	// BgpAsn: string, optional
	BgpAsn terra.StringValue `hcl:"bgp_asn,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InsideCidrBlocks: set of string, required
	InsideCidrBlocks terra.SetValue[terra.StringValue] `hcl:"inside_cidr_blocks,attr" validate:"required"`
	// PeerAddress: string, required
	PeerAddress terra.StringValue `hcl:"peer_address,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TransitGatewayAddress: string, optional
	TransitGatewayAddress terra.StringValue `hcl:"transit_gateway_address,attr"`
	// TransitGatewayAttachmentId: string, required
	TransitGatewayAttachmentId terra.StringValue `hcl:"transit_gateway_attachment_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *ec2transitgatewayconnectpeer.Timeouts `hcl:"timeouts,block"`
}
type ec2TransitGatewayConnectPeerAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ec2_transit_gateway_connect_peer.
func (etgcp ec2TransitGatewayConnectPeerAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(etgcp.ref.Append("arn"))
}

// BgpAsn returns a reference to field bgp_asn of aws_ec2_transit_gateway_connect_peer.
func (etgcp ec2TransitGatewayConnectPeerAttributes) BgpAsn() terra.StringValue {
	return terra.ReferenceAsString(etgcp.ref.Append("bgp_asn"))
}

// Id returns a reference to field id of aws_ec2_transit_gateway_connect_peer.
func (etgcp ec2TransitGatewayConnectPeerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(etgcp.ref.Append("id"))
}

// InsideCidrBlocks returns a reference to field inside_cidr_blocks of aws_ec2_transit_gateway_connect_peer.
func (etgcp ec2TransitGatewayConnectPeerAttributes) InsideCidrBlocks() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](etgcp.ref.Append("inside_cidr_blocks"))
}

// PeerAddress returns a reference to field peer_address of aws_ec2_transit_gateway_connect_peer.
func (etgcp ec2TransitGatewayConnectPeerAttributes) PeerAddress() terra.StringValue {
	return terra.ReferenceAsString(etgcp.ref.Append("peer_address"))
}

// Tags returns a reference to field tags of aws_ec2_transit_gateway_connect_peer.
func (etgcp ec2TransitGatewayConnectPeerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](etgcp.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_ec2_transit_gateway_connect_peer.
func (etgcp ec2TransitGatewayConnectPeerAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](etgcp.ref.Append("tags_all"))
}

// TransitGatewayAddress returns a reference to field transit_gateway_address of aws_ec2_transit_gateway_connect_peer.
func (etgcp ec2TransitGatewayConnectPeerAttributes) TransitGatewayAddress() terra.StringValue {
	return terra.ReferenceAsString(etgcp.ref.Append("transit_gateway_address"))
}

// TransitGatewayAttachmentId returns a reference to field transit_gateway_attachment_id of aws_ec2_transit_gateway_connect_peer.
func (etgcp ec2TransitGatewayConnectPeerAttributes) TransitGatewayAttachmentId() terra.StringValue {
	return terra.ReferenceAsString(etgcp.ref.Append("transit_gateway_attachment_id"))
}

func (etgcp ec2TransitGatewayConnectPeerAttributes) Timeouts() ec2transitgatewayconnectpeer.TimeoutsAttributes {
	return terra.ReferenceAsSingle[ec2transitgatewayconnectpeer.TimeoutsAttributes](etgcp.ref.Append("timeouts"))
}

type ec2TransitGatewayConnectPeerState struct {
	Arn                        string                                      `json:"arn"`
	BgpAsn                     string                                      `json:"bgp_asn"`
	Id                         string                                      `json:"id"`
	InsideCidrBlocks           []string                                    `json:"inside_cidr_blocks"`
	PeerAddress                string                                      `json:"peer_address"`
	Tags                       map[string]string                           `json:"tags"`
	TagsAll                    map[string]string                           `json:"tags_all"`
	TransitGatewayAddress      string                                      `json:"transit_gateway_address"`
	TransitGatewayAttachmentId string                                      `json:"transit_gateway_attachment_id"`
	Timeouts                   *ec2transitgatewayconnectpeer.TimeoutsState `json:"timeouts"`
}
