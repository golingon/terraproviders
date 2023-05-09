// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	keyvaultcertificateissuer "github.com/golingon/terraproviders/azurerm/3.55.0/keyvaultcertificateissuer"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKeyVaultCertificateIssuer creates a new instance of [KeyVaultCertificateIssuer].
func NewKeyVaultCertificateIssuer(name string, args KeyVaultCertificateIssuerArgs) *KeyVaultCertificateIssuer {
	return &KeyVaultCertificateIssuer{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KeyVaultCertificateIssuer)(nil)

// KeyVaultCertificateIssuer represents the Terraform resource azurerm_key_vault_certificate_issuer.
type KeyVaultCertificateIssuer struct {
	Name      string
	Args      KeyVaultCertificateIssuerArgs
	state     *keyVaultCertificateIssuerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KeyVaultCertificateIssuer].
func (kvci *KeyVaultCertificateIssuer) Type() string {
	return "azurerm_key_vault_certificate_issuer"
}

// LocalName returns the local name for [KeyVaultCertificateIssuer].
func (kvci *KeyVaultCertificateIssuer) LocalName() string {
	return kvci.Name
}

// Configuration returns the configuration (args) for [KeyVaultCertificateIssuer].
func (kvci *KeyVaultCertificateIssuer) Configuration() interface{} {
	return kvci.Args
}

// DependOn is used for other resources to depend on [KeyVaultCertificateIssuer].
func (kvci *KeyVaultCertificateIssuer) DependOn() terra.Reference {
	return terra.ReferenceResource(kvci)
}

// Dependencies returns the list of resources [KeyVaultCertificateIssuer] depends_on.
func (kvci *KeyVaultCertificateIssuer) Dependencies() terra.Dependencies {
	return kvci.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KeyVaultCertificateIssuer].
func (kvci *KeyVaultCertificateIssuer) LifecycleManagement() *terra.Lifecycle {
	return kvci.Lifecycle
}

// Attributes returns the attributes for [KeyVaultCertificateIssuer].
func (kvci *KeyVaultCertificateIssuer) Attributes() keyVaultCertificateIssuerAttributes {
	return keyVaultCertificateIssuerAttributes{ref: terra.ReferenceResource(kvci)}
}

// ImportState imports the given attribute values into [KeyVaultCertificateIssuer]'s state.
func (kvci *KeyVaultCertificateIssuer) ImportState(av io.Reader) error {
	kvci.state = &keyVaultCertificateIssuerState{}
	if err := json.NewDecoder(av).Decode(kvci.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kvci.Type(), kvci.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KeyVaultCertificateIssuer] has state.
func (kvci *KeyVaultCertificateIssuer) State() (*keyVaultCertificateIssuerState, bool) {
	return kvci.state, kvci.state != nil
}

// StateMust returns the state for [KeyVaultCertificateIssuer]. Panics if the state is nil.
func (kvci *KeyVaultCertificateIssuer) StateMust() *keyVaultCertificateIssuerState {
	if kvci.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kvci.Type(), kvci.LocalName()))
	}
	return kvci.state
}

// KeyVaultCertificateIssuerArgs contains the configurations for azurerm_key_vault_certificate_issuer.
type KeyVaultCertificateIssuerArgs struct {
	// AccountId: string, optional
	AccountId terra.StringValue `hcl:"account_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyVaultId: string, required
	KeyVaultId terra.StringValue `hcl:"key_vault_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// OrgId: string, optional
	OrgId terra.StringValue `hcl:"org_id,attr"`
	// Password: string, optional
	Password terra.StringValue `hcl:"password,attr"`
	// ProviderName: string, required
	ProviderName terra.StringValue `hcl:"provider_name,attr" validate:"required"`
	// Admin: min=0
	Admin []keyvaultcertificateissuer.Admin `hcl:"admin,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *keyvaultcertificateissuer.Timeouts `hcl:"timeouts,block"`
}
type keyVaultCertificateIssuerAttributes struct {
	ref terra.Reference
}

// AccountId returns a reference to field account_id of azurerm_key_vault_certificate_issuer.
func (kvci keyVaultCertificateIssuerAttributes) AccountId() terra.StringValue {
	return terra.ReferenceAsString(kvci.ref.Append("account_id"))
}

// Id returns a reference to field id of azurerm_key_vault_certificate_issuer.
func (kvci keyVaultCertificateIssuerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kvci.ref.Append("id"))
}

// KeyVaultId returns a reference to field key_vault_id of azurerm_key_vault_certificate_issuer.
func (kvci keyVaultCertificateIssuerAttributes) KeyVaultId() terra.StringValue {
	return terra.ReferenceAsString(kvci.ref.Append("key_vault_id"))
}

// Name returns a reference to field name of azurerm_key_vault_certificate_issuer.
func (kvci keyVaultCertificateIssuerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kvci.ref.Append("name"))
}

// OrgId returns a reference to field org_id of azurerm_key_vault_certificate_issuer.
func (kvci keyVaultCertificateIssuerAttributes) OrgId() terra.StringValue {
	return terra.ReferenceAsString(kvci.ref.Append("org_id"))
}

// Password returns a reference to field password of azurerm_key_vault_certificate_issuer.
func (kvci keyVaultCertificateIssuerAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(kvci.ref.Append("password"))
}

// ProviderName returns a reference to field provider_name of azurerm_key_vault_certificate_issuer.
func (kvci keyVaultCertificateIssuerAttributes) ProviderName() terra.StringValue {
	return terra.ReferenceAsString(kvci.ref.Append("provider_name"))
}

func (kvci keyVaultCertificateIssuerAttributes) Admin() terra.ListValue[keyvaultcertificateissuer.AdminAttributes] {
	return terra.ReferenceAsList[keyvaultcertificateissuer.AdminAttributes](kvci.ref.Append("admin"))
}

func (kvci keyVaultCertificateIssuerAttributes) Timeouts() keyvaultcertificateissuer.TimeoutsAttributes {
	return terra.ReferenceAsSingle[keyvaultcertificateissuer.TimeoutsAttributes](kvci.ref.Append("timeouts"))
}

type keyVaultCertificateIssuerState struct {
	AccountId    string                                   `json:"account_id"`
	Id           string                                   `json:"id"`
	KeyVaultId   string                                   `json:"key_vault_id"`
	Name         string                                   `json:"name"`
	OrgId        string                                   `json:"org_id"`
	Password     string                                   `json:"password"`
	ProviderName string                                   `json:"provider_name"`
	Admin        []keyvaultcertificateissuer.AdminState   `json:"admin"`
	Timeouts     *keyvaultcertificateissuer.TimeoutsState `json:"timeouts"`
}
