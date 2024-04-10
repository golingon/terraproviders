// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	synapselinkedservice "github.com/golingon/terraproviders/azurerm/3.98.0/synapselinkedservice"
	"io"
)

// NewSynapseLinkedService creates a new instance of [SynapseLinkedService].
func NewSynapseLinkedService(name string, args SynapseLinkedServiceArgs) *SynapseLinkedService {
	return &SynapseLinkedService{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SynapseLinkedService)(nil)

// SynapseLinkedService represents the Terraform resource azurerm_synapse_linked_service.
type SynapseLinkedService struct {
	Name      string
	Args      SynapseLinkedServiceArgs
	state     *synapseLinkedServiceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SynapseLinkedService].
func (sls *SynapseLinkedService) Type() string {
	return "azurerm_synapse_linked_service"
}

// LocalName returns the local name for [SynapseLinkedService].
func (sls *SynapseLinkedService) LocalName() string {
	return sls.Name
}

// Configuration returns the configuration (args) for [SynapseLinkedService].
func (sls *SynapseLinkedService) Configuration() interface{} {
	return sls.Args
}

// DependOn is used for other resources to depend on [SynapseLinkedService].
func (sls *SynapseLinkedService) DependOn() terra.Reference {
	return terra.ReferenceResource(sls)
}

// Dependencies returns the list of resources [SynapseLinkedService] depends_on.
func (sls *SynapseLinkedService) Dependencies() terra.Dependencies {
	return sls.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SynapseLinkedService].
func (sls *SynapseLinkedService) LifecycleManagement() *terra.Lifecycle {
	return sls.Lifecycle
}

// Attributes returns the attributes for [SynapseLinkedService].
func (sls *SynapseLinkedService) Attributes() synapseLinkedServiceAttributes {
	return synapseLinkedServiceAttributes{ref: terra.ReferenceResource(sls)}
}

// ImportState imports the given attribute values into [SynapseLinkedService]'s state.
func (sls *SynapseLinkedService) ImportState(av io.Reader) error {
	sls.state = &synapseLinkedServiceState{}
	if err := json.NewDecoder(av).Decode(sls.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sls.Type(), sls.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SynapseLinkedService] has state.
func (sls *SynapseLinkedService) State() (*synapseLinkedServiceState, bool) {
	return sls.state, sls.state != nil
}

// StateMust returns the state for [SynapseLinkedService]. Panics if the state is nil.
func (sls *SynapseLinkedService) StateMust() *synapseLinkedServiceState {
	if sls.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sls.Type(), sls.LocalName()))
	}
	return sls.state
}

// SynapseLinkedServiceArgs contains the configurations for azurerm_synapse_linked_service.
type SynapseLinkedServiceArgs struct {
	// AdditionalProperties: map of string, optional
	AdditionalProperties terra.MapValue[terra.StringValue] `hcl:"additional_properties,attr"`
	// Annotations: list of string, optional
	Annotations terra.ListValue[terra.StringValue] `hcl:"annotations,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parameters: map of string, optional
	Parameters terra.MapValue[terra.StringValue] `hcl:"parameters,attr"`
	// SynapseWorkspaceId: string, required
	SynapseWorkspaceId terra.StringValue `hcl:"synapse_workspace_id,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// TypePropertiesJson: string, required
	TypePropertiesJson terra.StringValue `hcl:"type_properties_json,attr" validate:"required"`
	// IntegrationRuntime: optional
	IntegrationRuntime *synapselinkedservice.IntegrationRuntime `hcl:"integration_runtime,block"`
	// Timeouts: optional
	Timeouts *synapselinkedservice.Timeouts `hcl:"timeouts,block"`
}
type synapseLinkedServiceAttributes struct {
	ref terra.Reference
}

// AdditionalProperties returns a reference to field additional_properties of azurerm_synapse_linked_service.
func (sls synapseLinkedServiceAttributes) AdditionalProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sls.ref.Append("additional_properties"))
}

// Annotations returns a reference to field annotations of azurerm_synapse_linked_service.
func (sls synapseLinkedServiceAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sls.ref.Append("annotations"))
}

// Description returns a reference to field description of azurerm_synapse_linked_service.
func (sls synapseLinkedServiceAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(sls.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_synapse_linked_service.
func (sls synapseLinkedServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sls.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_synapse_linked_service.
func (sls synapseLinkedServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sls.ref.Append("name"))
}

// Parameters returns a reference to field parameters of azurerm_synapse_linked_service.
func (sls synapseLinkedServiceAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sls.ref.Append("parameters"))
}

// SynapseWorkspaceId returns a reference to field synapse_workspace_id of azurerm_synapse_linked_service.
func (sls synapseLinkedServiceAttributes) SynapseWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(sls.ref.Append("synapse_workspace_id"))
}

// Type returns a reference to field type of azurerm_synapse_linked_service.
func (sls synapseLinkedServiceAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(sls.ref.Append("type"))
}

// TypePropertiesJson returns a reference to field type_properties_json of azurerm_synapse_linked_service.
func (sls synapseLinkedServiceAttributes) TypePropertiesJson() terra.StringValue {
	return terra.ReferenceAsString(sls.ref.Append("type_properties_json"))
}

func (sls synapseLinkedServiceAttributes) IntegrationRuntime() terra.ListValue[synapselinkedservice.IntegrationRuntimeAttributes] {
	return terra.ReferenceAsList[synapselinkedservice.IntegrationRuntimeAttributes](sls.ref.Append("integration_runtime"))
}

func (sls synapseLinkedServiceAttributes) Timeouts() synapselinkedservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[synapselinkedservice.TimeoutsAttributes](sls.ref.Append("timeouts"))
}

type synapseLinkedServiceState struct {
	AdditionalProperties map[string]string                              `json:"additional_properties"`
	Annotations          []string                                       `json:"annotations"`
	Description          string                                         `json:"description"`
	Id                   string                                         `json:"id"`
	Name                 string                                         `json:"name"`
	Parameters           map[string]string                              `json:"parameters"`
	SynapseWorkspaceId   string                                         `json:"synapse_workspace_id"`
	Type                 string                                         `json:"type"`
	TypePropertiesJson   string                                         `json:"type_properties_json"`
	IntegrationRuntime   []synapselinkedservice.IntegrationRuntimeState `json:"integration_runtime"`
	Timeouts             *synapselinkedservice.TimeoutsState            `json:"timeouts"`
}
