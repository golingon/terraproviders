// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	docdbeventsubscription "github.com/golingon/terraproviders/aws/5.8.0/docdbeventsubscription"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDocdbEventSubscription creates a new instance of [DocdbEventSubscription].
func NewDocdbEventSubscription(name string, args DocdbEventSubscriptionArgs) *DocdbEventSubscription {
	return &DocdbEventSubscription{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DocdbEventSubscription)(nil)

// DocdbEventSubscription represents the Terraform resource aws_docdb_event_subscription.
type DocdbEventSubscription struct {
	Name      string
	Args      DocdbEventSubscriptionArgs
	state     *docdbEventSubscriptionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DocdbEventSubscription].
func (des *DocdbEventSubscription) Type() string {
	return "aws_docdb_event_subscription"
}

// LocalName returns the local name for [DocdbEventSubscription].
func (des *DocdbEventSubscription) LocalName() string {
	return des.Name
}

// Configuration returns the configuration (args) for [DocdbEventSubscription].
func (des *DocdbEventSubscription) Configuration() interface{} {
	return des.Args
}

// DependOn is used for other resources to depend on [DocdbEventSubscription].
func (des *DocdbEventSubscription) DependOn() terra.Reference {
	return terra.ReferenceResource(des)
}

// Dependencies returns the list of resources [DocdbEventSubscription] depends_on.
func (des *DocdbEventSubscription) Dependencies() terra.Dependencies {
	return des.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DocdbEventSubscription].
func (des *DocdbEventSubscription) LifecycleManagement() *terra.Lifecycle {
	return des.Lifecycle
}

// Attributes returns the attributes for [DocdbEventSubscription].
func (des *DocdbEventSubscription) Attributes() docdbEventSubscriptionAttributes {
	return docdbEventSubscriptionAttributes{ref: terra.ReferenceResource(des)}
}

// ImportState imports the given attribute values into [DocdbEventSubscription]'s state.
func (des *DocdbEventSubscription) ImportState(av io.Reader) error {
	des.state = &docdbEventSubscriptionState{}
	if err := json.NewDecoder(av).Decode(des.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", des.Type(), des.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DocdbEventSubscription] has state.
func (des *DocdbEventSubscription) State() (*docdbEventSubscriptionState, bool) {
	return des.state, des.state != nil
}

// StateMust returns the state for [DocdbEventSubscription]. Panics if the state is nil.
func (des *DocdbEventSubscription) StateMust() *docdbEventSubscriptionState {
	if des.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", des.Type(), des.LocalName()))
	}
	return des.state
}

// DocdbEventSubscriptionArgs contains the configurations for aws_docdb_event_subscription.
type DocdbEventSubscriptionArgs struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// EventCategories: set of string, optional
	EventCategories terra.SetValue[terra.StringValue] `hcl:"event_categories,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NamePrefix: string, optional
	NamePrefix terra.StringValue `hcl:"name_prefix,attr"`
	// SnsTopicArn: string, required
	SnsTopicArn terra.StringValue `hcl:"sns_topic_arn,attr" validate:"required"`
	// SourceIds: set of string, optional
	SourceIds terra.SetValue[terra.StringValue] `hcl:"source_ids,attr"`
	// SourceType: string, optional
	SourceType terra.StringValue `hcl:"source_type,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Timeouts: optional
	Timeouts *docdbeventsubscription.Timeouts `hcl:"timeouts,block"`
}
type docdbEventSubscriptionAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_docdb_event_subscription.
func (des docdbEventSubscriptionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(des.ref.Append("arn"))
}

// CustomerAwsId returns a reference to field customer_aws_id of aws_docdb_event_subscription.
func (des docdbEventSubscriptionAttributes) CustomerAwsId() terra.StringValue {
	return terra.ReferenceAsString(des.ref.Append("customer_aws_id"))
}

// Enabled returns a reference to field enabled of aws_docdb_event_subscription.
func (des docdbEventSubscriptionAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(des.ref.Append("enabled"))
}

// EventCategories returns a reference to field event_categories of aws_docdb_event_subscription.
func (des docdbEventSubscriptionAttributes) EventCategories() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](des.ref.Append("event_categories"))
}

// Id returns a reference to field id of aws_docdb_event_subscription.
func (des docdbEventSubscriptionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(des.ref.Append("id"))
}

// Name returns a reference to field name of aws_docdb_event_subscription.
func (des docdbEventSubscriptionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(des.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_docdb_event_subscription.
func (des docdbEventSubscriptionAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(des.ref.Append("name_prefix"))
}

// SnsTopicArn returns a reference to field sns_topic_arn of aws_docdb_event_subscription.
func (des docdbEventSubscriptionAttributes) SnsTopicArn() terra.StringValue {
	return terra.ReferenceAsString(des.ref.Append("sns_topic_arn"))
}

// SourceIds returns a reference to field source_ids of aws_docdb_event_subscription.
func (des docdbEventSubscriptionAttributes) SourceIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](des.ref.Append("source_ids"))
}

// SourceType returns a reference to field source_type of aws_docdb_event_subscription.
func (des docdbEventSubscriptionAttributes) SourceType() terra.StringValue {
	return terra.ReferenceAsString(des.ref.Append("source_type"))
}

// Tags returns a reference to field tags of aws_docdb_event_subscription.
func (des docdbEventSubscriptionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](des.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_docdb_event_subscription.
func (des docdbEventSubscriptionAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](des.ref.Append("tags_all"))
}

func (des docdbEventSubscriptionAttributes) Timeouts() docdbeventsubscription.TimeoutsAttributes {
	return terra.ReferenceAsSingle[docdbeventsubscription.TimeoutsAttributes](des.ref.Append("timeouts"))
}

type docdbEventSubscriptionState struct {
	Arn             string                                `json:"arn"`
	CustomerAwsId   string                                `json:"customer_aws_id"`
	Enabled         bool                                  `json:"enabled"`
	EventCategories []string                              `json:"event_categories"`
	Id              string                                `json:"id"`
	Name            string                                `json:"name"`
	NamePrefix      string                                `json:"name_prefix"`
	SnsTopicArn     string                                `json:"sns_topic_arn"`
	SourceIds       []string                              `json:"source_ids"`
	SourceType      string                                `json:"source_type"`
	Tags            map[string]string                     `json:"tags"`
	TagsAll         map[string]string                     `json:"tags_all"`
	Timeouts        *docdbeventsubscription.TimeoutsState `json:"timeouts"`
}
