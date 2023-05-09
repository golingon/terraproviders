// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	trafficmanagernestedendpoint "github.com/golingon/terraproviders/azurerm/3.55.0/trafficmanagernestedendpoint"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewTrafficManagerNestedEndpoint creates a new instance of [TrafficManagerNestedEndpoint].
func NewTrafficManagerNestedEndpoint(name string, args TrafficManagerNestedEndpointArgs) *TrafficManagerNestedEndpoint {
	return &TrafficManagerNestedEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*TrafficManagerNestedEndpoint)(nil)

// TrafficManagerNestedEndpoint represents the Terraform resource azurerm_traffic_manager_nested_endpoint.
type TrafficManagerNestedEndpoint struct {
	Name      string
	Args      TrafficManagerNestedEndpointArgs
	state     *trafficManagerNestedEndpointState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [TrafficManagerNestedEndpoint].
func (tmne *TrafficManagerNestedEndpoint) Type() string {
	return "azurerm_traffic_manager_nested_endpoint"
}

// LocalName returns the local name for [TrafficManagerNestedEndpoint].
func (tmne *TrafficManagerNestedEndpoint) LocalName() string {
	return tmne.Name
}

// Configuration returns the configuration (args) for [TrafficManagerNestedEndpoint].
func (tmne *TrafficManagerNestedEndpoint) Configuration() interface{} {
	return tmne.Args
}

// DependOn is used for other resources to depend on [TrafficManagerNestedEndpoint].
func (tmne *TrafficManagerNestedEndpoint) DependOn() terra.Reference {
	return terra.ReferenceResource(tmne)
}

// Dependencies returns the list of resources [TrafficManagerNestedEndpoint] depends_on.
func (tmne *TrafficManagerNestedEndpoint) Dependencies() terra.Dependencies {
	return tmne.DependsOn
}

// LifecycleManagement returns the lifecycle block for [TrafficManagerNestedEndpoint].
func (tmne *TrafficManagerNestedEndpoint) LifecycleManagement() *terra.Lifecycle {
	return tmne.Lifecycle
}

// Attributes returns the attributes for [TrafficManagerNestedEndpoint].
func (tmne *TrafficManagerNestedEndpoint) Attributes() trafficManagerNestedEndpointAttributes {
	return trafficManagerNestedEndpointAttributes{ref: terra.ReferenceResource(tmne)}
}

// ImportState imports the given attribute values into [TrafficManagerNestedEndpoint]'s state.
func (tmne *TrafficManagerNestedEndpoint) ImportState(av io.Reader) error {
	tmne.state = &trafficManagerNestedEndpointState{}
	if err := json.NewDecoder(av).Decode(tmne.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", tmne.Type(), tmne.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [TrafficManagerNestedEndpoint] has state.
func (tmne *TrafficManagerNestedEndpoint) State() (*trafficManagerNestedEndpointState, bool) {
	return tmne.state, tmne.state != nil
}

// StateMust returns the state for [TrafficManagerNestedEndpoint]. Panics if the state is nil.
func (tmne *TrafficManagerNestedEndpoint) StateMust() *trafficManagerNestedEndpointState {
	if tmne.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", tmne.Type(), tmne.LocalName()))
	}
	return tmne.state
}

// TrafficManagerNestedEndpointArgs contains the configurations for azurerm_traffic_manager_nested_endpoint.
type TrafficManagerNestedEndpointArgs struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// EndpointLocation: string, optional
	EndpointLocation terra.StringValue `hcl:"endpoint_location,attr"`
	// GeoMappings: list of string, optional
	GeoMappings terra.ListValue[terra.StringValue] `hcl:"geo_mappings,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MinimumChildEndpoints: number, required
	MinimumChildEndpoints terra.NumberValue `hcl:"minimum_child_endpoints,attr" validate:"required"`
	// MinimumRequiredChildEndpointsIpv4: number, optional
	MinimumRequiredChildEndpointsIpv4 terra.NumberValue `hcl:"minimum_required_child_endpoints_ipv4,attr"`
	// MinimumRequiredChildEndpointsIpv6: number, optional
	MinimumRequiredChildEndpointsIpv6 terra.NumberValue `hcl:"minimum_required_child_endpoints_ipv6,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Priority: number, optional
	Priority terra.NumberValue `hcl:"priority,attr"`
	// ProfileId: string, required
	ProfileId terra.StringValue `hcl:"profile_id,attr" validate:"required"`
	// TargetResourceId: string, required
	TargetResourceId terra.StringValue `hcl:"target_resource_id,attr" validate:"required"`
	// Weight: number, optional
	Weight terra.NumberValue `hcl:"weight,attr"`
	// CustomHeader: min=0
	CustomHeader []trafficmanagernestedendpoint.CustomHeader `hcl:"custom_header,block" validate:"min=0"`
	// Subnet: min=0
	Subnet []trafficmanagernestedendpoint.Subnet `hcl:"subnet,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *trafficmanagernestedendpoint.Timeouts `hcl:"timeouts,block"`
}
type trafficManagerNestedEndpointAttributes struct {
	ref terra.Reference
}

// Enabled returns a reference to field enabled of azurerm_traffic_manager_nested_endpoint.
func (tmne trafficManagerNestedEndpointAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(tmne.ref.Append("enabled"))
}

// EndpointLocation returns a reference to field endpoint_location of azurerm_traffic_manager_nested_endpoint.
func (tmne trafficManagerNestedEndpointAttributes) EndpointLocation() terra.StringValue {
	return terra.ReferenceAsString(tmne.ref.Append("endpoint_location"))
}

// GeoMappings returns a reference to field geo_mappings of azurerm_traffic_manager_nested_endpoint.
func (tmne trafficManagerNestedEndpointAttributes) GeoMappings() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](tmne.ref.Append("geo_mappings"))
}

// Id returns a reference to field id of azurerm_traffic_manager_nested_endpoint.
func (tmne trafficManagerNestedEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(tmne.ref.Append("id"))
}

// MinimumChildEndpoints returns a reference to field minimum_child_endpoints of azurerm_traffic_manager_nested_endpoint.
func (tmne trafficManagerNestedEndpointAttributes) MinimumChildEndpoints() terra.NumberValue {
	return terra.ReferenceAsNumber(tmne.ref.Append("minimum_child_endpoints"))
}

// MinimumRequiredChildEndpointsIpv4 returns a reference to field minimum_required_child_endpoints_ipv4 of azurerm_traffic_manager_nested_endpoint.
func (tmne trafficManagerNestedEndpointAttributes) MinimumRequiredChildEndpointsIpv4() terra.NumberValue {
	return terra.ReferenceAsNumber(tmne.ref.Append("minimum_required_child_endpoints_ipv4"))
}

// MinimumRequiredChildEndpointsIpv6 returns a reference to field minimum_required_child_endpoints_ipv6 of azurerm_traffic_manager_nested_endpoint.
func (tmne trafficManagerNestedEndpointAttributes) MinimumRequiredChildEndpointsIpv6() terra.NumberValue {
	return terra.ReferenceAsNumber(tmne.ref.Append("minimum_required_child_endpoints_ipv6"))
}

// Name returns a reference to field name of azurerm_traffic_manager_nested_endpoint.
func (tmne trafficManagerNestedEndpointAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(tmne.ref.Append("name"))
}

// Priority returns a reference to field priority of azurerm_traffic_manager_nested_endpoint.
func (tmne trafficManagerNestedEndpointAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(tmne.ref.Append("priority"))
}

// ProfileId returns a reference to field profile_id of azurerm_traffic_manager_nested_endpoint.
func (tmne trafficManagerNestedEndpointAttributes) ProfileId() terra.StringValue {
	return terra.ReferenceAsString(tmne.ref.Append("profile_id"))
}

// TargetResourceId returns a reference to field target_resource_id of azurerm_traffic_manager_nested_endpoint.
func (tmne trafficManagerNestedEndpointAttributes) TargetResourceId() terra.StringValue {
	return terra.ReferenceAsString(tmne.ref.Append("target_resource_id"))
}

// Weight returns a reference to field weight of azurerm_traffic_manager_nested_endpoint.
func (tmne trafficManagerNestedEndpointAttributes) Weight() terra.NumberValue {
	return terra.ReferenceAsNumber(tmne.ref.Append("weight"))
}

func (tmne trafficManagerNestedEndpointAttributes) CustomHeader() terra.ListValue[trafficmanagernestedendpoint.CustomHeaderAttributes] {
	return terra.ReferenceAsList[trafficmanagernestedendpoint.CustomHeaderAttributes](tmne.ref.Append("custom_header"))
}

func (tmne trafficManagerNestedEndpointAttributes) Subnet() terra.ListValue[trafficmanagernestedendpoint.SubnetAttributes] {
	return terra.ReferenceAsList[trafficmanagernestedendpoint.SubnetAttributes](tmne.ref.Append("subnet"))
}

func (tmne trafficManagerNestedEndpointAttributes) Timeouts() trafficmanagernestedendpoint.TimeoutsAttributes {
	return terra.ReferenceAsSingle[trafficmanagernestedendpoint.TimeoutsAttributes](tmne.ref.Append("timeouts"))
}

type trafficManagerNestedEndpointState struct {
	Enabled                           bool                                             `json:"enabled"`
	EndpointLocation                  string                                           `json:"endpoint_location"`
	GeoMappings                       []string                                         `json:"geo_mappings"`
	Id                                string                                           `json:"id"`
	MinimumChildEndpoints             float64                                          `json:"minimum_child_endpoints"`
	MinimumRequiredChildEndpointsIpv4 float64                                          `json:"minimum_required_child_endpoints_ipv4"`
	MinimumRequiredChildEndpointsIpv6 float64                                          `json:"minimum_required_child_endpoints_ipv6"`
	Name                              string                                           `json:"name"`
	Priority                          float64                                          `json:"priority"`
	ProfileId                         string                                           `json:"profile_id"`
	TargetResourceId                  string                                           `json:"target_resource_id"`
	Weight                            float64                                          `json:"weight"`
	CustomHeader                      []trafficmanagernestedendpoint.CustomHeaderState `json:"custom_header"`
	Subnet                            []trafficmanagernestedendpoint.SubnetState       `json:"subnet"`
	Timeouts                          *trafficmanagernestedendpoint.TimeoutsState      `json:"timeouts"`
}
