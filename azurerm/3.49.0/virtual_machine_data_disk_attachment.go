// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	virtualmachinedatadiskattachment "github.com/golingon/terraproviders/azurerm/3.49.0/virtualmachinedatadiskattachment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVirtualMachineDataDiskAttachment creates a new instance of [VirtualMachineDataDiskAttachment].
func NewVirtualMachineDataDiskAttachment(name string, args VirtualMachineDataDiskAttachmentArgs) *VirtualMachineDataDiskAttachment {
	return &VirtualMachineDataDiskAttachment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VirtualMachineDataDiskAttachment)(nil)

// VirtualMachineDataDiskAttachment represents the Terraform resource azurerm_virtual_machine_data_disk_attachment.
type VirtualMachineDataDiskAttachment struct {
	Name      string
	Args      VirtualMachineDataDiskAttachmentArgs
	state     *virtualMachineDataDiskAttachmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VirtualMachineDataDiskAttachment].
func (vmdda *VirtualMachineDataDiskAttachment) Type() string {
	return "azurerm_virtual_machine_data_disk_attachment"
}

// LocalName returns the local name for [VirtualMachineDataDiskAttachment].
func (vmdda *VirtualMachineDataDiskAttachment) LocalName() string {
	return vmdda.Name
}

// Configuration returns the configuration (args) for [VirtualMachineDataDiskAttachment].
func (vmdda *VirtualMachineDataDiskAttachment) Configuration() interface{} {
	return vmdda.Args
}

// DependOn is used for other resources to depend on [VirtualMachineDataDiskAttachment].
func (vmdda *VirtualMachineDataDiskAttachment) DependOn() terra.Reference {
	return terra.ReferenceResource(vmdda)
}

// Dependencies returns the list of resources [VirtualMachineDataDiskAttachment] depends_on.
func (vmdda *VirtualMachineDataDiskAttachment) Dependencies() terra.Dependencies {
	return vmdda.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VirtualMachineDataDiskAttachment].
func (vmdda *VirtualMachineDataDiskAttachment) LifecycleManagement() *terra.Lifecycle {
	return vmdda.Lifecycle
}

// Attributes returns the attributes for [VirtualMachineDataDiskAttachment].
func (vmdda *VirtualMachineDataDiskAttachment) Attributes() virtualMachineDataDiskAttachmentAttributes {
	return virtualMachineDataDiskAttachmentAttributes{ref: terra.ReferenceResource(vmdda)}
}

// ImportState imports the given attribute values into [VirtualMachineDataDiskAttachment]'s state.
func (vmdda *VirtualMachineDataDiskAttachment) ImportState(av io.Reader) error {
	vmdda.state = &virtualMachineDataDiskAttachmentState{}
	if err := json.NewDecoder(av).Decode(vmdda.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vmdda.Type(), vmdda.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VirtualMachineDataDiskAttachment] has state.
func (vmdda *VirtualMachineDataDiskAttachment) State() (*virtualMachineDataDiskAttachmentState, bool) {
	return vmdda.state, vmdda.state != nil
}

// StateMust returns the state for [VirtualMachineDataDiskAttachment]. Panics if the state is nil.
func (vmdda *VirtualMachineDataDiskAttachment) StateMust() *virtualMachineDataDiskAttachmentState {
	if vmdda.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vmdda.Type(), vmdda.LocalName()))
	}
	return vmdda.state
}

// VirtualMachineDataDiskAttachmentArgs contains the configurations for azurerm_virtual_machine_data_disk_attachment.
type VirtualMachineDataDiskAttachmentArgs struct {
	// Caching: string, required
	Caching terra.StringValue `hcl:"caching,attr" validate:"required"`
	// CreateOption: string, optional
	CreateOption terra.StringValue `hcl:"create_option,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Lun: number, required
	Lun terra.NumberValue `hcl:"lun,attr" validate:"required"`
	// ManagedDiskId: string, required
	ManagedDiskId terra.StringValue `hcl:"managed_disk_id,attr" validate:"required"`
	// VirtualMachineId: string, required
	VirtualMachineId terra.StringValue `hcl:"virtual_machine_id,attr" validate:"required"`
	// WriteAcceleratorEnabled: bool, optional
	WriteAcceleratorEnabled terra.BoolValue `hcl:"write_accelerator_enabled,attr"`
	// Timeouts: optional
	Timeouts *virtualmachinedatadiskattachment.Timeouts `hcl:"timeouts,block"`
}
type virtualMachineDataDiskAttachmentAttributes struct {
	ref terra.Reference
}

// Caching returns a reference to field caching of azurerm_virtual_machine_data_disk_attachment.
func (vmdda virtualMachineDataDiskAttachmentAttributes) Caching() terra.StringValue {
	return terra.ReferenceAsString(vmdda.ref.Append("caching"))
}

// CreateOption returns a reference to field create_option of azurerm_virtual_machine_data_disk_attachment.
func (vmdda virtualMachineDataDiskAttachmentAttributes) CreateOption() terra.StringValue {
	return terra.ReferenceAsString(vmdda.ref.Append("create_option"))
}

// Id returns a reference to field id of azurerm_virtual_machine_data_disk_attachment.
func (vmdda virtualMachineDataDiskAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vmdda.ref.Append("id"))
}

// Lun returns a reference to field lun of azurerm_virtual_machine_data_disk_attachment.
func (vmdda virtualMachineDataDiskAttachmentAttributes) Lun() terra.NumberValue {
	return terra.ReferenceAsNumber(vmdda.ref.Append("lun"))
}

// ManagedDiskId returns a reference to field managed_disk_id of azurerm_virtual_machine_data_disk_attachment.
func (vmdda virtualMachineDataDiskAttachmentAttributes) ManagedDiskId() terra.StringValue {
	return terra.ReferenceAsString(vmdda.ref.Append("managed_disk_id"))
}

// VirtualMachineId returns a reference to field virtual_machine_id of azurerm_virtual_machine_data_disk_attachment.
func (vmdda virtualMachineDataDiskAttachmentAttributes) VirtualMachineId() terra.StringValue {
	return terra.ReferenceAsString(vmdda.ref.Append("virtual_machine_id"))
}

// WriteAcceleratorEnabled returns a reference to field write_accelerator_enabled of azurerm_virtual_machine_data_disk_attachment.
func (vmdda virtualMachineDataDiskAttachmentAttributes) WriteAcceleratorEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(vmdda.ref.Append("write_accelerator_enabled"))
}

func (vmdda virtualMachineDataDiskAttachmentAttributes) Timeouts() virtualmachinedatadiskattachment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[virtualmachinedatadiskattachment.TimeoutsAttributes](vmdda.ref.Append("timeouts"))
}

type virtualMachineDataDiskAttachmentState struct {
	Caching                 string                                          `json:"caching"`
	CreateOption            string                                          `json:"create_option"`
	Id                      string                                          `json:"id"`
	Lun                     float64                                         `json:"lun"`
	ManagedDiskId           string                                          `json:"managed_disk_id"`
	VirtualMachineId        string                                          `json:"virtual_machine_id"`
	WriteAcceleratorEnabled bool                                            `json:"write_accelerator_enabled"`
	Timeouts                *virtualmachinedatadiskattachment.TimeoutsState `json:"timeouts"`
}
