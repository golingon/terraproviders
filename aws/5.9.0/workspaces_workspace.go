// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	workspacesworkspace "github.com/golingon/terraproviders/aws/5.9.0/workspacesworkspace"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewWorkspacesWorkspace creates a new instance of [WorkspacesWorkspace].
func NewWorkspacesWorkspace(name string, args WorkspacesWorkspaceArgs) *WorkspacesWorkspace {
	return &WorkspacesWorkspace{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WorkspacesWorkspace)(nil)

// WorkspacesWorkspace represents the Terraform resource aws_workspaces_workspace.
type WorkspacesWorkspace struct {
	Name      string
	Args      WorkspacesWorkspaceArgs
	state     *workspacesWorkspaceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WorkspacesWorkspace].
func (ww *WorkspacesWorkspace) Type() string {
	return "aws_workspaces_workspace"
}

// LocalName returns the local name for [WorkspacesWorkspace].
func (ww *WorkspacesWorkspace) LocalName() string {
	return ww.Name
}

// Configuration returns the configuration (args) for [WorkspacesWorkspace].
func (ww *WorkspacesWorkspace) Configuration() interface{} {
	return ww.Args
}

// DependOn is used for other resources to depend on [WorkspacesWorkspace].
func (ww *WorkspacesWorkspace) DependOn() terra.Reference {
	return terra.ReferenceResource(ww)
}

// Dependencies returns the list of resources [WorkspacesWorkspace] depends_on.
func (ww *WorkspacesWorkspace) Dependencies() terra.Dependencies {
	return ww.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WorkspacesWorkspace].
func (ww *WorkspacesWorkspace) LifecycleManagement() *terra.Lifecycle {
	return ww.Lifecycle
}

// Attributes returns the attributes for [WorkspacesWorkspace].
func (ww *WorkspacesWorkspace) Attributes() workspacesWorkspaceAttributes {
	return workspacesWorkspaceAttributes{ref: terra.ReferenceResource(ww)}
}

// ImportState imports the given attribute values into [WorkspacesWorkspace]'s state.
func (ww *WorkspacesWorkspace) ImportState(av io.Reader) error {
	ww.state = &workspacesWorkspaceState{}
	if err := json.NewDecoder(av).Decode(ww.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ww.Type(), ww.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WorkspacesWorkspace] has state.
func (ww *WorkspacesWorkspace) State() (*workspacesWorkspaceState, bool) {
	return ww.state, ww.state != nil
}

// StateMust returns the state for [WorkspacesWorkspace]. Panics if the state is nil.
func (ww *WorkspacesWorkspace) StateMust() *workspacesWorkspaceState {
	if ww.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ww.Type(), ww.LocalName()))
	}
	return ww.state
}

// WorkspacesWorkspaceArgs contains the configurations for aws_workspaces_workspace.
type WorkspacesWorkspaceArgs struct {
	// BundleId: string, required
	BundleId terra.StringValue `hcl:"bundle_id,attr" validate:"required"`
	// DirectoryId: string, required
	DirectoryId terra.StringValue `hcl:"directory_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RootVolumeEncryptionEnabled: bool, optional
	RootVolumeEncryptionEnabled terra.BoolValue `hcl:"root_volume_encryption_enabled,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// UserName: string, required
	UserName terra.StringValue `hcl:"user_name,attr" validate:"required"`
	// UserVolumeEncryptionEnabled: bool, optional
	UserVolumeEncryptionEnabled terra.BoolValue `hcl:"user_volume_encryption_enabled,attr"`
	// VolumeEncryptionKey: string, optional
	VolumeEncryptionKey terra.StringValue `hcl:"volume_encryption_key,attr"`
	// Timeouts: optional
	Timeouts *workspacesworkspace.Timeouts `hcl:"timeouts,block"`
	// WorkspaceProperties: optional
	WorkspaceProperties *workspacesworkspace.WorkspaceProperties `hcl:"workspace_properties,block"`
}
type workspacesWorkspaceAttributes struct {
	ref terra.Reference
}

// BundleId returns a reference to field bundle_id of aws_workspaces_workspace.
func (ww workspacesWorkspaceAttributes) BundleId() terra.StringValue {
	return terra.ReferenceAsString(ww.ref.Append("bundle_id"))
}

// ComputerName returns a reference to field computer_name of aws_workspaces_workspace.
func (ww workspacesWorkspaceAttributes) ComputerName() terra.StringValue {
	return terra.ReferenceAsString(ww.ref.Append("computer_name"))
}

// DirectoryId returns a reference to field directory_id of aws_workspaces_workspace.
func (ww workspacesWorkspaceAttributes) DirectoryId() terra.StringValue {
	return terra.ReferenceAsString(ww.ref.Append("directory_id"))
}

// Id returns a reference to field id of aws_workspaces_workspace.
func (ww workspacesWorkspaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ww.ref.Append("id"))
}

// IpAddress returns a reference to field ip_address of aws_workspaces_workspace.
func (ww workspacesWorkspaceAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(ww.ref.Append("ip_address"))
}

// RootVolumeEncryptionEnabled returns a reference to field root_volume_encryption_enabled of aws_workspaces_workspace.
func (ww workspacesWorkspaceAttributes) RootVolumeEncryptionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ww.ref.Append("root_volume_encryption_enabled"))
}

// State returns a reference to field state of aws_workspaces_workspace.
func (ww workspacesWorkspaceAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(ww.ref.Append("state"))
}

// Tags returns a reference to field tags of aws_workspaces_workspace.
func (ww workspacesWorkspaceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ww.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_workspaces_workspace.
func (ww workspacesWorkspaceAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ww.ref.Append("tags_all"))
}

// UserName returns a reference to field user_name of aws_workspaces_workspace.
func (ww workspacesWorkspaceAttributes) UserName() terra.StringValue {
	return terra.ReferenceAsString(ww.ref.Append("user_name"))
}

// UserVolumeEncryptionEnabled returns a reference to field user_volume_encryption_enabled of aws_workspaces_workspace.
func (ww workspacesWorkspaceAttributes) UserVolumeEncryptionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ww.ref.Append("user_volume_encryption_enabled"))
}

// VolumeEncryptionKey returns a reference to field volume_encryption_key of aws_workspaces_workspace.
func (ww workspacesWorkspaceAttributes) VolumeEncryptionKey() terra.StringValue {
	return terra.ReferenceAsString(ww.ref.Append("volume_encryption_key"))
}

func (ww workspacesWorkspaceAttributes) Timeouts() workspacesworkspace.TimeoutsAttributes {
	return terra.ReferenceAsSingle[workspacesworkspace.TimeoutsAttributes](ww.ref.Append("timeouts"))
}

func (ww workspacesWorkspaceAttributes) WorkspaceProperties() terra.ListValue[workspacesworkspace.WorkspacePropertiesAttributes] {
	return terra.ReferenceAsList[workspacesworkspace.WorkspacePropertiesAttributes](ww.ref.Append("workspace_properties"))
}

type workspacesWorkspaceState struct {
	BundleId                    string                                         `json:"bundle_id"`
	ComputerName                string                                         `json:"computer_name"`
	DirectoryId                 string                                         `json:"directory_id"`
	Id                          string                                         `json:"id"`
	IpAddress                   string                                         `json:"ip_address"`
	RootVolumeEncryptionEnabled bool                                           `json:"root_volume_encryption_enabled"`
	State                       string                                         `json:"state"`
	Tags                        map[string]string                              `json:"tags"`
	TagsAll                     map[string]string                              `json:"tags_all"`
	UserName                    string                                         `json:"user_name"`
	UserVolumeEncryptionEnabled bool                                           `json:"user_volume_encryption_enabled"`
	VolumeEncryptionKey         string                                         `json:"volume_encryption_key"`
	Timeouts                    *workspacesworkspace.TimeoutsState             `json:"timeouts"`
	WorkspaceProperties         []workspacesworkspace.WorkspacePropertiesState `json:"workspace_properties"`
}
