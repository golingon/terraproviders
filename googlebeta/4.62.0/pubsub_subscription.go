// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	pubsubsubscription "github.com/golingon/terraproviders/googlebeta/4.62.0/pubsubsubscription"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPubsubSubscription creates a new instance of [PubsubSubscription].
func NewPubsubSubscription(name string, args PubsubSubscriptionArgs) *PubsubSubscription {
	return &PubsubSubscription{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PubsubSubscription)(nil)

// PubsubSubscription represents the Terraform resource google_pubsub_subscription.
type PubsubSubscription struct {
	Name      string
	Args      PubsubSubscriptionArgs
	state     *pubsubSubscriptionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PubsubSubscription].
func (ps *PubsubSubscription) Type() string {
	return "google_pubsub_subscription"
}

// LocalName returns the local name for [PubsubSubscription].
func (ps *PubsubSubscription) LocalName() string {
	return ps.Name
}

// Configuration returns the configuration (args) for [PubsubSubscription].
func (ps *PubsubSubscription) Configuration() interface{} {
	return ps.Args
}

// DependOn is used for other resources to depend on [PubsubSubscription].
func (ps *PubsubSubscription) DependOn() terra.Reference {
	return terra.ReferenceResource(ps)
}

// Dependencies returns the list of resources [PubsubSubscription] depends_on.
func (ps *PubsubSubscription) Dependencies() terra.Dependencies {
	return ps.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PubsubSubscription].
func (ps *PubsubSubscription) LifecycleManagement() *terra.Lifecycle {
	return ps.Lifecycle
}

// Attributes returns the attributes for [PubsubSubscription].
func (ps *PubsubSubscription) Attributes() pubsubSubscriptionAttributes {
	return pubsubSubscriptionAttributes{ref: terra.ReferenceResource(ps)}
}

// ImportState imports the given attribute values into [PubsubSubscription]'s state.
func (ps *PubsubSubscription) ImportState(av io.Reader) error {
	ps.state = &pubsubSubscriptionState{}
	if err := json.NewDecoder(av).Decode(ps.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ps.Type(), ps.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PubsubSubscription] has state.
func (ps *PubsubSubscription) State() (*pubsubSubscriptionState, bool) {
	return ps.state, ps.state != nil
}

// StateMust returns the state for [PubsubSubscription]. Panics if the state is nil.
func (ps *PubsubSubscription) StateMust() *pubsubSubscriptionState {
	if ps.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ps.Type(), ps.LocalName()))
	}
	return ps.state
}

// PubsubSubscriptionArgs contains the configurations for google_pubsub_subscription.
type PubsubSubscriptionArgs struct {
	// AckDeadlineSeconds: number, optional
	AckDeadlineSeconds terra.NumberValue `hcl:"ack_deadline_seconds,attr"`
	// EnableExactlyOnceDelivery: bool, optional
	EnableExactlyOnceDelivery terra.BoolValue `hcl:"enable_exactly_once_delivery,attr"`
	// EnableMessageOrdering: bool, optional
	EnableMessageOrdering terra.BoolValue `hcl:"enable_message_ordering,attr"`
	// Filter: string, optional
	Filter terra.StringValue `hcl:"filter,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// MessageRetentionDuration: string, optional
	MessageRetentionDuration terra.StringValue `hcl:"message_retention_duration,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// RetainAckedMessages: bool, optional
	RetainAckedMessages terra.BoolValue `hcl:"retain_acked_messages,attr"`
	// Topic: string, required
	Topic terra.StringValue `hcl:"topic,attr" validate:"required"`
	// BigqueryConfig: optional
	BigqueryConfig *pubsubsubscription.BigqueryConfig `hcl:"bigquery_config,block"`
	// DeadLetterPolicy: optional
	DeadLetterPolicy *pubsubsubscription.DeadLetterPolicy `hcl:"dead_letter_policy,block"`
	// ExpirationPolicy: optional
	ExpirationPolicy *pubsubsubscription.ExpirationPolicy `hcl:"expiration_policy,block"`
	// PushConfig: optional
	PushConfig *pubsubsubscription.PushConfig `hcl:"push_config,block"`
	// RetryPolicy: optional
	RetryPolicy *pubsubsubscription.RetryPolicy `hcl:"retry_policy,block"`
	// Timeouts: optional
	Timeouts *pubsubsubscription.Timeouts `hcl:"timeouts,block"`
}
type pubsubSubscriptionAttributes struct {
	ref terra.Reference
}

// AckDeadlineSeconds returns a reference to field ack_deadline_seconds of google_pubsub_subscription.
func (ps pubsubSubscriptionAttributes) AckDeadlineSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(ps.ref.Append("ack_deadline_seconds"))
}

// EnableExactlyOnceDelivery returns a reference to field enable_exactly_once_delivery of google_pubsub_subscription.
func (ps pubsubSubscriptionAttributes) EnableExactlyOnceDelivery() terra.BoolValue {
	return terra.ReferenceAsBool(ps.ref.Append("enable_exactly_once_delivery"))
}

// EnableMessageOrdering returns a reference to field enable_message_ordering of google_pubsub_subscription.
func (ps pubsubSubscriptionAttributes) EnableMessageOrdering() terra.BoolValue {
	return terra.ReferenceAsBool(ps.ref.Append("enable_message_ordering"))
}

// Filter returns a reference to field filter of google_pubsub_subscription.
func (ps pubsubSubscriptionAttributes) Filter() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("filter"))
}

// Id returns a reference to field id of google_pubsub_subscription.
func (ps pubsubSubscriptionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("id"))
}

// Labels returns a reference to field labels of google_pubsub_subscription.
func (ps pubsubSubscriptionAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ps.ref.Append("labels"))
}

// MessageRetentionDuration returns a reference to field message_retention_duration of google_pubsub_subscription.
func (ps pubsubSubscriptionAttributes) MessageRetentionDuration() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("message_retention_duration"))
}

// Name returns a reference to field name of google_pubsub_subscription.
func (ps pubsubSubscriptionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("name"))
}

// Project returns a reference to field project of google_pubsub_subscription.
func (ps pubsubSubscriptionAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("project"))
}

// RetainAckedMessages returns a reference to field retain_acked_messages of google_pubsub_subscription.
func (ps pubsubSubscriptionAttributes) RetainAckedMessages() terra.BoolValue {
	return terra.ReferenceAsBool(ps.ref.Append("retain_acked_messages"))
}

// Topic returns a reference to field topic of google_pubsub_subscription.
func (ps pubsubSubscriptionAttributes) Topic() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("topic"))
}

func (ps pubsubSubscriptionAttributes) BigqueryConfig() terra.ListValue[pubsubsubscription.BigqueryConfigAttributes] {
	return terra.ReferenceAsList[pubsubsubscription.BigqueryConfigAttributes](ps.ref.Append("bigquery_config"))
}

func (ps pubsubSubscriptionAttributes) DeadLetterPolicy() terra.ListValue[pubsubsubscription.DeadLetterPolicyAttributes] {
	return terra.ReferenceAsList[pubsubsubscription.DeadLetterPolicyAttributes](ps.ref.Append("dead_letter_policy"))
}

func (ps pubsubSubscriptionAttributes) ExpirationPolicy() terra.ListValue[pubsubsubscription.ExpirationPolicyAttributes] {
	return terra.ReferenceAsList[pubsubsubscription.ExpirationPolicyAttributes](ps.ref.Append("expiration_policy"))
}

func (ps pubsubSubscriptionAttributes) PushConfig() terra.ListValue[pubsubsubscription.PushConfigAttributes] {
	return terra.ReferenceAsList[pubsubsubscription.PushConfigAttributes](ps.ref.Append("push_config"))
}

func (ps pubsubSubscriptionAttributes) RetryPolicy() terra.ListValue[pubsubsubscription.RetryPolicyAttributes] {
	return terra.ReferenceAsList[pubsubsubscription.RetryPolicyAttributes](ps.ref.Append("retry_policy"))
}

func (ps pubsubSubscriptionAttributes) Timeouts() pubsubsubscription.TimeoutsAttributes {
	return terra.ReferenceAsSingle[pubsubsubscription.TimeoutsAttributes](ps.ref.Append("timeouts"))
}

type pubsubSubscriptionState struct {
	AckDeadlineSeconds        float64                                    `json:"ack_deadline_seconds"`
	EnableExactlyOnceDelivery bool                                       `json:"enable_exactly_once_delivery"`
	EnableMessageOrdering     bool                                       `json:"enable_message_ordering"`
	Filter                    string                                     `json:"filter"`
	Id                        string                                     `json:"id"`
	Labels                    map[string]string                          `json:"labels"`
	MessageRetentionDuration  string                                     `json:"message_retention_duration"`
	Name                      string                                     `json:"name"`
	Project                   string                                     `json:"project"`
	RetainAckedMessages       bool                                       `json:"retain_acked_messages"`
	Topic                     string                                     `json:"topic"`
	BigqueryConfig            []pubsubsubscription.BigqueryConfigState   `json:"bigquery_config"`
	DeadLetterPolicy          []pubsubsubscription.DeadLetterPolicyState `json:"dead_letter_policy"`
	ExpirationPolicy          []pubsubsubscription.ExpirationPolicyState `json:"expiration_policy"`
	PushConfig                []pubsubsubscription.PushConfigState       `json:"push_config"`
	RetryPolicy               []pubsubsubscription.RetryPolicyState      `json:"retry_policy"`
	Timeouts                  *pubsubsubscription.TimeoutsState          `json:"timeouts"`
}
