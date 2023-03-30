// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	sagemakeruserprofile "github.com/golingon/terraproviders/aws/4.60.0/sagemakeruserprofile"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewSagemakerUserProfile(name string, args SagemakerUserProfileArgs) *SagemakerUserProfile {
	return &SagemakerUserProfile{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SagemakerUserProfile)(nil)

type SagemakerUserProfile struct {
	Name  string
	Args  SagemakerUserProfileArgs
	state *sagemakerUserProfileState
}

func (sup *SagemakerUserProfile) Type() string {
	return "aws_sagemaker_user_profile"
}

func (sup *SagemakerUserProfile) LocalName() string {
	return sup.Name
}

func (sup *SagemakerUserProfile) Configuration() interface{} {
	return sup.Args
}

func (sup *SagemakerUserProfile) Attributes() sagemakerUserProfileAttributes {
	return sagemakerUserProfileAttributes{ref: terra.ReferenceResource(sup)}
}

func (sup *SagemakerUserProfile) ImportState(av io.Reader) error {
	sup.state = &sagemakerUserProfileState{}
	if err := json.NewDecoder(av).Decode(sup.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sup.Type(), sup.LocalName(), err)
	}
	return nil
}

func (sup *SagemakerUserProfile) State() (*sagemakerUserProfileState, bool) {
	return sup.state, sup.state != nil
}

func (sup *SagemakerUserProfile) StateMust() *sagemakerUserProfileState {
	if sup.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sup.Type(), sup.LocalName()))
	}
	return sup.state
}

func (sup *SagemakerUserProfile) DependOn() terra.Reference {
	return terra.ReferenceResource(sup)
}

type SagemakerUserProfileArgs struct {
	// DomainId: string, required
	DomainId terra.StringValue `hcl:"domain_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SingleSignOnUserIdentifier: string, optional
	SingleSignOnUserIdentifier terra.StringValue `hcl:"single_sign_on_user_identifier,attr"`
	// SingleSignOnUserValue: string, optional
	SingleSignOnUserValue terra.StringValue `hcl:"single_sign_on_user_value,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// UserProfileName: string, required
	UserProfileName terra.StringValue `hcl:"user_profile_name,attr" validate:"required"`
	// UserSettings: optional
	UserSettings *sagemakeruserprofile.UserSettings `hcl:"user_settings,block"`
	// DependsOn contains resources that SagemakerUserProfile depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type sagemakerUserProfileAttributes struct {
	ref terra.Reference
}

func (sup sagemakerUserProfileAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(sup.ref.Append("arn"))
}

func (sup sagemakerUserProfileAttributes) DomainId() terra.StringValue {
	return terra.ReferenceString(sup.ref.Append("domain_id"))
}

func (sup sagemakerUserProfileAttributes) HomeEfsFileSystemUid() terra.StringValue {
	return terra.ReferenceString(sup.ref.Append("home_efs_file_system_uid"))
}

func (sup sagemakerUserProfileAttributes) Id() terra.StringValue {
	return terra.ReferenceString(sup.ref.Append("id"))
}

func (sup sagemakerUserProfileAttributes) SingleSignOnUserIdentifier() terra.StringValue {
	return terra.ReferenceString(sup.ref.Append("single_sign_on_user_identifier"))
}

func (sup sagemakerUserProfileAttributes) SingleSignOnUserValue() terra.StringValue {
	return terra.ReferenceString(sup.ref.Append("single_sign_on_user_value"))
}

func (sup sagemakerUserProfileAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](sup.ref.Append("tags"))
}

func (sup sagemakerUserProfileAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](sup.ref.Append("tags_all"))
}

func (sup sagemakerUserProfileAttributes) UserProfileName() terra.StringValue {
	return terra.ReferenceString(sup.ref.Append("user_profile_name"))
}

func (sup sagemakerUserProfileAttributes) UserSettings() terra.ListValue[sagemakeruserprofile.UserSettingsAttributes] {
	return terra.ReferenceList[sagemakeruserprofile.UserSettingsAttributes](sup.ref.Append("user_settings"))
}

type sagemakerUserProfileState struct {
	Arn                        string                                   `json:"arn"`
	DomainId                   string                                   `json:"domain_id"`
	HomeEfsFileSystemUid       string                                   `json:"home_efs_file_system_uid"`
	Id                         string                                   `json:"id"`
	SingleSignOnUserIdentifier string                                   `json:"single_sign_on_user_identifier"`
	SingleSignOnUserValue      string                                   `json:"single_sign_on_user_value"`
	Tags                       map[string]string                        `json:"tags"`
	TagsAll                    map[string]string                        `json:"tags_all"`
	UserProfileName            string                                   `json:"user_profile_name"`
	UserSettings               []sagemakeruserprofile.UserSettingsState `json:"user_settings"`
}
