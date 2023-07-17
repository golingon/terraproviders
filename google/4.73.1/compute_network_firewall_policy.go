// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computenetworkfirewallpolicy "github.com/golingon/terraproviders/google/4.73.1/computenetworkfirewallpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeNetworkFirewallPolicy creates a new instance of [ComputeNetworkFirewallPolicy].
func NewComputeNetworkFirewallPolicy(name string, args ComputeNetworkFirewallPolicyArgs) *ComputeNetworkFirewallPolicy {
	return &ComputeNetworkFirewallPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeNetworkFirewallPolicy)(nil)

// ComputeNetworkFirewallPolicy represents the Terraform resource google_compute_network_firewall_policy.
type ComputeNetworkFirewallPolicy struct {
	Name      string
	Args      ComputeNetworkFirewallPolicyArgs
	state     *computeNetworkFirewallPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeNetworkFirewallPolicy].
func (cnfp *ComputeNetworkFirewallPolicy) Type() string {
	return "google_compute_network_firewall_policy"
}

// LocalName returns the local name for [ComputeNetworkFirewallPolicy].
func (cnfp *ComputeNetworkFirewallPolicy) LocalName() string {
	return cnfp.Name
}

// Configuration returns the configuration (args) for [ComputeNetworkFirewallPolicy].
func (cnfp *ComputeNetworkFirewallPolicy) Configuration() interface{} {
	return cnfp.Args
}

// DependOn is used for other resources to depend on [ComputeNetworkFirewallPolicy].
func (cnfp *ComputeNetworkFirewallPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(cnfp)
}

// Dependencies returns the list of resources [ComputeNetworkFirewallPolicy] depends_on.
func (cnfp *ComputeNetworkFirewallPolicy) Dependencies() terra.Dependencies {
	return cnfp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeNetworkFirewallPolicy].
func (cnfp *ComputeNetworkFirewallPolicy) LifecycleManagement() *terra.Lifecycle {
	return cnfp.Lifecycle
}

// Attributes returns the attributes for [ComputeNetworkFirewallPolicy].
func (cnfp *ComputeNetworkFirewallPolicy) Attributes() computeNetworkFirewallPolicyAttributes {
	return computeNetworkFirewallPolicyAttributes{ref: terra.ReferenceResource(cnfp)}
}

// ImportState imports the given attribute values into [ComputeNetworkFirewallPolicy]'s state.
func (cnfp *ComputeNetworkFirewallPolicy) ImportState(av io.Reader) error {
	cnfp.state = &computeNetworkFirewallPolicyState{}
	if err := json.NewDecoder(av).Decode(cnfp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cnfp.Type(), cnfp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeNetworkFirewallPolicy] has state.
func (cnfp *ComputeNetworkFirewallPolicy) State() (*computeNetworkFirewallPolicyState, bool) {
	return cnfp.state, cnfp.state != nil
}

// StateMust returns the state for [ComputeNetworkFirewallPolicy]. Panics if the state is nil.
func (cnfp *ComputeNetworkFirewallPolicy) StateMust() *computeNetworkFirewallPolicyState {
	if cnfp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cnfp.Type(), cnfp.LocalName()))
	}
	return cnfp.state
}

// ComputeNetworkFirewallPolicyArgs contains the configurations for google_compute_network_firewall_policy.
type ComputeNetworkFirewallPolicyArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *computenetworkfirewallpolicy.Timeouts `hcl:"timeouts,block"`
}
type computeNetworkFirewallPolicyAttributes struct {
	ref terra.Reference
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_network_firewall_policy.
func (cnfp computeNetworkFirewallPolicyAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(cnfp.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_network_firewall_policy.
func (cnfp computeNetworkFirewallPolicyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cnfp.ref.Append("description"))
}

// Fingerprint returns a reference to field fingerprint of google_compute_network_firewall_policy.
func (cnfp computeNetworkFirewallPolicyAttributes) Fingerprint() terra.StringValue {
	return terra.ReferenceAsString(cnfp.ref.Append("fingerprint"))
}

// Id returns a reference to field id of google_compute_network_firewall_policy.
func (cnfp computeNetworkFirewallPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cnfp.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_network_firewall_policy.
func (cnfp computeNetworkFirewallPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cnfp.ref.Append("name"))
}

// NetworkFirewallPolicyId returns a reference to field network_firewall_policy_id of google_compute_network_firewall_policy.
func (cnfp computeNetworkFirewallPolicyAttributes) NetworkFirewallPolicyId() terra.StringValue {
	return terra.ReferenceAsString(cnfp.ref.Append("network_firewall_policy_id"))
}

// Project returns a reference to field project of google_compute_network_firewall_policy.
func (cnfp computeNetworkFirewallPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cnfp.ref.Append("project"))
}

// RuleTupleCount returns a reference to field rule_tuple_count of google_compute_network_firewall_policy.
func (cnfp computeNetworkFirewallPolicyAttributes) RuleTupleCount() terra.NumberValue {
	return terra.ReferenceAsNumber(cnfp.ref.Append("rule_tuple_count"))
}

// SelfLink returns a reference to field self_link of google_compute_network_firewall_policy.
func (cnfp computeNetworkFirewallPolicyAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cnfp.ref.Append("self_link"))
}

// SelfLinkWithId returns a reference to field self_link_with_id of google_compute_network_firewall_policy.
func (cnfp computeNetworkFirewallPolicyAttributes) SelfLinkWithId() terra.StringValue {
	return terra.ReferenceAsString(cnfp.ref.Append("self_link_with_id"))
}

func (cnfp computeNetworkFirewallPolicyAttributes) Timeouts() computenetworkfirewallpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computenetworkfirewallpolicy.TimeoutsAttributes](cnfp.ref.Append("timeouts"))
}

type computeNetworkFirewallPolicyState struct {
	CreationTimestamp       string                                      `json:"creation_timestamp"`
	Description             string                                      `json:"description"`
	Fingerprint             string                                      `json:"fingerprint"`
	Id                      string                                      `json:"id"`
	Name                    string                                      `json:"name"`
	NetworkFirewallPolicyId string                                      `json:"network_firewall_policy_id"`
	Project                 string                                      `json:"project"`
	RuleTupleCount          float64                                     `json:"rule_tuple_count"`
	SelfLink                string                                      `json:"self_link"`
	SelfLinkWithId          string                                      `json:"self_link_with_id"`
	Timeouts                *computenetworkfirewallpolicy.TimeoutsState `json:"timeouts"`
}
