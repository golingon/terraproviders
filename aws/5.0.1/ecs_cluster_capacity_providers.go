// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ecsclustercapacityproviders "github.com/golingon/terraproviders/aws/5.0.1/ecsclustercapacityproviders"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEcsClusterCapacityProviders creates a new instance of [EcsClusterCapacityProviders].
func NewEcsClusterCapacityProviders(name string, args EcsClusterCapacityProvidersArgs) *EcsClusterCapacityProviders {
	return &EcsClusterCapacityProviders{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EcsClusterCapacityProviders)(nil)

// EcsClusterCapacityProviders represents the Terraform resource aws_ecs_cluster_capacity_providers.
type EcsClusterCapacityProviders struct {
	Name      string
	Args      EcsClusterCapacityProvidersArgs
	state     *ecsClusterCapacityProvidersState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EcsClusterCapacityProviders].
func (eccp *EcsClusterCapacityProviders) Type() string {
	return "aws_ecs_cluster_capacity_providers"
}

// LocalName returns the local name for [EcsClusterCapacityProviders].
func (eccp *EcsClusterCapacityProviders) LocalName() string {
	return eccp.Name
}

// Configuration returns the configuration (args) for [EcsClusterCapacityProviders].
func (eccp *EcsClusterCapacityProviders) Configuration() interface{} {
	return eccp.Args
}

// DependOn is used for other resources to depend on [EcsClusterCapacityProviders].
func (eccp *EcsClusterCapacityProviders) DependOn() terra.Reference {
	return terra.ReferenceResource(eccp)
}

// Dependencies returns the list of resources [EcsClusterCapacityProviders] depends_on.
func (eccp *EcsClusterCapacityProviders) Dependencies() terra.Dependencies {
	return eccp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EcsClusterCapacityProviders].
func (eccp *EcsClusterCapacityProviders) LifecycleManagement() *terra.Lifecycle {
	return eccp.Lifecycle
}

// Attributes returns the attributes for [EcsClusterCapacityProviders].
func (eccp *EcsClusterCapacityProviders) Attributes() ecsClusterCapacityProvidersAttributes {
	return ecsClusterCapacityProvidersAttributes{ref: terra.ReferenceResource(eccp)}
}

// ImportState imports the given attribute values into [EcsClusterCapacityProviders]'s state.
func (eccp *EcsClusterCapacityProviders) ImportState(av io.Reader) error {
	eccp.state = &ecsClusterCapacityProvidersState{}
	if err := json.NewDecoder(av).Decode(eccp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", eccp.Type(), eccp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EcsClusterCapacityProviders] has state.
func (eccp *EcsClusterCapacityProviders) State() (*ecsClusterCapacityProvidersState, bool) {
	return eccp.state, eccp.state != nil
}

// StateMust returns the state for [EcsClusterCapacityProviders]. Panics if the state is nil.
func (eccp *EcsClusterCapacityProviders) StateMust() *ecsClusterCapacityProvidersState {
	if eccp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", eccp.Type(), eccp.LocalName()))
	}
	return eccp.state
}

// EcsClusterCapacityProvidersArgs contains the configurations for aws_ecs_cluster_capacity_providers.
type EcsClusterCapacityProvidersArgs struct {
	// CapacityProviders: set of string, optional
	CapacityProviders terra.SetValue[terra.StringValue] `hcl:"capacity_providers,attr"`
	// ClusterName: string, required
	ClusterName terra.StringValue `hcl:"cluster_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// DefaultCapacityProviderStrategy: min=0
	DefaultCapacityProviderStrategy []ecsclustercapacityproviders.DefaultCapacityProviderStrategy `hcl:"default_capacity_provider_strategy,block" validate:"min=0"`
}
type ecsClusterCapacityProvidersAttributes struct {
	ref terra.Reference
}

// CapacityProviders returns a reference to field capacity_providers of aws_ecs_cluster_capacity_providers.
func (eccp ecsClusterCapacityProvidersAttributes) CapacityProviders() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](eccp.ref.Append("capacity_providers"))
}

// ClusterName returns a reference to field cluster_name of aws_ecs_cluster_capacity_providers.
func (eccp ecsClusterCapacityProvidersAttributes) ClusterName() terra.StringValue {
	return terra.ReferenceAsString(eccp.ref.Append("cluster_name"))
}

// Id returns a reference to field id of aws_ecs_cluster_capacity_providers.
func (eccp ecsClusterCapacityProvidersAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(eccp.ref.Append("id"))
}

func (eccp ecsClusterCapacityProvidersAttributes) DefaultCapacityProviderStrategy() terra.SetValue[ecsclustercapacityproviders.DefaultCapacityProviderStrategyAttributes] {
	return terra.ReferenceAsSet[ecsclustercapacityproviders.DefaultCapacityProviderStrategyAttributes](eccp.ref.Append("default_capacity_provider_strategy"))
}

type ecsClusterCapacityProvidersState struct {
	CapacityProviders               []string                                                           `json:"capacity_providers"`
	ClusterName                     string                                                             `json:"cluster_name"`
	Id                              string                                                             `json:"id"`
	DefaultCapacityProviderStrategy []ecsclustercapacityproviders.DefaultCapacityProviderStrategyState `json:"default_capacity_provider_strategy"`
}
