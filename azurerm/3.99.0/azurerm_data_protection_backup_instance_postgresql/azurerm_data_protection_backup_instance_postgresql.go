// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_data_protection_backup_instance_postgresql

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

// Resource represents the Terraform resource azurerm_data_protection_backup_instance_postgresql.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermDataProtectionBackupInstancePostgresqlState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (adpbip *Resource) Type() string {
	return "azurerm_data_protection_backup_instance_postgresql"
}

// LocalName returns the local name for [Resource].
func (adpbip *Resource) LocalName() string {
	return adpbip.Name
}

// Configuration returns the configuration (args) for [Resource].
func (adpbip *Resource) Configuration() interface{} {
	return adpbip.Args
}

// DependOn is used for other resources to depend on [Resource].
func (adpbip *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(adpbip)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (adpbip *Resource) Dependencies() terra.Dependencies {
	return adpbip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (adpbip *Resource) LifecycleManagement() *terra.Lifecycle {
	return adpbip.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (adpbip *Resource) Attributes() azurermDataProtectionBackupInstancePostgresqlAttributes {
	return azurermDataProtectionBackupInstancePostgresqlAttributes{ref: terra.ReferenceResource(adpbip)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (adpbip *Resource) ImportState(state io.Reader) error {
	adpbip.state = &azurermDataProtectionBackupInstancePostgresqlState{}
	if err := json.NewDecoder(state).Decode(adpbip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", adpbip.Type(), adpbip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (adpbip *Resource) State() (*azurermDataProtectionBackupInstancePostgresqlState, bool) {
	return adpbip.state, adpbip.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (adpbip *Resource) StateMust() *azurermDataProtectionBackupInstancePostgresqlState {
	if adpbip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", adpbip.Type(), adpbip.LocalName()))
	}
	return adpbip.state
}

// Args contains the configurations for azurerm_data_protection_backup_instance_postgresql.
type Args struct {
	// BackupPolicyId: string, required
	BackupPolicyId terra.StringValue `hcl:"backup_policy_id,attr" validate:"required"`
	// DatabaseCredentialKeyVaultSecretId: string, optional
	DatabaseCredentialKeyVaultSecretId terra.StringValue `hcl:"database_credential_key_vault_secret_id,attr"`
	// DatabaseId: string, required
	DatabaseId terra.StringValue `hcl:"database_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// VaultId: string, required
	VaultId terra.StringValue `hcl:"vault_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermDataProtectionBackupInstancePostgresqlAttributes struct {
	ref terra.Reference
}

// BackupPolicyId returns a reference to field backup_policy_id of azurerm_data_protection_backup_instance_postgresql.
func (adpbip azurermDataProtectionBackupInstancePostgresqlAttributes) BackupPolicyId() terra.StringValue {
	return terra.ReferenceAsString(adpbip.ref.Append("backup_policy_id"))
}

// DatabaseCredentialKeyVaultSecretId returns a reference to field database_credential_key_vault_secret_id of azurerm_data_protection_backup_instance_postgresql.
func (adpbip azurermDataProtectionBackupInstancePostgresqlAttributes) DatabaseCredentialKeyVaultSecretId() terra.StringValue {
	return terra.ReferenceAsString(adpbip.ref.Append("database_credential_key_vault_secret_id"))
}

// DatabaseId returns a reference to field database_id of azurerm_data_protection_backup_instance_postgresql.
func (adpbip azurermDataProtectionBackupInstancePostgresqlAttributes) DatabaseId() terra.StringValue {
	return terra.ReferenceAsString(adpbip.ref.Append("database_id"))
}

// Id returns a reference to field id of azurerm_data_protection_backup_instance_postgresql.
func (adpbip azurermDataProtectionBackupInstancePostgresqlAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adpbip.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_data_protection_backup_instance_postgresql.
func (adpbip azurermDataProtectionBackupInstancePostgresqlAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(adpbip.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_data_protection_backup_instance_postgresql.
func (adpbip azurermDataProtectionBackupInstancePostgresqlAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(adpbip.ref.Append("name"))
}

// VaultId returns a reference to field vault_id of azurerm_data_protection_backup_instance_postgresql.
func (adpbip azurermDataProtectionBackupInstancePostgresqlAttributes) VaultId() terra.StringValue {
	return terra.ReferenceAsString(adpbip.ref.Append("vault_id"))
}

func (adpbip azurermDataProtectionBackupInstancePostgresqlAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](adpbip.ref.Append("timeouts"))
}

type azurermDataProtectionBackupInstancePostgresqlState struct {
	BackupPolicyId                     string         `json:"backup_policy_id"`
	DatabaseCredentialKeyVaultSecretId string         `json:"database_credential_key_vault_secret_id"`
	DatabaseId                         string         `json:"database_id"`
	Id                                 string         `json:"id"`
	Location                           string         `json:"location"`
	Name                               string         `json:"name"`
	VaultId                            string         `json:"vault_id"`
	Timeouts                           *TimeoutsState `json:"timeouts"`
}
