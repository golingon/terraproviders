// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	routetable "github.com/golingon/terraproviders/azurerm/3.69.0/routetable"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRouteTable creates a new instance of [RouteTable].
func NewRouteTable(name string, args RouteTableArgs) *RouteTable {
	return &RouteTable{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RouteTable)(nil)

// RouteTable represents the Terraform resource azurerm_route_table.
type RouteTable struct {
	Name      string
	Args      RouteTableArgs
	state     *routeTableState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RouteTable].
func (rt *RouteTable) Type() string {
	return "azurerm_route_table"
}

// LocalName returns the local name for [RouteTable].
func (rt *RouteTable) LocalName() string {
	return rt.Name
}

// Configuration returns the configuration (args) for [RouteTable].
func (rt *RouteTable) Configuration() interface{} {
	return rt.Args
}

// DependOn is used for other resources to depend on [RouteTable].
func (rt *RouteTable) DependOn() terra.Reference {
	return terra.ReferenceResource(rt)
}

// Dependencies returns the list of resources [RouteTable] depends_on.
func (rt *RouteTable) Dependencies() terra.Dependencies {
	return rt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RouteTable].
func (rt *RouteTable) LifecycleManagement() *terra.Lifecycle {
	return rt.Lifecycle
}

// Attributes returns the attributes for [RouteTable].
func (rt *RouteTable) Attributes() routeTableAttributes {
	return routeTableAttributes{ref: terra.ReferenceResource(rt)}
}

// ImportState imports the given attribute values into [RouteTable]'s state.
func (rt *RouteTable) ImportState(av io.Reader) error {
	rt.state = &routeTableState{}
	if err := json.NewDecoder(av).Decode(rt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rt.Type(), rt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RouteTable] has state.
func (rt *RouteTable) State() (*routeTableState, bool) {
	return rt.state, rt.state != nil
}

// StateMust returns the state for [RouteTable]. Panics if the state is nil.
func (rt *RouteTable) StateMust() *routeTableState {
	if rt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rt.Type(), rt.LocalName()))
	}
	return rt.state
}

// RouteTableArgs contains the configurations for azurerm_route_table.
type RouteTableArgs struct {
	// DisableBgpRoutePropagation: bool, optional
	DisableBgpRoutePropagation terra.BoolValue `hcl:"disable_bgp_route_propagation,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Route: min=0
	Route []routetable.Route `hcl:"route,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *routetable.Timeouts `hcl:"timeouts,block"`
}
type routeTableAttributes struct {
	ref terra.Reference
}

// DisableBgpRoutePropagation returns a reference to field disable_bgp_route_propagation of azurerm_route_table.
func (rt routeTableAttributes) DisableBgpRoutePropagation() terra.BoolValue {
	return terra.ReferenceAsBool(rt.ref.Append("disable_bgp_route_propagation"))
}

// Id returns a reference to field id of azurerm_route_table.
func (rt routeTableAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rt.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_route_table.
func (rt routeTableAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(rt.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_route_table.
func (rt routeTableAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rt.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_route_table.
func (rt routeTableAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(rt.ref.Append("resource_group_name"))
}

// Subnets returns a reference to field subnets of azurerm_route_table.
func (rt routeTableAttributes) Subnets() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rt.ref.Append("subnets"))
}

// Tags returns a reference to field tags of azurerm_route_table.
func (rt routeTableAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rt.ref.Append("tags"))
}

func (rt routeTableAttributes) Route() terra.SetValue[routetable.RouteAttributes] {
	return terra.ReferenceAsSet[routetable.RouteAttributes](rt.ref.Append("route"))
}

func (rt routeTableAttributes) Timeouts() routetable.TimeoutsAttributes {
	return terra.ReferenceAsSingle[routetable.TimeoutsAttributes](rt.ref.Append("timeouts"))
}

type routeTableState struct {
	DisableBgpRoutePropagation bool                      `json:"disable_bgp_route_propagation"`
	Id                         string                    `json:"id"`
	Location                   string                    `json:"location"`
	Name                       string                    `json:"name"`
	ResourceGroupName          string                    `json:"resource_group_name"`
	Subnets                    []string                  `json:"subnets"`
	Tags                       map[string]string         `json:"tags"`
	Route                      []routetable.RouteState   `json:"route"`
	Timeouts                   *routetable.TimeoutsState `json:"timeouts"`
}
