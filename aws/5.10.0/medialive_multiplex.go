// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	medialivemultiplex "github.com/golingon/terraproviders/aws/5.10.0/medialivemultiplex"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMedialiveMultiplex creates a new instance of [MedialiveMultiplex].
func NewMedialiveMultiplex(name string, args MedialiveMultiplexArgs) *MedialiveMultiplex {
	return &MedialiveMultiplex{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MedialiveMultiplex)(nil)

// MedialiveMultiplex represents the Terraform resource aws_medialive_multiplex.
type MedialiveMultiplex struct {
	Name      string
	Args      MedialiveMultiplexArgs
	state     *medialiveMultiplexState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MedialiveMultiplex].
func (mm *MedialiveMultiplex) Type() string {
	return "aws_medialive_multiplex"
}

// LocalName returns the local name for [MedialiveMultiplex].
func (mm *MedialiveMultiplex) LocalName() string {
	return mm.Name
}

// Configuration returns the configuration (args) for [MedialiveMultiplex].
func (mm *MedialiveMultiplex) Configuration() interface{} {
	return mm.Args
}

// DependOn is used for other resources to depend on [MedialiveMultiplex].
func (mm *MedialiveMultiplex) DependOn() terra.Reference {
	return terra.ReferenceResource(mm)
}

// Dependencies returns the list of resources [MedialiveMultiplex] depends_on.
func (mm *MedialiveMultiplex) Dependencies() terra.Dependencies {
	return mm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MedialiveMultiplex].
func (mm *MedialiveMultiplex) LifecycleManagement() *terra.Lifecycle {
	return mm.Lifecycle
}

// Attributes returns the attributes for [MedialiveMultiplex].
func (mm *MedialiveMultiplex) Attributes() medialiveMultiplexAttributes {
	return medialiveMultiplexAttributes{ref: terra.ReferenceResource(mm)}
}

// ImportState imports the given attribute values into [MedialiveMultiplex]'s state.
func (mm *MedialiveMultiplex) ImportState(av io.Reader) error {
	mm.state = &medialiveMultiplexState{}
	if err := json.NewDecoder(av).Decode(mm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mm.Type(), mm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MedialiveMultiplex] has state.
func (mm *MedialiveMultiplex) State() (*medialiveMultiplexState, bool) {
	return mm.state, mm.state != nil
}

// StateMust returns the state for [MedialiveMultiplex]. Panics if the state is nil.
func (mm *MedialiveMultiplex) StateMust() *medialiveMultiplexState {
	if mm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mm.Type(), mm.LocalName()))
	}
	return mm.state
}

// MedialiveMultiplexArgs contains the configurations for aws_medialive_multiplex.
type MedialiveMultiplexArgs struct {
	// AvailabilityZones: list of string, required
	AvailabilityZones terra.ListValue[terra.StringValue] `hcl:"availability_zones,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// StartMultiplex: bool, optional
	StartMultiplex terra.BoolValue `hcl:"start_multiplex,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// MultiplexSettings: optional
	MultiplexSettings *medialivemultiplex.MultiplexSettings `hcl:"multiplex_settings,block"`
	// Timeouts: optional
	Timeouts *medialivemultiplex.Timeouts `hcl:"timeouts,block"`
}
type medialiveMultiplexAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_medialive_multiplex.
func (mm medialiveMultiplexAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(mm.ref.Append("arn"))
}

// AvailabilityZones returns a reference to field availability_zones of aws_medialive_multiplex.
func (mm medialiveMultiplexAttributes) AvailabilityZones() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](mm.ref.Append("availability_zones"))
}

// Id returns a reference to field id of aws_medialive_multiplex.
func (mm medialiveMultiplexAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mm.ref.Append("id"))
}

// Name returns a reference to field name of aws_medialive_multiplex.
func (mm medialiveMultiplexAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mm.ref.Append("name"))
}

// StartMultiplex returns a reference to field start_multiplex of aws_medialive_multiplex.
func (mm medialiveMultiplexAttributes) StartMultiplex() terra.BoolValue {
	return terra.ReferenceAsBool(mm.ref.Append("start_multiplex"))
}

// Tags returns a reference to field tags of aws_medialive_multiplex.
func (mm medialiveMultiplexAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mm.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_medialive_multiplex.
func (mm medialiveMultiplexAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mm.ref.Append("tags_all"))
}

func (mm medialiveMultiplexAttributes) MultiplexSettings() terra.ListValue[medialivemultiplex.MultiplexSettingsAttributes] {
	return terra.ReferenceAsList[medialivemultiplex.MultiplexSettingsAttributes](mm.ref.Append("multiplex_settings"))
}

func (mm medialiveMultiplexAttributes) Timeouts() medialivemultiplex.TimeoutsAttributes {
	return terra.ReferenceAsSingle[medialivemultiplex.TimeoutsAttributes](mm.ref.Append("timeouts"))
}

type medialiveMultiplexState struct {
	Arn               string                                      `json:"arn"`
	AvailabilityZones []string                                    `json:"availability_zones"`
	Id                string                                      `json:"id"`
	Name              string                                      `json:"name"`
	StartMultiplex    bool                                        `json:"start_multiplex"`
	Tags              map[string]string                           `json:"tags"`
	TagsAll           map[string]string                           `json:"tags_all"`
	MultiplexSettings []medialivemultiplex.MultiplexSettingsState `json:"multiplex_settings"`
	Timeouts          *medialivemultiplex.TimeoutsState           `json:"timeouts"`
}
