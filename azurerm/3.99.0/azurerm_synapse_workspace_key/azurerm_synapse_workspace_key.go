// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_synapse_workspace_key

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

// Resource represents the Terraform resource azurerm_synapse_workspace_key.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermSynapseWorkspaceKeyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aswk *Resource) Type() string {
	return "azurerm_synapse_workspace_key"
}

// LocalName returns the local name for [Resource].
func (aswk *Resource) LocalName() string {
	return aswk.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aswk *Resource) Configuration() interface{} {
	return aswk.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aswk *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aswk)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aswk *Resource) Dependencies() terra.Dependencies {
	return aswk.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aswk *Resource) LifecycleManagement() *terra.Lifecycle {
	return aswk.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aswk *Resource) Attributes() azurermSynapseWorkspaceKeyAttributes {
	return azurermSynapseWorkspaceKeyAttributes{ref: terra.ReferenceResource(aswk)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aswk *Resource) ImportState(state io.Reader) error {
	aswk.state = &azurermSynapseWorkspaceKeyState{}
	if err := json.NewDecoder(state).Decode(aswk.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aswk.Type(), aswk.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aswk *Resource) State() (*azurermSynapseWorkspaceKeyState, bool) {
	return aswk.state, aswk.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aswk *Resource) StateMust() *azurermSynapseWorkspaceKeyState {
	if aswk.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aswk.Type(), aswk.LocalName()))
	}
	return aswk.state
}

// Args contains the configurations for azurerm_synapse_workspace_key.
type Args struct {
	// Active: bool, required
	Active terra.BoolValue `hcl:"active,attr" validate:"required"`
	// CustomerManagedKeyName: string, required
	CustomerManagedKeyName terra.StringValue `hcl:"customer_managed_key_name,attr" validate:"required"`
	// CustomerManagedKeyVersionlessId: string, optional
	CustomerManagedKeyVersionlessId terra.StringValue `hcl:"customer_managed_key_versionless_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SynapseWorkspaceId: string, required
	SynapseWorkspaceId terra.StringValue `hcl:"synapse_workspace_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermSynapseWorkspaceKeyAttributes struct {
	ref terra.Reference
}

// Active returns a reference to field active of azurerm_synapse_workspace_key.
func (aswk azurermSynapseWorkspaceKeyAttributes) Active() terra.BoolValue {
	return terra.ReferenceAsBool(aswk.ref.Append("active"))
}

// CustomerManagedKeyName returns a reference to field customer_managed_key_name of azurerm_synapse_workspace_key.
func (aswk azurermSynapseWorkspaceKeyAttributes) CustomerManagedKeyName() terra.StringValue {
	return terra.ReferenceAsString(aswk.ref.Append("customer_managed_key_name"))
}

// CustomerManagedKeyVersionlessId returns a reference to field customer_managed_key_versionless_id of azurerm_synapse_workspace_key.
func (aswk azurermSynapseWorkspaceKeyAttributes) CustomerManagedKeyVersionlessId() terra.StringValue {
	return terra.ReferenceAsString(aswk.ref.Append("customer_managed_key_versionless_id"))
}

// Id returns a reference to field id of azurerm_synapse_workspace_key.
func (aswk azurermSynapseWorkspaceKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aswk.ref.Append("id"))
}

// SynapseWorkspaceId returns a reference to field synapse_workspace_id of azurerm_synapse_workspace_key.
func (aswk azurermSynapseWorkspaceKeyAttributes) SynapseWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(aswk.ref.Append("synapse_workspace_id"))
}

func (aswk azurermSynapseWorkspaceKeyAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](aswk.ref.Append("timeouts"))
}

type azurermSynapseWorkspaceKeyState struct {
	Active                          bool           `json:"active"`
	CustomerManagedKeyName          string         `json:"customer_managed_key_name"`
	CustomerManagedKeyVersionlessId string         `json:"customer_managed_key_versionless_id"`
	Id                              string         `json:"id"`
	SynapseWorkspaceId              string         `json:"synapse_workspace_id"`
	Timeouts                        *TimeoutsState `json:"timeouts"`
}
