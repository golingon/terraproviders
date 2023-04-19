// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewOktaAuthBackendUser creates a new instance of [OktaAuthBackendUser].
func NewOktaAuthBackendUser(name string, args OktaAuthBackendUserArgs) *OktaAuthBackendUser {
	return &OktaAuthBackendUser{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OktaAuthBackendUser)(nil)

// OktaAuthBackendUser represents the Terraform resource vault_okta_auth_backend_user.
type OktaAuthBackendUser struct {
	Name      string
	Args      OktaAuthBackendUserArgs
	state     *oktaAuthBackendUserState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [OktaAuthBackendUser].
func (oabu *OktaAuthBackendUser) Type() string {
	return "vault_okta_auth_backend_user"
}

// LocalName returns the local name for [OktaAuthBackendUser].
func (oabu *OktaAuthBackendUser) LocalName() string {
	return oabu.Name
}

// Configuration returns the configuration (args) for [OktaAuthBackendUser].
func (oabu *OktaAuthBackendUser) Configuration() interface{} {
	return oabu.Args
}

// DependOn is used for other resources to depend on [OktaAuthBackendUser].
func (oabu *OktaAuthBackendUser) DependOn() terra.Reference {
	return terra.ReferenceResource(oabu)
}

// Dependencies returns the list of resources [OktaAuthBackendUser] depends_on.
func (oabu *OktaAuthBackendUser) Dependencies() terra.Dependencies {
	return oabu.DependsOn
}

// LifecycleManagement returns the lifecycle block for [OktaAuthBackendUser].
func (oabu *OktaAuthBackendUser) LifecycleManagement() *terra.Lifecycle {
	return oabu.Lifecycle
}

// Attributes returns the attributes for [OktaAuthBackendUser].
func (oabu *OktaAuthBackendUser) Attributes() oktaAuthBackendUserAttributes {
	return oktaAuthBackendUserAttributes{ref: terra.ReferenceResource(oabu)}
}

// ImportState imports the given attribute values into [OktaAuthBackendUser]'s state.
func (oabu *OktaAuthBackendUser) ImportState(av io.Reader) error {
	oabu.state = &oktaAuthBackendUserState{}
	if err := json.NewDecoder(av).Decode(oabu.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", oabu.Type(), oabu.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [OktaAuthBackendUser] has state.
func (oabu *OktaAuthBackendUser) State() (*oktaAuthBackendUserState, bool) {
	return oabu.state, oabu.state != nil
}

// StateMust returns the state for [OktaAuthBackendUser]. Panics if the state is nil.
func (oabu *OktaAuthBackendUser) StateMust() *oktaAuthBackendUserState {
	if oabu.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", oabu.Type(), oabu.LocalName()))
	}
	return oabu.state
}

// OktaAuthBackendUserArgs contains the configurations for vault_okta_auth_backend_user.
type OktaAuthBackendUserArgs struct {
	// Groups: set of string, optional
	Groups terra.SetValue[terra.StringValue] `hcl:"groups,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Path: string, required
	Path terra.StringValue `hcl:"path,attr" validate:"required"`
	// Policies: set of string, optional
	Policies terra.SetValue[terra.StringValue] `hcl:"policies,attr"`
	// Username: string, required
	Username terra.StringValue `hcl:"username,attr" validate:"required"`
}
type oktaAuthBackendUserAttributes struct {
	ref terra.Reference
}

// Groups returns a reference to field groups of vault_okta_auth_backend_user.
func (oabu oktaAuthBackendUserAttributes) Groups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](oabu.ref.Append("groups"))
}

// Id returns a reference to field id of vault_okta_auth_backend_user.
func (oabu oktaAuthBackendUserAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(oabu.ref.Append("id"))
}

// Namespace returns a reference to field namespace of vault_okta_auth_backend_user.
func (oabu oktaAuthBackendUserAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(oabu.ref.Append("namespace"))
}

// Path returns a reference to field path of vault_okta_auth_backend_user.
func (oabu oktaAuthBackendUserAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(oabu.ref.Append("path"))
}

// Policies returns a reference to field policies of vault_okta_auth_backend_user.
func (oabu oktaAuthBackendUserAttributes) Policies() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](oabu.ref.Append("policies"))
}

// Username returns a reference to field username of vault_okta_auth_backend_user.
func (oabu oktaAuthBackendUserAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(oabu.ref.Append("username"))
}

type oktaAuthBackendUserState struct {
	Groups    []string `json:"groups"`
	Id        string   `json:"id"`
	Namespace string   `json:"namespace"`
	Path      string   `json:"path"`
	Policies  []string `json:"policies"`
	Username  string   `json:"username"`
}