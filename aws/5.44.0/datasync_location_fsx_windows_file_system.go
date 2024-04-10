// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewDatasyncLocationFsxWindowsFileSystem creates a new instance of [DatasyncLocationFsxWindowsFileSystem].
func NewDatasyncLocationFsxWindowsFileSystem(name string, args DatasyncLocationFsxWindowsFileSystemArgs) *DatasyncLocationFsxWindowsFileSystem {
	return &DatasyncLocationFsxWindowsFileSystem{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DatasyncLocationFsxWindowsFileSystem)(nil)

// DatasyncLocationFsxWindowsFileSystem represents the Terraform resource aws_datasync_location_fsx_windows_file_system.
type DatasyncLocationFsxWindowsFileSystem struct {
	Name      string
	Args      DatasyncLocationFsxWindowsFileSystemArgs
	state     *datasyncLocationFsxWindowsFileSystemState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DatasyncLocationFsxWindowsFileSystem].
func (dlfwfs *DatasyncLocationFsxWindowsFileSystem) Type() string {
	return "aws_datasync_location_fsx_windows_file_system"
}

// LocalName returns the local name for [DatasyncLocationFsxWindowsFileSystem].
func (dlfwfs *DatasyncLocationFsxWindowsFileSystem) LocalName() string {
	return dlfwfs.Name
}

// Configuration returns the configuration (args) for [DatasyncLocationFsxWindowsFileSystem].
func (dlfwfs *DatasyncLocationFsxWindowsFileSystem) Configuration() interface{} {
	return dlfwfs.Args
}

// DependOn is used for other resources to depend on [DatasyncLocationFsxWindowsFileSystem].
func (dlfwfs *DatasyncLocationFsxWindowsFileSystem) DependOn() terra.Reference {
	return terra.ReferenceResource(dlfwfs)
}

// Dependencies returns the list of resources [DatasyncLocationFsxWindowsFileSystem] depends_on.
func (dlfwfs *DatasyncLocationFsxWindowsFileSystem) Dependencies() terra.Dependencies {
	return dlfwfs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DatasyncLocationFsxWindowsFileSystem].
func (dlfwfs *DatasyncLocationFsxWindowsFileSystem) LifecycleManagement() *terra.Lifecycle {
	return dlfwfs.Lifecycle
}

// Attributes returns the attributes for [DatasyncLocationFsxWindowsFileSystem].
func (dlfwfs *DatasyncLocationFsxWindowsFileSystem) Attributes() datasyncLocationFsxWindowsFileSystemAttributes {
	return datasyncLocationFsxWindowsFileSystemAttributes{ref: terra.ReferenceResource(dlfwfs)}
}

// ImportState imports the given attribute values into [DatasyncLocationFsxWindowsFileSystem]'s state.
func (dlfwfs *DatasyncLocationFsxWindowsFileSystem) ImportState(av io.Reader) error {
	dlfwfs.state = &datasyncLocationFsxWindowsFileSystemState{}
	if err := json.NewDecoder(av).Decode(dlfwfs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dlfwfs.Type(), dlfwfs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DatasyncLocationFsxWindowsFileSystem] has state.
func (dlfwfs *DatasyncLocationFsxWindowsFileSystem) State() (*datasyncLocationFsxWindowsFileSystemState, bool) {
	return dlfwfs.state, dlfwfs.state != nil
}

// StateMust returns the state for [DatasyncLocationFsxWindowsFileSystem]. Panics if the state is nil.
func (dlfwfs *DatasyncLocationFsxWindowsFileSystem) StateMust() *datasyncLocationFsxWindowsFileSystemState {
	if dlfwfs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dlfwfs.Type(), dlfwfs.LocalName()))
	}
	return dlfwfs.state
}

// DatasyncLocationFsxWindowsFileSystemArgs contains the configurations for aws_datasync_location_fsx_windows_file_system.
type DatasyncLocationFsxWindowsFileSystemArgs struct {
	// Domain: string, optional
	Domain terra.StringValue `hcl:"domain,attr"`
	// FsxFilesystemArn: string, required
	FsxFilesystemArn terra.StringValue `hcl:"fsx_filesystem_arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Password: string, required
	Password terra.StringValue `hcl:"password,attr" validate:"required"`
	// SecurityGroupArns: set of string, required
	SecurityGroupArns terra.SetValue[terra.StringValue] `hcl:"security_group_arns,attr" validate:"required"`
	// Subdirectory: string, optional
	Subdirectory terra.StringValue `hcl:"subdirectory,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// User: string, required
	User terra.StringValue `hcl:"user,attr" validate:"required"`
}
type datasyncLocationFsxWindowsFileSystemAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_datasync_location_fsx_windows_file_system.
func (dlfwfs datasyncLocationFsxWindowsFileSystemAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(dlfwfs.ref.Append("arn"))
}

// CreationTime returns a reference to field creation_time of aws_datasync_location_fsx_windows_file_system.
func (dlfwfs datasyncLocationFsxWindowsFileSystemAttributes) CreationTime() terra.StringValue {
	return terra.ReferenceAsString(dlfwfs.ref.Append("creation_time"))
}

// Domain returns a reference to field domain of aws_datasync_location_fsx_windows_file_system.
func (dlfwfs datasyncLocationFsxWindowsFileSystemAttributes) Domain() terra.StringValue {
	return terra.ReferenceAsString(dlfwfs.ref.Append("domain"))
}

// FsxFilesystemArn returns a reference to field fsx_filesystem_arn of aws_datasync_location_fsx_windows_file_system.
func (dlfwfs datasyncLocationFsxWindowsFileSystemAttributes) FsxFilesystemArn() terra.StringValue {
	return terra.ReferenceAsString(dlfwfs.ref.Append("fsx_filesystem_arn"))
}

// Id returns a reference to field id of aws_datasync_location_fsx_windows_file_system.
func (dlfwfs datasyncLocationFsxWindowsFileSystemAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dlfwfs.ref.Append("id"))
}

// Password returns a reference to field password of aws_datasync_location_fsx_windows_file_system.
func (dlfwfs datasyncLocationFsxWindowsFileSystemAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(dlfwfs.ref.Append("password"))
}

// SecurityGroupArns returns a reference to field security_group_arns of aws_datasync_location_fsx_windows_file_system.
func (dlfwfs datasyncLocationFsxWindowsFileSystemAttributes) SecurityGroupArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dlfwfs.ref.Append("security_group_arns"))
}

// Subdirectory returns a reference to field subdirectory of aws_datasync_location_fsx_windows_file_system.
func (dlfwfs datasyncLocationFsxWindowsFileSystemAttributes) Subdirectory() terra.StringValue {
	return terra.ReferenceAsString(dlfwfs.ref.Append("subdirectory"))
}

// Tags returns a reference to field tags of aws_datasync_location_fsx_windows_file_system.
func (dlfwfs datasyncLocationFsxWindowsFileSystemAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dlfwfs.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_datasync_location_fsx_windows_file_system.
func (dlfwfs datasyncLocationFsxWindowsFileSystemAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dlfwfs.ref.Append("tags_all"))
}

// Uri returns a reference to field uri of aws_datasync_location_fsx_windows_file_system.
func (dlfwfs datasyncLocationFsxWindowsFileSystemAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(dlfwfs.ref.Append("uri"))
}

// User returns a reference to field user of aws_datasync_location_fsx_windows_file_system.
func (dlfwfs datasyncLocationFsxWindowsFileSystemAttributes) User() terra.StringValue {
	return terra.ReferenceAsString(dlfwfs.ref.Append("user"))
}

type datasyncLocationFsxWindowsFileSystemState struct {
	Arn               string            `json:"arn"`
	CreationTime      string            `json:"creation_time"`
	Domain            string            `json:"domain"`
	FsxFilesystemArn  string            `json:"fsx_filesystem_arn"`
	Id                string            `json:"id"`
	Password          string            `json:"password"`
	SecurityGroupArns []string          `json:"security_group_arns"`
	Subdirectory      string            `json:"subdirectory"`
	Tags              map[string]string `json:"tags"`
	TagsAll           map[string]string `json:"tags_all"`
	Uri               string            `json:"uri"`
	User              string            `json:"user"`
}
