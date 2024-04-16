// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vault_ssh_secret_backend_role

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

// Resource represents the Terraform resource vault_ssh_secret_backend_role.
type Resource struct {
	Name      string
	Args      Args
	state     *vaultSshSecretBackendRoleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (vssbr *Resource) Type() string {
	return "vault_ssh_secret_backend_role"
}

// LocalName returns the local name for [Resource].
func (vssbr *Resource) LocalName() string {
	return vssbr.Name
}

// Configuration returns the configuration (args) for [Resource].
func (vssbr *Resource) Configuration() interface{} {
	return vssbr.Args
}

// DependOn is used for other resources to depend on [Resource].
func (vssbr *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(vssbr)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (vssbr *Resource) Dependencies() terra.Dependencies {
	return vssbr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (vssbr *Resource) LifecycleManagement() *terra.Lifecycle {
	return vssbr.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (vssbr *Resource) Attributes() vaultSshSecretBackendRoleAttributes {
	return vaultSshSecretBackendRoleAttributes{ref: terra.ReferenceResource(vssbr)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (vssbr *Resource) ImportState(state io.Reader) error {
	vssbr.state = &vaultSshSecretBackendRoleState{}
	if err := json.NewDecoder(state).Decode(vssbr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vssbr.Type(), vssbr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (vssbr *Resource) State() (*vaultSshSecretBackendRoleState, bool) {
	return vssbr.state, vssbr.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (vssbr *Resource) StateMust() *vaultSshSecretBackendRoleState {
	if vssbr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vssbr.Type(), vssbr.LocalName()))
	}
	return vssbr.state
}

// Args contains the configurations for vault_ssh_secret_backend_role.
type Args struct {
	// AlgorithmSigner: string, optional
	AlgorithmSigner terra.StringValue `hcl:"algorithm_signer,attr"`
	// AllowBareDomains: bool, optional
	AllowBareDomains terra.BoolValue `hcl:"allow_bare_domains,attr"`
	// AllowHostCertificates: bool, optional
	AllowHostCertificates terra.BoolValue `hcl:"allow_host_certificates,attr"`
	// AllowSubdomains: bool, optional
	AllowSubdomains terra.BoolValue `hcl:"allow_subdomains,attr"`
	// AllowUserCertificates: bool, optional
	AllowUserCertificates terra.BoolValue `hcl:"allow_user_certificates,attr"`
	// AllowUserKeyIds: bool, optional
	AllowUserKeyIds terra.BoolValue `hcl:"allow_user_key_ids,attr"`
	// AllowedCriticalOptions: string, optional
	AllowedCriticalOptions terra.StringValue `hcl:"allowed_critical_options,attr"`
	// AllowedDomains: string, optional
	AllowedDomains terra.StringValue `hcl:"allowed_domains,attr"`
	// AllowedDomainsTemplate: bool, optional
	AllowedDomainsTemplate terra.BoolValue `hcl:"allowed_domains_template,attr"`
	// AllowedExtensions: string, optional
	AllowedExtensions terra.StringValue `hcl:"allowed_extensions,attr"`
	// AllowedUsers: string, optional
	AllowedUsers terra.StringValue `hcl:"allowed_users,attr"`
	// AllowedUsersTemplate: bool, optional
	AllowedUsersTemplate terra.BoolValue `hcl:"allowed_users_template,attr"`
	// Backend: string, required
	Backend terra.StringValue `hcl:"backend,attr" validate:"required"`
	// CidrList: string, optional
	CidrList terra.StringValue `hcl:"cidr_list,attr"`
	// DefaultCriticalOptions: map of string, optional
	DefaultCriticalOptions terra.MapValue[terra.StringValue] `hcl:"default_critical_options,attr"`
	// DefaultExtensions: map of string, optional
	DefaultExtensions terra.MapValue[terra.StringValue] `hcl:"default_extensions,attr"`
	// DefaultUser: string, optional
	DefaultUser terra.StringValue `hcl:"default_user,attr"`
	// DefaultUserTemplate: bool, optional
	DefaultUserTemplate terra.BoolValue `hcl:"default_user_template,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyIdFormat: string, optional
	KeyIdFormat terra.StringValue `hcl:"key_id_format,attr"`
	// KeyType: string, required
	KeyType terra.StringValue `hcl:"key_type,attr" validate:"required"`
	// MaxTtl: string, optional
	MaxTtl terra.StringValue `hcl:"max_ttl,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// NotBeforeDuration: string, optional
	NotBeforeDuration terra.StringValue `hcl:"not_before_duration,attr"`
	// Ttl: string, optional
	Ttl terra.StringValue `hcl:"ttl,attr"`
	// AllowedUserKeyConfig: min=0
	AllowedUserKeyConfig []AllowedUserKeyConfig `hcl:"allowed_user_key_config,block" validate:"min=0"`
}

type vaultSshSecretBackendRoleAttributes struct {
	ref terra.Reference
}

// AlgorithmSigner returns a reference to field algorithm_signer of vault_ssh_secret_backend_role.
func (vssbr vaultSshSecretBackendRoleAttributes) AlgorithmSigner() terra.StringValue {
	return terra.ReferenceAsString(vssbr.ref.Append("algorithm_signer"))
}

// AllowBareDomains returns a reference to field allow_bare_domains of vault_ssh_secret_backend_role.
func (vssbr vaultSshSecretBackendRoleAttributes) AllowBareDomains() terra.BoolValue {
	return terra.ReferenceAsBool(vssbr.ref.Append("allow_bare_domains"))
}

// AllowHostCertificates returns a reference to field allow_host_certificates of vault_ssh_secret_backend_role.
func (vssbr vaultSshSecretBackendRoleAttributes) AllowHostCertificates() terra.BoolValue {
	return terra.ReferenceAsBool(vssbr.ref.Append("allow_host_certificates"))
}

// AllowSubdomains returns a reference to field allow_subdomains of vault_ssh_secret_backend_role.
func (vssbr vaultSshSecretBackendRoleAttributes) AllowSubdomains() terra.BoolValue {
	return terra.ReferenceAsBool(vssbr.ref.Append("allow_subdomains"))
}

// AllowUserCertificates returns a reference to field allow_user_certificates of vault_ssh_secret_backend_role.
func (vssbr vaultSshSecretBackendRoleAttributes) AllowUserCertificates() terra.BoolValue {
	return terra.ReferenceAsBool(vssbr.ref.Append("allow_user_certificates"))
}

// AllowUserKeyIds returns a reference to field allow_user_key_ids of vault_ssh_secret_backend_role.
func (vssbr vaultSshSecretBackendRoleAttributes) AllowUserKeyIds() terra.BoolValue {
	return terra.ReferenceAsBool(vssbr.ref.Append("allow_user_key_ids"))
}

// AllowedCriticalOptions returns a reference to field allowed_critical_options of vault_ssh_secret_backend_role.
func (vssbr vaultSshSecretBackendRoleAttributes) AllowedCriticalOptions() terra.StringValue {
	return terra.ReferenceAsString(vssbr.ref.Append("allowed_critical_options"))
}

// AllowedDomains returns a reference to field allowed_domains of vault_ssh_secret_backend_role.
func (vssbr vaultSshSecretBackendRoleAttributes) AllowedDomains() terra.StringValue {
	return terra.ReferenceAsString(vssbr.ref.Append("allowed_domains"))
}

// AllowedDomainsTemplate returns a reference to field allowed_domains_template of vault_ssh_secret_backend_role.
func (vssbr vaultSshSecretBackendRoleAttributes) AllowedDomainsTemplate() terra.BoolValue {
	return terra.ReferenceAsBool(vssbr.ref.Append("allowed_domains_template"))
}

// AllowedExtensions returns a reference to field allowed_extensions of vault_ssh_secret_backend_role.
func (vssbr vaultSshSecretBackendRoleAttributes) AllowedExtensions() terra.StringValue {
	return terra.ReferenceAsString(vssbr.ref.Append("allowed_extensions"))
}

// AllowedUsers returns a reference to field allowed_users of vault_ssh_secret_backend_role.
func (vssbr vaultSshSecretBackendRoleAttributes) AllowedUsers() terra.StringValue {
	return terra.ReferenceAsString(vssbr.ref.Append("allowed_users"))
}

// AllowedUsersTemplate returns a reference to field allowed_users_template of vault_ssh_secret_backend_role.
func (vssbr vaultSshSecretBackendRoleAttributes) AllowedUsersTemplate() terra.BoolValue {
	return terra.ReferenceAsBool(vssbr.ref.Append("allowed_users_template"))
}

// Backend returns a reference to field backend of vault_ssh_secret_backend_role.
func (vssbr vaultSshSecretBackendRoleAttributes) Backend() terra.StringValue {
	return terra.ReferenceAsString(vssbr.ref.Append("backend"))
}

// CidrList returns a reference to field cidr_list of vault_ssh_secret_backend_role.
func (vssbr vaultSshSecretBackendRoleAttributes) CidrList() terra.StringValue {
	return terra.ReferenceAsString(vssbr.ref.Append("cidr_list"))
}

// DefaultCriticalOptions returns a reference to field default_critical_options of vault_ssh_secret_backend_role.
func (vssbr vaultSshSecretBackendRoleAttributes) DefaultCriticalOptions() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vssbr.ref.Append("default_critical_options"))
}

// DefaultExtensions returns a reference to field default_extensions of vault_ssh_secret_backend_role.
func (vssbr vaultSshSecretBackendRoleAttributes) DefaultExtensions() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vssbr.ref.Append("default_extensions"))
}

// DefaultUser returns a reference to field default_user of vault_ssh_secret_backend_role.
func (vssbr vaultSshSecretBackendRoleAttributes) DefaultUser() terra.StringValue {
	return terra.ReferenceAsString(vssbr.ref.Append("default_user"))
}

// DefaultUserTemplate returns a reference to field default_user_template of vault_ssh_secret_backend_role.
func (vssbr vaultSshSecretBackendRoleAttributes) DefaultUserTemplate() terra.BoolValue {
	return terra.ReferenceAsBool(vssbr.ref.Append("default_user_template"))
}

// Id returns a reference to field id of vault_ssh_secret_backend_role.
func (vssbr vaultSshSecretBackendRoleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vssbr.ref.Append("id"))
}

// KeyIdFormat returns a reference to field key_id_format of vault_ssh_secret_backend_role.
func (vssbr vaultSshSecretBackendRoleAttributes) KeyIdFormat() terra.StringValue {
	return terra.ReferenceAsString(vssbr.ref.Append("key_id_format"))
}

// KeyType returns a reference to field key_type of vault_ssh_secret_backend_role.
func (vssbr vaultSshSecretBackendRoleAttributes) KeyType() terra.StringValue {
	return terra.ReferenceAsString(vssbr.ref.Append("key_type"))
}

// MaxTtl returns a reference to field max_ttl of vault_ssh_secret_backend_role.
func (vssbr vaultSshSecretBackendRoleAttributes) MaxTtl() terra.StringValue {
	return terra.ReferenceAsString(vssbr.ref.Append("max_ttl"))
}

// Name returns a reference to field name of vault_ssh_secret_backend_role.
func (vssbr vaultSshSecretBackendRoleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vssbr.ref.Append("name"))
}

// Namespace returns a reference to field namespace of vault_ssh_secret_backend_role.
func (vssbr vaultSshSecretBackendRoleAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(vssbr.ref.Append("namespace"))
}

// NotBeforeDuration returns a reference to field not_before_duration of vault_ssh_secret_backend_role.
func (vssbr vaultSshSecretBackendRoleAttributes) NotBeforeDuration() terra.StringValue {
	return terra.ReferenceAsString(vssbr.ref.Append("not_before_duration"))
}

// Ttl returns a reference to field ttl of vault_ssh_secret_backend_role.
func (vssbr vaultSshSecretBackendRoleAttributes) Ttl() terra.StringValue {
	return terra.ReferenceAsString(vssbr.ref.Append("ttl"))
}

func (vssbr vaultSshSecretBackendRoleAttributes) AllowedUserKeyConfig() terra.SetValue[AllowedUserKeyConfigAttributes] {
	return terra.ReferenceAsSet[AllowedUserKeyConfigAttributes](vssbr.ref.Append("allowed_user_key_config"))
}

type vaultSshSecretBackendRoleState struct {
	AlgorithmSigner        string                      `json:"algorithm_signer"`
	AllowBareDomains       bool                        `json:"allow_bare_domains"`
	AllowHostCertificates  bool                        `json:"allow_host_certificates"`
	AllowSubdomains        bool                        `json:"allow_subdomains"`
	AllowUserCertificates  bool                        `json:"allow_user_certificates"`
	AllowUserKeyIds        bool                        `json:"allow_user_key_ids"`
	AllowedCriticalOptions string                      `json:"allowed_critical_options"`
	AllowedDomains         string                      `json:"allowed_domains"`
	AllowedDomainsTemplate bool                        `json:"allowed_domains_template"`
	AllowedExtensions      string                      `json:"allowed_extensions"`
	AllowedUsers           string                      `json:"allowed_users"`
	AllowedUsersTemplate   bool                        `json:"allowed_users_template"`
	Backend                string                      `json:"backend"`
	CidrList               string                      `json:"cidr_list"`
	DefaultCriticalOptions map[string]string           `json:"default_critical_options"`
	DefaultExtensions      map[string]string           `json:"default_extensions"`
	DefaultUser            string                      `json:"default_user"`
	DefaultUserTemplate    bool                        `json:"default_user_template"`
	Id                     string                      `json:"id"`
	KeyIdFormat            string                      `json:"key_id_format"`
	KeyType                string                      `json:"key_type"`
	MaxTtl                 string                      `json:"max_ttl"`
	Name                   string                      `json:"name"`
	Namespace              string                      `json:"namespace"`
	NotBeforeDuration      string                      `json:"not_before_duration"`
	Ttl                    string                      `json:"ttl"`
	AllowedUserKeyConfig   []AllowedUserKeyConfigState `json:"allowed_user_key_config"`
}
