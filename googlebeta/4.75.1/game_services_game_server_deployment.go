// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	gameservicesgameserverdeployment "github.com/golingon/terraproviders/googlebeta/4.75.1/gameservicesgameserverdeployment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGameServicesGameServerDeployment creates a new instance of [GameServicesGameServerDeployment].
func NewGameServicesGameServerDeployment(name string, args GameServicesGameServerDeploymentArgs) *GameServicesGameServerDeployment {
	return &GameServicesGameServerDeployment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GameServicesGameServerDeployment)(nil)

// GameServicesGameServerDeployment represents the Terraform resource google_game_services_game_server_deployment.
type GameServicesGameServerDeployment struct {
	Name      string
	Args      GameServicesGameServerDeploymentArgs
	state     *gameServicesGameServerDeploymentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GameServicesGameServerDeployment].
func (gsgsd *GameServicesGameServerDeployment) Type() string {
	return "google_game_services_game_server_deployment"
}

// LocalName returns the local name for [GameServicesGameServerDeployment].
func (gsgsd *GameServicesGameServerDeployment) LocalName() string {
	return gsgsd.Name
}

// Configuration returns the configuration (args) for [GameServicesGameServerDeployment].
func (gsgsd *GameServicesGameServerDeployment) Configuration() interface{} {
	return gsgsd.Args
}

// DependOn is used for other resources to depend on [GameServicesGameServerDeployment].
func (gsgsd *GameServicesGameServerDeployment) DependOn() terra.Reference {
	return terra.ReferenceResource(gsgsd)
}

// Dependencies returns the list of resources [GameServicesGameServerDeployment] depends_on.
func (gsgsd *GameServicesGameServerDeployment) Dependencies() terra.Dependencies {
	return gsgsd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GameServicesGameServerDeployment].
func (gsgsd *GameServicesGameServerDeployment) LifecycleManagement() *terra.Lifecycle {
	return gsgsd.Lifecycle
}

// Attributes returns the attributes for [GameServicesGameServerDeployment].
func (gsgsd *GameServicesGameServerDeployment) Attributes() gameServicesGameServerDeploymentAttributes {
	return gameServicesGameServerDeploymentAttributes{ref: terra.ReferenceResource(gsgsd)}
}

// ImportState imports the given attribute values into [GameServicesGameServerDeployment]'s state.
func (gsgsd *GameServicesGameServerDeployment) ImportState(av io.Reader) error {
	gsgsd.state = &gameServicesGameServerDeploymentState{}
	if err := json.NewDecoder(av).Decode(gsgsd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gsgsd.Type(), gsgsd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GameServicesGameServerDeployment] has state.
func (gsgsd *GameServicesGameServerDeployment) State() (*gameServicesGameServerDeploymentState, bool) {
	return gsgsd.state, gsgsd.state != nil
}

// StateMust returns the state for [GameServicesGameServerDeployment]. Panics if the state is nil.
func (gsgsd *GameServicesGameServerDeployment) StateMust() *gameServicesGameServerDeploymentState {
	if gsgsd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gsgsd.Type(), gsgsd.LocalName()))
	}
	return gsgsd.state
}

// GameServicesGameServerDeploymentArgs contains the configurations for google_game_services_game_server_deployment.
type GameServicesGameServerDeploymentArgs struct {
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
	// Timeouts: optional
	Timeouts *gameservicesgameserverdeployment.Timeouts `hcl:"timeouts,block"`
}
type gameServicesGameServerDeploymentAttributes struct {
	ref terra.Reference
}

// DeploymentId returns a reference to field deployment_id of google_game_services_game_server_deployment.
func (gsgsd gameServicesGameServerDeploymentAttributes) DeploymentId() terra.StringValue {
	return terra.ReferenceAsString(gsgsd.ref.Append("deployment_id"))
}

// Description returns a reference to field description of google_game_services_game_server_deployment.
func (gsgsd gameServicesGameServerDeploymentAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gsgsd.ref.Append("description"))
}

// Id returns a reference to field id of google_game_services_game_server_deployment.
func (gsgsd gameServicesGameServerDeploymentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gsgsd.ref.Append("id"))
}

// Labels returns a reference to field labels of google_game_services_game_server_deployment.
func (gsgsd gameServicesGameServerDeploymentAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gsgsd.ref.Append("labels"))
}

// Location returns a reference to field location of google_game_services_game_server_deployment.
func (gsgsd gameServicesGameServerDeploymentAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gsgsd.ref.Append("location"))
}

// Name returns a reference to field name of google_game_services_game_server_deployment.
func (gsgsd gameServicesGameServerDeploymentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gsgsd.ref.Append("name"))
}

// Project returns a reference to field project of google_game_services_game_server_deployment.
func (gsgsd gameServicesGameServerDeploymentAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gsgsd.ref.Append("project"))
}

func (gsgsd gameServicesGameServerDeploymentAttributes) Timeouts() gameservicesgameserverdeployment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[gameservicesgameserverdeployment.TimeoutsAttributes](gsgsd.ref.Append("timeouts"))
}

type gameServicesGameServerDeploymentState struct {
	DeploymentId string                                          `json:"deployment_id"`
	Description  string                                          `json:"description"`
	Id           string                                          `json:"id"`
	Labels       map[string]string                               `json:"labels"`
	Location     string                                          `json:"location"`
	Name         string                                          `json:"name"`
	Project      string                                          `json:"project"`
	Timeouts     *gameservicesgameserverdeployment.TimeoutsState `json:"timeouts"`
}
