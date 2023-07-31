// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	resourcepolicyexemption "github.com/golingon/terraproviders/azurerm/3.67.0/resourcepolicyexemption"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewResourcePolicyExemption creates a new instance of [ResourcePolicyExemption].
func NewResourcePolicyExemption(name string, args ResourcePolicyExemptionArgs) *ResourcePolicyExemption {
	return &ResourcePolicyExemption{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ResourcePolicyExemption)(nil)

// ResourcePolicyExemption represents the Terraform resource azurerm_resource_policy_exemption.
type ResourcePolicyExemption struct {
	Name      string
	Args      ResourcePolicyExemptionArgs
	state     *resourcePolicyExemptionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ResourcePolicyExemption].
func (rpe *ResourcePolicyExemption) Type() string {
	return "azurerm_resource_policy_exemption"
}

// LocalName returns the local name for [ResourcePolicyExemption].
func (rpe *ResourcePolicyExemption) LocalName() string {
	return rpe.Name
}

// Configuration returns the configuration (args) for [ResourcePolicyExemption].
func (rpe *ResourcePolicyExemption) Configuration() interface{} {
	return rpe.Args
}

// DependOn is used for other resources to depend on [ResourcePolicyExemption].
func (rpe *ResourcePolicyExemption) DependOn() terra.Reference {
	return terra.ReferenceResource(rpe)
}

// Dependencies returns the list of resources [ResourcePolicyExemption] depends_on.
func (rpe *ResourcePolicyExemption) Dependencies() terra.Dependencies {
	return rpe.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ResourcePolicyExemption].
func (rpe *ResourcePolicyExemption) LifecycleManagement() *terra.Lifecycle {
	return rpe.Lifecycle
}

// Attributes returns the attributes for [ResourcePolicyExemption].
func (rpe *ResourcePolicyExemption) Attributes() resourcePolicyExemptionAttributes {
	return resourcePolicyExemptionAttributes{ref: terra.ReferenceResource(rpe)}
}

// ImportState imports the given attribute values into [ResourcePolicyExemption]'s state.
func (rpe *ResourcePolicyExemption) ImportState(av io.Reader) error {
	rpe.state = &resourcePolicyExemptionState{}
	if err := json.NewDecoder(av).Decode(rpe.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rpe.Type(), rpe.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ResourcePolicyExemption] has state.
func (rpe *ResourcePolicyExemption) State() (*resourcePolicyExemptionState, bool) {
	return rpe.state, rpe.state != nil
}

// StateMust returns the state for [ResourcePolicyExemption]. Panics if the state is nil.
func (rpe *ResourcePolicyExemption) StateMust() *resourcePolicyExemptionState {
	if rpe.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rpe.Type(), rpe.LocalName()))
	}
	return rpe.state
}

// ResourcePolicyExemptionArgs contains the configurations for azurerm_resource_policy_exemption.
type ResourcePolicyExemptionArgs struct {
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
	// ResourceId: string, required
	ResourceId terra.StringValue `hcl:"resource_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *resourcepolicyexemption.Timeouts `hcl:"timeouts,block"`
}
type resourcePolicyExemptionAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_resource_policy_exemption.
func (rpe resourcePolicyExemptionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(rpe.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_resource_policy_exemption.
func (rpe resourcePolicyExemptionAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(rpe.ref.Append("display_name"))
}

// ExemptionCategory returns a reference to field exemption_category of azurerm_resource_policy_exemption.
func (rpe resourcePolicyExemptionAttributes) ExemptionCategory() terra.StringValue {
	return terra.ReferenceAsString(rpe.ref.Append("exemption_category"))
}

// ExpiresOn returns a reference to field expires_on of azurerm_resource_policy_exemption.
func (rpe resourcePolicyExemptionAttributes) ExpiresOn() terra.StringValue {
	return terra.ReferenceAsString(rpe.ref.Append("expires_on"))
}

// Id returns a reference to field id of azurerm_resource_policy_exemption.
func (rpe resourcePolicyExemptionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rpe.ref.Append("id"))
}

// Metadata returns a reference to field metadata of azurerm_resource_policy_exemption.
func (rpe resourcePolicyExemptionAttributes) Metadata() terra.StringValue {
	return terra.ReferenceAsString(rpe.ref.Append("metadata"))
}

// Name returns a reference to field name of azurerm_resource_policy_exemption.
func (rpe resourcePolicyExemptionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rpe.ref.Append("name"))
}

// PolicyAssignmentId returns a reference to field policy_assignment_id of azurerm_resource_policy_exemption.
func (rpe resourcePolicyExemptionAttributes) PolicyAssignmentId() terra.StringValue {
	return terra.ReferenceAsString(rpe.ref.Append("policy_assignment_id"))
}

// PolicyDefinitionReferenceIds returns a reference to field policy_definition_reference_ids of azurerm_resource_policy_exemption.
func (rpe resourcePolicyExemptionAttributes) PolicyDefinitionReferenceIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](rpe.ref.Append("policy_definition_reference_ids"))
}

// ResourceId returns a reference to field resource_id of azurerm_resource_policy_exemption.
func (rpe resourcePolicyExemptionAttributes) ResourceId() terra.StringValue {
	return terra.ReferenceAsString(rpe.ref.Append("resource_id"))
}

func (rpe resourcePolicyExemptionAttributes) Timeouts() resourcepolicyexemption.TimeoutsAttributes {
	return terra.ReferenceAsSingle[resourcepolicyexemption.TimeoutsAttributes](rpe.ref.Append("timeouts"))
}

type resourcePolicyExemptionState struct {
	Description                  string                                 `json:"description"`
	DisplayName                  string                                 `json:"display_name"`
	ExemptionCategory            string                                 `json:"exemption_category"`
	ExpiresOn                    string                                 `json:"expires_on"`
	Id                           string                                 `json:"id"`
	Metadata                     string                                 `json:"metadata"`
	Name                         string                                 `json:"name"`
	PolicyAssignmentId           string                                 `json:"policy_assignment_id"`
	PolicyDefinitionReferenceIds []string                               `json:"policy_definition_reference_ids"`
	ResourceId                   string                                 `json:"resource_id"`
	Timeouts                     *resourcepolicyexemption.TimeoutsState `json:"timeouts"`
}
