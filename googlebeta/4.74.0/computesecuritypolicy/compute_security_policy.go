// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package computesecuritypolicy

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AdaptiveProtectionConfig struct {
	// AutoDeployConfig: optional
	AutoDeployConfig *AutoDeployConfig `hcl:"auto_deploy_config,block"`
	// Layer7DdosDefenseConfig: optional
	Layer7DdosDefenseConfig *Layer7DdosDefenseConfig `hcl:"layer_7_ddos_defense_config,block"`
}

type AutoDeployConfig struct {
	// ConfidenceThreshold: number, optional
	ConfidenceThreshold terra.NumberValue `hcl:"confidence_threshold,attr"`
	// ExpirationSec: number, optional
	ExpirationSec terra.NumberValue `hcl:"expiration_sec,attr"`
	// ImpactedBaselineThreshold: number, optional
	ImpactedBaselineThreshold terra.NumberValue `hcl:"impacted_baseline_threshold,attr"`
	// LoadThreshold: number, optional
	LoadThreshold terra.NumberValue `hcl:"load_threshold,attr"`
}

type Layer7DdosDefenseConfig struct {
	// Enable: bool, optional
	Enable terra.BoolValue `hcl:"enable,attr"`
	// RuleVisibility: string, optional
	RuleVisibility terra.StringValue `hcl:"rule_visibility,attr"`
}

type AdvancedOptionsConfig struct {
	// JsonParsing: string, optional
	JsonParsing terra.StringValue `hcl:"json_parsing,attr"`
	// LogLevel: string, optional
	LogLevel terra.StringValue `hcl:"log_level,attr"`
	// JsonCustomConfig: optional
	JsonCustomConfig *JsonCustomConfig `hcl:"json_custom_config,block"`
}

type JsonCustomConfig struct {
	// ContentTypes: set of string, required
	ContentTypes terra.SetValue[terra.StringValue] `hcl:"content_types,attr" validate:"required"`
}

type RecaptchaOptionsConfig struct {
	// RedirectSiteKey: string, required
	RedirectSiteKey terra.StringValue `hcl:"redirect_site_key,attr" validate:"required"`
}

type Rule struct {
	// Action: string, required
	Action terra.StringValue `hcl:"action,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Preview: bool, optional
	Preview terra.BoolValue `hcl:"preview,attr"`
	// Priority: number, required
	Priority terra.NumberValue `hcl:"priority,attr" validate:"required"`
	// HeaderAction: optional
	HeaderAction *HeaderAction `hcl:"header_action,block"`
	// Match: required
	Match *Match `hcl:"match,block" validate:"required"`
	// PreconfiguredWafConfig: optional
	PreconfiguredWafConfig *PreconfiguredWafConfig `hcl:"preconfigured_waf_config,block"`
	// RateLimitOptions: optional
	RateLimitOptions *RateLimitOptions `hcl:"rate_limit_options,block"`
	// RedirectOptions: optional
	RedirectOptions *RedirectOptions `hcl:"redirect_options,block"`
}

type HeaderAction struct {
	// RequestHeadersToAdds: min=1
	RequestHeadersToAdds []RequestHeadersToAdds `hcl:"request_headers_to_adds,block" validate:"min=1"`
}

type RequestHeadersToAdds struct {
	// HeaderName: string, required
	HeaderName terra.StringValue `hcl:"header_name,attr" validate:"required"`
	// HeaderValue: string, optional
	HeaderValue terra.StringValue `hcl:"header_value,attr"`
}

type Match struct {
	// VersionedExpr: string, optional
	VersionedExpr terra.StringValue `hcl:"versioned_expr,attr"`
	// Config: optional
	Config *Config `hcl:"config,block"`
	// Expr: optional
	Expr *Expr `hcl:"expr,block"`
}

type Config struct {
	// SrcIpRanges: set of string, required
	SrcIpRanges terra.SetValue[terra.StringValue] `hcl:"src_ip_ranges,attr" validate:"required"`
}

type Expr struct {
	// Expression: string, required
	Expression terra.StringValue `hcl:"expression,attr" validate:"required"`
}

type PreconfiguredWafConfig struct {
	// Exclusion: min=0
	Exclusion []Exclusion `hcl:"exclusion,block" validate:"min=0"`
}

type Exclusion struct {
	// TargetRuleIds: set of string, optional
	TargetRuleIds terra.SetValue[terra.StringValue] `hcl:"target_rule_ids,attr"`
	// TargetRuleSet: string, required
	TargetRuleSet terra.StringValue `hcl:"target_rule_set,attr" validate:"required"`
	// RequestCookie: min=0
	RequestCookie []RequestCookie `hcl:"request_cookie,block" validate:"min=0"`
	// RequestHeader: min=0
	RequestHeader []RequestHeader `hcl:"request_header,block" validate:"min=0"`
	// RequestQueryParam: min=0
	RequestQueryParam []RequestQueryParam `hcl:"request_query_param,block" validate:"min=0"`
	// RequestUri: min=0
	RequestUri []RequestUri `hcl:"request_uri,block" validate:"min=0"`
}

type RequestCookie struct {
	// Operator: string, required
	Operator terra.StringValue `hcl:"operator,attr" validate:"required"`
	// Value: string, optional
	Value terra.StringValue `hcl:"value,attr"`
}

type RequestHeader struct {
	// Operator: string, required
	Operator terra.StringValue `hcl:"operator,attr" validate:"required"`
	// Value: string, optional
	Value terra.StringValue `hcl:"value,attr"`
}

type RequestQueryParam struct {
	// Operator: string, required
	Operator terra.StringValue `hcl:"operator,attr" validate:"required"`
	// Value: string, optional
	Value terra.StringValue `hcl:"value,attr"`
}

type RequestUri struct {
	// Operator: string, required
	Operator terra.StringValue `hcl:"operator,attr" validate:"required"`
	// Value: string, optional
	Value terra.StringValue `hcl:"value,attr"`
}

type RateLimitOptions struct {
	// BanDurationSec: number, optional
	BanDurationSec terra.NumberValue `hcl:"ban_duration_sec,attr"`
	// ConformAction: string, required
	ConformAction terra.StringValue `hcl:"conform_action,attr" validate:"required"`
	// EnforceOnKey: string, optional
	EnforceOnKey terra.StringValue `hcl:"enforce_on_key,attr"`
	// EnforceOnKeyName: string, optional
	EnforceOnKeyName terra.StringValue `hcl:"enforce_on_key_name,attr"`
	// ExceedAction: string, required
	ExceedAction terra.StringValue `hcl:"exceed_action,attr" validate:"required"`
	// BanThreshold: optional
	BanThreshold *BanThreshold `hcl:"ban_threshold,block"`
	// EnforceOnKeyConfigs: min=0
	EnforceOnKeyConfigs []EnforceOnKeyConfigs `hcl:"enforce_on_key_configs,block" validate:"min=0"`
	// ExceedRedirectOptions: optional
	ExceedRedirectOptions *ExceedRedirectOptions `hcl:"exceed_redirect_options,block"`
	// RateLimitThreshold: required
	RateLimitThreshold *RateLimitThreshold `hcl:"rate_limit_threshold,block" validate:"required"`
}

type BanThreshold struct {
	// Count: number, required
	Count terra.NumberValue `hcl:"count,attr" validate:"required"`
	// IntervalSec: number, required
	IntervalSec terra.NumberValue `hcl:"interval_sec,attr" validate:"required"`
}

type EnforceOnKeyConfigs struct {
	// EnforceOnKeyName: string, optional
	EnforceOnKeyName terra.StringValue `hcl:"enforce_on_key_name,attr"`
	// EnforceOnKeyType: string, optional
	EnforceOnKeyType terra.StringValue `hcl:"enforce_on_key_type,attr"`
}

type ExceedRedirectOptions struct {
	// Target: string, optional
	Target terra.StringValue `hcl:"target,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type RateLimitThreshold struct {
	// Count: number, required
	Count terra.NumberValue `hcl:"count,attr" validate:"required"`
	// IntervalSec: number, required
	IntervalSec terra.NumberValue `hcl:"interval_sec,attr" validate:"required"`
}

type RedirectOptions struct {
	// Target: string, optional
	Target terra.StringValue `hcl:"target,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type AdaptiveProtectionConfigAttributes struct {
	ref terra.Reference
}

func (apc AdaptiveProtectionConfigAttributes) InternalRef() (terra.Reference, error) {
	return apc.ref, nil
}

func (apc AdaptiveProtectionConfigAttributes) InternalWithRef(ref terra.Reference) AdaptiveProtectionConfigAttributes {
	return AdaptiveProtectionConfigAttributes{ref: ref}
}

func (apc AdaptiveProtectionConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return apc.ref.InternalTokens()
}

func (apc AdaptiveProtectionConfigAttributes) AutoDeployConfig() terra.ListValue[AutoDeployConfigAttributes] {
	return terra.ReferenceAsList[AutoDeployConfigAttributes](apc.ref.Append("auto_deploy_config"))
}

func (apc AdaptiveProtectionConfigAttributes) Layer7DdosDefenseConfig() terra.ListValue[Layer7DdosDefenseConfigAttributes] {
	return terra.ReferenceAsList[Layer7DdosDefenseConfigAttributes](apc.ref.Append("layer_7_ddos_defense_config"))
}

type AutoDeployConfigAttributes struct {
	ref terra.Reference
}

func (adc AutoDeployConfigAttributes) InternalRef() (terra.Reference, error) {
	return adc.ref, nil
}

func (adc AutoDeployConfigAttributes) InternalWithRef(ref terra.Reference) AutoDeployConfigAttributes {
	return AutoDeployConfigAttributes{ref: ref}
}

func (adc AutoDeployConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return adc.ref.InternalTokens()
}

func (adc AutoDeployConfigAttributes) ConfidenceThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(adc.ref.Append("confidence_threshold"))
}

func (adc AutoDeployConfigAttributes) ExpirationSec() terra.NumberValue {
	return terra.ReferenceAsNumber(adc.ref.Append("expiration_sec"))
}

func (adc AutoDeployConfigAttributes) ImpactedBaselineThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(adc.ref.Append("impacted_baseline_threshold"))
}

func (adc AutoDeployConfigAttributes) LoadThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(adc.ref.Append("load_threshold"))
}

type Layer7DdosDefenseConfigAttributes struct {
	ref terra.Reference
}

func (l7ddc Layer7DdosDefenseConfigAttributes) InternalRef() (terra.Reference, error) {
	return l7ddc.ref, nil
}

func (l7ddc Layer7DdosDefenseConfigAttributes) InternalWithRef(ref terra.Reference) Layer7DdosDefenseConfigAttributes {
	return Layer7DdosDefenseConfigAttributes{ref: ref}
}

func (l7ddc Layer7DdosDefenseConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return l7ddc.ref.InternalTokens()
}

func (l7ddc Layer7DdosDefenseConfigAttributes) Enable() terra.BoolValue {
	return terra.ReferenceAsBool(l7ddc.ref.Append("enable"))
}

func (l7ddc Layer7DdosDefenseConfigAttributes) RuleVisibility() terra.StringValue {
	return terra.ReferenceAsString(l7ddc.ref.Append("rule_visibility"))
}

type AdvancedOptionsConfigAttributes struct {
	ref terra.Reference
}

func (aoc AdvancedOptionsConfigAttributes) InternalRef() (terra.Reference, error) {
	return aoc.ref, nil
}

func (aoc AdvancedOptionsConfigAttributes) InternalWithRef(ref terra.Reference) AdvancedOptionsConfigAttributes {
	return AdvancedOptionsConfigAttributes{ref: ref}
}

func (aoc AdvancedOptionsConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return aoc.ref.InternalTokens()
}

func (aoc AdvancedOptionsConfigAttributes) JsonParsing() terra.StringValue {
	return terra.ReferenceAsString(aoc.ref.Append("json_parsing"))
}

func (aoc AdvancedOptionsConfigAttributes) LogLevel() terra.StringValue {
	return terra.ReferenceAsString(aoc.ref.Append("log_level"))
}

func (aoc AdvancedOptionsConfigAttributes) JsonCustomConfig() terra.ListValue[JsonCustomConfigAttributes] {
	return terra.ReferenceAsList[JsonCustomConfigAttributes](aoc.ref.Append("json_custom_config"))
}

type JsonCustomConfigAttributes struct {
	ref terra.Reference
}

func (jcc JsonCustomConfigAttributes) InternalRef() (terra.Reference, error) {
	return jcc.ref, nil
}

func (jcc JsonCustomConfigAttributes) InternalWithRef(ref terra.Reference) JsonCustomConfigAttributes {
	return JsonCustomConfigAttributes{ref: ref}
}

func (jcc JsonCustomConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return jcc.ref.InternalTokens()
}

func (jcc JsonCustomConfigAttributes) ContentTypes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](jcc.ref.Append("content_types"))
}

type RecaptchaOptionsConfigAttributes struct {
	ref terra.Reference
}

func (roc RecaptchaOptionsConfigAttributes) InternalRef() (terra.Reference, error) {
	return roc.ref, nil
}

func (roc RecaptchaOptionsConfigAttributes) InternalWithRef(ref terra.Reference) RecaptchaOptionsConfigAttributes {
	return RecaptchaOptionsConfigAttributes{ref: ref}
}

func (roc RecaptchaOptionsConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return roc.ref.InternalTokens()
}

func (roc RecaptchaOptionsConfigAttributes) RedirectSiteKey() terra.StringValue {
	return terra.ReferenceAsString(roc.ref.Append("redirect_site_key"))
}

type RuleAttributes struct {
	ref terra.Reference
}

func (r RuleAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RuleAttributes) InternalWithRef(ref terra.Reference) RuleAttributes {
	return RuleAttributes{ref: ref}
}

func (r RuleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RuleAttributes) Action() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("action"))
}

func (r RuleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("description"))
}

func (r RuleAttributes) Preview() terra.BoolValue {
	return terra.ReferenceAsBool(r.ref.Append("preview"))
}

func (r RuleAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(r.ref.Append("priority"))
}

func (r RuleAttributes) HeaderAction() terra.ListValue[HeaderActionAttributes] {
	return terra.ReferenceAsList[HeaderActionAttributes](r.ref.Append("header_action"))
}

func (r RuleAttributes) Match() terra.ListValue[MatchAttributes] {
	return terra.ReferenceAsList[MatchAttributes](r.ref.Append("match"))
}

func (r RuleAttributes) PreconfiguredWafConfig() terra.ListValue[PreconfiguredWafConfigAttributes] {
	return terra.ReferenceAsList[PreconfiguredWafConfigAttributes](r.ref.Append("preconfigured_waf_config"))
}

func (r RuleAttributes) RateLimitOptions() terra.ListValue[RateLimitOptionsAttributes] {
	return terra.ReferenceAsList[RateLimitOptionsAttributes](r.ref.Append("rate_limit_options"))
}

func (r RuleAttributes) RedirectOptions() terra.ListValue[RedirectOptionsAttributes] {
	return terra.ReferenceAsList[RedirectOptionsAttributes](r.ref.Append("redirect_options"))
}

type HeaderActionAttributes struct {
	ref terra.Reference
}

func (ha HeaderActionAttributes) InternalRef() (terra.Reference, error) {
	return ha.ref, nil
}

func (ha HeaderActionAttributes) InternalWithRef(ref terra.Reference) HeaderActionAttributes {
	return HeaderActionAttributes{ref: ref}
}

func (ha HeaderActionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ha.ref.InternalTokens()
}

func (ha HeaderActionAttributes) RequestHeadersToAdds() terra.ListValue[RequestHeadersToAddsAttributes] {
	return terra.ReferenceAsList[RequestHeadersToAddsAttributes](ha.ref.Append("request_headers_to_adds"))
}

type RequestHeadersToAddsAttributes struct {
	ref terra.Reference
}

func (rhta RequestHeadersToAddsAttributes) InternalRef() (terra.Reference, error) {
	return rhta.ref, nil
}

func (rhta RequestHeadersToAddsAttributes) InternalWithRef(ref terra.Reference) RequestHeadersToAddsAttributes {
	return RequestHeadersToAddsAttributes{ref: ref}
}

func (rhta RequestHeadersToAddsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rhta.ref.InternalTokens()
}

func (rhta RequestHeadersToAddsAttributes) HeaderName() terra.StringValue {
	return terra.ReferenceAsString(rhta.ref.Append("header_name"))
}

func (rhta RequestHeadersToAddsAttributes) HeaderValue() terra.StringValue {
	return terra.ReferenceAsString(rhta.ref.Append("header_value"))
}

type MatchAttributes struct {
	ref terra.Reference
}

func (m MatchAttributes) InternalRef() (terra.Reference, error) {
	return m.ref, nil
}

func (m MatchAttributes) InternalWithRef(ref terra.Reference) MatchAttributes {
	return MatchAttributes{ref: ref}
}

func (m MatchAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return m.ref.InternalTokens()
}

func (m MatchAttributes) VersionedExpr() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("versioned_expr"))
}

func (m MatchAttributes) Config() terra.ListValue[ConfigAttributes] {
	return terra.ReferenceAsList[ConfigAttributes](m.ref.Append("config"))
}

func (m MatchAttributes) Expr() terra.ListValue[ExprAttributes] {
	return terra.ReferenceAsList[ExprAttributes](m.ref.Append("expr"))
}

type ConfigAttributes struct {
	ref terra.Reference
}

func (c ConfigAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ConfigAttributes) InternalWithRef(ref terra.Reference) ConfigAttributes {
	return ConfigAttributes{ref: ref}
}

func (c ConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ConfigAttributes) SrcIpRanges() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](c.ref.Append("src_ip_ranges"))
}

type ExprAttributes struct {
	ref terra.Reference
}

func (e ExprAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e ExprAttributes) InternalWithRef(ref terra.Reference) ExprAttributes {
	return ExprAttributes{ref: ref}
}

func (e ExprAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return e.ref.InternalTokens()
}

func (e ExprAttributes) Expression() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("expression"))
}

type PreconfiguredWafConfigAttributes struct {
	ref terra.Reference
}

func (pwc PreconfiguredWafConfigAttributes) InternalRef() (terra.Reference, error) {
	return pwc.ref, nil
}

func (pwc PreconfiguredWafConfigAttributes) InternalWithRef(ref terra.Reference) PreconfiguredWafConfigAttributes {
	return PreconfiguredWafConfigAttributes{ref: ref}
}

func (pwc PreconfiguredWafConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pwc.ref.InternalTokens()
}

func (pwc PreconfiguredWafConfigAttributes) Exclusion() terra.ListValue[ExclusionAttributes] {
	return terra.ReferenceAsList[ExclusionAttributes](pwc.ref.Append("exclusion"))
}

type ExclusionAttributes struct {
	ref terra.Reference
}

func (e ExclusionAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e ExclusionAttributes) InternalWithRef(ref terra.Reference) ExclusionAttributes {
	return ExclusionAttributes{ref: ref}
}

func (e ExclusionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return e.ref.InternalTokens()
}

func (e ExclusionAttributes) TargetRuleIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](e.ref.Append("target_rule_ids"))
}

func (e ExclusionAttributes) TargetRuleSet() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("target_rule_set"))
}

func (e ExclusionAttributes) RequestCookie() terra.ListValue[RequestCookieAttributes] {
	return terra.ReferenceAsList[RequestCookieAttributes](e.ref.Append("request_cookie"))
}

func (e ExclusionAttributes) RequestHeader() terra.ListValue[RequestHeaderAttributes] {
	return terra.ReferenceAsList[RequestHeaderAttributes](e.ref.Append("request_header"))
}

func (e ExclusionAttributes) RequestQueryParam() terra.ListValue[RequestQueryParamAttributes] {
	return terra.ReferenceAsList[RequestQueryParamAttributes](e.ref.Append("request_query_param"))
}

func (e ExclusionAttributes) RequestUri() terra.ListValue[RequestUriAttributes] {
	return terra.ReferenceAsList[RequestUriAttributes](e.ref.Append("request_uri"))
}

type RequestCookieAttributes struct {
	ref terra.Reference
}

func (rc RequestCookieAttributes) InternalRef() (terra.Reference, error) {
	return rc.ref, nil
}

func (rc RequestCookieAttributes) InternalWithRef(ref terra.Reference) RequestCookieAttributes {
	return RequestCookieAttributes{ref: ref}
}

func (rc RequestCookieAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rc.ref.InternalTokens()
}

func (rc RequestCookieAttributes) Operator() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("operator"))
}

func (rc RequestCookieAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("value"))
}

type RequestHeaderAttributes struct {
	ref terra.Reference
}

func (rh RequestHeaderAttributes) InternalRef() (terra.Reference, error) {
	return rh.ref, nil
}

func (rh RequestHeaderAttributes) InternalWithRef(ref terra.Reference) RequestHeaderAttributes {
	return RequestHeaderAttributes{ref: ref}
}

func (rh RequestHeaderAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rh.ref.InternalTokens()
}

func (rh RequestHeaderAttributes) Operator() terra.StringValue {
	return terra.ReferenceAsString(rh.ref.Append("operator"))
}

func (rh RequestHeaderAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(rh.ref.Append("value"))
}

type RequestQueryParamAttributes struct {
	ref terra.Reference
}

func (rqp RequestQueryParamAttributes) InternalRef() (terra.Reference, error) {
	return rqp.ref, nil
}

func (rqp RequestQueryParamAttributes) InternalWithRef(ref terra.Reference) RequestQueryParamAttributes {
	return RequestQueryParamAttributes{ref: ref}
}

func (rqp RequestQueryParamAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rqp.ref.InternalTokens()
}

func (rqp RequestQueryParamAttributes) Operator() terra.StringValue {
	return terra.ReferenceAsString(rqp.ref.Append("operator"))
}

func (rqp RequestQueryParamAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(rqp.ref.Append("value"))
}

type RequestUriAttributes struct {
	ref terra.Reference
}

func (ru RequestUriAttributes) InternalRef() (terra.Reference, error) {
	return ru.ref, nil
}

func (ru RequestUriAttributes) InternalWithRef(ref terra.Reference) RequestUriAttributes {
	return RequestUriAttributes{ref: ref}
}

func (ru RequestUriAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ru.ref.InternalTokens()
}

func (ru RequestUriAttributes) Operator() terra.StringValue {
	return terra.ReferenceAsString(ru.ref.Append("operator"))
}

func (ru RequestUriAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(ru.ref.Append("value"))
}

type RateLimitOptionsAttributes struct {
	ref terra.Reference
}

func (rlo RateLimitOptionsAttributes) InternalRef() (terra.Reference, error) {
	return rlo.ref, nil
}

func (rlo RateLimitOptionsAttributes) InternalWithRef(ref terra.Reference) RateLimitOptionsAttributes {
	return RateLimitOptionsAttributes{ref: ref}
}

func (rlo RateLimitOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rlo.ref.InternalTokens()
}

func (rlo RateLimitOptionsAttributes) BanDurationSec() terra.NumberValue {
	return terra.ReferenceAsNumber(rlo.ref.Append("ban_duration_sec"))
}

func (rlo RateLimitOptionsAttributes) ConformAction() terra.StringValue {
	return terra.ReferenceAsString(rlo.ref.Append("conform_action"))
}

func (rlo RateLimitOptionsAttributes) EnforceOnKey() terra.StringValue {
	return terra.ReferenceAsString(rlo.ref.Append("enforce_on_key"))
}

func (rlo RateLimitOptionsAttributes) EnforceOnKeyName() terra.StringValue {
	return terra.ReferenceAsString(rlo.ref.Append("enforce_on_key_name"))
}

func (rlo RateLimitOptionsAttributes) ExceedAction() terra.StringValue {
	return terra.ReferenceAsString(rlo.ref.Append("exceed_action"))
}

func (rlo RateLimitOptionsAttributes) BanThreshold() terra.ListValue[BanThresholdAttributes] {
	return terra.ReferenceAsList[BanThresholdAttributes](rlo.ref.Append("ban_threshold"))
}

func (rlo RateLimitOptionsAttributes) EnforceOnKeyConfigs() terra.ListValue[EnforceOnKeyConfigsAttributes] {
	return terra.ReferenceAsList[EnforceOnKeyConfigsAttributes](rlo.ref.Append("enforce_on_key_configs"))
}

func (rlo RateLimitOptionsAttributes) ExceedRedirectOptions() terra.ListValue[ExceedRedirectOptionsAttributes] {
	return terra.ReferenceAsList[ExceedRedirectOptionsAttributes](rlo.ref.Append("exceed_redirect_options"))
}

func (rlo RateLimitOptionsAttributes) RateLimitThreshold() terra.ListValue[RateLimitThresholdAttributes] {
	return terra.ReferenceAsList[RateLimitThresholdAttributes](rlo.ref.Append("rate_limit_threshold"))
}

type BanThresholdAttributes struct {
	ref terra.Reference
}

func (bt BanThresholdAttributes) InternalRef() (terra.Reference, error) {
	return bt.ref, nil
}

func (bt BanThresholdAttributes) InternalWithRef(ref terra.Reference) BanThresholdAttributes {
	return BanThresholdAttributes{ref: ref}
}

func (bt BanThresholdAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bt.ref.InternalTokens()
}

func (bt BanThresholdAttributes) Count() terra.NumberValue {
	return terra.ReferenceAsNumber(bt.ref.Append("count"))
}

func (bt BanThresholdAttributes) IntervalSec() terra.NumberValue {
	return terra.ReferenceAsNumber(bt.ref.Append("interval_sec"))
}

type EnforceOnKeyConfigsAttributes struct {
	ref terra.Reference
}

func (eokc EnforceOnKeyConfigsAttributes) InternalRef() (terra.Reference, error) {
	return eokc.ref, nil
}

func (eokc EnforceOnKeyConfigsAttributes) InternalWithRef(ref terra.Reference) EnforceOnKeyConfigsAttributes {
	return EnforceOnKeyConfigsAttributes{ref: ref}
}

func (eokc EnforceOnKeyConfigsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return eokc.ref.InternalTokens()
}

func (eokc EnforceOnKeyConfigsAttributes) EnforceOnKeyName() terra.StringValue {
	return terra.ReferenceAsString(eokc.ref.Append("enforce_on_key_name"))
}

func (eokc EnforceOnKeyConfigsAttributes) EnforceOnKeyType() terra.StringValue {
	return terra.ReferenceAsString(eokc.ref.Append("enforce_on_key_type"))
}

type ExceedRedirectOptionsAttributes struct {
	ref terra.Reference
}

func (ero ExceedRedirectOptionsAttributes) InternalRef() (terra.Reference, error) {
	return ero.ref, nil
}

func (ero ExceedRedirectOptionsAttributes) InternalWithRef(ref terra.Reference) ExceedRedirectOptionsAttributes {
	return ExceedRedirectOptionsAttributes{ref: ref}
}

func (ero ExceedRedirectOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ero.ref.InternalTokens()
}

func (ero ExceedRedirectOptionsAttributes) Target() terra.StringValue {
	return terra.ReferenceAsString(ero.ref.Append("target"))
}

func (ero ExceedRedirectOptionsAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ero.ref.Append("type"))
}

type RateLimitThresholdAttributes struct {
	ref terra.Reference
}

func (rlt RateLimitThresholdAttributes) InternalRef() (terra.Reference, error) {
	return rlt.ref, nil
}

func (rlt RateLimitThresholdAttributes) InternalWithRef(ref terra.Reference) RateLimitThresholdAttributes {
	return RateLimitThresholdAttributes{ref: ref}
}

func (rlt RateLimitThresholdAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rlt.ref.InternalTokens()
}

func (rlt RateLimitThresholdAttributes) Count() terra.NumberValue {
	return terra.ReferenceAsNumber(rlt.ref.Append("count"))
}

func (rlt RateLimitThresholdAttributes) IntervalSec() terra.NumberValue {
	return terra.ReferenceAsNumber(rlt.ref.Append("interval_sec"))
}

type RedirectOptionsAttributes struct {
	ref terra.Reference
}

func (ro RedirectOptionsAttributes) InternalRef() (terra.Reference, error) {
	return ro.ref, nil
}

func (ro RedirectOptionsAttributes) InternalWithRef(ref terra.Reference) RedirectOptionsAttributes {
	return RedirectOptionsAttributes{ref: ref}
}

func (ro RedirectOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ro.ref.InternalTokens()
}

func (ro RedirectOptionsAttributes) Target() terra.StringValue {
	return terra.ReferenceAsString(ro.ref.Append("target"))
}

func (ro RedirectOptionsAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ro.ref.Append("type"))
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

type AdaptiveProtectionConfigState struct {
	AutoDeployConfig        []AutoDeployConfigState        `json:"auto_deploy_config"`
	Layer7DdosDefenseConfig []Layer7DdosDefenseConfigState `json:"layer_7_ddos_defense_config"`
}

type AutoDeployConfigState struct {
	ConfidenceThreshold       float64 `json:"confidence_threshold"`
	ExpirationSec             float64 `json:"expiration_sec"`
	ImpactedBaselineThreshold float64 `json:"impacted_baseline_threshold"`
	LoadThreshold             float64 `json:"load_threshold"`
}

type Layer7DdosDefenseConfigState struct {
	Enable         bool   `json:"enable"`
	RuleVisibility string `json:"rule_visibility"`
}

type AdvancedOptionsConfigState struct {
	JsonParsing      string                  `json:"json_parsing"`
	LogLevel         string                  `json:"log_level"`
	JsonCustomConfig []JsonCustomConfigState `json:"json_custom_config"`
}

type JsonCustomConfigState struct {
	ContentTypes []string `json:"content_types"`
}

type RecaptchaOptionsConfigState struct {
	RedirectSiteKey string `json:"redirect_site_key"`
}

type RuleState struct {
	Action                 string                        `json:"action"`
	Description            string                        `json:"description"`
	Preview                bool                          `json:"preview"`
	Priority               float64                       `json:"priority"`
	HeaderAction           []HeaderActionState           `json:"header_action"`
	Match                  []MatchState                  `json:"match"`
	PreconfiguredWafConfig []PreconfiguredWafConfigState `json:"preconfigured_waf_config"`
	RateLimitOptions       []RateLimitOptionsState       `json:"rate_limit_options"`
	RedirectOptions        []RedirectOptionsState        `json:"redirect_options"`
}

type HeaderActionState struct {
	RequestHeadersToAdds []RequestHeadersToAddsState `json:"request_headers_to_adds"`
}

type RequestHeadersToAddsState struct {
	HeaderName  string `json:"header_name"`
	HeaderValue string `json:"header_value"`
}

type MatchState struct {
	VersionedExpr string        `json:"versioned_expr"`
	Config        []ConfigState `json:"config"`
	Expr          []ExprState   `json:"expr"`
}

type ConfigState struct {
	SrcIpRanges []string `json:"src_ip_ranges"`
}

type ExprState struct {
	Expression string `json:"expression"`
}

type PreconfiguredWafConfigState struct {
	Exclusion []ExclusionState `json:"exclusion"`
}

type ExclusionState struct {
	TargetRuleIds     []string                 `json:"target_rule_ids"`
	TargetRuleSet     string                   `json:"target_rule_set"`
	RequestCookie     []RequestCookieState     `json:"request_cookie"`
	RequestHeader     []RequestHeaderState     `json:"request_header"`
	RequestQueryParam []RequestQueryParamState `json:"request_query_param"`
	RequestUri        []RequestUriState        `json:"request_uri"`
}

type RequestCookieState struct {
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type RequestHeaderState struct {
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type RequestQueryParamState struct {
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type RequestUriState struct {
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type RateLimitOptionsState struct {
	BanDurationSec        float64                      `json:"ban_duration_sec"`
	ConformAction         string                       `json:"conform_action"`
	EnforceOnKey          string                       `json:"enforce_on_key"`
	EnforceOnKeyName      string                       `json:"enforce_on_key_name"`
	ExceedAction          string                       `json:"exceed_action"`
	BanThreshold          []BanThresholdState          `json:"ban_threshold"`
	EnforceOnKeyConfigs   []EnforceOnKeyConfigsState   `json:"enforce_on_key_configs"`
	ExceedRedirectOptions []ExceedRedirectOptionsState `json:"exceed_redirect_options"`
	RateLimitThreshold    []RateLimitThresholdState    `json:"rate_limit_threshold"`
}

type BanThresholdState struct {
	Count       float64 `json:"count"`
	IntervalSec float64 `json:"interval_sec"`
}

type EnforceOnKeyConfigsState struct {
	EnforceOnKeyName string `json:"enforce_on_key_name"`
	EnforceOnKeyType string `json:"enforce_on_key_type"`
}

type ExceedRedirectOptionsState struct {
	Target string `json:"target"`
	Type   string `json:"type"`
}

type RateLimitThresholdState struct {
	Count       float64 `json:"count"`
	IntervalSec float64 `json:"interval_sec"`
}

type RedirectOptionsState struct {
	Target string `json:"target"`
	Type   string `json:"type"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
