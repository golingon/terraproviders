// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computeregionnetworkfirewallpolicy "github.com/golingon/terraproviders/google/4.74.0/computeregionnetworkfirewallpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeRegionNetworkFirewallPolicy creates a new instance of [ComputeRegionNetworkFirewallPolicy].
func NewComputeRegionNetworkFirewallPolicy(name string, args ComputeRegionNetworkFirewallPolicyArgs) *ComputeRegionNetworkFirewallPolicy {
	return &ComputeRegionNetworkFirewallPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeRegionNetworkFirewallPolicy)(nil)

// ComputeRegionNetworkFirewallPolicy represents the Terraform resource google_compute_region_network_firewall_policy.
type ComputeRegionNetworkFirewallPolicy struct {
	Name      string
	Args      ComputeRegionNetworkFirewallPolicyArgs
	state     *computeRegionNetworkFirewallPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeRegionNetworkFirewallPolicy].
func (crnfp *ComputeRegionNetworkFirewallPolicy) Type() string {
	return "google_compute_region_network_firewall_policy"
}

// LocalName returns the local name for [ComputeRegionNetworkFirewallPolicy].
func (crnfp *ComputeRegionNetworkFirewallPolicy) LocalName() string {
	return crnfp.Name
}

// Configuration returns the configuration (args) for [ComputeRegionNetworkFirewallPolicy].
func (crnfp *ComputeRegionNetworkFirewallPolicy) Configuration() interface{} {
	return crnfp.Args
}

// DependOn is used for other resources to depend on [ComputeRegionNetworkFirewallPolicy].
func (crnfp *ComputeRegionNetworkFirewallPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(crnfp)
}

// Dependencies returns the list of resources [ComputeRegionNetworkFirewallPolicy] depends_on.
func (crnfp *ComputeRegionNetworkFirewallPolicy) Dependencies() terra.Dependencies {
	return crnfp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeRegionNetworkFirewallPolicy].
func (crnfp *ComputeRegionNetworkFirewallPolicy) LifecycleManagement() *terra.Lifecycle {
	return crnfp.Lifecycle
}

// Attributes returns the attributes for [ComputeRegionNetworkFirewallPolicy].
func (crnfp *ComputeRegionNetworkFirewallPolicy) Attributes() computeRegionNetworkFirewallPolicyAttributes {
	return computeRegionNetworkFirewallPolicyAttributes{ref: terra.ReferenceResource(crnfp)}
}

// ImportState imports the given attribute values into [ComputeRegionNetworkFirewallPolicy]'s state.
func (crnfp *ComputeRegionNetworkFirewallPolicy) ImportState(av io.Reader) error {
	crnfp.state = &computeRegionNetworkFirewallPolicyState{}
	if err := json.NewDecoder(av).Decode(crnfp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crnfp.Type(), crnfp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeRegionNetworkFirewallPolicy] has state.
func (crnfp *ComputeRegionNetworkFirewallPolicy) State() (*computeRegionNetworkFirewallPolicyState, bool) {
	return crnfp.state, crnfp.state != nil
}

// StateMust returns the state for [ComputeRegionNetworkFirewallPolicy]. Panics if the state is nil.
func (crnfp *ComputeRegionNetworkFirewallPolicy) StateMust() *computeRegionNetworkFirewallPolicyState {
	if crnfp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crnfp.Type(), crnfp.LocalName()))
	}
	return crnfp.state
}

// ComputeRegionNetworkFirewallPolicyArgs contains the configurations for google_compute_region_network_firewall_policy.
type ComputeRegionNetworkFirewallPolicyArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Timeouts: optional
	Timeouts *computeregionnetworkfirewallpolicy.Timeouts `hcl:"timeouts,block"`
}
type computeRegionNetworkFirewallPolicyAttributes struct {
	ref terra.Reference
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_region_network_firewall_policy.
func (crnfp computeRegionNetworkFirewallPolicyAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(crnfp.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_region_network_firewall_policy.
func (crnfp computeRegionNetworkFirewallPolicyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(crnfp.ref.Append("description"))
}

// Fingerprint returns a reference to field fingerprint of google_compute_region_network_firewall_policy.
func (crnfp computeRegionNetworkFirewallPolicyAttributes) Fingerprint() terra.StringValue {
	return terra.ReferenceAsString(crnfp.ref.Append("fingerprint"))
}

// Id returns a reference to field id of google_compute_region_network_firewall_policy.
func (crnfp computeRegionNetworkFirewallPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crnfp.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_region_network_firewall_policy.
func (crnfp computeRegionNetworkFirewallPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crnfp.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_region_network_firewall_policy.
func (crnfp computeRegionNetworkFirewallPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crnfp.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_region_network_firewall_policy.
func (crnfp computeRegionNetworkFirewallPolicyAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(crnfp.ref.Append("region"))
}

// RegionNetworkFirewallPolicyId returns a reference to field region_network_firewall_policy_id of google_compute_region_network_firewall_policy.
func (crnfp computeRegionNetworkFirewallPolicyAttributes) RegionNetworkFirewallPolicyId() terra.StringValue {
	return terra.ReferenceAsString(crnfp.ref.Append("region_network_firewall_policy_id"))
}

// RuleTupleCount returns a reference to field rule_tuple_count of google_compute_region_network_firewall_policy.
func (crnfp computeRegionNetworkFirewallPolicyAttributes) RuleTupleCount() terra.NumberValue {
	return terra.ReferenceAsNumber(crnfp.ref.Append("rule_tuple_count"))
}

// SelfLink returns a reference to field self_link of google_compute_region_network_firewall_policy.
func (crnfp computeRegionNetworkFirewallPolicyAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(crnfp.ref.Append("self_link"))
}

// SelfLinkWithId returns a reference to field self_link_with_id of google_compute_region_network_firewall_policy.
func (crnfp computeRegionNetworkFirewallPolicyAttributes) SelfLinkWithId() terra.StringValue {
	return terra.ReferenceAsString(crnfp.ref.Append("self_link_with_id"))
}

func (crnfp computeRegionNetworkFirewallPolicyAttributes) Timeouts() computeregionnetworkfirewallpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeregionnetworkfirewallpolicy.TimeoutsAttributes](crnfp.ref.Append("timeouts"))
}

type computeRegionNetworkFirewallPolicyState struct {
	CreationTimestamp             string                                            `json:"creation_timestamp"`
	Description                   string                                            `json:"description"`
	Fingerprint                   string                                            `json:"fingerprint"`
	Id                            string                                            `json:"id"`
	Name                          string                                            `json:"name"`
	Project                       string                                            `json:"project"`
	Region                        string                                            `json:"region"`
	RegionNetworkFirewallPolicyId string                                            `json:"region_network_firewall_policy_id"`
	RuleTupleCount                float64                                           `json:"rule_tuple_count"`
	SelfLink                      string                                            `json:"self_link"`
	SelfLinkWithId                string                                            `json:"self_link_with_id"`
	Timeouts                      *computeregionnetworkfirewallpolicy.TimeoutsState `json:"timeouts"`
}
