// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ecstaskset "github.com/golingon/terraproviders/aws/5.6.2/ecstaskset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEcsTaskSet creates a new instance of [EcsTaskSet].
func NewEcsTaskSet(name string, args EcsTaskSetArgs) *EcsTaskSet {
	return &EcsTaskSet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EcsTaskSet)(nil)

// EcsTaskSet represents the Terraform resource aws_ecs_task_set.
type EcsTaskSet struct {
	Name      string
	Args      EcsTaskSetArgs
	state     *ecsTaskSetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EcsTaskSet].
func (ets *EcsTaskSet) Type() string {
	return "aws_ecs_task_set"
}

// LocalName returns the local name for [EcsTaskSet].
func (ets *EcsTaskSet) LocalName() string {
	return ets.Name
}

// Configuration returns the configuration (args) for [EcsTaskSet].
func (ets *EcsTaskSet) Configuration() interface{} {
	return ets.Args
}

// DependOn is used for other resources to depend on [EcsTaskSet].
func (ets *EcsTaskSet) DependOn() terra.Reference {
	return terra.ReferenceResource(ets)
}

// Dependencies returns the list of resources [EcsTaskSet] depends_on.
func (ets *EcsTaskSet) Dependencies() terra.Dependencies {
	return ets.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EcsTaskSet].
func (ets *EcsTaskSet) LifecycleManagement() *terra.Lifecycle {
	return ets.Lifecycle
}

// Attributes returns the attributes for [EcsTaskSet].
func (ets *EcsTaskSet) Attributes() ecsTaskSetAttributes {
	return ecsTaskSetAttributes{ref: terra.ReferenceResource(ets)}
}

// ImportState imports the given attribute values into [EcsTaskSet]'s state.
func (ets *EcsTaskSet) ImportState(av io.Reader) error {
	ets.state = &ecsTaskSetState{}
	if err := json.NewDecoder(av).Decode(ets.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ets.Type(), ets.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EcsTaskSet] has state.
func (ets *EcsTaskSet) State() (*ecsTaskSetState, bool) {
	return ets.state, ets.state != nil
}

// StateMust returns the state for [EcsTaskSet]. Panics if the state is nil.
func (ets *EcsTaskSet) StateMust() *ecsTaskSetState {
	if ets.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ets.Type(), ets.LocalName()))
	}
	return ets.state
}

// EcsTaskSetArgs contains the configurations for aws_ecs_task_set.
type EcsTaskSetArgs struct {
	// Cluster: string, required
	Cluster terra.StringValue `hcl:"cluster,attr" validate:"required"`
	// ExternalId: string, optional
	ExternalId terra.StringValue `hcl:"external_id,attr"`
	// ForceDelete: bool, optional
	ForceDelete terra.BoolValue `hcl:"force_delete,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LaunchType: string, optional
	LaunchType terra.StringValue `hcl:"launch_type,attr"`
	// PlatformVersion: string, optional
	PlatformVersion terra.StringValue `hcl:"platform_version,attr"`
	// Service: string, required
	Service terra.StringValue `hcl:"service,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TaskDefinition: string, required
	TaskDefinition terra.StringValue `hcl:"task_definition,attr" validate:"required"`
	// WaitUntilStable: bool, optional
	WaitUntilStable terra.BoolValue `hcl:"wait_until_stable,attr"`
	// WaitUntilStableTimeout: string, optional
	WaitUntilStableTimeout terra.StringValue `hcl:"wait_until_stable_timeout,attr"`
	// CapacityProviderStrategy: min=0
	CapacityProviderStrategy []ecstaskset.CapacityProviderStrategy `hcl:"capacity_provider_strategy,block" validate:"min=0"`
	// LoadBalancer: min=0
	LoadBalancer []ecstaskset.LoadBalancer `hcl:"load_balancer,block" validate:"min=0"`
	// NetworkConfiguration: optional
	NetworkConfiguration *ecstaskset.NetworkConfiguration `hcl:"network_configuration,block"`
	// Scale: optional
	Scale *ecstaskset.Scale `hcl:"scale,block"`
	// ServiceRegistries: optional
	ServiceRegistries *ecstaskset.ServiceRegistries `hcl:"service_registries,block"`
}
type ecsTaskSetAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ecs_task_set.
func (ets ecsTaskSetAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ets.ref.Append("arn"))
}

// Cluster returns a reference to field cluster of aws_ecs_task_set.
func (ets ecsTaskSetAttributes) Cluster() terra.StringValue {
	return terra.ReferenceAsString(ets.ref.Append("cluster"))
}

// ExternalId returns a reference to field external_id of aws_ecs_task_set.
func (ets ecsTaskSetAttributes) ExternalId() terra.StringValue {
	return terra.ReferenceAsString(ets.ref.Append("external_id"))
}

// ForceDelete returns a reference to field force_delete of aws_ecs_task_set.
func (ets ecsTaskSetAttributes) ForceDelete() terra.BoolValue {
	return terra.ReferenceAsBool(ets.ref.Append("force_delete"))
}

// Id returns a reference to field id of aws_ecs_task_set.
func (ets ecsTaskSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ets.ref.Append("id"))
}

// LaunchType returns a reference to field launch_type of aws_ecs_task_set.
func (ets ecsTaskSetAttributes) LaunchType() terra.StringValue {
	return terra.ReferenceAsString(ets.ref.Append("launch_type"))
}

// PlatformVersion returns a reference to field platform_version of aws_ecs_task_set.
func (ets ecsTaskSetAttributes) PlatformVersion() terra.StringValue {
	return terra.ReferenceAsString(ets.ref.Append("platform_version"))
}

// Service returns a reference to field service of aws_ecs_task_set.
func (ets ecsTaskSetAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(ets.ref.Append("service"))
}

// StabilityStatus returns a reference to field stability_status of aws_ecs_task_set.
func (ets ecsTaskSetAttributes) StabilityStatus() terra.StringValue {
	return terra.ReferenceAsString(ets.ref.Append("stability_status"))
}

// Status returns a reference to field status of aws_ecs_task_set.
func (ets ecsTaskSetAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(ets.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_ecs_task_set.
func (ets ecsTaskSetAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ets.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_ecs_task_set.
func (ets ecsTaskSetAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ets.ref.Append("tags_all"))
}

// TaskDefinition returns a reference to field task_definition of aws_ecs_task_set.
func (ets ecsTaskSetAttributes) TaskDefinition() terra.StringValue {
	return terra.ReferenceAsString(ets.ref.Append("task_definition"))
}

// TaskSetId returns a reference to field task_set_id of aws_ecs_task_set.
func (ets ecsTaskSetAttributes) TaskSetId() terra.StringValue {
	return terra.ReferenceAsString(ets.ref.Append("task_set_id"))
}

// WaitUntilStable returns a reference to field wait_until_stable of aws_ecs_task_set.
func (ets ecsTaskSetAttributes) WaitUntilStable() terra.BoolValue {
	return terra.ReferenceAsBool(ets.ref.Append("wait_until_stable"))
}

// WaitUntilStableTimeout returns a reference to field wait_until_stable_timeout of aws_ecs_task_set.
func (ets ecsTaskSetAttributes) WaitUntilStableTimeout() terra.StringValue {
	return terra.ReferenceAsString(ets.ref.Append("wait_until_stable_timeout"))
}

func (ets ecsTaskSetAttributes) CapacityProviderStrategy() terra.SetValue[ecstaskset.CapacityProviderStrategyAttributes] {
	return terra.ReferenceAsSet[ecstaskset.CapacityProviderStrategyAttributes](ets.ref.Append("capacity_provider_strategy"))
}

func (ets ecsTaskSetAttributes) LoadBalancer() terra.SetValue[ecstaskset.LoadBalancerAttributes] {
	return terra.ReferenceAsSet[ecstaskset.LoadBalancerAttributes](ets.ref.Append("load_balancer"))
}

func (ets ecsTaskSetAttributes) NetworkConfiguration() terra.ListValue[ecstaskset.NetworkConfigurationAttributes] {
	return terra.ReferenceAsList[ecstaskset.NetworkConfigurationAttributes](ets.ref.Append("network_configuration"))
}

func (ets ecsTaskSetAttributes) Scale() terra.ListValue[ecstaskset.ScaleAttributes] {
	return terra.ReferenceAsList[ecstaskset.ScaleAttributes](ets.ref.Append("scale"))
}

func (ets ecsTaskSetAttributes) ServiceRegistries() terra.ListValue[ecstaskset.ServiceRegistriesAttributes] {
	return terra.ReferenceAsList[ecstaskset.ServiceRegistriesAttributes](ets.ref.Append("service_registries"))
}

type ecsTaskSetState struct {
	Arn                      string                                     `json:"arn"`
	Cluster                  string                                     `json:"cluster"`
	ExternalId               string                                     `json:"external_id"`
	ForceDelete              bool                                       `json:"force_delete"`
	Id                       string                                     `json:"id"`
	LaunchType               string                                     `json:"launch_type"`
	PlatformVersion          string                                     `json:"platform_version"`
	Service                  string                                     `json:"service"`
	StabilityStatus          string                                     `json:"stability_status"`
	Status                   string                                     `json:"status"`
	Tags                     map[string]string                          `json:"tags"`
	TagsAll                  map[string]string                          `json:"tags_all"`
	TaskDefinition           string                                     `json:"task_definition"`
	TaskSetId                string                                     `json:"task_set_id"`
	WaitUntilStable          bool                                       `json:"wait_until_stable"`
	WaitUntilStableTimeout   string                                     `json:"wait_until_stable_timeout"`
	CapacityProviderStrategy []ecstaskset.CapacityProviderStrategyState `json:"capacity_provider_strategy"`
	LoadBalancer             []ecstaskset.LoadBalancerState             `json:"load_balancer"`
	NetworkConfiguration     []ecstaskset.NetworkConfigurationState     `json:"network_configuration"`
	Scale                    []ecstaskset.ScaleState                    `json:"scale"`
	ServiceRegistries        []ecstaskset.ServiceRegistriesState        `json:"service_registries"`
}
