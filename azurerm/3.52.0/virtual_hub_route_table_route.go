// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	virtualhubroutetableroute "github.com/golingon/terraproviders/azurerm/3.52.0/virtualhubroutetableroute"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVirtualHubRouteTableRoute creates a new instance of [VirtualHubRouteTableRoute].
func NewVirtualHubRouteTableRoute(name string, args VirtualHubRouteTableRouteArgs) *VirtualHubRouteTableRoute {
	return &VirtualHubRouteTableRoute{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VirtualHubRouteTableRoute)(nil)

// VirtualHubRouteTableRoute represents the Terraform resource azurerm_virtual_hub_route_table_route.
type VirtualHubRouteTableRoute struct {
	Name      string
	Args      VirtualHubRouteTableRouteArgs
	state     *virtualHubRouteTableRouteState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VirtualHubRouteTableRoute].
func (vhrtr *VirtualHubRouteTableRoute) Type() string {
	return "azurerm_virtual_hub_route_table_route"
}

// LocalName returns the local name for [VirtualHubRouteTableRoute].
func (vhrtr *VirtualHubRouteTableRoute) LocalName() string {
	return vhrtr.Name
}

// Configuration returns the configuration (args) for [VirtualHubRouteTableRoute].
func (vhrtr *VirtualHubRouteTableRoute) Configuration() interface{} {
	return vhrtr.Args
}

// DependOn is used for other resources to depend on [VirtualHubRouteTableRoute].
func (vhrtr *VirtualHubRouteTableRoute) DependOn() terra.Reference {
	return terra.ReferenceResource(vhrtr)
}

// Dependencies returns the list of resources [VirtualHubRouteTableRoute] depends_on.
func (vhrtr *VirtualHubRouteTableRoute) Dependencies() terra.Dependencies {
	return vhrtr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VirtualHubRouteTableRoute].
func (vhrtr *VirtualHubRouteTableRoute) LifecycleManagement() *terra.Lifecycle {
	return vhrtr.Lifecycle
}

// Attributes returns the attributes for [VirtualHubRouteTableRoute].
func (vhrtr *VirtualHubRouteTableRoute) Attributes() virtualHubRouteTableRouteAttributes {
	return virtualHubRouteTableRouteAttributes{ref: terra.ReferenceResource(vhrtr)}
}

// ImportState imports the given attribute values into [VirtualHubRouteTableRoute]'s state.
func (vhrtr *VirtualHubRouteTableRoute) ImportState(av io.Reader) error {
	vhrtr.state = &virtualHubRouteTableRouteState{}
	if err := json.NewDecoder(av).Decode(vhrtr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vhrtr.Type(), vhrtr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VirtualHubRouteTableRoute] has state.
func (vhrtr *VirtualHubRouteTableRoute) State() (*virtualHubRouteTableRouteState, bool) {
	return vhrtr.state, vhrtr.state != nil
}

// StateMust returns the state for [VirtualHubRouteTableRoute]. Panics if the state is nil.
func (vhrtr *VirtualHubRouteTableRoute) StateMust() *virtualHubRouteTableRouteState {
	if vhrtr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vhrtr.Type(), vhrtr.LocalName()))
	}
	return vhrtr.state
}

// VirtualHubRouteTableRouteArgs contains the configurations for azurerm_virtual_hub_route_table_route.
type VirtualHubRouteTableRouteArgs struct {
	// Destinations: set of string, required
	Destinations terra.SetValue[terra.StringValue] `hcl:"destinations,attr" validate:"required"`
	// DestinationsType: string, required
	DestinationsType terra.StringValue `hcl:"destinations_type,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NextHop: string, required
	NextHop terra.StringValue `hcl:"next_hop,attr" validate:"required"`
	// NextHopType: string, optional
	NextHopType terra.StringValue `hcl:"next_hop_type,attr"`
	// RouteTableId: string, required
	RouteTableId terra.StringValue `hcl:"route_table_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *virtualhubroutetableroute.Timeouts `hcl:"timeouts,block"`
}
type virtualHubRouteTableRouteAttributes struct {
	ref terra.Reference
}

// Destinations returns a reference to field destinations of azurerm_virtual_hub_route_table_route.
func (vhrtr virtualHubRouteTableRouteAttributes) Destinations() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vhrtr.ref.Append("destinations"))
}

// DestinationsType returns a reference to field destinations_type of azurerm_virtual_hub_route_table_route.
func (vhrtr virtualHubRouteTableRouteAttributes) DestinationsType() terra.StringValue {
	return terra.ReferenceAsString(vhrtr.ref.Append("destinations_type"))
}

// Id returns a reference to field id of azurerm_virtual_hub_route_table_route.
func (vhrtr virtualHubRouteTableRouteAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vhrtr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_virtual_hub_route_table_route.
func (vhrtr virtualHubRouteTableRouteAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vhrtr.ref.Append("name"))
}

// NextHop returns a reference to field next_hop of azurerm_virtual_hub_route_table_route.
func (vhrtr virtualHubRouteTableRouteAttributes) NextHop() terra.StringValue {
	return terra.ReferenceAsString(vhrtr.ref.Append("next_hop"))
}

// NextHopType returns a reference to field next_hop_type of azurerm_virtual_hub_route_table_route.
func (vhrtr virtualHubRouteTableRouteAttributes) NextHopType() terra.StringValue {
	return terra.ReferenceAsString(vhrtr.ref.Append("next_hop_type"))
}

// RouteTableId returns a reference to field route_table_id of azurerm_virtual_hub_route_table_route.
func (vhrtr virtualHubRouteTableRouteAttributes) RouteTableId() terra.StringValue {
	return terra.ReferenceAsString(vhrtr.ref.Append("route_table_id"))
}

func (vhrtr virtualHubRouteTableRouteAttributes) Timeouts() virtualhubroutetableroute.TimeoutsAttributes {
	return terra.ReferenceAsSingle[virtualhubroutetableroute.TimeoutsAttributes](vhrtr.ref.Append("timeouts"))
}

type virtualHubRouteTableRouteState struct {
	Destinations     []string                                 `json:"destinations"`
	DestinationsType string                                   `json:"destinations_type"`
	Id               string                                   `json:"id"`
	Name             string                                   `json:"name"`
	NextHop          string                                   `json:"next_hop"`
	NextHopType      string                                   `json:"next_hop_type"`
	RouteTableId     string                                   `json:"route_table_id"`
	Timeouts         *virtualhubroutetableroute.TimeoutsState `json:"timeouts"`
}
