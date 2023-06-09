// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	routeserverbgpconnection "github.com/golingon/terraproviders/azurerm/3.55.0/routeserverbgpconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRouteServerBgpConnection creates a new instance of [RouteServerBgpConnection].
func NewRouteServerBgpConnection(name string, args RouteServerBgpConnectionArgs) *RouteServerBgpConnection {
	return &RouteServerBgpConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RouteServerBgpConnection)(nil)

// RouteServerBgpConnection represents the Terraform resource azurerm_route_server_bgp_connection.
type RouteServerBgpConnection struct {
	Name      string
	Args      RouteServerBgpConnectionArgs
	state     *routeServerBgpConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RouteServerBgpConnection].
func (rsbc *RouteServerBgpConnection) Type() string {
	return "azurerm_route_server_bgp_connection"
}

// LocalName returns the local name for [RouteServerBgpConnection].
func (rsbc *RouteServerBgpConnection) LocalName() string {
	return rsbc.Name
}

// Configuration returns the configuration (args) for [RouteServerBgpConnection].
func (rsbc *RouteServerBgpConnection) Configuration() interface{} {
	return rsbc.Args
}

// DependOn is used for other resources to depend on [RouteServerBgpConnection].
func (rsbc *RouteServerBgpConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(rsbc)
}

// Dependencies returns the list of resources [RouteServerBgpConnection] depends_on.
func (rsbc *RouteServerBgpConnection) Dependencies() terra.Dependencies {
	return rsbc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RouteServerBgpConnection].
func (rsbc *RouteServerBgpConnection) LifecycleManagement() *terra.Lifecycle {
	return rsbc.Lifecycle
}

// Attributes returns the attributes for [RouteServerBgpConnection].
func (rsbc *RouteServerBgpConnection) Attributes() routeServerBgpConnectionAttributes {
	return routeServerBgpConnectionAttributes{ref: terra.ReferenceResource(rsbc)}
}

// ImportState imports the given attribute values into [RouteServerBgpConnection]'s state.
func (rsbc *RouteServerBgpConnection) ImportState(av io.Reader) error {
	rsbc.state = &routeServerBgpConnectionState{}
	if err := json.NewDecoder(av).Decode(rsbc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rsbc.Type(), rsbc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RouteServerBgpConnection] has state.
func (rsbc *RouteServerBgpConnection) State() (*routeServerBgpConnectionState, bool) {
	return rsbc.state, rsbc.state != nil
}

// StateMust returns the state for [RouteServerBgpConnection]. Panics if the state is nil.
func (rsbc *RouteServerBgpConnection) StateMust() *routeServerBgpConnectionState {
	if rsbc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rsbc.Type(), rsbc.LocalName()))
	}
	return rsbc.state
}

// RouteServerBgpConnectionArgs contains the configurations for azurerm_route_server_bgp_connection.
type RouteServerBgpConnectionArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PeerAsn: number, required
	PeerAsn terra.NumberValue `hcl:"peer_asn,attr" validate:"required"`
	// PeerIp: string, required
	PeerIp terra.StringValue `hcl:"peer_ip,attr" validate:"required"`
	// RouteServerId: string, required
	RouteServerId terra.StringValue `hcl:"route_server_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *routeserverbgpconnection.Timeouts `hcl:"timeouts,block"`
}
type routeServerBgpConnectionAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_route_server_bgp_connection.
func (rsbc routeServerBgpConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rsbc.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_route_server_bgp_connection.
func (rsbc routeServerBgpConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rsbc.ref.Append("name"))
}

// PeerAsn returns a reference to field peer_asn of azurerm_route_server_bgp_connection.
func (rsbc routeServerBgpConnectionAttributes) PeerAsn() terra.NumberValue {
	return terra.ReferenceAsNumber(rsbc.ref.Append("peer_asn"))
}

// PeerIp returns a reference to field peer_ip of azurerm_route_server_bgp_connection.
func (rsbc routeServerBgpConnectionAttributes) PeerIp() terra.StringValue {
	return terra.ReferenceAsString(rsbc.ref.Append("peer_ip"))
}

// RouteServerId returns a reference to field route_server_id of azurerm_route_server_bgp_connection.
func (rsbc routeServerBgpConnectionAttributes) RouteServerId() terra.StringValue {
	return terra.ReferenceAsString(rsbc.ref.Append("route_server_id"))
}

func (rsbc routeServerBgpConnectionAttributes) Timeouts() routeserverbgpconnection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[routeserverbgpconnection.TimeoutsAttributes](rsbc.ref.Append("timeouts"))
}

type routeServerBgpConnectionState struct {
	Id            string                                  `json:"id"`
	Name          string                                  `json:"name"`
	PeerAsn       float64                                 `json:"peer_asn"`
	PeerIp        string                                  `json:"peer_ip"`
	RouteServerId string                                  `json:"route_server_id"`
	Timeouts      *routeserverbgpconnection.TimeoutsState `json:"timeouts"`
}
