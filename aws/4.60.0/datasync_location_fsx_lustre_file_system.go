// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDatasyncLocationFsxLustreFileSystem creates a new instance of [DatasyncLocationFsxLustreFileSystem].
func NewDatasyncLocationFsxLustreFileSystem(name string, args DatasyncLocationFsxLustreFileSystemArgs) *DatasyncLocationFsxLustreFileSystem {
	return &DatasyncLocationFsxLustreFileSystem{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DatasyncLocationFsxLustreFileSystem)(nil)

// DatasyncLocationFsxLustreFileSystem represents the Terraform resource aws_datasync_location_fsx_lustre_file_system.
type DatasyncLocationFsxLustreFileSystem struct {
	Name      string
	Args      DatasyncLocationFsxLustreFileSystemArgs
	state     *datasyncLocationFsxLustreFileSystemState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DatasyncLocationFsxLustreFileSystem].
func (dlflfs *DatasyncLocationFsxLustreFileSystem) Type() string {
	return "aws_datasync_location_fsx_lustre_file_system"
}

// LocalName returns the local name for [DatasyncLocationFsxLustreFileSystem].
func (dlflfs *DatasyncLocationFsxLustreFileSystem) LocalName() string {
	return dlflfs.Name
}

// Configuration returns the configuration (args) for [DatasyncLocationFsxLustreFileSystem].
func (dlflfs *DatasyncLocationFsxLustreFileSystem) Configuration() interface{} {
	return dlflfs.Args
}

// DependOn is used for other resources to depend on [DatasyncLocationFsxLustreFileSystem].
func (dlflfs *DatasyncLocationFsxLustreFileSystem) DependOn() terra.Reference {
	return terra.ReferenceResource(dlflfs)
}

// Dependencies returns the list of resources [DatasyncLocationFsxLustreFileSystem] depends_on.
func (dlflfs *DatasyncLocationFsxLustreFileSystem) Dependencies() terra.Dependencies {
	return dlflfs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DatasyncLocationFsxLustreFileSystem].
func (dlflfs *DatasyncLocationFsxLustreFileSystem) LifecycleManagement() *terra.Lifecycle {
	return dlflfs.Lifecycle
}

// Attributes returns the attributes for [DatasyncLocationFsxLustreFileSystem].
func (dlflfs *DatasyncLocationFsxLustreFileSystem) Attributes() datasyncLocationFsxLustreFileSystemAttributes {
	return datasyncLocationFsxLustreFileSystemAttributes{ref: terra.ReferenceResource(dlflfs)}
}

// ImportState imports the given attribute values into [DatasyncLocationFsxLustreFileSystem]'s state.
func (dlflfs *DatasyncLocationFsxLustreFileSystem) ImportState(av io.Reader) error {
	dlflfs.state = &datasyncLocationFsxLustreFileSystemState{}
	if err := json.NewDecoder(av).Decode(dlflfs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dlflfs.Type(), dlflfs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DatasyncLocationFsxLustreFileSystem] has state.
func (dlflfs *DatasyncLocationFsxLustreFileSystem) State() (*datasyncLocationFsxLustreFileSystemState, bool) {
	return dlflfs.state, dlflfs.state != nil
}

// StateMust returns the state for [DatasyncLocationFsxLustreFileSystem]. Panics if the state is nil.
func (dlflfs *DatasyncLocationFsxLustreFileSystem) StateMust() *datasyncLocationFsxLustreFileSystemState {
	if dlflfs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dlflfs.Type(), dlflfs.LocalName()))
	}
	return dlflfs.state
}

// DatasyncLocationFsxLustreFileSystemArgs contains the configurations for aws_datasync_location_fsx_lustre_file_system.
type DatasyncLocationFsxLustreFileSystemArgs struct {
	// FsxFilesystemArn: string, required
	FsxFilesystemArn terra.StringValue `hcl:"fsx_filesystem_arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SecurityGroupArns: set of string, required
	SecurityGroupArns terra.SetValue[terra.StringValue] `hcl:"security_group_arns,attr" validate:"required"`
	// Subdirectory: string, optional
	Subdirectory terra.StringValue `hcl:"subdirectory,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}
type datasyncLocationFsxLustreFileSystemAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_datasync_location_fsx_lustre_file_system.
func (dlflfs datasyncLocationFsxLustreFileSystemAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(dlflfs.ref.Append("arn"))
}

// CreationTime returns a reference to field creation_time of aws_datasync_location_fsx_lustre_file_system.
func (dlflfs datasyncLocationFsxLustreFileSystemAttributes) CreationTime() terra.StringValue {
	return terra.ReferenceAsString(dlflfs.ref.Append("creation_time"))
}

// FsxFilesystemArn returns a reference to field fsx_filesystem_arn of aws_datasync_location_fsx_lustre_file_system.
func (dlflfs datasyncLocationFsxLustreFileSystemAttributes) FsxFilesystemArn() terra.StringValue {
	return terra.ReferenceAsString(dlflfs.ref.Append("fsx_filesystem_arn"))
}

// Id returns a reference to field id of aws_datasync_location_fsx_lustre_file_system.
func (dlflfs datasyncLocationFsxLustreFileSystemAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dlflfs.ref.Append("id"))
}

// SecurityGroupArns returns a reference to field security_group_arns of aws_datasync_location_fsx_lustre_file_system.
func (dlflfs datasyncLocationFsxLustreFileSystemAttributes) SecurityGroupArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dlflfs.ref.Append("security_group_arns"))
}

// Subdirectory returns a reference to field subdirectory of aws_datasync_location_fsx_lustre_file_system.
func (dlflfs datasyncLocationFsxLustreFileSystemAttributes) Subdirectory() terra.StringValue {
	return terra.ReferenceAsString(dlflfs.ref.Append("subdirectory"))
}

// Tags returns a reference to field tags of aws_datasync_location_fsx_lustre_file_system.
func (dlflfs datasyncLocationFsxLustreFileSystemAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dlflfs.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_datasync_location_fsx_lustre_file_system.
func (dlflfs datasyncLocationFsxLustreFileSystemAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dlflfs.ref.Append("tags_all"))
}

// Uri returns a reference to field uri of aws_datasync_location_fsx_lustre_file_system.
func (dlflfs datasyncLocationFsxLustreFileSystemAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(dlflfs.ref.Append("uri"))
}

type datasyncLocationFsxLustreFileSystemState struct {
	Arn               string            `json:"arn"`
	CreationTime      string            `json:"creation_time"`
	FsxFilesystemArn  string            `json:"fsx_filesystem_arn"`
	Id                string            `json:"id"`
	SecurityGroupArns []string          `json:"security_group_arns"`
	Subdirectory      string            `json:"subdirectory"`
	Tags              map[string]string `json:"tags"`
	TagsAll           map[string]string `json:"tags_all"`
	Uri               string            `json:"uri"`
}