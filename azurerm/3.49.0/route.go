// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	route "github.com/golingon/terraproviders/azurerm/3.49.0/route"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewRoute(name string, args RouteArgs) *Route {
	return &Route{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Route)(nil)

type Route struct {
	Name  string
	Args  RouteArgs
	state *routeState
}

func (r *Route) Type() string {
	return "azurerm_route"
}

func (r *Route) LocalName() string {
	return r.Name
}

func (r *Route) Configuration() interface{} {
	return r.Args
}

func (r *Route) Attributes() routeAttributes {
	return routeAttributes{ref: terra.ReferenceResource(r)}
}

func (r *Route) ImportState(av io.Reader) error {
	r.state = &routeState{}
	if err := json.NewDecoder(av).Decode(r.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", r.Type(), r.LocalName(), err)
	}
	return nil
}

func (r *Route) State() (*routeState, bool) {
	return r.state, r.state != nil
}

func (r *Route) StateMust() *routeState {
	if r.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", r.Type(), r.LocalName()))
	}
	return r.state
}

func (r *Route) DependOn() terra.Reference {
	return terra.ReferenceResource(r)
}

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
	// DependsOn contains resources that Route depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type routeAttributes struct {
	ref terra.Reference
}

func (r routeAttributes) AddressPrefix() terra.StringValue {
	return terra.ReferenceString(r.ref.Append("address_prefix"))
}

func (r routeAttributes) Id() terra.StringValue {
	return terra.ReferenceString(r.ref.Append("id"))
}

func (r routeAttributes) Name() terra.StringValue {
	return terra.ReferenceString(r.ref.Append("name"))
}

func (r routeAttributes) NextHopInIpAddress() terra.StringValue {
	return terra.ReferenceString(r.ref.Append("next_hop_in_ip_address"))
}

func (r routeAttributes) NextHopType() terra.StringValue {
	return terra.ReferenceString(r.ref.Append("next_hop_type"))
}

func (r routeAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(r.ref.Append("resource_group_name"))
}

func (r routeAttributes) RouteTableName() terra.StringValue {
	return terra.ReferenceString(r.ref.Append("route_table_name"))
}

func (r routeAttributes) Timeouts() route.TimeoutsAttributes {
	return terra.ReferenceSingle[route.TimeoutsAttributes](r.ref.Append("timeouts"))
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
