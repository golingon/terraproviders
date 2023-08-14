// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	managedlustrefilesystem "github.com/golingon/terraproviders/azurerm/3.69.0/managedlustrefilesystem"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewManagedLustreFileSystem creates a new instance of [ManagedLustreFileSystem].
func NewManagedLustreFileSystem(name string, args ManagedLustreFileSystemArgs) *ManagedLustreFileSystem {
	return &ManagedLustreFileSystem{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ManagedLustreFileSystem)(nil)

// ManagedLustreFileSystem represents the Terraform resource azurerm_managed_lustre_file_system.
type ManagedLustreFileSystem struct {
	Name      string
	Args      ManagedLustreFileSystemArgs
	state     *managedLustreFileSystemState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ManagedLustreFileSystem].
func (mlfs *ManagedLustreFileSystem) Type() string {
	return "azurerm_managed_lustre_file_system"
}

// LocalName returns the local name for [ManagedLustreFileSystem].
func (mlfs *ManagedLustreFileSystem) LocalName() string {
	return mlfs.Name
}

// Configuration returns the configuration (args) for [ManagedLustreFileSystem].
func (mlfs *ManagedLustreFileSystem) Configuration() interface{} {
	return mlfs.Args
}

// DependOn is used for other resources to depend on [ManagedLustreFileSystem].
func (mlfs *ManagedLustreFileSystem) DependOn() terra.Reference {
	return terra.ReferenceResource(mlfs)
}

// Dependencies returns the list of resources [ManagedLustreFileSystem] depends_on.
func (mlfs *ManagedLustreFileSystem) Dependencies() terra.Dependencies {
	return mlfs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ManagedLustreFileSystem].
func (mlfs *ManagedLustreFileSystem) LifecycleManagement() *terra.Lifecycle {
	return mlfs.Lifecycle
}

// Attributes returns the attributes for [ManagedLustreFileSystem].
func (mlfs *ManagedLustreFileSystem) Attributes() managedLustreFileSystemAttributes {
	return managedLustreFileSystemAttributes{ref: terra.ReferenceResource(mlfs)}
}

// ImportState imports the given attribute values into [ManagedLustreFileSystem]'s state.
func (mlfs *ManagedLustreFileSystem) ImportState(av io.Reader) error {
	mlfs.state = &managedLustreFileSystemState{}
	if err := json.NewDecoder(av).Decode(mlfs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mlfs.Type(), mlfs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ManagedLustreFileSystem] has state.
func (mlfs *ManagedLustreFileSystem) State() (*managedLustreFileSystemState, bool) {
	return mlfs.state, mlfs.state != nil
}

// StateMust returns the state for [ManagedLustreFileSystem]. Panics if the state is nil.
func (mlfs *ManagedLustreFileSystem) StateMust() *managedLustreFileSystemState {
	if mlfs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mlfs.Type(), mlfs.LocalName()))
	}
	return mlfs.state
}

// ManagedLustreFileSystemArgs contains the configurations for azurerm_managed_lustre_file_system.
type ManagedLustreFileSystemArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SkuName: string, required
	SkuName terra.StringValue `hcl:"sku_name,attr" validate:"required"`
	// StorageCapacityInTb: number, required
	StorageCapacityInTb terra.NumberValue `hcl:"storage_capacity_in_tb,attr" validate:"required"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Zones: set of string, required
	Zones terra.SetValue[terra.StringValue] `hcl:"zones,attr" validate:"required"`
	// EncryptionKey: optional
	EncryptionKey *managedlustrefilesystem.EncryptionKey `hcl:"encryption_key,block"`
	// HsmSetting: optional
	HsmSetting *managedlustrefilesystem.HsmSetting `hcl:"hsm_setting,block"`
	// Identity: optional
	Identity *managedlustrefilesystem.Identity `hcl:"identity,block"`
	// MaintenanceWindow: required
	MaintenanceWindow *managedlustrefilesystem.MaintenanceWindow `hcl:"maintenance_window,block" validate:"required"`
	// Timeouts: optional
	Timeouts *managedlustrefilesystem.Timeouts `hcl:"timeouts,block"`
}
type managedLustreFileSystemAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_managed_lustre_file_system.
func (mlfs managedLustreFileSystemAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mlfs.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_managed_lustre_file_system.
func (mlfs managedLustreFileSystemAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(mlfs.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_managed_lustre_file_system.
func (mlfs managedLustreFileSystemAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mlfs.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_managed_lustre_file_system.
func (mlfs managedLustreFileSystemAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(mlfs.ref.Append("resource_group_name"))
}

// SkuName returns a reference to field sku_name of azurerm_managed_lustre_file_system.
func (mlfs managedLustreFileSystemAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(mlfs.ref.Append("sku_name"))
}

// StorageCapacityInTb returns a reference to field storage_capacity_in_tb of azurerm_managed_lustre_file_system.
func (mlfs managedLustreFileSystemAttributes) StorageCapacityInTb() terra.NumberValue {
	return terra.ReferenceAsNumber(mlfs.ref.Append("storage_capacity_in_tb"))
}

// SubnetId returns a reference to field subnet_id of azurerm_managed_lustre_file_system.
func (mlfs managedLustreFileSystemAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(mlfs.ref.Append("subnet_id"))
}

// Tags returns a reference to field tags of azurerm_managed_lustre_file_system.
func (mlfs managedLustreFileSystemAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mlfs.ref.Append("tags"))
}

// Zones returns a reference to field zones of azurerm_managed_lustre_file_system.
func (mlfs managedLustreFileSystemAttributes) Zones() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](mlfs.ref.Append("zones"))
}

func (mlfs managedLustreFileSystemAttributes) EncryptionKey() terra.ListValue[managedlustrefilesystem.EncryptionKeyAttributes] {
	return terra.ReferenceAsList[managedlustrefilesystem.EncryptionKeyAttributes](mlfs.ref.Append("encryption_key"))
}

func (mlfs managedLustreFileSystemAttributes) HsmSetting() terra.ListValue[managedlustrefilesystem.HsmSettingAttributes] {
	return terra.ReferenceAsList[managedlustrefilesystem.HsmSettingAttributes](mlfs.ref.Append("hsm_setting"))
}

func (mlfs managedLustreFileSystemAttributes) Identity() terra.ListValue[managedlustrefilesystem.IdentityAttributes] {
	return terra.ReferenceAsList[managedlustrefilesystem.IdentityAttributes](mlfs.ref.Append("identity"))
}

func (mlfs managedLustreFileSystemAttributes) MaintenanceWindow() terra.ListValue[managedlustrefilesystem.MaintenanceWindowAttributes] {
	return terra.ReferenceAsList[managedlustrefilesystem.MaintenanceWindowAttributes](mlfs.ref.Append("maintenance_window"))
}

func (mlfs managedLustreFileSystemAttributes) Timeouts() managedlustrefilesystem.TimeoutsAttributes {
	return terra.ReferenceAsSingle[managedlustrefilesystem.TimeoutsAttributes](mlfs.ref.Append("timeouts"))
}

type managedLustreFileSystemState struct {
	Id                  string                                           `json:"id"`
	Location            string                                           `json:"location"`
	Name                string                                           `json:"name"`
	ResourceGroupName   string                                           `json:"resource_group_name"`
	SkuName             string                                           `json:"sku_name"`
	StorageCapacityInTb float64                                          `json:"storage_capacity_in_tb"`
	SubnetId            string                                           `json:"subnet_id"`
	Tags                map[string]string                                `json:"tags"`
	Zones               []string                                         `json:"zones"`
	EncryptionKey       []managedlustrefilesystem.EncryptionKeyState     `json:"encryption_key"`
	HsmSetting          []managedlustrefilesystem.HsmSettingState        `json:"hsm_setting"`
	Identity            []managedlustrefilesystem.IdentityState          `json:"identity"`
	MaintenanceWindow   []managedlustrefilesystem.MaintenanceWindowState `json:"maintenance_window"`
	Timeouts            *managedlustrefilesystem.TimeoutsState           `json:"timeouts"`
}
