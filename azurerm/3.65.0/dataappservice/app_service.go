// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataappservice

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ConnectionString struct{}

type SiteConfig struct {
	// Cors: min=0
	Cors []Cors `hcl:"cors,block" validate:"min=0"`
	// IpRestriction: min=0
	IpRestriction []IpRestriction `hcl:"ip_restriction,block" validate:"min=0"`
	// ScmIpRestriction: min=0
	ScmIpRestriction []ScmIpRestriction `hcl:"scm_ip_restriction,block" validate:"min=0"`
}

type Cors struct{}

type IpRestriction struct {
	// IpRestrictionHeaders: min=0
	Headers []IpRestrictionHeaders `hcl:"headers,block" validate:"min=0"`
}

type IpRestrictionHeaders struct{}

type ScmIpRestriction struct {
	// ScmIpRestrictionHeaders: min=0
	Headers []ScmIpRestrictionHeaders `hcl:"headers,block" validate:"min=0"`
}

type ScmIpRestrictionHeaders struct{}

type SiteCredential struct{}

type SourceControl struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
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

func (sc SiteConfigAttributes) AcrUseManagedIdentityCredentials() terra.BoolValue {
	return terra.ReferenceAsBool(sc.ref.Append("acr_use_managed_identity_credentials"))
}

func (sc SiteConfigAttributes) AcrUserManagedIdentityClientId() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("acr_user_managed_identity_client_id"))
}

func (sc SiteConfigAttributes) AlwaysOn() terra.BoolValue {
	return terra.ReferenceAsBool(sc.ref.Append("always_on"))
}

func (sc SiteConfigAttributes) AppCommandLine() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("app_command_line"))
}

func (sc SiteConfigAttributes) DefaultDocuments() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sc.ref.Append("default_documents"))
}

func (sc SiteConfigAttributes) DotnetFrameworkVersion() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("dotnet_framework_version"))
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

func (sc SiteConfigAttributes) JavaContainer() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("java_container"))
}

func (sc SiteConfigAttributes) JavaContainerVersion() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("java_container_version"))
}

func (sc SiteConfigAttributes) JavaVersion() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("java_version"))
}

func (sc SiteConfigAttributes) LinuxFxVersion() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("linux_fx_version"))
}

func (sc SiteConfigAttributes) LocalMysqlEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sc.ref.Append("local_mysql_enabled"))
}

func (sc SiteConfigAttributes) ManagedPipelineMode() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("managed_pipeline_mode"))
}

func (sc SiteConfigAttributes) MinTlsVersion() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("min_tls_version"))
}

func (sc SiteConfigAttributes) NumberOfWorkers() terra.NumberValue {
	return terra.ReferenceAsNumber(sc.ref.Append("number_of_workers"))
}

func (sc SiteConfigAttributes) PhpVersion() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("php_version"))
}

func (sc SiteConfigAttributes) PythonVersion() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("python_version"))
}

func (sc SiteConfigAttributes) RemoteDebuggingEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sc.ref.Append("remote_debugging_enabled"))
}

func (sc SiteConfigAttributes) RemoteDebuggingVersion() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("remote_debugging_version"))
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

func (sc SiteConfigAttributes) WindowsFxVersion() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("windows_fx_version"))
}

func (sc SiteConfigAttributes) Cors() terra.ListValue[CorsAttributes] {
	return terra.ReferenceAsList[CorsAttributes](sc.ref.Append("cors"))
}

func (sc SiteConfigAttributes) IpRestriction() terra.ListValue[IpRestrictionAttributes] {
	return terra.ReferenceAsList[IpRestrictionAttributes](sc.ref.Append("ip_restriction"))
}

func (sc SiteConfigAttributes) ScmIpRestriction() terra.ListValue[ScmIpRestrictionAttributes] {
	return terra.ReferenceAsList[ScmIpRestrictionAttributes](sc.ref.Append("scm_ip_restriction"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type ConnectionStringState struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type SiteConfigState struct {
	AcrUseManagedIdentityCredentials bool                    `json:"acr_use_managed_identity_credentials"`
	AcrUserManagedIdentityClientId   string                  `json:"acr_user_managed_identity_client_id"`
	AlwaysOn                         bool                    `json:"always_on"`
	AppCommandLine                   string                  `json:"app_command_line"`
	DefaultDocuments                 []string                `json:"default_documents"`
	DotnetFrameworkVersion           string                  `json:"dotnet_framework_version"`
	FtpsState                        string                  `json:"ftps_state"`
	HealthCheckPath                  string                  `json:"health_check_path"`
	Http2Enabled                     bool                    `json:"http2_enabled"`
	JavaContainer                    string                  `json:"java_container"`
	JavaContainerVersion             string                  `json:"java_container_version"`
	JavaVersion                      string                  `json:"java_version"`
	LinuxFxVersion                   string                  `json:"linux_fx_version"`
	LocalMysqlEnabled                bool                    `json:"local_mysql_enabled"`
	ManagedPipelineMode              string                  `json:"managed_pipeline_mode"`
	MinTlsVersion                    string                  `json:"min_tls_version"`
	NumberOfWorkers                  float64                 `json:"number_of_workers"`
	PhpVersion                       string                  `json:"php_version"`
	PythonVersion                    string                  `json:"python_version"`
	RemoteDebuggingEnabled           bool                    `json:"remote_debugging_enabled"`
	RemoteDebuggingVersion           string                  `json:"remote_debugging_version"`
	ScmType                          string                  `json:"scm_type"`
	ScmUseMainIpRestriction          bool                    `json:"scm_use_main_ip_restriction"`
	Use32BitWorkerProcess            bool                    `json:"use_32_bit_worker_process"`
	VnetRouteAllEnabled              bool                    `json:"vnet_route_all_enabled"`
	WebsocketsEnabled                bool                    `json:"websockets_enabled"`
	WindowsFxVersion                 string                  `json:"windows_fx_version"`
	Cors                             []CorsState             `json:"cors"`
	IpRestriction                    []IpRestrictionState    `json:"ip_restriction"`
	ScmIpRestriction                 []ScmIpRestrictionState `json:"scm_ip_restriction"`
}

type CorsState struct {
	AllowedOrigins     []string `json:"allowed_origins"`
	SupportCredentials bool     `json:"support_credentials"`
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

type SiteCredentialState struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type SourceControlState struct {
	Branch            string `json:"branch"`
	ManualIntegration bool   `json:"manual_integration"`
	RepoUrl           string `json:"repo_url"`
	RollbackEnabled   bool   `json:"rollback_enabled"`
	UseMercurial      bool   `json:"use_mercurial"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}