// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_lb_listener

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DefaultAction struct {
	// Order: number, optional
	Order terra.NumberValue `hcl:"order,attr"`
	// TargetGroupArn: string, optional
	TargetGroupArn terra.StringValue `hcl:"target_group_arn,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// DefaultActionAuthenticateCognito: optional
	AuthenticateCognito *DefaultActionAuthenticateCognito `hcl:"authenticate_cognito,block"`
	// DefaultActionAuthenticateOidc: optional
	AuthenticateOidc *DefaultActionAuthenticateOidc `hcl:"authenticate_oidc,block"`
	// DefaultActionFixedResponse: optional
	FixedResponse *DefaultActionFixedResponse `hcl:"fixed_response,block"`
	// DefaultActionForward: optional
	Forward *DefaultActionForward `hcl:"forward,block"`
	// DefaultActionRedirect: optional
	Redirect *DefaultActionRedirect `hcl:"redirect,block"`
}

type DefaultActionAuthenticateCognito struct {
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

type DefaultActionAuthenticateOidc struct {
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

type DefaultActionFixedResponse struct {
	// ContentType: string, required
	ContentType terra.StringValue `hcl:"content_type,attr" validate:"required"`
	// MessageBody: string, optional
	MessageBody terra.StringValue `hcl:"message_body,attr"`
	// StatusCode: string, optional
	StatusCode terra.StringValue `hcl:"status_code,attr"`
}

type DefaultActionForward struct {
	// DefaultActionForwardStickiness: optional
	Stickiness *DefaultActionForwardStickiness `hcl:"stickiness,block"`
	// DefaultActionForwardTargetGroup: min=1,max=5
	TargetGroup []DefaultActionForwardTargetGroup `hcl:"target_group,block" validate:"min=1,max=5"`
}

type DefaultActionForwardStickiness struct {
	// Duration: number, required
	Duration terra.NumberValue `hcl:"duration,attr" validate:"required"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
}

type DefaultActionForwardTargetGroup struct {
	// Arn: string, required
	Arn terra.StringValue `hcl:"arn,attr" validate:"required"`
	// Weight: number, optional
	Weight terra.NumberValue `hcl:"weight,attr"`
}

type DefaultActionRedirect struct {
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

type MutualAuthentication struct {
	// IgnoreClientCertificateExpiry: bool, optional
	IgnoreClientCertificateExpiry terra.BoolValue `hcl:"ignore_client_certificate_expiry,attr"`
	// Mode: string, required
	Mode terra.StringValue `hcl:"mode,attr" validate:"required"`
	// TrustStoreArn: string, optional
	TrustStoreArn terra.StringValue `hcl:"trust_store_arn,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type DefaultActionAttributes struct {
	ref terra.Reference
}

func (da DefaultActionAttributes) InternalRef() (terra.Reference, error) {
	return da.ref, nil
}

func (da DefaultActionAttributes) InternalWithRef(ref terra.Reference) DefaultActionAttributes {
	return DefaultActionAttributes{ref: ref}
}

func (da DefaultActionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return da.ref.InternalTokens()
}

func (da DefaultActionAttributes) Order() terra.NumberValue {
	return terra.ReferenceAsNumber(da.ref.Append("order"))
}

func (da DefaultActionAttributes) TargetGroupArn() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("target_group_arn"))
}

func (da DefaultActionAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("type"))
}

func (da DefaultActionAttributes) AuthenticateCognito() terra.ListValue[DefaultActionAuthenticateCognitoAttributes] {
	return terra.ReferenceAsList[DefaultActionAuthenticateCognitoAttributes](da.ref.Append("authenticate_cognito"))
}

func (da DefaultActionAttributes) AuthenticateOidc() terra.ListValue[DefaultActionAuthenticateOidcAttributes] {
	return terra.ReferenceAsList[DefaultActionAuthenticateOidcAttributes](da.ref.Append("authenticate_oidc"))
}

func (da DefaultActionAttributes) FixedResponse() terra.ListValue[DefaultActionFixedResponseAttributes] {
	return terra.ReferenceAsList[DefaultActionFixedResponseAttributes](da.ref.Append("fixed_response"))
}

func (da DefaultActionAttributes) Forward() terra.ListValue[DefaultActionForwardAttributes] {
	return terra.ReferenceAsList[DefaultActionForwardAttributes](da.ref.Append("forward"))
}

func (da DefaultActionAttributes) Redirect() terra.ListValue[DefaultActionRedirectAttributes] {
	return terra.ReferenceAsList[DefaultActionRedirectAttributes](da.ref.Append("redirect"))
}

type DefaultActionAuthenticateCognitoAttributes struct {
	ref terra.Reference
}

func (ac DefaultActionAuthenticateCognitoAttributes) InternalRef() (terra.Reference, error) {
	return ac.ref, nil
}

func (ac DefaultActionAuthenticateCognitoAttributes) InternalWithRef(ref terra.Reference) DefaultActionAuthenticateCognitoAttributes {
	return DefaultActionAuthenticateCognitoAttributes{ref: ref}
}

func (ac DefaultActionAuthenticateCognitoAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ac.ref.InternalTokens()
}

func (ac DefaultActionAuthenticateCognitoAttributes) AuthenticationRequestExtraParams() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ac.ref.Append("authentication_request_extra_params"))
}

func (ac DefaultActionAuthenticateCognitoAttributes) OnUnauthenticatedRequest() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("on_unauthenticated_request"))
}

func (ac DefaultActionAuthenticateCognitoAttributes) Scope() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("scope"))
}

func (ac DefaultActionAuthenticateCognitoAttributes) SessionCookieName() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("session_cookie_name"))
}

func (ac DefaultActionAuthenticateCognitoAttributes) SessionTimeout() terra.NumberValue {
	return terra.ReferenceAsNumber(ac.ref.Append("session_timeout"))
}

func (ac DefaultActionAuthenticateCognitoAttributes) UserPoolArn() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("user_pool_arn"))
}

func (ac DefaultActionAuthenticateCognitoAttributes) UserPoolClientId() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("user_pool_client_id"))
}

func (ac DefaultActionAuthenticateCognitoAttributes) UserPoolDomain() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("user_pool_domain"))
}

type DefaultActionAuthenticateOidcAttributes struct {
	ref terra.Reference
}

func (ao DefaultActionAuthenticateOidcAttributes) InternalRef() (terra.Reference, error) {
	return ao.ref, nil
}

func (ao DefaultActionAuthenticateOidcAttributes) InternalWithRef(ref terra.Reference) DefaultActionAuthenticateOidcAttributes {
	return DefaultActionAuthenticateOidcAttributes{ref: ref}
}

func (ao DefaultActionAuthenticateOidcAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ao.ref.InternalTokens()
}

func (ao DefaultActionAuthenticateOidcAttributes) AuthenticationRequestExtraParams() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ao.ref.Append("authentication_request_extra_params"))
}

func (ao DefaultActionAuthenticateOidcAttributes) AuthorizationEndpoint() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("authorization_endpoint"))
}

func (ao DefaultActionAuthenticateOidcAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("client_id"))
}

func (ao DefaultActionAuthenticateOidcAttributes) ClientSecret() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("client_secret"))
}

func (ao DefaultActionAuthenticateOidcAttributes) Issuer() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("issuer"))
}

func (ao DefaultActionAuthenticateOidcAttributes) OnUnauthenticatedRequest() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("on_unauthenticated_request"))
}

func (ao DefaultActionAuthenticateOidcAttributes) Scope() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("scope"))
}

func (ao DefaultActionAuthenticateOidcAttributes) SessionCookieName() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("session_cookie_name"))
}

func (ao DefaultActionAuthenticateOidcAttributes) SessionTimeout() terra.NumberValue {
	return terra.ReferenceAsNumber(ao.ref.Append("session_timeout"))
}

func (ao DefaultActionAuthenticateOidcAttributes) TokenEndpoint() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("token_endpoint"))
}

func (ao DefaultActionAuthenticateOidcAttributes) UserInfoEndpoint() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("user_info_endpoint"))
}

type DefaultActionFixedResponseAttributes struct {
	ref terra.Reference
}

func (fr DefaultActionFixedResponseAttributes) InternalRef() (terra.Reference, error) {
	return fr.ref, nil
}

func (fr DefaultActionFixedResponseAttributes) InternalWithRef(ref terra.Reference) DefaultActionFixedResponseAttributes {
	return DefaultActionFixedResponseAttributes{ref: ref}
}

func (fr DefaultActionFixedResponseAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fr.ref.InternalTokens()
}

func (fr DefaultActionFixedResponseAttributes) ContentType() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("content_type"))
}

func (fr DefaultActionFixedResponseAttributes) MessageBody() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("message_body"))
}

func (fr DefaultActionFixedResponseAttributes) StatusCode() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("status_code"))
}

type DefaultActionForwardAttributes struct {
	ref terra.Reference
}

func (f DefaultActionForwardAttributes) InternalRef() (terra.Reference, error) {
	return f.ref, nil
}

func (f DefaultActionForwardAttributes) InternalWithRef(ref terra.Reference) DefaultActionForwardAttributes {
	return DefaultActionForwardAttributes{ref: ref}
}

func (f DefaultActionForwardAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return f.ref.InternalTokens()
}

func (f DefaultActionForwardAttributes) Stickiness() terra.ListValue[DefaultActionForwardStickinessAttributes] {
	return terra.ReferenceAsList[DefaultActionForwardStickinessAttributes](f.ref.Append("stickiness"))
}

func (f DefaultActionForwardAttributes) TargetGroup() terra.SetValue[DefaultActionForwardTargetGroupAttributes] {
	return terra.ReferenceAsSet[DefaultActionForwardTargetGroupAttributes](f.ref.Append("target_group"))
}

type DefaultActionForwardStickinessAttributes struct {
	ref terra.Reference
}

func (s DefaultActionForwardStickinessAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s DefaultActionForwardStickinessAttributes) InternalWithRef(ref terra.Reference) DefaultActionForwardStickinessAttributes {
	return DefaultActionForwardStickinessAttributes{ref: ref}
}

func (s DefaultActionForwardStickinessAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s DefaultActionForwardStickinessAttributes) Duration() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("duration"))
}

func (s DefaultActionForwardStickinessAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("enabled"))
}

type DefaultActionForwardTargetGroupAttributes struct {
	ref terra.Reference
}

func (tg DefaultActionForwardTargetGroupAttributes) InternalRef() (terra.Reference, error) {
	return tg.ref, nil
}

func (tg DefaultActionForwardTargetGroupAttributes) InternalWithRef(ref terra.Reference) DefaultActionForwardTargetGroupAttributes {
	return DefaultActionForwardTargetGroupAttributes{ref: ref}
}

func (tg DefaultActionForwardTargetGroupAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tg.ref.InternalTokens()
}

func (tg DefaultActionForwardTargetGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(tg.ref.Append("arn"))
}

func (tg DefaultActionForwardTargetGroupAttributes) Weight() terra.NumberValue {
	return terra.ReferenceAsNumber(tg.ref.Append("weight"))
}

type DefaultActionRedirectAttributes struct {
	ref terra.Reference
}

func (r DefaultActionRedirectAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r DefaultActionRedirectAttributes) InternalWithRef(ref terra.Reference) DefaultActionRedirectAttributes {
	return DefaultActionRedirectAttributes{ref: ref}
}

func (r DefaultActionRedirectAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r DefaultActionRedirectAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("host"))
}

func (r DefaultActionRedirectAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("path"))
}

func (r DefaultActionRedirectAttributes) Port() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("port"))
}

func (r DefaultActionRedirectAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("protocol"))
}

func (r DefaultActionRedirectAttributes) Query() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("query"))
}

func (r DefaultActionRedirectAttributes) StatusCode() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("status_code"))
}

type MutualAuthenticationAttributes struct {
	ref terra.Reference
}

func (ma MutualAuthenticationAttributes) InternalRef() (terra.Reference, error) {
	return ma.ref, nil
}

func (ma MutualAuthenticationAttributes) InternalWithRef(ref terra.Reference) MutualAuthenticationAttributes {
	return MutualAuthenticationAttributes{ref: ref}
}

func (ma MutualAuthenticationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ma.ref.InternalTokens()
}

func (ma MutualAuthenticationAttributes) IgnoreClientCertificateExpiry() terra.BoolValue {
	return terra.ReferenceAsBool(ma.ref.Append("ignore_client_certificate_expiry"))
}

func (ma MutualAuthenticationAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("mode"))
}

func (ma MutualAuthenticationAttributes) TrustStoreArn() terra.StringValue {
	return terra.ReferenceAsString(ma.ref.Append("trust_store_arn"))
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

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type DefaultActionState struct {
	Order               float64                                 `json:"order"`
	TargetGroupArn      string                                  `json:"target_group_arn"`
	Type                string                                  `json:"type"`
	AuthenticateCognito []DefaultActionAuthenticateCognitoState `json:"authenticate_cognito"`
	AuthenticateOidc    []DefaultActionAuthenticateOidcState    `json:"authenticate_oidc"`
	FixedResponse       []DefaultActionFixedResponseState       `json:"fixed_response"`
	Forward             []DefaultActionForwardState             `json:"forward"`
	Redirect            []DefaultActionRedirectState            `json:"redirect"`
}

type DefaultActionAuthenticateCognitoState struct {
	AuthenticationRequestExtraParams map[string]string `json:"authentication_request_extra_params"`
	OnUnauthenticatedRequest         string            `json:"on_unauthenticated_request"`
	Scope                            string            `json:"scope"`
	SessionCookieName                string            `json:"session_cookie_name"`
	SessionTimeout                   float64           `json:"session_timeout"`
	UserPoolArn                      string            `json:"user_pool_arn"`
	UserPoolClientId                 string            `json:"user_pool_client_id"`
	UserPoolDomain                   string            `json:"user_pool_domain"`
}

type DefaultActionAuthenticateOidcState struct {
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

type DefaultActionFixedResponseState struct {
	ContentType string `json:"content_type"`
	MessageBody string `json:"message_body"`
	StatusCode  string `json:"status_code"`
}

type DefaultActionForwardState struct {
	Stickiness  []DefaultActionForwardStickinessState  `json:"stickiness"`
	TargetGroup []DefaultActionForwardTargetGroupState `json:"target_group"`
}

type DefaultActionForwardStickinessState struct {
	Duration float64 `json:"duration"`
	Enabled  bool    `json:"enabled"`
}

type DefaultActionForwardTargetGroupState struct {
	Arn    string  `json:"arn"`
	Weight float64 `json:"weight"`
}

type DefaultActionRedirectState struct {
	Host       string `json:"host"`
	Path       string `json:"path"`
	Port       string `json:"port"`
	Protocol   string `json:"protocol"`
	Query      string `json:"query"`
	StatusCode string `json:"status_code"`
}

type MutualAuthenticationState struct {
	IgnoreClientCertificateExpiry bool   `json:"ignore_client_certificate_expiry"`
	Mode                          string `json:"mode"`
	TrustStoreArn                 string `json:"trust_store_arn"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Update string `json:"update"`
}
