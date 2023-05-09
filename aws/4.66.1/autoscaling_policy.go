// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	autoscalingpolicy "github.com/golingon/terraproviders/aws/4.66.1/autoscalingpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAutoscalingPolicy creates a new instance of [AutoscalingPolicy].
func NewAutoscalingPolicy(name string, args AutoscalingPolicyArgs) *AutoscalingPolicy {
	return &AutoscalingPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AutoscalingPolicy)(nil)

// AutoscalingPolicy represents the Terraform resource aws_autoscaling_policy.
type AutoscalingPolicy struct {
	Name      string
	Args      AutoscalingPolicyArgs
	state     *autoscalingPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AutoscalingPolicy].
func (ap *AutoscalingPolicy) Type() string {
	return "aws_autoscaling_policy"
}

// LocalName returns the local name for [AutoscalingPolicy].
func (ap *AutoscalingPolicy) LocalName() string {
	return ap.Name
}

// Configuration returns the configuration (args) for [AutoscalingPolicy].
func (ap *AutoscalingPolicy) Configuration() interface{} {
	return ap.Args
}

// DependOn is used for other resources to depend on [AutoscalingPolicy].
func (ap *AutoscalingPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(ap)
}

// Dependencies returns the list of resources [AutoscalingPolicy] depends_on.
func (ap *AutoscalingPolicy) Dependencies() terra.Dependencies {
	return ap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AutoscalingPolicy].
func (ap *AutoscalingPolicy) LifecycleManagement() *terra.Lifecycle {
	return ap.Lifecycle
}

// Attributes returns the attributes for [AutoscalingPolicy].
func (ap *AutoscalingPolicy) Attributes() autoscalingPolicyAttributes {
	return autoscalingPolicyAttributes{ref: terra.ReferenceResource(ap)}
}

// ImportState imports the given attribute values into [AutoscalingPolicy]'s state.
func (ap *AutoscalingPolicy) ImportState(av io.Reader) error {
	ap.state = &autoscalingPolicyState{}
	if err := json.NewDecoder(av).Decode(ap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ap.Type(), ap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AutoscalingPolicy] has state.
func (ap *AutoscalingPolicy) State() (*autoscalingPolicyState, bool) {
	return ap.state, ap.state != nil
}

// StateMust returns the state for [AutoscalingPolicy]. Panics if the state is nil.
func (ap *AutoscalingPolicy) StateMust() *autoscalingPolicyState {
	if ap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ap.Type(), ap.LocalName()))
	}
	return ap.state
}

// AutoscalingPolicyArgs contains the configurations for aws_autoscaling_policy.
type AutoscalingPolicyArgs struct {
	// AdjustmentType: string, optional
	AdjustmentType terra.StringValue `hcl:"adjustment_type,attr"`
	// AutoscalingGroupName: string, required
	AutoscalingGroupName terra.StringValue `hcl:"autoscaling_group_name,attr" validate:"required"`
	// Cooldown: number, optional
	Cooldown terra.NumberValue `hcl:"cooldown,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// EstimatedInstanceWarmup: number, optional
	EstimatedInstanceWarmup terra.NumberValue `hcl:"estimated_instance_warmup,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MetricAggregationType: string, optional
	MetricAggregationType terra.StringValue `hcl:"metric_aggregation_type,attr"`
	// MinAdjustmentMagnitude: number, optional
	MinAdjustmentMagnitude terra.NumberValue `hcl:"min_adjustment_magnitude,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PolicyType: string, optional
	PolicyType terra.StringValue `hcl:"policy_type,attr"`
	// ScalingAdjustment: number, optional
	ScalingAdjustment terra.NumberValue `hcl:"scaling_adjustment,attr"`
	// PredictiveScalingConfiguration: optional
	PredictiveScalingConfiguration *autoscalingpolicy.PredictiveScalingConfiguration `hcl:"predictive_scaling_configuration,block"`
	// StepAdjustment: min=0
	StepAdjustment []autoscalingpolicy.StepAdjustment `hcl:"step_adjustment,block" validate:"min=0"`
	// TargetTrackingConfiguration: optional
	TargetTrackingConfiguration *autoscalingpolicy.TargetTrackingConfiguration `hcl:"target_tracking_configuration,block"`
}
type autoscalingPolicyAttributes struct {
	ref terra.Reference
}

// AdjustmentType returns a reference to field adjustment_type of aws_autoscaling_policy.
func (ap autoscalingPolicyAttributes) AdjustmentType() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("adjustment_type"))
}

// Arn returns a reference to field arn of aws_autoscaling_policy.
func (ap autoscalingPolicyAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("arn"))
}

// AutoscalingGroupName returns a reference to field autoscaling_group_name of aws_autoscaling_policy.
func (ap autoscalingPolicyAttributes) AutoscalingGroupName() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("autoscaling_group_name"))
}

// Cooldown returns a reference to field cooldown of aws_autoscaling_policy.
func (ap autoscalingPolicyAttributes) Cooldown() terra.NumberValue {
	return terra.ReferenceAsNumber(ap.ref.Append("cooldown"))
}

// Enabled returns a reference to field enabled of aws_autoscaling_policy.
func (ap autoscalingPolicyAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(ap.ref.Append("enabled"))
}

// EstimatedInstanceWarmup returns a reference to field estimated_instance_warmup of aws_autoscaling_policy.
func (ap autoscalingPolicyAttributes) EstimatedInstanceWarmup() terra.NumberValue {
	return terra.ReferenceAsNumber(ap.ref.Append("estimated_instance_warmup"))
}

// Id returns a reference to field id of aws_autoscaling_policy.
func (ap autoscalingPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("id"))
}

// MetricAggregationType returns a reference to field metric_aggregation_type of aws_autoscaling_policy.
func (ap autoscalingPolicyAttributes) MetricAggregationType() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("metric_aggregation_type"))
}

// MinAdjustmentMagnitude returns a reference to field min_adjustment_magnitude of aws_autoscaling_policy.
func (ap autoscalingPolicyAttributes) MinAdjustmentMagnitude() terra.NumberValue {
	return terra.ReferenceAsNumber(ap.ref.Append("min_adjustment_magnitude"))
}

// Name returns a reference to field name of aws_autoscaling_policy.
func (ap autoscalingPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("name"))
}

// PolicyType returns a reference to field policy_type of aws_autoscaling_policy.
func (ap autoscalingPolicyAttributes) PolicyType() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("policy_type"))
}

// ScalingAdjustment returns a reference to field scaling_adjustment of aws_autoscaling_policy.
func (ap autoscalingPolicyAttributes) ScalingAdjustment() terra.NumberValue {
	return terra.ReferenceAsNumber(ap.ref.Append("scaling_adjustment"))
}

func (ap autoscalingPolicyAttributes) PredictiveScalingConfiguration() terra.ListValue[autoscalingpolicy.PredictiveScalingConfigurationAttributes] {
	return terra.ReferenceAsList[autoscalingpolicy.PredictiveScalingConfigurationAttributes](ap.ref.Append("predictive_scaling_configuration"))
}

func (ap autoscalingPolicyAttributes) StepAdjustment() terra.SetValue[autoscalingpolicy.StepAdjustmentAttributes] {
	return terra.ReferenceAsSet[autoscalingpolicy.StepAdjustmentAttributes](ap.ref.Append("step_adjustment"))
}

func (ap autoscalingPolicyAttributes) TargetTrackingConfiguration() terra.ListValue[autoscalingpolicy.TargetTrackingConfigurationAttributes] {
	return terra.ReferenceAsList[autoscalingpolicy.TargetTrackingConfigurationAttributes](ap.ref.Append("target_tracking_configuration"))
}

type autoscalingPolicyState struct {
	AdjustmentType                 string                                                  `json:"adjustment_type"`
	Arn                            string                                                  `json:"arn"`
	AutoscalingGroupName           string                                                  `json:"autoscaling_group_name"`
	Cooldown                       float64                                                 `json:"cooldown"`
	Enabled                        bool                                                    `json:"enabled"`
	EstimatedInstanceWarmup        float64                                                 `json:"estimated_instance_warmup"`
	Id                             string                                                  `json:"id"`
	MetricAggregationType          string                                                  `json:"metric_aggregation_type"`
	MinAdjustmentMagnitude         float64                                                 `json:"min_adjustment_magnitude"`
	Name                           string                                                  `json:"name"`
	PolicyType                     string                                                  `json:"policy_type"`
	ScalingAdjustment              float64                                                 `json:"scaling_adjustment"`
	PredictiveScalingConfiguration []autoscalingpolicy.PredictiveScalingConfigurationState `json:"predictive_scaling_configuration"`
	StepAdjustment                 []autoscalingpolicy.StepAdjustmentState                 `json:"step_adjustment"`
	TargetTrackingConfiguration    []autoscalingpolicy.TargetTrackingConfigurationState    `json:"target_tracking_configuration"`
}
