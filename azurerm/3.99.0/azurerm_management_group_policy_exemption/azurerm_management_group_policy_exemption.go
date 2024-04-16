// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_management_group_policy_exemption

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

// Resource represents the Terraform resource azurerm_management_group_policy_exemption.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermManagementGroupPolicyExemptionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (amgpe *Resource) Type() string {
	return "azurerm_management_group_policy_exemption"
}

// LocalName returns the local name for [Resource].
func (amgpe *Resource) LocalName() string {
	return amgpe.Name
}

// Configuration returns the configuration (args) for [Resource].
func (amgpe *Resource) Configuration() interface{} {
	return amgpe.Args
}

// DependOn is used for other resources to depend on [Resource].
func (amgpe *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(amgpe)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (amgpe *Resource) Dependencies() terra.Dependencies {
	return amgpe.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (amgpe *Resource) LifecycleManagement() *terra.Lifecycle {
	return amgpe.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (amgpe *Resource) Attributes() azurermManagementGroupPolicyExemptionAttributes {
	return azurermManagementGroupPolicyExemptionAttributes{ref: terra.ReferenceResource(amgpe)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (amgpe *Resource) ImportState(state io.Reader) error {
	amgpe.state = &azurermManagementGroupPolicyExemptionState{}
	if err := json.NewDecoder(state).Decode(amgpe.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amgpe.Type(), amgpe.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (amgpe *Resource) State() (*azurermManagementGroupPolicyExemptionState, bool) {
	return amgpe.state, amgpe.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (amgpe *Resource) StateMust() *azurermManagementGroupPolicyExemptionState {
	if amgpe.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amgpe.Type(), amgpe.LocalName()))
	}
	return amgpe.state
}

// Args contains the configurations for azurerm_management_group_policy_exemption.
type Args struct {
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
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermManagementGroupPolicyExemptionAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_management_group_policy_exemption.
func (amgpe azurermManagementGroupPolicyExemptionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(amgpe.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_management_group_policy_exemption.
func (amgpe azurermManagementGroupPolicyExemptionAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(amgpe.ref.Append("display_name"))
}

// ExemptionCategory returns a reference to field exemption_category of azurerm_management_group_policy_exemption.
func (amgpe azurermManagementGroupPolicyExemptionAttributes) ExemptionCategory() terra.StringValue {
	return terra.ReferenceAsString(amgpe.ref.Append("exemption_category"))
}

// ExpiresOn returns a reference to field expires_on of azurerm_management_group_policy_exemption.
func (amgpe azurermManagementGroupPolicyExemptionAttributes) ExpiresOn() terra.StringValue {
	return terra.ReferenceAsString(amgpe.ref.Append("expires_on"))
}

// Id returns a reference to field id of azurerm_management_group_policy_exemption.
func (amgpe azurermManagementGroupPolicyExemptionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amgpe.ref.Append("id"))
}

// ManagementGroupId returns a reference to field management_group_id of azurerm_management_group_policy_exemption.
func (amgpe azurermManagementGroupPolicyExemptionAttributes) ManagementGroupId() terra.StringValue {
	return terra.ReferenceAsString(amgpe.ref.Append("management_group_id"))
}

// Metadata returns a reference to field metadata of azurerm_management_group_policy_exemption.
func (amgpe azurermManagementGroupPolicyExemptionAttributes) Metadata() terra.StringValue {
	return terra.ReferenceAsString(amgpe.ref.Append("metadata"))
}

// Name returns a reference to field name of azurerm_management_group_policy_exemption.
func (amgpe azurermManagementGroupPolicyExemptionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(amgpe.ref.Append("name"))
}

// PolicyAssignmentId returns a reference to field policy_assignment_id of azurerm_management_group_policy_exemption.
func (amgpe azurermManagementGroupPolicyExemptionAttributes) PolicyAssignmentId() terra.StringValue {
	return terra.ReferenceAsString(amgpe.ref.Append("policy_assignment_id"))
}

// PolicyDefinitionReferenceIds returns a reference to field policy_definition_reference_ids of azurerm_management_group_policy_exemption.
func (amgpe azurermManagementGroupPolicyExemptionAttributes) PolicyDefinitionReferenceIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](amgpe.ref.Append("policy_definition_reference_ids"))
}

func (amgpe azurermManagementGroupPolicyExemptionAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](amgpe.ref.Append("timeouts"))
}

type azurermManagementGroupPolicyExemptionState struct {
	Description                  string         `json:"description"`
	DisplayName                  string         `json:"display_name"`
	ExemptionCategory            string         `json:"exemption_category"`
	ExpiresOn                    string         `json:"expires_on"`
	Id                           string         `json:"id"`
	ManagementGroupId            string         `json:"management_group_id"`
	Metadata                     string         `json:"metadata"`
	Name                         string         `json:"name"`
	PolicyAssignmentId           string         `json:"policy_assignment_id"`
	PolicyDefinitionReferenceIds []string       `json:"policy_definition_reference_ids"`
	Timeouts                     *TimeoutsState `json:"timeouts"`
}
