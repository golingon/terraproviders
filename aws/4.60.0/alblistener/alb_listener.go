// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package alblistener

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type DefaultAction struct {
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

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
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

func (da DefaultActionAttributes) InternalTokens() hclwrite.Tokens {
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

func (da DefaultActionAttributes) AuthenticateCognito() terra.ListValue[AuthenticateCognitoAttributes] {
	return terra.ReferenceAsList[AuthenticateCognitoAttributes](da.ref.Append("authenticate_cognito"))
}

func (da DefaultActionAttributes) AuthenticateOidc() terra.ListValue[AuthenticateOidcAttributes] {
	return terra.ReferenceAsList[AuthenticateOidcAttributes](da.ref.Append("authenticate_oidc"))
}

func (da DefaultActionAttributes) FixedResponse() terra.ListValue[FixedResponseAttributes] {
	return terra.ReferenceAsList[FixedResponseAttributes](da.ref.Append("fixed_response"))
}

func (da DefaultActionAttributes) Forward() terra.ListValue[ForwardAttributes] {
	return terra.ReferenceAsList[ForwardAttributes](da.ref.Append("forward"))
}

func (da DefaultActionAttributes) Redirect() terra.ListValue[RedirectAttributes] {
	return terra.ReferenceAsList[RedirectAttributes](da.ref.Append("redirect"))
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

func (ac AuthenticateCognitoAttributes) InternalTokens() hclwrite.Tokens {
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

func (ao AuthenticateOidcAttributes) InternalTokens() hclwrite.Tokens {
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

func (fr FixedResponseAttributes) InternalTokens() hclwrite.Tokens {
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

func (f ForwardAttributes) InternalTokens() hclwrite.Tokens {
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

func (s StickinessAttributes) InternalTokens() hclwrite.Tokens {
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

func (tg TargetGroupAttributes) InternalTokens() hclwrite.Tokens {
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

func (r RedirectAttributes) InternalTokens() hclwrite.Tokens {
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

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type DefaultActionState struct {
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

type TimeoutsState struct {
	Read string `json:"read"`
}
