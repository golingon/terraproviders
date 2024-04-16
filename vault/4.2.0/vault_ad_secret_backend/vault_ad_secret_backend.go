// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vault_ad_secret_backend

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource vault_ad_secret_backend.
type Resource struct {
	Name      string
	Args      Args
	state     *vaultAdSecretBackendState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (vasb *Resource) Type() string {
	return "vault_ad_secret_backend"
}

// LocalName returns the local name for [Resource].
func (vasb *Resource) LocalName() string {
	return vasb.Name
}

// Configuration returns the configuration (args) for [Resource].
func (vasb *Resource) Configuration() interface{} {
	return vasb.Args
}

// DependOn is used for other resources to depend on [Resource].
func (vasb *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(vasb)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (vasb *Resource) Dependencies() terra.Dependencies {
	return vasb.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (vasb *Resource) LifecycleManagement() *terra.Lifecycle {
	return vasb.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (vasb *Resource) Attributes() vaultAdSecretBackendAttributes {
	return vaultAdSecretBackendAttributes{ref: terra.ReferenceResource(vasb)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (vasb *Resource) ImportState(state io.Reader) error {
	vasb.state = &vaultAdSecretBackendState{}
	if err := json.NewDecoder(state).Decode(vasb.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vasb.Type(), vasb.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (vasb *Resource) State() (*vaultAdSecretBackendState, bool) {
	return vasb.state, vasb.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (vasb *Resource) StateMust() *vaultAdSecretBackendState {
	if vasb.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vasb.Type(), vasb.LocalName()))
	}
	return vasb.state
}

// Args contains the configurations for vault_ad_secret_backend.
type Args struct {
	// AnonymousGroupSearch: bool, optional
	AnonymousGroupSearch terra.BoolValue `hcl:"anonymous_group_search,attr"`
	// Backend: string, optional
	Backend terra.StringValue `hcl:"backend,attr"`
	// Binddn: string, required
	Binddn terra.StringValue `hcl:"binddn,attr" validate:"required"`
	// Bindpass: string, required
	Bindpass terra.StringValue `hcl:"bindpass,attr" validate:"required"`
	// CaseSensitiveNames: bool, optional
	CaseSensitiveNames terra.BoolValue `hcl:"case_sensitive_names,attr"`
	// Certificate: string, optional
	Certificate terra.StringValue `hcl:"certificate,attr"`
	// ClientTlsCert: string, optional
	ClientTlsCert terra.StringValue `hcl:"client_tls_cert,attr"`
	// ClientTlsKey: string, optional
	ClientTlsKey terra.StringValue `hcl:"client_tls_key,attr"`
	// DefaultLeaseTtlSeconds: number, optional
	DefaultLeaseTtlSeconds terra.NumberValue `hcl:"default_lease_ttl_seconds,attr"`
	// DenyNullBind: bool, optional
	DenyNullBind terra.BoolValue `hcl:"deny_null_bind,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisableRemount: bool, optional
	DisableRemount terra.BoolValue `hcl:"disable_remount,attr"`
	// Discoverdn: bool, optional
	Discoverdn terra.BoolValue `hcl:"discoverdn,attr"`
	// Groupattr: string, optional
	Groupattr terra.StringValue `hcl:"groupattr,attr"`
	// Groupdn: string, optional
	Groupdn terra.StringValue `hcl:"groupdn,attr"`
	// Groupfilter: string, optional
	Groupfilter terra.StringValue `hcl:"groupfilter,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InsecureTls: bool, optional
	InsecureTls terra.BoolValue `hcl:"insecure_tls,attr"`
	// LastRotationTolerance: number, optional
	LastRotationTolerance terra.NumberValue `hcl:"last_rotation_tolerance,attr"`
	// Local: bool, optional
	Local terra.BoolValue `hcl:"local,attr"`
	// MaxLeaseTtlSeconds: number, optional
	MaxLeaseTtlSeconds terra.NumberValue `hcl:"max_lease_ttl_seconds,attr"`
	// MaxTtl: number, optional
	MaxTtl terra.NumberValue `hcl:"max_ttl,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// PasswordPolicy: string, optional
	PasswordPolicy terra.StringValue `hcl:"password_policy,attr"`
	// RequestTimeout: number, optional
	RequestTimeout terra.NumberValue `hcl:"request_timeout,attr"`
	// Starttls: bool, optional
	Starttls terra.BoolValue `hcl:"starttls,attr"`
	// TlsMaxVersion: string, optional
	TlsMaxVersion terra.StringValue `hcl:"tls_max_version,attr"`
	// TlsMinVersion: string, optional
	TlsMinVersion terra.StringValue `hcl:"tls_min_version,attr"`
	// Ttl: number, optional
	Ttl terra.NumberValue `hcl:"ttl,attr"`
	// Upndomain: string, optional
	Upndomain terra.StringValue `hcl:"upndomain,attr"`
	// Url: string, optional
	Url terra.StringValue `hcl:"url,attr"`
	// UsePre111GroupCnBehavior: bool, optional
	UsePre111GroupCnBehavior terra.BoolValue `hcl:"use_pre111_group_cn_behavior,attr"`
	// UseTokenGroups: bool, optional
	UseTokenGroups terra.BoolValue `hcl:"use_token_groups,attr"`
	// Userattr: string, optional
	Userattr terra.StringValue `hcl:"userattr,attr"`
	// Userdn: string, optional
	Userdn terra.StringValue `hcl:"userdn,attr"`
}

type vaultAdSecretBackendAttributes struct {
	ref terra.Reference
}

// AnonymousGroupSearch returns a reference to field anonymous_group_search of vault_ad_secret_backend.
func (vasb vaultAdSecretBackendAttributes) AnonymousGroupSearch() terra.BoolValue {
	return terra.ReferenceAsBool(vasb.ref.Append("anonymous_group_search"))
}

// Backend returns a reference to field backend of vault_ad_secret_backend.
func (vasb vaultAdSecretBackendAttributes) Backend() terra.StringValue {
	return terra.ReferenceAsString(vasb.ref.Append("backend"))
}

// Binddn returns a reference to field binddn of vault_ad_secret_backend.
func (vasb vaultAdSecretBackendAttributes) Binddn() terra.StringValue {
	return terra.ReferenceAsString(vasb.ref.Append("binddn"))
}

// Bindpass returns a reference to field bindpass of vault_ad_secret_backend.
func (vasb vaultAdSecretBackendAttributes) Bindpass() terra.StringValue {
	return terra.ReferenceAsString(vasb.ref.Append("bindpass"))
}

// CaseSensitiveNames returns a reference to field case_sensitive_names of vault_ad_secret_backend.
func (vasb vaultAdSecretBackendAttributes) CaseSensitiveNames() terra.BoolValue {
	return terra.ReferenceAsBool(vasb.ref.Append("case_sensitive_names"))
}

// Certificate returns a reference to field certificate of vault_ad_secret_backend.
func (vasb vaultAdSecretBackendAttributes) Certificate() terra.StringValue {
	return terra.ReferenceAsString(vasb.ref.Append("certificate"))
}

// ClientTlsCert returns a reference to field client_tls_cert of vault_ad_secret_backend.
func (vasb vaultAdSecretBackendAttributes) ClientTlsCert() terra.StringValue {
	return terra.ReferenceAsString(vasb.ref.Append("client_tls_cert"))
}

// ClientTlsKey returns a reference to field client_tls_key of vault_ad_secret_backend.
func (vasb vaultAdSecretBackendAttributes) ClientTlsKey() terra.StringValue {
	return terra.ReferenceAsString(vasb.ref.Append("client_tls_key"))
}

// DefaultLeaseTtlSeconds returns a reference to field default_lease_ttl_seconds of vault_ad_secret_backend.
func (vasb vaultAdSecretBackendAttributes) DefaultLeaseTtlSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(vasb.ref.Append("default_lease_ttl_seconds"))
}

// DenyNullBind returns a reference to field deny_null_bind of vault_ad_secret_backend.
func (vasb vaultAdSecretBackendAttributes) DenyNullBind() terra.BoolValue {
	return terra.ReferenceAsBool(vasb.ref.Append("deny_null_bind"))
}

// Description returns a reference to field description of vault_ad_secret_backend.
func (vasb vaultAdSecretBackendAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(vasb.ref.Append("description"))
}

// DisableRemount returns a reference to field disable_remount of vault_ad_secret_backend.
func (vasb vaultAdSecretBackendAttributes) DisableRemount() terra.BoolValue {
	return terra.ReferenceAsBool(vasb.ref.Append("disable_remount"))
}

// Discoverdn returns a reference to field discoverdn of vault_ad_secret_backend.
func (vasb vaultAdSecretBackendAttributes) Discoverdn() terra.BoolValue {
	return terra.ReferenceAsBool(vasb.ref.Append("discoverdn"))
}

// Groupattr returns a reference to field groupattr of vault_ad_secret_backend.
func (vasb vaultAdSecretBackendAttributes) Groupattr() terra.StringValue {
	return terra.ReferenceAsString(vasb.ref.Append("groupattr"))
}

// Groupdn returns a reference to field groupdn of vault_ad_secret_backend.
func (vasb vaultAdSecretBackendAttributes) Groupdn() terra.StringValue {
	return terra.ReferenceAsString(vasb.ref.Append("groupdn"))
}

// Groupfilter returns a reference to field groupfilter of vault_ad_secret_backend.
func (vasb vaultAdSecretBackendAttributes) Groupfilter() terra.StringValue {
	return terra.ReferenceAsString(vasb.ref.Append("groupfilter"))
}

// Id returns a reference to field id of vault_ad_secret_backend.
func (vasb vaultAdSecretBackendAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vasb.ref.Append("id"))
}

// InsecureTls returns a reference to field insecure_tls of vault_ad_secret_backend.
func (vasb vaultAdSecretBackendAttributes) InsecureTls() terra.BoolValue {
	return terra.ReferenceAsBool(vasb.ref.Append("insecure_tls"))
}

// LastRotationTolerance returns a reference to field last_rotation_tolerance of vault_ad_secret_backend.
func (vasb vaultAdSecretBackendAttributes) LastRotationTolerance() terra.NumberValue {
	return terra.ReferenceAsNumber(vasb.ref.Append("last_rotation_tolerance"))
}

// Local returns a reference to field local of vault_ad_secret_backend.
func (vasb vaultAdSecretBackendAttributes) Local() terra.BoolValue {
	return terra.ReferenceAsBool(vasb.ref.Append("local"))
}

// MaxLeaseTtlSeconds returns a reference to field max_lease_ttl_seconds of vault_ad_secret_backend.
func (vasb vaultAdSecretBackendAttributes) MaxLeaseTtlSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(vasb.ref.Append("max_lease_ttl_seconds"))
}

// MaxTtl returns a reference to field max_ttl of vault_ad_secret_backend.
func (vasb vaultAdSecretBackendAttributes) MaxTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(vasb.ref.Append("max_ttl"))
}

// Namespace returns a reference to field namespace of vault_ad_secret_backend.
func (vasb vaultAdSecretBackendAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(vasb.ref.Append("namespace"))
}

// PasswordPolicy returns a reference to field password_policy of vault_ad_secret_backend.
func (vasb vaultAdSecretBackendAttributes) PasswordPolicy() terra.StringValue {
	return terra.ReferenceAsString(vasb.ref.Append("password_policy"))
}

// RequestTimeout returns a reference to field request_timeout of vault_ad_secret_backend.
func (vasb vaultAdSecretBackendAttributes) RequestTimeout() terra.NumberValue {
	return terra.ReferenceAsNumber(vasb.ref.Append("request_timeout"))
}

// Starttls returns a reference to field starttls of vault_ad_secret_backend.
func (vasb vaultAdSecretBackendAttributes) Starttls() terra.BoolValue {
	return terra.ReferenceAsBool(vasb.ref.Append("starttls"))
}

// TlsMaxVersion returns a reference to field tls_max_version of vault_ad_secret_backend.
func (vasb vaultAdSecretBackendAttributes) TlsMaxVersion() terra.StringValue {
	return terra.ReferenceAsString(vasb.ref.Append("tls_max_version"))
}

// TlsMinVersion returns a reference to field tls_min_version of vault_ad_secret_backend.
func (vasb vaultAdSecretBackendAttributes) TlsMinVersion() terra.StringValue {
	return terra.ReferenceAsString(vasb.ref.Append("tls_min_version"))
}

// Ttl returns a reference to field ttl of vault_ad_secret_backend.
func (vasb vaultAdSecretBackendAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(vasb.ref.Append("ttl"))
}

// Upndomain returns a reference to field upndomain of vault_ad_secret_backend.
func (vasb vaultAdSecretBackendAttributes) Upndomain() terra.StringValue {
	return terra.ReferenceAsString(vasb.ref.Append("upndomain"))
}

// Url returns a reference to field url of vault_ad_secret_backend.
func (vasb vaultAdSecretBackendAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(vasb.ref.Append("url"))
}

// UsePre111GroupCnBehavior returns a reference to field use_pre111_group_cn_behavior of vault_ad_secret_backend.
func (vasb vaultAdSecretBackendAttributes) UsePre111GroupCnBehavior() terra.BoolValue {
	return terra.ReferenceAsBool(vasb.ref.Append("use_pre111_group_cn_behavior"))
}

// UseTokenGroups returns a reference to field use_token_groups of vault_ad_secret_backend.
func (vasb vaultAdSecretBackendAttributes) UseTokenGroups() terra.BoolValue {
	return terra.ReferenceAsBool(vasb.ref.Append("use_token_groups"))
}

// Userattr returns a reference to field userattr of vault_ad_secret_backend.
func (vasb vaultAdSecretBackendAttributes) Userattr() terra.StringValue {
	return terra.ReferenceAsString(vasb.ref.Append("userattr"))
}

// Userdn returns a reference to field userdn of vault_ad_secret_backend.
func (vasb vaultAdSecretBackendAttributes) Userdn() terra.StringValue {
	return terra.ReferenceAsString(vasb.ref.Append("userdn"))
}

type vaultAdSecretBackendState struct {
	AnonymousGroupSearch     bool    `json:"anonymous_group_search"`
	Backend                  string  `json:"backend"`
	Binddn                   string  `json:"binddn"`
	Bindpass                 string  `json:"bindpass"`
	CaseSensitiveNames       bool    `json:"case_sensitive_names"`
	Certificate              string  `json:"certificate"`
	ClientTlsCert            string  `json:"client_tls_cert"`
	ClientTlsKey             string  `json:"client_tls_key"`
	DefaultLeaseTtlSeconds   float64 `json:"default_lease_ttl_seconds"`
	DenyNullBind             bool    `json:"deny_null_bind"`
	Description              string  `json:"description"`
	DisableRemount           bool    `json:"disable_remount"`
	Discoverdn               bool    `json:"discoverdn"`
	Groupattr                string  `json:"groupattr"`
	Groupdn                  string  `json:"groupdn"`
	Groupfilter              string  `json:"groupfilter"`
	Id                       string  `json:"id"`
	InsecureTls              bool    `json:"insecure_tls"`
	LastRotationTolerance    float64 `json:"last_rotation_tolerance"`
	Local                    bool    `json:"local"`
	MaxLeaseTtlSeconds       float64 `json:"max_lease_ttl_seconds"`
	MaxTtl                   float64 `json:"max_ttl"`
	Namespace                string  `json:"namespace"`
	PasswordPolicy           string  `json:"password_policy"`
	RequestTimeout           float64 `json:"request_timeout"`
	Starttls                 bool    `json:"starttls"`
	TlsMaxVersion            string  `json:"tls_max_version"`
	TlsMinVersion            string  `json:"tls_min_version"`
	Ttl                      float64 `json:"ttl"`
	Upndomain                string  `json:"upndomain"`
	Url                      string  `json:"url"`
	UsePre111GroupCnBehavior bool    `json:"use_pre111_group_cn_behavior"`
	UseTokenGroups           bool    `json:"use_token_groups"`
	Userattr                 string  `json:"userattr"`
	Userdn                   string  `json:"userdn"`
}
