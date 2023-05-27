// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package springcloudservice

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type RequiredNetworkTrafficRules struct{}

type ConfigServerGitSetting struct {
	// Label: string, optional
	Label terra.StringValue `hcl:"label,attr"`
	// SearchPaths: list of string, optional
	SearchPaths terra.ListValue[terra.StringValue] `hcl:"search_paths,attr"`
	// Uri: string, required
	Uri terra.StringValue `hcl:"uri,attr" validate:"required"`
	// ConfigServerGitSettingHttpBasicAuth: optional
	HttpBasicAuth *ConfigServerGitSettingHttpBasicAuth `hcl:"http_basic_auth,block"`
	// Repository: min=0
	Repository []Repository `hcl:"repository,block" validate:"min=0"`
	// ConfigServerGitSettingSshAuth: optional
	SshAuth *ConfigServerGitSettingSshAuth `hcl:"ssh_auth,block"`
}

type ConfigServerGitSettingHttpBasicAuth struct {
	// Password: string, required
	Password terra.StringValue `hcl:"password,attr" validate:"required"`
	// Username: string, required
	Username terra.StringValue `hcl:"username,attr" validate:"required"`
}

type Repository struct {
	// Label: string, optional
	Label terra.StringValue `hcl:"label,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Pattern: list of string, optional
	Pattern terra.ListValue[terra.StringValue] `hcl:"pattern,attr"`
	// SearchPaths: list of string, optional
	SearchPaths terra.ListValue[terra.StringValue] `hcl:"search_paths,attr"`
	// Uri: string, required
	Uri terra.StringValue `hcl:"uri,attr" validate:"required"`
	// RepositoryHttpBasicAuth: optional
	HttpBasicAuth *RepositoryHttpBasicAuth `hcl:"http_basic_auth,block"`
	// RepositorySshAuth: optional
	SshAuth *RepositorySshAuth `hcl:"ssh_auth,block"`
}

type RepositoryHttpBasicAuth struct {
	// Password: string, required
	Password terra.StringValue `hcl:"password,attr" validate:"required"`
	// Username: string, required
	Username terra.StringValue `hcl:"username,attr" validate:"required"`
}

type RepositorySshAuth struct {
	// HostKey: string, optional
	HostKey terra.StringValue `hcl:"host_key,attr"`
	// HostKeyAlgorithm: string, optional
	HostKeyAlgorithm terra.StringValue `hcl:"host_key_algorithm,attr"`
	// PrivateKey: string, required
	PrivateKey terra.StringValue `hcl:"private_key,attr" validate:"required"`
	// StrictHostKeyCheckingEnabled: bool, optional
	StrictHostKeyCheckingEnabled terra.BoolValue `hcl:"strict_host_key_checking_enabled,attr"`
}

type ConfigServerGitSettingSshAuth struct {
	// HostKey: string, optional
	HostKey terra.StringValue `hcl:"host_key,attr"`
	// HostKeyAlgorithm: string, optional
	HostKeyAlgorithm terra.StringValue `hcl:"host_key_algorithm,attr"`
	// PrivateKey: string, required
	PrivateKey terra.StringValue `hcl:"private_key,attr" validate:"required"`
	// StrictHostKeyCheckingEnabled: bool, optional
	StrictHostKeyCheckingEnabled terra.BoolValue `hcl:"strict_host_key_checking_enabled,attr"`
}

type Network struct {
	// AppNetworkResourceGroup: string, optional
	AppNetworkResourceGroup terra.StringValue `hcl:"app_network_resource_group,attr"`
	// AppSubnetId: string, required
	AppSubnetId terra.StringValue `hcl:"app_subnet_id,attr" validate:"required"`
	// CidrRanges: list of string, required
	CidrRanges terra.ListValue[terra.StringValue] `hcl:"cidr_ranges,attr" validate:"required"`
	// ReadTimeoutSeconds: number, optional
	ReadTimeoutSeconds terra.NumberValue `hcl:"read_timeout_seconds,attr"`
	// ServiceRuntimeNetworkResourceGroup: string, optional
	ServiceRuntimeNetworkResourceGroup terra.StringValue `hcl:"service_runtime_network_resource_group,attr"`
	// ServiceRuntimeSubnetId: string, required
	ServiceRuntimeSubnetId terra.StringValue `hcl:"service_runtime_subnet_id,attr" validate:"required"`
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

type Trace struct {
	// ConnectionString: string, optional
	ConnectionString terra.StringValue `hcl:"connection_string,attr"`
	// SampleRate: number, optional
	SampleRate terra.NumberValue `hcl:"sample_rate,attr"`
}

type RequiredNetworkTrafficRulesAttributes struct {
	ref terra.Reference
}

func (rntr RequiredNetworkTrafficRulesAttributes) InternalRef() (terra.Reference, error) {
	return rntr.ref, nil
}

func (rntr RequiredNetworkTrafficRulesAttributes) InternalWithRef(ref terra.Reference) RequiredNetworkTrafficRulesAttributes {
	return RequiredNetworkTrafficRulesAttributes{ref: ref}
}

func (rntr RequiredNetworkTrafficRulesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rntr.ref.InternalTokens()
}

func (rntr RequiredNetworkTrafficRulesAttributes) Direction() terra.StringValue {
	return terra.ReferenceAsString(rntr.ref.Append("direction"))
}

func (rntr RequiredNetworkTrafficRulesAttributes) Fqdns() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](rntr.ref.Append("fqdns"))
}

func (rntr RequiredNetworkTrafficRulesAttributes) IpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](rntr.ref.Append("ip_addresses"))
}

func (rntr RequiredNetworkTrafficRulesAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(rntr.ref.Append("port"))
}

func (rntr RequiredNetworkTrafficRulesAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(rntr.ref.Append("protocol"))
}

type ConfigServerGitSettingAttributes struct {
	ref terra.Reference
}

func (csgs ConfigServerGitSettingAttributes) InternalRef() (terra.Reference, error) {
	return csgs.ref, nil
}

func (csgs ConfigServerGitSettingAttributes) InternalWithRef(ref terra.Reference) ConfigServerGitSettingAttributes {
	return ConfigServerGitSettingAttributes{ref: ref}
}

func (csgs ConfigServerGitSettingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return csgs.ref.InternalTokens()
}

func (csgs ConfigServerGitSettingAttributes) Label() terra.StringValue {
	return terra.ReferenceAsString(csgs.ref.Append("label"))
}

func (csgs ConfigServerGitSettingAttributes) SearchPaths() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](csgs.ref.Append("search_paths"))
}

func (csgs ConfigServerGitSettingAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(csgs.ref.Append("uri"))
}

func (csgs ConfigServerGitSettingAttributes) HttpBasicAuth() terra.ListValue[ConfigServerGitSettingHttpBasicAuthAttributes] {
	return terra.ReferenceAsList[ConfigServerGitSettingHttpBasicAuthAttributes](csgs.ref.Append("http_basic_auth"))
}

func (csgs ConfigServerGitSettingAttributes) Repository() terra.ListValue[RepositoryAttributes] {
	return terra.ReferenceAsList[RepositoryAttributes](csgs.ref.Append("repository"))
}

func (csgs ConfigServerGitSettingAttributes) SshAuth() terra.ListValue[ConfigServerGitSettingSshAuthAttributes] {
	return terra.ReferenceAsList[ConfigServerGitSettingSshAuthAttributes](csgs.ref.Append("ssh_auth"))
}

type ConfigServerGitSettingHttpBasicAuthAttributes struct {
	ref terra.Reference
}

func (hba ConfigServerGitSettingHttpBasicAuthAttributes) InternalRef() (terra.Reference, error) {
	return hba.ref, nil
}

func (hba ConfigServerGitSettingHttpBasicAuthAttributes) InternalWithRef(ref terra.Reference) ConfigServerGitSettingHttpBasicAuthAttributes {
	return ConfigServerGitSettingHttpBasicAuthAttributes{ref: ref}
}

func (hba ConfigServerGitSettingHttpBasicAuthAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hba.ref.InternalTokens()
}

func (hba ConfigServerGitSettingHttpBasicAuthAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(hba.ref.Append("password"))
}

func (hba ConfigServerGitSettingHttpBasicAuthAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(hba.ref.Append("username"))
}

type RepositoryAttributes struct {
	ref terra.Reference
}

func (r RepositoryAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RepositoryAttributes) InternalWithRef(ref terra.Reference) RepositoryAttributes {
	return RepositoryAttributes{ref: ref}
}

func (r RepositoryAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RepositoryAttributes) Label() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("label"))
}

func (r RepositoryAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("name"))
}

func (r RepositoryAttributes) Pattern() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](r.ref.Append("pattern"))
}

func (r RepositoryAttributes) SearchPaths() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](r.ref.Append("search_paths"))
}

func (r RepositoryAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("uri"))
}

func (r RepositoryAttributes) HttpBasicAuth() terra.ListValue[RepositoryHttpBasicAuthAttributes] {
	return terra.ReferenceAsList[RepositoryHttpBasicAuthAttributes](r.ref.Append("http_basic_auth"))
}

func (r RepositoryAttributes) SshAuth() terra.ListValue[RepositorySshAuthAttributes] {
	return terra.ReferenceAsList[RepositorySshAuthAttributes](r.ref.Append("ssh_auth"))
}

type RepositoryHttpBasicAuthAttributes struct {
	ref terra.Reference
}

func (hba RepositoryHttpBasicAuthAttributes) InternalRef() (terra.Reference, error) {
	return hba.ref, nil
}

func (hba RepositoryHttpBasicAuthAttributes) InternalWithRef(ref terra.Reference) RepositoryHttpBasicAuthAttributes {
	return RepositoryHttpBasicAuthAttributes{ref: ref}
}

func (hba RepositoryHttpBasicAuthAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hba.ref.InternalTokens()
}

func (hba RepositoryHttpBasicAuthAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(hba.ref.Append("password"))
}

func (hba RepositoryHttpBasicAuthAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(hba.ref.Append("username"))
}

type RepositorySshAuthAttributes struct {
	ref terra.Reference
}

func (sa RepositorySshAuthAttributes) InternalRef() (terra.Reference, error) {
	return sa.ref, nil
}

func (sa RepositorySshAuthAttributes) InternalWithRef(ref terra.Reference) RepositorySshAuthAttributes {
	return RepositorySshAuthAttributes{ref: ref}
}

func (sa RepositorySshAuthAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sa.ref.InternalTokens()
}

func (sa RepositorySshAuthAttributes) HostKey() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("host_key"))
}

func (sa RepositorySshAuthAttributes) HostKeyAlgorithm() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("host_key_algorithm"))
}

func (sa RepositorySshAuthAttributes) PrivateKey() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("private_key"))
}

func (sa RepositorySshAuthAttributes) StrictHostKeyCheckingEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sa.ref.Append("strict_host_key_checking_enabled"))
}

type ConfigServerGitSettingSshAuthAttributes struct {
	ref terra.Reference
}

func (sa ConfigServerGitSettingSshAuthAttributes) InternalRef() (terra.Reference, error) {
	return sa.ref, nil
}

func (sa ConfigServerGitSettingSshAuthAttributes) InternalWithRef(ref terra.Reference) ConfigServerGitSettingSshAuthAttributes {
	return ConfigServerGitSettingSshAuthAttributes{ref: ref}
}

func (sa ConfigServerGitSettingSshAuthAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sa.ref.InternalTokens()
}

func (sa ConfigServerGitSettingSshAuthAttributes) HostKey() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("host_key"))
}

func (sa ConfigServerGitSettingSshAuthAttributes) HostKeyAlgorithm() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("host_key_algorithm"))
}

func (sa ConfigServerGitSettingSshAuthAttributes) PrivateKey() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("private_key"))
}

func (sa ConfigServerGitSettingSshAuthAttributes) StrictHostKeyCheckingEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sa.ref.Append("strict_host_key_checking_enabled"))
}

type NetworkAttributes struct {
	ref terra.Reference
}

func (n NetworkAttributes) InternalRef() (terra.Reference, error) {
	return n.ref, nil
}

func (n NetworkAttributes) InternalWithRef(ref terra.Reference) NetworkAttributes {
	return NetworkAttributes{ref: ref}
}

func (n NetworkAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return n.ref.InternalTokens()
}

func (n NetworkAttributes) AppNetworkResourceGroup() terra.StringValue {
	return terra.ReferenceAsString(n.ref.Append("app_network_resource_group"))
}

func (n NetworkAttributes) AppSubnetId() terra.StringValue {
	return terra.ReferenceAsString(n.ref.Append("app_subnet_id"))
}

func (n NetworkAttributes) CidrRanges() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](n.ref.Append("cidr_ranges"))
}

func (n NetworkAttributes) ReadTimeoutSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(n.ref.Append("read_timeout_seconds"))
}

func (n NetworkAttributes) ServiceRuntimeNetworkResourceGroup() terra.StringValue {
	return terra.ReferenceAsString(n.ref.Append("service_runtime_network_resource_group"))
}

func (n NetworkAttributes) ServiceRuntimeSubnetId() terra.StringValue {
	return terra.ReferenceAsString(n.ref.Append("service_runtime_subnet_id"))
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

type TraceAttributes struct {
	ref terra.Reference
}

func (t TraceAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TraceAttributes) InternalWithRef(ref terra.Reference) TraceAttributes {
	return TraceAttributes{ref: ref}
}

func (t TraceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TraceAttributes) ConnectionString() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("connection_string"))
}

func (t TraceAttributes) SampleRate() terra.NumberValue {
	return terra.ReferenceAsNumber(t.ref.Append("sample_rate"))
}

type RequiredNetworkTrafficRulesState struct {
	Direction   string   `json:"direction"`
	Fqdns       []string `json:"fqdns"`
	IpAddresses []string `json:"ip_addresses"`
	Port        float64  `json:"port"`
	Protocol    string   `json:"protocol"`
}

type ConfigServerGitSettingState struct {
	Label         string                                     `json:"label"`
	SearchPaths   []string                                   `json:"search_paths"`
	Uri           string                                     `json:"uri"`
	HttpBasicAuth []ConfigServerGitSettingHttpBasicAuthState `json:"http_basic_auth"`
	Repository    []RepositoryState                          `json:"repository"`
	SshAuth       []ConfigServerGitSettingSshAuthState       `json:"ssh_auth"`
}

type ConfigServerGitSettingHttpBasicAuthState struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type RepositoryState struct {
	Label         string                         `json:"label"`
	Name          string                         `json:"name"`
	Pattern       []string                       `json:"pattern"`
	SearchPaths   []string                       `json:"search_paths"`
	Uri           string                         `json:"uri"`
	HttpBasicAuth []RepositoryHttpBasicAuthState `json:"http_basic_auth"`
	SshAuth       []RepositorySshAuthState       `json:"ssh_auth"`
}

type RepositoryHttpBasicAuthState struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type RepositorySshAuthState struct {
	HostKey                      string `json:"host_key"`
	HostKeyAlgorithm             string `json:"host_key_algorithm"`
	PrivateKey                   string `json:"private_key"`
	StrictHostKeyCheckingEnabled bool   `json:"strict_host_key_checking_enabled"`
}

type ConfigServerGitSettingSshAuthState struct {
	HostKey                      string `json:"host_key"`
	HostKeyAlgorithm             string `json:"host_key_algorithm"`
	PrivateKey                   string `json:"private_key"`
	StrictHostKeyCheckingEnabled bool   `json:"strict_host_key_checking_enabled"`
}

type NetworkState struct {
	AppNetworkResourceGroup            string   `json:"app_network_resource_group"`
	AppSubnetId                        string   `json:"app_subnet_id"`
	CidrRanges                         []string `json:"cidr_ranges"`
	ReadTimeoutSeconds                 float64  `json:"read_timeout_seconds"`
	ServiceRuntimeNetworkResourceGroup string   `json:"service_runtime_network_resource_group"`
	ServiceRuntimeSubnetId             string   `json:"service_runtime_subnet_id"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}

type TraceState struct {
	ConnectionString string  `json:"connection_string"`
	SampleRate       float64 `json:"sample_rate"`
}
