// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	virtualdesktopapplicationgroup "github.com/golingon/terraproviders/azurerm/3.52.0/virtualdesktopapplicationgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVirtualDesktopApplicationGroup creates a new instance of [VirtualDesktopApplicationGroup].
func NewVirtualDesktopApplicationGroup(name string, args VirtualDesktopApplicationGroupArgs) *VirtualDesktopApplicationGroup {
	return &VirtualDesktopApplicationGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VirtualDesktopApplicationGroup)(nil)

// VirtualDesktopApplicationGroup represents the Terraform resource azurerm_virtual_desktop_application_group.
type VirtualDesktopApplicationGroup struct {
	Name      string
	Args      VirtualDesktopApplicationGroupArgs
	state     *virtualDesktopApplicationGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VirtualDesktopApplicationGroup].
func (vdag *VirtualDesktopApplicationGroup) Type() string {
	return "azurerm_virtual_desktop_application_group"
}

// LocalName returns the local name for [VirtualDesktopApplicationGroup].
func (vdag *VirtualDesktopApplicationGroup) LocalName() string {
	return vdag.Name
}

// Configuration returns the configuration (args) for [VirtualDesktopApplicationGroup].
func (vdag *VirtualDesktopApplicationGroup) Configuration() interface{} {
	return vdag.Args
}

// DependOn is used for other resources to depend on [VirtualDesktopApplicationGroup].
func (vdag *VirtualDesktopApplicationGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(vdag)
}

// Dependencies returns the list of resources [VirtualDesktopApplicationGroup] depends_on.
func (vdag *VirtualDesktopApplicationGroup) Dependencies() terra.Dependencies {
	return vdag.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VirtualDesktopApplicationGroup].
func (vdag *VirtualDesktopApplicationGroup) LifecycleManagement() *terra.Lifecycle {
	return vdag.Lifecycle
}

// Attributes returns the attributes for [VirtualDesktopApplicationGroup].
func (vdag *VirtualDesktopApplicationGroup) Attributes() virtualDesktopApplicationGroupAttributes {
	return virtualDesktopApplicationGroupAttributes{ref: terra.ReferenceResource(vdag)}
}

// ImportState imports the given attribute values into [VirtualDesktopApplicationGroup]'s state.
func (vdag *VirtualDesktopApplicationGroup) ImportState(av io.Reader) error {
	vdag.state = &virtualDesktopApplicationGroupState{}
	if err := json.NewDecoder(av).Decode(vdag.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vdag.Type(), vdag.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VirtualDesktopApplicationGroup] has state.
func (vdag *VirtualDesktopApplicationGroup) State() (*virtualDesktopApplicationGroupState, bool) {
	return vdag.state, vdag.state != nil
}

// StateMust returns the state for [VirtualDesktopApplicationGroup]. Panics if the state is nil.
func (vdag *VirtualDesktopApplicationGroup) StateMust() *virtualDesktopApplicationGroupState {
	if vdag.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vdag.Type(), vdag.LocalName()))
	}
	return vdag.state
}

// VirtualDesktopApplicationGroupArgs contains the configurations for azurerm_virtual_desktop_application_group.
type VirtualDesktopApplicationGroupArgs struct {
	// DefaultDesktopDisplayName: string, optional
	DefaultDesktopDisplayName terra.StringValue `hcl:"default_desktop_display_name,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// FriendlyName: string, optional
	FriendlyName terra.StringValue `hcl:"friendly_name,attr"`
	// HostPoolId: string, required
	HostPoolId terra.StringValue `hcl:"host_pool_id,attr" validate:"required"`
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
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *virtualdesktopapplicationgroup.Timeouts `hcl:"timeouts,block"`
}
type virtualDesktopApplicationGroupAttributes struct {
	ref terra.Reference
}

// DefaultDesktopDisplayName returns a reference to field default_desktop_display_name of azurerm_virtual_desktop_application_group.
func (vdag virtualDesktopApplicationGroupAttributes) DefaultDesktopDisplayName() terra.StringValue {
	return terra.ReferenceAsString(vdag.ref.Append("default_desktop_display_name"))
}

// Description returns a reference to field description of azurerm_virtual_desktop_application_group.
func (vdag virtualDesktopApplicationGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(vdag.ref.Append("description"))
}

// FriendlyName returns a reference to field friendly_name of azurerm_virtual_desktop_application_group.
func (vdag virtualDesktopApplicationGroupAttributes) FriendlyName() terra.StringValue {
	return terra.ReferenceAsString(vdag.ref.Append("friendly_name"))
}

// HostPoolId returns a reference to field host_pool_id of azurerm_virtual_desktop_application_group.
func (vdag virtualDesktopApplicationGroupAttributes) HostPoolId() terra.StringValue {
	return terra.ReferenceAsString(vdag.ref.Append("host_pool_id"))
}

// Id returns a reference to field id of azurerm_virtual_desktop_application_group.
func (vdag virtualDesktopApplicationGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vdag.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_virtual_desktop_application_group.
func (vdag virtualDesktopApplicationGroupAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(vdag.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_virtual_desktop_application_group.
func (vdag virtualDesktopApplicationGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vdag.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_virtual_desktop_application_group.
func (vdag virtualDesktopApplicationGroupAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(vdag.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_virtual_desktop_application_group.
func (vdag virtualDesktopApplicationGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vdag.ref.Append("tags"))
}

// Type returns a reference to field type of azurerm_virtual_desktop_application_group.
func (vdag virtualDesktopApplicationGroupAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(vdag.ref.Append("type"))
}

func (vdag virtualDesktopApplicationGroupAttributes) Timeouts() virtualdesktopapplicationgroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[virtualdesktopapplicationgroup.TimeoutsAttributes](vdag.ref.Append("timeouts"))
}

type virtualDesktopApplicationGroupState struct {
	DefaultDesktopDisplayName string                                        `json:"default_desktop_display_name"`
	Description               string                                        `json:"description"`
	FriendlyName              string                                        `json:"friendly_name"`
	HostPoolId                string                                        `json:"host_pool_id"`
	Id                        string                                        `json:"id"`
	Location                  string                                        `json:"location"`
	Name                      string                                        `json:"name"`
	ResourceGroupName         string                                        `json:"resource_group_name"`
	Tags                      map[string]string                             `json:"tags"`
	Type                      string                                        `json:"type"`
	Timeouts                  *virtualdesktopapplicationgroup.TimeoutsState `json:"timeouts"`
}
