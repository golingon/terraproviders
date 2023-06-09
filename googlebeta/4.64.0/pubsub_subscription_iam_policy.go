// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPubsubSubscriptionIamPolicy creates a new instance of [PubsubSubscriptionIamPolicy].
func NewPubsubSubscriptionIamPolicy(name string, args PubsubSubscriptionIamPolicyArgs) *PubsubSubscriptionIamPolicy {
	return &PubsubSubscriptionIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PubsubSubscriptionIamPolicy)(nil)

// PubsubSubscriptionIamPolicy represents the Terraform resource google_pubsub_subscription_iam_policy.
type PubsubSubscriptionIamPolicy struct {
	Name      string
	Args      PubsubSubscriptionIamPolicyArgs
	state     *pubsubSubscriptionIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PubsubSubscriptionIamPolicy].
func (psip *PubsubSubscriptionIamPolicy) Type() string {
	return "google_pubsub_subscription_iam_policy"
}

// LocalName returns the local name for [PubsubSubscriptionIamPolicy].
func (psip *PubsubSubscriptionIamPolicy) LocalName() string {
	return psip.Name
}

// Configuration returns the configuration (args) for [PubsubSubscriptionIamPolicy].
func (psip *PubsubSubscriptionIamPolicy) Configuration() interface{} {
	return psip.Args
}

// DependOn is used for other resources to depend on [PubsubSubscriptionIamPolicy].
func (psip *PubsubSubscriptionIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(psip)
}

// Dependencies returns the list of resources [PubsubSubscriptionIamPolicy] depends_on.
func (psip *PubsubSubscriptionIamPolicy) Dependencies() terra.Dependencies {
	return psip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PubsubSubscriptionIamPolicy].
func (psip *PubsubSubscriptionIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return psip.Lifecycle
}

// Attributes returns the attributes for [PubsubSubscriptionIamPolicy].
func (psip *PubsubSubscriptionIamPolicy) Attributes() pubsubSubscriptionIamPolicyAttributes {
	return pubsubSubscriptionIamPolicyAttributes{ref: terra.ReferenceResource(psip)}
}

// ImportState imports the given attribute values into [PubsubSubscriptionIamPolicy]'s state.
func (psip *PubsubSubscriptionIamPolicy) ImportState(av io.Reader) error {
	psip.state = &pubsubSubscriptionIamPolicyState{}
	if err := json.NewDecoder(av).Decode(psip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", psip.Type(), psip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PubsubSubscriptionIamPolicy] has state.
func (psip *PubsubSubscriptionIamPolicy) State() (*pubsubSubscriptionIamPolicyState, bool) {
	return psip.state, psip.state != nil
}

// StateMust returns the state for [PubsubSubscriptionIamPolicy]. Panics if the state is nil.
func (psip *PubsubSubscriptionIamPolicy) StateMust() *pubsubSubscriptionIamPolicyState {
	if psip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", psip.Type(), psip.LocalName()))
	}
	return psip.state
}

// PubsubSubscriptionIamPolicyArgs contains the configurations for google_pubsub_subscription_iam_policy.
type PubsubSubscriptionIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Subscription: string, required
	Subscription terra.StringValue `hcl:"subscription,attr" validate:"required"`
}
type pubsubSubscriptionIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_pubsub_subscription_iam_policy.
func (psip pubsubSubscriptionIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(psip.ref.Append("etag"))
}

// Id returns a reference to field id of google_pubsub_subscription_iam_policy.
func (psip pubsubSubscriptionIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(psip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_pubsub_subscription_iam_policy.
func (psip pubsubSubscriptionIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(psip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_pubsub_subscription_iam_policy.
func (psip pubsubSubscriptionIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(psip.ref.Append("project"))
}

// Subscription returns a reference to field subscription of google_pubsub_subscription_iam_policy.
func (psip pubsubSubscriptionIamPolicyAttributes) Subscription() terra.StringValue {
	return terra.ReferenceAsString(psip.ref.Append("subscription"))
}

type pubsubSubscriptionIamPolicyState struct {
	Etag         string `json:"etag"`
	Id           string `json:"id"`
	PolicyData   string `json:"policy_data"`
	Project      string `json:"project"`
	Subscription string `json:"subscription"`
}
