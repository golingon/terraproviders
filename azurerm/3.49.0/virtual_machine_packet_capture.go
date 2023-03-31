// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	virtualmachinepacketcapture "github.com/golingon/terraproviders/azurerm/3.49.0/virtualmachinepacketcapture"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewVirtualMachinePacketCapture(name string, args VirtualMachinePacketCaptureArgs) *VirtualMachinePacketCapture {
	return &VirtualMachinePacketCapture{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VirtualMachinePacketCapture)(nil)

type VirtualMachinePacketCapture struct {
	Name  string
	Args  VirtualMachinePacketCaptureArgs
	state *virtualMachinePacketCaptureState
}

func (vmpc *VirtualMachinePacketCapture) Type() string {
	return "azurerm_virtual_machine_packet_capture"
}

func (vmpc *VirtualMachinePacketCapture) LocalName() string {
	return vmpc.Name
}

func (vmpc *VirtualMachinePacketCapture) Configuration() interface{} {
	return vmpc.Args
}

func (vmpc *VirtualMachinePacketCapture) Attributes() virtualMachinePacketCaptureAttributes {
	return virtualMachinePacketCaptureAttributes{ref: terra.ReferenceResource(vmpc)}
}

func (vmpc *VirtualMachinePacketCapture) ImportState(av io.Reader) error {
	vmpc.state = &virtualMachinePacketCaptureState{}
	if err := json.NewDecoder(av).Decode(vmpc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vmpc.Type(), vmpc.LocalName(), err)
	}
	return nil
}

func (vmpc *VirtualMachinePacketCapture) State() (*virtualMachinePacketCaptureState, bool) {
	return vmpc.state, vmpc.state != nil
}

func (vmpc *VirtualMachinePacketCapture) StateMust() *virtualMachinePacketCaptureState {
	if vmpc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vmpc.Type(), vmpc.LocalName()))
	}
	return vmpc.state
}

func (vmpc *VirtualMachinePacketCapture) DependOn() terra.Reference {
	return terra.ReferenceResource(vmpc)
}

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
	// DependsOn contains resources that VirtualMachinePacketCapture depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type virtualMachinePacketCaptureAttributes struct {
	ref terra.Reference
}

func (vmpc virtualMachinePacketCaptureAttributes) Id() terra.StringValue {
	return terra.ReferenceString(vmpc.ref.Append("id"))
}

func (vmpc virtualMachinePacketCaptureAttributes) MaximumBytesPerPacket() terra.NumberValue {
	return terra.ReferenceNumber(vmpc.ref.Append("maximum_bytes_per_packet"))
}

func (vmpc virtualMachinePacketCaptureAttributes) MaximumBytesPerSession() terra.NumberValue {
	return terra.ReferenceNumber(vmpc.ref.Append("maximum_bytes_per_session"))
}

func (vmpc virtualMachinePacketCaptureAttributes) MaximumCaptureDurationInSeconds() terra.NumberValue {
	return terra.ReferenceNumber(vmpc.ref.Append("maximum_capture_duration_in_seconds"))
}

func (vmpc virtualMachinePacketCaptureAttributes) Name() terra.StringValue {
	return terra.ReferenceString(vmpc.ref.Append("name"))
}

func (vmpc virtualMachinePacketCaptureAttributes) NetworkWatcherId() terra.StringValue {
	return terra.ReferenceString(vmpc.ref.Append("network_watcher_id"))
}

func (vmpc virtualMachinePacketCaptureAttributes) VirtualMachineId() terra.StringValue {
	return terra.ReferenceString(vmpc.ref.Append("virtual_machine_id"))
}

func (vmpc virtualMachinePacketCaptureAttributes) Filter() terra.ListValue[virtualmachinepacketcapture.FilterAttributes] {
	return terra.ReferenceList[virtualmachinepacketcapture.FilterAttributes](vmpc.ref.Append("filter"))
}

func (vmpc virtualMachinePacketCaptureAttributes) StorageLocation() terra.ListValue[virtualmachinepacketcapture.StorageLocationAttributes] {
	return terra.ReferenceList[virtualmachinepacketcapture.StorageLocationAttributes](vmpc.ref.Append("storage_location"))
}

func (vmpc virtualMachinePacketCaptureAttributes) Timeouts() virtualmachinepacketcapture.TimeoutsAttributes {
	return terra.ReferenceSingle[virtualmachinepacketcapture.TimeoutsAttributes](vmpc.ref.Append("timeouts"))
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
