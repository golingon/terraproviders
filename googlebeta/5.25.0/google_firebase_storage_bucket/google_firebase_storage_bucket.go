// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_firebase_storage_bucket

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource google_firebase_storage_bucket.
type Resource struct {
	Name      string
	Args      Args
	state     *googleFirebaseStorageBucketState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gfsb *Resource) Type() string {
	return "google_firebase_storage_bucket"
}

// LocalName returns the local name for [Resource].
func (gfsb *Resource) LocalName() string {
	return gfsb.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gfsb *Resource) Configuration() interface{} {
	return gfsb.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gfsb *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gfsb)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gfsb *Resource) Dependencies() terra.Dependencies {
	return gfsb.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gfsb *Resource) LifecycleManagement() *terra.Lifecycle {
	return gfsb.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gfsb *Resource) Attributes() googleFirebaseStorageBucketAttributes {
	return googleFirebaseStorageBucketAttributes{ref: terra.ReferenceResource(gfsb)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gfsb *Resource) ImportState(state io.Reader) error {
	gfsb.state = &googleFirebaseStorageBucketState{}
	if err := json.NewDecoder(state).Decode(gfsb.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gfsb.Type(), gfsb.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gfsb *Resource) State() (*googleFirebaseStorageBucketState, bool) {
	return gfsb.state, gfsb.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gfsb *Resource) StateMust() *googleFirebaseStorageBucketState {
	if gfsb.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gfsb.Type(), gfsb.LocalName()))
	}
	return gfsb.state
}

// Args contains the configurations for google_firebase_storage_bucket.
type Args struct {
	// BucketId: string, optional
	BucketId terra.StringValue `hcl:"bucket_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleFirebaseStorageBucketAttributes struct {
	ref terra.Reference
}

// BucketId returns a reference to field bucket_id of google_firebase_storage_bucket.
func (gfsb googleFirebaseStorageBucketAttributes) BucketId() terra.StringValue {
	return terra.ReferenceAsString(gfsb.ref.Append("bucket_id"))
}

// Id returns a reference to field id of google_firebase_storage_bucket.
func (gfsb googleFirebaseStorageBucketAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gfsb.ref.Append("id"))
}

// Name returns a reference to field name of google_firebase_storage_bucket.
func (gfsb googleFirebaseStorageBucketAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gfsb.ref.Append("name"))
}

// Project returns a reference to field project of google_firebase_storage_bucket.
func (gfsb googleFirebaseStorageBucketAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gfsb.ref.Append("project"))
}

func (gfsb googleFirebaseStorageBucketAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gfsb.ref.Append("timeouts"))
}

type googleFirebaseStorageBucketState struct {
	BucketId string         `json:"bucket_id"`
	Id       string         `json:"id"`
	Name     string         `json:"name"`
	Project  string         `json:"project"`
	Timeouts *TimeoutsState `json:"timeouts"`
}
