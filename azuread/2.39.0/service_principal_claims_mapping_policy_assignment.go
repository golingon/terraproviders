// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azuread

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewServicePrincipalClaimsMappingPolicyAssignment creates a new instance of [ServicePrincipalClaimsMappingPolicyAssignment].
func NewServicePrincipalClaimsMappingPolicyAssignment(name string, args ServicePrincipalClaimsMappingPolicyAssignmentArgs) *ServicePrincipalClaimsMappingPolicyAssignment {
	return &ServicePrincipalClaimsMappingPolicyAssignment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServicePrincipalClaimsMappingPolicyAssignment)(nil)

// ServicePrincipalClaimsMappingPolicyAssignment represents the Terraform resource azuread_service_principal_claims_mapping_policy_assignment.
type ServicePrincipalClaimsMappingPolicyAssignment struct {
	Name      string
	Args      ServicePrincipalClaimsMappingPolicyAssignmentArgs
	state     *servicePrincipalClaimsMappingPolicyAssignmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ServicePrincipalClaimsMappingPolicyAssignment].
func (spcmpa *ServicePrincipalClaimsMappingPolicyAssignment) Type() string {
	return "azuread_service_principal_claims_mapping_policy_assignment"
}

// LocalName returns the local name for [ServicePrincipalClaimsMappingPolicyAssignment].
func (spcmpa *ServicePrincipalClaimsMappingPolicyAssignment) LocalName() string {
	return spcmpa.Name
}

// Configuration returns the configuration (args) for [ServicePrincipalClaimsMappingPolicyAssignment].
func (spcmpa *ServicePrincipalClaimsMappingPolicyAssignment) Configuration() interface{} {
	return spcmpa.Args
}

// DependOn is used for other resources to depend on [ServicePrincipalClaimsMappingPolicyAssignment].
func (spcmpa *ServicePrincipalClaimsMappingPolicyAssignment) DependOn() terra.Reference {
	return terra.ReferenceResource(spcmpa)
}

// Dependencies returns the list of resources [ServicePrincipalClaimsMappingPolicyAssignment] depends_on.
func (spcmpa *ServicePrincipalClaimsMappingPolicyAssignment) Dependencies() terra.Dependencies {
	return spcmpa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ServicePrincipalClaimsMappingPolicyAssignment].
func (spcmpa *ServicePrincipalClaimsMappingPolicyAssignment) LifecycleManagement() *terra.Lifecycle {
	return spcmpa.Lifecycle
}

// Attributes returns the attributes for [ServicePrincipalClaimsMappingPolicyAssignment].
func (spcmpa *ServicePrincipalClaimsMappingPolicyAssignment) Attributes() servicePrincipalClaimsMappingPolicyAssignmentAttributes {
	return servicePrincipalClaimsMappingPolicyAssignmentAttributes{ref: terra.ReferenceResource(spcmpa)}
}

// ImportState imports the given attribute values into [ServicePrincipalClaimsMappingPolicyAssignment]'s state.
func (spcmpa *ServicePrincipalClaimsMappingPolicyAssignment) ImportState(av io.Reader) error {
	spcmpa.state = &servicePrincipalClaimsMappingPolicyAssignmentState{}
	if err := json.NewDecoder(av).Decode(spcmpa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", spcmpa.Type(), spcmpa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ServicePrincipalClaimsMappingPolicyAssignment] has state.
func (spcmpa *ServicePrincipalClaimsMappingPolicyAssignment) State() (*servicePrincipalClaimsMappingPolicyAssignmentState, bool) {
	return spcmpa.state, spcmpa.state != nil
}

// StateMust returns the state for [ServicePrincipalClaimsMappingPolicyAssignment]. Panics if the state is nil.
func (spcmpa *ServicePrincipalClaimsMappingPolicyAssignment) StateMust() *servicePrincipalClaimsMappingPolicyAssignmentState {
	if spcmpa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", spcmpa.Type(), spcmpa.LocalName()))
	}
	return spcmpa.state
}

// ServicePrincipalClaimsMappingPolicyAssignmentArgs contains the configurations for azuread_service_principal_claims_mapping_policy_assignment.
type ServicePrincipalClaimsMappingPolicyAssignmentArgs struct {
	// ClaimsMappingPolicyId: string, required
	ClaimsMappingPolicyId terra.StringValue `hcl:"claims_mapping_policy_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ServicePrincipalId: string, required
	ServicePrincipalId terra.StringValue `hcl:"service_principal_id,attr" validate:"required"`
}
type servicePrincipalClaimsMappingPolicyAssignmentAttributes struct {
	ref terra.Reference
}

// ClaimsMappingPolicyId returns a reference to field claims_mapping_policy_id of azuread_service_principal_claims_mapping_policy_assignment.
func (spcmpa servicePrincipalClaimsMappingPolicyAssignmentAttributes) ClaimsMappingPolicyId() terra.StringValue {
	return terra.ReferenceAsString(spcmpa.ref.Append("claims_mapping_policy_id"))
}

// Id returns a reference to field id of azuread_service_principal_claims_mapping_policy_assignment.
func (spcmpa servicePrincipalClaimsMappingPolicyAssignmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(spcmpa.ref.Append("id"))
}

// ServicePrincipalId returns a reference to field service_principal_id of azuread_service_principal_claims_mapping_policy_assignment.
func (spcmpa servicePrincipalClaimsMappingPolicyAssignmentAttributes) ServicePrincipalId() terra.StringValue {
	return terra.ReferenceAsString(spcmpa.ref.Append("service_principal_id"))
}

type servicePrincipalClaimsMappingPolicyAssignmentState struct {
	ClaimsMappingPolicyId string `json:"claims_mapping_policy_id"`
	Id                    string `json:"id"`
	ServicePrincipalId    string `json:"service_principal_id"`
}
