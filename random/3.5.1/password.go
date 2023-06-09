// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package random

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPassword creates a new instance of [Password].
func NewPassword(name string, args PasswordArgs) *Password {
	return &Password{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Password)(nil)

// Password represents the Terraform resource random_password.
type Password struct {
	Name      string
	Args      PasswordArgs
	state     *passwordState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Password].
func (p *Password) Type() string {
	return "random_password"
}

// LocalName returns the local name for [Password].
func (p *Password) LocalName() string {
	return p.Name
}

// Configuration returns the configuration (args) for [Password].
func (p *Password) Configuration() interface{} {
	return p.Args
}

// DependOn is used for other resources to depend on [Password].
func (p *Password) DependOn() terra.Reference {
	return terra.ReferenceResource(p)
}

// Dependencies returns the list of resources [Password] depends_on.
func (p *Password) Dependencies() terra.Dependencies {
	return p.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Password].
func (p *Password) LifecycleManagement() *terra.Lifecycle {
	return p.Lifecycle
}

// Attributes returns the attributes for [Password].
func (p *Password) Attributes() passwordAttributes {
	return passwordAttributes{ref: terra.ReferenceResource(p)}
}

// ImportState imports the given attribute values into [Password]'s state.
func (p *Password) ImportState(av io.Reader) error {
	p.state = &passwordState{}
	if err := json.NewDecoder(av).Decode(p.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", p.Type(), p.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Password] has state.
func (p *Password) State() (*passwordState, bool) {
	return p.state, p.state != nil
}

// StateMust returns the state for [Password]. Panics if the state is nil.
func (p *Password) StateMust() *passwordState {
	if p.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", p.Type(), p.LocalName()))
	}
	return p.state
}

// PasswordArgs contains the configurations for random_password.
type PasswordArgs struct {
	// Keepers: map of string, optional
	Keepers terra.MapValue[terra.StringValue] `hcl:"keepers,attr"`
	// Length: number, required
	Length terra.NumberValue `hcl:"length,attr" validate:"required"`
	// Lower: bool, optional
	Lower terra.BoolValue `hcl:"lower,attr"`
	// MinLower: number, optional
	MinLower terra.NumberValue `hcl:"min_lower,attr"`
	// MinNumeric: number, optional
	MinNumeric terra.NumberValue `hcl:"min_numeric,attr"`
	// MinSpecial: number, optional
	MinSpecial terra.NumberValue `hcl:"min_special,attr"`
	// MinUpper: number, optional
	MinUpper terra.NumberValue `hcl:"min_upper,attr"`
	// Number: bool, optional
	Number terra.BoolValue `hcl:"number,attr"`
	// Numeric: bool, optional
	Numeric terra.BoolValue `hcl:"numeric,attr"`
	// OverrideSpecial: string, optional
	OverrideSpecial terra.StringValue `hcl:"override_special,attr"`
	// Special: bool, optional
	Special terra.BoolValue `hcl:"special,attr"`
	// Upper: bool, optional
	Upper terra.BoolValue `hcl:"upper,attr"`
}
type passwordAttributes struct {
	ref terra.Reference
}

// BcryptHash returns a reference to field bcrypt_hash of random_password.
func (p passwordAttributes) BcryptHash() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("bcrypt_hash"))
}

// Id returns a reference to field id of random_password.
func (p passwordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("id"))
}

// Keepers returns a reference to field keepers of random_password.
func (p passwordAttributes) Keepers() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](p.ref.Append("keepers"))
}

// Length returns a reference to field length of random_password.
func (p passwordAttributes) Length() terra.NumberValue {
	return terra.ReferenceAsNumber(p.ref.Append("length"))
}

// Lower returns a reference to field lower of random_password.
func (p passwordAttributes) Lower() terra.BoolValue {
	return terra.ReferenceAsBool(p.ref.Append("lower"))
}

// MinLower returns a reference to field min_lower of random_password.
func (p passwordAttributes) MinLower() terra.NumberValue {
	return terra.ReferenceAsNumber(p.ref.Append("min_lower"))
}

// MinNumeric returns a reference to field min_numeric of random_password.
func (p passwordAttributes) MinNumeric() terra.NumberValue {
	return terra.ReferenceAsNumber(p.ref.Append("min_numeric"))
}

// MinSpecial returns a reference to field min_special of random_password.
func (p passwordAttributes) MinSpecial() terra.NumberValue {
	return terra.ReferenceAsNumber(p.ref.Append("min_special"))
}

// MinUpper returns a reference to field min_upper of random_password.
func (p passwordAttributes) MinUpper() terra.NumberValue {
	return terra.ReferenceAsNumber(p.ref.Append("min_upper"))
}

// Number returns a reference to field number of random_password.
func (p passwordAttributes) Number() terra.BoolValue {
	return terra.ReferenceAsBool(p.ref.Append("number"))
}

// Numeric returns a reference to field numeric of random_password.
func (p passwordAttributes) Numeric() terra.BoolValue {
	return terra.ReferenceAsBool(p.ref.Append("numeric"))
}

// OverrideSpecial returns a reference to field override_special of random_password.
func (p passwordAttributes) OverrideSpecial() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("override_special"))
}

// Result returns a reference to field result of random_password.
func (p passwordAttributes) Result() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("result"))
}

// Special returns a reference to field special of random_password.
func (p passwordAttributes) Special() terra.BoolValue {
	return terra.ReferenceAsBool(p.ref.Append("special"))
}

// Upper returns a reference to field upper of random_password.
func (p passwordAttributes) Upper() terra.BoolValue {
	return terra.ReferenceAsBool(p.ref.Append("upper"))
}

type passwordState struct {
	BcryptHash      string            `json:"bcrypt_hash"`
	Id              string            `json:"id"`
	Keepers         map[string]string `json:"keepers"`
	Length          float64           `json:"length"`
	Lower           bool              `json:"lower"`
	MinLower        float64           `json:"min_lower"`
	MinNumeric      float64           `json:"min_numeric"`
	MinSpecial      float64           `json:"min_special"`
	MinUpper        float64           `json:"min_upper"`
	Number          bool              `json:"number"`
	Numeric         bool              `json:"numeric"`
	OverrideSpecial string            `json:"override_special"`
	Result          string            `json:"result"`
	Special         bool              `json:"special"`
	Upper           bool              `json:"upper"`
}
