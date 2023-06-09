// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googleworkspace

import (
	"encoding/json"
	"fmt"
	groupsettings "github.com/golingon/terraproviders/googleworkspace/0.7.0/groupsettings"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGroupSettings creates a new instance of [GroupSettings].
func NewGroupSettings(name string, args GroupSettingsArgs) *GroupSettings {
	return &GroupSettings{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GroupSettings)(nil)

// GroupSettings represents the Terraform resource googleworkspace_group_settings.
type GroupSettings struct {
	Name      string
	Args      GroupSettingsArgs
	state     *groupSettingsState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GroupSettings].
func (gs *GroupSettings) Type() string {
	return "googleworkspace_group_settings"
}

// LocalName returns the local name for [GroupSettings].
func (gs *GroupSettings) LocalName() string {
	return gs.Name
}

// Configuration returns the configuration (args) for [GroupSettings].
func (gs *GroupSettings) Configuration() interface{} {
	return gs.Args
}

// DependOn is used for other resources to depend on [GroupSettings].
func (gs *GroupSettings) DependOn() terra.Reference {
	return terra.ReferenceResource(gs)
}

// Dependencies returns the list of resources [GroupSettings] depends_on.
func (gs *GroupSettings) Dependencies() terra.Dependencies {
	return gs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GroupSettings].
func (gs *GroupSettings) LifecycleManagement() *terra.Lifecycle {
	return gs.Lifecycle
}

// Attributes returns the attributes for [GroupSettings].
func (gs *GroupSettings) Attributes() groupSettingsAttributes {
	return groupSettingsAttributes{ref: terra.ReferenceResource(gs)}
}

// ImportState imports the given attribute values into [GroupSettings]'s state.
func (gs *GroupSettings) ImportState(av io.Reader) error {
	gs.state = &groupSettingsState{}
	if err := json.NewDecoder(av).Decode(gs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gs.Type(), gs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GroupSettings] has state.
func (gs *GroupSettings) State() (*groupSettingsState, bool) {
	return gs.state, gs.state != nil
}

// StateMust returns the state for [GroupSettings]. Panics if the state is nil.
func (gs *GroupSettings) StateMust() *groupSettingsState {
	if gs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gs.Type(), gs.LocalName()))
	}
	return gs.state
}

// GroupSettingsArgs contains the configurations for googleworkspace_group_settings.
type GroupSettingsArgs struct {
	// AllowExternalMembers: bool, optional
	AllowExternalMembers terra.BoolValue `hcl:"allow_external_members,attr"`
	// AllowWebPosting: bool, optional
	AllowWebPosting terra.BoolValue `hcl:"allow_web_posting,attr"`
	// ArchiveOnly: bool, optional
	ArchiveOnly terra.BoolValue `hcl:"archive_only,attr"`
	// CustomFooterText: string, optional
	CustomFooterText terra.StringValue `hcl:"custom_footer_text,attr"`
	// CustomReplyTo: string, optional
	CustomReplyTo terra.StringValue `hcl:"custom_reply_to,attr"`
	// DefaultMessageDenyNotificationText: string, optional
	DefaultMessageDenyNotificationText terra.StringValue `hcl:"default_message_deny_notification_text,attr"`
	// Email: string, required
	Email terra.StringValue `hcl:"email,attr" validate:"required"`
	// EnableCollaborativeInbox: bool, optional
	EnableCollaborativeInbox terra.BoolValue `hcl:"enable_collaborative_inbox,attr"`
	// IncludeCustomFooter: bool, optional
	IncludeCustomFooter terra.BoolValue `hcl:"include_custom_footer,attr"`
	// IncludeInGlobalAddressList: bool, optional
	IncludeInGlobalAddressList terra.BoolValue `hcl:"include_in_global_address_list,attr"`
	// IsArchived: bool, optional
	IsArchived terra.BoolValue `hcl:"is_archived,attr"`
	// MembersCanPostAsTheGroup: bool, optional
	MembersCanPostAsTheGroup terra.BoolValue `hcl:"members_can_post_as_the_group,attr"`
	// MessageModerationLevel: string, optional
	MessageModerationLevel terra.StringValue `hcl:"message_moderation_level,attr"`
	// PrimaryLanguage: string, optional
	PrimaryLanguage terra.StringValue `hcl:"primary_language,attr"`
	// ReplyTo: string, optional
	ReplyTo terra.StringValue `hcl:"reply_to,attr"`
	// SendMessageDenyNotification: bool, optional
	SendMessageDenyNotification terra.BoolValue `hcl:"send_message_deny_notification,attr"`
	// SpamModerationLevel: string, optional
	SpamModerationLevel terra.StringValue `hcl:"spam_moderation_level,attr"`
	// WhoCanAssistContent: string, optional
	WhoCanAssistContent terra.StringValue `hcl:"who_can_assist_content,attr"`
	// WhoCanContactOwner: string, optional
	WhoCanContactOwner terra.StringValue `hcl:"who_can_contact_owner,attr"`
	// WhoCanDiscoverGroup: string, optional
	WhoCanDiscoverGroup terra.StringValue `hcl:"who_can_discover_group,attr"`
	// WhoCanJoin: string, optional
	WhoCanJoin terra.StringValue `hcl:"who_can_join,attr"`
	// WhoCanLeaveGroup: string, optional
	WhoCanLeaveGroup terra.StringValue `hcl:"who_can_leave_group,attr"`
	// WhoCanModerateContent: string, optional
	WhoCanModerateContent terra.StringValue `hcl:"who_can_moderate_content,attr"`
	// WhoCanModerateMembers: string, optional
	WhoCanModerateMembers terra.StringValue `hcl:"who_can_moderate_members,attr"`
	// WhoCanPostMessage: string, optional
	WhoCanPostMessage terra.StringValue `hcl:"who_can_post_message,attr"`
	// WhoCanViewGroup: string, optional
	WhoCanViewGroup terra.StringValue `hcl:"who_can_view_group,attr"`
	// WhoCanViewMembership: string, optional
	WhoCanViewMembership terra.StringValue `hcl:"who_can_view_membership,attr"`
	// Timeouts: optional
	Timeouts *groupsettings.Timeouts `hcl:"timeouts,block"`
}
type groupSettingsAttributes struct {
	ref terra.Reference
}

// AllowExternalMembers returns a reference to field allow_external_members of googleworkspace_group_settings.
func (gs groupSettingsAttributes) AllowExternalMembers() terra.BoolValue {
	return terra.ReferenceAsBool(gs.ref.Append("allow_external_members"))
}

// AllowWebPosting returns a reference to field allow_web_posting of googleworkspace_group_settings.
func (gs groupSettingsAttributes) AllowWebPosting() terra.BoolValue {
	return terra.ReferenceAsBool(gs.ref.Append("allow_web_posting"))
}

// ArchiveOnly returns a reference to field archive_only of googleworkspace_group_settings.
func (gs groupSettingsAttributes) ArchiveOnly() terra.BoolValue {
	return terra.ReferenceAsBool(gs.ref.Append("archive_only"))
}

// CustomFooterText returns a reference to field custom_footer_text of googleworkspace_group_settings.
func (gs groupSettingsAttributes) CustomFooterText() terra.StringValue {
	return terra.ReferenceAsString(gs.ref.Append("custom_footer_text"))
}

// CustomReplyTo returns a reference to field custom_reply_to of googleworkspace_group_settings.
func (gs groupSettingsAttributes) CustomReplyTo() terra.StringValue {
	return terra.ReferenceAsString(gs.ref.Append("custom_reply_to"))
}

// CustomRolesEnabledForSettingsToBeMerged returns a reference to field custom_roles_enabled_for_settings_to_be_merged of googleworkspace_group_settings.
func (gs groupSettingsAttributes) CustomRolesEnabledForSettingsToBeMerged() terra.BoolValue {
	return terra.ReferenceAsBool(gs.ref.Append("custom_roles_enabled_for_settings_to_be_merged"))
}

// DefaultMessageDenyNotificationText returns a reference to field default_message_deny_notification_text of googleworkspace_group_settings.
func (gs groupSettingsAttributes) DefaultMessageDenyNotificationText() terra.StringValue {
	return terra.ReferenceAsString(gs.ref.Append("default_message_deny_notification_text"))
}

// Description returns a reference to field description of googleworkspace_group_settings.
func (gs groupSettingsAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gs.ref.Append("description"))
}

// Email returns a reference to field email of googleworkspace_group_settings.
func (gs groupSettingsAttributes) Email() terra.StringValue {
	return terra.ReferenceAsString(gs.ref.Append("email"))
}

// EnableCollaborativeInbox returns a reference to field enable_collaborative_inbox of googleworkspace_group_settings.
func (gs groupSettingsAttributes) EnableCollaborativeInbox() terra.BoolValue {
	return terra.ReferenceAsBool(gs.ref.Append("enable_collaborative_inbox"))
}

// Id returns a reference to field id of googleworkspace_group_settings.
func (gs groupSettingsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gs.ref.Append("id"))
}

// IncludeCustomFooter returns a reference to field include_custom_footer of googleworkspace_group_settings.
func (gs groupSettingsAttributes) IncludeCustomFooter() terra.BoolValue {
	return terra.ReferenceAsBool(gs.ref.Append("include_custom_footer"))
}

// IncludeInGlobalAddressList returns a reference to field include_in_global_address_list of googleworkspace_group_settings.
func (gs groupSettingsAttributes) IncludeInGlobalAddressList() terra.BoolValue {
	return terra.ReferenceAsBool(gs.ref.Append("include_in_global_address_list"))
}

// IsArchived returns a reference to field is_archived of googleworkspace_group_settings.
func (gs groupSettingsAttributes) IsArchived() terra.BoolValue {
	return terra.ReferenceAsBool(gs.ref.Append("is_archived"))
}

// MembersCanPostAsTheGroup returns a reference to field members_can_post_as_the_group of googleworkspace_group_settings.
func (gs groupSettingsAttributes) MembersCanPostAsTheGroup() terra.BoolValue {
	return terra.ReferenceAsBool(gs.ref.Append("members_can_post_as_the_group"))
}

// MessageModerationLevel returns a reference to field message_moderation_level of googleworkspace_group_settings.
func (gs groupSettingsAttributes) MessageModerationLevel() terra.StringValue {
	return terra.ReferenceAsString(gs.ref.Append("message_moderation_level"))
}

// Name returns a reference to field name of googleworkspace_group_settings.
func (gs groupSettingsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gs.ref.Append("name"))
}

// PrimaryLanguage returns a reference to field primary_language of googleworkspace_group_settings.
func (gs groupSettingsAttributes) PrimaryLanguage() terra.StringValue {
	return terra.ReferenceAsString(gs.ref.Append("primary_language"))
}

// ReplyTo returns a reference to field reply_to of googleworkspace_group_settings.
func (gs groupSettingsAttributes) ReplyTo() terra.StringValue {
	return terra.ReferenceAsString(gs.ref.Append("reply_to"))
}

// SendMessageDenyNotification returns a reference to field send_message_deny_notification of googleworkspace_group_settings.
func (gs groupSettingsAttributes) SendMessageDenyNotification() terra.BoolValue {
	return terra.ReferenceAsBool(gs.ref.Append("send_message_deny_notification"))
}

// SpamModerationLevel returns a reference to field spam_moderation_level of googleworkspace_group_settings.
func (gs groupSettingsAttributes) SpamModerationLevel() terra.StringValue {
	return terra.ReferenceAsString(gs.ref.Append("spam_moderation_level"))
}

// WhoCanAssistContent returns a reference to field who_can_assist_content of googleworkspace_group_settings.
func (gs groupSettingsAttributes) WhoCanAssistContent() terra.StringValue {
	return terra.ReferenceAsString(gs.ref.Append("who_can_assist_content"))
}

// WhoCanContactOwner returns a reference to field who_can_contact_owner of googleworkspace_group_settings.
func (gs groupSettingsAttributes) WhoCanContactOwner() terra.StringValue {
	return terra.ReferenceAsString(gs.ref.Append("who_can_contact_owner"))
}

// WhoCanDiscoverGroup returns a reference to field who_can_discover_group of googleworkspace_group_settings.
func (gs groupSettingsAttributes) WhoCanDiscoverGroup() terra.StringValue {
	return terra.ReferenceAsString(gs.ref.Append("who_can_discover_group"))
}

// WhoCanJoin returns a reference to field who_can_join of googleworkspace_group_settings.
func (gs groupSettingsAttributes) WhoCanJoin() terra.StringValue {
	return terra.ReferenceAsString(gs.ref.Append("who_can_join"))
}

// WhoCanLeaveGroup returns a reference to field who_can_leave_group of googleworkspace_group_settings.
func (gs groupSettingsAttributes) WhoCanLeaveGroup() terra.StringValue {
	return terra.ReferenceAsString(gs.ref.Append("who_can_leave_group"))
}

// WhoCanModerateContent returns a reference to field who_can_moderate_content of googleworkspace_group_settings.
func (gs groupSettingsAttributes) WhoCanModerateContent() terra.StringValue {
	return terra.ReferenceAsString(gs.ref.Append("who_can_moderate_content"))
}

// WhoCanModerateMembers returns a reference to field who_can_moderate_members of googleworkspace_group_settings.
func (gs groupSettingsAttributes) WhoCanModerateMembers() terra.StringValue {
	return terra.ReferenceAsString(gs.ref.Append("who_can_moderate_members"))
}

// WhoCanPostMessage returns a reference to field who_can_post_message of googleworkspace_group_settings.
func (gs groupSettingsAttributes) WhoCanPostMessage() terra.StringValue {
	return terra.ReferenceAsString(gs.ref.Append("who_can_post_message"))
}

// WhoCanViewGroup returns a reference to field who_can_view_group of googleworkspace_group_settings.
func (gs groupSettingsAttributes) WhoCanViewGroup() terra.StringValue {
	return terra.ReferenceAsString(gs.ref.Append("who_can_view_group"))
}

// WhoCanViewMembership returns a reference to field who_can_view_membership of googleworkspace_group_settings.
func (gs groupSettingsAttributes) WhoCanViewMembership() terra.StringValue {
	return terra.ReferenceAsString(gs.ref.Append("who_can_view_membership"))
}

func (gs groupSettingsAttributes) Timeouts() groupsettings.TimeoutsAttributes {
	return terra.ReferenceAsSingle[groupsettings.TimeoutsAttributes](gs.ref.Append("timeouts"))
}

type groupSettingsState struct {
	AllowExternalMembers                    bool                         `json:"allow_external_members"`
	AllowWebPosting                         bool                         `json:"allow_web_posting"`
	ArchiveOnly                             bool                         `json:"archive_only"`
	CustomFooterText                        string                       `json:"custom_footer_text"`
	CustomReplyTo                           string                       `json:"custom_reply_to"`
	CustomRolesEnabledForSettingsToBeMerged bool                         `json:"custom_roles_enabled_for_settings_to_be_merged"`
	DefaultMessageDenyNotificationText      string                       `json:"default_message_deny_notification_text"`
	Description                             string                       `json:"description"`
	Email                                   string                       `json:"email"`
	EnableCollaborativeInbox                bool                         `json:"enable_collaborative_inbox"`
	Id                                      string                       `json:"id"`
	IncludeCustomFooter                     bool                         `json:"include_custom_footer"`
	IncludeInGlobalAddressList              bool                         `json:"include_in_global_address_list"`
	IsArchived                              bool                         `json:"is_archived"`
	MembersCanPostAsTheGroup                bool                         `json:"members_can_post_as_the_group"`
	MessageModerationLevel                  string                       `json:"message_moderation_level"`
	Name                                    string                       `json:"name"`
	PrimaryLanguage                         string                       `json:"primary_language"`
	ReplyTo                                 string                       `json:"reply_to"`
	SendMessageDenyNotification             bool                         `json:"send_message_deny_notification"`
	SpamModerationLevel                     string                       `json:"spam_moderation_level"`
	WhoCanAssistContent                     string                       `json:"who_can_assist_content"`
	WhoCanContactOwner                      string                       `json:"who_can_contact_owner"`
	WhoCanDiscoverGroup                     string                       `json:"who_can_discover_group"`
	WhoCanJoin                              string                       `json:"who_can_join"`
	WhoCanLeaveGroup                        string                       `json:"who_can_leave_group"`
	WhoCanModerateContent                   string                       `json:"who_can_moderate_content"`
	WhoCanModerateMembers                   string                       `json:"who_can_moderate_members"`
	WhoCanPostMessage                       string                       `json:"who_can_post_message"`
	WhoCanViewGroup                         string                       `json:"who_can_view_group"`
	WhoCanViewMembership                    string                       `json:"who_can_view_membership"`
	Timeouts                                *groupsettings.TimeoutsState `json:"timeouts"`
}
