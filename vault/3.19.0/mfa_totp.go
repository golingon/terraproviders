// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMfaTotp creates a new instance of [MfaTotp].
func NewMfaTotp(name string, args MfaTotpArgs) *MfaTotp {
	return &MfaTotp{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MfaTotp)(nil)

// MfaTotp represents the Terraform resource vault_mfa_totp.
type MfaTotp struct {
	Name      string
	Args      MfaTotpArgs
	state     *mfaTotpState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MfaTotp].
func (mt *MfaTotp) Type() string {
	return "vault_mfa_totp"
}

// LocalName returns the local name for [MfaTotp].
func (mt *MfaTotp) LocalName() string {
	return mt.Name
}

// Configuration returns the configuration (args) for [MfaTotp].
func (mt *MfaTotp) Configuration() interface{} {
	return mt.Args
}

// DependOn is used for other resources to depend on [MfaTotp].
func (mt *MfaTotp) DependOn() terra.Reference {
	return terra.ReferenceResource(mt)
}

// Dependencies returns the list of resources [MfaTotp] depends_on.
func (mt *MfaTotp) Dependencies() terra.Dependencies {
	return mt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MfaTotp].
func (mt *MfaTotp) LifecycleManagement() *terra.Lifecycle {
	return mt.Lifecycle
}

// Attributes returns the attributes for [MfaTotp].
func (mt *MfaTotp) Attributes() mfaTotpAttributes {
	return mfaTotpAttributes{ref: terra.ReferenceResource(mt)}
}

// ImportState imports the given attribute values into [MfaTotp]'s state.
func (mt *MfaTotp) ImportState(av io.Reader) error {
	mt.state = &mfaTotpState{}
	if err := json.NewDecoder(av).Decode(mt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mt.Type(), mt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MfaTotp] has state.
func (mt *MfaTotp) State() (*mfaTotpState, bool) {
	return mt.state, mt.state != nil
}

// StateMust returns the state for [MfaTotp]. Panics if the state is nil.
func (mt *MfaTotp) StateMust() *mfaTotpState {
	if mt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mt.Type(), mt.LocalName()))
	}
	return mt.state
}

// MfaTotpArgs contains the configurations for vault_mfa_totp.
type MfaTotpArgs struct {
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
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Period: number, optional
	Period terra.NumberValue `hcl:"period,attr"`
	// QrSize: number, optional
	QrSize terra.NumberValue `hcl:"qr_size,attr"`
	// Skew: number, optional
	Skew terra.NumberValue `hcl:"skew,attr"`
}
type mfaTotpAttributes struct {
	ref terra.Reference
}

// Algorithm returns a reference to field algorithm of vault_mfa_totp.
func (mt mfaTotpAttributes) Algorithm() terra.StringValue {
	return terra.ReferenceAsString(mt.ref.Append("algorithm"))
}

// Digits returns a reference to field digits of vault_mfa_totp.
func (mt mfaTotpAttributes) Digits() terra.NumberValue {
	return terra.ReferenceAsNumber(mt.ref.Append("digits"))
}

// Id returns a reference to field id of vault_mfa_totp.
func (mt mfaTotpAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mt.ref.Append("id"))
}

// Issuer returns a reference to field issuer of vault_mfa_totp.
func (mt mfaTotpAttributes) Issuer() terra.StringValue {
	return terra.ReferenceAsString(mt.ref.Append("issuer"))
}

// KeySize returns a reference to field key_size of vault_mfa_totp.
func (mt mfaTotpAttributes) KeySize() terra.NumberValue {
	return terra.ReferenceAsNumber(mt.ref.Append("key_size"))
}

// Name returns a reference to field name of vault_mfa_totp.
func (mt mfaTotpAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mt.ref.Append("name"))
}

// Namespace returns a reference to field namespace of vault_mfa_totp.
func (mt mfaTotpAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(mt.ref.Append("namespace"))
}

// Period returns a reference to field period of vault_mfa_totp.
func (mt mfaTotpAttributes) Period() terra.NumberValue {
	return terra.ReferenceAsNumber(mt.ref.Append("period"))
}

// QrSize returns a reference to field qr_size of vault_mfa_totp.
func (mt mfaTotpAttributes) QrSize() terra.NumberValue {
	return terra.ReferenceAsNumber(mt.ref.Append("qr_size"))
}

// Skew returns a reference to field skew of vault_mfa_totp.
func (mt mfaTotpAttributes) Skew() terra.NumberValue {
	return terra.ReferenceAsNumber(mt.ref.Append("skew"))
}

type mfaTotpState struct {
	Algorithm string  `json:"algorithm"`
	Digits    float64 `json:"digits"`
	Id        string  `json:"id"`
	Issuer    string  `json:"issuer"`
	KeySize   float64 `json:"key_size"`
	Name      string  `json:"name"`
	Namespace string  `json:"namespace"`
	Period    float64 `json:"period"`
	QrSize    float64 `json:"qr_size"`
	Skew      float64 `json:"skew"`
}
