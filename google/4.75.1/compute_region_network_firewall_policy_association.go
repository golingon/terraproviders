// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computeregionnetworkfirewallpolicyassociation "github.com/golingon/terraproviders/google/4.75.1/computeregionnetworkfirewallpolicyassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeRegionNetworkFirewallPolicyAssociation creates a new instance of [ComputeRegionNetworkFirewallPolicyAssociation].
func NewComputeRegionNetworkFirewallPolicyAssociation(name string, args ComputeRegionNetworkFirewallPolicyAssociationArgs) *ComputeRegionNetworkFirewallPolicyAssociation {
	return &ComputeRegionNetworkFirewallPolicyAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeRegionNetworkFirewallPolicyAssociation)(nil)

// ComputeRegionNetworkFirewallPolicyAssociation represents the Terraform resource google_compute_region_network_firewall_policy_association.
type ComputeRegionNetworkFirewallPolicyAssociation struct {
	Name      string
	Args      ComputeRegionNetworkFirewallPolicyAssociationArgs
	state     *computeRegionNetworkFirewallPolicyAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeRegionNetworkFirewallPolicyAssociation].
func (crnfpa *ComputeRegionNetworkFirewallPolicyAssociation) Type() string {
	return "google_compute_region_network_firewall_policy_association"
}

// LocalName returns the local name for [ComputeRegionNetworkFirewallPolicyAssociation].
func (crnfpa *ComputeRegionNetworkFirewallPolicyAssociation) LocalName() string {
	return crnfpa.Name
}

// Configuration returns the configuration (args) for [ComputeRegionNetworkFirewallPolicyAssociation].
func (crnfpa *ComputeRegionNetworkFirewallPolicyAssociation) Configuration() interface{} {
	return crnfpa.Args
}

// DependOn is used for other resources to depend on [ComputeRegionNetworkFirewallPolicyAssociation].
func (crnfpa *ComputeRegionNetworkFirewallPolicyAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(crnfpa)
}

// Dependencies returns the list of resources [ComputeRegionNetworkFirewallPolicyAssociation] depends_on.
func (crnfpa *ComputeRegionNetworkFirewallPolicyAssociation) Dependencies() terra.Dependencies {
	return crnfpa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeRegionNetworkFirewallPolicyAssociation].
func (crnfpa *ComputeRegionNetworkFirewallPolicyAssociation) LifecycleManagement() *terra.Lifecycle {
	return crnfpa.Lifecycle
}

// Attributes returns the attributes for [ComputeRegionNetworkFirewallPolicyAssociation].
func (crnfpa *ComputeRegionNetworkFirewallPolicyAssociation) Attributes() computeRegionNetworkFirewallPolicyAssociationAttributes {
	return computeRegionNetworkFirewallPolicyAssociationAttributes{ref: terra.ReferenceResource(crnfpa)}
}

// ImportState imports the given attribute values into [ComputeRegionNetworkFirewallPolicyAssociation]'s state.
func (crnfpa *ComputeRegionNetworkFirewallPolicyAssociation) ImportState(av io.Reader) error {
	crnfpa.state = &computeRegionNetworkFirewallPolicyAssociationState{}
	if err := json.NewDecoder(av).Decode(crnfpa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crnfpa.Type(), crnfpa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeRegionNetworkFirewallPolicyAssociation] has state.
func (crnfpa *ComputeRegionNetworkFirewallPolicyAssociation) State() (*computeRegionNetworkFirewallPolicyAssociationState, bool) {
	return crnfpa.state, crnfpa.state != nil
}

// StateMust returns the state for [ComputeRegionNetworkFirewallPolicyAssociation]. Panics if the state is nil.
func (crnfpa *ComputeRegionNetworkFirewallPolicyAssociation) StateMust() *computeRegionNetworkFirewallPolicyAssociationState {
	if crnfpa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crnfpa.Type(), crnfpa.LocalName()))
	}
	return crnfpa.state
}

// ComputeRegionNetworkFirewallPolicyAssociationArgs contains the configurations for google_compute_region_network_firewall_policy_association.
type ComputeRegionNetworkFirewallPolicyAssociationArgs struct {
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
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Timeouts: optional
	Timeouts *computeregionnetworkfirewallpolicyassociation.Timeouts `hcl:"timeouts,block"`
}
type computeRegionNetworkFirewallPolicyAssociationAttributes struct {
	ref terra.Reference
}

// AttachmentTarget returns a reference to field attachment_target of google_compute_region_network_firewall_policy_association.
func (crnfpa computeRegionNetworkFirewallPolicyAssociationAttributes) AttachmentTarget() terra.StringValue {
	return terra.ReferenceAsString(crnfpa.ref.Append("attachment_target"))
}

// FirewallPolicy returns a reference to field firewall_policy of google_compute_region_network_firewall_policy_association.
func (crnfpa computeRegionNetworkFirewallPolicyAssociationAttributes) FirewallPolicy() terra.StringValue {
	return terra.ReferenceAsString(crnfpa.ref.Append("firewall_policy"))
}

// Id returns a reference to field id of google_compute_region_network_firewall_policy_association.
func (crnfpa computeRegionNetworkFirewallPolicyAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crnfpa.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_region_network_firewall_policy_association.
func (crnfpa computeRegionNetworkFirewallPolicyAssociationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crnfpa.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_region_network_firewall_policy_association.
func (crnfpa computeRegionNetworkFirewallPolicyAssociationAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crnfpa.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_region_network_firewall_policy_association.
func (crnfpa computeRegionNetworkFirewallPolicyAssociationAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(crnfpa.ref.Append("region"))
}

// ShortName returns a reference to field short_name of google_compute_region_network_firewall_policy_association.
func (crnfpa computeRegionNetworkFirewallPolicyAssociationAttributes) ShortName() terra.StringValue {
	return terra.ReferenceAsString(crnfpa.ref.Append("short_name"))
}

func (crnfpa computeRegionNetworkFirewallPolicyAssociationAttributes) Timeouts() computeregionnetworkfirewallpolicyassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeregionnetworkfirewallpolicyassociation.TimeoutsAttributes](crnfpa.ref.Append("timeouts"))
}

type computeRegionNetworkFirewallPolicyAssociationState struct {
	AttachmentTarget string                                                       `json:"attachment_target"`
	FirewallPolicy   string                                                       `json:"firewall_policy"`
	Id               string                                                       `json:"id"`
	Name             string                                                       `json:"name"`
	Project          string                                                       `json:"project"`
	Region           string                                                       `json:"region"`
	ShortName        string                                                       `json:"short_name"`
	Timeouts         *computeregionnetworkfirewallpolicyassociation.TimeoutsState `json:"timeouts"`
}