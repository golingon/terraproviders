// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package recaptchaenterprisekey

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
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

func (as AndroidSettingsAttributes) InternalRef() terra.Reference {
	return as.ref
}

func (as AndroidSettingsAttributes) InternalWithRef(ref terra.Reference) AndroidSettingsAttributes {
	return AndroidSettingsAttributes{ref: ref}
}

func (as AndroidSettingsAttributes) InternalTokens() hclwrite.Tokens {
	return as.ref.InternalTokens()
}

func (as AndroidSettingsAttributes) AllowAllPackageNames() terra.BoolValue {
	return terra.ReferenceBool(as.ref.Append("allow_all_package_names"))
}

func (as AndroidSettingsAttributes) AllowedPackageNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](as.ref.Append("allowed_package_names"))
}

type IosSettingsAttributes struct {
	ref terra.Reference
}

func (is IosSettingsAttributes) InternalRef() terra.Reference {
	return is.ref
}

func (is IosSettingsAttributes) InternalWithRef(ref terra.Reference) IosSettingsAttributes {
	return IosSettingsAttributes{ref: ref}
}

func (is IosSettingsAttributes) InternalTokens() hclwrite.Tokens {
	return is.ref.InternalTokens()
}

func (is IosSettingsAttributes) AllowAllBundleIds() terra.BoolValue {
	return terra.ReferenceBool(is.ref.Append("allow_all_bundle_ids"))
}

func (is IosSettingsAttributes) AllowedBundleIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](is.ref.Append("allowed_bundle_ids"))
}

type TestingOptionsAttributes struct {
	ref terra.Reference
}

func (to TestingOptionsAttributes) InternalRef() terra.Reference {
	return to.ref
}

func (to TestingOptionsAttributes) InternalWithRef(ref terra.Reference) TestingOptionsAttributes {
	return TestingOptionsAttributes{ref: ref}
}

func (to TestingOptionsAttributes) InternalTokens() hclwrite.Tokens {
	return to.ref.InternalTokens()
}

func (to TestingOptionsAttributes) TestingChallenge() terra.StringValue {
	return terra.ReferenceString(to.ref.Append("testing_challenge"))
}

func (to TestingOptionsAttributes) TestingScore() terra.NumberValue {
	return terra.ReferenceNumber(to.ref.Append("testing_score"))
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
	return terra.ReferenceString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("update"))
}

type WebSettingsAttributes struct {
	ref terra.Reference
}

func (ws WebSettingsAttributes) InternalRef() terra.Reference {
	return ws.ref
}

func (ws WebSettingsAttributes) InternalWithRef(ref terra.Reference) WebSettingsAttributes {
	return WebSettingsAttributes{ref: ref}
}

func (ws WebSettingsAttributes) InternalTokens() hclwrite.Tokens {
	return ws.ref.InternalTokens()
}

func (ws WebSettingsAttributes) AllowAllDomains() terra.BoolValue {
	return terra.ReferenceBool(ws.ref.Append("allow_all_domains"))
}

func (ws WebSettingsAttributes) AllowAmpTraffic() terra.BoolValue {
	return terra.ReferenceBool(ws.ref.Append("allow_amp_traffic"))
}

func (ws WebSettingsAttributes) AllowedDomains() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](ws.ref.Append("allowed_domains"))
}

func (ws WebSettingsAttributes) ChallengeSecurityPreference() terra.StringValue {
	return terra.ReferenceString(ws.ref.Append("challenge_security_preference"))
}

func (ws WebSettingsAttributes) IntegrationType() terra.StringValue {
	return terra.ReferenceString(ws.ref.Append("integration_type"))
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

type WebSettingsState struct {
	AllowAllDomains             bool     `json:"allow_all_domains"`
	AllowAmpTraffic             bool     `json:"allow_amp_traffic"`
	AllowedDomains              []string `json:"allowed_domains"`
	ChallengeSecurityPreference string   `json:"challenge_security_preference"`
	IntegrationType             string   `json:"integration_type"`
}
