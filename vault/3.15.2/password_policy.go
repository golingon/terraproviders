// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPasswordPolicy creates a new instance of [PasswordPolicy].
func NewPasswordPolicy(name string, args PasswordPolicyArgs) *PasswordPolicy {
	return &PasswordPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PasswordPolicy)(nil)

// PasswordPolicy represents the Terraform resource vault_password_policy.
type PasswordPolicy struct {
	Name      string
	Args      PasswordPolicyArgs
	state     *passwordPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PasswordPolicy].
func (pp *PasswordPolicy) Type() string {
	return "vault_password_policy"
}

// LocalName returns the local name for [PasswordPolicy].
func (pp *PasswordPolicy) LocalName() string {
	return pp.Name
}

// Configuration returns the configuration (args) for [PasswordPolicy].
func (pp *PasswordPolicy) Configuration() interface{} {
	return pp.Args
}

// DependOn is used for other resources to depend on [PasswordPolicy].
func (pp *PasswordPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(pp)
}

// Dependencies returns the list of resources [PasswordPolicy] depends_on.
func (pp *PasswordPolicy) Dependencies() terra.Dependencies {
	return pp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PasswordPolicy].
func (pp *PasswordPolicy) LifecycleManagement() *terra.Lifecycle {
	return pp.Lifecycle
}

// Attributes returns the attributes for [PasswordPolicy].
func (pp *PasswordPolicy) Attributes() passwordPolicyAttributes {
	return passwordPolicyAttributes{ref: terra.ReferenceResource(pp)}
}

// ImportState imports the given attribute values into [PasswordPolicy]'s state.
func (pp *PasswordPolicy) ImportState(av io.Reader) error {
	pp.state = &passwordPolicyState{}
	if err := json.NewDecoder(av).Decode(pp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pp.Type(), pp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PasswordPolicy] has state.
func (pp *PasswordPolicy) State() (*passwordPolicyState, bool) {
	return pp.state, pp.state != nil
}

// StateMust returns the state for [PasswordPolicy]. Panics if the state is nil.
func (pp *PasswordPolicy) StateMust() *passwordPolicyState {
	if pp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pp.Type(), pp.LocalName()))
	}
	return pp.state
}

// PasswordPolicyArgs contains the configurations for vault_password_policy.
type PasswordPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Policy: string, required
	Policy terra.StringValue `hcl:"policy,attr" validate:"required"`
}
type passwordPolicyAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of vault_password_policy.
func (pp passwordPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pp.ref.Append("id"))
}

// Name returns a reference to field name of vault_password_policy.
func (pp passwordPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pp.ref.Append("name"))
}

// Namespace returns a reference to field namespace of vault_password_policy.
func (pp passwordPolicyAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(pp.ref.Append("namespace"))
}

// Policy returns a reference to field policy of vault_password_policy.
func (pp passwordPolicyAttributes) Policy() terra.StringValue {
	return terra.ReferenceAsString(pp.ref.Append("policy"))
}

type passwordPolicyState struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Policy    string `json:"policy"`
}
