// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	computeorganizationsecuritypolicyassociation "github.com/golingon/terraproviders/googlebeta/5.24.0/computeorganizationsecuritypolicyassociation"
	"io"
)

// NewComputeOrganizationSecurityPolicyAssociation creates a new instance of [ComputeOrganizationSecurityPolicyAssociation].
func NewComputeOrganizationSecurityPolicyAssociation(name string, args ComputeOrganizationSecurityPolicyAssociationArgs) *ComputeOrganizationSecurityPolicyAssociation {
	return &ComputeOrganizationSecurityPolicyAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeOrganizationSecurityPolicyAssociation)(nil)

// ComputeOrganizationSecurityPolicyAssociation represents the Terraform resource google_compute_organization_security_policy_association.
type ComputeOrganizationSecurityPolicyAssociation struct {
	Name      string
	Args      ComputeOrganizationSecurityPolicyAssociationArgs
	state     *computeOrganizationSecurityPolicyAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeOrganizationSecurityPolicyAssociation].
func (cospa *ComputeOrganizationSecurityPolicyAssociation) Type() string {
	return "google_compute_organization_security_policy_association"
}

// LocalName returns the local name for [ComputeOrganizationSecurityPolicyAssociation].
func (cospa *ComputeOrganizationSecurityPolicyAssociation) LocalName() string {
	return cospa.Name
}

// Configuration returns the configuration (args) for [ComputeOrganizationSecurityPolicyAssociation].
func (cospa *ComputeOrganizationSecurityPolicyAssociation) Configuration() interface{} {
	return cospa.Args
}

// DependOn is used for other resources to depend on [ComputeOrganizationSecurityPolicyAssociation].
func (cospa *ComputeOrganizationSecurityPolicyAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(cospa)
}

// Dependencies returns the list of resources [ComputeOrganizationSecurityPolicyAssociation] depends_on.
func (cospa *ComputeOrganizationSecurityPolicyAssociation) Dependencies() terra.Dependencies {
	return cospa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeOrganizationSecurityPolicyAssociation].
func (cospa *ComputeOrganizationSecurityPolicyAssociation) LifecycleManagement() *terra.Lifecycle {
	return cospa.Lifecycle
}

// Attributes returns the attributes for [ComputeOrganizationSecurityPolicyAssociation].
func (cospa *ComputeOrganizationSecurityPolicyAssociation) Attributes() computeOrganizationSecurityPolicyAssociationAttributes {
	return computeOrganizationSecurityPolicyAssociationAttributes{ref: terra.ReferenceResource(cospa)}
}

// ImportState imports the given attribute values into [ComputeOrganizationSecurityPolicyAssociation]'s state.
func (cospa *ComputeOrganizationSecurityPolicyAssociation) ImportState(av io.Reader) error {
	cospa.state = &computeOrganizationSecurityPolicyAssociationState{}
	if err := json.NewDecoder(av).Decode(cospa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cospa.Type(), cospa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeOrganizationSecurityPolicyAssociation] has state.
func (cospa *ComputeOrganizationSecurityPolicyAssociation) State() (*computeOrganizationSecurityPolicyAssociationState, bool) {
	return cospa.state, cospa.state != nil
}

// StateMust returns the state for [ComputeOrganizationSecurityPolicyAssociation]. Panics if the state is nil.
func (cospa *ComputeOrganizationSecurityPolicyAssociation) StateMust() *computeOrganizationSecurityPolicyAssociationState {
	if cospa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cospa.Type(), cospa.LocalName()))
	}
	return cospa.state
}

// ComputeOrganizationSecurityPolicyAssociationArgs contains the configurations for google_compute_organization_security_policy_association.
type ComputeOrganizationSecurityPolicyAssociationArgs struct {
	// AttachmentId: string, required
	AttachmentId terra.StringValue `hcl:"attachment_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PolicyId: string, required
	PolicyId terra.StringValue `hcl:"policy_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *computeorganizationsecuritypolicyassociation.Timeouts `hcl:"timeouts,block"`
}
type computeOrganizationSecurityPolicyAssociationAttributes struct {
	ref terra.Reference
}

// AttachmentId returns a reference to field attachment_id of google_compute_organization_security_policy_association.
func (cospa computeOrganizationSecurityPolicyAssociationAttributes) AttachmentId() terra.StringValue {
	return terra.ReferenceAsString(cospa.ref.Append("attachment_id"))
}

// DisplayName returns a reference to field display_name of google_compute_organization_security_policy_association.
func (cospa computeOrganizationSecurityPolicyAssociationAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(cospa.ref.Append("display_name"))
}

// Id returns a reference to field id of google_compute_organization_security_policy_association.
func (cospa computeOrganizationSecurityPolicyAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cospa.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_organization_security_policy_association.
func (cospa computeOrganizationSecurityPolicyAssociationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cospa.ref.Append("name"))
}

// PolicyId returns a reference to field policy_id of google_compute_organization_security_policy_association.
func (cospa computeOrganizationSecurityPolicyAssociationAttributes) PolicyId() terra.StringValue {
	return terra.ReferenceAsString(cospa.ref.Append("policy_id"))
}

func (cospa computeOrganizationSecurityPolicyAssociationAttributes) Timeouts() computeorganizationsecuritypolicyassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeorganizationsecuritypolicyassociation.TimeoutsAttributes](cospa.ref.Append("timeouts"))
}

type computeOrganizationSecurityPolicyAssociationState struct {
	AttachmentId string                                                      `json:"attachment_id"`
	DisplayName  string                                                      `json:"display_name"`
	Id           string                                                      `json:"id"`
	Name         string                                                      `json:"name"`
	PolicyId     string                                                      `json:"policy_id"`
	Timeouts     *computeorganizationsecuritypolicyassociation.TimeoutsState `json:"timeouts"`
}
