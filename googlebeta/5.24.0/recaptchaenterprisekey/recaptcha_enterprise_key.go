// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package recaptchaenterprisekey

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type AndroidSettings struct {
	// AllowAllPackageNames: bool, optional
	AllowAllPackageNames terra.BoolValue `hcl:"allow_all_package_names,attr"`
	// AllowedPackageNames: list of string, optional
	AllowedPackageNames terra.ListValue[terra.StringValue] `hcl:"allowed_package_names,attr"`
}

type IosSettings struct {
	// AllowAllBundleIds: bool, optional
	AllowAllBundleIds terra.BoolValue `hcl:"allow_all_bundle_ids,attr"`
	// AllowedBundleIds: list of string, optional
	AllowedBundleIds terra.ListValue[terra.StringValue] `hcl:"allowed_bundle_ids,attr"`
}

type TestingOptions struct {
	// TestingChallenge: string, optional
	TestingChallenge terra.StringValue `hcl:"testing_challenge,attr"`
	// TestingScore: number, optional
	TestingScore terra.NumberValue `hcl:"testing_score,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type WafSettings struct {
	// WafFeature: string, required
	WafFeature terra.StringValue `hcl:"waf_feature,attr" validate:"required"`
	// WafService: string, required
	WafService terra.StringValue `hcl:"waf_service,attr" validate:"required"`
}

type WebSettings struct {
	// AllowAllDomains: bool, optional
	AllowAllDomains terra.BoolValue `hcl:"allow_all_domains,attr"`
	// AllowAmpTraffic: bool, optional
	AllowAmpTraffic terra.BoolValue `hcl:"allow_amp_traffic,attr"`
	// AllowedDomains: list of string, optional
	AllowedDomains terra.ListValue[terra.StringValue] `hcl:"allowed_domains,attr"`
	// ChallengeSecurityPreference: string, optional
	ChallengeSecurityPreference terra.StringValue `hcl:"challenge_security_preference,attr"`
	// IntegrationType: string, required
	IntegrationType terra.StringValue `hcl:"integration_type,attr" validate:"required"`
}

type AndroidSettingsAttributes struct {
	ref terra.Reference
}

func (as AndroidSettingsAttributes) InternalRef() (terra.Reference, error) {
	return as.ref, nil
}

func (as AndroidSettingsAttributes) InternalWithRef(ref terra.Reference) AndroidSettingsAttributes {
	return AndroidSettingsAttributes{ref: ref}
}

func (as AndroidSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return as.ref.InternalTokens()
}

func (as AndroidSettingsAttributes) AllowAllPackageNames() terra.BoolValue {
	return terra.ReferenceAsBool(as.ref.Append("allow_all_package_names"))
}

func (as AndroidSettingsAttributes) AllowedPackageNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](as.ref.Append("allowed_package_names"))
}

type IosSettingsAttributes struct {
	ref terra.Reference
}

func (is IosSettingsAttributes) InternalRef() (terra.Reference, error) {
	return is.ref, nil
}

func (is IosSettingsAttributes) InternalWithRef(ref terra.Reference) IosSettingsAttributes {
	return IosSettingsAttributes{ref: ref}
}

func (is IosSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return is.ref.InternalTokens()
}

func (is IosSettingsAttributes) AllowAllBundleIds() terra.BoolValue {
	return terra.ReferenceAsBool(is.ref.Append("allow_all_bundle_ids"))
}

func (is IosSettingsAttributes) AllowedBundleIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](is.ref.Append("allowed_bundle_ids"))
}

type TestingOptionsAttributes struct {
	ref terra.Reference
}

func (to TestingOptionsAttributes) InternalRef() (terra.Reference, error) {
	return to.ref, nil
}

func (to TestingOptionsAttributes) InternalWithRef(ref terra.Reference) TestingOptionsAttributes {
	return TestingOptionsAttributes{ref: ref}
}

func (to TestingOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return to.ref.InternalTokens()
}

func (to TestingOptionsAttributes) TestingChallenge() terra.StringValue {
	return terra.ReferenceAsString(to.ref.Append("testing_challenge"))
}

func (to TestingOptionsAttributes) TestingScore() terra.NumberValue {
	return terra.ReferenceAsNumber(to.ref.Append("testing_score"))
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

type WafSettingsAttributes struct {
	ref terra.Reference
}

func (ws WafSettingsAttributes) InternalRef() (terra.Reference, error) {
	return ws.ref, nil
}

func (ws WafSettingsAttributes) InternalWithRef(ref terra.Reference) WafSettingsAttributes {
	return WafSettingsAttributes{ref: ref}
}

func (ws WafSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ws.ref.InternalTokens()
}

func (ws WafSettingsAttributes) WafFeature() terra.StringValue {
	return terra.ReferenceAsString(ws.ref.Append("waf_feature"))
}

func (ws WafSettingsAttributes) WafService() terra.StringValue {
	return terra.ReferenceAsString(ws.ref.Append("waf_service"))
}

type WebSettingsAttributes struct {
	ref terra.Reference
}

func (ws WebSettingsAttributes) InternalRef() (terra.Reference, error) {
	return ws.ref, nil
}

func (ws WebSettingsAttributes) InternalWithRef(ref terra.Reference) WebSettingsAttributes {
	return WebSettingsAttributes{ref: ref}
}

func (ws WebSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ws.ref.InternalTokens()
}

func (ws WebSettingsAttributes) AllowAllDomains() terra.BoolValue {
	return terra.ReferenceAsBool(ws.ref.Append("allow_all_domains"))
}

func (ws WebSettingsAttributes) AllowAmpTraffic() terra.BoolValue {
	return terra.ReferenceAsBool(ws.ref.Append("allow_amp_traffic"))
}

func (ws WebSettingsAttributes) AllowedDomains() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ws.ref.Append("allowed_domains"))
}

func (ws WebSettingsAttributes) ChallengeSecurityPreference() terra.StringValue {
	return terra.ReferenceAsString(ws.ref.Append("challenge_security_preference"))
}

func (ws WebSettingsAttributes) IntegrationType() terra.StringValue {
	return terra.ReferenceAsString(ws.ref.Append("integration_type"))
}

type AndroidSettingsState struct {
	AllowAllPackageNames bool     `json:"allow_all_package_names"`
	AllowedPackageNames  []string `json:"allowed_package_names"`
}

type IosSettingsState struct {
	AllowAllBundleIds bool     `json:"allow_all_bundle_ids"`
	AllowedBundleIds  []string `json:"allowed_bundle_ids"`
}

type TestingOptionsState struct {
	TestingChallenge string  `json:"testing_challenge"`
	TestingScore     float64 `json:"testing_score"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

type WafSettingsState struct {
	WafFeature string `json:"waf_feature"`
	WafService string `json:"waf_service"`
}

type WebSettingsState struct {
	AllowAllDomains             bool     `json:"allow_all_domains"`
	AllowAmpTraffic             bool     `json:"allow_amp_traffic"`
	AllowedDomains              []string `json:"allowed_domains"`
	ChallengeSecurityPreference string   `json:"challenge_security_preference"`
	IntegrationType             string   `json:"integration_type"`
}
