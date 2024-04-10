// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewLdapAuthBackendUser creates a new instance of [LdapAuthBackendUser].
func NewLdapAuthBackendUser(name string, args LdapAuthBackendUserArgs) *LdapAuthBackendUser {
	return &LdapAuthBackendUser{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LdapAuthBackendUser)(nil)

// LdapAuthBackendUser represents the Terraform resource vault_ldap_auth_backend_user.
type LdapAuthBackendUser struct {
	Name      string
	Args      LdapAuthBackendUserArgs
	state     *ldapAuthBackendUserState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LdapAuthBackendUser].
func (labu *LdapAuthBackendUser) Type() string {
	return "vault_ldap_auth_backend_user"
}

// LocalName returns the local name for [LdapAuthBackendUser].
func (labu *LdapAuthBackendUser) LocalName() string {
	return labu.Name
}

// Configuration returns the configuration (args) for [LdapAuthBackendUser].
func (labu *LdapAuthBackendUser) Configuration() interface{} {
	return labu.Args
}

// DependOn is used for other resources to depend on [LdapAuthBackendUser].
func (labu *LdapAuthBackendUser) DependOn() terra.Reference {
	return terra.ReferenceResource(labu)
}

// Dependencies returns the list of resources [LdapAuthBackendUser] depends_on.
func (labu *LdapAuthBackendUser) Dependencies() terra.Dependencies {
	return labu.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LdapAuthBackendUser].
func (labu *LdapAuthBackendUser) LifecycleManagement() *terra.Lifecycle {
	return labu.Lifecycle
}

// Attributes returns the attributes for [LdapAuthBackendUser].
func (labu *LdapAuthBackendUser) Attributes() ldapAuthBackendUserAttributes {
	return ldapAuthBackendUserAttributes{ref: terra.ReferenceResource(labu)}
}

// ImportState imports the given attribute values into [LdapAuthBackendUser]'s state.
func (labu *LdapAuthBackendUser) ImportState(av io.Reader) error {
	labu.state = &ldapAuthBackendUserState{}
	if err := json.NewDecoder(av).Decode(labu.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", labu.Type(), labu.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LdapAuthBackendUser] has state.
func (labu *LdapAuthBackendUser) State() (*ldapAuthBackendUserState, bool) {
	return labu.state, labu.state != nil
}

// StateMust returns the state for [LdapAuthBackendUser]. Panics if the state is nil.
func (labu *LdapAuthBackendUser) StateMust() *ldapAuthBackendUserState {
	if labu.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", labu.Type(), labu.LocalName()))
	}
	return labu.state
}

// LdapAuthBackendUserArgs contains the configurations for vault_ldap_auth_backend_user.
type LdapAuthBackendUserArgs struct {
	// Backend: string, optional
	Backend terra.StringValue `hcl:"backend,attr"`
	// Groups: set of string, optional
	Groups terra.SetValue[terra.StringValue] `hcl:"groups,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Policies: set of string, optional
	Policies terra.SetValue[terra.StringValue] `hcl:"policies,attr"`
	// Username: string, required
	Username terra.StringValue `hcl:"username,attr" validate:"required"`
}
type ldapAuthBackendUserAttributes struct {
	ref terra.Reference
}

// Backend returns a reference to field backend of vault_ldap_auth_backend_user.
func (labu ldapAuthBackendUserAttributes) Backend() terra.StringValue {
	return terra.ReferenceAsString(labu.ref.Append("backend"))
}

// Groups returns a reference to field groups of vault_ldap_auth_backend_user.
func (labu ldapAuthBackendUserAttributes) Groups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](labu.ref.Append("groups"))
}

// Id returns a reference to field id of vault_ldap_auth_backend_user.
func (labu ldapAuthBackendUserAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(labu.ref.Append("id"))
}

// Namespace returns a reference to field namespace of vault_ldap_auth_backend_user.
func (labu ldapAuthBackendUserAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(labu.ref.Append("namespace"))
}

// Policies returns a reference to field policies of vault_ldap_auth_backend_user.
func (labu ldapAuthBackendUserAttributes) Policies() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](labu.ref.Append("policies"))
}

// Username returns a reference to field username of vault_ldap_auth_backend_user.
func (labu ldapAuthBackendUserAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(labu.ref.Append("username"))
}

type ldapAuthBackendUserState struct {
	Backend   string   `json:"backend"`
	Groups    []string `json:"groups"`
	Id        string   `json:"id"`
	Namespace string   `json:"namespace"`
	Policies  []string `json:"policies"`
	Username  string   `json:"username"`
}
