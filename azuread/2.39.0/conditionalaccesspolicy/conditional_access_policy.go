// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package conditionalaccesspolicy

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Conditions struct {
	// ClientAppTypes: list of string, required
	ClientAppTypes terra.ListValue[terra.StringValue] `hcl:"client_app_types,attr" validate:"required"`
	// SignInRiskLevels: list of string, optional
	SignInRiskLevels terra.ListValue[terra.StringValue] `hcl:"sign_in_risk_levels,attr"`
	// UserRiskLevels: list of string, optional
	UserRiskLevels terra.ListValue[terra.StringValue] `hcl:"user_risk_levels,attr"`
	// Applications: required
	Applications *Applications `hcl:"applications,block" validate:"required"`
	// Devices: optional
	Devices *Devices `hcl:"devices,block"`
	// Locations: optional
	Locations *Locations `hcl:"locations,block"`
	// Platforms: optional
	Platforms *Platforms `hcl:"platforms,block"`
	// Users: required
	Users *Users `hcl:"users,block" validate:"required"`
}

type Applications struct {
	// ExcludedApplications: list of string, optional
	ExcludedApplications terra.ListValue[terra.StringValue] `hcl:"excluded_applications,attr"`
	// IncludedApplications: list of string, optional
	IncludedApplications terra.ListValue[terra.StringValue] `hcl:"included_applications,attr"`
	// IncludedUserActions: list of string, optional
	IncludedUserActions terra.ListValue[terra.StringValue] `hcl:"included_user_actions,attr"`
}

type Devices struct {
	// Filter: optional
	Filter *Filter `hcl:"filter,block"`
}

type Filter struct {
	// Mode: string, required
	Mode terra.StringValue `hcl:"mode,attr" validate:"required"`
	// Rule: string, required
	Rule terra.StringValue `hcl:"rule,attr" validate:"required"`
}

type Locations struct {
	// ExcludedLocations: list of string, optional
	ExcludedLocations terra.ListValue[terra.StringValue] `hcl:"excluded_locations,attr"`
	// IncludedLocations: list of string, required
	IncludedLocations terra.ListValue[terra.StringValue] `hcl:"included_locations,attr" validate:"required"`
}

type Platforms struct {
	// ExcludedPlatforms: list of string, optional
	ExcludedPlatforms terra.ListValue[terra.StringValue] `hcl:"excluded_platforms,attr"`
	// IncludedPlatforms: list of string, required
	IncludedPlatforms terra.ListValue[terra.StringValue] `hcl:"included_platforms,attr" validate:"required"`
}

type Users struct {
	// ExcludedGroups: list of string, optional
	ExcludedGroups terra.ListValue[terra.StringValue] `hcl:"excluded_groups,attr"`
	// ExcludedRoles: list of string, optional
	ExcludedRoles terra.ListValue[terra.StringValue] `hcl:"excluded_roles,attr"`
	// ExcludedUsers: list of string, optional
	ExcludedUsers terra.ListValue[terra.StringValue] `hcl:"excluded_users,attr"`
	// IncludedGroups: list of string, optional
	IncludedGroups terra.ListValue[terra.StringValue] `hcl:"included_groups,attr"`
	// IncludedRoles: list of string, optional
	IncludedRoles terra.ListValue[terra.StringValue] `hcl:"included_roles,attr"`
	// IncludedUsers: list of string, optional
	IncludedUsers terra.ListValue[terra.StringValue] `hcl:"included_users,attr"`
}

type GrantControls struct {
	// BuiltInControls: list of string, required
	BuiltInControls terra.ListValue[terra.StringValue] `hcl:"built_in_controls,attr" validate:"required"`
	// CustomAuthenticationFactors: list of string, optional
	CustomAuthenticationFactors terra.ListValue[terra.StringValue] `hcl:"custom_authentication_factors,attr"`
	// Operator: string, required
	Operator terra.StringValue `hcl:"operator,attr" validate:"required"`
	// TermsOfUse: list of string, optional
	TermsOfUse terra.ListValue[terra.StringValue] `hcl:"terms_of_use,attr"`
}

type SessionControls struct {
	// ApplicationEnforcedRestrictionsEnabled: bool, optional
	ApplicationEnforcedRestrictionsEnabled terra.BoolValue `hcl:"application_enforced_restrictions_enabled,attr"`
	// CloudAppSecurityPolicy: string, optional
	CloudAppSecurityPolicy terra.StringValue `hcl:"cloud_app_security_policy,attr"`
	// PersistentBrowserMode: string, optional
	PersistentBrowserMode terra.StringValue `hcl:"persistent_browser_mode,attr"`
	// SignInFrequency: number, optional
	SignInFrequency terra.NumberValue `hcl:"sign_in_frequency,attr"`
	// SignInFrequencyPeriod: string, optional
	SignInFrequencyPeriod terra.StringValue `hcl:"sign_in_frequency_period,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type ConditionsAttributes struct {
	ref terra.Reference
}

func (c ConditionsAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ConditionsAttributes) InternalWithRef(ref terra.Reference) ConditionsAttributes {
	return ConditionsAttributes{ref: ref}
}

func (c ConditionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ConditionsAttributes) ClientAppTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](c.ref.Append("client_app_types"))
}

func (c ConditionsAttributes) SignInRiskLevels() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](c.ref.Append("sign_in_risk_levels"))
}

func (c ConditionsAttributes) UserRiskLevels() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](c.ref.Append("user_risk_levels"))
}

func (c ConditionsAttributes) Applications() terra.ListValue[ApplicationsAttributes] {
	return terra.ReferenceAsList[ApplicationsAttributes](c.ref.Append("applications"))
}

func (c ConditionsAttributes) Devices() terra.ListValue[DevicesAttributes] {
	return terra.ReferenceAsList[DevicesAttributes](c.ref.Append("devices"))
}

func (c ConditionsAttributes) Locations() terra.ListValue[LocationsAttributes] {
	return terra.ReferenceAsList[LocationsAttributes](c.ref.Append("locations"))
}

func (c ConditionsAttributes) Platforms() terra.ListValue[PlatformsAttributes] {
	return terra.ReferenceAsList[PlatformsAttributes](c.ref.Append("platforms"))
}

func (c ConditionsAttributes) Users() terra.ListValue[UsersAttributes] {
	return terra.ReferenceAsList[UsersAttributes](c.ref.Append("users"))
}

type ApplicationsAttributes struct {
	ref terra.Reference
}

func (a ApplicationsAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a ApplicationsAttributes) InternalWithRef(ref terra.Reference) ApplicationsAttributes {
	return ApplicationsAttributes{ref: ref}
}

func (a ApplicationsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a ApplicationsAttributes) ExcludedApplications() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](a.ref.Append("excluded_applications"))
}

func (a ApplicationsAttributes) IncludedApplications() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](a.ref.Append("included_applications"))
}

func (a ApplicationsAttributes) IncludedUserActions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](a.ref.Append("included_user_actions"))
}

type DevicesAttributes struct {
	ref terra.Reference
}

func (d DevicesAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d DevicesAttributes) InternalWithRef(ref terra.Reference) DevicesAttributes {
	return DevicesAttributes{ref: ref}
}

func (d DevicesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d DevicesAttributes) Filter() terra.ListValue[FilterAttributes] {
	return terra.ReferenceAsList[FilterAttributes](d.ref.Append("filter"))
}

type FilterAttributes struct {
	ref terra.Reference
}

func (f FilterAttributes) InternalRef() (terra.Reference, error) {
	return f.ref, nil
}

func (f FilterAttributes) InternalWithRef(ref terra.Reference) FilterAttributes {
	return FilterAttributes{ref: ref}
}

func (f FilterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return f.ref.InternalTokens()
}

func (f FilterAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("mode"))
}

func (f FilterAttributes) Rule() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("rule"))
}

type LocationsAttributes struct {
	ref terra.Reference
}

func (l LocationsAttributes) InternalRef() (terra.Reference, error) {
	return l.ref, nil
}

func (l LocationsAttributes) InternalWithRef(ref terra.Reference) LocationsAttributes {
	return LocationsAttributes{ref: ref}
}

func (l LocationsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return l.ref.InternalTokens()
}

func (l LocationsAttributes) ExcludedLocations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](l.ref.Append("excluded_locations"))
}

func (l LocationsAttributes) IncludedLocations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](l.ref.Append("included_locations"))
}

type PlatformsAttributes struct {
	ref terra.Reference
}

func (p PlatformsAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p PlatformsAttributes) InternalWithRef(ref terra.Reference) PlatformsAttributes {
	return PlatformsAttributes{ref: ref}
}

func (p PlatformsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p PlatformsAttributes) ExcludedPlatforms() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](p.ref.Append("excluded_platforms"))
}

func (p PlatformsAttributes) IncludedPlatforms() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](p.ref.Append("included_platforms"))
}

type UsersAttributes struct {
	ref terra.Reference
}

func (u UsersAttributes) InternalRef() (terra.Reference, error) {
	return u.ref, nil
}

func (u UsersAttributes) InternalWithRef(ref terra.Reference) UsersAttributes {
	return UsersAttributes{ref: ref}
}

func (u UsersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return u.ref.InternalTokens()
}

func (u UsersAttributes) ExcludedGroups() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](u.ref.Append("excluded_groups"))
}

func (u UsersAttributes) ExcludedRoles() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](u.ref.Append("excluded_roles"))
}

func (u UsersAttributes) ExcludedUsers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](u.ref.Append("excluded_users"))
}

func (u UsersAttributes) IncludedGroups() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](u.ref.Append("included_groups"))
}

func (u UsersAttributes) IncludedRoles() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](u.ref.Append("included_roles"))
}

func (u UsersAttributes) IncludedUsers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](u.ref.Append("included_users"))
}

type GrantControlsAttributes struct {
	ref terra.Reference
}

func (gc GrantControlsAttributes) InternalRef() (terra.Reference, error) {
	return gc.ref, nil
}

func (gc GrantControlsAttributes) InternalWithRef(ref terra.Reference) GrantControlsAttributes {
	return GrantControlsAttributes{ref: ref}
}

func (gc GrantControlsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return gc.ref.InternalTokens()
}

func (gc GrantControlsAttributes) BuiltInControls() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](gc.ref.Append("built_in_controls"))
}

func (gc GrantControlsAttributes) CustomAuthenticationFactors() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](gc.ref.Append("custom_authentication_factors"))
}

func (gc GrantControlsAttributes) Operator() terra.StringValue {
	return terra.ReferenceAsString(gc.ref.Append("operator"))
}

func (gc GrantControlsAttributes) TermsOfUse() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](gc.ref.Append("terms_of_use"))
}

type SessionControlsAttributes struct {
	ref terra.Reference
}

func (sc SessionControlsAttributes) InternalRef() (terra.Reference, error) {
	return sc.ref, nil
}

func (sc SessionControlsAttributes) InternalWithRef(ref terra.Reference) SessionControlsAttributes {
	return SessionControlsAttributes{ref: ref}
}

func (sc SessionControlsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sc.ref.InternalTokens()
}

func (sc SessionControlsAttributes) ApplicationEnforcedRestrictionsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sc.ref.Append("application_enforced_restrictions_enabled"))
}

func (sc SessionControlsAttributes) CloudAppSecurityPolicy() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("cloud_app_security_policy"))
}

func (sc SessionControlsAttributes) PersistentBrowserMode() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("persistent_browser_mode"))
}

func (sc SessionControlsAttributes) SignInFrequency() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("sign_in_frequency"))
}

func (sc SessionControlsAttributes) SignInFrequencyPeriod() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("sign_in_frequency_period"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type ConditionsState struct {
	ClientAppTypes   []string            `json:"client_app_types"`
	SignInRiskLevels []string            `json:"sign_in_risk_levels"`
	UserRiskLevels   []string            `json:"user_risk_levels"`
	Applications     []ApplicationsState `json:"applications"`
	Devices          []DevicesState      `json:"devices"`
	Locations        []LocationsState    `json:"locations"`
	Platforms        []PlatformsState    `json:"platforms"`
	Users            []UsersState        `json:"users"`
}

type ApplicationsState struct {
	ExcludedApplications []string `json:"excluded_applications"`
	IncludedApplications []string `json:"included_applications"`
	IncludedUserActions  []string `json:"included_user_actions"`
}

type DevicesState struct {
	Filter []FilterState `json:"filter"`
}

type FilterState struct {
	Mode string `json:"mode"`
	Rule string `json:"rule"`
}

type LocationsState struct {
	ExcludedLocations []string `json:"excluded_locations"`
	IncludedLocations []string `json:"included_locations"`
}

type PlatformsState struct {
	ExcludedPlatforms []string `json:"excluded_platforms"`
	IncludedPlatforms []string `json:"included_platforms"`
}

type UsersState struct {
	ExcludedGroups []string `json:"excluded_groups"`
	ExcludedRoles  []string `json:"excluded_roles"`
	ExcludedUsers  []string `json:"excluded_users"`
	IncludedGroups []string `json:"included_groups"`
	IncludedRoles  []string `json:"included_roles"`
	IncludedUsers  []string `json:"included_users"`
}

type GrantControlsState struct {
	BuiltInControls             []string `json:"built_in_controls"`
	CustomAuthenticationFactors []string `json:"custom_authentication_factors"`
	Operator                    string   `json:"operator"`
	TermsOfUse                  []string `json:"terms_of_use"`
}

type SessionControlsState struct {
	ApplicationEnforcedRestrictionsEnabled bool    `json:"application_enforced_restrictions_enabled"`
	CloudAppSecurityPolicy                 string  `json:"cloud_app_security_policy"`
	PersistentBrowserMode                  string  `json:"persistent_browser_mode"`
	SignInFrequency                        float64 `json:"sign_in_frequency"`
	SignInFrequencyPeriod                  string  `json:"sign_in_frequency_period"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}