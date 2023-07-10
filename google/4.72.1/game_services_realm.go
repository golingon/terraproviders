// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	gameservicesrealm "github.com/golingon/terraproviders/google/4.72.1/gameservicesrealm"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGameServicesRealm creates a new instance of [GameServicesRealm].
func NewGameServicesRealm(name string, args GameServicesRealmArgs) *GameServicesRealm {
	return &GameServicesRealm{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GameServicesRealm)(nil)

// GameServicesRealm represents the Terraform resource google_game_services_realm.
type GameServicesRealm struct {
	Name      string
	Args      GameServicesRealmArgs
	state     *gameServicesRealmState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GameServicesRealm].
func (gsr *GameServicesRealm) Type() string {
	return "google_game_services_realm"
}

// LocalName returns the local name for [GameServicesRealm].
func (gsr *GameServicesRealm) LocalName() string {
	return gsr.Name
}

// Configuration returns the configuration (args) for [GameServicesRealm].
func (gsr *GameServicesRealm) Configuration() interface{} {
	return gsr.Args
}

// DependOn is used for other resources to depend on [GameServicesRealm].
func (gsr *GameServicesRealm) DependOn() terra.Reference {
	return terra.ReferenceResource(gsr)
}

// Dependencies returns the list of resources [GameServicesRealm] depends_on.
func (gsr *GameServicesRealm) Dependencies() terra.Dependencies {
	return gsr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GameServicesRealm].
func (gsr *GameServicesRealm) LifecycleManagement() *terra.Lifecycle {
	return gsr.Lifecycle
}

// Attributes returns the attributes for [GameServicesRealm].
func (gsr *GameServicesRealm) Attributes() gameServicesRealmAttributes {
	return gameServicesRealmAttributes{ref: terra.ReferenceResource(gsr)}
}

// ImportState imports the given attribute values into [GameServicesRealm]'s state.
func (gsr *GameServicesRealm) ImportState(av io.Reader) error {
	gsr.state = &gameServicesRealmState{}
	if err := json.NewDecoder(av).Decode(gsr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gsr.Type(), gsr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GameServicesRealm] has state.
func (gsr *GameServicesRealm) State() (*gameServicesRealmState, bool) {
	return gsr.state, gsr.state != nil
}

// StateMust returns the state for [GameServicesRealm]. Panics if the state is nil.
func (gsr *GameServicesRealm) StateMust() *gameServicesRealmState {
	if gsr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gsr.Type(), gsr.LocalName()))
	}
	return gsr.state
}

// GameServicesRealmArgs contains the configurations for google_game_services_realm.
type GameServicesRealmArgs struct {
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
	// TimeZone: string, required
	TimeZone terra.StringValue `hcl:"time_zone,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *gameservicesrealm.Timeouts `hcl:"timeouts,block"`
}
type gameServicesRealmAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_game_services_realm.
func (gsr gameServicesRealmAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gsr.ref.Append("description"))
}

// Etag returns a reference to field etag of google_game_services_realm.
func (gsr gameServicesRealmAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gsr.ref.Append("etag"))
}

// Id returns a reference to field id of google_game_services_realm.
func (gsr gameServicesRealmAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gsr.ref.Append("id"))
}

// Labels returns a reference to field labels of google_game_services_realm.
func (gsr gameServicesRealmAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gsr.ref.Append("labels"))
}

// Location returns a reference to field location of google_game_services_realm.
func (gsr gameServicesRealmAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gsr.ref.Append("location"))
}

// Name returns a reference to field name of google_game_services_realm.
func (gsr gameServicesRealmAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gsr.ref.Append("name"))
}

// Project returns a reference to field project of google_game_services_realm.
func (gsr gameServicesRealmAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gsr.ref.Append("project"))
}

// RealmId returns a reference to field realm_id of google_game_services_realm.
func (gsr gameServicesRealmAttributes) RealmId() terra.StringValue {
	return terra.ReferenceAsString(gsr.ref.Append("realm_id"))
}

// TimeZone returns a reference to field time_zone of google_game_services_realm.
func (gsr gameServicesRealmAttributes) TimeZone() terra.StringValue {
	return terra.ReferenceAsString(gsr.ref.Append("time_zone"))
}

func (gsr gameServicesRealmAttributes) Timeouts() gameservicesrealm.TimeoutsAttributes {
	return terra.ReferenceAsSingle[gameservicesrealm.TimeoutsAttributes](gsr.ref.Append("timeouts"))
}

type gameServicesRealmState struct {
	Description string                           `json:"description"`
	Etag        string                           `json:"etag"`
	Id          string                           `json:"id"`
	Labels      map[string]string                `json:"labels"`
	Location    string                           `json:"location"`
	Name        string                           `json:"name"`
	Project     string                           `json:"project"`
	RealmId     string                           `json:"realm_id"`
	TimeZone    string                           `json:"time_zone"`
	Timeouts    *gameservicesrealm.TimeoutsState `json:"timeouts"`
}
