// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	firebasestoragebucket "github.com/golingon/terraproviders/googlebeta/4.63.1/firebasestoragebucket"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFirebaseStorageBucket creates a new instance of [FirebaseStorageBucket].
func NewFirebaseStorageBucket(name string, args FirebaseStorageBucketArgs) *FirebaseStorageBucket {
	return &FirebaseStorageBucket{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FirebaseStorageBucket)(nil)

// FirebaseStorageBucket represents the Terraform resource google_firebase_storage_bucket.
type FirebaseStorageBucket struct {
	Name      string
	Args      FirebaseStorageBucketArgs
	state     *firebaseStorageBucketState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FirebaseStorageBucket].
func (fsb *FirebaseStorageBucket) Type() string {
	return "google_firebase_storage_bucket"
}

// LocalName returns the local name for [FirebaseStorageBucket].
func (fsb *FirebaseStorageBucket) LocalName() string {
	return fsb.Name
}

// Configuration returns the configuration (args) for [FirebaseStorageBucket].
func (fsb *FirebaseStorageBucket) Configuration() interface{} {
	return fsb.Args
}

// DependOn is used for other resources to depend on [FirebaseStorageBucket].
func (fsb *FirebaseStorageBucket) DependOn() terra.Reference {
	return terra.ReferenceResource(fsb)
}

// Dependencies returns the list of resources [FirebaseStorageBucket] depends_on.
func (fsb *FirebaseStorageBucket) Dependencies() terra.Dependencies {
	return fsb.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FirebaseStorageBucket].
func (fsb *FirebaseStorageBucket) LifecycleManagement() *terra.Lifecycle {
	return fsb.Lifecycle
}

// Attributes returns the attributes for [FirebaseStorageBucket].
func (fsb *FirebaseStorageBucket) Attributes() firebaseStorageBucketAttributes {
	return firebaseStorageBucketAttributes{ref: terra.ReferenceResource(fsb)}
}

// ImportState imports the given attribute values into [FirebaseStorageBucket]'s state.
func (fsb *FirebaseStorageBucket) ImportState(av io.Reader) error {
	fsb.state = &firebaseStorageBucketState{}
	if err := json.NewDecoder(av).Decode(fsb.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fsb.Type(), fsb.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FirebaseStorageBucket] has state.
func (fsb *FirebaseStorageBucket) State() (*firebaseStorageBucketState, bool) {
	return fsb.state, fsb.state != nil
}

// StateMust returns the state for [FirebaseStorageBucket]. Panics if the state is nil.
func (fsb *FirebaseStorageBucket) StateMust() *firebaseStorageBucketState {
	if fsb.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fsb.Type(), fsb.LocalName()))
	}
	return fsb.state
}

// FirebaseStorageBucketArgs contains the configurations for google_firebase_storage_bucket.
type FirebaseStorageBucketArgs struct {
	// BucketId: string, optional
	BucketId terra.StringValue `hcl:"bucket_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *firebasestoragebucket.Timeouts `hcl:"timeouts,block"`
}
type firebaseStorageBucketAttributes struct {
	ref terra.Reference
}

// BucketId returns a reference to field bucket_id of google_firebase_storage_bucket.
func (fsb firebaseStorageBucketAttributes) BucketId() terra.StringValue {
	return terra.ReferenceAsString(fsb.ref.Append("bucket_id"))
}

// Id returns a reference to field id of google_firebase_storage_bucket.
func (fsb firebaseStorageBucketAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fsb.ref.Append("id"))
}

// Name returns a reference to field name of google_firebase_storage_bucket.
func (fsb firebaseStorageBucketAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(fsb.ref.Append("name"))
}

// Project returns a reference to field project of google_firebase_storage_bucket.
func (fsb firebaseStorageBucketAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(fsb.ref.Append("project"))
}

func (fsb firebaseStorageBucketAttributes) Timeouts() firebasestoragebucket.TimeoutsAttributes {
	return terra.ReferenceAsSingle[firebasestoragebucket.TimeoutsAttributes](fsb.ref.Append("timeouts"))
}

type firebaseStorageBucketState struct {
	BucketId string                               `json:"bucket_id"`
	Id       string                               `json:"id"`
	Name     string                               `json:"name"`
	Project  string                               `json:"project"`
	Timeouts *firebasestoragebucket.TimeoutsState `json:"timeouts"`
}
