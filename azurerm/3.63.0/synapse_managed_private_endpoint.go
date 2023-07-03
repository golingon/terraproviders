// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	synapsemanagedprivateendpoint "github.com/golingon/terraproviders/azurerm/3.63.0/synapsemanagedprivateendpoint"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSynapseManagedPrivateEndpoint creates a new instance of [SynapseManagedPrivateEndpoint].
func NewSynapseManagedPrivateEndpoint(name string, args SynapseManagedPrivateEndpointArgs) *SynapseManagedPrivateEndpoint {
	return &SynapseManagedPrivateEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SynapseManagedPrivateEndpoint)(nil)

// SynapseManagedPrivateEndpoint represents the Terraform resource azurerm_synapse_managed_private_endpoint.
type SynapseManagedPrivateEndpoint struct {
	Name      string
	Args      SynapseManagedPrivateEndpointArgs
	state     *synapseManagedPrivateEndpointState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SynapseManagedPrivateEndpoint].
func (smpe *SynapseManagedPrivateEndpoint) Type() string {
	return "azurerm_synapse_managed_private_endpoint"
}

// LocalName returns the local name for [SynapseManagedPrivateEndpoint].
func (smpe *SynapseManagedPrivateEndpoint) LocalName() string {
	return smpe.Name
}

// Configuration returns the configuration (args) for [SynapseManagedPrivateEndpoint].
func (smpe *SynapseManagedPrivateEndpoint) Configuration() interface{} {
	return smpe.Args
}

// DependOn is used for other resources to depend on [SynapseManagedPrivateEndpoint].
func (smpe *SynapseManagedPrivateEndpoint) DependOn() terra.Reference {
	return terra.ReferenceResource(smpe)
}

// Dependencies returns the list of resources [SynapseManagedPrivateEndpoint] depends_on.
func (smpe *SynapseManagedPrivateEndpoint) Dependencies() terra.Dependencies {
	return smpe.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SynapseManagedPrivateEndpoint].
func (smpe *SynapseManagedPrivateEndpoint) LifecycleManagement() *terra.Lifecycle {
	return smpe.Lifecycle
}

// Attributes returns the attributes for [SynapseManagedPrivateEndpoint].
func (smpe *SynapseManagedPrivateEndpoint) Attributes() synapseManagedPrivateEndpointAttributes {
	return synapseManagedPrivateEndpointAttributes{ref: terra.ReferenceResource(smpe)}
}

// ImportState imports the given attribute values into [SynapseManagedPrivateEndpoint]'s state.
func (smpe *SynapseManagedPrivateEndpoint) ImportState(av io.Reader) error {
	smpe.state = &synapseManagedPrivateEndpointState{}
	if err := json.NewDecoder(av).Decode(smpe.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", smpe.Type(), smpe.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SynapseManagedPrivateEndpoint] has state.
func (smpe *SynapseManagedPrivateEndpoint) State() (*synapseManagedPrivateEndpointState, bool) {
	return smpe.state, smpe.state != nil
}

// StateMust returns the state for [SynapseManagedPrivateEndpoint]. Panics if the state is nil.
func (smpe *SynapseManagedPrivateEndpoint) StateMust() *synapseManagedPrivateEndpointState {
	if smpe.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", smpe.Type(), smpe.LocalName()))
	}
	return smpe.state
}

// SynapseManagedPrivateEndpointArgs contains the configurations for azurerm_synapse_managed_private_endpoint.
type SynapseManagedPrivateEndpointArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SubresourceName: string, required
	SubresourceName terra.StringValue `hcl:"subresource_name,attr" validate:"required"`
	// SynapseWorkspaceId: string, required
	SynapseWorkspaceId terra.StringValue `hcl:"synapse_workspace_id,attr" validate:"required"`
	// TargetResourceId: string, required
	TargetResourceId terra.StringValue `hcl:"target_resource_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *synapsemanagedprivateendpoint.Timeouts `hcl:"timeouts,block"`
}
type synapseManagedPrivateEndpointAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_synapse_managed_private_endpoint.
func (smpe synapseManagedPrivateEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(smpe.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_synapse_managed_private_endpoint.
func (smpe synapseManagedPrivateEndpointAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(smpe.ref.Append("name"))
}

// SubresourceName returns a reference to field subresource_name of azurerm_synapse_managed_private_endpoint.
func (smpe synapseManagedPrivateEndpointAttributes) SubresourceName() terra.StringValue {
	return terra.ReferenceAsString(smpe.ref.Append("subresource_name"))
}

// SynapseWorkspaceId returns a reference to field synapse_workspace_id of azurerm_synapse_managed_private_endpoint.
func (smpe synapseManagedPrivateEndpointAttributes) SynapseWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(smpe.ref.Append("synapse_workspace_id"))
}

// TargetResourceId returns a reference to field target_resource_id of azurerm_synapse_managed_private_endpoint.
func (smpe synapseManagedPrivateEndpointAttributes) TargetResourceId() terra.StringValue {
	return terra.ReferenceAsString(smpe.ref.Append("target_resource_id"))
}

func (smpe synapseManagedPrivateEndpointAttributes) Timeouts() synapsemanagedprivateendpoint.TimeoutsAttributes {
	return terra.ReferenceAsSingle[synapsemanagedprivateendpoint.TimeoutsAttributes](smpe.ref.Append("timeouts"))
}

type synapseManagedPrivateEndpointState struct {
	Id                 string                                       `json:"id"`
	Name               string                                       `json:"name"`
	SubresourceName    string                                       `json:"subresource_name"`
	SynapseWorkspaceId string                                       `json:"synapse_workspace_id"`
	TargetResourceId   string                                       `json:"target_resource_id"`
	Timeouts           *synapsemanagedprivateendpoint.TimeoutsState `json:"timeouts"`
}