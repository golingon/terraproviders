// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAutoscalingNotification creates a new instance of [AutoscalingNotification].
func NewAutoscalingNotification(name string, args AutoscalingNotificationArgs) *AutoscalingNotification {
	return &AutoscalingNotification{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AutoscalingNotification)(nil)

// AutoscalingNotification represents the Terraform resource aws_autoscaling_notification.
type AutoscalingNotification struct {
	Name      string
	Args      AutoscalingNotificationArgs
	state     *autoscalingNotificationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AutoscalingNotification].
func (an *AutoscalingNotification) Type() string {
	return "aws_autoscaling_notification"
}

// LocalName returns the local name for [AutoscalingNotification].
func (an *AutoscalingNotification) LocalName() string {
	return an.Name
}

// Configuration returns the configuration (args) for [AutoscalingNotification].
func (an *AutoscalingNotification) Configuration() interface{} {
	return an.Args
}

// DependOn is used for other resources to depend on [AutoscalingNotification].
func (an *AutoscalingNotification) DependOn() terra.Reference {
	return terra.ReferenceResource(an)
}

// Dependencies returns the list of resources [AutoscalingNotification] depends_on.
func (an *AutoscalingNotification) Dependencies() terra.Dependencies {
	return an.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AutoscalingNotification].
func (an *AutoscalingNotification) LifecycleManagement() *terra.Lifecycle {
	return an.Lifecycle
}

// Attributes returns the attributes for [AutoscalingNotification].
func (an *AutoscalingNotification) Attributes() autoscalingNotificationAttributes {
	return autoscalingNotificationAttributes{ref: terra.ReferenceResource(an)}
}

// ImportState imports the given attribute values into [AutoscalingNotification]'s state.
func (an *AutoscalingNotification) ImportState(av io.Reader) error {
	an.state = &autoscalingNotificationState{}
	if err := json.NewDecoder(av).Decode(an.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", an.Type(), an.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AutoscalingNotification] has state.
func (an *AutoscalingNotification) State() (*autoscalingNotificationState, bool) {
	return an.state, an.state != nil
}

// StateMust returns the state for [AutoscalingNotification]. Panics if the state is nil.
func (an *AutoscalingNotification) StateMust() *autoscalingNotificationState {
	if an.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", an.Type(), an.LocalName()))
	}
	return an.state
}

// AutoscalingNotificationArgs contains the configurations for aws_autoscaling_notification.
type AutoscalingNotificationArgs struct {
	// GroupNames: set of string, required
	GroupNames terra.SetValue[terra.StringValue] `hcl:"group_names,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Notifications: set of string, required
	Notifications terra.SetValue[terra.StringValue] `hcl:"notifications,attr" validate:"required"`
	// TopicArn: string, required
	TopicArn terra.StringValue `hcl:"topic_arn,attr" validate:"required"`
}
type autoscalingNotificationAttributes struct {
	ref terra.Reference
}

// GroupNames returns a reference to field group_names of aws_autoscaling_notification.
func (an autoscalingNotificationAttributes) GroupNames() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](an.ref.Append("group_names"))
}

// Id returns a reference to field id of aws_autoscaling_notification.
func (an autoscalingNotificationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(an.ref.Append("id"))
}

// Notifications returns a reference to field notifications of aws_autoscaling_notification.
func (an autoscalingNotificationAttributes) Notifications() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](an.ref.Append("notifications"))
}

// TopicArn returns a reference to field topic_arn of aws_autoscaling_notification.
func (an autoscalingNotificationAttributes) TopicArn() terra.StringValue {
	return terra.ReferenceAsString(an.ref.Append("topic_arn"))
}

type autoscalingNotificationState struct {
	GroupNames    []string `json:"group_names"`
	Id            string   `json:"id"`
	Notifications []string `json:"notifications"`
	TopicArn      string   `json:"topic_arn"`
}