// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAutoscalingLifecycleHook creates a new instance of [AutoscalingLifecycleHook].
func NewAutoscalingLifecycleHook(name string, args AutoscalingLifecycleHookArgs) *AutoscalingLifecycleHook {
	return &AutoscalingLifecycleHook{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AutoscalingLifecycleHook)(nil)

// AutoscalingLifecycleHook represents the Terraform resource aws_autoscaling_lifecycle_hook.
type AutoscalingLifecycleHook struct {
	Name      string
	Args      AutoscalingLifecycleHookArgs
	state     *autoscalingLifecycleHookState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AutoscalingLifecycleHook].
func (alh *AutoscalingLifecycleHook) Type() string {
	return "aws_autoscaling_lifecycle_hook"
}

// LocalName returns the local name for [AutoscalingLifecycleHook].
func (alh *AutoscalingLifecycleHook) LocalName() string {
	return alh.Name
}

// Configuration returns the configuration (args) for [AutoscalingLifecycleHook].
func (alh *AutoscalingLifecycleHook) Configuration() interface{} {
	return alh.Args
}

// DependOn is used for other resources to depend on [AutoscalingLifecycleHook].
func (alh *AutoscalingLifecycleHook) DependOn() terra.Reference {
	return terra.ReferenceResource(alh)
}

// Dependencies returns the list of resources [AutoscalingLifecycleHook] depends_on.
func (alh *AutoscalingLifecycleHook) Dependencies() terra.Dependencies {
	return alh.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AutoscalingLifecycleHook].
func (alh *AutoscalingLifecycleHook) LifecycleManagement() *terra.Lifecycle {
	return alh.Lifecycle
}

// Attributes returns the attributes for [AutoscalingLifecycleHook].
func (alh *AutoscalingLifecycleHook) Attributes() autoscalingLifecycleHookAttributes {
	return autoscalingLifecycleHookAttributes{ref: terra.ReferenceResource(alh)}
}

// ImportState imports the given attribute values into [AutoscalingLifecycleHook]'s state.
func (alh *AutoscalingLifecycleHook) ImportState(av io.Reader) error {
	alh.state = &autoscalingLifecycleHookState{}
	if err := json.NewDecoder(av).Decode(alh.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", alh.Type(), alh.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AutoscalingLifecycleHook] has state.
func (alh *AutoscalingLifecycleHook) State() (*autoscalingLifecycleHookState, bool) {
	return alh.state, alh.state != nil
}

// StateMust returns the state for [AutoscalingLifecycleHook]. Panics if the state is nil.
func (alh *AutoscalingLifecycleHook) StateMust() *autoscalingLifecycleHookState {
	if alh.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", alh.Type(), alh.LocalName()))
	}
	return alh.state
}

// AutoscalingLifecycleHookArgs contains the configurations for aws_autoscaling_lifecycle_hook.
type AutoscalingLifecycleHookArgs struct {
	// AutoscalingGroupName: string, required
	AutoscalingGroupName terra.StringValue `hcl:"autoscaling_group_name,attr" validate:"required"`
	// DefaultResult: string, optional
	DefaultResult terra.StringValue `hcl:"default_result,attr"`
	// HeartbeatTimeout: number, optional
	HeartbeatTimeout terra.NumberValue `hcl:"heartbeat_timeout,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LifecycleTransition: string, required
	LifecycleTransition terra.StringValue `hcl:"lifecycle_transition,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NotificationMetadata: string, optional
	NotificationMetadata terra.StringValue `hcl:"notification_metadata,attr"`
	// NotificationTargetArn: string, optional
	NotificationTargetArn terra.StringValue `hcl:"notification_target_arn,attr"`
	// RoleArn: string, optional
	RoleArn terra.StringValue `hcl:"role_arn,attr"`
}
type autoscalingLifecycleHookAttributes struct {
	ref terra.Reference
}

// AutoscalingGroupName returns a reference to field autoscaling_group_name of aws_autoscaling_lifecycle_hook.
func (alh autoscalingLifecycleHookAttributes) AutoscalingGroupName() terra.StringValue {
	return terra.ReferenceAsString(alh.ref.Append("autoscaling_group_name"))
}

// DefaultResult returns a reference to field default_result of aws_autoscaling_lifecycle_hook.
func (alh autoscalingLifecycleHookAttributes) DefaultResult() terra.StringValue {
	return terra.ReferenceAsString(alh.ref.Append("default_result"))
}

// HeartbeatTimeout returns a reference to field heartbeat_timeout of aws_autoscaling_lifecycle_hook.
func (alh autoscalingLifecycleHookAttributes) HeartbeatTimeout() terra.NumberValue {
	return terra.ReferenceAsNumber(alh.ref.Append("heartbeat_timeout"))
}

// Id returns a reference to field id of aws_autoscaling_lifecycle_hook.
func (alh autoscalingLifecycleHookAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(alh.ref.Append("id"))
}

// LifecycleTransition returns a reference to field lifecycle_transition of aws_autoscaling_lifecycle_hook.
func (alh autoscalingLifecycleHookAttributes) LifecycleTransition() terra.StringValue {
	return terra.ReferenceAsString(alh.ref.Append("lifecycle_transition"))
}

// Name returns a reference to field name of aws_autoscaling_lifecycle_hook.
func (alh autoscalingLifecycleHookAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(alh.ref.Append("name"))
}

// NotificationMetadata returns a reference to field notification_metadata of aws_autoscaling_lifecycle_hook.
func (alh autoscalingLifecycleHookAttributes) NotificationMetadata() terra.StringValue {
	return terra.ReferenceAsString(alh.ref.Append("notification_metadata"))
}

// NotificationTargetArn returns a reference to field notification_target_arn of aws_autoscaling_lifecycle_hook.
func (alh autoscalingLifecycleHookAttributes) NotificationTargetArn() terra.StringValue {
	return terra.ReferenceAsString(alh.ref.Append("notification_target_arn"))
}

// RoleArn returns a reference to field role_arn of aws_autoscaling_lifecycle_hook.
func (alh autoscalingLifecycleHookAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(alh.ref.Append("role_arn"))
}

type autoscalingLifecycleHookState struct {
	AutoscalingGroupName  string  `json:"autoscaling_group_name"`
	DefaultResult         string  `json:"default_result"`
	HeartbeatTimeout      float64 `json:"heartbeat_timeout"`
	Id                    string  `json:"id"`
	LifecycleTransition   string  `json:"lifecycle_transition"`
	Name                  string  `json:"name"`
	NotificationMetadata  string  `json:"notification_metadata"`
	NotificationTargetArn string  `json:"notification_target_arn"`
	RoleArn               string  `json:"role_arn"`
}
