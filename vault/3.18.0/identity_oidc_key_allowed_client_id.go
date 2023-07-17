// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIdentityOidcKeyAllowedClientId creates a new instance of [IdentityOidcKeyAllowedClientId].
func NewIdentityOidcKeyAllowedClientId(name string, args IdentityOidcKeyAllowedClientIdArgs) *IdentityOidcKeyAllowedClientId {
	return &IdentityOidcKeyAllowedClientId{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IdentityOidcKeyAllowedClientId)(nil)

// IdentityOidcKeyAllowedClientId represents the Terraform resource vault_identity_oidc_key_allowed_client_id.
type IdentityOidcKeyAllowedClientId struct {
	Name      string
	Args      IdentityOidcKeyAllowedClientIdArgs
	state     *identityOidcKeyAllowedClientIdState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IdentityOidcKeyAllowedClientId].
func (iokaci *IdentityOidcKeyAllowedClientId) Type() string {
	return "vault_identity_oidc_key_allowed_client_id"
}

// LocalName returns the local name for [IdentityOidcKeyAllowedClientId].
func (iokaci *IdentityOidcKeyAllowedClientId) LocalName() string {
	return iokaci.Name
}

// Configuration returns the configuration (args) for [IdentityOidcKeyAllowedClientId].
func (iokaci *IdentityOidcKeyAllowedClientId) Configuration() interface{} {
	return iokaci.Args
}

// DependOn is used for other resources to depend on [IdentityOidcKeyAllowedClientId].
func (iokaci *IdentityOidcKeyAllowedClientId) DependOn() terra.Reference {
	return terra.ReferenceResource(iokaci)
}

// Dependencies returns the list of resources [IdentityOidcKeyAllowedClientId] depends_on.
func (iokaci *IdentityOidcKeyAllowedClientId) Dependencies() terra.Dependencies {
	return iokaci.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IdentityOidcKeyAllowedClientId].
func (iokaci *IdentityOidcKeyAllowedClientId) LifecycleManagement() *terra.Lifecycle {
	return iokaci.Lifecycle
}

// Attributes returns the attributes for [IdentityOidcKeyAllowedClientId].
func (iokaci *IdentityOidcKeyAllowedClientId) Attributes() identityOidcKeyAllowedClientIdAttributes {
	return identityOidcKeyAllowedClientIdAttributes{ref: terra.ReferenceResource(iokaci)}
}

// ImportState imports the given attribute values into [IdentityOidcKeyAllowedClientId]'s state.
func (iokaci *IdentityOidcKeyAllowedClientId) ImportState(av io.Reader) error {
	iokaci.state = &identityOidcKeyAllowedClientIdState{}
	if err := json.NewDecoder(av).Decode(iokaci.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iokaci.Type(), iokaci.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IdentityOidcKeyAllowedClientId] has state.
func (iokaci *IdentityOidcKeyAllowedClientId) State() (*identityOidcKeyAllowedClientIdState, bool) {
	return iokaci.state, iokaci.state != nil
}

// StateMust returns the state for [IdentityOidcKeyAllowedClientId]. Panics if the state is nil.
func (iokaci *IdentityOidcKeyAllowedClientId) StateMust() *identityOidcKeyAllowedClientIdState {
	if iokaci.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iokaci.Type(), iokaci.LocalName()))
	}
	return iokaci.state
}

// IdentityOidcKeyAllowedClientIdArgs contains the configurations for vault_identity_oidc_key_allowed_client_id.
type IdentityOidcKeyAllowedClientIdArgs struct {
	// AllowedClientId: string, required
	AllowedClientId terra.StringValue `hcl:"allowed_client_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyName: string, required
	KeyName terra.StringValue `hcl:"key_name,attr" validate:"required"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
}
type identityOidcKeyAllowedClientIdAttributes struct {
	ref terra.Reference
}

// AllowedClientId returns a reference to field allowed_client_id of vault_identity_oidc_key_allowed_client_id.
func (iokaci identityOidcKeyAllowedClientIdAttributes) AllowedClientId() terra.StringValue {
	return terra.ReferenceAsString(iokaci.ref.Append("allowed_client_id"))
}

// Id returns a reference to field id of vault_identity_oidc_key_allowed_client_id.
func (iokaci identityOidcKeyAllowedClientIdAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iokaci.ref.Append("id"))
}

// KeyName returns a reference to field key_name of vault_identity_oidc_key_allowed_client_id.
func (iokaci identityOidcKeyAllowedClientIdAttributes) KeyName() terra.StringValue {
	return terra.ReferenceAsString(iokaci.ref.Append("key_name"))
}

// Namespace returns a reference to field namespace of vault_identity_oidc_key_allowed_client_id.
func (iokaci identityOidcKeyAllowedClientIdAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(iokaci.ref.Append("namespace"))
}

type identityOidcKeyAllowedClientIdState struct {
	AllowedClientId string `json:"allowed_client_id"`
	Id              string `json:"id"`
	KeyName         string `json:"key_name"`
	Namespace       string `json:"namespace"`
}