// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package lblistenerrule

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
	// AuthenticateCognito: optional
	AuthenticateCognito *AuthenticateCognito `hcl:"authenticate_cognito,block"`
	// AuthenticateOidc: optional
	AuthenticateOidc *AuthenticateOidc `hcl:"authenticate_oidc,block"`
	// FixedResponse: optional
	FixedResponse *FixedResponse `hcl:"fixed_response,block"`
	// Forward: optional
	Forward *Forward `hcl:"forward,block"`
	// Redirect: optional
	Redirect *Redirect `hcl:"redirect,block"`
}

type AuthenticateCognito struct {
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

type AuthenticateOidc struct {
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

type FixedResponse struct {
	// ContentType: string, required
	ContentType terra.StringValue `hcl:"content_type,attr" validate:"required"`
	// MessageBody: string, optional
	MessageBody terra.StringValue `hcl:"message_body,attr"`
	// StatusCode: string, optional
	StatusCode terra.StringValue `hcl:"status_code,attr"`
}

type Forward struct {
	// Stickiness: optional
	Stickiness *Stickiness `hcl:"stickiness,block"`
	// TargetGroup: min=1,max=5
	TargetGroup []TargetGroup `hcl:"target_group,block" validate:"min=1,max=5"`
}

type Stickiness struct {
	// Duration: number, required
	Duration terra.NumberValue `hcl:"duration,attr" validate:"required"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
}

type TargetGroup struct {
	// Arn: string, required
	Arn terra.StringValue `hcl:"arn,attr" validate:"required"`
	// Weight: number, optional
	Weight terra.NumberValue `hcl:"weight,attr"`
}

type Redirect struct {
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
	// HostHeader: optional
	HostHeader *HostHeader `hcl:"host_header,block"`
	// HttpHeader: optional
	HttpHeader *HttpHeader `hcl:"http_header,block"`
	// HttpRequestMethod: optional
	HttpRequestMethod *HttpRequestMethod `hcl:"http_request_method,block"`
	// PathPattern: optional
	PathPattern *PathPattern `hcl:"path_pattern,block"`
	// QueryString: min=0
	QueryString []QueryString `hcl:"query_string,block" validate:"min=0"`
	// SourceIp: optional
	SourceIp *SourceIp `hcl:"source_ip,block"`
}

type HostHeader struct {
	// Values: set of string, required
	Values terra.SetValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type HttpHeader struct {
	// HttpHeaderName: string, required
	HttpHeaderName terra.StringValue `hcl:"http_header_name,attr" validate:"required"`
	// Values: set of string, required
	Values terra.SetValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type HttpRequestMethod struct {
	// Values: set of string, required
	Values terra.SetValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type PathPattern struct {
	// Values: set of string, required
	Values terra.SetValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type QueryString struct {
	// Key: string, optional
	Key terra.StringValue `hcl:"key,attr"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type SourceIp struct {
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

func (a ActionAttributes) AuthenticateCognito() terra.ListValue[AuthenticateCognitoAttributes] {
	return terra.ReferenceAsList[AuthenticateCognitoAttributes](a.ref.Append("authenticate_cognito"))
}

func (a ActionAttributes) AuthenticateOidc() terra.ListValue[AuthenticateOidcAttributes] {
	return terra.ReferenceAsList[AuthenticateOidcAttributes](a.ref.Append("authenticate_oidc"))
}

func (a ActionAttributes) FixedResponse() terra.ListValue[FixedResponseAttributes] {
	return terra.ReferenceAsList[FixedResponseAttributes](a.ref.Append("fixed_response"))
}

func (a ActionAttributes) Forward() terra.ListValue[ForwardAttributes] {
	return terra.ReferenceAsList[ForwardAttributes](a.ref.Append("forward"))
}

func (a ActionAttributes) Redirect() terra.ListValue[RedirectAttributes] {
	return terra.ReferenceAsList[RedirectAttributes](a.ref.Append("redirect"))
}

type AuthenticateCognitoAttributes struct {
	ref terra.Reference
}

func (ac AuthenticateCognitoAttributes) InternalRef() (terra.Reference, error) {
	return ac.ref, nil
}

func (ac AuthenticateCognitoAttributes) InternalWithRef(ref terra.Reference) AuthenticateCognitoAttributes {
	return AuthenticateCognitoAttributes{ref: ref}
}

func (ac AuthenticateCognitoAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ac.ref.InternalTokens()
}

func (ac AuthenticateCognitoAttributes) AuthenticationRequestExtraParams() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ac.ref.Append("authentication_request_extra_params"))
}

func (ac AuthenticateCognitoAttributes) OnUnauthenticatedRequest() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("on_unauthenticated_request"))
}

func (ac AuthenticateCognitoAttributes) Scope() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("scope"))
}

func (ac AuthenticateCognitoAttributes) SessionCookieName() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("session_cookie_name"))
}

func (ac AuthenticateCognitoAttributes) SessionTimeout() terra.NumberValue {
	return terra.ReferenceAsNumber(ac.ref.Append("session_timeout"))
}

func (ac AuthenticateCognitoAttributes) UserPoolArn() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("user_pool_arn"))
}

func (ac AuthenticateCognitoAttributes) UserPoolClientId() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("user_pool_client_id"))
}

func (ac AuthenticateCognitoAttributes) UserPoolDomain() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("user_pool_domain"))
}

type AuthenticateOidcAttributes struct {
	ref terra.Reference
}

func (ao AuthenticateOidcAttributes) InternalRef() (terra.Reference, error) {
	return ao.ref, nil
}

func (ao AuthenticateOidcAttributes) InternalWithRef(ref terra.Reference) AuthenticateOidcAttributes {
	return AuthenticateOidcAttributes{ref: ref}
}

func (ao AuthenticateOidcAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ao.ref.InternalTokens()
}

func (ao AuthenticateOidcAttributes) AuthenticationRequestExtraParams() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ao.ref.Append("authentication_request_extra_params"))
}

func (ao AuthenticateOidcAttributes) AuthorizationEndpoint() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("authorization_endpoint"))
}

func (ao AuthenticateOidcAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("client_id"))
}

func (ao AuthenticateOidcAttributes) ClientSecret() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("client_secret"))
}

func (ao AuthenticateOidcAttributes) Issuer() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("issuer"))
}

func (ao AuthenticateOidcAttributes) OnUnauthenticatedRequest() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("on_unauthenticated_request"))
}

func (ao AuthenticateOidcAttributes) Scope() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("scope"))
}

func (ao AuthenticateOidcAttributes) SessionCookieName() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("session_cookie_name"))
}

func (ao AuthenticateOidcAttributes) SessionTimeout() terra.NumberValue {
	return terra.ReferenceAsNumber(ao.ref.Append("session_timeout"))
}

func (ao AuthenticateOidcAttributes) TokenEndpoint() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("token_endpoint"))
}

func (ao AuthenticateOidcAttributes) UserInfoEndpoint() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("user_info_endpoint"))
}

type FixedResponseAttributes struct {
	ref terra.Reference
}

func (fr FixedResponseAttributes) InternalRef() (terra.Reference, error) {
	return fr.ref, nil
}

func (fr FixedResponseAttributes) InternalWithRef(ref terra.Reference) FixedResponseAttributes {
	return FixedResponseAttributes{ref: ref}
}

func (fr FixedResponseAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fr.ref.InternalTokens()
}

func (fr FixedResponseAttributes) ContentType() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("content_type"))
}

func (fr FixedResponseAttributes) MessageBody() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("message_body"))
}

func (fr FixedResponseAttributes) StatusCode() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("status_code"))
}

type ForwardAttributes struct {
	ref terra.Reference
}

func (f ForwardAttributes) InternalRef() (terra.Reference, error) {
	return f.ref, nil
}

func (f ForwardAttributes) InternalWithRef(ref terra.Reference) ForwardAttributes {
	return ForwardAttributes{ref: ref}
}

func (f ForwardAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return f.ref.InternalTokens()
}

func (f ForwardAttributes) Stickiness() terra.ListValue[StickinessAttributes] {
	return terra.ReferenceAsList[StickinessAttributes](f.ref.Append("stickiness"))
}

func (f ForwardAttributes) TargetGroup() terra.SetValue[TargetGroupAttributes] {
	return terra.ReferenceAsSet[TargetGroupAttributes](f.ref.Append("target_group"))
}

type StickinessAttributes struct {
	ref terra.Reference
}

func (s StickinessAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s StickinessAttributes) InternalWithRef(ref terra.Reference) StickinessAttributes {
	return StickinessAttributes{ref: ref}
}

func (s StickinessAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s StickinessAttributes) Duration() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("duration"))
}

func (s StickinessAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("enabled"))
}

type TargetGroupAttributes struct {
	ref terra.Reference
}

func (tg TargetGroupAttributes) InternalRef() (terra.Reference, error) {
	return tg.ref, nil
}

func (tg TargetGroupAttributes) InternalWithRef(ref terra.Reference) TargetGroupAttributes {
	return TargetGroupAttributes{ref: ref}
}

func (tg TargetGroupAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tg.ref.InternalTokens()
}

func (tg TargetGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(tg.ref.Append("arn"))
}

func (tg TargetGroupAttributes) Weight() terra.NumberValue {
	return terra.ReferenceAsNumber(tg.ref.Append("weight"))
}

type RedirectAttributes struct {
	ref terra.Reference
}

func (r RedirectAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RedirectAttributes) InternalWithRef(ref terra.Reference) RedirectAttributes {
	return RedirectAttributes{ref: ref}
}

func (r RedirectAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RedirectAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("host"))
}

func (r RedirectAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("path"))
}

func (r RedirectAttributes) Port() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("port"))
}

func (r RedirectAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("protocol"))
}

func (r RedirectAttributes) Query() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("query"))
}

func (r RedirectAttributes) StatusCode() terra.StringValue {
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

func (c ConditionAttributes) HostHeader() terra.ListValue[HostHeaderAttributes] {
	return terra.ReferenceAsList[HostHeaderAttributes](c.ref.Append("host_header"))
}

func (c ConditionAttributes) HttpHeader() terra.ListValue[HttpHeaderAttributes] {
	return terra.ReferenceAsList[HttpHeaderAttributes](c.ref.Append("http_header"))
}

func (c ConditionAttributes) HttpRequestMethod() terra.ListValue[HttpRequestMethodAttributes] {
	return terra.ReferenceAsList[HttpRequestMethodAttributes](c.ref.Append("http_request_method"))
}

func (c ConditionAttributes) PathPattern() terra.ListValue[PathPatternAttributes] {
	return terra.ReferenceAsList[PathPatternAttributes](c.ref.Append("path_pattern"))
}

func (c ConditionAttributes) QueryString() terra.SetValue[QueryStringAttributes] {
	return terra.ReferenceAsSet[QueryStringAttributes](c.ref.Append("query_string"))
}

func (c ConditionAttributes) SourceIp() terra.ListValue[SourceIpAttributes] {
	return terra.ReferenceAsList[SourceIpAttributes](c.ref.Append("source_ip"))
}

type HostHeaderAttributes struct {
	ref terra.Reference
}

func (hh HostHeaderAttributes) InternalRef() (terra.Reference, error) {
	return hh.ref, nil
}

func (hh HostHeaderAttributes) InternalWithRef(ref terra.Reference) HostHeaderAttributes {
	return HostHeaderAttributes{ref: ref}
}

func (hh HostHeaderAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hh.ref.InternalTokens()
}

func (hh HostHeaderAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](hh.ref.Append("values"))
}

type HttpHeaderAttributes struct {
	ref terra.Reference
}

func (hh HttpHeaderAttributes) InternalRef() (terra.Reference, error) {
	return hh.ref, nil
}

func (hh HttpHeaderAttributes) InternalWithRef(ref terra.Reference) HttpHeaderAttributes {
	return HttpHeaderAttributes{ref: ref}
}

func (hh HttpHeaderAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hh.ref.InternalTokens()
}

func (hh HttpHeaderAttributes) HttpHeaderName() terra.StringValue {
	return terra.ReferenceAsString(hh.ref.Append("http_header_name"))
}

func (hh HttpHeaderAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](hh.ref.Append("values"))
}

type HttpRequestMethodAttributes struct {
	ref terra.Reference
}

func (hrm HttpRequestMethodAttributes) InternalRef() (terra.Reference, error) {
	return hrm.ref, nil
}

func (hrm HttpRequestMethodAttributes) InternalWithRef(ref terra.Reference) HttpRequestMethodAttributes {
	return HttpRequestMethodAttributes{ref: ref}
}

func (hrm HttpRequestMethodAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hrm.ref.InternalTokens()
}

func (hrm HttpRequestMethodAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](hrm.ref.Append("values"))
}

type PathPatternAttributes struct {
	ref terra.Reference
}

func (pp PathPatternAttributes) InternalRef() (terra.Reference, error) {
	return pp.ref, nil
}

func (pp PathPatternAttributes) InternalWithRef(ref terra.Reference) PathPatternAttributes {
	return PathPatternAttributes{ref: ref}
}

func (pp PathPatternAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pp.ref.InternalTokens()
}

func (pp PathPatternAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](pp.ref.Append("values"))
}

type QueryStringAttributes struct {
	ref terra.Reference
}

func (qs QueryStringAttributes) InternalRef() (terra.Reference, error) {
	return qs.ref, nil
}

func (qs QueryStringAttributes) InternalWithRef(ref terra.Reference) QueryStringAttributes {
	return QueryStringAttributes{ref: ref}
}

func (qs QueryStringAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return qs.ref.InternalTokens()
}

func (qs QueryStringAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(qs.ref.Append("key"))
}

func (qs QueryStringAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(qs.ref.Append("value"))
}

type SourceIpAttributes struct {
	ref terra.Reference
}

func (si SourceIpAttributes) InternalRef() (terra.Reference, error) {
	return si.ref, nil
}

func (si SourceIpAttributes) InternalWithRef(ref terra.Reference) SourceIpAttributes {
	return SourceIpAttributes{ref: ref}
}

func (si SourceIpAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return si.ref.InternalTokens()
}

func (si SourceIpAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](si.ref.Append("values"))
}

type ActionState struct {
	Order               float64                    `json:"order"`
	TargetGroupArn      string                     `json:"target_group_arn"`
	Type                string                     `json:"type"`
	AuthenticateCognito []AuthenticateCognitoState `json:"authenticate_cognito"`
	AuthenticateOidc    []AuthenticateOidcState    `json:"authenticate_oidc"`
	FixedResponse       []FixedResponseState       `json:"fixed_response"`
	Forward             []ForwardState             `json:"forward"`
	Redirect            []RedirectState            `json:"redirect"`
}

type AuthenticateCognitoState struct {
	AuthenticationRequestExtraParams map[string]string `json:"authentication_request_extra_params"`
	OnUnauthenticatedRequest         string            `json:"on_unauthenticated_request"`
	Scope                            string            `json:"scope"`
	SessionCookieName                string            `json:"session_cookie_name"`
	SessionTimeout                   float64           `json:"session_timeout"`
	UserPoolArn                      string            `json:"user_pool_arn"`
	UserPoolClientId                 string            `json:"user_pool_client_id"`
	UserPoolDomain                   string            `json:"user_pool_domain"`
}

type AuthenticateOidcState struct {
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

type FixedResponseState struct {
	ContentType string `json:"content_type"`
	MessageBody string `json:"message_body"`
	StatusCode  string `json:"status_code"`
}

type ForwardState struct {
	Stickiness  []StickinessState  `json:"stickiness"`
	TargetGroup []TargetGroupState `json:"target_group"`
}

type StickinessState struct {
	Duration float64 `json:"duration"`
	Enabled  bool    `json:"enabled"`
}

type TargetGroupState struct {
	Arn    string  `json:"arn"`
	Weight float64 `json:"weight"`
}

type RedirectState struct {
	Host       string `json:"host"`
	Path       string `json:"path"`
	Port       string `json:"port"`
	Protocol   string `json:"protocol"`
	Query      string `json:"query"`
	StatusCode string `json:"status_code"`
}

type ConditionState struct {
	HostHeader        []HostHeaderState        `json:"host_header"`
	HttpHeader        []HttpHeaderState        `json:"http_header"`
	HttpRequestMethod []HttpRequestMethodState `json:"http_request_method"`
	PathPattern       []PathPatternState       `json:"path_pattern"`
	QueryString       []QueryStringState       `json:"query_string"`
	SourceIp          []SourceIpState          `json:"source_ip"`
}

type HostHeaderState struct {
	Values []string `json:"values"`
}

type HttpHeaderState struct {
	HttpHeaderName string   `json:"http_header_name"`
	Values         []string `json:"values"`
}

type HttpRequestMethodState struct {
	Values []string `json:"values"`
}

type PathPatternState struct {
	Values []string `json:"values"`
}

type QueryStringState struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type SourceIpState struct {
	Values []string `json:"values"`
}
