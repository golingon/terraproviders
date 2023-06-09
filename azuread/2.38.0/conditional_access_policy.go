// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azuread

import (
	"encoding/json"
	"fmt"
	conditionalaccesspolicy "github.com/golingon/terraproviders/azuread/2.38.0/conditionalaccesspolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewConditionalAccessPolicy creates a new instance of [ConditionalAccessPolicy].
func NewConditionalAccessPolicy(name string, args ConditionalAccessPolicyArgs) *ConditionalAccessPolicy {
	return &ConditionalAccessPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ConditionalAccessPolicy)(nil)

// ConditionalAccessPolicy represents the Terraform resource azuread_conditional_access_policy.
type ConditionalAccessPolicy struct {
	Name      string
	Args      ConditionalAccessPolicyArgs
	state     *conditionalAccessPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ConditionalAccessPolicy].
func (cap *ConditionalAccessPolicy) Type() string {
	return "azuread_conditional_access_policy"
}

// LocalName returns the local name for [ConditionalAccessPolicy].
func (cap *ConditionalAccessPolicy) LocalName() string {
	return cap.Name
}

// Configuration returns the configuration (args) for [ConditionalAccessPolicy].
func (cap *ConditionalAccessPolicy) Configuration() interface{} {
	return cap.Args
}

// DependOn is used for other resources to depend on [ConditionalAccessPolicy].
func (cap *ConditionalAccessPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(cap)
}

// Dependencies returns the list of resources [ConditionalAccessPolicy] depends_on.
func (cap *ConditionalAccessPolicy) Dependencies() terra.Dependencies {
	return cap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ConditionalAccessPolicy].
func (cap *ConditionalAccessPolicy) LifecycleManagement() *terra.Lifecycle {
	return cap.Lifecycle
}

// Attributes returns the attributes for [ConditionalAccessPolicy].
func (cap *ConditionalAccessPolicy) Attributes() conditionalAccessPolicyAttributes {
	return conditionalAccessPolicyAttributes{ref: terra.ReferenceResource(cap)}
}

// ImportState imports the given attribute values into [ConditionalAccessPolicy]'s state.
func (cap *ConditionalAccessPolicy) ImportState(av io.Reader) error {
	cap.state = &conditionalAccessPolicyState{}
	if err := json.NewDecoder(av).Decode(cap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cap.Type(), cap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ConditionalAccessPolicy] has state.
func (cap *ConditionalAccessPolicy) State() (*conditionalAccessPolicyState, bool) {
	return cap.state, cap.state != nil
}

// StateMust returns the state for [ConditionalAccessPolicy]. Panics if the state is nil.
func (cap *ConditionalAccessPolicy) StateMust() *conditionalAccessPolicyState {
	if cap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cap.Type(), cap.LocalName()))
	}
	return cap.state
}

// ConditionalAccessPolicyArgs contains the configurations for azuread_conditional_access_policy.
type ConditionalAccessPolicyArgs struct {
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// State: string, required
	State terra.StringValue `hcl:"state,attr" validate:"required"`
	// Conditions: required
	Conditions *conditionalaccesspolicy.Conditions `hcl:"conditions,block" validate:"required"`
	// GrantControls: required
	GrantControls *conditionalaccesspolicy.GrantControls `hcl:"grant_controls,block" validate:"required"`
	// SessionControls: optional
	SessionControls *conditionalaccesspolicy.SessionControls `hcl:"session_controls,block"`
	// Timeouts: optional
	Timeouts *conditionalaccesspolicy.Timeouts `hcl:"timeouts,block"`
}
type conditionalAccessPolicyAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of azuread_conditional_access_policy.
func (cap conditionalAccessPolicyAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(cap.ref.Append("display_name"))
}

// Id returns a reference to field id of azuread_conditional_access_policy.
func (cap conditionalAccessPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cap.ref.Append("id"))
}

// State returns a reference to field state of azuread_conditional_access_policy.
func (cap conditionalAccessPolicyAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(cap.ref.Append("state"))
}

func (cap conditionalAccessPolicyAttributes) Conditions() terra.ListValue[conditionalaccesspolicy.ConditionsAttributes] {
	return terra.ReferenceAsList[conditionalaccesspolicy.ConditionsAttributes](cap.ref.Append("conditions"))
}

func (cap conditionalAccessPolicyAttributes) GrantControls() terra.ListValue[conditionalaccesspolicy.GrantControlsAttributes] {
	return terra.ReferenceAsList[conditionalaccesspolicy.GrantControlsAttributes](cap.ref.Append("grant_controls"))
}

func (cap conditionalAccessPolicyAttributes) SessionControls() terra.ListValue[conditionalaccesspolicy.SessionControlsAttributes] {
	return terra.ReferenceAsList[conditionalaccesspolicy.SessionControlsAttributes](cap.ref.Append("session_controls"))
}

func (cap conditionalAccessPolicyAttributes) Timeouts() conditionalaccesspolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[conditionalaccesspolicy.TimeoutsAttributes](cap.ref.Append("timeouts"))
}

type conditionalAccessPolicyState struct {
	DisplayName     string                                         `json:"display_name"`
	Id              string                                         `json:"id"`
	State           string                                         `json:"state"`
	Conditions      []conditionalaccesspolicy.ConditionsState      `json:"conditions"`
	GrantControls   []conditionalaccesspolicy.GrantControlsState   `json:"grant_controls"`
	SessionControls []conditionalaccesspolicy.SessionControlsState `json:"session_controls"`
	Timeouts        *conditionalaccesspolicy.TimeoutsState         `json:"timeouts"`
}
