// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	synapseintegrationruntimeazure "github.com/golingon/terraproviders/azurerm/3.98.0/synapseintegrationruntimeazure"
	"io"
)

// NewSynapseIntegrationRuntimeAzure creates a new instance of [SynapseIntegrationRuntimeAzure].
func NewSynapseIntegrationRuntimeAzure(name string, args SynapseIntegrationRuntimeAzureArgs) *SynapseIntegrationRuntimeAzure {
	return &SynapseIntegrationRuntimeAzure{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SynapseIntegrationRuntimeAzure)(nil)

// SynapseIntegrationRuntimeAzure represents the Terraform resource azurerm_synapse_integration_runtime_azure.
type SynapseIntegrationRuntimeAzure struct {
	Name      string
	Args      SynapseIntegrationRuntimeAzureArgs
	state     *synapseIntegrationRuntimeAzureState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SynapseIntegrationRuntimeAzure].
func (sira *SynapseIntegrationRuntimeAzure) Type() string {
	return "azurerm_synapse_integration_runtime_azure"
}

// LocalName returns the local name for [SynapseIntegrationRuntimeAzure].
func (sira *SynapseIntegrationRuntimeAzure) LocalName() string {
	return sira.Name
}

// Configuration returns the configuration (args) for [SynapseIntegrationRuntimeAzure].
func (sira *SynapseIntegrationRuntimeAzure) Configuration() interface{} {
	return sira.Args
}

// DependOn is used for other resources to depend on [SynapseIntegrationRuntimeAzure].
func (sira *SynapseIntegrationRuntimeAzure) DependOn() terra.Reference {
	return terra.ReferenceResource(sira)
}

// Dependencies returns the list of resources [SynapseIntegrationRuntimeAzure] depends_on.
func (sira *SynapseIntegrationRuntimeAzure) Dependencies() terra.Dependencies {
	return sira.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SynapseIntegrationRuntimeAzure].
func (sira *SynapseIntegrationRuntimeAzure) LifecycleManagement() *terra.Lifecycle {
	return sira.Lifecycle
}

// Attributes returns the attributes for [SynapseIntegrationRuntimeAzure].
func (sira *SynapseIntegrationRuntimeAzure) Attributes() synapseIntegrationRuntimeAzureAttributes {
	return synapseIntegrationRuntimeAzureAttributes{ref: terra.ReferenceResource(sira)}
}

// ImportState imports the given attribute values into [SynapseIntegrationRuntimeAzure]'s state.
func (sira *SynapseIntegrationRuntimeAzure) ImportState(av io.Reader) error {
	sira.state = &synapseIntegrationRuntimeAzureState{}
	if err := json.NewDecoder(av).Decode(sira.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sira.Type(), sira.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SynapseIntegrationRuntimeAzure] has state.
func (sira *SynapseIntegrationRuntimeAzure) State() (*synapseIntegrationRuntimeAzureState, bool) {
	return sira.state, sira.state != nil
}

// StateMust returns the state for [SynapseIntegrationRuntimeAzure]. Panics if the state is nil.
func (sira *SynapseIntegrationRuntimeAzure) StateMust() *synapseIntegrationRuntimeAzureState {
	if sira.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sira.Type(), sira.LocalName()))
	}
	return sira.state
}

// SynapseIntegrationRuntimeAzureArgs contains the configurations for azurerm_synapse_integration_runtime_azure.
type SynapseIntegrationRuntimeAzureArgs struct {
	// ComputeType: string, optional
	ComputeType terra.StringValue `hcl:"compute_type,attr"`
	// CoreCount: number, optional
	CoreCount terra.NumberValue `hcl:"core_count,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SynapseWorkspaceId: string, required
	SynapseWorkspaceId terra.StringValue `hcl:"synapse_workspace_id,attr" validate:"required"`
	// TimeToLiveMin: number, optional
	TimeToLiveMin terra.NumberValue `hcl:"time_to_live_min,attr"`
	// Timeouts: optional
	Timeouts *synapseintegrationruntimeazure.Timeouts `hcl:"timeouts,block"`
}
type synapseIntegrationRuntimeAzureAttributes struct {
	ref terra.Reference
}

// ComputeType returns a reference to field compute_type of azurerm_synapse_integration_runtime_azure.
func (sira synapseIntegrationRuntimeAzureAttributes) ComputeType() terra.StringValue {
	return terra.ReferenceAsString(sira.ref.Append("compute_type"))
}

// CoreCount returns a reference to field core_count of azurerm_synapse_integration_runtime_azure.
func (sira synapseIntegrationRuntimeAzureAttributes) CoreCount() terra.NumberValue {
	return terra.ReferenceAsNumber(sira.ref.Append("core_count"))
}

// Description returns a reference to field description of azurerm_synapse_integration_runtime_azure.
func (sira synapseIntegrationRuntimeAzureAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(sira.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_synapse_integration_runtime_azure.
func (sira synapseIntegrationRuntimeAzureAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sira.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_synapse_integration_runtime_azure.
func (sira synapseIntegrationRuntimeAzureAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(sira.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_synapse_integration_runtime_azure.
func (sira synapseIntegrationRuntimeAzureAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sira.ref.Append("name"))
}

// SynapseWorkspaceId returns a reference to field synapse_workspace_id of azurerm_synapse_integration_runtime_azure.
func (sira synapseIntegrationRuntimeAzureAttributes) SynapseWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(sira.ref.Append("synapse_workspace_id"))
}

// TimeToLiveMin returns a reference to field time_to_live_min of azurerm_synapse_integration_runtime_azure.
func (sira synapseIntegrationRuntimeAzureAttributes) TimeToLiveMin() terra.NumberValue {
	return terra.ReferenceAsNumber(sira.ref.Append("time_to_live_min"))
}

func (sira synapseIntegrationRuntimeAzureAttributes) Timeouts() synapseintegrationruntimeazure.TimeoutsAttributes {
	return terra.ReferenceAsSingle[synapseintegrationruntimeazure.TimeoutsAttributes](sira.ref.Append("timeouts"))
}

type synapseIntegrationRuntimeAzureState struct {
	ComputeType        string                                        `json:"compute_type"`
	CoreCount          float64                                       `json:"core_count"`
	Description        string                                        `json:"description"`
	Id                 string                                        `json:"id"`
	Location           string                                        `json:"location"`
	Name               string                                        `json:"name"`
	SynapseWorkspaceId string                                        `json:"synapse_workspace_id"`
	TimeToLiveMin      float64                                       `json:"time_to_live_min"`
	Timeouts           *synapseintegrationruntimeazure.TimeoutsState `json:"timeouts"`
}
