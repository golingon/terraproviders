// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudwatchEventBus creates a new instance of [CloudwatchEventBus].
func NewCloudwatchEventBus(name string, args CloudwatchEventBusArgs) *CloudwatchEventBus {
	return &CloudwatchEventBus{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudwatchEventBus)(nil)

// CloudwatchEventBus represents the Terraform resource aws_cloudwatch_event_bus.
type CloudwatchEventBus struct {
	Name      string
	Args      CloudwatchEventBusArgs
	state     *cloudwatchEventBusState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudwatchEventBus].
func (ceb *CloudwatchEventBus) Type() string {
	return "aws_cloudwatch_event_bus"
}

// LocalName returns the local name for [CloudwatchEventBus].
func (ceb *CloudwatchEventBus) LocalName() string {
	return ceb.Name
}

// Configuration returns the configuration (args) for [CloudwatchEventBus].
func (ceb *CloudwatchEventBus) Configuration() interface{} {
	return ceb.Args
}

// DependOn is used for other resources to depend on [CloudwatchEventBus].
func (ceb *CloudwatchEventBus) DependOn() terra.Reference {
	return terra.ReferenceResource(ceb)
}

// Dependencies returns the list of resources [CloudwatchEventBus] depends_on.
func (ceb *CloudwatchEventBus) Dependencies() terra.Dependencies {
	return ceb.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudwatchEventBus].
func (ceb *CloudwatchEventBus) LifecycleManagement() *terra.Lifecycle {
	return ceb.Lifecycle
}

// Attributes returns the attributes for [CloudwatchEventBus].
func (ceb *CloudwatchEventBus) Attributes() cloudwatchEventBusAttributes {
	return cloudwatchEventBusAttributes{ref: terra.ReferenceResource(ceb)}
}

// ImportState imports the given attribute values into [CloudwatchEventBus]'s state.
func (ceb *CloudwatchEventBus) ImportState(av io.Reader) error {
	ceb.state = &cloudwatchEventBusState{}
	if err := json.NewDecoder(av).Decode(ceb.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ceb.Type(), ceb.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudwatchEventBus] has state.
func (ceb *CloudwatchEventBus) State() (*cloudwatchEventBusState, bool) {
	return ceb.state, ceb.state != nil
}

// StateMust returns the state for [CloudwatchEventBus]. Panics if the state is nil.
func (ceb *CloudwatchEventBus) StateMust() *cloudwatchEventBusState {
	if ceb.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ceb.Type(), ceb.LocalName()))
	}
	return ceb.state
}

// CloudwatchEventBusArgs contains the configurations for aws_cloudwatch_event_bus.
type CloudwatchEventBusArgs struct {
	// EventSourceName: string, optional
	EventSourceName terra.StringValue `hcl:"event_source_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}
type cloudwatchEventBusAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_cloudwatch_event_bus.
func (ceb cloudwatchEventBusAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ceb.ref.Append("arn"))
}

// EventSourceName returns a reference to field event_source_name of aws_cloudwatch_event_bus.
func (ceb cloudwatchEventBusAttributes) EventSourceName() terra.StringValue {
	return terra.ReferenceAsString(ceb.ref.Append("event_source_name"))
}

// Id returns a reference to field id of aws_cloudwatch_event_bus.
func (ceb cloudwatchEventBusAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ceb.ref.Append("id"))
}

// Name returns a reference to field name of aws_cloudwatch_event_bus.
func (ceb cloudwatchEventBusAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ceb.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_cloudwatch_event_bus.
func (ceb cloudwatchEventBusAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ceb.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_cloudwatch_event_bus.
func (ceb cloudwatchEventBusAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ceb.ref.Append("tags_all"))
}

type cloudwatchEventBusState struct {
	Arn             string            `json:"arn"`
	EventSourceName string            `json:"event_source_name"`
	Id              string            `json:"id"`
	Name            string            `json:"name"`
	Tags            map[string]string `json:"tags"`
	TagsAll         map[string]string `json:"tags_all"`
}
