// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	virtualhubbgpconnection "github.com/golingon/terraproviders/azurerm/3.52.0/virtualhubbgpconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVirtualHubBgpConnection creates a new instance of [VirtualHubBgpConnection].
func NewVirtualHubBgpConnection(name string, args VirtualHubBgpConnectionArgs) *VirtualHubBgpConnection {
	return &VirtualHubBgpConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VirtualHubBgpConnection)(nil)

// VirtualHubBgpConnection represents the Terraform resource azurerm_virtual_hub_bgp_connection.
type VirtualHubBgpConnection struct {
	Name      string
	Args      VirtualHubBgpConnectionArgs
	state     *virtualHubBgpConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VirtualHubBgpConnection].
func (vhbc *VirtualHubBgpConnection) Type() string {
	return "azurerm_virtual_hub_bgp_connection"
}

// LocalName returns the local name for [VirtualHubBgpConnection].
func (vhbc *VirtualHubBgpConnection) LocalName() string {
	return vhbc.Name
}

// Configuration returns the configuration (args) for [VirtualHubBgpConnection].
func (vhbc *VirtualHubBgpConnection) Configuration() interface{} {
	return vhbc.Args
}

// DependOn is used for other resources to depend on [VirtualHubBgpConnection].
func (vhbc *VirtualHubBgpConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(vhbc)
}

// Dependencies returns the list of resources [VirtualHubBgpConnection] depends_on.
func (vhbc *VirtualHubBgpConnection) Dependencies() terra.Dependencies {
	return vhbc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VirtualHubBgpConnection].
func (vhbc *VirtualHubBgpConnection) LifecycleManagement() *terra.Lifecycle {
	return vhbc.Lifecycle
}

// Attributes returns the attributes for [VirtualHubBgpConnection].
func (vhbc *VirtualHubBgpConnection) Attributes() virtualHubBgpConnectionAttributes {
	return virtualHubBgpConnectionAttributes{ref: terra.ReferenceResource(vhbc)}
}

// ImportState imports the given attribute values into [VirtualHubBgpConnection]'s state.
func (vhbc *VirtualHubBgpConnection) ImportState(av io.Reader) error {
	vhbc.state = &virtualHubBgpConnectionState{}
	if err := json.NewDecoder(av).Decode(vhbc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vhbc.Type(), vhbc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VirtualHubBgpConnection] has state.
func (vhbc *VirtualHubBgpConnection) State() (*virtualHubBgpConnectionState, bool) {
	return vhbc.state, vhbc.state != nil
}

// StateMust returns the state for [VirtualHubBgpConnection]. Panics if the state is nil.
func (vhbc *VirtualHubBgpConnection) StateMust() *virtualHubBgpConnectionState {
	if vhbc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vhbc.Type(), vhbc.LocalName()))
	}
	return vhbc.state
}

// VirtualHubBgpConnectionArgs contains the configurations for azurerm_virtual_hub_bgp_connection.
type VirtualHubBgpConnectionArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PeerAsn: number, required
	PeerAsn terra.NumberValue `hcl:"peer_asn,attr" validate:"required"`
	// PeerIp: string, required
	PeerIp terra.StringValue `hcl:"peer_ip,attr" validate:"required"`
	// VirtualHubId: string, required
	VirtualHubId terra.StringValue `hcl:"virtual_hub_id,attr" validate:"required"`
	// VirtualNetworkConnectionId: string, optional
	VirtualNetworkConnectionId terra.StringValue `hcl:"virtual_network_connection_id,attr"`
	// Timeouts: optional
	Timeouts *virtualhubbgpconnection.Timeouts `hcl:"timeouts,block"`
}
type virtualHubBgpConnectionAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_virtual_hub_bgp_connection.
func (vhbc virtualHubBgpConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vhbc.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_virtual_hub_bgp_connection.
func (vhbc virtualHubBgpConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vhbc.ref.Append("name"))
}

// PeerAsn returns a reference to field peer_asn of azurerm_virtual_hub_bgp_connection.
func (vhbc virtualHubBgpConnectionAttributes) PeerAsn() terra.NumberValue {
	return terra.ReferenceAsNumber(vhbc.ref.Append("peer_asn"))
}

// PeerIp returns a reference to field peer_ip of azurerm_virtual_hub_bgp_connection.
func (vhbc virtualHubBgpConnectionAttributes) PeerIp() terra.StringValue {
	return terra.ReferenceAsString(vhbc.ref.Append("peer_ip"))
}

// VirtualHubId returns a reference to field virtual_hub_id of azurerm_virtual_hub_bgp_connection.
func (vhbc virtualHubBgpConnectionAttributes) VirtualHubId() terra.StringValue {
	return terra.ReferenceAsString(vhbc.ref.Append("virtual_hub_id"))
}

// VirtualNetworkConnectionId returns a reference to field virtual_network_connection_id of azurerm_virtual_hub_bgp_connection.
func (vhbc virtualHubBgpConnectionAttributes) VirtualNetworkConnectionId() terra.StringValue {
	return terra.ReferenceAsString(vhbc.ref.Append("virtual_network_connection_id"))
}

func (vhbc virtualHubBgpConnectionAttributes) Timeouts() virtualhubbgpconnection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[virtualhubbgpconnection.TimeoutsAttributes](vhbc.ref.Append("timeouts"))
}

type virtualHubBgpConnectionState struct {
	Id                         string                                 `json:"id"`
	Name                       string                                 `json:"name"`
	PeerAsn                    float64                                `json:"peer_asn"`
	PeerIp                     string                                 `json:"peer_ip"`
	VirtualHubId               string                                 `json:"virtual_hub_id"`
	VirtualNetworkConnectionId string                                 `json:"virtual_network_connection_id"`
	Timeouts                   *virtualhubbgpconnection.TimeoutsState `json:"timeouts"`
}
