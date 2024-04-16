// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_vmware_netapp_volume_attachment

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource azurerm_vmware_netapp_volume_attachment.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermVmwareNetappVolumeAttachmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (avnva *Resource) Type() string {
	return "azurerm_vmware_netapp_volume_attachment"
}

// LocalName returns the local name for [Resource].
func (avnva *Resource) LocalName() string {
	return avnva.Name
}

// Configuration returns the configuration (args) for [Resource].
func (avnva *Resource) Configuration() interface{} {
	return avnva.Args
}

// DependOn is used for other resources to depend on [Resource].
func (avnva *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(avnva)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (avnva *Resource) Dependencies() terra.Dependencies {
	return avnva.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (avnva *Resource) LifecycleManagement() *terra.Lifecycle {
	return avnva.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (avnva *Resource) Attributes() azurermVmwareNetappVolumeAttachmentAttributes {
	return azurermVmwareNetappVolumeAttachmentAttributes{ref: terra.ReferenceResource(avnva)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (avnva *Resource) ImportState(state io.Reader) error {
	avnva.state = &azurermVmwareNetappVolumeAttachmentState{}
	if err := json.NewDecoder(state).Decode(avnva.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", avnva.Type(), avnva.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (avnva *Resource) State() (*azurermVmwareNetappVolumeAttachmentState, bool) {
	return avnva.state, avnva.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (avnva *Resource) StateMust() *azurermVmwareNetappVolumeAttachmentState {
	if avnva.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", avnva.Type(), avnva.LocalName()))
	}
	return avnva.state
}

// Args contains the configurations for azurerm_vmware_netapp_volume_attachment.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NetappVolumeId: string, required
	NetappVolumeId terra.StringValue `hcl:"netapp_volume_id,attr" validate:"required"`
	// VmwareClusterId: string, required
	VmwareClusterId terra.StringValue `hcl:"vmware_cluster_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermVmwareNetappVolumeAttachmentAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_vmware_netapp_volume_attachment.
func (avnva azurermVmwareNetappVolumeAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(avnva.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_vmware_netapp_volume_attachment.
func (avnva azurermVmwareNetappVolumeAttachmentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(avnva.ref.Append("name"))
}

// NetappVolumeId returns a reference to field netapp_volume_id of azurerm_vmware_netapp_volume_attachment.
func (avnva azurermVmwareNetappVolumeAttachmentAttributes) NetappVolumeId() terra.StringValue {
	return terra.ReferenceAsString(avnva.ref.Append("netapp_volume_id"))
}

// VmwareClusterId returns a reference to field vmware_cluster_id of azurerm_vmware_netapp_volume_attachment.
func (avnva azurermVmwareNetappVolumeAttachmentAttributes) VmwareClusterId() terra.StringValue {
	return terra.ReferenceAsString(avnva.ref.Append("vmware_cluster_id"))
}

func (avnva azurermVmwareNetappVolumeAttachmentAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](avnva.ref.Append("timeouts"))
}

type azurermVmwareNetappVolumeAttachmentState struct {
	Id              string         `json:"id"`
	Name            string         `json:"name"`
	NetappVolumeId  string         `json:"netapp_volume_id"`
	VmwareClusterId string         `json:"vmware_cluster_id"`
	Timeouts        *TimeoutsState `json:"timeouts"`
}
