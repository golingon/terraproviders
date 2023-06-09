// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataautoscalinggroup "github.com/golingon/terraproviders/aws/5.6.2/dataautoscalinggroup"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataAutoscalingGroup creates a new instance of [DataAutoscalingGroup].
func NewDataAutoscalingGroup(name string, args DataAutoscalingGroupArgs) *DataAutoscalingGroup {
	return &DataAutoscalingGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAutoscalingGroup)(nil)

// DataAutoscalingGroup represents the Terraform data resource aws_autoscaling_group.
type DataAutoscalingGroup struct {
	Name string
	Args DataAutoscalingGroupArgs
}

// DataSource returns the Terraform object type for [DataAutoscalingGroup].
func (ag *DataAutoscalingGroup) DataSource() string {
	return "aws_autoscaling_group"
}

// LocalName returns the local name for [DataAutoscalingGroup].
func (ag *DataAutoscalingGroup) LocalName() string {
	return ag.Name
}

// Configuration returns the configuration (args) for [DataAutoscalingGroup].
func (ag *DataAutoscalingGroup) Configuration() interface{} {
	return ag.Args
}

// Attributes returns the attributes for [DataAutoscalingGroup].
func (ag *DataAutoscalingGroup) Attributes() dataAutoscalingGroupAttributes {
	return dataAutoscalingGroupAttributes{ref: terra.ReferenceDataResource(ag)}
}

// DataAutoscalingGroupArgs contains the configurations for aws_autoscaling_group.
type DataAutoscalingGroupArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// LaunchTemplate: min=0
	LaunchTemplate []dataautoscalinggroup.LaunchTemplate `hcl:"launch_template,block" validate:"min=0"`
	// MixedInstancesPolicy: min=0
	MixedInstancesPolicy []dataautoscalinggroup.MixedInstancesPolicy `hcl:"mixed_instances_policy,block" validate:"min=0"`
	// Tag: min=0
	Tag []dataautoscalinggroup.Tag `hcl:"tag,block" validate:"min=0"`
	// TrafficSource: min=0
	TrafficSource []dataautoscalinggroup.TrafficSource `hcl:"traffic_source,block" validate:"min=0"`
	// WarmPool: min=0
	WarmPool []dataautoscalinggroup.WarmPool `hcl:"warm_pool,block" validate:"min=0"`
}
type dataAutoscalingGroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_autoscaling_group.
func (ag dataAutoscalingGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ag.ref.Append("arn"))
}

// AvailabilityZones returns a reference to field availability_zones of aws_autoscaling_group.
func (ag dataAutoscalingGroupAttributes) AvailabilityZones() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ag.ref.Append("availability_zones"))
}

// DefaultCooldown returns a reference to field default_cooldown of aws_autoscaling_group.
func (ag dataAutoscalingGroupAttributes) DefaultCooldown() terra.NumberValue {
	return terra.ReferenceAsNumber(ag.ref.Append("default_cooldown"))
}

// DesiredCapacity returns a reference to field desired_capacity of aws_autoscaling_group.
func (ag dataAutoscalingGroupAttributes) DesiredCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(ag.ref.Append("desired_capacity"))
}

// DesiredCapacityType returns a reference to field desired_capacity_type of aws_autoscaling_group.
func (ag dataAutoscalingGroupAttributes) DesiredCapacityType() terra.StringValue {
	return terra.ReferenceAsString(ag.ref.Append("desired_capacity_type"))
}

// EnabledMetrics returns a reference to field enabled_metrics of aws_autoscaling_group.
func (ag dataAutoscalingGroupAttributes) EnabledMetrics() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ag.ref.Append("enabled_metrics"))
}

// HealthCheckGracePeriod returns a reference to field health_check_grace_period of aws_autoscaling_group.
func (ag dataAutoscalingGroupAttributes) HealthCheckGracePeriod() terra.NumberValue {
	return terra.ReferenceAsNumber(ag.ref.Append("health_check_grace_period"))
}

// HealthCheckType returns a reference to field health_check_type of aws_autoscaling_group.
func (ag dataAutoscalingGroupAttributes) HealthCheckType() terra.StringValue {
	return terra.ReferenceAsString(ag.ref.Append("health_check_type"))
}

// Id returns a reference to field id of aws_autoscaling_group.
func (ag dataAutoscalingGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ag.ref.Append("id"))
}

// LaunchConfiguration returns a reference to field launch_configuration of aws_autoscaling_group.
func (ag dataAutoscalingGroupAttributes) LaunchConfiguration() terra.StringValue {
	return terra.ReferenceAsString(ag.ref.Append("launch_configuration"))
}

// LoadBalancers returns a reference to field load_balancers of aws_autoscaling_group.
func (ag dataAutoscalingGroupAttributes) LoadBalancers() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ag.ref.Append("load_balancers"))
}

// MaxInstanceLifetime returns a reference to field max_instance_lifetime of aws_autoscaling_group.
func (ag dataAutoscalingGroupAttributes) MaxInstanceLifetime() terra.NumberValue {
	return terra.ReferenceAsNumber(ag.ref.Append("max_instance_lifetime"))
}

// MaxSize returns a reference to field max_size of aws_autoscaling_group.
func (ag dataAutoscalingGroupAttributes) MaxSize() terra.NumberValue {
	return terra.ReferenceAsNumber(ag.ref.Append("max_size"))
}

// MinSize returns a reference to field min_size of aws_autoscaling_group.
func (ag dataAutoscalingGroupAttributes) MinSize() terra.NumberValue {
	return terra.ReferenceAsNumber(ag.ref.Append("min_size"))
}

// Name returns a reference to field name of aws_autoscaling_group.
func (ag dataAutoscalingGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ag.ref.Append("name"))
}

// NewInstancesProtectedFromScaleIn returns a reference to field new_instances_protected_from_scale_in of aws_autoscaling_group.
func (ag dataAutoscalingGroupAttributes) NewInstancesProtectedFromScaleIn() terra.BoolValue {
	return terra.ReferenceAsBool(ag.ref.Append("new_instances_protected_from_scale_in"))
}

// PlacementGroup returns a reference to field placement_group of aws_autoscaling_group.
func (ag dataAutoscalingGroupAttributes) PlacementGroup() terra.StringValue {
	return terra.ReferenceAsString(ag.ref.Append("placement_group"))
}

// PredictedCapacity returns a reference to field predicted_capacity of aws_autoscaling_group.
func (ag dataAutoscalingGroupAttributes) PredictedCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(ag.ref.Append("predicted_capacity"))
}

// ServiceLinkedRoleArn returns a reference to field service_linked_role_arn of aws_autoscaling_group.
func (ag dataAutoscalingGroupAttributes) ServiceLinkedRoleArn() terra.StringValue {
	return terra.ReferenceAsString(ag.ref.Append("service_linked_role_arn"))
}

// Status returns a reference to field status of aws_autoscaling_group.
func (ag dataAutoscalingGroupAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(ag.ref.Append("status"))
}

// SuspendedProcesses returns a reference to field suspended_processes of aws_autoscaling_group.
func (ag dataAutoscalingGroupAttributes) SuspendedProcesses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ag.ref.Append("suspended_processes"))
}

// TargetGroupArns returns a reference to field target_group_arns of aws_autoscaling_group.
func (ag dataAutoscalingGroupAttributes) TargetGroupArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ag.ref.Append("target_group_arns"))
}

// TerminationPolicies returns a reference to field termination_policies of aws_autoscaling_group.
func (ag dataAutoscalingGroupAttributes) TerminationPolicies() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ag.ref.Append("termination_policies"))
}

// VpcZoneIdentifier returns a reference to field vpc_zone_identifier of aws_autoscaling_group.
func (ag dataAutoscalingGroupAttributes) VpcZoneIdentifier() terra.StringValue {
	return terra.ReferenceAsString(ag.ref.Append("vpc_zone_identifier"))
}

// WarmPoolSize returns a reference to field warm_pool_size of aws_autoscaling_group.
func (ag dataAutoscalingGroupAttributes) WarmPoolSize() terra.NumberValue {
	return terra.ReferenceAsNumber(ag.ref.Append("warm_pool_size"))
}

func (ag dataAutoscalingGroupAttributes) LaunchTemplate() terra.ListValue[dataautoscalinggroup.LaunchTemplateAttributes] {
	return terra.ReferenceAsList[dataautoscalinggroup.LaunchTemplateAttributes](ag.ref.Append("launch_template"))
}

func (ag dataAutoscalingGroupAttributes) MixedInstancesPolicy() terra.ListValue[dataautoscalinggroup.MixedInstancesPolicyAttributes] {
	return terra.ReferenceAsList[dataautoscalinggroup.MixedInstancesPolicyAttributes](ag.ref.Append("mixed_instances_policy"))
}

func (ag dataAutoscalingGroupAttributes) Tag() terra.SetValue[dataautoscalinggroup.TagAttributes] {
	return terra.ReferenceAsSet[dataautoscalinggroup.TagAttributes](ag.ref.Append("tag"))
}

func (ag dataAutoscalingGroupAttributes) TrafficSource() terra.SetValue[dataautoscalinggroup.TrafficSourceAttributes] {
	return terra.ReferenceAsSet[dataautoscalinggroup.TrafficSourceAttributes](ag.ref.Append("traffic_source"))
}

func (ag dataAutoscalingGroupAttributes) WarmPool() terra.ListValue[dataautoscalinggroup.WarmPoolAttributes] {
	return terra.ReferenceAsList[dataautoscalinggroup.WarmPoolAttributes](ag.ref.Append("warm_pool"))
}
