// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	networkpacketcapture "github.com/golingon/terraproviders/azurerm/3.68.0/networkpacketcapture"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkPacketCapture creates a new instance of [NetworkPacketCapture].
func NewNetworkPacketCapture(name string, args NetworkPacketCaptureArgs) *NetworkPacketCapture {
	return &NetworkPacketCapture{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkPacketCapture)(nil)

// NetworkPacketCapture represents the Terraform resource azurerm_network_packet_capture.
type NetworkPacketCapture struct {
	Name      string
	Args      NetworkPacketCaptureArgs
	state     *networkPacketCaptureState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkPacketCapture].
func (npc *NetworkPacketCapture) Type() string {
	return "azurerm_network_packet_capture"
}

// LocalName returns the local name for [NetworkPacketCapture].
func (npc *NetworkPacketCapture) LocalName() string {
	return npc.Name
}

// Configuration returns the configuration (args) for [NetworkPacketCapture].
func (npc *NetworkPacketCapture) Configuration() interface{} {
	return npc.Args
}

// DependOn is used for other resources to depend on [NetworkPacketCapture].
func (npc *NetworkPacketCapture) DependOn() terra.Reference {
	return terra.ReferenceResource(npc)
}

// Dependencies returns the list of resources [NetworkPacketCapture] depends_on.
func (npc *NetworkPacketCapture) Dependencies() terra.Dependencies {
	return npc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkPacketCapture].
func (npc *NetworkPacketCapture) LifecycleManagement() *terra.Lifecycle {
	return npc.Lifecycle
}

// Attributes returns the attributes for [NetworkPacketCapture].
func (npc *NetworkPacketCapture) Attributes() networkPacketCaptureAttributes {
	return networkPacketCaptureAttributes{ref: terra.ReferenceResource(npc)}
}

// ImportState imports the given attribute values into [NetworkPacketCapture]'s state.
func (npc *NetworkPacketCapture) ImportState(av io.Reader) error {
	npc.state = &networkPacketCaptureState{}
	if err := json.NewDecoder(av).Decode(npc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", npc.Type(), npc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkPacketCapture] has state.
func (npc *NetworkPacketCapture) State() (*networkPacketCaptureState, bool) {
	return npc.state, npc.state != nil
}

// StateMust returns the state for [NetworkPacketCapture]. Panics if the state is nil.
func (npc *NetworkPacketCapture) StateMust() *networkPacketCaptureState {
	if npc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", npc.Type(), npc.LocalName()))
	}
	return npc.state
}

// NetworkPacketCaptureArgs contains the configurations for azurerm_network_packet_capture.
type NetworkPacketCaptureArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MaximumBytesPerPacket: number, optional
	MaximumBytesPerPacket terra.NumberValue `hcl:"maximum_bytes_per_packet,attr"`
	// MaximumBytesPerSession: number, optional
	MaximumBytesPerSession terra.NumberValue `hcl:"maximum_bytes_per_session,attr"`
	// MaximumCaptureDuration: number, optional
	MaximumCaptureDuration terra.NumberValue `hcl:"maximum_capture_duration,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NetworkWatcherName: string, required
	NetworkWatcherName terra.StringValue `hcl:"network_watcher_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// TargetResourceId: string, required
	TargetResourceId terra.StringValue `hcl:"target_resource_id,attr" validate:"required"`
	// Filter: min=0
	Filter []networkpacketcapture.Filter `hcl:"filter,block" validate:"min=0"`
	// StorageLocation: required
	StorageLocation *networkpacketcapture.StorageLocation `hcl:"storage_location,block" validate:"required"`
	// Timeouts: optional
	Timeouts *networkpacketcapture.Timeouts `hcl:"timeouts,block"`
}
type networkPacketCaptureAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_network_packet_capture.
func (npc networkPacketCaptureAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(npc.ref.Append("id"))
}

// MaximumBytesPerPacket returns a reference to field maximum_bytes_per_packet of azurerm_network_packet_capture.
func (npc networkPacketCaptureAttributes) MaximumBytesPerPacket() terra.NumberValue {
	return terra.ReferenceAsNumber(npc.ref.Append("maximum_bytes_per_packet"))
}

// MaximumBytesPerSession returns a reference to field maximum_bytes_per_session of azurerm_network_packet_capture.
func (npc networkPacketCaptureAttributes) MaximumBytesPerSession() terra.NumberValue {
	return terra.ReferenceAsNumber(npc.ref.Append("maximum_bytes_per_session"))
}

// MaximumCaptureDuration returns a reference to field maximum_capture_duration of azurerm_network_packet_capture.
func (npc networkPacketCaptureAttributes) MaximumCaptureDuration() terra.NumberValue {
	return terra.ReferenceAsNumber(npc.ref.Append("maximum_capture_duration"))
}

// Name returns a reference to field name of azurerm_network_packet_capture.
func (npc networkPacketCaptureAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(npc.ref.Append("name"))
}

// NetworkWatcherName returns a reference to field network_watcher_name of azurerm_network_packet_capture.
func (npc networkPacketCaptureAttributes) NetworkWatcherName() terra.StringValue {
	return terra.ReferenceAsString(npc.ref.Append("network_watcher_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_network_packet_capture.
func (npc networkPacketCaptureAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(npc.ref.Append("resource_group_name"))
}

// TargetResourceId returns a reference to field target_resource_id of azurerm_network_packet_capture.
func (npc networkPacketCaptureAttributes) TargetResourceId() terra.StringValue {
	return terra.ReferenceAsString(npc.ref.Append("target_resource_id"))
}

func (npc networkPacketCaptureAttributes) Filter() terra.ListValue[networkpacketcapture.FilterAttributes] {
	return terra.ReferenceAsList[networkpacketcapture.FilterAttributes](npc.ref.Append("filter"))
}

func (npc networkPacketCaptureAttributes) StorageLocation() terra.ListValue[networkpacketcapture.StorageLocationAttributes] {
	return terra.ReferenceAsList[networkpacketcapture.StorageLocationAttributes](npc.ref.Append("storage_location"))
}

func (npc networkPacketCaptureAttributes) Timeouts() networkpacketcapture.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkpacketcapture.TimeoutsAttributes](npc.ref.Append("timeouts"))
}

type networkPacketCaptureState struct {
	Id                     string                                      `json:"id"`
	MaximumBytesPerPacket  float64                                     `json:"maximum_bytes_per_packet"`
	MaximumBytesPerSession float64                                     `json:"maximum_bytes_per_session"`
	MaximumCaptureDuration float64                                     `json:"maximum_capture_duration"`
	Name                   string                                      `json:"name"`
	NetworkWatcherName     string                                      `json:"network_watcher_name"`
	ResourceGroupName      string                                      `json:"resource_group_name"`
	TargetResourceId       string                                      `json:"target_resource_id"`
	Filter                 []networkpacketcapture.FilterState          `json:"filter"`
	StorageLocation        []networkpacketcapture.StorageLocationState `json:"storage_location"`
	Timeouts               *networkpacketcapture.TimeoutsState         `json:"timeouts"`
}