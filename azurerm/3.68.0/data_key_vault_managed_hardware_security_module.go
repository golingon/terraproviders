// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datakeyvaultmanagedhardwaresecuritymodule "github.com/golingon/terraproviders/azurerm/3.68.0/datakeyvaultmanagedhardwaresecuritymodule"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataKeyVaultManagedHardwareSecurityModule creates a new instance of [DataKeyVaultManagedHardwareSecurityModule].
func NewDataKeyVaultManagedHardwareSecurityModule(name string, args DataKeyVaultManagedHardwareSecurityModuleArgs) *DataKeyVaultManagedHardwareSecurityModule {
	return &DataKeyVaultManagedHardwareSecurityModule{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataKeyVaultManagedHardwareSecurityModule)(nil)

// DataKeyVaultManagedHardwareSecurityModule represents the Terraform data resource azurerm_key_vault_managed_hardware_security_module.
type DataKeyVaultManagedHardwareSecurityModule struct {
	Name string
	Args DataKeyVaultManagedHardwareSecurityModuleArgs
}

// DataSource returns the Terraform object type for [DataKeyVaultManagedHardwareSecurityModule].
func (kvmhsm *DataKeyVaultManagedHardwareSecurityModule) DataSource() string {
	return "azurerm_key_vault_managed_hardware_security_module"
}

// LocalName returns the local name for [DataKeyVaultManagedHardwareSecurityModule].
func (kvmhsm *DataKeyVaultManagedHardwareSecurityModule) LocalName() string {
	return kvmhsm.Name
}

// Configuration returns the configuration (args) for [DataKeyVaultManagedHardwareSecurityModule].
func (kvmhsm *DataKeyVaultManagedHardwareSecurityModule) Configuration() interface{} {
	return kvmhsm.Args
}

// Attributes returns the attributes for [DataKeyVaultManagedHardwareSecurityModule].
func (kvmhsm *DataKeyVaultManagedHardwareSecurityModule) Attributes() dataKeyVaultManagedHardwareSecurityModuleAttributes {
	return dataKeyVaultManagedHardwareSecurityModuleAttributes{ref: terra.ReferenceDataResource(kvmhsm)}
}

// DataKeyVaultManagedHardwareSecurityModuleArgs contains the configurations for azurerm_key_vault_managed_hardware_security_module.
type DataKeyVaultManagedHardwareSecurityModuleArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datakeyvaultmanagedhardwaresecuritymodule.Timeouts `hcl:"timeouts,block"`
}
type dataKeyVaultManagedHardwareSecurityModuleAttributes struct {
	ref terra.Reference
}

// AdminObjectIds returns a reference to field admin_object_ids of azurerm_key_vault_managed_hardware_security_module.
func (kvmhsm dataKeyVaultManagedHardwareSecurityModuleAttributes) AdminObjectIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](kvmhsm.ref.Append("admin_object_ids"))
}

// HsmUri returns a reference to field hsm_uri of azurerm_key_vault_managed_hardware_security_module.
func (kvmhsm dataKeyVaultManagedHardwareSecurityModuleAttributes) HsmUri() terra.StringValue {
	return terra.ReferenceAsString(kvmhsm.ref.Append("hsm_uri"))
}

// Id returns a reference to field id of azurerm_key_vault_managed_hardware_security_module.
func (kvmhsm dataKeyVaultManagedHardwareSecurityModuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kvmhsm.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_key_vault_managed_hardware_security_module.
func (kvmhsm dataKeyVaultManagedHardwareSecurityModuleAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(kvmhsm.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_key_vault_managed_hardware_security_module.
func (kvmhsm dataKeyVaultManagedHardwareSecurityModuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kvmhsm.ref.Append("name"))
}

// PurgeProtectionEnabled returns a reference to field purge_protection_enabled of azurerm_key_vault_managed_hardware_security_module.
func (kvmhsm dataKeyVaultManagedHardwareSecurityModuleAttributes) PurgeProtectionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(kvmhsm.ref.Append("purge_protection_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_key_vault_managed_hardware_security_module.
func (kvmhsm dataKeyVaultManagedHardwareSecurityModuleAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(kvmhsm.ref.Append("resource_group_name"))
}

// SkuName returns a reference to field sku_name of azurerm_key_vault_managed_hardware_security_module.
func (kvmhsm dataKeyVaultManagedHardwareSecurityModuleAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(kvmhsm.ref.Append("sku_name"))
}

// SoftDeleteRetentionDays returns a reference to field soft_delete_retention_days of azurerm_key_vault_managed_hardware_security_module.
func (kvmhsm dataKeyVaultManagedHardwareSecurityModuleAttributes) SoftDeleteRetentionDays() terra.NumberValue {
	return terra.ReferenceAsNumber(kvmhsm.ref.Append("soft_delete_retention_days"))
}

// Tags returns a reference to field tags of azurerm_key_vault_managed_hardware_security_module.
func (kvmhsm dataKeyVaultManagedHardwareSecurityModuleAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](kvmhsm.ref.Append("tags"))
}

// TenantId returns a reference to field tenant_id of azurerm_key_vault_managed_hardware_security_module.
func (kvmhsm dataKeyVaultManagedHardwareSecurityModuleAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(kvmhsm.ref.Append("tenant_id"))
}

func (kvmhsm dataKeyVaultManagedHardwareSecurityModuleAttributes) Timeouts() datakeyvaultmanagedhardwaresecuritymodule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datakeyvaultmanagedhardwaresecuritymodule.TimeoutsAttributes](kvmhsm.ref.Append("timeouts"))
}
