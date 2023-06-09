// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	firebasehostingsite "github.com/golingon/terraproviders/googlebeta/4.62.0/firebasehostingsite"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFirebaseHostingSite creates a new instance of [FirebaseHostingSite].
func NewFirebaseHostingSite(name string, args FirebaseHostingSiteArgs) *FirebaseHostingSite {
	return &FirebaseHostingSite{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FirebaseHostingSite)(nil)

// FirebaseHostingSite represents the Terraform resource google_firebase_hosting_site.
type FirebaseHostingSite struct {
	Name      string
	Args      FirebaseHostingSiteArgs
	state     *firebaseHostingSiteState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FirebaseHostingSite].
func (fhs *FirebaseHostingSite) Type() string {
	return "google_firebase_hosting_site"
}

// LocalName returns the local name for [FirebaseHostingSite].
func (fhs *FirebaseHostingSite) LocalName() string {
	return fhs.Name
}

// Configuration returns the configuration (args) for [FirebaseHostingSite].
func (fhs *FirebaseHostingSite) Configuration() interface{} {
	return fhs.Args
}

// DependOn is used for other resources to depend on [FirebaseHostingSite].
func (fhs *FirebaseHostingSite) DependOn() terra.Reference {
	return terra.ReferenceResource(fhs)
}

// Dependencies returns the list of resources [FirebaseHostingSite] depends_on.
func (fhs *FirebaseHostingSite) Dependencies() terra.Dependencies {
	return fhs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FirebaseHostingSite].
func (fhs *FirebaseHostingSite) LifecycleManagement() *terra.Lifecycle {
	return fhs.Lifecycle
}

// Attributes returns the attributes for [FirebaseHostingSite].
func (fhs *FirebaseHostingSite) Attributes() firebaseHostingSiteAttributes {
	return firebaseHostingSiteAttributes{ref: terra.ReferenceResource(fhs)}
}

// ImportState imports the given attribute values into [FirebaseHostingSite]'s state.
func (fhs *FirebaseHostingSite) ImportState(av io.Reader) error {
	fhs.state = &firebaseHostingSiteState{}
	if err := json.NewDecoder(av).Decode(fhs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fhs.Type(), fhs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FirebaseHostingSite] has state.
func (fhs *FirebaseHostingSite) State() (*firebaseHostingSiteState, bool) {
	return fhs.state, fhs.state != nil
}

// StateMust returns the state for [FirebaseHostingSite]. Panics if the state is nil.
func (fhs *FirebaseHostingSite) StateMust() *firebaseHostingSiteState {
	if fhs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fhs.Type(), fhs.LocalName()))
	}
	return fhs.state
}

// FirebaseHostingSiteArgs contains the configurations for google_firebase_hosting_site.
type FirebaseHostingSiteArgs struct {
	// AppId: string, optional
	AppId terra.StringValue `hcl:"app_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// SiteId: string, optional
	SiteId terra.StringValue `hcl:"site_id,attr"`
	// Timeouts: optional
	Timeouts *firebasehostingsite.Timeouts `hcl:"timeouts,block"`
}
type firebaseHostingSiteAttributes struct {
	ref terra.Reference
}

// AppId returns a reference to field app_id of google_firebase_hosting_site.
func (fhs firebaseHostingSiteAttributes) AppId() terra.StringValue {
	return terra.ReferenceAsString(fhs.ref.Append("app_id"))
}

// DefaultUrl returns a reference to field default_url of google_firebase_hosting_site.
func (fhs firebaseHostingSiteAttributes) DefaultUrl() terra.StringValue {
	return terra.ReferenceAsString(fhs.ref.Append("default_url"))
}

// Id returns a reference to field id of google_firebase_hosting_site.
func (fhs firebaseHostingSiteAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fhs.ref.Append("id"))
}

// Name returns a reference to field name of google_firebase_hosting_site.
func (fhs firebaseHostingSiteAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(fhs.ref.Append("name"))
}

// Project returns a reference to field project of google_firebase_hosting_site.
func (fhs firebaseHostingSiteAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(fhs.ref.Append("project"))
}

// SiteId returns a reference to field site_id of google_firebase_hosting_site.
func (fhs firebaseHostingSiteAttributes) SiteId() terra.StringValue {
	return terra.ReferenceAsString(fhs.ref.Append("site_id"))
}

func (fhs firebaseHostingSiteAttributes) Timeouts() firebasehostingsite.TimeoutsAttributes {
	return terra.ReferenceAsSingle[firebasehostingsite.TimeoutsAttributes](fhs.ref.Append("timeouts"))
}

type firebaseHostingSiteState struct {
	AppId      string                             `json:"app_id"`
	DefaultUrl string                             `json:"default_url"`
	Id         string                             `json:"id"`
	Name       string                             `json:"name"`
	Project    string                             `json:"project"`
	SiteId     string                             `json:"site_id"`
	Timeouts   *firebasehostingsite.TimeoutsState `json:"timeouts"`
}
