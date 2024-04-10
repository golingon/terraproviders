// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	computenetworkfirewallpolicyassociation "github.com/golingon/terraproviders/googlebeta/5.24.0/computenetworkfirewallpolicyassociation"
	"io"
)

// NewComputeNetworkFirewallPolicyAssociation creates a new instance of [ComputeNetworkFirewallPolicyAssociation].
func NewComputeNetworkFirewallPolicyAssociation(name string, args ComputeNetworkFirewallPolicyAssociationArgs) *ComputeNetworkFirewallPolicyAssociation {
	return &ComputeNetworkFirewallPolicyAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeNetworkFirewallPolicyAssociation)(nil)

// ComputeNetworkFirewallPolicyAssociation represents the Terraform resource google_compute_network_firewall_policy_association.
type ComputeNetworkFirewallPolicyAssociation struct {
	Name      string
	Args      ComputeNetworkFirewallPolicyAssociationArgs
	state     *computeNetworkFirewallPolicyAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeNetworkFirewallPolicyAssociation].
func (cnfpa *ComputeNetworkFirewallPolicyAssociation) Type() string {
	return "google_compute_network_firewall_policy_association"
}

// LocalName returns the local name for [ComputeNetworkFirewallPolicyAssociation].
func (cnfpa *ComputeNetworkFirewallPolicyAssociation) LocalName() string {
	return cnfpa.Name
}

// Configuration returns the configuration (args) for [ComputeNetworkFirewallPolicyAssociation].
func (cnfpa *ComputeNetworkFirewallPolicyAssociation) Configuration() interface{} {
	return cnfpa.Args
}

// DependOn is used for other resources to depend on [ComputeNetworkFirewallPolicyAssociation].
func (cnfpa *ComputeNetworkFirewallPolicyAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(cnfpa)
}

// Dependencies returns the list of resources [ComputeNetworkFirewallPolicyAssociation] depends_on.
func (cnfpa *ComputeNetworkFirewallPolicyAssociation) Dependencies() terra.Dependencies {
	return cnfpa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeNetworkFirewallPolicyAssociation].
func (cnfpa *ComputeNetworkFirewallPolicyAssociation) LifecycleManagement() *terra.Lifecycle {
	return cnfpa.Lifecycle
}

// Attributes returns the attributes for [ComputeNetworkFirewallPolicyAssociation].
func (cnfpa *ComputeNetworkFirewallPolicyAssociation) Attributes() computeNetworkFirewallPolicyAssociationAttributes {
	return computeNetworkFirewallPolicyAssociationAttributes{ref: terra.ReferenceResource(cnfpa)}
}

// ImportState imports the given attribute values into [ComputeNetworkFirewallPolicyAssociation]'s state.
func (cnfpa *ComputeNetworkFirewallPolicyAssociation) ImportState(av io.Reader) error {
	cnfpa.state = &computeNetworkFirewallPolicyAssociationState{}
	if err := json.NewDecoder(av).Decode(cnfpa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cnfpa.Type(), cnfpa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeNetworkFirewallPolicyAssociation] has state.
func (cnfpa *ComputeNetworkFirewallPolicyAssociation) State() (*computeNetworkFirewallPolicyAssociationState, bool) {
	return cnfpa.state, cnfpa.state != nil
}

// StateMust returns the state for [ComputeNetworkFirewallPolicyAssociation]. Panics if the state is nil.
func (cnfpa *ComputeNetworkFirewallPolicyAssociation) StateMust() *computeNetworkFirewallPolicyAssociationState {
	if cnfpa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cnfpa.Type(), cnfpa.LocalName()))
	}
	return cnfpa.state
}

// ComputeNetworkFirewallPolicyAssociationArgs contains the configurations for google_compute_network_firewall_policy_association.
type ComputeNetworkFirewallPolicyAssociationArgs struct {
	// AttachmentTarget: string, required
	AttachmentTarget terra.StringValue `hcl:"attachment_target,attr" validate:"required"`
	// FirewallPolicy: string, required
	FirewallPolicy terra.StringValue `hcl:"firewall_policy,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *computenetworkfirewallpolicyassociation.Timeouts `hcl:"timeouts,block"`
}
type computeNetworkFirewallPolicyAssociationAttributes struct {
	ref terra.Reference
}

// AttachmentTarget returns a reference to field attachment_target of google_compute_network_firewall_policy_association.
func (cnfpa computeNetworkFirewallPolicyAssociationAttributes) AttachmentTarget() terra.StringValue {
	return terra.ReferenceAsString(cnfpa.ref.Append("attachment_target"))
}

// FirewallPolicy returns a reference to field firewall_policy of google_compute_network_firewall_policy_association.
func (cnfpa computeNetworkFirewallPolicyAssociationAttributes) FirewallPolicy() terra.StringValue {
	return terra.ReferenceAsString(cnfpa.ref.Append("firewall_policy"))
}

// Id returns a reference to field id of google_compute_network_firewall_policy_association.
func (cnfpa computeNetworkFirewallPolicyAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cnfpa.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_network_firewall_policy_association.
func (cnfpa computeNetworkFirewallPolicyAssociationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cnfpa.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_network_firewall_policy_association.
func (cnfpa computeNetworkFirewallPolicyAssociationAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cnfpa.ref.Append("project"))
}

// ShortName returns a reference to field short_name of google_compute_network_firewall_policy_association.
func (cnfpa computeNetworkFirewallPolicyAssociationAttributes) ShortName() terra.StringValue {
	return terra.ReferenceAsString(cnfpa.ref.Append("short_name"))
}

func (cnfpa computeNetworkFirewallPolicyAssociationAttributes) Timeouts() computenetworkfirewallpolicyassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computenetworkfirewallpolicyassociation.TimeoutsAttributes](cnfpa.ref.Append("timeouts"))
}

type computeNetworkFirewallPolicyAssociationState struct {
	AttachmentTarget string                                                 `json:"attachment_target"`
	FirewallPolicy   string                                                 `json:"firewall_policy"`
	Id               string                                                 `json:"id"`
	Name             string                                                 `json:"name"`
	Project          string                                                 `json:"project"`
	ShortName        string                                                 `json:"short_name"`
	Timeouts         *computenetworkfirewallpolicyassociation.TimeoutsState `json:"timeouts"`
}
