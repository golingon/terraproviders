// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPinpointEventStream creates a new instance of [PinpointEventStream].
func NewPinpointEventStream(name string, args PinpointEventStreamArgs) *PinpointEventStream {
	return &PinpointEventStream{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PinpointEventStream)(nil)

// PinpointEventStream represents the Terraform resource aws_pinpoint_event_stream.
type PinpointEventStream struct {
	Name      string
	Args      PinpointEventStreamArgs
	state     *pinpointEventStreamState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PinpointEventStream].
func (pes *PinpointEventStream) Type() string {
	return "aws_pinpoint_event_stream"
}

// LocalName returns the local name for [PinpointEventStream].
func (pes *PinpointEventStream) LocalName() string {
	return pes.Name
}

// Configuration returns the configuration (args) for [PinpointEventStream].
func (pes *PinpointEventStream) Configuration() interface{} {
	return pes.Args
}

// DependOn is used for other resources to depend on [PinpointEventStream].
func (pes *PinpointEventStream) DependOn() terra.Reference {
	return terra.ReferenceResource(pes)
}

// Dependencies returns the list of resources [PinpointEventStream] depends_on.
func (pes *PinpointEventStream) Dependencies() terra.Dependencies {
	return pes.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PinpointEventStream].
func (pes *PinpointEventStream) LifecycleManagement() *terra.Lifecycle {
	return pes.Lifecycle
}

// Attributes returns the attributes for [PinpointEventStream].
func (pes *PinpointEventStream) Attributes() pinpointEventStreamAttributes {
	return pinpointEventStreamAttributes{ref: terra.ReferenceResource(pes)}
}

// ImportState imports the given attribute values into [PinpointEventStream]'s state.
func (pes *PinpointEventStream) ImportState(av io.Reader) error {
	pes.state = &pinpointEventStreamState{}
	if err := json.NewDecoder(av).Decode(pes.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pes.Type(), pes.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PinpointEventStream] has state.
func (pes *PinpointEventStream) State() (*pinpointEventStreamState, bool) {
	return pes.state, pes.state != nil
}

// StateMust returns the state for [PinpointEventStream]. Panics if the state is nil.
func (pes *PinpointEventStream) StateMust() *pinpointEventStreamState {
	if pes.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pes.Type(), pes.LocalName()))
	}
	return pes.state
}

// PinpointEventStreamArgs contains the configurations for aws_pinpoint_event_stream.
type PinpointEventStreamArgs struct {
	// ApplicationId: string, required
	ApplicationId terra.StringValue `hcl:"application_id,attr" validate:"required"`
	// DestinationStreamArn: string, required
	DestinationStreamArn terra.StringValue `hcl:"destination_stream_arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
}
type pinpointEventStreamAttributes struct {
	ref terra.Reference
}

// ApplicationId returns a reference to field application_id of aws_pinpoint_event_stream.
func (pes pinpointEventStreamAttributes) ApplicationId() terra.StringValue {
	return terra.ReferenceAsString(pes.ref.Append("application_id"))
}

// DestinationStreamArn returns a reference to field destination_stream_arn of aws_pinpoint_event_stream.
func (pes pinpointEventStreamAttributes) DestinationStreamArn() terra.StringValue {
	return terra.ReferenceAsString(pes.ref.Append("destination_stream_arn"))
}

// Id returns a reference to field id of aws_pinpoint_event_stream.
func (pes pinpointEventStreamAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pes.ref.Append("id"))
}

// RoleArn returns a reference to field role_arn of aws_pinpoint_event_stream.
func (pes pinpointEventStreamAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(pes.ref.Append("role_arn"))
}

type pinpointEventStreamState struct {
	ApplicationId        string `json:"application_id"`
	DestinationStreamArn string `json:"destination_stream_arn"`
	Id                   string `json:"id"`
	RoleArn              string `json:"role_arn"`
}
