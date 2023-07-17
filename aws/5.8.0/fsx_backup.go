// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	fsxbackup "github.com/golingon/terraproviders/aws/5.8.0/fsxbackup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFsxBackup creates a new instance of [FsxBackup].
func NewFsxBackup(name string, args FsxBackupArgs) *FsxBackup {
	return &FsxBackup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FsxBackup)(nil)

// FsxBackup represents the Terraform resource aws_fsx_backup.
type FsxBackup struct {
	Name      string
	Args      FsxBackupArgs
	state     *fsxBackupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FsxBackup].
func (fb *FsxBackup) Type() string {
	return "aws_fsx_backup"
}

// LocalName returns the local name for [FsxBackup].
func (fb *FsxBackup) LocalName() string {
	return fb.Name
}

// Configuration returns the configuration (args) for [FsxBackup].
func (fb *FsxBackup) Configuration() interface{} {
	return fb.Args
}

// DependOn is used for other resources to depend on [FsxBackup].
func (fb *FsxBackup) DependOn() terra.Reference {
	return terra.ReferenceResource(fb)
}

// Dependencies returns the list of resources [FsxBackup] depends_on.
func (fb *FsxBackup) Dependencies() terra.Dependencies {
	return fb.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FsxBackup].
func (fb *FsxBackup) LifecycleManagement() *terra.Lifecycle {
	return fb.Lifecycle
}

// Attributes returns the attributes for [FsxBackup].
func (fb *FsxBackup) Attributes() fsxBackupAttributes {
	return fsxBackupAttributes{ref: terra.ReferenceResource(fb)}
}

// ImportState imports the given attribute values into [FsxBackup]'s state.
func (fb *FsxBackup) ImportState(av io.Reader) error {
	fb.state = &fsxBackupState{}
	if err := json.NewDecoder(av).Decode(fb.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fb.Type(), fb.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FsxBackup] has state.
func (fb *FsxBackup) State() (*fsxBackupState, bool) {
	return fb.state, fb.state != nil
}

// StateMust returns the state for [FsxBackup]. Panics if the state is nil.
func (fb *FsxBackup) StateMust() *fsxBackupState {
	if fb.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fb.Type(), fb.LocalName()))
	}
	return fb.state
}

// FsxBackupArgs contains the configurations for aws_fsx_backup.
type FsxBackupArgs struct {
	// FileSystemId: string, optional
	FileSystemId terra.StringValue `hcl:"file_system_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// VolumeId: string, optional
	VolumeId terra.StringValue `hcl:"volume_id,attr"`
	// Timeouts: optional
	Timeouts *fsxbackup.Timeouts `hcl:"timeouts,block"`
}
type fsxBackupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_fsx_backup.
func (fb fsxBackupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(fb.ref.Append("arn"))
}

// FileSystemId returns a reference to field file_system_id of aws_fsx_backup.
func (fb fsxBackupAttributes) FileSystemId() terra.StringValue {
	return terra.ReferenceAsString(fb.ref.Append("file_system_id"))
}

// Id returns a reference to field id of aws_fsx_backup.
func (fb fsxBackupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fb.ref.Append("id"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_fsx_backup.
func (fb fsxBackupAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(fb.ref.Append("kms_key_id"))
}

// OwnerId returns a reference to field owner_id of aws_fsx_backup.
func (fb fsxBackupAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(fb.ref.Append("owner_id"))
}

// Tags returns a reference to field tags of aws_fsx_backup.
func (fb fsxBackupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](fb.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_fsx_backup.
func (fb fsxBackupAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](fb.ref.Append("tags_all"))
}

// Type returns a reference to field type of aws_fsx_backup.
func (fb fsxBackupAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(fb.ref.Append("type"))
}

// VolumeId returns a reference to field volume_id of aws_fsx_backup.
func (fb fsxBackupAttributes) VolumeId() terra.StringValue {
	return terra.ReferenceAsString(fb.ref.Append("volume_id"))
}

func (fb fsxBackupAttributes) Timeouts() fsxbackup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[fsxbackup.TimeoutsAttributes](fb.ref.Append("timeouts"))
}

type fsxBackupState struct {
	Arn          string                   `json:"arn"`
	FileSystemId string                   `json:"file_system_id"`
	Id           string                   `json:"id"`
	KmsKeyId     string                   `json:"kms_key_id"`
	OwnerId      string                   `json:"owner_id"`
	Tags         map[string]string        `json:"tags"`
	TagsAll      map[string]string        `json:"tags_all"`
	Type         string                   `json:"type"`
	VolumeId     string                   `json:"volume_id"`
	Timeouts     *fsxbackup.TimeoutsState `json:"timeouts"`
}
