// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	expressroutecircuitauthorization "github.com/golingon/terraproviders/azurerm/3.65.0/expressroutecircuitauthorization"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewExpressRouteCircuitAuthorization creates a new instance of [ExpressRouteCircuitAuthorization].
func NewExpressRouteCircuitAuthorization(name string, args ExpressRouteCircuitAuthorizationArgs) *ExpressRouteCircuitAuthorization {
	return &ExpressRouteCircuitAuthorization{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ExpressRouteCircuitAuthorization)(nil)

// ExpressRouteCircuitAuthorization represents the Terraform resource azurerm_express_route_circuit_authorization.
type ExpressRouteCircuitAuthorization struct {
	Name      string
	Args      ExpressRouteCircuitAuthorizationArgs
	state     *expressRouteCircuitAuthorizationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ExpressRouteCircuitAuthorization].
func (erca *ExpressRouteCircuitAuthorization) Type() string {
	return "azurerm_express_route_circuit_authorization"
}

// LocalName returns the local name for [ExpressRouteCircuitAuthorization].
func (erca *ExpressRouteCircuitAuthorization) LocalName() string {
	return erca.Name
}

// Configuration returns the configuration (args) for [ExpressRouteCircuitAuthorization].
func (erca *ExpressRouteCircuitAuthorization) Configuration() interface{} {
	return erca.Args
}

// DependOn is used for other resources to depend on [ExpressRouteCircuitAuthorization].
func (erca *ExpressRouteCircuitAuthorization) DependOn() terra.Reference {
	return terra.ReferenceResource(erca)
}

// Dependencies returns the list of resources [ExpressRouteCircuitAuthorization] depends_on.
func (erca *ExpressRouteCircuitAuthorization) Dependencies() terra.Dependencies {
	return erca.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ExpressRouteCircuitAuthorization].
func (erca *ExpressRouteCircuitAuthorization) LifecycleManagement() *terra.Lifecycle {
	return erca.Lifecycle
}

// Attributes returns the attributes for [ExpressRouteCircuitAuthorization].
func (erca *ExpressRouteCircuitAuthorization) Attributes() expressRouteCircuitAuthorizationAttributes {
	return expressRouteCircuitAuthorizationAttributes{ref: terra.ReferenceResource(erca)}
}

// ImportState imports the given attribute values into [ExpressRouteCircuitAuthorization]'s state.
func (erca *ExpressRouteCircuitAuthorization) ImportState(av io.Reader) error {
	erca.state = &expressRouteCircuitAuthorizationState{}
	if err := json.NewDecoder(av).Decode(erca.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", erca.Type(), erca.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ExpressRouteCircuitAuthorization] has state.
func (erca *ExpressRouteCircuitAuthorization) State() (*expressRouteCircuitAuthorizationState, bool) {
	return erca.state, erca.state != nil
}

// StateMust returns the state for [ExpressRouteCircuitAuthorization]. Panics if the state is nil.
func (erca *ExpressRouteCircuitAuthorization) StateMust() *expressRouteCircuitAuthorizationState {
	if erca.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", erca.Type(), erca.LocalName()))
	}
	return erca.state
}

// ExpressRouteCircuitAuthorizationArgs contains the configurations for azurerm_express_route_circuit_authorization.
type ExpressRouteCircuitAuthorizationArgs struct {
	// ExpressRouteCircuitName: string, required
	ExpressRouteCircuitName terra.StringValue `hcl:"express_route_circuit_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *expressroutecircuitauthorization.Timeouts `hcl:"timeouts,block"`
}
type expressRouteCircuitAuthorizationAttributes struct {
	ref terra.Reference
}

// AuthorizationKey returns a reference to field authorization_key of azurerm_express_route_circuit_authorization.
func (erca expressRouteCircuitAuthorizationAttributes) AuthorizationKey() terra.StringValue {
	return terra.ReferenceAsString(erca.ref.Append("authorization_key"))
}

// AuthorizationUseStatus returns a reference to field authorization_use_status of azurerm_express_route_circuit_authorization.
func (erca expressRouteCircuitAuthorizationAttributes) AuthorizationUseStatus() terra.StringValue {
	return terra.ReferenceAsString(erca.ref.Append("authorization_use_status"))
}

// ExpressRouteCircuitName returns a reference to field express_route_circuit_name of azurerm_express_route_circuit_authorization.
func (erca expressRouteCircuitAuthorizationAttributes) ExpressRouteCircuitName() terra.StringValue {
	return terra.ReferenceAsString(erca.ref.Append("express_route_circuit_name"))
}

// Id returns a reference to field id of azurerm_express_route_circuit_authorization.
func (erca expressRouteCircuitAuthorizationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(erca.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_express_route_circuit_authorization.
func (erca expressRouteCircuitAuthorizationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(erca.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_express_route_circuit_authorization.
func (erca expressRouteCircuitAuthorizationAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(erca.ref.Append("resource_group_name"))
}

func (erca expressRouteCircuitAuthorizationAttributes) Timeouts() expressroutecircuitauthorization.TimeoutsAttributes {
	return terra.ReferenceAsSingle[expressroutecircuitauthorization.TimeoutsAttributes](erca.ref.Append("timeouts"))
}

type expressRouteCircuitAuthorizationState struct {
	AuthorizationKey        string                                          `json:"authorization_key"`
	AuthorizationUseStatus  string                                          `json:"authorization_use_status"`
	ExpressRouteCircuitName string                                          `json:"express_route_circuit_name"`
	Id                      string                                          `json:"id"`
	Name                    string                                          `json:"name"`
	ResourceGroupName       string                                          `json:"resource_group_name"`
	Timeouts                *expressroutecircuitauthorization.TimeoutsState `json:"timeouts"`
}