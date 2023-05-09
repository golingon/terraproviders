// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	networkserviceshttproute "github.com/golingon/terraproviders/googlebeta/4.63.1/networkserviceshttproute"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkServicesHttpRoute creates a new instance of [NetworkServicesHttpRoute].
func NewNetworkServicesHttpRoute(name string, args NetworkServicesHttpRouteArgs) *NetworkServicesHttpRoute {
	return &NetworkServicesHttpRoute{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkServicesHttpRoute)(nil)

// NetworkServicesHttpRoute represents the Terraform resource google_network_services_http_route.
type NetworkServicesHttpRoute struct {
	Name      string
	Args      NetworkServicesHttpRouteArgs
	state     *networkServicesHttpRouteState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkServicesHttpRoute].
func (nshr *NetworkServicesHttpRoute) Type() string {
	return "google_network_services_http_route"
}

// LocalName returns the local name for [NetworkServicesHttpRoute].
func (nshr *NetworkServicesHttpRoute) LocalName() string {
	return nshr.Name
}

// Configuration returns the configuration (args) for [NetworkServicesHttpRoute].
func (nshr *NetworkServicesHttpRoute) Configuration() interface{} {
	return nshr.Args
}

// DependOn is used for other resources to depend on [NetworkServicesHttpRoute].
func (nshr *NetworkServicesHttpRoute) DependOn() terra.Reference {
	return terra.ReferenceResource(nshr)
}

// Dependencies returns the list of resources [NetworkServicesHttpRoute] depends_on.
func (nshr *NetworkServicesHttpRoute) Dependencies() terra.Dependencies {
	return nshr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkServicesHttpRoute].
func (nshr *NetworkServicesHttpRoute) LifecycleManagement() *terra.Lifecycle {
	return nshr.Lifecycle
}

// Attributes returns the attributes for [NetworkServicesHttpRoute].
func (nshr *NetworkServicesHttpRoute) Attributes() networkServicesHttpRouteAttributes {
	return networkServicesHttpRouteAttributes{ref: terra.ReferenceResource(nshr)}
}

// ImportState imports the given attribute values into [NetworkServicesHttpRoute]'s state.
func (nshr *NetworkServicesHttpRoute) ImportState(av io.Reader) error {
	nshr.state = &networkServicesHttpRouteState{}
	if err := json.NewDecoder(av).Decode(nshr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nshr.Type(), nshr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkServicesHttpRoute] has state.
func (nshr *NetworkServicesHttpRoute) State() (*networkServicesHttpRouteState, bool) {
	return nshr.state, nshr.state != nil
}

// StateMust returns the state for [NetworkServicesHttpRoute]. Panics if the state is nil.
func (nshr *NetworkServicesHttpRoute) StateMust() *networkServicesHttpRouteState {
	if nshr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nshr.Type(), nshr.LocalName()))
	}
	return nshr.state
}

// NetworkServicesHttpRouteArgs contains the configurations for google_network_services_http_route.
type NetworkServicesHttpRouteArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Gateways: list of string, optional
	Gateways terra.ListValue[terra.StringValue] `hcl:"gateways,attr"`
	// Hostnames: list of string, required
	Hostnames terra.ListValue[terra.StringValue] `hcl:"hostnames,attr" validate:"required"`
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
	Rules []networkserviceshttproute.Rules `hcl:"rules,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *networkserviceshttproute.Timeouts `hcl:"timeouts,block"`
}
type networkServicesHttpRouteAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_network_services_http_route.
func (nshr networkServicesHttpRouteAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(nshr.ref.Append("create_time"))
}

// Description returns a reference to field description of google_network_services_http_route.
func (nshr networkServicesHttpRouteAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(nshr.ref.Append("description"))
}

// Gateways returns a reference to field gateways of google_network_services_http_route.
func (nshr networkServicesHttpRouteAttributes) Gateways() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nshr.ref.Append("gateways"))
}

// Hostnames returns a reference to field hostnames of google_network_services_http_route.
func (nshr networkServicesHttpRouteAttributes) Hostnames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nshr.ref.Append("hostnames"))
}

// Id returns a reference to field id of google_network_services_http_route.
func (nshr networkServicesHttpRouteAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nshr.ref.Append("id"))
}

// Labels returns a reference to field labels of google_network_services_http_route.
func (nshr networkServicesHttpRouteAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nshr.ref.Append("labels"))
}

// Meshes returns a reference to field meshes of google_network_services_http_route.
func (nshr networkServicesHttpRouteAttributes) Meshes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nshr.ref.Append("meshes"))
}

// Name returns a reference to field name of google_network_services_http_route.
func (nshr networkServicesHttpRouteAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nshr.ref.Append("name"))
}

// Project returns a reference to field project of google_network_services_http_route.
func (nshr networkServicesHttpRouteAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(nshr.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_network_services_http_route.
func (nshr networkServicesHttpRouteAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(nshr.ref.Append("self_link"))
}

// UpdateTime returns a reference to field update_time of google_network_services_http_route.
func (nshr networkServicesHttpRouteAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(nshr.ref.Append("update_time"))
}

func (nshr networkServicesHttpRouteAttributes) Rules() terra.ListValue[networkserviceshttproute.RulesAttributes] {
	return terra.ReferenceAsList[networkserviceshttproute.RulesAttributes](nshr.ref.Append("rules"))
}

func (nshr networkServicesHttpRouteAttributes) Timeouts() networkserviceshttproute.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkserviceshttproute.TimeoutsAttributes](nshr.ref.Append("timeouts"))
}

type networkServicesHttpRouteState struct {
	CreateTime  string                                  `json:"create_time"`
	Description string                                  `json:"description"`
	Gateways    []string                                `json:"gateways"`
	Hostnames   []string                                `json:"hostnames"`
	Id          string                                  `json:"id"`
	Labels      map[string]string                       `json:"labels"`
	Meshes      []string                                `json:"meshes"`
	Name        string                                  `json:"name"`
	Project     string                                  `json:"project"`
	SelfLink    string                                  `json:"self_link"`
	UpdateTime  string                                  `json:"update_time"`
	Rules       []networkserviceshttproute.RulesState   `json:"rules"`
	Timeouts    *networkserviceshttproute.TimeoutsState `json:"timeouts"`
}
