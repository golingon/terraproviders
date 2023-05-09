// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computefirewallpolicyassociation "github.com/golingon/terraproviders/google/4.63.1/computefirewallpolicyassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeFirewallPolicyAssociation creates a new instance of [ComputeFirewallPolicyAssociation].
func NewComputeFirewallPolicyAssociation(name string, args ComputeFirewallPolicyAssociationArgs) *ComputeFirewallPolicyAssociation {
	return &ComputeFirewallPolicyAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeFirewallPolicyAssociation)(nil)

// ComputeFirewallPolicyAssociation represents the Terraform resource google_compute_firewall_policy_association.
type ComputeFirewallPolicyAssociation struct {
	Name      string
	Args      ComputeFirewallPolicyAssociationArgs
	state     *computeFirewallPolicyAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeFirewallPolicyAssociation].
func (cfpa *ComputeFirewallPolicyAssociation) Type() string {
	return "google_compute_firewall_policy_association"
}

// LocalName returns the local name for [ComputeFirewallPolicyAssociation].
func (cfpa *ComputeFirewallPolicyAssociation) LocalName() string {
	return cfpa.Name
}

// Configuration returns the configuration (args) for [ComputeFirewallPolicyAssociation].
func (cfpa *ComputeFirewallPolicyAssociation) Configuration() interface{} {
	return cfpa.Args
}

// DependOn is used for other resources to depend on [ComputeFirewallPolicyAssociation].
func (cfpa *ComputeFirewallPolicyAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(cfpa)
}

// Dependencies returns the list of resources [ComputeFirewallPolicyAssociation] depends_on.
func (cfpa *ComputeFirewallPolicyAssociation) Dependencies() terra.Dependencies {
	return cfpa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeFirewallPolicyAssociation].
func (cfpa *ComputeFirewallPolicyAssociation) LifecycleManagement() *terra.Lifecycle {
	return cfpa.Lifecycle
}

// Attributes returns the attributes for [ComputeFirewallPolicyAssociation].
func (cfpa *ComputeFirewallPolicyAssociation) Attributes() computeFirewallPolicyAssociationAttributes {
	return computeFirewallPolicyAssociationAttributes{ref: terra.ReferenceResource(cfpa)}
}

// ImportState imports the given attribute values into [ComputeFirewallPolicyAssociation]'s state.
func (cfpa *ComputeFirewallPolicyAssociation) ImportState(av io.Reader) error {
	cfpa.state = &computeFirewallPolicyAssociationState{}
	if err := json.NewDecoder(av).Decode(cfpa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cfpa.Type(), cfpa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeFirewallPolicyAssociation] has state.
func (cfpa *ComputeFirewallPolicyAssociation) State() (*computeFirewallPolicyAssociationState, bool) {
	return cfpa.state, cfpa.state != nil
}

// StateMust returns the state for [ComputeFirewallPolicyAssociation]. Panics if the state is nil.
func (cfpa *ComputeFirewallPolicyAssociation) StateMust() *computeFirewallPolicyAssociationState {
	if cfpa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cfpa.Type(), cfpa.LocalName()))
	}
	return cfpa.state
}

// ComputeFirewallPolicyAssociationArgs contains the configurations for google_compute_firewall_policy_association.
type ComputeFirewallPolicyAssociationArgs struct {
	// AttachmentTarget: string, required
	AttachmentTarget terra.StringValue `hcl:"attachment_target,attr" validate:"required"`
	// FirewallPolicy: string, required
	FirewallPolicy terra.StringValue `hcl:"firewall_policy,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *computefirewallpolicyassociation.Timeouts `hcl:"timeouts,block"`
}
type computeFirewallPolicyAssociationAttributes struct {
	ref terra.Reference
}

// AttachmentTarget returns a reference to field attachment_target of google_compute_firewall_policy_association.
func (cfpa computeFirewallPolicyAssociationAttributes) AttachmentTarget() terra.StringValue {
	return terra.ReferenceAsString(cfpa.ref.Append("attachment_target"))
}

// FirewallPolicy returns a reference to field firewall_policy of google_compute_firewall_policy_association.
func (cfpa computeFirewallPolicyAssociationAttributes) FirewallPolicy() terra.StringValue {
	return terra.ReferenceAsString(cfpa.ref.Append("firewall_policy"))
}

// Id returns a reference to field id of google_compute_firewall_policy_association.
func (cfpa computeFirewallPolicyAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cfpa.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_firewall_policy_association.
func (cfpa computeFirewallPolicyAssociationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cfpa.ref.Append("name"))
}

// ShortName returns a reference to field short_name of google_compute_firewall_policy_association.
func (cfpa computeFirewallPolicyAssociationAttributes) ShortName() terra.StringValue {
	return terra.ReferenceAsString(cfpa.ref.Append("short_name"))
}

func (cfpa computeFirewallPolicyAssociationAttributes) Timeouts() computefirewallpolicyassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computefirewallpolicyassociation.TimeoutsAttributes](cfpa.ref.Append("timeouts"))
}

type computeFirewallPolicyAssociationState struct {
	AttachmentTarget string                                          `json:"attachment_target"`
	FirewallPolicy   string                                          `json:"firewall_policy"`
	Id               string                                          `json:"id"`
	Name             string                                          `json:"name"`
	ShortName        string                                          `json:"short_name"`
	Timeouts         *computefirewallpolicyassociation.TimeoutsState `json:"timeouts"`
}
