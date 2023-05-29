// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataconnectuser "github.com/golingon/terraproviders/aws/5.0.1/dataconnectuser"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataConnectUser creates a new instance of [DataConnectUser].
func NewDataConnectUser(name string, args DataConnectUserArgs) *DataConnectUser {
	return &DataConnectUser{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataConnectUser)(nil)

// DataConnectUser represents the Terraform data resource aws_connect_user.
type DataConnectUser struct {
	Name string
	Args DataConnectUserArgs
}

// DataSource returns the Terraform object type for [DataConnectUser].
func (cu *DataConnectUser) DataSource() string {
	return "aws_connect_user"
}

// LocalName returns the local name for [DataConnectUser].
func (cu *DataConnectUser) LocalName() string {
	return cu.Name
}

// Configuration returns the configuration (args) for [DataConnectUser].
func (cu *DataConnectUser) Configuration() interface{} {
	return cu.Args
}

// Attributes returns the attributes for [DataConnectUser].
func (cu *DataConnectUser) Attributes() dataConnectUserAttributes {
	return dataConnectUserAttributes{ref: terra.ReferenceDataResource(cu)}
}

// DataConnectUserArgs contains the configurations for aws_connect_user.
type DataConnectUserArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceId: string, required
	InstanceId terra.StringValue `hcl:"instance_id,attr" validate:"required"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// UserId: string, optional
	UserId terra.StringValue `hcl:"user_id,attr"`
	// IdentityInfo: min=0
	IdentityInfo []dataconnectuser.IdentityInfo `hcl:"identity_info,block" validate:"min=0"`
	// PhoneConfig: min=0
	PhoneConfig []dataconnectuser.PhoneConfig `hcl:"phone_config,block" validate:"min=0"`
}
type dataConnectUserAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_connect_user.
func (cu dataConnectUserAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cu.ref.Append("arn"))
}

// DirectoryUserId returns a reference to field directory_user_id of aws_connect_user.
func (cu dataConnectUserAttributes) DirectoryUserId() terra.StringValue {
	return terra.ReferenceAsString(cu.ref.Append("directory_user_id"))
}

// HierarchyGroupId returns a reference to field hierarchy_group_id of aws_connect_user.
func (cu dataConnectUserAttributes) HierarchyGroupId() terra.StringValue {
	return terra.ReferenceAsString(cu.ref.Append("hierarchy_group_id"))
}

// Id returns a reference to field id of aws_connect_user.
func (cu dataConnectUserAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cu.ref.Append("id"))
}

// InstanceId returns a reference to field instance_id of aws_connect_user.
func (cu dataConnectUserAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(cu.ref.Append("instance_id"))
}

// Name returns a reference to field name of aws_connect_user.
func (cu dataConnectUserAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cu.ref.Append("name"))
}

// RoutingProfileId returns a reference to field routing_profile_id of aws_connect_user.
func (cu dataConnectUserAttributes) RoutingProfileId() terra.StringValue {
	return terra.ReferenceAsString(cu.ref.Append("routing_profile_id"))
}

// SecurityProfileIds returns a reference to field security_profile_ids of aws_connect_user.
func (cu dataConnectUserAttributes) SecurityProfileIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cu.ref.Append("security_profile_ids"))
}

// Tags returns a reference to field tags of aws_connect_user.
func (cu dataConnectUserAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cu.ref.Append("tags"))
}

// UserId returns a reference to field user_id of aws_connect_user.
func (cu dataConnectUserAttributes) UserId() terra.StringValue {
	return terra.ReferenceAsString(cu.ref.Append("user_id"))
}

func (cu dataConnectUserAttributes) IdentityInfo() terra.ListValue[dataconnectuser.IdentityInfoAttributes] {
	return terra.ReferenceAsList[dataconnectuser.IdentityInfoAttributes](cu.ref.Append("identity_info"))
}

func (cu dataConnectUserAttributes) PhoneConfig() terra.ListValue[dataconnectuser.PhoneConfigAttributes] {
	return terra.ReferenceAsList[dataconnectuser.PhoneConfigAttributes](cu.ref.Append("phone_config"))
}