// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	managementgrouppolicyassignment "github.com/golingon/terraproviders/azurerm/3.64.0/managementgrouppolicyassignment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewManagementGroupPolicyAssignment creates a new instance of [ManagementGroupPolicyAssignment].
func NewManagementGroupPolicyAssignment(name string, args ManagementGroupPolicyAssignmentArgs) *ManagementGroupPolicyAssignment {
	return &ManagementGroupPolicyAssignment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ManagementGroupPolicyAssignment)(nil)

// ManagementGroupPolicyAssignment represents the Terraform resource azurerm_management_group_policy_assignment.
type ManagementGroupPolicyAssignment struct {
	Name      string
	Args      ManagementGroupPolicyAssignmentArgs
	state     *managementGroupPolicyAssignmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ManagementGroupPolicyAssignment].
func (mgpa *ManagementGroupPolicyAssignment) Type() string {
	return "azurerm_management_group_policy_assignment"
}

// LocalName returns the local name for [ManagementGroupPolicyAssignment].
func (mgpa *ManagementGroupPolicyAssignment) LocalName() string {
	return mgpa.Name
}

// Configuration returns the configuration (args) for [ManagementGroupPolicyAssignment].
func (mgpa *ManagementGroupPolicyAssignment) Configuration() interface{} {
	return mgpa.Args
}

// DependOn is used for other resources to depend on [ManagementGroupPolicyAssignment].
func (mgpa *ManagementGroupPolicyAssignment) DependOn() terra.Reference {
	return terra.ReferenceResource(mgpa)
}

// Dependencies returns the list of resources [ManagementGroupPolicyAssignment] depends_on.
func (mgpa *ManagementGroupPolicyAssignment) Dependencies() terra.Dependencies {
	return mgpa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ManagementGroupPolicyAssignment].
func (mgpa *ManagementGroupPolicyAssignment) LifecycleManagement() *terra.Lifecycle {
	return mgpa.Lifecycle
}

// Attributes returns the attributes for [ManagementGroupPolicyAssignment].
func (mgpa *ManagementGroupPolicyAssignment) Attributes() managementGroupPolicyAssignmentAttributes {
	return managementGroupPolicyAssignmentAttributes{ref: terra.ReferenceResource(mgpa)}
}

// ImportState imports the given attribute values into [ManagementGroupPolicyAssignment]'s state.
func (mgpa *ManagementGroupPolicyAssignment) ImportState(av io.Reader) error {
	mgpa.state = &managementGroupPolicyAssignmentState{}
	if err := json.NewDecoder(av).Decode(mgpa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mgpa.Type(), mgpa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ManagementGroupPolicyAssignment] has state.
func (mgpa *ManagementGroupPolicyAssignment) State() (*managementGroupPolicyAssignmentState, bool) {
	return mgpa.state, mgpa.state != nil
}

// StateMust returns the state for [ManagementGroupPolicyAssignment]. Panics if the state is nil.
func (mgpa *ManagementGroupPolicyAssignment) StateMust() *managementGroupPolicyAssignmentState {
	if mgpa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mgpa.Type(), mgpa.LocalName()))
	}
	return mgpa.state
}

// ManagementGroupPolicyAssignmentArgs contains the configurations for azurerm_management_group_policy_assignment.
type ManagementGroupPolicyAssignmentArgs struct {
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
	// ManagementGroupId: string, required
	ManagementGroupId terra.StringValue `hcl:"management_group_id,attr" validate:"required"`
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
	// Identity: optional
	Identity *managementgrouppolicyassignment.Identity `hcl:"identity,block"`
	// NonComplianceMessage: min=0
	NonComplianceMessage []managementgrouppolicyassignment.NonComplianceMessage `hcl:"non_compliance_message,block" validate:"min=0"`
	// Overrides: min=0
	Overrides []managementgrouppolicyassignment.Overrides `hcl:"overrides,block" validate:"min=0"`
	// ResourceSelectors: min=0
	ResourceSelectors []managementgrouppolicyassignment.ResourceSelectors `hcl:"resource_selectors,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *managementgrouppolicyassignment.Timeouts `hcl:"timeouts,block"`
}
type managementGroupPolicyAssignmentAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_management_group_policy_assignment.
func (mgpa managementGroupPolicyAssignmentAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(mgpa.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_management_group_policy_assignment.
func (mgpa managementGroupPolicyAssignmentAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(mgpa.ref.Append("display_name"))
}

// Enforce returns a reference to field enforce of azurerm_management_group_policy_assignment.
func (mgpa managementGroupPolicyAssignmentAttributes) Enforce() terra.BoolValue {
	return terra.ReferenceAsBool(mgpa.ref.Append("enforce"))
}

// Id returns a reference to field id of azurerm_management_group_policy_assignment.
func (mgpa managementGroupPolicyAssignmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mgpa.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_management_group_policy_assignment.
func (mgpa managementGroupPolicyAssignmentAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(mgpa.ref.Append("location"))
}

// ManagementGroupId returns a reference to field management_group_id of azurerm_management_group_policy_assignment.
func (mgpa managementGroupPolicyAssignmentAttributes) ManagementGroupId() terra.StringValue {
	return terra.ReferenceAsString(mgpa.ref.Append("management_group_id"))
}

// Metadata returns a reference to field metadata of azurerm_management_group_policy_assignment.
func (mgpa managementGroupPolicyAssignmentAttributes) Metadata() terra.StringValue {
	return terra.ReferenceAsString(mgpa.ref.Append("metadata"))
}

// Name returns a reference to field name of azurerm_management_group_policy_assignment.
func (mgpa managementGroupPolicyAssignmentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mgpa.ref.Append("name"))
}

// NotScopes returns a reference to field not_scopes of azurerm_management_group_policy_assignment.
func (mgpa managementGroupPolicyAssignmentAttributes) NotScopes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](mgpa.ref.Append("not_scopes"))
}

// Parameters returns a reference to field parameters of azurerm_management_group_policy_assignment.
func (mgpa managementGroupPolicyAssignmentAttributes) Parameters() terra.StringValue {
	return terra.ReferenceAsString(mgpa.ref.Append("parameters"))
}

// PolicyDefinitionId returns a reference to field policy_definition_id of azurerm_management_group_policy_assignment.
func (mgpa managementGroupPolicyAssignmentAttributes) PolicyDefinitionId() terra.StringValue {
	return terra.ReferenceAsString(mgpa.ref.Append("policy_definition_id"))
}

func (mgpa managementGroupPolicyAssignmentAttributes) Identity() terra.ListValue[managementgrouppolicyassignment.IdentityAttributes] {
	return terra.ReferenceAsList[managementgrouppolicyassignment.IdentityAttributes](mgpa.ref.Append("identity"))
}

func (mgpa managementGroupPolicyAssignmentAttributes) NonComplianceMessage() terra.ListValue[managementgrouppolicyassignment.NonComplianceMessageAttributes] {
	return terra.ReferenceAsList[managementgrouppolicyassignment.NonComplianceMessageAttributes](mgpa.ref.Append("non_compliance_message"))
}

func (mgpa managementGroupPolicyAssignmentAttributes) Overrides() terra.ListValue[managementgrouppolicyassignment.OverridesAttributes] {
	return terra.ReferenceAsList[managementgrouppolicyassignment.OverridesAttributes](mgpa.ref.Append("overrides"))
}

func (mgpa managementGroupPolicyAssignmentAttributes) ResourceSelectors() terra.ListValue[managementgrouppolicyassignment.ResourceSelectorsAttributes] {
	return terra.ReferenceAsList[managementgrouppolicyassignment.ResourceSelectorsAttributes](mgpa.ref.Append("resource_selectors"))
}

func (mgpa managementGroupPolicyAssignmentAttributes) Timeouts() managementgrouppolicyassignment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[managementgrouppolicyassignment.TimeoutsAttributes](mgpa.ref.Append("timeouts"))
}

type managementGroupPolicyAssignmentState struct {
	Description          string                                                      `json:"description"`
	DisplayName          string                                                      `json:"display_name"`
	Enforce              bool                                                        `json:"enforce"`
	Id                   string                                                      `json:"id"`
	Location             string                                                      `json:"location"`
	ManagementGroupId    string                                                      `json:"management_group_id"`
	Metadata             string                                                      `json:"metadata"`
	Name                 string                                                      `json:"name"`
	NotScopes            []string                                                    `json:"not_scopes"`
	Parameters           string                                                      `json:"parameters"`
	PolicyDefinitionId   string                                                      `json:"policy_definition_id"`
	Identity             []managementgrouppolicyassignment.IdentityState             `json:"identity"`
	NonComplianceMessage []managementgrouppolicyassignment.NonComplianceMessageState `json:"non_compliance_message"`
	Overrides            []managementgrouppolicyassignment.OverridesState            `json:"overrides"`
	ResourceSelectors    []managementgrouppolicyassignment.ResourceSelectorsState    `json:"resource_selectors"`
	Timeouts             *managementgrouppolicyassignment.TimeoutsState              `json:"timeouts"`
}
