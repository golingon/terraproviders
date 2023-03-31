// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	redshifteventsubscription "github.com/golingon/terraproviders/aws/4.60.0/redshifteventsubscription"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRedshiftEventSubscription creates a new instance of [RedshiftEventSubscription].
func NewRedshiftEventSubscription(name string, args RedshiftEventSubscriptionArgs) *RedshiftEventSubscription {
	return &RedshiftEventSubscription{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RedshiftEventSubscription)(nil)

// RedshiftEventSubscription represents the Terraform resource aws_redshift_event_subscription.
type RedshiftEventSubscription struct {
	Name      string
	Args      RedshiftEventSubscriptionArgs
	state     *redshiftEventSubscriptionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RedshiftEventSubscription].
func (res *RedshiftEventSubscription) Type() string {
	return "aws_redshift_event_subscription"
}

// LocalName returns the local name for [RedshiftEventSubscription].
func (res *RedshiftEventSubscription) LocalName() string {
	return res.Name
}

// Configuration returns the configuration (args) for [RedshiftEventSubscription].
func (res *RedshiftEventSubscription) Configuration() interface{} {
	return res.Args
}

// DependOn is used for other resources to depend on [RedshiftEventSubscription].
func (res *RedshiftEventSubscription) DependOn() terra.Reference {
	return terra.ReferenceResource(res)
}

// Dependencies returns the list of resources [RedshiftEventSubscription] depends_on.
func (res *RedshiftEventSubscription) Dependencies() terra.Dependencies {
	return res.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RedshiftEventSubscription].
func (res *RedshiftEventSubscription) LifecycleManagement() *terra.Lifecycle {
	return res.Lifecycle
}

// Attributes returns the attributes for [RedshiftEventSubscription].
func (res *RedshiftEventSubscription) Attributes() redshiftEventSubscriptionAttributes {
	return redshiftEventSubscriptionAttributes{ref: terra.ReferenceResource(res)}
}

// ImportState imports the given attribute values into [RedshiftEventSubscription]'s state.
func (res *RedshiftEventSubscription) ImportState(av io.Reader) error {
	res.state = &redshiftEventSubscriptionState{}
	if err := json.NewDecoder(av).Decode(res.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", res.Type(), res.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RedshiftEventSubscription] has state.
func (res *RedshiftEventSubscription) State() (*redshiftEventSubscriptionState, bool) {
	return res.state, res.state != nil
}

// StateMust returns the state for [RedshiftEventSubscription]. Panics if the state is nil.
func (res *RedshiftEventSubscription) StateMust() *redshiftEventSubscriptionState {
	if res.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", res.Type(), res.LocalName()))
	}
	return res.state
}

// RedshiftEventSubscriptionArgs contains the configurations for aws_redshift_event_subscription.
type RedshiftEventSubscriptionArgs struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// EventCategories: set of string, optional
	EventCategories terra.SetValue[terra.StringValue] `hcl:"event_categories,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Severity: string, optional
	Severity terra.StringValue `hcl:"severity,attr"`
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
	Timeouts *redshifteventsubscription.Timeouts `hcl:"timeouts,block"`
}
type redshiftEventSubscriptionAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_redshift_event_subscription.
func (res redshiftEventSubscriptionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(res.ref.Append("arn"))
}

// CustomerAwsId returns a reference to field customer_aws_id of aws_redshift_event_subscription.
func (res redshiftEventSubscriptionAttributes) CustomerAwsId() terra.StringValue {
	return terra.ReferenceAsString(res.ref.Append("customer_aws_id"))
}

// Enabled returns a reference to field enabled of aws_redshift_event_subscription.
func (res redshiftEventSubscriptionAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(res.ref.Append("enabled"))
}

// EventCategories returns a reference to field event_categories of aws_redshift_event_subscription.
func (res redshiftEventSubscriptionAttributes) EventCategories() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](res.ref.Append("event_categories"))
}

// Id returns a reference to field id of aws_redshift_event_subscription.
func (res redshiftEventSubscriptionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(res.ref.Append("id"))
}

// Name returns a reference to field name of aws_redshift_event_subscription.
func (res redshiftEventSubscriptionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(res.ref.Append("name"))
}

// Severity returns a reference to field severity of aws_redshift_event_subscription.
func (res redshiftEventSubscriptionAttributes) Severity() terra.StringValue {
	return terra.ReferenceAsString(res.ref.Append("severity"))
}

// SnsTopicArn returns a reference to field sns_topic_arn of aws_redshift_event_subscription.
func (res redshiftEventSubscriptionAttributes) SnsTopicArn() terra.StringValue {
	return terra.ReferenceAsString(res.ref.Append("sns_topic_arn"))
}

// SourceIds returns a reference to field source_ids of aws_redshift_event_subscription.
func (res redshiftEventSubscriptionAttributes) SourceIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](res.ref.Append("source_ids"))
}

// SourceType returns a reference to field source_type of aws_redshift_event_subscription.
func (res redshiftEventSubscriptionAttributes) SourceType() terra.StringValue {
	return terra.ReferenceAsString(res.ref.Append("source_type"))
}

// Status returns a reference to field status of aws_redshift_event_subscription.
func (res redshiftEventSubscriptionAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(res.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_redshift_event_subscription.
func (res redshiftEventSubscriptionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](res.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_redshift_event_subscription.
func (res redshiftEventSubscriptionAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](res.ref.Append("tags_all"))
}

func (res redshiftEventSubscriptionAttributes) Timeouts() redshifteventsubscription.TimeoutsAttributes {
	return terra.ReferenceAsSingle[redshifteventsubscription.TimeoutsAttributes](res.ref.Append("timeouts"))
}

type redshiftEventSubscriptionState struct {
	Arn             string                                   `json:"arn"`
	CustomerAwsId   string                                   `json:"customer_aws_id"`
	Enabled         bool                                     `json:"enabled"`
	EventCategories []string                                 `json:"event_categories"`
	Id              string                                   `json:"id"`
	Name            string                                   `json:"name"`
	Severity        string                                   `json:"severity"`
	SnsTopicArn     string                                   `json:"sns_topic_arn"`
	SourceIds       []string                                 `json:"source_ids"`
	SourceType      string                                   `json:"source_type"`
	Status          string                                   `json:"status"`
	Tags            map[string]string                        `json:"tags"`
	TagsAll         map[string]string                        `json:"tags_all"`
	Timeouts        *redshifteventsubscription.TimeoutsState `json:"timeouts"`
}
