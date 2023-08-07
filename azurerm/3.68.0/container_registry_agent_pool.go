// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	containerregistryagentpool "github.com/golingon/terraproviders/azurerm/3.68.0/containerregistryagentpool"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewContainerRegistryAgentPool creates a new instance of [ContainerRegistryAgentPool].
func NewContainerRegistryAgentPool(name string, args ContainerRegistryAgentPoolArgs) *ContainerRegistryAgentPool {
	return &ContainerRegistryAgentPool{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ContainerRegistryAgentPool)(nil)

// ContainerRegistryAgentPool represents the Terraform resource azurerm_container_registry_agent_pool.
type ContainerRegistryAgentPool struct {
	Name      string
	Args      ContainerRegistryAgentPoolArgs
	state     *containerRegistryAgentPoolState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ContainerRegistryAgentPool].
func (crap *ContainerRegistryAgentPool) Type() string {
	return "azurerm_container_registry_agent_pool"
}

// LocalName returns the local name for [ContainerRegistryAgentPool].
func (crap *ContainerRegistryAgentPool) LocalName() string {
	return crap.Name
}

// Configuration returns the configuration (args) for [ContainerRegistryAgentPool].
func (crap *ContainerRegistryAgentPool) Configuration() interface{} {
	return crap.Args
}

// DependOn is used for other resources to depend on [ContainerRegistryAgentPool].
func (crap *ContainerRegistryAgentPool) DependOn() terra.Reference {
	return terra.ReferenceResource(crap)
}

// Dependencies returns the list of resources [ContainerRegistryAgentPool] depends_on.
func (crap *ContainerRegistryAgentPool) Dependencies() terra.Dependencies {
	return crap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ContainerRegistryAgentPool].
func (crap *ContainerRegistryAgentPool) LifecycleManagement() *terra.Lifecycle {
	return crap.Lifecycle
}

// Attributes returns the attributes for [ContainerRegistryAgentPool].
func (crap *ContainerRegistryAgentPool) Attributes() containerRegistryAgentPoolAttributes {
	return containerRegistryAgentPoolAttributes{ref: terra.ReferenceResource(crap)}
}

// ImportState imports the given attribute values into [ContainerRegistryAgentPool]'s state.
func (crap *ContainerRegistryAgentPool) ImportState(av io.Reader) error {
	crap.state = &containerRegistryAgentPoolState{}
	if err := json.NewDecoder(av).Decode(crap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crap.Type(), crap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ContainerRegistryAgentPool] has state.
func (crap *ContainerRegistryAgentPool) State() (*containerRegistryAgentPoolState, bool) {
	return crap.state, crap.state != nil
}

// StateMust returns the state for [ContainerRegistryAgentPool]. Panics if the state is nil.
func (crap *ContainerRegistryAgentPool) StateMust() *containerRegistryAgentPoolState {
	if crap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crap.Type(), crap.LocalName()))
	}
	return crap.state
}

// ContainerRegistryAgentPoolArgs contains the configurations for azurerm_container_registry_agent_pool.
type ContainerRegistryAgentPoolArgs struct {
	// ContainerRegistryName: string, required
	ContainerRegistryName terra.StringValue `hcl:"container_registry_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceCount: number, optional
	InstanceCount terra.NumberValue `hcl:"instance_count,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Tier: string, optional
	Tier terra.StringValue `hcl:"tier,attr"`
	// VirtualNetworkSubnetId: string, optional
	VirtualNetworkSubnetId terra.StringValue `hcl:"virtual_network_subnet_id,attr"`
	// Timeouts: optional
	Timeouts *containerregistryagentpool.Timeouts `hcl:"timeouts,block"`
}
type containerRegistryAgentPoolAttributes struct {
	ref terra.Reference
}

// ContainerRegistryName returns a reference to field container_registry_name of azurerm_container_registry_agent_pool.
func (crap containerRegistryAgentPoolAttributes) ContainerRegistryName() terra.StringValue {
	return terra.ReferenceAsString(crap.ref.Append("container_registry_name"))
}

// Id returns a reference to field id of azurerm_container_registry_agent_pool.
func (crap containerRegistryAgentPoolAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crap.ref.Append("id"))
}

// InstanceCount returns a reference to field instance_count of azurerm_container_registry_agent_pool.
func (crap containerRegistryAgentPoolAttributes) InstanceCount() terra.NumberValue {
	return terra.ReferenceAsNumber(crap.ref.Append("instance_count"))
}

// Location returns a reference to field location of azurerm_container_registry_agent_pool.
func (crap containerRegistryAgentPoolAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(crap.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_container_registry_agent_pool.
func (crap containerRegistryAgentPoolAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crap.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_container_registry_agent_pool.
func (crap containerRegistryAgentPoolAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(crap.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_container_registry_agent_pool.
func (crap containerRegistryAgentPoolAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](crap.ref.Append("tags"))
}

// Tier returns a reference to field tier of azurerm_container_registry_agent_pool.
func (crap containerRegistryAgentPoolAttributes) Tier() terra.StringValue {
	return terra.ReferenceAsString(crap.ref.Append("tier"))
}

// VirtualNetworkSubnetId returns a reference to field virtual_network_subnet_id of azurerm_container_registry_agent_pool.
func (crap containerRegistryAgentPoolAttributes) VirtualNetworkSubnetId() terra.StringValue {
	return terra.ReferenceAsString(crap.ref.Append("virtual_network_subnet_id"))
}

func (crap containerRegistryAgentPoolAttributes) Timeouts() containerregistryagentpool.TimeoutsAttributes {
	return terra.ReferenceAsSingle[containerregistryagentpool.TimeoutsAttributes](crap.ref.Append("timeouts"))
}

type containerRegistryAgentPoolState struct {
	ContainerRegistryName  string                                    `json:"container_registry_name"`
	Id                     string                                    `json:"id"`
	InstanceCount          float64                                   `json:"instance_count"`
	Location               string                                    `json:"location"`
	Name                   string                                    `json:"name"`
	ResourceGroupName      string                                    `json:"resource_group_name"`
	Tags                   map[string]string                         `json:"tags"`
	Tier                   string                                    `json:"tier"`
	VirtualNetworkSubnetId string                                    `json:"virtual_network_subnet_id"`
	Timeouts               *containerregistryagentpool.TimeoutsState `json:"timeouts"`
}
