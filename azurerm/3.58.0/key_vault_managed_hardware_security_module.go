// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	keyvaultmanagedhardwaresecuritymodule "github.com/golingon/terraproviders/azurerm/3.58.0/keyvaultmanagedhardwaresecuritymodule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKeyVaultManagedHardwareSecurityModule creates a new instance of [KeyVaultManagedHardwareSecurityModule].
func NewKeyVaultManagedHardwareSecurityModule(name string, args KeyVaultManagedHardwareSecurityModuleArgs) *KeyVaultManagedHardwareSecurityModule {
	return &KeyVaultManagedHardwareSecurityModule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KeyVaultManagedHardwareSecurityModule)(nil)

// KeyVaultManagedHardwareSecurityModule represents the Terraform resource azurerm_key_vault_managed_hardware_security_module.
type KeyVaultManagedHardwareSecurityModule struct {
	Name      string
	Args      KeyVaultManagedHardwareSecurityModuleArgs
	state     *keyVaultManagedHardwareSecurityModuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KeyVaultManagedHardwareSecurityModule].
func (kvmhsm *KeyVaultManagedHardwareSecurityModule) Type() string {
	return "azurerm_key_vault_managed_hardware_security_module"
}

// LocalName returns the local name for [KeyVaultManagedHardwareSecurityModule].
func (kvmhsm *KeyVaultManagedHardwareSecurityModule) LocalName() string {
	return kvmhsm.Name
}

// Configuration returns the configuration (args) for [KeyVaultManagedHardwareSecurityModule].
func (kvmhsm *KeyVaultManagedHardwareSecurityModule) Configuration() interface{} {
	return kvmhsm.Args
}

// DependOn is used for other resources to depend on [KeyVaultManagedHardwareSecurityModule].
func (kvmhsm *KeyVaultManagedHardwareSecurityModule) DependOn() terra.Reference {
	return terra.ReferenceResource(kvmhsm)
}

// Dependencies returns the list of resources [KeyVaultManagedHardwareSecurityModule] depends_on.
func (kvmhsm *KeyVaultManagedHardwareSecurityModule) Dependencies() terra.Dependencies {
	return kvmhsm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KeyVaultManagedHardwareSecurityModule].
func (kvmhsm *KeyVaultManagedHardwareSecurityModule) LifecycleManagement() *terra.Lifecycle {
	return kvmhsm.Lifecycle
}

// Attributes returns the attributes for [KeyVaultManagedHardwareSecurityModule].
func (kvmhsm *KeyVaultManagedHardwareSecurityModule) Attributes() keyVaultManagedHardwareSecurityModuleAttributes {
	return keyVaultManagedHardwareSecurityModuleAttributes{ref: terra.ReferenceResource(kvmhsm)}
}

// ImportState imports the given attribute values into [KeyVaultManagedHardwareSecurityModule]'s state.
func (kvmhsm *KeyVaultManagedHardwareSecurityModule) ImportState(av io.Reader) error {
	kvmhsm.state = &keyVaultManagedHardwareSecurityModuleState{}
	if err := json.NewDecoder(av).Decode(kvmhsm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kvmhsm.Type(), kvmhsm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KeyVaultManagedHardwareSecurityModule] has state.
func (kvmhsm *KeyVaultManagedHardwareSecurityModule) State() (*keyVaultManagedHardwareSecurityModuleState, bool) {
	return kvmhsm.state, kvmhsm.state != nil
}

// StateMust returns the state for [KeyVaultManagedHardwareSecurityModule]. Panics if the state is nil.
func (kvmhsm *KeyVaultManagedHardwareSecurityModule) StateMust() *keyVaultManagedHardwareSecurityModuleState {
	if kvmhsm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kvmhsm.Type(), kvmhsm.LocalName()))
	}
	return kvmhsm.state
}

// KeyVaultManagedHardwareSecurityModuleArgs contains the configurations for azurerm_key_vault_managed_hardware_security_module.
type KeyVaultManagedHardwareSecurityModuleArgs struct {
	// AdminObjectIds: set of string, required
	AdminObjectIds terra.SetValue[terra.StringValue] `hcl:"admin_object_ids,attr" validate:"required"`
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
	// NetworkAcls: optional
	NetworkAcls *keyvaultmanagedhardwaresecuritymodule.NetworkAcls `hcl:"network_acls,block"`
	// Timeouts: optional
	Timeouts *keyvaultmanagedhardwaresecuritymodule.Timeouts `hcl:"timeouts,block"`
}
type keyVaultManagedHardwareSecurityModuleAttributes struct {
	ref terra.Reference
}

// AdminObjectIds returns a reference to field admin_object_ids of azurerm_key_vault_managed_hardware_security_module.
func (kvmhsm keyVaultManagedHardwareSecurityModuleAttributes) AdminObjectIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](kvmhsm.ref.Append("admin_object_ids"))
}

// HsmUri returns a reference to field hsm_uri of azurerm_key_vault_managed_hardware_security_module.
func (kvmhsm keyVaultManagedHardwareSecurityModuleAttributes) HsmUri() terra.StringValue {
	return terra.ReferenceAsString(kvmhsm.ref.Append("hsm_uri"))
}

// Id returns a reference to field id of azurerm_key_vault_managed_hardware_security_module.
func (kvmhsm keyVaultManagedHardwareSecurityModuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kvmhsm.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_key_vault_managed_hardware_security_module.
func (kvmhsm keyVaultManagedHardwareSecurityModuleAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(kvmhsm.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_key_vault_managed_hardware_security_module.
func (kvmhsm keyVaultManagedHardwareSecurityModuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kvmhsm.ref.Append("name"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_key_vault_managed_hardware_security_module.
func (kvmhsm keyVaultManagedHardwareSecurityModuleAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(kvmhsm.ref.Append("public_network_access_enabled"))
}

// PurgeProtectionEnabled returns a reference to field purge_protection_enabled of azurerm_key_vault_managed_hardware_security_module.
func (kvmhsm keyVaultManagedHardwareSecurityModuleAttributes) PurgeProtectionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(kvmhsm.ref.Append("purge_protection_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_key_vault_managed_hardware_security_module.
func (kvmhsm keyVaultManagedHardwareSecurityModuleAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(kvmhsm.ref.Append("resource_group_name"))
}

// SkuName returns a reference to field sku_name of azurerm_key_vault_managed_hardware_security_module.
func (kvmhsm keyVaultManagedHardwareSecurityModuleAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(kvmhsm.ref.Append("sku_name"))
}

// SoftDeleteRetentionDays returns a reference to field soft_delete_retention_days of azurerm_key_vault_managed_hardware_security_module.
func (kvmhsm keyVaultManagedHardwareSecurityModuleAttributes) SoftDeleteRetentionDays() terra.NumberValue {
	return terra.ReferenceAsNumber(kvmhsm.ref.Append("soft_delete_retention_days"))
}

// Tags returns a reference to field tags of azurerm_key_vault_managed_hardware_security_module.
func (kvmhsm keyVaultManagedHardwareSecurityModuleAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](kvmhsm.ref.Append("tags"))
}

// TenantId returns a reference to field tenant_id of azurerm_key_vault_managed_hardware_security_module.
func (kvmhsm keyVaultManagedHardwareSecurityModuleAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(kvmhsm.ref.Append("tenant_id"))
}

func (kvmhsm keyVaultManagedHardwareSecurityModuleAttributes) NetworkAcls() terra.ListValue[keyvaultmanagedhardwaresecuritymodule.NetworkAclsAttributes] {
	return terra.ReferenceAsList[keyvaultmanagedhardwaresecuritymodule.NetworkAclsAttributes](kvmhsm.ref.Append("network_acls"))
}

func (kvmhsm keyVaultManagedHardwareSecurityModuleAttributes) Timeouts() keyvaultmanagedhardwaresecuritymodule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[keyvaultmanagedhardwaresecuritymodule.TimeoutsAttributes](kvmhsm.ref.Append("timeouts"))
}

type keyVaultManagedHardwareSecurityModuleState struct {
	AdminObjectIds             []string                                                 `json:"admin_object_ids"`
	HsmUri                     string                                                   `json:"hsm_uri"`
	Id                         string                                                   `json:"id"`
	Location                   string                                                   `json:"location"`
	Name                       string                                                   `json:"name"`
	PublicNetworkAccessEnabled bool                                                     `json:"public_network_access_enabled"`
	PurgeProtectionEnabled     bool                                                     `json:"purge_protection_enabled"`
	ResourceGroupName          string                                                   `json:"resource_group_name"`
	SkuName                    string                                                   `json:"sku_name"`
	SoftDeleteRetentionDays    float64                                                  `json:"soft_delete_retention_days"`
	Tags                       map[string]string                                        `json:"tags"`
	TenantId                   string                                                   `json:"tenant_id"`
	NetworkAcls                []keyvaultmanagedhardwaresecuritymodule.NetworkAclsState `json:"network_acls"`
	Timeouts                   *keyvaultmanagedhardwaresecuritymodule.TimeoutsState     `json:"timeouts"`
}
