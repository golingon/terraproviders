// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataPubsubSubscriptionIamPolicy creates a new instance of [DataPubsubSubscriptionIamPolicy].
func NewDataPubsubSubscriptionIamPolicy(name string, args DataPubsubSubscriptionIamPolicyArgs) *DataPubsubSubscriptionIamPolicy {
	return &DataPubsubSubscriptionIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPubsubSubscriptionIamPolicy)(nil)

// DataPubsubSubscriptionIamPolicy represents the Terraform data resource google_pubsub_subscription_iam_policy.
type DataPubsubSubscriptionIamPolicy struct {
	Name string
	Args DataPubsubSubscriptionIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataPubsubSubscriptionIamPolicy].
func (psip *DataPubsubSubscriptionIamPolicy) DataSource() string {
	return "google_pubsub_subscription_iam_policy"
}

// LocalName returns the local name for [DataPubsubSubscriptionIamPolicy].
func (psip *DataPubsubSubscriptionIamPolicy) LocalName() string {
	return psip.Name
}

// Configuration returns the configuration (args) for [DataPubsubSubscriptionIamPolicy].
func (psip *DataPubsubSubscriptionIamPolicy) Configuration() interface{} {
	return psip.Args
}

// Attributes returns the attributes for [DataPubsubSubscriptionIamPolicy].
func (psip *DataPubsubSubscriptionIamPolicy) Attributes() dataPubsubSubscriptionIamPolicyAttributes {
	return dataPubsubSubscriptionIamPolicyAttributes{ref: terra.ReferenceDataResource(psip)}
}

// DataPubsubSubscriptionIamPolicyArgs contains the configurations for google_pubsub_subscription_iam_policy.
type DataPubsubSubscriptionIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Subscription: string, required
	Subscription terra.StringValue `hcl:"subscription,attr" validate:"required"`
}
type dataPubsubSubscriptionIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_pubsub_subscription_iam_policy.
func (psip dataPubsubSubscriptionIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(psip.ref.Append("etag"))
}

// Id returns a reference to field id of google_pubsub_subscription_iam_policy.
func (psip dataPubsubSubscriptionIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(psip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_pubsub_subscription_iam_policy.
func (psip dataPubsubSubscriptionIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(psip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_pubsub_subscription_iam_policy.
func (psip dataPubsubSubscriptionIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(psip.ref.Append("project"))
}

// Subscription returns a reference to field subscription of google_pubsub_subscription_iam_policy.
func (psip dataPubsubSubscriptionIamPolicyAttributes) Subscription() terra.StringValue {
	return terra.ReferenceAsString(psip.ref.Append("subscription"))
}