// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	expressroutegateway "github.com/golingon/terraproviders/azurerm/3.58.0/expressroutegateway"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewExpressRouteGateway creates a new instance of [ExpressRouteGateway].
func NewExpressRouteGateway(name string, args ExpressRouteGatewayArgs) *ExpressRouteGateway {
	return &ExpressRouteGateway{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ExpressRouteGateway)(nil)

// ExpressRouteGateway represents the Terraform resource azurerm_express_route_gateway.
type ExpressRouteGateway struct {
	Name      string
	Args      ExpressRouteGatewayArgs
	state     *expressRouteGatewayState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ExpressRouteGateway].
func (erg *ExpressRouteGateway) Type() string {
	return "azurerm_express_route_gateway"
}

// LocalName returns the local name for [ExpressRouteGateway].
func (erg *ExpressRouteGateway) LocalName() string {
	return erg.Name
}

// Configuration returns the configuration (args) for [ExpressRouteGateway].
func (erg *ExpressRouteGateway) Configuration() interface{} {
	return erg.Args
}

// DependOn is used for other resources to depend on [ExpressRouteGateway].
func (erg *ExpressRouteGateway) DependOn() terra.Reference {
	return terra.ReferenceResource(erg)
}

// Dependencies returns the list of resources [ExpressRouteGateway] depends_on.
func (erg *ExpressRouteGateway) Dependencies() terra.Dependencies {
	return erg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ExpressRouteGateway].
func (erg *ExpressRouteGateway) LifecycleManagement() *terra.Lifecycle {
	return erg.Lifecycle
}

// Attributes returns the attributes for [ExpressRouteGateway].
func (erg *ExpressRouteGateway) Attributes() expressRouteGatewayAttributes {
	return expressRouteGatewayAttributes{ref: terra.ReferenceResource(erg)}
}

// ImportState imports the given attribute values into [ExpressRouteGateway]'s state.
func (erg *ExpressRouteGateway) ImportState(av io.Reader) error {
	erg.state = &expressRouteGatewayState{}
	if err := json.NewDecoder(av).Decode(erg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", erg.Type(), erg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ExpressRouteGateway] has state.
func (erg *ExpressRouteGateway) State() (*expressRouteGatewayState, bool) {
	return erg.state, erg.state != nil
}

// StateMust returns the state for [ExpressRouteGateway]. Panics if the state is nil.
func (erg *ExpressRouteGateway) StateMust() *expressRouteGatewayState {
	if erg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", erg.Type(), erg.LocalName()))
	}
	return erg.state
}

// ExpressRouteGatewayArgs contains the configurations for azurerm_express_route_gateway.
type ExpressRouteGatewayArgs struct {
	// AllowNonVirtualWanTraffic: bool, optional
	AllowNonVirtualWanTraffic terra.BoolValue `hcl:"allow_non_virtual_wan_traffic,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ScaleUnits: number, required
	ScaleUnits terra.NumberValue `hcl:"scale_units,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// VirtualHubId: string, required
	VirtualHubId terra.StringValue `hcl:"virtual_hub_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *expressroutegateway.Timeouts `hcl:"timeouts,block"`
}
type expressRouteGatewayAttributes struct {
	ref terra.Reference
}

// AllowNonVirtualWanTraffic returns a reference to field allow_non_virtual_wan_traffic of azurerm_express_route_gateway.
func (erg expressRouteGatewayAttributes) AllowNonVirtualWanTraffic() terra.BoolValue {
	return terra.ReferenceAsBool(erg.ref.Append("allow_non_virtual_wan_traffic"))
}

// Id returns a reference to field id of azurerm_express_route_gateway.
func (erg expressRouteGatewayAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(erg.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_express_route_gateway.
func (erg expressRouteGatewayAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(erg.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_express_route_gateway.
func (erg expressRouteGatewayAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(erg.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_express_route_gateway.
func (erg expressRouteGatewayAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(erg.ref.Append("resource_group_name"))
}

// ScaleUnits returns a reference to field scale_units of azurerm_express_route_gateway.
func (erg expressRouteGatewayAttributes) ScaleUnits() terra.NumberValue {
	return terra.ReferenceAsNumber(erg.ref.Append("scale_units"))
}

// Tags returns a reference to field tags of azurerm_express_route_gateway.
func (erg expressRouteGatewayAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](erg.ref.Append("tags"))
}

// VirtualHubId returns a reference to field virtual_hub_id of azurerm_express_route_gateway.
func (erg expressRouteGatewayAttributes) VirtualHubId() terra.StringValue {
	return terra.ReferenceAsString(erg.ref.Append("virtual_hub_id"))
}

func (erg expressRouteGatewayAttributes) Timeouts() expressroutegateway.TimeoutsAttributes {
	return terra.ReferenceAsSingle[expressroutegateway.TimeoutsAttributes](erg.ref.Append("timeouts"))
}

type expressRouteGatewayState struct {
	AllowNonVirtualWanTraffic bool                               `json:"allow_non_virtual_wan_traffic"`
	Id                        string                             `json:"id"`
	Location                  string                             `json:"location"`
	Name                      string                             `json:"name"`
	ResourceGroupName         string                             `json:"resource_group_name"`
	ScaleUnits                float64                            `json:"scale_units"`
	Tags                      map[string]string                  `json:"tags"`
	VirtualHubId              string                             `json:"virtual_hub_id"`
	Timeouts                  *expressroutegateway.TimeoutsState `json:"timeouts"`
}