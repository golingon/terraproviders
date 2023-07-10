// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	workspacesdirectory "github.com/golingon/terraproviders/aws/5.7.0/workspacesdirectory"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewWorkspacesDirectory creates a new instance of [WorkspacesDirectory].
func NewWorkspacesDirectory(name string, args WorkspacesDirectoryArgs) *WorkspacesDirectory {
	return &WorkspacesDirectory{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WorkspacesDirectory)(nil)

// WorkspacesDirectory represents the Terraform resource aws_workspaces_directory.
type WorkspacesDirectory struct {
	Name      string
	Args      WorkspacesDirectoryArgs
	state     *workspacesDirectoryState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WorkspacesDirectory].
func (wd *WorkspacesDirectory) Type() string {
	return "aws_workspaces_directory"
}

// LocalName returns the local name for [WorkspacesDirectory].
func (wd *WorkspacesDirectory) LocalName() string {
	return wd.Name
}

// Configuration returns the configuration (args) for [WorkspacesDirectory].
func (wd *WorkspacesDirectory) Configuration() interface{} {
	return wd.Args
}

// DependOn is used for other resources to depend on [WorkspacesDirectory].
func (wd *WorkspacesDirectory) DependOn() terra.Reference {
	return terra.ReferenceResource(wd)
}

// Dependencies returns the list of resources [WorkspacesDirectory] depends_on.
func (wd *WorkspacesDirectory) Dependencies() terra.Dependencies {
	return wd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WorkspacesDirectory].
func (wd *WorkspacesDirectory) LifecycleManagement() *terra.Lifecycle {
	return wd.Lifecycle
}

// Attributes returns the attributes for [WorkspacesDirectory].
func (wd *WorkspacesDirectory) Attributes() workspacesDirectoryAttributes {
	return workspacesDirectoryAttributes{ref: terra.ReferenceResource(wd)}
}

// ImportState imports the given attribute values into [WorkspacesDirectory]'s state.
func (wd *WorkspacesDirectory) ImportState(av io.Reader) error {
	wd.state = &workspacesDirectoryState{}
	if err := json.NewDecoder(av).Decode(wd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wd.Type(), wd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WorkspacesDirectory] has state.
func (wd *WorkspacesDirectory) State() (*workspacesDirectoryState, bool) {
	return wd.state, wd.state != nil
}

// StateMust returns the state for [WorkspacesDirectory]. Panics if the state is nil.
func (wd *WorkspacesDirectory) StateMust() *workspacesDirectoryState {
	if wd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wd.Type(), wd.LocalName()))
	}
	return wd.state
}

// WorkspacesDirectoryArgs contains the configurations for aws_workspaces_directory.
type WorkspacesDirectoryArgs struct {
	// DirectoryId: string, required
	DirectoryId terra.StringValue `hcl:"directory_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpGroupIds: set of string, optional
	IpGroupIds terra.SetValue[terra.StringValue] `hcl:"ip_group_ids,attr"`
	// SubnetIds: set of string, optional
	SubnetIds terra.SetValue[terra.StringValue] `hcl:"subnet_ids,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// SelfServicePermissions: optional
	SelfServicePermissions *workspacesdirectory.SelfServicePermissions `hcl:"self_service_permissions,block"`
	// WorkspaceAccessProperties: optional
	WorkspaceAccessProperties *workspacesdirectory.WorkspaceAccessProperties `hcl:"workspace_access_properties,block"`
	// WorkspaceCreationProperties: optional
	WorkspaceCreationProperties *workspacesdirectory.WorkspaceCreationProperties `hcl:"workspace_creation_properties,block"`
}
type workspacesDirectoryAttributes struct {
	ref terra.Reference
}

// Alias returns a reference to field alias of aws_workspaces_directory.
func (wd workspacesDirectoryAttributes) Alias() terra.StringValue {
	return terra.ReferenceAsString(wd.ref.Append("alias"))
}

// CustomerUserName returns a reference to field customer_user_name of aws_workspaces_directory.
func (wd workspacesDirectoryAttributes) CustomerUserName() terra.StringValue {
	return terra.ReferenceAsString(wd.ref.Append("customer_user_name"))
}

// DirectoryId returns a reference to field directory_id of aws_workspaces_directory.
func (wd workspacesDirectoryAttributes) DirectoryId() terra.StringValue {
	return terra.ReferenceAsString(wd.ref.Append("directory_id"))
}

// DirectoryName returns a reference to field directory_name of aws_workspaces_directory.
func (wd workspacesDirectoryAttributes) DirectoryName() terra.StringValue {
	return terra.ReferenceAsString(wd.ref.Append("directory_name"))
}

// DirectoryType returns a reference to field directory_type of aws_workspaces_directory.
func (wd workspacesDirectoryAttributes) DirectoryType() terra.StringValue {
	return terra.ReferenceAsString(wd.ref.Append("directory_type"))
}

// DnsIpAddresses returns a reference to field dns_ip_addresses of aws_workspaces_directory.
func (wd workspacesDirectoryAttributes) DnsIpAddresses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](wd.ref.Append("dns_ip_addresses"))
}

// IamRoleId returns a reference to field iam_role_id of aws_workspaces_directory.
func (wd workspacesDirectoryAttributes) IamRoleId() terra.StringValue {
	return terra.ReferenceAsString(wd.ref.Append("iam_role_id"))
}

// Id returns a reference to field id of aws_workspaces_directory.
func (wd workspacesDirectoryAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wd.ref.Append("id"))
}

// IpGroupIds returns a reference to field ip_group_ids of aws_workspaces_directory.
func (wd workspacesDirectoryAttributes) IpGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](wd.ref.Append("ip_group_ids"))
}

// RegistrationCode returns a reference to field registration_code of aws_workspaces_directory.
func (wd workspacesDirectoryAttributes) RegistrationCode() terra.StringValue {
	return terra.ReferenceAsString(wd.ref.Append("registration_code"))
}

// SubnetIds returns a reference to field subnet_ids of aws_workspaces_directory.
func (wd workspacesDirectoryAttributes) SubnetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](wd.ref.Append("subnet_ids"))
}

// Tags returns a reference to field tags of aws_workspaces_directory.
func (wd workspacesDirectoryAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](wd.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_workspaces_directory.
func (wd workspacesDirectoryAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](wd.ref.Append("tags_all"))
}

// WorkspaceSecurityGroupId returns a reference to field workspace_security_group_id of aws_workspaces_directory.
func (wd workspacesDirectoryAttributes) WorkspaceSecurityGroupId() terra.StringValue {
	return terra.ReferenceAsString(wd.ref.Append("workspace_security_group_id"))
}

func (wd workspacesDirectoryAttributes) SelfServicePermissions() terra.ListValue[workspacesdirectory.SelfServicePermissionsAttributes] {
	return terra.ReferenceAsList[workspacesdirectory.SelfServicePermissionsAttributes](wd.ref.Append("self_service_permissions"))
}

func (wd workspacesDirectoryAttributes) WorkspaceAccessProperties() terra.ListValue[workspacesdirectory.WorkspaceAccessPropertiesAttributes] {
	return terra.ReferenceAsList[workspacesdirectory.WorkspaceAccessPropertiesAttributes](wd.ref.Append("workspace_access_properties"))
}

func (wd workspacesDirectoryAttributes) WorkspaceCreationProperties() terra.ListValue[workspacesdirectory.WorkspaceCreationPropertiesAttributes] {
	return terra.ReferenceAsList[workspacesdirectory.WorkspaceCreationPropertiesAttributes](wd.ref.Append("workspace_creation_properties"))
}

type workspacesDirectoryState struct {
	Alias                       string                                                 `json:"alias"`
	CustomerUserName            string                                                 `json:"customer_user_name"`
	DirectoryId                 string                                                 `json:"directory_id"`
	DirectoryName               string                                                 `json:"directory_name"`
	DirectoryType               string                                                 `json:"directory_type"`
	DnsIpAddresses              []string                                               `json:"dns_ip_addresses"`
	IamRoleId                   string                                                 `json:"iam_role_id"`
	Id                          string                                                 `json:"id"`
	IpGroupIds                  []string                                               `json:"ip_group_ids"`
	RegistrationCode            string                                                 `json:"registration_code"`
	SubnetIds                   []string                                               `json:"subnet_ids"`
	Tags                        map[string]string                                      `json:"tags"`
	TagsAll                     map[string]string                                      `json:"tags_all"`
	WorkspaceSecurityGroupId    string                                                 `json:"workspace_security_group_id"`
	SelfServicePermissions      []workspacesdirectory.SelfServicePermissionsState      `json:"self_service_permissions"`
	WorkspaceAccessProperties   []workspacesdirectory.WorkspaceAccessPropertiesState   `json:"workspace_access_properties"`
	WorkspaceCreationProperties []workspacesdirectory.WorkspaceCreationPropertiesState `json:"workspace_creation_properties"`
}
