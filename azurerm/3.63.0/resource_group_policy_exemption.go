// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	resourcegrouppolicyexemption "github.com/golingon/terraproviders/azurerm/3.63.0/resourcegrouppolicyexemption"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewResourceGroupPolicyExemption creates a new instance of [ResourceGroupPolicyExemption].
func NewResourceGroupPolicyExemption(name string, args ResourceGroupPolicyExemptionArgs) *ResourceGroupPolicyExemption {
	return &ResourceGroupPolicyExemption{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ResourceGroupPolicyExemption)(nil)

// ResourceGroupPolicyExemption represents the Terraform resource azurerm_resource_group_policy_exemption.
type ResourceGroupPolicyExemption struct {
	Name      string
	Args      ResourceGroupPolicyExemptionArgs
	state     *resourceGroupPolicyExemptionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ResourceGroupPolicyExemption].
func (rgpe *ResourceGroupPolicyExemption) Type() string {
	return "azurerm_resource_group_policy_exemption"
}

// LocalName returns the local name for [ResourceGroupPolicyExemption].
func (rgpe *ResourceGroupPolicyExemption) LocalName() string {
	return rgpe.Name
}

// Configuration returns the configuration (args) for [ResourceGroupPolicyExemption].
func (rgpe *ResourceGroupPolicyExemption) Configuration() interface{} {
	return rgpe.Args
}

// DependOn is used for other resources to depend on [ResourceGroupPolicyExemption].
func (rgpe *ResourceGroupPolicyExemption) DependOn() terra.Reference {
	return terra.ReferenceResource(rgpe)
}

// Dependencies returns the list of resources [ResourceGroupPolicyExemption] depends_on.
func (rgpe *ResourceGroupPolicyExemption) Dependencies() terra.Dependencies {
	return rgpe.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ResourceGroupPolicyExemption].
func (rgpe *ResourceGroupPolicyExemption) LifecycleManagement() *terra.Lifecycle {
	return rgpe.Lifecycle
}

// Attributes returns the attributes for [ResourceGroupPolicyExemption].
func (rgpe *ResourceGroupPolicyExemption) Attributes() resourceGroupPolicyExemptionAttributes {
	return resourceGroupPolicyExemptionAttributes{ref: terra.ReferenceResource(rgpe)}
}

// ImportState imports the given attribute values into [ResourceGroupPolicyExemption]'s state.
func (rgpe *ResourceGroupPolicyExemption) ImportState(av io.Reader) error {
	rgpe.state = &resourceGroupPolicyExemptionState{}
	if err := json.NewDecoder(av).Decode(rgpe.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rgpe.Type(), rgpe.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ResourceGroupPolicyExemption] has state.
func (rgpe *ResourceGroupPolicyExemption) State() (*resourceGroupPolicyExemptionState, bool) {
	return rgpe.state, rgpe.state != nil
}

// StateMust returns the state for [ResourceGroupPolicyExemption]. Panics if the state is nil.
func (rgpe *ResourceGroupPolicyExemption) StateMust() *resourceGroupPolicyExemptionState {
	if rgpe.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rgpe.Type(), rgpe.LocalName()))
	}
	return rgpe.state
}

// ResourceGroupPolicyExemptionArgs contains the configurations for azurerm_resource_group_policy_exemption.
type ResourceGroupPolicyExemptionArgs struct {
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
	// Metadata: string, optional
	Metadata terra.StringValue `hcl:"metadata,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PolicyAssignmentId: string, required
	PolicyAssignmentId terra.StringValue `hcl:"policy_assignment_id,attr" validate:"required"`
	// PolicyDefinitionReferenceIds: list of string, optional
	PolicyDefinitionReferenceIds terra.ListValue[terra.StringValue] `hcl:"policy_definition_reference_ids,attr"`
	// ResourceGroupId: string, required
	ResourceGroupId terra.StringValue `hcl:"resource_group_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *resourcegrouppolicyexemption.Timeouts `hcl:"timeouts,block"`
}
type resourceGroupPolicyExemptionAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_resource_group_policy_exemption.
func (rgpe resourceGroupPolicyExemptionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(rgpe.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_resource_group_policy_exemption.
func (rgpe resourceGroupPolicyExemptionAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(rgpe.ref.Append("display_name"))
}

// ExemptionCategory returns a reference to field exemption_category of azurerm_resource_group_policy_exemption.
func (rgpe resourceGroupPolicyExemptionAttributes) ExemptionCategory() terra.StringValue {
	return terra.ReferenceAsString(rgpe.ref.Append("exemption_category"))
}

// ExpiresOn returns a reference to field expires_on of azurerm_resource_group_policy_exemption.
func (rgpe resourceGroupPolicyExemptionAttributes) ExpiresOn() terra.StringValue {
	return terra.ReferenceAsString(rgpe.ref.Append("expires_on"))
}

// Id returns a reference to field id of azurerm_resource_group_policy_exemption.
func (rgpe resourceGroupPolicyExemptionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rgpe.ref.Append("id"))
}

// Metadata returns a reference to field metadata of azurerm_resource_group_policy_exemption.
func (rgpe resourceGroupPolicyExemptionAttributes) Metadata() terra.StringValue {
	return terra.ReferenceAsString(rgpe.ref.Append("metadata"))
}

// Name returns a reference to field name of azurerm_resource_group_policy_exemption.
func (rgpe resourceGroupPolicyExemptionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rgpe.ref.Append("name"))
}

// PolicyAssignmentId returns a reference to field policy_assignment_id of azurerm_resource_group_policy_exemption.
func (rgpe resourceGroupPolicyExemptionAttributes) PolicyAssignmentId() terra.StringValue {
	return terra.ReferenceAsString(rgpe.ref.Append("policy_assignment_id"))
}

// PolicyDefinitionReferenceIds returns a reference to field policy_definition_reference_ids of azurerm_resource_group_policy_exemption.
func (rgpe resourceGroupPolicyExemptionAttributes) PolicyDefinitionReferenceIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](rgpe.ref.Append("policy_definition_reference_ids"))
}

// ResourceGroupId returns a reference to field resource_group_id of azurerm_resource_group_policy_exemption.
func (rgpe resourceGroupPolicyExemptionAttributes) ResourceGroupId() terra.StringValue {
	return terra.ReferenceAsString(rgpe.ref.Append("resource_group_id"))
}

func (rgpe resourceGroupPolicyExemptionAttributes) Timeouts() resourcegrouppolicyexemption.TimeoutsAttributes {
	return terra.ReferenceAsSingle[resourcegrouppolicyexemption.TimeoutsAttributes](rgpe.ref.Append("timeouts"))
}

type resourceGroupPolicyExemptionState struct {
	Description                  string                                      `json:"description"`
	DisplayName                  string                                      `json:"display_name"`
	ExemptionCategory            string                                      `json:"exemption_category"`
	ExpiresOn                    string                                      `json:"expires_on"`
	Id                           string                                      `json:"id"`
	Metadata                     string                                      `json:"metadata"`
	Name                         string                                      `json:"name"`
	PolicyAssignmentId           string                                      `json:"policy_assignment_id"`
	PolicyDefinitionReferenceIds []string                                    `json:"policy_definition_reference_ids"`
	ResourceGroupId              string                                      `json:"resource_group_id"`
	Timeouts                     *resourcegrouppolicyexemption.TimeoutsState `json:"timeouts"`
}
