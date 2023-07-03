// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	routemap "github.com/golingon/terraproviders/azurerm/3.63.0/routemap"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRouteMap creates a new instance of [RouteMap].
func NewRouteMap(name string, args RouteMapArgs) *RouteMap {
	return &RouteMap{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RouteMap)(nil)

// RouteMap represents the Terraform resource azurerm_route_map.
type RouteMap struct {
	Name      string
	Args      RouteMapArgs
	state     *routeMapState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RouteMap].
func (rm *RouteMap) Type() string {
	return "azurerm_route_map"
}

// LocalName returns the local name for [RouteMap].
func (rm *RouteMap) LocalName() string {
	return rm.Name
}

// Configuration returns the configuration (args) for [RouteMap].
func (rm *RouteMap) Configuration() interface{} {
	return rm.Args
}

// DependOn is used for other resources to depend on [RouteMap].
func (rm *RouteMap) DependOn() terra.Reference {
	return terra.ReferenceResource(rm)
}

// Dependencies returns the list of resources [RouteMap] depends_on.
func (rm *RouteMap) Dependencies() terra.Dependencies {
	return rm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RouteMap].
func (rm *RouteMap) LifecycleManagement() *terra.Lifecycle {
	return rm.Lifecycle
}

// Attributes returns the attributes for [RouteMap].
func (rm *RouteMap) Attributes() routeMapAttributes {
	return routeMapAttributes{ref: terra.ReferenceResource(rm)}
}

// ImportState imports the given attribute values into [RouteMap]'s state.
func (rm *RouteMap) ImportState(av io.Reader) error {
	rm.state = &routeMapState{}
	if err := json.NewDecoder(av).Decode(rm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rm.Type(), rm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RouteMap] has state.
func (rm *RouteMap) State() (*routeMapState, bool) {
	return rm.state, rm.state != nil
}

// StateMust returns the state for [RouteMap]. Panics if the state is nil.
func (rm *RouteMap) StateMust() *routeMapState {
	if rm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rm.Type(), rm.LocalName()))
	}
	return rm.state
}

// RouteMapArgs contains the configurations for azurerm_route_map.
type RouteMapArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// VirtualHubId: string, required
	VirtualHubId terra.StringValue `hcl:"virtual_hub_id,attr" validate:"required"`
	// Rule: min=0
	Rule []routemap.Rule `hcl:"rule,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *routemap.Timeouts `hcl:"timeouts,block"`
}
type routeMapAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_route_map.
func (rm routeMapAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rm.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_route_map.
func (rm routeMapAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rm.ref.Append("name"))
}

// VirtualHubId returns a reference to field virtual_hub_id of azurerm_route_map.
func (rm routeMapAttributes) VirtualHubId() terra.StringValue {
	return terra.ReferenceAsString(rm.ref.Append("virtual_hub_id"))
}

func (rm routeMapAttributes) Rule() terra.ListValue[routemap.RuleAttributes] {
	return terra.ReferenceAsList[routemap.RuleAttributes](rm.ref.Append("rule"))
}

func (rm routeMapAttributes) Timeouts() routemap.TimeoutsAttributes {
	return terra.ReferenceAsSingle[routemap.TimeoutsAttributes](rm.ref.Append("timeouts"))
}

type routeMapState struct {
	Id           string                  `json:"id"`
	Name         string                  `json:"name"`
	VirtualHubId string                  `json:"virtual_hub_id"`
	Rule         []routemap.RuleState    `json:"rule"`
	Timeouts     *routemap.TimeoutsState `json:"timeouts"`
}
