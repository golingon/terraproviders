// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	dataprotectionbackupinstancedisk "github.com/golingon/terraproviders/azurerm/3.67.0/dataprotectionbackupinstancedisk"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataProtectionBackupInstanceDisk creates a new instance of [DataProtectionBackupInstanceDisk].
func NewDataProtectionBackupInstanceDisk(name string, args DataProtectionBackupInstanceDiskArgs) *DataProtectionBackupInstanceDisk {
	return &DataProtectionBackupInstanceDisk{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataProtectionBackupInstanceDisk)(nil)

// DataProtectionBackupInstanceDisk represents the Terraform resource azurerm_data_protection_backup_instance_disk.
type DataProtectionBackupInstanceDisk struct {
	Name      string
	Args      DataProtectionBackupInstanceDiskArgs
	state     *dataProtectionBackupInstanceDiskState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataProtectionBackupInstanceDisk].
func (dpbid *DataProtectionBackupInstanceDisk) Type() string {
	return "azurerm_data_protection_backup_instance_disk"
}

// LocalName returns the local name for [DataProtectionBackupInstanceDisk].
func (dpbid *DataProtectionBackupInstanceDisk) LocalName() string {
	return dpbid.Name
}

// Configuration returns the configuration (args) for [DataProtectionBackupInstanceDisk].
func (dpbid *DataProtectionBackupInstanceDisk) Configuration() interface{} {
	return dpbid.Args
}

// DependOn is used for other resources to depend on [DataProtectionBackupInstanceDisk].
func (dpbid *DataProtectionBackupInstanceDisk) DependOn() terra.Reference {
	return terra.ReferenceResource(dpbid)
}

// Dependencies returns the list of resources [DataProtectionBackupInstanceDisk] depends_on.
func (dpbid *DataProtectionBackupInstanceDisk) Dependencies() terra.Dependencies {
	return dpbid.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataProtectionBackupInstanceDisk].
func (dpbid *DataProtectionBackupInstanceDisk) LifecycleManagement() *terra.Lifecycle {
	return dpbid.Lifecycle
}

// Attributes returns the attributes for [DataProtectionBackupInstanceDisk].
func (dpbid *DataProtectionBackupInstanceDisk) Attributes() dataProtectionBackupInstanceDiskAttributes {
	return dataProtectionBackupInstanceDiskAttributes{ref: terra.ReferenceResource(dpbid)}
}

// ImportState imports the given attribute values into [DataProtectionBackupInstanceDisk]'s state.
func (dpbid *DataProtectionBackupInstanceDisk) ImportState(av io.Reader) error {
	dpbid.state = &dataProtectionBackupInstanceDiskState{}
	if err := json.NewDecoder(av).Decode(dpbid.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dpbid.Type(), dpbid.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataProtectionBackupInstanceDisk] has state.
func (dpbid *DataProtectionBackupInstanceDisk) State() (*dataProtectionBackupInstanceDiskState, bool) {
	return dpbid.state, dpbid.state != nil
}

// StateMust returns the state for [DataProtectionBackupInstanceDisk]. Panics if the state is nil.
func (dpbid *DataProtectionBackupInstanceDisk) StateMust() *dataProtectionBackupInstanceDiskState {
	if dpbid.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dpbid.Type(), dpbid.LocalName()))
	}
	return dpbid.state
}

// DataProtectionBackupInstanceDiskArgs contains the configurations for azurerm_data_protection_backup_instance_disk.
type DataProtectionBackupInstanceDiskArgs struct {
	// BackupPolicyId: string, required
	BackupPolicyId terra.StringValue `hcl:"backup_policy_id,attr" validate:"required"`
	// DiskId: string, required
	DiskId terra.StringValue `hcl:"disk_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SnapshotResourceGroupName: string, required
	SnapshotResourceGroupName terra.StringValue `hcl:"snapshot_resource_group_name,attr" validate:"required"`
	// VaultId: string, required
	VaultId terra.StringValue `hcl:"vault_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dataprotectionbackupinstancedisk.Timeouts `hcl:"timeouts,block"`
}
type dataProtectionBackupInstanceDiskAttributes struct {
	ref terra.Reference
}

// BackupPolicyId returns a reference to field backup_policy_id of azurerm_data_protection_backup_instance_disk.
func (dpbid dataProtectionBackupInstanceDiskAttributes) BackupPolicyId() terra.StringValue {
	return terra.ReferenceAsString(dpbid.ref.Append("backup_policy_id"))
}

// DiskId returns a reference to field disk_id of azurerm_data_protection_backup_instance_disk.
func (dpbid dataProtectionBackupInstanceDiskAttributes) DiskId() terra.StringValue {
	return terra.ReferenceAsString(dpbid.ref.Append("disk_id"))
}

// Id returns a reference to field id of azurerm_data_protection_backup_instance_disk.
func (dpbid dataProtectionBackupInstanceDiskAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dpbid.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_data_protection_backup_instance_disk.
func (dpbid dataProtectionBackupInstanceDiskAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dpbid.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_data_protection_backup_instance_disk.
func (dpbid dataProtectionBackupInstanceDiskAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dpbid.ref.Append("name"))
}

// SnapshotResourceGroupName returns a reference to field snapshot_resource_group_name of azurerm_data_protection_backup_instance_disk.
func (dpbid dataProtectionBackupInstanceDiskAttributes) SnapshotResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dpbid.ref.Append("snapshot_resource_group_name"))
}

// VaultId returns a reference to field vault_id of azurerm_data_protection_backup_instance_disk.
func (dpbid dataProtectionBackupInstanceDiskAttributes) VaultId() terra.StringValue {
	return terra.ReferenceAsString(dpbid.ref.Append("vault_id"))
}

func (dpbid dataProtectionBackupInstanceDiskAttributes) Timeouts() dataprotectionbackupinstancedisk.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataprotectionbackupinstancedisk.TimeoutsAttributes](dpbid.ref.Append("timeouts"))
}

type dataProtectionBackupInstanceDiskState struct {
	BackupPolicyId            string                                          `json:"backup_policy_id"`
	DiskId                    string                                          `json:"disk_id"`
	Id                        string                                          `json:"id"`
	Location                  string                                          `json:"location"`
	Name                      string                                          `json:"name"`
	SnapshotResourceGroupName string                                          `json:"snapshot_resource_group_name"`
	VaultId                   string                                          `json:"vault_id"`
	Timeouts                  *dataprotectionbackupinstancedisk.TimeoutsState `json:"timeouts"`
}
