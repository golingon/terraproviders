// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_key_vault

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

// Resource represents the Terraform resource azurerm_key_vault.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermKeyVaultState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (akv *Resource) Type() string {
	return "azurerm_key_vault"
}

// LocalName returns the local name for [Resource].
func (akv *Resource) LocalName() string {
	return akv.Name
}

// Configuration returns the configuration (args) for [Resource].
func (akv *Resource) Configuration() interface{} {
	return akv.Args
}

// DependOn is used for other resources to depend on [Resource].
func (akv *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(akv)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (akv *Resource) Dependencies() terra.Dependencies {
	return akv.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (akv *Resource) LifecycleManagement() *terra.Lifecycle {
	return akv.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (akv *Resource) Attributes() azurermKeyVaultAttributes {
	return azurermKeyVaultAttributes{ref: terra.ReferenceResource(akv)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (akv *Resource) ImportState(state io.Reader) error {
	akv.state = &azurermKeyVaultState{}
	if err := json.NewDecoder(state).Decode(akv.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", akv.Type(), akv.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (akv *Resource) State() (*azurermKeyVaultState, bool) {
	return akv.state, akv.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (akv *Resource) StateMust() *azurermKeyVaultState {
	if akv.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", akv.Type(), akv.LocalName()))
	}
	return akv.state
}

// Args contains the configurations for azurerm_key_vault.
type Args struct {
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
	AccessPolicy []AccessPolicy `hcl:"access_policy,block" validate:"min=0"`
	// Contact: min=0
	Contact []Contact `hcl:"contact,block" validate:"min=0"`
	// NetworkAcls: optional
	NetworkAcls *NetworkAcls `hcl:"network_acls,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermKeyVaultAttributes struct {
	ref terra.Reference
}

// EnableRbacAuthorization returns a reference to field enable_rbac_authorization of azurerm_key_vault.
func (akv azurermKeyVaultAttributes) EnableRbacAuthorization() terra.BoolValue {
	return terra.ReferenceAsBool(akv.ref.Append("enable_rbac_authorization"))
}

// EnabledForDeployment returns a reference to field enabled_for_deployment of azurerm_key_vault.
func (akv azurermKeyVaultAttributes) EnabledForDeployment() terra.BoolValue {
	return terra.ReferenceAsBool(akv.ref.Append("enabled_for_deployment"))
}

// EnabledForDiskEncryption returns a reference to field enabled_for_disk_encryption of azurerm_key_vault.
func (akv azurermKeyVaultAttributes) EnabledForDiskEncryption() terra.BoolValue {
	return terra.ReferenceAsBool(akv.ref.Append("enabled_for_disk_encryption"))
}

// EnabledForTemplateDeployment returns a reference to field enabled_for_template_deployment of azurerm_key_vault.
func (akv azurermKeyVaultAttributes) EnabledForTemplateDeployment() terra.BoolValue {
	return terra.ReferenceAsBool(akv.ref.Append("enabled_for_template_deployment"))
}

// Id returns a reference to field id of azurerm_key_vault.
func (akv azurermKeyVaultAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(akv.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_key_vault.
func (akv azurermKeyVaultAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(akv.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_key_vault.
func (akv azurermKeyVaultAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(akv.ref.Append("name"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_key_vault.
func (akv azurermKeyVaultAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(akv.ref.Append("public_network_access_enabled"))
}

// PurgeProtectionEnabled returns a reference to field purge_protection_enabled of azurerm_key_vault.
func (akv azurermKeyVaultAttributes) PurgeProtectionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(akv.ref.Append("purge_protection_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_key_vault.
func (akv azurermKeyVaultAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(akv.ref.Append("resource_group_name"))
}

// SkuName returns a reference to field sku_name of azurerm_key_vault.
func (akv azurermKeyVaultAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(akv.ref.Append("sku_name"))
}

// SoftDeleteRetentionDays returns a reference to field soft_delete_retention_days of azurerm_key_vault.
func (akv azurermKeyVaultAttributes) SoftDeleteRetentionDays() terra.NumberValue {
	return terra.ReferenceAsNumber(akv.ref.Append("soft_delete_retention_days"))
}

// Tags returns a reference to field tags of azurerm_key_vault.
func (akv azurermKeyVaultAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](akv.ref.Append("tags"))
}

// TenantId returns a reference to field tenant_id of azurerm_key_vault.
func (akv azurermKeyVaultAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(akv.ref.Append("tenant_id"))
}

// VaultUri returns a reference to field vault_uri of azurerm_key_vault.
func (akv azurermKeyVaultAttributes) VaultUri() terra.StringValue {
	return terra.ReferenceAsString(akv.ref.Append("vault_uri"))
}

func (akv azurermKeyVaultAttributes) AccessPolicy() terra.ListValue[AccessPolicyAttributes] {
	return terra.ReferenceAsList[AccessPolicyAttributes](akv.ref.Append("access_policy"))
}

func (akv azurermKeyVaultAttributes) Contact() terra.SetValue[ContactAttributes] {
	return terra.ReferenceAsSet[ContactAttributes](akv.ref.Append("contact"))
}

func (akv azurermKeyVaultAttributes) NetworkAcls() terra.ListValue[NetworkAclsAttributes] {
	return terra.ReferenceAsList[NetworkAclsAttributes](akv.ref.Append("network_acls"))
}

func (akv azurermKeyVaultAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](akv.ref.Append("timeouts"))
}

type azurermKeyVaultState struct {
	EnableRbacAuthorization      bool                `json:"enable_rbac_authorization"`
	EnabledForDeployment         bool                `json:"enabled_for_deployment"`
	EnabledForDiskEncryption     bool                `json:"enabled_for_disk_encryption"`
	EnabledForTemplateDeployment bool                `json:"enabled_for_template_deployment"`
	Id                           string              `json:"id"`
	Location                     string              `json:"location"`
	Name                         string              `json:"name"`
	PublicNetworkAccessEnabled   bool                `json:"public_network_access_enabled"`
	PurgeProtectionEnabled       bool                `json:"purge_protection_enabled"`
	ResourceGroupName            string              `json:"resource_group_name"`
	SkuName                      string              `json:"sku_name"`
	SoftDeleteRetentionDays      float64             `json:"soft_delete_retention_days"`
	Tags                         map[string]string   `json:"tags"`
	TenantId                     string              `json:"tenant_id"`
	VaultUri                     string              `json:"vault_uri"`
	AccessPolicy                 []AccessPolicyState `json:"access_policy"`
	Contact                      []ContactState      `json:"contact"`
	NetworkAcls                  []NetworkAclsState  `json:"network_acls"`
	Timeouts                     *TimeoutsState      `json:"timeouts"`
}
