// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	oktaauthbackend "github.com/golingon/terraproviders/vault/3.17.0/oktaauthbackend"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewOktaAuthBackend creates a new instance of [OktaAuthBackend].
func NewOktaAuthBackend(name string, args OktaAuthBackendArgs) *OktaAuthBackend {
	return &OktaAuthBackend{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OktaAuthBackend)(nil)

// OktaAuthBackend represents the Terraform resource vault_okta_auth_backend.
type OktaAuthBackend struct {
	Name      string
	Args      OktaAuthBackendArgs
	state     *oktaAuthBackendState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [OktaAuthBackend].
func (oab *OktaAuthBackend) Type() string {
	return "vault_okta_auth_backend"
}

// LocalName returns the local name for [OktaAuthBackend].
func (oab *OktaAuthBackend) LocalName() string {
	return oab.Name
}

// Configuration returns the configuration (args) for [OktaAuthBackend].
func (oab *OktaAuthBackend) Configuration() interface{} {
	return oab.Args
}

// DependOn is used for other resources to depend on [OktaAuthBackend].
func (oab *OktaAuthBackend) DependOn() terra.Reference {
	return terra.ReferenceResource(oab)
}

// Dependencies returns the list of resources [OktaAuthBackend] depends_on.
func (oab *OktaAuthBackend) Dependencies() terra.Dependencies {
	return oab.DependsOn
}

// LifecycleManagement returns the lifecycle block for [OktaAuthBackend].
func (oab *OktaAuthBackend) LifecycleManagement() *terra.Lifecycle {
	return oab.Lifecycle
}

// Attributes returns the attributes for [OktaAuthBackend].
func (oab *OktaAuthBackend) Attributes() oktaAuthBackendAttributes {
	return oktaAuthBackendAttributes{ref: terra.ReferenceResource(oab)}
}

// ImportState imports the given attribute values into [OktaAuthBackend]'s state.
func (oab *OktaAuthBackend) ImportState(av io.Reader) error {
	oab.state = &oktaAuthBackendState{}
	if err := json.NewDecoder(av).Decode(oab.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", oab.Type(), oab.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [OktaAuthBackend] has state.
func (oab *OktaAuthBackend) State() (*oktaAuthBackendState, bool) {
	return oab.state, oab.state != nil
}

// StateMust returns the state for [OktaAuthBackend]. Panics if the state is nil.
func (oab *OktaAuthBackend) StateMust() *oktaAuthBackendState {
	if oab.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", oab.Type(), oab.LocalName()))
	}
	return oab.state
}

// OktaAuthBackendArgs contains the configurations for vault_okta_auth_backend.
type OktaAuthBackendArgs struct {
	// BaseUrl: string, optional
	BaseUrl terra.StringValue `hcl:"base_url,attr"`
	// BypassOktaMfa: bool, optional
	BypassOktaMfa terra.BoolValue `hcl:"bypass_okta_mfa,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisableRemount: bool, optional
	DisableRemount terra.BoolValue `hcl:"disable_remount,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MaxTtl: string, optional
	MaxTtl terra.StringValue `hcl:"max_ttl,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Organization: string, required
	Organization terra.StringValue `hcl:"organization,attr" validate:"required"`
	// Path: string, optional
	Path terra.StringValue `hcl:"path,attr"`
	// Token: string, optional
	Token terra.StringValue `hcl:"token,attr"`
	// Ttl: string, optional
	Ttl terra.StringValue `hcl:"ttl,attr"`
	// Group: min=0
	Group []oktaauthbackend.Group `hcl:"group,block" validate:"min=0"`
	// User: min=0
	User []oktaauthbackend.User `hcl:"user,block" validate:"min=0"`
}
type oktaAuthBackendAttributes struct {
	ref terra.Reference
}

// Accessor returns a reference to field accessor of vault_okta_auth_backend.
func (oab oktaAuthBackendAttributes) Accessor() terra.StringValue {
	return terra.ReferenceAsString(oab.ref.Append("accessor"))
}

// BaseUrl returns a reference to field base_url of vault_okta_auth_backend.
func (oab oktaAuthBackendAttributes) BaseUrl() terra.StringValue {
	return terra.ReferenceAsString(oab.ref.Append("base_url"))
}

// BypassOktaMfa returns a reference to field bypass_okta_mfa of vault_okta_auth_backend.
func (oab oktaAuthBackendAttributes) BypassOktaMfa() terra.BoolValue {
	return terra.ReferenceAsBool(oab.ref.Append("bypass_okta_mfa"))
}

// Description returns a reference to field description of vault_okta_auth_backend.
func (oab oktaAuthBackendAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(oab.ref.Append("description"))
}

// DisableRemount returns a reference to field disable_remount of vault_okta_auth_backend.
func (oab oktaAuthBackendAttributes) DisableRemount() terra.BoolValue {
	return terra.ReferenceAsBool(oab.ref.Append("disable_remount"))
}

// Id returns a reference to field id of vault_okta_auth_backend.
func (oab oktaAuthBackendAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(oab.ref.Append("id"))
}

// MaxTtl returns a reference to field max_ttl of vault_okta_auth_backend.
func (oab oktaAuthBackendAttributes) MaxTtl() terra.StringValue {
	return terra.ReferenceAsString(oab.ref.Append("max_ttl"))
}

// Namespace returns a reference to field namespace of vault_okta_auth_backend.
func (oab oktaAuthBackendAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(oab.ref.Append("namespace"))
}

// Organization returns a reference to field organization of vault_okta_auth_backend.
func (oab oktaAuthBackendAttributes) Organization() terra.StringValue {
	return terra.ReferenceAsString(oab.ref.Append("organization"))
}

// Path returns a reference to field path of vault_okta_auth_backend.
func (oab oktaAuthBackendAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(oab.ref.Append("path"))
}

// Token returns a reference to field token of vault_okta_auth_backend.
func (oab oktaAuthBackendAttributes) Token() terra.StringValue {
	return terra.ReferenceAsString(oab.ref.Append("token"))
}

// Ttl returns a reference to field ttl of vault_okta_auth_backend.
func (oab oktaAuthBackendAttributes) Ttl() terra.StringValue {
	return terra.ReferenceAsString(oab.ref.Append("ttl"))
}

func (oab oktaAuthBackendAttributes) Group() terra.SetValue[oktaauthbackend.GroupAttributes] {
	return terra.ReferenceAsSet[oktaauthbackend.GroupAttributes](oab.ref.Append("group"))
}

func (oab oktaAuthBackendAttributes) User() terra.SetValue[oktaauthbackend.UserAttributes] {
	return terra.ReferenceAsSet[oktaauthbackend.UserAttributes](oab.ref.Append("user"))
}

type oktaAuthBackendState struct {
	Accessor       string                       `json:"accessor"`
	BaseUrl        string                       `json:"base_url"`
	BypassOktaMfa  bool                         `json:"bypass_okta_mfa"`
	Description    string                       `json:"description"`
	DisableRemount bool                         `json:"disable_remount"`
	Id             string                       `json:"id"`
	MaxTtl         string                       `json:"max_ttl"`
	Namespace      string                       `json:"namespace"`
	Organization   string                       `json:"organization"`
	Path           string                       `json:"path"`
	Token          string                       `json:"token"`
	Ttl            string                       `json:"ttl"`
	Group          []oktaauthbackend.GroupState `json:"group"`
	User           []oktaauthbackend.UserState  `json:"user"`
}
