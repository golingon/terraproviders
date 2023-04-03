// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	storageblobinventorypolicy "github.com/golingon/terraproviders/azurerm/3.49.0/storageblobinventorypolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStorageBlobInventoryPolicy creates a new instance of [StorageBlobInventoryPolicy].
func NewStorageBlobInventoryPolicy(name string, args StorageBlobInventoryPolicyArgs) *StorageBlobInventoryPolicy {
	return &StorageBlobInventoryPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StorageBlobInventoryPolicy)(nil)

// StorageBlobInventoryPolicy represents the Terraform resource azurerm_storage_blob_inventory_policy.
type StorageBlobInventoryPolicy struct {
	Name      string
	Args      StorageBlobInventoryPolicyArgs
	state     *storageBlobInventoryPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StorageBlobInventoryPolicy].
func (sbip *StorageBlobInventoryPolicy) Type() string {
	return "azurerm_storage_blob_inventory_policy"
}

// LocalName returns the local name for [StorageBlobInventoryPolicy].
func (sbip *StorageBlobInventoryPolicy) LocalName() string {
	return sbip.Name
}

// Configuration returns the configuration (args) for [StorageBlobInventoryPolicy].
func (sbip *StorageBlobInventoryPolicy) Configuration() interface{} {
	return sbip.Args
}

// DependOn is used for other resources to depend on [StorageBlobInventoryPolicy].
func (sbip *StorageBlobInventoryPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(sbip)
}

// Dependencies returns the list of resources [StorageBlobInventoryPolicy] depends_on.
func (sbip *StorageBlobInventoryPolicy) Dependencies() terra.Dependencies {
	return sbip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StorageBlobInventoryPolicy].
func (sbip *StorageBlobInventoryPolicy) LifecycleManagement() *terra.Lifecycle {
	return sbip.Lifecycle
}

// Attributes returns the attributes for [StorageBlobInventoryPolicy].
func (sbip *StorageBlobInventoryPolicy) Attributes() storageBlobInventoryPolicyAttributes {
	return storageBlobInventoryPolicyAttributes{ref: terra.ReferenceResource(sbip)}
}

// ImportState imports the given attribute values into [StorageBlobInventoryPolicy]'s state.
func (sbip *StorageBlobInventoryPolicy) ImportState(av io.Reader) error {
	sbip.state = &storageBlobInventoryPolicyState{}
	if err := json.NewDecoder(av).Decode(sbip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sbip.Type(), sbip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StorageBlobInventoryPolicy] has state.
func (sbip *StorageBlobInventoryPolicy) State() (*storageBlobInventoryPolicyState, bool) {
	return sbip.state, sbip.state != nil
}

// StateMust returns the state for [StorageBlobInventoryPolicy]. Panics if the state is nil.
func (sbip *StorageBlobInventoryPolicy) StateMust() *storageBlobInventoryPolicyState {
	if sbip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sbip.Type(), sbip.LocalName()))
	}
	return sbip.state
}

// StorageBlobInventoryPolicyArgs contains the configurations for azurerm_storage_blob_inventory_policy.
type StorageBlobInventoryPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// StorageAccountId: string, required
	StorageAccountId terra.StringValue `hcl:"storage_account_id,attr" validate:"required"`
	// Rules: min=1
	Rules []storageblobinventorypolicy.Rules `hcl:"rules,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *storageblobinventorypolicy.Timeouts `hcl:"timeouts,block"`
}
type storageBlobInventoryPolicyAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_storage_blob_inventory_policy.
func (sbip storageBlobInventoryPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sbip.ref.Append("id"))
}

// StorageAccountId returns a reference to field storage_account_id of azurerm_storage_blob_inventory_policy.
func (sbip storageBlobInventoryPolicyAttributes) StorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(sbip.ref.Append("storage_account_id"))
}

func (sbip storageBlobInventoryPolicyAttributes) Rules() terra.SetValue[storageblobinventorypolicy.RulesAttributes] {
	return terra.ReferenceAsSet[storageblobinventorypolicy.RulesAttributes](sbip.ref.Append("rules"))
}

func (sbip storageBlobInventoryPolicyAttributes) Timeouts() storageblobinventorypolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[storageblobinventorypolicy.TimeoutsAttributes](sbip.ref.Append("timeouts"))
}

type storageBlobInventoryPolicyState struct {
	Id               string                                    `json:"id"`
	StorageAccountId string                                    `json:"storage_account_id"`
	Rules            []storageblobinventorypolicy.RulesState   `json:"rules"`
	Timeouts         *storageblobinventorypolicy.TimeoutsState `json:"timeouts"`
}
