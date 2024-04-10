// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	virtualdesktophostpool "github.com/golingon/terraproviders/azurerm/3.98.0/virtualdesktophostpool"
	"io"
)

// NewVirtualDesktopHostPool creates a new instance of [VirtualDesktopHostPool].
func NewVirtualDesktopHostPool(name string, args VirtualDesktopHostPoolArgs) *VirtualDesktopHostPool {
	return &VirtualDesktopHostPool{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VirtualDesktopHostPool)(nil)

// VirtualDesktopHostPool represents the Terraform resource azurerm_virtual_desktop_host_pool.
type VirtualDesktopHostPool struct {
	Name      string
	Args      VirtualDesktopHostPoolArgs
	state     *virtualDesktopHostPoolState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VirtualDesktopHostPool].
func (vdhp *VirtualDesktopHostPool) Type() string {
	return "azurerm_virtual_desktop_host_pool"
}

// LocalName returns the local name for [VirtualDesktopHostPool].
func (vdhp *VirtualDesktopHostPool) LocalName() string {
	return vdhp.Name
}

// Configuration returns the configuration (args) for [VirtualDesktopHostPool].
func (vdhp *VirtualDesktopHostPool) Configuration() interface{} {
	return vdhp.Args
}

// DependOn is used for other resources to depend on [VirtualDesktopHostPool].
func (vdhp *VirtualDesktopHostPool) DependOn() terra.Reference {
	return terra.ReferenceResource(vdhp)
}

// Dependencies returns the list of resources [VirtualDesktopHostPool] depends_on.
func (vdhp *VirtualDesktopHostPool) Dependencies() terra.Dependencies {
	return vdhp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VirtualDesktopHostPool].
func (vdhp *VirtualDesktopHostPool) LifecycleManagement() *terra.Lifecycle {
	return vdhp.Lifecycle
}

// Attributes returns the attributes for [VirtualDesktopHostPool].
func (vdhp *VirtualDesktopHostPool) Attributes() virtualDesktopHostPoolAttributes {
	return virtualDesktopHostPoolAttributes{ref: terra.ReferenceResource(vdhp)}
}

// ImportState imports the given attribute values into [VirtualDesktopHostPool]'s state.
func (vdhp *VirtualDesktopHostPool) ImportState(av io.Reader) error {
	vdhp.state = &virtualDesktopHostPoolState{}
	if err := json.NewDecoder(av).Decode(vdhp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vdhp.Type(), vdhp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VirtualDesktopHostPool] has state.
func (vdhp *VirtualDesktopHostPool) State() (*virtualDesktopHostPoolState, bool) {
	return vdhp.state, vdhp.state != nil
}

// StateMust returns the state for [VirtualDesktopHostPool]. Panics if the state is nil.
func (vdhp *VirtualDesktopHostPool) StateMust() *virtualDesktopHostPoolState {
	if vdhp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vdhp.Type(), vdhp.LocalName()))
	}
	return vdhp.state
}

// VirtualDesktopHostPoolArgs contains the configurations for azurerm_virtual_desktop_host_pool.
type VirtualDesktopHostPoolArgs struct {
	// CustomRdpProperties: string, optional
	CustomRdpProperties terra.StringValue `hcl:"custom_rdp_properties,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// FriendlyName: string, optional
	FriendlyName terra.StringValue `hcl:"friendly_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LoadBalancerType: string, required
	LoadBalancerType terra.StringValue `hcl:"load_balancer_type,attr" validate:"required"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// MaximumSessionsAllowed: number, optional
	MaximumSessionsAllowed terra.NumberValue `hcl:"maximum_sessions_allowed,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PersonalDesktopAssignmentType: string, optional
	PersonalDesktopAssignmentType terra.StringValue `hcl:"personal_desktop_assignment_type,attr"`
	// PreferredAppGroupType: string, optional
	PreferredAppGroupType terra.StringValue `hcl:"preferred_app_group_type,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// StartVmOnConnect: bool, optional
	StartVmOnConnect terra.BoolValue `hcl:"start_vm_on_connect,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// ValidateEnvironment: bool, optional
	ValidateEnvironment terra.BoolValue `hcl:"validate_environment,attr"`
	// VmTemplate: string, optional
	VmTemplate terra.StringValue `hcl:"vm_template,attr"`
	// ScheduledAgentUpdates: optional
	ScheduledAgentUpdates *virtualdesktophostpool.ScheduledAgentUpdates `hcl:"scheduled_agent_updates,block"`
	// Timeouts: optional
	Timeouts *virtualdesktophostpool.Timeouts `hcl:"timeouts,block"`
}
type virtualDesktopHostPoolAttributes struct {
	ref terra.Reference
}

// CustomRdpProperties returns a reference to field custom_rdp_properties of azurerm_virtual_desktop_host_pool.
func (vdhp virtualDesktopHostPoolAttributes) CustomRdpProperties() terra.StringValue {
	return terra.ReferenceAsString(vdhp.ref.Append("custom_rdp_properties"))
}

// Description returns a reference to field description of azurerm_virtual_desktop_host_pool.
func (vdhp virtualDesktopHostPoolAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(vdhp.ref.Append("description"))
}

// FriendlyName returns a reference to field friendly_name of azurerm_virtual_desktop_host_pool.
func (vdhp virtualDesktopHostPoolAttributes) FriendlyName() terra.StringValue {
	return terra.ReferenceAsString(vdhp.ref.Append("friendly_name"))
}

// Id returns a reference to field id of azurerm_virtual_desktop_host_pool.
func (vdhp virtualDesktopHostPoolAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vdhp.ref.Append("id"))
}

// LoadBalancerType returns a reference to field load_balancer_type of azurerm_virtual_desktop_host_pool.
func (vdhp virtualDesktopHostPoolAttributes) LoadBalancerType() terra.StringValue {
	return terra.ReferenceAsString(vdhp.ref.Append("load_balancer_type"))
}

// Location returns a reference to field location of azurerm_virtual_desktop_host_pool.
func (vdhp virtualDesktopHostPoolAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(vdhp.ref.Append("location"))
}

// MaximumSessionsAllowed returns a reference to field maximum_sessions_allowed of azurerm_virtual_desktop_host_pool.
func (vdhp virtualDesktopHostPoolAttributes) MaximumSessionsAllowed() terra.NumberValue {
	return terra.ReferenceAsNumber(vdhp.ref.Append("maximum_sessions_allowed"))
}

// Name returns a reference to field name of azurerm_virtual_desktop_host_pool.
func (vdhp virtualDesktopHostPoolAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vdhp.ref.Append("name"))
}

// PersonalDesktopAssignmentType returns a reference to field personal_desktop_assignment_type of azurerm_virtual_desktop_host_pool.
func (vdhp virtualDesktopHostPoolAttributes) PersonalDesktopAssignmentType() terra.StringValue {
	return terra.ReferenceAsString(vdhp.ref.Append("personal_desktop_assignment_type"))
}

// PreferredAppGroupType returns a reference to field preferred_app_group_type of azurerm_virtual_desktop_host_pool.
func (vdhp virtualDesktopHostPoolAttributes) PreferredAppGroupType() terra.StringValue {
	return terra.ReferenceAsString(vdhp.ref.Append("preferred_app_group_type"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_virtual_desktop_host_pool.
func (vdhp virtualDesktopHostPoolAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(vdhp.ref.Append("resource_group_name"))
}

// StartVmOnConnect returns a reference to field start_vm_on_connect of azurerm_virtual_desktop_host_pool.
func (vdhp virtualDesktopHostPoolAttributes) StartVmOnConnect() terra.BoolValue {
	return terra.ReferenceAsBool(vdhp.ref.Append("start_vm_on_connect"))
}

// Tags returns a reference to field tags of azurerm_virtual_desktop_host_pool.
func (vdhp virtualDesktopHostPoolAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vdhp.ref.Append("tags"))
}

// Type returns a reference to field type of azurerm_virtual_desktop_host_pool.
func (vdhp virtualDesktopHostPoolAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(vdhp.ref.Append("type"))
}

// ValidateEnvironment returns a reference to field validate_environment of azurerm_virtual_desktop_host_pool.
func (vdhp virtualDesktopHostPoolAttributes) ValidateEnvironment() terra.BoolValue {
	return terra.ReferenceAsBool(vdhp.ref.Append("validate_environment"))
}

// VmTemplate returns a reference to field vm_template of azurerm_virtual_desktop_host_pool.
func (vdhp virtualDesktopHostPoolAttributes) VmTemplate() terra.StringValue {
	return terra.ReferenceAsString(vdhp.ref.Append("vm_template"))
}

func (vdhp virtualDesktopHostPoolAttributes) ScheduledAgentUpdates() terra.ListValue[virtualdesktophostpool.ScheduledAgentUpdatesAttributes] {
	return terra.ReferenceAsList[virtualdesktophostpool.ScheduledAgentUpdatesAttributes](vdhp.ref.Append("scheduled_agent_updates"))
}

func (vdhp virtualDesktopHostPoolAttributes) Timeouts() virtualdesktophostpool.TimeoutsAttributes {
	return terra.ReferenceAsSingle[virtualdesktophostpool.TimeoutsAttributes](vdhp.ref.Append("timeouts"))
}

type virtualDesktopHostPoolState struct {
	CustomRdpProperties           string                                              `json:"custom_rdp_properties"`
	Description                   string                                              `json:"description"`
	FriendlyName                  string                                              `json:"friendly_name"`
	Id                            string                                              `json:"id"`
	LoadBalancerType              string                                              `json:"load_balancer_type"`
	Location                      string                                              `json:"location"`
	MaximumSessionsAllowed        float64                                             `json:"maximum_sessions_allowed"`
	Name                          string                                              `json:"name"`
	PersonalDesktopAssignmentType string                                              `json:"personal_desktop_assignment_type"`
	PreferredAppGroupType         string                                              `json:"preferred_app_group_type"`
	ResourceGroupName             string                                              `json:"resource_group_name"`
	StartVmOnConnect              bool                                                `json:"start_vm_on_connect"`
	Tags                          map[string]string                                   `json:"tags"`
	Type                          string                                              `json:"type"`
	ValidateEnvironment           bool                                                `json:"validate_environment"`
	VmTemplate                    string                                              `json:"vm_template"`
	ScheduledAgentUpdates         []virtualdesktophostpool.ScheduledAgentUpdatesState `json:"scheduled_agent_updates"`
	Timeouts                      *virtualdesktophostpool.TimeoutsState               `json:"timeouts"`
}
