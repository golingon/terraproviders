// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	diskpoolmanageddiskattachment "github.com/golingon/terraproviders/azurerm/3.64.0/diskpoolmanageddiskattachment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDiskPoolManagedDiskAttachment creates a new instance of [DiskPoolManagedDiskAttachment].
func NewDiskPoolManagedDiskAttachment(name string, args DiskPoolManagedDiskAttachmentArgs) *DiskPoolManagedDiskAttachment {
	return &DiskPoolManagedDiskAttachment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DiskPoolManagedDiskAttachment)(nil)

// DiskPoolManagedDiskAttachment represents the Terraform resource azurerm_disk_pool_managed_disk_attachment.
type DiskPoolManagedDiskAttachment struct {
	Name      string
	Args      DiskPoolManagedDiskAttachmentArgs
	state     *diskPoolManagedDiskAttachmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DiskPoolManagedDiskAttachment].
func (dpmda *DiskPoolManagedDiskAttachment) Type() string {
	return "azurerm_disk_pool_managed_disk_attachment"
}

// LocalName returns the local name for [DiskPoolManagedDiskAttachment].
func (dpmda *DiskPoolManagedDiskAttachment) LocalName() string {
	return dpmda.Name
}

// Configuration returns the configuration (args) for [DiskPoolManagedDiskAttachment].
func (dpmda *DiskPoolManagedDiskAttachment) Configuration() interface{} {
	return dpmda.Args
}

// DependOn is used for other resources to depend on [DiskPoolManagedDiskAttachment].
func (dpmda *DiskPoolManagedDiskAttachment) DependOn() terra.Reference {
	return terra.ReferenceResource(dpmda)
}

// Dependencies returns the list of resources [DiskPoolManagedDiskAttachment] depends_on.
func (dpmda *DiskPoolManagedDiskAttachment) Dependencies() terra.Dependencies {
	return dpmda.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DiskPoolManagedDiskAttachment].
func (dpmda *DiskPoolManagedDiskAttachment) LifecycleManagement() *terra.Lifecycle {
	return dpmda.Lifecycle
}

// Attributes returns the attributes for [DiskPoolManagedDiskAttachment].
func (dpmda *DiskPoolManagedDiskAttachment) Attributes() diskPoolManagedDiskAttachmentAttributes {
	return diskPoolManagedDiskAttachmentAttributes{ref: terra.ReferenceResource(dpmda)}
}

// ImportState imports the given attribute values into [DiskPoolManagedDiskAttachment]'s state.
func (dpmda *DiskPoolManagedDiskAttachment) ImportState(av io.Reader) error {
	dpmda.state = &diskPoolManagedDiskAttachmentState{}
	if err := json.NewDecoder(av).Decode(dpmda.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dpmda.Type(), dpmda.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DiskPoolManagedDiskAttachment] has state.
func (dpmda *DiskPoolManagedDiskAttachment) State() (*diskPoolManagedDiskAttachmentState, bool) {
	return dpmda.state, dpmda.state != nil
}

// StateMust returns the state for [DiskPoolManagedDiskAttachment]. Panics if the state is nil.
func (dpmda *DiskPoolManagedDiskAttachment) StateMust() *diskPoolManagedDiskAttachmentState {
	if dpmda.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dpmda.Type(), dpmda.LocalName()))
	}
	return dpmda.state
}

// DiskPoolManagedDiskAttachmentArgs contains the configurations for azurerm_disk_pool_managed_disk_attachment.
type DiskPoolManagedDiskAttachmentArgs struct {
	// DiskPoolId: string, required
	DiskPoolId terra.StringValue `hcl:"disk_pool_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ManagedDiskId: string, required
	ManagedDiskId terra.StringValue `hcl:"managed_disk_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *diskpoolmanageddiskattachment.Timeouts `hcl:"timeouts,block"`
}
type diskPoolManagedDiskAttachmentAttributes struct {
	ref terra.Reference
}

// DiskPoolId returns a reference to field disk_pool_id of azurerm_disk_pool_managed_disk_attachment.
func (dpmda diskPoolManagedDiskAttachmentAttributes) DiskPoolId() terra.StringValue {
	return terra.ReferenceAsString(dpmda.ref.Append("disk_pool_id"))
}

// Id returns a reference to field id of azurerm_disk_pool_managed_disk_attachment.
func (dpmda diskPoolManagedDiskAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dpmda.ref.Append("id"))
}

// ManagedDiskId returns a reference to field managed_disk_id of azurerm_disk_pool_managed_disk_attachment.
func (dpmda diskPoolManagedDiskAttachmentAttributes) ManagedDiskId() terra.StringValue {
	return terra.ReferenceAsString(dpmda.ref.Append("managed_disk_id"))
}

func (dpmda diskPoolManagedDiskAttachmentAttributes) Timeouts() diskpoolmanageddiskattachment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[diskpoolmanageddiskattachment.TimeoutsAttributes](dpmda.ref.Append("timeouts"))
}

type diskPoolManagedDiskAttachmentState struct {
	DiskPoolId    string                                       `json:"disk_pool_id"`
	Id            string                                       `json:"id"`
	ManagedDiskId string                                       `json:"managed_disk_id"`
	Timeouts      *diskpoolmanageddiskattachment.TimeoutsState `json:"timeouts"`
}
