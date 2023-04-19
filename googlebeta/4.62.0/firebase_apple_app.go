// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	firebaseappleapp "github.com/golingon/terraproviders/googlebeta/4.62.0/firebaseappleapp"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFirebaseAppleApp creates a new instance of [FirebaseAppleApp].
func NewFirebaseAppleApp(name string, args FirebaseAppleAppArgs) *FirebaseAppleApp {
	return &FirebaseAppleApp{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FirebaseAppleApp)(nil)

// FirebaseAppleApp represents the Terraform resource google_firebase_apple_app.
type FirebaseAppleApp struct {
	Name      string
	Args      FirebaseAppleAppArgs
	state     *firebaseAppleAppState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FirebaseAppleApp].
func (faa *FirebaseAppleApp) Type() string {
	return "google_firebase_apple_app"
}

// LocalName returns the local name for [FirebaseAppleApp].
func (faa *FirebaseAppleApp) LocalName() string {
	return faa.Name
}

// Configuration returns the configuration (args) for [FirebaseAppleApp].
func (faa *FirebaseAppleApp) Configuration() interface{} {
	return faa.Args
}

// DependOn is used for other resources to depend on [FirebaseAppleApp].
func (faa *FirebaseAppleApp) DependOn() terra.Reference {
	return terra.ReferenceResource(faa)
}

// Dependencies returns the list of resources [FirebaseAppleApp] depends_on.
func (faa *FirebaseAppleApp) Dependencies() terra.Dependencies {
	return faa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FirebaseAppleApp].
func (faa *FirebaseAppleApp) LifecycleManagement() *terra.Lifecycle {
	return faa.Lifecycle
}

// Attributes returns the attributes for [FirebaseAppleApp].
func (faa *FirebaseAppleApp) Attributes() firebaseAppleAppAttributes {
	return firebaseAppleAppAttributes{ref: terra.ReferenceResource(faa)}
}

// ImportState imports the given attribute values into [FirebaseAppleApp]'s state.
func (faa *FirebaseAppleApp) ImportState(av io.Reader) error {
	faa.state = &firebaseAppleAppState{}
	if err := json.NewDecoder(av).Decode(faa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", faa.Type(), faa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FirebaseAppleApp] has state.
func (faa *FirebaseAppleApp) State() (*firebaseAppleAppState, bool) {
	return faa.state, faa.state != nil
}

// StateMust returns the state for [FirebaseAppleApp]. Panics if the state is nil.
func (faa *FirebaseAppleApp) StateMust() *firebaseAppleAppState {
	if faa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", faa.Type(), faa.LocalName()))
	}
	return faa.state
}

// FirebaseAppleAppArgs contains the configurations for google_firebase_apple_app.
type FirebaseAppleAppArgs struct {
	// AppStoreId: string, optional
	AppStoreId terra.StringValue `hcl:"app_store_id,attr"`
	// BundleId: string, optional
	BundleId terra.StringValue `hcl:"bundle_id,attr"`
	// DeletionPolicy: string, optional
	DeletionPolicy terra.StringValue `hcl:"deletion_policy,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// TeamId: string, optional
	TeamId terra.StringValue `hcl:"team_id,attr"`
	// Timeouts: optional
	Timeouts *firebaseappleapp.Timeouts `hcl:"timeouts,block"`
}
type firebaseAppleAppAttributes struct {
	ref terra.Reference
}

// AppId returns a reference to field app_id of google_firebase_apple_app.
func (faa firebaseAppleAppAttributes) AppId() terra.StringValue {
	return terra.ReferenceAsString(faa.ref.Append("app_id"))
}

// AppStoreId returns a reference to field app_store_id of google_firebase_apple_app.
func (faa firebaseAppleAppAttributes) AppStoreId() terra.StringValue {
	return terra.ReferenceAsString(faa.ref.Append("app_store_id"))
}

// BundleId returns a reference to field bundle_id of google_firebase_apple_app.
func (faa firebaseAppleAppAttributes) BundleId() terra.StringValue {
	return terra.ReferenceAsString(faa.ref.Append("bundle_id"))
}

// DeletionPolicy returns a reference to field deletion_policy of google_firebase_apple_app.
func (faa firebaseAppleAppAttributes) DeletionPolicy() terra.StringValue {
	return terra.ReferenceAsString(faa.ref.Append("deletion_policy"))
}

// DisplayName returns a reference to field display_name of google_firebase_apple_app.
func (faa firebaseAppleAppAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(faa.ref.Append("display_name"))
}

// Id returns a reference to field id of google_firebase_apple_app.
func (faa firebaseAppleAppAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(faa.ref.Append("id"))
}

// Name returns a reference to field name of google_firebase_apple_app.
func (faa firebaseAppleAppAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(faa.ref.Append("name"))
}

// Project returns a reference to field project of google_firebase_apple_app.
func (faa firebaseAppleAppAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(faa.ref.Append("project"))
}

// TeamId returns a reference to field team_id of google_firebase_apple_app.
func (faa firebaseAppleAppAttributes) TeamId() terra.StringValue {
	return terra.ReferenceAsString(faa.ref.Append("team_id"))
}

func (faa firebaseAppleAppAttributes) Timeouts() firebaseappleapp.TimeoutsAttributes {
	return terra.ReferenceAsSingle[firebaseappleapp.TimeoutsAttributes](faa.ref.Append("timeouts"))
}

type firebaseAppleAppState struct {
	AppId          string                          `json:"app_id"`
	AppStoreId     string                          `json:"app_store_id"`
	BundleId       string                          `json:"bundle_id"`
	DeletionPolicy string                          `json:"deletion_policy"`
	DisplayName    string                          `json:"display_name"`
	Id             string                          `json:"id"`
	Name           string                          `json:"name"`
	Project        string                          `json:"project"`
	TeamId         string                          `json:"team_id"`
	Timeouts       *firebaseappleapp.TimeoutsState `json:"timeouts"`
}