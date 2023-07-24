// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	gameservicesgameserverconfig "github.com/golingon/terraproviders/google/4.74.0/gameservicesgameserverconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGameServicesGameServerConfig creates a new instance of [GameServicesGameServerConfig].
func NewGameServicesGameServerConfig(name string, args GameServicesGameServerConfigArgs) *GameServicesGameServerConfig {
	return &GameServicesGameServerConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GameServicesGameServerConfig)(nil)

// GameServicesGameServerConfig represents the Terraform resource google_game_services_game_server_config.
type GameServicesGameServerConfig struct {
	Name      string
	Args      GameServicesGameServerConfigArgs
	state     *gameServicesGameServerConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GameServicesGameServerConfig].
func (gsgsc *GameServicesGameServerConfig) Type() string {
	return "google_game_services_game_server_config"
}

// LocalName returns the local name for [GameServicesGameServerConfig].
func (gsgsc *GameServicesGameServerConfig) LocalName() string {
	return gsgsc.Name
}

// Configuration returns the configuration (args) for [GameServicesGameServerConfig].
func (gsgsc *GameServicesGameServerConfig) Configuration() interface{} {
	return gsgsc.Args
}

// DependOn is used for other resources to depend on [GameServicesGameServerConfig].
func (gsgsc *GameServicesGameServerConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(gsgsc)
}

// Dependencies returns the list of resources [GameServicesGameServerConfig] depends_on.
func (gsgsc *GameServicesGameServerConfig) Dependencies() terra.Dependencies {
	return gsgsc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GameServicesGameServerConfig].
func (gsgsc *GameServicesGameServerConfig) LifecycleManagement() *terra.Lifecycle {
	return gsgsc.Lifecycle
}

// Attributes returns the attributes for [GameServicesGameServerConfig].
func (gsgsc *GameServicesGameServerConfig) Attributes() gameServicesGameServerConfigAttributes {
	return gameServicesGameServerConfigAttributes{ref: terra.ReferenceResource(gsgsc)}
}

// ImportState imports the given attribute values into [GameServicesGameServerConfig]'s state.
func (gsgsc *GameServicesGameServerConfig) ImportState(av io.Reader) error {
	gsgsc.state = &gameServicesGameServerConfigState{}
	if err := json.NewDecoder(av).Decode(gsgsc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gsgsc.Type(), gsgsc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GameServicesGameServerConfig] has state.
func (gsgsc *GameServicesGameServerConfig) State() (*gameServicesGameServerConfigState, bool) {
	return gsgsc.state, gsgsc.state != nil
}

// StateMust returns the state for [GameServicesGameServerConfig]. Panics if the state is nil.
func (gsgsc *GameServicesGameServerConfig) StateMust() *gameServicesGameServerConfigState {
	if gsgsc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gsgsc.Type(), gsgsc.LocalName()))
	}
	return gsgsc.state
}

// GameServicesGameServerConfigArgs contains the configurations for google_game_services_game_server_config.
type GameServicesGameServerConfigArgs struct {
	// ConfigId: string, required
	ConfigId terra.StringValue `hcl:"config_id,attr" validate:"required"`
	// DeploymentId: string, required
	DeploymentId terra.StringValue `hcl:"deployment_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// FleetConfigs: min=1
	FleetConfigs []gameservicesgameserverconfig.FleetConfigs `hcl:"fleet_configs,block" validate:"min=1"`
	// ScalingConfigs: min=0
	ScalingConfigs []gameservicesgameserverconfig.ScalingConfigs `hcl:"scaling_configs,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *gameservicesgameserverconfig.Timeouts `hcl:"timeouts,block"`
}
type gameServicesGameServerConfigAttributes struct {
	ref terra.Reference
}

// ConfigId returns a reference to field config_id of google_game_services_game_server_config.
func (gsgsc gameServicesGameServerConfigAttributes) ConfigId() terra.StringValue {
	return terra.ReferenceAsString(gsgsc.ref.Append("config_id"))
}

// DeploymentId returns a reference to field deployment_id of google_game_services_game_server_config.
func (gsgsc gameServicesGameServerConfigAttributes) DeploymentId() terra.StringValue {
	return terra.ReferenceAsString(gsgsc.ref.Append("deployment_id"))
}

// Description returns a reference to field description of google_game_services_game_server_config.
func (gsgsc gameServicesGameServerConfigAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gsgsc.ref.Append("description"))
}

// Id returns a reference to field id of google_game_services_game_server_config.
func (gsgsc gameServicesGameServerConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gsgsc.ref.Append("id"))
}

// Labels returns a reference to field labels of google_game_services_game_server_config.
func (gsgsc gameServicesGameServerConfigAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gsgsc.ref.Append("labels"))
}

// Location returns a reference to field location of google_game_services_game_server_config.
func (gsgsc gameServicesGameServerConfigAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gsgsc.ref.Append("location"))
}

// Name returns a reference to field name of google_game_services_game_server_config.
func (gsgsc gameServicesGameServerConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gsgsc.ref.Append("name"))
}

// Project returns a reference to field project of google_game_services_game_server_config.
func (gsgsc gameServicesGameServerConfigAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gsgsc.ref.Append("project"))
}

func (gsgsc gameServicesGameServerConfigAttributes) FleetConfigs() terra.ListValue[gameservicesgameserverconfig.FleetConfigsAttributes] {
	return terra.ReferenceAsList[gameservicesgameserverconfig.FleetConfigsAttributes](gsgsc.ref.Append("fleet_configs"))
}

func (gsgsc gameServicesGameServerConfigAttributes) ScalingConfigs() terra.ListValue[gameservicesgameserverconfig.ScalingConfigsAttributes] {
	return terra.ReferenceAsList[gameservicesgameserverconfig.ScalingConfigsAttributes](gsgsc.ref.Append("scaling_configs"))
}

func (gsgsc gameServicesGameServerConfigAttributes) Timeouts() gameservicesgameserverconfig.TimeoutsAttributes {
	return terra.ReferenceAsSingle[gameservicesgameserverconfig.TimeoutsAttributes](gsgsc.ref.Append("timeouts"))
}

type gameServicesGameServerConfigState struct {
	ConfigId       string                                             `json:"config_id"`
	DeploymentId   string                                             `json:"deployment_id"`
	Description    string                                             `json:"description"`
	Id             string                                             `json:"id"`
	Labels         map[string]string                                  `json:"labels"`
	Location       string                                             `json:"location"`
	Name           string                                             `json:"name"`
	Project        string                                             `json:"project"`
	FleetConfigs   []gameservicesgameserverconfig.FleetConfigsState   `json:"fleet_configs"`
	ScalingConfigs []gameservicesgameserverconfig.ScalingConfigsState `json:"scaling_configs"`
	Timeouts       *gameservicesgameserverconfig.TimeoutsState        `json:"timeouts"`
}
