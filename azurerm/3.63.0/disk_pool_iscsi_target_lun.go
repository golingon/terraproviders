// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	diskpooliscsitargetlun "github.com/golingon/terraproviders/azurerm/3.63.0/diskpooliscsitargetlun"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDiskPoolIscsiTargetLun creates a new instance of [DiskPoolIscsiTargetLun].
func NewDiskPoolIscsiTargetLun(name string, args DiskPoolIscsiTargetLunArgs) *DiskPoolIscsiTargetLun {
	return &DiskPoolIscsiTargetLun{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DiskPoolIscsiTargetLun)(nil)

// DiskPoolIscsiTargetLun represents the Terraform resource azurerm_disk_pool_iscsi_target_lun.
type DiskPoolIscsiTargetLun struct {
	Name      string
	Args      DiskPoolIscsiTargetLunArgs
	state     *diskPoolIscsiTargetLunState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DiskPoolIscsiTargetLun].
func (dpitl *DiskPoolIscsiTargetLun) Type() string {
	return "azurerm_disk_pool_iscsi_target_lun"
}

// LocalName returns the local name for [DiskPoolIscsiTargetLun].
func (dpitl *DiskPoolIscsiTargetLun) LocalName() string {
	return dpitl.Name
}

// Configuration returns the configuration (args) for [DiskPoolIscsiTargetLun].
func (dpitl *DiskPoolIscsiTargetLun) Configuration() interface{} {
	return dpitl.Args
}

// DependOn is used for other resources to depend on [DiskPoolIscsiTargetLun].
func (dpitl *DiskPoolIscsiTargetLun) DependOn() terra.Reference {
	return terra.ReferenceResource(dpitl)
}

// Dependencies returns the list of resources [DiskPoolIscsiTargetLun] depends_on.
func (dpitl *DiskPoolIscsiTargetLun) Dependencies() terra.Dependencies {
	return dpitl.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DiskPoolIscsiTargetLun].
func (dpitl *DiskPoolIscsiTargetLun) LifecycleManagement() *terra.Lifecycle {
	return dpitl.Lifecycle
}

// Attributes returns the attributes for [DiskPoolIscsiTargetLun].
func (dpitl *DiskPoolIscsiTargetLun) Attributes() diskPoolIscsiTargetLunAttributes {
	return diskPoolIscsiTargetLunAttributes{ref: terra.ReferenceResource(dpitl)}
}

// ImportState imports the given attribute values into [DiskPoolIscsiTargetLun]'s state.
func (dpitl *DiskPoolIscsiTargetLun) ImportState(av io.Reader) error {
	dpitl.state = &diskPoolIscsiTargetLunState{}
	if err := json.NewDecoder(av).Decode(dpitl.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dpitl.Type(), dpitl.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DiskPoolIscsiTargetLun] has state.
func (dpitl *DiskPoolIscsiTargetLun) State() (*diskPoolIscsiTargetLunState, bool) {
	return dpitl.state, dpitl.state != nil
}

// StateMust returns the state for [DiskPoolIscsiTargetLun]. Panics if the state is nil.
func (dpitl *DiskPoolIscsiTargetLun) StateMust() *diskPoolIscsiTargetLunState {
	if dpitl.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dpitl.Type(), dpitl.LocalName()))
	}
	return dpitl.state
}

// DiskPoolIscsiTargetLunArgs contains the configurations for azurerm_disk_pool_iscsi_target_lun.
type DiskPoolIscsiTargetLunArgs struct {
	// DiskPoolManagedDiskAttachmentId: string, required
	DiskPoolManagedDiskAttachmentId terra.StringValue `hcl:"disk_pool_managed_disk_attachment_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IscsiTargetId: string, required
	IscsiTargetId terra.StringValue `hcl:"iscsi_target_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *diskpooliscsitargetlun.Timeouts `hcl:"timeouts,block"`
}
type diskPoolIscsiTargetLunAttributes struct {
	ref terra.Reference
}

// DiskPoolManagedDiskAttachmentId returns a reference to field disk_pool_managed_disk_attachment_id of azurerm_disk_pool_iscsi_target_lun.
func (dpitl diskPoolIscsiTargetLunAttributes) DiskPoolManagedDiskAttachmentId() terra.StringValue {
	return terra.ReferenceAsString(dpitl.ref.Append("disk_pool_managed_disk_attachment_id"))
}

// Id returns a reference to field id of azurerm_disk_pool_iscsi_target_lun.
func (dpitl diskPoolIscsiTargetLunAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dpitl.ref.Append("id"))
}

// IscsiTargetId returns a reference to field iscsi_target_id of azurerm_disk_pool_iscsi_target_lun.
func (dpitl diskPoolIscsiTargetLunAttributes) IscsiTargetId() terra.StringValue {
	return terra.ReferenceAsString(dpitl.ref.Append("iscsi_target_id"))
}

// Lun returns a reference to field lun of azurerm_disk_pool_iscsi_target_lun.
func (dpitl diskPoolIscsiTargetLunAttributes) Lun() terra.NumberValue {
	return terra.ReferenceAsNumber(dpitl.ref.Append("lun"))
}

// Name returns a reference to field name of azurerm_disk_pool_iscsi_target_lun.
func (dpitl diskPoolIscsiTargetLunAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dpitl.ref.Append("name"))
}

func (dpitl diskPoolIscsiTargetLunAttributes) Timeouts() diskpooliscsitargetlun.TimeoutsAttributes {
	return terra.ReferenceAsSingle[diskpooliscsitargetlun.TimeoutsAttributes](dpitl.ref.Append("timeouts"))
}

type diskPoolIscsiTargetLunState struct {
	DiskPoolManagedDiskAttachmentId string                                `json:"disk_pool_managed_disk_attachment_id"`
	Id                              string                                `json:"id"`
	IscsiTargetId                   string                                `json:"iscsi_target_id"`
	Lun                             float64                               `json:"lun"`
	Name                            string                                `json:"name"`
	Timeouts                        *diskpooliscsitargetlun.TimeoutsState `json:"timeouts"`
}
