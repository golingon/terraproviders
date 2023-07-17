// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	containerregistrytoken "github.com/golingon/terraproviders/azurerm/3.65.0/containerregistrytoken"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewContainerRegistryToken creates a new instance of [ContainerRegistryToken].
func NewContainerRegistryToken(name string, args ContainerRegistryTokenArgs) *ContainerRegistryToken {
	return &ContainerRegistryToken{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ContainerRegistryToken)(nil)

// ContainerRegistryToken represents the Terraform resource azurerm_container_registry_token.
type ContainerRegistryToken struct {
	Name      string
	Args      ContainerRegistryTokenArgs
	state     *containerRegistryTokenState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ContainerRegistryToken].
func (crt *ContainerRegistryToken) Type() string {
	return "azurerm_container_registry_token"
}

// LocalName returns the local name for [ContainerRegistryToken].
func (crt *ContainerRegistryToken) LocalName() string {
	return crt.Name
}

// Configuration returns the configuration (args) for [ContainerRegistryToken].
func (crt *ContainerRegistryToken) Configuration() interface{} {
	return crt.Args
}

// DependOn is used for other resources to depend on [ContainerRegistryToken].
func (crt *ContainerRegistryToken) DependOn() terra.Reference {
	return terra.ReferenceResource(crt)
}

// Dependencies returns the list of resources [ContainerRegistryToken] depends_on.
func (crt *ContainerRegistryToken) Dependencies() terra.Dependencies {
	return crt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ContainerRegistryToken].
func (crt *ContainerRegistryToken) LifecycleManagement() *terra.Lifecycle {
	return crt.Lifecycle
}

// Attributes returns the attributes for [ContainerRegistryToken].
func (crt *ContainerRegistryToken) Attributes() containerRegistryTokenAttributes {
	return containerRegistryTokenAttributes{ref: terra.ReferenceResource(crt)}
}

// ImportState imports the given attribute values into [ContainerRegistryToken]'s state.
func (crt *ContainerRegistryToken) ImportState(av io.Reader) error {
	crt.state = &containerRegistryTokenState{}
	if err := json.NewDecoder(av).Decode(crt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crt.Type(), crt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ContainerRegistryToken] has state.
func (crt *ContainerRegistryToken) State() (*containerRegistryTokenState, bool) {
	return crt.state, crt.state != nil
}

// StateMust returns the state for [ContainerRegistryToken]. Panics if the state is nil.
func (crt *ContainerRegistryToken) StateMust() *containerRegistryTokenState {
	if crt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crt.Type(), crt.LocalName()))
	}
	return crt.state
}

// ContainerRegistryTokenArgs contains the configurations for azurerm_container_registry_token.
type ContainerRegistryTokenArgs struct {
	// ContainerRegistryName: string, required
	ContainerRegistryName terra.StringValue `hcl:"container_registry_name,attr" validate:"required"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ScopeMapId: string, required
	ScopeMapId terra.StringValue `hcl:"scope_map_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *containerregistrytoken.Timeouts `hcl:"timeouts,block"`
}
type containerRegistryTokenAttributes struct {
	ref terra.Reference
}

// ContainerRegistryName returns a reference to field container_registry_name of azurerm_container_registry_token.
func (crt containerRegistryTokenAttributes) ContainerRegistryName() terra.StringValue {
	return terra.ReferenceAsString(crt.ref.Append("container_registry_name"))
}

// Enabled returns a reference to field enabled of azurerm_container_registry_token.
func (crt containerRegistryTokenAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(crt.ref.Append("enabled"))
}

// Id returns a reference to field id of azurerm_container_registry_token.
func (crt containerRegistryTokenAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crt.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_container_registry_token.
func (crt containerRegistryTokenAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crt.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_container_registry_token.
func (crt containerRegistryTokenAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(crt.ref.Append("resource_group_name"))
}

// ScopeMapId returns a reference to field scope_map_id of azurerm_container_registry_token.
func (crt containerRegistryTokenAttributes) ScopeMapId() terra.StringValue {
	return terra.ReferenceAsString(crt.ref.Append("scope_map_id"))
}

func (crt containerRegistryTokenAttributes) Timeouts() containerregistrytoken.TimeoutsAttributes {
	return terra.ReferenceAsSingle[containerregistrytoken.TimeoutsAttributes](crt.ref.Append("timeouts"))
}

type containerRegistryTokenState struct {
	ContainerRegistryName string                                `json:"container_registry_name"`
	Enabled               bool                                  `json:"enabled"`
	Id                    string                                `json:"id"`
	Name                  string                                `json:"name"`
	ResourceGroupName     string                                `json:"resource_group_name"`
	ScopeMapId            string                                `json:"scope_map_id"`
	Timeouts              *containerregistrytoken.TimeoutsState `json:"timeouts"`
}
