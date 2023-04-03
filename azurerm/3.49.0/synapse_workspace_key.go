// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	synapseworkspacekey "github.com/golingon/terraproviders/azurerm/3.49.0/synapseworkspacekey"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSynapseWorkspaceKey creates a new instance of [SynapseWorkspaceKey].
func NewSynapseWorkspaceKey(name string, args SynapseWorkspaceKeyArgs) *SynapseWorkspaceKey {
	return &SynapseWorkspaceKey{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SynapseWorkspaceKey)(nil)

// SynapseWorkspaceKey represents the Terraform resource azurerm_synapse_workspace_key.
type SynapseWorkspaceKey struct {
	Name      string
	Args      SynapseWorkspaceKeyArgs
	state     *synapseWorkspaceKeyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SynapseWorkspaceKey].
func (swk *SynapseWorkspaceKey) Type() string {
	return "azurerm_synapse_workspace_key"
}

// LocalName returns the local name for [SynapseWorkspaceKey].
func (swk *SynapseWorkspaceKey) LocalName() string {
	return swk.Name
}

// Configuration returns the configuration (args) for [SynapseWorkspaceKey].
func (swk *SynapseWorkspaceKey) Configuration() interface{} {
	return swk.Args
}

// DependOn is used for other resources to depend on [SynapseWorkspaceKey].
func (swk *SynapseWorkspaceKey) DependOn() terra.Reference {
	return terra.ReferenceResource(swk)
}

// Dependencies returns the list of resources [SynapseWorkspaceKey] depends_on.
func (swk *SynapseWorkspaceKey) Dependencies() terra.Dependencies {
	return swk.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SynapseWorkspaceKey].
func (swk *SynapseWorkspaceKey) LifecycleManagement() *terra.Lifecycle {
	return swk.Lifecycle
}

// Attributes returns the attributes for [SynapseWorkspaceKey].
func (swk *SynapseWorkspaceKey) Attributes() synapseWorkspaceKeyAttributes {
	return synapseWorkspaceKeyAttributes{ref: terra.ReferenceResource(swk)}
}

// ImportState imports the given attribute values into [SynapseWorkspaceKey]'s state.
func (swk *SynapseWorkspaceKey) ImportState(av io.Reader) error {
	swk.state = &synapseWorkspaceKeyState{}
	if err := json.NewDecoder(av).Decode(swk.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", swk.Type(), swk.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SynapseWorkspaceKey] has state.
func (swk *SynapseWorkspaceKey) State() (*synapseWorkspaceKeyState, bool) {
	return swk.state, swk.state != nil
}

// StateMust returns the state for [SynapseWorkspaceKey]. Panics if the state is nil.
func (swk *SynapseWorkspaceKey) StateMust() *synapseWorkspaceKeyState {
	if swk.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", swk.Type(), swk.LocalName()))
	}
	return swk.state
}

// SynapseWorkspaceKeyArgs contains the configurations for azurerm_synapse_workspace_key.
type SynapseWorkspaceKeyArgs struct {
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
	Timeouts *synapseworkspacekey.Timeouts `hcl:"timeouts,block"`
}
type synapseWorkspaceKeyAttributes struct {
	ref terra.Reference
}

// Active returns a reference to field active of azurerm_synapse_workspace_key.
func (swk synapseWorkspaceKeyAttributes) Active() terra.BoolValue {
	return terra.ReferenceAsBool(swk.ref.Append("active"))
}

// CustomerManagedKeyName returns a reference to field customer_managed_key_name of azurerm_synapse_workspace_key.
func (swk synapseWorkspaceKeyAttributes) CustomerManagedKeyName() terra.StringValue {
	return terra.ReferenceAsString(swk.ref.Append("customer_managed_key_name"))
}

// CustomerManagedKeyVersionlessId returns a reference to field customer_managed_key_versionless_id of azurerm_synapse_workspace_key.
func (swk synapseWorkspaceKeyAttributes) CustomerManagedKeyVersionlessId() terra.StringValue {
	return terra.ReferenceAsString(swk.ref.Append("customer_managed_key_versionless_id"))
}

// Id returns a reference to field id of azurerm_synapse_workspace_key.
func (swk synapseWorkspaceKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(swk.ref.Append("id"))
}

// SynapseWorkspaceId returns a reference to field synapse_workspace_id of azurerm_synapse_workspace_key.
func (swk synapseWorkspaceKeyAttributes) SynapseWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(swk.ref.Append("synapse_workspace_id"))
}

func (swk synapseWorkspaceKeyAttributes) Timeouts() synapseworkspacekey.TimeoutsAttributes {
	return terra.ReferenceAsSingle[synapseworkspacekey.TimeoutsAttributes](swk.ref.Append("timeouts"))
}

type synapseWorkspaceKeyState struct {
	Active                          bool                               `json:"active"`
	CustomerManagedKeyName          string                             `json:"customer_managed_key_name"`
	CustomerManagedKeyVersionlessId string                             `json:"customer_managed_key_versionless_id"`
	Id                              string                             `json:"id"`
	SynapseWorkspaceId              string                             `json:"synapse_workspace_id"`
	Timeouts                        *synapseworkspacekey.TimeoutsState `json:"timeouts"`
}
