// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	databackuppolicyfileshare "github.com/golingon/terraproviders/azurerm/3.65.0/databackuppolicyfileshare"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataBackupPolicyFileShare creates a new instance of [DataBackupPolicyFileShare].
func NewDataBackupPolicyFileShare(name string, args DataBackupPolicyFileShareArgs) *DataBackupPolicyFileShare {
	return &DataBackupPolicyFileShare{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataBackupPolicyFileShare)(nil)

// DataBackupPolicyFileShare represents the Terraform data resource azurerm_backup_policy_file_share.
type DataBackupPolicyFileShare struct {
	Name string
	Args DataBackupPolicyFileShareArgs
}

// DataSource returns the Terraform object type for [DataBackupPolicyFileShare].
func (bpfs *DataBackupPolicyFileShare) DataSource() string {
	return "azurerm_backup_policy_file_share"
}

// LocalName returns the local name for [DataBackupPolicyFileShare].
func (bpfs *DataBackupPolicyFileShare) LocalName() string {
	return bpfs.Name
}

// Configuration returns the configuration (args) for [DataBackupPolicyFileShare].
func (bpfs *DataBackupPolicyFileShare) Configuration() interface{} {
	return bpfs.Args
}

// Attributes returns the attributes for [DataBackupPolicyFileShare].
func (bpfs *DataBackupPolicyFileShare) Attributes() dataBackupPolicyFileShareAttributes {
	return dataBackupPolicyFileShareAttributes{ref: terra.ReferenceDataResource(bpfs)}
}

// DataBackupPolicyFileShareArgs contains the configurations for azurerm_backup_policy_file_share.
type DataBackupPolicyFileShareArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RecoveryVaultName: string, required
	RecoveryVaultName terra.StringValue `hcl:"recovery_vault_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *databackuppolicyfileshare.Timeouts `hcl:"timeouts,block"`
}
type dataBackupPolicyFileShareAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_backup_policy_file_share.
func (bpfs dataBackupPolicyFileShareAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bpfs.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_backup_policy_file_share.
func (bpfs dataBackupPolicyFileShareAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bpfs.ref.Append("name"))
}

// RecoveryVaultName returns a reference to field recovery_vault_name of azurerm_backup_policy_file_share.
func (bpfs dataBackupPolicyFileShareAttributes) RecoveryVaultName() terra.StringValue {
	return terra.ReferenceAsString(bpfs.ref.Append("recovery_vault_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_backup_policy_file_share.
func (bpfs dataBackupPolicyFileShareAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(bpfs.ref.Append("resource_group_name"))
}

func (bpfs dataBackupPolicyFileShareAttributes) Timeouts() databackuppolicyfileshare.TimeoutsAttributes {
	return terra.ReferenceAsSingle[databackuppolicyfileshare.TimeoutsAttributes](bpfs.ref.Append("timeouts"))
}
