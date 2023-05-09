// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ivschatroom "github.com/golingon/terraproviders/aws/4.66.1/ivschatroom"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIvschatRoom creates a new instance of [IvschatRoom].
func NewIvschatRoom(name string, args IvschatRoomArgs) *IvschatRoom {
	return &IvschatRoom{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IvschatRoom)(nil)

// IvschatRoom represents the Terraform resource aws_ivschat_room.
type IvschatRoom struct {
	Name      string
	Args      IvschatRoomArgs
	state     *ivschatRoomState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IvschatRoom].
func (ir *IvschatRoom) Type() string {
	return "aws_ivschat_room"
}

// LocalName returns the local name for [IvschatRoom].
func (ir *IvschatRoom) LocalName() string {
	return ir.Name
}

// Configuration returns the configuration (args) for [IvschatRoom].
func (ir *IvschatRoom) Configuration() interface{} {
	return ir.Args
}

// DependOn is used for other resources to depend on [IvschatRoom].
func (ir *IvschatRoom) DependOn() terra.Reference {
	return terra.ReferenceResource(ir)
}

// Dependencies returns the list of resources [IvschatRoom] depends_on.
func (ir *IvschatRoom) Dependencies() terra.Dependencies {
	return ir.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IvschatRoom].
func (ir *IvschatRoom) LifecycleManagement() *terra.Lifecycle {
	return ir.Lifecycle
}

// Attributes returns the attributes for [IvschatRoom].
func (ir *IvschatRoom) Attributes() ivschatRoomAttributes {
	return ivschatRoomAttributes{ref: terra.ReferenceResource(ir)}
}

// ImportState imports the given attribute values into [IvschatRoom]'s state.
func (ir *IvschatRoom) ImportState(av io.Reader) error {
	ir.state = &ivschatRoomState{}
	if err := json.NewDecoder(av).Decode(ir.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ir.Type(), ir.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IvschatRoom] has state.
func (ir *IvschatRoom) State() (*ivschatRoomState, bool) {
	return ir.state, ir.state != nil
}

// StateMust returns the state for [IvschatRoom]. Panics if the state is nil.
func (ir *IvschatRoom) StateMust() *ivschatRoomState {
	if ir.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ir.Type(), ir.LocalName()))
	}
	return ir.state
}

// IvschatRoomArgs contains the configurations for aws_ivschat_room.
type IvschatRoomArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LoggingConfigurationIdentifiers: list of string, optional
	LoggingConfigurationIdentifiers terra.ListValue[terra.StringValue] `hcl:"logging_configuration_identifiers,attr"`
	// MaximumMessageLength: number, optional
	MaximumMessageLength terra.NumberValue `hcl:"maximum_message_length,attr"`
	// MaximumMessageRatePerSecond: number, optional
	MaximumMessageRatePerSecond terra.NumberValue `hcl:"maximum_message_rate_per_second,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// MessageReviewHandler: optional
	MessageReviewHandler *ivschatroom.MessageReviewHandler `hcl:"message_review_handler,block"`
	// Timeouts: optional
	Timeouts *ivschatroom.Timeouts `hcl:"timeouts,block"`
}
type ivschatRoomAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ivschat_room.
func (ir ivschatRoomAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("arn"))
}

// Id returns a reference to field id of aws_ivschat_room.
func (ir ivschatRoomAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("id"))
}

// LoggingConfigurationIdentifiers returns a reference to field logging_configuration_identifiers of aws_ivschat_room.
func (ir ivschatRoomAttributes) LoggingConfigurationIdentifiers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ir.ref.Append("logging_configuration_identifiers"))
}

// MaximumMessageLength returns a reference to field maximum_message_length of aws_ivschat_room.
func (ir ivschatRoomAttributes) MaximumMessageLength() terra.NumberValue {
	return terra.ReferenceAsNumber(ir.ref.Append("maximum_message_length"))
}

// MaximumMessageRatePerSecond returns a reference to field maximum_message_rate_per_second of aws_ivschat_room.
func (ir ivschatRoomAttributes) MaximumMessageRatePerSecond() terra.NumberValue {
	return terra.ReferenceAsNumber(ir.ref.Append("maximum_message_rate_per_second"))
}

// Name returns a reference to field name of aws_ivschat_room.
func (ir ivschatRoomAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_ivschat_room.
func (ir ivschatRoomAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ir.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_ivschat_room.
func (ir ivschatRoomAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ir.ref.Append("tags_all"))
}

func (ir ivschatRoomAttributes) MessageReviewHandler() terra.ListValue[ivschatroom.MessageReviewHandlerAttributes] {
	return terra.ReferenceAsList[ivschatroom.MessageReviewHandlerAttributes](ir.ref.Append("message_review_handler"))
}

func (ir ivschatRoomAttributes) Timeouts() ivschatroom.TimeoutsAttributes {
	return terra.ReferenceAsSingle[ivschatroom.TimeoutsAttributes](ir.ref.Append("timeouts"))
}

type ivschatRoomState struct {
	Arn                             string                                  `json:"arn"`
	Id                              string                                  `json:"id"`
	LoggingConfigurationIdentifiers []string                                `json:"logging_configuration_identifiers"`
	MaximumMessageLength            float64                                 `json:"maximum_message_length"`
	MaximumMessageRatePerSecond     float64                                 `json:"maximum_message_rate_per_second"`
	Name                            string                                  `json:"name"`
	Tags                            map[string]string                       `json:"tags"`
	TagsAll                         map[string]string                       `json:"tags_all"`
	MessageReviewHandler            []ivschatroom.MessageReviewHandlerState `json:"message_review_handler"`
	Timeouts                        *ivschatroom.TimeoutsState              `json:"timeouts"`
}
