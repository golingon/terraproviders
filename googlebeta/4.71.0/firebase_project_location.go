// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	firebaseprojectlocation "github.com/golingon/terraproviders/googlebeta/4.71.0/firebaseprojectlocation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFirebaseProjectLocation creates a new instance of [FirebaseProjectLocation].
func NewFirebaseProjectLocation(name string, args FirebaseProjectLocationArgs) *FirebaseProjectLocation {
	return &FirebaseProjectLocation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FirebaseProjectLocation)(nil)

// FirebaseProjectLocation represents the Terraform resource google_firebase_project_location.
type FirebaseProjectLocation struct {
	Name      string
	Args      FirebaseProjectLocationArgs
	state     *firebaseProjectLocationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FirebaseProjectLocation].
func (fpl *FirebaseProjectLocation) Type() string {
	return "google_firebase_project_location"
}

// LocalName returns the local name for [FirebaseProjectLocation].
func (fpl *FirebaseProjectLocation) LocalName() string {
	return fpl.Name
}

// Configuration returns the configuration (args) for [FirebaseProjectLocation].
func (fpl *FirebaseProjectLocation) Configuration() interface{} {
	return fpl.Args
}

// DependOn is used for other resources to depend on [FirebaseProjectLocation].
func (fpl *FirebaseProjectLocation) DependOn() terra.Reference {
	return terra.ReferenceResource(fpl)
}

// Dependencies returns the list of resources [FirebaseProjectLocation] depends_on.
func (fpl *FirebaseProjectLocation) Dependencies() terra.Dependencies {
	return fpl.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FirebaseProjectLocation].
func (fpl *FirebaseProjectLocation) LifecycleManagement() *terra.Lifecycle {
	return fpl.Lifecycle
}

// Attributes returns the attributes for [FirebaseProjectLocation].
func (fpl *FirebaseProjectLocation) Attributes() firebaseProjectLocationAttributes {
	return firebaseProjectLocationAttributes{ref: terra.ReferenceResource(fpl)}
}

// ImportState imports the given attribute values into [FirebaseProjectLocation]'s state.
func (fpl *FirebaseProjectLocation) ImportState(av io.Reader) error {
	fpl.state = &firebaseProjectLocationState{}
	if err := json.NewDecoder(av).Decode(fpl.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fpl.Type(), fpl.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FirebaseProjectLocation] has state.
func (fpl *FirebaseProjectLocation) State() (*firebaseProjectLocationState, bool) {
	return fpl.state, fpl.state != nil
}

// StateMust returns the state for [FirebaseProjectLocation]. Panics if the state is nil.
func (fpl *FirebaseProjectLocation) StateMust() *firebaseProjectLocationState {
	if fpl.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fpl.Type(), fpl.LocalName()))
	}
	return fpl.state
}

// FirebaseProjectLocationArgs contains the configurations for google_firebase_project_location.
type FirebaseProjectLocationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LocationId: string, required
	LocationId terra.StringValue `hcl:"location_id,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *firebaseprojectlocation.Timeouts `hcl:"timeouts,block"`
}
type firebaseProjectLocationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_firebase_project_location.
func (fpl firebaseProjectLocationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fpl.ref.Append("id"))
}

// LocationId returns a reference to field location_id of google_firebase_project_location.
func (fpl firebaseProjectLocationAttributes) LocationId() terra.StringValue {
	return terra.ReferenceAsString(fpl.ref.Append("location_id"))
}

// Project returns a reference to field project of google_firebase_project_location.
func (fpl firebaseProjectLocationAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(fpl.ref.Append("project"))
}

func (fpl firebaseProjectLocationAttributes) Timeouts() firebaseprojectlocation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[firebaseprojectlocation.TimeoutsAttributes](fpl.ref.Append("timeouts"))
}

type firebaseProjectLocationState struct {
	Id         string                                 `json:"id"`
	LocationId string                                 `json:"location_id"`
	Project    string                                 `json:"project"`
	Timeouts   *firebaseprojectlocation.TimeoutsState `json:"timeouts"`
}
