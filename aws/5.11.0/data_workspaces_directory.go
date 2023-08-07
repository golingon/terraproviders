// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataworkspacesdirectory "github.com/golingon/terraproviders/aws/5.11.0/dataworkspacesdirectory"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataWorkspacesDirectory creates a new instance of [DataWorkspacesDirectory].
func NewDataWorkspacesDirectory(name string, args DataWorkspacesDirectoryArgs) *DataWorkspacesDirectory {
	return &DataWorkspacesDirectory{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataWorkspacesDirectory)(nil)

// DataWorkspacesDirectory represents the Terraform data resource aws_workspaces_directory.
type DataWorkspacesDirectory struct {
	Name string
	Args DataWorkspacesDirectoryArgs
}

// DataSource returns the Terraform object type for [DataWorkspacesDirectory].
func (wd *DataWorkspacesDirectory) DataSource() string {
	return "aws_workspaces_directory"
}

// LocalName returns the local name for [DataWorkspacesDirectory].
func (wd *DataWorkspacesDirectory) LocalName() string {
	return wd.Name
}

// Configuration returns the configuration (args) for [DataWorkspacesDirectory].
func (wd *DataWorkspacesDirectory) Configuration() interface{} {
	return wd.Args
}

// Attributes returns the attributes for [DataWorkspacesDirectory].
func (wd *DataWorkspacesDirectory) Attributes() dataWorkspacesDirectoryAttributes {
	return dataWorkspacesDirectoryAttributes{ref: terra.ReferenceDataResource(wd)}
}

// DataWorkspacesDirectoryArgs contains the configurations for aws_workspaces_directory.
type DataWorkspacesDirectoryArgs struct {
	// DirectoryId: string, required
	DirectoryId terra.StringValue `hcl:"directory_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// SelfServicePermissions: min=0
	SelfServicePermissions []dataworkspacesdirectory.SelfServicePermissions `hcl:"self_service_permissions,block" validate:"min=0"`
	// WorkspaceAccessProperties: min=0
	WorkspaceAccessProperties []dataworkspacesdirectory.WorkspaceAccessProperties `hcl:"workspace_access_properties,block" validate:"min=0"`
	// WorkspaceCreationProperties: min=0
	WorkspaceCreationProperties []dataworkspacesdirectory.WorkspaceCreationProperties `hcl:"workspace_creation_properties,block" validate:"min=0"`
}
type dataWorkspacesDirectoryAttributes struct {
	ref terra.Reference
}

// Alias returns a reference to field alias of aws_workspaces_directory.
func (wd dataWorkspacesDirectoryAttributes) Alias() terra.StringValue {
	return terra.ReferenceAsString(wd.ref.Append("alias"))
}

// CustomerUserName returns a reference to field customer_user_name of aws_workspaces_directory.
func (wd dataWorkspacesDirectoryAttributes) CustomerUserName() terra.StringValue {
	return terra.ReferenceAsString(wd.ref.Append("customer_user_name"))
}

// DirectoryId returns a reference to field directory_id of aws_workspaces_directory.
func (wd dataWorkspacesDirectoryAttributes) DirectoryId() terra.StringValue {
	return terra.ReferenceAsString(wd.ref.Append("directory_id"))
}

// DirectoryName returns a reference to field directory_name of aws_workspaces_directory.
func (wd dataWorkspacesDirectoryAttributes) DirectoryName() terra.StringValue {
	return terra.ReferenceAsString(wd.ref.Append("directory_name"))
}

// DirectoryType returns a reference to field directory_type of aws_workspaces_directory.
func (wd dataWorkspacesDirectoryAttributes) DirectoryType() terra.StringValue {
	return terra.ReferenceAsString(wd.ref.Append("directory_type"))
}

// DnsIpAddresses returns a reference to field dns_ip_addresses of aws_workspaces_directory.
func (wd dataWorkspacesDirectoryAttributes) DnsIpAddresses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](wd.ref.Append("dns_ip_addresses"))
}

// IamRoleId returns a reference to field iam_role_id of aws_workspaces_directory.
func (wd dataWorkspacesDirectoryAttributes) IamRoleId() terra.StringValue {
	return terra.ReferenceAsString(wd.ref.Append("iam_role_id"))
}

// Id returns a reference to field id of aws_workspaces_directory.
func (wd dataWorkspacesDirectoryAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wd.ref.Append("id"))
}

// IpGroupIds returns a reference to field ip_group_ids of aws_workspaces_directory.
func (wd dataWorkspacesDirectoryAttributes) IpGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](wd.ref.Append("ip_group_ids"))
}

// RegistrationCode returns a reference to field registration_code of aws_workspaces_directory.
func (wd dataWorkspacesDirectoryAttributes) RegistrationCode() terra.StringValue {
	return terra.ReferenceAsString(wd.ref.Append("registration_code"))
}

// SubnetIds returns a reference to field subnet_ids of aws_workspaces_directory.
func (wd dataWorkspacesDirectoryAttributes) SubnetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](wd.ref.Append("subnet_ids"))
}

// Tags returns a reference to field tags of aws_workspaces_directory.
func (wd dataWorkspacesDirectoryAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](wd.ref.Append("tags"))
}

// WorkspaceSecurityGroupId returns a reference to field workspace_security_group_id of aws_workspaces_directory.
func (wd dataWorkspacesDirectoryAttributes) WorkspaceSecurityGroupId() terra.StringValue {
	return terra.ReferenceAsString(wd.ref.Append("workspace_security_group_id"))
}

func (wd dataWorkspacesDirectoryAttributes) SelfServicePermissions() terra.ListValue[dataworkspacesdirectory.SelfServicePermissionsAttributes] {
	return terra.ReferenceAsList[dataworkspacesdirectory.SelfServicePermissionsAttributes](wd.ref.Append("self_service_permissions"))
}

func (wd dataWorkspacesDirectoryAttributes) WorkspaceAccessProperties() terra.ListValue[dataworkspacesdirectory.WorkspaceAccessPropertiesAttributes] {
	return terra.ReferenceAsList[dataworkspacesdirectory.WorkspaceAccessPropertiesAttributes](wd.ref.Append("workspace_access_properties"))
}

func (wd dataWorkspacesDirectoryAttributes) WorkspaceCreationProperties() terra.ListValue[dataworkspacesdirectory.WorkspaceCreationPropertiesAttributes] {
	return terra.ReferenceAsList[dataworkspacesdirectory.WorkspaceCreationPropertiesAttributes](wd.ref.Append("workspace_creation_properties"))
}
