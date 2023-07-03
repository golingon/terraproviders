// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	gameservicesgameserverdeploymentrollout "github.com/golingon/terraproviders/google/4.71.0/gameservicesgameserverdeploymentrollout"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGameServicesGameServerDeploymentRollout creates a new instance of [GameServicesGameServerDeploymentRollout].
func NewGameServicesGameServerDeploymentRollout(name string, args GameServicesGameServerDeploymentRolloutArgs) *GameServicesGameServerDeploymentRollout {
	return &GameServicesGameServerDeploymentRollout{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GameServicesGameServerDeploymentRollout)(nil)

// GameServicesGameServerDeploymentRollout represents the Terraform resource google_game_services_game_server_deployment_rollout.
type GameServicesGameServerDeploymentRollout struct {
	Name      string
	Args      GameServicesGameServerDeploymentRolloutArgs
	state     *gameServicesGameServerDeploymentRolloutState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GameServicesGameServerDeploymentRollout].
func (gsgsdr *GameServicesGameServerDeploymentRollout) Type() string {
	return "google_game_services_game_server_deployment_rollout"
}

// LocalName returns the local name for [GameServicesGameServerDeploymentRollout].
func (gsgsdr *GameServicesGameServerDeploymentRollout) LocalName() string {
	return gsgsdr.Name
}

// Configuration returns the configuration (args) for [GameServicesGameServerDeploymentRollout].
func (gsgsdr *GameServicesGameServerDeploymentRollout) Configuration() interface{} {
	return gsgsdr.Args
}

// DependOn is used for other resources to depend on [GameServicesGameServerDeploymentRollout].
func (gsgsdr *GameServicesGameServerDeploymentRollout) DependOn() terra.Reference {
	return terra.ReferenceResource(gsgsdr)
}

// Dependencies returns the list of resources [GameServicesGameServerDeploymentRollout] depends_on.
func (gsgsdr *GameServicesGameServerDeploymentRollout) Dependencies() terra.Dependencies {
	return gsgsdr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GameServicesGameServerDeploymentRollout].
func (gsgsdr *GameServicesGameServerDeploymentRollout) LifecycleManagement() *terra.Lifecycle {
	return gsgsdr.Lifecycle
}

// Attributes returns the attributes for [GameServicesGameServerDeploymentRollout].
func (gsgsdr *GameServicesGameServerDeploymentRollout) Attributes() gameServicesGameServerDeploymentRolloutAttributes {
	return gameServicesGameServerDeploymentRolloutAttributes{ref: terra.ReferenceResource(gsgsdr)}
}

// ImportState imports the given attribute values into [GameServicesGameServerDeploymentRollout]'s state.
func (gsgsdr *GameServicesGameServerDeploymentRollout) ImportState(av io.Reader) error {
	gsgsdr.state = &gameServicesGameServerDeploymentRolloutState{}
	if err := json.NewDecoder(av).Decode(gsgsdr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gsgsdr.Type(), gsgsdr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GameServicesGameServerDeploymentRollout] has state.
func (gsgsdr *GameServicesGameServerDeploymentRollout) State() (*gameServicesGameServerDeploymentRolloutState, bool) {
	return gsgsdr.state, gsgsdr.state != nil
}

// StateMust returns the state for [GameServicesGameServerDeploymentRollout]. Panics if the state is nil.
func (gsgsdr *GameServicesGameServerDeploymentRollout) StateMust() *gameServicesGameServerDeploymentRolloutState {
	if gsgsdr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gsgsdr.Type(), gsgsdr.LocalName()))
	}
	return gsgsdr.state
}

// GameServicesGameServerDeploymentRolloutArgs contains the configurations for google_game_services_game_server_deployment_rollout.
type GameServicesGameServerDeploymentRolloutArgs struct {
	// DefaultGameServerConfig: string, required
	DefaultGameServerConfig terra.StringValue `hcl:"default_game_server_config,attr" validate:"required"`
	// DeploymentId: string, required
	DeploymentId terra.StringValue `hcl:"deployment_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// GameServerConfigOverrides: min=0
	GameServerConfigOverrides []gameservicesgameserverdeploymentrollout.GameServerConfigOverrides `hcl:"game_server_config_overrides,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *gameservicesgameserverdeploymentrollout.Timeouts `hcl:"timeouts,block"`
}
type gameServicesGameServerDeploymentRolloutAttributes struct {
	ref terra.Reference
}

// DefaultGameServerConfig returns a reference to field default_game_server_config of google_game_services_game_server_deployment_rollout.
func (gsgsdr gameServicesGameServerDeploymentRolloutAttributes) DefaultGameServerConfig() terra.StringValue {
	return terra.ReferenceAsString(gsgsdr.ref.Append("default_game_server_config"))
}

// DeploymentId returns a reference to field deployment_id of google_game_services_game_server_deployment_rollout.
func (gsgsdr gameServicesGameServerDeploymentRolloutAttributes) DeploymentId() terra.StringValue {
	return terra.ReferenceAsString(gsgsdr.ref.Append("deployment_id"))
}

// Id returns a reference to field id of google_game_services_game_server_deployment_rollout.
func (gsgsdr gameServicesGameServerDeploymentRolloutAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gsgsdr.ref.Append("id"))
}

// Name returns a reference to field name of google_game_services_game_server_deployment_rollout.
func (gsgsdr gameServicesGameServerDeploymentRolloutAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gsgsdr.ref.Append("name"))
}

// Project returns a reference to field project of google_game_services_game_server_deployment_rollout.
func (gsgsdr gameServicesGameServerDeploymentRolloutAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gsgsdr.ref.Append("project"))
}

func (gsgsdr gameServicesGameServerDeploymentRolloutAttributes) GameServerConfigOverrides() terra.ListValue[gameservicesgameserverdeploymentrollout.GameServerConfigOverridesAttributes] {
	return terra.ReferenceAsList[gameservicesgameserverdeploymentrollout.GameServerConfigOverridesAttributes](gsgsdr.ref.Append("game_server_config_overrides"))
}

func (gsgsdr gameServicesGameServerDeploymentRolloutAttributes) Timeouts() gameservicesgameserverdeploymentrollout.TimeoutsAttributes {
	return terra.ReferenceAsSingle[gameservicesgameserverdeploymentrollout.TimeoutsAttributes](gsgsdr.ref.Append("timeouts"))
}

type gameServicesGameServerDeploymentRolloutState struct {
	DefaultGameServerConfig   string                                                                   `json:"default_game_server_config"`
	DeploymentId              string                                                                   `json:"deployment_id"`
	Id                        string                                                                   `json:"id"`
	Name                      string                                                                   `json:"name"`
	Project                   string                                                                   `json:"project"`
	GameServerConfigOverrides []gameservicesgameserverdeploymentrollout.GameServerConfigOverridesState `json:"game_server_config_overrides"`
	Timeouts                  *gameservicesgameserverdeploymentrollout.TimeoutsState                   `json:"timeouts"`
}
