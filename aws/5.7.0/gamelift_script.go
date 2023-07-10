// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	gameliftscript "github.com/golingon/terraproviders/aws/5.7.0/gameliftscript"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGameliftScript creates a new instance of [GameliftScript].
func NewGameliftScript(name string, args GameliftScriptArgs) *GameliftScript {
	return &GameliftScript{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GameliftScript)(nil)

// GameliftScript represents the Terraform resource aws_gamelift_script.
type GameliftScript struct {
	Name      string
	Args      GameliftScriptArgs
	state     *gameliftScriptState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GameliftScript].
func (gs *GameliftScript) Type() string {
	return "aws_gamelift_script"
}

// LocalName returns the local name for [GameliftScript].
func (gs *GameliftScript) LocalName() string {
	return gs.Name
}

// Configuration returns the configuration (args) for [GameliftScript].
func (gs *GameliftScript) Configuration() interface{} {
	return gs.Args
}

// DependOn is used for other resources to depend on [GameliftScript].
func (gs *GameliftScript) DependOn() terra.Reference {
	return terra.ReferenceResource(gs)
}

// Dependencies returns the list of resources [GameliftScript] depends_on.
func (gs *GameliftScript) Dependencies() terra.Dependencies {
	return gs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GameliftScript].
func (gs *GameliftScript) LifecycleManagement() *terra.Lifecycle {
	return gs.Lifecycle
}

// Attributes returns the attributes for [GameliftScript].
func (gs *GameliftScript) Attributes() gameliftScriptAttributes {
	return gameliftScriptAttributes{ref: terra.ReferenceResource(gs)}
}

// ImportState imports the given attribute values into [GameliftScript]'s state.
func (gs *GameliftScript) ImportState(av io.Reader) error {
	gs.state = &gameliftScriptState{}
	if err := json.NewDecoder(av).Decode(gs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gs.Type(), gs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GameliftScript] has state.
func (gs *GameliftScript) State() (*gameliftScriptState, bool) {
	return gs.state, gs.state != nil
}

// StateMust returns the state for [GameliftScript]. Panics if the state is nil.
func (gs *GameliftScript) StateMust() *gameliftScriptState {
	if gs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gs.Type(), gs.LocalName()))
	}
	return gs.state
}

// GameliftScriptArgs contains the configurations for aws_gamelift_script.
type GameliftScriptArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Version: string, optional
	Version terra.StringValue `hcl:"version,attr"`
	// ZipFile: string, optional
	ZipFile terra.StringValue `hcl:"zip_file,attr"`
	// StorageLocation: optional
	StorageLocation *gameliftscript.StorageLocation `hcl:"storage_location,block"`
}
type gameliftScriptAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_gamelift_script.
func (gs gameliftScriptAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(gs.ref.Append("arn"))
}

// Id returns a reference to field id of aws_gamelift_script.
func (gs gameliftScriptAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gs.ref.Append("id"))
}

// Name returns a reference to field name of aws_gamelift_script.
func (gs gameliftScriptAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gs.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_gamelift_script.
func (gs gameliftScriptAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gs.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_gamelift_script.
func (gs gameliftScriptAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gs.ref.Append("tags_all"))
}

// Version returns a reference to field version of aws_gamelift_script.
func (gs gameliftScriptAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(gs.ref.Append("version"))
}

// ZipFile returns a reference to field zip_file of aws_gamelift_script.
func (gs gameliftScriptAttributes) ZipFile() terra.StringValue {
	return terra.ReferenceAsString(gs.ref.Append("zip_file"))
}

func (gs gameliftScriptAttributes) StorageLocation() terra.ListValue[gameliftscript.StorageLocationAttributes] {
	return terra.ReferenceAsList[gameliftscript.StorageLocationAttributes](gs.ref.Append("storage_location"))
}

type gameliftScriptState struct {
	Arn             string                                `json:"arn"`
	Id              string                                `json:"id"`
	Name            string                                `json:"name"`
	Tags            map[string]string                     `json:"tags"`
	TagsAll         map[string]string                     `json:"tags_all"`
	Version         string                                `json:"version"`
	ZipFile         string                                `json:"zip_file"`
	StorageLocation []gameliftscript.StorageLocationState `json:"storage_location"`
}
