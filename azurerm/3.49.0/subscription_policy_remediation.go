// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	subscriptionpolicyremediation "github.com/golingon/terraproviders/azurerm/3.49.0/subscriptionpolicyremediation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSubscriptionPolicyRemediation creates a new instance of [SubscriptionPolicyRemediation].
func NewSubscriptionPolicyRemediation(name string, args SubscriptionPolicyRemediationArgs) *SubscriptionPolicyRemediation {
	return &SubscriptionPolicyRemediation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SubscriptionPolicyRemediation)(nil)

// SubscriptionPolicyRemediation represents the Terraform resource azurerm_subscription_policy_remediation.
type SubscriptionPolicyRemediation struct {
	Name      string
	Args      SubscriptionPolicyRemediationArgs
	state     *subscriptionPolicyRemediationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SubscriptionPolicyRemediation].
func (spr *SubscriptionPolicyRemediation) Type() string {
	return "azurerm_subscription_policy_remediation"
}

// LocalName returns the local name for [SubscriptionPolicyRemediation].
func (spr *SubscriptionPolicyRemediation) LocalName() string {
	return spr.Name
}

// Configuration returns the configuration (args) for [SubscriptionPolicyRemediation].
func (spr *SubscriptionPolicyRemediation) Configuration() interface{} {
	return spr.Args
}

// DependOn is used for other resources to depend on [SubscriptionPolicyRemediation].
func (spr *SubscriptionPolicyRemediation) DependOn() terra.Reference {
	return terra.ReferenceResource(spr)
}

// Dependencies returns the list of resources [SubscriptionPolicyRemediation] depends_on.
func (spr *SubscriptionPolicyRemediation) Dependencies() terra.Dependencies {
	return spr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SubscriptionPolicyRemediation].
func (spr *SubscriptionPolicyRemediation) LifecycleManagement() *terra.Lifecycle {
	return spr.Lifecycle
}

// Attributes returns the attributes for [SubscriptionPolicyRemediation].
func (spr *SubscriptionPolicyRemediation) Attributes() subscriptionPolicyRemediationAttributes {
	return subscriptionPolicyRemediationAttributes{ref: terra.ReferenceResource(spr)}
}

// ImportState imports the given attribute values into [SubscriptionPolicyRemediation]'s state.
func (spr *SubscriptionPolicyRemediation) ImportState(av io.Reader) error {
	spr.state = &subscriptionPolicyRemediationState{}
	if err := json.NewDecoder(av).Decode(spr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", spr.Type(), spr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SubscriptionPolicyRemediation] has state.
func (spr *SubscriptionPolicyRemediation) State() (*subscriptionPolicyRemediationState, bool) {
	return spr.state, spr.state != nil
}

// StateMust returns the state for [SubscriptionPolicyRemediation]. Panics if the state is nil.
func (spr *SubscriptionPolicyRemediation) StateMust() *subscriptionPolicyRemediationState {
	if spr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", spr.Type(), spr.LocalName()))
	}
	return spr.state
}

// SubscriptionPolicyRemediationArgs contains the configurations for azurerm_subscription_policy_remediation.
type SubscriptionPolicyRemediationArgs struct {
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
	// SubscriptionId: string, required
	SubscriptionId terra.StringValue `hcl:"subscription_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *subscriptionpolicyremediation.Timeouts `hcl:"timeouts,block"`
}
type subscriptionPolicyRemediationAttributes struct {
	ref terra.Reference
}

// FailurePercentage returns a reference to field failure_percentage of azurerm_subscription_policy_remediation.
func (spr subscriptionPolicyRemediationAttributes) FailurePercentage() terra.NumberValue {
	return terra.ReferenceAsNumber(spr.ref.Append("failure_percentage"))
}

// Id returns a reference to field id of azurerm_subscription_policy_remediation.
func (spr subscriptionPolicyRemediationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(spr.ref.Append("id"))
}

// LocationFilters returns a reference to field location_filters of azurerm_subscription_policy_remediation.
func (spr subscriptionPolicyRemediationAttributes) LocationFilters() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](spr.ref.Append("location_filters"))
}

// Name returns a reference to field name of azurerm_subscription_policy_remediation.
func (spr subscriptionPolicyRemediationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(spr.ref.Append("name"))
}

// ParallelDeployments returns a reference to field parallel_deployments of azurerm_subscription_policy_remediation.
func (spr subscriptionPolicyRemediationAttributes) ParallelDeployments() terra.NumberValue {
	return terra.ReferenceAsNumber(spr.ref.Append("parallel_deployments"))
}

// PolicyAssignmentId returns a reference to field policy_assignment_id of azurerm_subscription_policy_remediation.
func (spr subscriptionPolicyRemediationAttributes) PolicyAssignmentId() terra.StringValue {
	return terra.ReferenceAsString(spr.ref.Append("policy_assignment_id"))
}

// PolicyDefinitionId returns a reference to field policy_definition_id of azurerm_subscription_policy_remediation.
func (spr subscriptionPolicyRemediationAttributes) PolicyDefinitionId() terra.StringValue {
	return terra.ReferenceAsString(spr.ref.Append("policy_definition_id"))
}

// PolicyDefinitionReferenceId returns a reference to field policy_definition_reference_id of azurerm_subscription_policy_remediation.
func (spr subscriptionPolicyRemediationAttributes) PolicyDefinitionReferenceId() terra.StringValue {
	return terra.ReferenceAsString(spr.ref.Append("policy_definition_reference_id"))
}

// ResourceCount returns a reference to field resource_count of azurerm_subscription_policy_remediation.
func (spr subscriptionPolicyRemediationAttributes) ResourceCount() terra.NumberValue {
	return terra.ReferenceAsNumber(spr.ref.Append("resource_count"))
}

// ResourceDiscoveryMode returns a reference to field resource_discovery_mode of azurerm_subscription_policy_remediation.
func (spr subscriptionPolicyRemediationAttributes) ResourceDiscoveryMode() terra.StringValue {
	return terra.ReferenceAsString(spr.ref.Append("resource_discovery_mode"))
}

// SubscriptionId returns a reference to field subscription_id of azurerm_subscription_policy_remediation.
func (spr subscriptionPolicyRemediationAttributes) SubscriptionId() terra.StringValue {
	return terra.ReferenceAsString(spr.ref.Append("subscription_id"))
}

func (spr subscriptionPolicyRemediationAttributes) Timeouts() subscriptionpolicyremediation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[subscriptionpolicyremediation.TimeoutsAttributes](spr.ref.Append("timeouts"))
}

type subscriptionPolicyRemediationState struct {
	FailurePercentage           float64                                      `json:"failure_percentage"`
	Id                          string                                       `json:"id"`
	LocationFilters             []string                                     `json:"location_filters"`
	Name                        string                                       `json:"name"`
	ParallelDeployments         float64                                      `json:"parallel_deployments"`
	PolicyAssignmentId          string                                       `json:"policy_assignment_id"`
	PolicyDefinitionId          string                                       `json:"policy_definition_id"`
	PolicyDefinitionReferenceId string                                       `json:"policy_definition_reference_id"`
	ResourceCount               float64                                      `json:"resource_count"`
	ResourceDiscoveryMode       string                                       `json:"resource_discovery_mode"`
	SubscriptionId              string                                       `json:"subscription_id"`
	Timeouts                    *subscriptionpolicyremediation.TimeoutsState `json:"timeouts"`
}
