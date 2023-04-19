// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	healthcareworkspace "github.com/golingon/terraproviders/azurerm/3.52.0/healthcareworkspace"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewHealthcareWorkspace creates a new instance of [HealthcareWorkspace].
func NewHealthcareWorkspace(name string, args HealthcareWorkspaceArgs) *HealthcareWorkspace {
	return &HealthcareWorkspace{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*HealthcareWorkspace)(nil)

// HealthcareWorkspace represents the Terraform resource azurerm_healthcare_workspace.
type HealthcareWorkspace struct {
	Name      string
	Args      HealthcareWorkspaceArgs
	state     *healthcareWorkspaceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [HealthcareWorkspace].
func (hw *HealthcareWorkspace) Type() string {
	return "azurerm_healthcare_workspace"
}

// LocalName returns the local name for [HealthcareWorkspace].
func (hw *HealthcareWorkspace) LocalName() string {
	return hw.Name
}

// Configuration returns the configuration (args) for [HealthcareWorkspace].
func (hw *HealthcareWorkspace) Configuration() interface{} {
	return hw.Args
}

// DependOn is used for other resources to depend on [HealthcareWorkspace].
func (hw *HealthcareWorkspace) DependOn() terra.Reference {
	return terra.ReferenceResource(hw)
}

// Dependencies returns the list of resources [HealthcareWorkspace] depends_on.
func (hw *HealthcareWorkspace) Dependencies() terra.Dependencies {
	return hw.DependsOn
}

// LifecycleManagement returns the lifecycle block for [HealthcareWorkspace].
func (hw *HealthcareWorkspace) LifecycleManagement() *terra.Lifecycle {
	return hw.Lifecycle
}

// Attributes returns the attributes for [HealthcareWorkspace].
func (hw *HealthcareWorkspace) Attributes() healthcareWorkspaceAttributes {
	return healthcareWorkspaceAttributes{ref: terra.ReferenceResource(hw)}
}

// ImportState imports the given attribute values into [HealthcareWorkspace]'s state.
func (hw *HealthcareWorkspace) ImportState(av io.Reader) error {
	hw.state = &healthcareWorkspaceState{}
	if err := json.NewDecoder(av).Decode(hw.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", hw.Type(), hw.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [HealthcareWorkspace] has state.
func (hw *HealthcareWorkspace) State() (*healthcareWorkspaceState, bool) {
	return hw.state, hw.state != nil
}

// StateMust returns the state for [HealthcareWorkspace]. Panics if the state is nil.
func (hw *HealthcareWorkspace) StateMust() *healthcareWorkspaceState {
	if hw.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", hw.Type(), hw.LocalName()))
	}
	return hw.state
}

// HealthcareWorkspaceArgs contains the configurations for azurerm_healthcare_workspace.
type HealthcareWorkspaceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// PrivateEndpointConnection: min=0
	PrivateEndpointConnection []healthcareworkspace.PrivateEndpointConnection `hcl:"private_endpoint_connection,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *healthcareworkspace.Timeouts `hcl:"timeouts,block"`
}
type healthcareWorkspaceAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_healthcare_workspace.
func (hw healthcareWorkspaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(hw.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_healthcare_workspace.
func (hw healthcareWorkspaceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(hw.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_healthcare_workspace.
func (hw healthcareWorkspaceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(hw.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_healthcare_workspace.
func (hw healthcareWorkspaceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(hw.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_healthcare_workspace.
func (hw healthcareWorkspaceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](hw.ref.Append("tags"))
}

func (hw healthcareWorkspaceAttributes) PrivateEndpointConnection() terra.SetValue[healthcareworkspace.PrivateEndpointConnectionAttributes] {
	return terra.ReferenceAsSet[healthcareworkspace.PrivateEndpointConnectionAttributes](hw.ref.Append("private_endpoint_connection"))
}

func (hw healthcareWorkspaceAttributes) Timeouts() healthcareworkspace.TimeoutsAttributes {
	return terra.ReferenceAsSingle[healthcareworkspace.TimeoutsAttributes](hw.ref.Append("timeouts"))
}

type healthcareWorkspaceState struct {
	Id                        string                                               `json:"id"`
	Location                  string                                               `json:"location"`
	Name                      string                                               `json:"name"`
	ResourceGroupName         string                                               `json:"resource_group_name"`
	Tags                      map[string]string                                    `json:"tags"`
	PrivateEndpointConnection []healthcareworkspace.PrivateEndpointConnectionState `json:"private_endpoint_connection"`
	Timeouts                  *healthcareworkspace.TimeoutsState                   `json:"timeouts"`
}
