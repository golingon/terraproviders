// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_db_event_subscription

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_db_event_subscription.
type Resource struct {
	Name      string
	Args      Args
	state     *awsDbEventSubscriptionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (ades *Resource) Type() string {
	return "aws_db_event_subscription"
}

// LocalName returns the local name for [Resource].
func (ades *Resource) LocalName() string {
	return ades.Name
}

// Configuration returns the configuration (args) for [Resource].
func (ades *Resource) Configuration() interface{} {
	return ades.Args
}

// DependOn is used for other resources to depend on [Resource].
func (ades *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(ades)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (ades *Resource) Dependencies() terra.Dependencies {
	return ades.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (ades *Resource) LifecycleManagement() *terra.Lifecycle {
	return ades.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (ades *Resource) Attributes() awsDbEventSubscriptionAttributes {
	return awsDbEventSubscriptionAttributes{ref: terra.ReferenceResource(ades)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (ades *Resource) ImportState(state io.Reader) error {
	ades.state = &awsDbEventSubscriptionState{}
	if err := json.NewDecoder(state).Decode(ades.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ades.Type(), ades.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (ades *Resource) State() (*awsDbEventSubscriptionState, bool) {
	return ades.state, ades.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (ades *Resource) StateMust() *awsDbEventSubscriptionState {
	if ades.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ades.Type(), ades.LocalName()))
	}
	return ades.state
}

// Args contains the configurations for aws_db_event_subscription.
type Args struct {
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
	// SnsTopic: string, required
	SnsTopic terra.StringValue `hcl:"sns_topic,attr" validate:"required"`
	// SourceIds: set of string, optional
	SourceIds terra.SetValue[terra.StringValue] `hcl:"source_ids,attr"`
	// SourceType: string, optional
	SourceType terra.StringValue `hcl:"source_type,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsDbEventSubscriptionAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_db_event_subscription.
func (ades awsDbEventSubscriptionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ades.ref.Append("arn"))
}

// CustomerAwsId returns a reference to field customer_aws_id of aws_db_event_subscription.
func (ades awsDbEventSubscriptionAttributes) CustomerAwsId() terra.StringValue {
	return terra.ReferenceAsString(ades.ref.Append("customer_aws_id"))
}

// Enabled returns a reference to field enabled of aws_db_event_subscription.
func (ades awsDbEventSubscriptionAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(ades.ref.Append("enabled"))
}

// EventCategories returns a reference to field event_categories of aws_db_event_subscription.
func (ades awsDbEventSubscriptionAttributes) EventCategories() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ades.ref.Append("event_categories"))
}

// Id returns a reference to field id of aws_db_event_subscription.
func (ades awsDbEventSubscriptionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ades.ref.Append("id"))
}

// Name returns a reference to field name of aws_db_event_subscription.
func (ades awsDbEventSubscriptionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ades.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_db_event_subscription.
func (ades awsDbEventSubscriptionAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(ades.ref.Append("name_prefix"))
}

// SnsTopic returns a reference to field sns_topic of aws_db_event_subscription.
func (ades awsDbEventSubscriptionAttributes) SnsTopic() terra.StringValue {
	return terra.ReferenceAsString(ades.ref.Append("sns_topic"))
}

// SourceIds returns a reference to field source_ids of aws_db_event_subscription.
func (ades awsDbEventSubscriptionAttributes) SourceIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ades.ref.Append("source_ids"))
}

// SourceType returns a reference to field source_type of aws_db_event_subscription.
func (ades awsDbEventSubscriptionAttributes) SourceType() terra.StringValue {
	return terra.ReferenceAsString(ades.ref.Append("source_type"))
}

// Tags returns a reference to field tags of aws_db_event_subscription.
func (ades awsDbEventSubscriptionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ades.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_db_event_subscription.
func (ades awsDbEventSubscriptionAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ades.ref.Append("tags_all"))
}

func (ades awsDbEventSubscriptionAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](ades.ref.Append("timeouts"))
}

type awsDbEventSubscriptionState struct {
	Arn             string            `json:"arn"`
	CustomerAwsId   string            `json:"customer_aws_id"`
	Enabled         bool              `json:"enabled"`
	EventCategories []string          `json:"event_categories"`
	Id              string            `json:"id"`
	Name            string            `json:"name"`
	NamePrefix      string            `json:"name_prefix"`
	SnsTopic        string            `json:"sns_topic"`
	SourceIds       []string          `json:"source_ids"`
	SourceType      string            `json:"source_type"`
	Tags            map[string]string `json:"tags"`
	TagsAll         map[string]string `json:"tags_all"`
	Timeouts        *TimeoutsState    `json:"timeouts"`
}
