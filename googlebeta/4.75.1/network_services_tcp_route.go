// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	networkservicestcproute "github.com/golingon/terraproviders/googlebeta/4.75.1/networkservicestcproute"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkServicesTcpRoute creates a new instance of [NetworkServicesTcpRoute].
func NewNetworkServicesTcpRoute(name string, args NetworkServicesTcpRouteArgs) *NetworkServicesTcpRoute {
	return &NetworkServicesTcpRoute{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkServicesTcpRoute)(nil)

// NetworkServicesTcpRoute represents the Terraform resource google_network_services_tcp_route.
type NetworkServicesTcpRoute struct {
	Name      string
	Args      NetworkServicesTcpRouteArgs
	state     *networkServicesTcpRouteState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkServicesTcpRoute].
func (nstr *NetworkServicesTcpRoute) Type() string {
	return "google_network_services_tcp_route"
}

// LocalName returns the local name for [NetworkServicesTcpRoute].
func (nstr *NetworkServicesTcpRoute) LocalName() string {
	return nstr.Name
}

// Configuration returns the configuration (args) for [NetworkServicesTcpRoute].
func (nstr *NetworkServicesTcpRoute) Configuration() interface{} {
	return nstr.Args
}

// DependOn is used for other resources to depend on [NetworkServicesTcpRoute].
func (nstr *NetworkServicesTcpRoute) DependOn() terra.Reference {
	return terra.ReferenceResource(nstr)
}

// Dependencies returns the list of resources [NetworkServicesTcpRoute] depends_on.
func (nstr *NetworkServicesTcpRoute) Dependencies() terra.Dependencies {
	return nstr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkServicesTcpRoute].
func (nstr *NetworkServicesTcpRoute) LifecycleManagement() *terra.Lifecycle {
	return nstr.Lifecycle
}

// Attributes returns the attributes for [NetworkServicesTcpRoute].
func (nstr *NetworkServicesTcpRoute) Attributes() networkServicesTcpRouteAttributes {
	return networkServicesTcpRouteAttributes{ref: terra.ReferenceResource(nstr)}
}

// ImportState imports the given attribute values into [NetworkServicesTcpRoute]'s state.
func (nstr *NetworkServicesTcpRoute) ImportState(av io.Reader) error {
	nstr.state = &networkServicesTcpRouteState{}
	if err := json.NewDecoder(av).Decode(nstr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nstr.Type(), nstr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkServicesTcpRoute] has state.
func (nstr *NetworkServicesTcpRoute) State() (*networkServicesTcpRouteState, bool) {
	return nstr.state, nstr.state != nil
}

// StateMust returns the state for [NetworkServicesTcpRoute]. Panics if the state is nil.
func (nstr *NetworkServicesTcpRoute) StateMust() *networkServicesTcpRouteState {
	if nstr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nstr.Type(), nstr.LocalName()))
	}
	return nstr.state
}

// NetworkServicesTcpRouteArgs contains the configurations for google_network_services_tcp_route.
type NetworkServicesTcpRouteArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Gateways: list of string, optional
	Gateways terra.ListValue[terra.StringValue] `hcl:"gateways,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Meshes: list of string, optional
	Meshes terra.ListValue[terra.StringValue] `hcl:"meshes,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Rules: min=1
	Rules []networkservicestcproute.Rules `hcl:"rules,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *networkservicestcproute.Timeouts `hcl:"timeouts,block"`
}
type networkServicesTcpRouteAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_network_services_tcp_route.
func (nstr networkServicesTcpRouteAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(nstr.ref.Append("create_time"))
}

// Description returns a reference to field description of google_network_services_tcp_route.
func (nstr networkServicesTcpRouteAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(nstr.ref.Append("description"))
}

// Gateways returns a reference to field gateways of google_network_services_tcp_route.
func (nstr networkServicesTcpRouteAttributes) Gateways() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nstr.ref.Append("gateways"))
}

// Id returns a reference to field id of google_network_services_tcp_route.
func (nstr networkServicesTcpRouteAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nstr.ref.Append("id"))
}

// Labels returns a reference to field labels of google_network_services_tcp_route.
func (nstr networkServicesTcpRouteAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nstr.ref.Append("labels"))
}

// Meshes returns a reference to field meshes of google_network_services_tcp_route.
func (nstr networkServicesTcpRouteAttributes) Meshes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nstr.ref.Append("meshes"))
}

// Name returns a reference to field name of google_network_services_tcp_route.
func (nstr networkServicesTcpRouteAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nstr.ref.Append("name"))
}

// Project returns a reference to field project of google_network_services_tcp_route.
func (nstr networkServicesTcpRouteAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(nstr.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_network_services_tcp_route.
func (nstr networkServicesTcpRouteAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(nstr.ref.Append("self_link"))
}

// UpdateTime returns a reference to field update_time of google_network_services_tcp_route.
func (nstr networkServicesTcpRouteAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(nstr.ref.Append("update_time"))
}

func (nstr networkServicesTcpRouteAttributes) Rules() terra.ListValue[networkservicestcproute.RulesAttributes] {
	return terra.ReferenceAsList[networkservicestcproute.RulesAttributes](nstr.ref.Append("rules"))
}

func (nstr networkServicesTcpRouteAttributes) Timeouts() networkservicestcproute.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkservicestcproute.TimeoutsAttributes](nstr.ref.Append("timeouts"))
}

type networkServicesTcpRouteState struct {
	CreateTime  string                                 `json:"create_time"`
	Description string                                 `json:"description"`
	Gateways    []string                               `json:"gateways"`
	Id          string                                 `json:"id"`
	Labels      map[string]string                      `json:"labels"`
	Meshes      []string                               `json:"meshes"`
	Name        string                                 `json:"name"`
	Project     string                                 `json:"project"`
	SelfLink    string                                 `json:"self_link"`
	UpdateTime  string                                 `json:"update_time"`
	Rules       []networkservicestcproute.RulesState   `json:"rules"`
	Timeouts    *networkservicestcproute.TimeoutsState `json:"timeouts"`
}