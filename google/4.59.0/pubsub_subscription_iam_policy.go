// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewPubsubSubscriptionIamPolicy(name string, args PubsubSubscriptionIamPolicyArgs) *PubsubSubscriptionIamPolicy {
	return &PubsubSubscriptionIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PubsubSubscriptionIamPolicy)(nil)

type PubsubSubscriptionIamPolicy struct {
	Name  string
	Args  PubsubSubscriptionIamPolicyArgs
	state *pubsubSubscriptionIamPolicyState
}

func (psip *PubsubSubscriptionIamPolicy) Type() string {
	return "google_pubsub_subscription_iam_policy"
}

func (psip *PubsubSubscriptionIamPolicy) LocalName() string {
	return psip.Name
}

func (psip *PubsubSubscriptionIamPolicy) Configuration() interface{} {
	return psip.Args
}

func (psip *PubsubSubscriptionIamPolicy) Attributes() pubsubSubscriptionIamPolicyAttributes {
	return pubsubSubscriptionIamPolicyAttributes{ref: terra.ReferenceResource(psip)}
}

func (psip *PubsubSubscriptionIamPolicy) ImportState(av io.Reader) error {
	psip.state = &pubsubSubscriptionIamPolicyState{}
	if err := json.NewDecoder(av).Decode(psip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", psip.Type(), psip.LocalName(), err)
	}
	return nil
}

func (psip *PubsubSubscriptionIamPolicy) State() (*pubsubSubscriptionIamPolicyState, bool) {
	return psip.state, psip.state != nil
}

func (psip *PubsubSubscriptionIamPolicy) StateMust() *pubsubSubscriptionIamPolicyState {
	if psip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", psip.Type(), psip.LocalName()))
	}
	return psip.state
}

func (psip *PubsubSubscriptionIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(psip)
}

type PubsubSubscriptionIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Subscription: string, required
	Subscription terra.StringValue `hcl:"subscription,attr" validate:"required"`
	// DependsOn contains resources that PubsubSubscriptionIamPolicy depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type pubsubSubscriptionIamPolicyAttributes struct {
	ref terra.Reference
}

func (psip pubsubSubscriptionIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceString(psip.ref.Append("etag"))
}

func (psip pubsubSubscriptionIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceString(psip.ref.Append("id"))
}

func (psip pubsubSubscriptionIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceString(psip.ref.Append("policy_data"))
}

func (psip pubsubSubscriptionIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceString(psip.ref.Append("project"))
}

func (psip pubsubSubscriptionIamPolicyAttributes) Subscription() terra.StringValue {
	return terra.ReferenceString(psip.ref.Append("subscription"))
}

type pubsubSubscriptionIamPolicyState struct {
	Etag         string `json:"etag"`
	Id           string `json:"id"`
	PolicyData   string `json:"policy_data"`
	Project      string `json:"project"`
	Subscription string `json:"subscription"`
}
