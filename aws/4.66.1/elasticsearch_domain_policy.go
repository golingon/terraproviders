// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	elasticsearchdomainpolicy "github.com/golingon/terraproviders/aws/4.66.1/elasticsearchdomainpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewElasticsearchDomainPolicy creates a new instance of [ElasticsearchDomainPolicy].
func NewElasticsearchDomainPolicy(name string, args ElasticsearchDomainPolicyArgs) *ElasticsearchDomainPolicy {
	return &ElasticsearchDomainPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ElasticsearchDomainPolicy)(nil)

// ElasticsearchDomainPolicy represents the Terraform resource aws_elasticsearch_domain_policy.
type ElasticsearchDomainPolicy struct {
	Name      string
	Args      ElasticsearchDomainPolicyArgs
	state     *elasticsearchDomainPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ElasticsearchDomainPolicy].
func (edp *ElasticsearchDomainPolicy) Type() string {
	return "aws_elasticsearch_domain_policy"
}

// LocalName returns the local name for [ElasticsearchDomainPolicy].
func (edp *ElasticsearchDomainPolicy) LocalName() string {
	return edp.Name
}

// Configuration returns the configuration (args) for [ElasticsearchDomainPolicy].
func (edp *ElasticsearchDomainPolicy) Configuration() interface{} {
	return edp.Args
}

// DependOn is used for other resources to depend on [ElasticsearchDomainPolicy].
func (edp *ElasticsearchDomainPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(edp)
}

// Dependencies returns the list of resources [ElasticsearchDomainPolicy] depends_on.
func (edp *ElasticsearchDomainPolicy) Dependencies() terra.Dependencies {
	return edp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ElasticsearchDomainPolicy].
func (edp *ElasticsearchDomainPolicy) LifecycleManagement() *terra.Lifecycle {
	return edp.Lifecycle
}

// Attributes returns the attributes for [ElasticsearchDomainPolicy].
func (edp *ElasticsearchDomainPolicy) Attributes() elasticsearchDomainPolicyAttributes {
	return elasticsearchDomainPolicyAttributes{ref: terra.ReferenceResource(edp)}
}

// ImportState imports the given attribute values into [ElasticsearchDomainPolicy]'s state.
func (edp *ElasticsearchDomainPolicy) ImportState(av io.Reader) error {
	edp.state = &elasticsearchDomainPolicyState{}
	if err := json.NewDecoder(av).Decode(edp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", edp.Type(), edp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ElasticsearchDomainPolicy] has state.
func (edp *ElasticsearchDomainPolicy) State() (*elasticsearchDomainPolicyState, bool) {
	return edp.state, edp.state != nil
}

// StateMust returns the state for [ElasticsearchDomainPolicy]. Panics if the state is nil.
func (edp *ElasticsearchDomainPolicy) StateMust() *elasticsearchDomainPolicyState {
	if edp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", edp.Type(), edp.LocalName()))
	}
	return edp.state
}

// ElasticsearchDomainPolicyArgs contains the configurations for aws_elasticsearch_domain_policy.
type ElasticsearchDomainPolicyArgs struct {
	// AccessPolicies: string, required
	AccessPolicies terra.StringValue `hcl:"access_policies,attr" validate:"required"`
	// DomainName: string, required
	DomainName terra.StringValue `hcl:"domain_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Timeouts: optional
	Timeouts *elasticsearchdomainpolicy.Timeouts `hcl:"timeouts,block"`
}
type elasticsearchDomainPolicyAttributes struct {
	ref terra.Reference
}

// AccessPolicies returns a reference to field access_policies of aws_elasticsearch_domain_policy.
func (edp elasticsearchDomainPolicyAttributes) AccessPolicies() terra.StringValue {
	return terra.ReferenceAsString(edp.ref.Append("access_policies"))
}

// DomainName returns a reference to field domain_name of aws_elasticsearch_domain_policy.
func (edp elasticsearchDomainPolicyAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(edp.ref.Append("domain_name"))
}

// Id returns a reference to field id of aws_elasticsearch_domain_policy.
func (edp elasticsearchDomainPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(edp.ref.Append("id"))
}

func (edp elasticsearchDomainPolicyAttributes) Timeouts() elasticsearchdomainpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[elasticsearchdomainpolicy.TimeoutsAttributes](edp.ref.Append("timeouts"))
}

type elasticsearchDomainPolicyState struct {
	AccessPolicies string                                   `json:"access_policies"`
	DomainName     string                                   `json:"domain_name"`
	Id             string                                   `json:"id"`
	Timeouts       *elasticsearchdomainpolicy.TimeoutsState `json:"timeouts"`
}