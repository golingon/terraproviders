// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	virtualmachinepacketcapture "github.com/golingon/terraproviders/azurerm/3.67.0/virtualmachinepacketcapture"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVirtualMachinePacketCapture creates a new instance of [VirtualMachinePacketCapture].
func NewVirtualMachinePacketCapture(name string, args VirtualMachinePacketCaptureArgs) *VirtualMachinePacketCapture {
	return &VirtualMachinePacketCapture{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VirtualMachinePacketCapture)(nil)

// VirtualMachinePacketCapture represents the Terraform resource azurerm_virtual_machine_packet_capture.
type VirtualMachinePacketCapture struct {
	Name      string
	Args      VirtualMachinePacketCaptureArgs
	state     *virtualMachinePacketCaptureState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VirtualMachinePacketCapture].
func (vmpc *VirtualMachinePacketCapture) Type() string {
	return "azurerm_virtual_machine_packet_capture"
}

// LocalName returns the local name for [VirtualMachinePacketCapture].
func (vmpc *VirtualMachinePacketCapture) LocalName() string {
	return vmpc.Name
}

// Configuration returns the configuration (args) for [VirtualMachinePacketCapture].
func (vmpc *VirtualMachinePacketCapture) Configuration() interface{} {
	return vmpc.Args
}

// DependOn is used for other resources to depend on [VirtualMachinePacketCapture].
func (vmpc *VirtualMachinePacketCapture) DependOn() terra.Reference {
	return terra.ReferenceResource(vmpc)
}

// Dependencies returns the list of resources [VirtualMachinePacketCapture] depends_on.
func (vmpc *VirtualMachinePacketCapture) Dependencies() terra.Dependencies {
	return vmpc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VirtualMachinePacketCapture].
func (vmpc *VirtualMachinePacketCapture) LifecycleManagement() *terra.Lifecycle {
	return vmpc.Lifecycle
}

// Attributes returns the attributes for [VirtualMachinePacketCapture].
func (vmpc *VirtualMachinePacketCapture) Attributes() virtualMachinePacketCaptureAttributes {
	return virtualMachinePacketCaptureAttributes{ref: terra.ReferenceResource(vmpc)}
}

// ImportState imports the given attribute values into [VirtualMachinePacketCapture]'s state.
func (vmpc *VirtualMachinePacketCapture) ImportState(av io.Reader) error {
	vmpc.state = &virtualMachinePacketCaptureState{}
	if err := json.NewDecoder(av).Decode(vmpc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vmpc.Type(), vmpc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VirtualMachinePacketCapture] has state.
func (vmpc *VirtualMachinePacketCapture) State() (*virtualMachinePacketCaptureState, bool) {
	return vmpc.state, vmpc.state != nil
}

// StateMust returns the state for [VirtualMachinePacketCapture]. Panics if the state is nil.
func (vmpc *VirtualMachinePacketCapture) StateMust() *virtualMachinePacketCaptureState {
	if vmpc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vmpc.Type(), vmpc.LocalName()))
	}
	return vmpc.state
}

// VirtualMachinePacketCaptureArgs contains the configurations for azurerm_virtual_machine_packet_capture.
type VirtualMachinePacketCaptureArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MaximumBytesPerPacket: number, optional
	MaximumBytesPerPacket terra.NumberValue `hcl:"maximum_bytes_per_packet,attr"`
	// MaximumBytesPerSession: number, optional
	MaximumBytesPerSession terra.NumberValue `hcl:"maximum_bytes_per_session,attr"`
	// MaximumCaptureDurationInSeconds: number, optional
	MaximumCaptureDurationInSeconds terra.NumberValue `hcl:"maximum_capture_duration_in_seconds,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NetworkWatcherId: string, required
	NetworkWatcherId terra.StringValue `hcl:"network_watcher_id,attr" validate:"required"`
	// VirtualMachineId: string, required
	VirtualMachineId terra.StringValue `hcl:"virtual_machine_id,attr" validate:"required"`
	// Filter: min=0
	Filter []virtualmachinepacketcapture.Filter `hcl:"filter,block" validate:"min=0"`
	// StorageLocation: required
	StorageLocation *virtualmachinepacketcapture.StorageLocation `hcl:"storage_location,block" validate:"required"`
	// Timeouts: optional
	Timeouts *virtualmachinepacketcapture.Timeouts `hcl:"timeouts,block"`
}
type virtualMachinePacketCaptureAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_virtual_machine_packet_capture.
func (vmpc virtualMachinePacketCaptureAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vmpc.ref.Append("id"))
}

// MaximumBytesPerPacket returns a reference to field maximum_bytes_per_packet of azurerm_virtual_machine_packet_capture.
func (vmpc virtualMachinePacketCaptureAttributes) MaximumBytesPerPacket() terra.NumberValue {
	return terra.ReferenceAsNumber(vmpc.ref.Append("maximum_bytes_per_packet"))
}

// MaximumBytesPerSession returns a reference to field maximum_bytes_per_session of azurerm_virtual_machine_packet_capture.
func (vmpc virtualMachinePacketCaptureAttributes) MaximumBytesPerSession() terra.NumberValue {
	return terra.ReferenceAsNumber(vmpc.ref.Append("maximum_bytes_per_session"))
}

// MaximumCaptureDurationInSeconds returns a reference to field maximum_capture_duration_in_seconds of azurerm_virtual_machine_packet_capture.
func (vmpc virtualMachinePacketCaptureAttributes) MaximumCaptureDurationInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(vmpc.ref.Append("maximum_capture_duration_in_seconds"))
}

// Name returns a reference to field name of azurerm_virtual_machine_packet_capture.
func (vmpc virtualMachinePacketCaptureAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vmpc.ref.Append("name"))
}

// NetworkWatcherId returns a reference to field network_watcher_id of azurerm_virtual_machine_packet_capture.
func (vmpc virtualMachinePacketCaptureAttributes) NetworkWatcherId() terra.StringValue {
	return terra.ReferenceAsString(vmpc.ref.Append("network_watcher_id"))
}

// VirtualMachineId returns a reference to field virtual_machine_id of azurerm_virtual_machine_packet_capture.
func (vmpc virtualMachinePacketCaptureAttributes) VirtualMachineId() terra.StringValue {
	return terra.ReferenceAsString(vmpc.ref.Append("virtual_machine_id"))
}

func (vmpc virtualMachinePacketCaptureAttributes) Filter() terra.ListValue[virtualmachinepacketcapture.FilterAttributes] {
	return terra.ReferenceAsList[virtualmachinepacketcapture.FilterAttributes](vmpc.ref.Append("filter"))
}

func (vmpc virtualMachinePacketCaptureAttributes) StorageLocation() terra.ListValue[virtualmachinepacketcapture.StorageLocationAttributes] {
	return terra.ReferenceAsList[virtualmachinepacketcapture.StorageLocationAttributes](vmpc.ref.Append("storage_location"))
}

func (vmpc virtualMachinePacketCaptureAttributes) Timeouts() virtualmachinepacketcapture.TimeoutsAttributes {
	return terra.ReferenceAsSingle[virtualmachinepacketcapture.TimeoutsAttributes](vmpc.ref.Append("timeouts"))
}

type virtualMachinePacketCaptureState struct {
	Id                              string                                             `json:"id"`
	MaximumBytesPerPacket           float64                                            `json:"maximum_bytes_per_packet"`
	MaximumBytesPerSession          float64                                            `json:"maximum_bytes_per_session"`
	MaximumCaptureDurationInSeconds float64                                            `json:"maximum_capture_duration_in_seconds"`
	Name                            string                                             `json:"name"`
	NetworkWatcherId                string                                             `json:"network_watcher_id"`
	VirtualMachineId                string                                             `json:"virtual_machine_id"`
	Filter                          []virtualmachinepacketcapture.FilterState          `json:"filter"`
	StorageLocation                 []virtualmachinepacketcapture.StorageLocationState `json:"storage_location"`
	Timeouts                        *virtualmachinepacketcapture.TimeoutsState         `json:"timeouts"`
}
