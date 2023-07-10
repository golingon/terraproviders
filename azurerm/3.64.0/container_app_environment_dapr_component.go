// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	containerappenvironmentdaprcomponent "github.com/golingon/terraproviders/azurerm/3.64.0/containerappenvironmentdaprcomponent"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewContainerAppEnvironmentDaprComponent creates a new instance of [ContainerAppEnvironmentDaprComponent].
func NewContainerAppEnvironmentDaprComponent(name string, args ContainerAppEnvironmentDaprComponentArgs) *ContainerAppEnvironmentDaprComponent {
	return &ContainerAppEnvironmentDaprComponent{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ContainerAppEnvironmentDaprComponent)(nil)

// ContainerAppEnvironmentDaprComponent represents the Terraform resource azurerm_container_app_environment_dapr_component.
type ContainerAppEnvironmentDaprComponent struct {
	Name      string
	Args      ContainerAppEnvironmentDaprComponentArgs
	state     *containerAppEnvironmentDaprComponentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ContainerAppEnvironmentDaprComponent].
func (caedc *ContainerAppEnvironmentDaprComponent) Type() string {
	return "azurerm_container_app_environment_dapr_component"
}

// LocalName returns the local name for [ContainerAppEnvironmentDaprComponent].
func (caedc *ContainerAppEnvironmentDaprComponent) LocalName() string {
	return caedc.Name
}

// Configuration returns the configuration (args) for [ContainerAppEnvironmentDaprComponent].
func (caedc *ContainerAppEnvironmentDaprComponent) Configuration() interface{} {
	return caedc.Args
}

// DependOn is used for other resources to depend on [ContainerAppEnvironmentDaprComponent].
func (caedc *ContainerAppEnvironmentDaprComponent) DependOn() terra.Reference {
	return terra.ReferenceResource(caedc)
}

// Dependencies returns the list of resources [ContainerAppEnvironmentDaprComponent] depends_on.
func (caedc *ContainerAppEnvironmentDaprComponent) Dependencies() terra.Dependencies {
	return caedc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ContainerAppEnvironmentDaprComponent].
func (caedc *ContainerAppEnvironmentDaprComponent) LifecycleManagement() *terra.Lifecycle {
	return caedc.Lifecycle
}

// Attributes returns the attributes for [ContainerAppEnvironmentDaprComponent].
func (caedc *ContainerAppEnvironmentDaprComponent) Attributes() containerAppEnvironmentDaprComponentAttributes {
	return containerAppEnvironmentDaprComponentAttributes{ref: terra.ReferenceResource(caedc)}
}

// ImportState imports the given attribute values into [ContainerAppEnvironmentDaprComponent]'s state.
func (caedc *ContainerAppEnvironmentDaprComponent) ImportState(av io.Reader) error {
	caedc.state = &containerAppEnvironmentDaprComponentState{}
	if err := json.NewDecoder(av).Decode(caedc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", caedc.Type(), caedc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ContainerAppEnvironmentDaprComponent] has state.
func (caedc *ContainerAppEnvironmentDaprComponent) State() (*containerAppEnvironmentDaprComponentState, bool) {
	return caedc.state, caedc.state != nil
}

// StateMust returns the state for [ContainerAppEnvironmentDaprComponent]. Panics if the state is nil.
func (caedc *ContainerAppEnvironmentDaprComponent) StateMust() *containerAppEnvironmentDaprComponentState {
	if caedc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", caedc.Type(), caedc.LocalName()))
	}
	return caedc.state
}

// ContainerAppEnvironmentDaprComponentArgs contains the configurations for azurerm_container_app_environment_dapr_component.
type ContainerAppEnvironmentDaprComponentArgs struct {
	// ComponentType: string, required
	ComponentType terra.StringValue `hcl:"component_type,attr" validate:"required"`
	// ContainerAppEnvironmentId: string, required
	ContainerAppEnvironmentId terra.StringValue `hcl:"container_app_environment_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IgnoreErrors: bool, optional
	IgnoreErrors terra.BoolValue `hcl:"ignore_errors,attr"`
	// InitTimeout: string, optional
	InitTimeout terra.StringValue `hcl:"init_timeout,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Scopes: list of string, optional
	Scopes terra.ListValue[terra.StringValue] `hcl:"scopes,attr"`
	// Version: string, required
	Version terra.StringValue `hcl:"version,attr" validate:"required"`
	// Metadata: min=0
	Metadata []containerappenvironmentdaprcomponent.Metadata `hcl:"metadata,block" validate:"min=0"`
	// Secret: min=0
	Secret []containerappenvironmentdaprcomponent.Secret `hcl:"secret,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *containerappenvironmentdaprcomponent.Timeouts `hcl:"timeouts,block"`
}
type containerAppEnvironmentDaprComponentAttributes struct {
	ref terra.Reference
}

// ComponentType returns a reference to field component_type of azurerm_container_app_environment_dapr_component.
func (caedc containerAppEnvironmentDaprComponentAttributes) ComponentType() terra.StringValue {
	return terra.ReferenceAsString(caedc.ref.Append("component_type"))
}

// ContainerAppEnvironmentId returns a reference to field container_app_environment_id of azurerm_container_app_environment_dapr_component.
func (caedc containerAppEnvironmentDaprComponentAttributes) ContainerAppEnvironmentId() terra.StringValue {
	return terra.ReferenceAsString(caedc.ref.Append("container_app_environment_id"))
}

// Id returns a reference to field id of azurerm_container_app_environment_dapr_component.
func (caedc containerAppEnvironmentDaprComponentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(caedc.ref.Append("id"))
}

// IgnoreErrors returns a reference to field ignore_errors of azurerm_container_app_environment_dapr_component.
func (caedc containerAppEnvironmentDaprComponentAttributes) IgnoreErrors() terra.BoolValue {
	return terra.ReferenceAsBool(caedc.ref.Append("ignore_errors"))
}

// InitTimeout returns a reference to field init_timeout of azurerm_container_app_environment_dapr_component.
func (caedc containerAppEnvironmentDaprComponentAttributes) InitTimeout() terra.StringValue {
	return terra.ReferenceAsString(caedc.ref.Append("init_timeout"))
}

// Name returns a reference to field name of azurerm_container_app_environment_dapr_component.
func (caedc containerAppEnvironmentDaprComponentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(caedc.ref.Append("name"))
}

// Scopes returns a reference to field scopes of azurerm_container_app_environment_dapr_component.
func (caedc containerAppEnvironmentDaprComponentAttributes) Scopes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](caedc.ref.Append("scopes"))
}

// Version returns a reference to field version of azurerm_container_app_environment_dapr_component.
func (caedc containerAppEnvironmentDaprComponentAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(caedc.ref.Append("version"))
}

func (caedc containerAppEnvironmentDaprComponentAttributes) Metadata() terra.ListValue[containerappenvironmentdaprcomponent.MetadataAttributes] {
	return terra.ReferenceAsList[containerappenvironmentdaprcomponent.MetadataAttributes](caedc.ref.Append("metadata"))
}

func (caedc containerAppEnvironmentDaprComponentAttributes) Secret() terra.SetValue[containerappenvironmentdaprcomponent.SecretAttributes] {
	return terra.ReferenceAsSet[containerappenvironmentdaprcomponent.SecretAttributes](caedc.ref.Append("secret"))
}

func (caedc containerAppEnvironmentDaprComponentAttributes) Timeouts() containerappenvironmentdaprcomponent.TimeoutsAttributes {
	return terra.ReferenceAsSingle[containerappenvironmentdaprcomponent.TimeoutsAttributes](caedc.ref.Append("timeouts"))
}

type containerAppEnvironmentDaprComponentState struct {
	ComponentType             string                                               `json:"component_type"`
	ContainerAppEnvironmentId string                                               `json:"container_app_environment_id"`
	Id                        string                                               `json:"id"`
	IgnoreErrors              bool                                                 `json:"ignore_errors"`
	InitTimeout               string                                               `json:"init_timeout"`
	Name                      string                                               `json:"name"`
	Scopes                    []string                                             `json:"scopes"`
	Version                   string                                               `json:"version"`
	Metadata                  []containerappenvironmentdaprcomponent.MetadataState `json:"metadata"`
	Secret                    []containerappenvironmentdaprcomponent.SecretState   `json:"secret"`
	Timeouts                  *containerappenvironmentdaprcomponent.TimeoutsState  `json:"timeouts"`
}
