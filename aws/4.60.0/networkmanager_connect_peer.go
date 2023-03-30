// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	networkmanagerconnectpeer "github.com/golingon/terraproviders/aws/4.60.0/networkmanagerconnectpeer"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewNetworkmanagerConnectPeer(name string, args NetworkmanagerConnectPeerArgs) *NetworkmanagerConnectPeer {
	return &NetworkmanagerConnectPeer{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkmanagerConnectPeer)(nil)

type NetworkmanagerConnectPeer struct {
	Name  string
	Args  NetworkmanagerConnectPeerArgs
	state *networkmanagerConnectPeerState
}

func (ncp *NetworkmanagerConnectPeer) Type() string {
	return "aws_networkmanager_connect_peer"
}

func (ncp *NetworkmanagerConnectPeer) LocalName() string {
	return ncp.Name
}

func (ncp *NetworkmanagerConnectPeer) Configuration() interface{} {
	return ncp.Args
}

func (ncp *NetworkmanagerConnectPeer) Attributes() networkmanagerConnectPeerAttributes {
	return networkmanagerConnectPeerAttributes{ref: terra.ReferenceResource(ncp)}
}

func (ncp *NetworkmanagerConnectPeer) ImportState(av io.Reader) error {
	ncp.state = &networkmanagerConnectPeerState{}
	if err := json.NewDecoder(av).Decode(ncp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ncp.Type(), ncp.LocalName(), err)
	}
	return nil
}

func (ncp *NetworkmanagerConnectPeer) State() (*networkmanagerConnectPeerState, bool) {
	return ncp.state, ncp.state != nil
}

func (ncp *NetworkmanagerConnectPeer) StateMust() *networkmanagerConnectPeerState {
	if ncp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ncp.Type(), ncp.LocalName()))
	}
	return ncp.state
}

func (ncp *NetworkmanagerConnectPeer) DependOn() terra.Reference {
	return terra.ReferenceResource(ncp)
}

type NetworkmanagerConnectPeerArgs struct {
	// ConnectAttachmentId: string, required
	ConnectAttachmentId terra.StringValue `hcl:"connect_attachment_id,attr" validate:"required"`
	// CoreNetworkAddress: string, optional
	CoreNetworkAddress terra.StringValue `hcl:"core_network_address,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InsideCidrBlocks: list of string, required
	InsideCidrBlocks terra.ListValue[terra.StringValue] `hcl:"inside_cidr_blocks,attr" validate:"required"`
	// PeerAddress: string, required
	PeerAddress terra.StringValue `hcl:"peer_address,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Configuration: min=0
	Configuration []networkmanagerconnectpeer.Configuration `hcl:"configuration,block" validate:"min=0"`
	// BgpOptions: optional
	BgpOptions *networkmanagerconnectpeer.BgpOptions `hcl:"bgp_options,block"`
	// Timeouts: optional
	Timeouts *networkmanagerconnectpeer.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that NetworkmanagerConnectPeer depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type networkmanagerConnectPeerAttributes struct {
	ref terra.Reference
}

func (ncp networkmanagerConnectPeerAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(ncp.ref.Append("arn"))
}

func (ncp networkmanagerConnectPeerAttributes) ConnectAttachmentId() terra.StringValue {
	return terra.ReferenceString(ncp.ref.Append("connect_attachment_id"))
}

func (ncp networkmanagerConnectPeerAttributes) ConnectPeerId() terra.StringValue {
	return terra.ReferenceString(ncp.ref.Append("connect_peer_id"))
}

func (ncp networkmanagerConnectPeerAttributes) CoreNetworkAddress() terra.StringValue {
	return terra.ReferenceString(ncp.ref.Append("core_network_address"))
}

func (ncp networkmanagerConnectPeerAttributes) CoreNetworkId() terra.StringValue {
	return terra.ReferenceString(ncp.ref.Append("core_network_id"))
}

func (ncp networkmanagerConnectPeerAttributes) CreatedAt() terra.StringValue {
	return terra.ReferenceString(ncp.ref.Append("created_at"))
}

func (ncp networkmanagerConnectPeerAttributes) EdgeLocation() terra.StringValue {
	return terra.ReferenceString(ncp.ref.Append("edge_location"))
}

func (ncp networkmanagerConnectPeerAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ncp.ref.Append("id"))
}

func (ncp networkmanagerConnectPeerAttributes) InsideCidrBlocks() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](ncp.ref.Append("inside_cidr_blocks"))
}

func (ncp networkmanagerConnectPeerAttributes) PeerAddress() terra.StringValue {
	return terra.ReferenceString(ncp.ref.Append("peer_address"))
}

func (ncp networkmanagerConnectPeerAttributes) State() terra.StringValue {
	return terra.ReferenceString(ncp.ref.Append("state"))
}

func (ncp networkmanagerConnectPeerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](ncp.ref.Append("tags"))
}

func (ncp networkmanagerConnectPeerAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](ncp.ref.Append("tags_all"))
}

func (ncp networkmanagerConnectPeerAttributes) Configuration() terra.ListValue[networkmanagerconnectpeer.ConfigurationAttributes] {
	return terra.ReferenceList[networkmanagerconnectpeer.ConfigurationAttributes](ncp.ref.Append("configuration"))
}

func (ncp networkmanagerConnectPeerAttributes) BgpOptions() terra.ListValue[networkmanagerconnectpeer.BgpOptionsAttributes] {
	return terra.ReferenceList[networkmanagerconnectpeer.BgpOptionsAttributes](ncp.ref.Append("bgp_options"))
}

func (ncp networkmanagerConnectPeerAttributes) Timeouts() networkmanagerconnectpeer.TimeoutsAttributes {
	return terra.ReferenceSingle[networkmanagerconnectpeer.TimeoutsAttributes](ncp.ref.Append("timeouts"))
}

type networkmanagerConnectPeerState struct {
	Arn                 string                                         `json:"arn"`
	ConnectAttachmentId string                                         `json:"connect_attachment_id"`
	ConnectPeerId       string                                         `json:"connect_peer_id"`
	CoreNetworkAddress  string                                         `json:"core_network_address"`
	CoreNetworkId       string                                         `json:"core_network_id"`
	CreatedAt           string                                         `json:"created_at"`
	EdgeLocation        string                                         `json:"edge_location"`
	Id                  string                                         `json:"id"`
	InsideCidrBlocks    []string                                       `json:"inside_cidr_blocks"`
	PeerAddress         string                                         `json:"peer_address"`
	State               string                                         `json:"state"`
	Tags                map[string]string                              `json:"tags"`
	TagsAll             map[string]string                              `json:"tags_all"`
	Configuration       []networkmanagerconnectpeer.ConfigurationState `json:"configuration"`
	BgpOptions          []networkmanagerconnectpeer.BgpOptionsState    `json:"bgp_options"`
	Timeouts            *networkmanagerconnectpeer.TimeoutsState       `json:"timeouts"`
}
