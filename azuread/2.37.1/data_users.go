// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azuread

import (
	datausers "github.com/golingon/terraproviders/azuread/2.37.1/datausers"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataUsers creates a new instance of [DataUsers].
func NewDataUsers(name string, args DataUsersArgs) *DataUsers {
	return &DataUsers{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataUsers)(nil)

// DataUsers represents the Terraform data resource azuread_users.
type DataUsers struct {
	Name string
	Args DataUsersArgs
}

// DataSource returns the Terraform object type for [DataUsers].
func (u *DataUsers) DataSource() string {
	return "azuread_users"
}

// LocalName returns the local name for [DataUsers].
func (u *DataUsers) LocalName() string {
	return u.Name
}

// Configuration returns the configuration (args) for [DataUsers].
func (u *DataUsers) Configuration() interface{} {
	return u.Args
}

// Attributes returns the attributes for [DataUsers].
func (u *DataUsers) Attributes() dataUsersAttributes {
	return dataUsersAttributes{ref: terra.ReferenceDataResource(u)}
}

// DataUsersArgs contains the configurations for azuread_users.
type DataUsersArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IgnoreMissing: bool, optional
	IgnoreMissing terra.BoolValue `hcl:"ignore_missing,attr"`
	// MailNicknames: list of string, optional
	MailNicknames terra.ListValue[terra.StringValue] `hcl:"mail_nicknames,attr"`
	// ObjectIds: list of string, optional
	ObjectIds terra.ListValue[terra.StringValue] `hcl:"object_ids,attr"`
	// ReturnAll: bool, optional
	ReturnAll terra.BoolValue `hcl:"return_all,attr"`
	// UserPrincipalNames: list of string, optional
	UserPrincipalNames terra.ListValue[terra.StringValue] `hcl:"user_principal_names,attr"`
	// Users: min=0
	Users []datausers.Users `hcl:"users,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datausers.Timeouts `hcl:"timeouts,block"`
}
type dataUsersAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azuread_users.
func (u dataUsersAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(u.ref.Append("id"))
}

// IgnoreMissing returns a reference to field ignore_missing of azuread_users.
func (u dataUsersAttributes) IgnoreMissing() terra.BoolValue {
	return terra.ReferenceAsBool(u.ref.Append("ignore_missing"))
}

// MailNicknames returns a reference to field mail_nicknames of azuread_users.
func (u dataUsersAttributes) MailNicknames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](u.ref.Append("mail_nicknames"))
}

// ObjectIds returns a reference to field object_ids of azuread_users.
func (u dataUsersAttributes) ObjectIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](u.ref.Append("object_ids"))
}

// ReturnAll returns a reference to field return_all of azuread_users.
func (u dataUsersAttributes) ReturnAll() terra.BoolValue {
	return terra.ReferenceAsBool(u.ref.Append("return_all"))
}

// UserPrincipalNames returns a reference to field user_principal_names of azuread_users.
func (u dataUsersAttributes) UserPrincipalNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](u.ref.Append("user_principal_names"))
}

func (u dataUsersAttributes) Users() terra.ListValue[datausers.UsersAttributes] {
	return terra.ReferenceAsList[datausers.UsersAttributes](u.ref.Append("users"))
}

func (u dataUsersAttributes) Timeouts() datausers.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datausers.TimeoutsAttributes](u.ref.Append("timeouts"))
}
