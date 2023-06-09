// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	medialivemultiplexprogram "github.com/golingon/terraproviders/aws/4.60.0/medialivemultiplexprogram"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMedialiveMultiplexProgram creates a new instance of [MedialiveMultiplexProgram].
func NewMedialiveMultiplexProgram(name string, args MedialiveMultiplexProgramArgs) *MedialiveMultiplexProgram {
	return &MedialiveMultiplexProgram{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MedialiveMultiplexProgram)(nil)

// MedialiveMultiplexProgram represents the Terraform resource aws_medialive_multiplex_program.
type MedialiveMultiplexProgram struct {
	Name      string
	Args      MedialiveMultiplexProgramArgs
	state     *medialiveMultiplexProgramState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MedialiveMultiplexProgram].
func (mmp *MedialiveMultiplexProgram) Type() string {
	return "aws_medialive_multiplex_program"
}

// LocalName returns the local name for [MedialiveMultiplexProgram].
func (mmp *MedialiveMultiplexProgram) LocalName() string {
	return mmp.Name
}

// Configuration returns the configuration (args) for [MedialiveMultiplexProgram].
func (mmp *MedialiveMultiplexProgram) Configuration() interface{} {
	return mmp.Args
}

// DependOn is used for other resources to depend on [MedialiveMultiplexProgram].
func (mmp *MedialiveMultiplexProgram) DependOn() terra.Reference {
	return terra.ReferenceResource(mmp)
}

// Dependencies returns the list of resources [MedialiveMultiplexProgram] depends_on.
func (mmp *MedialiveMultiplexProgram) Dependencies() terra.Dependencies {
	return mmp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MedialiveMultiplexProgram].
func (mmp *MedialiveMultiplexProgram) LifecycleManagement() *terra.Lifecycle {
	return mmp.Lifecycle
}

// Attributes returns the attributes for [MedialiveMultiplexProgram].
func (mmp *MedialiveMultiplexProgram) Attributes() medialiveMultiplexProgramAttributes {
	return medialiveMultiplexProgramAttributes{ref: terra.ReferenceResource(mmp)}
}

// ImportState imports the given attribute values into [MedialiveMultiplexProgram]'s state.
func (mmp *MedialiveMultiplexProgram) ImportState(av io.Reader) error {
	mmp.state = &medialiveMultiplexProgramState{}
	if err := json.NewDecoder(av).Decode(mmp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mmp.Type(), mmp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MedialiveMultiplexProgram] has state.
func (mmp *MedialiveMultiplexProgram) State() (*medialiveMultiplexProgramState, bool) {
	return mmp.state, mmp.state != nil
}

// StateMust returns the state for [MedialiveMultiplexProgram]. Panics if the state is nil.
func (mmp *MedialiveMultiplexProgram) StateMust() *medialiveMultiplexProgramState {
	if mmp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mmp.Type(), mmp.LocalName()))
	}
	return mmp.state
}

// MedialiveMultiplexProgramArgs contains the configurations for aws_medialive_multiplex_program.
type MedialiveMultiplexProgramArgs struct {
	// MultiplexId: string, required
	MultiplexId terra.StringValue `hcl:"multiplex_id,attr" validate:"required"`
	// ProgramName: string, required
	ProgramName terra.StringValue `hcl:"program_name,attr" validate:"required"`
	// MultiplexProgramSettings: min=0
	MultiplexProgramSettings []medialivemultiplexprogram.MultiplexProgramSettings `hcl:"multiplex_program_settings,block" validate:"min=0"`
}
type medialiveMultiplexProgramAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_medialive_multiplex_program.
func (mmp medialiveMultiplexProgramAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mmp.ref.Append("id"))
}

// MultiplexId returns a reference to field multiplex_id of aws_medialive_multiplex_program.
func (mmp medialiveMultiplexProgramAttributes) MultiplexId() terra.StringValue {
	return terra.ReferenceAsString(mmp.ref.Append("multiplex_id"))
}

// ProgramName returns a reference to field program_name of aws_medialive_multiplex_program.
func (mmp medialiveMultiplexProgramAttributes) ProgramName() terra.StringValue {
	return terra.ReferenceAsString(mmp.ref.Append("program_name"))
}

func (mmp medialiveMultiplexProgramAttributes) MultiplexProgramSettings() terra.ListValue[medialivemultiplexprogram.MultiplexProgramSettingsAttributes] {
	return terra.ReferenceAsList[medialivemultiplexprogram.MultiplexProgramSettingsAttributes](mmp.ref.Append("multiplex_program_settings"))
}

type medialiveMultiplexProgramState struct {
	Id                       string                                                    `json:"id"`
	MultiplexId              string                                                    `json:"multiplex_id"`
	ProgramName              string                                                    `json:"program_name"`
	MultiplexProgramSettings []medialivemultiplexprogram.MultiplexProgramSettingsState `json:"multiplex_program_settings"`
}
