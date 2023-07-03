// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLightsailLbHttpsRedirectionPolicy creates a new instance of [LightsailLbHttpsRedirectionPolicy].
func NewLightsailLbHttpsRedirectionPolicy(name string, args LightsailLbHttpsRedirectionPolicyArgs) *LightsailLbHttpsRedirectionPolicy {
	return &LightsailLbHttpsRedirectionPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LightsailLbHttpsRedirectionPolicy)(nil)

// LightsailLbHttpsRedirectionPolicy represents the Terraform resource aws_lightsail_lb_https_redirection_policy.
type LightsailLbHttpsRedirectionPolicy struct {
	Name      string
	Args      LightsailLbHttpsRedirectionPolicyArgs
	state     *lightsailLbHttpsRedirectionPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LightsailLbHttpsRedirectionPolicy].
func (llhrp *LightsailLbHttpsRedirectionPolicy) Type() string {
	return "aws_lightsail_lb_https_redirection_policy"
}

// LocalName returns the local name for [LightsailLbHttpsRedirectionPolicy].
func (llhrp *LightsailLbHttpsRedirectionPolicy) LocalName() string {
	return llhrp.Name
}

// Configuration returns the configuration (args) for [LightsailLbHttpsRedirectionPolicy].
func (llhrp *LightsailLbHttpsRedirectionPolicy) Configuration() interface{} {
	return llhrp.Args
}

// DependOn is used for other resources to depend on [LightsailLbHttpsRedirectionPolicy].
func (llhrp *LightsailLbHttpsRedirectionPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(llhrp)
}

// Dependencies returns the list of resources [LightsailLbHttpsRedirectionPolicy] depends_on.
func (llhrp *LightsailLbHttpsRedirectionPolicy) Dependencies() terra.Dependencies {
	return llhrp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LightsailLbHttpsRedirectionPolicy].
func (llhrp *LightsailLbHttpsRedirectionPolicy) LifecycleManagement() *terra.Lifecycle {
	return llhrp.Lifecycle
}

// Attributes returns the attributes for [LightsailLbHttpsRedirectionPolicy].
func (llhrp *LightsailLbHttpsRedirectionPolicy) Attributes() lightsailLbHttpsRedirectionPolicyAttributes {
	return lightsailLbHttpsRedirectionPolicyAttributes{ref: terra.ReferenceResource(llhrp)}
}

// ImportState imports the given attribute values into [LightsailLbHttpsRedirectionPolicy]'s state.
func (llhrp *LightsailLbHttpsRedirectionPolicy) ImportState(av io.Reader) error {
	llhrp.state = &lightsailLbHttpsRedirectionPolicyState{}
	if err := json.NewDecoder(av).Decode(llhrp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", llhrp.Type(), llhrp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LightsailLbHttpsRedirectionPolicy] has state.
func (llhrp *LightsailLbHttpsRedirectionPolicy) State() (*lightsailLbHttpsRedirectionPolicyState, bool) {
	return llhrp.state, llhrp.state != nil
}

// StateMust returns the state for [LightsailLbHttpsRedirectionPolicy]. Panics if the state is nil.
func (llhrp *LightsailLbHttpsRedirectionPolicy) StateMust() *lightsailLbHttpsRedirectionPolicyState {
	if llhrp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", llhrp.Type(), llhrp.LocalName()))
	}
	return llhrp.state
}

// LightsailLbHttpsRedirectionPolicyArgs contains the configurations for aws_lightsail_lb_https_redirection_policy.
type LightsailLbHttpsRedirectionPolicyArgs struct {
	// Enabled: bool, required
	Enabled terra.BoolValue `hcl:"enabled,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LbName: string, required
	LbName terra.StringValue `hcl:"lb_name,attr" validate:"required"`
}
type lightsailLbHttpsRedirectionPolicyAttributes struct {
	ref terra.Reference
}

// Enabled returns a reference to field enabled of aws_lightsail_lb_https_redirection_policy.
func (llhrp lightsailLbHttpsRedirectionPolicyAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(llhrp.ref.Append("enabled"))
}

// Id returns a reference to field id of aws_lightsail_lb_https_redirection_policy.
func (llhrp lightsailLbHttpsRedirectionPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(llhrp.ref.Append("id"))
}

// LbName returns a reference to field lb_name of aws_lightsail_lb_https_redirection_policy.
func (llhrp lightsailLbHttpsRedirectionPolicyAttributes) LbName() terra.StringValue {
	return terra.ReferenceAsString(llhrp.ref.Append("lb_name"))
}

type lightsailLbHttpsRedirectionPolicyState struct {
	Enabled bool   `json:"enabled"`
	Id      string `json:"id"`
	LbName  string `json:"lb_name"`
}