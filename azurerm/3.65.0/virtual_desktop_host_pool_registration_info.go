// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	virtualdesktophostpoolregistrationinfo "github.com/golingon/terraproviders/azurerm/3.65.0/virtualdesktophostpoolregistrationinfo"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVirtualDesktopHostPoolRegistrationInfo creates a new instance of [VirtualDesktopHostPoolRegistrationInfo].
func NewVirtualDesktopHostPoolRegistrationInfo(name string, args VirtualDesktopHostPoolRegistrationInfoArgs) *VirtualDesktopHostPoolRegistrationInfo {
	return &VirtualDesktopHostPoolRegistrationInfo{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VirtualDesktopHostPoolRegistrationInfo)(nil)

// VirtualDesktopHostPoolRegistrationInfo represents the Terraform resource azurerm_virtual_desktop_host_pool_registration_info.
type VirtualDesktopHostPoolRegistrationInfo struct {
	Name      string
	Args      VirtualDesktopHostPoolRegistrationInfoArgs
	state     *virtualDesktopHostPoolRegistrationInfoState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VirtualDesktopHostPoolRegistrationInfo].
func (vdhpri *VirtualDesktopHostPoolRegistrationInfo) Type() string {
	return "azurerm_virtual_desktop_host_pool_registration_info"
}

// LocalName returns the local name for [VirtualDesktopHostPoolRegistrationInfo].
func (vdhpri *VirtualDesktopHostPoolRegistrationInfo) LocalName() string {
	return vdhpri.Name
}

// Configuration returns the configuration (args) for [VirtualDesktopHostPoolRegistrationInfo].
func (vdhpri *VirtualDesktopHostPoolRegistrationInfo) Configuration() interface{} {
	return vdhpri.Args
}

// DependOn is used for other resources to depend on [VirtualDesktopHostPoolRegistrationInfo].
func (vdhpri *VirtualDesktopHostPoolRegistrationInfo) DependOn() terra.Reference {
	return terra.ReferenceResource(vdhpri)
}

// Dependencies returns the list of resources [VirtualDesktopHostPoolRegistrationInfo] depends_on.
func (vdhpri *VirtualDesktopHostPoolRegistrationInfo) Dependencies() terra.Dependencies {
	return vdhpri.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VirtualDesktopHostPoolRegistrationInfo].
func (vdhpri *VirtualDesktopHostPoolRegistrationInfo) LifecycleManagement() *terra.Lifecycle {
	return vdhpri.Lifecycle
}

// Attributes returns the attributes for [VirtualDesktopHostPoolRegistrationInfo].
func (vdhpri *VirtualDesktopHostPoolRegistrationInfo) Attributes() virtualDesktopHostPoolRegistrationInfoAttributes {
	return virtualDesktopHostPoolRegistrationInfoAttributes{ref: terra.ReferenceResource(vdhpri)}
}

// ImportState imports the given attribute values into [VirtualDesktopHostPoolRegistrationInfo]'s state.
func (vdhpri *VirtualDesktopHostPoolRegistrationInfo) ImportState(av io.Reader) error {
	vdhpri.state = &virtualDesktopHostPoolRegistrationInfoState{}
	if err := json.NewDecoder(av).Decode(vdhpri.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vdhpri.Type(), vdhpri.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VirtualDesktopHostPoolRegistrationInfo] has state.
func (vdhpri *VirtualDesktopHostPoolRegistrationInfo) State() (*virtualDesktopHostPoolRegistrationInfoState, bool) {
	return vdhpri.state, vdhpri.state != nil
}

// StateMust returns the state for [VirtualDesktopHostPoolRegistrationInfo]. Panics if the state is nil.
func (vdhpri *VirtualDesktopHostPoolRegistrationInfo) StateMust() *virtualDesktopHostPoolRegistrationInfoState {
	if vdhpri.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vdhpri.Type(), vdhpri.LocalName()))
	}
	return vdhpri.state
}

// VirtualDesktopHostPoolRegistrationInfoArgs contains the configurations for azurerm_virtual_desktop_host_pool_registration_info.
type VirtualDesktopHostPoolRegistrationInfoArgs struct {
	// ExpirationDate: string, required
	ExpirationDate terra.StringValue `hcl:"expiration_date,attr" validate:"required"`
	// HostpoolId: string, required
	HostpoolId terra.StringValue `hcl:"hostpool_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Timeouts: optional
	Timeouts *virtualdesktophostpoolregistrationinfo.Timeouts `hcl:"timeouts,block"`
}
type virtualDesktopHostPoolRegistrationInfoAttributes struct {
	ref terra.Reference
}

// ExpirationDate returns a reference to field expiration_date of azurerm_virtual_desktop_host_pool_registration_info.
func (vdhpri virtualDesktopHostPoolRegistrationInfoAttributes) ExpirationDate() terra.StringValue {
	return terra.ReferenceAsString(vdhpri.ref.Append("expiration_date"))
}

// HostpoolId returns a reference to field hostpool_id of azurerm_virtual_desktop_host_pool_registration_info.
func (vdhpri virtualDesktopHostPoolRegistrationInfoAttributes) HostpoolId() terra.StringValue {
	return terra.ReferenceAsString(vdhpri.ref.Append("hostpool_id"))
}

// Id returns a reference to field id of azurerm_virtual_desktop_host_pool_registration_info.
func (vdhpri virtualDesktopHostPoolRegistrationInfoAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vdhpri.ref.Append("id"))
}

// Token returns a reference to field token of azurerm_virtual_desktop_host_pool_registration_info.
func (vdhpri virtualDesktopHostPoolRegistrationInfoAttributes) Token() terra.StringValue {
	return terra.ReferenceAsString(vdhpri.ref.Append("token"))
}

func (vdhpri virtualDesktopHostPoolRegistrationInfoAttributes) Timeouts() virtualdesktophostpoolregistrationinfo.TimeoutsAttributes {
	return terra.ReferenceAsSingle[virtualdesktophostpoolregistrationinfo.TimeoutsAttributes](vdhpri.ref.Append("timeouts"))
}

type virtualDesktopHostPoolRegistrationInfoState struct {
	ExpirationDate string                                                `json:"expiration_date"`
	HostpoolId     string                                                `json:"hostpool_id"`
	Id             string                                                `json:"id"`
	Token          string                                                `json:"token"`
	Timeouts       *virtualdesktophostpoolregistrationinfo.TimeoutsState `json:"timeouts"`
}
