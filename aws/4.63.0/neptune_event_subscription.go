// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	neptuneeventsubscription "github.com/golingon/terraproviders/aws/4.63.0/neptuneeventsubscription"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNeptuneEventSubscription creates a new instance of [NeptuneEventSubscription].
func NewNeptuneEventSubscription(name string, args NeptuneEventSubscriptionArgs) *NeptuneEventSubscription {
	return &NeptuneEventSubscription{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NeptuneEventSubscription)(nil)

// NeptuneEventSubscription represents the Terraform resource aws_neptune_event_subscription.
type NeptuneEventSubscription struct {
	Name      string
	Args      NeptuneEventSubscriptionArgs
	state     *neptuneEventSubscriptionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NeptuneEventSubscription].
func (nes *NeptuneEventSubscription) Type() string {
	return "aws_neptune_event_subscription"
}

// LocalName returns the local name for [NeptuneEventSubscription].
func (nes *NeptuneEventSubscription) LocalName() string {
	return nes.Name
}

// Configuration returns the configuration (args) for [NeptuneEventSubscription].
func (nes *NeptuneEventSubscription) Configuration() interface{} {
	return nes.Args
}

// DependOn is used for other resources to depend on [NeptuneEventSubscription].
func (nes *NeptuneEventSubscription) DependOn() terra.Reference {
	return terra.ReferenceResource(nes)
}

// Dependencies returns the list of resources [NeptuneEventSubscription] depends_on.
func (nes *NeptuneEventSubscription) Dependencies() terra.Dependencies {
	return nes.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NeptuneEventSubscription].
func (nes *NeptuneEventSubscription) LifecycleManagement() *terra.Lifecycle {
	return nes.Lifecycle
}

// Attributes returns the attributes for [NeptuneEventSubscription].
func (nes *NeptuneEventSubscription) Attributes() neptuneEventSubscriptionAttributes {
	return neptuneEventSubscriptionAttributes{ref: terra.ReferenceResource(nes)}
}

// ImportState imports the given attribute values into [NeptuneEventSubscription]'s state.
func (nes *NeptuneEventSubscription) ImportState(av io.Reader) error {
	nes.state = &neptuneEventSubscriptionState{}
	if err := json.NewDecoder(av).Decode(nes.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nes.Type(), nes.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NeptuneEventSubscription] has state.
func (nes *NeptuneEventSubscription) State() (*neptuneEventSubscriptionState, bool) {
	return nes.state, nes.state != nil
}

// StateMust returns the state for [NeptuneEventSubscription]. Panics if the state is nil.
func (nes *NeptuneEventSubscription) StateMust() *neptuneEventSubscriptionState {
	if nes.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nes.Type(), nes.LocalName()))
	}
	return nes.state
}

// NeptuneEventSubscriptionArgs contains the configurations for aws_neptune_event_subscription.
type NeptuneEventSubscriptionArgs struct {
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
	Timeouts *neptuneeventsubscription.Timeouts `hcl:"timeouts,block"`
}
type neptuneEventSubscriptionAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_neptune_event_subscription.
func (nes neptuneEventSubscriptionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(nes.ref.Append("arn"))
}

// CustomerAwsId returns a reference to field customer_aws_id of aws_neptune_event_subscription.
func (nes neptuneEventSubscriptionAttributes) CustomerAwsId() terra.StringValue {
	return terra.ReferenceAsString(nes.ref.Append("customer_aws_id"))
}

// Enabled returns a reference to field enabled of aws_neptune_event_subscription.
func (nes neptuneEventSubscriptionAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(nes.ref.Append("enabled"))
}

// EventCategories returns a reference to field event_categories of aws_neptune_event_subscription.
func (nes neptuneEventSubscriptionAttributes) EventCategories() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](nes.ref.Append("event_categories"))
}

// Id returns a reference to field id of aws_neptune_event_subscription.
func (nes neptuneEventSubscriptionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nes.ref.Append("id"))
}

// Name returns a reference to field name of aws_neptune_event_subscription.
func (nes neptuneEventSubscriptionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nes.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_neptune_event_subscription.
func (nes neptuneEventSubscriptionAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(nes.ref.Append("name_prefix"))
}

// SnsTopicArn returns a reference to field sns_topic_arn of aws_neptune_event_subscription.
func (nes neptuneEventSubscriptionAttributes) SnsTopicArn() terra.StringValue {
	return terra.ReferenceAsString(nes.ref.Append("sns_topic_arn"))
}

// SourceIds returns a reference to field source_ids of aws_neptune_event_subscription.
func (nes neptuneEventSubscriptionAttributes) SourceIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](nes.ref.Append("source_ids"))
}

// SourceType returns a reference to field source_type of aws_neptune_event_subscription.
func (nes neptuneEventSubscriptionAttributes) SourceType() terra.StringValue {
	return terra.ReferenceAsString(nes.ref.Append("source_type"))
}

// Tags returns a reference to field tags of aws_neptune_event_subscription.
func (nes neptuneEventSubscriptionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nes.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_neptune_event_subscription.
func (nes neptuneEventSubscriptionAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nes.ref.Append("tags_all"))
}

func (nes neptuneEventSubscriptionAttributes) Timeouts() neptuneeventsubscription.TimeoutsAttributes {
	return terra.ReferenceAsSingle[neptuneeventsubscription.TimeoutsAttributes](nes.ref.Append("timeouts"))
}

type neptuneEventSubscriptionState struct {
	Arn             string                                  `json:"arn"`
	CustomerAwsId   string                                  `json:"customer_aws_id"`
	Enabled         bool                                    `json:"enabled"`
	EventCategories []string                                `json:"event_categories"`
	Id              string                                  `json:"id"`
	Name            string                                  `json:"name"`
	NamePrefix      string                                  `json:"name_prefix"`
	SnsTopicArn     string                                  `json:"sns_topic_arn"`
	SourceIds       []string                                `json:"source_ids"`
	SourceType      string                                  `json:"source_type"`
	Tags            map[string]string                       `json:"tags"`
	TagsAll         map[string]string                       `json:"tags_all"`
	Timeouts        *neptuneeventsubscription.TimeoutsState `json:"timeouts"`
}
