// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vault_identity_mfa_totp

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

// Resource represents the Terraform resource vault_identity_mfa_totp.
type Resource struct {
	Name      string
	Args      Args
	state     *vaultIdentityMfaTotpState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (vimt *Resource) Type() string {
	return "vault_identity_mfa_totp"
}

// LocalName returns the local name for [Resource].
func (vimt *Resource) LocalName() string {
	return vimt.Name
}

// Configuration returns the configuration (args) for [Resource].
func (vimt *Resource) Configuration() interface{} {
	return vimt.Args
}

// DependOn is used for other resources to depend on [Resource].
func (vimt *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(vimt)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (vimt *Resource) Dependencies() terra.Dependencies {
	return vimt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (vimt *Resource) LifecycleManagement() *terra.Lifecycle {
	return vimt.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (vimt *Resource) Attributes() vaultIdentityMfaTotpAttributes {
	return vaultIdentityMfaTotpAttributes{ref: terra.ReferenceResource(vimt)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (vimt *Resource) ImportState(state io.Reader) error {
	vimt.state = &vaultIdentityMfaTotpState{}
	if err := json.NewDecoder(state).Decode(vimt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vimt.Type(), vimt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (vimt *Resource) State() (*vaultIdentityMfaTotpState, bool) {
	return vimt.state, vimt.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (vimt *Resource) StateMust() *vaultIdentityMfaTotpState {
	if vimt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vimt.Type(), vimt.LocalName()))
	}
	return vimt.state
}

// Args contains the configurations for vault_identity_mfa_totp.
type Args struct {
	// Algorithm: string, optional
	Algorithm terra.StringValue `hcl:"algorithm,attr"`
	// Digits: number, optional
	Digits terra.NumberValue `hcl:"digits,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Issuer: string, required
	Issuer terra.StringValue `hcl:"issuer,attr" validate:"required"`
	// KeySize: number, optional
	KeySize terra.NumberValue `hcl:"key_size,attr"`
	// MaxValidationAttempts: number, optional
	MaxValidationAttempts terra.NumberValue `hcl:"max_validation_attempts,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Period: number, optional
	Period terra.NumberValue `hcl:"period,attr"`
	// QrSize: number, optional
	QrSize terra.NumberValue `hcl:"qr_size,attr"`
	// Skew: number, optional
	Skew terra.NumberValue `hcl:"skew,attr"`
}

type vaultIdentityMfaTotpAttributes struct {
	ref terra.Reference
}

// Algorithm returns a reference to field algorithm of vault_identity_mfa_totp.
func (vimt vaultIdentityMfaTotpAttributes) Algorithm() terra.StringValue {
	return terra.ReferenceAsString(vimt.ref.Append("algorithm"))
}

// Digits returns a reference to field digits of vault_identity_mfa_totp.
func (vimt vaultIdentityMfaTotpAttributes) Digits() terra.NumberValue {
	return terra.ReferenceAsNumber(vimt.ref.Append("digits"))
}

// Id returns a reference to field id of vault_identity_mfa_totp.
func (vimt vaultIdentityMfaTotpAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vimt.ref.Append("id"))
}

// Issuer returns a reference to field issuer of vault_identity_mfa_totp.
func (vimt vaultIdentityMfaTotpAttributes) Issuer() terra.StringValue {
	return terra.ReferenceAsString(vimt.ref.Append("issuer"))
}

// KeySize returns a reference to field key_size of vault_identity_mfa_totp.
func (vimt vaultIdentityMfaTotpAttributes) KeySize() terra.NumberValue {
	return terra.ReferenceAsNumber(vimt.ref.Append("key_size"))
}

// MaxValidationAttempts returns a reference to field max_validation_attempts of vault_identity_mfa_totp.
func (vimt vaultIdentityMfaTotpAttributes) MaxValidationAttempts() terra.NumberValue {
	return terra.ReferenceAsNumber(vimt.ref.Append("max_validation_attempts"))
}

// MethodId returns a reference to field method_id of vault_identity_mfa_totp.
func (vimt vaultIdentityMfaTotpAttributes) MethodId() terra.StringValue {
	return terra.ReferenceAsString(vimt.ref.Append("method_id"))
}

// MountAccessor returns a reference to field mount_accessor of vault_identity_mfa_totp.
func (vimt vaultIdentityMfaTotpAttributes) MountAccessor() terra.StringValue {
	return terra.ReferenceAsString(vimt.ref.Append("mount_accessor"))
}

// Name returns a reference to field name of vault_identity_mfa_totp.
func (vimt vaultIdentityMfaTotpAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vimt.ref.Append("name"))
}

// Namespace returns a reference to field namespace of vault_identity_mfa_totp.
func (vimt vaultIdentityMfaTotpAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(vimt.ref.Append("namespace"))
}

// NamespaceId returns a reference to field namespace_id of vault_identity_mfa_totp.
func (vimt vaultIdentityMfaTotpAttributes) NamespaceId() terra.StringValue {
	return terra.ReferenceAsString(vimt.ref.Append("namespace_id"))
}

// NamespacePath returns a reference to field namespace_path of vault_identity_mfa_totp.
func (vimt vaultIdentityMfaTotpAttributes) NamespacePath() terra.StringValue {
	return terra.ReferenceAsString(vimt.ref.Append("namespace_path"))
}

// Period returns a reference to field period of vault_identity_mfa_totp.
func (vimt vaultIdentityMfaTotpAttributes) Period() terra.NumberValue {
	return terra.ReferenceAsNumber(vimt.ref.Append("period"))
}

// QrSize returns a reference to field qr_size of vault_identity_mfa_totp.
func (vimt vaultIdentityMfaTotpAttributes) QrSize() terra.NumberValue {
	return terra.ReferenceAsNumber(vimt.ref.Append("qr_size"))
}

// Skew returns a reference to field skew of vault_identity_mfa_totp.
func (vimt vaultIdentityMfaTotpAttributes) Skew() terra.NumberValue {
	return terra.ReferenceAsNumber(vimt.ref.Append("skew"))
}

// Type returns a reference to field type of vault_identity_mfa_totp.
func (vimt vaultIdentityMfaTotpAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(vimt.ref.Append("type"))
}

// Uuid returns a reference to field uuid of vault_identity_mfa_totp.
func (vimt vaultIdentityMfaTotpAttributes) Uuid() terra.StringValue {
	return terra.ReferenceAsString(vimt.ref.Append("uuid"))
}

type vaultIdentityMfaTotpState struct {
	Algorithm             string  `json:"algorithm"`
	Digits                float64 `json:"digits"`
	Id                    string  `json:"id"`
	Issuer                string  `json:"issuer"`
	KeySize               float64 `json:"key_size"`
	MaxValidationAttempts float64 `json:"max_validation_attempts"`
	MethodId              string  `json:"method_id"`
	MountAccessor         string  `json:"mount_accessor"`
	Name                  string  `json:"name"`
	Namespace             string  `json:"namespace"`
	NamespaceId           string  `json:"namespace_id"`
	NamespacePath         string  `json:"namespace_path"`
	Period                float64 `json:"period"`
	QrSize                float64 `json:"qr_size"`
	Skew                  float64 `json:"skew"`
	Type                  string  `json:"type"`
	Uuid                  string  `json:"uuid"`
}
