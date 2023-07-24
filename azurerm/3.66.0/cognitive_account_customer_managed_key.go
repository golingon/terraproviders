// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	cognitiveaccountcustomermanagedkey "github.com/golingon/terraproviders/azurerm/3.66.0/cognitiveaccountcustomermanagedkey"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCognitiveAccountCustomerManagedKey creates a new instance of [CognitiveAccountCustomerManagedKey].
func NewCognitiveAccountCustomerManagedKey(name string, args CognitiveAccountCustomerManagedKeyArgs) *CognitiveAccountCustomerManagedKey {
	return &CognitiveAccountCustomerManagedKey{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CognitiveAccountCustomerManagedKey)(nil)

// CognitiveAccountCustomerManagedKey represents the Terraform resource azurerm_cognitive_account_customer_managed_key.
type CognitiveAccountCustomerManagedKey struct {
	Name      string
	Args      CognitiveAccountCustomerManagedKeyArgs
	state     *cognitiveAccountCustomerManagedKeyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CognitiveAccountCustomerManagedKey].
func (cacmk *CognitiveAccountCustomerManagedKey) Type() string {
	return "azurerm_cognitive_account_customer_managed_key"
}

// LocalName returns the local name for [CognitiveAccountCustomerManagedKey].
func (cacmk *CognitiveAccountCustomerManagedKey) LocalName() string {
	return cacmk.Name
}

// Configuration returns the configuration (args) for [CognitiveAccountCustomerManagedKey].
func (cacmk *CognitiveAccountCustomerManagedKey) Configuration() interface{} {
	return cacmk.Args
}

// DependOn is used for other resources to depend on [CognitiveAccountCustomerManagedKey].
func (cacmk *CognitiveAccountCustomerManagedKey) DependOn() terra.Reference {
	return terra.ReferenceResource(cacmk)
}

// Dependencies returns the list of resources [CognitiveAccountCustomerManagedKey] depends_on.
func (cacmk *CognitiveAccountCustomerManagedKey) Dependencies() terra.Dependencies {
	return cacmk.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CognitiveAccountCustomerManagedKey].
func (cacmk *CognitiveAccountCustomerManagedKey) LifecycleManagement() *terra.Lifecycle {
	return cacmk.Lifecycle
}

// Attributes returns the attributes for [CognitiveAccountCustomerManagedKey].
func (cacmk *CognitiveAccountCustomerManagedKey) Attributes() cognitiveAccountCustomerManagedKeyAttributes {
	return cognitiveAccountCustomerManagedKeyAttributes{ref: terra.ReferenceResource(cacmk)}
}

// ImportState imports the given attribute values into [CognitiveAccountCustomerManagedKey]'s state.
func (cacmk *CognitiveAccountCustomerManagedKey) ImportState(av io.Reader) error {
	cacmk.state = &cognitiveAccountCustomerManagedKeyState{}
	if err := json.NewDecoder(av).Decode(cacmk.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cacmk.Type(), cacmk.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CognitiveAccountCustomerManagedKey] has state.
func (cacmk *CognitiveAccountCustomerManagedKey) State() (*cognitiveAccountCustomerManagedKeyState, bool) {
	return cacmk.state, cacmk.state != nil
}

// StateMust returns the state for [CognitiveAccountCustomerManagedKey]. Panics if the state is nil.
func (cacmk *CognitiveAccountCustomerManagedKey) StateMust() *cognitiveAccountCustomerManagedKeyState {
	if cacmk.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cacmk.Type(), cacmk.LocalName()))
	}
	return cacmk.state
}

// CognitiveAccountCustomerManagedKeyArgs contains the configurations for azurerm_cognitive_account_customer_managed_key.
type CognitiveAccountCustomerManagedKeyArgs struct {
	// CognitiveAccountId: string, required
	CognitiveAccountId terra.StringValue `hcl:"cognitive_account_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IdentityClientId: string, optional
	IdentityClientId terra.StringValue `hcl:"identity_client_id,attr"`
	// KeyVaultKeyId: string, required
	KeyVaultKeyId terra.StringValue `hcl:"key_vault_key_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *cognitiveaccountcustomermanagedkey.Timeouts `hcl:"timeouts,block"`
}
type cognitiveAccountCustomerManagedKeyAttributes struct {
	ref terra.Reference
}

// CognitiveAccountId returns a reference to field cognitive_account_id of azurerm_cognitive_account_customer_managed_key.
func (cacmk cognitiveAccountCustomerManagedKeyAttributes) CognitiveAccountId() terra.StringValue {
	return terra.ReferenceAsString(cacmk.ref.Append("cognitive_account_id"))
}

// Id returns a reference to field id of azurerm_cognitive_account_customer_managed_key.
func (cacmk cognitiveAccountCustomerManagedKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cacmk.ref.Append("id"))
}

// IdentityClientId returns a reference to field identity_client_id of azurerm_cognitive_account_customer_managed_key.
func (cacmk cognitiveAccountCustomerManagedKeyAttributes) IdentityClientId() terra.StringValue {
	return terra.ReferenceAsString(cacmk.ref.Append("identity_client_id"))
}

// KeyVaultKeyId returns a reference to field key_vault_key_id of azurerm_cognitive_account_customer_managed_key.
func (cacmk cognitiveAccountCustomerManagedKeyAttributes) KeyVaultKeyId() terra.StringValue {
	return terra.ReferenceAsString(cacmk.ref.Append("key_vault_key_id"))
}

func (cacmk cognitiveAccountCustomerManagedKeyAttributes) Timeouts() cognitiveaccountcustomermanagedkey.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cognitiveaccountcustomermanagedkey.TimeoutsAttributes](cacmk.ref.Append("timeouts"))
}

type cognitiveAccountCustomerManagedKeyState struct {
	CognitiveAccountId string                                            `json:"cognitive_account_id"`
	Id                 string                                            `json:"id"`
	IdentityClientId   string                                            `json:"identity_client_id"`
	KeyVaultKeyId      string                                            `json:"key_vault_key_id"`
	Timeouts           *cognitiveaccountcustomermanagedkey.TimeoutsState `json:"timeouts"`
}
