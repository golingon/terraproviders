// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	databackuppolicyvm "github.com/golingon/terraproviders/azurerm/3.49.0/databackuppolicyvm"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataBackupPolicyVm(name string, args DataBackupPolicyVmArgs) *DataBackupPolicyVm {
	return &DataBackupPolicyVm{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataBackupPolicyVm)(nil)

type DataBackupPolicyVm struct {
	Name string
	Args DataBackupPolicyVmArgs
}

func (bpv *DataBackupPolicyVm) DataSource() string {
	return "azurerm_backup_policy_vm"
}

func (bpv *DataBackupPolicyVm) LocalName() string {
	return bpv.Name
}

func (bpv *DataBackupPolicyVm) Configuration() interface{} {
	return bpv.Args
}

func (bpv *DataBackupPolicyVm) Attributes() dataBackupPolicyVmAttributes {
	return dataBackupPolicyVmAttributes{ref: terra.ReferenceDataResource(bpv)}
}

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

func (bpv dataBackupPolicyVmAttributes) Id() terra.StringValue {
	return terra.ReferenceString(bpv.ref.Append("id"))
}

func (bpv dataBackupPolicyVmAttributes) Name() terra.StringValue {
	return terra.ReferenceString(bpv.ref.Append("name"))
}

func (bpv dataBackupPolicyVmAttributes) RecoveryVaultName() terra.StringValue {
	return terra.ReferenceString(bpv.ref.Append("recovery_vault_name"))
}

func (bpv dataBackupPolicyVmAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(bpv.ref.Append("resource_group_name"))
}

func (bpv dataBackupPolicyVmAttributes) Timeouts() databackuppolicyvm.TimeoutsAttributes {
	return terra.ReferenceSingle[databackuppolicyvm.TimeoutsAttributes](bpv.ref.Append("timeouts"))
}
