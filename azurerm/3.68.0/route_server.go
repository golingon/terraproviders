// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	routeserver "github.com/golingon/terraproviders/azurerm/3.68.0/routeserver"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRouteServer creates a new instance of [RouteServer].
func NewRouteServer(name string, args RouteServerArgs) *RouteServer {
	return &RouteServer{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RouteServer)(nil)

// RouteServer represents the Terraform resource azurerm_route_server.
type RouteServer struct {
	Name      string
	Args      RouteServerArgs
	state     *routeServerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RouteServer].
func (rs *RouteServer) Type() string {
	return "azurerm_route_server"
}

// LocalName returns the local name for [RouteServer].
func (rs *RouteServer) LocalName() string {
	return rs.Name
}

// Configuration returns the configuration (args) for [RouteServer].
func (rs *RouteServer) Configuration() interface{} {
	return rs.Args
}

// DependOn is used for other resources to depend on [RouteServer].
func (rs *RouteServer) DependOn() terra.Reference {
	return terra.ReferenceResource(rs)
}

// Dependencies returns the list of resources [RouteServer] depends_on.
func (rs *RouteServer) Dependencies() terra.Dependencies {
	return rs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RouteServer].
func (rs *RouteServer) LifecycleManagement() *terra.Lifecycle {
	return rs.Lifecycle
}

// Attributes returns the attributes for [RouteServer].
func (rs *RouteServer) Attributes() routeServerAttributes {
	return routeServerAttributes{ref: terra.ReferenceResource(rs)}
}

// ImportState imports the given attribute values into [RouteServer]'s state.
func (rs *RouteServer) ImportState(av io.Reader) error {
	rs.state = &routeServerState{}
	if err := json.NewDecoder(av).Decode(rs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rs.Type(), rs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RouteServer] has state.
func (rs *RouteServer) State() (*routeServerState, bool) {
	return rs.state, rs.state != nil
}

// StateMust returns the state for [RouteServer]. Panics if the state is nil.
func (rs *RouteServer) StateMust() *routeServerState {
	if rs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rs.Type(), rs.LocalName()))
	}
	return rs.state
}

// RouteServerArgs contains the configurations for azurerm_route_server.
type RouteServerArgs struct {
	// BranchToBranchTrafficEnabled: bool, optional
	BranchToBranchTrafficEnabled terra.BoolValue `hcl:"branch_to_branch_traffic_enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PublicIpAddressId: string, required
	PublicIpAddressId terra.StringValue `hcl:"public_ip_address_id,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Sku: string, required
	Sku terra.StringValue `hcl:"sku,attr" validate:"required"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *routeserver.Timeouts `hcl:"timeouts,block"`
}
type routeServerAttributes struct {
	ref terra.Reference
}

// BranchToBranchTrafficEnabled returns a reference to field branch_to_branch_traffic_enabled of azurerm_route_server.
func (rs routeServerAttributes) BranchToBranchTrafficEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(rs.ref.Append("branch_to_branch_traffic_enabled"))
}

// Id returns a reference to field id of azurerm_route_server.
func (rs routeServerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_route_server.
func (rs routeServerAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_route_server.
func (rs routeServerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("name"))
}

// PublicIpAddressId returns a reference to field public_ip_address_id of azurerm_route_server.
func (rs routeServerAttributes) PublicIpAddressId() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("public_ip_address_id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_route_server.
func (rs routeServerAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("resource_group_name"))
}

// RoutingState returns a reference to field routing_state of azurerm_route_server.
func (rs routeServerAttributes) RoutingState() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("routing_state"))
}

// Sku returns a reference to field sku of azurerm_route_server.
func (rs routeServerAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("sku"))
}

// SubnetId returns a reference to field subnet_id of azurerm_route_server.
func (rs routeServerAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("subnet_id"))
}

// Tags returns a reference to field tags of azurerm_route_server.
func (rs routeServerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rs.ref.Append("tags"))
}

// VirtualRouterAsn returns a reference to field virtual_router_asn of azurerm_route_server.
func (rs routeServerAttributes) VirtualRouterAsn() terra.NumberValue {
	return terra.ReferenceAsNumber(rs.ref.Append("virtual_router_asn"))
}

// VirtualRouterIps returns a reference to field virtual_router_ips of azurerm_route_server.
func (rs routeServerAttributes) VirtualRouterIps() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rs.ref.Append("virtual_router_ips"))
}

func (rs routeServerAttributes) Timeouts() routeserver.TimeoutsAttributes {
	return terra.ReferenceAsSingle[routeserver.TimeoutsAttributes](rs.ref.Append("timeouts"))
}

type routeServerState struct {
	BranchToBranchTrafficEnabled bool                       `json:"branch_to_branch_traffic_enabled"`
	Id                           string                     `json:"id"`
	Location                     string                     `json:"location"`
	Name                         string                     `json:"name"`
	PublicIpAddressId            string                     `json:"public_ip_address_id"`
	ResourceGroupName            string                     `json:"resource_group_name"`
	RoutingState                 string                     `json:"routing_state"`
	Sku                          string                     `json:"sku"`
	SubnetId                     string                     `json:"subnet_id"`
	Tags                         map[string]string          `json:"tags"`
	VirtualRouterAsn             float64                    `json:"virtual_router_asn"`
	VirtualRouterIps             []string                   `json:"virtual_router_ips"`
	Timeouts                     *routeserver.TimeoutsState `json:"timeouts"`
}
