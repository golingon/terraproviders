// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_ecs_cluster_capacity_providers

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_ecs_cluster_capacity_providers.
type Resource struct {
	Name      string
	Args      Args
	state     *awsEcsClusterCapacityProvidersState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aeccp *Resource) Type() string {
	return "aws_ecs_cluster_capacity_providers"
}

// LocalName returns the local name for [Resource].
func (aeccp *Resource) LocalName() string {
	return aeccp.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aeccp *Resource) Configuration() interface{} {
	return aeccp.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aeccp *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aeccp)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aeccp *Resource) Dependencies() terra.Dependencies {
	return aeccp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aeccp *Resource) LifecycleManagement() *terra.Lifecycle {
	return aeccp.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aeccp *Resource) Attributes() awsEcsClusterCapacityProvidersAttributes {
	return awsEcsClusterCapacityProvidersAttributes{ref: terra.ReferenceResource(aeccp)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aeccp *Resource) ImportState(state io.Reader) error {
	aeccp.state = &awsEcsClusterCapacityProvidersState{}
	if err := json.NewDecoder(state).Decode(aeccp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aeccp.Type(), aeccp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aeccp *Resource) State() (*awsEcsClusterCapacityProvidersState, bool) {
	return aeccp.state, aeccp.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aeccp *Resource) StateMust() *awsEcsClusterCapacityProvidersState {
	if aeccp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aeccp.Type(), aeccp.LocalName()))
	}
	return aeccp.state
}

// Args contains the configurations for aws_ecs_cluster_capacity_providers.
type Args struct {
	// CapacityProviders: set of string, optional
	CapacityProviders terra.SetValue[terra.StringValue] `hcl:"capacity_providers,attr"`
	// ClusterName: string, required
	ClusterName terra.StringValue `hcl:"cluster_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// DefaultCapacityProviderStrategy: min=0
	DefaultCapacityProviderStrategy []DefaultCapacityProviderStrategy `hcl:"default_capacity_provider_strategy,block" validate:"min=0"`
}

type awsEcsClusterCapacityProvidersAttributes struct {
	ref terra.Reference
}

// CapacityProviders returns a reference to field capacity_providers of aws_ecs_cluster_capacity_providers.
func (aeccp awsEcsClusterCapacityProvidersAttributes) CapacityProviders() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](aeccp.ref.Append("capacity_providers"))
}

// ClusterName returns a reference to field cluster_name of aws_ecs_cluster_capacity_providers.
func (aeccp awsEcsClusterCapacityProvidersAttributes) ClusterName() terra.StringValue {
	return terra.ReferenceAsString(aeccp.ref.Append("cluster_name"))
}

// Id returns a reference to field id of aws_ecs_cluster_capacity_providers.
func (aeccp awsEcsClusterCapacityProvidersAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aeccp.ref.Append("id"))
}

func (aeccp awsEcsClusterCapacityProvidersAttributes) DefaultCapacityProviderStrategy() terra.SetValue[DefaultCapacityProviderStrategyAttributes] {
	return terra.ReferenceAsSet[DefaultCapacityProviderStrategyAttributes](aeccp.ref.Append("default_capacity_provider_strategy"))
}

type awsEcsClusterCapacityProvidersState struct {
	CapacityProviders               []string                               `json:"capacity_providers"`
	ClusterName                     string                                 `json:"cluster_name"`
	Id                              string                                 `json:"id"`
	DefaultCapacityProviderStrategy []DefaultCapacityProviderStrategyState `json:"default_capacity_provider_strategy"`
}
