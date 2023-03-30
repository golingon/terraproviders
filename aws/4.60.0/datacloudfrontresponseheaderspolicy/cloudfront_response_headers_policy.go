// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datacloudfrontresponseheaderspolicy

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type CorsConfig struct {
	// AccessControlAllowHeaders: min=0
	AccessControlAllowHeaders []AccessControlAllowHeaders `hcl:"access_control_allow_headers,block" validate:"min=0"`
	// AccessControlAllowMethods: min=0
	AccessControlAllowMethods []AccessControlAllowMethods `hcl:"access_control_allow_methods,block" validate:"min=0"`
	// AccessControlAllowOrigins: min=0
	AccessControlAllowOrigins []AccessControlAllowOrigins `hcl:"access_control_allow_origins,block" validate:"min=0"`
	// AccessControlExposeHeaders: min=0
	AccessControlExposeHeaders []AccessControlExposeHeaders `hcl:"access_control_expose_headers,block" validate:"min=0"`
}

type AccessControlAllowHeaders struct{}

type AccessControlAllowMethods struct{}

type AccessControlAllowOrigins struct{}

type AccessControlExposeHeaders struct{}

type CustomHeadersConfig struct {
	// Items: min=0
	Items []Items `hcl:"items,block" validate:"min=0"`
}

type Items struct{}

type SecurityHeadersConfig struct {
	// ContentSecurityPolicy: min=0
	ContentSecurityPolicy []ContentSecurityPolicy `hcl:"content_security_policy,block" validate:"min=0"`
	// ContentTypeOptions: min=0
	ContentTypeOptions []ContentTypeOptions `hcl:"content_type_options,block" validate:"min=0"`
	// FrameOptions: min=0
	FrameOptions []FrameOptions `hcl:"frame_options,block" validate:"min=0"`
	// ReferrerPolicy: min=0
	ReferrerPolicy []ReferrerPolicy `hcl:"referrer_policy,block" validate:"min=0"`
	// StrictTransportSecurity: min=0
	StrictTransportSecurity []StrictTransportSecurity `hcl:"strict_transport_security,block" validate:"min=0"`
	// XssProtection: min=0
	XssProtection []XssProtection `hcl:"xss_protection,block" validate:"min=0"`
}

type ContentSecurityPolicy struct{}

type ContentTypeOptions struct{}

type FrameOptions struct{}

type ReferrerPolicy struct{}

type StrictTransportSecurity struct{}

type XssProtection struct{}

type ServerTimingHeadersConfig struct{}

type CorsConfigAttributes struct {
	ref terra.Reference
}

func (cc CorsConfigAttributes) InternalRef() terra.Reference {
	return cc.ref
}

func (cc CorsConfigAttributes) InternalWithRef(ref terra.Reference) CorsConfigAttributes {
	return CorsConfigAttributes{ref: ref}
}

func (cc CorsConfigAttributes) InternalTokens() hclwrite.Tokens {
	return cc.ref.InternalTokens()
}

func (cc CorsConfigAttributes) AccessControlAllowCredentials() terra.BoolValue {
	return terra.ReferenceBool(cc.ref.Append("access_control_allow_credentials"))
}

func (cc CorsConfigAttributes) AccessControlMaxAgeSec() terra.NumberValue {
	return terra.ReferenceNumber(cc.ref.Append("access_control_max_age_sec"))
}

func (cc CorsConfigAttributes) OriginOverride() terra.BoolValue {
	return terra.ReferenceBool(cc.ref.Append("origin_override"))
}

func (cc CorsConfigAttributes) AccessControlAllowHeaders() terra.ListValue[AccessControlAllowHeadersAttributes] {
	return terra.ReferenceList[AccessControlAllowHeadersAttributes](cc.ref.Append("access_control_allow_headers"))
}

func (cc CorsConfigAttributes) AccessControlAllowMethods() terra.ListValue[AccessControlAllowMethodsAttributes] {
	return terra.ReferenceList[AccessControlAllowMethodsAttributes](cc.ref.Append("access_control_allow_methods"))
}

func (cc CorsConfigAttributes) AccessControlAllowOrigins() terra.ListValue[AccessControlAllowOriginsAttributes] {
	return terra.ReferenceList[AccessControlAllowOriginsAttributes](cc.ref.Append("access_control_allow_origins"))
}

func (cc CorsConfigAttributes) AccessControlExposeHeaders() terra.ListValue[AccessControlExposeHeadersAttributes] {
	return terra.ReferenceList[AccessControlExposeHeadersAttributes](cc.ref.Append("access_control_expose_headers"))
}

type AccessControlAllowHeadersAttributes struct {
	ref terra.Reference
}

func (acah AccessControlAllowHeadersAttributes) InternalRef() terra.Reference {
	return acah.ref
}

func (acah AccessControlAllowHeadersAttributes) InternalWithRef(ref terra.Reference) AccessControlAllowHeadersAttributes {
	return AccessControlAllowHeadersAttributes{ref: ref}
}

func (acah AccessControlAllowHeadersAttributes) InternalTokens() hclwrite.Tokens {
	return acah.ref.InternalTokens()
}

func (acah AccessControlAllowHeadersAttributes) Items() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](acah.ref.Append("items"))
}

type AccessControlAllowMethodsAttributes struct {
	ref terra.Reference
}

func (acam AccessControlAllowMethodsAttributes) InternalRef() terra.Reference {
	return acam.ref
}

func (acam AccessControlAllowMethodsAttributes) InternalWithRef(ref terra.Reference) AccessControlAllowMethodsAttributes {
	return AccessControlAllowMethodsAttributes{ref: ref}
}

func (acam AccessControlAllowMethodsAttributes) InternalTokens() hclwrite.Tokens {
	return acam.ref.InternalTokens()
}

func (acam AccessControlAllowMethodsAttributes) Items() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](acam.ref.Append("items"))
}

type AccessControlAllowOriginsAttributes struct {
	ref terra.Reference
}

func (acao AccessControlAllowOriginsAttributes) InternalRef() terra.Reference {
	return acao.ref
}

func (acao AccessControlAllowOriginsAttributes) InternalWithRef(ref terra.Reference) AccessControlAllowOriginsAttributes {
	return AccessControlAllowOriginsAttributes{ref: ref}
}

func (acao AccessControlAllowOriginsAttributes) InternalTokens() hclwrite.Tokens {
	return acao.ref.InternalTokens()
}

func (acao AccessControlAllowOriginsAttributes) Items() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](acao.ref.Append("items"))
}

type AccessControlExposeHeadersAttributes struct {
	ref terra.Reference
}

func (aceh AccessControlExposeHeadersAttributes) InternalRef() terra.Reference {
	return aceh.ref
}

func (aceh AccessControlExposeHeadersAttributes) InternalWithRef(ref terra.Reference) AccessControlExposeHeadersAttributes {
	return AccessControlExposeHeadersAttributes{ref: ref}
}

func (aceh AccessControlExposeHeadersAttributes) InternalTokens() hclwrite.Tokens {
	return aceh.ref.InternalTokens()
}

func (aceh AccessControlExposeHeadersAttributes) Items() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](aceh.ref.Append("items"))
}

type CustomHeadersConfigAttributes struct {
	ref terra.Reference
}

func (chc CustomHeadersConfigAttributes) InternalRef() terra.Reference {
	return chc.ref
}

func (chc CustomHeadersConfigAttributes) InternalWithRef(ref terra.Reference) CustomHeadersConfigAttributes {
	return CustomHeadersConfigAttributes{ref: ref}
}

func (chc CustomHeadersConfigAttributes) InternalTokens() hclwrite.Tokens {
	return chc.ref.InternalTokens()
}

func (chc CustomHeadersConfigAttributes) Items() terra.SetValue[ItemsAttributes] {
	return terra.ReferenceSet[ItemsAttributes](chc.ref.Append("items"))
}

type ItemsAttributes struct {
	ref terra.Reference
}

func (i ItemsAttributes) InternalRef() terra.Reference {
	return i.ref
}

func (i ItemsAttributes) InternalWithRef(ref terra.Reference) ItemsAttributes {
	return ItemsAttributes{ref: ref}
}

func (i ItemsAttributes) InternalTokens() hclwrite.Tokens {
	return i.ref.InternalTokens()
}

func (i ItemsAttributes) Header() terra.StringValue {
	return terra.ReferenceString(i.ref.Append("header"))
}

func (i ItemsAttributes) Override() terra.BoolValue {
	return terra.ReferenceBool(i.ref.Append("override"))
}

func (i ItemsAttributes) Value() terra.StringValue {
	return terra.ReferenceString(i.ref.Append("value"))
}

type SecurityHeadersConfigAttributes struct {
	ref terra.Reference
}

func (shc SecurityHeadersConfigAttributes) InternalRef() terra.Reference {
	return shc.ref
}

func (shc SecurityHeadersConfigAttributes) InternalWithRef(ref terra.Reference) SecurityHeadersConfigAttributes {
	return SecurityHeadersConfigAttributes{ref: ref}
}

func (shc SecurityHeadersConfigAttributes) InternalTokens() hclwrite.Tokens {
	return shc.ref.InternalTokens()
}

func (shc SecurityHeadersConfigAttributes) ContentSecurityPolicy() terra.ListValue[ContentSecurityPolicyAttributes] {
	return terra.ReferenceList[ContentSecurityPolicyAttributes](shc.ref.Append("content_security_policy"))
}

func (shc SecurityHeadersConfigAttributes) ContentTypeOptions() terra.ListValue[ContentTypeOptionsAttributes] {
	return terra.ReferenceList[ContentTypeOptionsAttributes](shc.ref.Append("content_type_options"))
}

func (shc SecurityHeadersConfigAttributes) FrameOptions() terra.ListValue[FrameOptionsAttributes] {
	return terra.ReferenceList[FrameOptionsAttributes](shc.ref.Append("frame_options"))
}

func (shc SecurityHeadersConfigAttributes) ReferrerPolicy() terra.ListValue[ReferrerPolicyAttributes] {
	return terra.ReferenceList[ReferrerPolicyAttributes](shc.ref.Append("referrer_policy"))
}

func (shc SecurityHeadersConfigAttributes) StrictTransportSecurity() terra.ListValue[StrictTransportSecurityAttributes] {
	return terra.ReferenceList[StrictTransportSecurityAttributes](shc.ref.Append("strict_transport_security"))
}

func (shc SecurityHeadersConfigAttributes) XssProtection() terra.ListValue[XssProtectionAttributes] {
	return terra.ReferenceList[XssProtectionAttributes](shc.ref.Append("xss_protection"))
}

type ContentSecurityPolicyAttributes struct {
	ref terra.Reference
}

func (csp ContentSecurityPolicyAttributes) InternalRef() terra.Reference {
	return csp.ref
}

func (csp ContentSecurityPolicyAttributes) InternalWithRef(ref terra.Reference) ContentSecurityPolicyAttributes {
	return ContentSecurityPolicyAttributes{ref: ref}
}

func (csp ContentSecurityPolicyAttributes) InternalTokens() hclwrite.Tokens {
	return csp.ref.InternalTokens()
}

func (csp ContentSecurityPolicyAttributes) ContentSecurityPolicy() terra.StringValue {
	return terra.ReferenceString(csp.ref.Append("content_security_policy"))
}

func (csp ContentSecurityPolicyAttributes) Override() terra.BoolValue {
	return terra.ReferenceBool(csp.ref.Append("override"))
}

type ContentTypeOptionsAttributes struct {
	ref terra.Reference
}

func (cto ContentTypeOptionsAttributes) InternalRef() terra.Reference {
	return cto.ref
}

func (cto ContentTypeOptionsAttributes) InternalWithRef(ref terra.Reference) ContentTypeOptionsAttributes {
	return ContentTypeOptionsAttributes{ref: ref}
}

func (cto ContentTypeOptionsAttributes) InternalTokens() hclwrite.Tokens {
	return cto.ref.InternalTokens()
}

func (cto ContentTypeOptionsAttributes) Override() terra.BoolValue {
	return terra.ReferenceBool(cto.ref.Append("override"))
}

type FrameOptionsAttributes struct {
	ref terra.Reference
}

func (fo FrameOptionsAttributes) InternalRef() terra.Reference {
	return fo.ref
}

func (fo FrameOptionsAttributes) InternalWithRef(ref terra.Reference) FrameOptionsAttributes {
	return FrameOptionsAttributes{ref: ref}
}

func (fo FrameOptionsAttributes) InternalTokens() hclwrite.Tokens {
	return fo.ref.InternalTokens()
}

func (fo FrameOptionsAttributes) FrameOption() terra.StringValue {
	return terra.ReferenceString(fo.ref.Append("frame_option"))
}

func (fo FrameOptionsAttributes) Override() terra.BoolValue {
	return terra.ReferenceBool(fo.ref.Append("override"))
}

type ReferrerPolicyAttributes struct {
	ref terra.Reference
}

func (rp ReferrerPolicyAttributes) InternalRef() terra.Reference {
	return rp.ref
}

func (rp ReferrerPolicyAttributes) InternalWithRef(ref terra.Reference) ReferrerPolicyAttributes {
	return ReferrerPolicyAttributes{ref: ref}
}

func (rp ReferrerPolicyAttributes) InternalTokens() hclwrite.Tokens {
	return rp.ref.InternalTokens()
}

func (rp ReferrerPolicyAttributes) Override() terra.BoolValue {
	return terra.ReferenceBool(rp.ref.Append("override"))
}

func (rp ReferrerPolicyAttributes) ReferrerPolicy() terra.StringValue {
	return terra.ReferenceString(rp.ref.Append("referrer_policy"))
}

type StrictTransportSecurityAttributes struct {
	ref terra.Reference
}

func (sts StrictTransportSecurityAttributes) InternalRef() terra.Reference {
	return sts.ref
}

func (sts StrictTransportSecurityAttributes) InternalWithRef(ref terra.Reference) StrictTransportSecurityAttributes {
	return StrictTransportSecurityAttributes{ref: ref}
}

func (sts StrictTransportSecurityAttributes) InternalTokens() hclwrite.Tokens {
	return sts.ref.InternalTokens()
}

func (sts StrictTransportSecurityAttributes) AccessControlMaxAgeSec() terra.NumberValue {
	return terra.ReferenceNumber(sts.ref.Append("access_control_max_age_sec"))
}

func (sts StrictTransportSecurityAttributes) IncludeSubdomains() terra.BoolValue {
	return terra.ReferenceBool(sts.ref.Append("include_subdomains"))
}

func (sts StrictTransportSecurityAttributes) Override() terra.BoolValue {
	return terra.ReferenceBool(sts.ref.Append("override"))
}

func (sts StrictTransportSecurityAttributes) Preload() terra.BoolValue {
	return terra.ReferenceBool(sts.ref.Append("preload"))
}

type XssProtectionAttributes struct {
	ref terra.Reference
}

func (xp XssProtectionAttributes) InternalRef() terra.Reference {
	return xp.ref
}

func (xp XssProtectionAttributes) InternalWithRef(ref terra.Reference) XssProtectionAttributes {
	return XssProtectionAttributes{ref: ref}
}

func (xp XssProtectionAttributes) InternalTokens() hclwrite.Tokens {
	return xp.ref.InternalTokens()
}

func (xp XssProtectionAttributes) ModeBlock() terra.BoolValue {
	return terra.ReferenceBool(xp.ref.Append("mode_block"))
}

func (xp XssProtectionAttributes) Override() terra.BoolValue {
	return terra.ReferenceBool(xp.ref.Append("override"))
}

func (xp XssProtectionAttributes) Protection() terra.BoolValue {
	return terra.ReferenceBool(xp.ref.Append("protection"))
}

func (xp XssProtectionAttributes) ReportUri() terra.StringValue {
	return terra.ReferenceString(xp.ref.Append("report_uri"))
}

type ServerTimingHeadersConfigAttributes struct {
	ref terra.Reference
}

func (sthc ServerTimingHeadersConfigAttributes) InternalRef() terra.Reference {
	return sthc.ref
}

func (sthc ServerTimingHeadersConfigAttributes) InternalWithRef(ref terra.Reference) ServerTimingHeadersConfigAttributes {
	return ServerTimingHeadersConfigAttributes{ref: ref}
}

func (sthc ServerTimingHeadersConfigAttributes) InternalTokens() hclwrite.Tokens {
	return sthc.ref.InternalTokens()
}

func (sthc ServerTimingHeadersConfigAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceBool(sthc.ref.Append("enabled"))
}

func (sthc ServerTimingHeadersConfigAttributes) SamplingRate() terra.NumberValue {
	return terra.ReferenceNumber(sthc.ref.Append("sampling_rate"))
}

type CorsConfigState struct {
	AccessControlAllowCredentials bool                              `json:"access_control_allow_credentials"`
	AccessControlMaxAgeSec        float64                           `json:"access_control_max_age_sec"`
	OriginOverride                bool                              `json:"origin_override"`
	AccessControlAllowHeaders     []AccessControlAllowHeadersState  `json:"access_control_allow_headers"`
	AccessControlAllowMethods     []AccessControlAllowMethodsState  `json:"access_control_allow_methods"`
	AccessControlAllowOrigins     []AccessControlAllowOriginsState  `json:"access_control_allow_origins"`
	AccessControlExposeHeaders    []AccessControlExposeHeadersState `json:"access_control_expose_headers"`
}

type AccessControlAllowHeadersState struct {
	Items []string `json:"items"`
}

type AccessControlAllowMethodsState struct {
	Items []string `json:"items"`
}

type AccessControlAllowOriginsState struct {
	Items []string `json:"items"`
}

type AccessControlExposeHeadersState struct {
	Items []string `json:"items"`
}

type CustomHeadersConfigState struct {
	Items []ItemsState `json:"items"`
}

type ItemsState struct {
	Header   string `json:"header"`
	Override bool   `json:"override"`
	Value    string `json:"value"`
}

type SecurityHeadersConfigState struct {
	ContentSecurityPolicy   []ContentSecurityPolicyState   `json:"content_security_policy"`
	ContentTypeOptions      []ContentTypeOptionsState      `json:"content_type_options"`
	FrameOptions            []FrameOptionsState            `json:"frame_options"`
	ReferrerPolicy          []ReferrerPolicyState          `json:"referrer_policy"`
	StrictTransportSecurity []StrictTransportSecurityState `json:"strict_transport_security"`
	XssProtection           []XssProtectionState           `json:"xss_protection"`
}

type ContentSecurityPolicyState struct {
	ContentSecurityPolicy string `json:"content_security_policy"`
	Override              bool   `json:"override"`
}

type ContentTypeOptionsState struct {
	Override bool `json:"override"`
}

type FrameOptionsState struct {
	FrameOption string `json:"frame_option"`
	Override    bool   `json:"override"`
}

type ReferrerPolicyState struct {
	Override       bool   `json:"override"`
	ReferrerPolicy string `json:"referrer_policy"`
}

type StrictTransportSecurityState struct {
	AccessControlMaxAgeSec float64 `json:"access_control_max_age_sec"`
	IncludeSubdomains      bool    `json:"include_subdomains"`
	Override               bool    `json:"override"`
	Preload                bool    `json:"preload"`
}

type XssProtectionState struct {
	ModeBlock  bool   `json:"mode_block"`
	Override   bool   `json:"override"`
	Protection bool   `json:"protection"`
	ReportUri  string `json:"report_uri"`
}

type ServerTimingHeadersConfigState struct {
	Enabled      bool    `json:"enabled"`
	SamplingRate float64 `json:"sampling_rate"`
}
