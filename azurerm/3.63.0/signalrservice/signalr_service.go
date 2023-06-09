// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package signalrservice

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Cors struct {
	// AllowedOrigins: set of string, required
	AllowedOrigins terra.SetValue[terra.StringValue] `hcl:"allowed_origins,attr" validate:"required"`
}

type Identity struct {
	// IdentityIds: set of string, optional
	IdentityIds terra.SetValue[terra.StringValue] `hcl:"identity_ids,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type LiveTrace struct {
	// ConnectivityLogsEnabled: bool, optional
	ConnectivityLogsEnabled terra.BoolValue `hcl:"connectivity_logs_enabled,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// HttpRequestLogsEnabled: bool, optional
	HttpRequestLogsEnabled terra.BoolValue `hcl:"http_request_logs_enabled,attr"`
	// MessagingLogsEnabled: bool, optional
	MessagingLogsEnabled terra.BoolValue `hcl:"messaging_logs_enabled,attr"`
}

type Sku struct {
	// Capacity: number, required
	Capacity terra.NumberValue `hcl:"capacity,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type UpstreamEndpoint struct {
	// CategoryPattern: list of string, required
	CategoryPattern terra.ListValue[terra.StringValue] `hcl:"category_pattern,attr" validate:"required"`
	// EventPattern: list of string, required
	EventPattern terra.ListValue[terra.StringValue] `hcl:"event_pattern,attr" validate:"required"`
	// HubPattern: list of string, required
	HubPattern terra.ListValue[terra.StringValue] `hcl:"hub_pattern,attr" validate:"required"`
	// UrlTemplate: string, required
	UrlTemplate terra.StringValue `hcl:"url_template,attr" validate:"required"`
	// UserAssignedIdentityId: string, optional
	UserAssignedIdentityId terra.StringValue `hcl:"user_assigned_identity_id,attr"`
}

type CorsAttributes struct {
	ref terra.Reference
}

func (c CorsAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c CorsAttributes) InternalWithRef(ref terra.Reference) CorsAttributes {
	return CorsAttributes{ref: ref}
}

func (c CorsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c CorsAttributes) AllowedOrigins() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](c.ref.Append("allowed_origins"))
}

type IdentityAttributes struct {
	ref terra.Reference
}

func (i IdentityAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i IdentityAttributes) InternalWithRef(ref terra.Reference) IdentityAttributes {
	return IdentityAttributes{ref: ref}
}

func (i IdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i IdentityAttributes) IdentityIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](i.ref.Append("identity_ids"))
}

func (i IdentityAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("principal_id"))
}

func (i IdentityAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("tenant_id"))
}

func (i IdentityAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("type"))
}

type LiveTraceAttributes struct {
	ref terra.Reference
}

func (lt LiveTraceAttributes) InternalRef() (terra.Reference, error) {
	return lt.ref, nil
}

func (lt LiveTraceAttributes) InternalWithRef(ref terra.Reference) LiveTraceAttributes {
	return LiveTraceAttributes{ref: ref}
}

func (lt LiveTraceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lt.ref.InternalTokens()
}

func (lt LiveTraceAttributes) ConnectivityLogsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(lt.ref.Append("connectivity_logs_enabled"))
}

func (lt LiveTraceAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(lt.ref.Append("enabled"))
}

func (lt LiveTraceAttributes) HttpRequestLogsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(lt.ref.Append("http_request_logs_enabled"))
}

func (lt LiveTraceAttributes) MessagingLogsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(lt.ref.Append("messaging_logs_enabled"))
}

type SkuAttributes struct {
	ref terra.Reference
}

func (s SkuAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SkuAttributes) InternalWithRef(ref terra.Reference) SkuAttributes {
	return SkuAttributes{ref: ref}
}

func (s SkuAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SkuAttributes) Capacity() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("capacity"))
}

func (s SkuAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("name"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type UpstreamEndpointAttributes struct {
	ref terra.Reference
}

func (ue UpstreamEndpointAttributes) InternalRef() (terra.Reference, error) {
	return ue.ref, nil
}

func (ue UpstreamEndpointAttributes) InternalWithRef(ref terra.Reference) UpstreamEndpointAttributes {
	return UpstreamEndpointAttributes{ref: ref}
}

func (ue UpstreamEndpointAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ue.ref.InternalTokens()
}

func (ue UpstreamEndpointAttributes) CategoryPattern() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ue.ref.Append("category_pattern"))
}

func (ue UpstreamEndpointAttributes) EventPattern() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ue.ref.Append("event_pattern"))
}

func (ue UpstreamEndpointAttributes) HubPattern() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ue.ref.Append("hub_pattern"))
}

func (ue UpstreamEndpointAttributes) UrlTemplate() terra.StringValue {
	return terra.ReferenceAsString(ue.ref.Append("url_template"))
}

func (ue UpstreamEndpointAttributes) UserAssignedIdentityId() terra.StringValue {
	return terra.ReferenceAsString(ue.ref.Append("user_assigned_identity_id"))
}

type CorsState struct {
	AllowedOrigins []string `json:"allowed_origins"`
}

type IdentityState struct {
	IdentityIds []string `json:"identity_ids"`
	PrincipalId string   `json:"principal_id"`
	TenantId    string   `json:"tenant_id"`
	Type        string   `json:"type"`
}

type LiveTraceState struct {
	ConnectivityLogsEnabled bool `json:"connectivity_logs_enabled"`
	Enabled                 bool `json:"enabled"`
	HttpRequestLogsEnabled  bool `json:"http_request_logs_enabled"`
	MessagingLogsEnabled    bool `json:"messaging_logs_enabled"`
}

type SkuState struct {
	Capacity float64 `json:"capacity"`
	Name     string  `json:"name"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}

type UpstreamEndpointState struct {
	CategoryPattern        []string `json:"category_pattern"`
	EventPattern           []string `json:"event_pattern"`
	HubPattern             []string `json:"hub_pattern"`
	UrlTemplate            string   `json:"url_template"`
	UserAssignedIdentityId string   `json:"user_assigned_identity_id"`
}
