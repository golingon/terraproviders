// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	trafficmanagerazureendpoint "github.com/golingon/terraproviders/azurerm/3.69.0/trafficmanagerazureendpoint"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewTrafficManagerAzureEndpoint creates a new instance of [TrafficManagerAzureEndpoint].
func NewTrafficManagerAzureEndpoint(name string, args TrafficManagerAzureEndpointArgs) *TrafficManagerAzureEndpoint {
	return &TrafficManagerAzureEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*TrafficManagerAzureEndpoint)(nil)

// TrafficManagerAzureEndpoint represents the Terraform resource azurerm_traffic_manager_azure_endpoint.
type TrafficManagerAzureEndpoint struct {
	Name      string
	Args      TrafficManagerAzureEndpointArgs
	state     *trafficManagerAzureEndpointState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [TrafficManagerAzureEndpoint].
func (tmae *TrafficManagerAzureEndpoint) Type() string {
	return "azurerm_traffic_manager_azure_endpoint"
}

// LocalName returns the local name for [TrafficManagerAzureEndpoint].
func (tmae *TrafficManagerAzureEndpoint) LocalName() string {
	return tmae.Name
}

// Configuration returns the configuration (args) for [TrafficManagerAzureEndpoint].
func (tmae *TrafficManagerAzureEndpoint) Configuration() interface{} {
	return tmae.Args
}

// DependOn is used for other resources to depend on [TrafficManagerAzureEndpoint].
func (tmae *TrafficManagerAzureEndpoint) DependOn() terra.Reference {
	return terra.ReferenceResource(tmae)
}

// Dependencies returns the list of resources [TrafficManagerAzureEndpoint] depends_on.
func (tmae *TrafficManagerAzureEndpoint) Dependencies() terra.Dependencies {
	return tmae.DependsOn
}

// LifecycleManagement returns the lifecycle block for [TrafficManagerAzureEndpoint].
func (tmae *TrafficManagerAzureEndpoint) LifecycleManagement() *terra.Lifecycle {
	return tmae.Lifecycle
}

// Attributes returns the attributes for [TrafficManagerAzureEndpoint].
func (tmae *TrafficManagerAzureEndpoint) Attributes() trafficManagerAzureEndpointAttributes {
	return trafficManagerAzureEndpointAttributes{ref: terra.ReferenceResource(tmae)}
}

// ImportState imports the given attribute values into [TrafficManagerAzureEndpoint]'s state.
func (tmae *TrafficManagerAzureEndpoint) ImportState(av io.Reader) error {
	tmae.state = &trafficManagerAzureEndpointState{}
	if err := json.NewDecoder(av).Decode(tmae.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", tmae.Type(), tmae.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [TrafficManagerAzureEndpoint] has state.
func (tmae *TrafficManagerAzureEndpoint) State() (*trafficManagerAzureEndpointState, bool) {
	return tmae.state, tmae.state != nil
}

// StateMust returns the state for [TrafficManagerAzureEndpoint]. Panics if the state is nil.
func (tmae *TrafficManagerAzureEndpoint) StateMust() *trafficManagerAzureEndpointState {
	if tmae.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", tmae.Type(), tmae.LocalName()))
	}
	return tmae.state
}

// TrafficManagerAzureEndpointArgs contains the configurations for azurerm_traffic_manager_azure_endpoint.
type TrafficManagerAzureEndpointArgs struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
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
	// TargetResourceId: string, required
	TargetResourceId terra.StringValue `hcl:"target_resource_id,attr" validate:"required"`
	// Weight: number, optional
	Weight terra.NumberValue `hcl:"weight,attr"`
	// CustomHeader: min=0
	CustomHeader []trafficmanagerazureendpoint.CustomHeader `hcl:"custom_header,block" validate:"min=0"`
	// Subnet: min=0
	Subnet []trafficmanagerazureendpoint.Subnet `hcl:"subnet,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *trafficmanagerazureendpoint.Timeouts `hcl:"timeouts,block"`
}
type trafficManagerAzureEndpointAttributes struct {
	ref terra.Reference
}

// Enabled returns a reference to field enabled of azurerm_traffic_manager_azure_endpoint.
func (tmae trafficManagerAzureEndpointAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(tmae.ref.Append("enabled"))
}

// GeoMappings returns a reference to field geo_mappings of azurerm_traffic_manager_azure_endpoint.
func (tmae trafficManagerAzureEndpointAttributes) GeoMappings() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](tmae.ref.Append("geo_mappings"))
}

// Id returns a reference to field id of azurerm_traffic_manager_azure_endpoint.
func (tmae trafficManagerAzureEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(tmae.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_traffic_manager_azure_endpoint.
func (tmae trafficManagerAzureEndpointAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(tmae.ref.Append("name"))
}

// Priority returns a reference to field priority of azurerm_traffic_manager_azure_endpoint.
func (tmae trafficManagerAzureEndpointAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(tmae.ref.Append("priority"))
}

// ProfileId returns a reference to field profile_id of azurerm_traffic_manager_azure_endpoint.
func (tmae trafficManagerAzureEndpointAttributes) ProfileId() terra.StringValue {
	return terra.ReferenceAsString(tmae.ref.Append("profile_id"))
}

// TargetResourceId returns a reference to field target_resource_id of azurerm_traffic_manager_azure_endpoint.
func (tmae trafficManagerAzureEndpointAttributes) TargetResourceId() terra.StringValue {
	return terra.ReferenceAsString(tmae.ref.Append("target_resource_id"))
}

// Weight returns a reference to field weight of azurerm_traffic_manager_azure_endpoint.
func (tmae trafficManagerAzureEndpointAttributes) Weight() terra.NumberValue {
	return terra.ReferenceAsNumber(tmae.ref.Append("weight"))
}

func (tmae trafficManagerAzureEndpointAttributes) CustomHeader() terra.ListValue[trafficmanagerazureendpoint.CustomHeaderAttributes] {
	return terra.ReferenceAsList[trafficmanagerazureendpoint.CustomHeaderAttributes](tmae.ref.Append("custom_header"))
}

func (tmae trafficManagerAzureEndpointAttributes) Subnet() terra.ListValue[trafficmanagerazureendpoint.SubnetAttributes] {
	return terra.ReferenceAsList[trafficmanagerazureendpoint.SubnetAttributes](tmae.ref.Append("subnet"))
}

func (tmae trafficManagerAzureEndpointAttributes) Timeouts() trafficmanagerazureendpoint.TimeoutsAttributes {
	return terra.ReferenceAsSingle[trafficmanagerazureendpoint.TimeoutsAttributes](tmae.ref.Append("timeouts"))
}

type trafficManagerAzureEndpointState struct {
	Enabled          bool                                            `json:"enabled"`
	GeoMappings      []string                                        `json:"geo_mappings"`
	Id               string                                          `json:"id"`
	Name             string                                          `json:"name"`
	Priority         float64                                         `json:"priority"`
	ProfileId        string                                          `json:"profile_id"`
	TargetResourceId string                                          `json:"target_resource_id"`
	Weight           float64                                         `json:"weight"`
	CustomHeader     []trafficmanagerazureendpoint.CustomHeaderState `json:"custom_header"`
	Subnet           []trafficmanagerazureendpoint.SubnetState       `json:"subnet"`
	Timeouts         *trafficmanagerazureendpoint.TimeoutsState      `json:"timeouts"`
}
