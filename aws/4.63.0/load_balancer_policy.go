// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	loadbalancerpolicy "github.com/golingon/terraproviders/aws/4.63.0/loadbalancerpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLoadBalancerPolicy creates a new instance of [LoadBalancerPolicy].
func NewLoadBalancerPolicy(name string, args LoadBalancerPolicyArgs) *LoadBalancerPolicy {
	return &LoadBalancerPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LoadBalancerPolicy)(nil)

// LoadBalancerPolicy represents the Terraform resource aws_load_balancer_policy.
type LoadBalancerPolicy struct {
	Name      string
	Args      LoadBalancerPolicyArgs
	state     *loadBalancerPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LoadBalancerPolicy].
func (lbp *LoadBalancerPolicy) Type() string {
	return "aws_load_balancer_policy"
}

// LocalName returns the local name for [LoadBalancerPolicy].
func (lbp *LoadBalancerPolicy) LocalName() string {
	return lbp.Name
}

// Configuration returns the configuration (args) for [LoadBalancerPolicy].
func (lbp *LoadBalancerPolicy) Configuration() interface{} {
	return lbp.Args
}

// DependOn is used for other resources to depend on [LoadBalancerPolicy].
func (lbp *LoadBalancerPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(lbp)
}

// Dependencies returns the list of resources [LoadBalancerPolicy] depends_on.
func (lbp *LoadBalancerPolicy) Dependencies() terra.Dependencies {
	return lbp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LoadBalancerPolicy].
func (lbp *LoadBalancerPolicy) LifecycleManagement() *terra.Lifecycle {
	return lbp.Lifecycle
}

// Attributes returns the attributes for [LoadBalancerPolicy].
func (lbp *LoadBalancerPolicy) Attributes() loadBalancerPolicyAttributes {
	return loadBalancerPolicyAttributes{ref: terra.ReferenceResource(lbp)}
}

// ImportState imports the given attribute values into [LoadBalancerPolicy]'s state.
func (lbp *LoadBalancerPolicy) ImportState(av io.Reader) error {
	lbp.state = &loadBalancerPolicyState{}
	if err := json.NewDecoder(av).Decode(lbp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lbp.Type(), lbp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LoadBalancerPolicy] has state.
func (lbp *LoadBalancerPolicy) State() (*loadBalancerPolicyState, bool) {
	return lbp.state, lbp.state != nil
}

// StateMust returns the state for [LoadBalancerPolicy]. Panics if the state is nil.
func (lbp *LoadBalancerPolicy) StateMust() *loadBalancerPolicyState {
	if lbp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lbp.Type(), lbp.LocalName()))
	}
	return lbp.state
}

// LoadBalancerPolicyArgs contains the configurations for aws_load_balancer_policy.
type LoadBalancerPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LoadBalancerName: string, required
	LoadBalancerName terra.StringValue `hcl:"load_balancer_name,attr" validate:"required"`
	// PolicyName: string, required
	PolicyName terra.StringValue `hcl:"policy_name,attr" validate:"required"`
	// PolicyTypeName: string, required
	PolicyTypeName terra.StringValue `hcl:"policy_type_name,attr" validate:"required"`
	// PolicyAttribute: min=0
	PolicyAttribute []loadbalancerpolicy.PolicyAttribute `hcl:"policy_attribute,block" validate:"min=0"`
}
type loadBalancerPolicyAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_load_balancer_policy.
func (lbp loadBalancerPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lbp.ref.Append("id"))
}

// LoadBalancerName returns a reference to field load_balancer_name of aws_load_balancer_policy.
func (lbp loadBalancerPolicyAttributes) LoadBalancerName() terra.StringValue {
	return terra.ReferenceAsString(lbp.ref.Append("load_balancer_name"))
}

// PolicyName returns a reference to field policy_name of aws_load_balancer_policy.
func (lbp loadBalancerPolicyAttributes) PolicyName() terra.StringValue {
	return terra.ReferenceAsString(lbp.ref.Append("policy_name"))
}

// PolicyTypeName returns a reference to field policy_type_name of aws_load_balancer_policy.
func (lbp loadBalancerPolicyAttributes) PolicyTypeName() terra.StringValue {
	return terra.ReferenceAsString(lbp.ref.Append("policy_type_name"))
}

func (lbp loadBalancerPolicyAttributes) PolicyAttribute() terra.SetValue[loadbalancerpolicy.PolicyAttributeAttributes] {
	return terra.ReferenceAsSet[loadbalancerpolicy.PolicyAttributeAttributes](lbp.ref.Append("policy_attribute"))
}

type loadBalancerPolicyState struct {
	Id               string                                    `json:"id"`
	LoadBalancerName string                                    `json:"load_balancer_name"`
	PolicyName       string                                    `json:"policy_name"`
	PolicyTypeName   string                                    `json:"policy_type_name"`
	PolicyAttribute  []loadbalancerpolicy.PolicyAttributeState `json:"policy_attribute"`
}
