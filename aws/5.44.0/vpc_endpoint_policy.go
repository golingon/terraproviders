// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	vpcendpointpolicy "github.com/golingon/terraproviders/aws/5.44.0/vpcendpointpolicy"
	"io"
)

// NewVpcEndpointPolicy creates a new instance of [VpcEndpointPolicy].
func NewVpcEndpointPolicy(name string, args VpcEndpointPolicyArgs) *VpcEndpointPolicy {
	return &VpcEndpointPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpcEndpointPolicy)(nil)

// VpcEndpointPolicy represents the Terraform resource aws_vpc_endpoint_policy.
type VpcEndpointPolicy struct {
	Name      string
	Args      VpcEndpointPolicyArgs
	state     *vpcEndpointPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VpcEndpointPolicy].
func (vep *VpcEndpointPolicy) Type() string {
	return "aws_vpc_endpoint_policy"
}

// LocalName returns the local name for [VpcEndpointPolicy].
func (vep *VpcEndpointPolicy) LocalName() string {
	return vep.Name
}

// Configuration returns the configuration (args) for [VpcEndpointPolicy].
func (vep *VpcEndpointPolicy) Configuration() interface{} {
	return vep.Args
}

// DependOn is used for other resources to depend on [VpcEndpointPolicy].
func (vep *VpcEndpointPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(vep)
}

// Dependencies returns the list of resources [VpcEndpointPolicy] depends_on.
func (vep *VpcEndpointPolicy) Dependencies() terra.Dependencies {
	return vep.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VpcEndpointPolicy].
func (vep *VpcEndpointPolicy) LifecycleManagement() *terra.Lifecycle {
	return vep.Lifecycle
}

// Attributes returns the attributes for [VpcEndpointPolicy].
func (vep *VpcEndpointPolicy) Attributes() vpcEndpointPolicyAttributes {
	return vpcEndpointPolicyAttributes{ref: terra.ReferenceResource(vep)}
}

// ImportState imports the given attribute values into [VpcEndpointPolicy]'s state.
func (vep *VpcEndpointPolicy) ImportState(av io.Reader) error {
	vep.state = &vpcEndpointPolicyState{}
	if err := json.NewDecoder(av).Decode(vep.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vep.Type(), vep.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VpcEndpointPolicy] has state.
func (vep *VpcEndpointPolicy) State() (*vpcEndpointPolicyState, bool) {
	return vep.state, vep.state != nil
}

// StateMust returns the state for [VpcEndpointPolicy]. Panics if the state is nil.
func (vep *VpcEndpointPolicy) StateMust() *vpcEndpointPolicyState {
	if vep.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vep.Type(), vep.LocalName()))
	}
	return vep.state
}

// VpcEndpointPolicyArgs contains the configurations for aws_vpc_endpoint_policy.
type VpcEndpointPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Policy: string, optional
	Policy terra.StringValue `hcl:"policy,attr"`
	// VpcEndpointId: string, required
	VpcEndpointId terra.StringValue `hcl:"vpc_endpoint_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *vpcendpointpolicy.Timeouts `hcl:"timeouts,block"`
}
type vpcEndpointPolicyAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_vpc_endpoint_policy.
func (vep vpcEndpointPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vep.ref.Append("id"))
}

// Policy returns a reference to field policy of aws_vpc_endpoint_policy.
func (vep vpcEndpointPolicyAttributes) Policy() terra.StringValue {
	return terra.ReferenceAsString(vep.ref.Append("policy"))
}

// VpcEndpointId returns a reference to field vpc_endpoint_id of aws_vpc_endpoint_policy.
func (vep vpcEndpointPolicyAttributes) VpcEndpointId() terra.StringValue {
	return terra.ReferenceAsString(vep.ref.Append("vpc_endpoint_id"))
}

func (vep vpcEndpointPolicyAttributes) Timeouts() vpcendpointpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vpcendpointpolicy.TimeoutsAttributes](vep.ref.Append("timeouts"))
}

type vpcEndpointPolicyState struct {
	Id            string                           `json:"id"`
	Policy        string                           `json:"policy"`
	VpcEndpointId string                           `json:"vpc_endpoint_id"`
	Timeouts      *vpcendpointpolicy.TimeoutsState `json:"timeouts"`
}
