// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	monitorprivatelinkscope "github.com/golingon/terraproviders/azurerm/3.52.0/monitorprivatelinkscope"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMonitorPrivateLinkScope creates a new instance of [MonitorPrivateLinkScope].
func NewMonitorPrivateLinkScope(name string, args MonitorPrivateLinkScopeArgs) *MonitorPrivateLinkScope {
	return &MonitorPrivateLinkScope{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MonitorPrivateLinkScope)(nil)

// MonitorPrivateLinkScope represents the Terraform resource azurerm_monitor_private_link_scope.
type MonitorPrivateLinkScope struct {
	Name      string
	Args      MonitorPrivateLinkScopeArgs
	state     *monitorPrivateLinkScopeState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MonitorPrivateLinkScope].
func (mpls *MonitorPrivateLinkScope) Type() string {
	return "azurerm_monitor_private_link_scope"
}

// LocalName returns the local name for [MonitorPrivateLinkScope].
func (mpls *MonitorPrivateLinkScope) LocalName() string {
	return mpls.Name
}

// Configuration returns the configuration (args) for [MonitorPrivateLinkScope].
func (mpls *MonitorPrivateLinkScope) Configuration() interface{} {
	return mpls.Args
}

// DependOn is used for other resources to depend on [MonitorPrivateLinkScope].
func (mpls *MonitorPrivateLinkScope) DependOn() terra.Reference {
	return terra.ReferenceResource(mpls)
}

// Dependencies returns the list of resources [MonitorPrivateLinkScope] depends_on.
func (mpls *MonitorPrivateLinkScope) Dependencies() terra.Dependencies {
	return mpls.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MonitorPrivateLinkScope].
func (mpls *MonitorPrivateLinkScope) LifecycleManagement() *terra.Lifecycle {
	return mpls.Lifecycle
}

// Attributes returns the attributes for [MonitorPrivateLinkScope].
func (mpls *MonitorPrivateLinkScope) Attributes() monitorPrivateLinkScopeAttributes {
	return monitorPrivateLinkScopeAttributes{ref: terra.ReferenceResource(mpls)}
}

// ImportState imports the given attribute values into [MonitorPrivateLinkScope]'s state.
func (mpls *MonitorPrivateLinkScope) ImportState(av io.Reader) error {
	mpls.state = &monitorPrivateLinkScopeState{}
	if err := json.NewDecoder(av).Decode(mpls.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mpls.Type(), mpls.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MonitorPrivateLinkScope] has state.
func (mpls *MonitorPrivateLinkScope) State() (*monitorPrivateLinkScopeState, bool) {
	return mpls.state, mpls.state != nil
}

// StateMust returns the state for [MonitorPrivateLinkScope]. Panics if the state is nil.
func (mpls *MonitorPrivateLinkScope) StateMust() *monitorPrivateLinkScopeState {
	if mpls.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mpls.Type(), mpls.LocalName()))
	}
	return mpls.state
}

// MonitorPrivateLinkScopeArgs contains the configurations for azurerm_monitor_private_link_scope.
type MonitorPrivateLinkScopeArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *monitorprivatelinkscope.Timeouts `hcl:"timeouts,block"`
}
type monitorPrivateLinkScopeAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_monitor_private_link_scope.
func (mpls monitorPrivateLinkScopeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mpls.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_monitor_private_link_scope.
func (mpls monitorPrivateLinkScopeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mpls.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_monitor_private_link_scope.
func (mpls monitorPrivateLinkScopeAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(mpls.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_monitor_private_link_scope.
func (mpls monitorPrivateLinkScopeAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mpls.ref.Append("tags"))
}

func (mpls monitorPrivateLinkScopeAttributes) Timeouts() monitorprivatelinkscope.TimeoutsAttributes {
	return terra.ReferenceAsSingle[monitorprivatelinkscope.TimeoutsAttributes](mpls.ref.Append("timeouts"))
}

type monitorPrivateLinkScopeState struct {
	Id                string                                 `json:"id"`
	Name              string                                 `json:"name"`
	ResourceGroupName string                                 `json:"resource_group_name"`
	Tags              map[string]string                      `json:"tags"`
	Timeouts          *monitorprivatelinkscope.TimeoutsState `json:"timeouts"`
}
