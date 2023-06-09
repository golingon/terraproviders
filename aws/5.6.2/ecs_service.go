// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ecsservice "github.com/golingon/terraproviders/aws/5.6.2/ecsservice"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEcsService creates a new instance of [EcsService].
func NewEcsService(name string, args EcsServiceArgs) *EcsService {
	return &EcsService{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EcsService)(nil)

// EcsService represents the Terraform resource aws_ecs_service.
type EcsService struct {
	Name      string
	Args      EcsServiceArgs
	state     *ecsServiceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EcsService].
func (es *EcsService) Type() string {
	return "aws_ecs_service"
}

// LocalName returns the local name for [EcsService].
func (es *EcsService) LocalName() string {
	return es.Name
}

// Configuration returns the configuration (args) for [EcsService].
func (es *EcsService) Configuration() interface{} {
	return es.Args
}

// DependOn is used for other resources to depend on [EcsService].
func (es *EcsService) DependOn() terra.Reference {
	return terra.ReferenceResource(es)
}

// Dependencies returns the list of resources [EcsService] depends_on.
func (es *EcsService) Dependencies() terra.Dependencies {
	return es.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EcsService].
func (es *EcsService) LifecycleManagement() *terra.Lifecycle {
	return es.Lifecycle
}

// Attributes returns the attributes for [EcsService].
func (es *EcsService) Attributes() ecsServiceAttributes {
	return ecsServiceAttributes{ref: terra.ReferenceResource(es)}
}

// ImportState imports the given attribute values into [EcsService]'s state.
func (es *EcsService) ImportState(av io.Reader) error {
	es.state = &ecsServiceState{}
	if err := json.NewDecoder(av).Decode(es.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", es.Type(), es.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EcsService] has state.
func (es *EcsService) State() (*ecsServiceState, bool) {
	return es.state, es.state != nil
}

// StateMust returns the state for [EcsService]. Panics if the state is nil.
func (es *EcsService) StateMust() *ecsServiceState {
	if es.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", es.Type(), es.LocalName()))
	}
	return es.state
}

// EcsServiceArgs contains the configurations for aws_ecs_service.
type EcsServiceArgs struct {
	// Cluster: string, optional
	Cluster terra.StringValue `hcl:"cluster,attr"`
	// DeploymentMaximumPercent: number, optional
	DeploymentMaximumPercent terra.NumberValue `hcl:"deployment_maximum_percent,attr"`
	// DeploymentMinimumHealthyPercent: number, optional
	DeploymentMinimumHealthyPercent terra.NumberValue `hcl:"deployment_minimum_healthy_percent,attr"`
	// DesiredCount: number, optional
	DesiredCount terra.NumberValue `hcl:"desired_count,attr"`
	// EnableEcsManagedTags: bool, optional
	EnableEcsManagedTags terra.BoolValue `hcl:"enable_ecs_managed_tags,attr"`
	// EnableExecuteCommand: bool, optional
	EnableExecuteCommand terra.BoolValue `hcl:"enable_execute_command,attr"`
	// ForceNewDeployment: bool, optional
	ForceNewDeployment terra.BoolValue `hcl:"force_new_deployment,attr"`
	// HealthCheckGracePeriodSeconds: number, optional
	HealthCheckGracePeriodSeconds terra.NumberValue `hcl:"health_check_grace_period_seconds,attr"`
	// IamRole: string, optional
	IamRole terra.StringValue `hcl:"iam_role,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LaunchType: string, optional
	LaunchType terra.StringValue `hcl:"launch_type,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PlatformVersion: string, optional
	PlatformVersion terra.StringValue `hcl:"platform_version,attr"`
	// PropagateTags: string, optional
	PropagateTags terra.StringValue `hcl:"propagate_tags,attr"`
	// SchedulingStrategy: string, optional
	SchedulingStrategy terra.StringValue `hcl:"scheduling_strategy,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TaskDefinition: string, optional
	TaskDefinition terra.StringValue `hcl:"task_definition,attr"`
	// Triggers: map of string, optional
	Triggers terra.MapValue[terra.StringValue] `hcl:"triggers,attr"`
	// WaitForSteadyState: bool, optional
	WaitForSteadyState terra.BoolValue `hcl:"wait_for_steady_state,attr"`
	// Alarms: optional
	Alarms *ecsservice.Alarms `hcl:"alarms,block"`
	// CapacityProviderStrategy: min=0
	CapacityProviderStrategy []ecsservice.CapacityProviderStrategy `hcl:"capacity_provider_strategy,block" validate:"min=0"`
	// DeploymentCircuitBreaker: optional
	DeploymentCircuitBreaker *ecsservice.DeploymentCircuitBreaker `hcl:"deployment_circuit_breaker,block"`
	// DeploymentController: optional
	DeploymentController *ecsservice.DeploymentController `hcl:"deployment_controller,block"`
	// LoadBalancer: min=0
	LoadBalancer []ecsservice.LoadBalancer `hcl:"load_balancer,block" validate:"min=0"`
	// NetworkConfiguration: optional
	NetworkConfiguration *ecsservice.NetworkConfiguration `hcl:"network_configuration,block"`
	// OrderedPlacementStrategy: min=0,max=5
	OrderedPlacementStrategy []ecsservice.OrderedPlacementStrategy `hcl:"ordered_placement_strategy,block" validate:"min=0,max=5"`
	// PlacementConstraints: min=0,max=10
	PlacementConstraints []ecsservice.PlacementConstraints `hcl:"placement_constraints,block" validate:"min=0,max=10"`
	// ServiceConnectConfiguration: optional
	ServiceConnectConfiguration *ecsservice.ServiceConnectConfiguration `hcl:"service_connect_configuration,block"`
	// ServiceRegistries: optional
	ServiceRegistries *ecsservice.ServiceRegistries `hcl:"service_registries,block"`
	// Timeouts: optional
	Timeouts *ecsservice.Timeouts `hcl:"timeouts,block"`
}
type ecsServiceAttributes struct {
	ref terra.Reference
}

// Cluster returns a reference to field cluster of aws_ecs_service.
func (es ecsServiceAttributes) Cluster() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("cluster"))
}

// DeploymentMaximumPercent returns a reference to field deployment_maximum_percent of aws_ecs_service.
func (es ecsServiceAttributes) DeploymentMaximumPercent() terra.NumberValue {
	return terra.ReferenceAsNumber(es.ref.Append("deployment_maximum_percent"))
}

// DeploymentMinimumHealthyPercent returns a reference to field deployment_minimum_healthy_percent of aws_ecs_service.
func (es ecsServiceAttributes) DeploymentMinimumHealthyPercent() terra.NumberValue {
	return terra.ReferenceAsNumber(es.ref.Append("deployment_minimum_healthy_percent"))
}

// DesiredCount returns a reference to field desired_count of aws_ecs_service.
func (es ecsServiceAttributes) DesiredCount() terra.NumberValue {
	return terra.ReferenceAsNumber(es.ref.Append("desired_count"))
}

// EnableEcsManagedTags returns a reference to field enable_ecs_managed_tags of aws_ecs_service.
func (es ecsServiceAttributes) EnableEcsManagedTags() terra.BoolValue {
	return terra.ReferenceAsBool(es.ref.Append("enable_ecs_managed_tags"))
}

// EnableExecuteCommand returns a reference to field enable_execute_command of aws_ecs_service.
func (es ecsServiceAttributes) EnableExecuteCommand() terra.BoolValue {
	return terra.ReferenceAsBool(es.ref.Append("enable_execute_command"))
}

// ForceNewDeployment returns a reference to field force_new_deployment of aws_ecs_service.
func (es ecsServiceAttributes) ForceNewDeployment() terra.BoolValue {
	return terra.ReferenceAsBool(es.ref.Append("force_new_deployment"))
}

// HealthCheckGracePeriodSeconds returns a reference to field health_check_grace_period_seconds of aws_ecs_service.
func (es ecsServiceAttributes) HealthCheckGracePeriodSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(es.ref.Append("health_check_grace_period_seconds"))
}

// IamRole returns a reference to field iam_role of aws_ecs_service.
func (es ecsServiceAttributes) IamRole() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("iam_role"))
}

// Id returns a reference to field id of aws_ecs_service.
func (es ecsServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("id"))
}

// LaunchType returns a reference to field launch_type of aws_ecs_service.
func (es ecsServiceAttributes) LaunchType() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("launch_type"))
}

// Name returns a reference to field name of aws_ecs_service.
func (es ecsServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("name"))
}

// PlatformVersion returns a reference to field platform_version of aws_ecs_service.
func (es ecsServiceAttributes) PlatformVersion() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("platform_version"))
}

// PropagateTags returns a reference to field propagate_tags of aws_ecs_service.
func (es ecsServiceAttributes) PropagateTags() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("propagate_tags"))
}

// SchedulingStrategy returns a reference to field scheduling_strategy of aws_ecs_service.
func (es ecsServiceAttributes) SchedulingStrategy() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("scheduling_strategy"))
}

// Tags returns a reference to field tags of aws_ecs_service.
func (es ecsServiceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](es.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_ecs_service.
func (es ecsServiceAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](es.ref.Append("tags_all"))
}

// TaskDefinition returns a reference to field task_definition of aws_ecs_service.
func (es ecsServiceAttributes) TaskDefinition() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("task_definition"))
}

// Triggers returns a reference to field triggers of aws_ecs_service.
func (es ecsServiceAttributes) Triggers() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](es.ref.Append("triggers"))
}

// WaitForSteadyState returns a reference to field wait_for_steady_state of aws_ecs_service.
func (es ecsServiceAttributes) WaitForSteadyState() terra.BoolValue {
	return terra.ReferenceAsBool(es.ref.Append("wait_for_steady_state"))
}

func (es ecsServiceAttributes) Alarms() terra.ListValue[ecsservice.AlarmsAttributes] {
	return terra.ReferenceAsList[ecsservice.AlarmsAttributes](es.ref.Append("alarms"))
}

func (es ecsServiceAttributes) CapacityProviderStrategy() terra.SetValue[ecsservice.CapacityProviderStrategyAttributes] {
	return terra.ReferenceAsSet[ecsservice.CapacityProviderStrategyAttributes](es.ref.Append("capacity_provider_strategy"))
}

func (es ecsServiceAttributes) DeploymentCircuitBreaker() terra.ListValue[ecsservice.DeploymentCircuitBreakerAttributes] {
	return terra.ReferenceAsList[ecsservice.DeploymentCircuitBreakerAttributes](es.ref.Append("deployment_circuit_breaker"))
}

func (es ecsServiceAttributes) DeploymentController() terra.ListValue[ecsservice.DeploymentControllerAttributes] {
	return terra.ReferenceAsList[ecsservice.DeploymentControllerAttributes](es.ref.Append("deployment_controller"))
}

func (es ecsServiceAttributes) LoadBalancer() terra.SetValue[ecsservice.LoadBalancerAttributes] {
	return terra.ReferenceAsSet[ecsservice.LoadBalancerAttributes](es.ref.Append("load_balancer"))
}

func (es ecsServiceAttributes) NetworkConfiguration() terra.ListValue[ecsservice.NetworkConfigurationAttributes] {
	return terra.ReferenceAsList[ecsservice.NetworkConfigurationAttributes](es.ref.Append("network_configuration"))
}

func (es ecsServiceAttributes) OrderedPlacementStrategy() terra.ListValue[ecsservice.OrderedPlacementStrategyAttributes] {
	return terra.ReferenceAsList[ecsservice.OrderedPlacementStrategyAttributes](es.ref.Append("ordered_placement_strategy"))
}

func (es ecsServiceAttributes) PlacementConstraints() terra.SetValue[ecsservice.PlacementConstraintsAttributes] {
	return terra.ReferenceAsSet[ecsservice.PlacementConstraintsAttributes](es.ref.Append("placement_constraints"))
}

func (es ecsServiceAttributes) ServiceConnectConfiguration() terra.ListValue[ecsservice.ServiceConnectConfigurationAttributes] {
	return terra.ReferenceAsList[ecsservice.ServiceConnectConfigurationAttributes](es.ref.Append("service_connect_configuration"))
}

func (es ecsServiceAttributes) ServiceRegistries() terra.ListValue[ecsservice.ServiceRegistriesAttributes] {
	return terra.ReferenceAsList[ecsservice.ServiceRegistriesAttributes](es.ref.Append("service_registries"))
}

func (es ecsServiceAttributes) Timeouts() ecsservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[ecsservice.TimeoutsAttributes](es.ref.Append("timeouts"))
}

type ecsServiceState struct {
	Cluster                         string                                        `json:"cluster"`
	DeploymentMaximumPercent        float64                                       `json:"deployment_maximum_percent"`
	DeploymentMinimumHealthyPercent float64                                       `json:"deployment_minimum_healthy_percent"`
	DesiredCount                    float64                                       `json:"desired_count"`
	EnableEcsManagedTags            bool                                          `json:"enable_ecs_managed_tags"`
	EnableExecuteCommand            bool                                          `json:"enable_execute_command"`
	ForceNewDeployment              bool                                          `json:"force_new_deployment"`
	HealthCheckGracePeriodSeconds   float64                                       `json:"health_check_grace_period_seconds"`
	IamRole                         string                                        `json:"iam_role"`
	Id                              string                                        `json:"id"`
	LaunchType                      string                                        `json:"launch_type"`
	Name                            string                                        `json:"name"`
	PlatformVersion                 string                                        `json:"platform_version"`
	PropagateTags                   string                                        `json:"propagate_tags"`
	SchedulingStrategy              string                                        `json:"scheduling_strategy"`
	Tags                            map[string]string                             `json:"tags"`
	TagsAll                         map[string]string                             `json:"tags_all"`
	TaskDefinition                  string                                        `json:"task_definition"`
	Triggers                        map[string]string                             `json:"triggers"`
	WaitForSteadyState              bool                                          `json:"wait_for_steady_state"`
	Alarms                          []ecsservice.AlarmsState                      `json:"alarms"`
	CapacityProviderStrategy        []ecsservice.CapacityProviderStrategyState    `json:"capacity_provider_strategy"`
	DeploymentCircuitBreaker        []ecsservice.DeploymentCircuitBreakerState    `json:"deployment_circuit_breaker"`
	DeploymentController            []ecsservice.DeploymentControllerState        `json:"deployment_controller"`
	LoadBalancer                    []ecsservice.LoadBalancerState                `json:"load_balancer"`
	NetworkConfiguration            []ecsservice.NetworkConfigurationState        `json:"network_configuration"`
	OrderedPlacementStrategy        []ecsservice.OrderedPlacementStrategyState    `json:"ordered_placement_strategy"`
	PlacementConstraints            []ecsservice.PlacementConstraintsState        `json:"placement_constraints"`
	ServiceConnectConfiguration     []ecsservice.ServiceConnectConfigurationState `json:"service_connect_configuration"`
	ServiceRegistries               []ecsservice.ServiceRegistriesState           `json:"service_registries"`
	Timeouts                        *ecsservice.TimeoutsState                     `json:"timeouts"`
}
