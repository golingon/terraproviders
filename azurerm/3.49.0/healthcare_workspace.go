// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	healthcareworkspace "github.com/golingon/terraproviders/azurerm/3.49.0/healthcareworkspace"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewHealthcareWorkspace(name string, args HealthcareWorkspaceArgs) *HealthcareWorkspace {
	return &HealthcareWorkspace{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*HealthcareWorkspace)(nil)

type HealthcareWorkspace struct {
	Name  string
	Args  HealthcareWorkspaceArgs
	state *healthcareWorkspaceState
}

func (hw *HealthcareWorkspace) Type() string {
	return "azurerm_healthcare_workspace"
}

func (hw *HealthcareWorkspace) LocalName() string {
	return hw.Name
}

func (hw *HealthcareWorkspace) Configuration() interface{} {
	return hw.Args
}

func (hw *HealthcareWorkspace) Attributes() healthcareWorkspaceAttributes {
	return healthcareWorkspaceAttributes{ref: terra.ReferenceResource(hw)}
}

func (hw *HealthcareWorkspace) ImportState(av io.Reader) error {
	hw.state = &healthcareWorkspaceState{}
	if err := json.NewDecoder(av).Decode(hw.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", hw.Type(), hw.LocalName(), err)
	}
	return nil
}

func (hw *HealthcareWorkspace) State() (*healthcareWorkspaceState, bool) {
	return hw.state, hw.state != nil
}

func (hw *HealthcareWorkspace) StateMust() *healthcareWorkspaceState {
	if hw.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", hw.Type(), hw.LocalName()))
	}
	return hw.state
}

func (hw *HealthcareWorkspace) DependOn() terra.Reference {
	return terra.ReferenceResource(hw)
}

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
	// DependsOn contains resources that HealthcareWorkspace depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type healthcareWorkspaceAttributes struct {
	ref terra.Reference
}

func (hw healthcareWorkspaceAttributes) Id() terra.StringValue {
	return terra.ReferenceString(hw.ref.Append("id"))
}

func (hw healthcareWorkspaceAttributes) Location() terra.StringValue {
	return terra.ReferenceString(hw.ref.Append("location"))
}

func (hw healthcareWorkspaceAttributes) Name() terra.StringValue {
	return terra.ReferenceString(hw.ref.Append("name"))
}

func (hw healthcareWorkspaceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(hw.ref.Append("resource_group_name"))
}

func (hw healthcareWorkspaceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](hw.ref.Append("tags"))
}

func (hw healthcareWorkspaceAttributes) PrivateEndpointConnection() terra.SetValue[healthcareworkspace.PrivateEndpointConnectionAttributes] {
	return terra.ReferenceSet[healthcareworkspace.PrivateEndpointConnectionAttributes](hw.ref.Append("private_endpoint_connection"))
}

func (hw healthcareWorkspaceAttributes) Timeouts() healthcareworkspace.TimeoutsAttributes {
	return terra.ReferenceSingle[healthcareworkspace.TimeoutsAttributes](hw.ref.Append("timeouts"))
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
