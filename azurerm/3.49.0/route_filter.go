// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	routefilter "github.com/golingon/terraproviders/azurerm/3.49.0/routefilter"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewRouteFilter(name string, args RouteFilterArgs) *RouteFilter {
	return &RouteFilter{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RouteFilter)(nil)

type RouteFilter struct {
	Name  string
	Args  RouteFilterArgs
	state *routeFilterState
}

func (rf *RouteFilter) Type() string {
	return "azurerm_route_filter"
}

func (rf *RouteFilter) LocalName() string {
	return rf.Name
}

func (rf *RouteFilter) Configuration() interface{} {
	return rf.Args
}

func (rf *RouteFilter) Attributes() routeFilterAttributes {
	return routeFilterAttributes{ref: terra.ReferenceResource(rf)}
}

func (rf *RouteFilter) ImportState(av io.Reader) error {
	rf.state = &routeFilterState{}
	if err := json.NewDecoder(av).Decode(rf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rf.Type(), rf.LocalName(), err)
	}
	return nil
}

func (rf *RouteFilter) State() (*routeFilterState, bool) {
	return rf.state, rf.state != nil
}

func (rf *RouteFilter) StateMust() *routeFilterState {
	if rf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rf.Type(), rf.LocalName()))
	}
	return rf.state
}

func (rf *RouteFilter) DependOn() terra.Reference {
	return terra.ReferenceResource(rf)
}

type RouteFilterArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Rule: min=0
	Rule []routefilter.Rule `hcl:"rule,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *routefilter.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that RouteFilter depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type routeFilterAttributes struct {
	ref terra.Reference
}

func (rf routeFilterAttributes) Id() terra.StringValue {
	return terra.ReferenceString(rf.ref.Append("id"))
}

func (rf routeFilterAttributes) Location() terra.StringValue {
	return terra.ReferenceString(rf.ref.Append("location"))
}

func (rf routeFilterAttributes) Name() terra.StringValue {
	return terra.ReferenceString(rf.ref.Append("name"))
}

func (rf routeFilterAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(rf.ref.Append("resource_group_name"))
}

func (rf routeFilterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](rf.ref.Append("tags"))
}

func (rf routeFilterAttributes) Rule() terra.ListValue[routefilter.RuleAttributes] {
	return terra.ReferenceList[routefilter.RuleAttributes](rf.ref.Append("rule"))
}

func (rf routeFilterAttributes) Timeouts() routefilter.TimeoutsAttributes {
	return terra.ReferenceSingle[routefilter.TimeoutsAttributes](rf.ref.Append("timeouts"))
}

type routeFilterState struct {
	Id                string                     `json:"id"`
	Location          string                     `json:"location"`
	Name              string                     `json:"name"`
	ResourceGroupName string                     `json:"resource_group_name"`
	Tags              map[string]string          `json:"tags"`
	Rule              []routefilter.RuleState    `json:"rule"`
	Timeouts          *routefilter.TimeoutsState `json:"timeouts"`
}
