// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEgpPolicy creates a new instance of [EgpPolicy].
func NewEgpPolicy(name string, args EgpPolicyArgs) *EgpPolicy {
	return &EgpPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EgpPolicy)(nil)

// EgpPolicy represents the Terraform resource vault_egp_policy.
type EgpPolicy struct {
	Name      string
	Args      EgpPolicyArgs
	state     *egpPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EgpPolicy].
func (ep *EgpPolicy) Type() string {
	return "vault_egp_policy"
}

// LocalName returns the local name for [EgpPolicy].
func (ep *EgpPolicy) LocalName() string {
	return ep.Name
}

// Configuration returns the configuration (args) for [EgpPolicy].
func (ep *EgpPolicy) Configuration() interface{} {
	return ep.Args
}

// DependOn is used for other resources to depend on [EgpPolicy].
func (ep *EgpPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(ep)
}

// Dependencies returns the list of resources [EgpPolicy] depends_on.
func (ep *EgpPolicy) Dependencies() terra.Dependencies {
	return ep.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EgpPolicy].
func (ep *EgpPolicy) LifecycleManagement() *terra.Lifecycle {
	return ep.Lifecycle
}

// Attributes returns the attributes for [EgpPolicy].
func (ep *EgpPolicy) Attributes() egpPolicyAttributes {
	return egpPolicyAttributes{ref: terra.ReferenceResource(ep)}
}

// ImportState imports the given attribute values into [EgpPolicy]'s state.
func (ep *EgpPolicy) ImportState(av io.Reader) error {
	ep.state = &egpPolicyState{}
	if err := json.NewDecoder(av).Decode(ep.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ep.Type(), ep.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EgpPolicy] has state.
func (ep *EgpPolicy) State() (*egpPolicyState, bool) {
	return ep.state, ep.state != nil
}

// StateMust returns the state for [EgpPolicy]. Panics if the state is nil.
func (ep *EgpPolicy) StateMust() *egpPolicyState {
	if ep.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ep.Type(), ep.LocalName()))
	}
	return ep.state
}

// EgpPolicyArgs contains the configurations for vault_egp_policy.
type EgpPolicyArgs struct {
	// EnforcementLevel: string, required
	EnforcementLevel terra.StringValue `hcl:"enforcement_level,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Paths: list of string, required
	Paths terra.ListValue[terra.StringValue] `hcl:"paths,attr" validate:"required"`
	// Policy: string, required
	Policy terra.StringValue `hcl:"policy,attr" validate:"required"`
}
type egpPolicyAttributes struct {
	ref terra.Reference
}

// EnforcementLevel returns a reference to field enforcement_level of vault_egp_policy.
func (ep egpPolicyAttributes) EnforcementLevel() terra.StringValue {
	return terra.ReferenceAsString(ep.ref.Append("enforcement_level"))
}

// Id returns a reference to field id of vault_egp_policy.
func (ep egpPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ep.ref.Append("id"))
}

// Name returns a reference to field name of vault_egp_policy.
func (ep egpPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ep.ref.Append("name"))
}

// Namespace returns a reference to field namespace of vault_egp_policy.
func (ep egpPolicyAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(ep.ref.Append("namespace"))
}

// Paths returns a reference to field paths of vault_egp_policy.
func (ep egpPolicyAttributes) Paths() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ep.ref.Append("paths"))
}

// Policy returns a reference to field policy of vault_egp_policy.
func (ep egpPolicyAttributes) Policy() terra.StringValue {
	return terra.ReferenceAsString(ep.ref.Append("policy"))
}

type egpPolicyState struct {
	EnforcementLevel string   `json:"enforcement_level"`
	Id               string   `json:"id"`
	Name             string   `json:"name"`
	Namespace        string   `json:"namespace"`
	Paths            []string `json:"paths"`
	Policy           string   `json:"policy"`
}