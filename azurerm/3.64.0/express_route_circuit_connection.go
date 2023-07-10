// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	expressroutecircuitconnection "github.com/golingon/terraproviders/azurerm/3.64.0/expressroutecircuitconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewExpressRouteCircuitConnection creates a new instance of [ExpressRouteCircuitConnection].
func NewExpressRouteCircuitConnection(name string, args ExpressRouteCircuitConnectionArgs) *ExpressRouteCircuitConnection {
	return &ExpressRouteCircuitConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ExpressRouteCircuitConnection)(nil)

// ExpressRouteCircuitConnection represents the Terraform resource azurerm_express_route_circuit_connection.
type ExpressRouteCircuitConnection struct {
	Name      string
	Args      ExpressRouteCircuitConnectionArgs
	state     *expressRouteCircuitConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ExpressRouteCircuitConnection].
func (ercc *ExpressRouteCircuitConnection) Type() string {
	return "azurerm_express_route_circuit_connection"
}

// LocalName returns the local name for [ExpressRouteCircuitConnection].
func (ercc *ExpressRouteCircuitConnection) LocalName() string {
	return ercc.Name
}

// Configuration returns the configuration (args) for [ExpressRouteCircuitConnection].
func (ercc *ExpressRouteCircuitConnection) Configuration() interface{} {
	return ercc.Args
}

// DependOn is used for other resources to depend on [ExpressRouteCircuitConnection].
func (ercc *ExpressRouteCircuitConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(ercc)
}

// Dependencies returns the list of resources [ExpressRouteCircuitConnection] depends_on.
func (ercc *ExpressRouteCircuitConnection) Dependencies() terra.Dependencies {
	return ercc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ExpressRouteCircuitConnection].
func (ercc *ExpressRouteCircuitConnection) LifecycleManagement() *terra.Lifecycle {
	return ercc.Lifecycle
}

// Attributes returns the attributes for [ExpressRouteCircuitConnection].
func (ercc *ExpressRouteCircuitConnection) Attributes() expressRouteCircuitConnectionAttributes {
	return expressRouteCircuitConnectionAttributes{ref: terra.ReferenceResource(ercc)}
}

// ImportState imports the given attribute values into [ExpressRouteCircuitConnection]'s state.
func (ercc *ExpressRouteCircuitConnection) ImportState(av io.Reader) error {
	ercc.state = &expressRouteCircuitConnectionState{}
	if err := json.NewDecoder(av).Decode(ercc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ercc.Type(), ercc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ExpressRouteCircuitConnection] has state.
func (ercc *ExpressRouteCircuitConnection) State() (*expressRouteCircuitConnectionState, bool) {
	return ercc.state, ercc.state != nil
}

// StateMust returns the state for [ExpressRouteCircuitConnection]. Panics if the state is nil.
func (ercc *ExpressRouteCircuitConnection) StateMust() *expressRouteCircuitConnectionState {
	if ercc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ercc.Type(), ercc.LocalName()))
	}
	return ercc.state
}

// ExpressRouteCircuitConnectionArgs contains the configurations for azurerm_express_route_circuit_connection.
type ExpressRouteCircuitConnectionArgs struct {
	// AddressPrefixIpv4: string, required
	AddressPrefixIpv4 terra.StringValue `hcl:"address_prefix_ipv4,attr" validate:"required"`
	// AddressPrefixIpv6: string, optional
	AddressPrefixIpv6 terra.StringValue `hcl:"address_prefix_ipv6,attr"`
	// AuthorizationKey: string, optional
	AuthorizationKey terra.StringValue `hcl:"authorization_key,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PeerPeeringId: string, required
	PeerPeeringId terra.StringValue `hcl:"peer_peering_id,attr" validate:"required"`
	// PeeringId: string, required
	PeeringId terra.StringValue `hcl:"peering_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *expressroutecircuitconnection.Timeouts `hcl:"timeouts,block"`
}
type expressRouteCircuitConnectionAttributes struct {
	ref terra.Reference
}

// AddressPrefixIpv4 returns a reference to field address_prefix_ipv4 of azurerm_express_route_circuit_connection.
func (ercc expressRouteCircuitConnectionAttributes) AddressPrefixIpv4() terra.StringValue {
	return terra.ReferenceAsString(ercc.ref.Append("address_prefix_ipv4"))
}

// AddressPrefixIpv6 returns a reference to field address_prefix_ipv6 of azurerm_express_route_circuit_connection.
func (ercc expressRouteCircuitConnectionAttributes) AddressPrefixIpv6() terra.StringValue {
	return terra.ReferenceAsString(ercc.ref.Append("address_prefix_ipv6"))
}

// AuthorizationKey returns a reference to field authorization_key of azurerm_express_route_circuit_connection.
func (ercc expressRouteCircuitConnectionAttributes) AuthorizationKey() terra.StringValue {
	return terra.ReferenceAsString(ercc.ref.Append("authorization_key"))
}

// Id returns a reference to field id of azurerm_express_route_circuit_connection.
func (ercc expressRouteCircuitConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ercc.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_express_route_circuit_connection.
func (ercc expressRouteCircuitConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ercc.ref.Append("name"))
}

// PeerPeeringId returns a reference to field peer_peering_id of azurerm_express_route_circuit_connection.
func (ercc expressRouteCircuitConnectionAttributes) PeerPeeringId() terra.StringValue {
	return terra.ReferenceAsString(ercc.ref.Append("peer_peering_id"))
}

// PeeringId returns a reference to field peering_id of azurerm_express_route_circuit_connection.
func (ercc expressRouteCircuitConnectionAttributes) PeeringId() terra.StringValue {
	return terra.ReferenceAsString(ercc.ref.Append("peering_id"))
}

func (ercc expressRouteCircuitConnectionAttributes) Timeouts() expressroutecircuitconnection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[expressroutecircuitconnection.TimeoutsAttributes](ercc.ref.Append("timeouts"))
}

type expressRouteCircuitConnectionState struct {
	AddressPrefixIpv4 string                                       `json:"address_prefix_ipv4"`
	AddressPrefixIpv6 string                                       `json:"address_prefix_ipv6"`
	AuthorizationKey  string                                       `json:"authorization_key"`
	Id                string                                       `json:"id"`
	Name              string                                       `json:"name"`
	PeerPeeringId     string                                       `json:"peer_peering_id"`
	PeeringId         string                                       `json:"peering_id"`
	Timeouts          *expressroutecircuitconnection.TimeoutsState `json:"timeouts"`
}
