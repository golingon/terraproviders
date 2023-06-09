// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	opensearchdomainpolicy "github.com/golingon/terraproviders/aws/4.66.1/opensearchdomainpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewOpensearchDomainPolicy creates a new instance of [OpensearchDomainPolicy].
func NewOpensearchDomainPolicy(name string, args OpensearchDomainPolicyArgs) *OpensearchDomainPolicy {
	return &OpensearchDomainPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OpensearchDomainPolicy)(nil)

// OpensearchDomainPolicy represents the Terraform resource aws_opensearch_domain_policy.
type OpensearchDomainPolicy struct {
	Name      string
	Args      OpensearchDomainPolicyArgs
	state     *opensearchDomainPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [OpensearchDomainPolicy].
func (odp *OpensearchDomainPolicy) Type() string {
	return "aws_opensearch_domain_policy"
}

// LocalName returns the local name for [OpensearchDomainPolicy].
func (odp *OpensearchDomainPolicy) LocalName() string {
	return odp.Name
}

// Configuration returns the configuration (args) for [OpensearchDomainPolicy].
func (odp *OpensearchDomainPolicy) Configuration() interface{} {
	return odp.Args
}

// DependOn is used for other resources to depend on [OpensearchDomainPolicy].
func (odp *OpensearchDomainPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(odp)
}

// Dependencies returns the list of resources [OpensearchDomainPolicy] depends_on.
func (odp *OpensearchDomainPolicy) Dependencies() terra.Dependencies {
	return odp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [OpensearchDomainPolicy].
func (odp *OpensearchDomainPolicy) LifecycleManagement() *terra.Lifecycle {
	return odp.Lifecycle
}

// Attributes returns the attributes for [OpensearchDomainPolicy].
func (odp *OpensearchDomainPolicy) Attributes() opensearchDomainPolicyAttributes {
	return opensearchDomainPolicyAttributes{ref: terra.ReferenceResource(odp)}
}

// ImportState imports the given attribute values into [OpensearchDomainPolicy]'s state.
func (odp *OpensearchDomainPolicy) ImportState(av io.Reader) error {
	odp.state = &opensearchDomainPolicyState{}
	if err := json.NewDecoder(av).Decode(odp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", odp.Type(), odp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [OpensearchDomainPolicy] has state.
func (odp *OpensearchDomainPolicy) State() (*opensearchDomainPolicyState, bool) {
	return odp.state, odp.state != nil
}

// StateMust returns the state for [OpensearchDomainPolicy]. Panics if the state is nil.
func (odp *OpensearchDomainPolicy) StateMust() *opensearchDomainPolicyState {
	if odp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", odp.Type(), odp.LocalName()))
	}
	return odp.state
}

// OpensearchDomainPolicyArgs contains the configurations for aws_opensearch_domain_policy.
type OpensearchDomainPolicyArgs struct {
	// AccessPolicies: string, required
	AccessPolicies terra.StringValue `hcl:"access_policies,attr" validate:"required"`
	// DomainName: string, required
	DomainName terra.StringValue `hcl:"domain_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Timeouts: optional
	Timeouts *opensearchdomainpolicy.Timeouts `hcl:"timeouts,block"`
}
type opensearchDomainPolicyAttributes struct {
	ref terra.Reference
}

// AccessPolicies returns a reference to field access_policies of aws_opensearch_domain_policy.
func (odp opensearchDomainPolicyAttributes) AccessPolicies() terra.StringValue {
	return terra.ReferenceAsString(odp.ref.Append("access_policies"))
}

// DomainName returns a reference to field domain_name of aws_opensearch_domain_policy.
func (odp opensearchDomainPolicyAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(odp.ref.Append("domain_name"))
}

// Id returns a reference to field id of aws_opensearch_domain_policy.
func (odp opensearchDomainPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(odp.ref.Append("id"))
}

func (odp opensearchDomainPolicyAttributes) Timeouts() opensearchdomainpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[opensearchdomainpolicy.TimeoutsAttributes](odp.ref.Append("timeouts"))
}

type opensearchDomainPolicyState struct {
	AccessPolicies string                                `json:"access_policies"`
	DomainName     string                                `json:"domain_name"`
	Id             string                                `json:"id"`
	Timeouts       *opensearchdomainpolicy.TimeoutsState `json:"timeouts"`
}
