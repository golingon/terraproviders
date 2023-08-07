// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	storagemanagementpolicy "github.com/golingon/terraproviders/azurerm/3.68.0/storagemanagementpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStorageManagementPolicy creates a new instance of [StorageManagementPolicy].
func NewStorageManagementPolicy(name string, args StorageManagementPolicyArgs) *StorageManagementPolicy {
	return &StorageManagementPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StorageManagementPolicy)(nil)

// StorageManagementPolicy represents the Terraform resource azurerm_storage_management_policy.
type StorageManagementPolicy struct {
	Name      string
	Args      StorageManagementPolicyArgs
	state     *storageManagementPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StorageManagementPolicy].
func (smp *StorageManagementPolicy) Type() string {
	return "azurerm_storage_management_policy"
}

// LocalName returns the local name for [StorageManagementPolicy].
func (smp *StorageManagementPolicy) LocalName() string {
	return smp.Name
}

// Configuration returns the configuration (args) for [StorageManagementPolicy].
func (smp *StorageManagementPolicy) Configuration() interface{} {
	return smp.Args
}

// DependOn is used for other resources to depend on [StorageManagementPolicy].
func (smp *StorageManagementPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(smp)
}

// Dependencies returns the list of resources [StorageManagementPolicy] depends_on.
func (smp *StorageManagementPolicy) Dependencies() terra.Dependencies {
	return smp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StorageManagementPolicy].
func (smp *StorageManagementPolicy) LifecycleManagement() *terra.Lifecycle {
	return smp.Lifecycle
}

// Attributes returns the attributes for [StorageManagementPolicy].
func (smp *StorageManagementPolicy) Attributes() storageManagementPolicyAttributes {
	return storageManagementPolicyAttributes{ref: terra.ReferenceResource(smp)}
}

// ImportState imports the given attribute values into [StorageManagementPolicy]'s state.
func (smp *StorageManagementPolicy) ImportState(av io.Reader) error {
	smp.state = &storageManagementPolicyState{}
	if err := json.NewDecoder(av).Decode(smp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", smp.Type(), smp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StorageManagementPolicy] has state.
func (smp *StorageManagementPolicy) State() (*storageManagementPolicyState, bool) {
	return smp.state, smp.state != nil
}

// StateMust returns the state for [StorageManagementPolicy]. Panics if the state is nil.
func (smp *StorageManagementPolicy) StateMust() *storageManagementPolicyState {
	if smp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", smp.Type(), smp.LocalName()))
	}
	return smp.state
}

// StorageManagementPolicyArgs contains the configurations for azurerm_storage_management_policy.
type StorageManagementPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// StorageAccountId: string, required
	StorageAccountId terra.StringValue `hcl:"storage_account_id,attr" validate:"required"`
	// Rule: min=0
	Rule []storagemanagementpolicy.Rule `hcl:"rule,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *storagemanagementpolicy.Timeouts `hcl:"timeouts,block"`
}
type storageManagementPolicyAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_storage_management_policy.
func (smp storageManagementPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(smp.ref.Append("id"))
}

// StorageAccountId returns a reference to field storage_account_id of azurerm_storage_management_policy.
func (smp storageManagementPolicyAttributes) StorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(smp.ref.Append("storage_account_id"))
}

func (smp storageManagementPolicyAttributes) Rule() terra.ListValue[storagemanagementpolicy.RuleAttributes] {
	return terra.ReferenceAsList[storagemanagementpolicy.RuleAttributes](smp.ref.Append("rule"))
}

func (smp storageManagementPolicyAttributes) Timeouts() storagemanagementpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[storagemanagementpolicy.TimeoutsAttributes](smp.ref.Append("timeouts"))
}

type storageManagementPolicyState struct {
	Id               string                                 `json:"id"`
	StorageAccountId string                                 `json:"storage_account_id"`
	Rule             []storagemanagementpolicy.RuleState    `json:"rule"`
	Timeouts         *storagemanagementpolicy.TimeoutsState `json:"timeouts"`
}
