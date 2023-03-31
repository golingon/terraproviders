// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package mqbroker

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Instances struct{}

type Configuration struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Revision: number, optional
	Revision terra.NumberValue `hcl:"revision,attr"`
}

type EncryptionOptions struct {
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// UseAwsOwnedKey: bool, optional
	UseAwsOwnedKey terra.BoolValue `hcl:"use_aws_owned_key,attr"`
}

type LdapServerMetadata struct {
	// Hosts: list of string, optional
	Hosts terra.ListValue[terra.StringValue] `hcl:"hosts,attr"`
	// RoleBase: string, optional
	RoleBase terra.StringValue `hcl:"role_base,attr"`
	// RoleName: string, optional
	RoleName terra.StringValue `hcl:"role_name,attr"`
	// RoleSearchMatching: string, optional
	RoleSearchMatching terra.StringValue `hcl:"role_search_matching,attr"`
	// RoleSearchSubtree: bool, optional
	RoleSearchSubtree terra.BoolValue `hcl:"role_search_subtree,attr"`
	// ServiceAccountPassword: string, optional
	ServiceAccountPassword terra.StringValue `hcl:"service_account_password,attr"`
	// ServiceAccountUsername: string, optional
	ServiceAccountUsername terra.StringValue `hcl:"service_account_username,attr"`
	// UserBase: string, optional
	UserBase terra.StringValue `hcl:"user_base,attr"`
	// UserRoleName: string, optional
	UserRoleName terra.StringValue `hcl:"user_role_name,attr"`
	// UserSearchMatching: string, optional
	UserSearchMatching terra.StringValue `hcl:"user_search_matching,attr"`
	// UserSearchSubtree: bool, optional
	UserSearchSubtree terra.BoolValue `hcl:"user_search_subtree,attr"`
}

type Logs struct {
	// Audit: string, optional
	Audit terra.StringValue `hcl:"audit,attr"`
	// General: bool, optional
	General terra.BoolValue `hcl:"general,attr"`
}

type MaintenanceWindowStartTime struct {
	// DayOfWeek: string, required
	DayOfWeek terra.StringValue `hcl:"day_of_week,attr" validate:"required"`
	// TimeOfDay: string, required
	TimeOfDay terra.StringValue `hcl:"time_of_day,attr" validate:"required"`
	// TimeZone: string, required
	TimeZone terra.StringValue `hcl:"time_zone,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type User struct {
	// ConsoleAccess: bool, optional
	ConsoleAccess terra.BoolValue `hcl:"console_access,attr"`
	// Groups: set of string, optional
	Groups terra.SetValue[terra.StringValue] `hcl:"groups,attr"`
	// Password: string, required
	Password terra.StringValue `hcl:"password,attr" validate:"required"`
	// Username: string, required
	Username terra.StringValue `hcl:"username,attr" validate:"required"`
}

type InstancesAttributes struct {
	ref terra.Reference
}

func (i InstancesAttributes) InternalRef() terra.Reference {
	return i.ref
}

func (i InstancesAttributes) InternalWithRef(ref terra.Reference) InstancesAttributes {
	return InstancesAttributes{ref: ref}
}

func (i InstancesAttributes) InternalTokens() hclwrite.Tokens {
	return i.ref.InternalTokens()
}

func (i InstancesAttributes) ConsoleUrl() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("console_url"))
}

func (i InstancesAttributes) Endpoints() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](i.ref.Append("endpoints"))
}

func (i InstancesAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("ip_address"))
}

type ConfigurationAttributes struct {
	ref terra.Reference
}

func (c ConfigurationAttributes) InternalRef() terra.Reference {
	return c.ref
}

func (c ConfigurationAttributes) InternalWithRef(ref terra.Reference) ConfigurationAttributes {
	return ConfigurationAttributes{ref: ref}
}

func (c ConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return c.ref.InternalTokens()
}

func (c ConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("id"))
}

func (c ConfigurationAttributes) Revision() terra.NumberValue {
	return terra.ReferenceAsNumber(c.ref.Append("revision"))
}

type EncryptionOptionsAttributes struct {
	ref terra.Reference
}

func (eo EncryptionOptionsAttributes) InternalRef() terra.Reference {
	return eo.ref
}

func (eo EncryptionOptionsAttributes) InternalWithRef(ref terra.Reference) EncryptionOptionsAttributes {
	return EncryptionOptionsAttributes{ref: ref}
}

func (eo EncryptionOptionsAttributes) InternalTokens() hclwrite.Tokens {
	return eo.ref.InternalTokens()
}

func (eo EncryptionOptionsAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(eo.ref.Append("kms_key_id"))
}

func (eo EncryptionOptionsAttributes) UseAwsOwnedKey() terra.BoolValue {
	return terra.ReferenceAsBool(eo.ref.Append("use_aws_owned_key"))
}

type LdapServerMetadataAttributes struct {
	ref terra.Reference
}

func (lsm LdapServerMetadataAttributes) InternalRef() terra.Reference {
	return lsm.ref
}

func (lsm LdapServerMetadataAttributes) InternalWithRef(ref terra.Reference) LdapServerMetadataAttributes {
	return LdapServerMetadataAttributes{ref: ref}
}

func (lsm LdapServerMetadataAttributes) InternalTokens() hclwrite.Tokens {
	return lsm.ref.InternalTokens()
}

func (lsm LdapServerMetadataAttributes) Hosts() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](lsm.ref.Append("hosts"))
}

func (lsm LdapServerMetadataAttributes) RoleBase() terra.StringValue {
	return terra.ReferenceAsString(lsm.ref.Append("role_base"))
}

func (lsm LdapServerMetadataAttributes) RoleName() terra.StringValue {
	return terra.ReferenceAsString(lsm.ref.Append("role_name"))
}

func (lsm LdapServerMetadataAttributes) RoleSearchMatching() terra.StringValue {
	return terra.ReferenceAsString(lsm.ref.Append("role_search_matching"))
}

func (lsm LdapServerMetadataAttributes) RoleSearchSubtree() terra.BoolValue {
	return terra.ReferenceAsBool(lsm.ref.Append("role_search_subtree"))
}

func (lsm LdapServerMetadataAttributes) ServiceAccountPassword() terra.StringValue {
	return terra.ReferenceAsString(lsm.ref.Append("service_account_password"))
}

func (lsm LdapServerMetadataAttributes) ServiceAccountUsername() terra.StringValue {
	return terra.ReferenceAsString(lsm.ref.Append("service_account_username"))
}

func (lsm LdapServerMetadataAttributes) UserBase() terra.StringValue {
	return terra.ReferenceAsString(lsm.ref.Append("user_base"))
}

func (lsm LdapServerMetadataAttributes) UserRoleName() terra.StringValue {
	return terra.ReferenceAsString(lsm.ref.Append("user_role_name"))
}

func (lsm LdapServerMetadataAttributes) UserSearchMatching() terra.StringValue {
	return terra.ReferenceAsString(lsm.ref.Append("user_search_matching"))
}

func (lsm LdapServerMetadataAttributes) UserSearchSubtree() terra.BoolValue {
	return terra.ReferenceAsBool(lsm.ref.Append("user_search_subtree"))
}

type LogsAttributes struct {
	ref terra.Reference
}

func (l LogsAttributes) InternalRef() terra.Reference {
	return l.ref
}

func (l LogsAttributes) InternalWithRef(ref terra.Reference) LogsAttributes {
	return LogsAttributes{ref: ref}
}

func (l LogsAttributes) InternalTokens() hclwrite.Tokens {
	return l.ref.InternalTokens()
}

func (l LogsAttributes) Audit() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("audit"))
}

func (l LogsAttributes) General() terra.BoolValue {
	return terra.ReferenceAsBool(l.ref.Append("general"))
}

type MaintenanceWindowStartTimeAttributes struct {
	ref terra.Reference
}

func (mwst MaintenanceWindowStartTimeAttributes) InternalRef() terra.Reference {
	return mwst.ref
}

func (mwst MaintenanceWindowStartTimeAttributes) InternalWithRef(ref terra.Reference) MaintenanceWindowStartTimeAttributes {
	return MaintenanceWindowStartTimeAttributes{ref: ref}
}

func (mwst MaintenanceWindowStartTimeAttributes) InternalTokens() hclwrite.Tokens {
	return mwst.ref.InternalTokens()
}

func (mwst MaintenanceWindowStartTimeAttributes) DayOfWeek() terra.StringValue {
	return terra.ReferenceAsString(mwst.ref.Append("day_of_week"))
}

func (mwst MaintenanceWindowStartTimeAttributes) TimeOfDay() terra.StringValue {
	return terra.ReferenceAsString(mwst.ref.Append("time_of_day"))
}

func (mwst MaintenanceWindowStartTimeAttributes) TimeZone() terra.StringValue {
	return terra.ReferenceAsString(mwst.ref.Append("time_zone"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type UserAttributes struct {
	ref terra.Reference
}

func (u UserAttributes) InternalRef() terra.Reference {
	return u.ref
}

func (u UserAttributes) InternalWithRef(ref terra.Reference) UserAttributes {
	return UserAttributes{ref: ref}
}

func (u UserAttributes) InternalTokens() hclwrite.Tokens {
	return u.ref.InternalTokens()
}

func (u UserAttributes) ConsoleAccess() terra.BoolValue {
	return terra.ReferenceAsBool(u.ref.Append("console_access"))
}

func (u UserAttributes) Groups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](u.ref.Append("groups"))
}

func (u UserAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(u.ref.Append("password"))
}

func (u UserAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(u.ref.Append("username"))
}

type InstancesState struct {
	ConsoleUrl string   `json:"console_url"`
	Endpoints  []string `json:"endpoints"`
	IpAddress  string   `json:"ip_address"`
}

type ConfigurationState struct {
	Id       string  `json:"id"`
	Revision float64 `json:"revision"`
}

type EncryptionOptionsState struct {
	KmsKeyId       string `json:"kms_key_id"`
	UseAwsOwnedKey bool   `json:"use_aws_owned_key"`
}

type LdapServerMetadataState struct {
	Hosts                  []string `json:"hosts"`
	RoleBase               string   `json:"role_base"`
	RoleName               string   `json:"role_name"`
	RoleSearchMatching     string   `json:"role_search_matching"`
	RoleSearchSubtree      bool     `json:"role_search_subtree"`
	ServiceAccountPassword string   `json:"service_account_password"`
	ServiceAccountUsername string   `json:"service_account_username"`
	UserBase               string   `json:"user_base"`
	UserRoleName           string   `json:"user_role_name"`
	UserSearchMatching     string   `json:"user_search_matching"`
	UserSearchSubtree      bool     `json:"user_search_subtree"`
}

type LogsState struct {
	Audit   string `json:"audit"`
	General bool   `json:"general"`
}

type MaintenanceWindowStartTimeState struct {
	DayOfWeek string `json:"day_of_week"`
	TimeOfDay string `json:"time_of_day"`
	TimeZone  string `json:"time_zone"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

type UserState struct {
	ConsoleAccess bool     `json:"console_access"`
	Groups        []string `json:"groups"`
	Password      string   `json:"password"`
	Username      string   `json:"username"`
}
