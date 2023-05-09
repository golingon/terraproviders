// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	fsxopenzfssnapshot "github.com/golingon/terraproviders/aws/4.66.1/fsxopenzfssnapshot"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFsxOpenzfsSnapshot creates a new instance of [FsxOpenzfsSnapshot].
func NewFsxOpenzfsSnapshot(name string, args FsxOpenzfsSnapshotArgs) *FsxOpenzfsSnapshot {
	return &FsxOpenzfsSnapshot{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FsxOpenzfsSnapshot)(nil)

// FsxOpenzfsSnapshot represents the Terraform resource aws_fsx_openzfs_snapshot.
type FsxOpenzfsSnapshot struct {
	Name      string
	Args      FsxOpenzfsSnapshotArgs
	state     *fsxOpenzfsSnapshotState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FsxOpenzfsSnapshot].
func (fos *FsxOpenzfsSnapshot) Type() string {
	return "aws_fsx_openzfs_snapshot"
}

// LocalName returns the local name for [FsxOpenzfsSnapshot].
func (fos *FsxOpenzfsSnapshot) LocalName() string {
	return fos.Name
}

// Configuration returns the configuration (args) for [FsxOpenzfsSnapshot].
func (fos *FsxOpenzfsSnapshot) Configuration() interface{} {
	return fos.Args
}

// DependOn is used for other resources to depend on [FsxOpenzfsSnapshot].
func (fos *FsxOpenzfsSnapshot) DependOn() terra.Reference {
	return terra.ReferenceResource(fos)
}

// Dependencies returns the list of resources [FsxOpenzfsSnapshot] depends_on.
func (fos *FsxOpenzfsSnapshot) Dependencies() terra.Dependencies {
	return fos.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FsxOpenzfsSnapshot].
func (fos *FsxOpenzfsSnapshot) LifecycleManagement() *terra.Lifecycle {
	return fos.Lifecycle
}

// Attributes returns the attributes for [FsxOpenzfsSnapshot].
func (fos *FsxOpenzfsSnapshot) Attributes() fsxOpenzfsSnapshotAttributes {
	return fsxOpenzfsSnapshotAttributes{ref: terra.ReferenceResource(fos)}
}

// ImportState imports the given attribute values into [FsxOpenzfsSnapshot]'s state.
func (fos *FsxOpenzfsSnapshot) ImportState(av io.Reader) error {
	fos.state = &fsxOpenzfsSnapshotState{}
	if err := json.NewDecoder(av).Decode(fos.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fos.Type(), fos.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FsxOpenzfsSnapshot] has state.
func (fos *FsxOpenzfsSnapshot) State() (*fsxOpenzfsSnapshotState, bool) {
	return fos.state, fos.state != nil
}

// StateMust returns the state for [FsxOpenzfsSnapshot]. Panics if the state is nil.
func (fos *FsxOpenzfsSnapshot) StateMust() *fsxOpenzfsSnapshotState {
	if fos.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fos.Type(), fos.LocalName()))
	}
	return fos.state
}

// FsxOpenzfsSnapshotArgs contains the configurations for aws_fsx_openzfs_snapshot.
type FsxOpenzfsSnapshotArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// VolumeId: string, required
	VolumeId terra.StringValue `hcl:"volume_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *fsxopenzfssnapshot.Timeouts `hcl:"timeouts,block"`
}
type fsxOpenzfsSnapshotAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_fsx_openzfs_snapshot.
func (fos fsxOpenzfsSnapshotAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(fos.ref.Append("arn"))
}

// CreationTime returns a reference to field creation_time of aws_fsx_openzfs_snapshot.
func (fos fsxOpenzfsSnapshotAttributes) CreationTime() terra.StringValue {
	return terra.ReferenceAsString(fos.ref.Append("creation_time"))
}

// Id returns a reference to field id of aws_fsx_openzfs_snapshot.
func (fos fsxOpenzfsSnapshotAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fos.ref.Append("id"))
}

// Name returns a reference to field name of aws_fsx_openzfs_snapshot.
func (fos fsxOpenzfsSnapshotAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(fos.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_fsx_openzfs_snapshot.
func (fos fsxOpenzfsSnapshotAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](fos.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_fsx_openzfs_snapshot.
func (fos fsxOpenzfsSnapshotAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](fos.ref.Append("tags_all"))
}

// VolumeId returns a reference to field volume_id of aws_fsx_openzfs_snapshot.
func (fos fsxOpenzfsSnapshotAttributes) VolumeId() terra.StringValue {
	return terra.ReferenceAsString(fos.ref.Append("volume_id"))
}

func (fos fsxOpenzfsSnapshotAttributes) Timeouts() fsxopenzfssnapshot.TimeoutsAttributes {
	return terra.ReferenceAsSingle[fsxopenzfssnapshot.TimeoutsAttributes](fos.ref.Append("timeouts"))
}

type fsxOpenzfsSnapshotState struct {
	Arn          string                            `json:"arn"`
	CreationTime string                            `json:"creation_time"`
	Id           string                            `json:"id"`
	Name         string                            `json:"name"`
	Tags         map[string]string                 `json:"tags"`
	TagsAll      map[string]string                 `json:"tags_all"`
	VolumeId     string                            `json:"volume_id"`
	Timeouts     *fsxopenzfssnapshot.TimeoutsState `json:"timeouts"`
}
