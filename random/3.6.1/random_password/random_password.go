// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package random_password

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

// Resource represents the Terraform resource random_password.
type Resource struct {
	Name      string
	Args      Args
	state     *randomPasswordState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (rp *Resource) Type() string {
	return "random_password"
}

// LocalName returns the local name for [Resource].
func (rp *Resource) LocalName() string {
	return rp.Name
}

// Configuration returns the configuration (args) for [Resource].
func (rp *Resource) Configuration() interface{} {
	return rp.Args
}

// DependOn is used for other resources to depend on [Resource].
func (rp *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(rp)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (rp *Resource) Dependencies() terra.Dependencies {
	return rp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (rp *Resource) LifecycleManagement() *terra.Lifecycle {
	return rp.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (rp *Resource) Attributes() randomPasswordAttributes {
	return randomPasswordAttributes{ref: terra.ReferenceResource(rp)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (rp *Resource) ImportState(state io.Reader) error {
	rp.state = &randomPasswordState{}
	if err := json.NewDecoder(state).Decode(rp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rp.Type(), rp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (rp *Resource) State() (*randomPasswordState, bool) {
	return rp.state, rp.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (rp *Resource) StateMust() *randomPasswordState {
	if rp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rp.Type(), rp.LocalName()))
	}
	return rp.state
}

// Args contains the configurations for random_password.
type Args struct {
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

type randomPasswordAttributes struct {
	ref terra.Reference
}

// BcryptHash returns a reference to field bcrypt_hash of random_password.
func (rp randomPasswordAttributes) BcryptHash() terra.StringValue {
	return terra.ReferenceAsString(rp.ref.Append("bcrypt_hash"))
}

// Id returns a reference to field id of random_password.
func (rp randomPasswordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rp.ref.Append("id"))
}

// Keepers returns a reference to field keepers of random_password.
func (rp randomPasswordAttributes) Keepers() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rp.ref.Append("keepers"))
}

// Length returns a reference to field length of random_password.
func (rp randomPasswordAttributes) Length() terra.NumberValue {
	return terra.ReferenceAsNumber(rp.ref.Append("length"))
}

// Lower returns a reference to field lower of random_password.
func (rp randomPasswordAttributes) Lower() terra.BoolValue {
	return terra.ReferenceAsBool(rp.ref.Append("lower"))
}

// MinLower returns a reference to field min_lower of random_password.
func (rp randomPasswordAttributes) MinLower() terra.NumberValue {
	return terra.ReferenceAsNumber(rp.ref.Append("min_lower"))
}

// MinNumeric returns a reference to field min_numeric of random_password.
func (rp randomPasswordAttributes) MinNumeric() terra.NumberValue {
	return terra.ReferenceAsNumber(rp.ref.Append("min_numeric"))
}

// MinSpecial returns a reference to field min_special of random_password.
func (rp randomPasswordAttributes) MinSpecial() terra.NumberValue {
	return terra.ReferenceAsNumber(rp.ref.Append("min_special"))
}

// MinUpper returns a reference to field min_upper of random_password.
func (rp randomPasswordAttributes) MinUpper() terra.NumberValue {
	return terra.ReferenceAsNumber(rp.ref.Append("min_upper"))
}

// Number returns a reference to field number of random_password.
func (rp randomPasswordAttributes) Number() terra.BoolValue {
	return terra.ReferenceAsBool(rp.ref.Append("number"))
}

// Numeric returns a reference to field numeric of random_password.
func (rp randomPasswordAttributes) Numeric() terra.BoolValue {
	return terra.ReferenceAsBool(rp.ref.Append("numeric"))
}

// OverrideSpecial returns a reference to field override_special of random_password.
func (rp randomPasswordAttributes) OverrideSpecial() terra.StringValue {
	return terra.ReferenceAsString(rp.ref.Append("override_special"))
}

// Result returns a reference to field result of random_password.
func (rp randomPasswordAttributes) Result() terra.StringValue {
	return terra.ReferenceAsString(rp.ref.Append("result"))
}

// Special returns a reference to field special of random_password.
func (rp randomPasswordAttributes) Special() terra.BoolValue {
	return terra.ReferenceAsBool(rp.ref.Append("special"))
}

// Upper returns a reference to field upper of random_password.
func (rp randomPasswordAttributes) Upper() terra.BoolValue {
	return terra.ReferenceAsBool(rp.ref.Append("upper"))
}

type randomPasswordState struct {
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
