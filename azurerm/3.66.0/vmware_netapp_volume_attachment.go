// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	vmwarenetappvolumeattachment "github.com/golingon/terraproviders/azurerm/3.66.0/vmwarenetappvolumeattachment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVmwareNetappVolumeAttachment creates a new instance of [VmwareNetappVolumeAttachment].
func NewVmwareNetappVolumeAttachment(name string, args VmwareNetappVolumeAttachmentArgs) *VmwareNetappVolumeAttachment {
	return &VmwareNetappVolumeAttachment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VmwareNetappVolumeAttachment)(nil)

// VmwareNetappVolumeAttachment represents the Terraform resource azurerm_vmware_netapp_volume_attachment.
type VmwareNetappVolumeAttachment struct {
	Name      string
	Args      VmwareNetappVolumeAttachmentArgs
	state     *vmwareNetappVolumeAttachmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VmwareNetappVolumeAttachment].
func (vnva *VmwareNetappVolumeAttachment) Type() string {
	return "azurerm_vmware_netapp_volume_attachment"
}

// LocalName returns the local name for [VmwareNetappVolumeAttachment].
func (vnva *VmwareNetappVolumeAttachment) LocalName() string {
	return vnva.Name
}

// Configuration returns the configuration (args) for [VmwareNetappVolumeAttachment].
func (vnva *VmwareNetappVolumeAttachment) Configuration() interface{} {
	return vnva.Args
}

// DependOn is used for other resources to depend on [VmwareNetappVolumeAttachment].
func (vnva *VmwareNetappVolumeAttachment) DependOn() terra.Reference {
	return terra.ReferenceResource(vnva)
}

// Dependencies returns the list of resources [VmwareNetappVolumeAttachment] depends_on.
func (vnva *VmwareNetappVolumeAttachment) Dependencies() terra.Dependencies {
	return vnva.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VmwareNetappVolumeAttachment].
func (vnva *VmwareNetappVolumeAttachment) LifecycleManagement() *terra.Lifecycle {
	return vnva.Lifecycle
}

// Attributes returns the attributes for [VmwareNetappVolumeAttachment].
func (vnva *VmwareNetappVolumeAttachment) Attributes() vmwareNetappVolumeAttachmentAttributes {
	return vmwareNetappVolumeAttachmentAttributes{ref: terra.ReferenceResource(vnva)}
}

// ImportState imports the given attribute values into [VmwareNetappVolumeAttachment]'s state.
func (vnva *VmwareNetappVolumeAttachment) ImportState(av io.Reader) error {
	vnva.state = &vmwareNetappVolumeAttachmentState{}
	if err := json.NewDecoder(av).Decode(vnva.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vnva.Type(), vnva.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VmwareNetappVolumeAttachment] has state.
func (vnva *VmwareNetappVolumeAttachment) State() (*vmwareNetappVolumeAttachmentState, bool) {
	return vnva.state, vnva.state != nil
}

// StateMust returns the state for [VmwareNetappVolumeAttachment]. Panics if the state is nil.
func (vnva *VmwareNetappVolumeAttachment) StateMust() *vmwareNetappVolumeAttachmentState {
	if vnva.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vnva.Type(), vnva.LocalName()))
	}
	return vnva.state
}

// VmwareNetappVolumeAttachmentArgs contains the configurations for azurerm_vmware_netapp_volume_attachment.
type VmwareNetappVolumeAttachmentArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NetappVolumeId: string, required
	NetappVolumeId terra.StringValue `hcl:"netapp_volume_id,attr" validate:"required"`
	// VmwareClusterId: string, required
	VmwareClusterId terra.StringValue `hcl:"vmware_cluster_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *vmwarenetappvolumeattachment.Timeouts `hcl:"timeouts,block"`
}
type vmwareNetappVolumeAttachmentAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_vmware_netapp_volume_attachment.
func (vnva vmwareNetappVolumeAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vnva.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_vmware_netapp_volume_attachment.
func (vnva vmwareNetappVolumeAttachmentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vnva.ref.Append("name"))
}

// NetappVolumeId returns a reference to field netapp_volume_id of azurerm_vmware_netapp_volume_attachment.
func (vnva vmwareNetappVolumeAttachmentAttributes) NetappVolumeId() terra.StringValue {
	return terra.ReferenceAsString(vnva.ref.Append("netapp_volume_id"))
}

// VmwareClusterId returns a reference to field vmware_cluster_id of azurerm_vmware_netapp_volume_attachment.
func (vnva vmwareNetappVolumeAttachmentAttributes) VmwareClusterId() terra.StringValue {
	return terra.ReferenceAsString(vnva.ref.Append("vmware_cluster_id"))
}

func (vnva vmwareNetappVolumeAttachmentAttributes) Timeouts() vmwarenetappvolumeattachment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vmwarenetappvolumeattachment.TimeoutsAttributes](vnva.ref.Append("timeouts"))
}

type vmwareNetappVolumeAttachmentState struct {
	Id              string                                      `json:"id"`
	Name            string                                      `json:"name"`
	NetappVolumeId  string                                      `json:"netapp_volume_id"`
	VmwareClusterId string                                      `json:"vmware_cluster_id"`
	Timeouts        *vmwarenetappvolumeattachment.TimeoutsState `json:"timeouts"`
}
