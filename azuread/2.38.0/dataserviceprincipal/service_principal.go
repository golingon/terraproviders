// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataserviceprincipal

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AppRoles struct{}

type FeatureTags struct{}

type Features struct{}

type Oauth2PermissionScopes struct{}

type SamlSingleSignOn struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type AppRolesAttributes struct {
	ref terra.Reference
}

func (ar AppRolesAttributes) InternalRef() (terra.Reference, error) {
	return ar.ref, nil
}

func (ar AppRolesAttributes) InternalWithRef(ref terra.Reference) AppRolesAttributes {
	return AppRolesAttributes{ref: ref}
}

func (ar AppRolesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ar.ref.InternalTokens()
}

func (ar AppRolesAttributes) AllowedMemberTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ar.ref.Append("allowed_member_types"))
}

func (ar AppRolesAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("description"))
}

func (ar AppRolesAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("display_name"))
}

func (ar AppRolesAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(ar.ref.Append("enabled"))
}

func (ar AppRolesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("id"))
}

func (ar AppRolesAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("value"))
}

type FeatureTagsAttributes struct {
	ref terra.Reference
}

func (ft FeatureTagsAttributes) InternalRef() (terra.Reference, error) {
	return ft.ref, nil
}

func (ft FeatureTagsAttributes) InternalWithRef(ref terra.Reference) FeatureTagsAttributes {
	return FeatureTagsAttributes{ref: ref}
}

func (ft FeatureTagsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ft.ref.InternalTokens()
}

func (ft FeatureTagsAttributes) CustomSingleSignOn() terra.BoolValue {
	return terra.ReferenceAsBool(ft.ref.Append("custom_single_sign_on"))
}

func (ft FeatureTagsAttributes) Enterprise() terra.BoolValue {
	return terra.ReferenceAsBool(ft.ref.Append("enterprise"))
}

func (ft FeatureTagsAttributes) Gallery() terra.BoolValue {
	return terra.ReferenceAsBool(ft.ref.Append("gallery"))
}

func (ft FeatureTagsAttributes) Hide() terra.BoolValue {
	return terra.ReferenceAsBool(ft.ref.Append("hide"))
}

type FeaturesAttributes struct {
	ref terra.Reference
}

func (f FeaturesAttributes) InternalRef() (terra.Reference, error) {
	return f.ref, nil
}

func (f FeaturesAttributes) InternalWithRef(ref terra.Reference) FeaturesAttributes {
	return FeaturesAttributes{ref: ref}
}

func (f FeaturesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return f.ref.InternalTokens()
}

func (f FeaturesAttributes) CustomSingleSignOnApp() terra.BoolValue {
	return terra.ReferenceAsBool(f.ref.Append("custom_single_sign_on_app"))
}

func (f FeaturesAttributes) EnterpriseApplication() terra.BoolValue {
	return terra.ReferenceAsBool(f.ref.Append("enterprise_application"))
}

func (f FeaturesAttributes) GalleryApplication() terra.BoolValue {
	return terra.ReferenceAsBool(f.ref.Append("gallery_application"))
}

func (f FeaturesAttributes) VisibleToUsers() terra.BoolValue {
	return terra.ReferenceAsBool(f.ref.Append("visible_to_users"))
}

type Oauth2PermissionScopesAttributes struct {
	ref terra.Reference
}

func (ops Oauth2PermissionScopesAttributes) InternalRef() (terra.Reference, error) {
	return ops.ref, nil
}

func (ops Oauth2PermissionScopesAttributes) InternalWithRef(ref terra.Reference) Oauth2PermissionScopesAttributes {
	return Oauth2PermissionScopesAttributes{ref: ref}
}

func (ops Oauth2PermissionScopesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ops.ref.InternalTokens()
}

func (ops Oauth2PermissionScopesAttributes) AdminConsentDescription() terra.StringValue {
	return terra.ReferenceAsString(ops.ref.Append("admin_consent_description"))
}

func (ops Oauth2PermissionScopesAttributes) AdminConsentDisplayName() terra.StringValue {
	return terra.ReferenceAsString(ops.ref.Append("admin_consent_display_name"))
}

func (ops Oauth2PermissionScopesAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(ops.ref.Append("enabled"))
}

func (ops Oauth2PermissionScopesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ops.ref.Append("id"))
}

func (ops Oauth2PermissionScopesAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ops.ref.Append("type"))
}

func (ops Oauth2PermissionScopesAttributes) UserConsentDescription() terra.StringValue {
	return terra.ReferenceAsString(ops.ref.Append("user_consent_description"))
}

func (ops Oauth2PermissionScopesAttributes) UserConsentDisplayName() terra.StringValue {
	return terra.ReferenceAsString(ops.ref.Append("user_consent_display_name"))
}

func (ops Oauth2PermissionScopesAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(ops.ref.Append("value"))
}

type SamlSingleSignOnAttributes struct {
	ref terra.Reference
}

func (ssso SamlSingleSignOnAttributes) InternalRef() (terra.Reference, error) {
	return ssso.ref, nil
}

func (ssso SamlSingleSignOnAttributes) InternalWithRef(ref terra.Reference) SamlSingleSignOnAttributes {
	return SamlSingleSignOnAttributes{ref: ref}
}

func (ssso SamlSingleSignOnAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ssso.ref.InternalTokens()
}

func (ssso SamlSingleSignOnAttributes) RelayState() terra.StringValue {
	return terra.ReferenceAsString(ssso.ref.Append("relay_state"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type AppRolesState struct {
	AllowedMemberTypes []string `json:"allowed_member_types"`
	Description        string   `json:"description"`
	DisplayName        string   `json:"display_name"`
	Enabled            bool     `json:"enabled"`
	Id                 string   `json:"id"`
	Value              string   `json:"value"`
}

type FeatureTagsState struct {
	CustomSingleSignOn bool `json:"custom_single_sign_on"`
	Enterprise         bool `json:"enterprise"`
	Gallery            bool `json:"gallery"`
	Hide               bool `json:"hide"`
}

type FeaturesState struct {
	CustomSingleSignOnApp bool `json:"custom_single_sign_on_app"`
	EnterpriseApplication bool `json:"enterprise_application"`
	GalleryApplication    bool `json:"gallery_application"`
	VisibleToUsers        bool `json:"visible_to_users"`
}

type Oauth2PermissionScopesState struct {
	AdminConsentDescription string `json:"admin_consent_description"`
	AdminConsentDisplayName string `json:"admin_consent_display_name"`
	Enabled                 bool   `json:"enabled"`
	Id                      string `json:"id"`
	Type                    string `json:"type"`
	UserConsentDescription  string `json:"user_consent_description"`
	UserConsentDisplayName  string `json:"user_consent_display_name"`
	Value                   string `json:"value"`
}

type SamlSingleSignOnState struct {
	RelayState string `json:"relay_state"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
