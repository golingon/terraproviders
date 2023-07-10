// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	trafficmanagerexternalendpoint "github.com/golingon/terraproviders/azurerm/3.64.0/trafficmanagerexternalendpoint"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewTrafficManagerExternalEndpoint creates a new instance of [TrafficManagerExternalEndpoint].
func NewTrafficManagerExternalEndpoint(name string, args TrafficManagerExternalEndpointArgs) *TrafficManagerExternalEndpoint {
	return &TrafficManagerExternalEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*TrafficManagerExternalEndpoint)(nil)

// TrafficManagerExternalEndpoint represents the Terraform resource azurerm_traffic_manager_external_endpoint.
type TrafficManagerExternalEndpoint struct {
	Name      string
	Args      TrafficManagerExternalEndpointArgs
	state     *trafficManagerExternalEndpointState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [TrafficManagerExternalEndpoint].
func (tmee *TrafficManagerExternalEndpoint) Type() string {
	return "azurerm_traffic_manager_external_endpoint"
}

// LocalName returns the local name for [TrafficManagerExternalEndpoint].
func (tmee *TrafficManagerExternalEndpoint) LocalName() string {
	return tmee.Name
}

// Configuration returns the configuration (args) for [TrafficManagerExternalEndpoint].
func (tmee *TrafficManagerExternalEndpoint) Configuration() interface{} {
	return tmee.Args
}

// DependOn is used for other resources to depend on [TrafficManagerExternalEndpoint].
func (tmee *TrafficManagerExternalEndpoint) DependOn() terra.Reference {
	return terra.ReferenceResource(tmee)
}

// Dependencies returns the list of resources [TrafficManagerExternalEndpoint] depends_on.
func (tmee *TrafficManagerExternalEndpoint) Dependencies() terra.Dependencies {
	return tmee.DependsOn
}

// LifecycleManagement returns the lifecycle block for [TrafficManagerExternalEndpoint].
func (tmee *TrafficManagerExternalEndpoint) LifecycleManagement() *terra.Lifecycle {
	return tmee.Lifecycle
}

// Attributes returns the attributes for [TrafficManagerExternalEndpoint].
func (tmee *TrafficManagerExternalEndpoint) Attributes() trafficManagerExternalEndpointAttributes {
	return trafficManagerExternalEndpointAttributes{ref: terra.ReferenceResource(tmee)}
}

// ImportState imports the given attribute values into [TrafficManagerExternalEndpoint]'s state.
func (tmee *TrafficManagerExternalEndpoint) ImportState(av io.Reader) error {
	tmee.state = &trafficManagerExternalEndpointState{}
	if err := json.NewDecoder(av).Decode(tmee.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", tmee.Type(), tmee.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [TrafficManagerExternalEndpoint] has state.
func (tmee *TrafficManagerExternalEndpoint) State() (*trafficManagerExternalEndpointState, bool) {
	return tmee.state, tmee.state != nil
}

// StateMust returns the state for [TrafficManagerExternalEndpoint]. Panics if the state is nil.
func (tmee *TrafficManagerExternalEndpoint) StateMust() *trafficManagerExternalEndpointState {
	if tmee.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", tmee.Type(), tmee.LocalName()))
	}
	return tmee.state
}

// TrafficManagerExternalEndpointArgs contains the configurations for azurerm_traffic_manager_external_endpoint.
type TrafficManagerExternalEndpointArgs struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// EndpointLocation: string, optional
	EndpointLocation terra.StringValue `hcl:"endpoint_location,attr"`
	// GeoMappings: list of string, optional
	GeoMappings terra.ListValue[terra.StringValue] `hcl:"geo_mappings,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Priority: number, optional
	Priority terra.NumberValue `hcl:"priority,attr"`
	// ProfileId: string, required
	ProfileId terra.StringValue `hcl:"profile_id,attr" validate:"required"`
	// Target: string, required
	Target terra.StringValue `hcl:"target,attr" validate:"required"`
	// Weight: number, optional
	Weight terra.NumberValue `hcl:"weight,attr"`
	// CustomHeader: min=0
	CustomHeader []trafficmanagerexternalendpoint.CustomHeader `hcl:"custom_header,block" validate:"min=0"`
	// Subnet: min=0
	Subnet []trafficmanagerexternalendpoint.Subnet `hcl:"subnet,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *trafficmanagerexternalendpoint.Timeouts `hcl:"timeouts,block"`
}
type trafficManagerExternalEndpointAttributes struct {
	ref terra.Reference
}

// Enabled returns a reference to field enabled of azurerm_traffic_manager_external_endpoint.
func (tmee trafficManagerExternalEndpointAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(tmee.ref.Append("enabled"))
}

// EndpointLocation returns a reference to field endpoint_location of azurerm_traffic_manager_external_endpoint.
func (tmee trafficManagerExternalEndpointAttributes) EndpointLocation() terra.StringValue {
	return terra.ReferenceAsString(tmee.ref.Append("endpoint_location"))
}

// GeoMappings returns a reference to field geo_mappings of azurerm_traffic_manager_external_endpoint.
func (tmee trafficManagerExternalEndpointAttributes) GeoMappings() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](tmee.ref.Append("geo_mappings"))
}

// Id returns a reference to field id of azurerm_traffic_manager_external_endpoint.
func (tmee trafficManagerExternalEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(tmee.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_traffic_manager_external_endpoint.
func (tmee trafficManagerExternalEndpointAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(tmee.ref.Append("name"))
}

// Priority returns a reference to field priority of azurerm_traffic_manager_external_endpoint.
func (tmee trafficManagerExternalEndpointAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(tmee.ref.Append("priority"))
}

// ProfileId returns a reference to field profile_id of azurerm_traffic_manager_external_endpoint.
func (tmee trafficManagerExternalEndpointAttributes) ProfileId() terra.StringValue {
	return terra.ReferenceAsString(tmee.ref.Append("profile_id"))
}

// Target returns a reference to field target of azurerm_traffic_manager_external_endpoint.
func (tmee trafficManagerExternalEndpointAttributes) Target() terra.StringValue {
	return terra.ReferenceAsString(tmee.ref.Append("target"))
}

// Weight returns a reference to field weight of azurerm_traffic_manager_external_endpoint.
func (tmee trafficManagerExternalEndpointAttributes) Weight() terra.NumberValue {
	return terra.ReferenceAsNumber(tmee.ref.Append("weight"))
}

func (tmee trafficManagerExternalEndpointAttributes) CustomHeader() terra.ListValue[trafficmanagerexternalendpoint.CustomHeaderAttributes] {
	return terra.ReferenceAsList[trafficmanagerexternalendpoint.CustomHeaderAttributes](tmee.ref.Append("custom_header"))
}

func (tmee trafficManagerExternalEndpointAttributes) Subnet() terra.ListValue[trafficmanagerexternalendpoint.SubnetAttributes] {
	return terra.ReferenceAsList[trafficmanagerexternalendpoint.SubnetAttributes](tmee.ref.Append("subnet"))
}

func (tmee trafficManagerExternalEndpointAttributes) Timeouts() trafficmanagerexternalendpoint.TimeoutsAttributes {
	return terra.ReferenceAsSingle[trafficmanagerexternalendpoint.TimeoutsAttributes](tmee.ref.Append("timeouts"))
}

type trafficManagerExternalEndpointState struct {
	Enabled          bool                                               `json:"enabled"`
	EndpointLocation string                                             `json:"endpoint_location"`
	GeoMappings      []string                                           `json:"geo_mappings"`
	Id               string                                             `json:"id"`
	Name             string                                             `json:"name"`
	Priority         float64                                            `json:"priority"`
	ProfileId        string                                             `json:"profile_id"`
	Target           string                                             `json:"target"`
	Weight           float64                                            `json:"weight"`
	CustomHeader     []trafficmanagerexternalendpoint.CustomHeaderState `json:"custom_header"`
	Subnet           []trafficmanagerexternalendpoint.SubnetState       `json:"subnet"`
	Timeouts         *trafficmanagerexternalendpoint.TimeoutsState      `json:"timeouts"`
}
