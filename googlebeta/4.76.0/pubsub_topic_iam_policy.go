// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPubsubTopicIamPolicy creates a new instance of [PubsubTopicIamPolicy].
func NewPubsubTopicIamPolicy(name string, args PubsubTopicIamPolicyArgs) *PubsubTopicIamPolicy {
	return &PubsubTopicIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PubsubTopicIamPolicy)(nil)

// PubsubTopicIamPolicy represents the Terraform resource google_pubsub_topic_iam_policy.
type PubsubTopicIamPolicy struct {
	Name      string
	Args      PubsubTopicIamPolicyArgs
	state     *pubsubTopicIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PubsubTopicIamPolicy].
func (ptip *PubsubTopicIamPolicy) Type() string {
	return "google_pubsub_topic_iam_policy"
}

// LocalName returns the local name for [PubsubTopicIamPolicy].
func (ptip *PubsubTopicIamPolicy) LocalName() string {
	return ptip.Name
}

// Configuration returns the configuration (args) for [PubsubTopicIamPolicy].
func (ptip *PubsubTopicIamPolicy) Configuration() interface{} {
	return ptip.Args
}

// DependOn is used for other resources to depend on [PubsubTopicIamPolicy].
func (ptip *PubsubTopicIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(ptip)
}

// Dependencies returns the list of resources [PubsubTopicIamPolicy] depends_on.
func (ptip *PubsubTopicIamPolicy) Dependencies() terra.Dependencies {
	return ptip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PubsubTopicIamPolicy].
func (ptip *PubsubTopicIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return ptip.Lifecycle
}

// Attributes returns the attributes for [PubsubTopicIamPolicy].
func (ptip *PubsubTopicIamPolicy) Attributes() pubsubTopicIamPolicyAttributes {
	return pubsubTopicIamPolicyAttributes{ref: terra.ReferenceResource(ptip)}
}

// ImportState imports the given attribute values into [PubsubTopicIamPolicy]'s state.
func (ptip *PubsubTopicIamPolicy) ImportState(av io.Reader) error {
	ptip.state = &pubsubTopicIamPolicyState{}
	if err := json.NewDecoder(av).Decode(ptip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ptip.Type(), ptip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PubsubTopicIamPolicy] has state.
func (ptip *PubsubTopicIamPolicy) State() (*pubsubTopicIamPolicyState, bool) {
	return ptip.state, ptip.state != nil
}

// StateMust returns the state for [PubsubTopicIamPolicy]. Panics if the state is nil.
func (ptip *PubsubTopicIamPolicy) StateMust() *pubsubTopicIamPolicyState {
	if ptip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ptip.Type(), ptip.LocalName()))
	}
	return ptip.state
}

// PubsubTopicIamPolicyArgs contains the configurations for google_pubsub_topic_iam_policy.
type PubsubTopicIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Topic: string, required
	Topic terra.StringValue `hcl:"topic,attr" validate:"required"`
}
type pubsubTopicIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_pubsub_topic_iam_policy.
func (ptip pubsubTopicIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ptip.ref.Append("etag"))
}

// Id returns a reference to field id of google_pubsub_topic_iam_policy.
func (ptip pubsubTopicIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ptip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_pubsub_topic_iam_policy.
func (ptip pubsubTopicIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(ptip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_pubsub_topic_iam_policy.
func (ptip pubsubTopicIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ptip.ref.Append("project"))
}

// Topic returns a reference to field topic of google_pubsub_topic_iam_policy.
func (ptip pubsubTopicIamPolicyAttributes) Topic() terra.StringValue {
	return terra.ReferenceAsString(ptip.ref.Append("topic"))
}

type pubsubTopicIamPolicyState struct {
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	PolicyData string `json:"policy_data"`
	Project    string `json:"project"`
	Topic      string `json:"topic"`
}
