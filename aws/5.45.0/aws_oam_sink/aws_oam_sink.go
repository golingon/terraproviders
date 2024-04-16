// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_oam_sink

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_oam_sink.
type Resource struct {
	Name      string
	Args      Args
	state     *awsOamSinkState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aos *Resource) Type() string {
	return "aws_oam_sink"
}

// LocalName returns the local name for [Resource].
func (aos *Resource) LocalName() string {
	return aos.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aos *Resource) Configuration() interface{} {
	return aos.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aos *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aos)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aos *Resource) Dependencies() terra.Dependencies {
	return aos.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aos *Resource) LifecycleManagement() *terra.Lifecycle {
	return aos.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aos *Resource) Attributes() awsOamSinkAttributes {
	return awsOamSinkAttributes{ref: terra.ReferenceResource(aos)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aos *Resource) ImportState(state io.Reader) error {
	aos.state = &awsOamSinkState{}
	if err := json.NewDecoder(state).Decode(aos.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aos.Type(), aos.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aos *Resource) State() (*awsOamSinkState, bool) {
	return aos.state, aos.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aos *Resource) StateMust() *awsOamSinkState {
	if aos.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aos.Type(), aos.LocalName()))
	}
	return aos.state
}

// Args contains the configurations for aws_oam_sink.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsOamSinkAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_oam_sink.
func (aos awsOamSinkAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(aos.ref.Append("arn"))
}

// Id returns a reference to field id of aws_oam_sink.
func (aos awsOamSinkAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aos.ref.Append("id"))
}

// Name returns a reference to field name of aws_oam_sink.
func (aos awsOamSinkAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aos.ref.Append("name"))
}

// SinkId returns a reference to field sink_id of aws_oam_sink.
func (aos awsOamSinkAttributes) SinkId() terra.StringValue {
	return terra.ReferenceAsString(aos.ref.Append("sink_id"))
}

// Tags returns a reference to field tags of aws_oam_sink.
func (aos awsOamSinkAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aos.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_oam_sink.
func (aos awsOamSinkAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aos.ref.Append("tags_all"))
}

func (aos awsOamSinkAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](aos.ref.Append("timeouts"))
}

type awsOamSinkState struct {
	Arn      string            `json:"arn"`
	Id       string            `json:"id"`
	Name     string            `json:"name"`
	SinkId   string            `json:"sink_id"`
	Tags     map[string]string `json:"tags"`
	TagsAll  map[string]string `json:"tags_all"`
	Timeouts *TimeoutsState    `json:"timeouts"`
}
