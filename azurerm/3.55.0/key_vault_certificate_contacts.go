// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	keyvaultcertificatecontacts "github.com/golingon/terraproviders/azurerm/3.55.0/keyvaultcertificatecontacts"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKeyVaultCertificateContacts creates a new instance of [KeyVaultCertificateContacts].
func NewKeyVaultCertificateContacts(name string, args KeyVaultCertificateContactsArgs) *KeyVaultCertificateContacts {
	return &KeyVaultCertificateContacts{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KeyVaultCertificateContacts)(nil)

// KeyVaultCertificateContacts represents the Terraform resource azurerm_key_vault_certificate_contacts.
type KeyVaultCertificateContacts struct {
	Name      string
	Args      KeyVaultCertificateContactsArgs
	state     *keyVaultCertificateContactsState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KeyVaultCertificateContacts].
func (kvcc *KeyVaultCertificateContacts) Type() string {
	return "azurerm_key_vault_certificate_contacts"
}

// LocalName returns the local name for [KeyVaultCertificateContacts].
func (kvcc *KeyVaultCertificateContacts) LocalName() string {
	return kvcc.Name
}

// Configuration returns the configuration (args) for [KeyVaultCertificateContacts].
func (kvcc *KeyVaultCertificateContacts) Configuration() interface{} {
	return kvcc.Args
}

// DependOn is used for other resources to depend on [KeyVaultCertificateContacts].
func (kvcc *KeyVaultCertificateContacts) DependOn() terra.Reference {
	return terra.ReferenceResource(kvcc)
}

// Dependencies returns the list of resources [KeyVaultCertificateContacts] depends_on.
func (kvcc *KeyVaultCertificateContacts) Dependencies() terra.Dependencies {
	return kvcc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KeyVaultCertificateContacts].
func (kvcc *KeyVaultCertificateContacts) LifecycleManagement() *terra.Lifecycle {
	return kvcc.Lifecycle
}

// Attributes returns the attributes for [KeyVaultCertificateContacts].
func (kvcc *KeyVaultCertificateContacts) Attributes() keyVaultCertificateContactsAttributes {
	return keyVaultCertificateContactsAttributes{ref: terra.ReferenceResource(kvcc)}
}

// ImportState imports the given attribute values into [KeyVaultCertificateContacts]'s state.
func (kvcc *KeyVaultCertificateContacts) ImportState(av io.Reader) error {
	kvcc.state = &keyVaultCertificateContactsState{}
	if err := json.NewDecoder(av).Decode(kvcc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kvcc.Type(), kvcc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KeyVaultCertificateContacts] has state.
func (kvcc *KeyVaultCertificateContacts) State() (*keyVaultCertificateContactsState, bool) {
	return kvcc.state, kvcc.state != nil
}

// StateMust returns the state for [KeyVaultCertificateContacts]. Panics if the state is nil.
func (kvcc *KeyVaultCertificateContacts) StateMust() *keyVaultCertificateContactsState {
	if kvcc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kvcc.Type(), kvcc.LocalName()))
	}
	return kvcc.state
}

// KeyVaultCertificateContactsArgs contains the configurations for azurerm_key_vault_certificate_contacts.
type KeyVaultCertificateContactsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyVaultId: string, required
	KeyVaultId terra.StringValue `hcl:"key_vault_id,attr" validate:"required"`
	// Contact: min=1
	Contact []keyvaultcertificatecontacts.Contact `hcl:"contact,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *keyvaultcertificatecontacts.Timeouts `hcl:"timeouts,block"`
}
type keyVaultCertificateContactsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_key_vault_certificate_contacts.
func (kvcc keyVaultCertificateContactsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kvcc.ref.Append("id"))
}

// KeyVaultId returns a reference to field key_vault_id of azurerm_key_vault_certificate_contacts.
func (kvcc keyVaultCertificateContactsAttributes) KeyVaultId() terra.StringValue {
	return terra.ReferenceAsString(kvcc.ref.Append("key_vault_id"))
}

func (kvcc keyVaultCertificateContactsAttributes) Contact() terra.SetValue[keyvaultcertificatecontacts.ContactAttributes] {
	return terra.ReferenceAsSet[keyvaultcertificatecontacts.ContactAttributes](kvcc.ref.Append("contact"))
}

func (kvcc keyVaultCertificateContactsAttributes) Timeouts() keyvaultcertificatecontacts.TimeoutsAttributes {
	return terra.ReferenceAsSingle[keyvaultcertificatecontacts.TimeoutsAttributes](kvcc.ref.Append("timeouts"))
}

type keyVaultCertificateContactsState struct {
	Id         string                                     `json:"id"`
	KeyVaultId string                                     `json:"key_vault_id"`
	Contact    []keyvaultcertificatecontacts.ContactState `json:"contact"`
	Timeouts   *keyvaultcertificatecontacts.TimeoutsState `json:"timeouts"`
}
