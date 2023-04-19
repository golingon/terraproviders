// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	autoscalinggroup "github.com/golingon/terraproviders/aws/4.63.0/autoscalinggroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAutoscalingGroup creates a new instance of [AutoscalingGroup].
func NewAutoscalingGroup(name string, args AutoscalingGroupArgs) *AutoscalingGroup {
	return &AutoscalingGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AutoscalingGroup)(nil)

// AutoscalingGroup represents the Terraform resource aws_autoscaling_group.
type AutoscalingGroup struct {
	Name      string
	Args      AutoscalingGroupArgs
	state     *autoscalingGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AutoscalingGroup].
func (ag *AutoscalingGroup) Type() string {
	return "aws_autoscaling_group"
}

// LocalName returns the local name for [AutoscalingGroup].
func (ag *AutoscalingGroup) LocalName() string {
	return ag.Name
}

// Configuration returns the configuration (args) for [AutoscalingGroup].
func (ag *AutoscalingGroup) Configuration() interface{} {
	return ag.Args
}

// DependOn is used for other resources to depend on [AutoscalingGroup].
func (ag *AutoscalingGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(ag)
}

// Dependencies returns the list of resources [AutoscalingGroup] depends_on.
func (ag *AutoscalingGroup) Dependencies() terra.Dependencies {
	return ag.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AutoscalingGroup].
func (ag *AutoscalingGroup) LifecycleManagement() *terra.Lifecycle {
	return ag.Lifecycle
}

// Attributes returns the attributes for [AutoscalingGroup].
func (ag *AutoscalingGroup) Attributes() autoscalingGroupAttributes {
	return autoscalingGroupAttributes{ref: terra.ReferenceResource(ag)}
}

// ImportState imports the given attribute values into [AutoscalingGroup]'s state.
func (ag *AutoscalingGroup) ImportState(av io.Reader) error {
	ag.state = &autoscalingGroupState{}
	if err := json.NewDecoder(av).Decode(ag.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ag.Type(), ag.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AutoscalingGroup] has state.
func (ag *AutoscalingGroup) State() (*autoscalingGroupState, bool) {
	return ag.state, ag.state != nil
}

// StateMust returns the state for [AutoscalingGroup]. Panics if the state is nil.
func (ag *AutoscalingGroup) StateMust() *autoscalingGroupState {
	if ag.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ag.Type(), ag.LocalName()))
	}
	return ag.state
}

// AutoscalingGroupArgs contains the configurations for aws_autoscaling_group.
type AutoscalingGroupArgs struct {
	// AvailabilityZones: set of string, optional
	AvailabilityZones terra.SetValue[terra.StringValue] `hcl:"availability_zones,attr"`
	// CapacityRebalance: bool, optional
	CapacityRebalance terra.BoolValue `hcl:"capacity_rebalance,attr"`
	// Context: string, optional
	Context terra.StringValue `hcl:"context,attr"`
	// DefaultCooldown: number, optional
	DefaultCooldown terra.NumberValue `hcl:"default_cooldown,attr"`
	// DefaultInstanceWarmup: number, optional
	DefaultInstanceWarmup terra.NumberValue `hcl:"default_instance_warmup,attr"`
	// DesiredCapacity: number, optional
	DesiredCapacity terra.NumberValue `hcl:"desired_capacity,attr"`
	// DesiredCapacityType: string, optional
	DesiredCapacityType terra.StringValue `hcl:"desired_capacity_type,attr"`
	// EnabledMetrics: set of string, optional
	EnabledMetrics terra.SetValue[terra.StringValue] `hcl:"enabled_metrics,attr"`
	// ForceDelete: bool, optional
	ForceDelete terra.BoolValue `hcl:"force_delete,attr"`
	// ForceDeleteWarmPool: bool, optional
	ForceDeleteWarmPool terra.BoolValue `hcl:"force_delete_warm_pool,attr"`
	// HealthCheckGracePeriod: number, optional
	HealthCheckGracePeriod terra.NumberValue `hcl:"health_check_grace_period,attr"`
	// HealthCheckType: string, optional
	HealthCheckType terra.StringValue `hcl:"health_check_type,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LaunchConfiguration: string, optional
	LaunchConfiguration terra.StringValue `hcl:"launch_configuration,attr"`
	// LoadBalancers: set of string, optional
	LoadBalancers terra.SetValue[terra.StringValue] `hcl:"load_balancers,attr"`
	// MaxInstanceLifetime: number, optional
	MaxInstanceLifetime terra.NumberValue `hcl:"max_instance_lifetime,attr"`
	// MaxSize: number, required
	MaxSize terra.NumberValue `hcl:"max_size,attr" validate:"required"`
	// MetricsGranularity: string, optional
	MetricsGranularity terra.StringValue `hcl:"metrics_granularity,attr"`
	// MinElbCapacity: number, optional
	MinElbCapacity terra.NumberValue `hcl:"min_elb_capacity,attr"`
	// MinSize: number, required
	MinSize terra.NumberValue `hcl:"min_size,attr" validate:"required"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NamePrefix: string, optional
	NamePrefix terra.StringValue `hcl:"name_prefix,attr"`
	// PlacementGroup: string, optional
	PlacementGroup terra.StringValue `hcl:"placement_group,attr"`
	// ProtectFromScaleIn: bool, optional
	ProtectFromScaleIn terra.BoolValue `hcl:"protect_from_scale_in,attr"`
	// ServiceLinkedRoleArn: string, optional
	ServiceLinkedRoleArn terra.StringValue `hcl:"service_linked_role_arn,attr"`
	// SuspendedProcesses: set of string, optional
	SuspendedProcesses terra.SetValue[terra.StringValue] `hcl:"suspended_processes,attr"`
	// Tags: set of map of string, optional
	Tags terra.SetValue[terra.MapValue[terra.StringValue]] `hcl:"tags,attr"`
	// TargetGroupArns: set of string, optional
	TargetGroupArns terra.SetValue[terra.StringValue] `hcl:"target_group_arns,attr"`
	// TerminationPolicies: list of string, optional
	TerminationPolicies terra.ListValue[terra.StringValue] `hcl:"termination_policies,attr"`
	// VpcZoneIdentifier: set of string, optional
	VpcZoneIdentifier terra.SetValue[terra.StringValue] `hcl:"vpc_zone_identifier,attr"`
	// WaitForCapacityTimeout: string, optional
	WaitForCapacityTimeout terra.StringValue `hcl:"wait_for_capacity_timeout,attr"`
	// WaitForElbCapacity: number, optional
	WaitForElbCapacity terra.NumberValue `hcl:"wait_for_elb_capacity,attr"`
	// InitialLifecycleHook: min=0
	InitialLifecycleHook []autoscalinggroup.InitialLifecycleHook `hcl:"initial_lifecycle_hook,block" validate:"min=0"`
	// InstanceRefresh: optional
	InstanceRefresh *autoscalinggroup.InstanceRefresh `hcl:"instance_refresh,block"`
	// LaunchTemplate: optional
	LaunchTemplate *autoscalinggroup.LaunchTemplate `hcl:"launch_template,block"`
	// MixedInstancesPolicy: optional
	MixedInstancesPolicy *autoscalinggroup.MixedInstancesPolicy `hcl:"mixed_instances_policy,block"`
	// Tag: min=0
	Tag []autoscalinggroup.Tag `hcl:"tag,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *autoscalinggroup.Timeouts `hcl:"timeouts,block"`
	// WarmPool: optional
	WarmPool *autoscalinggroup.WarmPool `hcl:"warm_pool,block"`
}
type autoscalingGroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_autoscaling_group.
func (ag autoscalingGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ag.ref.Append("arn"))
}

// AvailabilityZones returns a reference to field availability_zones of aws_autoscaling_group.
func (ag autoscalingGroupAttributes) AvailabilityZones() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ag.ref.Append("availability_zones"))
}

// CapacityRebalance returns a reference to field capacity_rebalance of aws_autoscaling_group.
func (ag autoscalingGroupAttributes) CapacityRebalance() terra.BoolValue {
	return terra.ReferenceAsBool(ag.ref.Append("capacity_rebalance"))
}

// Context returns a reference to field context of aws_autoscaling_group.
func (ag autoscalingGroupAttributes) Context() terra.StringValue {
	return terra.ReferenceAsString(ag.ref.Append("context"))
}

// DefaultCooldown returns a reference to field default_cooldown of aws_autoscaling_group.
func (ag autoscalingGroupAttributes) DefaultCooldown() terra.NumberValue {
	return terra.ReferenceAsNumber(ag.ref.Append("default_cooldown"))
}

// DefaultInstanceWarmup returns a reference to field default_instance_warmup of aws_autoscaling_group.
func (ag autoscalingGroupAttributes) DefaultInstanceWarmup() terra.NumberValue {
	return terra.ReferenceAsNumber(ag.ref.Append("default_instance_warmup"))
}

// DesiredCapacity returns a reference to field desired_capacity of aws_autoscaling_group.
func (ag autoscalingGroupAttributes) DesiredCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(ag.ref.Append("desired_capacity"))
}

// DesiredCapacityType returns a reference to field desired_capacity_type of aws_autoscaling_group.
func (ag autoscalingGroupAttributes) DesiredCapacityType() terra.StringValue {
	return terra.ReferenceAsString(ag.ref.Append("desired_capacity_type"))
}

// EnabledMetrics returns a reference to field enabled_metrics of aws_autoscaling_group.
func (ag autoscalingGroupAttributes) EnabledMetrics() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ag.ref.Append("enabled_metrics"))
}

// ForceDelete returns a reference to field force_delete of aws_autoscaling_group.
func (ag autoscalingGroupAttributes) ForceDelete() terra.BoolValue {
	return terra.ReferenceAsBool(ag.ref.Append("force_delete"))
}

// ForceDeleteWarmPool returns a reference to field force_delete_warm_pool of aws_autoscaling_group.
func (ag autoscalingGroupAttributes) ForceDeleteWarmPool() terra.BoolValue {
	return terra.ReferenceAsBool(ag.ref.Append("force_delete_warm_pool"))
}

// HealthCheckGracePeriod returns a reference to field health_check_grace_period of aws_autoscaling_group.
func (ag autoscalingGroupAttributes) HealthCheckGracePeriod() terra.NumberValue {
	return terra.ReferenceAsNumber(ag.ref.Append("health_check_grace_period"))
}

// HealthCheckType returns a reference to field health_check_type of aws_autoscaling_group.
func (ag autoscalingGroupAttributes) HealthCheckType() terra.StringValue {
	return terra.ReferenceAsString(ag.ref.Append("health_check_type"))
}

// Id returns a reference to field id of aws_autoscaling_group.
func (ag autoscalingGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ag.ref.Append("id"))
}

// LaunchConfiguration returns a reference to field launch_configuration of aws_autoscaling_group.
func (ag autoscalingGroupAttributes) LaunchConfiguration() terra.StringValue {
	return terra.ReferenceAsString(ag.ref.Append("launch_configuration"))
}

// LoadBalancers returns a reference to field load_balancers of aws_autoscaling_group.
func (ag autoscalingGroupAttributes) LoadBalancers() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ag.ref.Append("load_balancers"))
}

// MaxInstanceLifetime returns a reference to field max_instance_lifetime of aws_autoscaling_group.
func (ag autoscalingGroupAttributes) MaxInstanceLifetime() terra.NumberValue {
	return terra.ReferenceAsNumber(ag.ref.Append("max_instance_lifetime"))
}

// MaxSize returns a reference to field max_size of aws_autoscaling_group.
func (ag autoscalingGroupAttributes) MaxSize() terra.NumberValue {
	return terra.ReferenceAsNumber(ag.ref.Append("max_size"))
}

// MetricsGranularity returns a reference to field metrics_granularity of aws_autoscaling_group.
func (ag autoscalingGroupAttributes) MetricsGranularity() terra.StringValue {
	return terra.ReferenceAsString(ag.ref.Append("metrics_granularity"))
}

// MinElbCapacity returns a reference to field min_elb_capacity of aws_autoscaling_group.
func (ag autoscalingGroupAttributes) MinElbCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(ag.ref.Append("min_elb_capacity"))
}

// MinSize returns a reference to field min_size of aws_autoscaling_group.
func (ag autoscalingGroupAttributes) MinSize() terra.NumberValue {
	return terra.ReferenceAsNumber(ag.ref.Append("min_size"))
}

// Name returns a reference to field name of aws_autoscaling_group.
func (ag autoscalingGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ag.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_autoscaling_group.
func (ag autoscalingGroupAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(ag.ref.Append("name_prefix"))
}

// PlacementGroup returns a reference to field placement_group of aws_autoscaling_group.
func (ag autoscalingGroupAttributes) PlacementGroup() terra.StringValue {
	return terra.ReferenceAsString(ag.ref.Append("placement_group"))
}

// ProtectFromScaleIn returns a reference to field protect_from_scale_in of aws_autoscaling_group.
func (ag autoscalingGroupAttributes) ProtectFromScaleIn() terra.BoolValue {
	return terra.ReferenceAsBool(ag.ref.Append("protect_from_scale_in"))
}

// ServiceLinkedRoleArn returns a reference to field service_linked_role_arn of aws_autoscaling_group.
func (ag autoscalingGroupAttributes) ServiceLinkedRoleArn() terra.StringValue {
	return terra.ReferenceAsString(ag.ref.Append("service_linked_role_arn"))
}

// SuspendedProcesses returns a reference to field suspended_processes of aws_autoscaling_group.
func (ag autoscalingGroupAttributes) SuspendedProcesses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ag.ref.Append("suspended_processes"))
}

// Tags returns a reference to field tags of aws_autoscaling_group.
func (ag autoscalingGroupAttributes) Tags() terra.SetValue[terra.MapValue[terra.StringValue]] {
	return terra.ReferenceAsSet[terra.MapValue[terra.StringValue]](ag.ref.Append("tags"))
}

// TargetGroupArns returns a reference to field target_group_arns of aws_autoscaling_group.
func (ag autoscalingGroupAttributes) TargetGroupArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ag.ref.Append("target_group_arns"))
}

// TerminationPolicies returns a reference to field termination_policies of aws_autoscaling_group.
func (ag autoscalingGroupAttributes) TerminationPolicies() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ag.ref.Append("termination_policies"))
}

// VpcZoneIdentifier returns a reference to field vpc_zone_identifier of aws_autoscaling_group.
func (ag autoscalingGroupAttributes) VpcZoneIdentifier() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ag.ref.Append("vpc_zone_identifier"))
}

// WaitForCapacityTimeout returns a reference to field wait_for_capacity_timeout of aws_autoscaling_group.
func (ag autoscalingGroupAttributes) WaitForCapacityTimeout() terra.StringValue {
	return terra.ReferenceAsString(ag.ref.Append("wait_for_capacity_timeout"))
}

// WaitForElbCapacity returns a reference to field wait_for_elb_capacity of aws_autoscaling_group.
func (ag autoscalingGroupAttributes) WaitForElbCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(ag.ref.Append("wait_for_elb_capacity"))
}

func (ag autoscalingGroupAttributes) InitialLifecycleHook() terra.SetValue[autoscalinggroup.InitialLifecycleHookAttributes] {
	return terra.ReferenceAsSet[autoscalinggroup.InitialLifecycleHookAttributes](ag.ref.Append("initial_lifecycle_hook"))
}

func (ag autoscalingGroupAttributes) InstanceRefresh() terra.ListValue[autoscalinggroup.InstanceRefreshAttributes] {
	return terra.ReferenceAsList[autoscalinggroup.InstanceRefreshAttributes](ag.ref.Append("instance_refresh"))
}

func (ag autoscalingGroupAttributes) LaunchTemplate() terra.ListValue[autoscalinggroup.LaunchTemplateAttributes] {
	return terra.ReferenceAsList[autoscalinggroup.LaunchTemplateAttributes](ag.ref.Append("launch_template"))
}

func (ag autoscalingGroupAttributes) MixedInstancesPolicy() terra.ListValue[autoscalinggroup.MixedInstancesPolicyAttributes] {
	return terra.ReferenceAsList[autoscalinggroup.MixedInstancesPolicyAttributes](ag.ref.Append("mixed_instances_policy"))
}

func (ag autoscalingGroupAttributes) Tag() terra.SetValue[autoscalinggroup.TagAttributes] {
	return terra.ReferenceAsSet[autoscalinggroup.TagAttributes](ag.ref.Append("tag"))
}

func (ag autoscalingGroupAttributes) Timeouts() autoscalinggroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[autoscalinggroup.TimeoutsAttributes](ag.ref.Append("timeouts"))
}

func (ag autoscalingGroupAttributes) WarmPool() terra.ListValue[autoscalinggroup.WarmPoolAttributes] {
	return terra.ReferenceAsList[autoscalinggroup.WarmPoolAttributes](ag.ref.Append("warm_pool"))
}

type autoscalingGroupState struct {
	Arn                    string                                       `json:"arn"`
	AvailabilityZones      []string                                     `json:"availability_zones"`
	CapacityRebalance      bool                                         `json:"capacity_rebalance"`
	Context                string                                       `json:"context"`
	DefaultCooldown        float64                                      `json:"default_cooldown"`
	DefaultInstanceWarmup  float64                                      `json:"default_instance_warmup"`
	DesiredCapacity        float64                                      `json:"desired_capacity"`
	DesiredCapacityType    string                                       `json:"desired_capacity_type"`
	EnabledMetrics         []string                                     `json:"enabled_metrics"`
	ForceDelete            bool                                         `json:"force_delete"`
	ForceDeleteWarmPool    bool                                         `json:"force_delete_warm_pool"`
	HealthCheckGracePeriod float64                                      `json:"health_check_grace_period"`
	HealthCheckType        string                                       `json:"health_check_type"`
	Id                     string                                       `json:"id"`
	LaunchConfiguration    string                                       `json:"launch_configuration"`
	LoadBalancers          []string                                     `json:"load_balancers"`
	MaxInstanceLifetime    float64                                      `json:"max_instance_lifetime"`
	MaxSize                float64                                      `json:"max_size"`
	MetricsGranularity     string                                       `json:"metrics_granularity"`
	MinElbCapacity         float64                                      `json:"min_elb_capacity"`
	MinSize                float64                                      `json:"min_size"`
	Name                   string                                       `json:"name"`
	NamePrefix             string                                       `json:"name_prefix"`
	PlacementGroup         string                                       `json:"placement_group"`
	ProtectFromScaleIn     bool                                         `json:"protect_from_scale_in"`
	ServiceLinkedRoleArn   string                                       `json:"service_linked_role_arn"`
	SuspendedProcesses     []string                                     `json:"suspended_processes"`
	Tags                   []map[string]string                          `json:"tags"`
	TargetGroupArns        []string                                     `json:"target_group_arns"`
	TerminationPolicies    []string                                     `json:"termination_policies"`
	VpcZoneIdentifier      []string                                     `json:"vpc_zone_identifier"`
	WaitForCapacityTimeout string                                       `json:"wait_for_capacity_timeout"`
	WaitForElbCapacity     float64                                      `json:"wait_for_elb_capacity"`
	InitialLifecycleHook   []autoscalinggroup.InitialLifecycleHookState `json:"initial_lifecycle_hook"`
	InstanceRefresh        []autoscalinggroup.InstanceRefreshState      `json:"instance_refresh"`
	LaunchTemplate         []autoscalinggroup.LaunchTemplateState       `json:"launch_template"`
	MixedInstancesPolicy   []autoscalinggroup.MixedInstancesPolicyState `json:"mixed_instances_policy"`
	Tag                    []autoscalinggroup.TagState                  `json:"tag"`
	Timeouts               *autoscalinggroup.TimeoutsState              `json:"timeouts"`
	WarmPool               []autoscalinggroup.WarmPoolState             `json:"warm_pool"`
}
