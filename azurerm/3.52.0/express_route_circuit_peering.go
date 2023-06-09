// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	expressroutecircuitpeering "github.com/golingon/terraproviders/azurerm/3.52.0/expressroutecircuitpeering"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewExpressRouteCircuitPeering creates a new instance of [ExpressRouteCircuitPeering].
func NewExpressRouteCircuitPeering(name string, args ExpressRouteCircuitPeeringArgs) *ExpressRouteCircuitPeering {
	return &ExpressRouteCircuitPeering{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ExpressRouteCircuitPeering)(nil)

// ExpressRouteCircuitPeering represents the Terraform resource azurerm_express_route_circuit_peering.
type ExpressRouteCircuitPeering struct {
	Name      string
	Args      ExpressRouteCircuitPeeringArgs
	state     *expressRouteCircuitPeeringState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ExpressRouteCircuitPeering].
func (ercp *ExpressRouteCircuitPeering) Type() string {
	return "azurerm_express_route_circuit_peering"
}

// LocalName returns the local name for [ExpressRouteCircuitPeering].
func (ercp *ExpressRouteCircuitPeering) LocalName() string {
	return ercp.Name
}

// Configuration returns the configuration (args) for [ExpressRouteCircuitPeering].
func (ercp *ExpressRouteCircuitPeering) Configuration() interface{} {
	return ercp.Args
}

// DependOn is used for other resources to depend on [ExpressRouteCircuitPeering].
func (ercp *ExpressRouteCircuitPeering) DependOn() terra.Reference {
	return terra.ReferenceResource(ercp)
}

// Dependencies returns the list of resources [ExpressRouteCircuitPeering] depends_on.
func (ercp *ExpressRouteCircuitPeering) Dependencies() terra.Dependencies {
	return ercp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ExpressRouteCircuitPeering].
func (ercp *ExpressRouteCircuitPeering) LifecycleManagement() *terra.Lifecycle {
	return ercp.Lifecycle
}

// Attributes returns the attributes for [ExpressRouteCircuitPeering].
func (ercp *ExpressRouteCircuitPeering) Attributes() expressRouteCircuitPeeringAttributes {
	return expressRouteCircuitPeeringAttributes{ref: terra.ReferenceResource(ercp)}
}

// ImportState imports the given attribute values into [ExpressRouteCircuitPeering]'s state.
func (ercp *ExpressRouteCircuitPeering) ImportState(av io.Reader) error {
	ercp.state = &expressRouteCircuitPeeringState{}
	if err := json.NewDecoder(av).Decode(ercp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ercp.Type(), ercp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ExpressRouteCircuitPeering] has state.
func (ercp *ExpressRouteCircuitPeering) State() (*expressRouteCircuitPeeringState, bool) {
	return ercp.state, ercp.state != nil
}

// StateMust returns the state for [ExpressRouteCircuitPeering]. Panics if the state is nil.
func (ercp *ExpressRouteCircuitPeering) StateMust() *expressRouteCircuitPeeringState {
	if ercp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ercp.Type(), ercp.LocalName()))
	}
	return ercp.state
}

// ExpressRouteCircuitPeeringArgs contains the configurations for azurerm_express_route_circuit_peering.
type ExpressRouteCircuitPeeringArgs struct {
	// ExpressRouteCircuitName: string, required
	ExpressRouteCircuitName terra.StringValue `hcl:"express_route_circuit_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Ipv4Enabled: bool, optional
	Ipv4Enabled terra.BoolValue `hcl:"ipv4_enabled,attr"`
	// PeerAsn: number, optional
	PeerAsn terra.NumberValue `hcl:"peer_asn,attr"`
	// PeeringType: string, required
	PeeringType terra.StringValue `hcl:"peering_type,attr" validate:"required"`
	// PrimaryPeerAddressPrefix: string, optional
	PrimaryPeerAddressPrefix terra.StringValue `hcl:"primary_peer_address_prefix,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// RouteFilterId: string, optional
	RouteFilterId terra.StringValue `hcl:"route_filter_id,attr"`
	// SecondaryPeerAddressPrefix: string, optional
	SecondaryPeerAddressPrefix terra.StringValue `hcl:"secondary_peer_address_prefix,attr"`
	// SharedKey: string, optional
	SharedKey terra.StringValue `hcl:"shared_key,attr"`
	// VlanId: number, required
	VlanId terra.NumberValue `hcl:"vlan_id,attr" validate:"required"`
	// Ipv6: optional
	Ipv6 *expressroutecircuitpeering.Ipv6 `hcl:"ipv6,block"`
	// MicrosoftPeeringConfig: optional
	MicrosoftPeeringConfig *expressroutecircuitpeering.MicrosoftPeeringConfig `hcl:"microsoft_peering_config,block"`
	// Timeouts: optional
	Timeouts *expressroutecircuitpeering.Timeouts `hcl:"timeouts,block"`
}
type expressRouteCircuitPeeringAttributes struct {
	ref terra.Reference
}

// AzureAsn returns a reference to field azure_asn of azurerm_express_route_circuit_peering.
func (ercp expressRouteCircuitPeeringAttributes) AzureAsn() terra.NumberValue {
	return terra.ReferenceAsNumber(ercp.ref.Append("azure_asn"))
}

// ExpressRouteCircuitName returns a reference to field express_route_circuit_name of azurerm_express_route_circuit_peering.
func (ercp expressRouteCircuitPeeringAttributes) ExpressRouteCircuitName() terra.StringValue {
	return terra.ReferenceAsString(ercp.ref.Append("express_route_circuit_name"))
}

// GatewayManagerEtag returns a reference to field gateway_manager_etag of azurerm_express_route_circuit_peering.
func (ercp expressRouteCircuitPeeringAttributes) GatewayManagerEtag() terra.StringValue {
	return terra.ReferenceAsString(ercp.ref.Append("gateway_manager_etag"))
}

// Id returns a reference to field id of azurerm_express_route_circuit_peering.
func (ercp expressRouteCircuitPeeringAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ercp.ref.Append("id"))
}

// Ipv4Enabled returns a reference to field ipv4_enabled of azurerm_express_route_circuit_peering.
func (ercp expressRouteCircuitPeeringAttributes) Ipv4Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(ercp.ref.Append("ipv4_enabled"))
}

// PeerAsn returns a reference to field peer_asn of azurerm_express_route_circuit_peering.
func (ercp expressRouteCircuitPeeringAttributes) PeerAsn() terra.NumberValue {
	return terra.ReferenceAsNumber(ercp.ref.Append("peer_asn"))
}

// PeeringType returns a reference to field peering_type of azurerm_express_route_circuit_peering.
func (ercp expressRouteCircuitPeeringAttributes) PeeringType() terra.StringValue {
	return terra.ReferenceAsString(ercp.ref.Append("peering_type"))
}

// PrimaryAzurePort returns a reference to field primary_azure_port of azurerm_express_route_circuit_peering.
func (ercp expressRouteCircuitPeeringAttributes) PrimaryAzurePort() terra.StringValue {
	return terra.ReferenceAsString(ercp.ref.Append("primary_azure_port"))
}

// PrimaryPeerAddressPrefix returns a reference to field primary_peer_address_prefix of azurerm_express_route_circuit_peering.
func (ercp expressRouteCircuitPeeringAttributes) PrimaryPeerAddressPrefix() terra.StringValue {
	return terra.ReferenceAsString(ercp.ref.Append("primary_peer_address_prefix"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_express_route_circuit_peering.
func (ercp expressRouteCircuitPeeringAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ercp.ref.Append("resource_group_name"))
}

// RouteFilterId returns a reference to field route_filter_id of azurerm_express_route_circuit_peering.
func (ercp expressRouteCircuitPeeringAttributes) RouteFilterId() terra.StringValue {
	return terra.ReferenceAsString(ercp.ref.Append("route_filter_id"))
}

// SecondaryAzurePort returns a reference to field secondary_azure_port of azurerm_express_route_circuit_peering.
func (ercp expressRouteCircuitPeeringAttributes) SecondaryAzurePort() terra.StringValue {
	return terra.ReferenceAsString(ercp.ref.Append("secondary_azure_port"))
}

// SecondaryPeerAddressPrefix returns a reference to field secondary_peer_address_prefix of azurerm_express_route_circuit_peering.
func (ercp expressRouteCircuitPeeringAttributes) SecondaryPeerAddressPrefix() terra.StringValue {
	return terra.ReferenceAsString(ercp.ref.Append("secondary_peer_address_prefix"))
}

// SharedKey returns a reference to field shared_key of azurerm_express_route_circuit_peering.
func (ercp expressRouteCircuitPeeringAttributes) SharedKey() terra.StringValue {
	return terra.ReferenceAsString(ercp.ref.Append("shared_key"))
}

// VlanId returns a reference to field vlan_id of azurerm_express_route_circuit_peering.
func (ercp expressRouteCircuitPeeringAttributes) VlanId() terra.NumberValue {
	return terra.ReferenceAsNumber(ercp.ref.Append("vlan_id"))
}

func (ercp expressRouteCircuitPeeringAttributes) Ipv6() terra.ListValue[expressroutecircuitpeering.Ipv6Attributes] {
	return terra.ReferenceAsList[expressroutecircuitpeering.Ipv6Attributes](ercp.ref.Append("ipv6"))
}

func (ercp expressRouteCircuitPeeringAttributes) MicrosoftPeeringConfig() terra.ListValue[expressroutecircuitpeering.MicrosoftPeeringConfigAttributes] {
	return terra.ReferenceAsList[expressroutecircuitpeering.MicrosoftPeeringConfigAttributes](ercp.ref.Append("microsoft_peering_config"))
}

func (ercp expressRouteCircuitPeeringAttributes) Timeouts() expressroutecircuitpeering.TimeoutsAttributes {
	return terra.ReferenceAsSingle[expressroutecircuitpeering.TimeoutsAttributes](ercp.ref.Append("timeouts"))
}

type expressRouteCircuitPeeringState struct {
	AzureAsn                   float64                                                  `json:"azure_asn"`
	ExpressRouteCircuitName    string                                                   `json:"express_route_circuit_name"`
	GatewayManagerEtag         string                                                   `json:"gateway_manager_etag"`
	Id                         string                                                   `json:"id"`
	Ipv4Enabled                bool                                                     `json:"ipv4_enabled"`
	PeerAsn                    float64                                                  `json:"peer_asn"`
	PeeringType                string                                                   `json:"peering_type"`
	PrimaryAzurePort           string                                                   `json:"primary_azure_port"`
	PrimaryPeerAddressPrefix   string                                                   `json:"primary_peer_address_prefix"`
	ResourceGroupName          string                                                   `json:"resource_group_name"`
	RouteFilterId              string                                                   `json:"route_filter_id"`
	SecondaryAzurePort         string                                                   `json:"secondary_azure_port"`
	SecondaryPeerAddressPrefix string                                                   `json:"secondary_peer_address_prefix"`
	SharedKey                  string                                                   `json:"shared_key"`
	VlanId                     float64                                                  `json:"vlan_id"`
	Ipv6                       []expressroutecircuitpeering.Ipv6State                   `json:"ipv6"`
	MicrosoftPeeringConfig     []expressroutecircuitpeering.MicrosoftPeeringConfigState `json:"microsoft_peering_config"`
	Timeouts                   *expressroutecircuitpeering.TimeoutsState                `json:"timeouts"`
}
