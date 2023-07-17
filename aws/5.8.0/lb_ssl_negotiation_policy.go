// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	lbsslnegotiationpolicy "github.com/golingon/terraproviders/aws/5.8.0/lbsslnegotiationpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLbSslNegotiationPolicy creates a new instance of [LbSslNegotiationPolicy].
func NewLbSslNegotiationPolicy(name string, args LbSslNegotiationPolicyArgs) *LbSslNegotiationPolicy {
	return &LbSslNegotiationPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LbSslNegotiationPolicy)(nil)

// LbSslNegotiationPolicy represents the Terraform resource aws_lb_ssl_negotiation_policy.
type LbSslNegotiationPolicy struct {
	Name      string
	Args      LbSslNegotiationPolicyArgs
	state     *lbSslNegotiationPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LbSslNegotiationPolicy].
func (lsnp *LbSslNegotiationPolicy) Type() string {
	return "aws_lb_ssl_negotiation_policy"
}

// LocalName returns the local name for [LbSslNegotiationPolicy].
func (lsnp *LbSslNegotiationPolicy) LocalName() string {
	return lsnp.Name
}

// Configuration returns the configuration (args) for [LbSslNegotiationPolicy].
func (lsnp *LbSslNegotiationPolicy) Configuration() interface{} {
	return lsnp.Args
}

// DependOn is used for other resources to depend on [LbSslNegotiationPolicy].
func (lsnp *LbSslNegotiationPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(lsnp)
}

// Dependencies returns the list of resources [LbSslNegotiationPolicy] depends_on.
func (lsnp *LbSslNegotiationPolicy) Dependencies() terra.Dependencies {
	return lsnp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LbSslNegotiationPolicy].
func (lsnp *LbSslNegotiationPolicy) LifecycleManagement() *terra.Lifecycle {
	return lsnp.Lifecycle
}

// Attributes returns the attributes for [LbSslNegotiationPolicy].
func (lsnp *LbSslNegotiationPolicy) Attributes() lbSslNegotiationPolicyAttributes {
	return lbSslNegotiationPolicyAttributes{ref: terra.ReferenceResource(lsnp)}
}

// ImportState imports the given attribute values into [LbSslNegotiationPolicy]'s state.
func (lsnp *LbSslNegotiationPolicy) ImportState(av io.Reader) error {
	lsnp.state = &lbSslNegotiationPolicyState{}
	if err := json.NewDecoder(av).Decode(lsnp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lsnp.Type(), lsnp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LbSslNegotiationPolicy] has state.
func (lsnp *LbSslNegotiationPolicy) State() (*lbSslNegotiationPolicyState, bool) {
	return lsnp.state, lsnp.state != nil
}

// StateMust returns the state for [LbSslNegotiationPolicy]. Panics if the state is nil.
func (lsnp *LbSslNegotiationPolicy) StateMust() *lbSslNegotiationPolicyState {
	if lsnp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lsnp.Type(), lsnp.LocalName()))
	}
	return lsnp.state
}

// LbSslNegotiationPolicyArgs contains the configurations for aws_lb_ssl_negotiation_policy.
type LbSslNegotiationPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LbPort: number, required
	LbPort terra.NumberValue `hcl:"lb_port,attr" validate:"required"`
	// LoadBalancer: string, required
	LoadBalancer terra.StringValue `hcl:"load_balancer,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Triggers: map of string, optional
	Triggers terra.MapValue[terra.StringValue] `hcl:"triggers,attr"`
	// Attribute: min=0
	Attribute []lbsslnegotiationpolicy.Attribute `hcl:"attribute,block" validate:"min=0"`
}
type lbSslNegotiationPolicyAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_lb_ssl_negotiation_policy.
func (lsnp lbSslNegotiationPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lsnp.ref.Append("id"))
}

// LbPort returns a reference to field lb_port of aws_lb_ssl_negotiation_policy.
func (lsnp lbSslNegotiationPolicyAttributes) LbPort() terra.NumberValue {
	return terra.ReferenceAsNumber(lsnp.ref.Append("lb_port"))
}

// LoadBalancer returns a reference to field load_balancer of aws_lb_ssl_negotiation_policy.
func (lsnp lbSslNegotiationPolicyAttributes) LoadBalancer() terra.StringValue {
	return terra.ReferenceAsString(lsnp.ref.Append("load_balancer"))
}

// Name returns a reference to field name of aws_lb_ssl_negotiation_policy.
func (lsnp lbSslNegotiationPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lsnp.ref.Append("name"))
}

// Triggers returns a reference to field triggers of aws_lb_ssl_negotiation_policy.
func (lsnp lbSslNegotiationPolicyAttributes) Triggers() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lsnp.ref.Append("triggers"))
}

func (lsnp lbSslNegotiationPolicyAttributes) Attribute() terra.SetValue[lbsslnegotiationpolicy.AttributeAttributes] {
	return terra.ReferenceAsSet[lbsslnegotiationpolicy.AttributeAttributes](lsnp.ref.Append("attribute"))
}

type lbSslNegotiationPolicyState struct {
	Id           string                                  `json:"id"`
	LbPort       float64                                 `json:"lb_port"`
	LoadBalancer string                                  `json:"load_balancer"`
	Name         string                                  `json:"name"`
	Triggers     map[string]string                       `json:"triggers"`
	Attribute    []lbsslnegotiationpolicy.AttributeState `json:"attribute"`
}
