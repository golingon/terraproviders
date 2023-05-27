// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	databackuppolicyvm "github.com/golingon/terraproviders/azurerm/3.58.0/databackuppolicyvm"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataBackupPolicyVm creates a new instance of [DataBackupPolicyVm].
func NewDataBackupPolicyVm(name string, args DataBackupPolicyVmArgs) *DataBackupPolicyVm {
	return &DataBackupPolicyVm{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataBackupPolicyVm)(nil)

// DataBackupPolicyVm represents the Terraform data resource azurerm_backup_policy_vm.
type DataBackupPolicyVm struct {
	Name string
	Args DataBackupPolicyVmArgs
}

// DataSource returns the Terraform object type for [DataBackupPolicyVm].
func (bpv *DataBackupPolicyVm) DataSource() string {
	return "azurerm_backup_policy_vm"
}

// LocalName returns the local name for [DataBackupPolicyVm].
func (bpv *DataBackupPolicyVm) LocalName() string {
	return bpv.Name
}

// Configuration returns the configuration (args) for [DataBackupPolicyVm].
func (bpv *DataBackupPolicyVm) Configuration() interface{} {
	return bpv.Args
}

// Attributes returns the attributes for [DataBackupPolicyVm].
func (bpv *DataBackupPolicyVm) Attributes() dataBackupPolicyVmAttributes {
	return dataBackupPolicyVmAttributes{ref: terra.ReferenceDataResource(bpv)}
}

// DataBackupPolicyVmArgs contains the configurations for azurerm_backup_policy_vm.
type DataBackupPolicyVmArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RecoveryVaultName: string, required
	RecoveryVaultName terra.StringValue `hcl:"recovery_vault_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *databackuppolicyvm.Timeouts `hcl:"timeouts,block"`
}
type dataBackupPolicyVmAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_backup_policy_vm.
func (bpv dataBackupPolicyVmAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bpv.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_backup_policy_vm.
func (bpv dataBackupPolicyVmAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bpv.ref.Append("name"))
}

// RecoveryVaultName returns a reference to field recovery_vault_name of azurerm_backup_policy_vm.
func (bpv dataBackupPolicyVmAttributes) RecoveryVaultName() terra.StringValue {
	return terra.ReferenceAsString(bpv.ref.Append("recovery_vault_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_backup_policy_vm.
func (bpv dataBackupPolicyVmAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(bpv.ref.Append("resource_group_name"))
}

func (bpv dataBackupPolicyVmAttributes) Timeouts() databackuppolicyvm.TimeoutsAttributes {
	return terra.ReferenceAsSingle[databackuppolicyvm.TimeoutsAttributes](bpv.ref.Append("timeouts"))
}
