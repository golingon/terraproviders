// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	virtualdesktopworkspace "github.com/golingon/terraproviders/azurerm/3.65.0/virtualdesktopworkspace"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVirtualDesktopWorkspace creates a new instance of [VirtualDesktopWorkspace].
func NewVirtualDesktopWorkspace(name string, args VirtualDesktopWorkspaceArgs) *VirtualDesktopWorkspace {
	return &VirtualDesktopWorkspace{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VirtualDesktopWorkspace)(nil)

// VirtualDesktopWorkspace represents the Terraform resource azurerm_virtual_desktop_workspace.
type VirtualDesktopWorkspace struct {
	Name      string
	Args      VirtualDesktopWorkspaceArgs
	state     *virtualDesktopWorkspaceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VirtualDesktopWorkspace].
func (vdw *VirtualDesktopWorkspace) Type() string {
	return "azurerm_virtual_desktop_workspace"
}

// LocalName returns the local name for [VirtualDesktopWorkspace].
func (vdw *VirtualDesktopWorkspace) LocalName() string {
	return vdw.Name
}

// Configuration returns the configuration (args) for [VirtualDesktopWorkspace].
func (vdw *VirtualDesktopWorkspace) Configuration() interface{} {
	return vdw.Args
}

// DependOn is used for other resources to depend on [VirtualDesktopWorkspace].
func (vdw *VirtualDesktopWorkspace) DependOn() terra.Reference {
	return terra.ReferenceResource(vdw)
}

// Dependencies returns the list of resources [VirtualDesktopWorkspace] depends_on.
func (vdw *VirtualDesktopWorkspace) Dependencies() terra.Dependencies {
	return vdw.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VirtualDesktopWorkspace].
func (vdw *VirtualDesktopWorkspace) LifecycleManagement() *terra.Lifecycle {
	return vdw.Lifecycle
}

// Attributes returns the attributes for [VirtualDesktopWorkspace].
func (vdw *VirtualDesktopWorkspace) Attributes() virtualDesktopWorkspaceAttributes {
	return virtualDesktopWorkspaceAttributes{ref: terra.ReferenceResource(vdw)}
}

// ImportState imports the given attribute values into [VirtualDesktopWorkspace]'s state.
func (vdw *VirtualDesktopWorkspace) ImportState(av io.Reader) error {
	vdw.state = &virtualDesktopWorkspaceState{}
	if err := json.NewDecoder(av).Decode(vdw.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vdw.Type(), vdw.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VirtualDesktopWorkspace] has state.
func (vdw *VirtualDesktopWorkspace) State() (*virtualDesktopWorkspaceState, bool) {
	return vdw.state, vdw.state != nil
}

// StateMust returns the state for [VirtualDesktopWorkspace]. Panics if the state is nil.
func (vdw *VirtualDesktopWorkspace) StateMust() *virtualDesktopWorkspaceState {
	if vdw.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vdw.Type(), vdw.LocalName()))
	}
	return vdw.state
}

// VirtualDesktopWorkspaceArgs contains the configurations for azurerm_virtual_desktop_workspace.
type VirtualDesktopWorkspaceArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// FriendlyName: string, optional
	FriendlyName terra.StringValue `hcl:"friendly_name,attr"`
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
	Timeouts *virtualdesktopworkspace.Timeouts `hcl:"timeouts,block"`
}
type virtualDesktopWorkspaceAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_virtual_desktop_workspace.
func (vdw virtualDesktopWorkspaceAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(vdw.ref.Append("description"))
}

// FriendlyName returns a reference to field friendly_name of azurerm_virtual_desktop_workspace.
func (vdw virtualDesktopWorkspaceAttributes) FriendlyName() terra.StringValue {
	return terra.ReferenceAsString(vdw.ref.Append("friendly_name"))
}

// Id returns a reference to field id of azurerm_virtual_desktop_workspace.
func (vdw virtualDesktopWorkspaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vdw.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_virtual_desktop_workspace.
func (vdw virtualDesktopWorkspaceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(vdw.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_virtual_desktop_workspace.
func (vdw virtualDesktopWorkspaceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vdw.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_virtual_desktop_workspace.
func (vdw virtualDesktopWorkspaceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(vdw.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_virtual_desktop_workspace.
func (vdw virtualDesktopWorkspaceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vdw.ref.Append("tags"))
}

func (vdw virtualDesktopWorkspaceAttributes) Timeouts() virtualdesktopworkspace.TimeoutsAttributes {
	return terra.ReferenceAsSingle[virtualdesktopworkspace.TimeoutsAttributes](vdw.ref.Append("timeouts"))
}

type virtualDesktopWorkspaceState struct {
	Description       string                                 `json:"description"`
	FriendlyName      string                                 `json:"friendly_name"`
	Id                string                                 `json:"id"`
	Location          string                                 `json:"location"`
	Name              string                                 `json:"name"`
	ResourceGroupName string                                 `json:"resource_group_name"`
	Tags              map[string]string                      `json:"tags"`
	Timeouts          *virtualdesktopworkspace.TimeoutsState `json:"timeouts"`
}
