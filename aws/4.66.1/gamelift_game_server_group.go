// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	gameliftgameservergroup "github.com/golingon/terraproviders/aws/4.66.1/gameliftgameservergroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGameliftGameServerGroup creates a new instance of [GameliftGameServerGroup].
func NewGameliftGameServerGroup(name string, args GameliftGameServerGroupArgs) *GameliftGameServerGroup {
	return &GameliftGameServerGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GameliftGameServerGroup)(nil)

// GameliftGameServerGroup represents the Terraform resource aws_gamelift_game_server_group.
type GameliftGameServerGroup struct {
	Name      string
	Args      GameliftGameServerGroupArgs
	state     *gameliftGameServerGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GameliftGameServerGroup].
func (ggsg *GameliftGameServerGroup) Type() string {
	return "aws_gamelift_game_server_group"
}

// LocalName returns the local name for [GameliftGameServerGroup].
func (ggsg *GameliftGameServerGroup) LocalName() string {
	return ggsg.Name
}

// Configuration returns the configuration (args) for [GameliftGameServerGroup].
func (ggsg *GameliftGameServerGroup) Configuration() interface{} {
	return ggsg.Args
}

// DependOn is used for other resources to depend on [GameliftGameServerGroup].
func (ggsg *GameliftGameServerGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(ggsg)
}

// Dependencies returns the list of resources [GameliftGameServerGroup] depends_on.
func (ggsg *GameliftGameServerGroup) Dependencies() terra.Dependencies {
	return ggsg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GameliftGameServerGroup].
func (ggsg *GameliftGameServerGroup) LifecycleManagement() *terra.Lifecycle {
	return ggsg.Lifecycle
}

// Attributes returns the attributes for [GameliftGameServerGroup].
func (ggsg *GameliftGameServerGroup) Attributes() gameliftGameServerGroupAttributes {
	return gameliftGameServerGroupAttributes{ref: terra.ReferenceResource(ggsg)}
}

// ImportState imports the given attribute values into [GameliftGameServerGroup]'s state.
func (ggsg *GameliftGameServerGroup) ImportState(av io.Reader) error {
	ggsg.state = &gameliftGameServerGroupState{}
	if err := json.NewDecoder(av).Decode(ggsg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ggsg.Type(), ggsg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GameliftGameServerGroup] has state.
func (ggsg *GameliftGameServerGroup) State() (*gameliftGameServerGroupState, bool) {
	return ggsg.state, ggsg.state != nil
}

// StateMust returns the state for [GameliftGameServerGroup]. Panics if the state is nil.
func (ggsg *GameliftGameServerGroup) StateMust() *gameliftGameServerGroupState {
	if ggsg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ggsg.Type(), ggsg.LocalName()))
	}
	return ggsg.state
}

// GameliftGameServerGroupArgs contains the configurations for aws_gamelift_game_server_group.
type GameliftGameServerGroupArgs struct {
	// BalancingStrategy: string, optional
	BalancingStrategy terra.StringValue `hcl:"balancing_strategy,attr"`
	// GameServerGroupName: string, required
	GameServerGroupName terra.StringValue `hcl:"game_server_group_name,attr" validate:"required"`
	// GameServerProtectionPolicy: string, optional
	GameServerProtectionPolicy terra.StringValue `hcl:"game_server_protection_policy,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MaxSize: number, required
	MaxSize terra.NumberValue `hcl:"max_size,attr" validate:"required"`
	// MinSize: number, required
	MinSize terra.NumberValue `hcl:"min_size,attr" validate:"required"`
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// VpcSubnets: set of string, optional
	VpcSubnets terra.SetValue[terra.StringValue] `hcl:"vpc_subnets,attr"`
	// AutoScalingPolicy: optional
	AutoScalingPolicy *gameliftgameservergroup.AutoScalingPolicy `hcl:"auto_scaling_policy,block"`
	// InstanceDefinition: min=2,max=20
	InstanceDefinition []gameliftgameservergroup.InstanceDefinition `hcl:"instance_definition,block" validate:"min=2,max=20"`
	// LaunchTemplate: required
	LaunchTemplate *gameliftgameservergroup.LaunchTemplate `hcl:"launch_template,block" validate:"required"`
	// Timeouts: optional
	Timeouts *gameliftgameservergroup.Timeouts `hcl:"timeouts,block"`
}
type gameliftGameServerGroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_gamelift_game_server_group.
func (ggsg gameliftGameServerGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ggsg.ref.Append("arn"))
}

// AutoScalingGroupArn returns a reference to field auto_scaling_group_arn of aws_gamelift_game_server_group.
func (ggsg gameliftGameServerGroupAttributes) AutoScalingGroupArn() terra.StringValue {
	return terra.ReferenceAsString(ggsg.ref.Append("auto_scaling_group_arn"))
}

// BalancingStrategy returns a reference to field balancing_strategy of aws_gamelift_game_server_group.
func (ggsg gameliftGameServerGroupAttributes) BalancingStrategy() terra.StringValue {
	return terra.ReferenceAsString(ggsg.ref.Append("balancing_strategy"))
}

// GameServerGroupName returns a reference to field game_server_group_name of aws_gamelift_game_server_group.
func (ggsg gameliftGameServerGroupAttributes) GameServerGroupName() terra.StringValue {
	return terra.ReferenceAsString(ggsg.ref.Append("game_server_group_name"))
}

// GameServerProtectionPolicy returns a reference to field game_server_protection_policy of aws_gamelift_game_server_group.
func (ggsg gameliftGameServerGroupAttributes) GameServerProtectionPolicy() terra.StringValue {
	return terra.ReferenceAsString(ggsg.ref.Append("game_server_protection_policy"))
}

// Id returns a reference to field id of aws_gamelift_game_server_group.
func (ggsg gameliftGameServerGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ggsg.ref.Append("id"))
}

// MaxSize returns a reference to field max_size of aws_gamelift_game_server_group.
func (ggsg gameliftGameServerGroupAttributes) MaxSize() terra.NumberValue {
	return terra.ReferenceAsNumber(ggsg.ref.Append("max_size"))
}

// MinSize returns a reference to field min_size of aws_gamelift_game_server_group.
func (ggsg gameliftGameServerGroupAttributes) MinSize() terra.NumberValue {
	return terra.ReferenceAsNumber(ggsg.ref.Append("min_size"))
}

// RoleArn returns a reference to field role_arn of aws_gamelift_game_server_group.
func (ggsg gameliftGameServerGroupAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(ggsg.ref.Append("role_arn"))
}

// Tags returns a reference to field tags of aws_gamelift_game_server_group.
func (ggsg gameliftGameServerGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ggsg.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_gamelift_game_server_group.
func (ggsg gameliftGameServerGroupAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ggsg.ref.Append("tags_all"))
}

// VpcSubnets returns a reference to field vpc_subnets of aws_gamelift_game_server_group.
func (ggsg gameliftGameServerGroupAttributes) VpcSubnets() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ggsg.ref.Append("vpc_subnets"))
}

func (ggsg gameliftGameServerGroupAttributes) AutoScalingPolicy() terra.ListValue[gameliftgameservergroup.AutoScalingPolicyAttributes] {
	return terra.ReferenceAsList[gameliftgameservergroup.AutoScalingPolicyAttributes](ggsg.ref.Append("auto_scaling_policy"))
}

func (ggsg gameliftGameServerGroupAttributes) InstanceDefinition() terra.SetValue[gameliftgameservergroup.InstanceDefinitionAttributes] {
	return terra.ReferenceAsSet[gameliftgameservergroup.InstanceDefinitionAttributes](ggsg.ref.Append("instance_definition"))
}

func (ggsg gameliftGameServerGroupAttributes) LaunchTemplate() terra.ListValue[gameliftgameservergroup.LaunchTemplateAttributes] {
	return terra.ReferenceAsList[gameliftgameservergroup.LaunchTemplateAttributes](ggsg.ref.Append("launch_template"))
}

func (ggsg gameliftGameServerGroupAttributes) Timeouts() gameliftgameservergroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[gameliftgameservergroup.TimeoutsAttributes](ggsg.ref.Append("timeouts"))
}

type gameliftGameServerGroupState struct {
	Arn                        string                                            `json:"arn"`
	AutoScalingGroupArn        string                                            `json:"auto_scaling_group_arn"`
	BalancingStrategy          string                                            `json:"balancing_strategy"`
	GameServerGroupName        string                                            `json:"game_server_group_name"`
	GameServerProtectionPolicy string                                            `json:"game_server_protection_policy"`
	Id                         string                                            `json:"id"`
	MaxSize                    float64                                           `json:"max_size"`
	MinSize                    float64                                           `json:"min_size"`
	RoleArn                    string                                            `json:"role_arn"`
	Tags                       map[string]string                                 `json:"tags"`
	TagsAll                    map[string]string                                 `json:"tags_all"`
	VpcSubnets                 []string                                          `json:"vpc_subnets"`
	AutoScalingPolicy          []gameliftgameservergroup.AutoScalingPolicyState  `json:"auto_scaling_policy"`
	InstanceDefinition         []gameliftgameservergroup.InstanceDefinitionState `json:"instance_definition"`
	LaunchTemplate             []gameliftgameservergroup.LaunchTemplateState     `json:"launch_template"`
	Timeouts                   *gameliftgameservergroup.TimeoutsState            `json:"timeouts"`
}
