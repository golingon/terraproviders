// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computefirewallpolicy "github.com/golingon/terraproviders/googlebeta/4.72.1/computefirewallpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeFirewallPolicy creates a new instance of [ComputeFirewallPolicy].
func NewComputeFirewallPolicy(name string, args ComputeFirewallPolicyArgs) *ComputeFirewallPolicy {
	return &ComputeFirewallPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeFirewallPolicy)(nil)

// ComputeFirewallPolicy represents the Terraform resource google_compute_firewall_policy.
type ComputeFirewallPolicy struct {
	Name      string
	Args      ComputeFirewallPolicyArgs
	state     *computeFirewallPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeFirewallPolicy].
func (cfp *ComputeFirewallPolicy) Type() string {
	return "google_compute_firewall_policy"
}

// LocalName returns the local name for [ComputeFirewallPolicy].
func (cfp *ComputeFirewallPolicy) LocalName() string {
	return cfp.Name
}

// Configuration returns the configuration (args) for [ComputeFirewallPolicy].
func (cfp *ComputeFirewallPolicy) Configuration() interface{} {
	return cfp.Args
}

// DependOn is used for other resources to depend on [ComputeFirewallPolicy].
func (cfp *ComputeFirewallPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(cfp)
}

// Dependencies returns the list of resources [ComputeFirewallPolicy] depends_on.
func (cfp *ComputeFirewallPolicy) Dependencies() terra.Dependencies {
	return cfp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeFirewallPolicy].
func (cfp *ComputeFirewallPolicy) LifecycleManagement() *terra.Lifecycle {
	return cfp.Lifecycle
}

// Attributes returns the attributes for [ComputeFirewallPolicy].
func (cfp *ComputeFirewallPolicy) Attributes() computeFirewallPolicyAttributes {
	return computeFirewallPolicyAttributes{ref: terra.ReferenceResource(cfp)}
}

// ImportState imports the given attribute values into [ComputeFirewallPolicy]'s state.
func (cfp *ComputeFirewallPolicy) ImportState(av io.Reader) error {
	cfp.state = &computeFirewallPolicyState{}
	if err := json.NewDecoder(av).Decode(cfp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cfp.Type(), cfp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeFirewallPolicy] has state.
func (cfp *ComputeFirewallPolicy) State() (*computeFirewallPolicyState, bool) {
	return cfp.state, cfp.state != nil
}

// StateMust returns the state for [ComputeFirewallPolicy]. Panics if the state is nil.
func (cfp *ComputeFirewallPolicy) StateMust() *computeFirewallPolicyState {
	if cfp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cfp.Type(), cfp.LocalName()))
	}
	return cfp.state
}

// ComputeFirewallPolicyArgs contains the configurations for google_compute_firewall_policy.
type ComputeFirewallPolicyArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Parent: string, required
	Parent terra.StringValue `hcl:"parent,attr" validate:"required"`
	// ShortName: string, required
	ShortName terra.StringValue `hcl:"short_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *computefirewallpolicy.Timeouts `hcl:"timeouts,block"`
}
type computeFirewallPolicyAttributes struct {
	ref terra.Reference
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_firewall_policy.
func (cfp computeFirewallPolicyAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(cfp.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_firewall_policy.
func (cfp computeFirewallPolicyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cfp.ref.Append("description"))
}

// Fingerprint returns a reference to field fingerprint of google_compute_firewall_policy.
func (cfp computeFirewallPolicyAttributes) Fingerprint() terra.StringValue {
	return terra.ReferenceAsString(cfp.ref.Append("fingerprint"))
}

// FirewallPolicyId returns a reference to field firewall_policy_id of google_compute_firewall_policy.
func (cfp computeFirewallPolicyAttributes) FirewallPolicyId() terra.StringValue {
	return terra.ReferenceAsString(cfp.ref.Append("firewall_policy_id"))
}

// Id returns a reference to field id of google_compute_firewall_policy.
func (cfp computeFirewallPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cfp.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_firewall_policy.
func (cfp computeFirewallPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cfp.ref.Append("name"))
}

// Parent returns a reference to field parent of google_compute_firewall_policy.
func (cfp computeFirewallPolicyAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(cfp.ref.Append("parent"))
}

// RuleTupleCount returns a reference to field rule_tuple_count of google_compute_firewall_policy.
func (cfp computeFirewallPolicyAttributes) RuleTupleCount() terra.NumberValue {
	return terra.ReferenceAsNumber(cfp.ref.Append("rule_tuple_count"))
}

// SelfLink returns a reference to field self_link of google_compute_firewall_policy.
func (cfp computeFirewallPolicyAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cfp.ref.Append("self_link"))
}

// SelfLinkWithId returns a reference to field self_link_with_id of google_compute_firewall_policy.
func (cfp computeFirewallPolicyAttributes) SelfLinkWithId() terra.StringValue {
	return terra.ReferenceAsString(cfp.ref.Append("self_link_with_id"))
}

// ShortName returns a reference to field short_name of google_compute_firewall_policy.
func (cfp computeFirewallPolicyAttributes) ShortName() terra.StringValue {
	return terra.ReferenceAsString(cfp.ref.Append("short_name"))
}

func (cfp computeFirewallPolicyAttributes) Timeouts() computefirewallpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computefirewallpolicy.TimeoutsAttributes](cfp.ref.Append("timeouts"))
}

type computeFirewallPolicyState struct {
	CreationTimestamp string                               `json:"creation_timestamp"`
	Description       string                               `json:"description"`
	Fingerprint       string                               `json:"fingerprint"`
	FirewallPolicyId  string                               `json:"firewall_policy_id"`
	Id                string                               `json:"id"`
	Name              string                               `json:"name"`
	Parent            string                               `json:"parent"`
	RuleTupleCount    float64                              `json:"rule_tuple_count"`
	SelfLink          string                               `json:"self_link"`
	SelfLinkWithId    string                               `json:"self_link_with_id"`
	ShortName         string                               `json:"short_name"`
	Timeouts          *computefirewallpolicy.TimeoutsState `json:"timeouts"`
}
