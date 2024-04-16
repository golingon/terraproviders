// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_compute_network_firewall_policy

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

// Resource represents the Terraform resource google_compute_network_firewall_policy.
type Resource struct {
	Name      string
	Args      Args
	state     *googleComputeNetworkFirewallPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gcnfp *Resource) Type() string {
	return "google_compute_network_firewall_policy"
}

// LocalName returns the local name for [Resource].
func (gcnfp *Resource) LocalName() string {
	return gcnfp.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gcnfp *Resource) Configuration() interface{} {
	return gcnfp.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gcnfp *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gcnfp)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gcnfp *Resource) Dependencies() terra.Dependencies {
	return gcnfp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gcnfp *Resource) LifecycleManagement() *terra.Lifecycle {
	return gcnfp.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gcnfp *Resource) Attributes() googleComputeNetworkFirewallPolicyAttributes {
	return googleComputeNetworkFirewallPolicyAttributes{ref: terra.ReferenceResource(gcnfp)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gcnfp *Resource) ImportState(state io.Reader) error {
	gcnfp.state = &googleComputeNetworkFirewallPolicyState{}
	if err := json.NewDecoder(state).Decode(gcnfp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gcnfp.Type(), gcnfp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gcnfp *Resource) State() (*googleComputeNetworkFirewallPolicyState, bool) {
	return gcnfp.state, gcnfp.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gcnfp *Resource) StateMust() *googleComputeNetworkFirewallPolicyState {
	if gcnfp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gcnfp.Type(), gcnfp.LocalName()))
	}
	return gcnfp.state
}

// Args contains the configurations for google_compute_network_firewall_policy.
type Args struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleComputeNetworkFirewallPolicyAttributes struct {
	ref terra.Reference
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_network_firewall_policy.
func (gcnfp googleComputeNetworkFirewallPolicyAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(gcnfp.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_network_firewall_policy.
func (gcnfp googleComputeNetworkFirewallPolicyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gcnfp.ref.Append("description"))
}

// Fingerprint returns a reference to field fingerprint of google_compute_network_firewall_policy.
func (gcnfp googleComputeNetworkFirewallPolicyAttributes) Fingerprint() terra.StringValue {
	return terra.ReferenceAsString(gcnfp.ref.Append("fingerprint"))
}

// Id returns a reference to field id of google_compute_network_firewall_policy.
func (gcnfp googleComputeNetworkFirewallPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gcnfp.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_network_firewall_policy.
func (gcnfp googleComputeNetworkFirewallPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gcnfp.ref.Append("name"))
}

// NetworkFirewallPolicyId returns a reference to field network_firewall_policy_id of google_compute_network_firewall_policy.
func (gcnfp googleComputeNetworkFirewallPolicyAttributes) NetworkFirewallPolicyId() terra.StringValue {
	return terra.ReferenceAsString(gcnfp.ref.Append("network_firewall_policy_id"))
}

// Project returns a reference to field project of google_compute_network_firewall_policy.
func (gcnfp googleComputeNetworkFirewallPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gcnfp.ref.Append("project"))
}

// RuleTupleCount returns a reference to field rule_tuple_count of google_compute_network_firewall_policy.
func (gcnfp googleComputeNetworkFirewallPolicyAttributes) RuleTupleCount() terra.NumberValue {
	return terra.ReferenceAsNumber(gcnfp.ref.Append("rule_tuple_count"))
}

// SelfLink returns a reference to field self_link of google_compute_network_firewall_policy.
func (gcnfp googleComputeNetworkFirewallPolicyAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(gcnfp.ref.Append("self_link"))
}

// SelfLinkWithId returns a reference to field self_link_with_id of google_compute_network_firewall_policy.
func (gcnfp googleComputeNetworkFirewallPolicyAttributes) SelfLinkWithId() terra.StringValue {
	return terra.ReferenceAsString(gcnfp.ref.Append("self_link_with_id"))
}

func (gcnfp googleComputeNetworkFirewallPolicyAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gcnfp.ref.Append("timeouts"))
}

type googleComputeNetworkFirewallPolicyState struct {
	CreationTimestamp       string         `json:"creation_timestamp"`
	Description             string         `json:"description"`
	Fingerprint             string         `json:"fingerprint"`
	Id                      string         `json:"id"`
	Name                    string         `json:"name"`
	NetworkFirewallPolicyId string         `json:"network_firewall_policy_id"`
	Project                 string         `json:"project"`
	RuleTupleCount          float64        `json:"rule_tuple_count"`
	SelfLink                string         `json:"self_link"`
	SelfLinkWithId          string         `json:"self_link_with_id"`
	Timeouts                *TimeoutsState `json:"timeouts"`
}
