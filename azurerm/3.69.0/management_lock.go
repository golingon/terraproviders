// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	managementlock "github.com/golingon/terraproviders/azurerm/3.69.0/managementlock"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewManagementLock creates a new instance of [ManagementLock].
func NewManagementLock(name string, args ManagementLockArgs) *ManagementLock {
	return &ManagementLock{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ManagementLock)(nil)

// ManagementLock represents the Terraform resource azurerm_management_lock.
type ManagementLock struct {
	Name      string
	Args      ManagementLockArgs
	state     *managementLockState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ManagementLock].
func (ml *ManagementLock) Type() string {
	return "azurerm_management_lock"
}

// LocalName returns the local name for [ManagementLock].
func (ml *ManagementLock) LocalName() string {
	return ml.Name
}

// Configuration returns the configuration (args) for [ManagementLock].
func (ml *ManagementLock) Configuration() interface{} {
	return ml.Args
}

// DependOn is used for other resources to depend on [ManagementLock].
func (ml *ManagementLock) DependOn() terra.Reference {
	return terra.ReferenceResource(ml)
}

// Dependencies returns the list of resources [ManagementLock] depends_on.
func (ml *ManagementLock) Dependencies() terra.Dependencies {
	return ml.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ManagementLock].
func (ml *ManagementLock) LifecycleManagement() *terra.Lifecycle {
	return ml.Lifecycle
}

// Attributes returns the attributes for [ManagementLock].
func (ml *ManagementLock) Attributes() managementLockAttributes {
	return managementLockAttributes{ref: terra.ReferenceResource(ml)}
}

// ImportState imports the given attribute values into [ManagementLock]'s state.
func (ml *ManagementLock) ImportState(av io.Reader) error {
	ml.state = &managementLockState{}
	if err := json.NewDecoder(av).Decode(ml.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ml.Type(), ml.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ManagementLock] has state.
func (ml *ManagementLock) State() (*managementLockState, bool) {
	return ml.state, ml.state != nil
}

// StateMust returns the state for [ManagementLock]. Panics if the state is nil.
func (ml *ManagementLock) StateMust() *managementLockState {
	if ml.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ml.Type(), ml.LocalName()))
	}
	return ml.state
}

// ManagementLockArgs contains the configurations for azurerm_management_lock.
type ManagementLockArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LockLevel: string, required
	LockLevel terra.StringValue `hcl:"lock_level,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Notes: string, optional
	Notes terra.StringValue `hcl:"notes,attr"`
	// Scope: string, required
	Scope terra.StringValue `hcl:"scope,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *managementlock.Timeouts `hcl:"timeouts,block"`
}
type managementLockAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_management_lock.
func (ml managementLockAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ml.ref.Append("id"))
}

// LockLevel returns a reference to field lock_level of azurerm_management_lock.
func (ml managementLockAttributes) LockLevel() terra.StringValue {
	return terra.ReferenceAsString(ml.ref.Append("lock_level"))
}

// Name returns a reference to field name of azurerm_management_lock.
func (ml managementLockAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ml.ref.Append("name"))
}

// Notes returns a reference to field notes of azurerm_management_lock.
func (ml managementLockAttributes) Notes() terra.StringValue {
	return terra.ReferenceAsString(ml.ref.Append("notes"))
}

// Scope returns a reference to field scope of azurerm_management_lock.
func (ml managementLockAttributes) Scope() terra.StringValue {
	return terra.ReferenceAsString(ml.ref.Append("scope"))
}

func (ml managementLockAttributes) Timeouts() managementlock.TimeoutsAttributes {
	return terra.ReferenceAsSingle[managementlock.TimeoutsAttributes](ml.ref.Append("timeouts"))
}

type managementLockState struct {
	Id        string                        `json:"id"`
	LockLevel string                        `json:"lock_level"`
	Name      string                        `json:"name"`
	Notes     string                        `json:"notes"`
	Scope     string                        `json:"scope"`
	Timeouts  *managementlock.TimeoutsState `json:"timeouts"`
}
