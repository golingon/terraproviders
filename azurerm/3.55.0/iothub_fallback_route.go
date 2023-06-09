// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	iothubfallbackroute "github.com/golingon/terraproviders/azurerm/3.55.0/iothubfallbackroute"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIothubFallbackRoute creates a new instance of [IothubFallbackRoute].
func NewIothubFallbackRoute(name string, args IothubFallbackRouteArgs) *IothubFallbackRoute {
	return &IothubFallbackRoute{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IothubFallbackRoute)(nil)

// IothubFallbackRoute represents the Terraform resource azurerm_iothub_fallback_route.
type IothubFallbackRoute struct {
	Name      string
	Args      IothubFallbackRouteArgs
	state     *iothubFallbackRouteState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IothubFallbackRoute].
func (ifr *IothubFallbackRoute) Type() string {
	return "azurerm_iothub_fallback_route"
}

// LocalName returns the local name for [IothubFallbackRoute].
func (ifr *IothubFallbackRoute) LocalName() string {
	return ifr.Name
}

// Configuration returns the configuration (args) for [IothubFallbackRoute].
func (ifr *IothubFallbackRoute) Configuration() interface{} {
	return ifr.Args
}

// DependOn is used for other resources to depend on [IothubFallbackRoute].
func (ifr *IothubFallbackRoute) DependOn() terra.Reference {
	return terra.ReferenceResource(ifr)
}

// Dependencies returns the list of resources [IothubFallbackRoute] depends_on.
func (ifr *IothubFallbackRoute) Dependencies() terra.Dependencies {
	return ifr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IothubFallbackRoute].
func (ifr *IothubFallbackRoute) LifecycleManagement() *terra.Lifecycle {
	return ifr.Lifecycle
}

// Attributes returns the attributes for [IothubFallbackRoute].
func (ifr *IothubFallbackRoute) Attributes() iothubFallbackRouteAttributes {
	return iothubFallbackRouteAttributes{ref: terra.ReferenceResource(ifr)}
}

// ImportState imports the given attribute values into [IothubFallbackRoute]'s state.
func (ifr *IothubFallbackRoute) ImportState(av io.Reader) error {
	ifr.state = &iothubFallbackRouteState{}
	if err := json.NewDecoder(av).Decode(ifr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ifr.Type(), ifr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IothubFallbackRoute] has state.
func (ifr *IothubFallbackRoute) State() (*iothubFallbackRouteState, bool) {
	return ifr.state, ifr.state != nil
}

// StateMust returns the state for [IothubFallbackRoute]. Panics if the state is nil.
func (ifr *IothubFallbackRoute) StateMust() *iothubFallbackRouteState {
	if ifr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ifr.Type(), ifr.LocalName()))
	}
	return ifr.state
}

// IothubFallbackRouteArgs contains the configurations for azurerm_iothub_fallback_route.
type IothubFallbackRouteArgs struct {
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
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Source: string, optional
	Source terra.StringValue `hcl:"source,attr"`
	// Timeouts: optional
	Timeouts *iothubfallbackroute.Timeouts `hcl:"timeouts,block"`
}
type iothubFallbackRouteAttributes struct {
	ref terra.Reference
}

// Condition returns a reference to field condition of azurerm_iothub_fallback_route.
func (ifr iothubFallbackRouteAttributes) Condition() terra.StringValue {
	return terra.ReferenceAsString(ifr.ref.Append("condition"))
}

// Enabled returns a reference to field enabled of azurerm_iothub_fallback_route.
func (ifr iothubFallbackRouteAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(ifr.ref.Append("enabled"))
}

// EndpointNames returns a reference to field endpoint_names of azurerm_iothub_fallback_route.
func (ifr iothubFallbackRouteAttributes) EndpointNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ifr.ref.Append("endpoint_names"))
}

// Id returns a reference to field id of azurerm_iothub_fallback_route.
func (ifr iothubFallbackRouteAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ifr.ref.Append("id"))
}

// IothubName returns a reference to field iothub_name of azurerm_iothub_fallback_route.
func (ifr iothubFallbackRouteAttributes) IothubName() terra.StringValue {
	return terra.ReferenceAsString(ifr.ref.Append("iothub_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_iothub_fallback_route.
func (ifr iothubFallbackRouteAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ifr.ref.Append("resource_group_name"))
}

// Source returns a reference to field source of azurerm_iothub_fallback_route.
func (ifr iothubFallbackRouteAttributes) Source() terra.StringValue {
	return terra.ReferenceAsString(ifr.ref.Append("source"))
}

func (ifr iothubFallbackRouteAttributes) Timeouts() iothubfallbackroute.TimeoutsAttributes {
	return terra.ReferenceAsSingle[iothubfallbackroute.TimeoutsAttributes](ifr.ref.Append("timeouts"))
}

type iothubFallbackRouteState struct {
	Condition         string                             `json:"condition"`
	Enabled           bool                               `json:"enabled"`
	EndpointNames     []string                           `json:"endpoint_names"`
	Id                string                             `json:"id"`
	IothubName        string                             `json:"iothub_name"`
	ResourceGroupName string                             `json:"resource_group_name"`
	Source            string                             `json:"source"`
	Timeouts          *iothubfallbackroute.TimeoutsState `json:"timeouts"`
}
