// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	networkmanagerconnectpeer "github.com/golingon/terraproviders/aws/4.63.0/networkmanagerconnectpeer"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkmanagerConnectPeer creates a new instance of [NetworkmanagerConnectPeer].
func NewNetworkmanagerConnectPeer(name string, args NetworkmanagerConnectPeerArgs) *NetworkmanagerConnectPeer {
	return &NetworkmanagerConnectPeer{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkmanagerConnectPeer)(nil)

// NetworkmanagerConnectPeer represents the Terraform resource aws_networkmanager_connect_peer.
type NetworkmanagerConnectPeer struct {
	Name      string
	Args      NetworkmanagerConnectPeerArgs
	state     *networkmanagerConnectPeerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkmanagerConnectPeer].
func (ncp *NetworkmanagerConnectPeer) Type() string {
	return "aws_networkmanager_connect_peer"
}

// LocalName returns the local name for [NetworkmanagerConnectPeer].
func (ncp *NetworkmanagerConnectPeer) LocalName() string {
	return ncp.Name
}

// Configuration returns the configuration (args) for [NetworkmanagerConnectPeer].
func (ncp *NetworkmanagerConnectPeer) Configuration() interface{} {
	return ncp.Args
}

// DependOn is used for other resources to depend on [NetworkmanagerConnectPeer].
func (ncp *NetworkmanagerConnectPeer) DependOn() terra.Reference {
	return terra.ReferenceResource(ncp)
}

// Dependencies returns the list of resources [NetworkmanagerConnectPeer] depends_on.
func (ncp *NetworkmanagerConnectPeer) Dependencies() terra.Dependencies {
	return ncp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkmanagerConnectPeer].
func (ncp *NetworkmanagerConnectPeer) LifecycleManagement() *terra.Lifecycle {
	return ncp.Lifecycle
}

// Attributes returns the attributes for [NetworkmanagerConnectPeer].
func (ncp *NetworkmanagerConnectPeer) Attributes() networkmanagerConnectPeerAttributes {
	return networkmanagerConnectPeerAttributes{ref: terra.ReferenceResource(ncp)}
}

// ImportState imports the given attribute values into [NetworkmanagerConnectPeer]'s state.
func (ncp *NetworkmanagerConnectPeer) ImportState(av io.Reader) error {
	ncp.state = &networkmanagerConnectPeerState{}
	if err := json.NewDecoder(av).Decode(ncp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ncp.Type(), ncp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkmanagerConnectPeer] has state.
func (ncp *NetworkmanagerConnectPeer) State() (*networkmanagerConnectPeerState, bool) {
	return ncp.state, ncp.state != nil
}

// StateMust returns the state for [NetworkmanagerConnectPeer]. Panics if the state is nil.
func (ncp *NetworkmanagerConnectPeer) StateMust() *networkmanagerConnectPeerState {
	if ncp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ncp.Type(), ncp.LocalName()))
	}
	return ncp.state
}

// NetworkmanagerConnectPeerArgs contains the configurations for aws_networkmanager_connect_peer.
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
}
type networkmanagerConnectPeerAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_networkmanager_connect_peer.
func (ncp networkmanagerConnectPeerAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ncp.ref.Append("arn"))
}

// ConnectAttachmentId returns a reference to field connect_attachment_id of aws_networkmanager_connect_peer.
func (ncp networkmanagerConnectPeerAttributes) ConnectAttachmentId() terra.StringValue {
	return terra.ReferenceAsString(ncp.ref.Append("connect_attachment_id"))
}

// ConnectPeerId returns a reference to field connect_peer_id of aws_networkmanager_connect_peer.
func (ncp networkmanagerConnectPeerAttributes) ConnectPeerId() terra.StringValue {
	return terra.ReferenceAsString(ncp.ref.Append("connect_peer_id"))
}

// CoreNetworkAddress returns a reference to field core_network_address of aws_networkmanager_connect_peer.
func (ncp networkmanagerConnectPeerAttributes) CoreNetworkAddress() terra.StringValue {
	return terra.ReferenceAsString(ncp.ref.Append("core_network_address"))
}

// CoreNetworkId returns a reference to field core_network_id of aws_networkmanager_connect_peer.
func (ncp networkmanagerConnectPeerAttributes) CoreNetworkId() terra.StringValue {
	return terra.ReferenceAsString(ncp.ref.Append("core_network_id"))
}

// CreatedAt returns a reference to field created_at of aws_networkmanager_connect_peer.
func (ncp networkmanagerConnectPeerAttributes) CreatedAt() terra.StringValue {
	return terra.ReferenceAsString(ncp.ref.Append("created_at"))
}

// EdgeLocation returns a reference to field edge_location of aws_networkmanager_connect_peer.
func (ncp networkmanagerConnectPeerAttributes) EdgeLocation() terra.StringValue {
	return terra.ReferenceAsString(ncp.ref.Append("edge_location"))
}

// Id returns a reference to field id of aws_networkmanager_connect_peer.
func (ncp networkmanagerConnectPeerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ncp.ref.Append("id"))
}

// InsideCidrBlocks returns a reference to field inside_cidr_blocks of aws_networkmanager_connect_peer.
func (ncp networkmanagerConnectPeerAttributes) InsideCidrBlocks() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ncp.ref.Append("inside_cidr_blocks"))
}

// PeerAddress returns a reference to field peer_address of aws_networkmanager_connect_peer.
func (ncp networkmanagerConnectPeerAttributes) PeerAddress() terra.StringValue {
	return terra.ReferenceAsString(ncp.ref.Append("peer_address"))
}

// State returns a reference to field state of aws_networkmanager_connect_peer.
func (ncp networkmanagerConnectPeerAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(ncp.ref.Append("state"))
}

// Tags returns a reference to field tags of aws_networkmanager_connect_peer.
func (ncp networkmanagerConnectPeerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ncp.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_networkmanager_connect_peer.
func (ncp networkmanagerConnectPeerAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ncp.ref.Append("tags_all"))
}

func (ncp networkmanagerConnectPeerAttributes) Configuration() terra.ListValue[networkmanagerconnectpeer.ConfigurationAttributes] {
	return terra.ReferenceAsList[networkmanagerconnectpeer.ConfigurationAttributes](ncp.ref.Append("configuration"))
}

func (ncp networkmanagerConnectPeerAttributes) BgpOptions() terra.ListValue[networkmanagerconnectpeer.BgpOptionsAttributes] {
	return terra.ReferenceAsList[networkmanagerconnectpeer.BgpOptionsAttributes](ncp.ref.Append("bgp_options"))
}

func (ncp networkmanagerConnectPeerAttributes) Timeouts() networkmanagerconnectpeer.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkmanagerconnectpeer.TimeoutsAttributes](ncp.ref.Append("timeouts"))
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
