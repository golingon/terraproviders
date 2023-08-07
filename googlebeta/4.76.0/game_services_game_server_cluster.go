// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	gameservicesgameservercluster "github.com/golingon/terraproviders/googlebeta/4.76.0/gameservicesgameservercluster"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGameServicesGameServerCluster creates a new instance of [GameServicesGameServerCluster].
func NewGameServicesGameServerCluster(name string, args GameServicesGameServerClusterArgs) *GameServicesGameServerCluster {
	return &GameServicesGameServerCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GameServicesGameServerCluster)(nil)

// GameServicesGameServerCluster represents the Terraform resource google_game_services_game_server_cluster.
type GameServicesGameServerCluster struct {
	Name      string
	Args      GameServicesGameServerClusterArgs
	state     *gameServicesGameServerClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GameServicesGameServerCluster].
func (gsgsc *GameServicesGameServerCluster) Type() string {
	return "google_game_services_game_server_cluster"
}

// LocalName returns the local name for [GameServicesGameServerCluster].
func (gsgsc *GameServicesGameServerCluster) LocalName() string {
	return gsgsc.Name
}

// Configuration returns the configuration (args) for [GameServicesGameServerCluster].
func (gsgsc *GameServicesGameServerCluster) Configuration() interface{} {
	return gsgsc.Args
}

// DependOn is used for other resources to depend on [GameServicesGameServerCluster].
func (gsgsc *GameServicesGameServerCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(gsgsc)
}

// Dependencies returns the list of resources [GameServicesGameServerCluster] depends_on.
func (gsgsc *GameServicesGameServerCluster) Dependencies() terra.Dependencies {
	return gsgsc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GameServicesGameServerCluster].
func (gsgsc *GameServicesGameServerCluster) LifecycleManagement() *terra.Lifecycle {
	return gsgsc.Lifecycle
}

// Attributes returns the attributes for [GameServicesGameServerCluster].
func (gsgsc *GameServicesGameServerCluster) Attributes() gameServicesGameServerClusterAttributes {
	return gameServicesGameServerClusterAttributes{ref: terra.ReferenceResource(gsgsc)}
}

// ImportState imports the given attribute values into [GameServicesGameServerCluster]'s state.
func (gsgsc *GameServicesGameServerCluster) ImportState(av io.Reader) error {
	gsgsc.state = &gameServicesGameServerClusterState{}
	if err := json.NewDecoder(av).Decode(gsgsc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gsgsc.Type(), gsgsc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GameServicesGameServerCluster] has state.
func (gsgsc *GameServicesGameServerCluster) State() (*gameServicesGameServerClusterState, bool) {
	return gsgsc.state, gsgsc.state != nil
}

// StateMust returns the state for [GameServicesGameServerCluster]. Panics if the state is nil.
func (gsgsc *GameServicesGameServerCluster) StateMust() *gameServicesGameServerClusterState {
	if gsgsc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gsgsc.Type(), gsgsc.LocalName()))
	}
	return gsgsc.state
}

// GameServicesGameServerClusterArgs contains the configurations for google_game_services_game_server_cluster.
type GameServicesGameServerClusterArgs struct {
	// ClusterId: string, required
	ClusterId terra.StringValue `hcl:"cluster_id,attr" validate:"required"`
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
	// RealmId: string, required
	RealmId terra.StringValue `hcl:"realm_id,attr" validate:"required"`
	// ConnectionInfo: required
	ConnectionInfo *gameservicesgameservercluster.ConnectionInfo `hcl:"connection_info,block" validate:"required"`
	// Timeouts: optional
	Timeouts *gameservicesgameservercluster.Timeouts `hcl:"timeouts,block"`
}
type gameServicesGameServerClusterAttributes struct {
	ref terra.Reference
}

// ClusterId returns a reference to field cluster_id of google_game_services_game_server_cluster.
func (gsgsc gameServicesGameServerClusterAttributes) ClusterId() terra.StringValue {
	return terra.ReferenceAsString(gsgsc.ref.Append("cluster_id"))
}

// Description returns a reference to field description of google_game_services_game_server_cluster.
func (gsgsc gameServicesGameServerClusterAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gsgsc.ref.Append("description"))
}

// Id returns a reference to field id of google_game_services_game_server_cluster.
func (gsgsc gameServicesGameServerClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gsgsc.ref.Append("id"))
}

// Labels returns a reference to field labels of google_game_services_game_server_cluster.
func (gsgsc gameServicesGameServerClusterAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gsgsc.ref.Append("labels"))
}

// Location returns a reference to field location of google_game_services_game_server_cluster.
func (gsgsc gameServicesGameServerClusterAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gsgsc.ref.Append("location"))
}

// Name returns a reference to field name of google_game_services_game_server_cluster.
func (gsgsc gameServicesGameServerClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gsgsc.ref.Append("name"))
}

// Project returns a reference to field project of google_game_services_game_server_cluster.
func (gsgsc gameServicesGameServerClusterAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gsgsc.ref.Append("project"))
}

// RealmId returns a reference to field realm_id of google_game_services_game_server_cluster.
func (gsgsc gameServicesGameServerClusterAttributes) RealmId() terra.StringValue {
	return terra.ReferenceAsString(gsgsc.ref.Append("realm_id"))
}

func (gsgsc gameServicesGameServerClusterAttributes) ConnectionInfo() terra.ListValue[gameservicesgameservercluster.ConnectionInfoAttributes] {
	return terra.ReferenceAsList[gameservicesgameservercluster.ConnectionInfoAttributes](gsgsc.ref.Append("connection_info"))
}

func (gsgsc gameServicesGameServerClusterAttributes) Timeouts() gameservicesgameservercluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[gameservicesgameservercluster.TimeoutsAttributes](gsgsc.ref.Append("timeouts"))
}

type gameServicesGameServerClusterState struct {
	ClusterId      string                                              `json:"cluster_id"`
	Description    string                                              `json:"description"`
	Id             string                                              `json:"id"`
	Labels         map[string]string                                   `json:"labels"`
	Location       string                                              `json:"location"`
	Name           string                                              `json:"name"`
	Project        string                                              `json:"project"`
	RealmId        string                                              `json:"realm_id"`
	ConnectionInfo []gameservicesgameservercluster.ConnectionInfoState `json:"connection_info"`
	Timeouts       *gameservicesgameservercluster.TimeoutsState        `json:"timeouts"`
}
