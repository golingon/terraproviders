// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	firebaseproject "github.com/golingon/terraproviders/googlebeta/4.64.0/firebaseproject"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFirebaseProject creates a new instance of [FirebaseProject].
func NewFirebaseProject(name string, args FirebaseProjectArgs) *FirebaseProject {
	return &FirebaseProject{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FirebaseProject)(nil)

// FirebaseProject represents the Terraform resource google_firebase_project.
type FirebaseProject struct {
	Name      string
	Args      FirebaseProjectArgs
	state     *firebaseProjectState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FirebaseProject].
func (fp *FirebaseProject) Type() string {
	return "google_firebase_project"
}

// LocalName returns the local name for [FirebaseProject].
func (fp *FirebaseProject) LocalName() string {
	return fp.Name
}

// Configuration returns the configuration (args) for [FirebaseProject].
func (fp *FirebaseProject) Configuration() interface{} {
	return fp.Args
}

// DependOn is used for other resources to depend on [FirebaseProject].
func (fp *FirebaseProject) DependOn() terra.Reference {
	return terra.ReferenceResource(fp)
}

// Dependencies returns the list of resources [FirebaseProject] depends_on.
func (fp *FirebaseProject) Dependencies() terra.Dependencies {
	return fp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FirebaseProject].
func (fp *FirebaseProject) LifecycleManagement() *terra.Lifecycle {
	return fp.Lifecycle
}

// Attributes returns the attributes for [FirebaseProject].
func (fp *FirebaseProject) Attributes() firebaseProjectAttributes {
	return firebaseProjectAttributes{ref: terra.ReferenceResource(fp)}
}

// ImportState imports the given attribute values into [FirebaseProject]'s state.
func (fp *FirebaseProject) ImportState(av io.Reader) error {
	fp.state = &firebaseProjectState{}
	if err := json.NewDecoder(av).Decode(fp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fp.Type(), fp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FirebaseProject] has state.
func (fp *FirebaseProject) State() (*firebaseProjectState, bool) {
	return fp.state, fp.state != nil
}

// StateMust returns the state for [FirebaseProject]. Panics if the state is nil.
func (fp *FirebaseProject) StateMust() *firebaseProjectState {
	if fp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fp.Type(), fp.LocalName()))
	}
	return fp.state
}

// FirebaseProjectArgs contains the configurations for google_firebase_project.
type FirebaseProjectArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *firebaseproject.Timeouts `hcl:"timeouts,block"`
}
type firebaseProjectAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of google_firebase_project.
func (fp firebaseProjectAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(fp.ref.Append("display_name"))
}

// Id returns a reference to field id of google_firebase_project.
func (fp firebaseProjectAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fp.ref.Append("id"))
}

// Project returns a reference to field project of google_firebase_project.
func (fp firebaseProjectAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(fp.ref.Append("project"))
}

// ProjectNumber returns a reference to field project_number of google_firebase_project.
func (fp firebaseProjectAttributes) ProjectNumber() terra.StringValue {
	return terra.ReferenceAsString(fp.ref.Append("project_number"))
}

func (fp firebaseProjectAttributes) Timeouts() firebaseproject.TimeoutsAttributes {
	return terra.ReferenceAsSingle[firebaseproject.TimeoutsAttributes](fp.ref.Append("timeouts"))
}

type firebaseProjectState struct {
	DisplayName   string                         `json:"display_name"`
	Id            string                         `json:"id"`
	Project       string                         `json:"project"`
	ProjectNumber string                         `json:"project_number"`
	Timeouts      *firebaseproject.TimeoutsState `json:"timeouts"`
}
