// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datavirtualdesktophostpool "github.com/golingon/terraproviders/azurerm/3.69.0/datavirtualdesktophostpool"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataVirtualDesktopHostPool creates a new instance of [DataVirtualDesktopHostPool].
func NewDataVirtualDesktopHostPool(name string, args DataVirtualDesktopHostPoolArgs) *DataVirtualDesktopHostPool {
	return &DataVirtualDesktopHostPool{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataVirtualDesktopHostPool)(nil)

// DataVirtualDesktopHostPool represents the Terraform data resource azurerm_virtual_desktop_host_pool.
type DataVirtualDesktopHostPool struct {
	Name string
	Args DataVirtualDesktopHostPoolArgs
}

// DataSource returns the Terraform object type for [DataVirtualDesktopHostPool].
func (vdhp *DataVirtualDesktopHostPool) DataSource() string {
	return "azurerm_virtual_desktop_host_pool"
}

// LocalName returns the local name for [DataVirtualDesktopHostPool].
func (vdhp *DataVirtualDesktopHostPool) LocalName() string {
	return vdhp.Name
}

// Configuration returns the configuration (args) for [DataVirtualDesktopHostPool].
func (vdhp *DataVirtualDesktopHostPool) Configuration() interface{} {
	return vdhp.Args
}

// Attributes returns the attributes for [DataVirtualDesktopHostPool].
func (vdhp *DataVirtualDesktopHostPool) Attributes() dataVirtualDesktopHostPoolAttributes {
	return dataVirtualDesktopHostPoolAttributes{ref: terra.ReferenceDataResource(vdhp)}
}

// DataVirtualDesktopHostPoolArgs contains the configurations for azurerm_virtual_desktop_host_pool.
type DataVirtualDesktopHostPoolArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ScheduledAgentUpdates: min=0
	ScheduledAgentUpdates []datavirtualdesktophostpool.ScheduledAgentUpdates `hcl:"scheduled_agent_updates,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datavirtualdesktophostpool.Timeouts `hcl:"timeouts,block"`
}
type dataVirtualDesktopHostPoolAttributes struct {
	ref terra.Reference
}

// CustomRdpProperties returns a reference to field custom_rdp_properties of azurerm_virtual_desktop_host_pool.
func (vdhp dataVirtualDesktopHostPoolAttributes) CustomRdpProperties() terra.StringValue {
	return terra.ReferenceAsString(vdhp.ref.Append("custom_rdp_properties"))
}

// Description returns a reference to field description of azurerm_virtual_desktop_host_pool.
func (vdhp dataVirtualDesktopHostPoolAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(vdhp.ref.Append("description"))
}

// FriendlyName returns a reference to field friendly_name of azurerm_virtual_desktop_host_pool.
func (vdhp dataVirtualDesktopHostPoolAttributes) FriendlyName() terra.StringValue {
	return terra.ReferenceAsString(vdhp.ref.Append("friendly_name"))
}

// Id returns a reference to field id of azurerm_virtual_desktop_host_pool.
func (vdhp dataVirtualDesktopHostPoolAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vdhp.ref.Append("id"))
}

// LoadBalancerType returns a reference to field load_balancer_type of azurerm_virtual_desktop_host_pool.
func (vdhp dataVirtualDesktopHostPoolAttributes) LoadBalancerType() terra.StringValue {
	return terra.ReferenceAsString(vdhp.ref.Append("load_balancer_type"))
}

// Location returns a reference to field location of azurerm_virtual_desktop_host_pool.
func (vdhp dataVirtualDesktopHostPoolAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(vdhp.ref.Append("location"))
}

// MaximumSessionsAllowed returns a reference to field maximum_sessions_allowed of azurerm_virtual_desktop_host_pool.
func (vdhp dataVirtualDesktopHostPoolAttributes) MaximumSessionsAllowed() terra.NumberValue {
	return terra.ReferenceAsNumber(vdhp.ref.Append("maximum_sessions_allowed"))
}

// Name returns a reference to field name of azurerm_virtual_desktop_host_pool.
func (vdhp dataVirtualDesktopHostPoolAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vdhp.ref.Append("name"))
}

// PersonalDesktopAssignmentType returns a reference to field personal_desktop_assignment_type of azurerm_virtual_desktop_host_pool.
func (vdhp dataVirtualDesktopHostPoolAttributes) PersonalDesktopAssignmentType() terra.StringValue {
	return terra.ReferenceAsString(vdhp.ref.Append("personal_desktop_assignment_type"))
}

// PreferredAppGroupType returns a reference to field preferred_app_group_type of azurerm_virtual_desktop_host_pool.
func (vdhp dataVirtualDesktopHostPoolAttributes) PreferredAppGroupType() terra.StringValue {
	return terra.ReferenceAsString(vdhp.ref.Append("preferred_app_group_type"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_virtual_desktop_host_pool.
func (vdhp dataVirtualDesktopHostPoolAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(vdhp.ref.Append("resource_group_name"))
}

// StartVmOnConnect returns a reference to field start_vm_on_connect of azurerm_virtual_desktop_host_pool.
func (vdhp dataVirtualDesktopHostPoolAttributes) StartVmOnConnect() terra.BoolValue {
	return terra.ReferenceAsBool(vdhp.ref.Append("start_vm_on_connect"))
}

// Tags returns a reference to field tags of azurerm_virtual_desktop_host_pool.
func (vdhp dataVirtualDesktopHostPoolAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vdhp.ref.Append("tags"))
}

// Type returns a reference to field type of azurerm_virtual_desktop_host_pool.
func (vdhp dataVirtualDesktopHostPoolAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(vdhp.ref.Append("type"))
}

// ValidateEnvironment returns a reference to field validate_environment of azurerm_virtual_desktop_host_pool.
func (vdhp dataVirtualDesktopHostPoolAttributes) ValidateEnvironment() terra.BoolValue {
	return terra.ReferenceAsBool(vdhp.ref.Append("validate_environment"))
}

func (vdhp dataVirtualDesktopHostPoolAttributes) ScheduledAgentUpdates() terra.ListValue[datavirtualdesktophostpool.ScheduledAgentUpdatesAttributes] {
	return terra.ReferenceAsList[datavirtualdesktophostpool.ScheduledAgentUpdatesAttributes](vdhp.ref.Append("scheduled_agent_updates"))
}

func (vdhp dataVirtualDesktopHostPoolAttributes) Timeouts() datavirtualdesktophostpool.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datavirtualdesktophostpool.TimeoutsAttributes](vdhp.ref.Append("timeouts"))
}
