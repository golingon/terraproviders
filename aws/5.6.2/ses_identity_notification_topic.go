// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSesIdentityNotificationTopic creates a new instance of [SesIdentityNotificationTopic].
func NewSesIdentityNotificationTopic(name string, args SesIdentityNotificationTopicArgs) *SesIdentityNotificationTopic {
	return &SesIdentityNotificationTopic{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SesIdentityNotificationTopic)(nil)

// SesIdentityNotificationTopic represents the Terraform resource aws_ses_identity_notification_topic.
type SesIdentityNotificationTopic struct {
	Name      string
	Args      SesIdentityNotificationTopicArgs
	state     *sesIdentityNotificationTopicState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SesIdentityNotificationTopic].
func (sint *SesIdentityNotificationTopic) Type() string {
	return "aws_ses_identity_notification_topic"
}

// LocalName returns the local name for [SesIdentityNotificationTopic].
func (sint *SesIdentityNotificationTopic) LocalName() string {
	return sint.Name
}

// Configuration returns the configuration (args) for [SesIdentityNotificationTopic].
func (sint *SesIdentityNotificationTopic) Configuration() interface{} {
	return sint.Args
}

// DependOn is used for other resources to depend on [SesIdentityNotificationTopic].
func (sint *SesIdentityNotificationTopic) DependOn() terra.Reference {
	return terra.ReferenceResource(sint)
}

// Dependencies returns the list of resources [SesIdentityNotificationTopic] depends_on.
func (sint *SesIdentityNotificationTopic) Dependencies() terra.Dependencies {
	return sint.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SesIdentityNotificationTopic].
func (sint *SesIdentityNotificationTopic) LifecycleManagement() *terra.Lifecycle {
	return sint.Lifecycle
}

// Attributes returns the attributes for [SesIdentityNotificationTopic].
func (sint *SesIdentityNotificationTopic) Attributes() sesIdentityNotificationTopicAttributes {
	return sesIdentityNotificationTopicAttributes{ref: terra.ReferenceResource(sint)}
}

// ImportState imports the given attribute values into [SesIdentityNotificationTopic]'s state.
func (sint *SesIdentityNotificationTopic) ImportState(av io.Reader) error {
	sint.state = &sesIdentityNotificationTopicState{}
	if err := json.NewDecoder(av).Decode(sint.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sint.Type(), sint.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SesIdentityNotificationTopic] has state.
func (sint *SesIdentityNotificationTopic) State() (*sesIdentityNotificationTopicState, bool) {
	return sint.state, sint.state != nil
}

// StateMust returns the state for [SesIdentityNotificationTopic]. Panics if the state is nil.
func (sint *SesIdentityNotificationTopic) StateMust() *sesIdentityNotificationTopicState {
	if sint.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sint.Type(), sint.LocalName()))
	}
	return sint.state
}

// SesIdentityNotificationTopicArgs contains the configurations for aws_ses_identity_notification_topic.
type SesIdentityNotificationTopicArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Identity: string, required
	Identity terra.StringValue `hcl:"identity,attr" validate:"required"`
	// IncludeOriginalHeaders: bool, optional
	IncludeOriginalHeaders terra.BoolValue `hcl:"include_original_headers,attr"`
	// NotificationType: string, required
	NotificationType terra.StringValue `hcl:"notification_type,attr" validate:"required"`
	// TopicArn: string, optional
	TopicArn terra.StringValue `hcl:"topic_arn,attr"`
}
type sesIdentityNotificationTopicAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ses_identity_notification_topic.
func (sint sesIdentityNotificationTopicAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sint.ref.Append("id"))
}

// Identity returns a reference to field identity of aws_ses_identity_notification_topic.
func (sint sesIdentityNotificationTopicAttributes) Identity() terra.StringValue {
	return terra.ReferenceAsString(sint.ref.Append("identity"))
}

// IncludeOriginalHeaders returns a reference to field include_original_headers of aws_ses_identity_notification_topic.
func (sint sesIdentityNotificationTopicAttributes) IncludeOriginalHeaders() terra.BoolValue {
	return terra.ReferenceAsBool(sint.ref.Append("include_original_headers"))
}

// NotificationType returns a reference to field notification_type of aws_ses_identity_notification_topic.
func (sint sesIdentityNotificationTopicAttributes) NotificationType() terra.StringValue {
	return terra.ReferenceAsString(sint.ref.Append("notification_type"))
}

// TopicArn returns a reference to field topic_arn of aws_ses_identity_notification_topic.
func (sint sesIdentityNotificationTopicAttributes) TopicArn() terra.StringValue {
	return terra.ReferenceAsString(sint.ref.Append("topic_arn"))
}

type sesIdentityNotificationTopicState struct {
	Id                     string `json:"id"`
	Identity               string `json:"identity"`
	IncludeOriginalHeaders bool   `json:"include_original_headers"`
	NotificationType       string `json:"notification_type"`
	TopicArn               string `json:"topic_arn"`
}
