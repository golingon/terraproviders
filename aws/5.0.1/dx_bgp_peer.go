// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	dxbgppeer "github.com/golingon/terraproviders/aws/5.0.1/dxbgppeer"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDxBgpPeer creates a new instance of [DxBgpPeer].
func NewDxBgpPeer(name string, args DxBgpPeerArgs) *DxBgpPeer {
	return &DxBgpPeer{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DxBgpPeer)(nil)

// DxBgpPeer represents the Terraform resource aws_dx_bgp_peer.
type DxBgpPeer struct {
	Name      string
	Args      DxBgpPeerArgs
	state     *dxBgpPeerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DxBgpPeer].
func (dbp *DxBgpPeer) Type() string {
	return "aws_dx_bgp_peer"
}

// LocalName returns the local name for [DxBgpPeer].
func (dbp *DxBgpPeer) LocalName() string {
	return dbp.Name
}

// Configuration returns the configuration (args) for [DxBgpPeer].
func (dbp *DxBgpPeer) Configuration() interface{} {
	return dbp.Args
}

// DependOn is used for other resources to depend on [DxBgpPeer].
func (dbp *DxBgpPeer) DependOn() terra.Reference {
	return terra.ReferenceResource(dbp)
}

// Dependencies returns the list of resources [DxBgpPeer] depends_on.
func (dbp *DxBgpPeer) Dependencies() terra.Dependencies {
	return dbp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DxBgpPeer].
func (dbp *DxBgpPeer) LifecycleManagement() *terra.Lifecycle {
	return dbp.Lifecycle
}

// Attributes returns the attributes for [DxBgpPeer].
func (dbp *DxBgpPeer) Attributes() dxBgpPeerAttributes {
	return dxBgpPeerAttributes{ref: terra.ReferenceResource(dbp)}
}

// ImportState imports the given attribute values into [DxBgpPeer]'s state.
func (dbp *DxBgpPeer) ImportState(av io.Reader) error {
	dbp.state = &dxBgpPeerState{}
	if err := json.NewDecoder(av).Decode(dbp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dbp.Type(), dbp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DxBgpPeer] has state.
func (dbp *DxBgpPeer) State() (*dxBgpPeerState, bool) {
	return dbp.state, dbp.state != nil
}

// StateMust returns the state for [DxBgpPeer]. Panics if the state is nil.
func (dbp *DxBgpPeer) StateMust() *dxBgpPeerState {
	if dbp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dbp.Type(), dbp.LocalName()))
	}
	return dbp.state
}

// DxBgpPeerArgs contains the configurations for aws_dx_bgp_peer.
type DxBgpPeerArgs struct {
	// AddressFamily: string, required
	AddressFamily terra.StringValue `hcl:"address_family,attr" validate:"required"`
	// AmazonAddress: string, optional
	AmazonAddress terra.StringValue `hcl:"amazon_address,attr"`
	// BgpAsn: number, required
	BgpAsn terra.NumberValue `hcl:"bgp_asn,attr" validate:"required"`
	// BgpAuthKey: string, optional
	BgpAuthKey terra.StringValue `hcl:"bgp_auth_key,attr"`
	// CustomerAddress: string, optional
	CustomerAddress terra.StringValue `hcl:"customer_address,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// VirtualInterfaceId: string, required
	VirtualInterfaceId terra.StringValue `hcl:"virtual_interface_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dxbgppeer.Timeouts `hcl:"timeouts,block"`
}
type dxBgpPeerAttributes struct {
	ref terra.Reference
}

// AddressFamily returns a reference to field address_family of aws_dx_bgp_peer.
func (dbp dxBgpPeerAttributes) AddressFamily() terra.StringValue {
	return terra.ReferenceAsString(dbp.ref.Append("address_family"))
}

// AmazonAddress returns a reference to field amazon_address of aws_dx_bgp_peer.
func (dbp dxBgpPeerAttributes) AmazonAddress() terra.StringValue {
	return terra.ReferenceAsString(dbp.ref.Append("amazon_address"))
}

// AwsDevice returns a reference to field aws_device of aws_dx_bgp_peer.
func (dbp dxBgpPeerAttributes) AwsDevice() terra.StringValue {
	return terra.ReferenceAsString(dbp.ref.Append("aws_device"))
}

// BgpAsn returns a reference to field bgp_asn of aws_dx_bgp_peer.
func (dbp dxBgpPeerAttributes) BgpAsn() terra.NumberValue {
	return terra.ReferenceAsNumber(dbp.ref.Append("bgp_asn"))
}

// BgpAuthKey returns a reference to field bgp_auth_key of aws_dx_bgp_peer.
func (dbp dxBgpPeerAttributes) BgpAuthKey() terra.StringValue {
	return terra.ReferenceAsString(dbp.ref.Append("bgp_auth_key"))
}

// BgpPeerId returns a reference to field bgp_peer_id of aws_dx_bgp_peer.
func (dbp dxBgpPeerAttributes) BgpPeerId() terra.StringValue {
	return terra.ReferenceAsString(dbp.ref.Append("bgp_peer_id"))
}

// BgpStatus returns a reference to field bgp_status of aws_dx_bgp_peer.
func (dbp dxBgpPeerAttributes) BgpStatus() terra.StringValue {
	return terra.ReferenceAsString(dbp.ref.Append("bgp_status"))
}

// CustomerAddress returns a reference to field customer_address of aws_dx_bgp_peer.
func (dbp dxBgpPeerAttributes) CustomerAddress() terra.StringValue {
	return terra.ReferenceAsString(dbp.ref.Append("customer_address"))
}

// Id returns a reference to field id of aws_dx_bgp_peer.
func (dbp dxBgpPeerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dbp.ref.Append("id"))
}

// VirtualInterfaceId returns a reference to field virtual_interface_id of aws_dx_bgp_peer.
func (dbp dxBgpPeerAttributes) VirtualInterfaceId() terra.StringValue {
	return terra.ReferenceAsString(dbp.ref.Append("virtual_interface_id"))
}

func (dbp dxBgpPeerAttributes) Timeouts() dxbgppeer.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dxbgppeer.TimeoutsAttributes](dbp.ref.Append("timeouts"))
}

type dxBgpPeerState struct {
	AddressFamily      string                   `json:"address_family"`
	AmazonAddress      string                   `json:"amazon_address"`
	AwsDevice          string                   `json:"aws_device"`
	BgpAsn             float64                  `json:"bgp_asn"`
	BgpAuthKey         string                   `json:"bgp_auth_key"`
	BgpPeerId          string                   `json:"bgp_peer_id"`
	BgpStatus          string                   `json:"bgp_status"`
	CustomerAddress    string                   `json:"customer_address"`
	Id                 string                   `json:"id"`
	VirtualInterfaceId string                   `json:"virtual_interface_id"`
	Timeouts           *dxbgppeer.TimeoutsState `json:"timeouts"`
}
