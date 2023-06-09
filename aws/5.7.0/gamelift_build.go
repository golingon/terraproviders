// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	gameliftbuild "github.com/golingon/terraproviders/aws/5.7.0/gameliftbuild"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGameliftBuild creates a new instance of [GameliftBuild].
func NewGameliftBuild(name string, args GameliftBuildArgs) *GameliftBuild {
	return &GameliftBuild{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GameliftBuild)(nil)

// GameliftBuild represents the Terraform resource aws_gamelift_build.
type GameliftBuild struct {
	Name      string
	Args      GameliftBuildArgs
	state     *gameliftBuildState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GameliftBuild].
func (gb *GameliftBuild) Type() string {
	return "aws_gamelift_build"
}

// LocalName returns the local name for [GameliftBuild].
func (gb *GameliftBuild) LocalName() string {
	return gb.Name
}

// Configuration returns the configuration (args) for [GameliftBuild].
func (gb *GameliftBuild) Configuration() interface{} {
	return gb.Args
}

// DependOn is used for other resources to depend on [GameliftBuild].
func (gb *GameliftBuild) DependOn() terra.Reference {
	return terra.ReferenceResource(gb)
}

// Dependencies returns the list of resources [GameliftBuild] depends_on.
func (gb *GameliftBuild) Dependencies() terra.Dependencies {
	return gb.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GameliftBuild].
func (gb *GameliftBuild) LifecycleManagement() *terra.Lifecycle {
	return gb.Lifecycle
}

// Attributes returns the attributes for [GameliftBuild].
func (gb *GameliftBuild) Attributes() gameliftBuildAttributes {
	return gameliftBuildAttributes{ref: terra.ReferenceResource(gb)}
}

// ImportState imports the given attribute values into [GameliftBuild]'s state.
func (gb *GameliftBuild) ImportState(av io.Reader) error {
	gb.state = &gameliftBuildState{}
	if err := json.NewDecoder(av).Decode(gb.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gb.Type(), gb.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GameliftBuild] has state.
func (gb *GameliftBuild) State() (*gameliftBuildState, bool) {
	return gb.state, gb.state != nil
}

// StateMust returns the state for [GameliftBuild]. Panics if the state is nil.
func (gb *GameliftBuild) StateMust() *gameliftBuildState {
	if gb.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gb.Type(), gb.LocalName()))
	}
	return gb.state
}

// GameliftBuildArgs contains the configurations for aws_gamelift_build.
type GameliftBuildArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// OperatingSystem: string, required
	OperatingSystem terra.StringValue `hcl:"operating_system,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Version: string, optional
	Version terra.StringValue `hcl:"version,attr"`
	// StorageLocation: required
	StorageLocation *gameliftbuild.StorageLocation `hcl:"storage_location,block" validate:"required"`
}
type gameliftBuildAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_gamelift_build.
func (gb gameliftBuildAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(gb.ref.Append("arn"))
}

// Id returns a reference to field id of aws_gamelift_build.
func (gb gameliftBuildAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gb.ref.Append("id"))
}

// Name returns a reference to field name of aws_gamelift_build.
func (gb gameliftBuildAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gb.ref.Append("name"))
}

// OperatingSystem returns a reference to field operating_system of aws_gamelift_build.
func (gb gameliftBuildAttributes) OperatingSystem() terra.StringValue {
	return terra.ReferenceAsString(gb.ref.Append("operating_system"))
}

// Tags returns a reference to field tags of aws_gamelift_build.
func (gb gameliftBuildAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gb.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_gamelift_build.
func (gb gameliftBuildAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gb.ref.Append("tags_all"))
}

// Version returns a reference to field version of aws_gamelift_build.
func (gb gameliftBuildAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(gb.ref.Append("version"))
}

func (gb gameliftBuildAttributes) StorageLocation() terra.ListValue[gameliftbuild.StorageLocationAttributes] {
	return terra.ReferenceAsList[gameliftbuild.StorageLocationAttributes](gb.ref.Append("storage_location"))
}

type gameliftBuildState struct {
	Arn             string                               `json:"arn"`
	Id              string                               `json:"id"`
	Name            string                               `json:"name"`
	OperatingSystem string                               `json:"operating_system"`
	Tags            map[string]string                    `json:"tags"`
	TagsAll         map[string]string                    `json:"tags_all"`
	Version         string                               `json:"version"`
	StorageLocation []gameliftbuild.StorageLocationState `json:"storage_location"`
}
