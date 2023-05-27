// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	kustoclustercustomermanagedkey "github.com/golingon/terraproviders/azurerm/3.58.0/kustoclustercustomermanagedkey"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKustoClusterCustomerManagedKey creates a new instance of [KustoClusterCustomerManagedKey].
func NewKustoClusterCustomerManagedKey(name string, args KustoClusterCustomerManagedKeyArgs) *KustoClusterCustomerManagedKey {
	return &KustoClusterCustomerManagedKey{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KustoClusterCustomerManagedKey)(nil)

// KustoClusterCustomerManagedKey represents the Terraform resource azurerm_kusto_cluster_customer_managed_key.
type KustoClusterCustomerManagedKey struct {
	Name      string
	Args      KustoClusterCustomerManagedKeyArgs
	state     *kustoClusterCustomerManagedKeyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KustoClusterCustomerManagedKey].
func (kccmk *KustoClusterCustomerManagedKey) Type() string {
	return "azurerm_kusto_cluster_customer_managed_key"
}

// LocalName returns the local name for [KustoClusterCustomerManagedKey].
func (kccmk *KustoClusterCustomerManagedKey) LocalName() string {
	return kccmk.Name
}

// Configuration returns the configuration (args) for [KustoClusterCustomerManagedKey].
func (kccmk *KustoClusterCustomerManagedKey) Configuration() interface{} {
	return kccmk.Args
}

// DependOn is used for other resources to depend on [KustoClusterCustomerManagedKey].
func (kccmk *KustoClusterCustomerManagedKey) DependOn() terra.Reference {
	return terra.ReferenceResource(kccmk)
}

// Dependencies returns the list of resources [KustoClusterCustomerManagedKey] depends_on.
func (kccmk *KustoClusterCustomerManagedKey) Dependencies() terra.Dependencies {
	return kccmk.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KustoClusterCustomerManagedKey].
func (kccmk *KustoClusterCustomerManagedKey) LifecycleManagement() *terra.Lifecycle {
	return kccmk.Lifecycle
}

// Attributes returns the attributes for [KustoClusterCustomerManagedKey].
func (kccmk *KustoClusterCustomerManagedKey) Attributes() kustoClusterCustomerManagedKeyAttributes {
	return kustoClusterCustomerManagedKeyAttributes{ref: terra.ReferenceResource(kccmk)}
}

// ImportState imports the given attribute values into [KustoClusterCustomerManagedKey]'s state.
func (kccmk *KustoClusterCustomerManagedKey) ImportState(av io.Reader) error {
	kccmk.state = &kustoClusterCustomerManagedKeyState{}
	if err := json.NewDecoder(av).Decode(kccmk.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kccmk.Type(), kccmk.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KustoClusterCustomerManagedKey] has state.
func (kccmk *KustoClusterCustomerManagedKey) State() (*kustoClusterCustomerManagedKeyState, bool) {
	return kccmk.state, kccmk.state != nil
}

// StateMust returns the state for [KustoClusterCustomerManagedKey]. Panics if the state is nil.
func (kccmk *KustoClusterCustomerManagedKey) StateMust() *kustoClusterCustomerManagedKeyState {
	if kccmk.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kccmk.Type(), kccmk.LocalName()))
	}
	return kccmk.state
}

// KustoClusterCustomerManagedKeyArgs contains the configurations for azurerm_kusto_cluster_customer_managed_key.
type KustoClusterCustomerManagedKeyArgs struct {
	// ClusterId: string, required
	ClusterId terra.StringValue `hcl:"cluster_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyName: string, required
	KeyName terra.StringValue `hcl:"key_name,attr" validate:"required"`
	// KeyVaultId: string, required
	KeyVaultId terra.StringValue `hcl:"key_vault_id,attr" validate:"required"`
	// KeyVersion: string, optional
	KeyVersion terra.StringValue `hcl:"key_version,attr"`
	// UserIdentity: string, optional
	UserIdentity terra.StringValue `hcl:"user_identity,attr"`
	// Timeouts: optional
	Timeouts *kustoclustercustomermanagedkey.Timeouts `hcl:"timeouts,block"`
}
type kustoClusterCustomerManagedKeyAttributes struct {
	ref terra.Reference
}

// ClusterId returns a reference to field cluster_id of azurerm_kusto_cluster_customer_managed_key.
func (kccmk kustoClusterCustomerManagedKeyAttributes) ClusterId() terra.StringValue {
	return terra.ReferenceAsString(kccmk.ref.Append("cluster_id"))
}

// Id returns a reference to field id of azurerm_kusto_cluster_customer_managed_key.
func (kccmk kustoClusterCustomerManagedKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kccmk.ref.Append("id"))
}

// KeyName returns a reference to field key_name of azurerm_kusto_cluster_customer_managed_key.
func (kccmk kustoClusterCustomerManagedKeyAttributes) KeyName() terra.StringValue {
	return terra.ReferenceAsString(kccmk.ref.Append("key_name"))
}

// KeyVaultId returns a reference to field key_vault_id of azurerm_kusto_cluster_customer_managed_key.
func (kccmk kustoClusterCustomerManagedKeyAttributes) KeyVaultId() terra.StringValue {
	return terra.ReferenceAsString(kccmk.ref.Append("key_vault_id"))
}

// KeyVersion returns a reference to field key_version of azurerm_kusto_cluster_customer_managed_key.
func (kccmk kustoClusterCustomerManagedKeyAttributes) KeyVersion() terra.StringValue {
	return terra.ReferenceAsString(kccmk.ref.Append("key_version"))
}

// UserIdentity returns a reference to field user_identity of azurerm_kusto_cluster_customer_managed_key.
func (kccmk kustoClusterCustomerManagedKeyAttributes) UserIdentity() terra.StringValue {
	return terra.ReferenceAsString(kccmk.ref.Append("user_identity"))
}

func (kccmk kustoClusterCustomerManagedKeyAttributes) Timeouts() kustoclustercustomermanagedkey.TimeoutsAttributes {
	return terra.ReferenceAsSingle[kustoclustercustomermanagedkey.TimeoutsAttributes](kccmk.ref.Append("timeouts"))
}

type kustoClusterCustomerManagedKeyState struct {
	ClusterId    string                                        `json:"cluster_id"`
	Id           string                                        `json:"id"`
	KeyName      string                                        `json:"key_name"`
	KeyVaultId   string                                        `json:"key_vault_id"`
	KeyVersion   string                                        `json:"key_version"`
	UserIdentity string                                        `json:"user_identity"`
	Timeouts     *kustoclustercustomermanagedkey.TimeoutsState `json:"timeouts"`
}
