// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRedshiftserverlessResourcePolicy creates a new instance of [RedshiftserverlessResourcePolicy].
func NewRedshiftserverlessResourcePolicy(name string, args RedshiftserverlessResourcePolicyArgs) *RedshiftserverlessResourcePolicy {
	return &RedshiftserverlessResourcePolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RedshiftserverlessResourcePolicy)(nil)

// RedshiftserverlessResourcePolicy represents the Terraform resource aws_redshiftserverless_resource_policy.
type RedshiftserverlessResourcePolicy struct {
	Name      string
	Args      RedshiftserverlessResourcePolicyArgs
	state     *redshiftserverlessResourcePolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RedshiftserverlessResourcePolicy].
func (rrp *RedshiftserverlessResourcePolicy) Type() string {
	return "aws_redshiftserverless_resource_policy"
}

// LocalName returns the local name for [RedshiftserverlessResourcePolicy].
func (rrp *RedshiftserverlessResourcePolicy) LocalName() string {
	return rrp.Name
}

// Configuration returns the configuration (args) for [RedshiftserverlessResourcePolicy].
func (rrp *RedshiftserverlessResourcePolicy) Configuration() interface{} {
	return rrp.Args
}

// DependOn is used for other resources to depend on [RedshiftserverlessResourcePolicy].
func (rrp *RedshiftserverlessResourcePolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(rrp)
}

// Dependencies returns the list of resources [RedshiftserverlessResourcePolicy] depends_on.
func (rrp *RedshiftserverlessResourcePolicy) Dependencies() terra.Dependencies {
	return rrp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RedshiftserverlessResourcePolicy].
func (rrp *RedshiftserverlessResourcePolicy) LifecycleManagement() *terra.Lifecycle {
	return rrp.Lifecycle
}

// Attributes returns the attributes for [RedshiftserverlessResourcePolicy].
func (rrp *RedshiftserverlessResourcePolicy) Attributes() redshiftserverlessResourcePolicyAttributes {
	return redshiftserverlessResourcePolicyAttributes{ref: terra.ReferenceResource(rrp)}
}

// ImportState imports the given attribute values into [RedshiftserverlessResourcePolicy]'s state.
func (rrp *RedshiftserverlessResourcePolicy) ImportState(av io.Reader) error {
	rrp.state = &redshiftserverlessResourcePolicyState{}
	if err := json.NewDecoder(av).Decode(rrp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rrp.Type(), rrp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RedshiftserverlessResourcePolicy] has state.
func (rrp *RedshiftserverlessResourcePolicy) State() (*redshiftserverlessResourcePolicyState, bool) {
	return rrp.state, rrp.state != nil
}

// StateMust returns the state for [RedshiftserverlessResourcePolicy]. Panics if the state is nil.
func (rrp *RedshiftserverlessResourcePolicy) StateMust() *redshiftserverlessResourcePolicyState {
	if rrp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rrp.Type(), rrp.LocalName()))
	}
	return rrp.state
}

// RedshiftserverlessResourcePolicyArgs contains the configurations for aws_redshiftserverless_resource_policy.
type RedshiftserverlessResourcePolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Policy: string, required
	Policy terra.StringValue `hcl:"policy,attr" validate:"required"`
	// ResourceArn: string, required
	ResourceArn terra.StringValue `hcl:"resource_arn,attr" validate:"required"`
}
type redshiftserverlessResourcePolicyAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_redshiftserverless_resource_policy.
func (rrp redshiftserverlessResourcePolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rrp.ref.Append("id"))
}

// Policy returns a reference to field policy of aws_redshiftserverless_resource_policy.
func (rrp redshiftserverlessResourcePolicyAttributes) Policy() terra.StringValue {
	return terra.ReferenceAsString(rrp.ref.Append("policy"))
}

// ResourceArn returns a reference to field resource_arn of aws_redshiftserverless_resource_policy.
func (rrp redshiftserverlessResourcePolicyAttributes) ResourceArn() terra.StringValue {
	return terra.ReferenceAsString(rrp.ref.Append("resource_arn"))
}

type redshiftserverlessResourcePolicyState struct {
	Id          string `json:"id"`
	Policy      string `json:"policy"`
	ResourceArn string `json:"resource_arn"`
}
