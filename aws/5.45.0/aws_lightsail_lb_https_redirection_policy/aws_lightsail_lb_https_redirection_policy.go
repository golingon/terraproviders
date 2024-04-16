// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_lightsail_lb_https_redirection_policy

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

// Resource represents the Terraform resource aws_lightsail_lb_https_redirection_policy.
type Resource struct {
	Name      string
	Args      Args
	state     *awsLightsailLbHttpsRedirectionPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (allhrp *Resource) Type() string {
	return "aws_lightsail_lb_https_redirection_policy"
}

// LocalName returns the local name for [Resource].
func (allhrp *Resource) LocalName() string {
	return allhrp.Name
}

// Configuration returns the configuration (args) for [Resource].
func (allhrp *Resource) Configuration() interface{} {
	return allhrp.Args
}

// DependOn is used for other resources to depend on [Resource].
func (allhrp *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(allhrp)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (allhrp *Resource) Dependencies() terra.Dependencies {
	return allhrp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (allhrp *Resource) LifecycleManagement() *terra.Lifecycle {
	return allhrp.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (allhrp *Resource) Attributes() awsLightsailLbHttpsRedirectionPolicyAttributes {
	return awsLightsailLbHttpsRedirectionPolicyAttributes{ref: terra.ReferenceResource(allhrp)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (allhrp *Resource) ImportState(state io.Reader) error {
	allhrp.state = &awsLightsailLbHttpsRedirectionPolicyState{}
	if err := json.NewDecoder(state).Decode(allhrp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", allhrp.Type(), allhrp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (allhrp *Resource) State() (*awsLightsailLbHttpsRedirectionPolicyState, bool) {
	return allhrp.state, allhrp.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (allhrp *Resource) StateMust() *awsLightsailLbHttpsRedirectionPolicyState {
	if allhrp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", allhrp.Type(), allhrp.LocalName()))
	}
	return allhrp.state
}

// Args contains the configurations for aws_lightsail_lb_https_redirection_policy.
type Args struct {
	// Enabled: bool, required
	Enabled terra.BoolValue `hcl:"enabled,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LbName: string, required
	LbName terra.StringValue `hcl:"lb_name,attr" validate:"required"`
}

type awsLightsailLbHttpsRedirectionPolicyAttributes struct {
	ref terra.Reference
}

// Enabled returns a reference to field enabled of aws_lightsail_lb_https_redirection_policy.
func (allhrp awsLightsailLbHttpsRedirectionPolicyAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(allhrp.ref.Append("enabled"))
}

// Id returns a reference to field id of aws_lightsail_lb_https_redirection_policy.
func (allhrp awsLightsailLbHttpsRedirectionPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(allhrp.ref.Append("id"))
}

// LbName returns a reference to field lb_name of aws_lightsail_lb_https_redirection_policy.
func (allhrp awsLightsailLbHttpsRedirectionPolicyAttributes) LbName() terra.StringValue {
	return terra.ReferenceAsString(allhrp.ref.Append("lb_name"))
}

type awsLightsailLbHttpsRedirectionPolicyState struct {
	Enabled bool   `json:"enabled"`
	Id      string `json:"id"`
	LbName  string `json:"lb_name"`
}
