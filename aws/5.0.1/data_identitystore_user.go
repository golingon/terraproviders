// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataidentitystoreuser "github.com/golingon/terraproviders/aws/5.0.1/dataidentitystoreuser"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataIdentitystoreUser creates a new instance of [DataIdentitystoreUser].
func NewDataIdentitystoreUser(name string, args DataIdentitystoreUserArgs) *DataIdentitystoreUser {
	return &DataIdentitystoreUser{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataIdentitystoreUser)(nil)

// DataIdentitystoreUser represents the Terraform data resource aws_identitystore_user.
type DataIdentitystoreUser struct {
	Name string
	Args DataIdentitystoreUserArgs
}

// DataSource returns the Terraform object type for [DataIdentitystoreUser].
func (iu *DataIdentitystoreUser) DataSource() string {
	return "aws_identitystore_user"
}

// LocalName returns the local name for [DataIdentitystoreUser].
func (iu *DataIdentitystoreUser) LocalName() string {
	return iu.Name
}

// Configuration returns the configuration (args) for [DataIdentitystoreUser].
func (iu *DataIdentitystoreUser) Configuration() interface{} {
	return iu.Args
}

// Attributes returns the attributes for [DataIdentitystoreUser].
func (iu *DataIdentitystoreUser) Attributes() dataIdentitystoreUserAttributes {
	return dataIdentitystoreUserAttributes{ref: terra.ReferenceDataResource(iu)}
}

// DataIdentitystoreUserArgs contains the configurations for aws_identitystore_user.
type DataIdentitystoreUserArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IdentityStoreId: string, required
	IdentityStoreId terra.StringValue `hcl:"identity_store_id,attr" validate:"required"`
	// UserId: string, optional
	UserId terra.StringValue `hcl:"user_id,attr"`
	// Addresses: min=0
	Addresses []dataidentitystoreuser.Addresses `hcl:"addresses,block" validate:"min=0"`
	// Emails: min=0
	Emails []dataidentitystoreuser.Emails `hcl:"emails,block" validate:"min=0"`
	// ExternalIds: min=0
	ExternalIds []dataidentitystoreuser.ExternalIds `hcl:"external_ids,block" validate:"min=0"`
	// Name: min=0
	Name []dataidentitystoreuser.Name `hcl:"name,block" validate:"min=0"`
	// PhoneNumbers: min=0
	PhoneNumbers []dataidentitystoreuser.PhoneNumbers `hcl:"phone_numbers,block" validate:"min=0"`
	// AlternateIdentifier: optional
	AlternateIdentifier *dataidentitystoreuser.AlternateIdentifier `hcl:"alternate_identifier,block"`
}
type dataIdentitystoreUserAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of aws_identitystore_user.
func (iu dataIdentitystoreUserAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(iu.ref.Append("display_name"))
}

// Id returns a reference to field id of aws_identitystore_user.
func (iu dataIdentitystoreUserAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iu.ref.Append("id"))
}

// IdentityStoreId returns a reference to field identity_store_id of aws_identitystore_user.
func (iu dataIdentitystoreUserAttributes) IdentityStoreId() terra.StringValue {
	return terra.ReferenceAsString(iu.ref.Append("identity_store_id"))
}

// Locale returns a reference to field locale of aws_identitystore_user.
func (iu dataIdentitystoreUserAttributes) Locale() terra.StringValue {
	return terra.ReferenceAsString(iu.ref.Append("locale"))
}

// Nickname returns a reference to field nickname of aws_identitystore_user.
func (iu dataIdentitystoreUserAttributes) Nickname() terra.StringValue {
	return terra.ReferenceAsString(iu.ref.Append("nickname"))
}

// PreferredLanguage returns a reference to field preferred_language of aws_identitystore_user.
func (iu dataIdentitystoreUserAttributes) PreferredLanguage() terra.StringValue {
	return terra.ReferenceAsString(iu.ref.Append("preferred_language"))
}

// ProfileUrl returns a reference to field profile_url of aws_identitystore_user.
func (iu dataIdentitystoreUserAttributes) ProfileUrl() terra.StringValue {
	return terra.ReferenceAsString(iu.ref.Append("profile_url"))
}

// Timezone returns a reference to field timezone of aws_identitystore_user.
func (iu dataIdentitystoreUserAttributes) Timezone() terra.StringValue {
	return terra.ReferenceAsString(iu.ref.Append("timezone"))
}

// Title returns a reference to field title of aws_identitystore_user.
func (iu dataIdentitystoreUserAttributes) Title() terra.StringValue {
	return terra.ReferenceAsString(iu.ref.Append("title"))
}

// UserId returns a reference to field user_id of aws_identitystore_user.
func (iu dataIdentitystoreUserAttributes) UserId() terra.StringValue {
	return terra.ReferenceAsString(iu.ref.Append("user_id"))
}

// UserName returns a reference to field user_name of aws_identitystore_user.
func (iu dataIdentitystoreUserAttributes) UserName() terra.StringValue {
	return terra.ReferenceAsString(iu.ref.Append("user_name"))
}

// UserType returns a reference to field user_type of aws_identitystore_user.
func (iu dataIdentitystoreUserAttributes) UserType() terra.StringValue {
	return terra.ReferenceAsString(iu.ref.Append("user_type"))
}

func (iu dataIdentitystoreUserAttributes) Addresses() terra.ListValue[dataidentitystoreuser.AddressesAttributes] {
	return terra.ReferenceAsList[dataidentitystoreuser.AddressesAttributes](iu.ref.Append("addresses"))
}

func (iu dataIdentitystoreUserAttributes) Emails() terra.ListValue[dataidentitystoreuser.EmailsAttributes] {
	return terra.ReferenceAsList[dataidentitystoreuser.EmailsAttributes](iu.ref.Append("emails"))
}

func (iu dataIdentitystoreUserAttributes) ExternalIds() terra.ListValue[dataidentitystoreuser.ExternalIdsAttributes] {
	return terra.ReferenceAsList[dataidentitystoreuser.ExternalIdsAttributes](iu.ref.Append("external_ids"))
}

func (iu dataIdentitystoreUserAttributes) Name() terra.ListValue[dataidentitystoreuser.NameAttributes] {
	return terra.ReferenceAsList[dataidentitystoreuser.NameAttributes](iu.ref.Append("name"))
}

func (iu dataIdentitystoreUserAttributes) PhoneNumbers() terra.ListValue[dataidentitystoreuser.PhoneNumbersAttributes] {
	return terra.ReferenceAsList[dataidentitystoreuser.PhoneNumbersAttributes](iu.ref.Append("phone_numbers"))
}

func (iu dataIdentitystoreUserAttributes) AlternateIdentifier() terra.ListValue[dataidentitystoreuser.AlternateIdentifierAttributes] {
	return terra.ReferenceAsList[dataidentitystoreuser.AlternateIdentifierAttributes](iu.ref.Append("alternate_identifier"))
}
