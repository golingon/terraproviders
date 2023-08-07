// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	pubsubsubscriptioniammember "github.com/golingon/terraproviders/google/4.76.0/pubsubsubscriptioniammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPubsubSubscriptionIamMember creates a new instance of [PubsubSubscriptionIamMember].
func NewPubsubSubscriptionIamMember(name string, args PubsubSubscriptionIamMemberArgs) *PubsubSubscriptionIamMember {
	return &PubsubSubscriptionIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PubsubSubscriptionIamMember)(nil)

// PubsubSubscriptionIamMember represents the Terraform resource google_pubsub_subscription_iam_member.
type PubsubSubscriptionIamMember struct {
	Name      string
	Args      PubsubSubscriptionIamMemberArgs
	state     *pubsubSubscriptionIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PubsubSubscriptionIamMember].
func (psim *PubsubSubscriptionIamMember) Type() string {
	return "google_pubsub_subscription_iam_member"
}

// LocalName returns the local name for [PubsubSubscriptionIamMember].
func (psim *PubsubSubscriptionIamMember) LocalName() string {
	return psim.Name
}

// Configuration returns the configuration (args) for [PubsubSubscriptionIamMember].
func (psim *PubsubSubscriptionIamMember) Configuration() interface{} {
	return psim.Args
}

// DependOn is used for other resources to depend on [PubsubSubscriptionIamMember].
func (psim *PubsubSubscriptionIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(psim)
}

// Dependencies returns the list of resources [PubsubSubscriptionIamMember] depends_on.
func (psim *PubsubSubscriptionIamMember) Dependencies() terra.Dependencies {
	return psim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PubsubSubscriptionIamMember].
func (psim *PubsubSubscriptionIamMember) LifecycleManagement() *terra.Lifecycle {
	return psim.Lifecycle
}

// Attributes returns the attributes for [PubsubSubscriptionIamMember].
func (psim *PubsubSubscriptionIamMember) Attributes() pubsubSubscriptionIamMemberAttributes {
	return pubsubSubscriptionIamMemberAttributes{ref: terra.ReferenceResource(psim)}
}

// ImportState imports the given attribute values into [PubsubSubscriptionIamMember]'s state.
func (psim *PubsubSubscriptionIamMember) ImportState(av io.Reader) error {
	psim.state = &pubsubSubscriptionIamMemberState{}
	if err := json.NewDecoder(av).Decode(psim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", psim.Type(), psim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PubsubSubscriptionIamMember] has state.
func (psim *PubsubSubscriptionIamMember) State() (*pubsubSubscriptionIamMemberState, bool) {
	return psim.state, psim.state != nil
}

// StateMust returns the state for [PubsubSubscriptionIamMember]. Panics if the state is nil.
func (psim *PubsubSubscriptionIamMember) StateMust() *pubsubSubscriptionIamMemberState {
	if psim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", psim.Type(), psim.LocalName()))
	}
	return psim.state
}

// PubsubSubscriptionIamMemberArgs contains the configurations for google_pubsub_subscription_iam_member.
type PubsubSubscriptionIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Subscription: string, required
	Subscription terra.StringValue `hcl:"subscription,attr" validate:"required"`
	// Condition: optional
	Condition *pubsubsubscriptioniammember.Condition `hcl:"condition,block"`
}
type pubsubSubscriptionIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_pubsub_subscription_iam_member.
func (psim pubsubSubscriptionIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(psim.ref.Append("etag"))
}

// Id returns a reference to field id of google_pubsub_subscription_iam_member.
func (psim pubsubSubscriptionIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(psim.ref.Append("id"))
}

// Member returns a reference to field member of google_pubsub_subscription_iam_member.
func (psim pubsubSubscriptionIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(psim.ref.Append("member"))
}

// Project returns a reference to field project of google_pubsub_subscription_iam_member.
func (psim pubsubSubscriptionIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(psim.ref.Append("project"))
}

// Role returns a reference to field role of google_pubsub_subscription_iam_member.
func (psim pubsubSubscriptionIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(psim.ref.Append("role"))
}

// Subscription returns a reference to field subscription of google_pubsub_subscription_iam_member.
func (psim pubsubSubscriptionIamMemberAttributes) Subscription() terra.StringValue {
	return terra.ReferenceAsString(psim.ref.Append("subscription"))
}

func (psim pubsubSubscriptionIamMemberAttributes) Condition() terra.ListValue[pubsubsubscriptioniammember.ConditionAttributes] {
	return terra.ReferenceAsList[pubsubsubscriptioniammember.ConditionAttributes](psim.ref.Append("condition"))
}

type pubsubSubscriptionIamMemberState struct {
	Etag         string                                       `json:"etag"`
	Id           string                                       `json:"id"`
	Member       string                                       `json:"member"`
	Project      string                                       `json:"project"`
	Role         string                                       `json:"role"`
	Subscription string                                       `json:"subscription"`
	Condition    []pubsubsubscriptioniammember.ConditionState `json:"condition"`
}
