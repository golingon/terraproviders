// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	resourcegrouppolicyassignment "github.com/golingon/terraproviders/azurerm/3.98.0/resourcegrouppolicyassignment"
	"io"
)

// NewResourceGroupPolicyAssignment creates a new instance of [ResourceGroupPolicyAssignment].
func NewResourceGroupPolicyAssignment(name string, args ResourceGroupPolicyAssignmentArgs) *ResourceGroupPolicyAssignment {
	return &ResourceGroupPolicyAssignment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ResourceGroupPolicyAssignment)(nil)

// ResourceGroupPolicyAssignment represents the Terraform resource azurerm_resource_group_policy_assignment.
type ResourceGroupPolicyAssignment struct {
	Name      string
	Args      ResourceGroupPolicyAssignmentArgs
	state     *resourceGroupPolicyAssignmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ResourceGroupPolicyAssignment].
func (rgpa *ResourceGroupPolicyAssignment) Type() string {
	return "azurerm_resource_group_policy_assignment"
}

// LocalName returns the local name for [ResourceGroupPolicyAssignment].
func (rgpa *ResourceGroupPolicyAssignment) LocalName() string {
	return rgpa.Name
}

// Configuration returns the configuration (args) for [ResourceGroupPolicyAssignment].
func (rgpa *ResourceGroupPolicyAssignment) Configuration() interface{} {
	return rgpa.Args
}

// DependOn is used for other resources to depend on [ResourceGroupPolicyAssignment].
func (rgpa *ResourceGroupPolicyAssignment) DependOn() terra.Reference {
	return terra.ReferenceResource(rgpa)
}

// Dependencies returns the list of resources [ResourceGroupPolicyAssignment] depends_on.
func (rgpa *ResourceGroupPolicyAssignment) Dependencies() terra.Dependencies {
	return rgpa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ResourceGroupPolicyAssignment].
func (rgpa *ResourceGroupPolicyAssignment) LifecycleManagement() *terra.Lifecycle {
	return rgpa.Lifecycle
}

// Attributes returns the attributes for [ResourceGroupPolicyAssignment].
func (rgpa *ResourceGroupPolicyAssignment) Attributes() resourceGroupPolicyAssignmentAttributes {
	return resourceGroupPolicyAssignmentAttributes{ref: terra.ReferenceResource(rgpa)}
}

// ImportState imports the given attribute values into [ResourceGroupPolicyAssignment]'s state.
func (rgpa *ResourceGroupPolicyAssignment) ImportState(av io.Reader) error {
	rgpa.state = &resourceGroupPolicyAssignmentState{}
	if err := json.NewDecoder(av).Decode(rgpa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rgpa.Type(), rgpa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ResourceGroupPolicyAssignment] has state.
func (rgpa *ResourceGroupPolicyAssignment) State() (*resourceGroupPolicyAssignmentState, bool) {
	return rgpa.state, rgpa.state != nil
}

// StateMust returns the state for [ResourceGroupPolicyAssignment]. Panics if the state is nil.
func (rgpa *ResourceGroupPolicyAssignment) StateMust() *resourceGroupPolicyAssignmentState {
	if rgpa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rgpa.Type(), rgpa.LocalName()))
	}
	return rgpa.state
}

// ResourceGroupPolicyAssignmentArgs contains the configurations for azurerm_resource_group_policy_assignment.
type ResourceGroupPolicyAssignmentArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Enforce: bool, optional
	Enforce terra.BoolValue `hcl:"enforce,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Metadata: string, optional
	Metadata terra.StringValue `hcl:"metadata,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NotScopes: list of string, optional
	NotScopes terra.ListValue[terra.StringValue] `hcl:"not_scopes,attr"`
	// Parameters: string, optional
	Parameters terra.StringValue `hcl:"parameters,attr"`
	// PolicyDefinitionId: string, required
	PolicyDefinitionId terra.StringValue `hcl:"policy_definition_id,attr" validate:"required"`
	// ResourceGroupId: string, required
	ResourceGroupId terra.StringValue `hcl:"resource_group_id,attr" validate:"required"`
	// Identity: optional
	Identity *resourcegrouppolicyassignment.Identity `hcl:"identity,block"`
	// NonComplianceMessage: min=0
	NonComplianceMessage []resourcegrouppolicyassignment.NonComplianceMessage `hcl:"non_compliance_message,block" validate:"min=0"`
	// Overrides: min=0
	Overrides []resourcegrouppolicyassignment.Overrides `hcl:"overrides,block" validate:"min=0"`
	// ResourceSelectors: min=0
	ResourceSelectors []resourcegrouppolicyassignment.ResourceSelectors `hcl:"resource_selectors,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *resourcegrouppolicyassignment.Timeouts `hcl:"timeouts,block"`
}
type resourceGroupPolicyAssignmentAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_resource_group_policy_assignment.
func (rgpa resourceGroupPolicyAssignmentAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(rgpa.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_resource_group_policy_assignment.
func (rgpa resourceGroupPolicyAssignmentAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(rgpa.ref.Append("display_name"))
}

// Enforce returns a reference to field enforce of azurerm_resource_group_policy_assignment.
func (rgpa resourceGroupPolicyAssignmentAttributes) Enforce() terra.BoolValue {
	return terra.ReferenceAsBool(rgpa.ref.Append("enforce"))
}

// Id returns a reference to field id of azurerm_resource_group_policy_assignment.
func (rgpa resourceGroupPolicyAssignmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rgpa.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_resource_group_policy_assignment.
func (rgpa resourceGroupPolicyAssignmentAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(rgpa.ref.Append("location"))
}

// Metadata returns a reference to field metadata of azurerm_resource_group_policy_assignment.
func (rgpa resourceGroupPolicyAssignmentAttributes) Metadata() terra.StringValue {
	return terra.ReferenceAsString(rgpa.ref.Append("metadata"))
}

// Name returns a reference to field name of azurerm_resource_group_policy_assignment.
func (rgpa resourceGroupPolicyAssignmentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rgpa.ref.Append("name"))
}

// NotScopes returns a reference to field not_scopes of azurerm_resource_group_policy_assignment.
func (rgpa resourceGroupPolicyAssignmentAttributes) NotScopes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](rgpa.ref.Append("not_scopes"))
}

// Parameters returns a reference to field parameters of azurerm_resource_group_policy_assignment.
func (rgpa resourceGroupPolicyAssignmentAttributes) Parameters() terra.StringValue {
	return terra.ReferenceAsString(rgpa.ref.Append("parameters"))
}

// PolicyDefinitionId returns a reference to field policy_definition_id of azurerm_resource_group_policy_assignment.
func (rgpa resourceGroupPolicyAssignmentAttributes) PolicyDefinitionId() terra.StringValue {
	return terra.ReferenceAsString(rgpa.ref.Append("policy_definition_id"))
}

// ResourceGroupId returns a reference to field resource_group_id of azurerm_resource_group_policy_assignment.
func (rgpa resourceGroupPolicyAssignmentAttributes) ResourceGroupId() terra.StringValue {
	return terra.ReferenceAsString(rgpa.ref.Append("resource_group_id"))
}

func (rgpa resourceGroupPolicyAssignmentAttributes) Identity() terra.ListValue[resourcegrouppolicyassignment.IdentityAttributes] {
	return terra.ReferenceAsList[resourcegrouppolicyassignment.IdentityAttributes](rgpa.ref.Append("identity"))
}

func (rgpa resourceGroupPolicyAssignmentAttributes) NonComplianceMessage() terra.ListValue[resourcegrouppolicyassignment.NonComplianceMessageAttributes] {
	return terra.ReferenceAsList[resourcegrouppolicyassignment.NonComplianceMessageAttributes](rgpa.ref.Append("non_compliance_message"))
}

func (rgpa resourceGroupPolicyAssignmentAttributes) Overrides() terra.ListValue[resourcegrouppolicyassignment.OverridesAttributes] {
	return terra.ReferenceAsList[resourcegrouppolicyassignment.OverridesAttributes](rgpa.ref.Append("overrides"))
}

func (rgpa resourceGroupPolicyAssignmentAttributes) ResourceSelectors() terra.ListValue[resourcegrouppolicyassignment.ResourceSelectorsAttributes] {
	return terra.ReferenceAsList[resourcegrouppolicyassignment.ResourceSelectorsAttributes](rgpa.ref.Append("resource_selectors"))
}

func (rgpa resourceGroupPolicyAssignmentAttributes) Timeouts() resourcegrouppolicyassignment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[resourcegrouppolicyassignment.TimeoutsAttributes](rgpa.ref.Append("timeouts"))
}

type resourceGroupPolicyAssignmentState struct {
	Description          string                                                    `json:"description"`
	DisplayName          string                                                    `json:"display_name"`
	Enforce              bool                                                      `json:"enforce"`
	Id                   string                                                    `json:"id"`
	Location             string                                                    `json:"location"`
	Metadata             string                                                    `json:"metadata"`
	Name                 string                                                    `json:"name"`
	NotScopes            []string                                                  `json:"not_scopes"`
	Parameters           string                                                    `json:"parameters"`
	PolicyDefinitionId   string                                                    `json:"policy_definition_id"`
	ResourceGroupId      string                                                    `json:"resource_group_id"`
	Identity             []resourcegrouppolicyassignment.IdentityState             `json:"identity"`
	NonComplianceMessage []resourcegrouppolicyassignment.NonComplianceMessageState `json:"non_compliance_message"`
	Overrides            []resourcegrouppolicyassignment.OverridesState            `json:"overrides"`
	ResourceSelectors    []resourcegrouppolicyassignment.ResourceSelectorsState    `json:"resource_selectors"`
	Timeouts             *resourcegrouppolicyassignment.TimeoutsState              `json:"timeouts"`
}
