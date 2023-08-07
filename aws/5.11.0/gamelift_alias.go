// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	gameliftalias "github.com/golingon/terraproviders/aws/5.11.0/gameliftalias"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGameliftAlias creates a new instance of [GameliftAlias].
func NewGameliftAlias(name string, args GameliftAliasArgs) *GameliftAlias {
	return &GameliftAlias{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GameliftAlias)(nil)

// GameliftAlias represents the Terraform resource aws_gamelift_alias.
type GameliftAlias struct {
	Name      string
	Args      GameliftAliasArgs
	state     *gameliftAliasState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GameliftAlias].
func (ga *GameliftAlias) Type() string {
	return "aws_gamelift_alias"
}

// LocalName returns the local name for [GameliftAlias].
func (ga *GameliftAlias) LocalName() string {
	return ga.Name
}

// Configuration returns the configuration (args) for [GameliftAlias].
func (ga *GameliftAlias) Configuration() interface{} {
	return ga.Args
}

// DependOn is used for other resources to depend on [GameliftAlias].
func (ga *GameliftAlias) DependOn() terra.Reference {
	return terra.ReferenceResource(ga)
}

// Dependencies returns the list of resources [GameliftAlias] depends_on.
func (ga *GameliftAlias) Dependencies() terra.Dependencies {
	return ga.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GameliftAlias].
func (ga *GameliftAlias) LifecycleManagement() *terra.Lifecycle {
	return ga.Lifecycle
}

// Attributes returns the attributes for [GameliftAlias].
func (ga *GameliftAlias) Attributes() gameliftAliasAttributes {
	return gameliftAliasAttributes{ref: terra.ReferenceResource(ga)}
}

// ImportState imports the given attribute values into [GameliftAlias]'s state.
func (ga *GameliftAlias) ImportState(av io.Reader) error {
	ga.state = &gameliftAliasState{}
	if err := json.NewDecoder(av).Decode(ga.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ga.Type(), ga.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GameliftAlias] has state.
func (ga *GameliftAlias) State() (*gameliftAliasState, bool) {
	return ga.state, ga.state != nil
}

// StateMust returns the state for [GameliftAlias]. Panics if the state is nil.
func (ga *GameliftAlias) StateMust() *gameliftAliasState {
	if ga.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ga.Type(), ga.LocalName()))
	}
	return ga.state
}

// GameliftAliasArgs contains the configurations for aws_gamelift_alias.
type GameliftAliasArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// RoutingStrategy: required
	RoutingStrategy *gameliftalias.RoutingStrategy `hcl:"routing_strategy,block" validate:"required"`
}
type gameliftAliasAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_gamelift_alias.
func (ga gameliftAliasAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ga.ref.Append("arn"))
}

// Description returns a reference to field description of aws_gamelift_alias.
func (ga gameliftAliasAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ga.ref.Append("description"))
}

// Id returns a reference to field id of aws_gamelift_alias.
func (ga gameliftAliasAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ga.ref.Append("id"))
}

// Name returns a reference to field name of aws_gamelift_alias.
func (ga gameliftAliasAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ga.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_gamelift_alias.
func (ga gameliftAliasAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ga.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_gamelift_alias.
func (ga gameliftAliasAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ga.ref.Append("tags_all"))
}

func (ga gameliftAliasAttributes) RoutingStrategy() terra.ListValue[gameliftalias.RoutingStrategyAttributes] {
	return terra.ReferenceAsList[gameliftalias.RoutingStrategyAttributes](ga.ref.Append("routing_strategy"))
}

type gameliftAliasState struct {
	Arn             string                               `json:"arn"`
	Description     string                               `json:"description"`
	Id              string                               `json:"id"`
	Name            string                               `json:"name"`
	Tags            map[string]string                    `json:"tags"`
	TagsAll         map[string]string                    `json:"tags_all"`
	RoutingStrategy []gameliftalias.RoutingStrategyState `json:"routing_strategy"`
}
