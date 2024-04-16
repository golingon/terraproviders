// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azuread_user

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource azuread_user.
type Resource struct {
	Name      string
	Args      Args
	state     *azureadUserState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (au *Resource) Type() string {
	return "azuread_user"
}

// LocalName returns the local name for [Resource].
func (au *Resource) LocalName() string {
	return au.Name
}

// Configuration returns the configuration (args) for [Resource].
func (au *Resource) Configuration() interface{} {
	return au.Args
}

// DependOn is used for other resources to depend on [Resource].
func (au *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(au)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (au *Resource) Dependencies() terra.Dependencies {
	return au.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (au *Resource) LifecycleManagement() *terra.Lifecycle {
	return au.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (au *Resource) Attributes() azureadUserAttributes {
	return azureadUserAttributes{ref: terra.ReferenceResource(au)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (au *Resource) ImportState(state io.Reader) error {
	au.state = &azureadUserState{}
	if err := json.NewDecoder(state).Decode(au.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", au.Type(), au.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (au *Resource) State() (*azureadUserState, bool) {
	return au.state, au.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (au *Resource) StateMust() *azureadUserState {
	if au.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", au.Type(), au.LocalName()))
	}
	return au.state
}

// Args contains the configurations for azuread_user.
type Args struct {
	// AccountEnabled: bool, optional
	AccountEnabled terra.BoolValue `hcl:"account_enabled,attr"`
	// AgeGroup: string, optional
	AgeGroup terra.StringValue `hcl:"age_group,attr"`
	// BusinessPhones: list of string, optional
	BusinessPhones terra.ListValue[terra.StringValue] `hcl:"business_phones,attr"`
	// City: string, optional
	City terra.StringValue `hcl:"city,attr"`
	// CompanyName: string, optional
	CompanyName terra.StringValue `hcl:"company_name,attr"`
	// ConsentProvidedForMinor: string, optional
	ConsentProvidedForMinor terra.StringValue `hcl:"consent_provided_for_minor,attr"`
	// CostCenter: string, optional
	CostCenter terra.StringValue `hcl:"cost_center,attr"`
	// Country: string, optional
	Country terra.StringValue `hcl:"country,attr"`
	// Department: string, optional
	Department terra.StringValue `hcl:"department,attr"`
	// DisablePasswordExpiration: bool, optional
	DisablePasswordExpiration terra.BoolValue `hcl:"disable_password_expiration,attr"`
	// DisableStrongPassword: bool, optional
	DisableStrongPassword terra.BoolValue `hcl:"disable_strong_password,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Division: string, optional
	Division terra.StringValue `hcl:"division,attr"`
	// EmployeeId: string, optional
	EmployeeId terra.StringValue `hcl:"employee_id,attr"`
	// EmployeeType: string, optional
	EmployeeType terra.StringValue `hcl:"employee_type,attr"`
	// FaxNumber: string, optional
	FaxNumber terra.StringValue `hcl:"fax_number,attr"`
	// ForcePasswordChange: bool, optional
	ForcePasswordChange terra.BoolValue `hcl:"force_password_change,attr"`
	// GivenName: string, optional
	GivenName terra.StringValue `hcl:"given_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// JobTitle: string, optional
	JobTitle terra.StringValue `hcl:"job_title,attr"`
	// Mail: string, optional
	Mail terra.StringValue `hcl:"mail,attr"`
	// MailNickname: string, optional
	MailNickname terra.StringValue `hcl:"mail_nickname,attr"`
	// ManagerId: string, optional
	ManagerId terra.StringValue `hcl:"manager_id,attr"`
	// MobilePhone: string, optional
	MobilePhone terra.StringValue `hcl:"mobile_phone,attr"`
	// OfficeLocation: string, optional
	OfficeLocation terra.StringValue `hcl:"office_location,attr"`
	// OnpremisesImmutableId: string, optional
	OnpremisesImmutableId terra.StringValue `hcl:"onpremises_immutable_id,attr"`
	// OtherMails: set of string, optional
	OtherMails terra.SetValue[terra.StringValue] `hcl:"other_mails,attr"`
	// Password: string, optional
	Password terra.StringValue `hcl:"password,attr"`
	// PostalCode: string, optional
	PostalCode terra.StringValue `hcl:"postal_code,attr"`
	// PreferredLanguage: string, optional
	PreferredLanguage terra.StringValue `hcl:"preferred_language,attr"`
	// ShowInAddressList: bool, optional
	ShowInAddressList terra.BoolValue `hcl:"show_in_address_list,attr"`
	// State: string, optional
	State terra.StringValue `hcl:"state,attr"`
	// StreetAddress: string, optional
	StreetAddress terra.StringValue `hcl:"street_address,attr"`
	// Surname: string, optional
	Surname terra.StringValue `hcl:"surname,attr"`
	// UsageLocation: string, optional
	UsageLocation terra.StringValue `hcl:"usage_location,attr"`
	// UserPrincipalName: string, required
	UserPrincipalName terra.StringValue `hcl:"user_principal_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azureadUserAttributes struct {
	ref terra.Reference
}

// AboutMe returns a reference to field about_me of azuread_user.
func (au azureadUserAttributes) AboutMe() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("about_me"))
}

// AccountEnabled returns a reference to field account_enabled of azuread_user.
func (au azureadUserAttributes) AccountEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(au.ref.Append("account_enabled"))
}

// AgeGroup returns a reference to field age_group of azuread_user.
func (au azureadUserAttributes) AgeGroup() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("age_group"))
}

// BusinessPhones returns a reference to field business_phones of azuread_user.
func (au azureadUserAttributes) BusinessPhones() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](au.ref.Append("business_phones"))
}

// City returns a reference to field city of azuread_user.
func (au azureadUserAttributes) City() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("city"))
}

// CompanyName returns a reference to field company_name of azuread_user.
func (au azureadUserAttributes) CompanyName() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("company_name"))
}

// ConsentProvidedForMinor returns a reference to field consent_provided_for_minor of azuread_user.
func (au azureadUserAttributes) ConsentProvidedForMinor() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("consent_provided_for_minor"))
}

// CostCenter returns a reference to field cost_center of azuread_user.
func (au azureadUserAttributes) CostCenter() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("cost_center"))
}

// Country returns a reference to field country of azuread_user.
func (au azureadUserAttributes) Country() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("country"))
}

// CreationType returns a reference to field creation_type of azuread_user.
func (au azureadUserAttributes) CreationType() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("creation_type"))
}

// Department returns a reference to field department of azuread_user.
func (au azureadUserAttributes) Department() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("department"))
}

// DisablePasswordExpiration returns a reference to field disable_password_expiration of azuread_user.
func (au azureadUserAttributes) DisablePasswordExpiration() terra.BoolValue {
	return terra.ReferenceAsBool(au.ref.Append("disable_password_expiration"))
}

// DisableStrongPassword returns a reference to field disable_strong_password of azuread_user.
func (au azureadUserAttributes) DisableStrongPassword() terra.BoolValue {
	return terra.ReferenceAsBool(au.ref.Append("disable_strong_password"))
}

// DisplayName returns a reference to field display_name of azuread_user.
func (au azureadUserAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("display_name"))
}

// Division returns a reference to field division of azuread_user.
func (au azureadUserAttributes) Division() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("division"))
}

// EmployeeId returns a reference to field employee_id of azuread_user.
func (au azureadUserAttributes) EmployeeId() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("employee_id"))
}

// EmployeeType returns a reference to field employee_type of azuread_user.
func (au azureadUserAttributes) EmployeeType() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("employee_type"))
}

// ExternalUserState returns a reference to field external_user_state of azuread_user.
func (au azureadUserAttributes) ExternalUserState() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("external_user_state"))
}

// FaxNumber returns a reference to field fax_number of azuread_user.
func (au azureadUserAttributes) FaxNumber() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("fax_number"))
}

// ForcePasswordChange returns a reference to field force_password_change of azuread_user.
func (au azureadUserAttributes) ForcePasswordChange() terra.BoolValue {
	return terra.ReferenceAsBool(au.ref.Append("force_password_change"))
}

// GivenName returns a reference to field given_name of azuread_user.
func (au azureadUserAttributes) GivenName() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("given_name"))
}

// Id returns a reference to field id of azuread_user.
func (au azureadUserAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("id"))
}

// ImAddresses returns a reference to field im_addresses of azuread_user.
func (au azureadUserAttributes) ImAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](au.ref.Append("im_addresses"))
}

// JobTitle returns a reference to field job_title of azuread_user.
func (au azureadUserAttributes) JobTitle() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("job_title"))
}

// Mail returns a reference to field mail of azuread_user.
func (au azureadUserAttributes) Mail() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("mail"))
}

// MailNickname returns a reference to field mail_nickname of azuread_user.
func (au azureadUserAttributes) MailNickname() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("mail_nickname"))
}

// ManagerId returns a reference to field manager_id of azuread_user.
func (au azureadUserAttributes) ManagerId() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("manager_id"))
}

// MobilePhone returns a reference to field mobile_phone of azuread_user.
func (au azureadUserAttributes) MobilePhone() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("mobile_phone"))
}

// ObjectId returns a reference to field object_id of azuread_user.
func (au azureadUserAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("object_id"))
}

// OfficeLocation returns a reference to field office_location of azuread_user.
func (au azureadUserAttributes) OfficeLocation() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("office_location"))
}

// OnpremisesDistinguishedName returns a reference to field onpremises_distinguished_name of azuread_user.
func (au azureadUserAttributes) OnpremisesDistinguishedName() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("onpremises_distinguished_name"))
}

// OnpremisesDomainName returns a reference to field onpremises_domain_name of azuread_user.
func (au azureadUserAttributes) OnpremisesDomainName() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("onpremises_domain_name"))
}

// OnpremisesImmutableId returns a reference to field onpremises_immutable_id of azuread_user.
func (au azureadUserAttributes) OnpremisesImmutableId() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("onpremises_immutable_id"))
}

// OnpremisesSamAccountName returns a reference to field onpremises_sam_account_name of azuread_user.
func (au azureadUserAttributes) OnpremisesSamAccountName() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("onpremises_sam_account_name"))
}

// OnpremisesSecurityIdentifier returns a reference to field onpremises_security_identifier of azuread_user.
func (au azureadUserAttributes) OnpremisesSecurityIdentifier() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("onpremises_security_identifier"))
}

// OnpremisesSyncEnabled returns a reference to field onpremises_sync_enabled of azuread_user.
func (au azureadUserAttributes) OnpremisesSyncEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(au.ref.Append("onpremises_sync_enabled"))
}

// OnpremisesUserPrincipalName returns a reference to field onpremises_user_principal_name of azuread_user.
func (au azureadUserAttributes) OnpremisesUserPrincipalName() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("onpremises_user_principal_name"))
}

// OtherMails returns a reference to field other_mails of azuread_user.
func (au azureadUserAttributes) OtherMails() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](au.ref.Append("other_mails"))
}

// Password returns a reference to field password of azuread_user.
func (au azureadUserAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("password"))
}

// PostalCode returns a reference to field postal_code of azuread_user.
func (au azureadUserAttributes) PostalCode() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("postal_code"))
}

// PreferredLanguage returns a reference to field preferred_language of azuread_user.
func (au azureadUserAttributes) PreferredLanguage() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("preferred_language"))
}

// ProxyAddresses returns a reference to field proxy_addresses of azuread_user.
func (au azureadUserAttributes) ProxyAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](au.ref.Append("proxy_addresses"))
}

// ShowInAddressList returns a reference to field show_in_address_list of azuread_user.
func (au azureadUserAttributes) ShowInAddressList() terra.BoolValue {
	return terra.ReferenceAsBool(au.ref.Append("show_in_address_list"))
}

// State returns a reference to field state of azuread_user.
func (au azureadUserAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("state"))
}

// StreetAddress returns a reference to field street_address of azuread_user.
func (au azureadUserAttributes) StreetAddress() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("street_address"))
}

// Surname returns a reference to field surname of azuread_user.
func (au azureadUserAttributes) Surname() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("surname"))
}

// UsageLocation returns a reference to field usage_location of azuread_user.
func (au azureadUserAttributes) UsageLocation() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("usage_location"))
}

// UserPrincipalName returns a reference to field user_principal_name of azuread_user.
func (au azureadUserAttributes) UserPrincipalName() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("user_principal_name"))
}

// UserType returns a reference to field user_type of azuread_user.
func (au azureadUserAttributes) UserType() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("user_type"))
}

func (au azureadUserAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](au.ref.Append("timeouts"))
}

type azureadUserState struct {
	AboutMe                      string         `json:"about_me"`
	AccountEnabled               bool           `json:"account_enabled"`
	AgeGroup                     string         `json:"age_group"`
	BusinessPhones               []string       `json:"business_phones"`
	City                         string         `json:"city"`
	CompanyName                  string         `json:"company_name"`
	ConsentProvidedForMinor      string         `json:"consent_provided_for_minor"`
	CostCenter                   string         `json:"cost_center"`
	Country                      string         `json:"country"`
	CreationType                 string         `json:"creation_type"`
	Department                   string         `json:"department"`
	DisablePasswordExpiration    bool           `json:"disable_password_expiration"`
	DisableStrongPassword        bool           `json:"disable_strong_password"`
	DisplayName                  string         `json:"display_name"`
	Division                     string         `json:"division"`
	EmployeeId                   string         `json:"employee_id"`
	EmployeeType                 string         `json:"employee_type"`
	ExternalUserState            string         `json:"external_user_state"`
	FaxNumber                    string         `json:"fax_number"`
	ForcePasswordChange          bool           `json:"force_password_change"`
	GivenName                    string         `json:"given_name"`
	Id                           string         `json:"id"`
	ImAddresses                  []string       `json:"im_addresses"`
	JobTitle                     string         `json:"job_title"`
	Mail                         string         `json:"mail"`
	MailNickname                 string         `json:"mail_nickname"`
	ManagerId                    string         `json:"manager_id"`
	MobilePhone                  string         `json:"mobile_phone"`
	ObjectId                     string         `json:"object_id"`
	OfficeLocation               string         `json:"office_location"`
	OnpremisesDistinguishedName  string         `json:"onpremises_distinguished_name"`
	OnpremisesDomainName         string         `json:"onpremises_domain_name"`
	OnpremisesImmutableId        string         `json:"onpremises_immutable_id"`
	OnpremisesSamAccountName     string         `json:"onpremises_sam_account_name"`
	OnpremisesSecurityIdentifier string         `json:"onpremises_security_identifier"`
	OnpremisesSyncEnabled        bool           `json:"onpremises_sync_enabled"`
	OnpremisesUserPrincipalName  string         `json:"onpremises_user_principal_name"`
	OtherMails                   []string       `json:"other_mails"`
	Password                     string         `json:"password"`
	PostalCode                   string         `json:"postal_code"`
	PreferredLanguage            string         `json:"preferred_language"`
	ProxyAddresses               []string       `json:"proxy_addresses"`
	ShowInAddressList            bool           `json:"show_in_address_list"`
	State                        string         `json:"state"`
	StreetAddress                string         `json:"street_address"`
	Surname                      string         `json:"surname"`
	UsageLocation                string         `json:"usage_location"`
	UserPrincipalName            string         `json:"user_principal_name"`
	UserType                     string         `json:"user_type"`
	Timeouts                     *TimeoutsState `json:"timeouts"`
}
