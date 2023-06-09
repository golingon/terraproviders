// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package functionapp

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type SiteCredential struct{}

type AuthSettings struct {
	// AdditionalLoginParams: map of string, optional
	AdditionalLoginParams terra.MapValue[terra.StringValue] `hcl:"additional_login_params,attr"`
	// AllowedExternalRedirectUrls: list of string, optional
	AllowedExternalRedirectUrls terra.ListValue[terra.StringValue] `hcl:"allowed_external_redirect_urls,attr"`
	// DefaultProvider: string, optional
	DefaultProvider terra.StringValue `hcl:"default_provider,attr"`
	// Enabled: bool, required
	Enabled terra.BoolValue `hcl:"enabled,attr" validate:"required"`
	// Issuer: string, optional
	Issuer terra.StringValue `hcl:"issuer,attr"`
	// RuntimeVersion: string, optional
	RuntimeVersion terra.StringValue `hcl:"runtime_version,attr"`
	// TokenRefreshExtensionHours: number, optional
	TokenRefreshExtensionHours terra.NumberValue `hcl:"token_refresh_extension_hours,attr"`
	// TokenStoreEnabled: bool, optional
	TokenStoreEnabled terra.BoolValue `hcl:"token_store_enabled,attr"`
	// UnauthenticatedClientAction: string, optional
	UnauthenticatedClientAction terra.StringValue `hcl:"unauthenticated_client_action,attr"`
	// ActiveDirectory: optional
	ActiveDirectory *ActiveDirectory `hcl:"active_directory,block"`
	// Facebook: optional
	Facebook *Facebook `hcl:"facebook,block"`
	// Google: optional
	Google *Google `hcl:"google,block"`
	// Microsoft: optional
	Microsoft *Microsoft `hcl:"microsoft,block"`
	// Twitter: optional
	Twitter *Twitter `hcl:"twitter,block"`
}

type ActiveDirectory struct {
	// AllowedAudiences: list of string, optional
	AllowedAudiences terra.ListValue[terra.StringValue] `hcl:"allowed_audiences,attr"`
	// ClientId: string, required
	ClientId terra.StringValue `hcl:"client_id,attr" validate:"required"`
	// ClientSecret: string, optional
	ClientSecret terra.StringValue `hcl:"client_secret,attr"`
}

type Facebook struct {
	// AppId: string, required
	AppId terra.StringValue `hcl:"app_id,attr" validate:"required"`
	// AppSecret: string, required
	AppSecret terra.StringValue `hcl:"app_secret,attr" validate:"required"`
	// OauthScopes: list of string, optional
	OauthScopes terra.ListValue[terra.StringValue] `hcl:"oauth_scopes,attr"`
}

type Google struct {
	// ClientId: string, required
	ClientId terra.StringValue `hcl:"client_id,attr" validate:"required"`
	// ClientSecret: string, required
	ClientSecret terra.StringValue `hcl:"client_secret,attr" validate:"required"`
	// OauthScopes: list of string, optional
	OauthScopes terra.ListValue[terra.StringValue] `hcl:"oauth_scopes,attr"`
}

type Microsoft struct {
	// ClientId: string, required
	ClientId terra.StringValue `hcl:"client_id,attr" validate:"required"`
	// ClientSecret: string, required
	ClientSecret terra.StringValue `hcl:"client_secret,attr" validate:"required"`
	// OauthScopes: list of string, optional
	OauthScopes terra.ListValue[terra.StringValue] `hcl:"oauth_scopes,attr"`
}

type Twitter struct {
	// ConsumerKey: string, required
	ConsumerKey terra.StringValue `hcl:"consumer_key,attr" validate:"required"`
	// ConsumerSecret: string, required
	ConsumerSecret terra.StringValue `hcl:"consumer_secret,attr" validate:"required"`
}

type ConnectionString struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type Identity struct {
	// IdentityIds: set of string, optional
	IdentityIds terra.SetValue[terra.StringValue] `hcl:"identity_ids,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type SiteConfig struct {
	// AlwaysOn: bool, optional
	AlwaysOn terra.BoolValue `hcl:"always_on,attr"`
	// AppScaleLimit: number, optional
	AppScaleLimit terra.NumberValue `hcl:"app_scale_limit,attr"`
	// AutoSwapSlotName: string, optional
	AutoSwapSlotName terra.StringValue `hcl:"auto_swap_slot_name,attr"`
	// DotnetFrameworkVersion: string, optional
	DotnetFrameworkVersion terra.StringValue `hcl:"dotnet_framework_version,attr"`
	// ElasticInstanceMinimum: number, optional
	ElasticInstanceMinimum terra.NumberValue `hcl:"elastic_instance_minimum,attr"`
	// FtpsState: string, optional
	FtpsState terra.StringValue `hcl:"ftps_state,attr"`
	// HealthCheckPath: string, optional
	HealthCheckPath terra.StringValue `hcl:"health_check_path,attr"`
	// Http2Enabled: bool, optional
	Http2Enabled terra.BoolValue `hcl:"http2_enabled,attr"`
	// JavaVersion: string, optional
	JavaVersion terra.StringValue `hcl:"java_version,attr"`
	// LinuxFxVersion: string, optional
	LinuxFxVersion terra.StringValue `hcl:"linux_fx_version,attr"`
	// MinTlsVersion: string, optional
	MinTlsVersion terra.StringValue `hcl:"min_tls_version,attr"`
	// PreWarmedInstanceCount: number, optional
	PreWarmedInstanceCount terra.NumberValue `hcl:"pre_warmed_instance_count,attr"`
	// RuntimeScaleMonitoringEnabled: bool, optional
	RuntimeScaleMonitoringEnabled terra.BoolValue `hcl:"runtime_scale_monitoring_enabled,attr"`
	// ScmType: string, optional
	ScmType terra.StringValue `hcl:"scm_type,attr"`
	// ScmUseMainIpRestriction: bool, optional
	ScmUseMainIpRestriction terra.BoolValue `hcl:"scm_use_main_ip_restriction,attr"`
	// Use32BitWorkerProcess: bool, optional
	Use32BitWorkerProcess terra.BoolValue `hcl:"use_32_bit_worker_process,attr"`
	// VnetRouteAllEnabled: bool, optional
	VnetRouteAllEnabled terra.BoolValue `hcl:"vnet_route_all_enabled,attr"`
	// WebsocketsEnabled: bool, optional
	WebsocketsEnabled terra.BoolValue `hcl:"websockets_enabled,attr"`
	// IpRestriction: min=0
	IpRestriction []IpRestriction `hcl:"ip_restriction,block" validate:"min=0"`
	// ScmIpRestriction: min=0
	ScmIpRestriction []ScmIpRestriction `hcl:"scm_ip_restriction,block" validate:"min=0"`
	// Cors: optional
	Cors *Cors `hcl:"cors,block"`
}

type IpRestriction struct {
	// Action: string, optional
	Action terra.StringValue `hcl:"action,attr"`
	// IpAddress: string, optional
	IpAddress terra.StringValue `hcl:"ip_address,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Priority: number, optional
	Priority terra.NumberValue `hcl:"priority,attr"`
	// ServiceTag: string, optional
	ServiceTag terra.StringValue `hcl:"service_tag,attr"`
	// VirtualNetworkSubnetId: string, optional
	VirtualNetworkSubnetId terra.StringValue `hcl:"virtual_network_subnet_id,attr"`
	// IpRestrictionHeaders: min=0
	Headers []IpRestrictionHeaders `hcl:"headers,block" validate:"min=0"`
}

type IpRestrictionHeaders struct {
	// XAzureFdid: set of string, optional
	XAzureFdid terra.SetValue[terra.StringValue] `hcl:"x_azure_fdid,attr"`
	// XFdHealthProbe: set of string, optional
	XFdHealthProbe terra.SetValue[terra.StringValue] `hcl:"x_fd_health_probe,attr"`
	// XForwardedFor: set of string, optional
	XForwardedFor terra.SetValue[terra.StringValue] `hcl:"x_forwarded_for,attr"`
	// XForwardedHost: set of string, optional
	XForwardedHost terra.SetValue[terra.StringValue] `hcl:"x_forwarded_host,attr"`
}

type ScmIpRestriction struct {
	// Action: string, optional
	Action terra.StringValue `hcl:"action,attr"`
	// IpAddress: string, optional
	IpAddress terra.StringValue `hcl:"ip_address,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Priority: number, optional
	Priority terra.NumberValue `hcl:"priority,attr"`
	// ServiceTag: string, optional
	ServiceTag terra.StringValue `hcl:"service_tag,attr"`
	// VirtualNetworkSubnetId: string, optional
	VirtualNetworkSubnetId terra.StringValue `hcl:"virtual_network_subnet_id,attr"`
	// ScmIpRestrictionHeaders: min=0
	Headers []ScmIpRestrictionHeaders `hcl:"headers,block" validate:"min=0"`
}

type ScmIpRestrictionHeaders struct {
	// XAzureFdid: set of string, optional
	XAzureFdid terra.SetValue[terra.StringValue] `hcl:"x_azure_fdid,attr"`
	// XFdHealthProbe: set of string, optional
	XFdHealthProbe terra.SetValue[terra.StringValue] `hcl:"x_fd_health_probe,attr"`
	// XForwardedFor: set of string, optional
	XForwardedFor terra.SetValue[terra.StringValue] `hcl:"x_forwarded_for,attr"`
	// XForwardedHost: set of string, optional
	XForwardedHost terra.SetValue[terra.StringValue] `hcl:"x_forwarded_host,attr"`
}

type Cors struct {
	// AllowedOrigins: set of string, required
	AllowedOrigins terra.SetValue[terra.StringValue] `hcl:"allowed_origins,attr" validate:"required"`
	// SupportCredentials: bool, optional
	SupportCredentials terra.BoolValue `hcl:"support_credentials,attr"`
}

type SourceControl struct {
	// Branch: string, optional
	Branch terra.StringValue `hcl:"branch,attr"`
	// ManualIntegration: bool, optional
	ManualIntegration terra.BoolValue `hcl:"manual_integration,attr"`
	// RepoUrl: string, optional
	RepoUrl terra.StringValue `hcl:"repo_url,attr"`
	// RollbackEnabled: bool, optional
	RollbackEnabled terra.BoolValue `hcl:"rollback_enabled,attr"`
	// UseMercurial: bool, optional
	UseMercurial terra.BoolValue `hcl:"use_mercurial,attr"`
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

type SiteCredentialAttributes struct {
	ref terra.Reference
}

func (sc SiteCredentialAttributes) InternalRef() (terra.Reference, error) {
	return sc.ref, nil
}

func (sc SiteCredentialAttributes) InternalWithRef(ref terra.Reference) SiteCredentialAttributes {
	return SiteCredentialAttributes{ref: ref}
}

func (sc SiteCredentialAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sc.ref.InternalTokens()
}

func (sc SiteCredentialAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("password"))
}

func (sc SiteCredentialAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("username"))
}

type AuthSettingsAttributes struct {
	ref terra.Reference
}

func (as AuthSettingsAttributes) InternalRef() (terra.Reference, error) {
	return as.ref, nil
}

func (as AuthSettingsAttributes) InternalWithRef(ref terra.Reference) AuthSettingsAttributes {
	return AuthSettingsAttributes{ref: ref}
}

func (as AuthSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return as.ref.InternalTokens()
}

func (as AuthSettingsAttributes) AdditionalLoginParams() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](as.ref.Append("additional_login_params"))
}

func (as AuthSettingsAttributes) AllowedExternalRedirectUrls() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](as.ref.Append("allowed_external_redirect_urls"))
}

func (as AuthSettingsAttributes) DefaultProvider() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("default_provider"))
}

func (as AuthSettingsAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(as.ref.Append("enabled"))
}

func (as AuthSettingsAttributes) Issuer() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("issuer"))
}

func (as AuthSettingsAttributes) RuntimeVersion() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("runtime_version"))
}

func (as AuthSettingsAttributes) TokenRefreshExtensionHours() terra.NumberValue {
	return terra.ReferenceAsNumber(as.ref.Append("token_refresh_extension_hours"))
}

func (as AuthSettingsAttributes) TokenStoreEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(as.ref.Append("token_store_enabled"))
}

func (as AuthSettingsAttributes) UnauthenticatedClientAction() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("unauthenticated_client_action"))
}

func (as AuthSettingsAttributes) ActiveDirectory() terra.ListValue[ActiveDirectoryAttributes] {
	return terra.ReferenceAsList[ActiveDirectoryAttributes](as.ref.Append("active_directory"))
}

func (as AuthSettingsAttributes) Facebook() terra.ListValue[FacebookAttributes] {
	return terra.ReferenceAsList[FacebookAttributes](as.ref.Append("facebook"))
}

func (as AuthSettingsAttributes) Google() terra.ListValue[GoogleAttributes] {
	return terra.ReferenceAsList[GoogleAttributes](as.ref.Append("google"))
}

func (as AuthSettingsAttributes) Microsoft() terra.ListValue[MicrosoftAttributes] {
	return terra.ReferenceAsList[MicrosoftAttributes](as.ref.Append("microsoft"))
}

func (as AuthSettingsAttributes) Twitter() terra.ListValue[TwitterAttributes] {
	return terra.ReferenceAsList[TwitterAttributes](as.ref.Append("twitter"))
}

type ActiveDirectoryAttributes struct {
	ref terra.Reference
}

func (ad ActiveDirectoryAttributes) InternalRef() (terra.Reference, error) {
	return ad.ref, nil
}

func (ad ActiveDirectoryAttributes) InternalWithRef(ref terra.Reference) ActiveDirectoryAttributes {
	return ActiveDirectoryAttributes{ref: ref}
}

func (ad ActiveDirectoryAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ad.ref.InternalTokens()
}

func (ad ActiveDirectoryAttributes) AllowedAudiences() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ad.ref.Append("allowed_audiences"))
}

func (ad ActiveDirectoryAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(ad.ref.Append("client_id"))
}

func (ad ActiveDirectoryAttributes) ClientSecret() terra.StringValue {
	return terra.ReferenceAsString(ad.ref.Append("client_secret"))
}

type FacebookAttributes struct {
	ref terra.Reference
}

func (f FacebookAttributes) InternalRef() (terra.Reference, error) {
	return f.ref, nil
}

func (f FacebookAttributes) InternalWithRef(ref terra.Reference) FacebookAttributes {
	return FacebookAttributes{ref: ref}
}

func (f FacebookAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return f.ref.InternalTokens()
}

func (f FacebookAttributes) AppId() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("app_id"))
}

func (f FacebookAttributes) AppSecret() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("app_secret"))
}

func (f FacebookAttributes) OauthScopes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](f.ref.Append("oauth_scopes"))
}

type GoogleAttributes struct {
	ref terra.Reference
}

func (g GoogleAttributes) InternalRef() (terra.Reference, error) {
	return g.ref, nil
}

func (g GoogleAttributes) InternalWithRef(ref terra.Reference) GoogleAttributes {
	return GoogleAttributes{ref: ref}
}

func (g GoogleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return g.ref.InternalTokens()
}

func (g GoogleAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("client_id"))
}

func (g GoogleAttributes) ClientSecret() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("client_secret"))
}

func (g GoogleAttributes) OauthScopes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](g.ref.Append("oauth_scopes"))
}

type MicrosoftAttributes struct {
	ref terra.Reference
}

func (m MicrosoftAttributes) InternalRef() (terra.Reference, error) {
	return m.ref, nil
}

func (m MicrosoftAttributes) InternalWithRef(ref terra.Reference) MicrosoftAttributes {
	return MicrosoftAttributes{ref: ref}
}

func (m MicrosoftAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return m.ref.InternalTokens()
}

func (m MicrosoftAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("client_id"))
}

func (m MicrosoftAttributes) ClientSecret() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("client_secret"))
}

func (m MicrosoftAttributes) OauthScopes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](m.ref.Append("oauth_scopes"))
}

type TwitterAttributes struct {
	ref terra.Reference
}

func (t TwitterAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TwitterAttributes) InternalWithRef(ref terra.Reference) TwitterAttributes {
	return TwitterAttributes{ref: ref}
}

func (t TwitterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TwitterAttributes) ConsumerKey() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("consumer_key"))
}

func (t TwitterAttributes) ConsumerSecret() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("consumer_secret"))
}

type ConnectionStringAttributes struct {
	ref terra.Reference
}

func (cs ConnectionStringAttributes) InternalRef() (terra.Reference, error) {
	return cs.ref, nil
}

func (cs ConnectionStringAttributes) InternalWithRef(ref terra.Reference) ConnectionStringAttributes {
	return ConnectionStringAttributes{ref: ref}
}

func (cs ConnectionStringAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cs.ref.InternalTokens()
}

func (cs ConnectionStringAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("name"))
}

func (cs ConnectionStringAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("type"))
}

func (cs ConnectionStringAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("value"))
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

type SiteConfigAttributes struct {
	ref terra.Reference
}

func (sc SiteConfigAttributes) InternalRef() (terra.Reference, error) {
	return sc.ref, nil
}

func (sc SiteConfigAttributes) InternalWithRef(ref terra.Reference) SiteConfigAttributes {
	return SiteConfigAttributes{ref: ref}
}

func (sc SiteConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sc.ref.InternalTokens()
}

func (sc SiteConfigAttributes) AlwaysOn() terra.BoolValue {
	return terra.ReferenceAsBool(sc.ref.Append("always_on"))
}

func (sc SiteConfigAttributes) AppScaleLimit() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("app_scale_limit"))
}

func (sc SiteConfigAttributes) AutoSwapSlotName() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("auto_swap_slot_name"))
}

func (sc SiteConfigAttributes) DotnetFrameworkVersion() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("dotnet_framework_version"))
}

func (sc SiteConfigAttributes) ElasticInstanceMinimum() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("elastic_instance_minimum"))
}

func (sc SiteConfigAttributes) FtpsState() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("ftps_state"))
}

func (sc SiteConfigAttributes) HealthCheckPath() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("health_check_path"))
}

func (sc SiteConfigAttributes) Http2Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(sc.ref.Append("http2_enabled"))
}

func (sc SiteConfigAttributes) JavaVersion() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("java_version"))
}

func (sc SiteConfigAttributes) LinuxFxVersion() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("linux_fx_version"))
}

func (sc SiteConfigAttributes) MinTlsVersion() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("min_tls_version"))
}

func (sc SiteConfigAttributes) PreWarmedInstanceCount() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("pre_warmed_instance_count"))
}

func (sc SiteConfigAttributes) RuntimeScaleMonitoringEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sc.ref.Append("runtime_scale_monitoring_enabled"))
}

func (sc SiteConfigAttributes) ScmType() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("scm_type"))
}

func (sc SiteConfigAttributes) ScmUseMainIpRestriction() terra.BoolValue {
	return terra.ReferenceAsBool(sc.ref.Append("scm_use_main_ip_restriction"))
}

func (sc SiteConfigAttributes) Use32BitWorkerProcess() terra.BoolValue {
	return terra.ReferenceAsBool(sc.ref.Append("use_32_bit_worker_process"))
}

func (sc SiteConfigAttributes) VnetRouteAllEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sc.ref.Append("vnet_route_all_enabled"))
}

func (sc SiteConfigAttributes) WebsocketsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sc.ref.Append("websockets_enabled"))
}

func (sc SiteConfigAttributes) IpRestriction() terra.ListValue[IpRestrictionAttributes] {
	return terra.ReferenceAsList[IpRestrictionAttributes](sc.ref.Append("ip_restriction"))
}

func (sc SiteConfigAttributes) ScmIpRestriction() terra.ListValue[ScmIpRestrictionAttributes] {
	return terra.ReferenceAsList[ScmIpRestrictionAttributes](sc.ref.Append("scm_ip_restriction"))
}

func (sc SiteConfigAttributes) Cors() terra.ListValue[CorsAttributes] {
	return terra.ReferenceAsList[CorsAttributes](sc.ref.Append("cors"))
}

type IpRestrictionAttributes struct {
	ref terra.Reference
}

func (ir IpRestrictionAttributes) InternalRef() (terra.Reference, error) {
	return ir.ref, nil
}

func (ir IpRestrictionAttributes) InternalWithRef(ref terra.Reference) IpRestrictionAttributes {
	return IpRestrictionAttributes{ref: ref}
}

func (ir IpRestrictionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ir.ref.InternalTokens()
}

func (ir IpRestrictionAttributes) Action() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("action"))
}

func (ir IpRestrictionAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("ip_address"))
}

func (ir IpRestrictionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("name"))
}

func (ir IpRestrictionAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(ir.ref.Append("priority"))
}

func (ir IpRestrictionAttributes) ServiceTag() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("service_tag"))
}

func (ir IpRestrictionAttributes) VirtualNetworkSubnetId() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("virtual_network_subnet_id"))
}

func (ir IpRestrictionAttributes) Headers() terra.ListValue[IpRestrictionHeadersAttributes] {
	return terra.ReferenceAsList[IpRestrictionHeadersAttributes](ir.ref.Append("headers"))
}

type IpRestrictionHeadersAttributes struct {
	ref terra.Reference
}

func (h IpRestrictionHeadersAttributes) InternalRef() (terra.Reference, error) {
	return h.ref, nil
}

func (h IpRestrictionHeadersAttributes) InternalWithRef(ref terra.Reference) IpRestrictionHeadersAttributes {
	return IpRestrictionHeadersAttributes{ref: ref}
}

func (h IpRestrictionHeadersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return h.ref.InternalTokens()
}

func (h IpRestrictionHeadersAttributes) XAzureFdid() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](h.ref.Append("x_azure_fdid"))
}

func (h IpRestrictionHeadersAttributes) XFdHealthProbe() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](h.ref.Append("x_fd_health_probe"))
}

func (h IpRestrictionHeadersAttributes) XForwardedFor() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](h.ref.Append("x_forwarded_for"))
}

func (h IpRestrictionHeadersAttributes) XForwardedHost() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](h.ref.Append("x_forwarded_host"))
}

type ScmIpRestrictionAttributes struct {
	ref terra.Reference
}

func (sir ScmIpRestrictionAttributes) InternalRef() (terra.Reference, error) {
	return sir.ref, nil
}

func (sir ScmIpRestrictionAttributes) InternalWithRef(ref terra.Reference) ScmIpRestrictionAttributes {
	return ScmIpRestrictionAttributes{ref: ref}
}

func (sir ScmIpRestrictionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sir.ref.InternalTokens()
}

func (sir ScmIpRestrictionAttributes) Action() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("action"))
}

func (sir ScmIpRestrictionAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("ip_address"))
}

func (sir ScmIpRestrictionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("name"))
}

func (sir ScmIpRestrictionAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(sir.ref.Append("priority"))
}

func (sir ScmIpRestrictionAttributes) ServiceTag() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("service_tag"))
}

func (sir ScmIpRestrictionAttributes) VirtualNetworkSubnetId() terra.StringValue {
	return terra.ReferenceAsString(sir.ref.Append("virtual_network_subnet_id"))
}

func (sir ScmIpRestrictionAttributes) Headers() terra.ListValue[ScmIpRestrictionHeadersAttributes] {
	return terra.ReferenceAsList[ScmIpRestrictionHeadersAttributes](sir.ref.Append("headers"))
}

type ScmIpRestrictionHeadersAttributes struct {
	ref terra.Reference
}

func (h ScmIpRestrictionHeadersAttributes) InternalRef() (terra.Reference, error) {
	return h.ref, nil
}

func (h ScmIpRestrictionHeadersAttributes) InternalWithRef(ref terra.Reference) ScmIpRestrictionHeadersAttributes {
	return ScmIpRestrictionHeadersAttributes{ref: ref}
}

func (h ScmIpRestrictionHeadersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return h.ref.InternalTokens()
}

func (h ScmIpRestrictionHeadersAttributes) XAzureFdid() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](h.ref.Append("x_azure_fdid"))
}

func (h ScmIpRestrictionHeadersAttributes) XFdHealthProbe() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](h.ref.Append("x_fd_health_probe"))
}

func (h ScmIpRestrictionHeadersAttributes) XForwardedFor() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](h.ref.Append("x_forwarded_for"))
}

func (h ScmIpRestrictionHeadersAttributes) XForwardedHost() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](h.ref.Append("x_forwarded_host"))
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

func (c CorsAttributes) SupportCredentials() terra.BoolValue {
	return terra.ReferenceAsBool(c.ref.Append("support_credentials"))
}

type SourceControlAttributes struct {
	ref terra.Reference
}

func (sc SourceControlAttributes) InternalRef() (terra.Reference, error) {
	return sc.ref, nil
}

func (sc SourceControlAttributes) InternalWithRef(ref terra.Reference) SourceControlAttributes {
	return SourceControlAttributes{ref: ref}
}

func (sc SourceControlAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sc.ref.InternalTokens()
}

func (sc SourceControlAttributes) Branch() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("branch"))
}

func (sc SourceControlAttributes) ManualIntegration() terra.BoolValue {
	return terra.ReferenceAsBool(sc.ref.Append("manual_integration"))
}

func (sc SourceControlAttributes) RepoUrl() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("repo_url"))
}

func (sc SourceControlAttributes) RollbackEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sc.ref.Append("rollback_enabled"))
}

func (sc SourceControlAttributes) UseMercurial() terra.BoolValue {
	return terra.ReferenceAsBool(sc.ref.Append("use_mercurial"))
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

type SiteCredentialState struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type AuthSettingsState struct {
	AdditionalLoginParams       map[string]string      `json:"additional_login_params"`
	AllowedExternalRedirectUrls []string               `json:"allowed_external_redirect_urls"`
	DefaultProvider             string                 `json:"default_provider"`
	Enabled                     bool                   `json:"enabled"`
	Issuer                      string                 `json:"issuer"`
	RuntimeVersion              string                 `json:"runtime_version"`
	TokenRefreshExtensionHours  float64                `json:"token_refresh_extension_hours"`
	TokenStoreEnabled           bool                   `json:"token_store_enabled"`
	UnauthenticatedClientAction string                 `json:"unauthenticated_client_action"`
	ActiveDirectory             []ActiveDirectoryState `json:"active_directory"`
	Facebook                    []FacebookState        `json:"facebook"`
	Google                      []GoogleState          `json:"google"`
	Microsoft                   []MicrosoftState       `json:"microsoft"`
	Twitter                     []TwitterState         `json:"twitter"`
}

type ActiveDirectoryState struct {
	AllowedAudiences []string `json:"allowed_audiences"`
	ClientId         string   `json:"client_id"`
	ClientSecret     string   `json:"client_secret"`
}

type FacebookState struct {
	AppId       string   `json:"app_id"`
	AppSecret   string   `json:"app_secret"`
	OauthScopes []string `json:"oauth_scopes"`
}

type GoogleState struct {
	ClientId     string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
	OauthScopes  []string `json:"oauth_scopes"`
}

type MicrosoftState struct {
	ClientId     string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
	OauthScopes  []string `json:"oauth_scopes"`
}

type TwitterState struct {
	ConsumerKey    string `json:"consumer_key"`
	ConsumerSecret string `json:"consumer_secret"`
}

type ConnectionStringState struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type IdentityState struct {
	IdentityIds []string `json:"identity_ids"`
	PrincipalId string   `json:"principal_id"`
	TenantId    string   `json:"tenant_id"`
	Type        string   `json:"type"`
}

type SiteConfigState struct {
	AlwaysOn                      bool                    `json:"always_on"`
	AppScaleLimit                 float64                 `json:"app_scale_limit"`
	AutoSwapSlotName              string                  `json:"auto_swap_slot_name"`
	DotnetFrameworkVersion        string                  `json:"dotnet_framework_version"`
	ElasticInstanceMinimum        float64                 `json:"elastic_instance_minimum"`
	FtpsState                     string                  `json:"ftps_state"`
	HealthCheckPath               string                  `json:"health_check_path"`
	Http2Enabled                  bool                    `json:"http2_enabled"`
	JavaVersion                   string                  `json:"java_version"`
	LinuxFxVersion                string                  `json:"linux_fx_version"`
	MinTlsVersion                 string                  `json:"min_tls_version"`
	PreWarmedInstanceCount        float64                 `json:"pre_warmed_instance_count"`
	RuntimeScaleMonitoringEnabled bool                    `json:"runtime_scale_monitoring_enabled"`
	ScmType                       string                  `json:"scm_type"`
	ScmUseMainIpRestriction       bool                    `json:"scm_use_main_ip_restriction"`
	Use32BitWorkerProcess         bool                    `json:"use_32_bit_worker_process"`
	VnetRouteAllEnabled           bool                    `json:"vnet_route_all_enabled"`
	WebsocketsEnabled             bool                    `json:"websockets_enabled"`
	IpRestriction                 []IpRestrictionState    `json:"ip_restriction"`
	ScmIpRestriction              []ScmIpRestrictionState `json:"scm_ip_restriction"`
	Cors                          []CorsState             `json:"cors"`
}

type IpRestrictionState struct {
	Action                 string                      `json:"action"`
	IpAddress              string                      `json:"ip_address"`
	Name                   string                      `json:"name"`
	Priority               float64                     `json:"priority"`
	ServiceTag             string                      `json:"service_tag"`
	VirtualNetworkSubnetId string                      `json:"virtual_network_subnet_id"`
	Headers                []IpRestrictionHeadersState `json:"headers"`
}

type IpRestrictionHeadersState struct {
	XAzureFdid     []string `json:"x_azure_fdid"`
	XFdHealthProbe []string `json:"x_fd_health_probe"`
	XForwardedFor  []string `json:"x_forwarded_for"`
	XForwardedHost []string `json:"x_forwarded_host"`
}

type ScmIpRestrictionState struct {
	Action                 string                         `json:"action"`
	IpAddress              string                         `json:"ip_address"`
	Name                   string                         `json:"name"`
	Priority               float64                        `json:"priority"`
	ServiceTag             string                         `json:"service_tag"`
	VirtualNetworkSubnetId string                         `json:"virtual_network_subnet_id"`
	Headers                []ScmIpRestrictionHeadersState `json:"headers"`
}

type ScmIpRestrictionHeadersState struct {
	XAzureFdid     []string `json:"x_azure_fdid"`
	XFdHealthProbe []string `json:"x_fd_health_probe"`
	XForwardedFor  []string `json:"x_forwarded_for"`
	XForwardedHost []string `json:"x_forwarded_host"`
}

type CorsState struct {
	AllowedOrigins     []string `json:"allowed_origins"`
	SupportCredentials bool     `json:"support_credentials"`
}

type SourceControlState struct {
	Branch            string `json:"branch"`
	ManualIntegration bool   `json:"manual_integration"`
	RepoUrl           string `json:"repo_url"`
	RollbackEnabled   bool   `json:"rollback_enabled"`
	UseMercurial      bool   `json:"use_mercurial"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
