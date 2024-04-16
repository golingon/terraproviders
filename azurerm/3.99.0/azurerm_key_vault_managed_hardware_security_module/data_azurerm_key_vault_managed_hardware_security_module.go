// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_key_vault_managed_hardware_security_module

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurerm_key_vault_managed_hardware_security_module.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (akvmhsm *DataSource) DataSource() string {
	return "azurerm_key_vault_managed_hardware_security_module"
}

// LocalName returns the local name for [DataSource].
func (akvmhsm *DataSource) LocalName() string {
	return akvmhsm.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (akvmhsm *DataSource) Configuration() interface{} {
	return akvmhsm.Args
}

// Attributes returns the attributes for [DataSource].
func (akvmhsm *DataSource) Attributes() dataAzurermKeyVaultManagedHardwareSecurityModuleAttributes {
	return dataAzurermKeyVaultManagedHardwareSecurityModuleAttributes{ref: terra.ReferenceDataSource(akvmhsm)}
}

// DataArgs contains the configurations for azurerm_key_vault_managed_hardware_security_module.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAzurermKeyVaultManagedHardwareSecurityModuleAttributes struct {
	ref terra.Reference
}

// AdminObjectIds returns a reference to field admin_object_ids of azurerm_key_vault_managed_hardware_security_module.
func (akvmhsm dataAzurermKeyVaultManagedHardwareSecurityModuleAttributes) AdminObjectIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](akvmhsm.ref.Append("admin_object_ids"))
}

// HsmUri returns a reference to field hsm_uri of azurerm_key_vault_managed_hardware_security_module.
func (akvmhsm dataAzurermKeyVaultManagedHardwareSecurityModuleAttributes) HsmUri() terra.StringValue {
	return terra.ReferenceAsString(akvmhsm.ref.Append("hsm_uri"))
}

// Id returns a reference to field id of azurerm_key_vault_managed_hardware_security_module.
func (akvmhsm dataAzurermKeyVaultManagedHardwareSecurityModuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(akvmhsm.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_key_vault_managed_hardware_security_module.
func (akvmhsm dataAzurermKeyVaultManagedHardwareSecurityModuleAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(akvmhsm.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_key_vault_managed_hardware_security_module.
func (akvmhsm dataAzurermKeyVaultManagedHardwareSecurityModuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(akvmhsm.ref.Append("name"))
}

// PurgeProtectionEnabled returns a reference to field purge_protection_enabled of azurerm_key_vault_managed_hardware_security_module.
func (akvmhsm dataAzurermKeyVaultManagedHardwareSecurityModuleAttributes) PurgeProtectionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(akvmhsm.ref.Append("purge_protection_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_key_vault_managed_hardware_security_module.
func (akvmhsm dataAzurermKeyVaultManagedHardwareSecurityModuleAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(akvmhsm.ref.Append("resource_group_name"))
}

// SkuName returns a reference to field sku_name of azurerm_key_vault_managed_hardware_security_module.
func (akvmhsm dataAzurermKeyVaultManagedHardwareSecurityModuleAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(akvmhsm.ref.Append("sku_name"))
}

// SoftDeleteRetentionDays returns a reference to field soft_delete_retention_days of azurerm_key_vault_managed_hardware_security_module.
func (akvmhsm dataAzurermKeyVaultManagedHardwareSecurityModuleAttributes) SoftDeleteRetentionDays() terra.NumberValue {
	return terra.ReferenceAsNumber(akvmhsm.ref.Append("soft_delete_retention_days"))
}

// Tags returns a reference to field tags of azurerm_key_vault_managed_hardware_security_module.
func (akvmhsm dataAzurermKeyVaultManagedHardwareSecurityModuleAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](akvmhsm.ref.Append("tags"))
}

// TenantId returns a reference to field tenant_id of azurerm_key_vault_managed_hardware_security_module.
func (akvmhsm dataAzurermKeyVaultManagedHardwareSecurityModuleAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(akvmhsm.ref.Append("tenant_id"))
}

func (akvmhsm dataAzurermKeyVaultManagedHardwareSecurityModuleAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](akvmhsm.ref.Append("timeouts"))
}
