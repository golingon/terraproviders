// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewLdapSecretBackendLibrarySet creates a new instance of [LdapSecretBackendLibrarySet].
func NewLdapSecretBackendLibrarySet(name string, args LdapSecretBackendLibrarySetArgs) *LdapSecretBackendLibrarySet {
	return &LdapSecretBackendLibrarySet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LdapSecretBackendLibrarySet)(nil)

// LdapSecretBackendLibrarySet represents the Terraform resource vault_ldap_secret_backend_library_set.
type LdapSecretBackendLibrarySet struct {
	Name      string
	Args      LdapSecretBackendLibrarySetArgs
	state     *ldapSecretBackendLibrarySetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LdapSecretBackendLibrarySet].
func (lsbls *LdapSecretBackendLibrarySet) Type() string {
	return "vault_ldap_secret_backend_library_set"
}

// LocalName returns the local name for [LdapSecretBackendLibrarySet].
func (lsbls *LdapSecretBackendLibrarySet) LocalName() string {
	return lsbls.Name
}

// Configuration returns the configuration (args) for [LdapSecretBackendLibrarySet].
func (lsbls *LdapSecretBackendLibrarySet) Configuration() interface{} {
	return lsbls.Args
}

// DependOn is used for other resources to depend on [LdapSecretBackendLibrarySet].
func (lsbls *LdapSecretBackendLibrarySet) DependOn() terra.Reference {
	return terra.ReferenceResource(lsbls)
}

// Dependencies returns the list of resources [LdapSecretBackendLibrarySet] depends_on.
func (lsbls *LdapSecretBackendLibrarySet) Dependencies() terra.Dependencies {
	return lsbls.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LdapSecretBackendLibrarySet].
func (lsbls *LdapSecretBackendLibrarySet) LifecycleManagement() *terra.Lifecycle {
	return lsbls.Lifecycle
}

// Attributes returns the attributes for [LdapSecretBackendLibrarySet].
func (lsbls *LdapSecretBackendLibrarySet) Attributes() ldapSecretBackendLibrarySetAttributes {
	return ldapSecretBackendLibrarySetAttributes{ref: terra.ReferenceResource(lsbls)}
}

// ImportState imports the given attribute values into [LdapSecretBackendLibrarySet]'s state.
func (lsbls *LdapSecretBackendLibrarySet) ImportState(av io.Reader) error {
	lsbls.state = &ldapSecretBackendLibrarySetState{}
	if err := json.NewDecoder(av).Decode(lsbls.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lsbls.Type(), lsbls.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LdapSecretBackendLibrarySet] has state.
func (lsbls *LdapSecretBackendLibrarySet) State() (*ldapSecretBackendLibrarySetState, bool) {
	return lsbls.state, lsbls.state != nil
}

// StateMust returns the state for [LdapSecretBackendLibrarySet]. Panics if the state is nil.
func (lsbls *LdapSecretBackendLibrarySet) StateMust() *ldapSecretBackendLibrarySetState {
	if lsbls.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lsbls.Type(), lsbls.LocalName()))
	}
	return lsbls.state
}

// LdapSecretBackendLibrarySetArgs contains the configurations for vault_ldap_secret_backend_library_set.
type LdapSecretBackendLibrarySetArgs struct {
	// DisableCheckInEnforcement: bool, optional
	DisableCheckInEnforcement terra.BoolValue `hcl:"disable_check_in_enforcement,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MaxTtl: number, optional
	MaxTtl terra.NumberValue `hcl:"max_ttl,attr"`
	// Mount: string, optional
	Mount terra.StringValue `hcl:"mount,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// ServiceAccountNames: list of string, required
	ServiceAccountNames terra.ListValue[terra.StringValue] `hcl:"service_account_names,attr" validate:"required"`
	// Ttl: number, optional
	Ttl terra.NumberValue `hcl:"ttl,attr"`
}
type ldapSecretBackendLibrarySetAttributes struct {
	ref terra.Reference
}

// DisableCheckInEnforcement returns a reference to field disable_check_in_enforcement of vault_ldap_secret_backend_library_set.
func (lsbls ldapSecretBackendLibrarySetAttributes) DisableCheckInEnforcement() terra.BoolValue {
	return terra.ReferenceAsBool(lsbls.ref.Append("disable_check_in_enforcement"))
}

// Id returns a reference to field id of vault_ldap_secret_backend_library_set.
func (lsbls ldapSecretBackendLibrarySetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lsbls.ref.Append("id"))
}

// MaxTtl returns a reference to field max_ttl of vault_ldap_secret_backend_library_set.
func (lsbls ldapSecretBackendLibrarySetAttributes) MaxTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(lsbls.ref.Append("max_ttl"))
}

// Mount returns a reference to field mount of vault_ldap_secret_backend_library_set.
func (lsbls ldapSecretBackendLibrarySetAttributes) Mount() terra.StringValue {
	return terra.ReferenceAsString(lsbls.ref.Append("mount"))
}

// Name returns a reference to field name of vault_ldap_secret_backend_library_set.
func (lsbls ldapSecretBackendLibrarySetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lsbls.ref.Append("name"))
}

// Namespace returns a reference to field namespace of vault_ldap_secret_backend_library_set.
func (lsbls ldapSecretBackendLibrarySetAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(lsbls.ref.Append("namespace"))
}

// ServiceAccountNames returns a reference to field service_account_names of vault_ldap_secret_backend_library_set.
func (lsbls ldapSecretBackendLibrarySetAttributes) ServiceAccountNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](lsbls.ref.Append("service_account_names"))
}

// Ttl returns a reference to field ttl of vault_ldap_secret_backend_library_set.
func (lsbls ldapSecretBackendLibrarySetAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(lsbls.ref.Append("ttl"))
}

type ldapSecretBackendLibrarySetState struct {
	DisableCheckInEnforcement bool     `json:"disable_check_in_enforcement"`
	Id                        string   `json:"id"`
	MaxTtl                    float64  `json:"max_ttl"`
	Mount                     string   `json:"mount"`
	Name                      string   `json:"name"`
	Namespace                 string   `json:"namespace"`
	ServiceAccountNames       []string `json:"service_account_names"`
	Ttl                       float64  `json:"ttl"`
}
