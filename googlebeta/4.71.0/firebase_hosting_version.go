// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	firebasehostingversion "github.com/golingon/terraproviders/googlebeta/4.71.0/firebasehostingversion"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFirebaseHostingVersion creates a new instance of [FirebaseHostingVersion].
func NewFirebaseHostingVersion(name string, args FirebaseHostingVersionArgs) *FirebaseHostingVersion {
	return &FirebaseHostingVersion{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FirebaseHostingVersion)(nil)

// FirebaseHostingVersion represents the Terraform resource google_firebase_hosting_version.
type FirebaseHostingVersion struct {
	Name      string
	Args      FirebaseHostingVersionArgs
	state     *firebaseHostingVersionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FirebaseHostingVersion].
func (fhv *FirebaseHostingVersion) Type() string {
	return "google_firebase_hosting_version"
}

// LocalName returns the local name for [FirebaseHostingVersion].
func (fhv *FirebaseHostingVersion) LocalName() string {
	return fhv.Name
}

// Configuration returns the configuration (args) for [FirebaseHostingVersion].
func (fhv *FirebaseHostingVersion) Configuration() interface{} {
	return fhv.Args
}

// DependOn is used for other resources to depend on [FirebaseHostingVersion].
func (fhv *FirebaseHostingVersion) DependOn() terra.Reference {
	return terra.ReferenceResource(fhv)
}

// Dependencies returns the list of resources [FirebaseHostingVersion] depends_on.
func (fhv *FirebaseHostingVersion) Dependencies() terra.Dependencies {
	return fhv.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FirebaseHostingVersion].
func (fhv *FirebaseHostingVersion) LifecycleManagement() *terra.Lifecycle {
	return fhv.Lifecycle
}

// Attributes returns the attributes for [FirebaseHostingVersion].
func (fhv *FirebaseHostingVersion) Attributes() firebaseHostingVersionAttributes {
	return firebaseHostingVersionAttributes{ref: terra.ReferenceResource(fhv)}
}

// ImportState imports the given attribute values into [FirebaseHostingVersion]'s state.
func (fhv *FirebaseHostingVersion) ImportState(av io.Reader) error {
	fhv.state = &firebaseHostingVersionState{}
	if err := json.NewDecoder(av).Decode(fhv.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fhv.Type(), fhv.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FirebaseHostingVersion] has state.
func (fhv *FirebaseHostingVersion) State() (*firebaseHostingVersionState, bool) {
	return fhv.state, fhv.state != nil
}

// StateMust returns the state for [FirebaseHostingVersion]. Panics if the state is nil.
func (fhv *FirebaseHostingVersion) StateMust() *firebaseHostingVersionState {
	if fhv.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fhv.Type(), fhv.LocalName()))
	}
	return fhv.state
}

// FirebaseHostingVersionArgs contains the configurations for google_firebase_hosting_version.
type FirebaseHostingVersionArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SiteId: string, required
	SiteId terra.StringValue `hcl:"site_id,attr" validate:"required"`
	// Config: optional
	Config *firebasehostingversion.Config `hcl:"config,block"`
	// Timeouts: optional
	Timeouts *firebasehostingversion.Timeouts `hcl:"timeouts,block"`
}
type firebaseHostingVersionAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_firebase_hosting_version.
func (fhv firebaseHostingVersionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fhv.ref.Append("id"))
}

// Name returns a reference to field name of google_firebase_hosting_version.
func (fhv firebaseHostingVersionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(fhv.ref.Append("name"))
}

// SiteId returns a reference to field site_id of google_firebase_hosting_version.
func (fhv firebaseHostingVersionAttributes) SiteId() terra.StringValue {
	return terra.ReferenceAsString(fhv.ref.Append("site_id"))
}

// VersionId returns a reference to field version_id of google_firebase_hosting_version.
func (fhv firebaseHostingVersionAttributes) VersionId() terra.StringValue {
	return terra.ReferenceAsString(fhv.ref.Append("version_id"))
}

func (fhv firebaseHostingVersionAttributes) Config() terra.ListValue[firebasehostingversion.ConfigAttributes] {
	return terra.ReferenceAsList[firebasehostingversion.ConfigAttributes](fhv.ref.Append("config"))
}

func (fhv firebaseHostingVersionAttributes) Timeouts() firebasehostingversion.TimeoutsAttributes {
	return terra.ReferenceAsSingle[firebasehostingversion.TimeoutsAttributes](fhv.ref.Append("timeouts"))
}

type firebaseHostingVersionState struct {
	Id        string                                `json:"id"`
	Name      string                                `json:"name"`
	SiteId    string                                `json:"site_id"`
	VersionId string                                `json:"version_id"`
	Config    []firebasehostingversion.ConfigState  `json:"config"`
	Timeouts  *firebasehostingversion.TimeoutsState `json:"timeouts"`
}
