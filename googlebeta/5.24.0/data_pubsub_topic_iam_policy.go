// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import "github.com/golingon/lingon/pkg/terra"

// NewDataPubsubTopicIamPolicy creates a new instance of [DataPubsubTopicIamPolicy].
func NewDataPubsubTopicIamPolicy(name string, args DataPubsubTopicIamPolicyArgs) *DataPubsubTopicIamPolicy {
	return &DataPubsubTopicIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPubsubTopicIamPolicy)(nil)

// DataPubsubTopicIamPolicy represents the Terraform data resource google_pubsub_topic_iam_policy.
type DataPubsubTopicIamPolicy struct {
	Name string
	Args DataPubsubTopicIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataPubsubTopicIamPolicy].
func (ptip *DataPubsubTopicIamPolicy) DataSource() string {
	return "google_pubsub_topic_iam_policy"
}

// LocalName returns the local name for [DataPubsubTopicIamPolicy].
func (ptip *DataPubsubTopicIamPolicy) LocalName() string {
	return ptip.Name
}

// Configuration returns the configuration (args) for [DataPubsubTopicIamPolicy].
func (ptip *DataPubsubTopicIamPolicy) Configuration() interface{} {
	return ptip.Args
}

// Attributes returns the attributes for [DataPubsubTopicIamPolicy].
func (ptip *DataPubsubTopicIamPolicy) Attributes() dataPubsubTopicIamPolicyAttributes {
	return dataPubsubTopicIamPolicyAttributes{ref: terra.ReferenceDataResource(ptip)}
}

// DataPubsubTopicIamPolicyArgs contains the configurations for google_pubsub_topic_iam_policy.
type DataPubsubTopicIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Topic: string, required
	Topic terra.StringValue `hcl:"topic,attr" validate:"required"`
}
type dataPubsubTopicIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_pubsub_topic_iam_policy.
func (ptip dataPubsubTopicIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ptip.ref.Append("etag"))
}

// Id returns a reference to field id of google_pubsub_topic_iam_policy.
func (ptip dataPubsubTopicIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ptip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_pubsub_topic_iam_policy.
func (ptip dataPubsubTopicIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(ptip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_pubsub_topic_iam_policy.
func (ptip dataPubsubTopicIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ptip.ref.Append("project"))
}

// Topic returns a reference to field topic of google_pubsub_topic_iam_policy.
func (ptip dataPubsubTopicIamPolicyAttributes) Topic() terra.StringValue {
	return terra.ReferenceAsString(ptip.ref.Append("topic"))
}
