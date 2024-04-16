// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_iothub_route

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource azurerm_iothub_route.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermIothubRouteState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (air *Resource) Type() string {
	return "azurerm_iothub_route"
}

// LocalName returns the local name for [Resource].
func (air *Resource) LocalName() string {
	return air.Name
}

// Configuration returns the configuration (args) for [Resource].
func (air *Resource) Configuration() interface{} {
	return air.Args
}

// DependOn is used for other resources to depend on [Resource].
func (air *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(air)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (air *Resource) Dependencies() terra.Dependencies {
	return air.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (air *Resource) LifecycleManagement() *terra.Lifecycle {
	return air.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (air *Resource) Attributes() azurermIothubRouteAttributes {
	return azurermIothubRouteAttributes{ref: terra.ReferenceResource(air)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (air *Resource) ImportState(state io.Reader) error {
	air.state = &azurermIothubRouteState{}
	if err := json.NewDecoder(state).Decode(air.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", air.Type(), air.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (air *Resource) State() (*azurermIothubRouteState, bool) {
	return air.state, air.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (air *Resource) StateMust() *azurermIothubRouteState {
	if air.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", air.Type(), air.LocalName()))
	}
	return air.state
}

// Args contains the configurations for azurerm_iothub_route.
type Args struct {
	// Condition: string, optional
	Condition terra.StringValue `hcl:"condition,attr"`
	// Enabled: bool, required
	Enabled terra.BoolValue `hcl:"enabled,attr" validate:"required"`
	// EndpointNames: list of string, required
	EndpointNames terra.ListValue[terra.StringValue] `hcl:"endpoint_names,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IothubName: string, required
	IothubName terra.StringValue `hcl:"iothub_name,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Source: string, required
	Source terra.StringValue `hcl:"source,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermIothubRouteAttributes struct {
	ref terra.Reference
}

// Condition returns a reference to field condition of azurerm_iothub_route.
func (air azurermIothubRouteAttributes) Condition() terra.StringValue {
	return terra.ReferenceAsString(air.ref.Append("condition"))
}

// Enabled returns a reference to field enabled of azurerm_iothub_route.
func (air azurermIothubRouteAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(air.ref.Append("enabled"))
}

// EndpointNames returns a reference to field endpoint_names of azurerm_iothub_route.
func (air azurermIothubRouteAttributes) EndpointNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](air.ref.Append("endpoint_names"))
}

// Id returns a reference to field id of azurerm_iothub_route.
func (air azurermIothubRouteAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(air.ref.Append("id"))
}

// IothubName returns a reference to field iothub_name of azurerm_iothub_route.
func (air azurermIothubRouteAttributes) IothubName() terra.StringValue {
	return terra.ReferenceAsString(air.ref.Append("iothub_name"))
}

// Name returns a reference to field name of azurerm_iothub_route.
func (air azurermIothubRouteAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(air.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_iothub_route.
func (air azurermIothubRouteAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(air.ref.Append("resource_group_name"))
}

// Source returns a reference to field source of azurerm_iothub_route.
func (air azurermIothubRouteAttributes) Source() terra.StringValue {
	return terra.ReferenceAsString(air.ref.Append("source"))
}

func (air azurermIothubRouteAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](air.ref.Append("timeouts"))
}

type azurermIothubRouteState struct {
	Condition         string         `json:"condition"`
	Enabled           bool           `json:"enabled"`
	EndpointNames     []string       `json:"endpoint_names"`
	Id                string         `json:"id"`
	IothubName        string         `json:"iothub_name"`
	Name              string         `json:"name"`
	ResourceGroupName string         `json:"resource_group_name"`
	Source            string         `json:"source"`
	Timeouts          *TimeoutsState `json:"timeouts"`
}
