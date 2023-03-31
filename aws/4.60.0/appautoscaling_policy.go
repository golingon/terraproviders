// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	appautoscalingpolicy "github.com/golingon/terraproviders/aws/4.60.0/appautoscalingpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppautoscalingPolicy creates a new instance of [AppautoscalingPolicy].
func NewAppautoscalingPolicy(name string, args AppautoscalingPolicyArgs) *AppautoscalingPolicy {
	return &AppautoscalingPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppautoscalingPolicy)(nil)

// AppautoscalingPolicy represents the Terraform resource aws_appautoscaling_policy.
type AppautoscalingPolicy struct {
	Name      string
	Args      AppautoscalingPolicyArgs
	state     *appautoscalingPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppautoscalingPolicy].
func (ap *AppautoscalingPolicy) Type() string {
	return "aws_appautoscaling_policy"
}

// LocalName returns the local name for [AppautoscalingPolicy].
func (ap *AppautoscalingPolicy) LocalName() string {
	return ap.Name
}

// Configuration returns the configuration (args) for [AppautoscalingPolicy].
func (ap *AppautoscalingPolicy) Configuration() interface{} {
	return ap.Args
}

// DependOn is used for other resources to depend on [AppautoscalingPolicy].
func (ap *AppautoscalingPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(ap)
}

// Dependencies returns the list of resources [AppautoscalingPolicy] depends_on.
func (ap *AppautoscalingPolicy) Dependencies() terra.Dependencies {
	return ap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppautoscalingPolicy].
func (ap *AppautoscalingPolicy) LifecycleManagement() *terra.Lifecycle {
	return ap.Lifecycle
}

// Attributes returns the attributes for [AppautoscalingPolicy].
func (ap *AppautoscalingPolicy) Attributes() appautoscalingPolicyAttributes {
	return appautoscalingPolicyAttributes{ref: terra.ReferenceResource(ap)}
}

// ImportState imports the given attribute values into [AppautoscalingPolicy]'s state.
func (ap *AppautoscalingPolicy) ImportState(av io.Reader) error {
	ap.state = &appautoscalingPolicyState{}
	if err := json.NewDecoder(av).Decode(ap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ap.Type(), ap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppautoscalingPolicy] has state.
func (ap *AppautoscalingPolicy) State() (*appautoscalingPolicyState, bool) {
	return ap.state, ap.state != nil
}

// StateMust returns the state for [AppautoscalingPolicy]. Panics if the state is nil.
func (ap *AppautoscalingPolicy) StateMust() *appautoscalingPolicyState {
	if ap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ap.Type(), ap.LocalName()))
	}
	return ap.state
}

// AppautoscalingPolicyArgs contains the configurations for aws_appautoscaling_policy.
type AppautoscalingPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PolicyType: string, optional
	PolicyType terra.StringValue `hcl:"policy_type,attr"`
	// ResourceId: string, required
	ResourceId terra.StringValue `hcl:"resource_id,attr" validate:"required"`
	// ScalableDimension: string, required
	ScalableDimension terra.StringValue `hcl:"scalable_dimension,attr" validate:"required"`
	// ServiceNamespace: string, required
	ServiceNamespace terra.StringValue `hcl:"service_namespace,attr" validate:"required"`
	// StepScalingPolicyConfiguration: optional
	StepScalingPolicyConfiguration *appautoscalingpolicy.StepScalingPolicyConfiguration `hcl:"step_scaling_policy_configuration,block"`
	// TargetTrackingScalingPolicyConfiguration: optional
	TargetTrackingScalingPolicyConfiguration *appautoscalingpolicy.TargetTrackingScalingPolicyConfiguration `hcl:"target_tracking_scaling_policy_configuration,block"`
}
type appautoscalingPolicyAttributes struct {
	ref terra.Reference
}

// AlarmArns returns a reference to field alarm_arns of aws_appautoscaling_policy.
func (ap appautoscalingPolicyAttributes) AlarmArns() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ap.ref.Append("alarm_arns"))
}

// Arn returns a reference to field arn of aws_appautoscaling_policy.
func (ap appautoscalingPolicyAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("arn"))
}

// Id returns a reference to field id of aws_appautoscaling_policy.
func (ap appautoscalingPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("id"))
}

// Name returns a reference to field name of aws_appautoscaling_policy.
func (ap appautoscalingPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("name"))
}

// PolicyType returns a reference to field policy_type of aws_appautoscaling_policy.
func (ap appautoscalingPolicyAttributes) PolicyType() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("policy_type"))
}

// ResourceId returns a reference to field resource_id of aws_appautoscaling_policy.
func (ap appautoscalingPolicyAttributes) ResourceId() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("resource_id"))
}

// ScalableDimension returns a reference to field scalable_dimension of aws_appautoscaling_policy.
func (ap appautoscalingPolicyAttributes) ScalableDimension() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("scalable_dimension"))
}

// ServiceNamespace returns a reference to field service_namespace of aws_appautoscaling_policy.
func (ap appautoscalingPolicyAttributes) ServiceNamespace() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("service_namespace"))
}

func (ap appautoscalingPolicyAttributes) StepScalingPolicyConfiguration() terra.ListValue[appautoscalingpolicy.StepScalingPolicyConfigurationAttributes] {
	return terra.ReferenceAsList[appautoscalingpolicy.StepScalingPolicyConfigurationAttributes](ap.ref.Append("step_scaling_policy_configuration"))
}

func (ap appautoscalingPolicyAttributes) TargetTrackingScalingPolicyConfiguration() terra.ListValue[appautoscalingpolicy.TargetTrackingScalingPolicyConfigurationAttributes] {
	return terra.ReferenceAsList[appautoscalingpolicy.TargetTrackingScalingPolicyConfigurationAttributes](ap.ref.Append("target_tracking_scaling_policy_configuration"))
}

type appautoscalingPolicyState struct {
	AlarmArns                                []string                                                             `json:"alarm_arns"`
	Arn                                      string                                                               `json:"arn"`
	Id                                       string                                                               `json:"id"`
	Name                                     string                                                               `json:"name"`
	PolicyType                               string                                                               `json:"policy_type"`
	ResourceId                               string                                                               `json:"resource_id"`
	ScalableDimension                        string                                                               `json:"scalable_dimension"`
	ServiceNamespace                         string                                                               `json:"service_namespace"`
	StepScalingPolicyConfiguration           []appautoscalingpolicy.StepScalingPolicyConfigurationState           `json:"step_scaling_policy_configuration"`
	TargetTrackingScalingPolicyConfiguration []appautoscalingpolicy.TargetTrackingScalingPolicyConfigurationState `json:"target_tracking_scaling_policy_configuration"`
}
