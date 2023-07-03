// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewProxyProtocolPolicy creates a new instance of [ProxyProtocolPolicy].
func NewProxyProtocolPolicy(name string, args ProxyProtocolPolicyArgs) *ProxyProtocolPolicy {
	return &ProxyProtocolPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ProxyProtocolPolicy)(nil)

// ProxyProtocolPolicy represents the Terraform resource aws_proxy_protocol_policy.
type ProxyProtocolPolicy struct {
	Name      string
	Args      ProxyProtocolPolicyArgs
	state     *proxyProtocolPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ProxyProtocolPolicy].
func (ppp *ProxyProtocolPolicy) Type() string {
	return "aws_proxy_protocol_policy"
}

// LocalName returns the local name for [ProxyProtocolPolicy].
func (ppp *ProxyProtocolPolicy) LocalName() string {
	return ppp.Name
}

// Configuration returns the configuration (args) for [ProxyProtocolPolicy].
func (ppp *ProxyProtocolPolicy) Configuration() interface{} {
	return ppp.Args
}

// DependOn is used for other resources to depend on [ProxyProtocolPolicy].
func (ppp *ProxyProtocolPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(ppp)
}

// Dependencies returns the list of resources [ProxyProtocolPolicy] depends_on.
func (ppp *ProxyProtocolPolicy) Dependencies() terra.Dependencies {
	return ppp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ProxyProtocolPolicy].
func (ppp *ProxyProtocolPolicy) LifecycleManagement() *terra.Lifecycle {
	return ppp.Lifecycle
}

// Attributes returns the attributes for [ProxyProtocolPolicy].
func (ppp *ProxyProtocolPolicy) Attributes() proxyProtocolPolicyAttributes {
	return proxyProtocolPolicyAttributes{ref: terra.ReferenceResource(ppp)}
}

// ImportState imports the given attribute values into [ProxyProtocolPolicy]'s state.
func (ppp *ProxyProtocolPolicy) ImportState(av io.Reader) error {
	ppp.state = &proxyProtocolPolicyState{}
	if err := json.NewDecoder(av).Decode(ppp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ppp.Type(), ppp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ProxyProtocolPolicy] has state.
func (ppp *ProxyProtocolPolicy) State() (*proxyProtocolPolicyState, bool) {
	return ppp.state, ppp.state != nil
}

// StateMust returns the state for [ProxyProtocolPolicy]. Panics if the state is nil.
func (ppp *ProxyProtocolPolicy) StateMust() *proxyProtocolPolicyState {
	if ppp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ppp.Type(), ppp.LocalName()))
	}
	return ppp.state
}

// ProxyProtocolPolicyArgs contains the configurations for aws_proxy_protocol_policy.
type ProxyProtocolPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstancePorts: set of string, required
	InstancePorts terra.SetValue[terra.StringValue] `hcl:"instance_ports,attr" validate:"required"`
	// LoadBalancer: string, required
	LoadBalancer terra.StringValue `hcl:"load_balancer,attr" validate:"required"`
}
type proxyProtocolPolicyAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_proxy_protocol_policy.
func (ppp proxyProtocolPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ppp.ref.Append("id"))
}

// InstancePorts returns a reference to field instance_ports of aws_proxy_protocol_policy.
func (ppp proxyProtocolPolicyAttributes) InstancePorts() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ppp.ref.Append("instance_ports"))
}

// LoadBalancer returns a reference to field load_balancer of aws_proxy_protocol_policy.
func (ppp proxyProtocolPolicyAttributes) LoadBalancer() terra.StringValue {
	return terra.ReferenceAsString(ppp.ref.Append("load_balancer"))
}

type proxyProtocolPolicyState struct {
	Id            string   `json:"id"`
	InstancePorts []string `json:"instance_ports"`
	LoadBalancer  string   `json:"load_balancer"`
}