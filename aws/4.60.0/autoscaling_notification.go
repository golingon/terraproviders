// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewAutoscalingNotification(name string, args AutoscalingNotificationArgs) *AutoscalingNotification {
	return &AutoscalingNotification{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AutoscalingNotification)(nil)

type AutoscalingNotification struct {
	Name  string
	Args  AutoscalingNotificationArgs
	state *autoscalingNotificationState
}

func (an *AutoscalingNotification) Type() string {
	return "aws_autoscaling_notification"
}

func (an *AutoscalingNotification) LocalName() string {
	return an.Name
}

func (an *AutoscalingNotification) Configuration() interface{} {
	return an.Args
}

func (an *AutoscalingNotification) Attributes() autoscalingNotificationAttributes {
	return autoscalingNotificationAttributes{ref: terra.ReferenceResource(an)}
}

func (an *AutoscalingNotification) ImportState(av io.Reader) error {
	an.state = &autoscalingNotificationState{}
	if err := json.NewDecoder(av).Decode(an.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", an.Type(), an.LocalName(), err)
	}
	return nil
}

func (an *AutoscalingNotification) State() (*autoscalingNotificationState, bool) {
	return an.state, an.state != nil
}

func (an *AutoscalingNotification) StateMust() *autoscalingNotificationState {
	if an.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", an.Type(), an.LocalName()))
	}
	return an.state
}

func (an *AutoscalingNotification) DependOn() terra.Reference {
	return terra.ReferenceResource(an)
}

type AutoscalingNotificationArgs struct {
	// GroupNames: set of string, required
	GroupNames terra.SetValue[terra.StringValue] `hcl:"group_names,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Notifications: set of string, required
	Notifications terra.SetValue[terra.StringValue] `hcl:"notifications,attr" validate:"required"`
	// TopicArn: string, required
	TopicArn terra.StringValue `hcl:"topic_arn,attr" validate:"required"`
	// DependsOn contains resources that AutoscalingNotification depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type autoscalingNotificationAttributes struct {
	ref terra.Reference
}

func (an autoscalingNotificationAttributes) GroupNames() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](an.ref.Append("group_names"))
}

func (an autoscalingNotificationAttributes) Id() terra.StringValue {
	return terra.ReferenceString(an.ref.Append("id"))
}

func (an autoscalingNotificationAttributes) Notifications() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](an.ref.Append("notifications"))
}

func (an autoscalingNotificationAttributes) TopicArn() terra.StringValue {
	return terra.ReferenceString(an.ref.Append("topic_arn"))
}

type autoscalingNotificationState struct {
	GroupNames    []string `json:"group_names"`
	Id            string   `json:"id"`
	Notifications []string `json:"notifications"`
	TopicArn      string   `json:"topic_arn"`
}
