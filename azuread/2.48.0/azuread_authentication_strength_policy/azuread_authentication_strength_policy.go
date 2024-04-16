// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azuread_authentication_strength_policy

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

// Resource represents the Terraform resource azuread_authentication_strength_policy.
type Resource struct {
	Name      string
	Args      Args
	state     *azureadAuthenticationStrengthPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aasp *Resource) Type() string {
	return "azuread_authentication_strength_policy"
}

// LocalName returns the local name for [Resource].
func (aasp *Resource) LocalName() string {
	return aasp.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aasp *Resource) Configuration() interface{} {
	return aasp.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aasp *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aasp)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aasp *Resource) Dependencies() terra.Dependencies {
	return aasp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aasp *Resource) LifecycleManagement() *terra.Lifecycle {
	return aasp.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aasp *Resource) Attributes() azureadAuthenticationStrengthPolicyAttributes {
	return azureadAuthenticationStrengthPolicyAttributes{ref: terra.ReferenceResource(aasp)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aasp *Resource) ImportState(state io.Reader) error {
	aasp.state = &azureadAuthenticationStrengthPolicyState{}
	if err := json.NewDecoder(state).Decode(aasp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aasp.Type(), aasp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aasp *Resource) State() (*azureadAuthenticationStrengthPolicyState, bool) {
	return aasp.state, aasp.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aasp *Resource) StateMust() *azureadAuthenticationStrengthPolicyState {
	if aasp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aasp.Type(), aasp.LocalName()))
	}
	return aasp.state
}

// Args contains the configurations for azuread_authentication_strength_policy.
type Args struct {
	// AllowedCombinations: set of string, required
	AllowedCombinations terra.SetValue[terra.StringValue] `hcl:"allowed_combinations,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azureadAuthenticationStrengthPolicyAttributes struct {
	ref terra.Reference
}

// AllowedCombinations returns a reference to field allowed_combinations of azuread_authentication_strength_policy.
func (aasp azureadAuthenticationStrengthPolicyAttributes) AllowedCombinations() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](aasp.ref.Append("allowed_combinations"))
}

// Description returns a reference to field description of azuread_authentication_strength_policy.
func (aasp azureadAuthenticationStrengthPolicyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(aasp.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azuread_authentication_strength_policy.
func (aasp azureadAuthenticationStrengthPolicyAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(aasp.ref.Append("display_name"))
}

// Id returns a reference to field id of azuread_authentication_strength_policy.
func (aasp azureadAuthenticationStrengthPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aasp.ref.Append("id"))
}

func (aasp azureadAuthenticationStrengthPolicyAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](aasp.ref.Append("timeouts"))
}

type azureadAuthenticationStrengthPolicyState struct {
	AllowedCombinations []string       `json:"allowed_combinations"`
	Description         string         `json:"description"`
	DisplayName         string         `json:"display_name"`
	Id                  string         `json:"id"`
	Timeouts            *TimeoutsState `json:"timeouts"`
}
