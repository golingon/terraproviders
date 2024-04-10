// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package datacloudfrontresponseheaderspolicy

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
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
	// CustomHeadersConfigItems: min=0
	Items []CustomHeadersConfigItems `hcl:"items,block" validate:"min=0"`
}

type CustomHeadersConfigItems struct{}

type RemoveHeadersConfig struct {
	// RemoveHeadersConfigItems: min=0
	Items []RemoveHeadersConfigItems `hcl:"items,block" validate:"min=0"`
}

type RemoveHeadersConfigItems struct{}

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

func (cc CorsConfigAttributes) InternalRef() (terra.Reference, error) {
	return cc.ref, nil
}

func (cc CorsConfigAttributes) InternalWithRef(ref terra.Reference) CorsConfigAttributes {
	return CorsConfigAttributes{ref: ref}
}

func (cc CorsConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cc.ref.InternalTokens()
}

func (cc CorsConfigAttributes) AccessControlAllowCredentials() terra.BoolValue {
	return terra.ReferenceAsBool(cc.ref.Append("access_control_allow_credentials"))
}

func (cc CorsConfigAttributes) AccessControlMaxAgeSec() terra.NumberValue {
	return terra.ReferenceAsNumber(cc.ref.Append("access_control_max_age_sec"))
}

func (cc CorsConfigAttributes) OriginOverride() terra.BoolValue {
	return terra.ReferenceAsBool(cc.ref.Append("origin_override"))
}

func (cc CorsConfigAttributes) AccessControlAllowHeaders() terra.ListValue[AccessControlAllowHeadersAttributes] {
	return terra.ReferenceAsList[AccessControlAllowHeadersAttributes](cc.ref.Append("access_control_allow_headers"))
}

func (cc CorsConfigAttributes) AccessControlAllowMethods() terra.ListValue[AccessControlAllowMethodsAttributes] {
	return terra.ReferenceAsList[AccessControlAllowMethodsAttributes](cc.ref.Append("access_control_allow_methods"))
}

func (cc CorsConfigAttributes) AccessControlAllowOrigins() terra.ListValue[AccessControlAllowOriginsAttributes] {
	return terra.ReferenceAsList[AccessControlAllowOriginsAttributes](cc.ref.Append("access_control_allow_origins"))
}

func (cc CorsConfigAttributes) AccessControlExposeHeaders() terra.ListValue[AccessControlExposeHeadersAttributes] {
	return terra.ReferenceAsList[AccessControlExposeHeadersAttributes](cc.ref.Append("access_control_expose_headers"))
}

type AccessControlAllowHeadersAttributes struct {
	ref terra.Reference
}

func (acah AccessControlAllowHeadersAttributes) InternalRef() (terra.Reference, error) {
	return acah.ref, nil
}

func (acah AccessControlAllowHeadersAttributes) InternalWithRef(ref terra.Reference) AccessControlAllowHeadersAttributes {
	return AccessControlAllowHeadersAttributes{ref: ref}
}

func (acah AccessControlAllowHeadersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return acah.ref.InternalTokens()
}

func (acah AccessControlAllowHeadersAttributes) Items() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](acah.ref.Append("items"))
}

type AccessControlAllowMethodsAttributes struct {
	ref terra.Reference
}

func (acam AccessControlAllowMethodsAttributes) InternalRef() (terra.Reference, error) {
	return acam.ref, nil
}

func (acam AccessControlAllowMethodsAttributes) InternalWithRef(ref terra.Reference) AccessControlAllowMethodsAttributes {
	return AccessControlAllowMethodsAttributes{ref: ref}
}

func (acam AccessControlAllowMethodsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return acam.ref.InternalTokens()
}

func (acam AccessControlAllowMethodsAttributes) Items() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](acam.ref.Append("items"))
}

type AccessControlAllowOriginsAttributes struct {
	ref terra.Reference
}

func (acao AccessControlAllowOriginsAttributes) InternalRef() (terra.Reference, error) {
	return acao.ref, nil
}

func (acao AccessControlAllowOriginsAttributes) InternalWithRef(ref terra.Reference) AccessControlAllowOriginsAttributes {
	return AccessControlAllowOriginsAttributes{ref: ref}
}

func (acao AccessControlAllowOriginsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return acao.ref.InternalTokens()
}

func (acao AccessControlAllowOriginsAttributes) Items() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](acao.ref.Append("items"))
}

type AccessControlExposeHeadersAttributes struct {
	ref terra.Reference
}

func (aceh AccessControlExposeHeadersAttributes) InternalRef() (terra.Reference, error) {
	return aceh.ref, nil
}

func (aceh AccessControlExposeHeadersAttributes) InternalWithRef(ref terra.Reference) AccessControlExposeHeadersAttributes {
	return AccessControlExposeHeadersAttributes{ref: ref}
}

func (aceh AccessControlExposeHeadersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return aceh.ref.InternalTokens()
}

func (aceh AccessControlExposeHeadersAttributes) Items() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](aceh.ref.Append("items"))
}

type CustomHeadersConfigAttributes struct {
	ref terra.Reference
}

func (chc CustomHeadersConfigAttributes) InternalRef() (terra.Reference, error) {
	return chc.ref, nil
}

func (chc CustomHeadersConfigAttributes) InternalWithRef(ref terra.Reference) CustomHeadersConfigAttributes {
	return CustomHeadersConfigAttributes{ref: ref}
}

func (chc CustomHeadersConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return chc.ref.InternalTokens()
}

func (chc CustomHeadersConfigAttributes) Items() terra.SetValue[CustomHeadersConfigItemsAttributes] {
	return terra.ReferenceAsSet[CustomHeadersConfigItemsAttributes](chc.ref.Append("items"))
}

type CustomHeadersConfigItemsAttributes struct {
	ref terra.Reference
}

func (i CustomHeadersConfigItemsAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i CustomHeadersConfigItemsAttributes) InternalWithRef(ref terra.Reference) CustomHeadersConfigItemsAttributes {
	return CustomHeadersConfigItemsAttributes{ref: ref}
}

func (i CustomHeadersConfigItemsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i CustomHeadersConfigItemsAttributes) Header() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("header"))
}

func (i CustomHeadersConfigItemsAttributes) Override() terra.BoolValue {
	return terra.ReferenceAsBool(i.ref.Append("override"))
}

func (i CustomHeadersConfigItemsAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("value"))
}

type RemoveHeadersConfigAttributes struct {
	ref terra.Reference
}

func (rhc RemoveHeadersConfigAttributes) InternalRef() (terra.Reference, error) {
	return rhc.ref, nil
}

func (rhc RemoveHeadersConfigAttributes) InternalWithRef(ref terra.Reference) RemoveHeadersConfigAttributes {
	return RemoveHeadersConfigAttributes{ref: ref}
}

func (rhc RemoveHeadersConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rhc.ref.InternalTokens()
}

func (rhc RemoveHeadersConfigAttributes) Items() terra.SetValue[RemoveHeadersConfigItemsAttributes] {
	return terra.ReferenceAsSet[RemoveHeadersConfigItemsAttributes](rhc.ref.Append("items"))
}

type RemoveHeadersConfigItemsAttributes struct {
	ref terra.Reference
}

func (i RemoveHeadersConfigItemsAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i RemoveHeadersConfigItemsAttributes) InternalWithRef(ref terra.Reference) RemoveHeadersConfigItemsAttributes {
	return RemoveHeadersConfigItemsAttributes{ref: ref}
}

func (i RemoveHeadersConfigItemsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i RemoveHeadersConfigItemsAttributes) Header() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("header"))
}

type SecurityHeadersConfigAttributes struct {
	ref terra.Reference
}

func (shc SecurityHeadersConfigAttributes) InternalRef() (terra.Reference, error) {
	return shc.ref, nil
}

func (shc SecurityHeadersConfigAttributes) InternalWithRef(ref terra.Reference) SecurityHeadersConfigAttributes {
	return SecurityHeadersConfigAttributes{ref: ref}
}

func (shc SecurityHeadersConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return shc.ref.InternalTokens()
}

func (shc SecurityHeadersConfigAttributes) ContentSecurityPolicy() terra.ListValue[ContentSecurityPolicyAttributes] {
	return terra.ReferenceAsList[ContentSecurityPolicyAttributes](shc.ref.Append("content_security_policy"))
}

func (shc SecurityHeadersConfigAttributes) ContentTypeOptions() terra.ListValue[ContentTypeOptionsAttributes] {
	return terra.ReferenceAsList[ContentTypeOptionsAttributes](shc.ref.Append("content_type_options"))
}

func (shc SecurityHeadersConfigAttributes) FrameOptions() terra.ListValue[FrameOptionsAttributes] {
	return terra.ReferenceAsList[FrameOptionsAttributes](shc.ref.Append("frame_options"))
}

func (shc SecurityHeadersConfigAttributes) ReferrerPolicy() terra.ListValue[ReferrerPolicyAttributes] {
	return terra.ReferenceAsList[ReferrerPolicyAttributes](shc.ref.Append("referrer_policy"))
}

func (shc SecurityHeadersConfigAttributes) StrictTransportSecurity() terra.ListValue[StrictTransportSecurityAttributes] {
	return terra.ReferenceAsList[StrictTransportSecurityAttributes](shc.ref.Append("strict_transport_security"))
}

func (shc SecurityHeadersConfigAttributes) XssProtection() terra.ListValue[XssProtectionAttributes] {
	return terra.ReferenceAsList[XssProtectionAttributes](shc.ref.Append("xss_protection"))
}

type ContentSecurityPolicyAttributes struct {
	ref terra.Reference
}

func (csp ContentSecurityPolicyAttributes) InternalRef() (terra.Reference, error) {
	return csp.ref, nil
}

func (csp ContentSecurityPolicyAttributes) InternalWithRef(ref terra.Reference) ContentSecurityPolicyAttributes {
	return ContentSecurityPolicyAttributes{ref: ref}
}

func (csp ContentSecurityPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return csp.ref.InternalTokens()
}

func (csp ContentSecurityPolicyAttributes) ContentSecurityPolicy() terra.StringValue {
	return terra.ReferenceAsString(csp.ref.Append("content_security_policy"))
}

func (csp ContentSecurityPolicyAttributes) Override() terra.BoolValue {
	return terra.ReferenceAsBool(csp.ref.Append("override"))
}

type ContentTypeOptionsAttributes struct {
	ref terra.Reference
}

func (cto ContentTypeOptionsAttributes) InternalRef() (terra.Reference, error) {
	return cto.ref, nil
}

func (cto ContentTypeOptionsAttributes) InternalWithRef(ref terra.Reference) ContentTypeOptionsAttributes {
	return ContentTypeOptionsAttributes{ref: ref}
}

func (cto ContentTypeOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cto.ref.InternalTokens()
}

func (cto ContentTypeOptionsAttributes) Override() terra.BoolValue {
	return terra.ReferenceAsBool(cto.ref.Append("override"))
}

type FrameOptionsAttributes struct {
	ref terra.Reference
}

func (fo FrameOptionsAttributes) InternalRef() (terra.Reference, error) {
	return fo.ref, nil
}

func (fo FrameOptionsAttributes) InternalWithRef(ref terra.Reference) FrameOptionsAttributes {
	return FrameOptionsAttributes{ref: ref}
}

func (fo FrameOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fo.ref.InternalTokens()
}

func (fo FrameOptionsAttributes) FrameOption() terra.StringValue {
	return terra.ReferenceAsString(fo.ref.Append("frame_option"))
}

func (fo FrameOptionsAttributes) Override() terra.BoolValue {
	return terra.ReferenceAsBool(fo.ref.Append("override"))
}

type ReferrerPolicyAttributes struct {
	ref terra.Reference
}

func (rp ReferrerPolicyAttributes) InternalRef() (terra.Reference, error) {
	return rp.ref, nil
}

func (rp ReferrerPolicyAttributes) InternalWithRef(ref terra.Reference) ReferrerPolicyAttributes {
	return ReferrerPolicyAttributes{ref: ref}
}

func (rp ReferrerPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rp.ref.InternalTokens()
}

func (rp ReferrerPolicyAttributes) Override() terra.BoolValue {
	return terra.ReferenceAsBool(rp.ref.Append("override"))
}

func (rp ReferrerPolicyAttributes) ReferrerPolicy() terra.StringValue {
	return terra.ReferenceAsString(rp.ref.Append("referrer_policy"))
}

type StrictTransportSecurityAttributes struct {
	ref terra.Reference
}

func (sts StrictTransportSecurityAttributes) InternalRef() (terra.Reference, error) {
	return sts.ref, nil
}

func (sts StrictTransportSecurityAttributes) InternalWithRef(ref terra.Reference) StrictTransportSecurityAttributes {
	return StrictTransportSecurityAttributes{ref: ref}
}

func (sts StrictTransportSecurityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sts.ref.InternalTokens()
}

func (sts StrictTransportSecurityAttributes) AccessControlMaxAgeSec() terra.NumberValue {
	return terra.ReferenceAsNumber(sts.ref.Append("access_control_max_age_sec"))
}

func (sts StrictTransportSecurityAttributes) IncludeSubdomains() terra.BoolValue {
	return terra.ReferenceAsBool(sts.ref.Append("include_subdomains"))
}

func (sts StrictTransportSecurityAttributes) Override() terra.BoolValue {
	return terra.ReferenceAsBool(sts.ref.Append("override"))
}

func (sts StrictTransportSecurityAttributes) Preload() terra.BoolValue {
	return terra.ReferenceAsBool(sts.ref.Append("preload"))
}

type XssProtectionAttributes struct {
	ref terra.Reference
}

func (xp XssProtectionAttributes) InternalRef() (terra.Reference, error) {
	return xp.ref, nil
}

func (xp XssProtectionAttributes) InternalWithRef(ref terra.Reference) XssProtectionAttributes {
	return XssProtectionAttributes{ref: ref}
}

func (xp XssProtectionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return xp.ref.InternalTokens()
}

func (xp XssProtectionAttributes) ModeBlock() terra.BoolValue {
	return terra.ReferenceAsBool(xp.ref.Append("mode_block"))
}

func (xp XssProtectionAttributes) Override() terra.BoolValue {
	return terra.ReferenceAsBool(xp.ref.Append("override"))
}

func (xp XssProtectionAttributes) Protection() terra.BoolValue {
	return terra.ReferenceAsBool(xp.ref.Append("protection"))
}

func (xp XssProtectionAttributes) ReportUri() terra.StringValue {
	return terra.ReferenceAsString(xp.ref.Append("report_uri"))
}

type ServerTimingHeadersConfigAttributes struct {
	ref terra.Reference
}

func (sthc ServerTimingHeadersConfigAttributes) InternalRef() (terra.Reference, error) {
	return sthc.ref, nil
}

func (sthc ServerTimingHeadersConfigAttributes) InternalWithRef(ref terra.Reference) ServerTimingHeadersConfigAttributes {
	return ServerTimingHeadersConfigAttributes{ref: ref}
}

func (sthc ServerTimingHeadersConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sthc.ref.InternalTokens()
}

func (sthc ServerTimingHeadersConfigAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(sthc.ref.Append("enabled"))
}

func (sthc ServerTimingHeadersConfigAttributes) SamplingRate() terra.NumberValue {
	return terra.ReferenceAsNumber(sthc.ref.Append("sampling_rate"))
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
	Items []CustomHeadersConfigItemsState `json:"items"`
}

type CustomHeadersConfigItemsState struct {
	Header   string `json:"header"`
	Override bool   `json:"override"`
	Value    string `json:"value"`
}

type RemoveHeadersConfigState struct {
	Items []RemoveHeadersConfigItemsState `json:"items"`
}

type RemoveHeadersConfigItemsState struct {
	Header string `json:"header"`
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
