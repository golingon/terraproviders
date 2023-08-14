// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	firebasehostingrelease "github.com/golingon/terraproviders/googlebeta/4.77.0/firebasehostingrelease"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFirebaseHostingRelease creates a new instance of [FirebaseHostingRelease].
func NewFirebaseHostingRelease(name string, args FirebaseHostingReleaseArgs) *FirebaseHostingRelease {
	return &FirebaseHostingRelease{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FirebaseHostingRelease)(nil)

// FirebaseHostingRelease represents the Terraform resource google_firebase_hosting_release.
type FirebaseHostingRelease struct {
	Name      string
	Args      FirebaseHostingReleaseArgs
	state     *firebaseHostingReleaseState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FirebaseHostingRelease].
func (fhr *FirebaseHostingRelease) Type() string {
	return "google_firebase_hosting_release"
}

// LocalName returns the local name for [FirebaseHostingRelease].
func (fhr *FirebaseHostingRelease) LocalName() string {
	return fhr.Name
}

// Configuration returns the configuration (args) for [FirebaseHostingRelease].
func (fhr *FirebaseHostingRelease) Configuration() interface{} {
	return fhr.Args
}

// DependOn is used for other resources to depend on [FirebaseHostingRelease].
func (fhr *FirebaseHostingRelease) DependOn() terra.Reference {
	return terra.ReferenceResource(fhr)
}

// Dependencies returns the list of resources [FirebaseHostingRelease] depends_on.
func (fhr *FirebaseHostingRelease) Dependencies() terra.Dependencies {
	return fhr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FirebaseHostingRelease].
func (fhr *FirebaseHostingRelease) LifecycleManagement() *terra.Lifecycle {
	return fhr.Lifecycle
}

// Attributes returns the attributes for [FirebaseHostingRelease].
func (fhr *FirebaseHostingRelease) Attributes() firebaseHostingReleaseAttributes {
	return firebaseHostingReleaseAttributes{ref: terra.ReferenceResource(fhr)}
}

// ImportState imports the given attribute values into [FirebaseHostingRelease]'s state.
func (fhr *FirebaseHostingRelease) ImportState(av io.Reader) error {
	fhr.state = &firebaseHostingReleaseState{}
	if err := json.NewDecoder(av).Decode(fhr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fhr.Type(), fhr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FirebaseHostingRelease] has state.
func (fhr *FirebaseHostingRelease) State() (*firebaseHostingReleaseState, bool) {
	return fhr.state, fhr.state != nil
}

// StateMust returns the state for [FirebaseHostingRelease]. Panics if the state is nil.
func (fhr *FirebaseHostingRelease) StateMust() *firebaseHostingReleaseState {
	if fhr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fhr.Type(), fhr.LocalName()))
	}
	return fhr.state
}

// FirebaseHostingReleaseArgs contains the configurations for google_firebase_hosting_release.
type FirebaseHostingReleaseArgs struct {
	// ChannelId: string, optional
	ChannelId terra.StringValue `hcl:"channel_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Message: string, optional
	Message terra.StringValue `hcl:"message,attr"`
	// SiteId: string, required
	SiteId terra.StringValue `hcl:"site_id,attr" validate:"required"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// VersionName: string, optional
	VersionName terra.StringValue `hcl:"version_name,attr"`
	// Timeouts: optional
	Timeouts *firebasehostingrelease.Timeouts `hcl:"timeouts,block"`
}
type firebaseHostingReleaseAttributes struct {
	ref terra.Reference
}

// ChannelId returns a reference to field channel_id of google_firebase_hosting_release.
func (fhr firebaseHostingReleaseAttributes) ChannelId() terra.StringValue {
	return terra.ReferenceAsString(fhr.ref.Append("channel_id"))
}

// Id returns a reference to field id of google_firebase_hosting_release.
func (fhr firebaseHostingReleaseAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fhr.ref.Append("id"))
}

// Message returns a reference to field message of google_firebase_hosting_release.
func (fhr firebaseHostingReleaseAttributes) Message() terra.StringValue {
	return terra.ReferenceAsString(fhr.ref.Append("message"))
}

// Name returns a reference to field name of google_firebase_hosting_release.
func (fhr firebaseHostingReleaseAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(fhr.ref.Append("name"))
}

// ReleaseId returns a reference to field release_id of google_firebase_hosting_release.
func (fhr firebaseHostingReleaseAttributes) ReleaseId() terra.StringValue {
	return terra.ReferenceAsString(fhr.ref.Append("release_id"))
}

// SiteId returns a reference to field site_id of google_firebase_hosting_release.
func (fhr firebaseHostingReleaseAttributes) SiteId() terra.StringValue {
	return terra.ReferenceAsString(fhr.ref.Append("site_id"))
}

// Type returns a reference to field type of google_firebase_hosting_release.
func (fhr firebaseHostingReleaseAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(fhr.ref.Append("type"))
}

// VersionName returns a reference to field version_name of google_firebase_hosting_release.
func (fhr firebaseHostingReleaseAttributes) VersionName() terra.StringValue {
	return terra.ReferenceAsString(fhr.ref.Append("version_name"))
}

func (fhr firebaseHostingReleaseAttributes) Timeouts() firebasehostingrelease.TimeoutsAttributes {
	return terra.ReferenceAsSingle[firebasehostingrelease.TimeoutsAttributes](fhr.ref.Append("timeouts"))
}

type firebaseHostingReleaseState struct {
	ChannelId   string                                `json:"channel_id"`
	Id          string                                `json:"id"`
	Message     string                                `json:"message"`
	Name        string                                `json:"name"`
	ReleaseId   string                                `json:"release_id"`
	SiteId      string                                `json:"site_id"`
	Type        string                                `json:"type"`
	VersionName string                                `json:"version_name"`
	Timeouts    *firebasehostingrelease.TimeoutsState `json:"timeouts"`
}
