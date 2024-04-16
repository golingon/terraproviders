// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vault_identity_oidc_provider

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

// Resource represents the Terraform resource vault_identity_oidc_provider.
type Resource struct {
	Name      string
	Args      Args
	state     *vaultIdentityOidcProviderState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (viop *Resource) Type() string {
	return "vault_identity_oidc_provider"
}

// LocalName returns the local name for [Resource].
func (viop *Resource) LocalName() string {
	return viop.Name
}

// Configuration returns the configuration (args) for [Resource].
func (viop *Resource) Configuration() interface{} {
	return viop.Args
}

// DependOn is used for other resources to depend on [Resource].
func (viop *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(viop)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (viop *Resource) Dependencies() terra.Dependencies {
	return viop.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (viop *Resource) LifecycleManagement() *terra.Lifecycle {
	return viop.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (viop *Resource) Attributes() vaultIdentityOidcProviderAttributes {
	return vaultIdentityOidcProviderAttributes{ref: terra.ReferenceResource(viop)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (viop *Resource) ImportState(state io.Reader) error {
	viop.state = &vaultIdentityOidcProviderState{}
	if err := json.NewDecoder(state).Decode(viop.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", viop.Type(), viop.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (viop *Resource) State() (*vaultIdentityOidcProviderState, bool) {
	return viop.state, viop.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (viop *Resource) StateMust() *vaultIdentityOidcProviderState {
	if viop.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", viop.Type(), viop.LocalName()))
	}
	return viop.state
}

// Args contains the configurations for vault_identity_oidc_provider.
type Args struct {
	// AllowedClientIds: set of string, optional
	AllowedClientIds terra.SetValue[terra.StringValue] `hcl:"allowed_client_ids,attr"`
	// HttpsEnabled: bool, optional
	HttpsEnabled terra.BoolValue `hcl:"https_enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IssuerHost: string, optional
	IssuerHost terra.StringValue `hcl:"issuer_host,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// ScopesSupported: set of string, optional
	ScopesSupported terra.SetValue[terra.StringValue] `hcl:"scopes_supported,attr"`
}

type vaultIdentityOidcProviderAttributes struct {
	ref terra.Reference
}

// AllowedClientIds returns a reference to field allowed_client_ids of vault_identity_oidc_provider.
func (viop vaultIdentityOidcProviderAttributes) AllowedClientIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](viop.ref.Append("allowed_client_ids"))
}

// HttpsEnabled returns a reference to field https_enabled of vault_identity_oidc_provider.
func (viop vaultIdentityOidcProviderAttributes) HttpsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(viop.ref.Append("https_enabled"))
}

// Id returns a reference to field id of vault_identity_oidc_provider.
func (viop vaultIdentityOidcProviderAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(viop.ref.Append("id"))
}

// Issuer returns a reference to field issuer of vault_identity_oidc_provider.
func (viop vaultIdentityOidcProviderAttributes) Issuer() terra.StringValue {
	return terra.ReferenceAsString(viop.ref.Append("issuer"))
}

// IssuerHost returns a reference to field issuer_host of vault_identity_oidc_provider.
func (viop vaultIdentityOidcProviderAttributes) IssuerHost() terra.StringValue {
	return terra.ReferenceAsString(viop.ref.Append("issuer_host"))
}

// Name returns a reference to field name of vault_identity_oidc_provider.
func (viop vaultIdentityOidcProviderAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(viop.ref.Append("name"))
}

// Namespace returns a reference to field namespace of vault_identity_oidc_provider.
func (viop vaultIdentityOidcProviderAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(viop.ref.Append("namespace"))
}

// ScopesSupported returns a reference to field scopes_supported of vault_identity_oidc_provider.
func (viop vaultIdentityOidcProviderAttributes) ScopesSupported() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](viop.ref.Append("scopes_supported"))
}

type vaultIdentityOidcProviderState struct {
	AllowedClientIds []string `json:"allowed_client_ids"`
	HttpsEnabled     bool     `json:"https_enabled"`
	Id               string   `json:"id"`
	Issuer           string   `json:"issuer"`
	IssuerHost       string   `json:"issuer_host"`
	Name             string   `json:"name"`
	Namespace        string   `json:"namespace"`
	ScopesSupported  []string `json:"scopes_supported"`
}
