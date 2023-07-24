// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	dataprotectionbackuppolicyblobstorage "github.com/golingon/terraproviders/azurerm/3.66.0/dataprotectionbackuppolicyblobstorage"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataProtectionBackupPolicyBlobStorage creates a new instance of [DataProtectionBackupPolicyBlobStorage].
func NewDataProtectionBackupPolicyBlobStorage(name string, args DataProtectionBackupPolicyBlobStorageArgs) *DataProtectionBackupPolicyBlobStorage {
	return &DataProtectionBackupPolicyBlobStorage{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataProtectionBackupPolicyBlobStorage)(nil)

// DataProtectionBackupPolicyBlobStorage represents the Terraform resource azurerm_data_protection_backup_policy_blob_storage.
type DataProtectionBackupPolicyBlobStorage struct {
	Name      string
	Args      DataProtectionBackupPolicyBlobStorageArgs
	state     *dataProtectionBackupPolicyBlobStorageState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataProtectionBackupPolicyBlobStorage].
func (dpbpbs *DataProtectionBackupPolicyBlobStorage) Type() string {
	return "azurerm_data_protection_backup_policy_blob_storage"
}

// LocalName returns the local name for [DataProtectionBackupPolicyBlobStorage].
func (dpbpbs *DataProtectionBackupPolicyBlobStorage) LocalName() string {
	return dpbpbs.Name
}

// Configuration returns the configuration (args) for [DataProtectionBackupPolicyBlobStorage].
func (dpbpbs *DataProtectionBackupPolicyBlobStorage) Configuration() interface{} {
	return dpbpbs.Args
}

// DependOn is used for other resources to depend on [DataProtectionBackupPolicyBlobStorage].
func (dpbpbs *DataProtectionBackupPolicyBlobStorage) DependOn() terra.Reference {
	return terra.ReferenceResource(dpbpbs)
}

// Dependencies returns the list of resources [DataProtectionBackupPolicyBlobStorage] depends_on.
func (dpbpbs *DataProtectionBackupPolicyBlobStorage) Dependencies() terra.Dependencies {
	return dpbpbs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataProtectionBackupPolicyBlobStorage].
func (dpbpbs *DataProtectionBackupPolicyBlobStorage) LifecycleManagement() *terra.Lifecycle {
	return dpbpbs.Lifecycle
}

// Attributes returns the attributes for [DataProtectionBackupPolicyBlobStorage].
func (dpbpbs *DataProtectionBackupPolicyBlobStorage) Attributes() dataProtectionBackupPolicyBlobStorageAttributes {
	return dataProtectionBackupPolicyBlobStorageAttributes{ref: terra.ReferenceResource(dpbpbs)}
}

// ImportState imports the given attribute values into [DataProtectionBackupPolicyBlobStorage]'s state.
func (dpbpbs *DataProtectionBackupPolicyBlobStorage) ImportState(av io.Reader) error {
	dpbpbs.state = &dataProtectionBackupPolicyBlobStorageState{}
	if err := json.NewDecoder(av).Decode(dpbpbs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dpbpbs.Type(), dpbpbs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataProtectionBackupPolicyBlobStorage] has state.
func (dpbpbs *DataProtectionBackupPolicyBlobStorage) State() (*dataProtectionBackupPolicyBlobStorageState, bool) {
	return dpbpbs.state, dpbpbs.state != nil
}

// StateMust returns the state for [DataProtectionBackupPolicyBlobStorage]. Panics if the state is nil.
func (dpbpbs *DataProtectionBackupPolicyBlobStorage) StateMust() *dataProtectionBackupPolicyBlobStorageState {
	if dpbpbs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dpbpbs.Type(), dpbpbs.LocalName()))
	}
	return dpbpbs.state
}

// DataProtectionBackupPolicyBlobStorageArgs contains the configurations for azurerm_data_protection_backup_policy_blob_storage.
type DataProtectionBackupPolicyBlobStorageArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RetentionDuration: string, required
	RetentionDuration terra.StringValue `hcl:"retention_duration,attr" validate:"required"`
	// VaultId: string, required
	VaultId terra.StringValue `hcl:"vault_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dataprotectionbackuppolicyblobstorage.Timeouts `hcl:"timeouts,block"`
}
type dataProtectionBackupPolicyBlobStorageAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_data_protection_backup_policy_blob_storage.
func (dpbpbs dataProtectionBackupPolicyBlobStorageAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dpbpbs.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_data_protection_backup_policy_blob_storage.
func (dpbpbs dataProtectionBackupPolicyBlobStorageAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dpbpbs.ref.Append("name"))
}

// RetentionDuration returns a reference to field retention_duration of azurerm_data_protection_backup_policy_blob_storage.
func (dpbpbs dataProtectionBackupPolicyBlobStorageAttributes) RetentionDuration() terra.StringValue {
	return terra.ReferenceAsString(dpbpbs.ref.Append("retention_duration"))
}

// VaultId returns a reference to field vault_id of azurerm_data_protection_backup_policy_blob_storage.
func (dpbpbs dataProtectionBackupPolicyBlobStorageAttributes) VaultId() terra.StringValue {
	return terra.ReferenceAsString(dpbpbs.ref.Append("vault_id"))
}

func (dpbpbs dataProtectionBackupPolicyBlobStorageAttributes) Timeouts() dataprotectionbackuppolicyblobstorage.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataprotectionbackuppolicyblobstorage.TimeoutsAttributes](dpbpbs.ref.Append("timeouts"))
}

type dataProtectionBackupPolicyBlobStorageState struct {
	Id                string                                               `json:"id"`
	Name              string                                               `json:"name"`
	RetentionDuration string                                               `json:"retention_duration"`
	VaultId           string                                               `json:"vault_id"`
	Timeouts          *dataprotectionbackuppolicyblobstorage.TimeoutsState `json:"timeouts"`
}
