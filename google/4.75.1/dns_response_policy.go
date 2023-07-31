// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	dnsresponsepolicy "github.com/golingon/terraproviders/google/4.75.1/dnsresponsepolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDnsResponsePolicy creates a new instance of [DnsResponsePolicy].
func NewDnsResponsePolicy(name string, args DnsResponsePolicyArgs) *DnsResponsePolicy {
	return &DnsResponsePolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DnsResponsePolicy)(nil)

// DnsResponsePolicy represents the Terraform resource google_dns_response_policy.
type DnsResponsePolicy struct {
	Name      string
	Args      DnsResponsePolicyArgs
	state     *dnsResponsePolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DnsResponsePolicy].
func (drp *DnsResponsePolicy) Type() string {
	return "google_dns_response_policy"
}

// LocalName returns the local name for [DnsResponsePolicy].
func (drp *DnsResponsePolicy) LocalName() string {
	return drp.Name
}

// Configuration returns the configuration (args) for [DnsResponsePolicy].
func (drp *DnsResponsePolicy) Configuration() interface{} {
	return drp.Args
}

// DependOn is used for other resources to depend on [DnsResponsePolicy].
func (drp *DnsResponsePolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(drp)
}

// Dependencies returns the list of resources [DnsResponsePolicy] depends_on.
func (drp *DnsResponsePolicy) Dependencies() terra.Dependencies {
	return drp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DnsResponsePolicy].
func (drp *DnsResponsePolicy) LifecycleManagement() *terra.Lifecycle {
	return drp.Lifecycle
}

// Attributes returns the attributes for [DnsResponsePolicy].
func (drp *DnsResponsePolicy) Attributes() dnsResponsePolicyAttributes {
	return dnsResponsePolicyAttributes{ref: terra.ReferenceResource(drp)}
}

// ImportState imports the given attribute values into [DnsResponsePolicy]'s state.
func (drp *DnsResponsePolicy) ImportState(av io.Reader) error {
	drp.state = &dnsResponsePolicyState{}
	if err := json.NewDecoder(av).Decode(drp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", drp.Type(), drp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DnsResponsePolicy] has state.
func (drp *DnsResponsePolicy) State() (*dnsResponsePolicyState, bool) {
	return drp.state, drp.state != nil
}

// StateMust returns the state for [DnsResponsePolicy]. Panics if the state is nil.
func (drp *DnsResponsePolicy) StateMust() *dnsResponsePolicyState {
	if drp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", drp.Type(), drp.LocalName()))
	}
	return drp.state
}

// DnsResponsePolicyArgs contains the configurations for google_dns_response_policy.
type DnsResponsePolicyArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// ResponsePolicyName: string, required
	ResponsePolicyName terra.StringValue `hcl:"response_policy_name,attr" validate:"required"`
	// GkeClusters: min=0
	GkeClusters []dnsresponsepolicy.GkeClusters `hcl:"gke_clusters,block" validate:"min=0"`
	// Networks: min=0
	Networks []dnsresponsepolicy.Networks `hcl:"networks,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dnsresponsepolicy.Timeouts `hcl:"timeouts,block"`
}
type dnsResponsePolicyAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_dns_response_policy.
func (drp dnsResponsePolicyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(drp.ref.Append("description"))
}

// Id returns a reference to field id of google_dns_response_policy.
func (drp dnsResponsePolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(drp.ref.Append("id"))
}

// Project returns a reference to field project of google_dns_response_policy.
func (drp dnsResponsePolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(drp.ref.Append("project"))
}

// ResponsePolicyName returns a reference to field response_policy_name of google_dns_response_policy.
func (drp dnsResponsePolicyAttributes) ResponsePolicyName() terra.StringValue {
	return terra.ReferenceAsString(drp.ref.Append("response_policy_name"))
}

func (drp dnsResponsePolicyAttributes) GkeClusters() terra.ListValue[dnsresponsepolicy.GkeClustersAttributes] {
	return terra.ReferenceAsList[dnsresponsepolicy.GkeClustersAttributes](drp.ref.Append("gke_clusters"))
}

func (drp dnsResponsePolicyAttributes) Networks() terra.ListValue[dnsresponsepolicy.NetworksAttributes] {
	return terra.ReferenceAsList[dnsresponsepolicy.NetworksAttributes](drp.ref.Append("networks"))
}

func (drp dnsResponsePolicyAttributes) Timeouts() dnsresponsepolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dnsresponsepolicy.TimeoutsAttributes](drp.ref.Append("timeouts"))
}

type dnsResponsePolicyState struct {
	Description        string                               `json:"description"`
	Id                 string                               `json:"id"`
	Project            string                               `json:"project"`
	ResponsePolicyName string                               `json:"response_policy_name"`
	GkeClusters        []dnsresponsepolicy.GkeClustersState `json:"gke_clusters"`
	Networks           []dnsresponsepolicy.NetworksState    `json:"networks"`
	Timeouts           *dnsresponsepolicy.TimeoutsState     `json:"timeouts"`
}
