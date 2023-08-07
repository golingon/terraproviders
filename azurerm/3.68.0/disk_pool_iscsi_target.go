// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	diskpooliscsitarget "github.com/golingon/terraproviders/azurerm/3.68.0/diskpooliscsitarget"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDiskPoolIscsiTarget creates a new instance of [DiskPoolIscsiTarget].
func NewDiskPoolIscsiTarget(name string, args DiskPoolIscsiTargetArgs) *DiskPoolIscsiTarget {
	return &DiskPoolIscsiTarget{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DiskPoolIscsiTarget)(nil)

// DiskPoolIscsiTarget represents the Terraform resource azurerm_disk_pool_iscsi_target.
type DiskPoolIscsiTarget struct {
	Name      string
	Args      DiskPoolIscsiTargetArgs
	state     *diskPoolIscsiTargetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DiskPoolIscsiTarget].
func (dpit *DiskPoolIscsiTarget) Type() string {
	return "azurerm_disk_pool_iscsi_target"
}

// LocalName returns the local name for [DiskPoolIscsiTarget].
func (dpit *DiskPoolIscsiTarget) LocalName() string {
	return dpit.Name
}

// Configuration returns the configuration (args) for [DiskPoolIscsiTarget].
func (dpit *DiskPoolIscsiTarget) Configuration() interface{} {
	return dpit.Args
}

// DependOn is used for other resources to depend on [DiskPoolIscsiTarget].
func (dpit *DiskPoolIscsiTarget) DependOn() terra.Reference {
	return terra.ReferenceResource(dpit)
}

// Dependencies returns the list of resources [DiskPoolIscsiTarget] depends_on.
func (dpit *DiskPoolIscsiTarget) Dependencies() terra.Dependencies {
	return dpit.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DiskPoolIscsiTarget].
func (dpit *DiskPoolIscsiTarget) LifecycleManagement() *terra.Lifecycle {
	return dpit.Lifecycle
}

// Attributes returns the attributes for [DiskPoolIscsiTarget].
func (dpit *DiskPoolIscsiTarget) Attributes() diskPoolIscsiTargetAttributes {
	return diskPoolIscsiTargetAttributes{ref: terra.ReferenceResource(dpit)}
}

// ImportState imports the given attribute values into [DiskPoolIscsiTarget]'s state.
func (dpit *DiskPoolIscsiTarget) ImportState(av io.Reader) error {
	dpit.state = &diskPoolIscsiTargetState{}
	if err := json.NewDecoder(av).Decode(dpit.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dpit.Type(), dpit.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DiskPoolIscsiTarget] has state.
func (dpit *DiskPoolIscsiTarget) State() (*diskPoolIscsiTargetState, bool) {
	return dpit.state, dpit.state != nil
}

// StateMust returns the state for [DiskPoolIscsiTarget]. Panics if the state is nil.
func (dpit *DiskPoolIscsiTarget) StateMust() *diskPoolIscsiTargetState {
	if dpit.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dpit.Type(), dpit.LocalName()))
	}
	return dpit.state
}

// DiskPoolIscsiTargetArgs contains the configurations for azurerm_disk_pool_iscsi_target.
type DiskPoolIscsiTargetArgs struct {
	// AclMode: string, required
	AclMode terra.StringValue `hcl:"acl_mode,attr" validate:"required"`
	// DisksPoolId: string, required
	DisksPoolId terra.StringValue `hcl:"disks_pool_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// TargetIqn: string, optional
	TargetIqn terra.StringValue `hcl:"target_iqn,attr"`
	// Timeouts: optional
	Timeouts *diskpooliscsitarget.Timeouts `hcl:"timeouts,block"`
}
type diskPoolIscsiTargetAttributes struct {
	ref terra.Reference
}

// AclMode returns a reference to field acl_mode of azurerm_disk_pool_iscsi_target.
func (dpit diskPoolIscsiTargetAttributes) AclMode() terra.StringValue {
	return terra.ReferenceAsString(dpit.ref.Append("acl_mode"))
}

// DisksPoolId returns a reference to field disks_pool_id of azurerm_disk_pool_iscsi_target.
func (dpit diskPoolIscsiTargetAttributes) DisksPoolId() terra.StringValue {
	return terra.ReferenceAsString(dpit.ref.Append("disks_pool_id"))
}

// Endpoints returns a reference to field endpoints of azurerm_disk_pool_iscsi_target.
func (dpit diskPoolIscsiTargetAttributes) Endpoints() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dpit.ref.Append("endpoints"))
}

// Id returns a reference to field id of azurerm_disk_pool_iscsi_target.
func (dpit diskPoolIscsiTargetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dpit.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_disk_pool_iscsi_target.
func (dpit diskPoolIscsiTargetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dpit.ref.Append("name"))
}

// Port returns a reference to field port of azurerm_disk_pool_iscsi_target.
func (dpit diskPoolIscsiTargetAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(dpit.ref.Append("port"))
}

// TargetIqn returns a reference to field target_iqn of azurerm_disk_pool_iscsi_target.
func (dpit diskPoolIscsiTargetAttributes) TargetIqn() terra.StringValue {
	return terra.ReferenceAsString(dpit.ref.Append("target_iqn"))
}

func (dpit diskPoolIscsiTargetAttributes) Timeouts() diskpooliscsitarget.TimeoutsAttributes {
	return terra.ReferenceAsSingle[diskpooliscsitarget.TimeoutsAttributes](dpit.ref.Append("timeouts"))
}

type diskPoolIscsiTargetState struct {
	AclMode     string                             `json:"acl_mode"`
	DisksPoolId string                             `json:"disks_pool_id"`
	Endpoints   []string                           `json:"endpoints"`
	Id          string                             `json:"id"`
	Name        string                             `json:"name"`
	Port        float64                            `json:"port"`
	TargetIqn   string                             `json:"target_iqn"`
	Timeouts    *diskpooliscsitarget.TimeoutsState `json:"timeouts"`
}
