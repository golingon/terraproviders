// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	oamsink "github.com/golingon/terraproviders/aws/5.10.0/oamsink"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewOamSink creates a new instance of [OamSink].
func NewOamSink(name string, args OamSinkArgs) *OamSink {
	return &OamSink{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OamSink)(nil)

// OamSink represents the Terraform resource aws_oam_sink.
type OamSink struct {
	Name      string
	Args      OamSinkArgs
	state     *oamSinkState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [OamSink].
func (os *OamSink) Type() string {
	return "aws_oam_sink"
}

// LocalName returns the local name for [OamSink].
func (os *OamSink) LocalName() string {
	return os.Name
}

// Configuration returns the configuration (args) for [OamSink].
func (os *OamSink) Configuration() interface{} {
	return os.Args
}

// DependOn is used for other resources to depend on [OamSink].
func (os *OamSink) DependOn() terra.Reference {
	return terra.ReferenceResource(os)
}

// Dependencies returns the list of resources [OamSink] depends_on.
func (os *OamSink) Dependencies() terra.Dependencies {
	return os.DependsOn
}

// LifecycleManagement returns the lifecycle block for [OamSink].
func (os *OamSink) LifecycleManagement() *terra.Lifecycle {
	return os.Lifecycle
}

// Attributes returns the attributes for [OamSink].
func (os *OamSink) Attributes() oamSinkAttributes {
	return oamSinkAttributes{ref: terra.ReferenceResource(os)}
}

// ImportState imports the given attribute values into [OamSink]'s state.
func (os *OamSink) ImportState(av io.Reader) error {
	os.state = &oamSinkState{}
	if err := json.NewDecoder(av).Decode(os.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", os.Type(), os.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [OamSink] has state.
func (os *OamSink) State() (*oamSinkState, bool) {
	return os.state, os.state != nil
}

// StateMust returns the state for [OamSink]. Panics if the state is nil.
func (os *OamSink) StateMust() *oamSinkState {
	if os.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", os.Type(), os.LocalName()))
	}
	return os.state
}

// OamSinkArgs contains the configurations for aws_oam_sink.
type OamSinkArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Timeouts: optional
	Timeouts *oamsink.Timeouts `hcl:"timeouts,block"`
}
type oamSinkAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_oam_sink.
func (os oamSinkAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(os.ref.Append("arn"))
}

// Id returns a reference to field id of aws_oam_sink.
func (os oamSinkAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(os.ref.Append("id"))
}

// Name returns a reference to field name of aws_oam_sink.
func (os oamSinkAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(os.ref.Append("name"))
}

// SinkId returns a reference to field sink_id of aws_oam_sink.
func (os oamSinkAttributes) SinkId() terra.StringValue {
	return terra.ReferenceAsString(os.ref.Append("sink_id"))
}

// Tags returns a reference to field tags of aws_oam_sink.
func (os oamSinkAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](os.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_oam_sink.
func (os oamSinkAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](os.ref.Append("tags_all"))
}

func (os oamSinkAttributes) Timeouts() oamsink.TimeoutsAttributes {
	return terra.ReferenceAsSingle[oamsink.TimeoutsAttributes](os.ref.Append("timeouts"))
}

type oamSinkState struct {
	Arn      string                 `json:"arn"`
	Id       string                 `json:"id"`
	Name     string                 `json:"name"`
	SinkId   string                 `json:"sink_id"`
	Tags     map[string]string      `json:"tags"`
	TagsAll  map[string]string      `json:"tags_all"`
	Timeouts *oamsink.TimeoutsState `json:"timeouts"`
}
