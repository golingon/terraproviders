// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataspringcloudservice

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ConfigServerGitSetting struct {
	// ConfigServerGitSettingHttpBasicAuth: min=0
	HttpBasicAuth []ConfigServerGitSettingHttpBasicAuth `hcl:"http_basic_auth,block" validate:"min=0"`
	// Repository: min=0
	Repository []Repository `hcl:"repository,block" validate:"min=0"`
	// ConfigServerGitSettingSshAuth: min=0
	SshAuth []ConfigServerGitSettingSshAuth `hcl:"ssh_auth,block" validate:"min=0"`
}

type ConfigServerGitSettingHttpBasicAuth struct{}

type Repository struct {
	// RepositoryHttpBasicAuth: min=0
	HttpBasicAuth []RepositoryHttpBasicAuth `hcl:"http_basic_auth,block" validate:"min=0"`
	// RepositorySshAuth: min=0
	SshAuth []RepositorySshAuth `hcl:"ssh_auth,block" validate:"min=0"`
}

type RepositoryHttpBasicAuth struct{}

type RepositorySshAuth struct{}

type ConfigServerGitSettingSshAuth struct{}

type RequiredNetworkTrafficRules struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
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

type RequiredNetworkTrafficRulesState struct {
	Direction   string   `json:"direction"`
	Fqdns       []string `json:"fqdns"`
	IpAddresses []string `json:"ip_addresses"`
	Port        float64  `json:"port"`
	Protocol    string   `json:"protocol"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
