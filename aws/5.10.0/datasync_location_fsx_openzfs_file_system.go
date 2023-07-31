// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	datasynclocationfsxopenzfsfilesystem "github.com/golingon/terraproviders/aws/5.10.0/datasynclocationfsxopenzfsfilesystem"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDatasyncLocationFsxOpenzfsFileSystem creates a new instance of [DatasyncLocationFsxOpenzfsFileSystem].
func NewDatasyncLocationFsxOpenzfsFileSystem(name string, args DatasyncLocationFsxOpenzfsFileSystemArgs) *DatasyncLocationFsxOpenzfsFileSystem {
	return &DatasyncLocationFsxOpenzfsFileSystem{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DatasyncLocationFsxOpenzfsFileSystem)(nil)

// DatasyncLocationFsxOpenzfsFileSystem represents the Terraform resource aws_datasync_location_fsx_openzfs_file_system.
type DatasyncLocationFsxOpenzfsFileSystem struct {
	Name      string
	Args      DatasyncLocationFsxOpenzfsFileSystemArgs
	state     *datasyncLocationFsxOpenzfsFileSystemState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DatasyncLocationFsxOpenzfsFileSystem].
func (dlfofs *DatasyncLocationFsxOpenzfsFileSystem) Type() string {
	return "aws_datasync_location_fsx_openzfs_file_system"
}

// LocalName returns the local name for [DatasyncLocationFsxOpenzfsFileSystem].
func (dlfofs *DatasyncLocationFsxOpenzfsFileSystem) LocalName() string {
	return dlfofs.Name
}

// Configuration returns the configuration (args) for [DatasyncLocationFsxOpenzfsFileSystem].
func (dlfofs *DatasyncLocationFsxOpenzfsFileSystem) Configuration() interface{} {
	return dlfofs.Args
}

// DependOn is used for other resources to depend on [DatasyncLocationFsxOpenzfsFileSystem].
func (dlfofs *DatasyncLocationFsxOpenzfsFileSystem) DependOn() terra.Reference {
	return terra.ReferenceResource(dlfofs)
}

// Dependencies returns the list of resources [DatasyncLocationFsxOpenzfsFileSystem] depends_on.
func (dlfofs *DatasyncLocationFsxOpenzfsFileSystem) Dependencies() terra.Dependencies {
	return dlfofs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DatasyncLocationFsxOpenzfsFileSystem].
func (dlfofs *DatasyncLocationFsxOpenzfsFileSystem) LifecycleManagement() *terra.Lifecycle {
	return dlfofs.Lifecycle
}

// Attributes returns the attributes for [DatasyncLocationFsxOpenzfsFileSystem].
func (dlfofs *DatasyncLocationFsxOpenzfsFileSystem) Attributes() datasyncLocationFsxOpenzfsFileSystemAttributes {
	return datasyncLocationFsxOpenzfsFileSystemAttributes{ref: terra.ReferenceResource(dlfofs)}
}

// ImportState imports the given attribute values into [DatasyncLocationFsxOpenzfsFileSystem]'s state.
func (dlfofs *DatasyncLocationFsxOpenzfsFileSystem) ImportState(av io.Reader) error {
	dlfofs.state = &datasyncLocationFsxOpenzfsFileSystemState{}
	if err := json.NewDecoder(av).Decode(dlfofs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dlfofs.Type(), dlfofs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DatasyncLocationFsxOpenzfsFileSystem] has state.
func (dlfofs *DatasyncLocationFsxOpenzfsFileSystem) State() (*datasyncLocationFsxOpenzfsFileSystemState, bool) {
	return dlfofs.state, dlfofs.state != nil
}

// StateMust returns the state for [DatasyncLocationFsxOpenzfsFileSystem]. Panics if the state is nil.
func (dlfofs *DatasyncLocationFsxOpenzfsFileSystem) StateMust() *datasyncLocationFsxOpenzfsFileSystemState {
	if dlfofs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dlfofs.Type(), dlfofs.LocalName()))
	}
	return dlfofs.state
}

// DatasyncLocationFsxOpenzfsFileSystemArgs contains the configurations for aws_datasync_location_fsx_openzfs_file_system.
type DatasyncLocationFsxOpenzfsFileSystemArgs struct {
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
	// Protocol: required
	Protocol *datasynclocationfsxopenzfsfilesystem.Protocol `hcl:"protocol,block" validate:"required"`
}
type datasyncLocationFsxOpenzfsFileSystemAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_datasync_location_fsx_openzfs_file_system.
func (dlfofs datasyncLocationFsxOpenzfsFileSystemAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(dlfofs.ref.Append("arn"))
}

// CreationTime returns a reference to field creation_time of aws_datasync_location_fsx_openzfs_file_system.
func (dlfofs datasyncLocationFsxOpenzfsFileSystemAttributes) CreationTime() terra.StringValue {
	return terra.ReferenceAsString(dlfofs.ref.Append("creation_time"))
}

// FsxFilesystemArn returns a reference to field fsx_filesystem_arn of aws_datasync_location_fsx_openzfs_file_system.
func (dlfofs datasyncLocationFsxOpenzfsFileSystemAttributes) FsxFilesystemArn() terra.StringValue {
	return terra.ReferenceAsString(dlfofs.ref.Append("fsx_filesystem_arn"))
}

// Id returns a reference to field id of aws_datasync_location_fsx_openzfs_file_system.
func (dlfofs datasyncLocationFsxOpenzfsFileSystemAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dlfofs.ref.Append("id"))
}

// SecurityGroupArns returns a reference to field security_group_arns of aws_datasync_location_fsx_openzfs_file_system.
func (dlfofs datasyncLocationFsxOpenzfsFileSystemAttributes) SecurityGroupArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dlfofs.ref.Append("security_group_arns"))
}

// Subdirectory returns a reference to field subdirectory of aws_datasync_location_fsx_openzfs_file_system.
func (dlfofs datasyncLocationFsxOpenzfsFileSystemAttributes) Subdirectory() terra.StringValue {
	return terra.ReferenceAsString(dlfofs.ref.Append("subdirectory"))
}

// Tags returns a reference to field tags of aws_datasync_location_fsx_openzfs_file_system.
func (dlfofs datasyncLocationFsxOpenzfsFileSystemAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dlfofs.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_datasync_location_fsx_openzfs_file_system.
func (dlfofs datasyncLocationFsxOpenzfsFileSystemAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dlfofs.ref.Append("tags_all"))
}

// Uri returns a reference to field uri of aws_datasync_location_fsx_openzfs_file_system.
func (dlfofs datasyncLocationFsxOpenzfsFileSystemAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(dlfofs.ref.Append("uri"))
}

func (dlfofs datasyncLocationFsxOpenzfsFileSystemAttributes) Protocol() terra.ListValue[datasynclocationfsxopenzfsfilesystem.ProtocolAttributes] {
	return terra.ReferenceAsList[datasynclocationfsxopenzfsfilesystem.ProtocolAttributes](dlfofs.ref.Append("protocol"))
}

type datasyncLocationFsxOpenzfsFileSystemState struct {
	Arn               string                                               `json:"arn"`
	CreationTime      string                                               `json:"creation_time"`
	FsxFilesystemArn  string                                               `json:"fsx_filesystem_arn"`
	Id                string                                               `json:"id"`
	SecurityGroupArns []string                                             `json:"security_group_arns"`
	Subdirectory      string                                               `json:"subdirectory"`
	Tags              map[string]string                                    `json:"tags"`
	TagsAll           map[string]string                                    `json:"tags_all"`
	Uri               string                                               `json:"uri"`
	Protocol          []datasynclocationfsxopenzfsfilesystem.ProtocolState `json:"protocol"`
}
