// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	pubsubtopiciambinding "github.com/golingon/terraproviders/googlebeta/4.71.0/pubsubtopiciambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPubsubTopicIamBinding creates a new instance of [PubsubTopicIamBinding].
func NewPubsubTopicIamBinding(name string, args PubsubTopicIamBindingArgs) *PubsubTopicIamBinding {
	return &PubsubTopicIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PubsubTopicIamBinding)(nil)

// PubsubTopicIamBinding represents the Terraform resource google_pubsub_topic_iam_binding.
type PubsubTopicIamBinding struct {
	Name      string
	Args      PubsubTopicIamBindingArgs
	state     *pubsubTopicIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PubsubTopicIamBinding].
func (ptib *PubsubTopicIamBinding) Type() string {
	return "google_pubsub_topic_iam_binding"
}

// LocalName returns the local name for [PubsubTopicIamBinding].
func (ptib *PubsubTopicIamBinding) LocalName() string {
	return ptib.Name
}

// Configuration returns the configuration (args) for [PubsubTopicIamBinding].
func (ptib *PubsubTopicIamBinding) Configuration() interface{} {
	return ptib.Args
}

// DependOn is used for other resources to depend on [PubsubTopicIamBinding].
func (ptib *PubsubTopicIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(ptib)
}

// Dependencies returns the list of resources [PubsubTopicIamBinding] depends_on.
func (ptib *PubsubTopicIamBinding) Dependencies() terra.Dependencies {
	return ptib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PubsubTopicIamBinding].
func (ptib *PubsubTopicIamBinding) LifecycleManagement() *terra.Lifecycle {
	return ptib.Lifecycle
}

// Attributes returns the attributes for [PubsubTopicIamBinding].
func (ptib *PubsubTopicIamBinding) Attributes() pubsubTopicIamBindingAttributes {
	return pubsubTopicIamBindingAttributes{ref: terra.ReferenceResource(ptib)}
}

// ImportState imports the given attribute values into [PubsubTopicIamBinding]'s state.
func (ptib *PubsubTopicIamBinding) ImportState(av io.Reader) error {
	ptib.state = &pubsubTopicIamBindingState{}
	if err := json.NewDecoder(av).Decode(ptib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ptib.Type(), ptib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PubsubTopicIamBinding] has state.
func (ptib *PubsubTopicIamBinding) State() (*pubsubTopicIamBindingState, bool) {
	return ptib.state, ptib.state != nil
}

// StateMust returns the state for [PubsubTopicIamBinding]. Panics if the state is nil.
func (ptib *PubsubTopicIamBinding) StateMust() *pubsubTopicIamBindingState {
	if ptib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ptib.Type(), ptib.LocalName()))
	}
	return ptib.state
}

// PubsubTopicIamBindingArgs contains the configurations for google_pubsub_topic_iam_binding.
type PubsubTopicIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Topic: string, required
	Topic terra.StringValue `hcl:"topic,attr" validate:"required"`
	// Condition: optional
	Condition *pubsubtopiciambinding.Condition `hcl:"condition,block"`
}
type pubsubTopicIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_pubsub_topic_iam_binding.
func (ptib pubsubTopicIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ptib.ref.Append("etag"))
}

// Id returns a reference to field id of google_pubsub_topic_iam_binding.
func (ptib pubsubTopicIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ptib.ref.Append("id"))
}

// Members returns a reference to field members of google_pubsub_topic_iam_binding.
func (ptib pubsubTopicIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ptib.ref.Append("members"))
}

// Project returns a reference to field project of google_pubsub_topic_iam_binding.
func (ptib pubsubTopicIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ptib.ref.Append("project"))
}

// Role returns a reference to field role of google_pubsub_topic_iam_binding.
func (ptib pubsubTopicIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(ptib.ref.Append("role"))
}

// Topic returns a reference to field topic of google_pubsub_topic_iam_binding.
func (ptib pubsubTopicIamBindingAttributes) Topic() terra.StringValue {
	return terra.ReferenceAsString(ptib.ref.Append("topic"))
}

func (ptib pubsubTopicIamBindingAttributes) Condition() terra.ListValue[pubsubtopiciambinding.ConditionAttributes] {
	return terra.ReferenceAsList[pubsubtopiciambinding.ConditionAttributes](ptib.ref.Append("condition"))
}

type pubsubTopicIamBindingState struct {
	Etag      string                                 `json:"etag"`
	Id        string                                 `json:"id"`
	Members   []string                               `json:"members"`
	Project   string                                 `json:"project"`
	Role      string                                 `json:"role"`
	Topic     string                                 `json:"topic"`
	Condition []pubsubtopiciambinding.ConditionState `json:"condition"`
}
