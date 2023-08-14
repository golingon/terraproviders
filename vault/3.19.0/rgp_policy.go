// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRgpPolicy creates a new instance of [RgpPolicy].
func NewRgpPolicy(name string, args RgpPolicyArgs) *RgpPolicy {
	return &RgpPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RgpPolicy)(nil)

// RgpPolicy represents the Terraform resource vault_rgp_policy.
type RgpPolicy struct {
	Name      string
	Args      RgpPolicyArgs
	state     *rgpPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RgpPolicy].
func (rp *RgpPolicy) Type() string {
	return "vault_rgp_policy"
}

// LocalName returns the local name for [RgpPolicy].
func (rp *RgpPolicy) LocalName() string {
	return rp.Name
}

// Configuration returns the configuration (args) for [RgpPolicy].
func (rp *RgpPolicy) Configuration() interface{} {
	return rp.Args
}

// DependOn is used for other resources to depend on [RgpPolicy].
func (rp *RgpPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(rp)
}

// Dependencies returns the list of resources [RgpPolicy] depends_on.
func (rp *RgpPolicy) Dependencies() terra.Dependencies {
	return rp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RgpPolicy].
func (rp *RgpPolicy) LifecycleManagement() *terra.Lifecycle {
	return rp.Lifecycle
}

// Attributes returns the attributes for [RgpPolicy].
func (rp *RgpPolicy) Attributes() rgpPolicyAttributes {
	return rgpPolicyAttributes{ref: terra.ReferenceResource(rp)}
}

// ImportState imports the given attribute values into [RgpPolicy]'s state.
func (rp *RgpPolicy) ImportState(av io.Reader) error {
	rp.state = &rgpPolicyState{}
	if err := json.NewDecoder(av).Decode(rp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rp.Type(), rp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RgpPolicy] has state.
func (rp *RgpPolicy) State() (*rgpPolicyState, bool) {
	return rp.state, rp.state != nil
}

// StateMust returns the state for [RgpPolicy]. Panics if the state is nil.
func (rp *RgpPolicy) StateMust() *rgpPolicyState {
	if rp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rp.Type(), rp.LocalName()))
	}
	return rp.state
}

// RgpPolicyArgs contains the configurations for vault_rgp_policy.
type RgpPolicyArgs struct {
	// EnforcementLevel: string, required
	EnforcementLevel terra.StringValue `hcl:"enforcement_level,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Policy: string, required
	Policy terra.StringValue `hcl:"policy,attr" validate:"required"`
}
type rgpPolicyAttributes struct {
	ref terra.Reference
}

// EnforcementLevel returns a reference to field enforcement_level of vault_rgp_policy.
func (rp rgpPolicyAttributes) EnforcementLevel() terra.StringValue {
	return terra.ReferenceAsString(rp.ref.Append("enforcement_level"))
}

// Id returns a reference to field id of vault_rgp_policy.
func (rp rgpPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rp.ref.Append("id"))
}

// Name returns a reference to field name of vault_rgp_policy.
func (rp rgpPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rp.ref.Append("name"))
}

// Namespace returns a reference to field namespace of vault_rgp_policy.
func (rp rgpPolicyAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(rp.ref.Append("namespace"))
}

// Policy returns a reference to field policy of vault_rgp_policy.
func (rp rgpPolicyAttributes) Policy() terra.StringValue {
	return terra.ReferenceAsString(rp.ref.Append("policy"))
}

type rgpPolicyState struct {
	EnforcementLevel string `json:"enforcement_level"`
	Id               string `json:"id"`
	Name             string `json:"name"`
	Namespace        string `json:"namespace"`
	Policy           string `json:"policy"`
}