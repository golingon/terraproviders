// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSesIdentityPolicy creates a new instance of [SesIdentityPolicy].
func NewSesIdentityPolicy(name string, args SesIdentityPolicyArgs) *SesIdentityPolicy {
	return &SesIdentityPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SesIdentityPolicy)(nil)

// SesIdentityPolicy represents the Terraform resource aws_ses_identity_policy.
type SesIdentityPolicy struct {
	Name      string
	Args      SesIdentityPolicyArgs
	state     *sesIdentityPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SesIdentityPolicy].
func (sip *SesIdentityPolicy) Type() string {
	return "aws_ses_identity_policy"
}

// LocalName returns the local name for [SesIdentityPolicy].
func (sip *SesIdentityPolicy) LocalName() string {
	return sip.Name
}

// Configuration returns the configuration (args) for [SesIdentityPolicy].
func (sip *SesIdentityPolicy) Configuration() interface{} {
	return sip.Args
}

// DependOn is used for other resources to depend on [SesIdentityPolicy].
func (sip *SesIdentityPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(sip)
}

// Dependencies returns the list of resources [SesIdentityPolicy] depends_on.
func (sip *SesIdentityPolicy) Dependencies() terra.Dependencies {
	return sip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SesIdentityPolicy].
func (sip *SesIdentityPolicy) LifecycleManagement() *terra.Lifecycle {
	return sip.Lifecycle
}

// Attributes returns the attributes for [SesIdentityPolicy].
func (sip *SesIdentityPolicy) Attributes() sesIdentityPolicyAttributes {
	return sesIdentityPolicyAttributes{ref: terra.ReferenceResource(sip)}
}

// ImportState imports the given attribute values into [SesIdentityPolicy]'s state.
func (sip *SesIdentityPolicy) ImportState(av io.Reader) error {
	sip.state = &sesIdentityPolicyState{}
	if err := json.NewDecoder(av).Decode(sip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sip.Type(), sip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SesIdentityPolicy] has state.
func (sip *SesIdentityPolicy) State() (*sesIdentityPolicyState, bool) {
	return sip.state, sip.state != nil
}

// StateMust returns the state for [SesIdentityPolicy]. Panics if the state is nil.
func (sip *SesIdentityPolicy) StateMust() *sesIdentityPolicyState {
	if sip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sip.Type(), sip.LocalName()))
	}
	return sip.state
}

// SesIdentityPolicyArgs contains the configurations for aws_ses_identity_policy.
type SesIdentityPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Identity: string, required
	Identity terra.StringValue `hcl:"identity,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Policy: string, required
	Policy terra.StringValue `hcl:"policy,attr" validate:"required"`
}
type sesIdentityPolicyAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ses_identity_policy.
func (sip sesIdentityPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sip.ref.Append("id"))
}

// Identity returns a reference to field identity of aws_ses_identity_policy.
func (sip sesIdentityPolicyAttributes) Identity() terra.StringValue {
	return terra.ReferenceAsString(sip.ref.Append("identity"))
}

// Name returns a reference to field name of aws_ses_identity_policy.
func (sip sesIdentityPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sip.ref.Append("name"))
}

// Policy returns a reference to field policy of aws_ses_identity_policy.
func (sip sesIdentityPolicyAttributes) Policy() terra.StringValue {
	return terra.ReferenceAsString(sip.ref.Append("policy"))
}

type sesIdentityPolicyState struct {
	Id       string `json:"id"`
	Identity string `json:"identity"`
	Name     string `json:"name"`
	Policy   string `json:"policy"`
}
