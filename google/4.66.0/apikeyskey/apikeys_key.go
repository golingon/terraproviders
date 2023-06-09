// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package apikeyskey

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Restrictions struct {
	// AndroidKeyRestrictions: optional
	AndroidKeyRestrictions *AndroidKeyRestrictions `hcl:"android_key_restrictions,block"`
	// ApiTargets: min=0
	ApiTargets []ApiTargets `hcl:"api_targets,block" validate:"min=0"`
	// BrowserKeyRestrictions: optional
	BrowserKeyRestrictions *BrowserKeyRestrictions `hcl:"browser_key_restrictions,block"`
	// IosKeyRestrictions: optional
	IosKeyRestrictions *IosKeyRestrictions `hcl:"ios_key_restrictions,block"`
	// ServerKeyRestrictions: optional
	ServerKeyRestrictions *ServerKeyRestrictions `hcl:"server_key_restrictions,block"`
}

type AndroidKeyRestrictions struct {
	// AllowedApplications: min=1
	AllowedApplications []AllowedApplications `hcl:"allowed_applications,block" validate:"min=1"`
}

type AllowedApplications struct {
	// PackageName: string, required
	PackageName terra.StringValue `hcl:"package_name,attr" validate:"required"`
	// Sha1Fingerprint: string, required
	Sha1Fingerprint terra.StringValue `hcl:"sha1_fingerprint,attr" validate:"required"`
}

type ApiTargets struct {
	// Methods: list of string, optional
	Methods terra.ListValue[terra.StringValue] `hcl:"methods,attr"`
	// Service: string, required
	Service terra.StringValue `hcl:"service,attr" validate:"required"`
}

type BrowserKeyRestrictions struct {
	// AllowedReferrers: list of string, required
	AllowedReferrers terra.ListValue[terra.StringValue] `hcl:"allowed_referrers,attr" validate:"required"`
}

type IosKeyRestrictions struct {
	// AllowedBundleIds: list of string, required
	AllowedBundleIds terra.ListValue[terra.StringValue] `hcl:"allowed_bundle_ids,attr" validate:"required"`
}

type ServerKeyRestrictions struct {
	// AllowedIps: list of string, required
	AllowedIps terra.ListValue[terra.StringValue] `hcl:"allowed_ips,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type RestrictionsAttributes struct {
	ref terra.Reference
}

func (r RestrictionsAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RestrictionsAttributes) InternalWithRef(ref terra.Reference) RestrictionsAttributes {
	return RestrictionsAttributes{ref: ref}
}

func (r RestrictionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RestrictionsAttributes) AndroidKeyRestrictions() terra.ListValue[AndroidKeyRestrictionsAttributes] {
	return terra.ReferenceAsList[AndroidKeyRestrictionsAttributes](r.ref.Append("android_key_restrictions"))
}

func (r RestrictionsAttributes) ApiTargets() terra.ListValue[ApiTargetsAttributes] {
	return terra.ReferenceAsList[ApiTargetsAttributes](r.ref.Append("api_targets"))
}

func (r RestrictionsAttributes) BrowserKeyRestrictions() terra.ListValue[BrowserKeyRestrictionsAttributes] {
	return terra.ReferenceAsList[BrowserKeyRestrictionsAttributes](r.ref.Append("browser_key_restrictions"))
}

func (r RestrictionsAttributes) IosKeyRestrictions() terra.ListValue[IosKeyRestrictionsAttributes] {
	return terra.ReferenceAsList[IosKeyRestrictionsAttributes](r.ref.Append("ios_key_restrictions"))
}

func (r RestrictionsAttributes) ServerKeyRestrictions() terra.ListValue[ServerKeyRestrictionsAttributes] {
	return terra.ReferenceAsList[ServerKeyRestrictionsAttributes](r.ref.Append("server_key_restrictions"))
}

type AndroidKeyRestrictionsAttributes struct {
	ref terra.Reference
}

func (akr AndroidKeyRestrictionsAttributes) InternalRef() (terra.Reference, error) {
	return akr.ref, nil
}

func (akr AndroidKeyRestrictionsAttributes) InternalWithRef(ref terra.Reference) AndroidKeyRestrictionsAttributes {
	return AndroidKeyRestrictionsAttributes{ref: ref}
}

func (akr AndroidKeyRestrictionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return akr.ref.InternalTokens()
}

func (akr AndroidKeyRestrictionsAttributes) AllowedApplications() terra.ListValue[AllowedApplicationsAttributes] {
	return terra.ReferenceAsList[AllowedApplicationsAttributes](akr.ref.Append("allowed_applications"))
}

type AllowedApplicationsAttributes struct {
	ref terra.Reference
}

func (aa AllowedApplicationsAttributes) InternalRef() (terra.Reference, error) {
	return aa.ref, nil
}

func (aa AllowedApplicationsAttributes) InternalWithRef(ref terra.Reference) AllowedApplicationsAttributes {
	return AllowedApplicationsAttributes{ref: ref}
}

func (aa AllowedApplicationsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return aa.ref.InternalTokens()
}

func (aa AllowedApplicationsAttributes) PackageName() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("package_name"))
}

func (aa AllowedApplicationsAttributes) Sha1Fingerprint() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("sha1_fingerprint"))
}

type ApiTargetsAttributes struct {
	ref terra.Reference
}

func (at ApiTargetsAttributes) InternalRef() (terra.Reference, error) {
	return at.ref, nil
}

func (at ApiTargetsAttributes) InternalWithRef(ref terra.Reference) ApiTargetsAttributes {
	return ApiTargetsAttributes{ref: ref}
}

func (at ApiTargetsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return at.ref.InternalTokens()
}

func (at ApiTargetsAttributes) Methods() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](at.ref.Append("methods"))
}

func (at ApiTargetsAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(at.ref.Append("service"))
}

type BrowserKeyRestrictionsAttributes struct {
	ref terra.Reference
}

func (bkr BrowserKeyRestrictionsAttributes) InternalRef() (terra.Reference, error) {
	return bkr.ref, nil
}

func (bkr BrowserKeyRestrictionsAttributes) InternalWithRef(ref terra.Reference) BrowserKeyRestrictionsAttributes {
	return BrowserKeyRestrictionsAttributes{ref: ref}
}

func (bkr BrowserKeyRestrictionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bkr.ref.InternalTokens()
}

func (bkr BrowserKeyRestrictionsAttributes) AllowedReferrers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](bkr.ref.Append("allowed_referrers"))
}

type IosKeyRestrictionsAttributes struct {
	ref terra.Reference
}

func (ikr IosKeyRestrictionsAttributes) InternalRef() (terra.Reference, error) {
	return ikr.ref, nil
}

func (ikr IosKeyRestrictionsAttributes) InternalWithRef(ref terra.Reference) IosKeyRestrictionsAttributes {
	return IosKeyRestrictionsAttributes{ref: ref}
}

func (ikr IosKeyRestrictionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ikr.ref.InternalTokens()
}

func (ikr IosKeyRestrictionsAttributes) AllowedBundleIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ikr.ref.Append("allowed_bundle_ids"))
}

type ServerKeyRestrictionsAttributes struct {
	ref terra.Reference
}

func (skr ServerKeyRestrictionsAttributes) InternalRef() (terra.Reference, error) {
	return skr.ref, nil
}

func (skr ServerKeyRestrictionsAttributes) InternalWithRef(ref terra.Reference) ServerKeyRestrictionsAttributes {
	return ServerKeyRestrictionsAttributes{ref: ref}
}

func (skr ServerKeyRestrictionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return skr.ref.InternalTokens()
}

func (skr ServerKeyRestrictionsAttributes) AllowedIps() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](skr.ref.Append("allowed_ips"))
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

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type RestrictionsState struct {
	AndroidKeyRestrictions []AndroidKeyRestrictionsState `json:"android_key_restrictions"`
	ApiTargets             []ApiTargetsState             `json:"api_targets"`
	BrowserKeyRestrictions []BrowserKeyRestrictionsState `json:"browser_key_restrictions"`
	IosKeyRestrictions     []IosKeyRestrictionsState     `json:"ios_key_restrictions"`
	ServerKeyRestrictions  []ServerKeyRestrictionsState  `json:"server_key_restrictions"`
}

type AndroidKeyRestrictionsState struct {
	AllowedApplications []AllowedApplicationsState `json:"allowed_applications"`
}

type AllowedApplicationsState struct {
	PackageName     string `json:"package_name"`
	Sha1Fingerprint string `json:"sha1_fingerprint"`
}

type ApiTargetsState struct {
	Methods []string `json:"methods"`
	Service string   `json:"service"`
}

type BrowserKeyRestrictionsState struct {
	AllowedReferrers []string `json:"allowed_referrers"`
}

type IosKeyRestrictionsState struct {
	AllowedBundleIds []string `json:"allowed_bundle_ids"`
}

type ServerKeyRestrictionsState struct {
	AllowedIps []string `json:"allowed_ips"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
