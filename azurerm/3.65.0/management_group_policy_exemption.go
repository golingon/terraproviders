// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	managementgrouppolicyexemption "github.com/golingon/terraproviders/azurerm/3.65.0/managementgrouppolicyexemption"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewManagementGroupPolicyExemption creates a new instance of [ManagementGroupPolicyExemption].
func NewManagementGroupPolicyExemption(name string, args ManagementGroupPolicyExemptionArgs) *ManagementGroupPolicyExemption {
	return &ManagementGroupPolicyExemption{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ManagementGroupPolicyExemption)(nil)

// ManagementGroupPolicyExemption represents the Terraform resource azurerm_management_group_policy_exemption.
type ManagementGroupPolicyExemption struct {
	Name      string
	Args      ManagementGroupPolicyExemptionArgs
	state     *managementGroupPolicyExemptionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ManagementGroupPolicyExemption].
func (mgpe *ManagementGroupPolicyExemption) Type() string {
	return "azurerm_management_group_policy_exemption"
}

// LocalName returns the local name for [ManagementGroupPolicyExemption].
func (mgpe *ManagementGroupPolicyExemption) LocalName() string {
	return mgpe.Name
}

// Configuration returns the configuration (args) for [ManagementGroupPolicyExemption].
func (mgpe *ManagementGroupPolicyExemption) Configuration() interface{} {
	return mgpe.Args
}

// DependOn is used for other resources to depend on [ManagementGroupPolicyExemption].
func (mgpe *ManagementGroupPolicyExemption) DependOn() terra.Reference {
	return terra.ReferenceResource(mgpe)
}

// Dependencies returns the list of resources [ManagementGroupPolicyExemption] depends_on.
func (mgpe *ManagementGroupPolicyExemption) Dependencies() terra.Dependencies {
	return mgpe.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ManagementGroupPolicyExemption].
func (mgpe *ManagementGroupPolicyExemption) LifecycleManagement() *terra.Lifecycle {
	return mgpe.Lifecycle
}

// Attributes returns the attributes for [ManagementGroupPolicyExemption].
func (mgpe *ManagementGroupPolicyExemption) Attributes() managementGroupPolicyExemptionAttributes {
	return managementGroupPolicyExemptionAttributes{ref: terra.ReferenceResource(mgpe)}
}

// ImportState imports the given attribute values into [ManagementGroupPolicyExemption]'s state.
func (mgpe *ManagementGroupPolicyExemption) ImportState(av io.Reader) error {
	mgpe.state = &managementGroupPolicyExemptionState{}
	if err := json.NewDecoder(av).Decode(mgpe.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mgpe.Type(), mgpe.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ManagementGroupPolicyExemption] has state.
func (mgpe *ManagementGroupPolicyExemption) State() (*managementGroupPolicyExemptionState, bool) {
	return mgpe.state, mgpe.state != nil
}

// StateMust returns the state for [ManagementGroupPolicyExemption]. Panics if the state is nil.
func (mgpe *ManagementGroupPolicyExemption) StateMust() *managementGroupPolicyExemptionState {
	if mgpe.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mgpe.Type(), mgpe.LocalName()))
	}
	return mgpe.state
}

// ManagementGroupPolicyExemptionArgs contains the configurations for azurerm_management_group_policy_exemption.
type ManagementGroupPolicyExemptionArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// ExemptionCategory: string, required
	ExemptionCategory terra.StringValue `hcl:"exemption_category,attr" validate:"required"`
	// ExpiresOn: string, optional
	ExpiresOn terra.StringValue `hcl:"expires_on,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ManagementGroupId: string, required
	ManagementGroupId terra.StringValue `hcl:"management_group_id,attr" validate:"required"`
	// Metadata: string, optional
	Metadata terra.StringValue `hcl:"metadata,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PolicyAssignmentId: string, required
	PolicyAssignmentId terra.StringValue `hcl:"policy_assignment_id,attr" validate:"required"`
	// PolicyDefinitionReferenceIds: list of string, optional
	PolicyDefinitionReferenceIds terra.ListValue[terra.StringValue] `hcl:"policy_definition_reference_ids,attr"`
	// Timeouts: optional
	Timeouts *managementgrouppolicyexemption.Timeouts `hcl:"timeouts,block"`
}
type managementGroupPolicyExemptionAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_management_group_policy_exemption.
func (mgpe managementGroupPolicyExemptionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(mgpe.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_management_group_policy_exemption.
func (mgpe managementGroupPolicyExemptionAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(mgpe.ref.Append("display_name"))
}

// ExemptionCategory returns a reference to field exemption_category of azurerm_management_group_policy_exemption.
func (mgpe managementGroupPolicyExemptionAttributes) ExemptionCategory() terra.StringValue {
	return terra.ReferenceAsString(mgpe.ref.Append("exemption_category"))
}

// ExpiresOn returns a reference to field expires_on of azurerm_management_group_policy_exemption.
func (mgpe managementGroupPolicyExemptionAttributes) ExpiresOn() terra.StringValue {
	return terra.ReferenceAsString(mgpe.ref.Append("expires_on"))
}

// Id returns a reference to field id of azurerm_management_group_policy_exemption.
func (mgpe managementGroupPolicyExemptionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mgpe.ref.Append("id"))
}

// ManagementGroupId returns a reference to field management_group_id of azurerm_management_group_policy_exemption.
func (mgpe managementGroupPolicyExemptionAttributes) ManagementGroupId() terra.StringValue {
	return terra.ReferenceAsString(mgpe.ref.Append("management_group_id"))
}

// Metadata returns a reference to field metadata of azurerm_management_group_policy_exemption.
func (mgpe managementGroupPolicyExemptionAttributes) Metadata() terra.StringValue {
	return terra.ReferenceAsString(mgpe.ref.Append("metadata"))
}

// Name returns a reference to field name of azurerm_management_group_policy_exemption.
func (mgpe managementGroupPolicyExemptionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mgpe.ref.Append("name"))
}

// PolicyAssignmentId returns a reference to field policy_assignment_id of azurerm_management_group_policy_exemption.
func (mgpe managementGroupPolicyExemptionAttributes) PolicyAssignmentId() terra.StringValue {
	return terra.ReferenceAsString(mgpe.ref.Append("policy_assignment_id"))
}

// PolicyDefinitionReferenceIds returns a reference to field policy_definition_reference_ids of azurerm_management_group_policy_exemption.
func (mgpe managementGroupPolicyExemptionAttributes) PolicyDefinitionReferenceIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](mgpe.ref.Append("policy_definition_reference_ids"))
}

func (mgpe managementGroupPolicyExemptionAttributes) Timeouts() managementgrouppolicyexemption.TimeoutsAttributes {
	return terra.ReferenceAsSingle[managementgrouppolicyexemption.TimeoutsAttributes](mgpe.ref.Append("timeouts"))
}

type managementGroupPolicyExemptionState struct {
	Description                  string                                        `json:"description"`
	DisplayName                  string                                        `json:"display_name"`
	ExemptionCategory            string                                        `json:"exemption_category"`
	ExpiresOn                    string                                        `json:"expires_on"`
	Id                           string                                        `json:"id"`
	ManagementGroupId            string                                        `json:"management_group_id"`
	Metadata                     string                                        `json:"metadata"`
	Name                         string                                        `json:"name"`
	PolicyAssignmentId           string                                        `json:"policy_assignment_id"`
	PolicyDefinitionReferenceIds []string                                      `json:"policy_definition_reference_ids"`
	Timeouts                     *managementgrouppolicyexemption.TimeoutsState `json:"timeouts"`
}
