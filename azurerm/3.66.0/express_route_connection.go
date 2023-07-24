// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	expressrouteconnection "github.com/golingon/terraproviders/azurerm/3.66.0/expressrouteconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewExpressRouteConnection creates a new instance of [ExpressRouteConnection].
func NewExpressRouteConnection(name string, args ExpressRouteConnectionArgs) *ExpressRouteConnection {
	return &ExpressRouteConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ExpressRouteConnection)(nil)

// ExpressRouteConnection represents the Terraform resource azurerm_express_route_connection.
type ExpressRouteConnection struct {
	Name      string
	Args      ExpressRouteConnectionArgs
	state     *expressRouteConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ExpressRouteConnection].
func (erc *ExpressRouteConnection) Type() string {
	return "azurerm_express_route_connection"
}

// LocalName returns the local name for [ExpressRouteConnection].
func (erc *ExpressRouteConnection) LocalName() string {
	return erc.Name
}

// Configuration returns the configuration (args) for [ExpressRouteConnection].
func (erc *ExpressRouteConnection) Configuration() interface{} {
	return erc.Args
}

// DependOn is used for other resources to depend on [ExpressRouteConnection].
func (erc *ExpressRouteConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(erc)
}

// Dependencies returns the list of resources [ExpressRouteConnection] depends_on.
func (erc *ExpressRouteConnection) Dependencies() terra.Dependencies {
	return erc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ExpressRouteConnection].
func (erc *ExpressRouteConnection) LifecycleManagement() *terra.Lifecycle {
	return erc.Lifecycle
}

// Attributes returns the attributes for [ExpressRouteConnection].
func (erc *ExpressRouteConnection) Attributes() expressRouteConnectionAttributes {
	return expressRouteConnectionAttributes{ref: terra.ReferenceResource(erc)}
}

// ImportState imports the given attribute values into [ExpressRouteConnection]'s state.
func (erc *ExpressRouteConnection) ImportState(av io.Reader) error {
	erc.state = &expressRouteConnectionState{}
	if err := json.NewDecoder(av).Decode(erc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", erc.Type(), erc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ExpressRouteConnection] has state.
func (erc *ExpressRouteConnection) State() (*expressRouteConnectionState, bool) {
	return erc.state, erc.state != nil
}

// StateMust returns the state for [ExpressRouteConnection]. Panics if the state is nil.
func (erc *ExpressRouteConnection) StateMust() *expressRouteConnectionState {
	if erc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", erc.Type(), erc.LocalName()))
	}
	return erc.state
}

// ExpressRouteConnectionArgs contains the configurations for azurerm_express_route_connection.
type ExpressRouteConnectionArgs struct {
	// AuthorizationKey: string, optional
	AuthorizationKey terra.StringValue `hcl:"authorization_key,attr"`
	// EnableInternetSecurity: bool, optional
	EnableInternetSecurity terra.BoolValue `hcl:"enable_internet_security,attr"`
	// ExpressRouteCircuitPeeringId: string, required
	ExpressRouteCircuitPeeringId terra.StringValue `hcl:"express_route_circuit_peering_id,attr" validate:"required"`
	// ExpressRouteGatewayBypassEnabled: bool, optional
	ExpressRouteGatewayBypassEnabled terra.BoolValue `hcl:"express_route_gateway_bypass_enabled,attr"`
	// ExpressRouteGatewayId: string, required
	ExpressRouteGatewayId terra.StringValue `hcl:"express_route_gateway_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RoutingWeight: number, optional
	RoutingWeight terra.NumberValue `hcl:"routing_weight,attr"`
	// Routing: optional
	Routing *expressrouteconnection.Routing `hcl:"routing,block"`
	// Timeouts: optional
	Timeouts *expressrouteconnection.Timeouts `hcl:"timeouts,block"`
}
type expressRouteConnectionAttributes struct {
	ref terra.Reference
}

// AuthorizationKey returns a reference to field authorization_key of azurerm_express_route_connection.
func (erc expressRouteConnectionAttributes) AuthorizationKey() terra.StringValue {
	return terra.ReferenceAsString(erc.ref.Append("authorization_key"))
}

// EnableInternetSecurity returns a reference to field enable_internet_security of azurerm_express_route_connection.
func (erc expressRouteConnectionAttributes) EnableInternetSecurity() terra.BoolValue {
	return terra.ReferenceAsBool(erc.ref.Append("enable_internet_security"))
}

// ExpressRouteCircuitPeeringId returns a reference to field express_route_circuit_peering_id of azurerm_express_route_connection.
func (erc expressRouteConnectionAttributes) ExpressRouteCircuitPeeringId() terra.StringValue {
	return terra.ReferenceAsString(erc.ref.Append("express_route_circuit_peering_id"))
}

// ExpressRouteGatewayBypassEnabled returns a reference to field express_route_gateway_bypass_enabled of azurerm_express_route_connection.
func (erc expressRouteConnectionAttributes) ExpressRouteGatewayBypassEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(erc.ref.Append("express_route_gateway_bypass_enabled"))
}

// ExpressRouteGatewayId returns a reference to field express_route_gateway_id of azurerm_express_route_connection.
func (erc expressRouteConnectionAttributes) ExpressRouteGatewayId() terra.StringValue {
	return terra.ReferenceAsString(erc.ref.Append("express_route_gateway_id"))
}

// Id returns a reference to field id of azurerm_express_route_connection.
func (erc expressRouteConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(erc.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_express_route_connection.
func (erc expressRouteConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(erc.ref.Append("name"))
}

// RoutingWeight returns a reference to field routing_weight of azurerm_express_route_connection.
func (erc expressRouteConnectionAttributes) RoutingWeight() terra.NumberValue {
	return terra.ReferenceAsNumber(erc.ref.Append("routing_weight"))
}

func (erc expressRouteConnectionAttributes) Routing() terra.ListValue[expressrouteconnection.RoutingAttributes] {
	return terra.ReferenceAsList[expressrouteconnection.RoutingAttributes](erc.ref.Append("routing"))
}

func (erc expressRouteConnectionAttributes) Timeouts() expressrouteconnection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[expressrouteconnection.TimeoutsAttributes](erc.ref.Append("timeouts"))
}

type expressRouteConnectionState struct {
	AuthorizationKey                 string                                `json:"authorization_key"`
	EnableInternetSecurity           bool                                  `json:"enable_internet_security"`
	ExpressRouteCircuitPeeringId     string                                `json:"express_route_circuit_peering_id"`
	ExpressRouteGatewayBypassEnabled bool                                  `json:"express_route_gateway_bypass_enabled"`
	ExpressRouteGatewayId            string                                `json:"express_route_gateway_id"`
	Id                               string                                `json:"id"`
	Name                             string                                `json:"name"`
	RoutingWeight                    float64                               `json:"routing_weight"`
	Routing                          []expressrouteconnection.RoutingState `json:"routing"`
	Timeouts                         *expressrouteconnection.TimeoutsState `json:"timeouts"`
}
