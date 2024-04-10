// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewIamUserPolicy creates a new instance of [IamUserPolicy].
func NewIamUserPolicy(name string, args IamUserPolicyArgs) *IamUserPolicy {
	return &IamUserPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IamUserPolicy)(nil)

// IamUserPolicy represents the Terraform resource aws_iam_user_policy.
type IamUserPolicy struct {
	Name      string
	Args      IamUserPolicyArgs
	state     *iamUserPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IamUserPolicy].
func (iup *IamUserPolicy) Type() string {
	return "aws_iam_user_policy"
}

// LocalName returns the local name for [IamUserPolicy].
func (iup *IamUserPolicy) LocalName() string {
	return iup.Name
}

// Configuration returns the configuration (args) for [IamUserPolicy].
func (iup *IamUserPolicy) Configuration() interface{} {
	return iup.Args
}

// DependOn is used for other resources to depend on [IamUserPolicy].
func (iup *IamUserPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(iup)
}

// Dependencies returns the list of resources [IamUserPolicy] depends_on.
func (iup *IamUserPolicy) Dependencies() terra.Dependencies {
	return iup.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IamUserPolicy].
func (iup *IamUserPolicy) LifecycleManagement() *terra.Lifecycle {
	return iup.Lifecycle
}

// Attributes returns the attributes for [IamUserPolicy].
func (iup *IamUserPolicy) Attributes() iamUserPolicyAttributes {
	return iamUserPolicyAttributes{ref: terra.ReferenceResource(iup)}
}

// ImportState imports the given attribute values into [IamUserPolicy]'s state.
func (iup *IamUserPolicy) ImportState(av io.Reader) error {
	iup.state = &iamUserPolicyState{}
	if err := json.NewDecoder(av).Decode(iup.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iup.Type(), iup.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IamUserPolicy] has state.
func (iup *IamUserPolicy) State() (*iamUserPolicyState, bool) {
	return iup.state, iup.state != nil
}

// StateMust returns the state for [IamUserPolicy]. Panics if the state is nil.
func (iup *IamUserPolicy) StateMust() *iamUserPolicyState {
	if iup.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iup.Type(), iup.LocalName()))
	}
	return iup.state
}

// IamUserPolicyArgs contains the configurations for aws_iam_user_policy.
type IamUserPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NamePrefix: string, optional
	NamePrefix terra.StringValue `hcl:"name_prefix,attr"`
	// Policy: string, required
	Policy terra.StringValue `hcl:"policy,attr" validate:"required"`
	// User: string, required
	User terra.StringValue `hcl:"user,attr" validate:"required"`
}
type iamUserPolicyAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_iam_user_policy.
func (iup iamUserPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iup.ref.Append("id"))
}

// Name returns a reference to field name of aws_iam_user_policy.
func (iup iamUserPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(iup.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_iam_user_policy.
func (iup iamUserPolicyAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(iup.ref.Append("name_prefix"))
}

// Policy returns a reference to field policy of aws_iam_user_policy.
func (iup iamUserPolicyAttributes) Policy() terra.StringValue {
	return terra.ReferenceAsString(iup.ref.Append("policy"))
}

// User returns a reference to field user of aws_iam_user_policy.
func (iup iamUserPolicyAttributes) User() terra.StringValue {
	return terra.ReferenceAsString(iup.ref.Append("user"))
}

type iamUserPolicyState struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	NamePrefix string `json:"name_prefix"`
	Policy     string `json:"policy"`
	User       string `json:"user"`
}
