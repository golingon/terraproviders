// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	resourcegrouppolicyremediation "github.com/golingon/terraproviders/azurerm/3.63.0/resourcegrouppolicyremediation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewResourceGroupPolicyRemediation creates a new instance of [ResourceGroupPolicyRemediation].
func NewResourceGroupPolicyRemediation(name string, args ResourceGroupPolicyRemediationArgs) *ResourceGroupPolicyRemediation {
	return &ResourceGroupPolicyRemediation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ResourceGroupPolicyRemediation)(nil)

// ResourceGroupPolicyRemediation represents the Terraform resource azurerm_resource_group_policy_remediation.
type ResourceGroupPolicyRemediation struct {
	Name      string
	Args      ResourceGroupPolicyRemediationArgs
	state     *resourceGroupPolicyRemediationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ResourceGroupPolicyRemediation].
func (rgpr *ResourceGroupPolicyRemediation) Type() string {
	return "azurerm_resource_group_policy_remediation"
}

// LocalName returns the local name for [ResourceGroupPolicyRemediation].
func (rgpr *ResourceGroupPolicyRemediation) LocalName() string {
	return rgpr.Name
}

// Configuration returns the configuration (args) for [ResourceGroupPolicyRemediation].
func (rgpr *ResourceGroupPolicyRemediation) Configuration() interface{} {
	return rgpr.Args
}

// DependOn is used for other resources to depend on [ResourceGroupPolicyRemediation].
func (rgpr *ResourceGroupPolicyRemediation) DependOn() terra.Reference {
	return terra.ReferenceResource(rgpr)
}

// Dependencies returns the list of resources [ResourceGroupPolicyRemediation] depends_on.
func (rgpr *ResourceGroupPolicyRemediation) Dependencies() terra.Dependencies {
	return rgpr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ResourceGroupPolicyRemediation].
func (rgpr *ResourceGroupPolicyRemediation) LifecycleManagement() *terra.Lifecycle {
	return rgpr.Lifecycle
}

// Attributes returns the attributes for [ResourceGroupPolicyRemediation].
func (rgpr *ResourceGroupPolicyRemediation) Attributes() resourceGroupPolicyRemediationAttributes {
	return resourceGroupPolicyRemediationAttributes{ref: terra.ReferenceResource(rgpr)}
}

// ImportState imports the given attribute values into [ResourceGroupPolicyRemediation]'s state.
func (rgpr *ResourceGroupPolicyRemediation) ImportState(av io.Reader) error {
	rgpr.state = &resourceGroupPolicyRemediationState{}
	if err := json.NewDecoder(av).Decode(rgpr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rgpr.Type(), rgpr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ResourceGroupPolicyRemediation] has state.
func (rgpr *ResourceGroupPolicyRemediation) State() (*resourceGroupPolicyRemediationState, bool) {
	return rgpr.state, rgpr.state != nil
}

// StateMust returns the state for [ResourceGroupPolicyRemediation]. Panics if the state is nil.
func (rgpr *ResourceGroupPolicyRemediation) StateMust() *resourceGroupPolicyRemediationState {
	if rgpr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rgpr.Type(), rgpr.LocalName()))
	}
	return rgpr.state
}

// ResourceGroupPolicyRemediationArgs contains the configurations for azurerm_resource_group_policy_remediation.
type ResourceGroupPolicyRemediationArgs struct {
	// FailurePercentage: number, optional
	FailurePercentage terra.NumberValue `hcl:"failure_percentage,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LocationFilters: list of string, optional
	LocationFilters terra.ListValue[terra.StringValue] `hcl:"location_filters,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ParallelDeployments: number, optional
	ParallelDeployments terra.NumberValue `hcl:"parallel_deployments,attr"`
	// PolicyAssignmentId: string, required
	PolicyAssignmentId terra.StringValue `hcl:"policy_assignment_id,attr" validate:"required"`
	// PolicyDefinitionId: string, optional
	PolicyDefinitionId terra.StringValue `hcl:"policy_definition_id,attr"`
	// PolicyDefinitionReferenceId: string, optional
	PolicyDefinitionReferenceId terra.StringValue `hcl:"policy_definition_reference_id,attr"`
	// ResourceCount: number, optional
	ResourceCount terra.NumberValue `hcl:"resource_count,attr"`
	// ResourceDiscoveryMode: string, optional
	ResourceDiscoveryMode terra.StringValue `hcl:"resource_discovery_mode,attr"`
	// ResourceGroupId: string, required
	ResourceGroupId terra.StringValue `hcl:"resource_group_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *resourcegrouppolicyremediation.Timeouts `hcl:"timeouts,block"`
}
type resourceGroupPolicyRemediationAttributes struct {
	ref terra.Reference
}

// FailurePercentage returns a reference to field failure_percentage of azurerm_resource_group_policy_remediation.
func (rgpr resourceGroupPolicyRemediationAttributes) FailurePercentage() terra.NumberValue {
	return terra.ReferenceAsNumber(rgpr.ref.Append("failure_percentage"))
}

// Id returns a reference to field id of azurerm_resource_group_policy_remediation.
func (rgpr resourceGroupPolicyRemediationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rgpr.ref.Append("id"))
}

// LocationFilters returns a reference to field location_filters of azurerm_resource_group_policy_remediation.
func (rgpr resourceGroupPolicyRemediationAttributes) LocationFilters() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](rgpr.ref.Append("location_filters"))
}

// Name returns a reference to field name of azurerm_resource_group_policy_remediation.
func (rgpr resourceGroupPolicyRemediationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rgpr.ref.Append("name"))
}

// ParallelDeployments returns a reference to field parallel_deployments of azurerm_resource_group_policy_remediation.
func (rgpr resourceGroupPolicyRemediationAttributes) ParallelDeployments() terra.NumberValue {
	return terra.ReferenceAsNumber(rgpr.ref.Append("parallel_deployments"))
}

// PolicyAssignmentId returns a reference to field policy_assignment_id of azurerm_resource_group_policy_remediation.
func (rgpr resourceGroupPolicyRemediationAttributes) PolicyAssignmentId() terra.StringValue {
	return terra.ReferenceAsString(rgpr.ref.Append("policy_assignment_id"))
}

// PolicyDefinitionId returns a reference to field policy_definition_id of azurerm_resource_group_policy_remediation.
func (rgpr resourceGroupPolicyRemediationAttributes) PolicyDefinitionId() terra.StringValue {
	return terra.ReferenceAsString(rgpr.ref.Append("policy_definition_id"))
}

// PolicyDefinitionReferenceId returns a reference to field policy_definition_reference_id of azurerm_resource_group_policy_remediation.
func (rgpr resourceGroupPolicyRemediationAttributes) PolicyDefinitionReferenceId() terra.StringValue {
	return terra.ReferenceAsString(rgpr.ref.Append("policy_definition_reference_id"))
}

// ResourceCount returns a reference to field resource_count of azurerm_resource_group_policy_remediation.
func (rgpr resourceGroupPolicyRemediationAttributes) ResourceCount() terra.NumberValue {
	return terra.ReferenceAsNumber(rgpr.ref.Append("resource_count"))
}

// ResourceDiscoveryMode returns a reference to field resource_discovery_mode of azurerm_resource_group_policy_remediation.
func (rgpr resourceGroupPolicyRemediationAttributes) ResourceDiscoveryMode() terra.StringValue {
	return terra.ReferenceAsString(rgpr.ref.Append("resource_discovery_mode"))
}

// ResourceGroupId returns a reference to field resource_group_id of azurerm_resource_group_policy_remediation.
func (rgpr resourceGroupPolicyRemediationAttributes) ResourceGroupId() terra.StringValue {
	return terra.ReferenceAsString(rgpr.ref.Append("resource_group_id"))
}

func (rgpr resourceGroupPolicyRemediationAttributes) Timeouts() resourcegrouppolicyremediation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[resourcegrouppolicyremediation.TimeoutsAttributes](rgpr.ref.Append("timeouts"))
}

type resourceGroupPolicyRemediationState struct {
	FailurePercentage           float64                                       `json:"failure_percentage"`
	Id                          string                                        `json:"id"`
	LocationFilters             []string                                      `json:"location_filters"`
	Name                        string                                        `json:"name"`
	ParallelDeployments         float64                                       `json:"parallel_deployments"`
	PolicyAssignmentId          string                                        `json:"policy_assignment_id"`
	PolicyDefinitionId          string                                        `json:"policy_definition_id"`
	PolicyDefinitionReferenceId string                                        `json:"policy_definition_reference_id"`
	ResourceCount               float64                                       `json:"resource_count"`
	ResourceDiscoveryMode       string                                        `json:"resource_discovery_mode"`
	ResourceGroupId             string                                        `json:"resource_group_id"`
	Timeouts                    *resourcegrouppolicyremediation.TimeoutsState `json:"timeouts"`
}
