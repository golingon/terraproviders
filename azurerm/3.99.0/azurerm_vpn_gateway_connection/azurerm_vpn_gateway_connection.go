// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_vpn_gateway_connection

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

// Resource represents the Terraform resource azurerm_vpn_gateway_connection.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermVpnGatewayConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (avgc *Resource) Type() string {
	return "azurerm_vpn_gateway_connection"
}

// LocalName returns the local name for [Resource].
func (avgc *Resource) LocalName() string {
	return avgc.Name
}

// Configuration returns the configuration (args) for [Resource].
func (avgc *Resource) Configuration() interface{} {
	return avgc.Args
}

// DependOn is used for other resources to depend on [Resource].
func (avgc *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(avgc)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (avgc *Resource) Dependencies() terra.Dependencies {
	return avgc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (avgc *Resource) LifecycleManagement() *terra.Lifecycle {
	return avgc.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (avgc *Resource) Attributes() azurermVpnGatewayConnectionAttributes {
	return azurermVpnGatewayConnectionAttributes{ref: terra.ReferenceResource(avgc)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (avgc *Resource) ImportState(state io.Reader) error {
	avgc.state = &azurermVpnGatewayConnectionState{}
	if err := json.NewDecoder(state).Decode(avgc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", avgc.Type(), avgc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (avgc *Resource) State() (*azurermVpnGatewayConnectionState, bool) {
	return avgc.state, avgc.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (avgc *Resource) StateMust() *azurermVpnGatewayConnectionState {
	if avgc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", avgc.Type(), avgc.LocalName()))
	}
	return avgc.state
}

// Args contains the configurations for azurerm_vpn_gateway_connection.
type Args struct {
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
	Routing *Routing `hcl:"routing,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
	// TrafficSelectorPolicy: min=0
	TrafficSelectorPolicy []TrafficSelectorPolicy `hcl:"traffic_selector_policy,block" validate:"min=0"`
	// VpnLink: min=1
	VpnLink []VpnLink `hcl:"vpn_link,block" validate:"min=1"`
}

type azurermVpnGatewayConnectionAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_vpn_gateway_connection.
func (avgc azurermVpnGatewayConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(avgc.ref.Append("id"))
}

// InternetSecurityEnabled returns a reference to field internet_security_enabled of azurerm_vpn_gateway_connection.
func (avgc azurermVpnGatewayConnectionAttributes) InternetSecurityEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(avgc.ref.Append("internet_security_enabled"))
}

// Name returns a reference to field name of azurerm_vpn_gateway_connection.
func (avgc azurermVpnGatewayConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(avgc.ref.Append("name"))
}

// RemoteVpnSiteId returns a reference to field remote_vpn_site_id of azurerm_vpn_gateway_connection.
func (avgc azurermVpnGatewayConnectionAttributes) RemoteVpnSiteId() terra.StringValue {
	return terra.ReferenceAsString(avgc.ref.Append("remote_vpn_site_id"))
}

// VpnGatewayId returns a reference to field vpn_gateway_id of azurerm_vpn_gateway_connection.
func (avgc azurermVpnGatewayConnectionAttributes) VpnGatewayId() terra.StringValue {
	return terra.ReferenceAsString(avgc.ref.Append("vpn_gateway_id"))
}

func (avgc azurermVpnGatewayConnectionAttributes) Routing() terra.ListValue[RoutingAttributes] {
	return terra.ReferenceAsList[RoutingAttributes](avgc.ref.Append("routing"))
}

func (avgc azurermVpnGatewayConnectionAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](avgc.ref.Append("timeouts"))
}

func (avgc azurermVpnGatewayConnectionAttributes) TrafficSelectorPolicy() terra.SetValue[TrafficSelectorPolicyAttributes] {
	return terra.ReferenceAsSet[TrafficSelectorPolicyAttributes](avgc.ref.Append("traffic_selector_policy"))
}

func (avgc azurermVpnGatewayConnectionAttributes) VpnLink() terra.ListValue[VpnLinkAttributes] {
	return terra.ReferenceAsList[VpnLinkAttributes](avgc.ref.Append("vpn_link"))
}

type azurermVpnGatewayConnectionState struct {
	Id                      string                       `json:"id"`
	InternetSecurityEnabled bool                         `json:"internet_security_enabled"`
	Name                    string                       `json:"name"`
	RemoteVpnSiteId         string                       `json:"remote_vpn_site_id"`
	VpnGatewayId            string                       `json:"vpn_gateway_id"`
	Routing                 []RoutingState               `json:"routing"`
	Timeouts                *TimeoutsState               `json:"timeouts"`
	TrafficSelectorPolicy   []TrafficSelectorPolicyState `json:"traffic_selector_policy"`
	VpnLink                 []VpnLinkState               `json:"vpn_link"`
}
