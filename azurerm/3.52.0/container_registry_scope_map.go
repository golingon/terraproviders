// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	containerregistryscopemap "github.com/golingon/terraproviders/azurerm/3.52.0/containerregistryscopemap"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewContainerRegistryScopeMap creates a new instance of [ContainerRegistryScopeMap].
func NewContainerRegistryScopeMap(name string, args ContainerRegistryScopeMapArgs) *ContainerRegistryScopeMap {
	return &ContainerRegistryScopeMap{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ContainerRegistryScopeMap)(nil)

// ContainerRegistryScopeMap represents the Terraform resource azurerm_container_registry_scope_map.
type ContainerRegistryScopeMap struct {
	Name      string
	Args      ContainerRegistryScopeMapArgs
	state     *containerRegistryScopeMapState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ContainerRegistryScopeMap].
func (crsm *ContainerRegistryScopeMap) Type() string {
	return "azurerm_container_registry_scope_map"
}

// LocalName returns the local name for [ContainerRegistryScopeMap].
func (crsm *ContainerRegistryScopeMap) LocalName() string {
	return crsm.Name
}

// Configuration returns the configuration (args) for [ContainerRegistryScopeMap].
func (crsm *ContainerRegistryScopeMap) Configuration() interface{} {
	return crsm.Args
}

// DependOn is used for other resources to depend on [ContainerRegistryScopeMap].
func (crsm *ContainerRegistryScopeMap) DependOn() terra.Reference {
	return terra.ReferenceResource(crsm)
}

// Dependencies returns the list of resources [ContainerRegistryScopeMap] depends_on.
func (crsm *ContainerRegistryScopeMap) Dependencies() terra.Dependencies {
	return crsm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ContainerRegistryScopeMap].
func (crsm *ContainerRegistryScopeMap) LifecycleManagement() *terra.Lifecycle {
	return crsm.Lifecycle
}

// Attributes returns the attributes for [ContainerRegistryScopeMap].
func (crsm *ContainerRegistryScopeMap) Attributes() containerRegistryScopeMapAttributes {
	return containerRegistryScopeMapAttributes{ref: terra.ReferenceResource(crsm)}
}

// ImportState imports the given attribute values into [ContainerRegistryScopeMap]'s state.
func (crsm *ContainerRegistryScopeMap) ImportState(av io.Reader) error {
	crsm.state = &containerRegistryScopeMapState{}
	if err := json.NewDecoder(av).Decode(crsm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crsm.Type(), crsm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ContainerRegistryScopeMap] has state.
func (crsm *ContainerRegistryScopeMap) State() (*containerRegistryScopeMapState, bool) {
	return crsm.state, crsm.state != nil
}

// StateMust returns the state for [ContainerRegistryScopeMap]. Panics if the state is nil.
func (crsm *ContainerRegistryScopeMap) StateMust() *containerRegistryScopeMapState {
	if crsm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crsm.Type(), crsm.LocalName()))
	}
	return crsm.state
}

// ContainerRegistryScopeMapArgs contains the configurations for azurerm_container_registry_scope_map.
type ContainerRegistryScopeMapArgs struct {
	// Actions: list of string, required
	Actions terra.ListValue[terra.StringValue] `hcl:"actions,attr" validate:"required"`
	// ContainerRegistryName: string, required
	ContainerRegistryName terra.StringValue `hcl:"container_registry_name,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *containerregistryscopemap.Timeouts `hcl:"timeouts,block"`
}
type containerRegistryScopeMapAttributes struct {
	ref terra.Reference
}

// Actions returns a reference to field actions of azurerm_container_registry_scope_map.
func (crsm containerRegistryScopeMapAttributes) Actions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](crsm.ref.Append("actions"))
}

// ContainerRegistryName returns a reference to field container_registry_name of azurerm_container_registry_scope_map.
func (crsm containerRegistryScopeMapAttributes) ContainerRegistryName() terra.StringValue {
	return terra.ReferenceAsString(crsm.ref.Append("container_registry_name"))
}

// Description returns a reference to field description of azurerm_container_registry_scope_map.
func (crsm containerRegistryScopeMapAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(crsm.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_container_registry_scope_map.
func (crsm containerRegistryScopeMapAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crsm.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_container_registry_scope_map.
func (crsm containerRegistryScopeMapAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crsm.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_container_registry_scope_map.
func (crsm containerRegistryScopeMapAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(crsm.ref.Append("resource_group_name"))
}

func (crsm containerRegistryScopeMapAttributes) Timeouts() containerregistryscopemap.TimeoutsAttributes {
	return terra.ReferenceAsSingle[containerregistryscopemap.TimeoutsAttributes](crsm.ref.Append("timeouts"))
}

type containerRegistryScopeMapState struct {
	Actions               []string                                 `json:"actions"`
	ContainerRegistryName string                                   `json:"container_registry_name"`
	Description           string                                   `json:"description"`
	Id                    string                                   `json:"id"`
	Name                  string                                   `json:"name"`
	ResourceGroupName     string                                   `json:"resource_group_name"`
	Timeouts              *containerregistryscopemap.TimeoutsState `json:"timeouts"`
}