// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	sshsecretbackendrole "github.com/golingon/terraproviders/vault/3.18.0/sshsecretbackendrole"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSshSecretBackendRole creates a new instance of [SshSecretBackendRole].
func NewSshSecretBackendRole(name string, args SshSecretBackendRoleArgs) *SshSecretBackendRole {
	return &SshSecretBackendRole{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SshSecretBackendRole)(nil)

// SshSecretBackendRole represents the Terraform resource vault_ssh_secret_backend_role.
type SshSecretBackendRole struct {
	Name      string
	Args      SshSecretBackendRoleArgs
	state     *sshSecretBackendRoleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SshSecretBackendRole].
func (ssbr *SshSecretBackendRole) Type() string {
	return "vault_ssh_secret_backend_role"
}

// LocalName returns the local name for [SshSecretBackendRole].
func (ssbr *SshSecretBackendRole) LocalName() string {
	return ssbr.Name
}

// Configuration returns the configuration (args) for [SshSecretBackendRole].
func (ssbr *SshSecretBackendRole) Configuration() interface{} {
	return ssbr.Args
}

// DependOn is used for other resources to depend on [SshSecretBackendRole].
func (ssbr *SshSecretBackendRole) DependOn() terra.Reference {
	return terra.ReferenceResource(ssbr)
}

// Dependencies returns the list of resources [SshSecretBackendRole] depends_on.
func (ssbr *SshSecretBackendRole) Dependencies() terra.Dependencies {
	return ssbr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SshSecretBackendRole].
func (ssbr *SshSecretBackendRole) LifecycleManagement() *terra.Lifecycle {
	return ssbr.Lifecycle
}

// Attributes returns the attributes for [SshSecretBackendRole].
func (ssbr *SshSecretBackendRole) Attributes() sshSecretBackendRoleAttributes {
	return sshSecretBackendRoleAttributes{ref: terra.ReferenceResource(ssbr)}
}

// ImportState imports the given attribute values into [SshSecretBackendRole]'s state.
func (ssbr *SshSecretBackendRole) ImportState(av io.Reader) error {
	ssbr.state = &sshSecretBackendRoleState{}
	if err := json.NewDecoder(av).Decode(ssbr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ssbr.Type(), ssbr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SshSecretBackendRole] has state.
func (ssbr *SshSecretBackendRole) State() (*sshSecretBackendRoleState, bool) {
	return ssbr.state, ssbr.state != nil
}

// StateMust returns the state for [SshSecretBackendRole]. Panics if the state is nil.
func (ssbr *SshSecretBackendRole) StateMust() *sshSecretBackendRoleState {
	if ssbr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ssbr.Type(), ssbr.LocalName()))
	}
	return ssbr.state
}

// SshSecretBackendRoleArgs contains the configurations for vault_ssh_secret_backend_role.
type SshSecretBackendRoleArgs struct {
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
	// AllowedUserKeyLengths: map of number, optional
	AllowedUserKeyLengths terra.MapValue[terra.NumberValue] `hcl:"allowed_user_key_lengths,attr"`
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
	// Ttl: string, optional
	Ttl terra.StringValue `hcl:"ttl,attr"`
	// AllowedUserKeyConfig: min=0
	AllowedUserKeyConfig []sshsecretbackendrole.AllowedUserKeyConfig `hcl:"allowed_user_key_config,block" validate:"min=0"`
}
type sshSecretBackendRoleAttributes struct {
	ref terra.Reference
}

// AlgorithmSigner returns a reference to field algorithm_signer of vault_ssh_secret_backend_role.
func (ssbr sshSecretBackendRoleAttributes) AlgorithmSigner() terra.StringValue {
	return terra.ReferenceAsString(ssbr.ref.Append("algorithm_signer"))
}

// AllowBareDomains returns a reference to field allow_bare_domains of vault_ssh_secret_backend_role.
func (ssbr sshSecretBackendRoleAttributes) AllowBareDomains() terra.BoolValue {
	return terra.ReferenceAsBool(ssbr.ref.Append("allow_bare_domains"))
}

// AllowHostCertificates returns a reference to field allow_host_certificates of vault_ssh_secret_backend_role.
func (ssbr sshSecretBackendRoleAttributes) AllowHostCertificates() terra.BoolValue {
	return terra.ReferenceAsBool(ssbr.ref.Append("allow_host_certificates"))
}

// AllowSubdomains returns a reference to field allow_subdomains of vault_ssh_secret_backend_role.
func (ssbr sshSecretBackendRoleAttributes) AllowSubdomains() terra.BoolValue {
	return terra.ReferenceAsBool(ssbr.ref.Append("allow_subdomains"))
}

// AllowUserCertificates returns a reference to field allow_user_certificates of vault_ssh_secret_backend_role.
func (ssbr sshSecretBackendRoleAttributes) AllowUserCertificates() terra.BoolValue {
	return terra.ReferenceAsBool(ssbr.ref.Append("allow_user_certificates"))
}

// AllowUserKeyIds returns a reference to field allow_user_key_ids of vault_ssh_secret_backend_role.
func (ssbr sshSecretBackendRoleAttributes) AllowUserKeyIds() terra.BoolValue {
	return terra.ReferenceAsBool(ssbr.ref.Append("allow_user_key_ids"))
}

// AllowedCriticalOptions returns a reference to field allowed_critical_options of vault_ssh_secret_backend_role.
func (ssbr sshSecretBackendRoleAttributes) AllowedCriticalOptions() terra.StringValue {
	return terra.ReferenceAsString(ssbr.ref.Append("allowed_critical_options"))
}

// AllowedDomains returns a reference to field allowed_domains of vault_ssh_secret_backend_role.
func (ssbr sshSecretBackendRoleAttributes) AllowedDomains() terra.StringValue {
	return terra.ReferenceAsString(ssbr.ref.Append("allowed_domains"))
}

// AllowedDomainsTemplate returns a reference to field allowed_domains_template of vault_ssh_secret_backend_role.
func (ssbr sshSecretBackendRoleAttributes) AllowedDomainsTemplate() terra.BoolValue {
	return terra.ReferenceAsBool(ssbr.ref.Append("allowed_domains_template"))
}

// AllowedExtensions returns a reference to field allowed_extensions of vault_ssh_secret_backend_role.
func (ssbr sshSecretBackendRoleAttributes) AllowedExtensions() terra.StringValue {
	return terra.ReferenceAsString(ssbr.ref.Append("allowed_extensions"))
}

// AllowedUserKeyLengths returns a reference to field allowed_user_key_lengths of vault_ssh_secret_backend_role.
func (ssbr sshSecretBackendRoleAttributes) AllowedUserKeyLengths() terra.MapValue[terra.NumberValue] {
	return terra.ReferenceAsMap[terra.NumberValue](ssbr.ref.Append("allowed_user_key_lengths"))
}

// AllowedUsers returns a reference to field allowed_users of vault_ssh_secret_backend_role.
func (ssbr sshSecretBackendRoleAttributes) AllowedUsers() terra.StringValue {
	return terra.ReferenceAsString(ssbr.ref.Append("allowed_users"))
}

// AllowedUsersTemplate returns a reference to field allowed_users_template of vault_ssh_secret_backend_role.
func (ssbr sshSecretBackendRoleAttributes) AllowedUsersTemplate() terra.BoolValue {
	return terra.ReferenceAsBool(ssbr.ref.Append("allowed_users_template"))
}

// Backend returns a reference to field backend of vault_ssh_secret_backend_role.
func (ssbr sshSecretBackendRoleAttributes) Backend() terra.StringValue {
	return terra.ReferenceAsString(ssbr.ref.Append("backend"))
}

// CidrList returns a reference to field cidr_list of vault_ssh_secret_backend_role.
func (ssbr sshSecretBackendRoleAttributes) CidrList() terra.StringValue {
	return terra.ReferenceAsString(ssbr.ref.Append("cidr_list"))
}

// DefaultCriticalOptions returns a reference to field default_critical_options of vault_ssh_secret_backend_role.
func (ssbr sshSecretBackendRoleAttributes) DefaultCriticalOptions() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ssbr.ref.Append("default_critical_options"))
}

// DefaultExtensions returns a reference to field default_extensions of vault_ssh_secret_backend_role.
func (ssbr sshSecretBackendRoleAttributes) DefaultExtensions() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ssbr.ref.Append("default_extensions"))
}

// DefaultUser returns a reference to field default_user of vault_ssh_secret_backend_role.
func (ssbr sshSecretBackendRoleAttributes) DefaultUser() terra.StringValue {
	return terra.ReferenceAsString(ssbr.ref.Append("default_user"))
}

// DefaultUserTemplate returns a reference to field default_user_template of vault_ssh_secret_backend_role.
func (ssbr sshSecretBackendRoleAttributes) DefaultUserTemplate() terra.BoolValue {
	return terra.ReferenceAsBool(ssbr.ref.Append("default_user_template"))
}

// Id returns a reference to field id of vault_ssh_secret_backend_role.
func (ssbr sshSecretBackendRoleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ssbr.ref.Append("id"))
}

// KeyIdFormat returns a reference to field key_id_format of vault_ssh_secret_backend_role.
func (ssbr sshSecretBackendRoleAttributes) KeyIdFormat() terra.StringValue {
	return terra.ReferenceAsString(ssbr.ref.Append("key_id_format"))
}

// KeyType returns a reference to field key_type of vault_ssh_secret_backend_role.
func (ssbr sshSecretBackendRoleAttributes) KeyType() terra.StringValue {
	return terra.ReferenceAsString(ssbr.ref.Append("key_type"))
}

// MaxTtl returns a reference to field max_ttl of vault_ssh_secret_backend_role.
func (ssbr sshSecretBackendRoleAttributes) MaxTtl() terra.StringValue {
	return terra.ReferenceAsString(ssbr.ref.Append("max_ttl"))
}

// Name returns a reference to field name of vault_ssh_secret_backend_role.
func (ssbr sshSecretBackendRoleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ssbr.ref.Append("name"))
}

// Namespace returns a reference to field namespace of vault_ssh_secret_backend_role.
func (ssbr sshSecretBackendRoleAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(ssbr.ref.Append("namespace"))
}

// Ttl returns a reference to field ttl of vault_ssh_secret_backend_role.
func (ssbr sshSecretBackendRoleAttributes) Ttl() terra.StringValue {
	return terra.ReferenceAsString(ssbr.ref.Append("ttl"))
}

func (ssbr sshSecretBackendRoleAttributes) AllowedUserKeyConfig() terra.SetValue[sshsecretbackendrole.AllowedUserKeyConfigAttributes] {
	return terra.ReferenceAsSet[sshsecretbackendrole.AllowedUserKeyConfigAttributes](ssbr.ref.Append("allowed_user_key_config"))
}

type sshSecretBackendRoleState struct {
	AlgorithmSigner        string                                           `json:"algorithm_signer"`
	AllowBareDomains       bool                                             `json:"allow_bare_domains"`
	AllowHostCertificates  bool                                             `json:"allow_host_certificates"`
	AllowSubdomains        bool                                             `json:"allow_subdomains"`
	AllowUserCertificates  bool                                             `json:"allow_user_certificates"`
	AllowUserKeyIds        bool                                             `json:"allow_user_key_ids"`
	AllowedCriticalOptions string                                           `json:"allowed_critical_options"`
	AllowedDomains         string                                           `json:"allowed_domains"`
	AllowedDomainsTemplate bool                                             `json:"allowed_domains_template"`
	AllowedExtensions      string                                           `json:"allowed_extensions"`
	AllowedUserKeyLengths  map[string]float64                               `json:"allowed_user_key_lengths"`
	AllowedUsers           string                                           `json:"allowed_users"`
	AllowedUsersTemplate   bool                                             `json:"allowed_users_template"`
	Backend                string                                           `json:"backend"`
	CidrList               string                                           `json:"cidr_list"`
	DefaultCriticalOptions map[string]string                                `json:"default_critical_options"`
	DefaultExtensions      map[string]string                                `json:"default_extensions"`
	DefaultUser            string                                           `json:"default_user"`
	DefaultUserTemplate    bool                                             `json:"default_user_template"`
	Id                     string                                           `json:"id"`
	KeyIdFormat            string                                           `json:"key_id_format"`
	KeyType                string                                           `json:"key_type"`
	MaxTtl                 string                                           `json:"max_ttl"`
	Name                   string                                           `json:"name"`
	Namespace              string                                           `json:"namespace"`
	Ttl                    string                                           `json:"ttl"`
	AllowedUserKeyConfig   []sshsecretbackendrole.AllowedUserKeyConfigState `json:"allowed_user_key_config"`
}
