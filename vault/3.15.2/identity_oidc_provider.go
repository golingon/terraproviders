// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIdentityOidcProvider creates a new instance of [IdentityOidcProvider].
func NewIdentityOidcProvider(name string, args IdentityOidcProviderArgs) *IdentityOidcProvider {
	return &IdentityOidcProvider{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IdentityOidcProvider)(nil)

// IdentityOidcProvider represents the Terraform resource vault_identity_oidc_provider.
type IdentityOidcProvider struct {
	Name      string
	Args      IdentityOidcProviderArgs
	state     *identityOidcProviderState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IdentityOidcProvider].
func (iop *IdentityOidcProvider) Type() string {
	return "vault_identity_oidc_provider"
}

// LocalName returns the local name for [IdentityOidcProvider].
func (iop *IdentityOidcProvider) LocalName() string {
	return iop.Name
}

// Configuration returns the configuration (args) for [IdentityOidcProvider].
func (iop *IdentityOidcProvider) Configuration() interface{} {
	return iop.Args
}

// DependOn is used for other resources to depend on [IdentityOidcProvider].
func (iop *IdentityOidcProvider) DependOn() terra.Reference {
	return terra.ReferenceResource(iop)
}

// Dependencies returns the list of resources [IdentityOidcProvider] depends_on.
func (iop *IdentityOidcProvider) Dependencies() terra.Dependencies {
	return iop.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IdentityOidcProvider].
func (iop *IdentityOidcProvider) LifecycleManagement() *terra.Lifecycle {
	return iop.Lifecycle
}

// Attributes returns the attributes for [IdentityOidcProvider].
func (iop *IdentityOidcProvider) Attributes() identityOidcProviderAttributes {
	return identityOidcProviderAttributes{ref: terra.ReferenceResource(iop)}
}

// ImportState imports the given attribute values into [IdentityOidcProvider]'s state.
func (iop *IdentityOidcProvider) ImportState(av io.Reader) error {
	iop.state = &identityOidcProviderState{}
	if err := json.NewDecoder(av).Decode(iop.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iop.Type(), iop.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IdentityOidcProvider] has state.
func (iop *IdentityOidcProvider) State() (*identityOidcProviderState, bool) {
	return iop.state, iop.state != nil
}

// StateMust returns the state for [IdentityOidcProvider]. Panics if the state is nil.
func (iop *IdentityOidcProvider) StateMust() *identityOidcProviderState {
	if iop.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iop.Type(), iop.LocalName()))
	}
	return iop.state
}

// IdentityOidcProviderArgs contains the configurations for vault_identity_oidc_provider.
type IdentityOidcProviderArgs struct {
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
type identityOidcProviderAttributes struct {
	ref terra.Reference
}

// AllowedClientIds returns a reference to field allowed_client_ids of vault_identity_oidc_provider.
func (iop identityOidcProviderAttributes) AllowedClientIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](iop.ref.Append("allowed_client_ids"))
}

// HttpsEnabled returns a reference to field https_enabled of vault_identity_oidc_provider.
func (iop identityOidcProviderAttributes) HttpsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(iop.ref.Append("https_enabled"))
}

// Id returns a reference to field id of vault_identity_oidc_provider.
func (iop identityOidcProviderAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iop.ref.Append("id"))
}

// Issuer returns a reference to field issuer of vault_identity_oidc_provider.
func (iop identityOidcProviderAttributes) Issuer() terra.StringValue {
	return terra.ReferenceAsString(iop.ref.Append("issuer"))
}

// IssuerHost returns a reference to field issuer_host of vault_identity_oidc_provider.
func (iop identityOidcProviderAttributes) IssuerHost() terra.StringValue {
	return terra.ReferenceAsString(iop.ref.Append("issuer_host"))
}

// Name returns a reference to field name of vault_identity_oidc_provider.
func (iop identityOidcProviderAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(iop.ref.Append("name"))
}

// Namespace returns a reference to field namespace of vault_identity_oidc_provider.
func (iop identityOidcProviderAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(iop.ref.Append("namespace"))
}

// ScopesSupported returns a reference to field scopes_supported of vault_identity_oidc_provider.
func (iop identityOidcProviderAttributes) ScopesSupported() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](iop.ref.Append("scopes_supported"))
}

type identityOidcProviderState struct {
	AllowedClientIds []string `json:"allowed_client_ids"`
	HttpsEnabled     bool     `json:"https_enabled"`
	Id               string   `json:"id"`
	Issuer           string   `json:"issuer"`
	IssuerHost       string   `json:"issuer_host"`
	Name             string   `json:"name"`
	Namespace        string   `json:"namespace"`
	ScopesSupported  []string `json:"scopes_supported"`
}