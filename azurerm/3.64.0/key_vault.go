// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	keyvault "github.com/golingon/terraproviders/azurerm/3.64.0/keyvault"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKeyVault creates a new instance of [KeyVault].
func NewKeyVault(name string, args KeyVaultArgs) *KeyVault {
	return &KeyVault{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KeyVault)(nil)

// KeyVault represents the Terraform resource azurerm_key_vault.
type KeyVault struct {
	Name      string
	Args      KeyVaultArgs
	state     *keyVaultState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KeyVault].
func (kv *KeyVault) Type() string {
	return "azurerm_key_vault"
}

// LocalName returns the local name for [KeyVault].
func (kv *KeyVault) LocalName() string {
	return kv.Name
}

// Configuration returns the configuration (args) for [KeyVault].
func (kv *KeyVault) Configuration() interface{} {
	return kv.Args
}

// DependOn is used for other resources to depend on [KeyVault].
func (kv *KeyVault) DependOn() terra.Reference {
	return terra.ReferenceResource(kv)
}

// Dependencies returns the list of resources [KeyVault] depends_on.
func (kv *KeyVault) Dependencies() terra.Dependencies {
	return kv.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KeyVault].
func (kv *KeyVault) LifecycleManagement() *terra.Lifecycle {
	return kv.Lifecycle
}

// Attributes returns the attributes for [KeyVault].
func (kv *KeyVault) Attributes() keyVaultAttributes {
	return keyVaultAttributes{ref: terra.ReferenceResource(kv)}
}

// ImportState imports the given attribute values into [KeyVault]'s state.
func (kv *KeyVault) ImportState(av io.Reader) error {
	kv.state = &keyVaultState{}
	if err := json.NewDecoder(av).Decode(kv.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kv.Type(), kv.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KeyVault] has state.
func (kv *KeyVault) State() (*keyVaultState, bool) {
	return kv.state, kv.state != nil
}

// StateMust returns the state for [KeyVault]. Panics if the state is nil.
func (kv *KeyVault) StateMust() *keyVaultState {
	if kv.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kv.Type(), kv.LocalName()))
	}
	return kv.state
}

// KeyVaultArgs contains the configurations for azurerm_key_vault.
type KeyVaultArgs struct {
	// EnableRbacAuthorization: bool, optional
	EnableRbacAuthorization terra.BoolValue `hcl:"enable_rbac_authorization,attr"`
	// EnabledForDeployment: bool, optional
	EnabledForDeployment terra.BoolValue `hcl:"enabled_for_deployment,attr"`
	// EnabledForDiskEncryption: bool, optional
	EnabledForDiskEncryption terra.BoolValue `hcl:"enabled_for_disk_encryption,attr"`
	// EnabledForTemplateDeployment: bool, optional
	EnabledForTemplateDeployment terra.BoolValue `hcl:"enabled_for_template_deployment,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PublicNetworkAccessEnabled: bool, optional
	PublicNetworkAccessEnabled terra.BoolValue `hcl:"public_network_access_enabled,attr"`
	// PurgeProtectionEnabled: bool, optional
	PurgeProtectionEnabled terra.BoolValue `hcl:"purge_protection_enabled,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SkuName: string, required
	SkuName terra.StringValue `hcl:"sku_name,attr" validate:"required"`
	// SoftDeleteRetentionDays: number, optional
	SoftDeleteRetentionDays terra.NumberValue `hcl:"soft_delete_retention_days,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TenantId: string, required
	TenantId terra.StringValue `hcl:"tenant_id,attr" validate:"required"`
	// AccessPolicy: min=0
	AccessPolicy []keyvault.AccessPolicy `hcl:"access_policy,block" validate:"min=0"`
	// Contact: min=0
	Contact []keyvault.Contact `hcl:"contact,block" validate:"min=0"`
	// NetworkAcls: optional
	NetworkAcls *keyvault.NetworkAcls `hcl:"network_acls,block"`
	// Timeouts: optional
	Timeouts *keyvault.Timeouts `hcl:"timeouts,block"`
}
type keyVaultAttributes struct {
	ref terra.Reference
}

// EnableRbacAuthorization returns a reference to field enable_rbac_authorization of azurerm_key_vault.
func (kv keyVaultAttributes) EnableRbacAuthorization() terra.BoolValue {
	return terra.ReferenceAsBool(kv.ref.Append("enable_rbac_authorization"))
}

// EnabledForDeployment returns a reference to field enabled_for_deployment of azurerm_key_vault.
func (kv keyVaultAttributes) EnabledForDeployment() terra.BoolValue {
	return terra.ReferenceAsBool(kv.ref.Append("enabled_for_deployment"))
}

// EnabledForDiskEncryption returns a reference to field enabled_for_disk_encryption of azurerm_key_vault.
func (kv keyVaultAttributes) EnabledForDiskEncryption() terra.BoolValue {
	return terra.ReferenceAsBool(kv.ref.Append("enabled_for_disk_encryption"))
}

// EnabledForTemplateDeployment returns a reference to field enabled_for_template_deployment of azurerm_key_vault.
func (kv keyVaultAttributes) EnabledForTemplateDeployment() terra.BoolValue {
	return terra.ReferenceAsBool(kv.ref.Append("enabled_for_template_deployment"))
}

// Id returns a reference to field id of azurerm_key_vault.
func (kv keyVaultAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kv.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_key_vault.
func (kv keyVaultAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(kv.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_key_vault.
func (kv keyVaultAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kv.ref.Append("name"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_key_vault.
func (kv keyVaultAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(kv.ref.Append("public_network_access_enabled"))
}

// PurgeProtectionEnabled returns a reference to field purge_protection_enabled of azurerm_key_vault.
func (kv keyVaultAttributes) PurgeProtectionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(kv.ref.Append("purge_protection_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_key_vault.
func (kv keyVaultAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(kv.ref.Append("resource_group_name"))
}

// SkuName returns a reference to field sku_name of azurerm_key_vault.
func (kv keyVaultAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(kv.ref.Append("sku_name"))
}

// SoftDeleteRetentionDays returns a reference to field soft_delete_retention_days of azurerm_key_vault.
func (kv keyVaultAttributes) SoftDeleteRetentionDays() terra.NumberValue {
	return terra.ReferenceAsNumber(kv.ref.Append("soft_delete_retention_days"))
}

// Tags returns a reference to field tags of azurerm_key_vault.
func (kv keyVaultAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](kv.ref.Append("tags"))
}

// TenantId returns a reference to field tenant_id of azurerm_key_vault.
func (kv keyVaultAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(kv.ref.Append("tenant_id"))
}

// VaultUri returns a reference to field vault_uri of azurerm_key_vault.
func (kv keyVaultAttributes) VaultUri() terra.StringValue {
	return terra.ReferenceAsString(kv.ref.Append("vault_uri"))
}

func (kv keyVaultAttributes) AccessPolicy() terra.ListValue[keyvault.AccessPolicyAttributes] {
	return terra.ReferenceAsList[keyvault.AccessPolicyAttributes](kv.ref.Append("access_policy"))
}

func (kv keyVaultAttributes) Contact() terra.SetValue[keyvault.ContactAttributes] {
	return terra.ReferenceAsSet[keyvault.ContactAttributes](kv.ref.Append("contact"))
}

func (kv keyVaultAttributes) NetworkAcls() terra.ListValue[keyvault.NetworkAclsAttributes] {
	return terra.ReferenceAsList[keyvault.NetworkAclsAttributes](kv.ref.Append("network_acls"))
}

func (kv keyVaultAttributes) Timeouts() keyvault.TimeoutsAttributes {
	return terra.ReferenceAsSingle[keyvault.TimeoutsAttributes](kv.ref.Append("timeouts"))
}

type keyVaultState struct {
	EnableRbacAuthorization      bool                         `json:"enable_rbac_authorization"`
	EnabledForDeployment         bool                         `json:"enabled_for_deployment"`
	EnabledForDiskEncryption     bool                         `json:"enabled_for_disk_encryption"`
	EnabledForTemplateDeployment bool                         `json:"enabled_for_template_deployment"`
	Id                           string                       `json:"id"`
	Location                     string                       `json:"location"`
	Name                         string                       `json:"name"`
	PublicNetworkAccessEnabled   bool                         `json:"public_network_access_enabled"`
	PurgeProtectionEnabled       bool                         `json:"purge_protection_enabled"`
	ResourceGroupName            string                       `json:"resource_group_name"`
	SkuName                      string                       `json:"sku_name"`
	SoftDeleteRetentionDays      float64                      `json:"soft_delete_retention_days"`
	Tags                         map[string]string            `json:"tags"`
	TenantId                     string                       `json:"tenant_id"`
	VaultUri                     string                       `json:"vault_uri"`
	AccessPolicy                 []keyvault.AccessPolicyState `json:"access_policy"`
	Contact                      []keyvault.ContactState      `json:"contact"`
	NetworkAcls                  []keyvault.NetworkAclsState  `json:"network_acls"`
	Timeouts                     *keyvault.TimeoutsState      `json:"timeouts"`
}
