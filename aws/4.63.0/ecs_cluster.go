// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ecscluster "github.com/golingon/terraproviders/aws/4.63.0/ecscluster"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEcsCluster creates a new instance of [EcsCluster].
func NewEcsCluster(name string, args EcsClusterArgs) *EcsCluster {
	return &EcsCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EcsCluster)(nil)

// EcsCluster represents the Terraform resource aws_ecs_cluster.
type EcsCluster struct {
	Name      string
	Args      EcsClusterArgs
	state     *ecsClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EcsCluster].
func (ec *EcsCluster) Type() string {
	return "aws_ecs_cluster"
}

// LocalName returns the local name for [EcsCluster].
func (ec *EcsCluster) LocalName() string {
	return ec.Name
}

// Configuration returns the configuration (args) for [EcsCluster].
func (ec *EcsCluster) Configuration() interface{} {
	return ec.Args
}

// DependOn is used for other resources to depend on [EcsCluster].
func (ec *EcsCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(ec)
}

// Dependencies returns the list of resources [EcsCluster] depends_on.
func (ec *EcsCluster) Dependencies() terra.Dependencies {
	return ec.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EcsCluster].
func (ec *EcsCluster) LifecycleManagement() *terra.Lifecycle {
	return ec.Lifecycle
}

// Attributes returns the attributes for [EcsCluster].
func (ec *EcsCluster) Attributes() ecsClusterAttributes {
	return ecsClusterAttributes{ref: terra.ReferenceResource(ec)}
}

// ImportState imports the given attribute values into [EcsCluster]'s state.
func (ec *EcsCluster) ImportState(av io.Reader) error {
	ec.state = &ecsClusterState{}
	if err := json.NewDecoder(av).Decode(ec.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ec.Type(), ec.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EcsCluster] has state.
func (ec *EcsCluster) State() (*ecsClusterState, bool) {
	return ec.state, ec.state != nil
}

// StateMust returns the state for [EcsCluster]. Panics if the state is nil.
func (ec *EcsCluster) StateMust() *ecsClusterState {
	if ec.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ec.Type(), ec.LocalName()))
	}
	return ec.state
}

// EcsClusterArgs contains the configurations for aws_ecs_cluster.
type EcsClusterArgs struct {
	// CapacityProviders: set of string, optional
	CapacityProviders terra.SetValue[terra.StringValue] `hcl:"capacity_providers,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Configuration: optional
	Configuration *ecscluster.Configuration `hcl:"configuration,block"`
	// DefaultCapacityProviderStrategy: min=0
	DefaultCapacityProviderStrategy []ecscluster.DefaultCapacityProviderStrategy `hcl:"default_capacity_provider_strategy,block" validate:"min=0"`
	// ServiceConnectDefaults: optional
	ServiceConnectDefaults *ecscluster.ServiceConnectDefaults `hcl:"service_connect_defaults,block"`
	// Setting: min=0
	Setting []ecscluster.Setting `hcl:"setting,block" validate:"min=0"`
}
type ecsClusterAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ecs_cluster.
func (ec ecsClusterAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("arn"))
}

// CapacityProviders returns a reference to field capacity_providers of aws_ecs_cluster.
func (ec ecsClusterAttributes) CapacityProviders() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ec.ref.Append("capacity_providers"))
}

// Id returns a reference to field id of aws_ecs_cluster.
func (ec ecsClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("id"))
}

// Name returns a reference to field name of aws_ecs_cluster.
func (ec ecsClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_ecs_cluster.
func (ec ecsClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ec.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_ecs_cluster.
func (ec ecsClusterAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ec.ref.Append("tags_all"))
}

func (ec ecsClusterAttributes) Configuration() terra.ListValue[ecscluster.ConfigurationAttributes] {
	return terra.ReferenceAsList[ecscluster.ConfigurationAttributes](ec.ref.Append("configuration"))
}

func (ec ecsClusterAttributes) DefaultCapacityProviderStrategy() terra.SetValue[ecscluster.DefaultCapacityProviderStrategyAttributes] {
	return terra.ReferenceAsSet[ecscluster.DefaultCapacityProviderStrategyAttributes](ec.ref.Append("default_capacity_provider_strategy"))
}

func (ec ecsClusterAttributes) ServiceConnectDefaults() terra.ListValue[ecscluster.ServiceConnectDefaultsAttributes] {
	return terra.ReferenceAsList[ecscluster.ServiceConnectDefaultsAttributes](ec.ref.Append("service_connect_defaults"))
}

func (ec ecsClusterAttributes) Setting() terra.SetValue[ecscluster.SettingAttributes] {
	return terra.ReferenceAsSet[ecscluster.SettingAttributes](ec.ref.Append("setting"))
}

type ecsClusterState struct {
	Arn                             string                                            `json:"arn"`
	CapacityProviders               []string                                          `json:"capacity_providers"`
	Id                              string                                            `json:"id"`
	Name                            string                                            `json:"name"`
	Tags                            map[string]string                                 `json:"tags"`
	TagsAll                         map[string]string                                 `json:"tags_all"`
	Configuration                   []ecscluster.ConfigurationState                   `json:"configuration"`
	DefaultCapacityProviderStrategy []ecscluster.DefaultCapacityProviderStrategyState `json:"default_capacity_provider_strategy"`
	ServiceConnectDefaults          []ecscluster.ServiceConnectDefaultsState          `json:"service_connect_defaults"`
	Setting                         []ecscluster.SettingState                         `json:"setting"`
}