// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	dataprotectionbackupinstanceblobstorage "github.com/golingon/terraproviders/azurerm/3.69.0/dataprotectionbackupinstanceblobstorage"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataProtectionBackupInstanceBlobStorage creates a new instance of [DataProtectionBackupInstanceBlobStorage].
func NewDataProtectionBackupInstanceBlobStorage(name string, args DataProtectionBackupInstanceBlobStorageArgs) *DataProtectionBackupInstanceBlobStorage {
	return &DataProtectionBackupInstanceBlobStorage{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataProtectionBackupInstanceBlobStorage)(nil)

// DataProtectionBackupInstanceBlobStorage represents the Terraform resource azurerm_data_protection_backup_instance_blob_storage.
type DataProtectionBackupInstanceBlobStorage struct {
	Name      string
	Args      DataProtectionBackupInstanceBlobStorageArgs
	state     *dataProtectionBackupInstanceBlobStorageState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataProtectionBackupInstanceBlobStorage].
func (dpbibs *DataProtectionBackupInstanceBlobStorage) Type() string {
	return "azurerm_data_protection_backup_instance_blob_storage"
}

// LocalName returns the local name for [DataProtectionBackupInstanceBlobStorage].
func (dpbibs *DataProtectionBackupInstanceBlobStorage) LocalName() string {
	return dpbibs.Name
}

// Configuration returns the configuration (args) for [DataProtectionBackupInstanceBlobStorage].
func (dpbibs *DataProtectionBackupInstanceBlobStorage) Configuration() interface{} {
	return dpbibs.Args
}

// DependOn is used for other resources to depend on [DataProtectionBackupInstanceBlobStorage].
func (dpbibs *DataProtectionBackupInstanceBlobStorage) DependOn() terra.Reference {
	return terra.ReferenceResource(dpbibs)
}

// Dependencies returns the list of resources [DataProtectionBackupInstanceBlobStorage] depends_on.
func (dpbibs *DataProtectionBackupInstanceBlobStorage) Dependencies() terra.Dependencies {
	return dpbibs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataProtectionBackupInstanceBlobStorage].
func (dpbibs *DataProtectionBackupInstanceBlobStorage) LifecycleManagement() *terra.Lifecycle {
	return dpbibs.Lifecycle
}

// Attributes returns the attributes for [DataProtectionBackupInstanceBlobStorage].
func (dpbibs *DataProtectionBackupInstanceBlobStorage) Attributes() dataProtectionBackupInstanceBlobStorageAttributes {
	return dataProtectionBackupInstanceBlobStorageAttributes{ref: terra.ReferenceResource(dpbibs)}
}

// ImportState imports the given attribute values into [DataProtectionBackupInstanceBlobStorage]'s state.
func (dpbibs *DataProtectionBackupInstanceBlobStorage) ImportState(av io.Reader) error {
	dpbibs.state = &dataProtectionBackupInstanceBlobStorageState{}
	if err := json.NewDecoder(av).Decode(dpbibs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dpbibs.Type(), dpbibs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataProtectionBackupInstanceBlobStorage] has state.
func (dpbibs *DataProtectionBackupInstanceBlobStorage) State() (*dataProtectionBackupInstanceBlobStorageState, bool) {
	return dpbibs.state, dpbibs.state != nil
}

// StateMust returns the state for [DataProtectionBackupInstanceBlobStorage]. Panics if the state is nil.
func (dpbibs *DataProtectionBackupInstanceBlobStorage) StateMust() *dataProtectionBackupInstanceBlobStorageState {
	if dpbibs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dpbibs.Type(), dpbibs.LocalName()))
	}
	return dpbibs.state
}

// DataProtectionBackupInstanceBlobStorageArgs contains the configurations for azurerm_data_protection_backup_instance_blob_storage.
type DataProtectionBackupInstanceBlobStorageArgs struct {
	// BackupPolicyId: string, required
	BackupPolicyId terra.StringValue `hcl:"backup_policy_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// StorageAccountId: string, required
	StorageAccountId terra.StringValue `hcl:"storage_account_id,attr" validate:"required"`
	// VaultId: string, required
	VaultId terra.StringValue `hcl:"vault_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dataprotectionbackupinstanceblobstorage.Timeouts `hcl:"timeouts,block"`
}
type dataProtectionBackupInstanceBlobStorageAttributes struct {
	ref terra.Reference
}

// BackupPolicyId returns a reference to field backup_policy_id of azurerm_data_protection_backup_instance_blob_storage.
func (dpbibs dataProtectionBackupInstanceBlobStorageAttributes) BackupPolicyId() terra.StringValue {
	return terra.ReferenceAsString(dpbibs.ref.Append("backup_policy_id"))
}

// Id returns a reference to field id of azurerm_data_protection_backup_instance_blob_storage.
func (dpbibs dataProtectionBackupInstanceBlobStorageAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dpbibs.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_data_protection_backup_instance_blob_storage.
func (dpbibs dataProtectionBackupInstanceBlobStorageAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dpbibs.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_data_protection_backup_instance_blob_storage.
func (dpbibs dataProtectionBackupInstanceBlobStorageAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dpbibs.ref.Append("name"))
}

// StorageAccountId returns a reference to field storage_account_id of azurerm_data_protection_backup_instance_blob_storage.
func (dpbibs dataProtectionBackupInstanceBlobStorageAttributes) StorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(dpbibs.ref.Append("storage_account_id"))
}

// VaultId returns a reference to field vault_id of azurerm_data_protection_backup_instance_blob_storage.
func (dpbibs dataProtectionBackupInstanceBlobStorageAttributes) VaultId() terra.StringValue {
	return terra.ReferenceAsString(dpbibs.ref.Append("vault_id"))
}

func (dpbibs dataProtectionBackupInstanceBlobStorageAttributes) Timeouts() dataprotectionbackupinstanceblobstorage.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataprotectionbackupinstanceblobstorage.TimeoutsAttributes](dpbibs.ref.Append("timeouts"))
}

type dataProtectionBackupInstanceBlobStorageState struct {
	BackupPolicyId   string                                                 `json:"backup_policy_id"`
	Id               string                                                 `json:"id"`
	Location         string                                                 `json:"location"`
	Name             string                                                 `json:"name"`
	StorageAccountId string                                                 `json:"storage_account_id"`
	VaultId          string                                                 `json:"vault_id"`
	Timeouts         *dataprotectionbackupinstanceblobstorage.TimeoutsState `json:"timeouts"`
}
