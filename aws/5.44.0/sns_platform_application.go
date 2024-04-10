// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewSnsPlatformApplication creates a new instance of [SnsPlatformApplication].
func NewSnsPlatformApplication(name string, args SnsPlatformApplicationArgs) *SnsPlatformApplication {
	return &SnsPlatformApplication{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SnsPlatformApplication)(nil)

// SnsPlatformApplication represents the Terraform resource aws_sns_platform_application.
type SnsPlatformApplication struct {
	Name      string
	Args      SnsPlatformApplicationArgs
	state     *snsPlatformApplicationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SnsPlatformApplication].
func (spa *SnsPlatformApplication) Type() string {
	return "aws_sns_platform_application"
}

// LocalName returns the local name for [SnsPlatformApplication].
func (spa *SnsPlatformApplication) LocalName() string {
	return spa.Name
}

// Configuration returns the configuration (args) for [SnsPlatformApplication].
func (spa *SnsPlatformApplication) Configuration() interface{} {
	return spa.Args
}

// DependOn is used for other resources to depend on [SnsPlatformApplication].
func (spa *SnsPlatformApplication) DependOn() terra.Reference {
	return terra.ReferenceResource(spa)
}

// Dependencies returns the list of resources [SnsPlatformApplication] depends_on.
func (spa *SnsPlatformApplication) Dependencies() terra.Dependencies {
	return spa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SnsPlatformApplication].
func (spa *SnsPlatformApplication) LifecycleManagement() *terra.Lifecycle {
	return spa.Lifecycle
}

// Attributes returns the attributes for [SnsPlatformApplication].
func (spa *SnsPlatformApplication) Attributes() snsPlatformApplicationAttributes {
	return snsPlatformApplicationAttributes{ref: terra.ReferenceResource(spa)}
}

// ImportState imports the given attribute values into [SnsPlatformApplication]'s state.
func (spa *SnsPlatformApplication) ImportState(av io.Reader) error {
	spa.state = &snsPlatformApplicationState{}
	if err := json.NewDecoder(av).Decode(spa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", spa.Type(), spa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SnsPlatformApplication] has state.
func (spa *SnsPlatformApplication) State() (*snsPlatformApplicationState, bool) {
	return spa.state, spa.state != nil
}

// StateMust returns the state for [SnsPlatformApplication]. Panics if the state is nil.
func (spa *SnsPlatformApplication) StateMust() *snsPlatformApplicationState {
	if spa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", spa.Type(), spa.LocalName()))
	}
	return spa.state
}

// SnsPlatformApplicationArgs contains the configurations for aws_sns_platform_application.
type SnsPlatformApplicationArgs struct {
	// ApplePlatformBundleId: string, optional
	ApplePlatformBundleId terra.StringValue `hcl:"apple_platform_bundle_id,attr"`
	// ApplePlatformTeamId: string, optional
	ApplePlatformTeamId terra.StringValue `hcl:"apple_platform_team_id,attr"`
	// EventDeliveryFailureTopicArn: string, optional
	EventDeliveryFailureTopicArn terra.StringValue `hcl:"event_delivery_failure_topic_arn,attr"`
	// EventEndpointCreatedTopicArn: string, optional
	EventEndpointCreatedTopicArn terra.StringValue `hcl:"event_endpoint_created_topic_arn,attr"`
	// EventEndpointDeletedTopicArn: string, optional
	EventEndpointDeletedTopicArn terra.StringValue `hcl:"event_endpoint_deleted_topic_arn,attr"`
	// EventEndpointUpdatedTopicArn: string, optional
	EventEndpointUpdatedTopicArn terra.StringValue `hcl:"event_endpoint_updated_topic_arn,attr"`
	// FailureFeedbackRoleArn: string, optional
	FailureFeedbackRoleArn terra.StringValue `hcl:"failure_feedback_role_arn,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Platform: string, required
	Platform terra.StringValue `hcl:"platform,attr" validate:"required"`
	// PlatformCredential: string, required
	PlatformCredential terra.StringValue `hcl:"platform_credential,attr" validate:"required"`
	// PlatformPrincipal: string, optional
	PlatformPrincipal terra.StringValue `hcl:"platform_principal,attr"`
	// SuccessFeedbackRoleArn: string, optional
	SuccessFeedbackRoleArn terra.StringValue `hcl:"success_feedback_role_arn,attr"`
	// SuccessFeedbackSampleRate: string, optional
	SuccessFeedbackSampleRate terra.StringValue `hcl:"success_feedback_sample_rate,attr"`
}
type snsPlatformApplicationAttributes struct {
	ref terra.Reference
}

// ApplePlatformBundleId returns a reference to field apple_platform_bundle_id of aws_sns_platform_application.
func (spa snsPlatformApplicationAttributes) ApplePlatformBundleId() terra.StringValue {
	return terra.ReferenceAsString(spa.ref.Append("apple_platform_bundle_id"))
}

// ApplePlatformTeamId returns a reference to field apple_platform_team_id of aws_sns_platform_application.
func (spa snsPlatformApplicationAttributes) ApplePlatformTeamId() terra.StringValue {
	return terra.ReferenceAsString(spa.ref.Append("apple_platform_team_id"))
}

// Arn returns a reference to field arn of aws_sns_platform_application.
func (spa snsPlatformApplicationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(spa.ref.Append("arn"))
}

// EventDeliveryFailureTopicArn returns a reference to field event_delivery_failure_topic_arn of aws_sns_platform_application.
func (spa snsPlatformApplicationAttributes) EventDeliveryFailureTopicArn() terra.StringValue {
	return terra.ReferenceAsString(spa.ref.Append("event_delivery_failure_topic_arn"))
}

// EventEndpointCreatedTopicArn returns a reference to field event_endpoint_created_topic_arn of aws_sns_platform_application.
func (spa snsPlatformApplicationAttributes) EventEndpointCreatedTopicArn() terra.StringValue {
	return terra.ReferenceAsString(spa.ref.Append("event_endpoint_created_topic_arn"))
}

// EventEndpointDeletedTopicArn returns a reference to field event_endpoint_deleted_topic_arn of aws_sns_platform_application.
func (spa snsPlatformApplicationAttributes) EventEndpointDeletedTopicArn() terra.StringValue {
	return terra.ReferenceAsString(spa.ref.Append("event_endpoint_deleted_topic_arn"))
}

// EventEndpointUpdatedTopicArn returns a reference to field event_endpoint_updated_topic_arn of aws_sns_platform_application.
func (spa snsPlatformApplicationAttributes) EventEndpointUpdatedTopicArn() terra.StringValue {
	return terra.ReferenceAsString(spa.ref.Append("event_endpoint_updated_topic_arn"))
}

// FailureFeedbackRoleArn returns a reference to field failure_feedback_role_arn of aws_sns_platform_application.
func (spa snsPlatformApplicationAttributes) FailureFeedbackRoleArn() terra.StringValue {
	return terra.ReferenceAsString(spa.ref.Append("failure_feedback_role_arn"))
}

// Id returns a reference to field id of aws_sns_platform_application.
func (spa snsPlatformApplicationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(spa.ref.Append("id"))
}

// Name returns a reference to field name of aws_sns_platform_application.
func (spa snsPlatformApplicationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(spa.ref.Append("name"))
}

// Platform returns a reference to field platform of aws_sns_platform_application.
func (spa snsPlatformApplicationAttributes) Platform() terra.StringValue {
	return terra.ReferenceAsString(spa.ref.Append("platform"))
}

// PlatformCredential returns a reference to field platform_credential of aws_sns_platform_application.
func (spa snsPlatformApplicationAttributes) PlatformCredential() terra.StringValue {
	return terra.ReferenceAsString(spa.ref.Append("platform_credential"))
}

// PlatformPrincipal returns a reference to field platform_principal of aws_sns_platform_application.
func (spa snsPlatformApplicationAttributes) PlatformPrincipal() terra.StringValue {
	return terra.ReferenceAsString(spa.ref.Append("platform_principal"))
}

// SuccessFeedbackRoleArn returns a reference to field success_feedback_role_arn of aws_sns_platform_application.
func (spa snsPlatformApplicationAttributes) SuccessFeedbackRoleArn() terra.StringValue {
	return terra.ReferenceAsString(spa.ref.Append("success_feedback_role_arn"))
}

// SuccessFeedbackSampleRate returns a reference to field success_feedback_sample_rate of aws_sns_platform_application.
func (spa snsPlatformApplicationAttributes) SuccessFeedbackSampleRate() terra.StringValue {
	return terra.ReferenceAsString(spa.ref.Append("success_feedback_sample_rate"))
}

type snsPlatformApplicationState struct {
	ApplePlatformBundleId        string `json:"apple_platform_bundle_id"`
	ApplePlatformTeamId          string `json:"apple_platform_team_id"`
	Arn                          string `json:"arn"`
	EventDeliveryFailureTopicArn string `json:"event_delivery_failure_topic_arn"`
	EventEndpointCreatedTopicArn string `json:"event_endpoint_created_topic_arn"`
	EventEndpointDeletedTopicArn string `json:"event_endpoint_deleted_topic_arn"`
	EventEndpointUpdatedTopicArn string `json:"event_endpoint_updated_topic_arn"`
	FailureFeedbackRoleArn       string `json:"failure_feedback_role_arn"`
	Id                           string `json:"id"`
	Name                         string `json:"name"`
	Platform                     string `json:"platform"`
	PlatformCredential           string `json:"platform_credential"`
	PlatformPrincipal            string `json:"platform_principal"`
	SuccessFeedbackRoleArn       string `json:"success_feedback_role_arn"`
	SuccessFeedbackSampleRate    string `json:"success_feedback_sample_rate"`
}
