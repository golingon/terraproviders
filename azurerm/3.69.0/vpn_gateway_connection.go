// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	vpngatewayconnection "github.com/golingon/terraproviders/azurerm/3.69.0/vpngatewayconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVpnGatewayConnection creates a new instance of [VpnGatewayConnection].
func NewVpnGatewayConnection(name string, args VpnGatewayConnectionArgs) *VpnGatewayConnection {
	return &VpnGatewayConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpnGatewayConnection)(nil)

// VpnGatewayConnection represents the Terraform resource azurerm_vpn_gateway_connection.
type VpnGatewayConnection struct {
	Name      string
	Args      VpnGatewayConnectionArgs
	state     *vpnGatewayConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VpnGatewayConnection].
func (vgc *VpnGatewayConnection) Type() string {
	return "azurerm_vpn_gateway_connection"
}

// LocalName returns the local name for [VpnGatewayConnection].
func (vgc *VpnGatewayConnection) LocalName() string {
	return vgc.Name
}

// Configuration returns the configuration (args) for [VpnGatewayConnection].
func (vgc *VpnGatewayConnection) Configuration() interface{} {
	return vgc.Args
}

// DependOn is used for other resources to depend on [VpnGatewayConnection].
func (vgc *VpnGatewayConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(vgc)
}

// Dependencies returns the list of resources [VpnGatewayConnection] depends_on.
func (vgc *VpnGatewayConnection) Dependencies() terra.Dependencies {
	return vgc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VpnGatewayConnection].
func (vgc *VpnGatewayConnection) LifecycleManagement() *terra.Lifecycle {
	return vgc.Lifecycle
}

// Attributes returns the attributes for [VpnGatewayConnection].
func (vgc *VpnGatewayConnection) Attributes() vpnGatewayConnectionAttributes {
	return vpnGatewayConnectionAttributes{ref: terra.ReferenceResource(vgc)}
}

// ImportState imports the given attribute values into [VpnGatewayConnection]'s state.
func (vgc *VpnGatewayConnection) ImportState(av io.Reader) error {
	vgc.state = &vpnGatewayConnectionState{}
	if err := json.NewDecoder(av).Decode(vgc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vgc.Type(), vgc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VpnGatewayConnection] has state.
func (vgc *VpnGatewayConnection) State() (*vpnGatewayConnectionState, bool) {
	return vgc.state, vgc.state != nil
}

// StateMust returns the state for [VpnGatewayConnection]. Panics if the state is nil.
func (vgc *VpnGatewayConnection) StateMust() *vpnGatewayConnectionState {
	if vgc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vgc.Type(), vgc.LocalName()))
	}
	return vgc.state
}

// VpnGatewayConnectionArgs contains the configurations for azurerm_vpn_gateway_connection.
type VpnGatewayConnectionArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InternetSecurityEnabled: bool, optional
	InternetSecurityEnabled terra.BoolValue `hcl:"internet_security_enabled,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RemoteVpnSiteId: string, required
	RemoteVpnSiteId terra.StringValue `hcl:"remote_vpn_site_id,attr" validate:"required"`
	// VpnGatewayId: string, required
	VpnGatewayId terra.StringValue `hcl:"vpn_gateway_id,attr" validate:"required"`
	// Routing: optional
	Routing *vpngatewayconnection.Routing `hcl:"routing,block"`
	// Timeouts: optional
	Timeouts *vpngatewayconnection.Timeouts `hcl:"timeouts,block"`
	// TrafficSelectorPolicy: min=0
	TrafficSelectorPolicy []vpngatewayconnection.TrafficSelectorPolicy `hcl:"traffic_selector_policy,block" validate:"min=0"`
	// VpnLink: min=1
	VpnLink []vpngatewayconnection.VpnLink `hcl:"vpn_link,block" validate:"min=1"`
}
type vpnGatewayConnectionAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_vpn_gateway_connection.
func (vgc vpnGatewayConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vgc.ref.Append("id"))
}

// InternetSecurityEnabled returns a reference to field internet_security_enabled of azurerm_vpn_gateway_connection.
func (vgc vpnGatewayConnectionAttributes) InternetSecurityEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(vgc.ref.Append("internet_security_enabled"))
}

// Name returns a reference to field name of azurerm_vpn_gateway_connection.
func (vgc vpnGatewayConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vgc.ref.Append("name"))
}

// RemoteVpnSiteId returns a reference to field remote_vpn_site_id of azurerm_vpn_gateway_connection.
func (vgc vpnGatewayConnectionAttributes) RemoteVpnSiteId() terra.StringValue {
	return terra.ReferenceAsString(vgc.ref.Append("remote_vpn_site_id"))
}

// VpnGatewayId returns a reference to field vpn_gateway_id of azurerm_vpn_gateway_connection.
func (vgc vpnGatewayConnectionAttributes) VpnGatewayId() terra.StringValue {
	return terra.ReferenceAsString(vgc.ref.Append("vpn_gateway_id"))
}

func (vgc vpnGatewayConnectionAttributes) Routing() terra.ListValue[vpngatewayconnection.RoutingAttributes] {
	return terra.ReferenceAsList[vpngatewayconnection.RoutingAttributes](vgc.ref.Append("routing"))
}

func (vgc vpnGatewayConnectionAttributes) Timeouts() vpngatewayconnection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vpngatewayconnection.TimeoutsAttributes](vgc.ref.Append("timeouts"))
}

func (vgc vpnGatewayConnectionAttributes) TrafficSelectorPolicy() terra.SetValue[vpngatewayconnection.TrafficSelectorPolicyAttributes] {
	return terra.ReferenceAsSet[vpngatewayconnection.TrafficSelectorPolicyAttributes](vgc.ref.Append("traffic_selector_policy"))
}

func (vgc vpnGatewayConnectionAttributes) VpnLink() terra.ListValue[vpngatewayconnection.VpnLinkAttributes] {
	return terra.ReferenceAsList[vpngatewayconnection.VpnLinkAttributes](vgc.ref.Append("vpn_link"))
}

type vpnGatewayConnectionState struct {
	Id                      string                                            `json:"id"`
	InternetSecurityEnabled bool                                              `json:"internet_security_enabled"`
	Name                    string                                            `json:"name"`
	RemoteVpnSiteId         string                                            `json:"remote_vpn_site_id"`
	VpnGatewayId            string                                            `json:"vpn_gateway_id"`
	Routing                 []vpngatewayconnection.RoutingState               `json:"routing"`
	Timeouts                *vpngatewayconnection.TimeoutsState               `json:"timeouts"`
	TrafficSelectorPolicy   []vpngatewayconnection.TrafficSelectorPolicyState `json:"traffic_selector_policy"`
	VpnLink                 []vpngatewayconnection.VpnLinkState               `json:"vpn_link"`
}
