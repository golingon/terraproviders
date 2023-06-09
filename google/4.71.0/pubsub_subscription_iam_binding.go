// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	pubsubsubscriptioniambinding "github.com/golingon/terraproviders/google/4.71.0/pubsubsubscriptioniambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPubsubSubscriptionIamBinding creates a new instance of [PubsubSubscriptionIamBinding].
func NewPubsubSubscriptionIamBinding(name string, args PubsubSubscriptionIamBindingArgs) *PubsubSubscriptionIamBinding {
	return &PubsubSubscriptionIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PubsubSubscriptionIamBinding)(nil)

// PubsubSubscriptionIamBinding represents the Terraform resource google_pubsub_subscription_iam_binding.
type PubsubSubscriptionIamBinding struct {
	Name      string
	Args      PubsubSubscriptionIamBindingArgs
	state     *pubsubSubscriptionIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PubsubSubscriptionIamBinding].
func (psib *PubsubSubscriptionIamBinding) Type() string {
	return "google_pubsub_subscription_iam_binding"
}

// LocalName returns the local name for [PubsubSubscriptionIamBinding].
func (psib *PubsubSubscriptionIamBinding) LocalName() string {
	return psib.Name
}

// Configuration returns the configuration (args) for [PubsubSubscriptionIamBinding].
func (psib *PubsubSubscriptionIamBinding) Configuration() interface{} {
	return psib.Args
}

// DependOn is used for other resources to depend on [PubsubSubscriptionIamBinding].
func (psib *PubsubSubscriptionIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(psib)
}

// Dependencies returns the list of resources [PubsubSubscriptionIamBinding] depends_on.
func (psib *PubsubSubscriptionIamBinding) Dependencies() terra.Dependencies {
	return psib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PubsubSubscriptionIamBinding].
func (psib *PubsubSubscriptionIamBinding) LifecycleManagement() *terra.Lifecycle {
	return psib.Lifecycle
}

// Attributes returns the attributes for [PubsubSubscriptionIamBinding].
func (psib *PubsubSubscriptionIamBinding) Attributes() pubsubSubscriptionIamBindingAttributes {
	return pubsubSubscriptionIamBindingAttributes{ref: terra.ReferenceResource(psib)}
}

// ImportState imports the given attribute values into [PubsubSubscriptionIamBinding]'s state.
func (psib *PubsubSubscriptionIamBinding) ImportState(av io.Reader) error {
	psib.state = &pubsubSubscriptionIamBindingState{}
	if err := json.NewDecoder(av).Decode(psib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", psib.Type(), psib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PubsubSubscriptionIamBinding] has state.
func (psib *PubsubSubscriptionIamBinding) State() (*pubsubSubscriptionIamBindingState, bool) {
	return psib.state, psib.state != nil
}

// StateMust returns the state for [PubsubSubscriptionIamBinding]. Panics if the state is nil.
func (psib *PubsubSubscriptionIamBinding) StateMust() *pubsubSubscriptionIamBindingState {
	if psib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", psib.Type(), psib.LocalName()))
	}
	return psib.state
}

// PubsubSubscriptionIamBindingArgs contains the configurations for google_pubsub_subscription_iam_binding.
type PubsubSubscriptionIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Subscription: string, required
	Subscription terra.StringValue `hcl:"subscription,attr" validate:"required"`
	// Condition: optional
	Condition *pubsubsubscriptioniambinding.Condition `hcl:"condition,block"`
}
type pubsubSubscriptionIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_pubsub_subscription_iam_binding.
func (psib pubsubSubscriptionIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(psib.ref.Append("etag"))
}

// Id returns a reference to field id of google_pubsub_subscription_iam_binding.
func (psib pubsubSubscriptionIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(psib.ref.Append("id"))
}

// Members returns a reference to field members of google_pubsub_subscription_iam_binding.
func (psib pubsubSubscriptionIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](psib.ref.Append("members"))
}

// Project returns a reference to field project of google_pubsub_subscription_iam_binding.
func (psib pubsubSubscriptionIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(psib.ref.Append("project"))
}

// Role returns a reference to field role of google_pubsub_subscription_iam_binding.
func (psib pubsubSubscriptionIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(psib.ref.Append("role"))
}

// Subscription returns a reference to field subscription of google_pubsub_subscription_iam_binding.
func (psib pubsubSubscriptionIamBindingAttributes) Subscription() terra.StringValue {
	return terra.ReferenceAsString(psib.ref.Append("subscription"))
}

func (psib pubsubSubscriptionIamBindingAttributes) Condition() terra.ListValue[pubsubsubscriptioniambinding.ConditionAttributes] {
	return terra.ReferenceAsList[pubsubsubscriptioniambinding.ConditionAttributes](psib.ref.Append("condition"))
}

type pubsubSubscriptionIamBindingState struct {
	Etag         string                                        `json:"etag"`
	Id           string                                        `json:"id"`
	Members      []string                                      `json:"members"`
	Project      string                                        `json:"project"`
	Role         string                                        `json:"role"`
	Subscription string                                        `json:"subscription"`
	Condition    []pubsubsubscriptioniambinding.ConditionState `json:"condition"`
}
