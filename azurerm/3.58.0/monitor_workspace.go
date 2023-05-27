// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	monitorworkspace "github.com/golingon/terraproviders/azurerm/3.58.0/monitorworkspace"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMonitorWorkspace creates a new instance of [MonitorWorkspace].
func NewMonitorWorkspace(name string, args MonitorWorkspaceArgs) *MonitorWorkspace {
	return &MonitorWorkspace{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MonitorWorkspace)(nil)

// MonitorWorkspace represents the Terraform resource azurerm_monitor_workspace.
type MonitorWorkspace struct {
	Name      string
	Args      MonitorWorkspaceArgs
	state     *monitorWorkspaceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MonitorWorkspace].
func (mw *MonitorWorkspace) Type() string {
	return "azurerm_monitor_workspace"
}

// LocalName returns the local name for [MonitorWorkspace].
func (mw *MonitorWorkspace) LocalName() string {
	return mw.Name
}

// Configuration returns the configuration (args) for [MonitorWorkspace].
func (mw *MonitorWorkspace) Configuration() interface{} {
	return mw.Args
}

// DependOn is used for other resources to depend on [MonitorWorkspace].
func (mw *MonitorWorkspace) DependOn() terra.Reference {
	return terra.ReferenceResource(mw)
}

// Dependencies returns the list of resources [MonitorWorkspace] depends_on.
func (mw *MonitorWorkspace) Dependencies() terra.Dependencies {
	return mw.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MonitorWorkspace].
func (mw *MonitorWorkspace) LifecycleManagement() *terra.Lifecycle {
	return mw.Lifecycle
}

// Attributes returns the attributes for [MonitorWorkspace].
func (mw *MonitorWorkspace) Attributes() monitorWorkspaceAttributes {
	return monitorWorkspaceAttributes{ref: terra.ReferenceResource(mw)}
}

// ImportState imports the given attribute values into [MonitorWorkspace]'s state.
func (mw *MonitorWorkspace) ImportState(av io.Reader) error {
	mw.state = &monitorWorkspaceState{}
	if err := json.NewDecoder(av).Decode(mw.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mw.Type(), mw.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MonitorWorkspace] has state.
func (mw *MonitorWorkspace) State() (*monitorWorkspaceState, bool) {
	return mw.state, mw.state != nil
}

// StateMust returns the state for [MonitorWorkspace]. Panics if the state is nil.
func (mw *MonitorWorkspace) StateMust() *monitorWorkspaceState {
	if mw.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mw.Type(), mw.LocalName()))
	}
	return mw.state
}

// MonitorWorkspaceArgs contains the configurations for azurerm_monitor_workspace.
type MonitorWorkspaceArgs struct {
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
	// Timeouts: optional
	Timeouts *monitorworkspace.Timeouts `hcl:"timeouts,block"`
}
type monitorWorkspaceAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_monitor_workspace.
func (mw monitorWorkspaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mw.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_monitor_workspace.
func (mw monitorWorkspaceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(mw.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_monitor_workspace.
func (mw monitorWorkspaceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mw.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_monitor_workspace.
func (mw monitorWorkspaceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(mw.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_monitor_workspace.
func (mw monitorWorkspaceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mw.ref.Append("tags"))
}

func (mw monitorWorkspaceAttributes) Timeouts() monitorworkspace.TimeoutsAttributes {
	return terra.ReferenceAsSingle[monitorworkspace.TimeoutsAttributes](mw.ref.Append("timeouts"))
}

type monitorWorkspaceState struct {
	Id                string                          `json:"id"`
	Location          string                          `json:"location"`
	Name              string                          `json:"name"`
	ResourceGroupName string                          `json:"resource_group_name"`
	Tags              map[string]string               `json:"tags"`
	Timeouts          *monitorworkspace.TimeoutsState `json:"timeouts"`
}
