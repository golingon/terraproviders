// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_key_vault

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurerm_key_vault.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (akv *DataSource) DataSource() string {
	return "azurerm_key_vault"
}

// LocalName returns the local name for [DataSource].
func (akv *DataSource) LocalName() string {
	return akv.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (akv *DataSource) Configuration() interface{} {
	return akv.Args
}

// Attributes returns the attributes for [DataSource].
func (akv *DataSource) Attributes() dataAzurermKeyVaultAttributes {
	return dataAzurermKeyVaultAttributes{ref: terra.ReferenceDataSource(akv)}
}

// DataArgs contains the configurations for azurerm_key_vault.
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

type dataAzurermKeyVaultAttributes struct {
	ref terra.Reference
}

// EnableRbacAuthorization returns a reference to field enable_rbac_authorization of azurerm_key_vault.
func (akv dataAzurermKeyVaultAttributes) EnableRbacAuthorization() terra.BoolValue {
	return terra.ReferenceAsBool(akv.ref.Append("enable_rbac_authorization"))
}

// EnabledForDeployment returns a reference to field enabled_for_deployment of azurerm_key_vault.
func (akv dataAzurermKeyVaultAttributes) EnabledForDeployment() terra.BoolValue {
	return terra.ReferenceAsBool(akv.ref.Append("enabled_for_deployment"))
}

// EnabledForDiskEncryption returns a reference to field enabled_for_disk_encryption of azurerm_key_vault.
func (akv dataAzurermKeyVaultAttributes) EnabledForDiskEncryption() terra.BoolValue {
	return terra.ReferenceAsBool(akv.ref.Append("enabled_for_disk_encryption"))
}

// EnabledForTemplateDeployment returns a reference to field enabled_for_template_deployment of azurerm_key_vault.
func (akv dataAzurermKeyVaultAttributes) EnabledForTemplateDeployment() terra.BoolValue {
	return terra.ReferenceAsBool(akv.ref.Append("enabled_for_template_deployment"))
}

// Id returns a reference to field id of azurerm_key_vault.
func (akv dataAzurermKeyVaultAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(akv.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_key_vault.
func (akv dataAzurermKeyVaultAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(akv.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_key_vault.
func (akv dataAzurermKeyVaultAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(akv.ref.Append("name"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_key_vault.
func (akv dataAzurermKeyVaultAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(akv.ref.Append("public_network_access_enabled"))
}

// PurgeProtectionEnabled returns a reference to field purge_protection_enabled of azurerm_key_vault.
func (akv dataAzurermKeyVaultAttributes) PurgeProtectionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(akv.ref.Append("purge_protection_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_key_vault.
func (akv dataAzurermKeyVaultAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(akv.ref.Append("resource_group_name"))
}

// SkuName returns a reference to field sku_name of azurerm_key_vault.
func (akv dataAzurermKeyVaultAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(akv.ref.Append("sku_name"))
}

// Tags returns a reference to field tags of azurerm_key_vault.
func (akv dataAzurermKeyVaultAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](akv.ref.Append("tags"))
}

// TenantId returns a reference to field tenant_id of azurerm_key_vault.
func (akv dataAzurermKeyVaultAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(akv.ref.Append("tenant_id"))
}

// VaultUri returns a reference to field vault_uri of azurerm_key_vault.
func (akv dataAzurermKeyVaultAttributes) VaultUri() terra.StringValue {
	return terra.ReferenceAsString(akv.ref.Append("vault_uri"))
}

func (akv dataAzurermKeyVaultAttributes) AccessPolicy() terra.ListValue[DataAccessPolicyAttributes] {
	return terra.ReferenceAsList[DataAccessPolicyAttributes](akv.ref.Append("access_policy"))
}

func (akv dataAzurermKeyVaultAttributes) NetworkAcls() terra.ListValue[DataNetworkAclsAttributes] {
	return terra.ReferenceAsList[DataNetworkAclsAttributes](akv.ref.Append("network_acls"))
}

func (akv dataAzurermKeyVaultAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](akv.ref.Append("timeouts"))
}
