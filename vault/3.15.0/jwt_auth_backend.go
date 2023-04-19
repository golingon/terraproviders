// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	jwtauthbackend "github.com/golingon/terraproviders/vault/3.15.0/jwtauthbackend"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewJwtAuthBackend creates a new instance of [JwtAuthBackend].
func NewJwtAuthBackend(name string, args JwtAuthBackendArgs) *JwtAuthBackend {
	return &JwtAuthBackend{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*JwtAuthBackend)(nil)

// JwtAuthBackend represents the Terraform resource vault_jwt_auth_backend.
type JwtAuthBackend struct {
	Name      string
	Args      JwtAuthBackendArgs
	state     *jwtAuthBackendState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [JwtAuthBackend].
func (jab *JwtAuthBackend) Type() string {
	return "vault_jwt_auth_backend"
}

// LocalName returns the local name for [JwtAuthBackend].
func (jab *JwtAuthBackend) LocalName() string {
	return jab.Name
}

// Configuration returns the configuration (args) for [JwtAuthBackend].
func (jab *JwtAuthBackend) Configuration() interface{} {
	return jab.Args
}

// DependOn is used for other resources to depend on [JwtAuthBackend].
func (jab *JwtAuthBackend) DependOn() terra.Reference {
	return terra.ReferenceResource(jab)
}

// Dependencies returns the list of resources [JwtAuthBackend] depends_on.
func (jab *JwtAuthBackend) Dependencies() terra.Dependencies {
	return jab.DependsOn
}

// LifecycleManagement returns the lifecycle block for [JwtAuthBackend].
func (jab *JwtAuthBackend) LifecycleManagement() *terra.Lifecycle {
	return jab.Lifecycle
}

// Attributes returns the attributes for [JwtAuthBackend].
func (jab *JwtAuthBackend) Attributes() jwtAuthBackendAttributes {
	return jwtAuthBackendAttributes{ref: terra.ReferenceResource(jab)}
}

// ImportState imports the given attribute values into [JwtAuthBackend]'s state.
func (jab *JwtAuthBackend) ImportState(av io.Reader) error {
	jab.state = &jwtAuthBackendState{}
	if err := json.NewDecoder(av).Decode(jab.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", jab.Type(), jab.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [JwtAuthBackend] has state.
func (jab *JwtAuthBackend) State() (*jwtAuthBackendState, bool) {
	return jab.state, jab.state != nil
}

// StateMust returns the state for [JwtAuthBackend]. Panics if the state is nil.
func (jab *JwtAuthBackend) StateMust() *jwtAuthBackendState {
	if jab.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", jab.Type(), jab.LocalName()))
	}
	return jab.state
}

// JwtAuthBackendArgs contains the configurations for vault_jwt_auth_backend.
type JwtAuthBackendArgs struct {
	// BoundIssuer: string, optional
	BoundIssuer terra.StringValue `hcl:"bound_issuer,attr"`
	// DefaultRole: string, optional
	DefaultRole terra.StringValue `hcl:"default_role,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisableRemount: bool, optional
	DisableRemount terra.BoolValue `hcl:"disable_remount,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// JwksCaPem: string, optional
	JwksCaPem terra.StringValue `hcl:"jwks_ca_pem,attr"`
	// JwksUrl: string, optional
	JwksUrl terra.StringValue `hcl:"jwks_url,attr"`
	// JwtSupportedAlgs: list of string, optional
	JwtSupportedAlgs terra.ListValue[terra.StringValue] `hcl:"jwt_supported_algs,attr"`
	// JwtValidationPubkeys: list of string, optional
	JwtValidationPubkeys terra.ListValue[terra.StringValue] `hcl:"jwt_validation_pubkeys,attr"`
	// Local: bool, optional
	Local terra.BoolValue `hcl:"local,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// NamespaceInState: bool, optional
	NamespaceInState terra.BoolValue `hcl:"namespace_in_state,attr"`
	// OidcClientId: string, optional
	OidcClientId terra.StringValue `hcl:"oidc_client_id,attr"`
	// OidcClientSecret: string, optional
	OidcClientSecret terra.StringValue `hcl:"oidc_client_secret,attr"`
	// OidcDiscoveryCaPem: string, optional
	OidcDiscoveryCaPem terra.StringValue `hcl:"oidc_discovery_ca_pem,attr"`
	// OidcDiscoveryUrl: string, optional
	OidcDiscoveryUrl terra.StringValue `hcl:"oidc_discovery_url,attr"`
	// OidcResponseMode: string, optional
	OidcResponseMode terra.StringValue `hcl:"oidc_response_mode,attr"`
	// OidcResponseTypes: list of string, optional
	OidcResponseTypes terra.ListValue[terra.StringValue] `hcl:"oidc_response_types,attr"`
	// Path: string, optional
	Path terra.StringValue `hcl:"path,attr"`
	// ProviderConfig: map of string, optional
	ProviderConfig terra.MapValue[terra.StringValue] `hcl:"provider_config,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// Tune: min=0
	Tune []jwtauthbackend.Tune `hcl:"tune,block" validate:"min=0"`
}
type jwtAuthBackendAttributes struct {
	ref terra.Reference
}

// Accessor returns a reference to field accessor of vault_jwt_auth_backend.
func (jab jwtAuthBackendAttributes) Accessor() terra.StringValue {
	return terra.ReferenceAsString(jab.ref.Append("accessor"))
}

// BoundIssuer returns a reference to field bound_issuer of vault_jwt_auth_backend.
func (jab jwtAuthBackendAttributes) BoundIssuer() terra.StringValue {
	return terra.ReferenceAsString(jab.ref.Append("bound_issuer"))
}

// DefaultRole returns a reference to field default_role of vault_jwt_auth_backend.
func (jab jwtAuthBackendAttributes) DefaultRole() terra.StringValue {
	return terra.ReferenceAsString(jab.ref.Append("default_role"))
}

// Description returns a reference to field description of vault_jwt_auth_backend.
func (jab jwtAuthBackendAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(jab.ref.Append("description"))
}

// DisableRemount returns a reference to field disable_remount of vault_jwt_auth_backend.
func (jab jwtAuthBackendAttributes) DisableRemount() terra.BoolValue {
	return terra.ReferenceAsBool(jab.ref.Append("disable_remount"))
}

// Id returns a reference to field id of vault_jwt_auth_backend.
func (jab jwtAuthBackendAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(jab.ref.Append("id"))
}

// JwksCaPem returns a reference to field jwks_ca_pem of vault_jwt_auth_backend.
func (jab jwtAuthBackendAttributes) JwksCaPem() terra.StringValue {
	return terra.ReferenceAsString(jab.ref.Append("jwks_ca_pem"))
}

// JwksUrl returns a reference to field jwks_url of vault_jwt_auth_backend.
func (jab jwtAuthBackendAttributes) JwksUrl() terra.StringValue {
	return terra.ReferenceAsString(jab.ref.Append("jwks_url"))
}

// JwtSupportedAlgs returns a reference to field jwt_supported_algs of vault_jwt_auth_backend.
func (jab jwtAuthBackendAttributes) JwtSupportedAlgs() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](jab.ref.Append("jwt_supported_algs"))
}

// JwtValidationPubkeys returns a reference to field jwt_validation_pubkeys of vault_jwt_auth_backend.
func (jab jwtAuthBackendAttributes) JwtValidationPubkeys() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](jab.ref.Append("jwt_validation_pubkeys"))
}

// Local returns a reference to field local of vault_jwt_auth_backend.
func (jab jwtAuthBackendAttributes) Local() terra.BoolValue {
	return terra.ReferenceAsBool(jab.ref.Append("local"))
}

// Namespace returns a reference to field namespace of vault_jwt_auth_backend.
func (jab jwtAuthBackendAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(jab.ref.Append("namespace"))
}

// NamespaceInState returns a reference to field namespace_in_state of vault_jwt_auth_backend.
func (jab jwtAuthBackendAttributes) NamespaceInState() terra.BoolValue {
	return terra.ReferenceAsBool(jab.ref.Append("namespace_in_state"))
}

// OidcClientId returns a reference to field oidc_client_id of vault_jwt_auth_backend.
func (jab jwtAuthBackendAttributes) OidcClientId() terra.StringValue {
	return terra.ReferenceAsString(jab.ref.Append("oidc_client_id"))
}

// OidcClientSecret returns a reference to field oidc_client_secret of vault_jwt_auth_backend.
func (jab jwtAuthBackendAttributes) OidcClientSecret() terra.StringValue {
	return terra.ReferenceAsString(jab.ref.Append("oidc_client_secret"))
}

// OidcDiscoveryCaPem returns a reference to field oidc_discovery_ca_pem of vault_jwt_auth_backend.
func (jab jwtAuthBackendAttributes) OidcDiscoveryCaPem() terra.StringValue {
	return terra.ReferenceAsString(jab.ref.Append("oidc_discovery_ca_pem"))
}

// OidcDiscoveryUrl returns a reference to field oidc_discovery_url of vault_jwt_auth_backend.
func (jab jwtAuthBackendAttributes) OidcDiscoveryUrl() terra.StringValue {
	return terra.ReferenceAsString(jab.ref.Append("oidc_discovery_url"))
}

// OidcResponseMode returns a reference to field oidc_response_mode of vault_jwt_auth_backend.
func (jab jwtAuthBackendAttributes) OidcResponseMode() terra.StringValue {
	return terra.ReferenceAsString(jab.ref.Append("oidc_response_mode"))
}

// OidcResponseTypes returns a reference to field oidc_response_types of vault_jwt_auth_backend.
func (jab jwtAuthBackendAttributes) OidcResponseTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](jab.ref.Append("oidc_response_types"))
}

// Path returns a reference to field path of vault_jwt_auth_backend.
func (jab jwtAuthBackendAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(jab.ref.Append("path"))
}

// ProviderConfig returns a reference to field provider_config of vault_jwt_auth_backend.
func (jab jwtAuthBackendAttributes) ProviderConfig() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](jab.ref.Append("provider_config"))
}

// Type returns a reference to field type of vault_jwt_auth_backend.
func (jab jwtAuthBackendAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(jab.ref.Append("type"))
}

func (jab jwtAuthBackendAttributes) Tune() terra.SetValue[jwtauthbackend.TuneAttributes] {
	return terra.ReferenceAsSet[jwtauthbackend.TuneAttributes](jab.ref.Append("tune"))
}

type jwtAuthBackendState struct {
	Accessor             string                     `json:"accessor"`
	BoundIssuer          string                     `json:"bound_issuer"`
	DefaultRole          string                     `json:"default_role"`
	Description          string                     `json:"description"`
	DisableRemount       bool                       `json:"disable_remount"`
	Id                   string                     `json:"id"`
	JwksCaPem            string                     `json:"jwks_ca_pem"`
	JwksUrl              string                     `json:"jwks_url"`
	JwtSupportedAlgs     []string                   `json:"jwt_supported_algs"`
	JwtValidationPubkeys []string                   `json:"jwt_validation_pubkeys"`
	Local                bool                       `json:"local"`
	Namespace            string                     `json:"namespace"`
	NamespaceInState     bool                       `json:"namespace_in_state"`
	OidcClientId         string                     `json:"oidc_client_id"`
	OidcClientSecret     string                     `json:"oidc_client_secret"`
	OidcDiscoveryCaPem   string                     `json:"oidc_discovery_ca_pem"`
	OidcDiscoveryUrl     string                     `json:"oidc_discovery_url"`
	OidcResponseMode     string                     `json:"oidc_response_mode"`
	OidcResponseTypes    []string                   `json:"oidc_response_types"`
	Path                 string                     `json:"path"`
	ProviderConfig       map[string]string          `json:"provider_config"`
	Type                 string                     `json:"type"`
	Tune                 []jwtauthbackend.TuneState `json:"tune"`
}
