// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azuread_claims_mapping_policy

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

// Resource represents the Terraform resource azuread_claims_mapping_policy.
type Resource struct {
	Name      string
	Args      Args
	state     *azureadClaimsMappingPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (acmp *Resource) Type() string {
	return "azuread_claims_mapping_policy"
}

// LocalName returns the local name for [Resource].
func (acmp *Resource) LocalName() string {
	return acmp.Name
}

// Configuration returns the configuration (args) for [Resource].
func (acmp *Resource) Configuration() interface{} {
	return acmp.Args
}

// DependOn is used for other resources to depend on [Resource].
func (acmp *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(acmp)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (acmp *Resource) Dependencies() terra.Dependencies {
	return acmp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (acmp *Resource) LifecycleManagement() *terra.Lifecycle {
	return acmp.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (acmp *Resource) Attributes() azureadClaimsMappingPolicyAttributes {
	return azureadClaimsMappingPolicyAttributes{ref: terra.ReferenceResource(acmp)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (acmp *Resource) ImportState(state io.Reader) error {
	acmp.state = &azureadClaimsMappingPolicyState{}
	if err := json.NewDecoder(state).Decode(acmp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acmp.Type(), acmp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (acmp *Resource) State() (*azureadClaimsMappingPolicyState, bool) {
	return acmp.state, acmp.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (acmp *Resource) StateMust() *azureadClaimsMappingPolicyState {
	if acmp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acmp.Type(), acmp.LocalName()))
	}
	return acmp.state
}

// Args contains the configurations for azuread_claims_mapping_policy.
type Args struct {
	// Definition: list of string, required
	Definition terra.ListValue[terra.StringValue] `hcl:"definition,attr" validate:"required"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}

type azureadClaimsMappingPolicyAttributes struct {
	ref terra.Reference
}

// Definition returns a reference to field definition of azuread_claims_mapping_policy.
func (acmp azureadClaimsMappingPolicyAttributes) Definition() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](acmp.ref.Append("definition"))
}

// DisplayName returns a reference to field display_name of azuread_claims_mapping_policy.
func (acmp azureadClaimsMappingPolicyAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(acmp.ref.Append("display_name"))
}

// Id returns a reference to field id of azuread_claims_mapping_policy.
func (acmp azureadClaimsMappingPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acmp.ref.Append("id"))
}

type azureadClaimsMappingPolicyState struct {
	Definition  []string `json:"definition"`
	DisplayName string   `json:"display_name"`
	Id          string   `json:"id"`
}
