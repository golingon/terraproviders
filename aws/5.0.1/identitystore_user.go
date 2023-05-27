// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	identitystoreuser "github.com/golingon/terraproviders/aws/5.0.1/identitystoreuser"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIdentitystoreUser creates a new instance of [IdentitystoreUser].
func NewIdentitystoreUser(name string, args IdentitystoreUserArgs) *IdentitystoreUser {
	return &IdentitystoreUser{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IdentitystoreUser)(nil)

// IdentitystoreUser represents the Terraform resource aws_identitystore_user.
type IdentitystoreUser struct {
	Name      string
	Args      IdentitystoreUserArgs
	state     *identitystoreUserState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IdentitystoreUser].
func (iu *IdentitystoreUser) Type() string {
	return "aws_identitystore_user"
}

// LocalName returns the local name for [IdentitystoreUser].
func (iu *IdentitystoreUser) LocalName() string {
	return iu.Name
}

// Configuration returns the configuration (args) for [IdentitystoreUser].
func (iu *IdentitystoreUser) Configuration() interface{} {
	return iu.Args
}

// DependOn is used for other resources to depend on [IdentitystoreUser].
func (iu *IdentitystoreUser) DependOn() terra.Reference {
	return terra.ReferenceResource(iu)
}

// Dependencies returns the list of resources [IdentitystoreUser] depends_on.
func (iu *IdentitystoreUser) Dependencies() terra.Dependencies {
	return iu.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IdentitystoreUser].
func (iu *IdentitystoreUser) LifecycleManagement() *terra.Lifecycle {
	return iu.Lifecycle
}

// Attributes returns the attributes for [IdentitystoreUser].
func (iu *IdentitystoreUser) Attributes() identitystoreUserAttributes {
	return identitystoreUserAttributes{ref: terra.ReferenceResource(iu)}
}

// ImportState imports the given attribute values into [IdentitystoreUser]'s state.
func (iu *IdentitystoreUser) ImportState(av io.Reader) error {
	iu.state = &identitystoreUserState{}
	if err := json.NewDecoder(av).Decode(iu.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iu.Type(), iu.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IdentitystoreUser] has state.
func (iu *IdentitystoreUser) State() (*identitystoreUserState, bool) {
	return iu.state, iu.state != nil
}

// StateMust returns the state for [IdentitystoreUser]. Panics if the state is nil.
func (iu *IdentitystoreUser) StateMust() *identitystoreUserState {
	if iu.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iu.Type(), iu.LocalName()))
	}
	return iu.state
}

// IdentitystoreUserArgs contains the configurations for aws_identitystore_user.
type IdentitystoreUserArgs struct {
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IdentityStoreId: string, required
	IdentityStoreId terra.StringValue `hcl:"identity_store_id,attr" validate:"required"`
	// Locale: string, optional
	Locale terra.StringValue `hcl:"locale,attr"`
	// Nickname: string, optional
	Nickname terra.StringValue `hcl:"nickname,attr"`
	// PreferredLanguage: string, optional
	PreferredLanguage terra.StringValue `hcl:"preferred_language,attr"`
	// ProfileUrl: string, optional
	ProfileUrl terra.StringValue `hcl:"profile_url,attr"`
	// Timezone: string, optional
	Timezone terra.StringValue `hcl:"timezone,attr"`
	// Title: string, optional
	Title terra.StringValue `hcl:"title,attr"`
	// UserName: string, required
	UserName terra.StringValue `hcl:"user_name,attr" validate:"required"`
	// UserType: string, optional
	UserType terra.StringValue `hcl:"user_type,attr"`
	// ExternalIds: min=0
	ExternalIds []identitystoreuser.ExternalIds `hcl:"external_ids,block" validate:"min=0"`
	// Addresses: optional
	Addresses *identitystoreuser.Addresses `hcl:"addresses,block"`
	// Emails: optional
	Emails *identitystoreuser.Emails `hcl:"emails,block"`
	// Name: required
	Name *identitystoreuser.Name `hcl:"name,block" validate:"required"`
	// PhoneNumbers: optional
	PhoneNumbers *identitystoreuser.PhoneNumbers `hcl:"phone_numbers,block"`
}
type identitystoreUserAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of aws_identitystore_user.
func (iu identitystoreUserAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(iu.ref.Append("display_name"))
}

// Id returns a reference to field id of aws_identitystore_user.
func (iu identitystoreUserAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iu.ref.Append("id"))
}

// IdentityStoreId returns a reference to field identity_store_id of aws_identitystore_user.
func (iu identitystoreUserAttributes) IdentityStoreId() terra.StringValue {
	return terra.ReferenceAsString(iu.ref.Append("identity_store_id"))
}

// Locale returns a reference to field locale of aws_identitystore_user.
func (iu identitystoreUserAttributes) Locale() terra.StringValue {
	return terra.ReferenceAsString(iu.ref.Append("locale"))
}

// Nickname returns a reference to field nickname of aws_identitystore_user.
func (iu identitystoreUserAttributes) Nickname() terra.StringValue {
	return terra.ReferenceAsString(iu.ref.Append("nickname"))
}

// PreferredLanguage returns a reference to field preferred_language of aws_identitystore_user.
func (iu identitystoreUserAttributes) PreferredLanguage() terra.StringValue {
	return terra.ReferenceAsString(iu.ref.Append("preferred_language"))
}

// ProfileUrl returns a reference to field profile_url of aws_identitystore_user.
func (iu identitystoreUserAttributes) ProfileUrl() terra.StringValue {
	return terra.ReferenceAsString(iu.ref.Append("profile_url"))
}

// Timezone returns a reference to field timezone of aws_identitystore_user.
func (iu identitystoreUserAttributes) Timezone() terra.StringValue {
	return terra.ReferenceAsString(iu.ref.Append("timezone"))
}

// Title returns a reference to field title of aws_identitystore_user.
func (iu identitystoreUserAttributes) Title() terra.StringValue {
	return terra.ReferenceAsString(iu.ref.Append("title"))
}

// UserId returns a reference to field user_id of aws_identitystore_user.
func (iu identitystoreUserAttributes) UserId() terra.StringValue {
	return terra.ReferenceAsString(iu.ref.Append("user_id"))
}

// UserName returns a reference to field user_name of aws_identitystore_user.
func (iu identitystoreUserAttributes) UserName() terra.StringValue {
	return terra.ReferenceAsString(iu.ref.Append("user_name"))
}

// UserType returns a reference to field user_type of aws_identitystore_user.
func (iu identitystoreUserAttributes) UserType() terra.StringValue {
	return terra.ReferenceAsString(iu.ref.Append("user_type"))
}

func (iu identitystoreUserAttributes) ExternalIds() terra.ListValue[identitystoreuser.ExternalIdsAttributes] {
	return terra.ReferenceAsList[identitystoreuser.ExternalIdsAttributes](iu.ref.Append("external_ids"))
}

func (iu identitystoreUserAttributes) Addresses() terra.ListValue[identitystoreuser.AddressesAttributes] {
	return terra.ReferenceAsList[identitystoreuser.AddressesAttributes](iu.ref.Append("addresses"))
}

func (iu identitystoreUserAttributes) Emails() terra.ListValue[identitystoreuser.EmailsAttributes] {
	return terra.ReferenceAsList[identitystoreuser.EmailsAttributes](iu.ref.Append("emails"))
}

func (iu identitystoreUserAttributes) Name() terra.ListValue[identitystoreuser.NameAttributes] {
	return terra.ReferenceAsList[identitystoreuser.NameAttributes](iu.ref.Append("name"))
}

func (iu identitystoreUserAttributes) PhoneNumbers() terra.ListValue[identitystoreuser.PhoneNumbersAttributes] {
	return terra.ReferenceAsList[identitystoreuser.PhoneNumbersAttributes](iu.ref.Append("phone_numbers"))
}

type identitystoreUserState struct {
	DisplayName       string                                `json:"display_name"`
	Id                string                                `json:"id"`
	IdentityStoreId   string                                `json:"identity_store_id"`
	Locale            string                                `json:"locale"`
	Nickname          string                                `json:"nickname"`
	PreferredLanguage string                                `json:"preferred_language"`
	ProfileUrl        string                                `json:"profile_url"`
	Timezone          string                                `json:"timezone"`
	Title             string                                `json:"title"`
	UserId            string                                `json:"user_id"`
	UserName          string                                `json:"user_name"`
	UserType          string                                `json:"user_type"`
	ExternalIds       []identitystoreuser.ExternalIdsState  `json:"external_ids"`
	Addresses         []identitystoreuser.AddressesState    `json:"addresses"`
	Emails            []identitystoreuser.EmailsState       `json:"emails"`
	Name              []identitystoreuser.NameState         `json:"name"`
	PhoneNumbers      []identitystoreuser.PhoneNumbersState `json:"phone_numbers"`
}
