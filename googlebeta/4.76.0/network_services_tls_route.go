// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	networkservicestlsroute "github.com/golingon/terraproviders/googlebeta/4.76.0/networkservicestlsroute"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkServicesTlsRoute creates a new instance of [NetworkServicesTlsRoute].
func NewNetworkServicesTlsRoute(name string, args NetworkServicesTlsRouteArgs) *NetworkServicesTlsRoute {
	return &NetworkServicesTlsRoute{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkServicesTlsRoute)(nil)

// NetworkServicesTlsRoute represents the Terraform resource google_network_services_tls_route.
type NetworkServicesTlsRoute struct {
	Name      string
	Args      NetworkServicesTlsRouteArgs
	state     *networkServicesTlsRouteState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkServicesTlsRoute].
func (nstr *NetworkServicesTlsRoute) Type() string {
	return "google_network_services_tls_route"
}

// LocalName returns the local name for [NetworkServicesTlsRoute].
func (nstr *NetworkServicesTlsRoute) LocalName() string {
	return nstr.Name
}

// Configuration returns the configuration (args) for [NetworkServicesTlsRoute].
func (nstr *NetworkServicesTlsRoute) Configuration() interface{} {
	return nstr.Args
}

// DependOn is used for other resources to depend on [NetworkServicesTlsRoute].
func (nstr *NetworkServicesTlsRoute) DependOn() terra.Reference {
	return terra.ReferenceResource(nstr)
}

// Dependencies returns the list of resources [NetworkServicesTlsRoute] depends_on.
func (nstr *NetworkServicesTlsRoute) Dependencies() terra.Dependencies {
	return nstr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkServicesTlsRoute].
func (nstr *NetworkServicesTlsRoute) LifecycleManagement() *terra.Lifecycle {
	return nstr.Lifecycle
}

// Attributes returns the attributes for [NetworkServicesTlsRoute].
func (nstr *NetworkServicesTlsRoute) Attributes() networkServicesTlsRouteAttributes {
	return networkServicesTlsRouteAttributes{ref: terra.ReferenceResource(nstr)}
}

// ImportState imports the given attribute values into [NetworkServicesTlsRoute]'s state.
func (nstr *NetworkServicesTlsRoute) ImportState(av io.Reader) error {
	nstr.state = &networkServicesTlsRouteState{}
	if err := json.NewDecoder(av).Decode(nstr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nstr.Type(), nstr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkServicesTlsRoute] has state.
func (nstr *NetworkServicesTlsRoute) State() (*networkServicesTlsRouteState, bool) {
	return nstr.state, nstr.state != nil
}

// StateMust returns the state for [NetworkServicesTlsRoute]. Panics if the state is nil.
func (nstr *NetworkServicesTlsRoute) StateMust() *networkServicesTlsRouteState {
	if nstr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nstr.Type(), nstr.LocalName()))
	}
	return nstr.state
}

// NetworkServicesTlsRouteArgs contains the configurations for google_network_services_tls_route.
type NetworkServicesTlsRouteArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Gateways: list of string, optional
	Gateways terra.ListValue[terra.StringValue] `hcl:"gateways,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Meshes: list of string, optional
	Meshes terra.ListValue[terra.StringValue] `hcl:"meshes,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Rules: min=1
	Rules []networkservicestlsroute.Rules `hcl:"rules,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *networkservicestlsroute.Timeouts `hcl:"timeouts,block"`
}
type networkServicesTlsRouteAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_network_services_tls_route.
func (nstr networkServicesTlsRouteAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(nstr.ref.Append("create_time"))
}

// Description returns a reference to field description of google_network_services_tls_route.
func (nstr networkServicesTlsRouteAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(nstr.ref.Append("description"))
}

// Gateways returns a reference to field gateways of google_network_services_tls_route.
func (nstr networkServicesTlsRouteAttributes) Gateways() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nstr.ref.Append("gateways"))
}

// Id returns a reference to field id of google_network_services_tls_route.
func (nstr networkServicesTlsRouteAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nstr.ref.Append("id"))
}

// Meshes returns a reference to field meshes of google_network_services_tls_route.
func (nstr networkServicesTlsRouteAttributes) Meshes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nstr.ref.Append("meshes"))
}

// Name returns a reference to field name of google_network_services_tls_route.
func (nstr networkServicesTlsRouteAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nstr.ref.Append("name"))
}

// Project returns a reference to field project of google_network_services_tls_route.
func (nstr networkServicesTlsRouteAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(nstr.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_network_services_tls_route.
func (nstr networkServicesTlsRouteAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(nstr.ref.Append("self_link"))
}

// UpdateTime returns a reference to field update_time of google_network_services_tls_route.
func (nstr networkServicesTlsRouteAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(nstr.ref.Append("update_time"))
}

func (nstr networkServicesTlsRouteAttributes) Rules() terra.ListValue[networkservicestlsroute.RulesAttributes] {
	return terra.ReferenceAsList[networkservicestlsroute.RulesAttributes](nstr.ref.Append("rules"))
}

func (nstr networkServicesTlsRouteAttributes) Timeouts() networkservicestlsroute.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkservicestlsroute.TimeoutsAttributes](nstr.ref.Append("timeouts"))
}

type networkServicesTlsRouteState struct {
	CreateTime  string                                 `json:"create_time"`
	Description string                                 `json:"description"`
	Gateways    []string                               `json:"gateways"`
	Id          string                                 `json:"id"`
	Meshes      []string                               `json:"meshes"`
	Name        string                                 `json:"name"`
	Project     string                                 `json:"project"`
	SelfLink    string                                 `json:"self_link"`
	UpdateTime  string                                 `json:"update_time"`
	Rules       []networkservicestlsroute.RulesState   `json:"rules"`
	Timeouts    *networkservicestlsroute.TimeoutsState `json:"timeouts"`
}
