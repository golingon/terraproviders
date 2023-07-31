// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	iothubroute "github.com/golingon/terraproviders/azurerm/3.67.0/iothubroute"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIothubRoute creates a new instance of [IothubRoute].
func NewIothubRoute(name string, args IothubRouteArgs) *IothubRoute {
	return &IothubRoute{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IothubRoute)(nil)

// IothubRoute represents the Terraform resource azurerm_iothub_route.
type IothubRoute struct {
	Name      string
	Args      IothubRouteArgs
	state     *iothubRouteState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IothubRoute].
func (ir *IothubRoute) Type() string {
	return "azurerm_iothub_route"
}

// LocalName returns the local name for [IothubRoute].
func (ir *IothubRoute) LocalName() string {
	return ir.Name
}

// Configuration returns the configuration (args) for [IothubRoute].
func (ir *IothubRoute) Configuration() interface{} {
	return ir.Args
}

// DependOn is used for other resources to depend on [IothubRoute].
func (ir *IothubRoute) DependOn() terra.Reference {
	return terra.ReferenceResource(ir)
}

// Dependencies returns the list of resources [IothubRoute] depends_on.
func (ir *IothubRoute) Dependencies() terra.Dependencies {
	return ir.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IothubRoute].
func (ir *IothubRoute) LifecycleManagement() *terra.Lifecycle {
	return ir.Lifecycle
}

// Attributes returns the attributes for [IothubRoute].
func (ir *IothubRoute) Attributes() iothubRouteAttributes {
	return iothubRouteAttributes{ref: terra.ReferenceResource(ir)}
}

// ImportState imports the given attribute values into [IothubRoute]'s state.
func (ir *IothubRoute) ImportState(av io.Reader) error {
	ir.state = &iothubRouteState{}
	if err := json.NewDecoder(av).Decode(ir.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ir.Type(), ir.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IothubRoute] has state.
func (ir *IothubRoute) State() (*iothubRouteState, bool) {
	return ir.state, ir.state != nil
}

// StateMust returns the state for [IothubRoute]. Panics if the state is nil.
func (ir *IothubRoute) StateMust() *iothubRouteState {
	if ir.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ir.Type(), ir.LocalName()))
	}
	return ir.state
}

// IothubRouteArgs contains the configurations for azurerm_iothub_route.
type IothubRouteArgs struct {
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
	Timeouts *iothubroute.Timeouts `hcl:"timeouts,block"`
}
type iothubRouteAttributes struct {
	ref terra.Reference
}

// Condition returns a reference to field condition of azurerm_iothub_route.
func (ir iothubRouteAttributes) Condition() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("condition"))
}

// Enabled returns a reference to field enabled of azurerm_iothub_route.
func (ir iothubRouteAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(ir.ref.Append("enabled"))
}

// EndpointNames returns a reference to field endpoint_names of azurerm_iothub_route.
func (ir iothubRouteAttributes) EndpointNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ir.ref.Append("endpoint_names"))
}

// Id returns a reference to field id of azurerm_iothub_route.
func (ir iothubRouteAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("id"))
}

// IothubName returns a reference to field iothub_name of azurerm_iothub_route.
func (ir iothubRouteAttributes) IothubName() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("iothub_name"))
}

// Name returns a reference to field name of azurerm_iothub_route.
func (ir iothubRouteAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_iothub_route.
func (ir iothubRouteAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("resource_group_name"))
}

// Source returns a reference to field source of azurerm_iothub_route.
func (ir iothubRouteAttributes) Source() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("source"))
}

func (ir iothubRouteAttributes) Timeouts() iothubroute.TimeoutsAttributes {
	return terra.ReferenceAsSingle[iothubroute.TimeoutsAttributes](ir.ref.Append("timeouts"))
}

type iothubRouteState struct {
	Condition         string                     `json:"condition"`
	Enabled           bool                       `json:"enabled"`
	EndpointNames     []string                   `json:"endpoint_names"`
	Id                string                     `json:"id"`
	IothubName        string                     `json:"iothub_name"`
	Name              string                     `json:"name"`
	ResourceGroupName string                     `json:"resource_group_name"`
	Source            string                     `json:"source"`
	Timeouts          *iothubroute.TimeoutsState `json:"timeouts"`
}
