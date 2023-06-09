// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	cloudtrail "github.com/golingon/terraproviders/aws/4.63.0/cloudtrail"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudtrail creates a new instance of [Cloudtrail].
func NewCloudtrail(name string, args CloudtrailArgs) *Cloudtrail {
	return &Cloudtrail{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Cloudtrail)(nil)

// Cloudtrail represents the Terraform resource aws_cloudtrail.
type Cloudtrail struct {
	Name      string
	Args      CloudtrailArgs
	state     *cloudtrailState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Cloudtrail].
func (c *Cloudtrail) Type() string {
	return "aws_cloudtrail"
}

// LocalName returns the local name for [Cloudtrail].
func (c *Cloudtrail) LocalName() string {
	return c.Name
}

// Configuration returns the configuration (args) for [Cloudtrail].
func (c *Cloudtrail) Configuration() interface{} {
	return c.Args
}

// DependOn is used for other resources to depend on [Cloudtrail].
func (c *Cloudtrail) DependOn() terra.Reference {
	return terra.ReferenceResource(c)
}

// Dependencies returns the list of resources [Cloudtrail] depends_on.
func (c *Cloudtrail) Dependencies() terra.Dependencies {
	return c.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Cloudtrail].
func (c *Cloudtrail) LifecycleManagement() *terra.Lifecycle {
	return c.Lifecycle
}

// Attributes returns the attributes for [Cloudtrail].
func (c *Cloudtrail) Attributes() cloudtrailAttributes {
	return cloudtrailAttributes{ref: terra.ReferenceResource(c)}
}

// ImportState imports the given attribute values into [Cloudtrail]'s state.
func (c *Cloudtrail) ImportState(av io.Reader) error {
	c.state = &cloudtrailState{}
	if err := json.NewDecoder(av).Decode(c.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", c.Type(), c.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Cloudtrail] has state.
func (c *Cloudtrail) State() (*cloudtrailState, bool) {
	return c.state, c.state != nil
}

// StateMust returns the state for [Cloudtrail]. Panics if the state is nil.
func (c *Cloudtrail) StateMust() *cloudtrailState {
	if c.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", c.Type(), c.LocalName()))
	}
	return c.state
}

// CloudtrailArgs contains the configurations for aws_cloudtrail.
type CloudtrailArgs struct {
	// CloudWatchLogsGroupArn: string, optional
	CloudWatchLogsGroupArn terra.StringValue `hcl:"cloud_watch_logs_group_arn,attr"`
	// CloudWatchLogsRoleArn: string, optional
	CloudWatchLogsRoleArn terra.StringValue `hcl:"cloud_watch_logs_role_arn,attr"`
	// EnableLogFileValidation: bool, optional
	EnableLogFileValidation terra.BoolValue `hcl:"enable_log_file_validation,attr"`
	// EnableLogging: bool, optional
	EnableLogging terra.BoolValue `hcl:"enable_logging,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IncludeGlobalServiceEvents: bool, optional
	IncludeGlobalServiceEvents terra.BoolValue `hcl:"include_global_service_events,attr"`
	// IsMultiRegionTrail: bool, optional
	IsMultiRegionTrail terra.BoolValue `hcl:"is_multi_region_trail,attr"`
	// IsOrganizationTrail: bool, optional
	IsOrganizationTrail terra.BoolValue `hcl:"is_organization_trail,attr"`
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// S3BucketName: string, required
	S3BucketName terra.StringValue `hcl:"s3_bucket_name,attr" validate:"required"`
	// S3KeyPrefix: string, optional
	S3KeyPrefix terra.StringValue `hcl:"s3_key_prefix,attr"`
	// SnsTopicName: string, optional
	SnsTopicName terra.StringValue `hcl:"sns_topic_name,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// AdvancedEventSelector: min=0
	AdvancedEventSelector []cloudtrail.AdvancedEventSelector `hcl:"advanced_event_selector,block" validate:"min=0"`
	// EventSelector: min=0,max=5
	EventSelector []cloudtrail.EventSelector `hcl:"event_selector,block" validate:"min=0,max=5"`
	// InsightSelector: min=0
	InsightSelector []cloudtrail.InsightSelector `hcl:"insight_selector,block" validate:"min=0"`
}
type cloudtrailAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_cloudtrail.
func (c cloudtrailAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("arn"))
}

// CloudWatchLogsGroupArn returns a reference to field cloud_watch_logs_group_arn of aws_cloudtrail.
func (c cloudtrailAttributes) CloudWatchLogsGroupArn() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("cloud_watch_logs_group_arn"))
}

// CloudWatchLogsRoleArn returns a reference to field cloud_watch_logs_role_arn of aws_cloudtrail.
func (c cloudtrailAttributes) CloudWatchLogsRoleArn() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("cloud_watch_logs_role_arn"))
}

// EnableLogFileValidation returns a reference to field enable_log_file_validation of aws_cloudtrail.
func (c cloudtrailAttributes) EnableLogFileValidation() terra.BoolValue {
	return terra.ReferenceAsBool(c.ref.Append("enable_log_file_validation"))
}

// EnableLogging returns a reference to field enable_logging of aws_cloudtrail.
func (c cloudtrailAttributes) EnableLogging() terra.BoolValue {
	return terra.ReferenceAsBool(c.ref.Append("enable_logging"))
}

// HomeRegion returns a reference to field home_region of aws_cloudtrail.
func (c cloudtrailAttributes) HomeRegion() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("home_region"))
}

// Id returns a reference to field id of aws_cloudtrail.
func (c cloudtrailAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("id"))
}

// IncludeGlobalServiceEvents returns a reference to field include_global_service_events of aws_cloudtrail.
func (c cloudtrailAttributes) IncludeGlobalServiceEvents() terra.BoolValue {
	return terra.ReferenceAsBool(c.ref.Append("include_global_service_events"))
}

// IsMultiRegionTrail returns a reference to field is_multi_region_trail of aws_cloudtrail.
func (c cloudtrailAttributes) IsMultiRegionTrail() terra.BoolValue {
	return terra.ReferenceAsBool(c.ref.Append("is_multi_region_trail"))
}

// IsOrganizationTrail returns a reference to field is_organization_trail of aws_cloudtrail.
func (c cloudtrailAttributes) IsOrganizationTrail() terra.BoolValue {
	return terra.ReferenceAsBool(c.ref.Append("is_organization_trail"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_cloudtrail.
func (c cloudtrailAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("kms_key_id"))
}

// Name returns a reference to field name of aws_cloudtrail.
func (c cloudtrailAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("name"))
}

// S3BucketName returns a reference to field s3_bucket_name of aws_cloudtrail.
func (c cloudtrailAttributes) S3BucketName() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("s3_bucket_name"))
}

// S3KeyPrefix returns a reference to field s3_key_prefix of aws_cloudtrail.
func (c cloudtrailAttributes) S3KeyPrefix() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("s3_key_prefix"))
}

// SnsTopicName returns a reference to field sns_topic_name of aws_cloudtrail.
func (c cloudtrailAttributes) SnsTopicName() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("sns_topic_name"))
}

// Tags returns a reference to field tags of aws_cloudtrail.
func (c cloudtrailAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](c.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_cloudtrail.
func (c cloudtrailAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](c.ref.Append("tags_all"))
}

func (c cloudtrailAttributes) AdvancedEventSelector() terra.ListValue[cloudtrail.AdvancedEventSelectorAttributes] {
	return terra.ReferenceAsList[cloudtrail.AdvancedEventSelectorAttributes](c.ref.Append("advanced_event_selector"))
}

func (c cloudtrailAttributes) EventSelector() terra.ListValue[cloudtrail.EventSelectorAttributes] {
	return terra.ReferenceAsList[cloudtrail.EventSelectorAttributes](c.ref.Append("event_selector"))
}

func (c cloudtrailAttributes) InsightSelector() terra.ListValue[cloudtrail.InsightSelectorAttributes] {
	return terra.ReferenceAsList[cloudtrail.InsightSelectorAttributes](c.ref.Append("insight_selector"))
}

type cloudtrailState struct {
	Arn                        string                                  `json:"arn"`
	CloudWatchLogsGroupArn     string                                  `json:"cloud_watch_logs_group_arn"`
	CloudWatchLogsRoleArn      string                                  `json:"cloud_watch_logs_role_arn"`
	EnableLogFileValidation    bool                                    `json:"enable_log_file_validation"`
	EnableLogging              bool                                    `json:"enable_logging"`
	HomeRegion                 string                                  `json:"home_region"`
	Id                         string                                  `json:"id"`
	IncludeGlobalServiceEvents bool                                    `json:"include_global_service_events"`
	IsMultiRegionTrail         bool                                    `json:"is_multi_region_trail"`
	IsOrganizationTrail        bool                                    `json:"is_organization_trail"`
	KmsKeyId                   string                                  `json:"kms_key_id"`
	Name                       string                                  `json:"name"`
	S3BucketName               string                                  `json:"s3_bucket_name"`
	S3KeyPrefix                string                                  `json:"s3_key_prefix"`
	SnsTopicName               string                                  `json:"sns_topic_name"`
	Tags                       map[string]string                       `json:"tags"`
	TagsAll                    map[string]string                       `json:"tags_all"`
	AdvancedEventSelector      []cloudtrail.AdvancedEventSelectorState `json:"advanced_event_selector"`
	EventSelector              []cloudtrail.EventSelectorState         `json:"event_selector"`
	InsightSelector            []cloudtrail.InsightSelectorState       `json:"insight_selector"`
}
