// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_alb_listener_rule

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Action struct {
	// Order: number, optional
	Order terra.NumberValue `hcl:"order,attr"`
	// TargetGroupArn: string, optional
	TargetGroupArn terra.StringValue `hcl:"target_group_arn,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// ActionAuthenticateCognito: optional
	AuthenticateCognito *ActionAuthenticateCognito `hcl:"authenticate_cognito,block"`
	// ActionAuthenticateOidc: optional
	AuthenticateOidc *ActionAuthenticateOidc `hcl:"authenticate_oidc,block"`
	// ActionFixedResponse: optional
	FixedResponse *ActionFixedResponse `hcl:"fixed_response,block"`
	// ActionForward: optional
	Forward *ActionForward `hcl:"forward,block"`
	// ActionRedirect: optional
	Redirect *ActionRedirect `hcl:"redirect,block"`
}

type ActionAuthenticateCognito struct {
	// AuthenticationRequestExtraParams: map of string, optional
	AuthenticationRequestExtraParams terra.MapValue[terra.StringValue] `hcl:"authentication_request_extra_params,attr"`
	// OnUnauthenticatedRequest: string, optional
	OnUnauthenticatedRequest terra.StringValue `hcl:"on_unauthenticated_request,attr"`
	// Scope: string, optional
	Scope terra.StringValue `hcl:"scope,attr"`
	// SessionCookieName: string, optional
	SessionCookieName terra.StringValue `hcl:"session_cookie_name,attr"`
	// SessionTimeout: number, optional
	SessionTimeout terra.NumberValue `hcl:"session_timeout,attr"`
	// UserPoolArn: string, required
	UserPoolArn terra.StringValue `hcl:"user_pool_arn,attr" validate:"required"`
	// UserPoolClientId: string, required
	UserPoolClientId terra.StringValue `hcl:"user_pool_client_id,attr" validate:"required"`
	// UserPoolDomain: string, required
	UserPoolDomain terra.StringValue `hcl:"user_pool_domain,attr" validate:"required"`
}

type ActionAuthenticateOidc struct {
	// AuthenticationRequestExtraParams: map of string, optional
	AuthenticationRequestExtraParams terra.MapValue[terra.StringValue] `hcl:"authentication_request_extra_params,attr"`
	// AuthorizationEndpoint: string, required
	AuthorizationEndpoint terra.StringValue `hcl:"authorization_endpoint,attr" validate:"required"`
	// ClientId: string, required
	ClientId terra.StringValue `hcl:"client_id,attr" validate:"required"`
	// ClientSecret: string, required
	ClientSecret terra.StringValue `hcl:"client_secret,attr" validate:"required"`
	// Issuer: string, required
	Issuer terra.StringValue `hcl:"issuer,attr" validate:"required"`
	// OnUnauthenticatedRequest: string, optional
	OnUnauthenticatedRequest terra.StringValue `hcl:"on_unauthenticated_request,attr"`
	// Scope: string, optional
	Scope terra.StringValue `hcl:"scope,attr"`
	// SessionCookieName: string, optional
	SessionCookieName terra.StringValue `hcl:"session_cookie_name,attr"`
	// SessionTimeout: number, optional
	SessionTimeout terra.NumberValue `hcl:"session_timeout,attr"`
	// TokenEndpoint: string, required
	TokenEndpoint terra.StringValue `hcl:"token_endpoint,attr" validate:"required"`
	// UserInfoEndpoint: string, required
	UserInfoEndpoint terra.StringValue `hcl:"user_info_endpoint,attr" validate:"required"`
}

type ActionFixedResponse struct {
	// ContentType: string, required
	ContentType terra.StringValue `hcl:"content_type,attr" validate:"required"`
	// MessageBody: string, optional
	MessageBody terra.StringValue `hcl:"message_body,attr"`
	// StatusCode: string, optional
	StatusCode terra.StringValue `hcl:"status_code,attr"`
}

type ActionForward struct {
	// ActionForwardStickiness: optional
	Stickiness *ActionForwardStickiness `hcl:"stickiness,block"`
	// ActionForwardTargetGroup: min=1,max=5
	TargetGroup []ActionForwardTargetGroup `hcl:"target_group,block" validate:"min=1,max=5"`
}

type ActionForwardStickiness struct {
	// Duration: number, required
	Duration terra.NumberValue `hcl:"duration,attr" validate:"required"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
}

type ActionForwardTargetGroup struct {
	// Arn: string, required
	Arn terra.StringValue `hcl:"arn,attr" validate:"required"`
	// Weight: number, optional
	Weight terra.NumberValue `hcl:"weight,attr"`
}

type ActionRedirect struct {
	// Host: string, optional
	Host terra.StringValue `hcl:"host,attr"`
	// Path: string, optional
	Path terra.StringValue `hcl:"path,attr"`
	// Port: string, optional
	Port terra.StringValue `hcl:"port,attr"`
	// Protocol: string, optional
	Protocol terra.StringValue `hcl:"protocol,attr"`
	// Query: string, optional
	Query terra.StringValue `hcl:"query,attr"`
	// StatusCode: string, required
	StatusCode terra.StringValue `hcl:"status_code,attr" validate:"required"`
}

type Condition struct {
	// ConditionHostHeader: optional
	HostHeader *ConditionHostHeader `hcl:"host_header,block"`
	// ConditionHttpHeader: optional
	HttpHeader *ConditionHttpHeader `hcl:"http_header,block"`
	// ConditionHttpRequestMethod: optional
	HttpRequestMethod *ConditionHttpRequestMethod `hcl:"http_request_method,block"`
	// ConditionPathPattern: optional
	PathPattern *ConditionPathPattern `hcl:"path_pattern,block"`
	// ConditionQueryString: min=0
	QueryString []ConditionQueryString `hcl:"query_string,block" validate:"min=0"`
	// ConditionSourceIp: optional
	SourceIp *ConditionSourceIp `hcl:"source_ip,block"`
}

type ConditionHostHeader struct {
	// Values: set of string, required
	Values terra.SetValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type ConditionHttpHeader struct {
	// HttpHeaderName: string, required
	HttpHeaderName terra.StringValue `hcl:"http_header_name,attr" validate:"required"`
	// Values: set of string, required
	Values terra.SetValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type ConditionHttpRequestMethod struct {
	// Values: set of string, required
	Values terra.SetValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type ConditionPathPattern struct {
	// Values: set of string, required
	Values terra.SetValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type ConditionQueryString struct {
	// Key: string, optional
	Key terra.StringValue `hcl:"key,attr"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type ConditionSourceIp struct {
	// Values: set of string, required
	Values terra.SetValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type ActionAttributes struct {
	ref terra.Reference
}

func (a ActionAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a ActionAttributes) InternalWithRef(ref terra.Reference) ActionAttributes {
	return ActionAttributes{ref: ref}
}

func (a ActionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a ActionAttributes) Order() terra.NumberValue {
	return terra.ReferenceAsNumber(a.ref.Append("order"))
}

func (a ActionAttributes) TargetGroupArn() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("target_group_arn"))
}

func (a ActionAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("type"))
}

func (a ActionAttributes) AuthenticateCognito() terra.ListValue[ActionAuthenticateCognitoAttributes] {
	return terra.ReferenceAsList[ActionAuthenticateCognitoAttributes](a.ref.Append("authenticate_cognito"))
}

func (a ActionAttributes) AuthenticateOidc() terra.ListValue[ActionAuthenticateOidcAttributes] {
	return terra.ReferenceAsList[ActionAuthenticateOidcAttributes](a.ref.Append("authenticate_oidc"))
}

func (a ActionAttributes) FixedResponse() terra.ListValue[ActionFixedResponseAttributes] {
	return terra.ReferenceAsList[ActionFixedResponseAttributes](a.ref.Append("fixed_response"))
}

func (a ActionAttributes) Forward() terra.ListValue[ActionForwardAttributes] {
	return terra.ReferenceAsList[ActionForwardAttributes](a.ref.Append("forward"))
}

func (a ActionAttributes) Redirect() terra.ListValue[ActionRedirectAttributes] {
	return terra.ReferenceAsList[ActionRedirectAttributes](a.ref.Append("redirect"))
}

type ActionAuthenticateCognitoAttributes struct {
	ref terra.Reference
}

func (ac ActionAuthenticateCognitoAttributes) InternalRef() (terra.Reference, error) {
	return ac.ref, nil
}

func (ac ActionAuthenticateCognitoAttributes) InternalWithRef(ref terra.Reference) ActionAuthenticateCognitoAttributes {
	return ActionAuthenticateCognitoAttributes{ref: ref}
}

func (ac ActionAuthenticateCognitoAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ac.ref.InternalTokens()
}

func (ac ActionAuthenticateCognitoAttributes) AuthenticationRequestExtraParams() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ac.ref.Append("authentication_request_extra_params"))
}

func (ac ActionAuthenticateCognitoAttributes) OnUnauthenticatedRequest() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("on_unauthenticated_request"))
}

func (ac ActionAuthenticateCognitoAttributes) Scope() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("scope"))
}

func (ac ActionAuthenticateCognitoAttributes) SessionCookieName() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("session_cookie_name"))
}

func (ac ActionAuthenticateCognitoAttributes) SessionTimeout() terra.NumberValue {
	return terra.ReferenceAsNumber(ac.ref.Append("session_timeout"))
}

func (ac ActionAuthenticateCognitoAttributes) UserPoolArn() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("user_pool_arn"))
}

func (ac ActionAuthenticateCognitoAttributes) UserPoolClientId() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("user_pool_client_id"))
}

func (ac ActionAuthenticateCognitoAttributes) UserPoolDomain() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("user_pool_domain"))
}

type ActionAuthenticateOidcAttributes struct {
	ref terra.Reference
}

func (ao ActionAuthenticateOidcAttributes) InternalRef() (terra.Reference, error) {
	return ao.ref, nil
}

func (ao ActionAuthenticateOidcAttributes) InternalWithRef(ref terra.Reference) ActionAuthenticateOidcAttributes {
	return ActionAuthenticateOidcAttributes{ref: ref}
}

func (ao ActionAuthenticateOidcAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ao.ref.InternalTokens()
}

func (ao ActionAuthenticateOidcAttributes) AuthenticationRequestExtraParams() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ao.ref.Append("authentication_request_extra_params"))
}

func (ao ActionAuthenticateOidcAttributes) AuthorizationEndpoint() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("authorization_endpoint"))
}

func (ao ActionAuthenticateOidcAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("client_id"))
}

func (ao ActionAuthenticateOidcAttributes) ClientSecret() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("client_secret"))
}

func (ao ActionAuthenticateOidcAttributes) Issuer() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("issuer"))
}

func (ao ActionAuthenticateOidcAttributes) OnUnauthenticatedRequest() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("on_unauthenticated_request"))
}

func (ao ActionAuthenticateOidcAttributes) Scope() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("scope"))
}

func (ao ActionAuthenticateOidcAttributes) SessionCookieName() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("session_cookie_name"))
}

func (ao ActionAuthenticateOidcAttributes) SessionTimeout() terra.NumberValue {
	return terra.ReferenceAsNumber(ao.ref.Append("session_timeout"))
}

func (ao ActionAuthenticateOidcAttributes) TokenEndpoint() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("token_endpoint"))
}

func (ao ActionAuthenticateOidcAttributes) UserInfoEndpoint() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("user_info_endpoint"))
}

type ActionFixedResponseAttributes struct {
	ref terra.Reference
}

func (fr ActionFixedResponseAttributes) InternalRef() (terra.Reference, error) {
	return fr.ref, nil
}

func (fr ActionFixedResponseAttributes) InternalWithRef(ref terra.Reference) ActionFixedResponseAttributes {
	return ActionFixedResponseAttributes{ref: ref}
}

func (fr ActionFixedResponseAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fr.ref.InternalTokens()
}

func (fr ActionFixedResponseAttributes) ContentType() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("content_type"))
}

func (fr ActionFixedResponseAttributes) MessageBody() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("message_body"))
}

func (fr ActionFixedResponseAttributes) StatusCode() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("status_code"))
}

type ActionForwardAttributes struct {
	ref terra.Reference
}

func (f ActionForwardAttributes) InternalRef() (terra.Reference, error) {
	return f.ref, nil
}

func (f ActionForwardAttributes) InternalWithRef(ref terra.Reference) ActionForwardAttributes {
	return ActionForwardAttributes{ref: ref}
}

func (f ActionForwardAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return f.ref.InternalTokens()
}

func (f ActionForwardAttributes) Stickiness() terra.ListValue[ActionForwardStickinessAttributes] {
	return terra.ReferenceAsList[ActionForwardStickinessAttributes](f.ref.Append("stickiness"))
}

func (f ActionForwardAttributes) TargetGroup() terra.SetValue[ActionForwardTargetGroupAttributes] {
	return terra.ReferenceAsSet[ActionForwardTargetGroupAttributes](f.ref.Append("target_group"))
}

type ActionForwardStickinessAttributes struct {
	ref terra.Reference
}

func (s ActionForwardStickinessAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s ActionForwardStickinessAttributes) InternalWithRef(ref terra.Reference) ActionForwardStickinessAttributes {
	return ActionForwardStickinessAttributes{ref: ref}
}

func (s ActionForwardStickinessAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s ActionForwardStickinessAttributes) Duration() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("duration"))
}

func (s ActionForwardStickinessAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("enabled"))
}

type ActionForwardTargetGroupAttributes struct {
	ref terra.Reference
}

func (tg ActionForwardTargetGroupAttributes) InternalRef() (terra.Reference, error) {
	return tg.ref, nil
}

func (tg ActionForwardTargetGroupAttributes) InternalWithRef(ref terra.Reference) ActionForwardTargetGroupAttributes {
	return ActionForwardTargetGroupAttributes{ref: ref}
}

func (tg ActionForwardTargetGroupAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tg.ref.InternalTokens()
}

func (tg ActionForwardTargetGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(tg.ref.Append("arn"))
}

func (tg ActionForwardTargetGroupAttributes) Weight() terra.NumberValue {
	return terra.ReferenceAsNumber(tg.ref.Append("weight"))
}

type ActionRedirectAttributes struct {
	ref terra.Reference
}

func (r ActionRedirectAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r ActionRedirectAttributes) InternalWithRef(ref terra.Reference) ActionRedirectAttributes {
	return ActionRedirectAttributes{ref: ref}
}

func (r ActionRedirectAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r ActionRedirectAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("host"))
}

func (r ActionRedirectAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("path"))
}

func (r ActionRedirectAttributes) Port() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("port"))
}

func (r ActionRedirectAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("protocol"))
}

func (r ActionRedirectAttributes) Query() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("query"))
}

func (r ActionRedirectAttributes) StatusCode() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("status_code"))
}

type ConditionAttributes struct {
	ref terra.Reference
}

func (c ConditionAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ConditionAttributes) InternalWithRef(ref terra.Reference) ConditionAttributes {
	return ConditionAttributes{ref: ref}
}

func (c ConditionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ConditionAttributes) HostHeader() terra.ListValue[ConditionHostHeaderAttributes] {
	return terra.ReferenceAsList[ConditionHostHeaderAttributes](c.ref.Append("host_header"))
}

func (c ConditionAttributes) HttpHeader() terra.ListValue[ConditionHttpHeaderAttributes] {
	return terra.ReferenceAsList[ConditionHttpHeaderAttributes](c.ref.Append("http_header"))
}

func (c ConditionAttributes) HttpRequestMethod() terra.ListValue[ConditionHttpRequestMethodAttributes] {
	return terra.ReferenceAsList[ConditionHttpRequestMethodAttributes](c.ref.Append("http_request_method"))
}

func (c ConditionAttributes) PathPattern() terra.ListValue[ConditionPathPatternAttributes] {
	return terra.ReferenceAsList[ConditionPathPatternAttributes](c.ref.Append("path_pattern"))
}

func (c ConditionAttributes) QueryString() terra.SetValue[ConditionQueryStringAttributes] {
	return terra.ReferenceAsSet[ConditionQueryStringAttributes](c.ref.Append("query_string"))
}

func (c ConditionAttributes) SourceIp() terra.ListValue[ConditionSourceIpAttributes] {
	return terra.ReferenceAsList[ConditionSourceIpAttributes](c.ref.Append("source_ip"))
}

type ConditionHostHeaderAttributes struct {
	ref terra.Reference
}

func (hh ConditionHostHeaderAttributes) InternalRef() (terra.Reference, error) {
	return hh.ref, nil
}

func (hh ConditionHostHeaderAttributes) InternalWithRef(ref terra.Reference) ConditionHostHeaderAttributes {
	return ConditionHostHeaderAttributes{ref: ref}
}

func (hh ConditionHostHeaderAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hh.ref.InternalTokens()
}

func (hh ConditionHostHeaderAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](hh.ref.Append("values"))
}

type ConditionHttpHeaderAttributes struct {
	ref terra.Reference
}

func (hh ConditionHttpHeaderAttributes) InternalRef() (terra.Reference, error) {
	return hh.ref, nil
}

func (hh ConditionHttpHeaderAttributes) InternalWithRef(ref terra.Reference) ConditionHttpHeaderAttributes {
	return ConditionHttpHeaderAttributes{ref: ref}
}

func (hh ConditionHttpHeaderAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hh.ref.InternalTokens()
}

func (hh ConditionHttpHeaderAttributes) HttpHeaderName() terra.StringValue {
	return terra.ReferenceAsString(hh.ref.Append("http_header_name"))
}

func (hh ConditionHttpHeaderAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](hh.ref.Append("values"))
}

type ConditionHttpRequestMethodAttributes struct {
	ref terra.Reference
}

func (hrm ConditionHttpRequestMethodAttributes) InternalRef() (terra.Reference, error) {
	return hrm.ref, nil
}

func (hrm ConditionHttpRequestMethodAttributes) InternalWithRef(ref terra.Reference) ConditionHttpRequestMethodAttributes {
	return ConditionHttpRequestMethodAttributes{ref: ref}
}

func (hrm ConditionHttpRequestMethodAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hrm.ref.InternalTokens()
}

func (hrm ConditionHttpRequestMethodAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](hrm.ref.Append("values"))
}

type ConditionPathPatternAttributes struct {
	ref terra.Reference
}

func (pp ConditionPathPatternAttributes) InternalRef() (terra.Reference, error) {
	return pp.ref, nil
}

func (pp ConditionPathPatternAttributes) InternalWithRef(ref terra.Reference) ConditionPathPatternAttributes {
	return ConditionPathPatternAttributes{ref: ref}
}

func (pp ConditionPathPatternAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pp.ref.InternalTokens()
}

func (pp ConditionPathPatternAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](pp.ref.Append("values"))
}

type ConditionQueryStringAttributes struct {
	ref terra.Reference
}

func (qs ConditionQueryStringAttributes) InternalRef() (terra.Reference, error) {
	return qs.ref, nil
}

func (qs ConditionQueryStringAttributes) InternalWithRef(ref terra.Reference) ConditionQueryStringAttributes {
	return ConditionQueryStringAttributes{ref: ref}
}

func (qs ConditionQueryStringAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return qs.ref.InternalTokens()
}

func (qs ConditionQueryStringAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(qs.ref.Append("key"))
}

func (qs ConditionQueryStringAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(qs.ref.Append("value"))
}

type ConditionSourceIpAttributes struct {
	ref terra.Reference
}

func (si ConditionSourceIpAttributes) InternalRef() (terra.Reference, error) {
	return si.ref, nil
}

func (si ConditionSourceIpAttributes) InternalWithRef(ref terra.Reference) ConditionSourceIpAttributes {
	return ConditionSourceIpAttributes{ref: ref}
}

func (si ConditionSourceIpAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return si.ref.InternalTokens()
}

func (si ConditionSourceIpAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](si.ref.Append("values"))
}

type ActionState struct {
	Order               float64                          `json:"order"`
	TargetGroupArn      string                           `json:"target_group_arn"`
	Type                string                           `json:"type"`
	AuthenticateCognito []ActionAuthenticateCognitoState `json:"authenticate_cognito"`
	AuthenticateOidc    []ActionAuthenticateOidcState    `json:"authenticate_oidc"`
	FixedResponse       []ActionFixedResponseState       `json:"fixed_response"`
	Forward             []ActionForwardState             `json:"forward"`
	Redirect            []ActionRedirectState            `json:"redirect"`
}

type ActionAuthenticateCognitoState struct {
	AuthenticationRequestExtraParams map[string]string `json:"authentication_request_extra_params"`
	OnUnauthenticatedRequest         string            `json:"on_unauthenticated_request"`
	Scope                            string            `json:"scope"`
	SessionCookieName                string            `json:"session_cookie_name"`
	SessionTimeout                   float64           `json:"session_timeout"`
	UserPoolArn                      string            `json:"user_pool_arn"`
	UserPoolClientId                 string            `json:"user_pool_client_id"`
	UserPoolDomain                   string            `json:"user_pool_domain"`
}

type ActionAuthenticateOidcState struct {
	AuthenticationRequestExtraParams map[string]string `json:"authentication_request_extra_params"`
	AuthorizationEndpoint            string            `json:"authorization_endpoint"`
	ClientId                         string            `json:"client_id"`
	ClientSecret                     string            `json:"client_secret"`
	Issuer                           string            `json:"issuer"`
	OnUnauthenticatedRequest         string            `json:"on_unauthenticated_request"`
	Scope                            string            `json:"scope"`
	SessionCookieName                string            `json:"session_cookie_name"`
	SessionTimeout                   float64           `json:"session_timeout"`
	TokenEndpoint                    string            `json:"token_endpoint"`
	UserInfoEndpoint                 string            `json:"user_info_endpoint"`
}

type ActionFixedResponseState struct {
	ContentType string `json:"content_type"`
	MessageBody string `json:"message_body"`
	StatusCode  string `json:"status_code"`
}

type ActionForwardState struct {
	Stickiness  []ActionForwardStickinessState  `json:"stickiness"`
	TargetGroup []ActionForwardTargetGroupState `json:"target_group"`
}

type ActionForwardStickinessState struct {
	Duration float64 `json:"duration"`
	Enabled  bool    `json:"enabled"`
}

type ActionForwardTargetGroupState struct {
	Arn    string  `json:"arn"`
	Weight float64 `json:"weight"`
}

type ActionRedirectState struct {
	Host       string `json:"host"`
	Path       string `json:"path"`
	Port       string `json:"port"`
	Protocol   string `json:"protocol"`
	Query      string `json:"query"`
	StatusCode string `json:"status_code"`
}

type ConditionState struct {
	HostHeader        []ConditionHostHeaderState        `json:"host_header"`
	HttpHeader        []ConditionHttpHeaderState        `json:"http_header"`
	HttpRequestMethod []ConditionHttpRequestMethodState `json:"http_request_method"`
	PathPattern       []ConditionPathPatternState       `json:"path_pattern"`
	QueryString       []ConditionQueryStringState       `json:"query_string"`
	SourceIp          []ConditionSourceIpState          `json:"source_ip"`
}

type ConditionHostHeaderState struct {
	Values []string `json:"values"`
}

type ConditionHttpHeaderState struct {
	HttpHeaderName string   `json:"http_header_name"`
	Values         []string `json:"values"`
}

type ConditionHttpRequestMethodState struct {
	Values []string `json:"values"`
}

type ConditionPathPatternState struct {
	Values []string `json:"values"`
}

type ConditionQueryStringState struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ConditionSourceIpState struct {
	Values []string `json:"values"`
}
