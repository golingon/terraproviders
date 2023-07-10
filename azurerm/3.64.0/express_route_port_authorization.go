// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	expressrouteportauthorization "github.com/golingon/terraproviders/azurerm/3.64.0/expressrouteportauthorization"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewExpressRoutePortAuthorization creates a new instance of [ExpressRoutePortAuthorization].
func NewExpressRoutePortAuthorization(name string, args ExpressRoutePortAuthorizationArgs) *ExpressRoutePortAuthorization {
	return &ExpressRoutePortAuthorization{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ExpressRoutePortAuthorization)(nil)

// ExpressRoutePortAuthorization represents the Terraform resource azurerm_express_route_port_authorization.
type ExpressRoutePortAuthorization struct {
	Name      string
	Args      ExpressRoutePortAuthorizationArgs
	state     *expressRoutePortAuthorizationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ExpressRoutePortAuthorization].
func (erpa *ExpressRoutePortAuthorization) Type() string {
	return "azurerm_express_route_port_authorization"
}

// LocalName returns the local name for [ExpressRoutePortAuthorization].
func (erpa *ExpressRoutePortAuthorization) LocalName() string {
	return erpa.Name
}

// Configuration returns the configuration (args) for [ExpressRoutePortAuthorization].
func (erpa *ExpressRoutePortAuthorization) Configuration() interface{} {
	return erpa.Args
}

// DependOn is used for other resources to depend on [ExpressRoutePortAuthorization].
func (erpa *ExpressRoutePortAuthorization) DependOn() terra.Reference {
	return terra.ReferenceResource(erpa)
}

// Dependencies returns the list of resources [ExpressRoutePortAuthorization] depends_on.
func (erpa *ExpressRoutePortAuthorization) Dependencies() terra.Dependencies {
	return erpa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ExpressRoutePortAuthorization].
func (erpa *ExpressRoutePortAuthorization) LifecycleManagement() *terra.Lifecycle {
	return erpa.Lifecycle
}

// Attributes returns the attributes for [ExpressRoutePortAuthorization].
func (erpa *ExpressRoutePortAuthorization) Attributes() expressRoutePortAuthorizationAttributes {
	return expressRoutePortAuthorizationAttributes{ref: terra.ReferenceResource(erpa)}
}

// ImportState imports the given attribute values into [ExpressRoutePortAuthorization]'s state.
func (erpa *ExpressRoutePortAuthorization) ImportState(av io.Reader) error {
	erpa.state = &expressRoutePortAuthorizationState{}
	if err := json.NewDecoder(av).Decode(erpa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", erpa.Type(), erpa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ExpressRoutePortAuthorization] has state.
func (erpa *ExpressRoutePortAuthorization) State() (*expressRoutePortAuthorizationState, bool) {
	return erpa.state, erpa.state != nil
}

// StateMust returns the state for [ExpressRoutePortAuthorization]. Panics if the state is nil.
func (erpa *ExpressRoutePortAuthorization) StateMust() *expressRoutePortAuthorizationState {
	if erpa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", erpa.Type(), erpa.LocalName()))
	}
	return erpa.state
}

// ExpressRoutePortAuthorizationArgs contains the configurations for azurerm_express_route_port_authorization.
type ExpressRoutePortAuthorizationArgs struct {
	// ExpressRoutePortName: string, required
	ExpressRoutePortName terra.StringValue `hcl:"express_route_port_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *expressrouteportauthorization.Timeouts `hcl:"timeouts,block"`
}
type expressRoutePortAuthorizationAttributes struct {
	ref terra.Reference
}

// AuthorizationKey returns a reference to field authorization_key of azurerm_express_route_port_authorization.
func (erpa expressRoutePortAuthorizationAttributes) AuthorizationKey() terra.StringValue {
	return terra.ReferenceAsString(erpa.ref.Append("authorization_key"))
}

// AuthorizationUseStatus returns a reference to field authorization_use_status of azurerm_express_route_port_authorization.
func (erpa expressRoutePortAuthorizationAttributes) AuthorizationUseStatus() terra.StringValue {
	return terra.ReferenceAsString(erpa.ref.Append("authorization_use_status"))
}

// ExpressRoutePortName returns a reference to field express_route_port_name of azurerm_express_route_port_authorization.
func (erpa expressRoutePortAuthorizationAttributes) ExpressRoutePortName() terra.StringValue {
	return terra.ReferenceAsString(erpa.ref.Append("express_route_port_name"))
}

// Id returns a reference to field id of azurerm_express_route_port_authorization.
func (erpa expressRoutePortAuthorizationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(erpa.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_express_route_port_authorization.
func (erpa expressRoutePortAuthorizationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(erpa.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_express_route_port_authorization.
func (erpa expressRoutePortAuthorizationAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(erpa.ref.Append("resource_group_name"))
}

func (erpa expressRoutePortAuthorizationAttributes) Timeouts() expressrouteportauthorization.TimeoutsAttributes {
	return terra.ReferenceAsSingle[expressrouteportauthorization.TimeoutsAttributes](erpa.ref.Append("timeouts"))
}

type expressRoutePortAuthorizationState struct {
	AuthorizationKey       string                                       `json:"authorization_key"`
	AuthorizationUseStatus string                                       `json:"authorization_use_status"`
	ExpressRoutePortName   string                                       `json:"express_route_port_name"`
	Id                     string                                       `json:"id"`
	Name                   string                                       `json:"name"`
	ResourceGroupName      string                                       `json:"resource_group_name"`
	Timeouts               *expressrouteportauthorization.TimeoutsState `json:"timeouts"`
}
