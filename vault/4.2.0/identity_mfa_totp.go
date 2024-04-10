// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewIdentityMfaTotp creates a new instance of [IdentityMfaTotp].
func NewIdentityMfaTotp(name string, args IdentityMfaTotpArgs) *IdentityMfaTotp {
	return &IdentityMfaTotp{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IdentityMfaTotp)(nil)

// IdentityMfaTotp represents the Terraform resource vault_identity_mfa_totp.
type IdentityMfaTotp struct {
	Name      string
	Args      IdentityMfaTotpArgs
	state     *identityMfaTotpState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IdentityMfaTotp].
func (imt *IdentityMfaTotp) Type() string {
	return "vault_identity_mfa_totp"
}

// LocalName returns the local name for [IdentityMfaTotp].
func (imt *IdentityMfaTotp) LocalName() string {
	return imt.Name
}

// Configuration returns the configuration (args) for [IdentityMfaTotp].
func (imt *IdentityMfaTotp) Configuration() interface{} {
	return imt.Args
}

// DependOn is used for other resources to depend on [IdentityMfaTotp].
func (imt *IdentityMfaTotp) DependOn() terra.Reference {
	return terra.ReferenceResource(imt)
}

// Dependencies returns the list of resources [IdentityMfaTotp] depends_on.
func (imt *IdentityMfaTotp) Dependencies() terra.Dependencies {
	return imt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IdentityMfaTotp].
func (imt *IdentityMfaTotp) LifecycleManagement() *terra.Lifecycle {
	return imt.Lifecycle
}

// Attributes returns the attributes for [IdentityMfaTotp].
func (imt *IdentityMfaTotp) Attributes() identityMfaTotpAttributes {
	return identityMfaTotpAttributes{ref: terra.ReferenceResource(imt)}
}

// ImportState imports the given attribute values into [IdentityMfaTotp]'s state.
func (imt *IdentityMfaTotp) ImportState(av io.Reader) error {
	imt.state = &identityMfaTotpState{}
	if err := json.NewDecoder(av).Decode(imt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", imt.Type(), imt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IdentityMfaTotp] has state.
func (imt *IdentityMfaTotp) State() (*identityMfaTotpState, bool) {
	return imt.state, imt.state != nil
}

// StateMust returns the state for [IdentityMfaTotp]. Panics if the state is nil.
func (imt *IdentityMfaTotp) StateMust() *identityMfaTotpState {
	if imt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", imt.Type(), imt.LocalName()))
	}
	return imt.state
}

// IdentityMfaTotpArgs contains the configurations for vault_identity_mfa_totp.
type IdentityMfaTotpArgs struct {
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
type identityMfaTotpAttributes struct {
	ref terra.Reference
}

// Algorithm returns a reference to field algorithm of vault_identity_mfa_totp.
func (imt identityMfaTotpAttributes) Algorithm() terra.StringValue {
	return terra.ReferenceAsString(imt.ref.Append("algorithm"))
}

// Digits returns a reference to field digits of vault_identity_mfa_totp.
func (imt identityMfaTotpAttributes) Digits() terra.NumberValue {
	return terra.ReferenceAsNumber(imt.ref.Append("digits"))
}

// Id returns a reference to field id of vault_identity_mfa_totp.
func (imt identityMfaTotpAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(imt.ref.Append("id"))
}

// Issuer returns a reference to field issuer of vault_identity_mfa_totp.
func (imt identityMfaTotpAttributes) Issuer() terra.StringValue {
	return terra.ReferenceAsString(imt.ref.Append("issuer"))
}

// KeySize returns a reference to field key_size of vault_identity_mfa_totp.
func (imt identityMfaTotpAttributes) KeySize() terra.NumberValue {
	return terra.ReferenceAsNumber(imt.ref.Append("key_size"))
}

// MaxValidationAttempts returns a reference to field max_validation_attempts of vault_identity_mfa_totp.
func (imt identityMfaTotpAttributes) MaxValidationAttempts() terra.NumberValue {
	return terra.ReferenceAsNumber(imt.ref.Append("max_validation_attempts"))
}

// MethodId returns a reference to field method_id of vault_identity_mfa_totp.
func (imt identityMfaTotpAttributes) MethodId() terra.StringValue {
	return terra.ReferenceAsString(imt.ref.Append("method_id"))
}

// MountAccessor returns a reference to field mount_accessor of vault_identity_mfa_totp.
func (imt identityMfaTotpAttributes) MountAccessor() terra.StringValue {
	return terra.ReferenceAsString(imt.ref.Append("mount_accessor"))
}

// Name returns a reference to field name of vault_identity_mfa_totp.
func (imt identityMfaTotpAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(imt.ref.Append("name"))
}

// Namespace returns a reference to field namespace of vault_identity_mfa_totp.
func (imt identityMfaTotpAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(imt.ref.Append("namespace"))
}

// NamespaceId returns a reference to field namespace_id of vault_identity_mfa_totp.
func (imt identityMfaTotpAttributes) NamespaceId() terra.StringValue {
	return terra.ReferenceAsString(imt.ref.Append("namespace_id"))
}

// NamespacePath returns a reference to field namespace_path of vault_identity_mfa_totp.
func (imt identityMfaTotpAttributes) NamespacePath() terra.StringValue {
	return terra.ReferenceAsString(imt.ref.Append("namespace_path"))
}

// Period returns a reference to field period of vault_identity_mfa_totp.
func (imt identityMfaTotpAttributes) Period() terra.NumberValue {
	return terra.ReferenceAsNumber(imt.ref.Append("period"))
}

// QrSize returns a reference to field qr_size of vault_identity_mfa_totp.
func (imt identityMfaTotpAttributes) QrSize() terra.NumberValue {
	return terra.ReferenceAsNumber(imt.ref.Append("qr_size"))
}

// Skew returns a reference to field skew of vault_identity_mfa_totp.
func (imt identityMfaTotpAttributes) Skew() terra.NumberValue {
	return terra.ReferenceAsNumber(imt.ref.Append("skew"))
}

// Type returns a reference to field type of vault_identity_mfa_totp.
func (imt identityMfaTotpAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(imt.ref.Append("type"))
}

// Uuid returns a reference to field uuid of vault_identity_mfa_totp.
func (imt identityMfaTotpAttributes) Uuid() terra.StringValue {
	return terra.ReferenceAsString(imt.ref.Append("uuid"))
}

type identityMfaTotpState struct {
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
