// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	gameliftgamesessionqueue "github.com/golingon/terraproviders/aws/5.7.0/gameliftgamesessionqueue"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGameliftGameSessionQueue creates a new instance of [GameliftGameSessionQueue].
func NewGameliftGameSessionQueue(name string, args GameliftGameSessionQueueArgs) *GameliftGameSessionQueue {
	return &GameliftGameSessionQueue{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GameliftGameSessionQueue)(nil)

// GameliftGameSessionQueue represents the Terraform resource aws_gamelift_game_session_queue.
type GameliftGameSessionQueue struct {
	Name      string
	Args      GameliftGameSessionQueueArgs
	state     *gameliftGameSessionQueueState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GameliftGameSessionQueue].
func (ggsq *GameliftGameSessionQueue) Type() string {
	return "aws_gamelift_game_session_queue"
}

// LocalName returns the local name for [GameliftGameSessionQueue].
func (ggsq *GameliftGameSessionQueue) LocalName() string {
	return ggsq.Name
}

// Configuration returns the configuration (args) for [GameliftGameSessionQueue].
func (ggsq *GameliftGameSessionQueue) Configuration() interface{} {
	return ggsq.Args
}

// DependOn is used for other resources to depend on [GameliftGameSessionQueue].
func (ggsq *GameliftGameSessionQueue) DependOn() terra.Reference {
	return terra.ReferenceResource(ggsq)
}

// Dependencies returns the list of resources [GameliftGameSessionQueue] depends_on.
func (ggsq *GameliftGameSessionQueue) Dependencies() terra.Dependencies {
	return ggsq.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GameliftGameSessionQueue].
func (ggsq *GameliftGameSessionQueue) LifecycleManagement() *terra.Lifecycle {
	return ggsq.Lifecycle
}

// Attributes returns the attributes for [GameliftGameSessionQueue].
func (ggsq *GameliftGameSessionQueue) Attributes() gameliftGameSessionQueueAttributes {
	return gameliftGameSessionQueueAttributes{ref: terra.ReferenceResource(ggsq)}
}

// ImportState imports the given attribute values into [GameliftGameSessionQueue]'s state.
func (ggsq *GameliftGameSessionQueue) ImportState(av io.Reader) error {
	ggsq.state = &gameliftGameSessionQueueState{}
	if err := json.NewDecoder(av).Decode(ggsq.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ggsq.Type(), ggsq.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GameliftGameSessionQueue] has state.
func (ggsq *GameliftGameSessionQueue) State() (*gameliftGameSessionQueueState, bool) {
	return ggsq.state, ggsq.state != nil
}

// StateMust returns the state for [GameliftGameSessionQueue]. Panics if the state is nil.
func (ggsq *GameliftGameSessionQueue) StateMust() *gameliftGameSessionQueueState {
	if ggsq.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ggsq.Type(), ggsq.LocalName()))
	}
	return ggsq.state
}

// GameliftGameSessionQueueArgs contains the configurations for aws_gamelift_game_session_queue.
type GameliftGameSessionQueueArgs struct {
	// CustomEventData: string, optional
	CustomEventData terra.StringValue `hcl:"custom_event_data,attr"`
	// Destinations: list of string, optional
	Destinations terra.ListValue[terra.StringValue] `hcl:"destinations,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NotificationTarget: string, optional
	NotificationTarget terra.StringValue `hcl:"notification_target,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TimeoutInSeconds: number, optional
	TimeoutInSeconds terra.NumberValue `hcl:"timeout_in_seconds,attr"`
	// PlayerLatencyPolicy: min=0
	PlayerLatencyPolicy []gameliftgamesessionqueue.PlayerLatencyPolicy `hcl:"player_latency_policy,block" validate:"min=0"`
}
type gameliftGameSessionQueueAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_gamelift_game_session_queue.
func (ggsq gameliftGameSessionQueueAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ggsq.ref.Append("arn"))
}

// CustomEventData returns a reference to field custom_event_data of aws_gamelift_game_session_queue.
func (ggsq gameliftGameSessionQueueAttributes) CustomEventData() terra.StringValue {
	return terra.ReferenceAsString(ggsq.ref.Append("custom_event_data"))
}

// Destinations returns a reference to field destinations of aws_gamelift_game_session_queue.
func (ggsq gameliftGameSessionQueueAttributes) Destinations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ggsq.ref.Append("destinations"))
}

// Id returns a reference to field id of aws_gamelift_game_session_queue.
func (ggsq gameliftGameSessionQueueAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ggsq.ref.Append("id"))
}

// Name returns a reference to field name of aws_gamelift_game_session_queue.
func (ggsq gameliftGameSessionQueueAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ggsq.ref.Append("name"))
}

// NotificationTarget returns a reference to field notification_target of aws_gamelift_game_session_queue.
func (ggsq gameliftGameSessionQueueAttributes) NotificationTarget() terra.StringValue {
	return terra.ReferenceAsString(ggsq.ref.Append("notification_target"))
}

// Tags returns a reference to field tags of aws_gamelift_game_session_queue.
func (ggsq gameliftGameSessionQueueAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ggsq.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_gamelift_game_session_queue.
func (ggsq gameliftGameSessionQueueAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ggsq.ref.Append("tags_all"))
}

// TimeoutInSeconds returns a reference to field timeout_in_seconds of aws_gamelift_game_session_queue.
func (ggsq gameliftGameSessionQueueAttributes) TimeoutInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(ggsq.ref.Append("timeout_in_seconds"))
}

func (ggsq gameliftGameSessionQueueAttributes) PlayerLatencyPolicy() terra.ListValue[gameliftgamesessionqueue.PlayerLatencyPolicyAttributes] {
	return terra.ReferenceAsList[gameliftgamesessionqueue.PlayerLatencyPolicyAttributes](ggsq.ref.Append("player_latency_policy"))
}

type gameliftGameSessionQueueState struct {
	Arn                 string                                              `json:"arn"`
	CustomEventData     string                                              `json:"custom_event_data"`
	Destinations        []string                                            `json:"destinations"`
	Id                  string                                              `json:"id"`
	Name                string                                              `json:"name"`
	NotificationTarget  string                                              `json:"notification_target"`
	Tags                map[string]string                                   `json:"tags"`
	TagsAll             map[string]string                                   `json:"tags_all"`
	TimeoutInSeconds    float64                                             `json:"timeout_in_seconds"`
	PlayerLatencyPolicy []gameliftgamesessionqueue.PlayerLatencyPolicyState `json:"player_latency_policy"`
}
