// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	firebasewebapp "github.com/golingon/terraproviders/googlebeta/4.74.0/firebasewebapp"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFirebaseWebApp creates a new instance of [FirebaseWebApp].
func NewFirebaseWebApp(name string, args FirebaseWebAppArgs) *FirebaseWebApp {
	return &FirebaseWebApp{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FirebaseWebApp)(nil)

// FirebaseWebApp represents the Terraform resource google_firebase_web_app.
type FirebaseWebApp struct {
	Name      string
	Args      FirebaseWebAppArgs
	state     *firebaseWebAppState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FirebaseWebApp].
func (fwa *FirebaseWebApp) Type() string {
	return "google_firebase_web_app"
}

// LocalName returns the local name for [FirebaseWebApp].
func (fwa *FirebaseWebApp) LocalName() string {
	return fwa.Name
}

// Configuration returns the configuration (args) for [FirebaseWebApp].
func (fwa *FirebaseWebApp) Configuration() interface{} {
	return fwa.Args
}

// DependOn is used for other resources to depend on [FirebaseWebApp].
func (fwa *FirebaseWebApp) DependOn() terra.Reference {
	return terra.ReferenceResource(fwa)
}

// Dependencies returns the list of resources [FirebaseWebApp] depends_on.
func (fwa *FirebaseWebApp) Dependencies() terra.Dependencies {
	return fwa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FirebaseWebApp].
func (fwa *FirebaseWebApp) LifecycleManagement() *terra.Lifecycle {
	return fwa.Lifecycle
}

// Attributes returns the attributes for [FirebaseWebApp].
func (fwa *FirebaseWebApp) Attributes() firebaseWebAppAttributes {
	return firebaseWebAppAttributes{ref: terra.ReferenceResource(fwa)}
}

// ImportState imports the given attribute values into [FirebaseWebApp]'s state.
func (fwa *FirebaseWebApp) ImportState(av io.Reader) error {
	fwa.state = &firebaseWebAppState{}
	if err := json.NewDecoder(av).Decode(fwa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fwa.Type(), fwa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FirebaseWebApp] has state.
func (fwa *FirebaseWebApp) State() (*firebaseWebAppState, bool) {
	return fwa.state, fwa.state != nil
}

// StateMust returns the state for [FirebaseWebApp]. Panics if the state is nil.
func (fwa *FirebaseWebApp) StateMust() *firebaseWebAppState {
	if fwa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fwa.Type(), fwa.LocalName()))
	}
	return fwa.state
}

// FirebaseWebAppArgs contains the configurations for google_firebase_web_app.
type FirebaseWebAppArgs struct {
	// DeletionPolicy: string, optional
	DeletionPolicy terra.StringValue `hcl:"deletion_policy,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *firebasewebapp.Timeouts `hcl:"timeouts,block"`
}
type firebaseWebAppAttributes struct {
	ref terra.Reference
}

// AppId returns a reference to field app_id of google_firebase_web_app.
func (fwa firebaseWebAppAttributes) AppId() terra.StringValue {
	return terra.ReferenceAsString(fwa.ref.Append("app_id"))
}

// AppUrls returns a reference to field app_urls of google_firebase_web_app.
func (fwa firebaseWebAppAttributes) AppUrls() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](fwa.ref.Append("app_urls"))
}

// DeletionPolicy returns a reference to field deletion_policy of google_firebase_web_app.
func (fwa firebaseWebAppAttributes) DeletionPolicy() terra.StringValue {
	return terra.ReferenceAsString(fwa.ref.Append("deletion_policy"))
}

// DisplayName returns a reference to field display_name of google_firebase_web_app.
func (fwa firebaseWebAppAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(fwa.ref.Append("display_name"))
}

// Id returns a reference to field id of google_firebase_web_app.
func (fwa firebaseWebAppAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fwa.ref.Append("id"))
}

// Name returns a reference to field name of google_firebase_web_app.
func (fwa firebaseWebAppAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(fwa.ref.Append("name"))
}

// Project returns a reference to field project of google_firebase_web_app.
func (fwa firebaseWebAppAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(fwa.ref.Append("project"))
}

func (fwa firebaseWebAppAttributes) Timeouts() firebasewebapp.TimeoutsAttributes {
	return terra.ReferenceAsSingle[firebasewebapp.TimeoutsAttributes](fwa.ref.Append("timeouts"))
}

type firebaseWebAppState struct {
	AppId          string                        `json:"app_id"`
	AppUrls        []string                      `json:"app_urls"`
	DeletionPolicy string                        `json:"deletion_policy"`
	DisplayName    string                        `json:"display_name"`
	Id             string                        `json:"id"`
	Name           string                        `json:"name"`
	Project        string                        `json:"project"`
	Timeouts       *firebasewebapp.TimeoutsState `json:"timeouts"`
}
