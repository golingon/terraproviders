// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	connectuser "github.com/golingon/terraproviders/aws/4.60.0/connectuser"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewConnectUser creates a new instance of [ConnectUser].
func NewConnectUser(name string, args ConnectUserArgs) *ConnectUser {
	return &ConnectUser{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ConnectUser)(nil)

// ConnectUser represents the Terraform resource aws_connect_user.
type ConnectUser struct {
	Name      string
	Args      ConnectUserArgs
	state     *connectUserState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ConnectUser].
func (cu *ConnectUser) Type() string {
	return "aws_connect_user"
}

// LocalName returns the local name for [ConnectUser].
func (cu *ConnectUser) LocalName() string {
	return cu.Name
}

// Configuration returns the configuration (args) for [ConnectUser].
func (cu *ConnectUser) Configuration() interface{} {
	return cu.Args
}

// DependOn is used for other resources to depend on [ConnectUser].
func (cu *ConnectUser) DependOn() terra.Reference {
	return terra.ReferenceResource(cu)
}

// Dependencies returns the list of resources [ConnectUser] depends_on.
func (cu *ConnectUser) Dependencies() terra.Dependencies {
	return cu.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ConnectUser].
func (cu *ConnectUser) LifecycleManagement() *terra.Lifecycle {
	return cu.Lifecycle
}

// Attributes returns the attributes for [ConnectUser].
func (cu *ConnectUser) Attributes() connectUserAttributes {
	return connectUserAttributes{ref: terra.ReferenceResource(cu)}
}

// ImportState imports the given attribute values into [ConnectUser]'s state.
func (cu *ConnectUser) ImportState(av io.Reader) error {
	cu.state = &connectUserState{}
	if err := json.NewDecoder(av).Decode(cu.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cu.Type(), cu.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ConnectUser] has state.
func (cu *ConnectUser) State() (*connectUserState, bool) {
	return cu.state, cu.state != nil
}

// StateMust returns the state for [ConnectUser]. Panics if the state is nil.
func (cu *ConnectUser) StateMust() *connectUserState {
	if cu.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cu.Type(), cu.LocalName()))
	}
	return cu.state
}

// ConnectUserArgs contains the configurations for aws_connect_user.
type ConnectUserArgs struct {
	// DirectoryUserId: string, optional
	DirectoryUserId terra.StringValue `hcl:"directory_user_id,attr"`
	// HierarchyGroupId: string, optional
	HierarchyGroupId terra.StringValue `hcl:"hierarchy_group_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceId: string, required
	InstanceId terra.StringValue `hcl:"instance_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Password: string, optional
	Password terra.StringValue `hcl:"password,attr"`
	// RoutingProfileId: string, required
	RoutingProfileId terra.StringValue `hcl:"routing_profile_id,attr" validate:"required"`
	// SecurityProfileIds: set of string, required
	SecurityProfileIds terra.SetValue[terra.StringValue] `hcl:"security_profile_ids,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// IdentityInfo: optional
	IdentityInfo *connectuser.IdentityInfo `hcl:"identity_info,block"`
	// PhoneConfig: required
	PhoneConfig *connectuser.PhoneConfig `hcl:"phone_config,block" validate:"required"`
}
type connectUserAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_connect_user.
func (cu connectUserAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cu.ref.Append("arn"))
}

// DirectoryUserId returns a reference to field directory_user_id of aws_connect_user.
func (cu connectUserAttributes) DirectoryUserId() terra.StringValue {
	return terra.ReferenceAsString(cu.ref.Append("directory_user_id"))
}

// HierarchyGroupId returns a reference to field hierarchy_group_id of aws_connect_user.
func (cu connectUserAttributes) HierarchyGroupId() terra.StringValue {
	return terra.ReferenceAsString(cu.ref.Append("hierarchy_group_id"))
}

// Id returns a reference to field id of aws_connect_user.
func (cu connectUserAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cu.ref.Append("id"))
}

// InstanceId returns a reference to field instance_id of aws_connect_user.
func (cu connectUserAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(cu.ref.Append("instance_id"))
}

// Name returns a reference to field name of aws_connect_user.
func (cu connectUserAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cu.ref.Append("name"))
}

// Password returns a reference to field password of aws_connect_user.
func (cu connectUserAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(cu.ref.Append("password"))
}

// RoutingProfileId returns a reference to field routing_profile_id of aws_connect_user.
func (cu connectUserAttributes) RoutingProfileId() terra.StringValue {
	return terra.ReferenceAsString(cu.ref.Append("routing_profile_id"))
}

// SecurityProfileIds returns a reference to field security_profile_ids of aws_connect_user.
func (cu connectUserAttributes) SecurityProfileIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cu.ref.Append("security_profile_ids"))
}

// Tags returns a reference to field tags of aws_connect_user.
func (cu connectUserAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cu.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_connect_user.
func (cu connectUserAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cu.ref.Append("tags_all"))
}

// UserId returns a reference to field user_id of aws_connect_user.
func (cu connectUserAttributes) UserId() terra.StringValue {
	return terra.ReferenceAsString(cu.ref.Append("user_id"))
}

func (cu connectUserAttributes) IdentityInfo() terra.ListValue[connectuser.IdentityInfoAttributes] {
	return terra.ReferenceAsList[connectuser.IdentityInfoAttributes](cu.ref.Append("identity_info"))
}

func (cu connectUserAttributes) PhoneConfig() terra.ListValue[connectuser.PhoneConfigAttributes] {
	return terra.ReferenceAsList[connectuser.PhoneConfigAttributes](cu.ref.Append("phone_config"))
}

type connectUserState struct {
	Arn                string                          `json:"arn"`
	DirectoryUserId    string                          `json:"directory_user_id"`
	HierarchyGroupId   string                          `json:"hierarchy_group_id"`
	Id                 string                          `json:"id"`
	InstanceId         string                          `json:"instance_id"`
	Name               string                          `json:"name"`
	Password           string                          `json:"password"`
	RoutingProfileId   string                          `json:"routing_profile_id"`
	SecurityProfileIds []string                        `json:"security_profile_ids"`
	Tags               map[string]string               `json:"tags"`
	TagsAll            map[string]string               `json:"tags_all"`
	UserId             string                          `json:"user_id"`
	IdentityInfo       []connectuser.IdentityInfoState `json:"identity_info"`
	PhoneConfig        []connectuser.PhoneConfigState  `json:"phone_config"`
}
