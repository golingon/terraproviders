// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	datapubsubsubscription "github.com/golingon/terraproviders/google/4.63.1/datapubsubsubscription"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataPubsubSubscription creates a new instance of [DataPubsubSubscription].
func NewDataPubsubSubscription(name string, args DataPubsubSubscriptionArgs) *DataPubsubSubscription {
	return &DataPubsubSubscription{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPubsubSubscription)(nil)

// DataPubsubSubscription represents the Terraform data resource google_pubsub_subscription.
type DataPubsubSubscription struct {
	Name string
	Args DataPubsubSubscriptionArgs
}

// DataSource returns the Terraform object type for [DataPubsubSubscription].
func (ps *DataPubsubSubscription) DataSource() string {
	return "google_pubsub_subscription"
}

// LocalName returns the local name for [DataPubsubSubscription].
func (ps *DataPubsubSubscription) LocalName() string {
	return ps.Name
}

// Configuration returns the configuration (args) for [DataPubsubSubscription].
func (ps *DataPubsubSubscription) Configuration() interface{} {
	return ps.Args
}

// Attributes returns the attributes for [DataPubsubSubscription].
func (ps *DataPubsubSubscription) Attributes() dataPubsubSubscriptionAttributes {
	return dataPubsubSubscriptionAttributes{ref: terra.ReferenceDataResource(ps)}
}

// DataPubsubSubscriptionArgs contains the configurations for google_pubsub_subscription.
type DataPubsubSubscriptionArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// BigqueryConfig: min=0
	BigqueryConfig []datapubsubsubscription.BigqueryConfig `hcl:"bigquery_config,block" validate:"min=0"`
	// DeadLetterPolicy: min=0
	DeadLetterPolicy []datapubsubsubscription.DeadLetterPolicy `hcl:"dead_letter_policy,block" validate:"min=0"`
	// ExpirationPolicy: min=0
	ExpirationPolicy []datapubsubsubscription.ExpirationPolicy `hcl:"expiration_policy,block" validate:"min=0"`
	// PushConfig: min=0
	PushConfig []datapubsubsubscription.PushConfig `hcl:"push_config,block" validate:"min=0"`
	// RetryPolicy: min=0
	RetryPolicy []datapubsubsubscription.RetryPolicy `hcl:"retry_policy,block" validate:"min=0"`
}
type dataPubsubSubscriptionAttributes struct {
	ref terra.Reference
}

// AckDeadlineSeconds returns a reference to field ack_deadline_seconds of google_pubsub_subscription.
func (ps dataPubsubSubscriptionAttributes) AckDeadlineSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(ps.ref.Append("ack_deadline_seconds"))
}

// EnableExactlyOnceDelivery returns a reference to field enable_exactly_once_delivery of google_pubsub_subscription.
func (ps dataPubsubSubscriptionAttributes) EnableExactlyOnceDelivery() terra.BoolValue {
	return terra.ReferenceAsBool(ps.ref.Append("enable_exactly_once_delivery"))
}

// EnableMessageOrdering returns a reference to field enable_message_ordering of google_pubsub_subscription.
func (ps dataPubsubSubscriptionAttributes) EnableMessageOrdering() terra.BoolValue {
	return terra.ReferenceAsBool(ps.ref.Append("enable_message_ordering"))
}

// Filter returns a reference to field filter of google_pubsub_subscription.
func (ps dataPubsubSubscriptionAttributes) Filter() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("filter"))
}

// Id returns a reference to field id of google_pubsub_subscription.
func (ps dataPubsubSubscriptionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("id"))
}

// Labels returns a reference to field labels of google_pubsub_subscription.
func (ps dataPubsubSubscriptionAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ps.ref.Append("labels"))
}

// MessageRetentionDuration returns a reference to field message_retention_duration of google_pubsub_subscription.
func (ps dataPubsubSubscriptionAttributes) MessageRetentionDuration() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("message_retention_duration"))
}

// Name returns a reference to field name of google_pubsub_subscription.
func (ps dataPubsubSubscriptionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("name"))
}

// Project returns a reference to field project of google_pubsub_subscription.
func (ps dataPubsubSubscriptionAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("project"))
}

// RetainAckedMessages returns a reference to field retain_acked_messages of google_pubsub_subscription.
func (ps dataPubsubSubscriptionAttributes) RetainAckedMessages() terra.BoolValue {
	return terra.ReferenceAsBool(ps.ref.Append("retain_acked_messages"))
}

// Topic returns a reference to field topic of google_pubsub_subscription.
func (ps dataPubsubSubscriptionAttributes) Topic() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("topic"))
}

func (ps dataPubsubSubscriptionAttributes) BigqueryConfig() terra.ListValue[datapubsubsubscription.BigqueryConfigAttributes] {
	return terra.ReferenceAsList[datapubsubsubscription.BigqueryConfigAttributes](ps.ref.Append("bigquery_config"))
}

func (ps dataPubsubSubscriptionAttributes) DeadLetterPolicy() terra.ListValue[datapubsubsubscription.DeadLetterPolicyAttributes] {
	return terra.ReferenceAsList[datapubsubsubscription.DeadLetterPolicyAttributes](ps.ref.Append("dead_letter_policy"))
}

func (ps dataPubsubSubscriptionAttributes) ExpirationPolicy() terra.ListValue[datapubsubsubscription.ExpirationPolicyAttributes] {
	return terra.ReferenceAsList[datapubsubsubscription.ExpirationPolicyAttributes](ps.ref.Append("expiration_policy"))
}

func (ps dataPubsubSubscriptionAttributes) PushConfig() terra.ListValue[datapubsubsubscription.PushConfigAttributes] {
	return terra.ReferenceAsList[datapubsubsubscription.PushConfigAttributes](ps.ref.Append("push_config"))
}

func (ps dataPubsubSubscriptionAttributes) RetryPolicy() terra.ListValue[datapubsubsubscription.RetryPolicyAttributes] {
	return terra.ReferenceAsList[datapubsubsubscription.RetryPolicyAttributes](ps.ref.Append("retry_policy"))
}
