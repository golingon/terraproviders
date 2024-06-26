// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_cloudsearch_domain_service_access_policy

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

// Resource represents the Terraform resource aws_cloudsearch_domain_service_access_policy.
type Resource struct {
	Name      string
	Args      Args
	state     *awsCloudsearchDomainServiceAccessPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (acdsap *Resource) Type() string {
	return "aws_cloudsearch_domain_service_access_policy"
}

// LocalName returns the local name for [Resource].
func (acdsap *Resource) LocalName() string {
	return acdsap.Name
}

// Configuration returns the configuration (args) for [Resource].
func (acdsap *Resource) Configuration() interface{} {
	return acdsap.Args
}

// DependOn is used for other resources to depend on [Resource].
func (acdsap *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(acdsap)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (acdsap *Resource) Dependencies() terra.Dependencies {
	return acdsap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (acdsap *Resource) LifecycleManagement() *terra.Lifecycle {
	return acdsap.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (acdsap *Resource) Attributes() awsCloudsearchDomainServiceAccessPolicyAttributes {
	return awsCloudsearchDomainServiceAccessPolicyAttributes{ref: terra.ReferenceResource(acdsap)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (acdsap *Resource) ImportState(state io.Reader) error {
	acdsap.state = &awsCloudsearchDomainServiceAccessPolicyState{}
	if err := json.NewDecoder(state).Decode(acdsap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acdsap.Type(), acdsap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (acdsap *Resource) State() (*awsCloudsearchDomainServiceAccessPolicyState, bool) {
	return acdsap.state, acdsap.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (acdsap *Resource) StateMust() *awsCloudsearchDomainServiceAccessPolicyState {
	if acdsap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acdsap.Type(), acdsap.LocalName()))
	}
	return acdsap.state
}

// Args contains the configurations for aws_cloudsearch_domain_service_access_policy.
type Args struct {
	// AccessPolicy: string, required
	AccessPolicy terra.StringValue `hcl:"access_policy,attr" validate:"required"`
	// DomainName: string, required
	DomainName terra.StringValue `hcl:"domain_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsCloudsearchDomainServiceAccessPolicyAttributes struct {
	ref terra.Reference
}

// AccessPolicy returns a reference to field access_policy of aws_cloudsearch_domain_service_access_policy.
func (acdsap awsCloudsearchDomainServiceAccessPolicyAttributes) AccessPolicy() terra.StringValue {
	return terra.ReferenceAsString(acdsap.ref.Append("access_policy"))
}

// DomainName returns a reference to field domain_name of aws_cloudsearch_domain_service_access_policy.
func (acdsap awsCloudsearchDomainServiceAccessPolicyAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(acdsap.ref.Append("domain_name"))
}

// Id returns a reference to field id of aws_cloudsearch_domain_service_access_policy.
func (acdsap awsCloudsearchDomainServiceAccessPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acdsap.ref.Append("id"))
}

func (acdsap awsCloudsearchDomainServiceAccessPolicyAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](acdsap.ref.Append("timeouts"))
}

type awsCloudsearchDomainServiceAccessPolicyState struct {
	AccessPolicy string         `json:"access_policy"`
	DomainName   string         `json:"domain_name"`
	Id           string         `json:"id"`
	Timeouts     *TimeoutsState `json:"timeouts"`
}
