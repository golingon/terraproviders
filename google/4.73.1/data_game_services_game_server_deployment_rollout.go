// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	datagameservicesgameserverdeploymentrollout "github.com/golingon/terraproviders/google/4.73.1/datagameservicesgameserverdeploymentrollout"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataGameServicesGameServerDeploymentRollout creates a new instance of [DataGameServicesGameServerDeploymentRollout].
func NewDataGameServicesGameServerDeploymentRollout(name string, args DataGameServicesGameServerDeploymentRolloutArgs) *DataGameServicesGameServerDeploymentRollout {
	return &DataGameServicesGameServerDeploymentRollout{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataGameServicesGameServerDeploymentRollout)(nil)

// DataGameServicesGameServerDeploymentRollout represents the Terraform data resource google_game_services_game_server_deployment_rollout.
type DataGameServicesGameServerDeploymentRollout struct {
	Name string
	Args DataGameServicesGameServerDeploymentRolloutArgs
}

// DataSource returns the Terraform object type for [DataGameServicesGameServerDeploymentRollout].
func (gsgsdr *DataGameServicesGameServerDeploymentRollout) DataSource() string {
	return "google_game_services_game_server_deployment_rollout"
}

// LocalName returns the local name for [DataGameServicesGameServerDeploymentRollout].
func (gsgsdr *DataGameServicesGameServerDeploymentRollout) LocalName() string {
	return gsgsdr.Name
}

// Configuration returns the configuration (args) for [DataGameServicesGameServerDeploymentRollout].
func (gsgsdr *DataGameServicesGameServerDeploymentRollout) Configuration() interface{} {
	return gsgsdr.Args
}

// Attributes returns the attributes for [DataGameServicesGameServerDeploymentRollout].
func (gsgsdr *DataGameServicesGameServerDeploymentRollout) Attributes() dataGameServicesGameServerDeploymentRolloutAttributes {
	return dataGameServicesGameServerDeploymentRolloutAttributes{ref: terra.ReferenceDataResource(gsgsdr)}
}

// DataGameServicesGameServerDeploymentRolloutArgs contains the configurations for google_game_services_game_server_deployment_rollout.
type DataGameServicesGameServerDeploymentRolloutArgs struct {
	// DeploymentId: string, required
	DeploymentId terra.StringValue `hcl:"deployment_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// GameServerConfigOverrides: min=0
	GameServerConfigOverrides []datagameservicesgameserverdeploymentrollout.GameServerConfigOverrides `hcl:"game_server_config_overrides,block" validate:"min=0"`
}
type dataGameServicesGameServerDeploymentRolloutAttributes struct {
	ref terra.Reference
}

// DefaultGameServerConfig returns a reference to field default_game_server_config of google_game_services_game_server_deployment_rollout.
func (gsgsdr dataGameServicesGameServerDeploymentRolloutAttributes) DefaultGameServerConfig() terra.StringValue {
	return terra.ReferenceAsString(gsgsdr.ref.Append("default_game_server_config"))
}

// DeploymentId returns a reference to field deployment_id of google_game_services_game_server_deployment_rollout.
func (gsgsdr dataGameServicesGameServerDeploymentRolloutAttributes) DeploymentId() terra.StringValue {
	return terra.ReferenceAsString(gsgsdr.ref.Append("deployment_id"))
}

// Id returns a reference to field id of google_game_services_game_server_deployment_rollout.
func (gsgsdr dataGameServicesGameServerDeploymentRolloutAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gsgsdr.ref.Append("id"))
}

// Name returns a reference to field name of google_game_services_game_server_deployment_rollout.
func (gsgsdr dataGameServicesGameServerDeploymentRolloutAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gsgsdr.ref.Append("name"))
}

// Project returns a reference to field project of google_game_services_game_server_deployment_rollout.
func (gsgsdr dataGameServicesGameServerDeploymentRolloutAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gsgsdr.ref.Append("project"))
}

func (gsgsdr dataGameServicesGameServerDeploymentRolloutAttributes) GameServerConfigOverrides() terra.ListValue[datagameservicesgameserverdeploymentrollout.GameServerConfigOverridesAttributes] {
	return terra.ReferenceAsList[datagameservicesgameserverdeploymentrollout.GameServerConfigOverridesAttributes](gsgsdr.ref.Append("game_server_config_overrides"))
}
