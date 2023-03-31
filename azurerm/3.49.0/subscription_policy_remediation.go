// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	subscriptionpolicyremediation "github.com/golingon/terraproviders/azurerm/3.49.0/subscriptionpolicyremediation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewSubscriptionPolicyRemediation(name string, args SubscriptionPolicyRemediationArgs) *SubscriptionPolicyRemediation {
	return &SubscriptionPolicyRemediation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SubscriptionPolicyRemediation)(nil)

type SubscriptionPolicyRemediation struct {
	Name  string
	Args  SubscriptionPolicyRemediationArgs
	state *subscriptionPolicyRemediationState
}

func (spr *SubscriptionPolicyRemediation) Type() string {
	return "azurerm_subscription_policy_remediation"
}

func (spr *SubscriptionPolicyRemediation) LocalName() string {
	return spr.Name
}

func (spr *SubscriptionPolicyRemediation) Configuration() interface{} {
	return spr.Args
}

func (spr *SubscriptionPolicyRemediation) Attributes() subscriptionPolicyRemediationAttributes {
	return subscriptionPolicyRemediationAttributes{ref: terra.ReferenceResource(spr)}
}

func (spr *SubscriptionPolicyRemediation) ImportState(av io.Reader) error {
	spr.state = &subscriptionPolicyRemediationState{}
	if err := json.NewDecoder(av).Decode(spr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", spr.Type(), spr.LocalName(), err)
	}
	return nil
}

func (spr *SubscriptionPolicyRemediation) State() (*subscriptionPolicyRemediationState, bool) {
	return spr.state, spr.state != nil
}

func (spr *SubscriptionPolicyRemediation) StateMust() *subscriptionPolicyRemediationState {
	if spr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", spr.Type(), spr.LocalName()))
	}
	return spr.state
}

func (spr *SubscriptionPolicyRemediation) DependOn() terra.Reference {
	return terra.ReferenceResource(spr)
}

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
	// DependsOn contains resources that SubscriptionPolicyRemediation depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type subscriptionPolicyRemediationAttributes struct {
	ref terra.Reference
}

func (spr subscriptionPolicyRemediationAttributes) FailurePercentage() terra.NumberValue {
	return terra.ReferenceNumber(spr.ref.Append("failure_percentage"))
}

func (spr subscriptionPolicyRemediationAttributes) Id() terra.StringValue {
	return terra.ReferenceString(spr.ref.Append("id"))
}

func (spr subscriptionPolicyRemediationAttributes) LocationFilters() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](spr.ref.Append("location_filters"))
}

func (spr subscriptionPolicyRemediationAttributes) Name() terra.StringValue {
	return terra.ReferenceString(spr.ref.Append("name"))
}

func (spr subscriptionPolicyRemediationAttributes) ParallelDeployments() terra.NumberValue {
	return terra.ReferenceNumber(spr.ref.Append("parallel_deployments"))
}

func (spr subscriptionPolicyRemediationAttributes) PolicyAssignmentId() terra.StringValue {
	return terra.ReferenceString(spr.ref.Append("policy_assignment_id"))
}

func (spr subscriptionPolicyRemediationAttributes) PolicyDefinitionId() terra.StringValue {
	return terra.ReferenceString(spr.ref.Append("policy_definition_id"))
}

func (spr subscriptionPolicyRemediationAttributes) PolicyDefinitionReferenceId() terra.StringValue {
	return terra.ReferenceString(spr.ref.Append("policy_definition_reference_id"))
}

func (spr subscriptionPolicyRemediationAttributes) ResourceCount() terra.NumberValue {
	return terra.ReferenceNumber(spr.ref.Append("resource_count"))
}

func (spr subscriptionPolicyRemediationAttributes) ResourceDiscoveryMode() terra.StringValue {
	return terra.ReferenceString(spr.ref.Append("resource_discovery_mode"))
}

func (spr subscriptionPolicyRemediationAttributes) SubscriptionId() terra.StringValue {
	return terra.ReferenceString(spr.ref.Append("subscription_id"))
}

func (spr subscriptionPolicyRemediationAttributes) Timeouts() subscriptionpolicyremediation.TimeoutsAttributes {
	return terra.ReferenceSingle[subscriptionpolicyremediation.TimeoutsAttributes](spr.ref.Append("timeouts"))
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
