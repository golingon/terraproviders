// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	pubsubtopiciammember "github.com/golingon/terraproviders/googlebeta/4.63.1/pubsubtopiciammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPubsubTopicIamMember creates a new instance of [PubsubTopicIamMember].
func NewPubsubTopicIamMember(name string, args PubsubTopicIamMemberArgs) *PubsubTopicIamMember {
	return &PubsubTopicIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PubsubTopicIamMember)(nil)

// PubsubTopicIamMember represents the Terraform resource google_pubsub_topic_iam_member.
type PubsubTopicIamMember struct {
	Name      string
	Args      PubsubTopicIamMemberArgs
	state     *pubsubTopicIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PubsubTopicIamMember].
func (ptim *PubsubTopicIamMember) Type() string {
	return "google_pubsub_topic_iam_member"
}

// LocalName returns the local name for [PubsubTopicIamMember].
func (ptim *PubsubTopicIamMember) LocalName() string {
	return ptim.Name
}

// Configuration returns the configuration (args) for [PubsubTopicIamMember].
func (ptim *PubsubTopicIamMember) Configuration() interface{} {
	return ptim.Args
}

// DependOn is used for other resources to depend on [PubsubTopicIamMember].
func (ptim *PubsubTopicIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(ptim)
}

// Dependencies returns the list of resources [PubsubTopicIamMember] depends_on.
func (ptim *PubsubTopicIamMember) Dependencies() terra.Dependencies {
	return ptim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PubsubTopicIamMember].
func (ptim *PubsubTopicIamMember) LifecycleManagement() *terra.Lifecycle {
	return ptim.Lifecycle
}

// Attributes returns the attributes for [PubsubTopicIamMember].
func (ptim *PubsubTopicIamMember) Attributes() pubsubTopicIamMemberAttributes {
	return pubsubTopicIamMemberAttributes{ref: terra.ReferenceResource(ptim)}
}

// ImportState imports the given attribute values into [PubsubTopicIamMember]'s state.
func (ptim *PubsubTopicIamMember) ImportState(av io.Reader) error {
	ptim.state = &pubsubTopicIamMemberState{}
	if err := json.NewDecoder(av).Decode(ptim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ptim.Type(), ptim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PubsubTopicIamMember] has state.
func (ptim *PubsubTopicIamMember) State() (*pubsubTopicIamMemberState, bool) {
	return ptim.state, ptim.state != nil
}

// StateMust returns the state for [PubsubTopicIamMember]. Panics if the state is nil.
func (ptim *PubsubTopicIamMember) StateMust() *pubsubTopicIamMemberState {
	if ptim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ptim.Type(), ptim.LocalName()))
	}
	return ptim.state
}

// PubsubTopicIamMemberArgs contains the configurations for google_pubsub_topic_iam_member.
type PubsubTopicIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Topic: string, required
	Topic terra.StringValue `hcl:"topic,attr" validate:"required"`
	// Condition: optional
	Condition *pubsubtopiciammember.Condition `hcl:"condition,block"`
}
type pubsubTopicIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_pubsub_topic_iam_member.
func (ptim pubsubTopicIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ptim.ref.Append("etag"))
}

// Id returns a reference to field id of google_pubsub_topic_iam_member.
func (ptim pubsubTopicIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ptim.ref.Append("id"))
}

// Member returns a reference to field member of google_pubsub_topic_iam_member.
func (ptim pubsubTopicIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(ptim.ref.Append("member"))
}

// Project returns a reference to field project of google_pubsub_topic_iam_member.
func (ptim pubsubTopicIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ptim.ref.Append("project"))
}

// Role returns a reference to field role of google_pubsub_topic_iam_member.
func (ptim pubsubTopicIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(ptim.ref.Append("role"))
}

// Topic returns a reference to field topic of google_pubsub_topic_iam_member.
func (ptim pubsubTopicIamMemberAttributes) Topic() terra.StringValue {
	return terra.ReferenceAsString(ptim.ref.Append("topic"))
}

func (ptim pubsubTopicIamMemberAttributes) Condition() terra.ListValue[pubsubtopiciammember.ConditionAttributes] {
	return terra.ReferenceAsList[pubsubtopiciammember.ConditionAttributes](ptim.ref.Append("condition"))
}

type pubsubTopicIamMemberState struct {
	Etag      string                                `json:"etag"`
	Id        string                                `json:"id"`
	Member    string                                `json:"member"`
	Project   string                                `json:"project"`
	Role      string                                `json:"role"`
	Topic     string                                `json:"topic"`
	Condition []pubsubtopiciammember.ConditionState `json:"condition"`
}
