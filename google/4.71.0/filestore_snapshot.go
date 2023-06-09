// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	filestoresnapshot "github.com/golingon/terraproviders/google/4.71.0/filestoresnapshot"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFilestoreSnapshot creates a new instance of [FilestoreSnapshot].
func NewFilestoreSnapshot(name string, args FilestoreSnapshotArgs) *FilestoreSnapshot {
	return &FilestoreSnapshot{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FilestoreSnapshot)(nil)

// FilestoreSnapshot represents the Terraform resource google_filestore_snapshot.
type FilestoreSnapshot struct {
	Name      string
	Args      FilestoreSnapshotArgs
	state     *filestoreSnapshotState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FilestoreSnapshot].
func (fs *FilestoreSnapshot) Type() string {
	return "google_filestore_snapshot"
}

// LocalName returns the local name for [FilestoreSnapshot].
func (fs *FilestoreSnapshot) LocalName() string {
	return fs.Name
}

// Configuration returns the configuration (args) for [FilestoreSnapshot].
func (fs *FilestoreSnapshot) Configuration() interface{} {
	return fs.Args
}

// DependOn is used for other resources to depend on [FilestoreSnapshot].
func (fs *FilestoreSnapshot) DependOn() terra.Reference {
	return terra.ReferenceResource(fs)
}

// Dependencies returns the list of resources [FilestoreSnapshot] depends_on.
func (fs *FilestoreSnapshot) Dependencies() terra.Dependencies {
	return fs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FilestoreSnapshot].
func (fs *FilestoreSnapshot) LifecycleManagement() *terra.Lifecycle {
	return fs.Lifecycle
}

// Attributes returns the attributes for [FilestoreSnapshot].
func (fs *FilestoreSnapshot) Attributes() filestoreSnapshotAttributes {
	return filestoreSnapshotAttributes{ref: terra.ReferenceResource(fs)}
}

// ImportState imports the given attribute values into [FilestoreSnapshot]'s state.
func (fs *FilestoreSnapshot) ImportState(av io.Reader) error {
	fs.state = &filestoreSnapshotState{}
	if err := json.NewDecoder(av).Decode(fs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fs.Type(), fs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FilestoreSnapshot] has state.
func (fs *FilestoreSnapshot) State() (*filestoreSnapshotState, bool) {
	return fs.state, fs.state != nil
}

// StateMust returns the state for [FilestoreSnapshot]. Panics if the state is nil.
func (fs *FilestoreSnapshot) StateMust() *filestoreSnapshotState {
	if fs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fs.Type(), fs.LocalName()))
	}
	return fs.state
}

// FilestoreSnapshotArgs contains the configurations for google_filestore_snapshot.
type FilestoreSnapshotArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Instance: string, required
	Instance terra.StringValue `hcl:"instance,attr" validate:"required"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *filestoresnapshot.Timeouts `hcl:"timeouts,block"`
}
type filestoreSnapshotAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_filestore_snapshot.
func (fs filestoreSnapshotAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(fs.ref.Append("create_time"))
}

// Description returns a reference to field description of google_filestore_snapshot.
func (fs filestoreSnapshotAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(fs.ref.Append("description"))
}

// FilesystemUsedBytes returns a reference to field filesystem_used_bytes of google_filestore_snapshot.
func (fs filestoreSnapshotAttributes) FilesystemUsedBytes() terra.StringValue {
	return terra.ReferenceAsString(fs.ref.Append("filesystem_used_bytes"))
}

// Id returns a reference to field id of google_filestore_snapshot.
func (fs filestoreSnapshotAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fs.ref.Append("id"))
}

// Instance returns a reference to field instance of google_filestore_snapshot.
func (fs filestoreSnapshotAttributes) Instance() terra.StringValue {
	return terra.ReferenceAsString(fs.ref.Append("instance"))
}

// Labels returns a reference to field labels of google_filestore_snapshot.
func (fs filestoreSnapshotAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](fs.ref.Append("labels"))
}

// Location returns a reference to field location of google_filestore_snapshot.
func (fs filestoreSnapshotAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(fs.ref.Append("location"))
}

// Name returns a reference to field name of google_filestore_snapshot.
func (fs filestoreSnapshotAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(fs.ref.Append("name"))
}

// Project returns a reference to field project of google_filestore_snapshot.
func (fs filestoreSnapshotAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(fs.ref.Append("project"))
}

// State returns a reference to field state of google_filestore_snapshot.
func (fs filestoreSnapshotAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(fs.ref.Append("state"))
}

func (fs filestoreSnapshotAttributes) Timeouts() filestoresnapshot.TimeoutsAttributes {
	return terra.ReferenceAsSingle[filestoresnapshot.TimeoutsAttributes](fs.ref.Append("timeouts"))
}

type filestoreSnapshotState struct {
	CreateTime          string                           `json:"create_time"`
	Description         string                           `json:"description"`
	FilesystemUsedBytes string                           `json:"filesystem_used_bytes"`
	Id                  string                           `json:"id"`
	Instance            string                           `json:"instance"`
	Labels              map[string]string                `json:"labels"`
	Location            string                           `json:"location"`
	Name                string                           `json:"name"`
	Project             string                           `json:"project"`
	State               string                           `json:"state"`
	Timeouts            *filestoresnapshot.TimeoutsState `json:"timeouts"`
}
