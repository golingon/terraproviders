// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	route "github.com/golingon/terraproviders/azurerm/3.55.0/route"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRoute creates a new instance of [Route].
func NewRoute(name string, args RouteArgs) *Route {
	return &Route{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Route)(nil)

// Route represents the Terraform resource azurerm_route.
type Route struct {
	Name      string
	Args      RouteArgs
	state     *routeState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Route].
func (r *Route) Type() string {
	return "azurerm_route"
}

// LocalName returns the local name for [Route].
func (r *Route) LocalName() string {
	return r.Name
}

// Configuration returns the configuration (args) for [Route].
func (r *Route) Configuration() interface{} {
	return r.Args
}

// DependOn is used for other resources to depend on [Route].
func (r *Route) DependOn() terra.Reference {
	return terra.ReferenceResource(r)
}

// Dependencies returns the list of resources [Route] depends_on.
func (r *Route) Dependencies() terra.Dependencies {
	return r.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Route].
func (r *Route) LifecycleManagement() *terra.Lifecycle {
	return r.Lifecycle
}

// Attributes returns the attributes for [Route].
func (r *Route) Attributes() routeAttributes {
	return routeAttributes{ref: terra.ReferenceResource(r)}
}

// ImportState imports the given attribute values into [Route]'s state.
func (r *Route) ImportState(av io.Reader) error {
	r.state = &routeState{}
	if err := json.NewDecoder(av).Decode(r.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", r.Type(), r.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Route] has state.
func (r *Route) State() (*routeState, bool) {
	return r.state, r.state != nil
}

// StateMust returns the state for [Route]. Panics if the state is nil.
func (r *Route) StateMust() *routeState {
	if r.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", r.Type(), r.LocalName()))
	}
	return r.state
}

// RouteArgs contains the configurations for azurerm_route.
type RouteArgs struct {
	// AddressPrefix: string, required
	AddressPrefix terra.StringValue `hcl:"address_prefix,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NextHopInIpAddress: string, optional
	NextHopInIpAddress terra.StringValue `hcl:"next_hop_in_ip_address,attr"`
	// NextHopType: string, required
	NextHopType terra.StringValue `hcl:"next_hop_type,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// RouteTableName: string, required
	RouteTableName terra.StringValue `hcl:"route_table_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *route.Timeouts `hcl:"timeouts,block"`
}
type routeAttributes struct {
	ref terra.Reference
}

// AddressPrefix returns a reference to field address_prefix of azurerm_route.
func (r routeAttributes) AddressPrefix() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("address_prefix"))
}

// Id returns a reference to field id of azurerm_route.
func (r routeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_route.
func (r routeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("name"))
}

// NextHopInIpAddress returns a reference to field next_hop_in_ip_address of azurerm_route.
func (r routeAttributes) NextHopInIpAddress() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("next_hop_in_ip_address"))
}

// NextHopType returns a reference to field next_hop_type of azurerm_route.
func (r routeAttributes) NextHopType() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("next_hop_type"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_route.
func (r routeAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("resource_group_name"))
}

// RouteTableName returns a reference to field route_table_name of azurerm_route.
func (r routeAttributes) RouteTableName() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("route_table_name"))
}

func (r routeAttributes) Timeouts() route.TimeoutsAttributes {
	return terra.ReferenceAsSingle[route.TimeoutsAttributes](r.ref.Append("timeouts"))
}

type routeState struct {
	AddressPrefix      string               `json:"address_prefix"`
	Id                 string               `json:"id"`
	Name               string               `json:"name"`
	NextHopInIpAddress string               `json:"next_hop_in_ip_address"`
	NextHopType        string               `json:"next_hop_type"`
	ResourceGroupName  string               `json:"resource_group_name"`
	RouteTableName     string               `json:"route_table_name"`
	Timeouts           *route.TimeoutsState `json:"timeouts"`
}
