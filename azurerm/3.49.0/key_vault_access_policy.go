// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	keyvaultaccesspolicy "github.com/golingon/terraproviders/azurerm/3.49.0/keyvaultaccesspolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKeyVaultAccessPolicy creates a new instance of [KeyVaultAccessPolicy].
func NewKeyVaultAccessPolicy(name string, args KeyVaultAccessPolicyArgs) *KeyVaultAccessPolicy {
	return &KeyVaultAccessPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KeyVaultAccessPolicy)(nil)

// KeyVaultAccessPolicy represents the Terraform resource azurerm_key_vault_access_policy.
type KeyVaultAccessPolicy struct {
	Name      string
	Args      KeyVaultAccessPolicyArgs
	state     *keyVaultAccessPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KeyVaultAccessPolicy].
func (kvap *KeyVaultAccessPolicy) Type() string {
	return "azurerm_key_vault_access_policy"
}

// LocalName returns the local name for [KeyVaultAccessPolicy].
func (kvap *KeyVaultAccessPolicy) LocalName() string {
	return kvap.Name
}

// Configuration returns the configuration (args) for [KeyVaultAccessPolicy].
func (kvap *KeyVaultAccessPolicy) Configuration() interface{} {
	return kvap.Args
}

// DependOn is used for other resources to depend on [KeyVaultAccessPolicy].
func (kvap *KeyVaultAccessPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(kvap)
}

// Dependencies returns the list of resources [KeyVaultAccessPolicy] depends_on.
func (kvap *KeyVaultAccessPolicy) Dependencies() terra.Dependencies {
	return kvap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KeyVaultAccessPolicy].
func (kvap *KeyVaultAccessPolicy) LifecycleManagement() *terra.Lifecycle {
	return kvap.Lifecycle
}

// Attributes returns the attributes for [KeyVaultAccessPolicy].
func (kvap *KeyVaultAccessPolicy) Attributes() keyVaultAccessPolicyAttributes {
	return keyVaultAccessPolicyAttributes{ref: terra.ReferenceResource(kvap)}
}

// ImportState imports the given attribute values into [KeyVaultAccessPolicy]'s state.
func (kvap *KeyVaultAccessPolicy) ImportState(av io.Reader) error {
	kvap.state = &keyVaultAccessPolicyState{}
	if err := json.NewDecoder(av).Decode(kvap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kvap.Type(), kvap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KeyVaultAccessPolicy] has state.
func (kvap *KeyVaultAccessPolicy) State() (*keyVaultAccessPolicyState, bool) {
	return kvap.state, kvap.state != nil
}

// StateMust returns the state for [KeyVaultAccessPolicy]. Panics if the state is nil.
func (kvap *KeyVaultAccessPolicy) StateMust() *keyVaultAccessPolicyState {
	if kvap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kvap.Type(), kvap.LocalName()))
	}
	return kvap.state
}

// KeyVaultAccessPolicyArgs contains the configurations for azurerm_key_vault_access_policy.
type KeyVaultAccessPolicyArgs struct {
	// ApplicationId: string, optional
	ApplicationId terra.StringValue `hcl:"application_id,attr"`
	// CertificatePermissions: list of string, optional
	CertificatePermissions terra.ListValue[terra.StringValue] `hcl:"certificate_permissions,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyPermissions: list of string, optional
	KeyPermissions terra.ListValue[terra.StringValue] `hcl:"key_permissions,attr"`
	// KeyVaultId: string, required
	KeyVaultId terra.StringValue `hcl:"key_vault_id,attr" validate:"required"`
	// ObjectId: string, required
	ObjectId terra.StringValue `hcl:"object_id,attr" validate:"required"`
	// SecretPermissions: list of string, optional
	SecretPermissions terra.ListValue[terra.StringValue] `hcl:"secret_permissions,attr"`
	// StoragePermissions: list of string, optional
	StoragePermissions terra.ListValue[terra.StringValue] `hcl:"storage_permissions,attr"`
	// TenantId: string, required
	TenantId terra.StringValue `hcl:"tenant_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *keyvaultaccesspolicy.Timeouts `hcl:"timeouts,block"`
}
type keyVaultAccessPolicyAttributes struct {
	ref terra.Reference
}

// ApplicationId returns a reference to field application_id of azurerm_key_vault_access_policy.
func (kvap keyVaultAccessPolicyAttributes) ApplicationId() terra.StringValue {
	return terra.ReferenceAsString(kvap.ref.Append("application_id"))
}

// CertificatePermissions returns a reference to field certificate_permissions of azurerm_key_vault_access_policy.
func (kvap keyVaultAccessPolicyAttributes) CertificatePermissions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](kvap.ref.Append("certificate_permissions"))
}

// Id returns a reference to field id of azurerm_key_vault_access_policy.
func (kvap keyVaultAccessPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kvap.ref.Append("id"))
}

// KeyPermissions returns a reference to field key_permissions of azurerm_key_vault_access_policy.
func (kvap keyVaultAccessPolicyAttributes) KeyPermissions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](kvap.ref.Append("key_permissions"))
}

// KeyVaultId returns a reference to field key_vault_id of azurerm_key_vault_access_policy.
func (kvap keyVaultAccessPolicyAttributes) KeyVaultId() terra.StringValue {
	return terra.ReferenceAsString(kvap.ref.Append("key_vault_id"))
}

// ObjectId returns a reference to field object_id of azurerm_key_vault_access_policy.
func (kvap keyVaultAccessPolicyAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceAsString(kvap.ref.Append("object_id"))
}

// SecretPermissions returns a reference to field secret_permissions of azurerm_key_vault_access_policy.
func (kvap keyVaultAccessPolicyAttributes) SecretPermissions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](kvap.ref.Append("secret_permissions"))
}

// StoragePermissions returns a reference to field storage_permissions of azurerm_key_vault_access_policy.
func (kvap keyVaultAccessPolicyAttributes) StoragePermissions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](kvap.ref.Append("storage_permissions"))
}

// TenantId returns a reference to field tenant_id of azurerm_key_vault_access_policy.
func (kvap keyVaultAccessPolicyAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(kvap.ref.Append("tenant_id"))
}

func (kvap keyVaultAccessPolicyAttributes) Timeouts() keyvaultaccesspolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[keyvaultaccesspolicy.TimeoutsAttributes](kvap.ref.Append("timeouts"))
}

type keyVaultAccessPolicyState struct {
	ApplicationId          string                              `json:"application_id"`
	CertificatePermissions []string                            `json:"certificate_permissions"`
	Id                     string                              `json:"id"`
	KeyPermissions         []string                            `json:"key_permissions"`
	KeyVaultId             string                              `json:"key_vault_id"`
	ObjectId               string                              `json:"object_id"`
	SecretPermissions      []string                            `json:"secret_permissions"`
	StoragePermissions     []string                            `json:"storage_permissions"`
	TenantId               string                              `json:"tenant_id"`
	Timeouts               *keyvaultaccesspolicy.TimeoutsState `json:"timeouts"`
}
