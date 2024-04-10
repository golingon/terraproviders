// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	networkservicesgrpcroute "github.com/golingon/terraproviders/googlebeta/5.24.0/networkservicesgrpcroute"
	"io"
)

// NewNetworkServicesGrpcRoute creates a new instance of [NetworkServicesGrpcRoute].
func NewNetworkServicesGrpcRoute(name string, args NetworkServicesGrpcRouteArgs) *NetworkServicesGrpcRoute {
	return &NetworkServicesGrpcRoute{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkServicesGrpcRoute)(nil)

// NetworkServicesGrpcRoute represents the Terraform resource google_network_services_grpc_route.
type NetworkServicesGrpcRoute struct {
	Name      string
	Args      NetworkServicesGrpcRouteArgs
	state     *networkServicesGrpcRouteState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkServicesGrpcRoute].
func (nsgr *NetworkServicesGrpcRoute) Type() string {
	return "google_network_services_grpc_route"
}

// LocalName returns the local name for [NetworkServicesGrpcRoute].
func (nsgr *NetworkServicesGrpcRoute) LocalName() string {
	return nsgr.Name
}

// Configuration returns the configuration (args) for [NetworkServicesGrpcRoute].
func (nsgr *NetworkServicesGrpcRoute) Configuration() interface{} {
	return nsgr.Args
}

// DependOn is used for other resources to depend on [NetworkServicesGrpcRoute].
func (nsgr *NetworkServicesGrpcRoute) DependOn() terra.Reference {
	return terra.ReferenceResource(nsgr)
}

// Dependencies returns the list of resources [NetworkServicesGrpcRoute] depends_on.
func (nsgr *NetworkServicesGrpcRoute) Dependencies() terra.Dependencies {
	return nsgr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkServicesGrpcRoute].
func (nsgr *NetworkServicesGrpcRoute) LifecycleManagement() *terra.Lifecycle {
	return nsgr.Lifecycle
}

// Attributes returns the attributes for [NetworkServicesGrpcRoute].
func (nsgr *NetworkServicesGrpcRoute) Attributes() networkServicesGrpcRouteAttributes {
	return networkServicesGrpcRouteAttributes{ref: terra.ReferenceResource(nsgr)}
}

// ImportState imports the given attribute values into [NetworkServicesGrpcRoute]'s state.
func (nsgr *NetworkServicesGrpcRoute) ImportState(av io.Reader) error {
	nsgr.state = &networkServicesGrpcRouteState{}
	if err := json.NewDecoder(av).Decode(nsgr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nsgr.Type(), nsgr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkServicesGrpcRoute] has state.
func (nsgr *NetworkServicesGrpcRoute) State() (*networkServicesGrpcRouteState, bool) {
	return nsgr.state, nsgr.state != nil
}

// StateMust returns the state for [NetworkServicesGrpcRoute]. Panics if the state is nil.
func (nsgr *NetworkServicesGrpcRoute) StateMust() *networkServicesGrpcRouteState {
	if nsgr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nsgr.Type(), nsgr.LocalName()))
	}
	return nsgr.state
}

// NetworkServicesGrpcRouteArgs contains the configurations for google_network_services_grpc_route.
type NetworkServicesGrpcRouteArgs struct {
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
	Rules []networkservicesgrpcroute.Rules `hcl:"rules,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *networkservicesgrpcroute.Timeouts `hcl:"timeouts,block"`
}
type networkServicesGrpcRouteAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_network_services_grpc_route.
func (nsgr networkServicesGrpcRouteAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(nsgr.ref.Append("create_time"))
}

// Description returns a reference to field description of google_network_services_grpc_route.
func (nsgr networkServicesGrpcRouteAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(nsgr.ref.Append("description"))
}

// EffectiveLabels returns a reference to field effective_labels of google_network_services_grpc_route.
func (nsgr networkServicesGrpcRouteAttributes) EffectiveLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nsgr.ref.Append("effective_labels"))
}

// Gateways returns a reference to field gateways of google_network_services_grpc_route.
func (nsgr networkServicesGrpcRouteAttributes) Gateways() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nsgr.ref.Append("gateways"))
}

// Hostnames returns a reference to field hostnames of google_network_services_grpc_route.
func (nsgr networkServicesGrpcRouteAttributes) Hostnames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nsgr.ref.Append("hostnames"))
}

// Id returns a reference to field id of google_network_services_grpc_route.
func (nsgr networkServicesGrpcRouteAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nsgr.ref.Append("id"))
}

// Labels returns a reference to field labels of google_network_services_grpc_route.
func (nsgr networkServicesGrpcRouteAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nsgr.ref.Append("labels"))
}

// Meshes returns a reference to field meshes of google_network_services_grpc_route.
func (nsgr networkServicesGrpcRouteAttributes) Meshes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nsgr.ref.Append("meshes"))
}

// Name returns a reference to field name of google_network_services_grpc_route.
func (nsgr networkServicesGrpcRouteAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nsgr.ref.Append("name"))
}

// Project returns a reference to field project of google_network_services_grpc_route.
func (nsgr networkServicesGrpcRouteAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(nsgr.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_network_services_grpc_route.
func (nsgr networkServicesGrpcRouteAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(nsgr.ref.Append("self_link"))
}

// TerraformLabels returns a reference to field terraform_labels of google_network_services_grpc_route.
func (nsgr networkServicesGrpcRouteAttributes) TerraformLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nsgr.ref.Append("terraform_labels"))
}

// UpdateTime returns a reference to field update_time of google_network_services_grpc_route.
func (nsgr networkServicesGrpcRouteAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(nsgr.ref.Append("update_time"))
}

func (nsgr networkServicesGrpcRouteAttributes) Rules() terra.ListValue[networkservicesgrpcroute.RulesAttributes] {
	return terra.ReferenceAsList[networkservicesgrpcroute.RulesAttributes](nsgr.ref.Append("rules"))
}

func (nsgr networkServicesGrpcRouteAttributes) Timeouts() networkservicesgrpcroute.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkservicesgrpcroute.TimeoutsAttributes](nsgr.ref.Append("timeouts"))
}

type networkServicesGrpcRouteState struct {
	CreateTime      string                                  `json:"create_time"`
	Description     string                                  `json:"description"`
	EffectiveLabels map[string]string                       `json:"effective_labels"`
	Gateways        []string                                `json:"gateways"`
	Hostnames       []string                                `json:"hostnames"`
	Id              string                                  `json:"id"`
	Labels          map[string]string                       `json:"labels"`
	Meshes          []string                                `json:"meshes"`
	Name            string                                  `json:"name"`
	Project         string                                  `json:"project"`
	SelfLink        string                                  `json:"self_link"`
	TerraformLabels map[string]string                       `json:"terraform_labels"`
	UpdateTime      string                                  `json:"update_time"`
	Rules           []networkservicesgrpcroute.RulesState   `json:"rules"`
	Timeouts        *networkservicesgrpcroute.TimeoutsState `json:"timeouts"`
}
