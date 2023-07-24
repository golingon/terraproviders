// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datadataprotectionbackupvault "github.com/golingon/terraproviders/azurerm/3.66.0/datadataprotectionbackupvault"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataDataProtectionBackupVault creates a new instance of [DataDataProtectionBackupVault].
func NewDataDataProtectionBackupVault(name string, args DataDataProtectionBackupVaultArgs) *DataDataProtectionBackupVault {
	return &DataDataProtectionBackupVault{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDataProtectionBackupVault)(nil)

// DataDataProtectionBackupVault represents the Terraform data resource azurerm_data_protection_backup_vault.
type DataDataProtectionBackupVault struct {
	Name string
	Args DataDataProtectionBackupVaultArgs
}

// DataSource returns the Terraform object type for [DataDataProtectionBackupVault].
func (dpbv *DataDataProtectionBackupVault) DataSource() string {
	return "azurerm_data_protection_backup_vault"
}

// LocalName returns the local name for [DataDataProtectionBackupVault].
func (dpbv *DataDataProtectionBackupVault) LocalName() string {
	return dpbv.Name
}

// Configuration returns the configuration (args) for [DataDataProtectionBackupVault].
func (dpbv *DataDataProtectionBackupVault) Configuration() interface{} {
	return dpbv.Args
}

// Attributes returns the attributes for [DataDataProtectionBackupVault].
func (dpbv *DataDataProtectionBackupVault) Attributes() dataDataProtectionBackupVaultAttributes {
	return dataDataProtectionBackupVaultAttributes{ref: terra.ReferenceDataResource(dpbv)}
}

// DataDataProtectionBackupVaultArgs contains the configurations for azurerm_data_protection_backup_vault.
type DataDataProtectionBackupVaultArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Identity: min=0
	Identity []datadataprotectionbackupvault.Identity `hcl:"identity,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datadataprotectionbackupvault.Timeouts `hcl:"timeouts,block"`
}
type dataDataProtectionBackupVaultAttributes struct {
	ref terra.Reference
}

// DatastoreType returns a reference to field datastore_type of azurerm_data_protection_backup_vault.
func (dpbv dataDataProtectionBackupVaultAttributes) DatastoreType() terra.StringValue {
	return terra.ReferenceAsString(dpbv.ref.Append("datastore_type"))
}

// Id returns a reference to field id of azurerm_data_protection_backup_vault.
func (dpbv dataDataProtectionBackupVaultAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dpbv.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_data_protection_backup_vault.
func (dpbv dataDataProtectionBackupVaultAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dpbv.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_data_protection_backup_vault.
func (dpbv dataDataProtectionBackupVaultAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dpbv.ref.Append("name"))
}

// Redundancy returns a reference to field redundancy of azurerm_data_protection_backup_vault.
func (dpbv dataDataProtectionBackupVaultAttributes) Redundancy() terra.StringValue {
	return terra.ReferenceAsString(dpbv.ref.Append("redundancy"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_data_protection_backup_vault.
func (dpbv dataDataProtectionBackupVaultAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dpbv.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_data_protection_backup_vault.
func (dpbv dataDataProtectionBackupVaultAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dpbv.ref.Append("tags"))
}

func (dpbv dataDataProtectionBackupVaultAttributes) Identity() terra.ListValue[datadataprotectionbackupvault.IdentityAttributes] {
	return terra.ReferenceAsList[datadataprotectionbackupvault.IdentityAttributes](dpbv.ref.Append("identity"))
}

func (dpbv dataDataProtectionBackupVaultAttributes) Timeouts() datadataprotectionbackupvault.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datadataprotectionbackupvault.TimeoutsAttributes](dpbv.ref.Append("timeouts"))
}
