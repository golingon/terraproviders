// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	vpclatticeauthpolicy "github.com/golingon/terraproviders/aws/5.0.1/vpclatticeauthpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVpclatticeAuthPolicy creates a new instance of [VpclatticeAuthPolicy].
func NewVpclatticeAuthPolicy(name string, args VpclatticeAuthPolicyArgs) *VpclatticeAuthPolicy {
	return &VpclatticeAuthPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpclatticeAuthPolicy)(nil)

// VpclatticeAuthPolicy represents the Terraform resource aws_vpclattice_auth_policy.
type VpclatticeAuthPolicy struct {
	Name      string
	Args      VpclatticeAuthPolicyArgs
	state     *vpclatticeAuthPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VpclatticeAuthPolicy].
func (vap *VpclatticeAuthPolicy) Type() string {
	return "aws_vpclattice_auth_policy"
}

// LocalName returns the local name for [VpclatticeAuthPolicy].
func (vap *VpclatticeAuthPolicy) LocalName() string {
	return vap.Name
}

// Configuration returns the configuration (args) for [VpclatticeAuthPolicy].
func (vap *VpclatticeAuthPolicy) Configuration() interface{} {
	return vap.Args
}

// DependOn is used for other resources to depend on [VpclatticeAuthPolicy].
func (vap *VpclatticeAuthPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(vap)
}

// Dependencies returns the list of resources [VpclatticeAuthPolicy] depends_on.
func (vap *VpclatticeAuthPolicy) Dependencies() terra.Dependencies {
	return vap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VpclatticeAuthPolicy].
func (vap *VpclatticeAuthPolicy) LifecycleManagement() *terra.Lifecycle {
	return vap.Lifecycle
}

// Attributes returns the attributes for [VpclatticeAuthPolicy].
func (vap *VpclatticeAuthPolicy) Attributes() vpclatticeAuthPolicyAttributes {
	return vpclatticeAuthPolicyAttributes{ref: terra.ReferenceResource(vap)}
}

// ImportState imports the given attribute values into [VpclatticeAuthPolicy]'s state.
func (vap *VpclatticeAuthPolicy) ImportState(av io.Reader) error {
	vap.state = &vpclatticeAuthPolicyState{}
	if err := json.NewDecoder(av).Decode(vap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vap.Type(), vap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VpclatticeAuthPolicy] has state.
func (vap *VpclatticeAuthPolicy) State() (*vpclatticeAuthPolicyState, bool) {
	return vap.state, vap.state != nil
}

// StateMust returns the state for [VpclatticeAuthPolicy]. Panics if the state is nil.
func (vap *VpclatticeAuthPolicy) StateMust() *vpclatticeAuthPolicyState {
	if vap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vap.Type(), vap.LocalName()))
	}
	return vap.state
}

// VpclatticeAuthPolicyArgs contains the configurations for aws_vpclattice_auth_policy.
type VpclatticeAuthPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Policy: string, required
	Policy terra.StringValue `hcl:"policy,attr" validate:"required"`
	// ResourceIdentifier: string, required
	ResourceIdentifier terra.StringValue `hcl:"resource_identifier,attr" validate:"required"`
	// State: string, optional
	State terra.StringValue `hcl:"state,attr"`
	// Timeouts: optional
	Timeouts *vpclatticeauthpolicy.Timeouts `hcl:"timeouts,block"`
}
type vpclatticeAuthPolicyAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_vpclattice_auth_policy.
func (vap vpclatticeAuthPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vap.ref.Append("id"))
}

// Policy returns a reference to field policy of aws_vpclattice_auth_policy.
func (vap vpclatticeAuthPolicyAttributes) Policy() terra.StringValue {
	return terra.ReferenceAsString(vap.ref.Append("policy"))
}

// ResourceIdentifier returns a reference to field resource_identifier of aws_vpclattice_auth_policy.
func (vap vpclatticeAuthPolicyAttributes) ResourceIdentifier() terra.StringValue {
	return terra.ReferenceAsString(vap.ref.Append("resource_identifier"))
}

// State returns a reference to field state of aws_vpclattice_auth_policy.
func (vap vpclatticeAuthPolicyAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(vap.ref.Append("state"))
}

func (vap vpclatticeAuthPolicyAttributes) Timeouts() vpclatticeauthpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vpclatticeauthpolicy.TimeoutsAttributes](vap.ref.Append("timeouts"))
}

type vpclatticeAuthPolicyState struct {
	Id                 string                              `json:"id"`
	Policy             string                              `json:"policy"`
	ResourceIdentifier string                              `json:"resource_identifier"`
	State              string                              `json:"state"`
	Timeouts           *vpclatticeauthpolicy.TimeoutsState `json:"timeouts"`
}
