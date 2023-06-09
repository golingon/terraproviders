// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package cloudschedulerjob

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AppEngineHttpTarget struct {
	// Body: string, optional
	Body terra.StringValue `hcl:"body,attr"`
	// Headers: map of string, optional
	Headers terra.MapValue[terra.StringValue] `hcl:"headers,attr"`
	// HttpMethod: string, optional
	HttpMethod terra.StringValue `hcl:"http_method,attr"`
	// RelativeUri: string, required
	RelativeUri terra.StringValue `hcl:"relative_uri,attr" validate:"required"`
	// AppEngineRouting: optional
	AppEngineRouting *AppEngineRouting `hcl:"app_engine_routing,block"`
}

type AppEngineRouting struct {
	// Instance: string, optional
	Instance terra.StringValue `hcl:"instance,attr"`
	// Service: string, optional
	Service terra.StringValue `hcl:"service,attr"`
	// Version: string, optional
	Version terra.StringValue `hcl:"version,attr"`
}

type HttpTarget struct {
	// Body: string, optional
	Body terra.StringValue `hcl:"body,attr"`
	// Headers: map of string, optional
	Headers terra.MapValue[terra.StringValue] `hcl:"headers,attr"`
	// HttpMethod: string, optional
	HttpMethod terra.StringValue `hcl:"http_method,attr"`
	// Uri: string, required
	Uri terra.StringValue `hcl:"uri,attr" validate:"required"`
	// OauthToken: optional
	OauthToken *OauthToken `hcl:"oauth_token,block"`
	// OidcToken: optional
	OidcToken *OidcToken `hcl:"oidc_token,block"`
}

type OauthToken struct {
	// Scope: string, optional
	Scope terra.StringValue `hcl:"scope,attr"`
	// ServiceAccountEmail: string, required
	ServiceAccountEmail terra.StringValue `hcl:"service_account_email,attr" validate:"required"`
}

type OidcToken struct {
	// Audience: string, optional
	Audience terra.StringValue `hcl:"audience,attr"`
	// ServiceAccountEmail: string, required
	ServiceAccountEmail terra.StringValue `hcl:"service_account_email,attr" validate:"required"`
}

type PubsubTarget struct {
	// Attributes: map of string, optional
	Attributes terra.MapValue[terra.StringValue] `hcl:"attributes,attr"`
	// Data: string, optional
	Data terra.StringValue `hcl:"data,attr"`
	// TopicName: string, required
	TopicName terra.StringValue `hcl:"topic_name,attr" validate:"required"`
}

type RetryConfig struct {
	// MaxBackoffDuration: string, optional
	MaxBackoffDuration terra.StringValue `hcl:"max_backoff_duration,attr"`
	// MaxDoublings: number, optional
	MaxDoublings terra.NumberValue `hcl:"max_doublings,attr"`
	// MaxRetryDuration: string, optional
	MaxRetryDuration terra.StringValue `hcl:"max_retry_duration,attr"`
	// MinBackoffDuration: string, optional
	MinBackoffDuration terra.StringValue `hcl:"min_backoff_duration,attr"`
	// RetryCount: number, optional
	RetryCount terra.NumberValue `hcl:"retry_count,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type AppEngineHttpTargetAttributes struct {
	ref terra.Reference
}

func (aeht AppEngineHttpTargetAttributes) InternalRef() (terra.Reference, error) {
	return aeht.ref, nil
}

func (aeht AppEngineHttpTargetAttributes) InternalWithRef(ref terra.Reference) AppEngineHttpTargetAttributes {
	return AppEngineHttpTargetAttributes{ref: ref}
}

func (aeht AppEngineHttpTargetAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return aeht.ref.InternalTokens()
}

func (aeht AppEngineHttpTargetAttributes) Body() terra.StringValue {
	return terra.ReferenceAsString(aeht.ref.Append("body"))
}

func (aeht AppEngineHttpTargetAttributes) Headers() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aeht.ref.Append("headers"))
}

func (aeht AppEngineHttpTargetAttributes) HttpMethod() terra.StringValue {
	return terra.ReferenceAsString(aeht.ref.Append("http_method"))
}

func (aeht AppEngineHttpTargetAttributes) RelativeUri() terra.StringValue {
	return terra.ReferenceAsString(aeht.ref.Append("relative_uri"))
}

func (aeht AppEngineHttpTargetAttributes) AppEngineRouting() terra.ListValue[AppEngineRoutingAttributes] {
	return terra.ReferenceAsList[AppEngineRoutingAttributes](aeht.ref.Append("app_engine_routing"))
}

type AppEngineRoutingAttributes struct {
	ref terra.Reference
}

func (aer AppEngineRoutingAttributes) InternalRef() (terra.Reference, error) {
	return aer.ref, nil
}

func (aer AppEngineRoutingAttributes) InternalWithRef(ref terra.Reference) AppEngineRoutingAttributes {
	return AppEngineRoutingAttributes{ref: ref}
}

func (aer AppEngineRoutingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return aer.ref.InternalTokens()
}

func (aer AppEngineRoutingAttributes) Instance() terra.StringValue {
	return terra.ReferenceAsString(aer.ref.Append("instance"))
}

func (aer AppEngineRoutingAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(aer.ref.Append("service"))
}

func (aer AppEngineRoutingAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(aer.ref.Append("version"))
}

type HttpTargetAttributes struct {
	ref terra.Reference
}

func (ht HttpTargetAttributes) InternalRef() (terra.Reference, error) {
	return ht.ref, nil
}

func (ht HttpTargetAttributes) InternalWithRef(ref terra.Reference) HttpTargetAttributes {
	return HttpTargetAttributes{ref: ref}
}

func (ht HttpTargetAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ht.ref.InternalTokens()
}

func (ht HttpTargetAttributes) Body() terra.StringValue {
	return terra.ReferenceAsString(ht.ref.Append("body"))
}

func (ht HttpTargetAttributes) Headers() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ht.ref.Append("headers"))
}

func (ht HttpTargetAttributes) HttpMethod() terra.StringValue {
	return terra.ReferenceAsString(ht.ref.Append("http_method"))
}

func (ht HttpTargetAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(ht.ref.Append("uri"))
}

func (ht HttpTargetAttributes) OauthToken() terra.ListValue[OauthTokenAttributes] {
	return terra.ReferenceAsList[OauthTokenAttributes](ht.ref.Append("oauth_token"))
}

func (ht HttpTargetAttributes) OidcToken() terra.ListValue[OidcTokenAttributes] {
	return terra.ReferenceAsList[OidcTokenAttributes](ht.ref.Append("oidc_token"))
}

type OauthTokenAttributes struct {
	ref terra.Reference
}

func (ot OauthTokenAttributes) InternalRef() (terra.Reference, error) {
	return ot.ref, nil
}

func (ot OauthTokenAttributes) InternalWithRef(ref terra.Reference) OauthTokenAttributes {
	return OauthTokenAttributes{ref: ref}
}

func (ot OauthTokenAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ot.ref.InternalTokens()
}

func (ot OauthTokenAttributes) Scope() terra.StringValue {
	return terra.ReferenceAsString(ot.ref.Append("scope"))
}

func (ot OauthTokenAttributes) ServiceAccountEmail() terra.StringValue {
	return terra.ReferenceAsString(ot.ref.Append("service_account_email"))
}

type OidcTokenAttributes struct {
	ref terra.Reference
}

func (ot OidcTokenAttributes) InternalRef() (terra.Reference, error) {
	return ot.ref, nil
}

func (ot OidcTokenAttributes) InternalWithRef(ref terra.Reference) OidcTokenAttributes {
	return OidcTokenAttributes{ref: ref}
}

func (ot OidcTokenAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ot.ref.InternalTokens()
}

func (ot OidcTokenAttributes) Audience() terra.StringValue {
	return terra.ReferenceAsString(ot.ref.Append("audience"))
}

func (ot OidcTokenAttributes) ServiceAccountEmail() terra.StringValue {
	return terra.ReferenceAsString(ot.ref.Append("service_account_email"))
}

type PubsubTargetAttributes struct {
	ref terra.Reference
}

func (pt PubsubTargetAttributes) InternalRef() (terra.Reference, error) {
	return pt.ref, nil
}

func (pt PubsubTargetAttributes) InternalWithRef(ref terra.Reference) PubsubTargetAttributes {
	return PubsubTargetAttributes{ref: ref}
}

func (pt PubsubTargetAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pt.ref.InternalTokens()
}

func (pt PubsubTargetAttributes) Attributes() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pt.ref.Append("attributes"))
}

func (pt PubsubTargetAttributes) Data() terra.StringValue {
	return terra.ReferenceAsString(pt.ref.Append("data"))
}

func (pt PubsubTargetAttributes) TopicName() terra.StringValue {
	return terra.ReferenceAsString(pt.ref.Append("topic_name"))
}

type RetryConfigAttributes struct {
	ref terra.Reference
}

func (rc RetryConfigAttributes) InternalRef() (terra.Reference, error) {
	return rc.ref, nil
}

func (rc RetryConfigAttributes) InternalWithRef(ref terra.Reference) RetryConfigAttributes {
	return RetryConfigAttributes{ref: ref}
}

func (rc RetryConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rc.ref.InternalTokens()
}

func (rc RetryConfigAttributes) MaxBackoffDuration() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("max_backoff_duration"))
}

func (rc RetryConfigAttributes) MaxDoublings() terra.NumberValue {
	return terra.ReferenceAsNumber(rc.ref.Append("max_doublings"))
}

func (rc RetryConfigAttributes) MaxRetryDuration() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("max_retry_duration"))
}

func (rc RetryConfigAttributes) MinBackoffDuration() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("min_backoff_duration"))
}

func (rc RetryConfigAttributes) RetryCount() terra.NumberValue {
	return terra.ReferenceAsNumber(rc.ref.Append("retry_count"))
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

type AppEngineHttpTargetState struct {
	Body             string                  `json:"body"`
	Headers          map[string]string       `json:"headers"`
	HttpMethod       string                  `json:"http_method"`
	RelativeUri      string                  `json:"relative_uri"`
	AppEngineRouting []AppEngineRoutingState `json:"app_engine_routing"`
}

type AppEngineRoutingState struct {
	Instance string `json:"instance"`
	Service  string `json:"service"`
	Version  string `json:"version"`
}

type HttpTargetState struct {
	Body       string            `json:"body"`
	Headers    map[string]string `json:"headers"`
	HttpMethod string            `json:"http_method"`
	Uri        string            `json:"uri"`
	OauthToken []OauthTokenState `json:"oauth_token"`
	OidcToken  []OidcTokenState  `json:"oidc_token"`
}

type OauthTokenState struct {
	Scope               string `json:"scope"`
	ServiceAccountEmail string `json:"service_account_email"`
}

type OidcTokenState struct {
	Audience            string `json:"audience"`
	ServiceAccountEmail string `json:"service_account_email"`
}

type PubsubTargetState struct {
	Attributes map[string]string `json:"attributes"`
	Data       string            `json:"data"`
	TopicName  string            `json:"topic_name"`
}

type RetryConfigState struct {
	MaxBackoffDuration string  `json:"max_backoff_duration"`
	MaxDoublings       float64 `json:"max_doublings"`
	MaxRetryDuration   string  `json:"max_retry_duration"`
	MinBackoffDuration string  `json:"min_backoff_duration"`
	RetryCount         float64 `json:"retry_count"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
