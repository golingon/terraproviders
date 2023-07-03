// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	pubsublitetopic "github.com/golingon/terraproviders/googlebeta/4.71.0/pubsublitetopic"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPubsubLiteTopic creates a new instance of [PubsubLiteTopic].
func NewPubsubLiteTopic(name string, args PubsubLiteTopicArgs) *PubsubLiteTopic {
	return &PubsubLiteTopic{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PubsubLiteTopic)(nil)

// PubsubLiteTopic represents the Terraform resource google_pubsub_lite_topic.
type PubsubLiteTopic struct {
	Name      string
	Args      PubsubLiteTopicArgs
	state     *pubsubLiteTopicState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PubsubLiteTopic].
func (plt *PubsubLiteTopic) Type() string {
	return "google_pubsub_lite_topic"
}

// LocalName returns the local name for [PubsubLiteTopic].
func (plt *PubsubLiteTopic) LocalName() string {
	return plt.Name
}

// Configuration returns the configuration (args) for [PubsubLiteTopic].
func (plt *PubsubLiteTopic) Configuration() interface{} {
	return plt.Args
}

// DependOn is used for other resources to depend on [PubsubLiteTopic].
func (plt *PubsubLiteTopic) DependOn() terra.Reference {
	return terra.ReferenceResource(plt)
}

// Dependencies returns the list of resources [PubsubLiteTopic] depends_on.
func (plt *PubsubLiteTopic) Dependencies() terra.Dependencies {
	return plt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PubsubLiteTopic].
func (plt *PubsubLiteTopic) LifecycleManagement() *terra.Lifecycle {
	return plt.Lifecycle
}

// Attributes returns the attributes for [PubsubLiteTopic].
func (plt *PubsubLiteTopic) Attributes() pubsubLiteTopicAttributes {
	return pubsubLiteTopicAttributes{ref: terra.ReferenceResource(plt)}
}

// ImportState imports the given attribute values into [PubsubLiteTopic]'s state.
func (plt *PubsubLiteTopic) ImportState(av io.Reader) error {
	plt.state = &pubsubLiteTopicState{}
	if err := json.NewDecoder(av).Decode(plt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", plt.Type(), plt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PubsubLiteTopic] has state.
func (plt *PubsubLiteTopic) State() (*pubsubLiteTopicState, bool) {
	return plt.state, plt.state != nil
}

// StateMust returns the state for [PubsubLiteTopic]. Panics if the state is nil.
func (plt *PubsubLiteTopic) StateMust() *pubsubLiteTopicState {
	if plt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", plt.Type(), plt.LocalName()))
	}
	return plt.state
}

// PubsubLiteTopicArgs contains the configurations for google_pubsub_lite_topic.
type PubsubLiteTopicArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// PartitionConfig: optional
	PartitionConfig *pubsublitetopic.PartitionConfig `hcl:"partition_config,block"`
	// ReservationConfig: optional
	ReservationConfig *pubsublitetopic.ReservationConfig `hcl:"reservation_config,block"`
	// RetentionConfig: optional
	RetentionConfig *pubsublitetopic.RetentionConfig `hcl:"retention_config,block"`
	// Timeouts: optional
	Timeouts *pubsublitetopic.Timeouts `hcl:"timeouts,block"`
}
type pubsubLiteTopicAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_pubsub_lite_topic.
func (plt pubsubLiteTopicAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(plt.ref.Append("id"))
}

// Name returns a reference to field name of google_pubsub_lite_topic.
func (plt pubsubLiteTopicAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(plt.ref.Append("name"))
}

// Project returns a reference to field project of google_pubsub_lite_topic.
func (plt pubsubLiteTopicAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(plt.ref.Append("project"))
}

// Region returns a reference to field region of google_pubsub_lite_topic.
func (plt pubsubLiteTopicAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(plt.ref.Append("region"))
}

// Zone returns a reference to field zone of google_pubsub_lite_topic.
func (plt pubsubLiteTopicAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(plt.ref.Append("zone"))
}

func (plt pubsubLiteTopicAttributes) PartitionConfig() terra.ListValue[pubsublitetopic.PartitionConfigAttributes] {
	return terra.ReferenceAsList[pubsublitetopic.PartitionConfigAttributes](plt.ref.Append("partition_config"))
}

func (plt pubsubLiteTopicAttributes) ReservationConfig() terra.ListValue[pubsublitetopic.ReservationConfigAttributes] {
	return terra.ReferenceAsList[pubsublitetopic.ReservationConfigAttributes](plt.ref.Append("reservation_config"))
}

func (plt pubsubLiteTopicAttributes) RetentionConfig() terra.ListValue[pubsublitetopic.RetentionConfigAttributes] {
	return terra.ReferenceAsList[pubsublitetopic.RetentionConfigAttributes](plt.ref.Append("retention_config"))
}

func (plt pubsubLiteTopicAttributes) Timeouts() pubsublitetopic.TimeoutsAttributes {
	return terra.ReferenceAsSingle[pubsublitetopic.TimeoutsAttributes](plt.ref.Append("timeouts"))
}

type pubsubLiteTopicState struct {
	Id                string                                   `json:"id"`
	Name              string                                   `json:"name"`
	Project           string                                   `json:"project"`
	Region            string                                   `json:"region"`
	Zone              string                                   `json:"zone"`
	PartitionConfig   []pubsublitetopic.PartitionConfigState   `json:"partition_config"`
	ReservationConfig []pubsublitetopic.ReservationConfigState `json:"reservation_config"`
	RetentionConfig   []pubsublitetopic.RetentionConfigState   `json:"retention_config"`
	Timeouts          *pubsublitetopic.TimeoutsState           `json:"timeouts"`
}
